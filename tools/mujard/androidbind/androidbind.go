// Package androidbind exposes the Mujard checker to the Android app as a gomobile-bindable package, the
// same lexer, parser and checker 'mujard check' runs - the phone gets the authoritative diagnostics the
// web editor gets from mujard.wasm, not a Kotlin reimplementation. Build it into a .aar/.so the app
// loads: 'gomobile bind -target android -o mujard.aar ./androidbind' (needs the Android NDK).
package androidbind

import (
	"encoding/json"

	"mujard/lang"
)

// diagOut is one diagnostic, matching the JSON the wasm checker (tools/mujard/wasm) returns so web,
// Studio and Android all consume the same schema.
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

// Check runs the Mujard checker on source and returns the diagnostics as a JSON string
// [{code, file?, line, col, warning, message}] - the same shape mujard.wasm's mujardCheck produces.
// modulesJson is an optional {"motors.muj": "source"} object so 'import motors' resolves against the
// project's other files; pass "" when there are none.
func Check(source, modulesJson string) string {
	modules := map[string]string{}
	if modulesJson != "" {
		_ = json.Unmarshal([]byte(modulesJson), &modules)
	}
	_, diags := lang.CheckSourceModules(source, modules)
	return diagsJSON(diags)
}

// CompletionList runs the shared Mujard editor intelligence at (line, col) - 0-based like the LSP - and
// returns the completion JSON ({isIncomplete, items:[{label,kind,detail}]}). Receiver-scoped after 'x.':
// 'led = pin 13 as output' -> led. gives the output members, the same as the web/Studio LSP.
func CompletionList(source string, line, col int) string {
	out := lang.NewEditorService(lang.NewEditorInfo()).Complete(source, lang.LspPos{Line: line, Character: col})
	b, _ := json.Marshal(out)
	return string(b)
}

// HoverText is the hover markdown at (line, col) - the word's definition(s) - or "" when the word is unknown.
func HoverText(source string, line, col int) string {
	out := lang.NewEditorService(lang.NewEditorInfo()).Hover(source, lang.LspPos{Line: line, Character: col})
	if out == nil {
		return ""
	}
	b, _ := json.Marshal(out)
	return string(b)
}

// Definitions at (line, col): where the word is declared, as JSON Locations [{uri, range}].
func Definitions(source, uri string, line, col int) string {
	out := lang.NewEditorService(lang.NewEditorInfo()).Definitions(source, uri, lang.LspPos{Line: line, Character: col})
	b, _ := json.Marshal(out)
	return string(b)
}

// References is every use of the word at (line, col) in source, as JSON Locations.
func References(source, uri string, line, col int) string {
	out := lang.NewEditorService(lang.NewEditorInfo()).References(source, uri, lang.LspPos{Line: line, Character: col})
	b, _ := json.Marshal(out)
	return string(b)
}

// RenameEdits is a WorkspaceEdit JSON renaming the word at (line, col) to newName in source.
func RenameEdits(source, uri string, line, col int, newName string) string {
	out := lang.NewEditorService(lang.NewEditorInfo()).Rename(source, uri, lang.LspPos{Line: line, Character: col}, newName)
	b, _ := json.Marshal(out)
	return string(b)
}

// Describe wraps one pasted C++ header's text into the '.mlib' JSON a 'use "name"' line binds - the
// auto-wrapper: a .h in, a LibDesc out, the same parseHeader as 'mujard describe' and the wasm
// mujardDescribe, so the phone wraps a header exactly like the browser and the CLI. name is what
// 'use "name"' looks up; header is the '#include <header>' the generated code needs.
// Returns {"error":"..."} when nothing describable is in the header.
func Describe(name, header, src string) string {
	out, err := lang.DescribeHeaderSourceJSON(name, header, src)
	if err != nil {
		b, _ := json.Marshal(map[string]string{"error": err.Error()})
		return string(b)
	}
	return out
}

// Secrets returns the build-secret names the program needs as a JSON string
// ["HOME_WIFI_SSID", ...] - names only, never values, so the app can flag which the project still lacks.
// [] when the program doesn't check clean.
func Secrets(source, modulesJson string) string {
	modules := map[string]string{}
	if modulesJson != "" {
		_ = json.Unmarshal([]byte(modulesJson), &modules)
	}
	res := lang.CheckProgramModules(source, modules)
	if len(res.Secrets) == 0 {
		return "[]"
	}
	b, _ := json.Marshal(res.Secrets)
	return string(b)
}
