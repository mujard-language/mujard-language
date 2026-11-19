package lang

import (
	"strconv"
	"strings"
	"unicode"
)

// maxDiagnostics stops a hopeless file from producing thousands of messages.
const maxDiagnostics = 100

// Lex turns a .muj source into tokens (spec section 2). It never stops at the first problem: it reports every lexer
// error it can (up to maxDiagnostics) and still returns a usable token stream, ending in EOF.
func Lex(src string) ([]Token, []Diagnostic) {
	src = strings.ReplaceAll(src, "\r\n", "\n")
	src = strings.TrimPrefix(src, string(rune(0xFEFF))) // a byte-order mark some editors add
	lx := &lexer{indents: []int{0}}
	for _, l := range strings.Split(src, "\n") {
		lx.lines = append(lx.lines, []rune(l))
	}
	lx.run()
	return lx.toks, lx.errs
}

type bracket struct {
	ch  rune
	pos Pos
}

type lexer struct {
	lines     [][]rune
	toks      []Token
	errs      []Diagnostic
	indents   []int
	brackets  []bracket // open ( and [ across lines
	seenFirst bool      // a non-blank logical line has been seen
	lineToks  int       // tokens produced so far on the current logical line
	lastEnd   Pos       // just after the last token of the current logical line
}

func (lx *lexer) errorf(code string, pos Pos, args ...string) {
	if len(lx.errs) < maxDiagnostics {
		lx.errs = append(lx.errs, diag(code, pos, args...))
	}
}

func (lx *lexer) run() {
	for i := 0; i < len(lx.lines); i++ {
		r := lx.lines[i]
		line := i + 1
		start := 0
		if len(lx.brackets) == 0 {
			// The start of a logical line: indentation decides INDENT / DEDENT.
			indent, tabAt := 0, -1
			for start < len(r) && (r[start] == ' ' || r[start] == '\t') {
				if r[start] == '\t' {
					if tabAt < 0 {
						tabAt = start
					}
					indent += 4
				} else {
					indent++
				}
				start++
			}
			if start == len(r) || r[start] == '#' {
				continue // blank or comment-only lines don't count
			}
			if tabAt >= 0 {
				lx.errorf("L101", Pos{line, tabAt + 1})
			}
			if !lx.seenFirst && indent > 0 {
				lx.errorf("L109", Pos{line, start + 1})
				indent = 0
			}
			lx.seenFirst = true
			lx.indentTo(indent, Pos{line, start + 1})
			lx.lineToks = 0
		}
		sc := &scanner{lx: lx, r: r, line: line, out: &lx.toks, brackets: &lx.brackets}
		sc.scan(start)
		if len(lx.brackets) > 0 {
			continue // the logical line goes on inside ( or [
		}
		if lx.lineToks == 0 {
			continue
		}
		// "cpp:" at the end of a logical line: the indented lines below are raw C++.
		if n := len(lx.toks); lx.lineToks >= 2 && lx.toks[n-1].Is(":") && lx.toks[n-2].Is("cpp") {
			i = lx.captureRaw(i)
		}
		lx.toks = append(lx.toks, Token{Kind: NEWLINE, Pos: lx.lastEnd})
		lx.lineToks = 0
	}
	end := Pos{len(lx.lines) + 1, 1}
	if n := len(lx.brackets); n > 0 {
		b := lx.brackets[n-1]
		lx.errorf("L110", b.pos, "open", string(b.ch))
		lx.brackets = nil
		lx.toks = append(lx.toks, Token{Kind: NEWLINE, Pos: lx.lastEnd})
	}
	for len(lx.indents) > 1 {
		lx.indents = lx.indents[:len(lx.indents)-1]
		lx.toks = append(lx.toks, Token{Kind: DEDENT, Pos: end})
	}
	lx.toks = append(lx.toks, Token{Kind: EOF, Pos: end})
}

func (lx *lexer) indentTo(indent int, pos Pos) {
	top := lx.indents[len(lx.indents)-1]
	if indent > top {
		lx.indents = append(lx.indents, indent)
		lx.toks = append(lx.toks, Token{Kind: INDENT, Pos: pos})
		return
	}
	for indent < lx.indents[len(lx.indents)-1] {
		lx.indents = lx.indents[:len(lx.indents)-1]
		lx.toks = append(lx.toks, Token{Kind: DEDENT, Pos: pos})
	}
	if indent != lx.indents[len(lx.indents)-1] {
		lx.errorf("L102", pos)
		// Recover by treating this line as a block of its own level.
		lx.indents = append(lx.indents, indent)
		lx.toks = append(lx.toks, Token{Kind: INDENT, Pos: pos})
	}
}

