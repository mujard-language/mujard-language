package lang

// A library description generator: reads an installed Arduino library's folder (library.properties
// plus its headers) and produces the LibDesc a '.mlib' file holds. Like LibSymbols on the server it
// is a plain reading of the header text - no compiler - so templates and macro-generated members can
// be missed; what it finds is accurate, and a missed member only means that call stays unchecked.

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	libGenMaxHeaders = 12
	libGenMaxHeader  = 300 * 1024
	libGenMaxMembers = 160
)

// DescribeLibrary reads the Arduino library in dir and returns the description of what it offers.
// library.properties gives the name, version, architectures and header list; the headers give the
// classes, functions and constants.
func DescribeLibrary(dir string) (*LibDesc, error) {
	info, err := os.Stat(dir)
	if err != nil || !info.IsDir() {
		return nil, fmt.Errorf("no library folder at %s", dir)
	}
	props := readProperties(filepath.Join(dir, "library.properties"))
	pioName, pioArch, pioHeaders := readPioLibraryJSON(filepath.Join(dir, "library.json"))
	d := &LibDesc{Name: props["name"]}
	if d.Name == "" {
		d.Name = pioName
	}
	if d.Name == "" {
		d.Name = strings.ReplaceAll(filepath.Base(dir), "_", " ")
	}
	arch := props["architectures"]
	if arch == "" {
		arch = pioArch
	}
	if arch != "" && arch != "*" {
		for _, a := range strings.Split(arch, ",") {
			if a = strings.TrimSpace(strings.ToLower(a)); a != "" {
				d.Architectures = append(d.Architectures, a)
			}
		}
	}
	includes := props["includes"]
	if includes == "" {
		includes = pioHeaders
	}
	headers := libHeaders(dir, includes)
	if len(headers) == 0 {
		return nil, fmt.Errorf("no headers in %s", dir)
	}
	d.Header = filepath.Base(headers[0])
	classes := map[string]*LibClass{}
	classNames := map[string]bool{}
	fns := map[string]*LibFn{}
	consts := map[string]string{}
	for _, h := range headers {
		b, err := os.ReadFile(h)
		if err != nil || len(b) > libGenMaxHeader {
			continue
		}
		parseHeader(string(b), classes, classNames, fns, consts)
	}
	if len(classes) > 0 {
		d.Classes = classes
	}
	if len(fns) > 0 {
		d.Functions = fns
	}
	if len(consts) > 0 {
		d.Constants = consts
	}
	return d, nil
}

// DescribeLibraryJSON is DescribeLibrary's result as the '.mlib' file's text.
func DescribeLibraryJSON(dir string) (string, error) {
	d, err := DescribeLibrary(dir)
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b) + "\n", nil
}

// DescribeHeaderSource parses one C++ header's text - not a library folder but a pasted or uploaded '.h' - into the
// LibDesc a '.mlib' file holds. It's the auto-wrapper for a single header the editor is handed: wasm and the
// desktop's import both go here so the browser and the phone wrap a header the same way 'mujard describe' does a
// library folder. name is what 'use "name"' looks up; header is the '#include <header>' the generated code needs.
func DescribeHeaderSource(name, header, src string) (*LibDesc, error) {
	if len(src) > libGenMaxHeader {
		return nil, fmt.Errorf("%s is too big to be a library header (over %d bytes)", header, libGenMaxHeader)
	}
	d := &LibDesc{Name: name, Header: header}
	if d.Name == "" {
		d.Name = strings.TrimSuffix(header, filepath.Ext(header))
	}
	if d.Name == "" {
		return nil, fmt.Errorf("name the library (what 'use \"name\"' will ask for)")
	}
	classes := map[string]*LibClass{}
	classNames := map[string]bool{}
	fns := map[string]*LibFn{}
	consts := map[string]string{}
	parseHeader(src, classes, classNames, fns, consts)
	if len(classes) == 0 && len(fns) == 0 && len(consts) == 0 {
		return nil, fmt.Errorf("no classes, functions or constants to offer in %s", header)
	}
	d.Classes, d.Functions, d.Constants = classes, fns, consts
	return d, nil
}

