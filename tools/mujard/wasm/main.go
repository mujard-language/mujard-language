//go:build js && wasm

// The Mujard language inside the browser: the same lexer, parser and checker 'mujard check' runs, compiled to
// WebAssembly so the web editor checks .muj/.mjd programs as you type - offline, before anything is built.
// Shipped to the web app as its wasm bundle; wasm_exec.js comes from the Go toolchain.
//
//	GOOS=js GOARCH=wasm go build -o mujard.wasm ./wasm
package main

import (
	"encoding/json"
	"syscall/js"

	"mujard/lang"
)

type diagOut struct {
	Code    string `json:"code"`
	File    string `json:"file,omitempty"`
	Line    int    `json:"line"`
	Col     int    `json:"col"`
	Warning bool   `json:"warning"`
	Message string `json:"message"`
}

func diagsJSON(diags []lang.Diagnostic) string {
	out := make([]diagOut, 0, len(diags))
	for _, d := range diags {
		out = append(out, diagOut{Code: d.Code, File: d.File, Line: d.Pos.Line, Col: d.Pos.Col, Warning: d.IsWarning(), Message: d.Message})
	}
	b, _ := json.Marshal(out)
	return string(b)
}

// mujardCheck(source, modulesJson) -> JSON [{code, file?, line, col, warning, message}]. modulesJson is an
// optional {"motors.muj": "source"} object so 'import motors' resolves against the project's other files.
func check(_ js.Value, args []js.Value) any {
	if len(args) == 0 || args[0].Type() != js.TypeString {
		return "[]"
	}
	modules := map[string]string{}
	if len(args) > 1 && args[1].Type() == js.TypeString {
		_ = json.Unmarshal([]byte(args[1].String()), &modules)
	}
	_, diags := lang.CheckSourceModules(args[0].String(), modules)
	return diagsJSON(diags)
}

// mujardSecrets(source, modulesJson) -> JSON ["HOME_WIFI_SSID", ...]: the build-secret names the program needs
// (names only - the editor shows which the project still lacks). [] when it doesn't check clean.
func secrets(_ js.Value, args []js.Value) any {
	if len(args) == 0 || args[0].Type() != js.TypeString {
		return "[]"
	}
	modules := map[string]string{}
	if len(args) > 1 && args[1].Type() == js.TypeString {
		_ = json.Unmarshal([]byte(args[1].String()), &modules)
	}
	names := lang.CheckProgramModules(args[0].String(), modules).Secrets
	if names == nil {
		names = []string{}
	}
	b, _ := json.Marshal(names)
	return string(b)
}

// mujardFormat(source) -> JSON {ok, text, diags}; text is the source in 'mujard fmt' layout, "" when the file
// can't be read (diags then says why).
func format(_ js.Value, args []js.Value) any {
	if len(args) == 0 || args[0].Type() != js.TypeString {
		return `{"ok":false,"text":""}`
	}
	text, diags := lang.Format(args[0].String())
	b, _ := json.Marshal(map[string]any{"ok": len(diags) == 0, "text": text, "diags": json.RawMessage(diagsJSON(diags))})
	return string(b)
}

// mujardInfo() -> the EditorInfo JSON (keywords, units, built-ins, member tables, boards).
func info(_ js.Value, _ []js.Value) any {
	b, _ := json.Marshal(lang.NewEditorInfo())
	return string(b)
}

// The editor intelligence - completion/hover/definition/references/rename - from lang.EditorService, the
// same service the LSP and the Android .so run, so web/Studio's CodeMirror can offer the receiver-scoped
// completion and navigation the phone gets, not just the word lists.

// mujardComplete(source, line, col) -> {isIncomplete, items:[{label,kind,detail}]} - receiver-scoped after 'x.'.
func complete(_ js.Value, args []js.Value) any {
	if len(args) < 3 || args[0].Type() != js.TypeString {
		return `{"isIncomplete":false,"items":[]}`
	}
	out := lang.NewEditorService(lang.NewEditorInfo()).Complete(args[0].String(), lang.LspPos{Line: args[1].Int(), Character: args[2].Int()})
	b, _ := json.Marshal(out)
	return string(b)
}

// mujardHover(source, line, col) -> hover {contents:{kind,value}} or null.
func hover(_ js.Value, args []js.Value) any {
	if len(args) < 3 || args[0].Type() != js.TypeString {
		return nil
	}
	out := lang.NewEditorService(lang.NewEditorInfo()).Hover(args[0].String(), lang.LspPos{Line: args[1].Int(), Character: args[2].Int()})
	b, _ := json.Marshal(out)
	return string(b)
}

// mujardDefinition(source, uri, line, col) -> JSON Locations [{uri,range}].
func definition(_ js.Value, args []js.Value) any {
	if len(args) < 4 || args[0].Type() != js.TypeString {
		return "[]"
	}
	out := lang.NewEditorService(lang.NewEditorInfo()).Definitions(args[0].String(), args[1].String(), lang.LspPos{Line: args[2].Int(), Character: args[3].Int()})
	b, _ := json.Marshal(out)
	return string(b)
}

// mujardReferences(source, uri, line, col) -> JSON Locations.
func references(_ js.Value, args []js.Value) any {
	if len(args) < 4 || args[0].Type() != js.TypeString {
		return "[]"
	}
	out := lang.NewEditorService(lang.NewEditorInfo()).References(args[0].String(), args[1].String(), lang.LspPos{Line: args[2].Int(), Character: args[3].Int()})
	b, _ := json.Marshal(out)
	return string(b)
}

// mujardRename(source, uri, line, col, newName) -> a WorkspaceEdit JSON.
func rename(_ js.Value, args []js.Value) any {
	if len(args) < 5 || args[0].Type() != js.TypeString {
		return `{}`
	}
	out := lang.NewEditorService(lang.NewEditorInfo()).Rename(args[0].String(), args[1].String(), lang.LspPos{Line: args[2].Int(), Character: args[3].Int()}, args[4].String())
	b, _ := json.Marshal(out)
	return string(b)
}

// mujardDescribe(name, header, src) -> the '.mlib' JSON the 'use "name"' line binds. src is a pasted C++ header
// (the auto-wrapper - a .h in, a LibDesc out) - the same parseHeader as 'mujard describe' for one .h.
// Returns {"error":"..."} when nothing describable is in the header.
func describe(_ js.Value, args []js.Value) any {
	if len(args) < 3 || args[0].Type() != js.TypeString || args[2].Type() != js.TypeString {
		return `{"error":"mujardDescribe needs the library name, its .h name and the header text"}`
	}
	out, err := lang.DescribeHeaderSourceJSON(args[0].String(), args[1].String(), args[2].String())
	if err != nil {
		b, _ := json.Marshal(map[string]string{"error": err.Error()})
		return string(b)
	}
	return out
}

func main() {
	js.Global().Set("mujardCheck", js.FuncOf(check))
	js.Global().Set("mujardSecrets", js.FuncOf(secrets))
	js.Global().Set("mujardFormat", js.FuncOf(format))
	js.Global().Set("mujardInfo", js.FuncOf(info))
	js.Global().Set("mujardComplete", js.FuncOf(complete))
	js.Global().Set("mujardHover", js.FuncOf(hover))
	js.Global().Set("mujardDefinition", js.FuncOf(definition))
	js.Global().Set("mujardReferences", js.FuncOf(references))
	js.Global().Set("mujardRename", js.FuncOf(rename))
	js.Global().Set("mujardDescribe", js.FuncOf(describe))
	select {}
}