// captureRaw takes the lines after "cpp:" that are indented more than the current block, emits them as one RAW token
// (with their common indentation removed) and returns the index of the last line taken.
func (lx *lexer) captureRaw(i int) int {
	base := lx.indents[len(lx.indents)-1]
	last := i
	minIndent := -1
	for j := i + 1; j < len(lx.lines); j++ {
		r := lx.lines[j]
		ind := 0
		for ind < len(r) && (r[ind] == ' ' || r[ind] == '\t') {
			ind++
		}
		if ind == len(r) {
			continue // blank lines belong to the block only if more raw lines follow
		}
		if ind <= base {
			break
		}
		if minIndent < 0 || ind < minIndent {
			minIndent = ind
		}
		last = j
	}
	if last == i {
		lx.toks = append(lx.toks, Token{Kind: RAW, Pos: lx.lastEnd})
		return i
	}
	var b strings.Builder
	for j := i + 1; j <= last; j++ {
		r := lx.lines[j]
		if len(r) > minIndent {
			b.WriteString(string(r[minIndent:]))
		}
		b.WriteByte('\n')
	}
	lx.toks = append(lx.toks, Token{Kind: RAW, Text: b.String(), Pos: Pos{i + 2, minIndent + 1}})
	lx.lastEnd = Pos{last + 1, len(lx.lines[last]) + 1} // the logical line ends after the raw block
	return last
}

// scanner reads the tokens of one line (or of one {interpolation} inside text).
type scanner struct {
	lx       *lexer
	r        []rune
	line     int
	colBase  int // added to rune indexes to get columns (for interpolations)
	out      *[]Token
	brackets *[]bracket
}

func (sc *scanner) pos(i int) Pos { return Pos{sc.line, sc.colBase + i + 1} }

func (sc *scanner) emit(t Token, end int) {
	*sc.out = append(*sc.out, t)
	if sc.out == &sc.lx.toks {
		sc.lx.lineToks++
		sc.lx.lastEnd = sc.pos(end)
	}
}

func (sc *scanner) prev() *Token {
	if n := len(*sc.out); n > 0 {
		return &(*sc.out)[n-1]
	}
	return nil
}

func isLetter(c rune) bool { return c == '_' || (c < 128 && unicode.IsLetter(c)) }
func isDigit(c rune) bool  { return c >= '0' && c <= '9' }
func isWord(c rune) bool   { return isLetter(c) || isDigit(c) }

func (sc *scanner) at(i int) rune {
	if i >= 0 && i < len(sc.r) {
		return sc.r[i]
	}
	return 0
}

func (sc *scanner) scan(i int) {
	r := sc.r
	for i < len(r) {
		c := r[i]
		switch {
		case c == ' ' || c == '\t':
			i++
		case c == '#':
			return
		case isLetter(c):
			j := i
			for j < len(r) && isWord(r[j]) {
				j++
			}
			word := string(r[i:j])
			kind := NAME
			if p := sc.prev(); Keywords[word] && !(p != nil && p.Is(".")) {
				kind = KEYWORD
			}
			sc.emit(Token{Kind: kind, Text: word, Pos: sc.pos(i)}, j)
			i = j
		case isDigit(c):
			i = sc.number(i)
		case c == '.' && isDigit(sc.at(i+1)) && !sc.afterValue():
			// ".5": a number needs a digit before the point.
			j := i + 1
			for j < len(r) && (isWord(r[j]) || r[j] == '.') {
				j++
			}
			sc.lx.errorf("L105", sc.pos(i), "text", string(r[i:j]))
			sc.emit(Token{Kind: FLOAT, Text: string(r[i:j]), Pos: sc.pos(i)}, j)
			i = j
		case c == '"':
			i = sc.text(i)
		case c == '\'':
			sc.lx.errorf("L106", sc.pos(i))
			j := i + 1
			for j < len(r) && r[j] != '\'' {
				j++
			}
			sc.emit(Token{Kind: STRING, Text: string(r[i:min(j+1, len(r))]), Pos: sc.pos(i), Parts: []StringPart{{Text: string(r[i+1 : j]), Pos: sc.pos(i)}}}, min(j+1, len(r)))
			i = j + 1
		case c == '(' || c == '[':
			*sc.brackets = append(*sc.brackets, bracket{c, sc.pos(i)})
			sc.emit(Token{Kind: OP, Text: string(c), Pos: sc.pos(i)}, i+1)
			i++
		case c == ')' || c == ']':
			open := '('
			if c == ']' {
				open = '['
			}
			n := len(*sc.brackets)
			if n == 0 || (*sc.brackets)[n-1].ch != open {
				sc.lx.errorf("L111", sc.pos(i), "open", string(open), "close", string(c))
			} else {
				*sc.brackets = (*sc.brackets)[:n-1]
				sc.emit(Token{Kind: OP, Text: string(c), Pos: sc.pos(i)}, i+1)
			}
			i++
		case c == '%':
			// '%' is only a unit: straight after a number (handled in number->unit), or after 'as'
			// in a conversion ('d as %'). Anywhere else it isn't a Mujard character (L104).
			if p := sc.prev(); p != nil && p.Is("as") {
				sc.emit(Token{Kind: UNIT, Text: "%", Pos: sc.pos(i)}, i+1)
			} else {
				sc.lx.errorf("L104", sc.pos(i), "c", "%")
			}
			i++
		default:
			if op := sc.operator(i); op != "" {
				sc.emit(Token{Kind: OP, Text: op, Pos: sc.pos(i)}, i+len(op))
				i += len(op)
				continue
			}
			sc.lx.errorf("L104", sc.pos(i), "c", string(c))
			i++
		}
	}
}

