// EditorService is the Mujard editor intelligence - completion, hover, definition, rename and references
// - over a parsed program, independent of any transport. The LSP server calls it (lsp.go keeps only the
// JSON-RPC plumbing); the Android .so binds the same functions via tools/mujard/androidbind, so the phone
// gets the same receiver-scoped completion and navigation the web/Studio editors get - one implementation.
package lang

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"
	"unicode/utf16"
)

func isWordRune(r rune) bool { return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r) }

func memberKind(m EditorMember) int {
	switch m.Kind {
	case "field":
		return 5
	case "event":
		return 23
	}
	return 2
}

// LspPos is a 0-based (line, character) cursor position, character counted in UTF-16 code units like the
// LSP. (The checker uses Pos, 1-based byte columns - the editors use this LSP convention.)
type LspPos struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

// LspRange is an LSP span: start and end positions.
type LspRange struct {
	Start LspPos `json:"start"`
	End   LspPos `json:"end"`
}

// EditorService provides completion/hover/navigation over a Mujard program against the word + member +
// unit lists in NewEditorInfo. Stateless apart from info - a service per document, like the LSP does.
type EditorService struct {
	info *EditorInfo
}

// NewEditorService builds a service against the language's editor metadata.
func NewEditorService(info *EditorInfo) *EditorService {
	return &EditorService{info: info}
}

// wordBefore is the word ending at the cursor and whether a '.' comes right before it (a member is wanted).
func wordBefore(text string, pos LspPos) (word string, dot bool, line string) {
	lines := strings.Split(text, "\n")
	if pos.Line < 0 || pos.Line >= len(lines) {
		return "", false, ""
	}
	line = strings.TrimRight(lines[pos.Line], "\r")
	u := utf16.Encode([]rune(line))
	if pos.Character < len(u) {
		u = u[:pos.Character]
	}
	before := []rune(string(utf16.Decode(u)))
	i := len(before)
	for i > 0 && isWordRune(before[i-1]) {
		i--
	}
	return string(before[i:]), i > 0 && before[i-1] == '.', line
}

