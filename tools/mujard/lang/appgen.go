package lang

// The app built-ins (Phase 9): show, plot, notify, control and ask, generated as calls into the runtime's app
// protocol (docs/lang/MUJARD_APP_PROTOCOL.md, mujard_rt app_*). Numbers go to the app in one display unit per
// dimension, so tiles, gauges, sliders and plots agree.

import (
	"fmt"
	"hash/fnv"
	"strings"
)

// displayUnits are the units the app shows each dimension in, with how many base units one of them is.
var displayUnits = map[string]struct {
	unit string
	per  float64
}{
	"ratio": {"%", 0.01}, "temperature": {"C", 1}, "voltage": {"V", 1}, "current": {"mA", 1e-3},
	"length": {"cm", 0.01}, "angle": {"deg", 1}, "speed": {"m/s", 1}, "memory": {"bytes", 1}, "baud": {"baud", 1},
	"time": {"s", 1e6}, "frequency": {"Hz", 1e3},
}

// appNumber is a value as a float in its display unit, with that unit; ok is false for values without a number
// (text, enums, lists).
func (g *gen) appNumber(n Node) (expr, unit string, ok bool) {
	t := nonOpt(g.typeOf(n))
	x := g.expr(n)
	switch t.K {
	case TInt, TFloat:
		return "(float)(" + x + ")", "", true
	case TBool:
		return "((" + x + ") ? 1.0f : 0.0f)", "", true
	case TQty:
		d, known := displayUnits[t.Dim]
		if !known {
			return "(float)(" + x + ")", "", true
		}
		switch t.Dim {
		case "time":
			return "mujard::secs(" + x + ")", d.unit, true
		case "frequency":
			return "((float)(" + x + ") / 1000.0f)", d.unit, true
		}
		if d.per == 1 {
			return "(float)(" + x + ")", d.unit, true
		}
		return fmt.Sprintf("((float)(%s) / %s)", x, cfloat(d.per)), d.unit, true
	}
	return "", "", false
}

// appCall generates show / plot / notify / ask; ok is false for anything else.
func (g *gen) appCall(c *Call, name string) (string, bool) {
	switch name {
	case "notify":
		g.usesApp = true
		g.usesNotify = true
		a := argMap(c, "message", "urgent")
		urgent := "false"
		if a["urgent"] != nil {
			urgent = "(bool)(" + g.expr(a["urgent"]) + ")"
		}
		return fmt.Sprintf("mujard::app_notify(%s, %s, %d)", g.textValue(a["message"]), urgent, c.At.Line), true
	case "plot":
		g.usesApp = true
		var b strings.Builder
		b.WriteString("mujard::AppPlot()")
		for _, arg := range c.Args {
			num, _, ok := g.appNumber(arg.Value)
			if !ok {
				return g.later(arg.Value.Start(), "plot of a value that isn't a number", phaseLater), true
			}
			name := arg.Name
			if name == "" {
				name = sourceText(arg.Value)
			}
			fmt.Fprintf(&b, ".add(%s, %s)", cString(name), num)
		}
		b.WriteString(".end()")
		return b.String(), true
	case "show":
		g.usesApp = true
		return g.showCall(c), true
	case "ask":
		if id, hoisted := g.askIDs[c]; hoisted {
			return fmt.Sprintf("mujard::app_ask_result(%d)", id), true
		}
		return g.later(c.At, "ask here (it works as 'if ask …:', 'var x = ask …' or 'x = ask …' inside a task)", phaseLater), true
	}
	return "", false
}

