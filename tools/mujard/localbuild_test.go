package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"
	"testing"
)

// countingServer answers nothing useful and counts every request: an offline build must make none.
func countingServer(t *testing.T) (*httptest.Server, *int32) {
	t.Helper()
	var n int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&n, 1)
		w.WriteHeader(500)
	}))
	t.Cleanup(srv.Close)
	return srv, &n
}

func envelope(t *testing.T, out string) map[string]any {
	t.Helper()
	var doc map[string]any
	line := strings.TrimSpace(out)
	if i := strings.LastIndex(line, "\n"); i >= 0 {
		line = line[i+1:]
	}
	if err := json.Unmarshal([]byte(line), &doc); err != nil {
		t.Fatalf("not one JSON envelope: %v\n%s", err, out)
	}
	return doc
}

// noDirectBackend keeps mujard-build out of reach too (noLocalTools already hides arduino-cli and the repo).
func noDirectBackend(t *testing.T) {
	t.Helper()
	t.Setenv("MUJARD_BUILD", filepath.Join(t.TempDir(), "mujard-build.exe"))
}

func TestOfflineBuildNeverCallsTheServer(t *testing.T) {
	isolate(t)
	noLocalTools(t)
	noDirectBackend(t)
	srv, calls := countingServer(t)
	dir := t.TempDir()
	writeFile(t, dir, "blink.ino", "void setup(){}\nvoid loop(){}\n")
	src := writeFile(t, dir, "lamp.muj", okMujard)

	for _, args := range [][]string{
		{"uno", dir, "--offline", "--format", "json", "--api", srv.URL, "--key", "mjd_test_x"},
		{src, "--offline", "--format", "json", "--api", srv.URL, "--key", "mjd_test_x"},
	} {
		var buf bytes.Buffer
		code := runBuild(args, &buf)
		doc := envelope(t, buf.String())
		e, _ := doc["error"].(map[string]any)
		if code != 1 || doc["ok"] != false || e == nil || e["code"] != "no_local_toolchain" {
			t.Errorf("%v: with nothing on this PC to build with, --offline says so (exit %d):\n%s", args, code, buf.String())
		}
		if msg, _ := e["message"].(string); !strings.Contains(msg, "offline build pack") {
			t.Errorf("the message says what to do: %q", msg)
		}
	}
	if n := atomic.LoadInt32(calls); n != 0 {
		t.Fatalf("an offline build sent %d request(s) to the server", n)
	}
}

func TestOfflineMujardCompileErrorsAreJSON(t *testing.T) {
	isolate(t)
	noLocalTools(t)
	noDirectBackend(t)
	dir := t.TempDir()
	src := writeFile(t, dir, "bad.muj", badMujard)
	var buf bytes.Buffer
	code := runBuild([]string{src, "--offline", "--format", "json"}, &buf)
	doc := envelope(t, buf.String())
	e, _ := doc["error"].(map[string]any)
	data, _ := doc["data"].(map[string]any)
	errs, _ := data["errors"].([]any)
	if code != 1 || e == nil || e["code"] != "compile_failed" || len(errs) == 0 {
		t.Fatalf("a program that doesn't compile is one JSON failure with its messages (exit %d):\n%s", code, buf.String())
	}
	var first map[string]any
	for _, x := range errs {
		if m, _ := x.(map[string]any); m != nil && m["severity"] == "error" {
			first = m
			break
		}
	}
	if first == nil || first["file"] != "bad.muj" || first["line"] != float64(3) || first["code"] == "" {
		t.Errorf("the error points at the .muj line, with its code: %v", errs)
	}
	if strings.Count(strings.TrimSpace(buf.String()), "\n") != 0 {
		t.Errorf("nothing but the envelope is printed in JSON mode:\n%s", buf.String())
	}
}

func TestGccDiags(t *testing.T) {
	log := strings.Join([]string{
		`C:\Users\me\AppData\Local\Temp\mujard-build1\sketch\sketch.ino:12:5: error: 'ledd' was not declared in this scope`,
		`C:\Users\me\AppData\Local\Temp\mujard-build1\sketch\sketch.ino:12:5: error: 'ledd' was not declared in this scope`, // repeated by make
		`/tmp/x/sketch/motors.muj:4: warning: unused variable 'speed'`,
		`In file included from C:\x\sketch\sketch.ino:1:0:`,
		`C:\x\sketch\helper.h:2:10: fatal error: Missing.h: No such file or directory`,
		`compilation terminated.`,
	}, "\r\n")
	d := gccDiags(log)
	if len(d) != 3 {
		t.Fatalf("three distinct messages, got %d: %+v", len(d), d)
	}
	if d[0].File != "sketch.ino" || d[0].Line != 12 || d[0].Column != 5 || d[0].Severity != "error" || !strings.Contains(d[0].Message, "ledd") {
		t.Errorf("a Windows path keeps only the file name: %+v", d[0])
	}
	if d[1].File != "motors.muj" || d[1].Line != 4 || d[1].Column != 0 || d[1].Severity != "warning" {
		t.Errorf("a message without a column, on a .muj line: %+v", d[1])
	}
	if d[2].File != "helper.h" || d[2].Severity != "error" {
		t.Errorf("a fatal error is an error: %+v", d[2])
	}
}

