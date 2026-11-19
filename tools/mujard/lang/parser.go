package lang

import "strings"

// ParseSource lexes and parses a .muj program. Lexer diagnostics come first, then the parser's.
func ParseSource(src string) (*File, []Diagnostic) {
	toks, lexErrs := Lex(src)
	f, parseErrs := Parse(toks)
	return f, append(lexErrs, parseErrs...)
}

// ParseModuleSource lexes and parses a module: a .muj file loaded by 'import name'. A module doesn't need
// a 'board' line (the file may still have one).
func ParseModuleSource(src string) (*File, []Diagnostic) {
	toks, lexErrs := Lex(src)
	f, parseErrs := parseModule(toks)
	return f, append(lexErrs, parseErrs...)
}

// Parse builds the syntax tree from tokens (spec section 3). Like the lexer, it keeps going after an error - it skips to
// the next line or block - so one run reports every syntax error (up to maxDiagnostics).
func Parse(toks []Token) (*File, []Diagnostic) {
	return parse(toks, false)
}

func parseModule(toks []Token) (*File, []Diagnostic) {
	return parse(toks, true)
}

func parse(toks []Token, module bool) (*File, []Diagnostic) {
	if len(toks) == 0 || toks[len(toks)-1].Kind != EOF {
		toks = append(toks, Token{Kind: EOF})
	}
	p := &parser{toks: toks, module: module}
	f := p.file()
	return f, p.errs
}

type parser struct {
	toks   []Token
	i      int
	errs   []Diagnostic
	module bool // parsing a module file: 'board' is optional
}

var pinRoles = map[string]bool{"input": true, "output": true, "pwm": true}
var pinOptions = map[string]bool{"pullup": true, "pulldown": true, "inverted": true}
var compareOps = map[string]bool{"==": true, "!=": true, "<": true, "<=": true, ">": true, ">=": true}
var assignOps = map[string]bool{"=": true, "+=": true, "-=": true, "*=": true, "/=": true}

// ---- token helpers

func (p *parser) peek() Token { return p.toks[p.i] }

func (p *parser) peekAt(n int) Token {
	if p.i+n < len(p.toks) {
		return p.toks[p.i+n]
	}
	return p.toks[len(p.toks)-1]
}

func (p *parser) next() Token {
	t := p.toks[p.i]
	if t.Kind != EOF {
		p.i++
	}
	return t
}

func (p *parser) isName(text string) bool { t := p.peek(); return t.Kind == NAME && t.Text == text }

func (p *parser) errorf(code string, pos Pos, args ...string) { p.add(diag(code, pos, args...)) }

func (p *parser) add(d Diagnostic) {
	for _, e := range p.errs {
		if e.Pos == d.Pos {
			return // one message per place: the first is the useful one
		}
	}
	if len(p.errs) < maxDiagnostics {
		p.errs = append(p.errs, d)
	}
}

// tokText is how a token is named in "Didn't expect '...' here".
func tokText(t Token) string {
	switch t.Kind {
	case NEWLINE:
		return "end of line"
	case INDENT:
		return "indentation"
	case DEDENT:
		return "end of block"
	case EOF:
		return "end of file"
	}
	return t.Text
}

// skipBlock skips an indented block (the current token is its INDENT).
func (p *parser) skipBlock() {
	depth := 0
	for {
		t := p.next()
		switch t.Kind {
		case INDENT:
			depth++
		case DEDENT:
			depth--
			if depth <= 0 {
				return
			}
		case EOF:
			return
		}
	}
}

// syncLine skips the rest of the line, and the block under it if there is one.
func (p *parser) syncLine() {
	for {
		t := p.peek()
		switch t.Kind {
		case EOF:
			return
		case NEWLINE:
			p.next()
			if p.peek().Kind == INDENT {
				p.skipBlock()
			}
			return
		case INDENT:
			p.skipBlock()
			return
		case DEDENT:
			return
		}
		p.next()
	}
}

// endLine expects the end of a line after a complete statement.
func (p *parser) endLine() {
	t := p.peek()
	switch t.Kind {
	case NEWLINE:
		p.next()
		if p.peek().Kind == INDENT {
			p.errorf("P203", p.peek().Pos)
			p.skipBlock()
		}
	case EOF, DEDENT:
	default:
		p.errorf("P205", t.Pos, "tok", tokText(t))
		p.syncLine()
	}
}

func (p *parser) expectOp(op string) bool {
	if p.peek().Is(op) {
		p.next()
		return true
	}
	p.errorf("P205", p.peek().Pos, "tok", tokText(p.peek()))
	return false
}

// name reads a name; a keyword in its place is P211.
func (p *parser) name() (string, Pos) {
	t := p.peek()
	switch t.Kind {
	case NAME:
		p.next()
		return t.Text, t.Pos
	case KEYWORD:
		p.errorf("P211", t.Pos, "kw", t.Text)
		p.next()
		return t.Text, t.Pos
	}
	p.errorf("P204", t.Pos)
	return "", t.Pos
}

