package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

// TestBoardCaps: 'mujard boards caps' prints the compiler's own capability table, locally and without a key.
func TestBoardCaps(t *testing.T) {
	t.Setenv("MUJARD_API_KEY", "")
	var out bytes.Buffer
	if code := runBoards([]string{"caps", "nano"}, &out); code != 0 {
		t.Fatalf("exit %d: %s", code, out.String())
	}
	s := out.String()
	for _, want := range []string{"Arduino Nano (nano)", "output", "hardware", "seen by the owner", "wifi", "none", "Not in Mujard yet: DMA"} {
		if !strings.Contains(s, want) {
			t.Errorf("missing %q in:\n%s", want, s)
		}
	}
	out.Reset()
	if code := runBoards([]string{"caps", "esp32", "--format", "json"}, &out); code != 0 {
		t.Fatalf("json: exit %d: %s", code, out.String())
	}
	var doc struct {
		Board        string `json:"board"`
		Capabilities []struct {
			Capability, Level string
		} `json:"capabilities"`
	}
	if err := json.Unmarshal(out.Bytes(), &doc); err != nil || doc.Board != "esp32" || len(doc.Capabilities) == 0 {
		t.Fatalf("json: %v %s", err, out.String())
	}
	for _, c := range doc.Capabilities {
		if c.Capability == "wifi" && c.Level == "none" {
			t.Error("an ESP32 has Wi-Fi")
		}
	}
	out.Reset()
	if code := runBoards([]string{"caps", "toaster"}, &out); code != 2 || !strings.Contains(out.String(), "doesn't know") {
		t.Errorf("unknown board: %d %s", code, out.String())
	}
}
