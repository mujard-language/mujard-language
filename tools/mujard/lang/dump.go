package lang

import (
	"fmt"
	"strings"
)

// Dump prints a syntax tree as indented S-expressions, for tests and "mujard check --tree".
func Dump(n Node) string {
	var b strings.Builder
	dump(&b, n, 0)
	return b.String()
}

func dump(b *strings.Builder, n Node, depth int) {
	pad := strings.Repeat("  ", depth)
	line := func(format string, args ...any) { fmt.Fprintf(b, pad+format+"\n", args...) }
	block := func(bl *Block) {
		if bl == nil {
			line("  (no block)")
			return
		}
		for _, s := range bl.Stmts {
			dump(b, s, depth+1)
		}
	}
	switch n := n.(type) {
	case *File:
		line("(file %s", n.Board)
		for _, it := range n.Items {
			dump(b, it, depth+1)
		}
		line(")")
	case *UseDecl:
		line("(use %s %q version=%q as=%q)", n.From, n.Path, n.Version, n.Alias)
	case *ImportDecl:
		line("(import %s)", n.Name)
	case *NeedsDecl:
		line("(needs %s)", strings.Join(n.Names, " "))
	case *VarDecl:
		line("(%s %s%s = %s)", n.Kind, n.Name, typeSuffix(n.Type), expr(n.Value))
	case *DeviceDecl:
		var pins []string
		for _, p := range n.Pins {
			pins = append(pins, p.Text)
		}
		if n.Kind == "object" {
			line("(device %s %s)", n.Name, expr(n.Object))
		} else {
			line("(device %s %s [%s] %s %v to_end=%v)", n.Name, n.Kind, strings.Join(pins, " "), n.Role, n.Options, n.ToEnd)
		}
	case *FnDecl:
		var ps []string
		for _, p := range n.Params {
			s := p.Name + typeSuffix(p.Type)
			if p.Default != nil {
				s += "=" + expr(p.Default)
			}
			ps = append(ps, s)
		}
		line("(fn %s (%s)%s", n.Name, strings.Join(ps, ", "), typeSuffix(n.Result))
		block(n.Body)
		line(")")
	case *StructDecl:
		var fs []string
		for _, f := range n.Fields {
			fs = append(fs, f.Name+typeSuffix(f.Type))
		}
		line("(struct %s %s)", n.Name, strings.Join(fs, " "))
	case *EnumDecl:
		line("(enum %s %s)", n.Name, strings.Join(n.Values, " "))
	case *EventBlock:
		head := n.Kind
		switch n.Kind {
		case "time":
			head = fmt.Sprintf("at %02d:%02d", n.Hour, n.Min)
		case "start":
			head = "at start"
		default:
			head += " " + expr(n.Expr)
		}
		if n.Fast {
			head += " fast"
		}
		if n.As != "" {
			head += " as " + n.As
		}
		if n.Priority != 0 {
			head += " priority " + fmt.Sprint(n.Priority)
		}
		line("(%s", head)
		block(n.Body)
		line(")")
	case *TaskDecl:
		head := "(task " + n.Name
		if n.Priority != 0 {
			head += " priority " + fmt.Sprint(n.Priority)
		}
		line("%s", head)
		block(n.Body)
		line(")")
	case *SpecialBlock:
		line("(%s %q", n.Kind, n.Name)
		block(n.Body)
		line(")")
	case *LimitDecl:
		line("(limit %s %s)", expr(n.Target), expr(n.Range))
	case *ControlDecl:
		switch {
		case n.Hold:
			line("(control %q %s hold)", n.Label, n.Var)
		case n.Var2 != "":
			line("(control %q %s %s %s)", n.Label, n.Var, n.Var2, expr(n.Range))
		default:
			line("(control %q %s %s)", n.Label, n.Var, expr(n.Range))
		}
	case *CppBlock:
		line("(cpp %q)", n.Code)
	case *IfStmt:
		for i, c := range n.Conds {
			kw := "if"
			if i > 0 {
				kw = "elif"
			}
			line("(%s %s", kw, expr(c))
			block(n.Thens[i])
			line(")")
		}
		if n.Else != nil {
			line("(else")
			block(n.Else)
			line(")")
		}
	case *WhileStmt:
		line("(while %s", expr(n.Cond))
		block(n.Body)
		line(")")
	case *ForStmt:
		line("(for %s in %s", n.Var, expr(n.Iter))
		block(n.Body)
		line(")")
	case *MatchStmt:
		line("(match %s", expr(n.Value))
		for _, c := range n.Cases {
			if c.IsElse {
				line("  (else")
			} else {
				var ps []string
				for _, p := range c.Patterns {
					ps = append(ps, expr(p))
				}
				line("  (case %s", strings.Join(ps, ", "))
			}
			dumpBlock(b, c.Body, depth+2)
			line("  )")
		}
		line(")")
	case *SimpleStmt:
		if n.Kind == "wait" && n.Until {
			line("(wait-until %s)", expr(n.Value))
		} else if n.Value != nil {
			line("(%s %s)", n.Kind, expr(n.Value))
		} else {
			line("(%s)", n.Kind)
		}
	case *ExpectStmt:
		var sb strings.Builder
		for _, a := range n.Args {
			if a.Name != "" {
				fmt.Fprintf(&sb, " %s: %s", a.Name, expr(a.Value))
			} else {
				fmt.Fprintf(&sb, " %s", expr(a.Value))
			}
		}
		line("(expect %s%s)", expr(n.Cond), sb.String())
	case *AssignStmt:
		line("(%s %s %s)", n.Op, expr(n.Target), expr(n.Value))
	case *ExprStmt:
		line("(do %s)", expr(n.X))
	default:
		line("%s", expr(n))
	}
}