// ---- the file

func (p *parser) file() *File {
	f := &File{}
	for p.peek().Kind == NEWLINE {
		p.next()
	}
	t := p.peek()
	if t.Is("board") {
		p.next()
		f.BoardPos = t.Pos
		switch b := p.peek(); b.Kind {
		case NAME:
			f.Board = b.Text
			p.next()
		case STRING:
			f.Board = stringText(b)
			p.next()
		default:
			p.errorf("P204", b.Pos)
		}
		p.endLine()
		// 'safety strict' right after the board line (spec 6.10). Only exactly 'safety <word>' on its own line: a device
		// or variable called safety ('safety = pin 3 ...') is still an ordinary declaration.
		for p.peek().Kind == NEWLINE {
			p.next()
		}
		if s := p.peek(); p.isName("safety") && p.peekAt(1).Kind == NAME && (p.peekAt(2).Kind == NEWLINE || p.peekAt(2).Kind == EOF) {
			p.next()
			f.Safety, f.SafetyPos = p.peek().Text, s.Pos
			p.next()
			p.endLine()
		}
	} else if !p.module {
		p.errorf("P206", t.Pos)
	}
	for p.peek().Kind != EOF {
		start := p.i
		if item := p.topItem(); item != nil {
			f.Items = append(f.Items, item)
		}
		if p.i == start {
			p.next() // never stall
		}
	}
	return f
}

func stringText(t Token) string {
	var b strings.Builder
	for _, part := range t.Parts {
		b.WriteString(part.Text)
	}
	return b.String()
}

func (p *parser) topItem() Node {
	t := p.peek()
	switch t.Kind {
	case NEWLINE, DEDENT:
		p.next()
		return nil
	case INDENT:
		p.errorf("P203", t.Pos)
		p.skipBlock()
		return nil
	case KEYWORD:
		switch t.Text {
		case "use":
			return p.useDecl()
		case "import":
			p.next()
			n, _ := p.name()
			d := &ImportDecl{At: t.Pos, Name: n}
			if p.peek().Is("(") { // import sonar(trig: pin 7, echo: pin 8)
				p.next()
				for !p.peek().Is(")") {
					pn, at := p.name()
					if pn == "" || !p.expectOp(":") {
						p.syncLine()
						return d
					}
					if !p.peek().Is("pin") {
						p.errorf("P208", p.peek().Pos)
						p.syncLine()
						return d
					}
					p.next()
					d.Pins = append(d.Pins, ImportPin{At: at, Name: pn, Ref: p.pinRef()})
					if !p.peek().Is(",") {
						break
					}
					p.next()
				}
				if !p.expectOp(")") {
					p.syncLine()
					return d
				}
			}
			p.endLine()
			return d
		case "const", "var":
			return p.varDecl("")
		case "saved":
			p.next()
			if !p.peek().Is("var") {
				p.errorf("P205", p.peek().Pos, "tok", tokText(p.peek()))
				p.syncLine()
				return nil
			}
			return p.varDecl("saved")
		case "fn":
			return p.fnDecl()
		case "struct":
			return p.structDecl()
		case "enum":
			return p.enumDecl()
		case "every", "after", "when", "on", "at":
			return p.eventBlock(true)
		case "task":
			p.next()
			n, _ := p.name()
			return &TaskDecl{At: t.Pos, Name: n, Priority: p.priority(), Core: p.core(), Body: p.block()}
		case "cpp":
			return p.cppBlock()
		case "elif", "else":
			p.errorf("P207", t.Pos, "kw", t.Text)
			p.syncLine()
			return nil
		}
	case NAME:
		next := p.peekAt(1)
		switch {
		case t.Text == "safe_state" && next.Is(":"):
			p.next()
			return &SpecialBlock{At: t.Pos, Kind: "safe_state", Body: p.block()}
		case t.Text == "test" && next.Kind == STRING:
			p.next()
			p.next()
			return &SpecialBlock{At: t.Pos, Kind: "test", Name: stringText(next), Body: p.block()}
		case t.Text == "control" && next.Kind == STRING:
			return p.controlDecl()
		case t.Text == "limit" && !next.Is("=") && next.Kind != NEWLINE:
			return p.limitDecl()
		case t.Text == "needs" && (next.Kind == NAME || next.Is("pin")): // a device called needs ('needs = ...') is still a device
			return p.needsDecl()
		case next.Is("="):
			return p.deviceDecl()
		}
	}
	p.errorf("P216", t.Pos)
	p.syncLine()
	return nil
}

// ---- top-level declarations

