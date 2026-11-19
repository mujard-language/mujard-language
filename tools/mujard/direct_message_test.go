package main

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

// A failed direct build with no compiler message says what the log says went wrong, not just "The build failed." - Studio
// shows this message in its notice (an esptool that could not unpack itself was the case that showed nothing useful).
func TestDirectFailureNamesTheProblem(t *testing.T) {
	log := "=== bootloader\nC:\\tc\\esptool.exe --chip esp32 elf2image -o out.bin boot.elf\n" +
		"[PYI-1108:ERROR] Failed to extract ucrtbase.dll: failed to open target file!\nfopen: Permission denied\n"
	if got := lastProblemLine(log); got != "fopen: Permission denied" {
		t.Fatalf("lastProblemLine = %q", got)
	}
	if got := lastProblemLine("=== link\nC:\\gcc\\g++.exe -Werror=return-type -o a.elf\n"); got != "" {
		t.Fatalf("a command line is not a problem: %q", got)
	}

	var buf bytes.Buffer
	res := &directResult{Log: log}
	reportDirectJSON(res, "esp32", "blink", filepath.Join(t.TempDir(), "x.bin"), &buf)
	var env struct {
		Error struct{ Code, Message string } `json:"error"`
	}
	if err := json.Unmarshal(buf.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if env.Error.Code != "build_failed" || !strings.Contains(env.Error.Message, "Permission denied") {
		t.Fatalf("message %+v", env.Error)
	}

	// With compiler messages, those are the answer and the message stays short.
	buf.Reset()
	res = &directResult{Log: "sketch.ino:3:1: error: 'x' was not declared in this scope\n"}
	reportDirectJSON(res, "esp32", "blink", filepath.Join(t.TempDir(), "x.bin"), &buf)
	if err := json.Unmarshal(buf.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if env.Error.Message != "The build failed." {
		t.Fatalf("with diagnostics the message is %q", env.Error.Message)
	}
}
