package lang

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// go test ./lang -run TestGoldenTrees -update rewrites testdata/trees/*.tree. Review the diff before keeping it.
var update = flag.Bool("update", false, "rewrite golden files")

func TestOkSamplesParse(t *testing.T) {
	for _, s := range loadSamples(t, "ok") {
		f, errs := ParseSource(s.src)
		if len(errs) > 0 {
			t.Errorf("%s: unexpected %v", s.name, errs)
		}
		if f.Board == "" || len(f.Items) == 0 {
			t.Errorf("%s: empty tree", s.name)
		}
	}
}

func TestBadSamplesParser(t *testing.T) {
	parserSamples := 0
	for _, s := range loadSamples(t, "bad") {
		_, errs := ParseSource(s.src)
		switch s.code[0] {
		case 'L':
			continue // the lexer's (TestBadSamplesLexer)
		case 'P':
			parserSamples++
			if len(errs) == 0 {
				t.Errorf("%s: expected %s at %s, got no error", s.name, s.code, s.pos)
			} else if errs[0].Code != s.code || errs[0].Pos != s.pos {
				t.Errorf("%s: expected %s at %s, got %v", s.name, s.code, s.pos, errs)
			} else if strings.Contains(errs[0].Message, "{") && s.code != "P206" {
				t.Errorf("%s: message not filled in: %q", s.name, errs[0].Message)
			}
		default:
			if len(errs) > 0 {
				t.Errorf("%s: the parser should accept this (its error is %s, a later phase), but reported %v", s.name, s.code, errs)
			}
		}
	}
	if parserSamples != 10 {
		t.Errorf("expected 10 parser (P) samples, found %d", parserSamples)
	}
}

func TestGoldenTrees(t *testing.T) {
	dir := filepath.Join("testdata", "trees")
	for _, s := range loadSamples(t, "ok") {
		f, _ := ParseSource(s.src)
		got := Dump(f)
		path := filepath.Join(dir, strings.TrimSuffix(s.name, ".muj")+".tree")
		if *update {
			if err := os.MkdirAll(dir, 0o755); err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(path, []byte(got), 0o644); err != nil {
				t.Fatal(err)
			}
			continue
		}
		want, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("%s: no golden tree (run with -update and review it): %v", s.name, err)
			continue
		}
		if got != strings.ReplaceAll(string(want), "\r\n", "\n") {
			t.Errorf("%s: tree changed\n--- got\n%s\n--- want\n%s", s.name, got, want)
		}
	}
}

// TestModuleFiles: a file loaded by 'import' parses like a program, except 'board' is optional (spec 10.1).
func TestModuleFiles(t *testing.T) {
	f, errs := ParseModuleSource("const top_speed = 800\nfn faster(v: int) -> int:\n    return v + 1\n")
	if FirstError(errs) != nil {
		t.Fatalf("a module needs no board line: %v", errs)
	}
	if f.Board != "" || len(f.Items) != 2 {
		t.Errorf("boardless module: %v", Dump(f))
	}
	f, errs = ParseModuleSource("board uno\nimport motors\n")
	if FirstError(errs) != nil {
		t.Fatalf("a board line and an import parse (the checker rejects the import): %v", errs)
	}
	if f.Board != "uno" || len(f.Items) != 1 {
		t.Errorf("board+import module: %v", Dump(f))
	}
	// ... while a program still insists on 'board'.
	if _, errs := ParseSource("const x = 1\n"); FirstError(errs) == nil || errs[0].Code != "P206" {
		t.Errorf("a program without 'board' must report P206: %v", errs)
	}
}

func parseOne(t *testing.T, src string) (*File, []Diagnostic) {
	t.Helper()
	return ParseSource("board uno\n" + src)
}