// needsDecl: needs i2c, pwm, pin trig (spec 10.5).
func (p *parser) needsDecl() Node {
	d := &NeedsDecl{At: p.next().Pos}
	for {
		isPin := p.peek().Is("pin")
		if isPin {
			p.next()
		}
		n, at := p.name()
		if n == "" {
			p.syncLine()
			return d
		}
		if isPin {
			d.PinNames, d.PinPos = append(d.PinNames, n), append(d.PinPos, at)
		} else {
			d.Names, d.Pos = append(d.Names, n), append(d.Pos, at)
		}
		if !p.peek().Is(",") {
			break
		}
		p.next()
	}
	p.endLine()
	return d
}

func (p *parser) useDecl() Node {
	at := p.next().Pos
	u := &UseDecl{At: at}
	t := p.peek()
	switch {
	case t.Kind == STRING:
		u.From, u.Path = "arduino", stringText(t)
		p.next()
	case t.Kind == NAME && (t.Text == "idf" || t.Text == "pio") && p.peekAt(1).Kind == STRING:
		u.From, u.Path = t.Text, stringText(p.peekAt(1))
		p.next()
		p.next()
	case t.Kind == NAME:
		parts := []string{t.Text}
		p.next()
		for p.peek().Is(".") && p.peekAt(1).Kind == NAME {
			p.next()
			parts = append(parts, p.next().Text)
		}
		u.Path = strings.Join(parts, ".")
		u.From = "module"
		if parts[0] == "mujard" {
			u.From = "mujard"
		}
	default:
		p.errorf("P204", t.Pos)
		p.syncLine()
		return u
	}
	if p.isName("version") && p.peekAt(1).Kind == STRING {
		p.next()
		u.Version = stringText(p.next())
	}
	if p.peek().Is("as") {
		p.next()
		u.Alias, _ = p.name()
	}
	p.endLine()
	return u
}

func (p *parser) varDecl(kind string) Node {
	t := p.next() // const / var
	if kind == "" {
		kind = t.Text
	}
	d := &VarDecl{At: t.Pos, Kind: kind}
	d.Name, _ = p.name()
	if p.peek().Is(":") {
		p.next()
		d.Type = p.typeExpr()
	}
	if !p.peek().Is("=") {
		p.errorf("P205", p.peek().Pos, "tok", tokText(p.peek()))
		p.syncLine()
		return d
	}
	p.next()
	d.Value = p.expr()
	p.endLine()
	return d
}

func (p *parser) typeExpr() *TypeExpr {
	t := p.peek()
	te := &TypeExpr{At: t.Pos}
	te.Name, _ = p.name()
	if p.peek().Is("[") {
		p.next()
		for {
			if a := p.peek(); a.Kind == INT {
				p.next()
				te.Args = append(te.Args, &TypeExpr{At: a.Pos, Name: a.Text})
			} else {
				te.Args = append(te.Args, p.typeExpr())
			}
			if !p.peek().Is(",") {
				break
			}
			p.next()
		}
		p.expectOp("]")
	}
	if p.peek().Is("?") {
		p.next()
		te.Optional = true
	}
	return te
}

func (p *parser) deviceDecl() Node {
	nameTok := p.next()
	p.next() // '='
	d := &DeviceDecl{At: nameTok.Pos, Name: nameTok.Text}
	t := p.peek()
	switch {
	case t.Is("pin"):
		p.next()
		d.Kind = "pin"
		d.Pins = []PinRef{p.pinRef()}
		p.pinRole(d)
	case t.Kind == NAME && t.Text == "pins":
		p.next()
		d.Kind = "pins"
		first := p.pinRef()
		if p.peek().Is("..") || p.peek().Is("..=") {
			d.ToEnd = p.next().Text == "..="
		} else {
			p.errorf("P208", p.peek().Pos)
		}
		d.Pins = []PinRef{first, p.pinRef()}
		p.pinRole(d)
	case t.Is("[") && p.peekAt(1).Is("pin"):
		p.next()
		d.Kind = "pinlist"
		for {
			if !p.peek().Is("pin") {
				p.errorf("P208", p.peek().Pos)
				break
			}
			p.next()
			d.Pins = append(d.Pins, p.pinRef())
			if !p.peek().Is(",") {
				break
			}
			p.next()
		}
		p.expectOp("]")
		p.pinRole(d)
	case t.Is("analog"):
		p.next()
		d.Kind = "analog"
		d.Pins = []PinRef{p.pinRef()}
	default:
		x := p.expr()
		if _, ok := x.(*Call); !ok {
			// "x = 5" at the top level: values belong in var/const, code in a block.
			p.errorf("P216", nameTok.Pos)
			p.syncLine()
			return nil
		}
		d.Kind, d.Object = "object", x
	}
	p.endLine()
	return d
}

