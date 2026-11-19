package main

// Phase 3 tests: "mujard simulate" and "mujard ai --mode/--file/--model" hand over to the desktop app, which asks
// before it spends. Here a stand-in desktop (see fakeDesktop) proves what mujard does and does not pass on:
// the developer key never goes to the desktop, an old desktop is explained before anything is sent, and
// plain "mujard ai <prompt>" still goes to the developer API as it always did.
import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestIsDesktopAIRequest(t *testing.T) {
	yes := [][]string{{"--mode", "fix"}, {"--mode=fix", "x"}, {"hi", "--file", "a.ino"}, {"--model=a:b", "hi"}}
	for _, a := range yes {
		if !isDesktopAIRequest(a) {
			t.Errorf("%q should be a desktop AI request", a)
		}
	}
	no := [][]string{{}, {"why is my LED off"}, {"what does --mode do in a sketch"}, {"--", "--mode"}, {"model", "list"}, {"--format", "json", "hi"}}
	for _, a := range no {
		if isDesktopAIRequest(a) {
			t.Errorf("%q is a plain prompt or another command", a)
		}
	}
}

func TestWithoutFlagsDropsTheKeyAndAddress(t *testing.T) {
	got := withoutFlags([]string{"--key", "mjd_live_SECRET", "--mode", "fix", "--api=http://x", "--board", "uno", "do it"}, "api", "key")
	if !reflect.DeepEqual(got, []string{"--mode", "fix", "--board", "uno", "do it"}) {
		t.Errorf("got %q", got)
	}
	for _, a := range got {
		if strings.Contains(a, "SECRET") {
			t.Error("the developer key must never be passed on")
		}
	}
}

func TestAIWithModeGoesToTheDesktopWithoutTheKey(t *testing.T) {
	set, args := fakeDesktop(t)
	set("--help", "  MujardServer.exe ai --mode <m> [--file <sketch|folder>] --board <id> \"<prompt>\"\r\n", 0)
	set("ai", "ok\r\n", 0)
	t.Setenv("MUJARD_CONFIG", filepath.Join(t.TempDir(), "none.json"))
	var out bytes.Buffer
	code := runAI([]string{"--key", "mjd_live_SECRET", "--api", "http://127.0.0.1:1", "--mode", "explain", "--file", "blink.ino", "--board", "uno", "why", "--yes"}, &out)
	if code != 0 {
		t.Fatalf("exit %d: %s", code, out.String())
	}
	if got := args(); got != "ai --mode explain --file blink.ino --board uno why --yes" {
		t.Errorf("the desktop gets the mode flags and the prompt, and never the key or address: %q", got)
	}
}

func TestPlainAIPromptStillUsesTheDeveloperAPI(t *testing.T) {
	_, args := fakeDesktop(t)
	t.Setenv("MUJARD_CONFIG", filepath.Join(t.TempDir(), "none.json"))
	t.Setenv("MUJARD_KEY", "")
	var out bytes.Buffer
	if code := runAI([]string{"why is my LED off"}, &out); code != 2 || !strings.Contains(out.String(), "needs a developer key") {
		t.Errorf("a plain prompt without a key is the usual refusal: %d %s", code, out.String())
	}
	if args() != "" {
		t.Errorf("a plain prompt must never reach the desktop app: %q", args())
	}
}

func TestOldDesktopIsExplainedBeforeAnythingIsSent(t *testing.T) {
	set, args := fakeDesktop(t)
	set("--help", "MujardServer.exe list  the boards detected right now\r\n", 0) // an app from before Phase 3
	set("ai", "this must not run\r\n", 0)
	set("simulate", "this must not run\r\n", 0)
	var out bytes.Buffer
	if code := runAI([]string{"--mode", "chat", "--board", "uno", "hi"}, &out); code != 2 || !strings.Contains(out.String(), "needs a newer Mujard desktop app") {
		t.Errorf("ai: %d %s", code, out.String())
	}
	out.Reset()
	if code := runSimulate([]string{"blink.ino", "--board", "uno"}, &out); code != 2 || !strings.Contains(out.String(), "needs a newer") {
		t.Errorf("simulate: %d %s", code, out.String())
	}
	if args() != "--help" {
		t.Errorf("only the harmless --help was ever sent to the old app: %q", args())
	}
	out.Reset()
	if code := runSimulate([]string{"blink.ino", "--board", "uno", "--format", "json"}, &out); code != 2 {
		t.Fatalf("exit %d", code)
	}
	var env map[string]any
	if err := json.Unmarshal(out.Bytes(), &env); err != nil || env["ok"] != false || env["error"].(map[string]any)["code"] != "desktop_too_old" {
		t.Errorf("JSON callers get an envelope: %v %s", err, out.String())
	}
}

func TestSimulateForwardsItsArgumentsAndNeedsASketch(t *testing.T) {
	set, args := fakeDesktop(t)
	set("--help", "  MujardServer.exe simulate <sketch|folder> --board uno\r\n", 0)
	set("simulate", "serial output\r\n", 0)
	var out bytes.Buffer
	if code := runSimulate(nil, &out); code != 2 || !strings.Contains(out.String(), "Usage: mujard simulate") {
		t.Errorf("bare simulate: %d %s", code, out.String())
	}
	if args() != "" {
		t.Errorf("a usage error must not start the desktop app: %q", args())
	}
	out.Reset()
	if code := runSimulate([]string{"blink.ino", "--board", "uno", "--yes"}, &out); code != 0 {
		t.Errorf("exit %d", code)
	}
	if args() != "simulate blink.ino --board uno --yes" {
		t.Errorf("args %q", args())
	}
}

func TestPhase3HelpSaysWhoPays(t *testing.T) {
	var out bytes.Buffer
	printHelp(&out, "ai")
	for _, want := range []string{"developer key's coins", "THIS DESKTOP's", "Your files are never changed"} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("help ai lacks %q:\n%s", want, out.String())
		}
	}
	out.Reset()
	printHelp(&out, "simulate")
	if !strings.Contains(out.String(), "counts as a build") {
		t.Errorf("help simulate: %s", out.String())
	}
	if findBuiltin("simulate") == nil || operatorOnly["simulate"] {
		t.Error("simulate is a user command")
	}
	_ = os.Stdout
}
