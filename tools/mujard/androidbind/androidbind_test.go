// Verify the gomobile-bound functions produce the same diagnostics as the checker they wrap: Check and
// Secrets call lang.CheckSourceModules/CheckProgramModules, so the JSON the Android .so returns matches
// what mujard.wasm and 'mujard check' produce for the same program. libgojni's native load is on-device,
// but the binding contract is host-verifiable - the Go function the .so exposes is directly testable.
package androidbind

import (
	"encoding/json"
	"mujard/lang"
	"strings"
	"testing"
)

func TestCheckGoodProgram(t *testing.T) {
	got := Check("board uno\nled = pin 13 as output\nevery 500 ms:\n    led.toggle()\n", "")
	if strings.Contains(got, "\"code\":\"C") || strings.Contains(got, "\"code\":\"P") || strings.Contains(got, "\"code\":\"L") {
		t.Errorf("a clean program should produce no parse/semantic errors, got %s", got)
	}
}

func TestCheckErrorProgram(t *testing.T) {
	got := Check("board uno\nat start:\n    print bogus\n", "")
	var diags []diagOut
	if err := json.Unmarshal([]byte(got), &diags); err != nil {
		t.Fatalf("Check should return JSON: %v (%s)", err, got)
	}
	if len(diags) == 0 {
		t.Fatal("an undefined-ident program should produce a diagnostic")
	}
	d := diags[0]
	if d.Code == "" || d.Message == "" || d.Line < 1 || d.Col < 1 {
		t.Errorf("each diagnostic should carry code/line/col/message, got %+v", d)
	}
	if d.Warning {
		t.Errorf("an undefined-ident should be an error, not a warning, got %+v", d)
	}
}

func TestCheckModules(t *testing.T) {
	// 'import motors' resolves against the modules map, like the browser passes the project's other files.
	src := "board uno\nimport motors\nat start:\n    print motors.top\n"
	withMod := `{"motors.muj":"const top = 800\n"}`
	got := Check(src, withMod)
	if strings.Contains(got, "\"code\":\"C") {
		t.Errorf("with the module imported cleanly the program should check clean: %s", got)
	}
}

func TestCheckSchemaParity(t *testing.T) {
	// The JSON shape every client consumes: {code, file?, line, col, warning, message}.
	got := Check("board uno\nat start:\n    print bogus\n", "")
	var raw []map[string]any
	if err := json.Unmarshal([]byte(got), &raw); err != nil {
		t.Fatalf("Check should return a JSON array: %v", err)
	}
	for _, key := range []string{"code", "line", "col", "warning", "message"} {
		if _, ok := raw[0][key]; !ok {
			t.Errorf("the diagnostic JSON should have a %q key, got %v", key, raw[0])
		}
	}
}

func TestSecrets(t *testing.T) {
	got := Secrets("board uno\nat start:\n    var x = 0\n    print x\n", "")
	var names []string
	if err := json.Unmarshal([]byte(got), &names); err != nil {
		t.Fatalf("Secrets should return a JSON array: %v (%s)", err, got)
	}
	// A program needing no secrets returns [] - names only, never values.
}

// Editor-services conformance: the .so-bound completionList/hoverText/definitions/references/renameEdits
// must produce the same output as the shared lang.EditorService they wrap - the phone's editor
// intelligence is the same code the LSP runs, host-verifiable like the diagnostics contract.
func TestCompletionConformance(t *testing.T) {
	src := "board uno\nled = pin 13 as output\nevery 500 ms:\n    led.\n"
	// 'led.' at line 3 (0-based), col 8 - receiver-scoped: the output members, not every object's.
	got := CompletionList(src, 3, 8)
	var out map[string]any
	if err := json.Unmarshal([]byte(got), &out); err != nil {
		t.Fatalf("CompletionList should return JSON: %v", err)
	}
	items, _ := out["items"].([]any)
	labels := map[string]bool{}
	for _, it := range items {
		labels[it.(map[string]any)["label"].(string)] = true
	}
	for _, want := range []string{"on", "off", "toggle"} {
		if !labels[want] {
			t.Errorf("'led.' should offer the output member %q, got %v", want, labels)
		}
	}
	// The receiver-scoped result shouldn't have every object's members - 'notify' belongs to a different object.
	if labels["notify"] {
		t.Errorf("'led.' should be receiver-scoped (output members), got notify too: %v", labels)
	}
}

func TestCompletionParityWithLang(t *testing.T) {
	// The .so function returns the same JSON lang.EditorService.Complete produces for the same input.
	src := "board uno\nled = pin 13 as output\nat start:\n    led.\n"
	bound := CompletionList(src, 3, 8)
	direct := lang.NewEditorService(lang.NewEditorInfo()).Complete(src, lang.LspPos{Line: 3, Character: 8})
	want, _ := json.Marshal(direct)
	if bound != string(want) {
		t.Errorf("the .so completion should match lang.EditorService:\n  .so: %s\n  lang: %s", bound, want)
	}
}

func TestHoverConformance(t *testing.T) {
	src := "board uno\nled = pin 13 as output\nat start:\n    led.toggle()\n"
	got := HoverText(src, 3, 8)
	if got == "" {
		t.Fatal("hover over 'led' should describe it")
	}
	var out map[string]any
	if err := json.Unmarshal([]byte(got), &out); err != nil {
		t.Fatalf("HoverText should return JSON: %v (%s)", err, got)
	}
	contents, _ := out["contents"].(map[string]any)
	if contents == nil || contents["value"] == "" {
		t.Errorf("hover should carry contents.value markdown, got %v", out)
	}
}

func TestDefinitionConformance(t *testing.T) {
	src := "board uno\nled = pin 13 as output\nat start:\n    led.on()\n"
	got := Definitions(src, "main.muj", 3, 4)
	var locs []map[string]any
	if err := json.Unmarshal([]byte(got), &locs); err != nil {
		t.Fatalf("Definitions should return a JSON array: %v (%s)", err, got)
	}
	if len(locs) == 0 {
		t.Fatal("'led' should resolve to its declaration")
	}
}

func TestReferencesConformance(t *testing.T) {
	src := "board uno\nled = pin 13 as output\nat start:\n    led.on()\n    led.off()\n"
	got := References(src, "main.muj", 3, 4)
	var locs []any
	if err := json.Unmarshal([]byte(got), &locs); err != nil {
		t.Fatalf("References should return a JSON array: %v (%s)", err, got)
	}
	if len(locs) < 3 {
		t.Errorf("'led' is used 3 times (decl + on + off), got %d", len(locs))
	}
}

func TestRenameConformance(t *testing.T) {
	src := "board uno\nled = pin 13 as output\nat start:\n    led.on()\n"
	got := RenameEdits(src, "main.muj", 3, 4, "lamp")
	var out map[string]any
	if err := json.Unmarshal([]byte(got), &out); err != nil {
		t.Fatalf("RenameEdits should return JSON: %v (%s)", err, got)
	}
	changes, _ := out["changes"].(map[string]any)
	if len(changes["main.muj"].([]any)) == 0 {
		t.Errorf("renaming 'led' should produce edits, got %v", out)
	}
}