func (p *parser) pinRef() PinRef {
	t := p.peek()
	if t.Kind == INT || t.Kind == NAME {
		p.next()
		return PinRef{At: t.Pos, Text: t.Text}
	}
	p.errorf("P208", t.Pos)
	return PinRef{At: t.Pos}
}

func (p *parser) pinRole(d *DeviceDecl) {
	if !p.peek().Is("as") {
		p.errorf("P208", p.peek().Pos)
		return
	}
	p.next()
	t := p.peek()
	d.RoleAt = t.Pos
	if t.Kind != NAME {
		p.errorf("P208", t.Pos)
		return
	}
	p.next()
	if !pinRoles[t.Text] {
		p.errorf("P209", t.Pos, "x", t.Text)
		return
	}
	d.Role = t.Text
	for p.peek().Kind == NAME && pinOptions[p.peek().Text] {
		d.Options = append(d.Options, p.next().Text)
	}
}

func (p *parser) fnDecl() Node {
	at := p.next().Pos
	fn := &FnDecl{At: at}
	fn.Name, _ = p.name()
	if p.expectOp("(") {
		for !p.peek().Is(")") && p.peek().Kind != NEWLINE && p.peek().Kind != EOF {
			fn.Params = append(fn.Params, p.param())
			if !p.peek().Is(",") {
				break
			}
			p.next()
		}
		p.expectOp(")")
	}
	if p.peek().Is("->") {
		p.next()
		fn.Result = p.typeExpr()
	}
	fn.Body = p.block()
	return fn
}

// param reads "name: type [= default]" (also used for struct fields).
func (p *parser) param() Param {
	pr := Param{At: p.peek().Pos}
	pr.Name, _ = p.name()
	if p.expectOp(":") {
		pr.Type = p.typeExpr()
	}
	if p.peek().Is("=") {
		p.next()
		pr.Default = p.expr()
	}
	return pr
}

func (p *parser) structDecl() Node {
	at := p.next().Pos
	s := &StructDecl{At: at}
	s.Name, _ = p.name()
	if !p.peek().Is(":") {
		p.errorf("P201", p.peek().Pos)
	} else {
		p.next()
	}
	if p.peek().Kind == NEWLINE {
		p.next()
	}
	if p.peek().Kind != INDENT {
		p.errorf("P202", p.peek().Pos)
		return s
	}
	p.next()
	for p.peek().Kind != DEDENT && p.peek().Kind != EOF {
		if p.peek().Kind == NEWLINE {
			p.next()
			continue
		}
		s.Fields = append(s.Fields, p.param())
		p.endLine()
	}
	p.next()
	return s
}

func (p *parser) enumDecl() Node {
	at := p.next().Pos
	e := &EnumDecl{At: at}
	e.Name, _ = p.name()
	p.expectOp(":")
	for {
		v, _ := p.name()
		e.Values = append(e.Values, v)
		if !p.peek().Is(",") {
			break
		}
		p.next()
	}
	p.endLine()
	return e
}

func (p *parser) eventBlock(top bool) Node {
	t := p.next()
	ev := &EventBlock{At: t.Pos, Kind: t.Text}
	switch t.Text {
	case "at":
		switch a := p.peek(); {
		case a.Kind == NAME && a.Text == "start":
			p.next()
			ev.Kind = "start"
		case a.Kind == TIMEOFDAY:
			p.next()
			ev.Kind, ev.Hour, ev.Min = "time", a.Hour, a.Min
		default:
			p.errorf("P205", a.Pos, "tok", tokText(a))
			p.syncLine()
			return nil
		}
	case "on":
		ev.Expr = p.postfix()
		if p.isName("fast") {
			p.next()
			ev.Fast = true
		}
	default:
		ev.Expr = p.expr()
		if t.Text == "when" {
			p.noAssignInCondition()
		}
	}
	if p.peek().Is("as") && p.peekAt(1).Kind == NAME {
		p.next()
		ev.As = p.next().Text
	}
	if top && (ev.Kind == "every" || ev.Kind == "after" || ev.Kind == "when") {
		ev.Priority = p.priority()
		ev.Core = p.core()
	}
	ev.Body = p.block()
	return ev
}

// core reads 'on core N' in a top-level event or task header (0: the main core, the default). 'on' is a keyword,
// 'core' a name; the checker keeps N on a real secondary core.
func (p *parser) core() int {
	if !(p.peek().Is("on") && p.peekAt(1).Kind == NAME && p.peekAt(1).Text == "core") {
		return 0
	}
	p.next()
	p.next()
	t := p.peek()
	if t.Kind != INT {
		p.errorf("P205", t.Pos, "tok", tokText(t))
		return 0
	}
	p.next()
	return int(t.Int)
}

// priority reads 'priority N' in a top-level event or task header (0: the default).
func (p *parser) priority() uint64 {
	if !p.isName("priority") {
		return 0
	}
	p.next()
	t := p.peek()
	if t.Kind != INT {
		p.errorf("P205", t.Pos, "tok", tokText(t))
		return 0
	}
	p.next()
	return t.Int
}