// showCall is show "Label", value [, as: number | light | gauge(lo..hi)].
func (g *gen) showCall(c *Call) string {
	a := argMap(c, "label", "value", "as")
	v := a["value"]
	if v == nil {
		return g.later(c.At, "show without a value", phaseLater)
	}
	t := nonOpt(g.typeOf(v))
	text := g.textValue(v)
	if t.K == TBool {
		text = fmt.Sprintf("mujard::Text((%s) ? \"on\" : \"off\")", g.expr(v))
	} else if t.K != TText {
		text = "mujard::Text()" + g.addPart(v)
	}
	num, unit, hasNum := g.appNumber(v)
	style, hasRange, lo, hi := 0, false, "0", "0"
	if t.K == TBool {
		style = 2
	}
	if as := a["as"]; as != nil {
		switch x := as.(type) {
		case *Ident:
			switch x.Name {
			case "light":
				style = 2
			case "number":
				style = 0
			default:
				return g.later(as.Start(), "show as "+x.Name, phaseLater)
			}
		case *Call:
			id, isIdent := x.Fn.(*Ident)
			lo2, hi2, _, isRange := Node(nil), Node(nil), false, false
			if len(x.Args) == 1 {
				lo2, hi2, _, isRange = rangeEnds(x.Args[0].Value)
			}
			if !isIdent || id.Name != "gauge" || !isRange || !hasNum {
				return g.later(as.Start(), "this kind of show style", phaseLater)
			}
			style, hasRange = 1, true
			lo, _, _ = g.appNumber(lo2)
			hi, _, _ = g.appNumber(hi2)
		default:
			return g.later(as.Start(), "this kind of show style", phaseLater)
		}
	}
	if !hasNum {
		num = "0.0f"
	}
	return fmt.Sprintf("mujard::app_show(%s, %s, %t, %s, %s, %d, %t, %s, %s)",
		g.cstr(a["label"]), text, hasNum, num, cString(unit), style, hasRange, lo, hi)
}

// controlDecl registers control "Label", var[, lo..hi] with the runtime (in setup, before app_begin). A hold button
// and a joystick's two axes are marked held right after (app_control_hold): the board puts them back at rest when
// the app stops refreshing them - the failsafe for anything that drives.
func (g *gen) controlDecl(d *ControlDecl) {
	g.usesApp = true
	switch {
	case d.Hold:
		if g.oneControl(d, d.Var) {
			g.setup.add(d.At.Line, "    mujard::app_control_hold(mujard::APP_HOLD_BUTTON);")
			g.controlLabels = append(g.controlLabels, "hold")
		}
	case d.Var2 != "":
		if g.oneControl(d, d.Var) {
			g.setup.add(d.At.Line, "    mujard::app_control_hold(mujard::APP_HOLD_STICK_X);")
			if g.oneControl(d, d.Var2) {
				g.setup.add(d.At.Line, "    mujard::app_control_hold(mujard::APP_HOLD_STICK_Y);")
				g.controlLabels = append(g.controlLabels, "stick")
			}
		}
	default:
		g.oneControl(d, d.Var)
	}
}

