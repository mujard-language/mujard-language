package main

// Tests for the Phase 1 command language: get/models aliases, auth, config, builds, webhooks,
// boards search, check, diagnose and "did you mean". The old commands' output is asserted unchanged.
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

const testBuildID = "bld_0123456789abcdef"

// fakeV1 serves what the real Api.php serves for the routes Phase 1 uses, and records the requests.
func fakeV1(t *testing.T, key string, hits *[]string) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*hits = append(*hits, r.Method+" "+r.URL.RequestURI())
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/health.php" {
			_, _ = w.Write([]byte(`{"ok":true,"maintenance":false,"compilerReady":true,"aiReady":false,"cores":["arduino:avr"],"serverId":"x"}`))
			return
		}
		if r.Header.Get("Authorization") != "Bearer "+key {
			w.WriteHeader(401)
			_, _ = w.Write([]byte(`{"error":{"code":"invalid_api_key","message":"bad key","request_id":"R1"}}`))
			return
		}
		p := strings.TrimPrefix(r.URL.Path, "/api/v1")
		switch {
		case r.Method == "GET" && p == "/me":
			_, _ = w.Write([]byte(`{"key":{"id":1,"name":"ci","prefix":"mjd_test_x","mode":"test","scopes":["builds:read","boards:read"],"expires_at":"2030-01-01T00:00:00Z","rate_per_minute":60,"daily_quota":100},"developer":{"company":"T","monthly_spend_cap":50,"spent_this_month":4}}`))
		case r.Method == "GET" && p == "/boards":
			_, _ = w.Write([]byte(`{"data":[{"id":"nano","name":"Arduino Nano","family":"avr","available":true},{"id":"uno","name":"Arduino Uno","family":"avr","available":false},{"id":"esp32","name":"ESP32 Dev Module","family":"esp32","available":true}]}`))
		case r.Method == "GET" && p == "/boards/nano":
			_, _ = w.Write([]byte(`{"id":"nano","name":"Arduino Nano","family":"avr","available":true}`))
		case r.Method == "GET" && p == "/boards/uno":
			_, _ = w.Write([]byte(`{"id":"uno","name":"Arduino Uno","family":"avr","available":false}`))
		case r.Method == "GET" && p == "/boards/ghost":
			w.WriteHeader(404)
			_, _ = w.Write([]byte(`{"error":{"code":"not_found","message":"No such board.","request_id":"R2"}}`))
		case r.Method == "GET" && p == "/builds":
			if r.URL.Query().Get("state") == "failed" {
				_, _ = w.Write([]byte(`{"data":[],"has_more":false,"next_cursor":null}`))
				return
			}
			_, _ = w.Write([]byte(`{"data":[{"id":"` + testBuildID + `","board":"nano","state":"done","mode":"test","cost":0,"created_at":"2026-10-17T01:00:00+00:00"}],"has_more":true,"next_cursor":"` + testBuildID + `"}`))
		case r.Method == "GET" && p == "/builds/"+testBuildID:
			_, _ = w.Write([]byte(`{"id":"` + testBuildID + `","board":"nano","state":"done","mode":"live","cost":2,"refunded":false,"created_at":"2026-10-17T01:00:00+00:00","has_artifact":true,"result":{"success":true,"format":"hex","sizes":{"text":924}}}`))
		case r.Method == "GET" && p == "/builds/"+testBuildID+"/artifact":
			_, _ = w.Write([]byte(":100000000C9435000C945D000C945D000C945D00CA\r\n"))
		case r.Method == "DELETE" && p == "/builds/"+testBuildID:
			_, _ = w.Write([]byte(`{"id":"` + testBuildID + `","state":"cancelled"}`))
		case r.Method == "GET" && p == "/webhooks/3":
			_, _ = w.Write([]byte(`{"id":3,"url":"https://example.com/hook","enabled":true}`))
		case r.Method == "POST" && p == "/webhooks/3/enable":
			_, _ = w.Write([]byte(`{"id":3,"enabled":true}`))
		case r.Method == "POST" && p == "/webhooks/3/rotate-secret":
			_, _ = w.Write([]byte(`{"id":3,"secret":"whsec_new_value","secret_note":"shown once"}`))
		case r.Method == "GET" && p == "/usage":
			w.WriteHeader(403)
			_, _ = w.Write([]byte(`{"error":{"code":"insufficient_scope","message":"No scope.","request_id":"R3","details":{"required_scope":"usage:read"}}}`))
		default:
			w.WriteHeader(404)
			_, _ = w.Write([]byte(`{"error":{"code":"not_found","message":"No such endpoint.","request_id":"R9"}}`))
		}
	}))
	t.Cleanup(srv.Close)
	return srv
}

