package lang

import (
	"sort"
	"strings"
)

// The language as an editor needs it: keywords, units, the built-in names, every built-in object's members and the
// board table. The web editor's WASM build hands this to the completion and hover code; the Android app reads the
// same JSON from the file docs/lang/editor-info.json so both editors describe the language identically.
type EditorInfo struct {
	// LanguageVersion is the language this word list describes (LanguageVersion()): an app compares it with the server's.
	LanguageVersion int      `json:"languageVersion"`
	Keywords        []string `json:"keywords"`
	// Statement words that aren't reserved keywords: the ones the parser accepts at the start of a line.
	Statements []string                  `json:"statements"`
	Units      []string                  `json:"units"`
	Globals    []EditorGlobal            `json:"globals"`
	Members    map[string][]EditorMember `json:"members"`
	Boards     []EditorBoard             `json:"boards"`
}

// EditorGlobal is one built-in name (a function, an object, or a constant).
type EditorGlobal struct {
	Name   string `json:"name"`
	Kind   string `json:"kind"` // "function", "object" or "constant"
	Detail string `json:"detail"`
	Lib    string `json:"lib,omitempty"`  // needs 'use mujard.<lib>' first; empty = always there
	Test   bool   `json:"test,omitempty"` // a testing built-in (advance, pretend, ...)
}

// EditorMember is one member of a built-in object or of a plain value (the "text", "number" and "list" classes).
type EditorMember struct {
	Name   string `json:"name"`
	Kind   string `json:"kind"` // "method", "field" or "event" (usable after 'on')
	Detail string `json:"detail"`
	// Type is the member's result as a Members key ("text", "number", "list", an object or enum name):
	// it lets completion resolve 'a.b.' chains one hop at a time. Empty = not a class with members.
	Type string `json:"type,omitempty"`
}

// EditorBoard is one board's pin table for completion and hover.
type EditorBoard struct {
	Name    string         `json:"name"`
	Display string         `json:"display"`
	Pins    map[string]int `json:"pins"`   // names as written in programs: LED_BUILTIN, A0, ...
	Valid   []int          `json:"valid"`  // usable pin numbers
	PWM     []int          `json:"pwm"`    // pins that can do pwm
	Analog  []int          `json:"analog"` // pins an 'analog' device may read
}

func typeShort(t *Type) string {
	if t == nil {
		return "any"
	}
	s := "?"
	switch t.K {
	case TAny:
		s = "any"
	case TVoid:
		s = "void"
	case TInt:
		s = "int"
	case TFloat:
		s = "float"
	case TBool:
		s = "bool"
	case TText:
		s = "text"
	case TQty:
		s = t.Dim
	case TList:
		s = "list"
	case TRange:
		s = "range"
	case TStruct, TEnum, TDevice, TObject:
		s = t.Name
	case TNone:
		s = "none"
	}
	if t.Opt {
		s += "?"
	}
	return s
}

// sig renders a function's call shape for a completion detail: "blink(rate, for?, on?)" or "name(...)" for
// unchecked (C++ library) functions.
func sig(name string, s *fnSig) string {
	if s == nil {
		return name
	}
	if s.params == nil {
		return name + "(...)"
	}
	var b strings.Builder
	b.WriteString(name + "(")
	for i, p := range s.params {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(p.name)
		if p.optional {
			b.WriteString("?")
		}
	}
	b.WriteString(")")
	if s.result != nil && s.result.K != TVoid {
		b.WriteString(" -> " + typeShort(s.result))
	}
	return b.String()
}

// memberTypeName maps a member's result type to the Members key completion should open next:
// text/number/list for the plain-value classes, the enum's own name (a Members key too), or the
// lowercase object/device name. Anything else has no members to offer.
func memberTypeName(t *Type) string {
	if t == nil {
		return ""
	}
	switch t.K {
	case TText:
		return "text"
	case TInt, TFloat, TQty:
		return "number"
	case TList:
		return "list"
	case TEnum, TStruct:
		return t.Name
	case TObject, TDevice:
		return strings.ToLower(t.Name)
	}
	return ""
}

func memberOf(name string, m *member) EditorMember {
	switch {
	case m.event:
		d := name
		if m.eventArg != nil {
			d += "(" + typeShort(m.eventArg) + ")"
		}
		return EditorMember{name, "event", d, ""} // an event isn't a value you can chain from
	case m.field != nil:
		return EditorMember{name, "field", name + ": " + typeShort(m.field), memberTypeName(m.field)}
	default:
		var res *Type
		if m.method != nil {
			res = m.method.result
		}
		return EditorMember{name, "method", sig(name, m.method), memberTypeName(res)}
	}
}

// NewEditorInfo is the editor-facing description of the whole language.
func NewEditorInfo() *EditorInfo {
	kw := make([]string, 0, len(Keywords))
	for k := range Keywords {
		kw = append(kw, k)
	}
	sort.Strings(kw)

	un := make([]string, 0, len(Units))
	for u := range Units {
		un = append(un, u)
	}
	sort.Strings(un)

	gs := make([]EditorGlobal, 0, len(globals))
	for name, g := range globals {
		e := EditorGlobal{Name: name, Lib: g.lib}
		switch {
		case g.class != "":
			e.Kind = "object"
			e.Detail = g.class
			if g.fn != nil {
				e.Detail = sig(name, g.fn) + " - the " + g.class + " object"
			}
		case g.value != nil:
			e.Kind = "constant"
			e.Detail = typeShort(g.value)
		default:
			e.Kind = "function"
			e.Detail = sig(name, g.fn)
		}
		if g.fn != nil && g.fn.testOnly {
			e.Test = true
		}
		gs = append(gs, e)
	}
	sort.Slice(gs, func(i, j int) bool { return gs[i].Name < gs[j].Name })

	mems := map[string][]EditorMember{}
	add := func(class string, c class) {
		list := make([]EditorMember, 0, len(c))
		for name, m := range c {
			list = append(list, memberOf(name, m))
		}
		sort.Slice(list, func(i, j int) bool { return list[i].Name < list[j].Name })
		mems[class] = list
	}
	for cl, c := range classes {
		add(cl, c)
	}
	add("text", textMembers)
	add("number", intMembers)
	add("list", listMembers)

	var bs []EditorBoard
	for _, n := range BoardNames() {
		b := boards[n]
		bs = append(bs, EditorBoard{Name: n, Display: b.Display, Pins: b.Names, Valid: b.ValidPins, PWM: b.PWM, Analog: b.AnalogPins()})
	}
	return &EditorInfo{LanguageVersion: LanguageVersion(), Keywords: kw, Statements: []string{"pass", "print"}, Units: un, Globals: gs, Members: mems, Boards: bs}
}
