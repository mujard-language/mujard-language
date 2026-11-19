package main

// Phase 4: the compatibility contract, checked against the REAL built program (main(), the dispatcher, the
// operator gate, the exit codes) rather than the handler functions. These lock down what scripts and people
// already rely on: command spellings, exit codes, the raw JSON of me/usage/boards, build's JSON envelope,
// the operator tier, "mujard ai <prompt>" staying a prompt - plus that no command prints a developer key.
// Nothing here touches hardware, a real server or billing: the API is a local fake, and no desktop app is
// found (the delegated commands are expected to say so).
import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	builtMu    sync.Mutex
	builtPaths = map[string]string{}
)

// buildProgram builds this package once per build-tag set per test run and returns the executable.
// "" is the public build (what users get); "operator" is the server-side build.
func buildProgram(t *testing.T, tags string) string {
	t.Helper()
	builtMu.Lock()
	defer builtMu.Unlock()
	if p, ok := builtPaths[tags]; ok {
		return p
	}
	dir, err := os.MkdirTemp("", "mujard-compat-bin")
	if err != nil {
		t.Skipf("no temp folder: %v", err)
	}
	name := "mujard"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	path := filepath.Join(dir, name)
	goBin := filepath.Join(runtime.GOROOT(), "bin", "go")
	if runtime.GOOS == "windows" {
		goBin += ".exe"
	}
	args := []string{"build", "-o", path}
	if tags != "" {
		args = append(args, "-tags", tags)
	}
	if out, err := exec.Command(goBin, append(args, ".")...).CombinedOutput(); err != nil {
		t.Skipf("cannot build the program for the compatibility tests: %v\n%s", err, out)
	}
	builtPaths[tags] = path
	return path
}

// mujardBinary is the public build.
func mujardBinary(t *testing.T) string { return buildProgram(t, "") }

// runMujard runs the public build in a clean environment: no key, no saved config, no desktop app, no operator gate.
func runMujard(t *testing.T, env map[string]string, args ...string) (stdout, stderr string, code int) {
	t.Helper()
	return runMujardBin(t, mujardBinary(t), env, args...)
}

func runMujardBin(t *testing.T, bin string, env map[string]string, args ...string) (stdout, stderr string, code int) {
	t.Helper()
	home := t.TempDir()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.WaitDelay = 15 * time.Second // abandon output pipes a lingering grandchild may hold
	cmd.Env = []string{
		"MUJARD_HOME=" + home, "MUJARD_CONFIG=" + filepath.Join(home, "config.json"),
		"PATH=" + home, "LOCALAPPDATA=" + home, "USERPROFILE=" + home, "HOME=" + home, "SystemRoot=" + os.Getenv("SystemRoot"),
		"MUJARD_ROOT=" + home, "MUJARD_COMMANDS=" + filepath.Join(home, "commands"),
	}
	for k, v := range env {
		cmd.Env = append(cmd.Env, k+"="+v)
	}
	var so, se bytes.Buffer
	cmd.Stdout, cmd.Stderr = &so, &se
	err := cmd.Run()
	code = 0
	if ee, ok := err.(*exec.ExitError); ok {
		code = ee.ExitCode()
	} else if err != nil {
		t.Fatalf("running mujard %v: %v", args, err)
	}
	return so.String(), se.String(), code
}

