package lang

// Format is the one standard layout for .muj source - what "mujard fmt" writes and what the editors run on save.
// It works on the token stream, not the tree, so comments stay where the author put them: 4 spaces per indent
// level, one space between tokens where the grammar needs it, none around . ( ) [ ] .. or before , : ?,
// two spaces before a comment that was touching the code. Lines continued inside ( or [ are left alone.
import (
	"strings"
)

// Format returns src in the standard layout, or "" plus the lexer's diagnostics when the text can't be read
// (fmt refuses to guess a broken file). Leading tabs count as indentation and are fixed, not refused.
func Format(src string) (string, []Diagnostic) {
	lines := normalizeLines(src)
	toks, diags := Lex(strings.Join(lines, "\n"))
	if len(diags) > 0 {
		return "", diags
	}
	// The depth of each logical line comes from the INDENT/DEDENT tokens; the text is re-spaced per line.
	var out []string
	depth := 0
	phys := 0 // physical lines already emitted (blank and comment-only lines pass through verbatim)
	emitGap := func(upto int) {
		for phys < upto {
			out = append(out, lines[phys])
			phys++
		}
	}
	i := 0
	for i < len(toks) {
		switch toks[i].Kind {
		case INDENT:
			depth++
		case DEDENT:
			depth--
		}
		if toks[i].Kind != NAME && toks[i].Kind != KEYWORD && toks[i].Kind != INT && toks[i].Kind != FLOAT &&
			toks[i].Kind != UNIT && toks[i].Kind != STRING && toks[i].Kind != TIMEOFDAY && toks[i].Kind != OP &&
			toks[i].Kind != RAW {
			i++
			continue
		}
		// Collect one logical line: tokens until its NEWLINE.
		var lt []Token
		for i < len(toks) && toks[i].Kind != NEWLINE {
			lt = append(lt, toks[i])
			i++
		}
		var nl Token
		if i < len(toks) {
			nl = toks[i]
			i++
		}
		if len(lt) == 0 {
			continue
		}
		first, last := lt[0].Pos.Line, nl.Pos.Line
		emitGap(first - 1)
		if n := len(lt); n >= 3 && lt[n-1].Kind == RAW && lt[n-2].Is(":") && lt[0].Is("cpp") {
			// A cpp: block keeps its C++ as written; only the block's indentation is made canonical.
			colon := lt[n-2]
			out = append(out, emitLine(lt[:n-1], depth, comment(lines[colon.Pos.Line-1], colon)))
			if lt[n-1].Text != "" {
				for _, raw := range strings.Split(strings.TrimSuffix(lt[n-1].Text, "\n"), "\n") {
					if strings.TrimSpace(raw) == "" {
						out = append(out, "")
					} else {
						out = append(out, strings.Repeat(" ", 4*(depth+1))+raw)
					}
				}
			}
		} else if first != last {
			// A logical line continued inside ( or [ keeps its own line breaks.
			out = append(out, lines[first-1:last]...)
		} else {
			out = append(out, emitLine(lt, depth, comment(lines[first-1], lt[len(lt)-1])))
		}
		phys = last
	}
	emitGap(len(lines))
	// One trailing newline, no blank tail.
	for len(out) > 0 && strings.TrimSpace(out[len(out)-1]) == "" {
		out = out[:len(out)-1]
	}
	return strings.Join(out, "\n") + "\n", nil
}

// normalizeLines makes the text the formatter and the lexer agree on: \n endings, no BOM, a leading tab
// spelled as the 4 spaces the lexer counts it as, and no whitespace at a line's end.
func normalizeLines(src string) []string {
	src = strings.ReplaceAll(src, "\r\n", "\n")
	src = strings.TrimPrefix(src, "ï»¿") // a byte-order mark some editors add
	src = strings.ReplaceAll(src, "\r", "\n")
	lines := strings.Split(src, "\n")
	for i, l := range lines {
		indent := 0
		j := 0
		for j < len(l) && (l[j] == ' ' || l[j] == '\t') {
			if l[j] == '\t' {
				indent += 4
			} else {
				indent++
			}
			j++
		}
		lines[i] = strings.Repeat(" ", indent) + strings.TrimRight(l[j:], " \t")
	}
	return lines
}

// comment is the trailing "# ..." on the physical line after the logical line's last token, or "". The gap is
// kept when the author left one (aligned comments stay aligned); a touching one is given two spaces.
func comment(line string, last Token) string {
	r := []rune(line)
	end := last.Pos.Col - 1 + len([]rune(last.Text))
	if end >= len(r) {
		return ""
	}
	gap := string(r[end:])
	rest := strings.TrimLeft(gap, " ")
	if !strings.HasPrefix(rest, "#") {
		return ""
	}
	gap = gap[:len(gap)-len(rest)]
	if gap == "" {
		gap = "  "
	}
	return gap + rest
}

// emitLine rewrites one logical line: canonical indent, canonical spacing between tokens.
func emitLine(toks []Token, depth int, cmt string) string {
	var b strings.Builder
	b.WriteString(strings.Repeat(" ", 4*depth))
	for i, t := range toks {
		if i > 0 && spaceBetween(toks, i) {
			b.WriteByte(' ')
		}
		b.WriteString(t.Text)
	}
	return b.String() + cmt
}

// endsValue reports whether t closes something that yields a value (so a following ( or [ is a call or an
// index, and a following - or + is binary, not a sign).
func endsValue(t Token) bool {
	switch t.Kind {
	case NAME, INT, FLOAT, STRING, TIMEOFDAY, UNIT:
		return true
	case KEYWORD:
		return t.Is("true") || t.Is("false") || t.Is("none")
	}
	return t.Is(")") || t.Is("]") || t.Is("?")
}

// spaceBetween is the one place the spacing rules live: toks[i-1] and toks[i] are the neighbours; toks[:i-1]
// is the lookback that tells a unary sign from subtraction.
func spaceBetween(toks []Token, i int) bool {
	prev, cur := toks[i-1], toks[i]
	sign := func(t Token) bool { return t.Is("-") || t.Is("+") || t.Is("~") }
	switch {
	case cur.Is(")") || cur.Is("]") || cur.Is(",") || cur.Is(":") || cur.Is("?"):
		return false // "f(1)", "x[0]", "a, b", "value: 1", "int?"
	case cur.Is(".") || prev.Is("."):
		return false // "led.on"
	case cur.Is("..") || cur.Is("..=") || prev.Is("..") || prev.Is("..="):
		return false // "1..5"
	case prev.Is("(") || prev.Is("["):
		return false // "(a + b)", "[1, 2]"
	case (cur.Is("(") || cur.Is("[")) && endsValue(prev):
		return false // "f(1)", "x[0]" - but "= (a + b)", "= [1, 2]" keep their space
	case sign(prev) && (i < 2 || !endsValue(toks[i-2])):
		return false // "-1", "~flags", "+x": a unary sign hugs its value
	}
	return true
}
