package lang

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// examplesDir is docs/lang/examples, the language's shared test programs.
var examplesDir = filepath.Join("..", "..", "..", "docs", "lang", "examples")

type sample struct {
	name, src string
	dir       string // the folder the sample file is in (where 'import' looks for modules)
	code      string // expected first diagnostic for bad samples
	pos       Pos
}

var expectRx = regexp.MustCompile(`^# expect: ([A-Z]\d{3}) (\d+):(\d+)`)

func loadSamples(t testing.TB, kind string) []sample {
	t.Helper()
	files, err := filepath.Glob(filepath.Join(examplesDir, kind, "*.muj"))
	if err != nil || len(files) == 0 {
		t.Fatalf("no %s samples in %s", kind, examplesDir)
	}
	var out []sample
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		s := sample{name: filepath.Base(f), src: string(b), dir: filepath.Dir(f)}
		if m := expectRx.FindStringSubmatch(s.src); m != nil {
			s.code = m[1]
			fmt.Sscanf(m[2]+" "+m[3], "%d %d", &s.pos.Line, &s.pos.Col)
		}
		out = append(out, s)
	}
	return out
}

func TestOkSamplesLexClean(t *testing.T) {
	samples := loadSamples(t, "ok")
	if len(samples) != 42 {
		t.Errorf("expected 42 ok samples, found %d", len(samples))
	}
	for _, s := range samples {
		toks, errs := Lex(s.src)
		if len(errs) > 0 {
			t.Errorf("%s: unexpected %v", s.name, errs)
			continue
		}
		checkShape(t, s.name, toks)
	}
}

func TestBadSamplesLexer(t *testing.T) {
	samples := loadSamples(t, "bad")
	if len(samples) != 34 {
		t.Errorf("expected 34 bad samples, found %d", len(samples))
	}
	lexerSamples := 0
	for _, s := range samples {
		if s.code == "" {
			t.Errorf("%s: no '# expect:' line", s.name)
			continue
		}
		toks, errs := Lex(s.src)
		checkShape(t, s.name, toks)
		if !strings.HasPrefix(s.code, "L") {
			if len(errs) > 0 {
				t.Errorf("%s: the lexer should accept this (its error is %s, a later phase), but reported %v", s.name, s.code, errs)
			}
			continue
		}
		lexerSamples++
		if len(errs) == 0 {
			t.Errorf("%s: expected %s at %s, got no error", s.name, s.code, s.pos)
			continue
		}
		if errs[0].Code != s.code || errs[0].Pos != s.pos {
			t.Errorf("%s: expected %s at %s, got %s", s.name, s.code, s.pos, errs[0])
		}
		if errs[0].Message == "" || strings.Contains(errs[0].Message, "{c}") || strings.Contains(errs[0].Message, "{text}") {
			t.Errorf("%s: message not filled in: %q", s.name, errs[0].Message)
		}
	}
	if lexerSamples != 12 {
		t.Errorf("expected 12 lexer (L) samples, found %d", lexerSamples)
	}
}

// checkShape checks what every token stream must look like: ends in EOF, INDENT/DEDENT balanced, never-decreasing lines.
func checkShape(t testing.TB, name string, toks []Token) {
	t.Helper()
	if len(toks) == 0 || toks[len(toks)-1].Kind != EOF {
		t.Errorf("%s: token stream doesn't end in EOF", name)
		return
	}
	depth, line := 0, 0
	for _, tk := range toks {
		switch tk.Kind {
		case INDENT:
			depth++
		case DEDENT:
			depth--
		}
		if depth < 0 {
			t.Errorf("%s: DEDENT without INDENT at %s", name, tk.Pos)
			return
		}
		if tk.Pos.Line < line {
			t.Errorf("%s: token %s goes back to an earlier line", name, tk)
			return
		}
		line = tk.Pos.Line
	}
	if depth != 0 {
		t.Errorf("%s: %d INDENT(s) never closed", name, depth)
	}
}

// show renders tokens compactly: KIND:text, e.g. "KEYWORD:every INT:500 UNIT:ms OP:: NEWLINE".
func show(toks []Token) string {
	var b []string
	for _, t := range toks {
		switch t.Kind {
		case NEWLINE, INDENT, DEDENT, EOF:
			b = append(b, t.Kind.String())
		default:
			b = append(b, t.Kind.String()+":"+t.Text)
		}
	}
	return strings.Join(b, " ")
}

