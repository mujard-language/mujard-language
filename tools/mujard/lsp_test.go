package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// lspSession sends the messages to the server as an editor would and returns everything it sent back.
func lspSession(t *testing.T, msgs ...map[string]any) (code int, replies []map[string]any) {
	t.Helper()
	var in bytes.Buffer
	for _, m := range msgs {
		m["jsonrpc"] = "2.0"
		b, _ := json.Marshal(m)
		fmt.Fprintf(&in, "Content-Length: %d\r\n\r\n%s", len(b), b)
	}
	var out bytes.Buffer
	code = lspServe(&in, &out)
	r := bufio.NewReader(&out)
	for {
		body, err := lspRead(r)
		if err != nil {
			break
		}
		var m map[string]any
		if err := json.Unmarshal(body, &m); err != nil {
			t.Fatalf("the server sent something that isn't JSON: %s", body)
		}
		replies = append(replies, m)
	}
	return code, replies
}

func fileURI(p string) string {
	p = filepath.ToSlash(p)
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	return (&url.URL{Scheme: "file", Path: p}).String()
}

func reply(replies []map[string]any, id int) map[string]any {
	for _, r := range replies {
		if v, ok := r["id"].(float64); ok && int(v) == id {
			return r
		}
	}
	return nil
}

func diagnosticsFor(replies []map[string]any, uri string) []any {
	var last []any
	for _, r := range replies {
		if r["method"] == "textDocument/publishDiagnostics" {
			p := r["params"].(map[string]any)
			if p["uri"] == uri {
				last = p["diagnostics"].([]any)
			}
		}
	}
	return last
}

func TestLspSession(t *testing.T) {
	dir := t.TempDir()
	program := filepath.Join(dir, "sketch.muj")
	mod := filepath.Join(dir, "motors.muj")
	// On disk the module has no 'top': only the editor's unsaved copy does, and the program must be checked against that.
	os.WriteFile(mod, []byte("const other = 1\n"), 0o644)
	bad := "board uno\nled = pin 13 as output\nevery 500 ms:\n    led.togle()\n"
	good := "board uno\nimport motors\nat start:\n    print motors.top\n"
	messy := "board uno\nled = pin 13 as output\nevery   500 ms:\n      led.toggle( )\n"
	mainURI, modURI := fileURI(program), fileURI(mod)
	open := func(uri, text string) map[string]any {
		return map[string]any{"method": "textDocument/didOpen", "params": map[string]any{"textDocument": map[string]any{"uri": uri, "languageId": "mujard", "version": 1, "text": text}}}
	}
	change := func(uri, text string) map[string]any {
		return map[string]any{"method": "textDocument/didChange", "params": map[string]any{"textDocument": map[string]any{"uri": uri, "version": 2}, "contentChanges": []any{map[string]any{"text": text}}}}
	}
	at := func(id int, method, uri string, line, ch int) map[string]any {
		return map[string]any{"id": id, "method": method, "params": map[string]any{"textDocument": map[string]any{"uri": uri}, "position": map[string]any{"line": line, "character": ch}}}
	}
	code, replies := lspSession(t,
		map[string]any{"id": 1, "method": "initialize", "params": map[string]any{"capabilities": map[string]any{}}},
		map[string]any{"method": "initialized", "params": map[string]any{}},
		open(mainURI, bad),
		at(2, "textDocument/completion", mainURI, 3, 8), // after "    led."
		at(3, "textDocument/hover", mainURI, 2, 2),      // on "every"
		open(modURI, "const top = 800\n"),
		change(mainURI, good),
		change(mainURI, messy),
		map[string]any{"id": 4, "method": "textDocument/formatting", "params": map[string]any{"textDocument": map[string]any{"uri": mainURI}, "options": map[string]any{"tabSize": 4, "insertSpaces": true}}},
		map[string]any{"id": 5, "method": "nonsense/request", "params": map[string]any{}},
		map[string]any{"id": 6, "method": "shutdown"},
		map[string]any{"method": "exit"},
	)
	if code != 0 {
		t.Errorf("exit after shutdown should be 0, got %d", code)
	}
	hello := reply(replies, 1)
	caps, _ := hello["result"].(map[string]any)["capabilities"].(map[string]any)
	if caps["textDocumentSync"] != float64(1) || caps["hoverProvider"] != true || caps["documentFormattingProvider"] != true {
		t.Errorf("initialize should announce full sync, hover and formatting: %v", hello)
	}

	// The first diagnostics are for the typo, on line 4 (0-based 3), underlining the misspelt member.
	var first []any
	for _, r := range replies {
		if r["method"] == "textDocument/publishDiagnostics" && r["params"].(map[string]any)["uri"] == mainURI {
			first = r["params"].(map[string]any)["diagnostics"].([]any)
			break
		}
	}
	if len(first) == 0 {
		t.Fatal("a misspelt member should give a diagnostic")
	}
	d := first[0].(map[string]any)
	rng := d["range"].(map[string]any)
	start := rng["start"].(map[string]any)
	if d["code"] != "C308" || d["severity"] != float64(1) || d["source"] != "mujard" || start["line"] != float64(3) {
		t.Errorf("want C308 as an error on line 3: %v", d)
	}

	items := reply(replies, 2)["result"].(map[string]any)["items"].([]any)
	labels := map[string]bool{}
	for _, it := range items {
		labels[it.(map[string]any)["label"].(string)] = true
	}
	if !labels["toggle"] || labels["every"] {
		t.Errorf("after 'led.' completion should offer members (toggle), not keywords: %d items", len(items))
	}
	if h := reply(replies, 3)["result"]; h == nil || !strings.Contains(fmt.Sprint(h), "every") {
		t.Errorf("hover on 'every' should describe it: %v", h)
	}

	// After the module was opened with 'top' and the program changed to use it, the program is clean: the unsaved
	// module counted, not the file on disk.
	var afterGood []any
	seenGood := false
	for _, r := range replies {
		if r["method"] == "textDocument/publishDiagnostics" && r["params"].(map[string]any)["uri"] == mainURI {
			if seenGood {
				break
			}
			afterGood = r["params"].(map[string]any)["diagnostics"].([]any)
			if len(afterGood) == 0 {
				seenGood = true
			}
		}
	}
	if !seenGood {
		t.Errorf("the program importing the unsaved module should check clean at some point")
	}

	edits, _ := reply(replies, 4)["result"].([]any)
	if len(edits) != 1 || !strings.Contains(edits[0].(map[string]any)["newText"].(string), "every 500 ms:\n    led.toggle()") {
		t.Errorf("formatting should give one edit with the standard layout: %v", edits)
	}
	if e := reply(replies, 5)["error"]; e == nil {
		t.Error("an unknown request gets an error answer")
	}
	if diagnosticsFor(replies, modURI) == nil {
		t.Error("the opened module is checked too")
	}
}

func TestLspExitWithoutShutdown(t *testing.T) {
	if code, _ := lspSession(t, map[string]any{"method": "exit"}); code != 1 {
		t.Errorf("exit without shutdown should be 1, got %d", code)
	}
}

func TestLspTakesNoArguments(t *testing.T) {
	var out bytes.Buffer
	if runLsp([]string{"--tcp"}, &out) != 2 {
		t.Error("only stdio is offered")
	}
}