// afterValue reports whether the previous token ends a value (so a following '.' is member access).
func (sc *scanner) afterValue() bool {
	p := sc.prev()
	return p != nil && (p.Kind == NAME || p.Is(")") || p.Is("]"))
}

func (sc *scanner) operator(i int) string {
	for _, op := range operators {
		if strings.HasPrefix(string(sc.r[i:min(i+len(op), len(sc.r))]), op) {
			return op
		}
	}
	return ""
}

// number reads a number, an optional unit after it, or a time of day after 'at'. It returns the index after them.
func (sc *scanner) number(i int) int {
	r := sc.r
	// 07:30 after 'at'
	if p := sc.prev(); p != nil && p.Is("at") {
		j := i
		for j < len(r) && isDigit(r[j]) {
			j++
		}
		if j-i <= 2 && sc.at(j) == ':' && isDigit(sc.at(j+1)) && isDigit(sc.at(j+2)) && !isDigit(sc.at(j+3)) {
			h, _ := strconv.Atoi(string(r[i:j]))
			m, _ := strconv.Atoi(string(r[j+1 : j+3]))
			text := string(r[i : j+3])
			if h > 23 || m > 59 {
				sc.lx.errorf("L112", sc.pos(i), "text", text)
			}
			sc.emit(Token{Kind: TIMEOFDAY, Text: text, Pos: sc.pos(i), Hour: h, Min: m}, j+3)
			return j + 3
		}
	}

	j := i
	base := 10
	isFloat := false
	if r[i] == '0' && (sc.at(i+1) == 'x' || sc.at(i+1) == 'X') {
		base, j = 16, i+2
		for j < len(r) && (isDigit(r[j]) || r[j] == '_' || strings.ContainsRune("abcdefABCDEF", r[j])) {
			j++
		}
	} else if r[i] == '0' && (sc.at(i+1) == 'b' || sc.at(i+1) == 'B') {
		base, j = 2, i+2
		for j < len(r) && (r[j] == '0' || r[j] == '1' || r[j] == '_') {
			j++
		}
	} else {
		for j < len(r) && (isDigit(r[j]) || r[j] == '_') {
			j++
		}
		if sc.at(j) == '.' && isDigit(sc.at(j+1)) {
			isFloat = true
			j++
			for j < len(r) && (isDigit(r[j]) || r[j] == '_') {
				j++
			}
		} else if sc.at(j) == '.' && sc.at(j+1) != '.' && !isLetter(sc.at(j+1)) {
			j++ // "5." - reported below
			return sc.badNumber(i, j)
		}
		if (sc.at(j) == 'e' || sc.at(j) == 'E') && (isDigit(sc.at(j+1)) || ((sc.at(j+1) == '+' || sc.at(j+1) == '-') && isDigit(sc.at(j+2)))) {
			isFloat = true
			j += 2
			for j < len(r) && isDigit(r[j]) {
				j++
			}
		}
	}
	text := string(r[i:j])
	digits := text
	if base != 10 {
		digits = text[2:]
	}
	tok := Token{Kind: INT, Text: text, Pos: sc.pos(i)}
	ok := digits != "" && !strings.Contains(digits, "__") && !strings.HasPrefix(digits, "_") && !strings.HasSuffix(digits, "_") &&
		!strings.Contains(digits, "_.") && !strings.Contains(digits, "._")
	clean := strings.ReplaceAll(digits, "_", "")
	if ok {
		if isFloat {
			f, err := strconv.ParseFloat(clean, 64)
			ok = err == nil
			tok.Kind, tok.Float = FLOAT, f
		} else {
			v, err := strconv.ParseUint(clean, base, 64)
			ok = err == nil
			tok.Int = v
		}
	}
	// Letters straight after the digits: a unit written without a space (500ms), or a broken number (5x7).
	if isLetter(sc.at(j)) {
		k := j
		for k < len(r) && isWord(r[k]) {
			k++
		}
		if _, unit := Units[string(r[j:k])]; !unit || sc.at(k) == '(' {
			return sc.badNumber(i, k)
		}
	}
	if !ok {
		return sc.badNumber(i, j)
	}
	sc.emit(tok, j)
	return sc.unit(j)
}