// DescribeHeaderSourceJSON is DescribeHeaderSource's result as the '.mlib' file's text.
func DescribeHeaderSourceJSON(name, header, src string) (string, error) {
	d, err := DescribeHeaderSource(name, header, src)
	if err != nil {
		return "", err
	}
	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b) + "\n", nil
}

// readProperties reads the flat 'key=value' lines of library.properties.
func readProperties(path string) map[string]string {
	out := map[string]string{}
	f, err := os.Open(path)
	if err != nil {
		return out
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if i := strings.IndexByte(line, '='); i > 0 {
			out[strings.TrimSpace(line[:i])] = strings.TrimSpace(line[i+1:])
		}
	}
	return out
}

// readPioLibraryJSON reads a PlatformIO library.json manifest (name, platforms, headers).
func readPioLibraryJSON(path string) (name, arch, includes string) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", "", ""
	}
	var raw struct {
		Name      string `json:"name"`
		Platforms any    `json:"platforms"`
		Headers   any    `json:"headers"`
	}
	if json.Unmarshal(b, &raw) != nil {
		return "", "", ""
	}
	pioMap := map[string]string{
		"atmelavr": "avr", "espressif32": "esp32", "espressif8266": "esp8266",
		"raspberrypi": "rp2040", "ststm32": "stm32", "atmelsam": "samd", "renesas-ra": "renesas_uno",
	}
	addPlat := func(p string, list *[]string) {
		p = strings.TrimSpace(strings.ToLower(p))
		if p == "" || p == "*" {
			return
		}
		if mapped, ok := pioMap[p]; ok {
			p = mapped
		}
		*list = append(*list, p)
	}
	var plats []string
	switch v := raw.Platforms.(type) {
	case string:
		for _, p := range strings.Split(v, ",") {
			addPlat(p, &plats)
		}
	case []any:
		for _, item := range v {
			if s, ok := item.(string); ok {
				addPlat(s, &plats)
			}
		}
	}
	var hdrs []string
	switch v := raw.Headers.(type) {
	case string:
		if v = strings.TrimSpace(v); v != "" {
			hdrs = append(hdrs, v)
		}
	case []any:
		for _, item := range v {
			if s, ok := item.(string); ok && strings.TrimSpace(s) != "" {
				hdrs = append(hdrs, strings.TrimSpace(s))
			}
		}
	}
	return strings.TrimSpace(raw.Name), strings.Join(plats, ","), strings.Join(hdrs, ",")
}

// libHeaders picks the headers to read: the 'includes=' list from library.properties (or library.json)
// when it is there, else every header in include/, src/, or the folder root for a flat library.
func libHeaders(dir, includes string) []string {
	var out []string
	add := func(f string) {
		if len(out) < libGenMaxHeaders {
			out = append(out, f)
		}
	}
	if includes != "" {
		for _, h := range strings.Split(includes, ",") {
			h = strings.TrimSpace(h)
			for _, base := range []string{filepath.Join(dir, "include", h), filepath.Join(dir, "src", h), filepath.Join(dir, h)} {
				if st, err := os.Stat(base); err == nil && !st.IsDir() {
					add(base)
					break
				}
			}
		}
		return out
	}
	for _, base := range []string{filepath.Join(dir, "include"), filepath.Join(dir, "src"), dir} {
		entries, err := os.ReadDir(base)
		if err != nil {
			continue
		}
		for _, e := range entries {
			lower := strings.ToLower(e.Name())
			if e.IsDir() || (!strings.HasSuffix(lower, ".h") && !strings.HasSuffix(lower, ".hpp")) {
				continue
			}
			if strings.HasPrefix(e.Name(), "mujard_") { // our own runtime files are not a library
				continue
			}
			add(filepath.Join(base, e.Name()))
		}
		if len(out) > 0 {
			break
		}
	}
	return out
}

