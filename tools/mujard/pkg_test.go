package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// A stand-in registry: what the real one (backend ApiPackages) answers, including a hostile package.
func fakeRegistry(t *testing.T, published *map[string]any) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == "GET" && r.URL.Path == "/api/v1/packages/blinker/latest":
			io.WriteString(w, `{"name":"blinker","version":"1.2.0","board":"uno","files":{"blinker.muj":"fn twice(x: int) -> int:\n    return x * 2\n","blinker.mlib":"{}"}}`)
		case r.Method == "GET" && r.URL.Path == "/api/v1/packages/evil/latest":
			io.WriteString(w, `{"name":"evil","version":"1.0.0","files":{"../../outside.muj":"x","evil.muj":"y"}}`)
		case r.Method == "POST" && r.URL.Path == "/api/v1/packages":
			var body map[string]any
			json.NewDecoder(r.Body).Decode(&body)
			*published = body
			if body["version"] == "9.9.9" {
				w.WriteHeader(422)
				io.WriteString(w, `{"error":{"code":"check_failed","message":"doesn't pass the Mujard check","details":{"errors":["mine.muj:2 C303 unknown name 'y'"]}}}`)
				return
			}
			w.WriteHeader(201)
			io.WriteString(w, `{"name":"mine","version":"1.0.0","board":"uno","files":1}`)
		default:
			w.WriteHeader(404)
			io.WriteString(w, `{"error":{"code":"not_found","message":"No such package."}}`)
		}
	}))
}

func TestPkgInstallAndPublish(t *testing.T) {
	var published map[string]any
	srv := fakeRegistry(t, &published)
	defer srv.Close()
	keyed := func(args ...string) []string { return append(args, "--api", srv.URL, "--key", "mjd_test_abc") }
	dir := t.TempDir()

	var out bytes.Buffer
	if code := runPkg(keyed("install", "blinker", "--dir", dir), &out); code != 0 {
		t.Fatalf("install failed: %s", out.String())
	}
	if b, _ := os.ReadFile(filepath.Join(dir, "blinker.muj")); !strings.Contains(string(b), "fn twice") {
		t.Error("the module should be written next to the program")
	}
	var lock map[string]string
	b, _ := os.ReadFile(filepath.Join(dir, pkgLockFile))
	if json.Unmarshal(b, &lock) != nil || lock["blinker"] != "1.2.0" {
		t.Errorf("%s should record the version installed: %s", pkgLockFile, b)
	}

	// A changed local copy is never overwritten silently; --force replaces it.
	os.WriteFile(filepath.Join(dir, "blinker.muj"), []byte("# my edits\n"), 0o644)
	out.Reset()
	if runPkg(keyed("install", "blinker", "--dir", dir), &out) == 0 || !strings.Contains(out.String(), "--force") {
		t.Errorf("a differing file should stop the install: %s", out.String())
	}
	if runPkg(keyed("install", "blinker", "--dir", dir, "--force"), &out) != 0 {
		t.Error("--force replaces it")
	}

	// A package naming a path outside the folder writes nothing at all.
	evil := t.TempDir()
	out.Reset()
	if runPkg(keyed("install", "evil", "--dir", evil), &out) == 0 {
		t.Error("a file name with a path must be refused")
	}
	if entries, _ := os.ReadDir(evil); len(entries) != 0 {
		t.Errorf("nothing may be written when one name is bad: %v", entries)
	}
	if _, err := os.Stat(filepath.Join(filepath.Dir(evil), "outside.muj")); err == nil {
		t.Error("a file escaped the folder")
	}

	// Publish sends the library, not the folder's own program or other files.
	lib := filepath.Join(t.TempDir(), "mine")
	os.MkdirAll(lib, 0o755)
	os.WriteFile(filepath.Join(lib, "mine.muj"), []byte("const top = 1\n"), 0o644)
	os.WriteFile(filepath.Join(lib, "sketch.muj"), []byte("board uno\nimport mine\n"), 0o644)
	os.WriteFile(filepath.Join(lib, "notes.txt"), []byte("private"), 0o644)
	out.Reset()
	if code := runPkg(keyed("publish", lib, "--version", "1.0.0"), &out); code != 0 {
		t.Fatalf("publish failed: %s", out.String())
	}
	files, _ := published["files"].(map[string]any)
	if published["name"] != "mine" || len(files) != 1 || files["mine.muj"] == nil {
		t.Errorf("publish should send only mine.muj under the folder's name: %v", published)
	}
	out.Reset()
	if runPkg(keyed("publish", lib, "--version", "9.9.9"), &out) == 0 || !strings.Contains(out.String(), "C303") {
		t.Errorf("a refused publish shows the compiler's errors: %s", out.String())
	}
	if runPkg(keyed("publish", lib, "--version", "one"), &out) != 2 {
		t.Error("a bad version is a usage error")
	}
}

func TestPkgNeedsAKey(t *testing.T) {
	t.Setenv("MUJARD_KEY", "")
	t.Setenv("MUJARD_CONFIG_DIR", t.TempDir())
	var out bytes.Buffer
	if runPkg([]string{"search", "--key", ""}, &out) == 0 {
		t.Error("the registry needs a developer key")
	}
}