func TestTokenLists(t *testing.T) {
	cases := []struct{ src, want string }{
		{"board uno\nevery 500 ms:\n    led.toggle()\n",
			"KEYWORD:board NAME:uno NEWLINE KEYWORD:every INT:500 UNIT:ms OP:: NEWLINE INDENT NAME:led OP:. NAME:toggle OP:( OP:) NEWLINE DEDENT EOF"},
		// a keyword straight after '.' is a member name
		{"on button.pressed:\n    led.on()\n",
			"KEYWORD:on NAME:button OP:. NAME:pressed OP:: NEWLINE INDENT NAME:led OP:. NAME:on OP:( OP:) NEWLINE DEDENT EOF"},
		// units: written with or without a space, compound, percent (remainder is the word 'mod')
		{"a = 500ms", "NAME:a OP:= INT:500 UNIT:ms NEWLINE EOF"},
		{"d = t * 343 m/s / 2", "NAME:d OP:= NAME:t OP:* INT:343 UNIT:m/s OP:/ INT:2 NEWLINE EOF"},
		{"x = 50 %", "NAME:x OP:= INT:50 UNIT:% NEWLINE EOF"},
		{"x = 10 mod 3", "NAME:x OP:= INT:10 NAME:mod INT:3 NEWLINE EOF"},
		{"f(50 %)", "NAME:f OP:( INT:50 UNIT:% OP:) NEWLINE EOF"},
		{"r = 0 %..100 %", "NAME:r OP:= INT:0 UNIT:% OP:.. INT:100 UNIT:% NEWLINE EOF"},
		{"d = x as %", "NAME:d OP:= NAME:x KEYWORD:as UNIT:% NEWLINE EOF"},
		{"v = 0.5 V..1.5 V", "NAME:v OP:= FLOAT:0.5 UNIT:V OP:.. FLOAT:1.5 UNIT:V NEWLINE EOF"},
		// a unit name used as a function is not a unit; a unit name elsewhere is a name
		{"x = 3 min(a, b)", "NAME:x OP:= INT:3 NAME:min OP:( NAME:a OP:, NAME:b OP:) NEWLINE EOF"},
		{"var m = 3", "KEYWORD:var NAME:m OP:= INT:3 NEWLINE EOF"},
		// ranges are not decimal points
		{"for i in 0..=10:", "KEYWORD:for NAME:i KEYWORD:in INT:0 OP:..= INT:10 OP:: NEWLINE EOF"},
		// numbers
		{"a = 0xFF + 0b1010_0001 + 1_000 + 2.5e-3", "NAME:a OP:= INT:0xFF OP:+ INT:0b1010_0001 OP:+ INT:1_000 OP:+ FLOAT:2.5e-3 NEWLINE EOF"},
		// time of day only after 'at'
		{"at 07:30:", "KEYWORD:at TIMEOFDAY:07:30 OP:: NEWLINE EOF"},
		// optional types and arrows
		{"fn f(a: int) -> time?:", "KEYWORD:fn NAME:f OP:( NAME:a OP:: NAME:int OP:) OP:-> NAME:time OP:? OP:: NEWLINE EOF"},
		// newlines inside brackets don't end the line, and indentation there doesn't count
		{"x = f(1,\n      2)\ny = 3", "NAME:x OP:= NAME:f OP:( INT:1 OP:, INT:2 OP:) NEWLINE NAME:y OP:= INT:3 NEWLINE EOF"},
		// blank and comment lines don't affect indentation
		{"if a:\n\n    # note\n    b = 1\nc = 2", "KEYWORD:if NAME:a OP:: NEWLINE INDENT NAME:b OP:= INT:1 NEWLINE DEDENT NAME:c OP:= INT:2 NEWLINE EOF"},
		// several blocks closed at once
		{"a:\n    b:\n        c\nd", "NAME:a OP:: NEWLINE INDENT NAME:b OP:: NEWLINE INDENT NAME:c NEWLINE DEDENT DEDENT NAME:d NEWLINE EOF"},
	}
	for _, c := range cases {
		toks, errs := Lex(c.src)
		if len(errs) > 0 {
			t.Errorf("%q: unexpected %v", c.src, errs)
		}
		if got := show(toks); got != c.want {
			t.Errorf("%q\n got: %s\nwant: %s", c.src, got, c.want)
		}
	}
}

func TestPositions(t *testing.T) {
	toks, _ := Lex("board uno\nevery 1 s\n")
	// NEWLINE sits just after the last token of its line (where "Add ':'" points)
	for _, tk := range toks {
		if tk.Kind == NEWLINE && tk.Pos.Line == 2 && tk.Pos.Col != 10 {
			t.Errorf("NEWLINE of line 2 at %s, want 2:10", tk.Pos)
		}
	}
	toks, _ = Lex("x = 1\nif x:\n    y = 2\n")
	for _, tk := range toks {
		if tk.Kind == INDENT && tk.Pos != (Pos{3, 5}) {
			t.Errorf("INDENT at %s, want 3:5", tk.Pos)
		}
	}
}