func (p *parser) controlDecl() Node {
	at := p.next().Pos
	c := &ControlDecl{At: at, Label: stringText(p.next())}
	p.expectOp(",")
	c.Var, _ = p.name()
	if p.peek().Is(",") {
		p.next()
		c.Range = p.expr()
		if p.peek().Is(",") { // control "Drive", steer, throttle, range: the first extra name is the stick's y axis
			p.next()
			if id, ok := c.Range.(*Ident); ok {
				c.Var2 = id.Name
			} else {
				p.errorf("P205", p.peek().Pos, "tok", ",")
			}
			c.Range = p.expr()
		} else if id, ok := c.Range.(*Ident); ok && id.Name == "hold" {
			c.Hold, c.Range = true, nil // control "Horn", horn, hold: true only while it's pressed
		}
	}
	p.endLine()
	return c
}

func (p *parser) limitDecl() Node {
	at := p.next().Pos
	l := &LimitDecl{At: at, Target: p.postfix()}
	if !p.isName("to") {
		p.errorf("P205", p.peek().Pos, "tok", tokText(p.peek()))
		p.syncLine()
		return l
	}
	p.next()
	l.Range = p.expr()
	p.endLine()
	return l
}

func (p *parser) cppBlock() Node {
	at := p.next().Pos
	c := &CppBlock{At: at}
	if !p.peek().Is(":") {
		p.errorf("P201", p.peek().Pos)
		p.syncLine()
		return c
	}
	p.next()
	if p.peek().Kind == RAW {
		raw := p.next()
		if raw.Text == "" {
			p.errorf("P202", raw.Pos)
		}
		c.Code = raw.Text
	}
	p.endLine()
	return c
}

// ---- blocks and statements

// block reads ": NEWLINE INDENT statements DEDENT" after a block header.
func (p *parser) block() *Block {
	b := &Block{At: p.peek().Pos}
	if p.peek().Is(":") {
		p.next()
	} else {
		p.errorf("P201", p.peek().Pos)
	}
	if t := p.peek(); t.Kind != NEWLINE && t.Kind != EOF {
		p.errorf("P205", t.Pos, "tok", tokText(t))
		for p.peek().Kind != NEWLINE && p.peek().Kind != EOF {
			p.next()
		}
	}
	if p.peek().Kind == NEWLINE {
		p.next()
	}
	if p.peek().Kind != INDENT {
		p.errorf("P202", p.peek().Pos)
		return b
	}
	b.At = p.next().Pos
	for p.peek().Kind != DEDENT && p.peek().Kind != EOF {
		start := p.i
		if s := p.stmt(); s != nil {
			b.Stmts = append(b.Stmts, s)
		}
		if p.i == start {
			p.next()
		}
	}
	p.next() // DEDENT
	return b
}

// topOnly are the declarations that can't appear inside a block.
var topOnly = map[string]bool{"use": true, "import": true, "saved": true, "fn": true, "struct": true, "enum": true,
	"every": true, "on": true, "at": true, "task": true}