// wordAt is the whole word under the cursor: the part before it and the part after.
func wordAt(text string, pos LspPos) string {
	word, _, line := wordBefore(text, pos)
	u := utf16.Encode([]rune(line))
	at := len(utf16.Decode(u[:minInt(pos.Character, len(u))])) // the cursor as a rune index
	runes := []rune(line)
	for ; at < len(runes) && isWordRune(runes[at]); at++ {
		word += string(runes[at])
	}
	return word
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Complete is the completion list at pos: after 'x.', the members of x's type (a 'led = pin 13 as output'
// gives the output members, not every object's); otherwise keywords/statements/globals/units.
func (s *EditorService) Complete(text string, pos LspPos) map[string]any {
	prefix, dot, _ := wordBefore(text, pos)
	items := []map[string]any{}
	seen := map[string]bool{}
	add := func(label string, kind int, detail string) {
		if seen[label] || !strings.HasPrefix(strings.ToLower(label), strings.ToLower(prefix)) {
			return
		}
		seen[label] = true
		items = append(items, map[string]any{"label": label, "kind": kind, "detail": detail})
	}
	if dot {
		// receiverPath is the dotted name before the '.' ('a.b' for 'a.b.|'); each hop's EditorMember
		// Type resolves the next, so 't.split.' opens the list members. When a hop can't be resolved
		// it falls back to every object's members.
		if path := receiverPath(text, pos); len(path) > 0 {
			if tn := s.pathType(text, path); tn != "" {
				ms := s.info.Members[tn]
				if len(ms) == 0 {
					ms = s.info.Members[strings.ToLower(tn)]
				}
				if len(ms) > 0 {
					for _, m := range ms {
						add(m.Name, memberKind(m), m.Detail)
					}
					return map[string]any{"isIncomplete": false, "items": items}
				}
			}
		}
		owners := make([]string, 0, len(s.info.Members))
		for o := range s.info.Members {
			owners = append(owners, o)
		}
		sort.Strings(owners)
		for _, o := range owners {
			for _, m := range s.info.Members[o] {
				add(m.Name, memberKind(m), m.Detail)
			}
		}
		return map[string]any{"isIncomplete": false, "items": items}
	}
	for _, k := range s.info.Keywords {
		add(k, 14, "keyword")
	}
	for _, k := range s.info.Statements {
		add(k, 14, "statement")
	}
	for _, g := range s.info.Globals {
		kind := 3 // function
		switch g.Kind {
		case "object":
			kind = 6
		case "constant":
			kind = 21
		}
		detail := g.Detail
		if g.Lib != "" {
			detail += "  (use mujard." + g.Lib + ")"
		}
		add(g.Name, kind, detail)
	}
	for _, u := range s.info.Units {
		add(u, 11, "unit")
	}
	return map[string]any{"isIncomplete": false, "items": items}
}

// Hover is the hover text at pos: the word's definition(s) as markdown, or nil when the word is unknown.
func (s *EditorService) Hover(text string, pos LspPos) any {
	word := wordAt(text, pos)
	if word == "" {
		return nil
	}
	var lines []string
	for _, g := range s.info.Globals {
		if g.Name == word {
			l := "**" + g.Name + "** - " + g.Kind + ": `" + g.Detail + "`"
			if g.Lib != "" {
				l += " (needs `use mujard." + g.Lib + "`)"
			}
			lines = append(lines, l)
		}
	}
	owners := make([]string, 0, len(s.info.Members))
	for o := range s.info.Members {
		owners = append(owners, o)
	}
	sort.Strings(owners)
	for _, o := range owners {
		for _, m := range s.info.Members[o] {
			if m.Name == word {
				lines = append(lines, "**"+o+"."+m.Name+"** - "+m.Kind+": `"+m.Detail+"`")
			}
		}
	}
	for _, k := range s.info.Keywords {
		if k == word {
			lines = append(lines, "**"+word+"** - a Mujard keyword")
		}
	}
	if len(lines) == 0 {
		return nil
	}
	if len(lines) > 8 {
		lines = append(lines[:8], fmt.Sprintf("... and %d more", len(lines)-8))
	}
	return map[string]any{"contents": map[string]any{"kind": "markdown", "value": strings.Join(lines, "\n\n")}}
}

// lspDeclWords: the words that introduce a named thing in Mujard ('pin led', 'fn go()', 'let total = 0').
var lspDeclWords = regexp.MustCompile(`^\s*(?:fn|pin|let|var|dev|device|const|state|type|flag|every|on|when|module|sensor|block)\s+`)

// lspColRange: the utf-16 columns of a whole-word hit at byte offset at..at+width on line i.
func lspColRange(line string, i, at, width int) LspRange {
	runes := []rune(line[:at])
	start := utf16Len(runes)
	return LspRange{LspPos{i, start}, LspPos{i, start + utf16Len([]rune(line[at:at+width]))}}
}

func utf16Len(rs []rune) int { return len(utf16.Encode(rs)) }

// Definitions: where the word under the cursor is declared in this doc - scanned locally because the
// language's internal symbol table isn't exported. A decl line is '<decl word> NAME' at any indent.
func (s *EditorService) Definitions(text, uri string, pos LspPos) []map[string]any {
	word := wordAt(text, pos)
	if word == "" {
		return nil
	}
	var out []map[string]any
	// A device decl like 'led = pin 13 as output' names the word before the '='; a decl-word form like
	// 'fn name'/'var name' puts it after the keyword. Both mark the word's declaration.
	nameEq := regexp.MustCompile(`^\s*` + regexp.QuoteMeta(word) + `\s*=`)
	for i, raw := range strings.Split(text, "\n") {
		line := strings.TrimRight(raw, "\r")
		if loc := nameEq.FindStringIndex(line); loc != nil {
			// The word is at line-start (after indent) before '=' - its declaration.
			at := strings.Index(line, word)
			out = append(out, map[string]any{"uri": uri, "range": lspColRange(line, i, at, len(word))})
			continue
		}
		kw := lspDeclWords.FindStringIndex(line)
		if kw == nil {
			continue
		}
		idx := strings.Index(line[kw[1]:], word)
		if idx < 0 {
			continue
		}
		at := kw[1] + idx
		if end := at + len(word); end <= len(line) && (end == len(line) || !isWordRune(rune(line[end]))) && (at == 0 || !isWordRune(rune(line[at-1]))) {
			out = append(out, map[string]any{"uri": uri, "range": lspColRange(line, i, at, len(word))})
		}
	}
	return out
}

// Rename: a WorkspaceEdit changing every whole-word occurrence of the name under the cursor in this doc.
func (s *EditorService) Rename(text, uri string, pos LspPos, newName string) any {
	empty := map[string]any{"changes": map[string]any{uri: []any{}}}
	word := wordAt(text, pos)
	if word == "" || newName == "" {
		return empty
	}
	wordRe := regexp.MustCompile(`\b` + regexp.QuoteMeta(word) + `\b`)
	var edits []any
	for i, raw := range strings.Split(text, "\n") {
		line := strings.TrimRight(raw, "\r")
		for _, m := range wordRe.FindAllStringIndex(line, -1) {
			edits = append(edits, map[string]any{"range": lspColRange(line, i, m[0], m[1]-m[0]), "newText": newName})
		}
	}
	return map[string]any{"changes": map[string]any{uri: edits}}
}

// References answers "where is this symbol used" - every whole-word occurrence in the open doc, like
// rename's scan but returning Locations, not edits.
func (s *EditorService) References(text, uri string, pos LspPos) any {
	word := wordAt(text, pos)
	if word == "" {
		return []any{}
	}
	wordRe := regexp.MustCompile(`\b` + regexp.QuoteMeta(word) + `\b`)
	var locs []any
	for i, raw := range strings.Split(text, "\n") {
		line := strings.TrimRight(raw, "\r")
		for _, m := range wordRe.FindAllStringIndex(line, -1) {
			locs = append(locs, map[string]any{"uri": uri, "range": lspColRange(line, i, m[0], m[1]-m[0])})
		}
	}
	return locs
}

var memberPathRe = regexp.MustCompile(`([A-Za-z_][A-Za-z0-9_]*(?:\.[A-Za-z_][A-Za-z0-9_]*)*)\.$`)

// receiverPath is the dotted name ending in the '.' before the cursor - ['a','b'] for 'a.b.|'.
func receiverPath(text string, pos LspPos) []string {
	lines := strings.Split(text, "\n")
	if pos.Line < 0 || pos.Line >= len(lines) {
		return nil
	}
	line := strings.TrimRight(lines[pos.Line], "\r")
	u := utf16.Encode([]rune(line))
	if pos.Character < len(u) {
		u = u[:pos.Character]
	}
	m := memberPathRe.FindStringSubmatch(string(utf16.Decode(u)))
	if m == nil {
		return nil
	}
	return strings.Split(m[1], ".")
}

// pathType resolves a dotted path: the first word through receiverType (a declared name or built-in
// object), each next word through the Type its member advertises. Empty when any hop can't resolve.
func (s *EditorService) pathType(text string, path []string) string {
	cur := s.receiverType(text, path[0])
	for _, name := range path[1:] {
		if cur == "" {
			return ""
		}
		ms := s.info.Members[cur]
		if len(ms) == 0 {
			ms = s.info.Members[strings.ToLower(cur)]
		}
		cur = ""
		for _, m := range ms {
			if m.Name == name {
				cur = m.Type
				break
			}
		}
	}
	return cur
}

// receiverType resolves the type of the word before '.' - 'led' in 'led = pin 13 as output' is an
// 'output', so 'led.' offers the output members, not every object's. Empty when x can't be resolved.
func (s *EditorService) receiverType(text, x string) string {
	if x == "" {
		return ""
	}
	// A built-in object named x (store, board, serial, wifi, ...) is already its own class.
	lx := strings.ToLower(x)
	if _, ok := s.info.Members[lx]; ok {
		return lx
	}
	f, _ := CheckSource(text)
	if f == nil {
		return ""
	}
	for _, it := range f.Items {
		if tn := s.declType(it, x); tn != "" {
			return tn
		}
	}
	return ""
}

// declType resolves x's declared type when item declares it - a DeviceDecl's role, a VarDecl's type, a
// fn's param type, or an object-style constructor. Empty when the item doesn't declare x or x's type
// isn't a builtin object (an unresolvable chain hop falls back to all members).
func (s *EditorService) declType(item Node, x string) string {
	switch d := item.(type) {
	case *DeviceDecl:
		if d.Name == x {
			if d.Role != "" {
				return strings.ToLower(d.Role) // pin 13 as output -> output
			}
			if t := ctorName(d.Object); t != "" {
				return t // led = Servo(pin 6) -> servo
			}
		}
	case *VarDecl:
		if d.Name == x && d.Type != nil {
			return strings.ToLower(d.Type.Name) // var/const/saved t: timer -> timer
		}
	case *FnDecl:
		for _, p := range d.Params {
			if p.Name == x && p.Type != nil {
				return strings.ToLower(p.Type.Name)
			}
		}
	}
	return ""
}

// ctorName is the class an object-style declaration names: Servo(pin 6) -> servo, timer 2 s -> timer.
func ctorName(n Node) string {
	switch e := n.(type) {
	case *Ident:
		return strings.ToLower(e.Name)
	case *Call:
		return ctorName(e.Fn)
	}
	return ""
}