func isolate(t *testing.T) {
	t.Helper()
	t.Setenv("MUJARD_CONFIG", filepath.Join(t.TempDir(), "config.json"))
	t.Setenv("MUJARD_KEY", "")
	t.Setenv("MUJARD_API", "")
}

func TestRewriteGet(t *testing.T) {
	cases := map[string][]string{
		"boards --format json": {"boards", "--format", "json"},
		"board nano":           {"board", "info", "nano"},
		"models":               {"models", "list"},
		"model foo":            {"models", "info", "foo"},
		"builds --state done":  {"builds", "list", "--state", "done"},
		"build " + testBuildID: {"builds", "get", testBuildID},
		"webhook 3":            {"webhooks", "info", "3"},
		"usage":                {"usage"},
	}
	for in, want := range cases {
		got, ok := rewriteGet(strings.Fields(in))
		if !ok || !reflect.DeepEqual(got, want) {
			t.Errorf("get %s -> %q (ok %v), want %q", in, got, ok, want)
		}
	}
	for _, bad := range []string{"", "doctor", "status", "admin", "frobs"} {
		if _, ok := rewriteGet(strings.Fields(bad)); ok {
			t.Errorf("get %q must not be a thing", bad)
		}
	}
	// get must never reach an operator command, whatever the table says
	for noun, canon := range getNouns {
		if operatorOnly[canon[0]] {
			t.Errorf("get %s points at operator command %s", noun, canon[0])
		}
		if findBuiltin(canon[0]) == nil {
			t.Errorf("get %s points at %s, which is not a built-in", noun, canon[0])
		}
	}
}

func TestGetIsByteIdenticalToTheCommandItStandsFor(t *testing.T) {
	isolate(t)
	var hits []string
	srv := fakeV1(t, "mjd_test_abc", &hits)
	base := []string{"--api", srv.URL, "--key", "mjd_test_abc"}
	for _, pair := range [][2][]string{
		{{"boards"}, {"boards"}},
		{{"boards", "--format", "json"}, {"boards", "--format", "json"}},
		{{"board", "nano"}, {"board", "info", "nano"}},
		{{"me"}, {"me"}},
		{{"builds", "--state", "done"}, {"builds", "list", "--state", "done"}},
		{{"build", testBuildID}, {"builds", "get", testBuildID}},
		{{"webhook", "3"}, {"webhooks", "info", "3"}},
	} {
		var viaGet, direct bytes.Buffer
		c1 := runGet(append(append([]string{}, pair[0]...), base...), &viaGet)
		canon := findBuiltin(pair[1][0])
		c2 := canon.run(append(append([]string{}, pair[1][1:]...), base...), &direct)
		if c1 != c2 || viaGet.String() != direct.String() {
			t.Errorf("get %v differs from %v:\n%s\n---\n%s", pair[0], pair[1], viaGet.String(), direct.String())
		}
	}
	var out bytes.Buffer
	if code := runGet([]string{"nonsense"}, &out); code != 2 || !strings.Contains(out.String(), "boards") {
		t.Errorf("get <unknown> is a usage error that names the choices (exit %d): %s", code, out.String())
	}
}

func TestSuggest(t *testing.T) {
	names := []string{"boards", "build", "builds", "models", "diagnose", "webhooks", "login"}
	for in, want := range map[string]string{"bords": "boards", "model": "models", "diagnos": "diagnose", "logn": "login", "webhook": "webhooks"} {
		got := suggest(in, names)
		if len(got) == 0 || got[0] != want {
			t.Errorf("suggest(%q) = %q, want first %q", in, got, want)
		}
	}
	if got := suggest("zzzzzz", names); len(got) != 0 {
		t.Errorf("nothing is close to zzzzzz: %q", got)
	}
}

