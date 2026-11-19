package main

// Tests for the unified command line: the API client, the build pipeline, config and plugins.
// The fake developer API checks the Bearer key and the Idempotency-Key the way the real one does.
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestApiV1Base(t *testing.T) {
	for in, want := range map[string]string{
		"http://127.0.0.1/mujard-api":         "http://127.0.0.1/mujard-api/api/v1",
		"https://h.example/mujard-api/":       "https://h.example/mujard-api/api/v1",
		"http://127.0.0.1:8760":               "http://127.0.0.1:8760/api/v1",
		"https://h.example/mujard-api/api/v1": "https://h.example/mujard-api/api/v1",
	} {
		if got := apiV1(in); got != want {
			t.Errorf("apiV1(%q) = %q, want %q", in, got, want)
		}
	}
}

// fakeDevAPI is the /v1 side the real Api.php speaks: Bearer key, idempotent builds, job states.
func fakeDevAPI(t *testing.T, key string) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header.Get("Authorization") != "Bearer "+key {
			w.WriteHeader(401)
			_, _ = w.Write([]byte(`{"error":{"code":"invalid_api_key","message":"bad key","request_id":"R1"}}`))
			return
		}
		switch {
		case r.Method == "POST" && r.URL.Path == "/api/v1/builds":
			if r.Header.Get("Idempotency-Key") == "" {
				w.WriteHeader(400)
				_, _ = w.Write([]byte(`{"error":{"code":"idempotency_key_required","message":"need it","request_id":"R2"}}`))
				return
			}
			var body map[string]any
			_ = json.NewDecoder(r.Body).Decode(&body)
			if body["board"] != "nano" {
				w.WriteHeader(422)
				_, _ = w.Write([]byte(`{"error":{"code":"invalid_board","message":"Unknown board.","request_id":"R3"}}`))
				return
			}
			files, _ := body["files"].(map[string]any)
			if files["sketch.ino"] == nil {
				w.WriteHeader(422)
				_, _ = w.Write([]byte(`{"error":{"code":"invalid_files","message":"need sketch.ino","request_id":"R4"}}`))
				return
			}
			w.WriteHeader(202)
			w.Header().Set("Location", "/v1/builds/bld_0123456789abcdef")
			_, _ = w.Write([]byte(`{"id":"bld_0123456789abcdef","board":"nano","state":"queued","cost":2,"has_artifact":false}`))
		case r.Method == "GET" && r.URL.Path == "/api/v1/builds/bld_0123456789abcdef":
			_, _ = w.Write([]byte(`{"id":"bld_0123456789abcdef","board":"nano","state":"done","cost":2,"has_artifact":true,"result":{"success":true,"format":"hex","sizes":{"text":924}}}`))
		case r.Method == "GET" && r.URL.Path == "/api/v1/builds/bld_0123456789abcdef/artifact":
			w.Header().Set("Content-Type", "application/octet-stream")
			_, _ = w.Write([]byte(":100000000C9435000C945D000C945D000C945D00CA\r\n"))
		case r.Method == "GET" && r.URL.Path == "/api/v1/boards":
			_, _ = w.Write([]byte(`{"data":[{"id":"nano","name":"Arduino Nano","family":"avr","available":true},{"id":"uno","name":"Arduino Uno","family":"avr","available":false}]}`))
		case r.Method == "GET" && r.URL.Path == "/api/v1/boards/nano":
			_, _ = w.Write([]byte(`{"data":{"id":"nano","name":"Arduino Nano","family":"avr","available":true}}`))
		case r.Method == "GET" && r.URL.Path == "/api/v1/me":
			_, _ = w.Write([]byte(`{"key":{"id":1,"name":"ci","prefix":"mjd_test_x","mode":"test","scopes":["builds:write","boards:read"],"expires_at":"2030-01-01T00:00:00Z","rate_per_minute":60,"daily_quota":100},"developer":{"company":"T","monthly_spend_cap":50,"spent_this_month":4}}`))
		default:
			w.WriteHeader(404)
			_, _ = w.Write([]byte(`{"error":{"code":"not_found","message":"No such endpoint.","request_id":"R9"}}`))
		}
	}))
	t.Cleanup(srv.Close)
	return srv
}

func writeFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	p := filepath.Join(dir, name)
	if err := os.WriteFile(p, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestGatherSketch(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "blink.ino", "void setup(){}\nvoid loop(){}\n")
	writeFile(t, dir, "helper.h", "#pragma once\n")
	writeFile(t, dir, "notes.txt", "not a source file - ignored\n")
	files, base, err := gatherSketch(dir)
	if err != nil {
		t.Fatal(err)
	}
	if base != "blink" || files["sketch.ino"] == "" || files["helper.h"] == "" {
		t.Fatalf("files=%v base=%q", files, base)
	}
	if _, ok := files["notes.txt"]; ok {
		t.Error("notes.txt is not a source file and must not be sent")
	}
	if _, ok := files["blink.ino"]; ok {
		t.Error("the entry file must be renamed to sketch.ino - that is the API's contract")
	}

	// a bare .ino file becomes sketch.ino too
	one := writeFile(t, dir, "solo.ino", "void setup(){}\nvoid loop(){}\n")
	files, base, err = gatherSketch(one)
	if err != nil || base != "solo" || len(files) != 1 || files["sketch.ino"] == "" {
		t.Fatalf("single file: %v %q %v", err, base, files)
	}

	// two .ino in one folder is ambiguous
	writeFile(t, dir, "other.ino", "void setup(){}\nvoid loop(){}\n")
	if _, _, err := gatherSketch(dir); err == nil || !strings.Contains(err.Error(), "two .ino") {
		t.Errorf("two .ino files must be refused: %v", err)
	}
}

func TestCheckFilesMirrorsTheServer(t *testing.T) {
	good := map[string]string{"sketch.ino": "void setup(){}\nvoid loop(){}\n"}
	if err := checkFiles(good); err != nil {
		t.Errorf("a valid sketch must pass: %v", err)
	}
	for name, bad := range map[string]map[string]string{
		"empty":         {},
		"folders":       {"sketch.ino": "x", "sub/a.h": "x"},
		"bad ending":    {"sketch.ino": "x", "evil.exe": "x"},
		"too big":       {"sketch.ino": strings.Repeat("x", maxFileBytes+1)},
		"not utf8":      {"sketch.ino": "\xff\xfe"},
		"total too big": {"sketch.ino": strings.Repeat("a", 200_000), "a.h": strings.Repeat("b", 200_000), "b.h": strings.Repeat("c", 200_000)},
	} {
		if err := checkFiles(bad); err == nil {
			t.Errorf("%s must be refused", name)
		}
	}
}

func TestBuildEndToEnd(t *testing.T) {
	key := "mjd_test_abc"
	srv := fakeDevAPI(t, key)
	t.Setenv("MUJARD_CONFIG", filepath.Join(t.TempDir(), "none.json"))
	dir := t.TempDir()
	writeFile(t, dir, "blink.ino", "void setup(){}\nvoid loop(){}\n")
	out := filepath.Join(dir, "out.hex")

	var buf bytes.Buffer
	code := runBuild([]string{"nano", filepath.Join(dir, "blink.ino"), "--api", srv.URL, "--key", key, "--out", out}, &buf)
	if code != 0 {
		t.Fatalf("exit %d:\n%s", code, buf.String())
	}
	b, err := os.ReadFile(out)
	if err != nil || !strings.HasPrefix(string(b), ":100000") {
		t.Fatalf("the artifact must be the downloaded hex: %v %q", err, b)
	}
	if !strings.Contains(buf.String(), "Built") || !strings.Contains(buf.String(), "coin") {
		t.Errorf("the summary must say what happened and what it cost:\n%s", buf.String())
	}
}

