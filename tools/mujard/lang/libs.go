package lang

// Library descriptions (.mlib files, JSON): what a 'use "…"' line's library offers - classes with their
// constructors and methods, free functions and constants - so calls into C++ libraries are checked like
// Mujard's own. A generator writes one per installed library version; the compiler reads any '.mlib' file
// that sits next to the program (or is handed over in the modules map). Without one, 'use' warns W611 and
// the calls go through unchecked.

import (
	"embed"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// descs/*.mlib are the checked descriptions the compiler ships: the common Arduino libraries the
// implementation plan lists, so 'use "Adafruit NeoPixel"' checks calls without any project file.
// A '.mlib' next to the program (or in the modules map) overrides these.
//
//go:embed descs/*.mlib
var descsFS embed.FS

// builtinDescs is every embedded description, parsed once at startup.
var builtinDescs = loadDescs()

// DescribedLibs is the shipped .mlib descriptions, for generated documentation ('mujard doc').
func DescribedLibs() []*LibDesc { return builtinDescs }

func loadDescs() []*LibDesc {
	entries, err := descsFS.ReadDir("descs")
	if err != nil {
		return nil
	}
	var out []*LibDesc
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".mlib") {
			continue
		}
		if b, err := descsFS.ReadFile("descs/" + e.Name()); err == nil {
			if d, err := ParseLibDesc(string(b)); err == nil {
				out = append(out, d)
			}
		}
	}
	return out
}

// LibDesc is one library's description: what 'use "Name"' makes available.
type LibDesc struct {
	Name          string               `json:"name"`                    // as written in 'use "Name"'
	Header        string               `json:"header"`                  // "Adafruit_NeoPixel.h" - the #include
	Headers       []string             `json:"headers,omitempty"`       // more #includes for classes in their own header (MultiStepper.h)
	Architectures []string             `json:"architectures,omitempty"` // from library.properties; empty: every board
	Pins          []int                `json:"pins,omitempty"`          // pins the library needs to itself (W609)
	Takes         []string             `json:"takes,omitempty"`         // board claims it uses: "i2c", "spi", "servo" (W609, W613)
	Classes       map[string]*LibClass `json:"classes,omitempty"`
	Functions     map[string]*LibFn    `json:"functions,omitempty"`
	Constants     map[string]string    `json:"constants,omitempty"` // name -> type ("int", ...)
	// Families are generated classes too many to list one by one: U8g2's ~4,900 display classes are
	// U8G2_<display>_<buffer>_<bus>, each a U8G2 with a constructor fixed by the bus. A family's name
	// (for example "U8G2") is the class every member extends.
	Families map[string]*LibFamily `json:"families,omitempty"`
}

// LibClass is one C++ class: constructor signatures, methods and fields. Extends names another class of the same
// description whose methods and fields it has too (its own win), as a C++ subclass does. An abstract class is
// only there to be extended (Adafruit_ADS1X15 under Adafruit_ADS1115): a program never names it, whereas a class
// with no constructors is otherwise the library's ready-made object ('WiFi').
type LibClass struct {
	Extends  string `json:"extends,omitempty"`
	Abstract bool   `json:"abstract,omitempty"`
	// Value: a type the library hands out (SD.open()'s File, gps.location's TinyGPSLocation) - a program gets one from a
	// method or field, so it is neither made with a constructor nor a ready-made object by its class name.
	Value   bool              `json:"value,omitempty"`
	Ctors   []LibFn           `json:"ctors,omitempty"`
	Methods map[string]*LibFn `json:"methods,omitempty"`
	Fields  map[string]string `json:"fields,omitempty"`
}

// LibFamily describes classes named Prefix + stem + "_" + mid + "_" + variant: the stems a variant lists, each with
// every mid ("F", "1", "2" - U8g2's buffer sizes); the variant ("HW_I2C", "4W_SW_SPI", ...) fixes the constructor.
// Only names the library really has match, so a misspelt display is "unknown name", not a C++ error later.
type LibFamily struct {
	Prefix   string                 `json:"prefix"`
	Mids     []string               `json:"mids,omitempty"`
	Variants map[string]*LibVariant `json:"variants"`
}