func TestCompatOfflineBehaviourIsUnchanged(t *testing.T) {
	out, _, code := runMujard(t, nil, "version")
	if code != 0 || strings.TrimSpace(out) != version {
		t.Errorf("version: %d %q", code, out)
	}
	for _, v := range []string{"--version", "-v"} {
		if out, _, code := runMujard(t, nil, v); code != 0 || strings.TrimSpace(out) != version {
			t.Errorf("%s: %d %q", v, code, out)
		}
	}
	// from a pipe, bare mujard is still the help with exit 2
	if out, _, code := runMujard(t, nil); code != 2 || !strings.Contains(out, "mujard - the Mujard command line") {
		t.Errorf("bare mujard from a pipe: %d %q", code, out)
	}
	out, _, code = runMujard(t, nil, "help")
	if code != 0 || !strings.HasSuffix(strings.TrimSpace(out), "Exit codes: 0 done, 1 failed, 2 bad command line.") {
		t.Errorf("help: %d", code)
	}
	for _, hidden := range []string{"loadtest", "admin ", "doctor ", "status "} {
		if strings.Contains(out, "\n  "+hidden) {
			t.Errorf("operator command %q must not be in a user's help", hidden)
		}
	}
	// every pre-existing command is still listed
	for _, name := range []string{"login", "logout", "config", "me", "usage", "boards", "board", "build", "ai", "webhooks", "devices", "list", "drivers", "adapters", "flash", "burn", "run", "upload", "setup", "shell"} {
		if !strings.Contains(out, "\n  "+name+" ") {
			t.Errorf("help lost %q", name)
		}
	}
	out, _, code = runMujard(t, nil, "help", "build")
	want := "mujard build <board> [sketch|folder] [--out f] [--wait 10m]\n\n  compile a sketch for a board: mujard build nano blink.ino\n\nExamples\n"
	if code != 0 || !strings.HasPrefix(strings.ReplaceAll(out, "\r\n", "\n"), want) {
		t.Errorf("help build changed: %d %q", code, out)
	}
	if _, se, code := runMujard(t, nil, "frobnicate"); code != 2 || !strings.Contains(se, `unknown command "frobnicate"`) {
		t.Errorf("unknown verb: %d %q", code, se)
	}
	for _, c := range []struct {
		args []string
		code int
		has  string
	}{
		{[]string{"me"}, 2, "This command needs a developer key."},
		{[]string{"usage"}, 2, "needs a developer key"},
		{[]string{"boards"}, 2, "needs a developer key"},
		{[]string{"boards", "list"}, 2, "needs a developer key"},
		{[]string{"board", "info", "nano"}, 2, "needs a developer key"},
		{[]string{"ai", "why is my LED off"}, 2, "needs a developer key"},
		{[]string{"webhooks", "list"}, 2, "needs a developer key"},
		{[]string{"build", "nano", "x.ino"}, 2, "needs a developer key"},
		{[]string{"build", "BAD BOARD"}, 2, "is not a board id"},
		{[]string{"build"}, 2, "Usage: mujard build"},
		{[]string{"login"}, 2, "Give the key:"},
		{[]string{"config"}, 0, "Nothing saved."},
		{[]string{"board"}, 2, "Say which board"},
		{[]string{"boards", "list", "extra"}, 2, "takes nothing after"},
		{[]string{"ai", "--format", "json", "hi"}, 2, "flag provided but not defined"},
		{[]string{"webhooks", "bogus"}, 2, "needs a developer key"}, // the key is checked before the sub-command, as it always was
		{[]string{"shell", "x"}, 2, "takes no arguments"},
		// delegated commands with no desktop app installed say so and exit 1, as before
		{[]string{"devices"}, 1, "is done by the desktop app"},
		{[]string{"list"}, 1, "is done by the desktop app"},
		{[]string{"drivers"}, 1, "is done by the desktop app"},
		{[]string{"adapters"}, 1, "is done by the desktop app"},
		{[]string{"flash", "x.hex", "--board", "uno"}, 1, "is done by the desktop app"},
		{[]string{"burn", "uno", "--port", "COM3"}, 1, "is done by the desktop app"},
		{[]string{"run", "x.ino", "--board", "uno", "--port", "COM3"}, 1, "is done by the desktop app"},
		{[]string{"upload", "x.ino", "--board", "uno", "--port", "COM3"}, 1, "is done by the desktop app"},
		{[]string{"ai", "model", "list"}, 1, "is done by the desktop app"},
		{[]string{"ai", "model", "use", "x"}, 1, "is done by the desktop app"},
	} {
		so, se, got := runMujard(t, nil, c.args...)
		if got != c.code || !strings.Contains(so+se, c.has) {
			t.Errorf("mujard %v: exit %d (want %d), output %q", c.args, got, c.code, so+se)
		}
	}
	// (the operator tier is covered, much harder, by the operator build's own tests)
}

