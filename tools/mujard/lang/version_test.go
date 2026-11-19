package lang

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// The language version follows the language: when its sources change, the number must move, so every copy of Mujard can tell
// an older language from a newer one. MUJARD_GOLDEN=1 (update-mujard.ps1) refreshes version.json and bumps the number.
func TestLanguageVersionIsCurrent(t *testing.T) {
	fp, err := languageFingerprint(".")
	if err != nil {
		t.Fatal(err)
	}
	var v versionFile
	if err := json.Unmarshal(versionJSON, &v); err != nil || v.Version < 1 {
		t.Fatalf("version.json is unreadable or has no version: %v", err)
	}
	if v.Fingerprint == fp {
		return
	}
	if os.Getenv("MUJARD_GOLDEN") == "1" {
		next := versionFile{Version: v.Version + 1, Fingerprint: fp}
		if v.Fingerprint == "" { // the first fingerprint: nothing was released under the old one
			next.Version = v.Version
		}
		data, _ := json.MarshalIndent(next, "", "  ")
		if err := os.WriteFile(filepath.Join(".", "version.json"), append(data, '\n'), 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("language version %d -> %d", v.Version, next.Version)
		return
	}
	t.Fatalf("the language changed since version %d was recorded: run .\\update-mujard.ps1 (or MUJARD_GOLDEN=1 go test ./lang -run TestLanguageVersionIsCurrent) to give it a new version number", v.Version)
}

// A comment or formatting change is not a new language; a code change is.
func TestLanguageFingerprintIgnoresComments(t *testing.T) {
	dir := t.TempDir()
	write := func(src string) string {
		t.Helper()
		if err := os.WriteFile(filepath.Join(dir, "a.go"), []byte(src), 0o644); err != nil {
			t.Fatal(err)
		}
		fp, err := languageFingerprint(dir)
		if err != nil {
			t.Fatal(err)
		}
		return fp
	}
	base := write("package x\n\nfunc F() int { return 1 }\n")
	if write("package x\n\n// F returns one.\nfunc F() int {\n\treturn 1 // one\n}\n") != base {
		t.Error("a comment or layout change moved the fingerprint")
	}
	if write("package x\n\nfunc F() int { return 2 }\n") == base {
		t.Error("a code change did not move the fingerprint")
	}
	if LanguageVersion() < 1 {
		t.Error("LanguageVersion is not set")
	}
}

// capabilities.json fingerprints on which capabilities each board has proven, not on the dates or descriptions.
// A matrix run that only re-dates a proven capability must not bump the language version; a capability being added
// or dropped still must - W618 warns on presence, never on the date.
func TestCapStructureIgnoresDates(t *testing.T) {
	v1 := `{"build":{"uno":{"output":"2026-10-30","pwm":"2026-10-30"},"pico":{"output":"2026-10-30"}},"hardware":{"nano":{"output":"2026-10-29: ran blink"}}}`
	v2 := `{"build":{"uno":{"output":"2026-11-01","pwm":"2026-11-01"},"pico":{"output":"2026-11-01"}},"hardware":{"nano":{"output":"2026-11-01: ran blink on the real board, saw it"}}}`
	v3 := `{"build":{"uno":{"output":"2026-11-01"},"pico":{"output":"2026-11-01","servo":"2026-11-01"}},"hardware":{"nano":{"output":"2026-11-01: ran blink"}}}`

	a, err := capStructure([]byte(v1))
	if err != nil {
		t.Fatal(err)
	}
	b, err := capStructure([]byte(v2))
	if err != nil {
		t.Fatal(err)
	}
	c, err := capStructure([]byte(v3))
	if err != nil {
		t.Fatal(err)
	}
	if string(a) != string(b) {
		t.Error("re-dating a proven capability moved the fingerprint - evidence-only changes must not bump the version")
	}
	if string(a) == string(c) {
		t.Error("an added/removed capability did not move the fingerprint - W618-visible changes must bump the version")
	}
}
