package lang

import "strings"

// The rules for 'on … fast' handlers (spec 6.5, C314). A fast handler runs as a hardware interrupt, in the middle of
// whatever else is running, so it may only touch variables, pins and a few pure built-ins: no printing, no waiting,
// no libraries, no text. The rules are checked after the whole program, when every expression's type is known, and
// through every function the handler calls.

// fastBuiltins are the built-in functions a fast handler may call: pure, quick, and safe inside an interrupt.
var fastBuiltins = map[string]bool{"abs": true, "min": true, "max": true, "clamp": true, "round": true,
	"between": true, "now": true, "since": true}

// fastPinMethods are the pin and pin-group methods a fast handler may call.
var fastPinMethods = map[string]bool{"on": true, "off": true, "toggle": true, "write": true, "read": true}

// fastIntMethods are the bit methods of an int.
var fastIntMethods = map[string]bool{"bit": true, "set_bit": true, "clear_bit": true, "toggle_bit": true}

// fastRules reports C314 for everything in a fast handler's body that may not run inside an interrupt.
func (c *checker) fastRules(b *Block) {
	f := &fastCheck{c: c, fns: map[*FnDecl]string{}}
	f.report = func(at Pos, what string) { c.errorf("C314", at, "what", what) }
	f.block(b, map[string]bool{})
}

type fastCheck struct {
	c      *checker
	fns    map[*FnDecl]string // the first thing each function does that a fast handler may not ("" = nothing)
	report func(at Pos, what string)
}

func (f *fastCheck) block(b *Block, local map[string]bool) {
	if b == nil {
		return
	}
	for _, s := range b.Stmts {
		f.stmt(s, local)
	}
}

func (f *fastCheck) stmt(n Node, local map[string]bool) {
	switch s := n.(type) {
	case *VarDecl:
		local[s.Name] = true
		f.expr(s.Value)
	case *AssignStmt:
		if t := f.c.types[s.Target]; t != nil && t.K == TText {
			f.report(leftmost(s.Target), "text")
			return
		}
		switch t := s.Target.(type) {
		case *Ident:
			if sym := f.c.global.syms[t.Name]; !local[t.Name] && sym != nil && sym.decl != nil && sym.decl.Kind == "saved" {
				f.report(t.At, "changing the saved var '"+t.Name+"'")
			}
		case *Index:
			f.report(leftmost(t), "changing a list")
		}
		f.expr(s.Target)
		f.expr(s.Value)
	case *ExprStmt:
		f.expr(s.X)
	case *SimpleStmt:
		switch s.Kind {
		case "print", "wait":
			f.report(s.At, "'"+s.Kind+"'")
		default:
			f.expr(s.Value)
		}
	case *ExpectStmt:
		f.report(s.At, "'expect'")
	case *IfStmt:
		for _, x := range s.Conds {
			f.expr(x)
		}
		for _, t := range s.Thens {
			f.block(t, local)
		}
		f.block(s.Else, local)
	case *MatchStmt:
		f.expr(s.Value)
		for _, mc := range s.Cases {
			f.block(mc.Body, local)
		}
	case *ForStmt:
		if _, ok := s.Iter.(*RangeExpr); !ok {
			f.report(s.At, "a for loop over anything but a range")
		}
		f.expr(s.Iter)
		local[s.Var] = true
		f.block(s.Body, local)
	case *WhileStmt:
		f.report(s.At, "'while'")
	case *EventBlock:
		f.report(s.At, "'"+s.Kind+"'")
	}
	// cpp: blocks are left to the programmer, as everywhere (W604)
}

func (f *fastCheck) expr(n Node) {
	if n == nil {
		return
	}
	if t := f.c.types[n]; t != nil && t.K == TText {
		f.report(leftmost(n), "text")
		return
	}
	switch e := n.(type) {
	case *StringLit:
		f.report(e.At, "text")
	case *ListLit:
		f.report(e.At, "a list")
	case *Unary:
		f.expr(e.X)
	case *Binary:
		f.expr(e.X)
		f.expr(e.Y)
	case *RangeExpr:
		f.expr(e.Lo)
		f.expr(e.Hi)
	case *Conversion:
		f.expr(e.X)
	case *CondExpr:
		f.expr(e.Then)
		f.expr(e.Cond)
		f.expr(e.Else)
	case *Index:
		f.expr(e.X)
		f.expr(e.Ix)
	case *Member:
		// a module's variable, an enum value, a struct field and .number are fine; a device's or an object's
		// properties (analog.value, board.uptime, a stopwatch) are not
		if xt := f.c.types[e.X]; xt != nil && (xt.K == TDevice || (xt.K == TObject && !strings.HasPrefix(xt.Name, "mod:") && !strings.HasPrefix(xt.Name, "enum:"))) {
			f.report(leftmost(e), "'"+nodeName(e)+"'")
			return
		}
		f.expr(e.X)
	case *Call:
		f.call(e)
	}
}

func (f *fastCheck) call(e *Call) {
	at := leftmost(e.Fn)
	switch fn := e.Fn.(type) {
	case *Ident:
		if sym := f.c.global.syms[fn.Name]; sym != nil && sym.fn != nil {
			f.args(e)
			if what := f.fnUse(sym.fn); what != "" {
				f.report(at, "'"+fn.Name+"' (it uses "+what+")")
			}
			return
		}
		if fastBuiltins[fn.Name] && f.c.global.syms[fn.Name] == nil {
			f.args(e)
			return
		}
	case *Member:
		xt := f.c.types[fn.X]
		switch {
		case xt == nil:
		case xt.K == TDevice && xt.Name != "analog" && fastPinMethods[fn.Name]:
			f.expr(fn.X) // leds[i].on(): the index
			f.args(e)
			return
		case xt.K == TInt && fastIntMethods[fn.Name]:
			f.expr(fn.X)
			f.args(e)
			return
		case xt.K == TObject && xt.Name == "board" && fn.Name == "halt":
			// the reason must be plain text written out: no text is built inside the interrupt
			if len(e.Args) == 1 {
				if s, ok := e.Args[0].Value.(*StringLit); ok && !hasInterpolation(s) {
					return
				}
			}
			f.args(e)
			return
		}
	}
	f.report(at, "'"+nodeName(e.Fn)+"'")
}

func (f *fastCheck) args(e *Call) {
	for _, a := range e.Args {
		f.expr(a.Value)
	}
}

// fnUse gives the first thing a function does that a fast handler may not ("" if nothing), checking it once.
func (f *fastCheck) fnUse(fd *FnDecl) string {
	if what, ok := f.fns[fd]; ok {
		return what
	}
	f.fns[fd] = "" // a recursive call is judged by the rest of the body
	what := ""
	sub := &fastCheck{c: f.c, fns: f.fns, report: func(_ Pos, w string) {
		if what == "" {
			what = w
		}
	}}
	local := map[string]bool{}
	for _, p := range fd.Params {
		local[p.Name] = true
	}
	sub.block(fd.Body, local)
	f.fns[fd] = what
	return what
}

// leftmost is where an expression starts in the source: the 'log' of log.info, not its '.'.
func leftmost(n Node) Pos {
	switch e := n.(type) {
	case *Member:
		return leftmost(e.X)
	case *Call:
		return leftmost(e.Fn)
	case *Index:
		return leftmost(e.X)
	}
	return n.Start()
}

func hasInterpolation(s *StringLit) bool {
	for _, p := range s.Parts {
		if p.Expr != nil {
			return true
		}
	}
	return false
}