var (
	classRe  = regexp.MustCompile(`\b(class|struct)\s+(?:[A-Z_][A-Z0-9_]*\s+)?([A-Za-z_]\w*)\s*(?:final\s*)?(?::[^{;]*)?\{`)
	defineRe = regexp.MustCompile(`(?m)^[ \t]*#[ \t]*define[ \t]+([A-Za-z_]\w*)(?:[ \t]+([^/\n]+))?`)
	labelRe  = regexp.MustCompile(`\b(public|private|protected)\s*:`)
	declRe   = regexp.MustCompile(`^(?:[A-Za-z_][\w:<>,\s\*&~]*?[\*&]*\s+)?(~?[A-Za-z_]\w*)\s*\(([^()]*)\)\s*(?:const)?\s*(?:override|noexcept)?\s*(?:=\s*(?:0|default|delete))?\s*(?:\{\})?$`)
	fieldRe  = regexp.MustCompile(`^([A-Za-z_][\w:<>,\s\*&]*?[\*&]*)\s+([A-Za-z_]\w*)\s*(?:\[[^\]]*\])?\s*(?:=.*)?$`)
	prepRe   = regexp.MustCompile(`(?m)^[ \t]*#.*$`)
	qualRe   = regexp.MustCompile(`^(?:(?:virtual|static|inline|explicit|extern|constexpr|friend)\s+)+`)
)

// parseHeader adds one header's classes, functions and constants to the tables. Later headers only
// fill names that are still free - the include-first header wins, matching the compiler's view.
func parseHeader(text string, classes map[string]*LibClass, classNames map[string]bool, fns map[string]*LibFn, consts map[string]string) {
	text = stripComments(text)
	for _, m := range defineRe.FindAllStringSubmatch(text, -1) {
		name := m[1]
		if strings.HasPrefix(name, "_") || strings.HasSuffix(name, "_H") || strings.HasSuffix(name, "_HPP") || consts[name] != "" {
			continue
		}
		if len(consts) < libGenMaxMembers {
			consts[name] = constType(m[2])
		}
	}
	flat := prepRe.ReplaceAllString(text, "")
	topStart, top := 0, strings.Builder{}
	for _, loc := range classRe.FindAllStringSubmatchIndex(flat, -1) {
		start, name := loc[0], flat[loc[4]:loc[5]]
		open := loc[1] - 1
		close := matchingBrace(flat, open)
		if close < 0 {
			continue
		}
		top.WriteString(flat[topStart:start])
		topStart = close + 1
		if !classNames[name] && len(classes) < libGenMaxMembers {
			classNames[name] = true
			lc := &LibClass{Methods: map[string]*LibFn{}, Fields: map[string]string{}}
			parseClassBody(name, flat[open+1:close], strings.HasPrefix(flat[loc[2]:loc[3]], "struct"), lc)
			if len(lc.Ctors) > 0 || len(lc.Methods) > 0 || len(lc.Fields) > 0 {
				classes[name] = lc
			}
		}
	}
	top.WriteString(flat[topStart:])
	for _, d := range declList(top.String()) {
		if len(fns) >= libGenMaxMembers || fns[d.name] != nil {
			continue
		}
		switch d.name {
		case "if", "for", "while", "switch", "return", "sizeof", "typedef", "using":
			continue
		}
		if fn := sigFromDecl(d); fn != nil {
			fns[d.name] = fn
		}
	}
}

// parseClassBody reads the public chunk of one class body into lc: constructors, methods, fields.
func parseClassBody(className, body string, isStruct bool, lc *LibClass) {
	flat := blankBodies(body)
	parts := labelRe.Split(flat, -1)
	kinds := labelRe.FindAllString(flat, -1)
	public := isStruct
	for i, part := range parts {
		if i > 0 {
			public = strings.HasPrefix(kinds[i-1], "public")
		}
		if !public {
			continue
		}
		for _, d := range declList(part) {
			if d.name == "" {
				continue
			}
			if d.name == className { // a constructor
				if fn := sigFromDecl(d); fn != nil && len(lc.Ctors) < 8 {
					lc.Ctors = append(lc.Ctors, *fn)
				}
				continue
			}
			if strings.HasPrefix(d.name, "~") || strings.HasPrefix(d.name, "operator") || len(lc.Methods) >= libGenMaxMembers {
				continue
			}
			if fn := sigFromDecl(d); fn != nil {
				if old := lc.Methods[d.name]; old == nil {
					lc.Methods[d.name] = fn
				} else {
					mergeOverload(old, fn)
				}
				continue
			}
			// A field: "type name;" or "type name = x;".
			if m := fieldRe.FindStringSubmatch(qualRe.ReplaceAllString(d.head, "")); m != nil {
				if t := mujardType(m[1], nil); t != "void" && len(lc.Fields) < libGenMaxMembers {
					lc.Fields[m[2]] = t
				}
			}
		}
	}
	// A C++ class with public members and no explicit constructor still has an implicit default constructor.
	if len(lc.Ctors) == 0 && (len(lc.Methods) > 0 || len(lc.Fields) > 0) {
		lc.Ctors = []LibFn{{Result: className}}
	}
}