// LibVariant is one constructor shape of a family and the stems that come in it.
type LibVariant struct {
	Ctor  LibFn    `json:"ctor"`
	Stems []string `json:"stems"`
}

// familyClass finds the family a class name belongs to: the base class it extends and its constructor.
func (d *LibDesc) familyClass(name string) (base string, ctor *LibFn) {
	for b, f := range d.Families {
		if !strings.HasPrefix(name, f.Prefix) {
			continue
		}
		for vname, v := range f.Variants {
			rest, ok := strings.CutSuffix(name[len(f.Prefix):], "_"+vname)
			if !ok {
				continue
			}
			stem := ""
			if len(f.Mids) == 0 {
				stem = rest
			} else {
				for _, m := range f.Mids {
					if s, ok := strings.CutSuffix(rest, "_"+m); ok {
						stem = s
						break
					}
				}
			}
			if stem == "" {
				continue
			}
			for _, s := range v.Stems {
				if s == stem {
					return b, &v.Ctor
				}
			}
		}
	}
	return "", nil
}

// LibFn is one callable signature: parameters and a result, types as Mujard names
// ("int", "text", "bool", "float", "time", "list<int>", a class name, "...?" for optional).
type LibFn struct {
	Params []LibParam `json:"params"`
	Result string     `json:"result"`
}

// LibParam is one argument of a LibFn.
type LibParam struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Optional bool   `json:"optional,omitempty"`
}

// LibDep is a library a program needs installed on the build server, recorded for the 'mujard_libs.json'
// manifest next to the generated sketch.
type LibDep struct {
	From    string `json:"from"` // "arduino", "idf" or "pio"
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
	Header  string `json:"header,omitempty"` // the include that provides it, when known
}

// ParseLibDesc reads one description. The first hard problem is reported; anything unparseable in a type
// name degrades to 'any' rather than failing, so a loosely generated file still helps.
func ParseLibDesc(data string) (*LibDesc, error) {
	var d LibDesc
	if err := json.Unmarshal([]byte(data), &d); err != nil {
		return nil, err
	}
	if d.Name == "" {
		return nil, errNoLibName
	}
	// A description can come with a project, so what it makes the sketch #include must be a library header's name,
	// never a path out of the library folders (a server file's text would come back in a compiler error).
	for _, h := range append([]string{d.Header}, d.Headers...) {
		if h != "" && !libHeaderName.MatchString(h) || strings.Contains(h, "..") {
			return nil, jsonErr("the description's header " + strconv.Quote(h) + " isn't a library header name")
		}
	}
	return &d, nil
}

// libHeaderName is a header a library provides: "Adafruit_GFX.h", "utility/twi.h" - letters, digits and _ . + -,
// folders allowed below the library, never absolute or with a drive.
var libHeaderName = regexp.MustCompile(`^[A-Za-z0-9_][A-Za-z0-9_.+\-]*(?:/[A-Za-z0-9_][A-Za-z0-9_.+\-]*)*\.(?:h|hh|hpp)$`)

var errNoLibName = jsonErr("a library description needs a \"name\"")

type jsonErr string

func (e jsonErr) Error() string { return string(e) }

// libAlias normalises a name for the snake_case aliasing the spec gives C++ libraries:
// 'read_temperature' and 'readTemperature' are the same name, and so are 'Adafruit NeoPixel'
// and 'Adafruit_NeoPixel'.
func libAlias(name string) string {
	return strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(name, "_", ""), " ", ""))
}