func dumpBlock(b *strings.Builder, bl *Block, depth int) {
	if bl == nil {
		return
	}
	for _, s := range bl.Stmts {
		dump(b, s, depth)
	}
}

func typeSuffix(t *TypeExpr) string {
	if t == nil {
		return ""
	}
	return ":" + typeString(t)
}

func typeString(t *TypeExpr) string {
	s := t.Name
	if len(t.Args) > 0 {
		var as []string
		for _, a := range t.Args {
			as = append(as, typeString(a))
		}
		s += "[" + strings.Join(as, ",") + "]"
	}
	if t.Optional {
		s += "?"
	}
	return s
}

// expr prints an expression on one line.
func expr(n Node) string {
	switch n := n.(type) {
	case nil:
		return "_"
	case *Ident:
		return n.Name
	case *NumberLit:
		if n.Unit != "" {
			return n.Text + n.Unit
		}
		return n.Text
	case *StringLit:
		var ps []string
		for _, p := range n.Parts {
			if p.Expr != nil {
				ps = append(ps, "{"+expr(p.Expr)+"}")
			} else {
				ps = append(ps, fmt.Sprintf("%q", p.Text))
			}
		}
		return "(text " + strings.Join(ps, " ") + ")"
	case *Literal:
		return n.Value
	case *ListLit:
		var is []string
		for _, i := range n.Items {
			is = append(is, expr(i))
		}
		return "[" + strings.Join(is, " ") + "]"
	case *PinLit:
		if n.Analog {
			return "(analog " + n.Ref.Text + ")"
		}
		return "(pin " + n.Ref.Text + ")"
	case *Unary:
		return "(" + n.Op + " " + expr(n.X) + ")"
	case *Binary:
		return "(" + n.Op + " " + expr(n.X) + " " + expr(n.Y) + ")"
	case *RangeExpr:
		op := ".."
		if n.Inclusive {
			op = "..="
		}
		return "(" + op + " " + expr(n.Lo) + " " + expr(n.Hi) + ")"
	case *Conversion:
		return "(as " + expr(n.X) + " " + n.Unit + ")"
	case *CondExpr:
		return "(ifelse " + expr(n.Cond) + " " + expr(n.Then) + " " + expr(n.Else) + ")"
	case *Call:
		s := "(call " + expr(n.Fn)
		for _, a := range n.Args {
			if a.Name != "" {
				s += " " + a.Name + ":" + expr(a.Value)
			} else {
				s += " " + expr(a.Value)
			}
		}
		return s + ")"
	case *Member:
		return expr(n.X) + "." + n.Name
	case *Index:
		return expr(n.X) + "[" + expr(n.Ix) + "]"
	}
	return fmt.Sprintf("<%T>", n)
}