func TestSuggestionsNeverOfferOperatorCommands(t *testing.T) {
	t.Setenv("MUJARD_ADMIN", "")
	for _, n := range visibleVerbs() {
		if operatorOnly[n] {
			t.Errorf("%s is operator-only and must not be suggested to a user", n)
		}
	}
	t.Setenv("MUJARD_ADMIN", "1")
	found := false
	for _, n := range visibleVerbs() {
		found = found || n == "status"
	}
	if found != operatorBuild {
		t.Errorf("operator commands visible = %v, but operatorBuild = %v: only the operator build may list them", found, operatorBuild)
	}
}

func TestAuthStatusAndTest(t *testing.T) {
	isolate(t)
	var hits []string
	srv := fakeV1(t, "mjd_test_abcdefghijkl", &hits)
	var out bytes.Buffer
	if code := runAuth([]string{"status"}, &out); code != 1 || !strings.Contains(out.String(), "Not signed in") {
		t.Errorf("no key: status says so and exits 1 (exit %d): %s", code, out.String())
	}
	out.Reset()
	hits = nil
	if code := runAuth([]string{"status", "--key", "mjd_test_abcdefghijkl", "--api", srv.URL}, &out); code != 0 {
		t.Fatalf("exit %d: %s", code, out.String())
	}
	if strings.Contains(out.String(), "abcdefghijkl") || !strings.Contains(out.String(), "test key") || !strings.Contains(out.String(), "--key") {
		t.Errorf("status shows the masked key, its mode and where it came from - never the key:\n%s", out.String())
	}
	if len(hits) != 0 {
		t.Errorf("auth status must not touch the network: %v", hits)
	}
	out.Reset()
	if code := runAuth([]string{"test", "--key", "mjd_test_abcdefghijkl", "--api", srv.URL}, &out); code != 0 || !strings.Contains(out.String(), "accepts this key") {
		t.Errorf("test with a good key (exit %d): %s", code, out.String())
	}
	out.Reset()
	if code := runAuth([]string{"test", "--key", "mjd_test_wrong", "--api", srv.URL}, &out); code != 1 || !strings.Contains(out.String(), "invalid_api_key") {
		t.Errorf("test with a bad key exits 1 and says why (exit %d): %s", code, out.String())
	}
	out.Reset()
	if code := runAuth([]string{"test"}, &out); code != 2 {
		t.Errorf("test with no key is a usage error (exit %d)", code)
	}
	out.Reset()
	if code := runAuth([]string{"status", "--format", "json", "--key", "mjd_test_abcdefghijkl", "--api", srv.URL}, &out); code != 0 {
		t.Fatalf("json status exit %d", code)
	}
	var env map[string]any
	if err := json.Unmarshal(out.Bytes(), &env); err != nil || env["schema"] != "mujard-cli.v1" || env["ok"] != true || env["command"] != "auth status" {
		t.Errorf("auth status --format json is the mujard-cli.v1 envelope: %v %s", err, out.String())
	}
	if strings.Contains(out.String(), "abcdefghijkl") {
		t.Error("the JSON must not carry the key")
	}
	if code := runAuth([]string{"bogus"}, &out); code != 2 {
		t.Errorf("unknown auth sub-command is exit 2, got %d", code)
	}
}

func TestConfigSubcommands(t *testing.T) {
	isolate(t)
	var out bytes.Buffer
	// bare config is unchanged
	if code := runConfig(nil, &out); code != 0 || !strings.Contains(out.String(), "Nothing saved") {
		t.Errorf("bare config: %d %s", code, out.String())
	}
	out.Reset()
	if code := runConfig([]string{"set", "api", "https://h.example/mujard-api/"}, &out); code != 0 {
		t.Fatalf("set api: %d %s", code, out.String())
	}
	if loadConfig().API != "https://h.example/mujard-api" {
		t.Errorf("saved %q", loadConfig().API)
	}
	out.Reset()
	runConfig([]string{"get", "api"}, &out)
	if strings.TrimSpace(out.String()) != "https://h.example/mujard-api" {
		t.Errorf("get api: %q", out.String())
	}
	for _, bad := range [][]string{{"set", "api", "not a url"}, {"set", "api", "ftp://x/y"}, {"set", "key", "mjd_live_secret"}, {"get", "nothing"}, {"set"}, {"path", "extra"}} {
		out.Reset()
		if code := runConfig(bad, &out); code != 2 {
			t.Errorf("config %v must be a usage error, got %d: %s", bad, code, out.String())
		}
		if strings.Contains(out.String(), "mjd_live_secret") {
			t.Error("a key typed after config set must not be echoed")
		}
	}
	// the key is only ever shown masked
	if err := saveConfig(cliConfig{API: "https://h.example/mujard-api", Key: "mjd_live_abcdefghijklmnop"}); err != nil {
		t.Fatal(err)
	}
	out.Reset()
	runConfig([]string{"get", "key"}, &out)
	if strings.Contains(out.String(), "abcdefghijklmnop") || !strings.Contains(out.String(), "mjd_live") {
		t.Errorf("get key must be masked: %q", out.String())
	}
	out.Reset()
	runConfig([]string{"path"}, &out)
	if strings.TrimSpace(out.String()) != configPath() {
		t.Errorf("path: %q", out.String())
	}
	// "config <anything else>" still prints the settings, as it always did
	out.Reset()
	if code := runConfig([]string{"whatever"}, &out); code != 0 || !strings.Contains(out.String(), "file ") {
		t.Errorf("unknown words after config keep the old behaviour: %d %s", code, out.String())
	}
}