// oneControl emits the app_control call for one variable of a control line; false when it couldn't.
func (g *gen) oneControl(d *ControlDecl, varName string) bool {
	var decl *VarDecl
	for _, it := range g.f.Items {
		if v, ok := it.(*VarDecl); ok && v.Name == varName && v.Kind != "const" {
			decl = v
		}
	}
	if decl == nil {
		g.later(d.At, "control of this variable", phaseLater)
		return false
	}
	t := g.typeOf(decl.Value)
	if decl.Type != nil {
		if dt := g.declTypeFromTypeExpr(decl.Type); dt != nil {
			t = dt
		}
	}
	if t.Opt {
		g.later(d.At, "a control on a variable that can be none", phaseLater)
		return false
	}
	name := "&" + cname(varName)
	label := cString(d.Label)
	if t.K == TBool {
		g.setup.add(d.At.Line, "    mujard::app_control(%s, true, 0, 1, 1, \"\", 1, %s, mujard::APP_BOOL);", label, name)
		g.controlLabels = append(g.controlLabels, d.Label)
		return true
	}
	lo, hi, _, ok := rangeEnds(d.Range)
	if !ok {
		g.later(d.At, "a control on a number without a range (write: control \"Label\", x, 0..100)", phaseLater)
		return false
	}
	unit, per := "", 1.0
	vtype := "mujard::APP_FLOAT"
	switch {
	case t.K == TInt:
		vtype = "mujard::APP_INT"
	case t.K == TQty:
		if du, known := displayUnits[t.Dim]; known {
			unit, per = du.unit, du.per
		}
		if n, isNum := lo.(*NumberLit); isNum && n.Unit != "" {
			unit, per = n.Unit, unitFactor(n.Unit) // the unit the range is written in: 0 %..100 % shows as %
		}
		if t.Dim == "time" || t.Dim == "frequency" {
			vtype = "mujard::APP_I64"
		}
	case t.K != TFloat:
		g.later(d.At, "a control on this kind of variable", phaseLater)
		return false
	}
	minX := fmt.Sprintf("(float)(%s) / %s", g.expr(lo), cfloat(per))
	maxX := fmt.Sprintf("(float)(%s) / %s", g.expr(hi), cfloat(per))
	step := "1"
	if t.K != TInt {
		step = fmt.Sprintf("((float)(%s) - (float)(%s)) / %s / 100.0f", g.expr(hi), g.expr(lo), cfloat(per))
	}
	g.setup.add(d.At.Line, "    mujard::app_control(%s, false, %s, %s, %s, %s, %s, %s, %s);", label, minX, maxX, step, cString(unit), cfloat(per), name, vtype)
	g.controlLabels = append(g.controlLabels, d.Label)
	return true
}

// appHash names this program's dashboard: a different program, or changed controls, starts the app's dashboard fresh.
func (g *gen) appHash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(g.src))
	for _, l := range g.controlLabels {
		h.Write([]byte{0})
		h.Write([]byte(l))
	}
	return h.Sum32()
}

// ---- waits inside a task: a call that can suspend (ask, a line read, an http request, a fn that waits) runs
// before the statement that holds it: its begin/send is emitted, the task waits at that point (like 'wait until'),
// and the expression then reads the result.

// hoistWaits emits, before a statement in a task, the start-and-wait for each call it evaluates unconditionally.
func (g *gen) hoistWaits(out *section, n Node, depth int) {
	if g.inTask == nil {
		return
	}
	var roots []Node
	switch s := n.(type) {
	case *IfStmt:
		roots = s.Conds[:1] // an 'elif' condition only runs when the ones above were false
	case *VarDecl:
		roots = []Node{s.Value}
	case *AssignStmt:
		roots = []Node{s.Value}
	case *ExprStmt:
		roots = []Node{s.X}
	case *SimpleStmt:
		if s.Kind == "print" {
			roots = []Node{s.Value}
		}
	}
	p := ind(depth)
	for _, r := range roots {
		for _, c := range g.waitsIn(r) {
			g.inTask.waits++
			w := g.inTask.waits
			if g.waitResults == nil {
				g.waitResults = map[*Call]string{}
			}
			g.waitFor(out, c, w, p)
		}
	}
}

