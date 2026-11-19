package lang

// The formatter is the standard layout from the examples: 4-space indents, canonical spacing, comments kept.
import (
	"testing"
)

func mustFormat(t *testing.T, src string) string {
	t.Helper()
	out, diags := Format(src)
	if len(diags) > 0 {
		t.Fatalf("Format(%q) refused: %v", src, diags)
	}
	return out
}

func TestFormatSpacing(t *testing.T) {
	cases := map[string]string{
		"board   uno":          "board uno",
		"led=pin 13 as output": "led = pin 13 as output",
		"every 500ms:":         "every 500 ms:",
		"led.toggle( )":        "led.toggle()",
		"if x==1:":             "if x == 1:",
		"x=-1":                 "x = -1",
		"x = a*-b":             "x = a * -b",
		"var r = Reading(value:20 C,taken:now())": "var r = Reading(value: 20 C, taken: now())",
		"fn f(n:int)->int:":                       "fn f(n: int) -> int:",
		"for i in 1..=n:":                         "for i in 1..=n:",
		"var l = [1,2 ,3]":                        "var l = [1, 2, 3]",
		"x = y[0]":                                "x = y[0]",
		"fan.write(clamp(v,0 %..100 %))":          "fan.write(clamp(v, 0 %..100 %))",
		"when soil.percent<30 %:":                 "when soil.percent < 30 %:",
		"print \"hi {x}\"":                        "print \"hi {x}\"",
	}
	for in, want := range cases {
		if got := mustFormat(t, in+"\n"); got != want+"\n" {
			t.Errorf("Format(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestFormatIndentation(t *testing.T) {
	src := "board uno\n\nevery 500 ms:\n  led.on()\n  if x == 1:\n            led.off()\n"
	want := "board uno\n\nevery 500 ms:\n    led.on()\n    if x == 1:\n        led.off()\n"
	if got := mustFormat(t, src); got != want {
		t.Errorf("Format re-indented wrongly:\n%s", got)
	}
	// tabs are indentation; fmt fixes them rather than refusing
	if got := mustFormat(t, "board uno\nat start:\n\tled.on()\n"); got != "board uno\nat start:\n    led.on()\n" {
		t.Errorf("tabs must become 4 spaces:\n%q", got)
	}
}

func TestFormatKeepsCommentsAndBlankLines(t *testing.T) {
	src := "# header\nboard uno\n\nled = pin LED as output        # the built-in one\n\n# a block comment\nevery 500 ms:\n    led.toggle()# blinking\n"
	want := "# header\nboard uno\n\nled = pin LED as output        # the built-in one\n\n# a block comment\nevery 500 ms:\n    led.toggle()  # blinking\n"
	if got := mustFormat(t, src); got != want {
		t.Errorf("Format lost a comment or blank line:\n%s", got)
	}
}

func TestFormatKeepsContinuedLines(t *testing.T) {
	src := "board uno\nvar x = f(a,\n         b)\n"
	if got := mustFormat(t, src); got != src {
		t.Errorf("a line continued inside ( keeps its breaks:\n%q", got)
	}
}

func TestFormatCppBlock(t *testing.T) {
	src := "board uno\nat start:\n    cpp:\n        Serial.println(1);\n          int x = 2;\n"
	want := "board uno\nat start:\n    cpp:\n        Serial.println(1);\n          int x = 2;\n"
	if got := mustFormat(t, src); got != want {
		t.Errorf("cpp: block: the keyword re-indents, the C++ keeps its relative indent:\n%s", got)
	}
}

func TestFormatRefusesBrokenSource(t *testing.T) {
	if out, diags := Format("board uno\nx = \"unclosed\n"); len(diags) == 0 || out != "" {
		t.Errorf("a file the lexer can't read must not be formatted: %q %v", out, diags)
	}
}

// Every shipped example formats to a file that still checks clean and is a fixed point.
func TestFormatExamplesAreFixedPoints(t *testing.T) {
	for _, s := range loadSamples(t, "ok") {
		out, diags := Format(s.src)
		if len(diags) > 0 {
			t.Fatalf("%s: the example must lex clean: %v", s.name, diags)
		}
		if _, d := CheckSourceIn(out, s.dir); FirstError(d) != nil {
			t.Errorf("%s: the formatted file must still check clean: %v", s.name, d)
		}
		if again := mustFormat(t, out); again != out {
			t.Errorf("%s: fmt is not idempotent:\n--- first\n%s\n--- second\n%s", s.name, out, again)
		}
	}
}