func (p *parser) stmt() Node {
	t := p.peek()
	switch t.Kind {
	case NEWLINE:
		p.next()
		return nil
	case INDENT:
		p.errorf("P203", t.Pos)
		p.skipBlock()
		return nil
	case KEYWORD:
		switch t.Text {
		case "var", "const":
			return p.varDecl("")
		case "if":
			return p.ifStmt()
		case "elif", "else":
			p.errorf("P207", t.Pos, "kw", t.Text)
			p.syncLine()
			return nil
		case "while":
			p.next()
			w := &WhileStmt{At: t.Pos, Cond: p.expr()}
			p.noAssignInCondition()
			w.Body = p.block()
			return w
		case "for":
			p.next()
			f := &ForStmt{At: t.Pos}
			f.Var, _ = p.name()
			if !p.peek().Is("in") {
				p.errorf("P205", p.peek().Pos, "tok", tokText(p.peek()))
				p.syncLine()
				return nil
			}
			p.next()
			f.Iter = p.expr()
			f.Body = p.block()
			return f
		case "match":
			return p.matchStmt()
		case "break", "continue":
			p.next()
			p.endLine()
			return &SimpleStmt{At: t.Pos, Kind: t.Text}
		case "return":
			p.next()
			s := &SimpleStmt{At: t.Pos, Kind: "return"}
			if k := p.peek().Kind; k != NEWLINE && k != EOF && k != DEDENT {
				s.Value = p.expr()
			}
			p.endLine()
			return s
		case "wait":
			p.next()
			s := &SimpleStmt{At: t.Pos, Kind: "wait"}
			if p.isName("until") && startsExpr(p.peekAt(1)) { // a bare 'wait until' waits on a variable called until
				p.next()
				s.Until = true
			}
			s.Value = p.expr()
			p.endLine()
			return s
		case "after", "when":
			return p.eventBlock(false)
		case "cpp":
			return p.cppBlock()
		}
		if topOnly[t.Text] {
			p.errorf("P215", t.Pos, "kw", t.Text)
			p.syncLine()
			return nil
		}
	case NAME:
		next := p.peekAt(1)
		// 'print x' and 'expect x' only when a value follows; otherwise print/expect is an ordinary name ('print = 3').
		statementWord := startsExpr(next)
		switch {
		case t.Text == "pass" && next.Kind == NEWLINE:
			p.next()
			p.endLine()
			return &SimpleStmt{At: t.Pos, Kind: "pass"}
		case t.Text == "print" && statementWord:
			p.next()
			s := &SimpleStmt{At: t.Pos, Kind: "print", Value: p.expr()}
			p.endLine()
			return s
		case t.Text == "println" && statementWord:
			p.next()
			s := &SimpleStmt{At: t.Pos, Kind: "println", Value: p.expr()}
			p.endLine()
			return s
		case t.Text == "expect" && statementWord:
			p.next()
			e := &ExpectStmt{At: t.Pos, Cond: p.expr()}
			for p.peek().Is(",") {
				p.next()
				e.Args = append(e.Args, p.arg())
			}
			p.endLine()
			return e
		case (t.Text == "safe_state" && next.Is(":")) || ((t.Text == "test" || t.Text == "control") && next.Kind == STRING):
			p.errorf("P215", t.Pos, "kw", t.Text)
			p.syncLine()
			return nil
		}
		if c := p.command(); c != nil {
			return c
		}
	}
	if t.Is("board") {
		if c := p.command(); c != nil {
			return c
		}
	}
	return p.simpleStmt()
}

func (p *parser) simpleStmt() Node {
	at := p.peek().Pos
	x := p.expr()
	if op := p.peek(); op.Kind == OP && assignOps[op.Text] {
		p.next()
		switch x.(type) {
		case *Ident, *Member, *Index:
		default:
			p.errorf("P213", at)
		}
		s := &AssignStmt{At: at, Target: x, Op: op.Text, Value: p.expr()}
		p.endLine()
		return s
	}
	p.endLine()
	return &ExprStmt{At: at, X: x}
}

// noAssignInCondition reports "if x = 3:" and skips the "= 3".
func (p *parser) noAssignInCondition() {
	if t := p.peek(); t.Is("=") {
		p.errorf("P214", t.Pos)
		p.next()
		p.expr()
	}
}

func (p *parser) ifStmt() Node {
	s := &IfStmt{At: p.next().Pos}
	s.Conds = append(s.Conds, p.expr())
	p.noAssignInCondition()
	s.Thens = append(s.Thens, p.block())
	for {
		t := p.peek()
		if t.Is("elif") {
			p.next()
			s.Conds = append(s.Conds, p.expr())
			p.noAssignInCondition()
			s.Thens = append(s.Thens, p.block())
			continue
		}
		if t.Is("else") {
			p.next()
			s.Else = p.block()
		}
		return s
	}
}

func (p *parser) matchStmt() Node {
	m := &MatchStmt{At: p.next().Pos, Value: p.expr()}
	if p.peek().Is(":") {
		p.next()
	} else {
		p.errorf("P201", p.peek().Pos)
	}
	if p.peek().Kind == NEWLINE {
		p.next()
	}
	if p.peek().Kind != INDENT {
		p.errorf("P202", p.peek().Pos)
		return m
	}
	p.next()
	for p.peek().Kind != DEDENT && p.peek().Kind != EOF {
		if p.peek().Kind == NEWLINE {
			p.next()
			continue
		}
		c := MatchCase{At: p.peek().Pos}
		if p.peek().Is("else") {
			p.next()
			c.IsElse = true
		} else {
			for {
				c.Patterns = append(c.Patterns, p.expr())
				if !p.peek().Is(",") {
					break
				}
				p.next()
			}
		}
		start := p.i
		c.Body = p.block()
		m.Cases = append(m.Cases, c)
		if p.i == start {
			p.next()
		}
	}
	p.next()
	return m
}

// ---- expressions (spec section 3, from the loosest operator to the tightest)

func (p *parser) expr() Node { return p.condExpr() }

func (p *parser) condExpr() Node {
	x := p.orExpr()
	if t := p.peek(); t.Is("if") {
		p.next()
		c := &CondExpr{At: t.Pos, Then: x, Cond: p.orExpr()}
		if p.expectOp("else") {
			c.Else = p.condExpr()
		}
		return c
	}
	return x
}