// snakeName is the spelling the spec promises for camelCase library names: 'readTemperature' may also be
// written 'read_temperature'.
func snakeName(name string) string {
	if !strings.ContainsAny(name, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return name
	}
	var b strings.Builder
	for i, r := range name {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				b.WriteByte('_')
			}
			b.WriteRune(r + 'a' - 'A')
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// realName finds the member in a class table matching a name as written: exact first, then the
// snake_case alias ('set_pixel_color' -> 'setPixelColor'). Returns "" on no match.
func realName(written string, cl class) string {
	if _, ok := cl[written]; ok {
		return written
	}
	want := libAlias(written)
	for n := range cl {
		if libAlias(n) == want {
			return n
		}
	}
	return ""
}

// libType turns a description's type name into a *Type. Unknown names become 'any' - a generated
// description may name a C++ type Mujard has no word for, and those calls just aren't type-checked.
func libType(s string) *Type {
	s = strings.TrimSpace(s)
	if s == "" {
		return tAny
	}
	optSuffix := strings.HasSuffix(s, "?")
	if optSuffix {
		s = s[:len(s)-1]
	}
	var t *Type
	switch s {
	case "void":
		t = tVoid
	case "any":
		t = tAny
	case "bool":
		t = tBool
	case "int", "byte":
		t = tInt
	case "float", "double":
		t = tFloat
	case "text", "string":
		t = tText
	default:
		if strings.HasPrefix(s, "list<") && strings.HasSuffix(s, ">") {
			t = listOf(libType(s[5 : len(s)-1]))
		} else if dims[s] {
			t = qty(s)
		} else {
			// A class name: an object of the described library (or an unknown one - 'any' keeps it free).
			t = object(s)
		}
	}
	if optSuffix {
		t = opt(t)
	}
	return t
}

// dims are the quantity names a description may use ('500 ms' is time, '90 deg' angle).
var dims = map[string]bool{
	"time": true, "frequency": true, "length": true, "voltage": true, "current": true,
	"angle": true, "temperature": true, "ratio": true, "speed": true, "memory": true,
}

// libSig converts a LibFn into the checker's fnSig. A nil sig (no signature written) means "don't check".
func libSig(fn *LibFn) *fnSig {
	if fn == nil {
		return nil
	}
	ps := make([]param, 0, len(fn.Params))
	for _, p := range fn.Params {
		t := libType(p.Type)
		if p.Type == "" {
			t = nil // a bare name: any value
		}
		ps = append(ps, param{name: p.Name, t: t, optional: p.Optional})
	}
	return &fnSig{params: ps, result: libType(fn.Result)}
}

// class builds the checker's member table for one described class.
func (d *LibDesc) classTable() map[string]class {
	out := map[string]class{}
	for name := range d.Classes {
		out[name] = d.members(name)
	}
	return out
}

// members is one class's member table, with what it inherits through 'extends' (a loop just stops).
func (d *LibDesc) members(name string) class {
	cl := class{}
	seen := map[string]bool{}
	for lc := d.Classes[name]; lc != nil && !seen[name]; lc = d.Classes[name] {
		seen[name] = true
		for m, f := range lc.Methods {
			if cl[m] == nil {
				cl[m] = &member{method: libSig(f)}
			}
		}
		for f, t := range lc.Fields {
			if cl[f] == nil {
				cl[f] = &member{field: libType(t)}
			}
		}
		name = lc.Extends
	}
	return cl
}

// familyGlobal is a family member's constructor and member table, for a name a program writes
// ('U8G2_SSD1306_128X64_NONAME_F_HW_I2C(...)'); ok is false when the name is no member of a family.
func (d *LibDesc) familyGlobal(name string) (g *global, cl class, ok bool) {
	base, ctor := d.familyClass(name)
	if ctor == nil {
		return nil, nil, false
	}
	sig := libSig(ctor)
	sig.result = object(name)
	return &global{fn: sig, class: name, emit: name}, d.members(base), true
}

// globals builds the checker's top-level names for one described library: class constructors, free
// functions and constants, plus their snake_case aliases.
func (d *LibDesc) globalsTable() map[string]*global {
	out := map[string]*global{}
	put := func(name string, g *global) {
		g.emit = name // the C++ spelling, for when the program writes the snake_case alias
		out[name] = g
		if s := snakeName(name); s != name {
			out[s] = g // 'read_temperature' finds 'readTemperature'
		}
	}
	for name, lc := range d.Classes {
		if d.Families[name] != nil || lc.Abstract || lc.Value {
			continue // a base (U8G2, Adafruit_ADS1X15), or a type values come in (File): never a name a program writes
		}
		if len(lc.Ctors) > 0 {
			sig := libSig(&lc.Ctors[0])
			sig.result = object(name) // a constructor returns an object of its class
			put(name, &global{fn: sig, class: name})
		} else {
			put(name, &global{value: object(name)})
		}
	}
	for name, f := range d.Functions {
		put(name, &global{fn: libSig(f)})
	}
	for key, t := range d.Constants {
		// "AccelStepper::DRIVER": a constant inside a class - the program writes DRIVER, the C++ gets the full name
		name := key
		if i := strings.LastIndex(key, "::"); i >= 0 {
			name = key[i+2:]
		}
		g := &global{value: libType(t)}
		put(name, g)
		g.emit = key
	}
	return out
}

// ---- where descriptions and project files come from

// libEnv is what a checker knows about the project's surroundings: the library descriptions in play and
// whether a relative name is a project file.
type libEnv struct {
	descs []*LibDesc
	has   func(name string) bool // "crc.h" or "lib/crc.h" is a file in the project
}

// dirEnv builds a libEnv for a directory: every '*.mlib' there is a description, and 'use "x.h"' can look
// for the file on disk.
func dirEnv(dir string) *libEnv {
	env := &libEnv{}
	if dir == "" {
		return env // no files around the program
	}
	env.has = func(name string) bool {
		_, err := os.Stat(filepath.Join(dir, filepath.FromSlash(name)))
		return err == nil
	}
	env.descs = descDirEntries(dir)
	// MUJARD_LIB_DESCS is the build server's description cache: generated '.mlib' files kept next to
	// the toolchain so every compiled program sees them without shipping them in the project.
	if extra := os.Getenv("MUJARD_LIB_DESCS"); extra != "" {
		for _, e := range descDirEntries(extra) {
			env.descs = append(env.descs, e)
		}
	}
	return env
}

// descDirEntries parses every '*.mlib' in a directory.
func descDirEntries(dir string) []*LibDesc {
	var out []*LibDesc
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(strings.ToLower(e.Name()), ".mlib") {
			continue
		}
		if b, err := os.ReadFile(filepath.Join(dir, e.Name())); err == nil {
			if d, err := ParseLibDesc(string(b)); err == nil {
				out = append(out, d)
			}
		}
	}
	return out
}