func (sc *scanner) badNumber(i, j int) int {
	text := string(sc.r[i:j])
	sc.lx.errorf("L105", sc.pos(i), "text", text)
	sc.emit(Token{Kind: INT, Text: text, Pos: sc.pos(i)}, j)
	return j
}

// unit reads a unit straight after a number, if there is one (spec 2.7).
func (sc *scanner) unit(i int) int {
	r := sc.r
	j := i
	for j < len(r) && r[j] == ' ' {
		j++
	}
	if sc.at(j) == '%' {
		// After a number, % is always the ratio unit (remainder is the word 'mod').
		sc.emit(Token{Kind: UNIT, Text: "%", Pos: sc.pos(j)}, j+1)
		return j + 1
	}
	if !isLetter(sc.at(j)) {
		return i
	}
	k := j
	for k < len(r) && isWord(r[k]) {
		k++
	}
	word := string(r[j:k])
	if _, ok := Units[word]; !ok || sc.at(k) == '(' {
		return i
	}
	// A compound unit: m/s, mm/s (no spaces around '/').
	if sc.at(k) == '/' && isLetter(sc.at(k+1)) {
		m := k + 1
		for m < len(r) && isWord(r[m]) {
			m++
		}
		if _, ok := Units[string(r[k+1:m])]; ok {
			word, k = string(r[j:m]), m
		}
	}
	sc.emit(Token{Kind: UNIT, Text: word, Pos: sc.pos(j)}, k)
	return k
}

// text reads a "..." literal starting at i, with escapes and {expression} parts. It returns the index after it.
func (sc *scanner) text(i int) int {
	r := sc.r
	var parts []StringPart
	var lit strings.Builder
	litPos := sc.pos(i + 1)
	flush := func() {
		if lit.Len() > 0 {
			parts = append(parts, StringPart{Text: lit.String(), Pos: litPos})
			lit.Reset()
		}
	}
	j := i + 1
	for {
		if j >= len(r) {
			sc.lx.errorf("L103", sc.pos(i))
			flush()
			sc.emit(Token{Kind: STRING, Text: string(r[i:]), Pos: sc.pos(i), Parts: parts}, len(r))
			return len(r)
		}
		c := r[j]
		switch {
		case c == '"':
			flush()
			sc.emit(Token{Kind: STRING, Text: string(r[i : j+1]), Pos: sc.pos(i), Parts: parts}, j+1)
			return j + 1
		case c == '\\':
			e := sc.at(j + 1)
			switch e {
			case 'n':
				lit.WriteRune('\n')
			case 't':
				lit.WriteRune('\t')
			case '"', '\\', '{', '}':
				lit.WriteRune(e)
			case 0:
				j++
				continue // "\" at the end of the line: the text is unclosed, reported above on the next turn
			default:
				sc.lx.errorf("L108", sc.pos(j), "c", string(e))
				lit.WriteRune(e)
			}
			j += 2
		case c == '{':
			end := sc.interpolationEnd(j)
			if end < 0 {
				sc.lx.errorf("L107", sc.pos(j))
				flush()
				sc.emit(Token{Kind: STRING, Text: string(r[i:]), Pos: sc.pos(i), Parts: parts}, len(r))
				return len(r)
			}
			flush()
			var toks []Token
			var br []bracket
			sub := &scanner{lx: sc.lx, r: r[j+1 : end], line: sc.line, colBase: sc.colBase + j + 1, out: &toks, brackets: &br}
			sub.scan(0)
			if n := len(br); n > 0 {
				sc.lx.errorf("L110", br[n-1].pos, "open", string(br[n-1].ch))
			}
			parts = append(parts, StringPart{Tokens: toks, Pos: sc.pos(j), IsExpr: true})
			j = end + 1
			litPos = sc.pos(j)
		default:
			if lit.Len() == 0 {
				litPos = sc.pos(j)
			}
			lit.WriteRune(c)
			j++
		}
	}
}

// interpolationEnd finds the '}' that closes the '{' at i, skipping nested braces and nested "text". -1 if there is none
// on this line (or the text ends first).
func (sc *scanner) interpolationEnd(i int) int {
	r := sc.r
	depth := 1
	for j := i + 1; j < len(r); j++ {
		switch r[j] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return j
			}
		case '"':
			// A nested text inside the expression: it must close on this line.
			k := j + 1
			for k < len(r) && r[k] != '"' {
				if r[k] == '\\' {
					k++
				}
				k++
			}
			if k >= len(r) {
				return -1
			}
			j = k
		}
	}
	return -1
}