func (p *parser) orExpr() Node {
	x := p.andExpr()
	for t := p.peek(); t.Is("or"); t = p.peek() {
		p.next()
		x = &Binary{At: t.Pos, Op: "or", X: x, Y: p.andExpr()}
	}
	return x
}

func (p *parser) andExpr() Node {
	x := p.notExpr()
	for t := p.peek(); t.Is("and"); t = p.peek() {
		p.next()
		x = &Binary{At: t.Pos, Op: "and", X: x, Y: p.notExpr()}
	}
	return x
}

func (p *parser) notExpr() Node {
	if t := p.peek(); t.Is("not") {
		p.next()
		return &Unary{At: t.Pos, Op: "not", X: p.notExpr()}
	}
	return p.compare()
}

func (p *parser) compare() Node {
	x := p.rangeExpr()
	if t := p.peek(); (t.Kind == OP && compareOps[t.Text]) || t.Is("in") {
		p.next()
		x = &Binary{At: t.Pos, Op: t.Text, X: x, Y: p.rangeExpr()}
		if t2 := p.peek(); (t2.Kind == OP && compareOps[t2.Text]) || t2.Is("in") {
			// a < b < c: one comparison at a time (x in a..b for a range)
			p.errorf("P212", t2.Pos)
			p.next()
			p.rangeExpr()
		}
	}
	return x
}

func (p *parser) rangeExpr() Node {
	x := p.conversion()
	if t := p.peek(); t.Is("..") || t.Is("..=") {
		p.next()
		return &RangeExpr{At: t.Pos, Lo: x, Hi: p.conversion(), Inclusive: t.Text == "..="}
	}
	return x
}

func (p *parser) conversion() Node {
	x := p.binary(0)
	if t := p.peek(); t.Is("as") {
		u := p.peekAt(1)
		if _, isUnit := Units[u.Text]; isUnit && (u.Kind == NAME || u.Kind == UNIT) {
			p.next()
			p.next()
			return &Conversion{At: t.Pos, X: x, Unit: u.Text}
		}
	}
	return x
}

// binaryLevels are the left-associative operator levels below conversion, loosest first. The product
// level also takes the context word 'mod' (the remainder operator, spec 2.5).
var binaryLevels = [][]string{{"|"}, {"^"}, {"&"}, {"<<", ">>"}, {"+", "-"}, {"*", "/"}}

func (p *parser) binary(level int) Node {
	if level == len(binaryLevels) {
		return p.unary()
	}
	x := p.binary(level + 1)
	for {
		t := p.peek()
		matched := false
		if t.Kind == OP {
			for _, op := range binaryLevels[level] {
				if t.Text == op {
					matched = true
				}
			}
		} else if level == len(binaryLevels)-1 && t.Kind == NAME && t.Text == "mod" {
			matched = true
		}
		if !matched {
			return x
		}
		p.next()
		x = &Binary{At: t.Pos, Op: t.Text, X: x, Y: p.binary(level + 1)}
	}
}

func (p *parser) unary() Node {
	if t := p.peek(); t.Is("-") || t.Is("~") {
		p.next()
		return &Unary{At: t.Pos, Op: t.Text, X: p.unary()}
	}
	return p.postfix()
}

func (p *parser) postfix() Node {
	x := p.primary()
	for {
		t := p.peek()
		switch {
		case t.Is("."):
			p.next()
			n := p.peek()
			if n.Kind != NAME {
				p.errorf("P204", n.Pos)
				return x
			}
			p.next()
			x = &Member{At: n.Pos, X: x, Name: n.Text}
		case t.Is("("):
			p.next()
			c := &Call{At: t.Pos, Fn: x}
			for !p.peek().Is(")") && p.peek().Kind != EOF && p.peek().Kind != NEWLINE {
				c.Args = append(c.Args, p.arg())
				if !p.peek().Is(",") {
					break
				}
				p.next()
			}
			p.expectOp(")")
			x = c
		case t.Is("["):
			p.next()
			x = &Index{At: t.Pos, X: x, Ix: p.expr()}
			p.expectOp("]")
		default:
			return x
		}
	}
}

// arg reads one call argument; "word: value" is a named argument, and the word may be a keyword (on: 60 C).
func (p *parser) arg() Arg {
	a := Arg{}
	if t := p.peek(); (t.Kind == NAME || t.Kind == KEYWORD) && p.peekAt(1).Is(":") {
		a.Name = t.Text
		p.next()
		p.next()
	}
	a.Value = p.expr()
	return a
}

// startsCommandArg reports whether a token can start the first argument of a call written without parentheses.
func startsCommandArg(t Token) bool {
	switch t.Kind {
	case STRING, INT, FLOAT:
		return true
	case NAME:
		return t.Text != "mod" // 'total mod i' is a remainder, not total(mod ...)
	case KEYWORD:
		return t.Text == "true" || t.Text == "false" || t.Text == "none" || t.Text == "not" || t.Text == "pin" || t.Text == "analog"
	}
	return false
}