// waitFor emits the start of one suspending call and the wait for its end, and records where its result is read.
func (g *gen) waitFor(out *section, c *Call, w int, p string) {
	switch fn := c.Fn.(type) {
	case *Ident:
		if d := g.fnDecls[fn.Name]; d != nil && g.waitingFn(d) {
			if g.cycFn[d] {
				g.later(c.At, "a function that waits and calls itself, even through another function", phaseLater)
				return
			}
			obj := g.fnInsts[c]
			if obj == "" {
				g.later(c.At, "a function that waits here (use it as 'var x = …' or 'x = …' inside a task)", phaseLater)
				return
			}
			m := argMap(c, paramNames(d)...)
			for _, prm := range d.Params {
				pt := g.declTypeFromTypeExpr(prm.Type)
				if v, ok := m[prm.Name]; ok {
					g.setWant(v, pt)
					out.add(c.At.Line, "%s%s.%s = %s;", p, obj, cname(prm.Name), g.expr(v))
				} else if prm.Default != nil {
					g.setWant(prm.Default, pt)
					out.add(c.At.Line, "%s%s.%s = %s;", p, obj, cname(prm.Name), g.expr(prm.Default))
				}
			}
			out.add(c.At.Line, "%s%s.restart();", p, obj)
			out.add(c.At.Line, "%sMUJARD_WAIT_UNTIL((%s.run(), !%s.running()), %d);", p, obj, obj, w)
			if d.Result != nil {
				g.waitResults[c] = obj + ".result_"
			} else {
				g.waitResults[c] = "(void)0"
			}
			return
		}
		// ask
		if g.askIDs == nil {
			g.askIDs = map[*Call]int{}
		}
		g.usesApp = true
		g.askN++
		id := g.askN
		g.askIDs[c] = id
		a := argMap(c, "question", "timeout")
		timeout := "0"
		if a["timeout"] != nil {
			timeout = g.expr(a["timeout"])
		}
		out.add(c.At.Line, "%smujard::app_ask_begin(%d, %s, %s);", p, id, g.textValue(a["question"]), timeout)
		out.add(c.At.Line, "%sMUJARD_WAIT_UNTIL(mujard::app_ask_done(%d), %d);", p, id, w)
	case *Member:
		xt := nonOpt(g.typeOf(fn.X))
		obj := g.expr(fn.X)
		switch {
		case xt.K == TObject && xt.Name == "serial" && fn.Name == "read_line":
			g.serialLines = true
			out.add(c.At.Line, "%sMUJARD_WAIT_UNTIL(mujard::serial_has_line(), %d);", p, w)
			g.waitResults[c] = "mujard::serial_take_line()"
		case xt.K == TObject && xt.Name == "uart" && fn.Name == "read_line":
			g.uartLines = append(g.uartLines, obj)
			out.add(c.At.Line, "%sMUJARD_WAIT_UNTIL(%s.has_line(), %d);", p, obj, w)
			g.waitResults[c] = obj + ".take_line()"
		case xt.K == TObject && xt.Name == "http" && (fn.Name == "get" || fn.Name == "post"):
			a := argMap(c, "url", "data")
			post := "false"
			body := "mujard::Text()"
			if fn.Name == "post" {
				post = "true"
				if a["data"] != nil {
					body = g.text(a["data"])
				}
			}
			g.httpN++
			id := g.httpN // each call site has its own number: the runtime keeps the one-request rule honest
			out.add(c.At.Line, "%smujard::http_begin(%d, %s, %s, %s);", p, id, post, g.textValue(a["url"]), body)
			out.add(c.At.Line, "%sMUJARD_WAIT_UNTIL(mujard::http_done(%d), %d);", p, id, w)
			g.waitResults[c] = fmt.Sprintf("mujard::http_result(%d)", id)
		}
	}
}

func paramNames(d *FnDecl) []string {
	var names []string
	for _, p := range d.Params {
		names = append(names, p.Name)
	}
	return names
}

// waitsIn finds the suspending calls an expression always evaluates, in evaluation order: not in the right side of
// and/or or in either branch of 'a if c else b', where starting one would start it even when the program wouldn't.
func (g *gen) waitsIn(n Node) []*Call {
	var out []*Call
	var walk func(Node)
	walk = func(n Node) {
		switch e := n.(type) {
		case *Call:
			walk(e.Fn)
			for _, a := range e.Args {
				walk(a.Value)
			}
			if g.needsWait(e) {
				out = append(out, e)
			}
		case *Binary:
			walk(e.X)
			if e.Op != "and" && e.Op != "or" {
				walk(e.Y)
			}
		case *Unary:
			walk(e.X)
		case *Member:
			walk(e.X)
		case *Index:
			walk(e.X)
			walk(e.Ix)
		case *Conversion:
			walk(e.X)
		case *CondExpr:
			walk(e.Cond)
		case *StringLit:
			for _, p := range e.Parts {
				if p.Expr != nil {
					walk(p.Expr)
				}
			}
		}
	}
	walk(n)
	return out
}
