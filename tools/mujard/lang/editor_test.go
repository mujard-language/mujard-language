package lang

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// editorInfoPath is docs/lang/editor-info.json: the language description the Android app ships in its assets.
var editorInfoPath = filepath.Join("..", "..", "..", "docs", "lang", "editor-info.json")

func editorInfoJSON(t *testing.T) []byte {
	t.Helper()
	b, err := json.MarshalIndent(NewEditorInfo(), "", "  ")
	if err != nil {
		t.Fatalf("marshal EditorInfo: %v", err)
	}
	return append(b, '\n')
}

// The info blob itself: it has to describe the whole language or the editors complete against a stale picture.
func TestEditorInfoCoversTheLanguage(t *testing.T) {
	info := NewEditorInfo()

	for _, kw := range []string{"board", "every", "task", "pin", "on", "cpp", "use"} {
		found := false
		for _, k := range info.Keywords {
			if k == kw {
				found = true
			}
		}
		if !found {
			t.Errorf("keywords missing %q", kw)
		}
	}
	if len(info.Boards) == 0 {
		t.Fatal("no boards")
	}
	for _, b := range info.Boards {
		if b.Pins == nil || len(b.Valid) == 0 {
			t.Errorf("board %s has no pin table", b.Name)
		}
	}
	has := func(name string) *EditorGlobal {
		for i := range info.Globals {
			if info.Globals[i].Name == name {
				return &info.Globals[i]
			}
		}
		return nil
	}
	if g := has("notify"); g == nil || g.Kind != "function" {
		t.Error("notify should be a function global")
	}
	if g := has("serial"); g == nil || g.Kind != "object" {
		t.Error("serial should be an object global")
	}
	if g := has("mqtt"); g == nil || g.Lib == "" {
		t.Error("mqtt should name its mujard library")
	}
	// A pin declared 'as output' is what editors complete 'led.' against.
	out, ok := info.Members["output"]
	if !ok || len(out) == 0 {
		t.Fatal("no 'output' member table")
	}
	var isEvent bool
	for _, m := range info.Members["input"] {
		if m.Kind == "event" {
			isEvent = true
		}
	}
	if !isEvent {
		t.Error("an input pin should have events for 'on led.'")
	}
}

// TestCompleteMemberChains: 'a.b.' completes against the member's own result type - 't.split.' opens
// the list members (split returns a list), not text's and not every member (v66 chain resolution).
func TestCompleteMemberChains(t *testing.T) {
	svc := NewEditorService(NewEditorInfo())
	labels := func(text string, pos LspPos) map[string]bool {
		out := map[string]bool{}
		for _, it := range svc.Complete(text, pos)["items"].([]map[string]any) {
			out[it["label"].(string)] = true
		}
		return out
	}
	// '    t.split.' - the cursor is right after the dot (char 12).
	l := labels("board uno\nvar t: text = \"a,b\"\nevery 500 ms:\n    t.split.", LspPos{Line: 3, Character: 12})
	if !l["add"] || !l["clear"] || l["split"] || l["toggle"] {
		t.Errorf("'t.split.' should offer the list members only: %v", l)
	}
	// '    t.split.length.' - a two-hop chain: text -> split -> list -> length -> number.
	l = labels("board uno\nvar t: text = \"a,b\"\nevery 500 ms:\n    t.split.length.", LspPos{Line: 3, Character: 19})
	if !l["bit"] || l["add"] || l["toggle"] {
		t.Errorf("'t.split.length.' should offer the number members only: %v", l)
	}
	// An unresolvable hop still falls back to every member, like an unknown receiver always did.
	l = labels("board uno\nevery 500 ms:\n    zzz.foo.", LspPos{Line: 2, Character: 12})
	if !l["toggle"] {
		t.Errorf("an unresolvable chain should fall back to all members: %v", l)
	}
}

// docs/lang/editor-info.json is generated from NewEditorInfo; MUJARD_GOLDEN=1 rewrites it after a language change.
// androidEditorInfoPath is the Android app's copy (its editor's highlighting and completion read it as an asset).
var androidEditorInfoPath = filepath.Join("..", "..", "..", "app", "src", "main", "assets", "editor-info.json")

func TestEditorInfoFileIsCurrent(t *testing.T) {
	want := editorInfoJSON(t)
	paths := []string{editorInfoPath}
	if _, err := os.Stat(filepath.Dir(androidEditorInfoPath)); err == nil {
		paths = append(paths, androidEditorInfoPath) // the app's checkout is only in the monorepo
	}
	if os.Getenv("MUJARD_GOLDEN") == "1" {
		for _, p := range paths {
			if err := os.WriteFile(p, want, 0o644); err != nil {
				t.Fatalf("write %s: %v", p, err)
			}
		}
		return
	}
	for _, p := range paths {
		got, err := os.ReadFile(p)
		if err != nil {
			t.Fatalf("read %s: %v (run with MUJARD_GOLDEN=1 to generate)", p, err)
		}
		if string(got) != string(want) {
			t.Errorf("%s is stale; regenerate with MUJARD_GOLDEN=1 go test ./lang", p)
		}
	}
}
