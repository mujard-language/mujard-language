package main

// Phase 2 tests: the desktop-delegated read-only verbs. A tiny batch file stands in for MujardServer.exe
// (it prints canned output, records its arguments and exits with a canned code), so no desktop app, board
// or server is involved. Windows only, like the desktop itself.
import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// fakeDesktop returns a function that sets what the stand-in prints and exits with for one verb.
func fakeDesktop(t *testing.T) (setVerb func(verb, out string, code int), lastArgs func() string) {
	t.Helper()
	if runtime.GOOS != "windows" {
		t.Skip("the stand-in desktop is a .bat file")
	}
	dir := t.TempDir()
	bat := filepath.Join(dir, "MujardServer.bat")
	script := "@echo off\r\necho %* > \"%~dp0args.txt\"\r\nif exist \"%~dp0out_%1.txt\" type \"%~dp0out_%1.txt\"\r\nset /p C=<\"%~dp0code_%1.txt\"\r\nexit /b %C%\r\n"
	if err := os.WriteFile(bat, []byte(script), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("MUJARD_DESKTOP", bat)
	setVerb = func(verb, out string, code int) {
		if err := os.WriteFile(filepath.Join(dir, "out_"+verb+".txt"), []byte(out), 0o644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "code_"+verb+".txt"), []byte(itoa(code)), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	lastArgs = func() string {
		b, _ := os.ReadFile(filepath.Join(dir, "args.txt"))
		return strings.TrimSpace(string(b))
	}
	return
}

func itoa(n int) string { return strings.TrimSpace(strings.Repeat(" ", 0) + string(rune('0'+n%10))) }

func TestUnknownDesktopVerbDetection(t *testing.T) {
	if !isUnknownDesktopVerb("Unknown command \"libraries\".\r\n\r\nMujard 1.0.0", "libraries") {
		t.Error("the text form")
	}
	if !isUnknownDesktopVerb(`{"schema":"mujard-cli.v1","command":"libraries","ok":false,"error":{"code":"usage_error","message":"Unknown command \"libraries\"."}}`, "libraries") {
		t.Error("the JSON form")
	}
	if isUnknownDesktopVerb("Unknown command \"flash\".", "libraries") || isUnknownDesktopVerb("libraries needs search", "libraries") {
		t.Error("only that exact complaint counts")
	}
}

func TestFormatArgsAndDesktopBoardsRouting(t *testing.T) {
	if got := formatArgs([]string{"--key", "k", "--format", "json", "--api", "x"}); !reflect.DeepEqual(got, []string{"--format", "json"}) {
		t.Errorf("only --format goes to the desktop: %q", got)
	}
	if got := formatArgs([]string{"--format=json"}); !reflect.DeepEqual(got, []string{"--format=json"}) {
		t.Errorf("%q", got)
	}
	if formatArgs([]string{"--key", "k"}) != nil {
		t.Error("nothing to forward")
	}
	if !wantsJSONFormat([]string{"--format", "json"}) || !wantsJSONFormat([]string{"--format=json"}) || wantsJSONFormat([]string{"--format", "text"}) {
		t.Error("wantsJSONFormat")
	}
}

func TestLibrariesNeedsAnArgumentBeforeAnythingRuns(t *testing.T) {
	t.Setenv("MUJARD_DESKTOP", filepath.Join(t.TempDir(), "nothing.exe"))
	var out bytes.Buffer
	if code := runLibraries(nil, &out); code != 2 || !strings.Contains(out.String(), "mujard libraries search") {
		t.Errorf("bare libraries: %d %s", code, out.String())
	}
	out.Reset()
	if code := runLibraries([]string{"--format", "json"}, &out); code != 2 {
		t.Errorf("flags alone are not a sub-command: %d", code)
	}
}

func TestDesktopVerbsPassThroughAndOldDesktopsAreExplained(t *testing.T) {
	set, args := fakeDesktop(t)
	set("libraries", "name  version\r\nServo 1.2\r\n", 0)
	var out bytes.Buffer
	if code := runLibraries([]string{"search", "servo"}, &out); code != 0 || !strings.Contains(out.String(), "Servo 1.2") {
		t.Errorf("pass-through: %d %q", code, out.String())
	}
	if args() != "libraries search servo" {
		t.Errorf("the arguments reach the desktop unchanged: %q", args())
	}
	set("libraries", "boom\r\n", 1)
	out.Reset()
	if code := runLibraries([]string{"installed"}, &out); code != 1 || !strings.Contains(out.String(), "boom") {
		t.Errorf("the desktop's exit code is kept: %d %q", code, out.String())
	}
	set("libraries", "Unknown command \"libraries\".\r\n\r\nMujard 1.0.0 - usage...\r\n", 2)
	out.Reset()
	if code := runLibraries([]string{"installed"}, &out); code != 2 || !strings.Contains(out.String(), "needs a newer Mujard desktop app") || strings.Contains(out.String(), "usage...") {
		t.Errorf("an old desktop gets one sentence, not its usage dump: %d %q", code, out.String())
	}
	out.Reset()
	if code := runLibraries([]string{"installed", "--format", "json"}, &out); code != 2 {
		t.Errorf("exit %d", code)
	}
	var env map[string]any
	if err := json.Unmarshal(out.Bytes(), &env); err != nil || env["ok"] != false || env["schema"] != "mujard-cli.v1" {
		t.Errorf("with --format json the explanation is an envelope: %v %q", err, out.String())
	}
}

func TestUsageWindowsForwardsOnlyTheFormat(t *testing.T) {
	set, args := fakeDesktop(t)
	set("usage", "Allowance for THIS DESKTOP\r\n", 0)
	var out bytes.Buffer
	if code := runUsage([]string{"windows", "--key", "mjd_test_secret", "--api", "http://x", "--format", "json"}, &out); code != 0 || !strings.Contains(out.String(), "THIS DESKTOP") {
		t.Fatalf("%d %q", code, out.String())
	}
	if got := args(); got != "usage windows --format json" {
		t.Errorf("the developer key and API address must never reach the desktop process: %q", got)
	}
	if strings.Contains(args(), "mjd_test_secret") {
		t.Error("key leaked into the desktop's arguments")
	}
}

func TestBoardsDesktopRouting(t *testing.T) {
	set, args := fakeDesktop(t)
	set("board", "uno details\r\n", 0)
	set("boards", "all boards\r\n", 0)
	t.Setenv("MUJARD_CONFIG", filepath.Join(t.TempDir(), "none.json"))
	t.Setenv("MUJARD_KEY", "")
	var out bytes.Buffer
	// no developer key is needed: this is the desktop's own catalog
	// (the pass-through verbs write to the process's own stdout, as they always have; the arguments are what matter)
	if code := runBoardInfo([]string{"info", "uno", "--desktop", "--format", "json"}, &out); code != 0 {
		t.Errorf("board info --desktop: %d %q", code, out.String())
	}
	if args() != "board uno --format json" {
		t.Errorf("args %q", args())
	}
	out.Reset()
	if code := runBoards([]string{"--desktop"}, &out); code != 0 || args() != "boards" {
		t.Errorf("boards --desktop: %d %q args %q", code, out.String(), args())
	}
	out.Reset()
	if code := runBoards([]string{"--desktop", "a", "b"}, &out); code != 2 {
		t.Errorf("too many words: %d", code)
	}
	// without --desktop nothing changed: a missing key is still the usual refusal
	out.Reset()
	if code := runBoardInfo([]string{"info", "uno"}, &out); code != 2 || !strings.Contains(out.String(), "needs a developer key") {
		t.Errorf("plain board info still wants a key: %d %q", code, out.String())
	}
}

func TestDevicesInfoAndAdaptersInfo(t *testing.T) {
	set, _ := fakeDesktop(t)
	set("list", `{"schema":"mujard-cli.v1","command":"list","ok":true,"data":{"boards":[{"port":"COM5","name":"Arduino Uno","chip":"ATmega16U2","boardId":"uno","vid":"2341","pid":"0043","kind":"serial"},{"port":"COM7","name":"CH340 adapter","chip":"CH340","boardId":null,"vid":"1A86","pid":"7523","kind":"serial"}]},"error":null}`+"\r\n", 0)
	set("adapters", `{"schema":"mujard-cli.v1","command":"adapters","ok":true,"data":{"adapters":[{"id":"serial-flash","version":"1.2.0","protocols":["stk500v1","uf2"],"boards":["uno","nano","pico"],"runsOn":"desktop"}],"phoneOnly":["jtag2updi"]},"error":null}`+"\r\n", 0)

	var out bytes.Buffer
	if code := runDeviceList([]string{"info", "com5"}, &out); code != 0 || !strings.Contains(out.String(), "ATmega16U2") || !strings.Contains(out.String(), "--board uno --port COM5") {
		t.Errorf("devices info (port match ignores case): %d\n%s", code, out.String())
	}
	out.Reset()
	if code := runDeviceList([]string{"info", "COM7"}, &out); code != 0 || !strings.Contains(out.String(), "pass --board yourself") {
		t.Errorf("an adapter-only board says so: %d\n%s", code, out.String())
	}
	out.Reset()
	if code := runDeviceList([]string{"info", "COM9"}, &out); code != 1 || !strings.Contains(out.String(), "Nothing is plugged into COM9") {
		t.Errorf("unknown port: %d\n%s", code, out.String())
	}
	out.Reset()
	if code := runDeviceList([]string{"info", "COM5", "--format", "json"}, &out); code != 0 {
		t.Fatalf("json: %d", code)
	}
	var env map[string]any
	if err := json.Unmarshal(out.Bytes(), &env); err != nil || env["command"] != "devices info" || env["ok"] != true || env["data"].(map[string]any)["port"] != "COM5" {
		t.Errorf("devices info --format json: %v %s", err, out.String())
	}
	out.Reset()
	if code := runDeviceList([]string{"info", "COM9", "--format", "json"}, &out); code != 1 || !strings.Contains(out.String(), `"code":"not_found"`) {
		t.Errorf("json not found: %d %s", code, out.String())
	}
	for _, bad := range [][]string{{"info"}, {"info", "not a port"}, {"info", "COM5", "COM6"}, {"info", "COM5", "--format", "yaml"}} {
		out.Reset()
		if code := runDeviceList(bad, &out); code != 2 {
			t.Errorf("devices %v must be a usage error, got %d", bad, code)
		}
	}

	out.Reset()
	if code := runDeviceAdapters([]string{"info", "serial-flash"}, &out); code != 0 || !strings.Contains(out.String(), "stk500v1, uf2") || !strings.Contains(out.String(), "boards     3: uno, nano, pico") {
		t.Errorf("adapters info: %d\n%s", code, out.String())
	}
	out.Reset()
	if code := runDeviceAdapters([]string{"info", "jtag2updi"}, &out); code != 1 || !strings.Contains(out.String(), "phone app") {
		t.Errorf("phone-only adapter: %d\n%s", code, out.String())
	}
	out.Reset()
	if code := runDeviceAdapters([]string{"info", "nope", "--format", "json"}, &out); code != 1 || !strings.Contains(out.String(), "not_found") {
		t.Errorf("unknown adapter: %d %s", code, out.String())
	}
	// the old spellings are untouched
	set("list", "No board detected\r\n", 0)
	out.Reset()
	if code := runDeviceList([]string{"list"}, &out); code != 0 {
		t.Errorf("devices list: %d", code)
	}
}

func TestPhase2CommandsAreListedAndSelfDescribing(t *testing.T) {
	c := findBuiltin("libraries")
	if c == nil || operatorOnly["libraries"] || !selfHelp["libraries"] {
		t.Fatal("libraries is a user command that answers --help itself")
	}
	var out bytes.Buffer
	printHelp(&out, "libraries")
	if !strings.Contains(out.String(), "Installing a library is not offered") {
		t.Errorf("help should say what is deliberately missing:\n%s", out.String())
	}
	out.Reset()
	printHelp(&out, "usage")
	if !strings.Contains(out.String(), "usage windows") {
		t.Errorf("help usage: %s", out.String())
	}
}