func TestBoardsSearchAndInfo(t *testing.T) {
	isolate(t)
	var hits []string
	srv := fakeV1(t, "k", &hits)
	base := []string{"--api", srv.URL, "--key", "k"}
	var out bytes.Buffer
	if code := runBoards(append([]string{"search", "esp"}, base...), &out); code != 0 || !strings.Contains(out.String(), "esp32") || strings.Contains(out.String(), "Arduino Nano") {
		t.Errorf("search esp: %d\n%s", code, out.String())
	}
	out.Reset()
	if code := runBoards(append([]string{"search", "arduino", "--available"}, base...), &out); code != 0 || !strings.Contains(out.String(), "nano") || strings.Contains(out.String(), "Arduino Uno") {
		t.Errorf("--available hides boards that cannot be built now: %d\n%s", code, out.String())
	}
	out.Reset()
	runBoards(append([]string{"search", "esp", "--format", "json"}, base...), &out)
	var doc struct {
		Data []map[string]any `json:"data"`
	}
	if err := json.Unmarshal(out.Bytes(), &doc); err != nil || len(doc.Data) != 1 || doc.Data[0]["id"] != "esp32" {
		t.Errorf("search --format json keeps the API's {data:[...]} shape: %v %s", err, out.String())
	}
	out.Reset()
	if code := runBoards(append([]string{"search"}, base...), &out); code != 2 {
		t.Errorf("search with no text is a usage error (exit %d)", code)
	}
	out.Reset()
	if code := runBoards(append([]string{"search", "zzz"}, base...), &out); code != 0 || !strings.Contains(out.String(), "No board matches") {
		t.Errorf("no match is not an error: %d %s", code, out.String())
	}
	var a, b bytes.Buffer
	runBoards(append([]string{"info", "nano"}, base...), &a)
	runBoards(append([]string{"nano"}, base...), &b)
	if a.String() != b.String() || !strings.Contains(a.String(), "Arduino Nano") {
		t.Errorf("boards info nano == boards nano:\n%s\n%s", a.String(), b.String())
	}
	out.Reset()
	if code := runBoards(append([]string{"info"}, base...), &out); code != 2 {
		t.Errorf("boards info with no id is a usage error (exit %d)", code)
	}
}