// fakeArduinoCli is a stand-in arduino-cli: "compile --fqbn F --output-dir D [--config-file C] S" leaves a firmware in D.
func fakeArduinoCli(t *testing.T) string {
	t.Helper()
	if runtime.GOOS == "windows" {
		p := filepath.Join(t.TempDir(), "arduino-cli.bat")
		// It also copies the build's secrets header to MUJARD_TEST_SEEN (when set), so a test can see what the compiler got.
		script := "@echo off\r\nmkdir \"%~5\" 2>nul\r\necho :00000001FF> \"%~5\\sketch.ino.hex\"\r\nif defined MUJARD_TEST_SEEN if exist \"%~6\\mujard_secrets.h\" copy /y \"%~6\\mujard_secrets.h\" \"%MUJARD_TEST_SEEN%\" >nul\r\nexit /b 0\r\n"
		if err := os.WriteFile(p, []byte(script), 0o700); err != nil {
			t.Fatal(err)
		}
		return p
	}
	p := filepath.Join(t.TempDir(), "arduino-cli")
	script := "#!/bin/sh\nmkdir -p \"$5\"\necho ':00000001FF' > \"$5/sketch.ino.hex\"\nif [ -n \"$MUJARD_TEST_SEEN\" ] && [ -f \"$6/mujard_secrets.h\" ]; then cp \"$6/mujard_secrets.h\" \"$MUJARD_TEST_SEEN\"; fi\nexit 0\n"
	if err := os.WriteFile(p, []byte(script), 0o700); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestOfflineBuildWithArduinoCli(t *testing.T) {
	isolate(t)
	noLocalTools(t)
	noDirectBackend(t)
	t.Setenv("MUJARD_ARDUINO_CLI", fakeArduinoCli(t))
	srv, calls := countingServer(t)
	dir := t.TempDir()
	src := writeFile(t, dir, "lamp.muj", okMujard)
	out := filepath.Join(dir, "lamp.hex")

	var buf bytes.Buffer
	code := runBuild([]string{src, "--offline", "--format", "json", "--out", out, "--api", srv.URL}, &buf)
	doc := envelope(t, buf.String())
	data, _ := doc["data"].(map[string]any)
	if code != 0 || doc["ok"] != true || data["backend"] != "arduino-cli" || data["format"] != "hex" {
		t.Fatalf("a .muj program builds on this PC with arduino-cli (exit %d):\n%s", code, buf.String())
	}
	art, _ := data["artifact"].(map[string]any)
	if art["path"] != out || art["bytes"] == float64(0) || len(art["sha256"].(string)) != 64 {
		t.Errorf("the artifact is reported with its path, size and hash: %v", art)
	}
	if b, err := os.ReadFile(out); err != nil || !strings.HasPrefix(string(b), ":00000001FF") {
		t.Errorf("the firmware is saved to --out: %v", err)
	}

	// A .ino sketch with a full board name builds the same way; a short id can't (only a .muj knows its FQBN).
	sk := t.TempDir()
	writeFile(t, sk, "blink.ino", "void setup(){}\nvoid loop(){}\n")
	buf.Reset()
	if code := runBuild([]string{"arduino:avr:uno", sk, "--offline", "--format", "json", "--out", filepath.Join(sk, "b.hex")}, &buf); code != 0 || envelope(t, buf.String())["ok"] != true {
		t.Errorf("a .ino with an FQBN builds offline with arduino-cli (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runBuild([]string{"uno", sk, "--offline", "--format", "json"}, &buf); code != 1 || !strings.Contains(buf.String(), "no_local_toolchain") {
		t.Errorf("a short id with only arduino-cli is reported, not guessed (exit %d):\n%s", code, buf.String())
	}
	if n := atomic.LoadInt32(calls); n != 0 {
		t.Fatalf("an offline build sent %d request(s) to the server", n)
	}
}

func TestOfflineSecretsStayInTheBuildCopy(t *testing.T) {
	isolate(t)
	noLocalTools(t)
	noDirectBackend(t)
	t.Setenv("MUJARD_ARDUINO_CLI", fakeArduinoCli(t))
	t.Setenv("MUJARD_SECRET_HOME_WIFI_SSID", "MyNet")
	t.Setenv("MUJARD_SECRET_HOME_WIFI_PASSWORD", "pa\"ss")
	seen := filepath.Join(t.TempDir(), "seen.h")
	t.Setenv("MUJARD_TEST_SEEN", seen)
	dir := t.TempDir()
	prog := "board esp32\n\nat start:\n    wifi.connect(secret(\"home_wifi\"))\n"
	src := writeFile(t, dir, "net.muj", prog)
	var buf bytes.Buffer
	code := runBuild([]string{src, "--offline", "--format", "json", "--out", filepath.Join(dir, "net.bin")}, &buf)
	if code != 0 {
		t.Fatalf("a program with a secret builds offline (exit %d):\n%s", code, buf.String())
	}
	h, err := os.ReadFile(seen)
	if err != nil {
		t.Fatalf("the build got a secrets header: %v", err)
	}
	if !strings.Contains(string(h), `#define HOME_WIFI_SSID "MyNet"`) || !strings.Contains(string(h), `#define HOME_WIFI_PASSWORD "pa\"ss"`) {
		t.Errorf("the header carries the values from MUJARD_SECRET_*, escaped:\n%s", h)
	}
	if strings.Contains(buf.String(), "MyNet") {
		t.Errorf("a secret's value never appears in the output:\n%s", buf.String())
	}
	for _, e := range mustList(t, dir) {
		if strings.Contains(e, "secrets") {
			t.Errorf("no secrets file is left in the project folder: %s", e)
		}
	}
}

func mustList(t *testing.T, dir string) []string {
	t.Helper()
	ents, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, e := range ents {
		names = append(names, e.Name())
	}
	return names
}