func TestCompatJSONContractsAgainstALocalAPI(t *testing.T) {
	key := "mjd_test_compatkey0001"
	srv := fakeDevAPI(t, key)
	env := map[string]string{"MUJARD_API": srv.URL, "MUJARD_KEY": key}

	// me / usage / boards / board: --format json is the server's document, byte for byte, no envelope.
	boardsRaw, _, code := runMujard(t, env, "boards", "--format", "json")
	if code != 0 || !strings.HasPrefix(boardsRaw, `{"data":[{"id":"nano"`) || strings.Contains(boardsRaw, "mujard-cli") {
		t.Errorf("boards --format json must stay the raw document: %d %q", code, boardsRaw)
	}
	viaGet, _, code2 := runMujard(t, env, "get", "boards", "--format", "json")
	if code2 != 0 || viaGet != boardsRaw {
		t.Errorf("get boards --format json differs from boards --format json:\n%q\n%q", viaGet, boardsRaw)
	}
	for _, args := range [][]string{{"boards"}, {"boards", "list"}, {"board", "info", "nano"}, {"board", "nano"}, {"boards", "nano"}, {"me"}} {
		_, se, code := runMujard(t, env, args...)
		if code != 0 {
			t.Errorf("mujard %v: exit %d %q", args, code, se)
		}
	}
	if me, _, _ := runMujard(t, env, "me", "--format", "json"); !strings.HasPrefix(me, `{"key":{"id":1`) {
		t.Errorf("me --format json stays raw: %q", me)
	}
	a, _, _ := runMujard(t, env, "boards")
	b, _, _ := runMujard(t, env, "get", "boards")
	if a != b {
		t.Errorf("get boards must print what boards prints:\n%s\n%s", a, b)
	}
	if a, b := func() (string, string) {
		x, _, _ := runMujard(t, env, "board", "info", "nano")
		y, _, _ := runMujard(t, env, "get", "board", "nano")
		return x, y
	}(); a != b {
		t.Errorf("get board nano must print what board info nano prints")
	}

	// build --format json: the slim envelope - schema, command, ok, data (success) / error (failure); no null keys.
	dir := t.TempDir()
	sk := writeFile(t, dir, "blink.ino", "void setup(){}\nvoid loop(){}\n")
	out, _, code := runMujard(t, env, "build", "nano", sk, "--out", filepath.Join(dir, "o.hex"), "--format", "json")
	var okDoc map[string]any
	if err := json.Unmarshal([]byte(strings.TrimSpace(out)), &okDoc); err != nil || code != 0 {
		t.Fatalf("build json: %d %v %q", code, err, out)
	}
	if len(okDoc) != 4 || okDoc["schema"] != "mujard-cli.v1" || okDoc["command"] != "build" || okDoc["ok"] != true || okDoc["data"] == nil {
		t.Errorf("build success envelope changed: %q", out)
	}
	data := okDoc["data"].(map[string]any)
	if data["build"] == nil || data["artifact"] == nil {
		t.Errorf("build success data lost build/artifact: %q", out)
	}
	out, _, code = runMujard(t, env, "build", "uno", sk, "--format", "json")
	var failDoc map[string]any
	if err := json.Unmarshal([]byte(strings.TrimSpace(out)), &failDoc); err != nil || code != 1 {
		t.Fatalf("build failure json: %d %v %q", code, err, out)
	}
	if len(failDoc) != 4 || failDoc["ok"] != false || failDoc["error"].(map[string]any)["code"] != "invalid_board" {
		t.Errorf("build failure envelope changed: %q", out)
	}
	if _, has := failDoc["data"]; has {
		t.Errorf("build failure has no data key: %q", out)
	}
}

func TestCompatNoCommandPrintsADeveloperKey(t *testing.T) {
	secret := "mjd_live_SUPERSECRETVALUE0123456789"
	srv := fakeV1Static(t, secret)
	env := map[string]string{"MUJARD_API": srv.URL, "MUJARD_KEY": secret}
	dir := t.TempDir()
	sk := writeFile(t, dir, "blink.ino", "void setup(){}\nvoid loop(){}\n")
	for _, args := range [][]string{
		{"auth", "status"}, {"auth", "test"}, {"auth", "status", "--format", "json"}, {"config"}, {"config", "get", "key"},
		{"me"}, {"me", "--format", "json"}, {"usage"}, {"boards"}, {"builds", "list"}, {"check", "sketch", "nano", sk},
		{"check", "server"}, {"diagnose"}, {"webhooks", "list"}, {"build", "uno", sk}, {"help"}, {"commands"},
		{"login", "--key", secret, "--api", srv.URL}, {"config"}, {"config", "get", "key"}, {"auth", "status"},
		{"builds", "get", "bld_ffffffffffffffff"}, {"webhooks", "info", "99"},
	} {
		so, se, _ := runMujard(t, env, args...)
		if strings.Contains(so+se, "SUPERSECRETVALUE") {
			t.Errorf("mujard %v printed the key: %q", args, so+se)
		}
	}
}

// fakeV1Static is a tiny fixed API for the key-leak sweep: every call answers, nothing is stored.
func fakeV1Static(t *testing.T, key string) *httptest.Server {
	var hits []string
	return fakeV1(t, key, &hits)
}

func TestREADMEMentionsEveryCommand(t *testing.T) {
	b, err := os.ReadFile("README.md")
	if err != nil {
		t.Skip("README.md not found")
	}
	readme := string(b)
	for _, c := range builtins {
		if !strings.Contains(readme, "mujard "+c.name) {
			t.Errorf("README.md does not document %q", c.name)
		}
		if strings.TrimSpace(c.summary) == "" || c.group == "" {
			t.Errorf("%s needs a summary and a help group", c.name)
		}
	}
}

func TestEveryBuiltinHasWorkingHelp(t *testing.T) {
	t.Setenv("MUJARD_ADMIN", "1") // the operator sees them all
	for _, c := range builtins {
		var out bytes.Buffer
		printHelp(&out, c.name)
		if !strings.HasPrefix(out.String(), "mujard "+c.name) {
			t.Errorf("help %s: %q", c.name, out.String())
		}
	}
}