func TestBuildsCommands(t *testing.T) {
	isolate(t)
	var hits []string
	srv := fakeV1(t, "k", &hits)
	base := []string{"--api", srv.URL, "--key", "k"}
	run := func(args ...string) (int, string) {
		var out bytes.Buffer
		code := runBuilds(append(append([]string{}, args...), base...), &out)
		return code, out.String()
	}
	code, out := run("list", "--limit", "5", "--state", "done", "--mode", "test")
	if code != 0 || !strings.Contains(out, testBuildID) || !strings.Contains(out, "--after "+testBuildID) {
		t.Errorf("list: %d\n%s", code, out)
	}
	last := hits[len(hits)-1]
	if !strings.Contains(last, "limit=5") || !strings.Contains(last, "state=done") || !strings.Contains(last, "mode=test") {
		t.Errorf("filters reach the API: %s", last)
	}
	if code, out = run("list", "--state", "failed"); code != 0 || !strings.Contains(out, "No builds match") {
		t.Errorf("empty list: %d %s", code, out)
	}
	for _, bad := range [][]string{{"list", "--limit", "0"}, {"list", "--limit", "101"}, {"list", "--state", "weird"}, {"list", "--mode", "x"}, {"list", "--after", "nope"}, {"list", "extra"}, {"get"}, {"get", "nano"}, {"artifact", "../x"}, {"frob"}} {
		hits = nil
		if code, out = run(bad...); code != 2 {
			t.Errorf("builds %v must be a usage error, got %d: %s", bad, code, out)
		}
		if len(hits) != 0 {
			t.Errorf("builds %v must be refused before any request: %v", bad, hits)
		}
	}
	if code, out = run("get", testBuildID); code != 0 || !strings.Contains(out, "state    done") || !strings.Contains(out, "builds artifact") {
		t.Errorf("get: %d\n%s", code, out)
	}
	if code, out = run("get", testBuildID, "--format", "json"); code != 0 || !strings.HasPrefix(strings.TrimSpace(out), `{"id":"`+testBuildID) {
		t.Errorf("get --format json is the raw document: %d %s", code, out)
	}

	dir := t.TempDir()
	dest := filepath.Join(dir, "f.hex")
	if code, out = run("artifact", testBuildID, "--out", dest); code != 0 {
		t.Fatalf("artifact: %d %s", code, out)
	}
	if b, _ := os.ReadFile(dest); !strings.HasPrefix(string(b), ":100000") {
		t.Errorf("artifact bytes: %q", b)
	}
	if code, out = run("artifact", testBuildID, "--out", dest); code != 1 || !strings.Contains(out, "already exists") {
		t.Errorf("an existing file is never replaced silently: %d %s", code, out)
	}
	if code, _ = run("artifact", testBuildID, "--out", dest, "--force"); code != 0 {
		t.Errorf("--force replaces: %d", code)
	}

	// cancel asks first; end-of-input is "no"
	stdinReader = strings.NewReader("")
	defer func() { stdinReader = os.Stdin }()
	hits = nil
	if code, out = run("cancel", testBuildID); code != 1 || !strings.Contains(out, "Not cancelled") {
		t.Errorf("cancel without an answer cancels nothing: %d %s", code, out)
	}
	for _, h := range hits {
		if strings.HasPrefix(h, "DELETE") {
			t.Error("no DELETE may be sent without a yes")
		}
	}
	stdinReader = strings.NewReader("y\n")
	if code, out = run("cancel", testBuildID); code != 0 || !strings.Contains(out, "cancelled") {
		t.Errorf("cancel with y: %d %s", code, out)
	}
	if code, out = run("cancel", testBuildID, "--yes"); code != 0 {
		t.Errorf("--yes: %d %s", code, out)
	}
}

func TestBuildSubwordsOnlyWhenTheArgumentIsABuildID(t *testing.T) {
	isolate(t)
	var hits []string
	srv := fakeV1(t, "k", &hits)
	var out bytes.Buffer
	code := runBuild([]string{"status", testBuildID, "--api", srv.URL, "--key", "k"}, &out)
	if code != 0 || !strings.Contains(out.String(), "state    done") {
		t.Errorf("build status <id>: %d\n%s", code, out.String())
	}
	// "get" followed by something that is not a build id is still a board named "get": a usage error from build
	hits = nil
	out.Reset()
	if code := runBuild([]string{"get", "nano", "--api", srv.URL, "--key", "k"}, &out); code == 0 {
		t.Errorf("build get nano must not succeed: %s", out.String())
	}
	for _, h := range hits {
		if strings.HasPrefix(h, "GET /api/v1/builds") {
			t.Errorf("build get nano must not be read as builds get: %v", hits)
		}
	}
}

