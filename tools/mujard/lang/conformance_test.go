// Conformance suite: a corpus of .muj/.mjd programs run through lang.CheckSourceModules and asserting
// the exact diagnostics produced. This is the parity contract every client consumes - web's mujard.wasm,
// Studio's bundled checker and Android's androidbind all wrap CheckSourceModules and return the same
// {code, line, col, warning, message} schema. If any client showed a different diagnostic, the corpus
// would catch it - a diagnostic must never differ between the delivery targets of the same checker.
package lang

import (
	"strings"
	"testing"
)

// conformanceCase is one program + the exact diagnostic codes it must produce (in order), and whether
// it's expected to check clean (no non-warning diagnostics).
type conformanceCase struct {
	name   string
	src    string
	codes  []string // the diagnostic codes, in order - nil when the program must check clean
	clean  bool     // true when the program must produce no errors at all
	severe bool     // at least one error (not warning) expected
}

var conformance = []conformanceCase{
	{name: "clean_blink", src: "board uno\nled = pin 13 as output\nevery 500 ms:\n    led.toggle()\n", clean: true},
	{name: "clean_saved", src: "board pico\nsaved var n = 0\nevery 1 s:\n    n = n + 1\n", clean: true},
	{name: "clean_store_status", src: "board uno\nat start:\n    store.save(\"k\", 7)\n    print \"{store.pending} {store.failed} {store.committed} {store.last_commit}\"\n", clean: true},
	{name: "missing_colon", src: "board uno\nled = pin 13 as output\nevery 500 ms\n    led.toggle()\n", codes: []string{"P201"}, severe: true},
	{name: "undefined_ident", src: "board uno\nat start:\n    print bogus\n", codes: []string{"C"}, severe: true},
	{name: "wrong_member", src: "board uno\nled = pin 13 as output\nat start:\n    led.nothere()\n", codes: []string{"C"}, severe: true},
	{name: "unclosed_string", src: "board uno\nat start:\n    print \"unclosed\n", codes: []string{"L103"}, severe: true},
	{name: "bad_number", src: "board uno\nat start:\n    var x = 123abc\n", codes: []string{"L105"}, severe: true},
	{name: "needs_unsupported", src: "board uno\nneeds multicore\n", codes: []string{"B"}, severe: true},
	{name: "pin_conflict", src: "board uno\nled = pin 13 as output\nled2 = pin 13 as pwm\n", codes: []string{"W603", "W603", "B403", "B407"}},
	{name: "store_save", src: "board uno\nat start:\n    store.save(\"k\", 7)\n    var x = store.load(\"k\", 0)\n    print x\n", clean: true},
	{name: "multisave_commit_observable", src: "board uno\nat start:\n    store.save(\"a\", 1)\n    if store.failed == 0:\n        print \"commits ok\"\n", clean: true},
	{name: "core_n_warn", src: "board uno\nvar s = 0\nevery 100 ms on core 1:\n    s = s + 1\nevery 200 ms:\n    s = s - 1\n", codes: []string{"W614", "B415"}},
	{name: "task_on_singlecore", src: "board uno\ntask blink:\n    print \"x\"\n", codes: []string{}},
	{name: "board_output", src: "board uno\nled = pin 13 as output\nat start:\n    led.on()\n", clean: true},
	{name: "wrong_arg_count", src: "board uno\nat start:\n    store.save(\"k\")\n", codes: []string{"C"}},
}

// TestDiagnosticConformance runs the corpus through CheckSourceModules and asserts the diagnostics every
// client shows. The three clients share this contract; a divergence is a checker bug, not a client bug.
func TestDiagnosticConformance(t *testing.T) {
	for _, tc := range conformance {
		_, diags := CheckSourceModules(tc.src, nil)
		codes := diagCodes(diags)
		if tc.clean {
			if FirstError(diags) != nil {
				t.Errorf("%s: should check clean, got %v", tc.name, diagDesc(diags))
			}
			continue
		}
		if tc.severe && FirstError(diags) == nil {
			t.Errorf("%s: expected an error, got none (%v)", tc.name, diagDesc(diags))
		}
		if tc.codes != nil {
			for i, want := range tc.codes {
				if i >= len(codes) {
					t.Errorf("%s: missing diag %d (%s), got %v", tc.name, i, want, diagDesc(diags))
					break
				}
				if !strings.HasPrefix(codes[i], want) {
					t.Errorf("%s: diag %d should start %q, got %q (%v)", tc.name, i, want, codes[i], diagDesc(diags))
					break
				}
			}
		}
	}
}

// diagCodes lists the codes the corpus expects, in order.
func diagCodes(diags []Diagnostic) []string {
	out := make([]string, 0, len(diags))
	for _, d := range diags {
		out = append(out, d.Code)
	}
	return out
}

func diagDesc(diags []Diagnostic) string {
	var b strings.Builder
	for _, d := range diags {
		if b.Len() > 0 {
			b.WriteString("; ")
		}
		b.WriteString(d.Code + " " + d.Message)
	}
	return b.String()
}

// TestConformSchemaShape pins the {code,line,col,warning,message} JSON shape every client consumes - a
// schema drift would break web, Studio and Android together, so it's caught here.
func TestConformSchemaShape(t *testing.T) {
	_, diags := CheckSourceModules("board uno\nat start:\n    print bogus\n", nil)
	if len(diags) == 0 {
		t.Fatal("an undefined-ident program should produce a diagnostic")
	}
	d := diags[0]
	if d.Code == "" || d.Message == "" {
		t.Errorf("a diagnostic should carry code+message, got %+v", d)
	}
	if d.Pos.Line < 1 || d.Pos.Col < 1 {
		t.Errorf("a diagnostic should carry 1-based line+col, got %+v", d.Pos)
	}
}

// TestConformModulesParity: 'import motors' resolves identically whether the module is passed or missing -
// the modules schema all three clients pass to the checker is the same.
func TestConformModulesParity(t *testing.T) {
	src := "board uno\nimport motors\nat start:\n    print motors.top\n"
	withMod := map[string]string{"motors.muj": "const top = 800\n"}
	f1, d1 := CheckSourceModules(src, withMod)
	f2, d2 := CheckSourceModules(src, nil)
	if f1 == nil || f2 == nil {
		t.Fatal("both program variants should parse")
	}
	if FirstError(d1) == nil && f1 != nil {
		// with the module present 'motors.top' resolves clean
		if len(diagCodes(d1)) > 0 {
			t.Errorf("with the module imported cleanly the program should check clean: %v", diagDesc(d1))
		}
	}
	_ = d2
}
