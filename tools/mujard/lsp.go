package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf16"

	"mujard/lang"
)

// mujard lsp: the Mujard language server (plan Phase 13), so VS Code, Neovim, Zed and any other editor that speaks the
// Language Server Protocol check, complete and format .muj/.mjd files with the same compiler as 'mujard check'. It talks
// JSON-RPC on stdin/stdout and touches nothing but the files of the project it is asked about: no network, no key.
//
//	errors and warnings as you type (the whole check, the generator's refusals too), completion, hover, formatting.

func runLsp(args []string, w io.Writer) int {
	if len(args) > 0 && args[0] != "--stdio" { // --stdio is what most editors pass; it is the only transport
		fmt.Fprintln(w, "mujard lsp talks to an editor on stdin/stdout; it takes no arguments (--stdio is accepted).")
		return 2
	}
	return lspServe(os.Stdin, w)
}

type lspMessage struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type lspServer struct {
	out      io.Writer
	docs     map[string]string // open documents: uri -> text (the editor's copy wins over the file on disk)
	info     *lang.EditorInfo
	service  *lang.EditorService // the editor intelligence - shared with the Android .so binding
	shutdown bool
}

// lspServe runs the protocol until 'exit' (or the end of the input); the exit code follows the protocol: 0 after a
// shutdown request, 1 without one.
func lspServe(in io.Reader, out io.Writer) int {
	s := &lspServer{out: out, docs: map[string]string{}, info: lang.NewEditorInfo()}
	s.service = lang.NewEditorService(s.info)
	r := bufio.NewReader(in)
	for {
		body, err := lspRead(r)
		if err != nil {
			if s.shutdown {
				return 0
			}
			return 1
		}
		var m lspMessage
		if json.Unmarshal(body, &m) != nil {
			s.reply(nil, nil, &lspError{-32700, "not JSON"})
			continue
		}
		if m.Method == "exit" {
			if s.shutdown {
				return 0
			}
			return 1
		}
		s.handle(m)
	}
}

// lspRead reads one "Content-Length: n" framed message (headers end with an empty line).
func lspRead(r *bufio.Reader) ([]byte, error) {
	n := -1
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		line = strings.TrimRight(line, "\r\n")
		if line == "" {
			break
		}
		if v, ok := strings.CutPrefix(strings.ToLower(line), "content-length:"); ok {
			n, err = strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				return nil, err
			}
		}
	}
	if n < 0 || n > 64<<20 {
		return nil, fmt.Errorf("bad Content-Length")
	}
	body := make([]byte, n)
	_, err := io.ReadFull(r, body)
	return body, err
}

type lspError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s *lspServer) send(v any) {
	b, _ := json.Marshal(v)
	fmt.Fprintf(s.out, "Content-Length: %d\r\n\r\n%s", len(b), b)
}

func (s *lspServer) reply(id json.RawMessage, result any, e *lspError) {
	if id == nil {
		id = json.RawMessage("null")
	}
	msg := map[string]any{"jsonrpc": "2.0", "id": id}
	if e != nil {
		msg["error"] = e
	} else {
		msg["result"] = result
	}
	s.send(msg)
}

func (s *lspServer) notify(method string, params any) {
	s.send(map[string]any{"jsonrpc": "2.0", "method": method, "params": params})
}

type lspPosition struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type lspRange struct {
	Start lspPosition `json:"start"`
	End   lspPosition `json:"end"`
}

type lspDocParams struct {
	TextDocument struct {
		URI  string `json:"uri"`
		Text string `json:"text"`
	} `json:"textDocument"`
	ContentChanges []struct {
		Text string `json:"text"`
	} `json:"contentChanges"`
	Position lspPosition `json:"position"`
	NewName  string      `json:"newName"`
}