func TestWebhookExtras(t *testing.T) {
	isolate(t)
	var hits []string
	srv := fakeV1(t, "k", &hits)
	base := []string{"--api", srv.URL, "--key", "k"}
	run := func(args ...string) (int, string) {
		var out bytes.Buffer
		code := runWebhooks(append(append([]string{}, args...), base...), &out)
		return code, out.String()
	}
	if code, out := run("info", "3"); code != 0 || !strings.Contains(out, "example.com") {
		t.Errorf("info: %d %s", code, out)
	}
	if code, _ := run("get", "3"); code != 0 {
		t.Errorf("get is info: %d", code)
	}
	if code, _ := run("enable", "3"); code != 0 {
		t.Errorf("enable: %d", code)
	}
	for _, bad := range [][]string{{"info"}, {"info", "abc"}, {"enable", "3/4"}, {"rotate-secret", "../x"}} {
		if code, _ := run(bad...); code != 2 {
			t.Errorf("webhooks %v must be a usage error, got %d", bad, code)
		}
	}
	stdinReader = strings.NewReader("")
	defer func() { stdinReader = os.Stdin }()
	hits = nil
	if code, out := run("rotate-secret", "3"); code != 1 || !strings.Contains(out, "Not changed") {
		t.Errorf("rotating needs a yes: %d %s", code, out)
	}
	for _, h := range hits {
		if strings.Contains(h, "rotate-secret") {
			t.Error("nothing is sent without a yes")
		}
	}
	code, out := run("rotate-secret", "3", "--yes")
	if code != 0 || !strings.Contains(out, "whsec_new_value") || !strings.Contains(out, "only this once") {
		t.Errorf("rotate with --yes shows the new secret once: %d %s", code, out)
	}
	// the existing sub-commands still work as before
	if code, out := run("add"); code != 2 || !strings.Contains(out, "Usage: mujard webhooks add") {
		t.Errorf("add usage unchanged: %d %s", code, out)
	}
}

func TestInsufficientScopeNamesTheScope(t *testing.T) {
	isolate(t)
	var hits []string
	srv := fakeV1(t, "k", &hits)
	var out bytes.Buffer
	if code := runUsage([]string{"--api", srv.URL, "--key", "k"}, &out); code != 1 || !strings.Contains(out.String(), `needs the "usage:read" scope`) {
		t.Errorf("a 403 names the scope: %d %s", code, out.String())
	}
	// an error without details prints exactly as before
	if got := apiErrText(map[string]any{"error": map[string]any{"code": "rate_limited", "message": "slow down", "request_id": "ABC"}}); got != "rate_limited: slow down (request id ABC - quote it if you ask for help)" {
		t.Errorf("old wording changed: %q", got)
	}
}

func TestCheckSketch(t *testing.T) {
	isolate(t)
	var hits []string
	srv := fakeV1(t, "k", &hits)
	dir := t.TempDir()
	sk := writeFile(t, dir, "blink.ino", "void setup(){}\nvoid loop(){}\n")
	run := func(args ...string) (int, string) {
		var out bytes.Buffer
		return runCheck(args, &out), out.String()
	}
	hits = nil
	code, out := run("sketch", "nano", sk, "--api", srv.URL, "--key", "k")
	if code != 0 || !strings.Contains(out, "can be built now") || !strings.Contains(out, "Nothing is compiled") {
		t.Errorf("good sketch: %d\n%s", code, out)
	}
	for _, h := range hits {
		if strings.HasPrefix(h, "POST") {
			t.Errorf("check sketch must never start a build: %v", hits)
		}
	}
	if code, out = run("sketch", "uno", sk, "--api", srv.URL, "--key", "k"); code != 1 || !strings.Contains(out, "not on the build server") {
		t.Errorf("a board that cannot be built now: %d %s", code, out)
	}
	if code, out = run("sketch", "ghost", sk, "--api", srv.URL, "--key", "k"); code != 1 || !strings.Contains(out, "does not know") {
		t.Errorf("unknown board: %d %s", code, out)
	}
	if code, out = run("sketch", "nano", sk, "--api", srv.URL); code != 0 || !strings.Contains(out, "not checked") {
		t.Errorf("without a key only the files are checked: %d %s", code, out)
	}
	bad := writeFile(t, dir, "big.ino", strings.Repeat("x", maxFileBytes+1))
	if code, _ = run("sketch", "nano", bad, "--api", srv.URL, "--key", "k"); code != 2 {
		t.Errorf("oversized sketch is exit 2, got %d", code)
	}
	if code, _ = run("sketch", "BAD BOARD", sk); code != 2 {
		t.Errorf("bad board id exit 2, got %d", code)
	}
	if code, _ = run("sketch"); code != 2 {
		t.Errorf("usage exit 2, got %d", code)
	}
	code, out = run("sketch", "nano", sk, "--api", srv.URL, "--key", "k", "--format", "json")
	var env map[string]any
	if err := json.Unmarshal([]byte(out), &env); err != nil || env["ok"] != true || env["command"] != "check sketch" || code != 0 {
		t.Errorf("json: %d %v %s", code, err, out)
	}
}