// startsExpr reports whether a token can start an expression (spec 3, 'primary', 'unary' and 'not_expr').
func startsExpr(t Token) bool {
	switch t.Kind {
	case STRING, INT, FLOAT, NAME:
		return true
	case KEYWORD:
		switch t.Text {
		case "true", "false", "none", "not", "pin", "analog", "board":
			return true
		}
	case OP:
		return t.Text == "(" || t.Text == "[" || t.Text == "-" || t.Text == "~"
	}
	return false
}

// command parses a call written as a whole statement without parentheses: notify "Hi", show "Temp", t, log.info "x".
// It returns nil (consuming nothing) when the statement isn't one.
func (p *parser) command() Node {
	j := p.i
	if p.toks[j].Kind != NAME && !p.toks[j].Is("board") {
		return nil
	}
	j++
	for p.toks[j].Is(".") && p.toks[j+1].Kind == NAME {
		j += 2
	}
	if !startsCommandArg(p.toks[j]) {
		return nil
	}
	at := p.peek().Pos
	var fn Node = &Ident{At: p.toks[p.i].Pos, Name: p.toks[p.i].Text}
	p.next()
	for p.i < j {
		p.next()
		n := p.next()
		fn = &Member{At: n.Pos, X: fn, Name: n.Text}
	}
	c := &Call{At: at, Fn: fn}
	for {
		c.Args = append(c.Args, p.arg())
		if !p.peek().Is(",") {
			break
		}
		p.next()
	}
	p.endLine()
	return &ExprStmt{At: at, X: c}
}

func (p *parser) primary() Node {
	t := p.peek()
	switch {
	case t.Is("board") && p.peekAt(1).Is("."):
		// "board" is a keyword, but also the built-in object board.uptime, board.halt(...)
		p.next()
		return &Ident{At: t.Pos, Name: "board"}
	case t.Kind == NAME && p.peekAt(1).Kind == STRING:
		// A name followed by text is a call with that text: if ask "Ready?":
		p.next()
		s := p.next()
		return &Call{At: t.Pos, Fn: &Ident{At: t.Pos, Name: t.Text}, Args: []Arg{{Value: p.stringLit(s)}}}
	case t.Kind == INT || t.Kind == FLOAT:
		p.next()
		n := &NumberLit{At: t.Pos, Text: t.Text, IsFloat: t.Kind == FLOAT, Int: t.Int, Float: t.Float}
		if p.peek().Kind == UNIT {
			n.Unit = p.next().Text
		}
		return n
	case t.Kind == STRING:
		p.next()
		return p.stringLit(t)
	case t.Kind == NAME:
		p.next()
		return &Ident{At: t.Pos, Name: t.Text}
	case t.Is("true") || t.Is("false") || t.Is("none"):
		p.next()
		return &Literal{At: t.Pos, Value: t.Text}
	case t.Is("pin") || t.Is("analog"):
		p.next()
		return &PinLit{At: t.Pos, Analog: t.Text == "analog", Ref: p.pinRef()}
	case t.Is("("):
		p.next()
		x := p.expr()
		p.expectOp(")")
		return x
	case t.Is("["):
		p.next()
		l := &ListLit{At: t.Pos}
		for !p.peek().Is("]") && p.peek().Kind != EOF && p.peek().Kind != NEWLINE {
			l.Items = append(l.Items, p.expr())
			if !p.peek().Is(",") {
				break
			}
			p.next()
		}
		p.expectOp("]")
		return l
	case t.Kind == KEYWORD:
		// A keyword where a value should be ("var every = 3" is caught earlier; this is e.g. "x = every").
		p.errorf("P211", t.Pos, "kw", t.Text)
		p.next()
		return &Ident{At: t.Pos, Name: t.Text}
	}
	p.errorf("P204", t.Pos)
	return &Ident{At: t.Pos, Name: ""}
}

// stringLit parses the {expression} parts of a text literal.
func (p *parser) stringLit(t Token) Node {
	s := &StringLit{At: t.Pos}
	for _, part := range t.Parts {
		if !part.IsExpr {
			s.Parts = append(s.Parts, StringLitPart{Text: part.Text})
			continue
		}
		if len(part.Tokens) == 0 {
			p.errorf("P204", part.Pos)
			continue
		}
		toks := append(append([]Token{}, part.Tokens...), Token{Kind: EOF, Pos: part.Pos})
		sub := &parser{toks: toks}
		x := sub.expr()
		if e := sub.peek(); e.Kind != EOF {
			sub.errorf("P205", e.Pos, "tok", tokText(e))
		}
		for _, d := range sub.errs {
			p.add(d)
		}
		s.Parts = append(s.Parts, StringLitPart{Expr: x})
	}
	return s
}