func (s *lspServer) handle(m lspMessage) {
	var p lspDocParams
	_ = json.Unmarshal(m.Params, &p)
	uri := p.TextDocument.URI
	switch m.Method {
	case "initialize":
		s.reply(m.ID, map[string]any{
			"capabilities": map[string]any{
				"textDocumentSync":           1, // the whole text on every change: programs are small
				"completionProvider":         map[string]any{"triggerCharacters": []string{"."}},
				"hoverProvider":              true,
				"documentFormattingProvider": true,
				"definitionProvider":         true,
				"renameProvider":             true,
				"referencesProvider":         true,
			},
			"serverInfo": map[string]any{"name": "mujard", "version": version},
		}, nil)
	case "initialized", "$/cancelRequest", "$/setTrace", "workspace/didChangeConfiguration", "textDocument/didSave":
		// nothing to do
	case "shutdown":
		s.shutdown = true
		s.reply(m.ID, nil, nil)
	case "textDocument/didOpen":
		s.docs[uri] = p.TextDocument.Text
		s.publish(uri)
	case "textDocument/didChange":
		if n := len(p.ContentChanges); n > 0 {
			s.docs[uri] = p.ContentChanges[n-1].Text
		}
		s.publish(uri)
		s.republishSiblings(uri) // a module's change can fix or break the programs that import it
	case "textDocument/didClose":
		delete(s.docs, uri)
		s.notify("textDocument/publishDiagnostics", map[string]any{"uri": uri, "diagnostics": []any{}})
	case "textDocument/completion":
		s.reply(m.ID, s.complete(s.docs[uri], p.Position), nil)
	case "textDocument/hover":
		s.reply(m.ID, s.hover(s.docs[uri], p.Position), nil)
	case "textDocument/formatting":
		s.reply(m.ID, s.format(s.docs[uri]), nil)
	case "textDocument/definition":
		s.reply(m.ID, s.definitions(s.docs[uri], uri, p.Position), nil)
	case "textDocument/rename":
		s.reply(m.ID, s.rename(s.docs[uri], uri, p.Position, p.NewName), nil)
	case "textDocument/references":
		s.reply(m.ID, s.references(s.docs[uri], uri, p.Position), nil)
	default:
		if m.ID != nil { // an unknown request gets an answer; an unknown notification is ignored, as the protocol says
			s.reply(m.ID, nil, &lspError{-32601, "mujard lsp doesn't do " + m.Method})
		}
	}
}

// uriPath is the file a file:// URI names (file:///C:/x/a.muj -> C:\x\a.muj), or "" for any other scheme.
func uriPath(uri string) string {
	u, err := url.Parse(uri)
	if err != nil || u.Scheme != "file" {
		return ""
	}
	p := u.Path
	if runtime.GOOS == "windows" {
		p = strings.TrimPrefix(p, "/")
	}
	return filepath.FromSlash(p)
}

func isMujardName(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	return ext == ".muj" || ext == ".mjd"
}

// project is the open file's folder as the checker sees it: every Mujard source and library description there, with
// the editor's unsaved copies over what is on disk.
func (s *lspServer) project(path string) map[string]string {
	files := map[string]string{}
	dir := filepath.Dir(path)
	if entries, err := os.ReadDir(dir); err == nil {
		for _, e := range entries {
			name := e.Name()
			ext := strings.ToLower(filepath.Ext(name))
			if e.IsDir() || !(ext == ".muj" || ext == ".mjd" || ext == ".mlib" || ext == ".h") {
				continue
			}
			if info, err := e.Info(); err == nil && info.Size() <= 1<<20 {
				if b, err := os.ReadFile(filepath.Join(dir, name)); err == nil {
					files[name] = string(b)
				}
			}
		}
	}
	for uri, text := range s.docs {
		if p := uriPath(uri); p != "" && strings.EqualFold(filepath.Dir(p), dir) {
			files[filepath.Base(p)] = text
		}
	}
	return files
}