func TestCheckServer(t *testing.T) {
	isolate(t)
	var hits []string
	srv := fakeV1(t, "k", &hits)
	var out bytes.Buffer
	if code := runCheck([]string{"server", "--api", srv.URL}, &out); code != 0 || !strings.Contains(out.String(), "builds are on") || !strings.Contains(out.String(), "WARN") {
		t.Errorf("a healthy server with the AI off warns but passes: %d\n%s", code, out.String())
	}
	if strings.Contains(out.String(), "serverId") || strings.Contains(out.String(), "arduino:avr") {
		t.Error("server internals from the health page are not repeated")
	}
	if hits[len(hits)-1] != "GET /health.php" {
		t.Errorf("check server uses the public health page only: %v", hits)
	}
	out.Reset()
	if code := runCheck([]string{"server", "--api", "http://127.0.0.1:1"}, &out); code != 1 || !strings.Contains(out.String(), "FAIL") {
		t.Errorf("an unreachable server fails: %d %s", code, out.String())
	}
	// an --api that already ends in /api/v1 still finds health.php
	if apiRoot("http://h/x/api/v1/") != "http://h/x" {
		t.Error("apiRoot")
	}
}

func TestCheckDesktop(t *testing.T) {
	t.Setenv("MUJARD_DESKTOP", filepath.Join(t.TempDir(), "missing.exe"))
	t.Setenv("PATH", t.TempDir())
	t.Setenv("MUJARD_ROOT", t.TempDir())
	t.Setenv("LOCALAPPDATA", t.TempDir())
	rs := desktopChecks()
	if len(rs) != 1 || (rs[0].lvl != fail && rs[0].lvl != ok && rs[0].lvl != warn) {
		t.Fatalf("one result: %+v", rs)
	}
}

func TestModelsAndDeviceAliasesForwardTheRightArguments(t *testing.T) {
	if got := dropList([]string{"list"}); len(got) != 0 {
		t.Errorf("devices list -> bare list verb: %q", got)
	}
	if got := dropList([]string{"--format", "json"}); !reflect.DeepEqual(got, []string{"--format", "json"}) {
		t.Errorf("other arguments pass through: %q", got)
	}
	if got := dropList([]string{"list", "--format", "json"}); !reflect.DeepEqual(got, []string{"--format", "json"}) {
		t.Errorf("devices list --format json: %q", got)
	}
}

func TestPhase1CommandsAreUserTierAndListed(t *testing.T) {
	for _, n := range []string{"auth", "builds", "check", "diagnose", "get", "models"} {
		c := findBuiltin(n)
		if c == nil {
			t.Fatalf("%s must be built in", n)
		}
		if operatorOnly[n] {
			t.Errorf("%s is for users", n)
		}
		if !selfHelp[n] && n != "auth" {
			t.Errorf("%s should answer --help itself", n)
		}
	}
	var out bytes.Buffer
	t.Setenv("MUJARD_ADMIN", "")
	printHelp(&out, "")
	for _, want := range []string{"builds ", "check ", "diagnose ", "auth ", "models "} {
		if !strings.Contains(out.String(), want) {
			t.Errorf("help lacks %q", want)
		}
	}
	for _, hidden := range []string{"loadtest", "admin ", "status ", "doctor "} {
		if strings.Contains(out.String(), "\n  "+hidden) {
			t.Errorf("operator command %q must stay hidden:\n%s", hidden, out.String())
		}
	}
	out.Reset()
	printHelp(&out, "get")
	if !strings.Contains(out.String(), "Examples") || !strings.Contains(out.String(), "mujard get boards") {
		t.Errorf("help get has examples: %s", out.String())
	}
}

func TestWantsHelp(t *testing.T) {
	if !wantsHelp([]string{"sketch", "--help"}) || !wantsHelp([]string{"-h"}) || wantsHelp([]string{"sketch"}) || wantsHelp([]string{"--", "--help"}) {
		t.Error("wantsHelp")
	}
}

func TestEnvelopeShape(t *testing.T) {
	var out bytes.Buffer
	envFail(&out, "x", "code", "msg", 2)
	var env map[string]any
	_ = json.Unmarshal(out.Bytes(), &env)
	for _, k := range []string{"schema", "command", "ok", "data", "error"} {
		if _, ok := env[k]; !ok {
			t.Errorf("envelope lacks %q: %s", k, out.String())
		}
	}
}
