// Test plans for the emulated chip (sim-web): a test block's steps as data a driver executes against a real
// (emulated) board - not the PC fake inside a test build. `pretend` drives a GPIO input, `advance` steps the
// emulated clock, `expect` reads a GPIO output / the UART / the halted flag. Steps a real board can't show
// (an expect on program internals) mark the test emulatable:false honestly.
package lang

import "strings"

// TestStep is one thing a driver does or checks on the emulated board.
type TestStep struct {
	Op     string `json:"op"`               // pretend | advance | expect_pin | expect_printed | expect_notified | expect_stopped | answer
	Pin    int    `json:"pin,omitempty"`    // a GPIO the op drives or reads
	Active bool   `json:"active,omitempty"` // the pin's level to pretend or expect
	Us     int64  `json:"us,omitempty"`     // advance: microseconds to step the clock
	Text   string `json:"text,omitempty"`   // printed/notified text, or the expect's message
	Value  string `json:"value,omitempty"`  // a pretend/answer value the driver evaluates ("500 ms", "0.7 V", 1)
	Msg    string `json:"msg,omitempty"`    // the expect's failure message
}

// TestPlan is one test's steps plus whether they all run on a real board.
type TestPlan struct {
	Name       string     `json:"name"`
	Line       int        `json:"line"`
	Emulatable bool       `json:"emulatable"` // every step is something the emulated board can show
	Steps      []TestStep `json:"steps"`
}

// testPlan serializes one test block's body for the emulated-board driver. d.Body is the block's statements;
// g.devices resolves a pretend/read target's pin.
func (g *gen) testPlan(d *SpecialBlock) TestPlan {
	p := TestPlan{Name: d.Name, Line: d.At.Line, Emulatable: true}
	for _, s := range d.Body.Stmts {
		p.Steps = append(p.Steps, g.planStep(s, &p)...)
	}
	return p
}

// planStep turns one test statement into zero or more plan steps.
func (g *gen) planStep(n Node, p *TestPlan) []TestStep {
	switch s := n.(type) {
	case *ExpectStmt:
		return g.planExpect(s, p)
	case *ExprStmt:
		if c, ok := s.X.(*Call); ok {
			return g.planCall(c, p)
		}
	case *Call:
		return g.planCall(s, p)
	}
	p.Emulatable = false // a test statement the driver can't express
	return nil
}

// planCall handles a bare call in a test: advance(t), dev.pretend(v), answer(v).
func (g *gen) planCall(c *Call, p *TestPlan) []TestStep {
	if id, ok := c.Fn.(*Ident); ok && id.Name == "advance" && len(c.Args) == 1 {
		return []TestStep{{Op: "advance", Value: g.srcText(c.Args[0].Value)}}
	}
	if id, ok := c.Fn.(*Ident); ok && id.Name == "answer" && len(c.Args) == 1 {
		return []TestStep{{Op: "answer", Value: g.srcText(c.Args[0].Value)}}
	}
	if m, ok := c.Fn.(*Member); ok && m.Name == "pretend" && len(c.Args) == 1 {
		if obj, ok := m.X.(*Ident); ok {
			if pin, ok := g.devicePin(obj.Name); ok {
				return []TestStep{{Op: "pretend", Pin: pin, Value: g.srcText(c.Args[0].Value)}}
			}
		}
	}
	p.Emulatable = false
	return nil
}

// planExpect turns `expect cond, "msg"` into a check the emulated board can run.
func (g *gen) planExpect(s *ExpectStmt, p *TestPlan) []TestStep {
	msg := ""
	if m := s.Message(); m != nil {
		msg = g.srcText(m)
	}
	c, ok := s.Cond.(*Call)
	wantActive := true
	if !ok {
		// `expect !led.read()` - a negated pin read is still an expect on a pin, expecting it low.
		if u, isUn := s.Cond.(*Unary); isUn && (u.Op == "!" || u.Op == "not") {
			if cc, isCall := u.X.(*Call); isCall {
				c, wantActive = cc, false
				ok = true
			}
		}
	}
	if !ok {
		p.Emulatable = false
		return nil
	}
	if id, ok := c.Fn.(*Ident); ok {
		switch id.Name {
		case "stopped", "halted":
			return []TestStep{{Op: "expect_stopped", Msg: msg}}
		case "printed":
			return []TestStep{{Op: "expect_printed", Text: g.argText(c), Msg: msg}}
		case "notified":
			return []TestStep{{Op: "expect_notified", Text: g.argText(c), Msg: msg}}
		}
	}
	if m, ok := c.Fn.(*Member); ok && m.Name == "read" {
		if obj, ok := m.X.(*Ident); ok {
			if pin, ok := g.devicePin(obj.Name); ok {
				return []TestStep{{Op: "expect_pin", Pin: pin, Active: wantActive, Msg: msg}}
			}
		}
	}
	// An expect on program internals (watch.running, battery.percent, a variable) has no board output to read.
	p.Emulatable = false
	return nil
}

// devicePin resolves a device name to its GPIO number for pretend/read ops - pin devices only (a pretend/read on an
// object like a Servo or a library class isn't a GPIO op the driver maps).
func (g *gen) devicePin(name string) (int, bool) {
	d, ok := g.devices[name]
	if !ok || d == nil || len(d.Pins) == 0 {
		return 0, false
	}
	if n, ok := g.board.ResolvePin(d.Pins[0].Text); ok {
		return n, true
	}
	return 0, false
}

// argText is the first arg's literal text ("Door opened") for printed/notified.
func (g *gen) argText(c *Call) string {
	if len(c.Args) == 0 {
		return ""
	}
	return g.srcText(c.Args[0].Value)
}

// srcText is an arg's form for a driver to evaluate: a literal's text, else the generated expression.
func (g *gen) srcText(n Node) string {
	switch v := n.(type) {
	case *Literal:
		return strings.Trim(v.Value, `"`)
	case *PinLit:
		return v.Ref.Text
	}
	return g.expr(n)
}