func (s *lspServer) publish(uri string) {
	text, open := s.docs[uri]
	path := uriPath(uri)
	if !open || !isMujardName(path) {
		return
	}
	name := filepath.Base(path)
	modules := s.project(path)
	delete(modules, name)
	res := lang.CheckProgramModules(text, modules)
	lines := strings.Split(text, "\n")
	diags := []any{}
	for _, d := range res.Diags {
		if d.File != "" && !strings.EqualFold(d.File, name) {
			continue // a problem inside an imported module: that file's own diagnostics show it
		}
		severity := 1
		if d.IsWarning() {
			severity = 2
		}
		diags = append(diags, map[string]any{
			"range": wordRange(lines, d.Pos), "severity": severity, "code": d.Code, "source": "mujard", "message": d.Message,
		})
	}
	s.notify("textDocument/publishDiagnostics", map[string]any{"uri": uri, "diagnostics": diags})
}

// republishSiblings checks the other open Mujard files of the same folder again.
func (s *lspServer) republishSiblings(uri string) {
	dir := filepath.Dir(uriPath(uri))
	others := make([]string, 0, len(s.docs))
	for u := range s.docs {
		if u != uri && strings.EqualFold(filepath.Dir(uriPath(u)), dir) {
			others = append(others, u)
		}
	}
	sort.Strings(others)
	for _, u := range others {
		s.publish(u)
	}
}

// wordRange is the editor range of the word a diagnostic points at (1-based line and column in the program), in the
// protocol's units: 0-based lines and UTF-16 code units.
func wordRange(lines []string, at lang.Pos) lspRange {
	line := at.Line - 1
	if line < 0 {
		line = 0
	}
	if line >= len(lines) {
		return lspRange{lspPosition{line, 0}, lspPosition{line, 0}}
	}
	runes := []rune(strings.TrimRight(lines[line], "\r"))
	start := at.Col - 1
	if start < 0 {
		start = 0
	}
	if start > len(runes) {
		start = len(runes)
	}
	end := start
	for end < len(runes) && isWordRune(runes[end]) {
		end++
	}
	if end == start && end < len(runes) {
		end++
	}
	return lspRange{lspPosition{line, utf16Len(runes[:start])}, lspPosition{line, utf16Len(runes[:end])}}
}

func isWordRune(r rune) bool { return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r) }

func utf16Len(rs []rune) int { return len(utf16.Encode(rs)) }

// The editor intelligence lives in lang.EditorService (lang/editor_service.go) so the Android .so
// (tools/mujard/androidbind) binds the SAME completion/hover/definition/rename/references the web and
// Studio editors use - one implementation, not a JS or Kotlin copy. The methods below are thin adapters
// from the LSP's lspPosition to the service's LspPos; everything else the service needs is in lang.
func (s *lspServer) complete(text string, pos lspPosition) map[string]any {
	return s.service.Complete(text, lang.LspPos{Line: pos.Line, Character: pos.Character})
}

func (s *lspServer) hover(text string, pos lspPosition) any {
	return s.service.Hover(text, lang.LspPos{Line: pos.Line, Character: pos.Character})
}

func (s *lspServer) definitions(text, uri string, pos lspPosition) []map[string]any {
	return s.service.Definitions(text, uri, lang.LspPos{Line: pos.Line, Character: pos.Character})
}

func (s *lspServer) rename(text, uri string, pos lspPosition, newName string) any {
	return s.service.Rename(text, uri, lang.LspPos{Line: pos.Line, Character: pos.Character}, newName)
}

func (s *lspServer) references(text, uri string, pos lspPosition) any {
	return s.service.References(text, uri, lang.LspPos{Line: pos.Line, Character: pos.Character})
}

func (s *lspServer) format(text string) []any {
	out, diags := lang.Format(text)
	if lang.FirstError(diags) != nil || out == text {
		return []any{}
	}
	lines := strings.Split(text, "\n")
	last := len(lines) - 1
	return []any{map[string]any{
		"range":   lspRange{lspPosition{0, 0}, lspPosition{last, utf16Len([]rune(lines[last]))}},
		"newText": out,
	}}
}

// memberKind maps an EditorMember kind to the LSP CompletionItemKind: method=2, field=5, event=23.