func TestExpressions(t *testing.T) {
	cases := []struct{ src, want string }{
		{"1 + 2 * 3", "(+ 1 (* 2 3))"},
		{"(1 + 2) * 3", "(* (+ 1 2) 3)"},
		{"a - b - c", "(- (- a b) c)"},
		{"not a and b or c", "(or (and (not a) b) c)"},
		{"a < b + 1", "(< a (+ b 1))"},
		{"x & 0xF0 | y << 2", "(| (& x 0xF0) (<< y 2))"},
		{"-x * 2", "(* (- x) 2)"},
		{"t * 343 m/s / 2 as cm", "(as (/ (* t 343m/s) 2) cm)"},
		{"0 %..100 %", "(.. 0% 100%)"},
		{"x in 1..10", "(in x (.. 1 10))"},
		{"t in 20 C..=30 C", "(in t (..= 20C 30C))"},
		{"1..=6", "(..= 1 6)"},
		{"a if c else b", "(ifelse c a b)"},
		{"led.blink(1 Hz, on: 100 ms)", "(call led.blink 1Hz on:100ms)"},
		{"leds[i].write(i < roll)", "(call leds[i].write (< i roll))"},
		{"Servo(pin 6)", "(call Servo (pin 6))"},
		{"f(analog A1, divider: 2)", "(call f (analog A1) divider:2)"},
		{"board.uptime > 12 h", "(> board.uptime 12h)"},
		{`"T: {round(t, 1)} C"`, `(text "T: " {(call round t 1)} " C")`},
		{"[1, 2, 3]", "[1 2 3]"},
		{`ask "Ready?"`, `(call ask (text "Ready?"))`},
		{"x mod 3", "(mod x 3)"},
	}
	for _, c := range cases {
		f, errs := parseOne(t, "var x = "+c.src+"\n")
		if len(errs) > 0 {
			t.Errorf("%s: %v", c.src, errs)
			continue
		}
		got := expr(f.Items[0].(*VarDecl).Value)
		if got != c.want {
			t.Errorf("%s\n got: %s\nwant: %s", c.src, got, c.want)
		}
	}
}

func TestStatements(t *testing.T) {
	src := `every 1 s as ticker:
    notify "Fire!", urgent: true
    show "Temperature", t
    log.info "Up {board.uptime}"
    serial 9600 baud
    board.halt("done")
    x += 1
    if a:
        b()
    elif c:
        d()
    else:
        e()
    after 3 s:
        f()
on serial.line as cmd:
    print cmd
`
	f, errs := parseOne(t, src)
	if len(errs) > 0 {
		t.Fatal(errs)
	}
	got := Dump(f)
	want := `(file uno
  (every 1s as ticker
    (do (call notify (text "Fire!") urgent:true))
    (do (call show (text "Temperature") t))
    (do (call log.info (text "Up " {board.uptime})))
    (do (call serial 9600baud))
    (do (call board.halt (text "done")))
    (+= x 1)
    (if a
      (do (call b))
    )
    (elif c
      (do (call d))
    )
    (else
      (do (call e))
    )
    (after 3s
      (do (call f))
    )
  )
  (on serial.line as cmd
    (print cmd)
  )
)
`
	if got != want {
		t.Errorf("\n got:\n%s\nwant:\n%s", got, want)
	}
}

func TestParserErrors(t *testing.T) {
	cases := []struct {
		src  string
		code string
		pos  Pos
	}{
		{"every 1 s:\n    fn f():\n        x()\n", "P215", Pos{3, 5}},
		{"f() = 3\n", "P216", Pos{2, 1}},
		{"at start:\n    f() = 3\n", "P213", Pos{3, 5}},
		{"x = 5\n", "P216", Pos{2, 1}},
		{"led = pin 13\n", "P208", Pos{2, 13}},
		{"at start:\n    var x = \n", "P204", Pos{3, 12}},
		{"at start:\n    print \"{}\"\n", "P204", Pos{3, 12}},
		{"at noon:\n    x()\n", "P205", Pos{2, 4}},
	}
	for _, c := range cases {
		_, errs := parseOne(t, c.src)
		if len(errs) == 0 || errs[0].Code != c.code || errs[0].Pos != c.pos {
			t.Errorf("%q: want %s at %s, got %v", c.src, c.code, c.pos, errs)
		}
	}
	// Recovery: every broken line is reported, and the good code around them still parses.
	f, errs := parseOne(t, "every 1 s:\n    x = = 1\n    y()\nevery 2 s\n    z()\nat start:\n    w()\n")
	if len(errs) != 2 {
		t.Errorf("want 2 errors, got %v", errs)
	}
	if len(f.Items) != 3 {
		t.Errorf("want 3 blocks after recovery, got %d:\n%s", len(f.Items), Dump(f))
	}
}

func FuzzParse(f *testing.F) {
	for _, kind := range []string{"ok", "bad"} {
		files, _ := filepath.Glob(filepath.Join(examplesDir, kind, "*.muj"))
		for _, file := range files {
			if b, err := os.ReadFile(file); err == nil {
				f.Add(string(b))
			}
		}
	}
	f.Fuzz(func(t *testing.T, src string) {
		file, errs := ParseSource(src)
		if file == nil {
			t.Fatal("nil file")
		}
		if len(errs) > 2*maxDiagnostics {
			t.Errorf("%d diagnostics", len(errs))
		}
		_ = Dump(file)
	})
}