// decl is one raw statement that might be a function declaration or a field.
type decl struct {
	head string // the statement text before '('
	name string
	args string
}

// declList splits a text chunk into declarations: things that look like 'ret name(args)' keep their
// signature; the rest stay as possible fields via .head.
func declList(text string) []decl {
	flat := blankBodies(text)
	var out []decl
	for _, stmt := range strings.Split(flat, ";") {
		s := strings.Join(strings.Fields(stmt), " ")
		if s == "" || strings.HasPrefix(s, "typedef") || strings.HasPrefix(s, "using ") || strings.HasPrefix(s, "return") {
			continue
		}
		if strings.HasPrefix(s, "template") || strings.HasPrefix(s, "friend ") || strings.HasPrefix(s, "enum ") {
			continue
		}
		s = qualRe.ReplaceAllString(s, "")
		if !strings.Contains(s, "(") {
			out = append(out, decl{head: s})
			continue
		}
		if m := declRe.FindStringSubmatch(s); m != nil {
			name := m[1]
			if name == "operator" {
				continue
			}
			out = append(out, decl{head: s, name: name, args: m[2]})
			continue
		}
		out = append(out, decl{head: s}) // not a declaration - maybe a field with an initializer
	}
	return out
}

// sigFromDecl turns a parsed declaration into a LibFn: parameters split and typed, the result read
// from the text before the name. Returns nil for declarations the type table can't make sense of.
func sigFromDecl(d decl) *LibFn {
	fn := &LibFn{Result: "void"}
	if i := strings.LastIndex(d.head, d.name); i > 0 {
		ret := strings.TrimSpace(d.head[:i])
		ret = qualRe.ReplaceAllString(ret, "")
		if ret != "" {
			fn.Result = mujardType(ret, nil)
		}
	}
	fn.Params = parseParams(d.args)
	return fn
}

// parseParams splits an argument list into typed parameters; a 'x = default' marks it optional.
func parseParams(args string) []LibParam {
	var out []LibParam
	optionalSeen := false
	for _, part := range splitArgs(args) {
		optional := optionalSeen
		if i := strings.IndexByte(part, '='); i >= 0 {
			part, optional, optionalSeen = part[:i], true, true
		}
		part = strings.Join(strings.Fields(strings.TrimSpace(part)), " ")
		if part == "" || part == "void" {
			continue
		}
		// "const String& name", "uint8_t pin = 13", function pointers "( *f)(int)".
		if strings.Contains(part, "(*") {
			out = append(out, LibParam{Name: "callback", Type: "any", Optional: optional})
			continue
		}
		name := part
		typ := ""
		if i := strings.LastIndexAny(part, " *&"); i >= 0 {
			name, typ = part[i+1:], strings.TrimSpace(strings.TrimRight(part[:i+1], "*&"))
		} else if i := strings.LastIndexByte(part, ' '); i >= 0 {
			name, typ = part[i+1:], part[:i]
		}
		if name == "" {
			name = "arg" + fmt.Sprint(len(out)+1)
		}
		out = append(out, LibParam{Name: name, Type: mujardType(typ, nil), Optional: optional})
	}
	return out
}

// mergeOverload folds another signature for the same name into fn: parameters past the shorter
// list's end become optional, so 'setPixelColor(n, color)' and 'setPixelColor(n, r, g, b)' both fit.
func mergeOverload(fn, alt *LibFn) {
	short := len(fn.Params)
	if len(alt.Params) < short {
		short = len(alt.Params)
	}
	for i := 0; i < short; i++ {
		if alt.Params[i].Optional {
			fn.Params[i].Optional = true
		}
	}
	if len(alt.Params) > len(fn.Params) {
		extra := alt.Params[len(fn.Params):]
		for i := range extra {
			extra[i].Optional = true
		}
		fn.Params = append(fn.Params, extra...)
	}
	if fn.Result == "void" && alt.Result != "void" {
		fn.Result = alt.Result
	}
}