func TestText(t *testing.T) {
	toks, errs := Lex(`print "T: {t * 2} C\n\t\"q\" \{x\}"`)
	if len(errs) > 0 {
		t.Fatal(errs)
	}
	s := toks[1]
	if s.Kind != STRING || len(s.Parts) != 3 {
		t.Fatalf("parts: %+v", s.Parts)
	}
	if s.Parts[0].Text != "T: " || !s.Parts[1].IsExpr || s.Parts[2].Text != " C\n\t\"q\" {x}" {
		t.Errorf("parts: %+v", s.Parts)
	}
	if got := show(s.Parts[1].Tokens); got != "NAME:t OP:* INT:2" {
		t.Errorf("interpolation tokens: %s", got)
	}
	if s.Parts[1].Tokens[0].Pos != (Pos{1, 12}) {
		t.Errorf("interpolated token at %s, want 1:12", s.Parts[1].Tokens[0].Pos)
	}
	// nested text inside an interpolation
	toks, errs = Lex(`x = "a{f("b")}c"`)
	if len(errs) > 0 || len(toks[2].Parts) != 3 || show(toks[2].Parts[1].Tokens) != `NAME:f OP:( STRING:"b" OP:)` {
		t.Errorf("nested text: %v %+v", errs, toks[2].Parts)
	}
}

func TestCppBlock(t *testing.T) {
	src := "board uno\ncpp:\n    int a = 1;\n\n      if (a) { a++; }\nevery 1 s:\n    x = 1\n"
	toks, errs := Lex(src)
	if len(errs) > 0 {
		t.Fatal(errs)
	}
	got := show(toks)
	want := "KEYWORD:board NAME:uno NEWLINE KEYWORD:cpp OP:: RAW:int a = 1;\n\n  if (a) { a++; }\n NEWLINE KEYWORD:every INT:1 UNIT:s OP:: NEWLINE INDENT NAME:x OP:= INT:1 NEWLINE DEDENT EOF"
	if got != want {
		t.Errorf("\n got: %q\nwant: %q", got, want)
	}
	// inside a block: raw lines must be deeper than the cpp: line
	toks, errs = Lex("every 1 s:\n    cpp:\n        foo();\n    x = 1\n")
	if len(errs) > 0 || !strings.Contains(show(toks), "RAW:foo();\n NEWLINE NAME:x") {
		t.Errorf("nested cpp: %v %s", errs, show(toks))
	}
}

func TestLexerErrors(t *testing.T) {
	cases := []struct {
		src  string
		code string
		pos  Pos
	}{
		{"at 25:00:", "L112", Pos{1, 4}},
		{"x = 5.", "L105", Pos{1, 5}},
		{"x = 5x7", "L105", Pos{1, 5}},
		{"x = 0x", "L105", Pos{1, 5}},
		{"x = 99999999999999999999", "L105", Pos{1, 5}},
		{"x = ]", "L111", Pos{1, 5}},
		{"x = (1]", "L111", Pos{1, 7}},
		{"x = {", "L104", Pos{1, 5}},
		{"x = a % b", "L104", Pos{1, 7}},
		{"x = \"a {b\"", "L107", Pos{1, 8}},
		{"x = \"{f(1}\"", "L110", Pos{1, 8}},
		{"x = \"é\" + ü", "L104", Pos{1, 11}},
	}
	for _, c := range cases {
		_, errs := Lex(c.src)
		if len(errs) == 0 || errs[0].Code != c.code || errs[0].Pos != c.pos {
			t.Errorf("%q: want %s at %s, got %v", c.src, c.code, c.pos, errs)
		}
	}
	// many errors are all reported, not just the first
	_, errs := Lex("a = $\nb = $\nc = $\n")
	if len(errs) != 3 {
		t.Errorf("want 3 errors, got %v", errs)
	}
}

func FuzzLex(f *testing.F) {
	for _, kind := range []string{"ok", "bad"} {
		files, _ := filepath.Glob(filepath.Join(examplesDir, kind, "*.muj"))
		for _, file := range files {
			if b, err := os.ReadFile(file); err == nil {
				f.Add(string(b))
			}
		}
	}
	f.Add("\"{\"{\"{")
	f.Add("cpp:\n\tx\n  y")
	f.Fuzz(func(t *testing.T, src string) {
		toks, errs := Lex(src)
		checkShape(t, "fuzz", toks)
		if len(errs) > maxDiagnostics {
			t.Errorf("%d diagnostics", len(errs))
		}
	})
}