// mapEnv is dirEnv for in-memory project files (the browser editor): '*.mlib' entries are descriptions and
// every key is a project file.
func mapEnv(files map[string]string) *libEnv {
	env := &libEnv{has: func(name string) bool {
		_, ok := files[name]
		return ok
	}}
	for name, text := range files {
		if !strings.HasSuffix(strings.ToLower(name), ".mlib") {
			continue
		}
		if d, err := ParseLibDesc(text); err == nil {
			env.descs = append(env.descs, d)
		}
	}
	return env
}

// findLib matches a 'use "…"' path against the descriptions: project files first, then the
// descriptions the compiler ships. By name, by header, or by the loose snake_case/spelling form.
func (e *libEnv) findLib(path string) *LibDesc {
	want := libAlias(strings.TrimSuffix(strings.ToLower(path), ".h"))
	match := func(descs []*LibDesc) *LibDesc {
		for _, d := range descs {
			if libAlias(d.Name) == want || libAlias(strings.TrimSuffix(strings.ToLower(d.Header), ".h")) == want {
				return d
			}
		}
		return nil
	}
	if e != nil {
		if d := match(e.descs); d != nil {
			return d
		}
	}
	return match(builtinDescs)
}

// hasFile reports whether a 'use "x.h"' path is a file in the project (a quoted include, not a library).
func (e *libEnv) hasFile(path string) bool {
	return e != nil && e.has != nil && e.has(path)
}

// isFilePath reports whether a 'use "…"' path is written like a file ('lib/crc.h', 'crc.h'): a path
// separator, or a C/C++ extension.
func isFilePath(path string) bool {
	if strings.ContainsAny(path, "/\\") {
		return true
	}
	switch strings.ToLower(filepath.Ext(path)) {
	case ".h", ".hpp", ".c", ".cpp", ".hh", ".hxx":
		return true
	}
	return false
}

// headerFromName is the conventional include for a library named like "Adafruit NeoPixel" or
// "bblanchon/ArduinoJson": spaces to underscores, ".h" when there isn't one already.
func headerFromName(name string) string {
	h := strings.ReplaceAll(name, " ", "_")
	if i := strings.LastIndex(h, "/"); i >= 0 {
		h = h[i+1:]
	}
	if l := strings.ToLower(h); !strings.HasSuffix(l, ".h") && !strings.HasSuffix(l, ".hpp") {
		h += ".h"
	}
	return h
}