// mujardType maps a C++ type spelling to a description type name.
func mujardType(t string, classNames map[string]bool) string {
	t = strings.Join(strings.Fields(strings.TrimSpace(t)), " ")
	t = strings.TrimPrefix(t, "const ")
	t = strings.TrimPrefix(t, "typename ")
	ptr := strings.ContainsAny(t, "*[")
	t = strings.TrimSpace(strings.TrimRight(t, "*&[] "))
	if t == "" {
		return "any"
	}
	if i := strings.LastIndex(t, "::"); i >= 0 {
		t = t[i+2:]
	}
	if strings.ContainsAny(t, "<>") {
		return "any"
	}
	switch t {
	case "void":
		return "void"
	case "bool", "boolean":
		return "bool"
	case "String", "string", "std::string", "__FlashStringHelper":
		return "text"
	case "char":
		if ptr {
			return "text" // char* / char[] is a string; a bare char is a number
		}
		return "int"
	case "float", "double":
		return "float"
	}
	if ptr {
		return "any" // any other pointer is a buffer or object the checker can't shape
	}
	switch {
	case t == "size_t" || t == "ssize_t",
		strings.HasPrefix(t, "int") && strings.HasSuffix(t, "_t"),
		strings.HasPrefix(t, "uint") && strings.HasSuffix(t, "_t"),
		t == "int" || t == "long" || t == "short" || t == "byte" || t == "word" || t == "unsigned":
		return "int"
	}
	if classNames[t] {
		return t
	}
	return "any"
}

// constType guesses a #define's value type from its text.
func constType(v string) string {
	v = strings.TrimSpace(v)
	if v == "" {
		return "int"
	}
	if strings.HasPrefix(v, "\"") || strings.HasPrefix(v, "F(") {
		return "text"
	}
	if strings.ContainsAny(v, ".eE") && !strings.ContainsAny(v, "xX") {
		return "float"
	}
	return "int"
}

// stripComments removes // and block comments.
func stripComments(t string) string {
	var b strings.Builder
	b.Grow(len(t))
	for i := 0; i < len(t); {
		if i+1 < len(t) && t[i] == '/' && t[i+1] == '/' {
			for i < len(t) && t[i] != '\n' {
				i++
			}
			continue
		}
		if i+1 < len(t) && t[i] == '/' && t[i+1] == '*' {
			i += 2
			for i+1 < len(t) && !(t[i] == '*' && t[i+1] == '/') {
				i++
			}
			i += 2
			b.WriteByte(' ')
			continue
		}
		b.WriteByte(t[i])
		i++
	}
	return b.String()
}

// blankBodies replaces every { ... } with '{}' so inline code cannot look like declarations (and a
// ';' inside a body cannot split one).
func blankBodies(t string) string {
	var b strings.Builder
	b.Grow(len(t))
	depth := 0
	for i := 0; i < len(t); i++ {
		switch t[i] {
		case '{':
			if depth == 0 {
				b.WriteString("{};")
			}
			depth++
		case '}':
			if depth > 0 {
				depth--
			}
		default:
			if depth == 0 {
				b.WriteByte(t[i])
			}
		}
	}
	return b.String()
}

// matchingBrace returns the index of the '}' closing the '{' at open, or -1.
func matchingBrace(t string, open int) int {
	depth := 0
	for i := open; i < len(t); i++ {
		switch t[i] {
		case '{':
			depth++
		case '}':
			if depth--; depth == 0 {
				return i
			}
		}
	}
	return -1
}

// splitArgs splits an argument list on top-level commas (parens, angles and brackets don't count).
func splitArgs(args string) []string {
	var out []string
	depth, start := 0, 0
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case '(', '<', '[':
			depth++
		case ')', '>', ']':
			if depth > 0 {
				depth--
			}
		case ',':
			if depth == 0 {
				out = append(out, args[start:i])
				start = i + 1
			}
		}
	}
	if tail := strings.TrimSpace(args[start:]); tail != "" {
		out = append(out, tail)
	}
	return out
}