func TestBuildFailures(t *testing.T) {
	key := "mjd_test_abc"
	srv := fakeDevAPI(t, key)
	dir := t.TempDir()
	writeFile(t, dir, "blink.ino", "void setup(){}\n")
	// an empty saved config and no MUJARD_KEY: no key may leak in from the machine running the tests
	t.Setenv("MUJARD_CONFIG", filepath.Join(t.TempDir(), "none.json"))
	t.Setenv("MUJARD_KEY", "")

	var buf bytes.Buffer
	// "uno" is a valid-looking id the fake refuses -> the server's refusal is shown, exit 1
	if code := runBuild([]string{"uno", filepath.Join(dir, "blink.ino"), "--api", srv.URL, "--key", key}, &buf); code != 1 || !strings.Contains(buf.String(), "invalid_board") {
		t.Errorf("server refusal must be shown (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runBuild([]string{"nano", filepath.Join(dir, "blink.ino"), "--api", srv.URL}, &buf); code != 2 || !strings.Contains(buf.String(), "needs a developer key") {
		t.Errorf("no key must say what to do (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runBuild([]string{}, &buf); code != 2 {
		t.Errorf("no board is a usage error (exit %d)", code)
	}
}

func TestLoginConfigRoundTrip(t *testing.T) {
	key := "mjd_test_abc"
	srv := fakeDevAPI(t, key)
	cfg := filepath.Join(t.TempDir(), "config.json")
	t.Setenv("MUJARD_CONFIG", cfg)

	var buf bytes.Buffer
	if code := runLogin([]string{"--api", srv.URL, "--key", key}, &buf); code != 0 {
		t.Fatalf("login exit %d:\n%s", code, buf.String())
	}
	c := loadConfig()
	if c.Key != key || !strings.HasSuffix(c.API, srv.URL) {
		t.Fatalf("config not saved: %+v", c)
	}
	buf.Reset()
	runConfig(nil, &buf)
	if strings.Contains(buf.String(), key) {
		t.Errorf("config show must mask the key:\n%s", buf.String())
	}
	buf.Reset()
	if code := runLogin([]string{"--api", srv.URL, "--key", "mjd_test_wrong"}, &buf); code != 1 {
		t.Errorf("a key the server refuses must not be saved (exit %d):\n%s", code, buf.String())
	}
}

func TestPluginDiscovery(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("MUJARD_COMMANDS", dir)
	ext := ".ps1"
	if runtime.GOOS != "windows" {
		ext = ".sh"
	}
	writeFile(t, dir, "mujard-hello"+ext, "echo hi")
	writeFile(t, dir, "not-a-plugin.txt", "x")
	writeFile(t, dir, "mujard-status.ps1", "echo shadowed") // must not shadow a built-in

	if got := findPlugin("hello"); !strings.HasSuffix(got, "mujard-hello"+ext) {
		t.Errorf("findPlugin(hello) = %q", got)
	}
	for _, p := range listPlugins() {
		if p.name == "status" {
			t.Error("a plugin must never shadow a built-in command")
		}
		if p.name == "not-a-plugin" {
			t.Error("non mujard-* files are not plugins")
		}
	}
	if findPlugin("../escape") != "" || findPlugin("-x") != "" || findPlugin("a/b") != "" {
		t.Error("paths and flags must not resolve as plugins")
	}
	if findPlugin("definitely-not-here-xyz") != "" {
		t.Error("missing plugin should be empty")
	}
}

// Phase 1 aliases: "boards list" is "boards", "board info <id>" is "boards <id>", "upload" is run.
func TestAliasNormalization(t *testing.T) {
	key := "mjd_test_abc"
	srv := fakeDevAPI(t, key)
	t.Setenv("MUJARD_CONFIG", filepath.Join(t.TempDir(), "none.json"))
	base := []string{"--api", srv.URL, "--key", key}

	var buf bytes.Buffer
	if code := runBoards(append([]string{"list"}, base...), &buf); code != 0 || !strings.Contains(buf.String(), "nano") {
		t.Fatalf("boards list must list boards (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runBoards(append([]string{"list", "extra"}, base...), &buf); code != 2 {
		t.Errorf("boards list <extra> is a usage error (exit %d)", code)
	}
	buf.Reset()
	if code := runBoardInfo(append([]string{"info", "nano"}, base...), &buf); code != 0 || !strings.Contains(buf.String(), "Arduino Nano") {
		t.Errorf("board info nano must show the board (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runBoardInfo(append([]string{"nano"}, base...), &buf); code != 0 || !strings.Contains(buf.String(), "Arduino Nano") {
		t.Errorf("board <id> without \"info\" works too (exit %d):\n%s", code, buf.String())
	}
	buf.Reset()
	if code := runBoardInfo(base, &buf); code != 2 {
		t.Errorf("board with no id is a usage error (exit %d)", code)
	}
	// The alias verbs exist and route where they claim to.
	for _, v := range []string{"upload", "board"} {
		c := findBuiltin(v)
		if c == nil {
			t.Fatalf("the %q command must be built in", v)
		}
	}
	if got := findBuiltin("upload"); got == nil || !strings.Contains(got.summary, "run") {
		t.Error("upload must present itself as run's alias")
	}
}

// "mujard ai model list|info" goes to the desktop app; every other "mujard ai ..." stays a prompt.
func TestAIModelRouting(t *testing.T) {
	// Includes the near-misses that must never become a paid prompt: bare "model", "model use x".
	for _, args := range [][]string{{"model", "list"}, {"model", "info", "zhipu:glm-4.5"}, {"model", "list", "--format", "json"}, {"model", "info"}, {"model"}, {"model", "use", "x"}} {
		if !isModelCommand(args) {
			t.Errorf("%q must be a model command", args)
		}
	}
	for _, args := range [][]string{{"model of the esp32 pinout?"}, {"model", "of", "the", "esp32"}, {"why", "is", "my", "LED", "off"}, {"list", "model"}, {}} {
		if isModelCommand(args) {
			t.Errorf("%q must stay an AI prompt", args)
		}
	}
}

func TestApiErrorText(t *testing.T) {
	doc := map[string]any{"error": map[string]any{"code": "rate_limited", "message": "slow down", "request_id": "ABC"}}
	s := apiErrText(doc)
	if !strings.Contains(s, "rate_limited") || !strings.Contains(s, "ABC") {
		t.Errorf("code and request id must be in the message: %q", s)
	}
	if got := apiErrText(map[string]any{}); !strings.Contains(got, "unexpected") {
		t.Errorf("a body that is not an error envelope still needs a line: %q", got)
	}
}
