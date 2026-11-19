package main

// Phase 5: the .muj verbs - check, build, run, fmt, new - on top of the lang package.
// The fake developer API is the same one the .ino tests use; nothing here touches hardware.
import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"mujard/lang"
)

const okMujard = "board nano\n\nled = pin LED as output\n\nevery 500 ms:\n    led.toggle()\n"
const badMujard = "board nano\n\nled = pin 99 as output\n" // pin 99 is not a Nano pin

func TestCheckMujardFile(t *testing.T) {
	isolate(t)
	dir := t.TempDir()
	ok := writeFile(t, dir, "blink.muj", okMujard)
	bad := writeFile(t, dir, "bad.muj", badMujard)

	var buf bytes.Buffer
	if code := runCheck([]string{ok}, &buf); code != 0 || !strings.Contains(buf.String(), "checks clean") {
		t.Fatalf("a good program checks clean (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runCheck([]string{bad}, &buf); code != 1 || !strings.Contains(buf.String(), "B4") {
		t.Errorf("a bad program fails with its diagnostic code (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runCheck([]string{bad, "--format", "json"}, &buf); code != 1 || !strings.Contains(buf.String(), `"ok":false`) || !strings.Contains(buf.String(), `"severity":"error"`) {
		t.Errorf("json check reports the diagnostics (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runCheck([]string{ok, "--tree"}, &buf); code != 0 || !strings.Contains(buf.String(), "(file nano") {
		t.Errorf("--tree prints the parsed tree (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runCheck([]string{filepath.Join(dir, "missing.muj")}, &buf); code != 2 {
		t.Errorf("a missing file is a usage error (exit %d):\n%s", code, buf.String())
	}
}

func TestMujardFmt(t *testing.T) {
	isolate(t)
	dir := t.TempDir()
	messy := writeFile(t, dir, "messy.muj", "board   nano\nled=pin LED as output# the led\nevery 500ms:\n\tled.toggle( )\n")

	var buf bytes.Buffer
	if code := runMujardFmt([]string{messy}, &buf); code != 0 {
		t.Fatalf("fmt exit %d:\n%s", code, buf.String())
	}
	b, _ := os.ReadFile(messy)
	want := "board nano\nled = pin LED as output  # the led\nevery 500 ms:\n    led.toggle()\n"
	if string(b) != want {
		t.Errorf("fmt must rewrite to the standard layout:\n%s", b)
	}
	buf.Reset()
	if code := runMujardFmt([]string{messy, "--check"}, &buf); code != 0 {
		t.Errorf("--check on a formatted file exits 0, got %d: %s", code, buf.String())
	}
	other := writeFile(t, dir, "other.muj", "board   nano\n")
	buf.Reset()
	if code := runMujardFmt([]string{other, "--check"}, &buf); code != 1 || !strings.Contains(buf.String(), "not in the standard layout") {
		t.Errorf("--check flags an unformatted file without writing it (exit %d): %s", code, buf.String())
	}
	if b, _ := os.ReadFile(other); string(b) != "board   nano\n" {
		t.Error("--check must not write")
	}
	buf.Reset()
	if code := runMujardFmt([]string{other, "--print"}, &buf); code != 0 || buf.String() != "board nano\n" {
		t.Errorf("--print writes the formatted source to stdout: %d %q", code, buf.String())
	}
	buf.Reset()
	broken := writeFile(t, dir, "broken.muj", "board nano\nx = \"unclosed\n")
	if code := runMujardFmt([]string{broken}, &buf); code != 1 || !strings.Contains(buf.String(), "L103") {
		t.Errorf("a broken file is refused, not guessed (exit %d):\n%s", code, buf.String())
	}
	if code := runMujardFmt([]string{}, &buf); code != 2 {
		t.Errorf("no file is a usage error (exit %d)", code)
	}
}

func TestMujardNew(t *testing.T) {
	isolate(t)
	dir := t.TempDir()
	old, _ := os.Getwd()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(old) })

	var buf bytes.Buffer
	if code := runMujardNew([]string{"blink"}, &buf); code != 0 {
		t.Fatalf("new exit %d:\n%s", code, buf.String())
	}
	b, err := os.ReadFile(filepath.Join(dir, "blink", "blink.muj"))
	if err != nil {
		t.Fatalf("the project file must exist: %v", err)
	}
	if _, diags := lang.CheckSource(string(b)); lang.FirstError(diags) != nil {
		t.Fatalf("the template must check clean: %v", diags)
	}
	buf.Reset()
	if code := runMujardNew([]string{"blink"}, &buf); code != 1 || !strings.Contains(buf.String(), "already exists") {
		t.Errorf("an existing folder is refused (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runMujardNew([]string{"9bad"}, &buf); code != 2 {
		t.Errorf("a name that can't be a file is a usage error (exit %d)", code)
	}
	buf.Reset()
	if code := runMujardNew([]string{"x", "--board", "nonexistent"}, &buf); code != 2 {
		t.Errorf("an unknown board is a usage error (exit %d)", code)
	}
}

// noLocalTools empties PATH and points the desktop/arduino finders at nothing, so a test can never
// reach a real MujardServer.exe or arduino-cli installed on the machine it happens to run on.
func noLocalTools(t *testing.T) {
	t.Helper()
	empty := t.TempDir()
	t.Setenv("PATH", empty)
	t.Setenv("MUJARD_ROOT", empty)
	t.Setenv("LOCALAPPDATA", empty)
	t.Setenv("MUJARD_DESKTOP", filepath.Join(empty, "MujardServer.exe"))
	t.Setenv("MUJARD_ARDUINO_CLI", filepath.Join(empty, "arduino-cli.exe"))
	t.Setenv("MUJARD_ARDUINO_CONFIG", "")
}

func TestBuildMujardOnTheServer(t *testing.T) {
	isolate(t)
	noLocalTools(t) // force the server path: no arduino-cli to run
	key := "mjd_test_abc"
	srv := fakeDevAPI(t, key)
	dir := t.TempDir()
	src := writeFile(t, dir, "blink.muj", okMujard)
	out := filepath.Join(dir, "out.hex")

	var buf bytes.Buffer
	code := runBuild([]string{src, "--board", "nano", "--api", srv.URL, "--key", key, "--out", out}, &buf)
	if code != 0 {
		t.Fatalf("mujard build file.muj on the server (exit %d):\n%s", code, buf.String())
	}
	if b, err := os.ReadFile(out); err != nil || !strings.HasPrefix(string(b), ":100000") {
		t.Fatalf("the artifact must be the downloaded hex: %v", err)
	}
}

func TestBuildMujardRefusesWrongBoard(t *testing.T) {
	isolate(t)
	dir := t.TempDir()
	src := writeFile(t, dir, "blink.muj", okMujard)
	var buf bytes.Buffer
	if code := runBuild([]string{src, "--board", "esp32"}, &buf); code != 2 || !strings.Contains(buf.String(), "board") {
		t.Errorf("--board must match the program's board line (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runBuild([]string{writeFile(t, dir, "bad.muj", badMujard)}, &buf); code != 1 || !strings.Contains(buf.String(), "B4") {
		t.Errorf("a program that doesn't compile never reaches a build (exit %d):\n%s", code, buf.String())
	}
}

func TestBuildMujardShowCpp(t *testing.T) {
	isolate(t)
	noLocalTools(t)
	key := "mjd_test_abc"
	srv := fakeDevAPI(t, key)
	dir := t.TempDir()
	src := writeFile(t, dir, "blink.muj", okMujard)
	var buf bytes.Buffer
	code := runBuild([]string{src, "--board", "nano", "--show-cpp", "--api", srv.URL, "--key", key,
		"--out", filepath.Join(dir, "o.hex")}, &buf)
	if code != 0 {
		t.Fatalf("exit %d:\n%s", code, buf.String())
	}
	for _, name := range []string{"sketch.ino", "mujard_program.cpp", "mujard_rt.h", "mujard_rt.cpp"} {
		if _, err := os.Stat(filepath.Join(dir, "blink.gen", name)); err != nil {
			t.Errorf("--show-cpp must write %s: %v", name, err)
		}
	}
}

func TestRunMujardDelegatesToTheDesktop(t *testing.T) {
	isolate(t)
	noLocalTools(t)
	for _, name := range desktopExeNames {
		if _, err := os.Stat(`C:\Program Files\Mujard\` + name); err == nil {
			t.Skip("Mujard is installed on this machine; it would really be run")
		}
	}
	dir := t.TempDir()
	src := writeFile(t, dir, "blink.muj", okMujard)
	var buf bytes.Buffer
	// no desktop app here: the .muj must be compiled and handed over - the refusal proves it got that far
	if code := runDeviceRun([]string{src, "--port", "COM5"}, &buf); code != 1 ||
		!strings.Contains(buf.String(), "Compiled blink.muj") || !strings.Contains(buf.String(), "desktop app") {
		t.Errorf("run file.muj compiles then delegates (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runDeviceRun([]string{writeFile(t, dir, "bad.muj", badMujard), "--port", "COM5"}, &buf); code != 1 ||
		strings.Contains(buf.String(), "desktop app") {
		t.Errorf("a program that doesn't compile never reaches the desktop (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runDeviceRun([]string{"blink.ino"}, &buf); code != 1 || strings.Contains(buf.String(), "Compiled") ||
		!strings.Contains(buf.String(), "desktop app") {
		t.Errorf("a .ino goes straight through unchanged (exit %d):\n%s", code, buf.String())
	}
}

func TestPhase5CommandsAreUserTierAndListed(t *testing.T) {
	for _, n := range []string{"fmt", "new"} {
		c := findBuiltin(n)
		if c == nil {
			t.Fatalf("%s must be built in", n)
		}
		if operatorOnly[n] {
			t.Errorf("%s is for users", n)
		}
		if !selfHelp[n] {
			t.Errorf("%s should answer --help itself", n)
		}
	}
	var out bytes.Buffer
	t.Setenv("MUJARD_ADMIN", "")
	printHelp(&out, "")
	for _, want := range []string{"fmt ", "new "} {
		if !strings.Contains(out.String(), "\n  "+want) {
			t.Errorf("help lacks %q:\n%s", want, out.String())
		}
	}
}
