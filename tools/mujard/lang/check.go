package lang

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// CheckSource lexes, parses and checks a program. All diagnostics come back sorted by position, errors and warnings
// together.
func CheckSource(src string) (*File, []Diagnostic) {
	return CheckSourceIn(src, "")
}

// CheckSourceIn is CheckSource with a directory for resolving 'import name' to name.muj next to the program.
// With an empty dir, 'import' still reports C320 (the program has no files around it to load).
func CheckSourceIn(src, dir string) (*File, []Diagnostic) {
	r := CheckProgramIn(src, dir)
	return r.File, r.Diags
}

// CheckSourceModules is CheckSource with 'import name' resolved against sources held in memory - the web
// editor's WASM build passes the project's other .muj/.mjd files this way (file name -> source). A name
// absent from the map is "can't find the module" (C321).
func CheckSourceModules(src string, modules map[string]string) (*File, []Diagnostic) {
	r := CheckProgramModules(src, modules)
	return r.File, r.Diags
}

// CheckResult is a check and what the build will need from the project.
type CheckResult struct {
	File  *File
	Diags []Diagnostic // sorted by position, errors and warnings together
	// Secrets are the build-secret names the program's secret() calls compile to ("HOME_WIFI_SSID",
	// "HOME_WIFI_PASSWORD" for secret("home_wifi") in wifi.connect), sorted - what the project's secrets must
	// hold for a build. Only for a program that checks clean (nil otherwise). Names, never values.
	Secrets []string
}

// CheckProgramIn is CheckSourceIn with the secrets the program needs.
func CheckProgramIn(src, dir string) CheckResult {
	return checkProgram(src, moduleLoader(dir), dirEnv(dir))
}

// CheckProgramModules is CheckSourceModules with the secrets the program needs.
func CheckProgramModules(src string, modules map[string]string) CheckResult {
	return checkProgram(src, moduleLoaderMap(modules), mapEnv(modules))
}

func checkProgram(src string, load func(name string) (*File, []Diagnostic, string), env *libEnv) CheckResult {
	f, diags := ParseSource(src)
	r := CheckResult{File: f}
	if len(diags) == 0 {
		var info *Info
		info, diags = analyzeWith(f, load, env)
		diags, r.Secrets = withGenerated(f, info, diags)
	}
	r.Diags = sortDiags(diags)
	return r
}

// withGenerated adds what the C++ generator would refuse (C320, a feature that comes later) to a clean check, so a
// program that checks clean also builds: the CLI, the server's check.php and the editors' live checker all report the
// generator's refusals before a build does. Nothing generated is kept but the secret names. A refusal the checker
// already made (same code, same place) isn't repeated.
func withGenerated(f *File, info *Info, diags []Diagnostic) (out []Diagnostic, secrets []string) {
	if info == nil || info.Board == nil || FirstError(diags) != nil {
		return diags, nil
	}
	defer func() {
		// The editors check on every keystroke: a generator fault must not take the checker down with it (the
		// build still runs the generator and reports the fault there).
		if recover() != nil {
			out, secrets = diags, nil
		}
	}()
	seen := map[string]bool{}
	key := func(d Diagnostic) string { return fmt.Sprintf("%s|%s|%d|%d", d.Code, d.File, d.Pos.Line, d.Pos.Col) }
	for _, d := range diags {
		seen[key(d)] = true
	}
	sk, gdiags := Generate(f, info, "sketch.muj")
	out = diags
	for _, d := range gdiags {
		if !seen[key(d)] {
			seen[key(d)] = true
			out = append(out, d)
		}
	}
	if FirstError(out) == nil && sk != nil {
		secrets = sk.Secrets
	}
	return out, secrets
}

// FirstError is the first diagnostic that isn't a warning, or nil.
func FirstError(diags []Diagnostic) *Diagnostic {
	for i := range diags {
		if !diags[i].IsWarning() {
			return &diags[i]
		}
	}
	return nil
}

// Check checks a parsed program against the spec: names, types, units, none, events and tasks, and pins against the
// board (spec sections 4-11).
func Check(f *File) []Diagnostic {
	_, d := Analyze(f)
	return d
}

// Info is what the checker learned about a program, for the C++ generator.
type Info struct {
	Board     *Board
	Types     map[Node]*Type      // the type of every expression (an optional name narrowed by an if is non-optional here)
	Declared  map[*Ident]*Type    // what each name was declared as (to tell a narrowed optional from a plain value)
	Libs      map[string]bool     // mujard libraries in use
	Unchecked bool                // a C++ library or cpp: block is in use
	Modules   map[string]*checker // 'import name' -> the module's own check result (spec 10.3)

	// Libraries described to the checker by '.mlib' files (Phase 8): each 'use' matched to its description,
	// each 'use' that is a project file, the names those libraries offer, and the libraries the build needs.
	UseLibs    map[*UseDecl]*LibDesc
	UseFiles   map[*UseDecl]bool
	LibDeps    []LibDep
	LibClasses map[string]class   // described classes, by C++ name
	LibGlobals map[string]*global // described top-level names, as written (aliases included)
	LibReal    map[string]string  // a name as written -> its C++ spelling ("read_temperature" -> "readTemperature", "Class.member")

	Caps []string // what the program asks of its board, by capability name (capabilities.go)
}

// Analyze checks a program and returns what it learned about it.
func Analyze(f *File) (*Info, []Diagnostic) {
	return analyzeWith(f, nil, nil)
}

// AnalyzeIn is Analyze with a directory for 'import' resolution (like CheckSourceIn).
func AnalyzeIn(f *File, dir string) (*Info, []Diagnostic) {
	return analyzeWith(f, moduleLoader(dir), dirEnv(dir))
}

// SourceExts are the file extensions of Mujard programs: .muj, and .mjd as the alternative.
var SourceExts = []string{".muj", ".mjd"}

// IsSource reports whether a file name has a Mujard source extension.
func IsSource(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	for _, e := range SourceExts {
		if ext == e {
			return true
		}
	}
	return false
}

// moduleLoader resolves 'import name' to dir/name.muj (or dir/name.mjd) and returns the file name it read, "" when
// there is none. dir "" means no files: the loader is nil and imports keep reporting C320.
func moduleLoader(dir string) func(name string) (*File, []Diagnostic, string) {
	if dir == "" {
		return nil
	}
	return func(name string) (*File, []Diagnostic, string) {
		for _, ext := range SourceExts {
			b, err := os.ReadFile(filepath.Join(dir, name+ext))
			if err != nil {
				continue
			}
			f, diags := ParseModuleSource(string(b))
			return f, diags, name + ext
		}
		return nil, nil, ""
	}
}

// moduleLoaderMap resolves 'import name' like moduleLoader, but against sources kept in memory instead of a
// folder: the name plus each source extension is looked up in modules.
func moduleLoaderMap(modules map[string]string) func(name string) (*File, []Diagnostic, string) {
	return func(name string) (*File, []Diagnostic, string) {
		for _, ext := range SourceExts {
			b, ok := modules[name+ext]
			if !ok {
				continue
			}
			f, diags := ParseModuleSource(b)
			return f, diags, name + ext
		}
		return nil, nil, ""
	}
}

func analyzeWith(f *File, load func(name string) (*File, []Diagnostic, string), env *libEnv) (*Info, []Diagnostic) {
	c := &checker{file: f, global: newScope(nil), libs: map[string]bool{}, pins: map[int]pinUse{},
		types: map[Node]*Type{}, declared: map[*Ident]*Type{}, load: load, mods: map[string]*checker{},
		pendingEnum: map[*Ident]pendingEnum{}, env: env}
	c.run()
	return &Info{Board: c.board, Types: c.types, Declared: c.declared, Libs: c.libs, Unchecked: c.unchecked,
		Modules: c.mods, UseLibs: c.useLibs, UseFiles: c.useFiles, LibDeps: c.libDeps,
		LibClasses: c.libClasses, LibGlobals: c.libGlobals, LibReal: c.libReal, Caps: c.programCaps()}, applySafety(f, c.errs)
}

// strictWarnings are the warnings 'safety strict' turns into errors (spec 6.10): each marks code the compiler can't
// check (raw C++, an undescribed library) or a resource the program may not really have (a pin, a timer, the ADC,
// RAM, a register written behind a device's back).
var strictWarnings = map[string]bool{"W604": true, "W609": true, "W611": true, "W613": true, "W614": true, "W615": true, "W616": true, "W617": true}

// applySafety checks the program's safety line and, under 'safety strict', adds an error (C328) for each warning that
// level doesn't allow. The warnings stay, so the reason is still there in the editor.
func applySafety(f *File, diags []Diagnostic) []Diagnostic {
	switch f.Safety {
	case "", "standard":
		return diags
	case "strict":
	default:
		return append(diags, diag("C329", f.SafetyPos, "level", f.Safety))
	}
	out := diags
	for _, d := range diags {
		if strictWarnings[d.Code] {
			e := diag("C328", d.Pos, "warning", d.Code+" "+d.Message)
			e.File = d.File
			out = append(out, e)
		}
	}
	return out
}

type symKind int

const (
	symVar symKind = iota
	symConst
	symDevice
	symFn
	symStruct
	symEnum
	symBlock // a task or a named block
	symParam
	symModule // an imported file: motors.forward(), motors.top_speed
)

type symbol struct {
	kind    symKind
	name    string
	t       *Type
	pos     Pos
	fn      *FnDecl
	strukt  *StructDecl
	enum    *EnumDecl
	used    bool
	waits   bool // a function that waits (only callable from tasks)
	checked bool
	decl    *VarDecl
}

type scope struct {
	parent   *scope
	syms     map[string]*symbol
	narrowed map[string]bool // optional names known not to be none here
	barrier  bool            // a nested 'after' block: the locals around it are gone when it runs
}

// crossesBarrier reports whether name is a local of a block outside a nested 'after' (spec 6.3, C319).
func (s *scope) crossesBarrier(name string) bool {
	crossed := false
	for sc := s; sc != nil; sc = sc.parent {
		if _, ok := sc.syms[name]; ok {
			return crossed && sc.parent != nil // globals are fine
		}
		if sc.barrier {
			crossed = true
		}
	}
	return false
}

func newScope(parent *scope) *scope {
	return &scope{parent: parent, syms: map[string]*symbol{}, narrowed: map[string]bool{}}
}

func (s *scope) lookup(name string) (*symbol, bool) {
	narrowed := false
	for sc := s; sc != nil; sc = sc.parent {
		if sc.narrowed[name] {
			narrowed = true
		}
		if sym, ok := sc.syms[name]; ok {
			return sym, narrowed
		}
	}
	return nil, false
}

func (s *scope) narrow(names []string) *scope {
	n := newScope(s)
	for _, name := range names {
		n.narrowed[name] = true
	}
	return n
}

type pinUse struct {
	name string
	line int
}

// regWrite is a register write/set/clear on a GPIO port register (regsCall), kept for regsConflicts (W616).
type regWrite struct {
	at   Pos
	reg  string
	pins []int
}

// analogPin is one analog read site (an 'analog' device or an analog pin literal): where, which pin, how it's labeled.
type analogPin struct {
	at    Pos
	pin   int
	label string
}

// ctx is where the checker is: what is allowed depends on it.
type ctx struct {
	fn     *FnDecl
	fnSym  *symbol
	inTask bool
	inTest bool
	inSafe bool
	inFast bool
	core   int // the secondary core of the enclosing 'on core N' block (0 = the main core)
	loops  int
	block  string // the event kind of the enclosing block, for messages
}

// waitCall is a call of a user fn from a place that can't suspend; the callee's waits flag is final only after all
// bodies are checked, so the C311 is reported by propagateWaits.
type waitCall struct {
	at     Pos
	callee string
	sym    *symbol
}

type checker struct {
	file      *File
	board     *Board
	global    *scope
	errs      []Diagnostic
	libs      map[string]bool // mujard libraries in use
	unchecked bool            // a C++ library is in use: unknown names belong to it
	pins      map[int]pinUse
	types     map[Node]*Type
	declared  map[*Ident]*Type

	regsWrites []regWrite // GPIO-port register writes, for regsConflicts (W616)
	regsNoChip bool       // the board's chip has no register table yet (B413 once)

	waitDeps  map[*symbol][]*symbol // a fn -> the fns it calls (for wait propagation, order-free)
	waitCalls []waitCall            // calls of user fns from outside tasks: C311 once waits are final

	load     func(name string) (*File, []Diagnostic, string) // resolves 'import name' to a file; nil = no files
	mods     map[string]*checker                             // imported modules by name
	inModule bool                                            // checking a module file
	modName  string                                          // the module's import name
	modFile  string                                          // the module's file name ("motors.muj")
	failed   bool                                            // a module with errors: members stay unknown

	enumMembers map[string][]string    // enum member name -> the enum types that have it (built on demand)
	pendingEnum map[*Ident]pendingEnum // bare enum member names waiting for the context's type (spec 4.6)

	// Described libraries (Phase 8): '.mlib' descriptions in play, the names they offer, which 'use' got
	// which, and the pins they keep (W609).
	env        *libEnv
	libGlobals map[string]*global
	libClasses map[string]class
	libPins    map[int]string
	useLibs    map[*UseDecl]*LibDesc
	useFiles   map[*UseDecl]bool
	libDeps    []LibDep
	libReal    map[string]string

	secretOK    map[*Call]bool    // secret("…") calls written where a secret may go (secretConsumers)
	claimUsers  map[string]string // board claim ("i2c", "spi", "servo") -> the first described library that takes it
	claimAt     map[string]Pos    // where claimUsers' library was named (its 'use' line)
	pinOwner    *DeviceDecl       // the object declaration being checked, which owns the pins written in it
	unboundPins map[string]bool   // a driver module's 'needs pin' names its importer didn't give (C331 already said so)
	usesTone    bool              // a pin's tone() or melody() is called somewhere
	toneAt      Pos               // the first tone() or melody() call
	usesWifi    bool              // a network method or a wifi/mqtt/http member is used somewhere (wifiADC, W615)
	usesTimer   bool              // a Timer() object is created somewhere (timerOwners, B414)
	timerAt     Pos               // the first Timer() call
	timerCount  int               // how many Timer() objects exist (each needs its own hardware timer)
	coreWrites  map[*symbol]Pos   // global variables written inside an 'on core N' block (coreShareCheck, W620)
	coreReads   map[*symbol]Pos   // global variables read inside an 'on core N' block (coreShareCheck, W620)
	mainUses    map[*symbol]bool  // global variables read on the main core (coreShareCheck, W620)
	mainWrites  map[*symbol]bool  // global variables written on the main core (coreShareCheck, W620)
	sliceTones  map[int]struct {
		pin int
		hz  float64
	} // the pin and frequency of a tone() call per PWM slice (RP-family, W623)
	analogPins []analogPin // analog device/pin uses, for W615 after Wi-Fi use is known
	siteBytes  int         // RAM kept by the call sites of smooth, peak, ... (ramEstimate)
}

// pendingEnum is a bare name that might be an enum member (spec 4.6): where its C301 sits in errs, and
// the enum types that have a member of that name.
type pendingEnum struct {
	err   int
	enums []string
}

func (c *checker) errorf(code string, pos Pos, args ...string) {
	for _, e := range c.errs {
		if e.Pos == pos && e.Code == code {
			return
		}
	}
	if len(c.errs) < maxDiagnostics {
		c.errs = append(c.errs, diag(code, pos, args...))
	}
}

// ---- the program

func (c *checker) run() {
	f := c.file
	if f.Board != "" {
		if c.board = LookupBoard(f.Board); c.board == nil {
			c.errorf("B401", f.BoardPos, "b", f.Board)
		}
	}
	// 1. libraries, then every top-level name (functions, devices, ... may be used before their line)
	for _, it := range f.Items {
		if u, ok := it.(*UseDecl); ok {
			c.useDecl(u)
		}
	}
	for _, it := range f.Items {
		c.declare(it)
	}
	// 2. global variables in source order, then everything else
	for _, it := range f.Items {
		if v, ok := it.(*VarDecl); ok {
			c.globalVar(v)
		}
	}
	for _, it := range f.Items {
		c.topItem(it)
	}
	c.propagateWaits()
	// fast handlers, now that every expression's type is known (spec 6.5), and pins that libraries take over
	for _, it := range f.Items {
		if d, ok := it.(*EventBlock); ok && d.Fast {
			c.fastRules(d.Body)
		}
	}
	c.libraryClaims()
	c.timerOwners()
	c.wifiADC()
	c.regsConflicts()
	c.ramEstimate()
	c.recursionCheck()
	c.serialBudget()
	c.coreShareCheck()
	c.capabilityCheck()
	// 3. unused globals (a module's names are its exports: the importer uses them, not the module)
	if !c.inModule {
		for _, it := range f.Items {
			switch d := it.(type) {
			case *VarDecl:
				if s := c.global.syms[d.Name]; s != nil && !s.used && d.Kind != "saved" {
					c.errorf("W603", d.At, "name", d.Name)
				}
			case *DeviceDecl:
				if s := c.global.syms[d.Name]; s != nil && !s.used {
					c.errorf("W603", d.At, "name", d.Name)
				}
			}
		}
	}
	// Bare enum members that found their type left empty tombstones in errs; the rest keep their C301.
	kept := c.errs[:0]
	for _, e := range c.errs {
		if e.Code != "" {
			kept = append(kept, e)
		}
	}
	c.errs = kept
}

// propagateWaits finishes 'fn that waits' marking: a fn waits when it holds 'wait', calls a blocking built-in
// (marked during checking), or calls a fn that waits - and calling one outside a task is C311.
func (c *checker) propagateWaits() {
	changed := true
	for changed {
		changed = false
		for caller, callees := range c.waitDeps {
			if caller.waits {
				continue
			}
			for _, callee := range callees {
				if callee.waits {
					caller.waits = true
					changed = true
					break
				}
			}
		}
	}
	for _, wc := range c.waitCalls {
		if wc.sym.waits {
			c.errorf("C311", wc.at, "call", wc.callee)
		}
	}
}

func (c *checker) define(sc *scope, s *symbol) {
	if old, ok := sc.syms[s.name]; ok {
		c.errorf("C302", s.pos, "name", s.name, "line", strconv.Itoa(old.pos.Line))
		return
	}
	sc.syms[s.name] = s
}

// useDecl processes one 'use' line: Mujard libraries gate names, described C/C++ libraries check their
// calls, project files become quoted includes, and everything else goes through unchecked (W611).
// safeUsePath matches what a 'use "…"' may name. The text ends up in an #include, so it must never reach outside
// the project ('../../config.php', 'C:/…', '/etc/…') or carry characters that mean something to the preprocessor.
var safeUsePath = regexp.MustCompile(`^[A-Za-z0-9_][A-Za-z0-9_ .+/-]*$`)

func usePathOK(p string) bool {
	if !safeUsePath.MatchString(p) || strings.Contains(p, "//") {
		return false
	}
	for _, seg := range strings.Split(p, "/") {
		if seg == ".." || seg == "." {
			return false
		}
	}
	return true
}

func (c *checker) useDecl(u *UseDecl) {
	if (u.From == "arduino" || u.From == "idf" || u.From == "pio") && !usePathOK(u.Path) {
		c.errorf("C324", u.At, "path", u.Path)
		return
	}
	desc := c.env.findLib(u.Path)
	switch u.From {
	case "mujard":
		c.libs[strings.TrimPrefix(u.Path, "mujard.")] = true
	case "idf":
		if c.board != nil && c.board.Arch != "esp32" {
			c.errorf("B412", u.At, "lib", u.Path, "board", c.board.Display)
		}
		if desc != nil {
			c.addLib(desc, u)
		} else {
			c.unchecked = true
			c.errorf("W611", u.At, "lib", u.Path)
			c.libDeps = append(c.libDeps, LibDep{From: "idf", Name: u.Path, Version: u.Version})
		}
	case "pio":
		if desc != nil {
			c.addLib(desc, u)
		} else {
			c.unchecked = true
			c.errorf("W611", u.At, "lib", u.Path)
			c.libDeps = append(c.libDeps, LibDep{From: "pio", Name: u.Path, Version: u.Version, Header: headerFromName(u.Path)})
		}
	case "arduino":
		switch {
		case desc != nil:
			c.addLib(desc, u)
		case c.env.hasFile(u.Path) || strings.ContainsAny(u.Path, "/\\"):
			// A C/C++ file in the project ('use "lib/crc.h"'): included with quotes, not a library.
			if c.useFiles == nil {
				c.useFiles = map[*UseDecl]bool{}
			}
			c.useFiles[u] = true
		default:
			c.unchecked = true
			c.errorf("W611", u.At, "lib", u.Path)
			hdr := u.Path
			if !isFilePath(hdr) {
				hdr = headerFromName(hdr)
			}
			c.libDeps = append(c.libDeps, LibDep{From: "arduino", Name: u.Path, Version: u.Version, Header: hdr})
		}
	default: // "module" and anything else keeps its unchecked meaning
		c.unchecked = true
	}
}

// addLib registers a described library: its globals, classes and pins go into the checks, and the build
// gets told the library (and version) is needed.
func (c *checker) addLib(d *LibDesc, u *UseDecl) {
	if c.board != nil && len(d.Architectures) > 0 {
		ok := false
		for _, a := range d.Architectures {
			if a == c.board.Arch {
				ok = true
			}
		}
		if !ok {
			c.errorf("B412", u.At, "lib", d.Name, "board", c.board.Display)
		}
	}
	if c.useLibs == nil {
		c.useLibs = map[*UseDecl]*LibDesc{}
		c.libGlobals = map[string]*global{}
		c.libClasses = map[string]class{}
		c.libPins = map[int]string{}
		c.libReal = map[string]string{}
	}
	c.useLibs[u] = d
	for n, g := range d.globalsTable() {
		if _, taken := c.libGlobals[n]; !taken {
			c.libGlobals[n] = g
		}
	}
	for n, cl := range d.classTable() {
		if _, taken := c.libClasses[n]; !taken {
			c.libClasses[n] = cl
		}
	}
	for _, pin := range d.Pins {
		c.libPins[pin] = d.Name
	}
	for _, claim := range d.Takes {
		if c.claimUsers == nil {
			c.claimUsers = map[string]string{}
			c.claimAt = map[string]Pos{}
		}
		if _, taken := c.claimUsers[claim]; !taken {
			c.claimUsers[claim] = d.Name
			c.claimAt[claim] = u.At
		}
	}
	c.libDeps = append(c.libDeps, LibDep{From: u.From, Name: d.Name, Version: u.Version, Header: d.Header})
}

// libGlobal is a top-level name a described library in use offers. A member of a class family (U8g2's displays)
// is found by its name's shape the first time it is written, and then kept like any other class.
func (c *checker) libGlobal(name string) (*global, bool) {
	if g, ok := c.libGlobals[name]; ok {
		return g, true
	}
	for _, d := range c.useLibs {
		if g, cl, ok := d.familyGlobal(name); ok {
			c.libGlobals[name] = g
			if _, taken := c.libClasses[name]; !taken {
				c.libClasses[name] = cl
			}
			return g, true
		}
	}
	return nil, false
}

// libraryClaims warns about pins the board's buses and timers give to a library in use (spec 7): the I2C and SPI
// pins (W609) when mujard.i2c, spi or a described library that takes the bus is in the program, and on AVR boards
// the PWM pins whose timer a Servo library or tone() takes (W613). The pins come from the board database (claims).
func (c *checker) libraryClaims() {
	if c.board == nil || len(c.board.Claims) == 0 {
		return
	}
	users := map[string]string{} // claim -> the library to name in the message
	for claim, lib := range c.claimUsers {
		users[claim] = lib
	}
	for _, lib := range []string{"i2c", "spi", "servo"} {
		if c.libs[lib] {
			users[lib] = "mujard." + lib
		}
	}
	if c.libs["rtc"] {
		if _, taken := users["i2c"]; !taken {
			users["i2c"] = "mujard.rtc" // the clock chip (DS3231) sits on the I2C pins
		}
	}
	if c.usesTone {
		users["tone"] = "tone"
	}
	for _, it := range c.file.Items {
		d, ok := it.(*DeviceDecl)
		if !ok || d.Kind == "object" {
			continue
		}
		for _, r := range c.declPins(d) {
			n, ok := c.board.ResolvePin(r.Text)
			if !ok {
				continue
			}
			label := r.Text
			if c.board.IsESP() {
				label = c.board.PinLabel(n)
			}
			for claim, lib := range users {
				if !has(c.board.Claims[claim], n) {
					continue
				}
				if claim == "servo" || claim == "tone" {
					if d.Role == "pwm" {
						c.errorf("W613", r.At, "p", label, "board", c.board.Display, "lib", lib)
					}
				} else {
					c.errorf("W609", r.At, "p", label, "lib", lib)
				}
			}
		}
	}
}

// timerOwners reports two users of one hardware timer (B414): each brings that timer's interrupt, so the linker
// would stop with "multiple definition of __vector_7". A timer has one owner, like a pin (B407).
func (c *checker) timerOwners() {
	if c.board == nil {
		return
	}
	type owner struct {
		who string
		at  Pos
	}
	var owners []owner
	timerOf := map[string]string{}
	add := func(claim, who string, at Pos) {
		if t := c.board.TimerOf(claim); t != "" {
			owners = append(owners, owner{who, at})
			timerOf[who] = t
		}
	}
	if c.usesTone {
		add("tone", "tone", c.toneAt)
	}
	if c.usesTimer {
		add("timer", "Timer", c.timerAt)
	}
	// AVR pwm pins sit on named timers: pin 9/10 on an Uno is Timer 1, which Timer() or Servo may already own. ESP32
	// ledc and RP2040 slices allocate at runtime, so there's no static claim to make for them (PwmTimerOf is AVR-only).
	for _, it := range c.file.Items {
		d, ok := it.(*DeviceDecl)
		if !ok || d.Role != "pwm" {
			continue
		}
		for _, p := range c.declPins(d) {
			n, okN := c.board.ResolvePin(p.Text)
			if !okN {
				continue
			}
			if t := c.board.PwmTimerOf(n); t != "" {
				who := fmt.Sprintf("pwm %d", n)
				owners = append(owners, owner{who, d.At})
				timerOf[who] = t
			}
		}
	}
	if c.timerCount > 1 {
		c.errorf("B414", c.timerAt, "a", "Timer", "b", "Timer", "board", c.board.Display, "timer", "a hardware timer")
	}
	for claim, lib := range c.claimUsers {
		add(claim, lib, c.claimAt[claim])
	}
	sort.Slice(owners, func(i, j int) bool { // the later one is the one to report
		a, b := owners[i].at, owners[j].at
		return a.Line < b.Line || a.Line == b.Line && a.Col < b.Col
	})
	first := map[string]owner{} // timer -> its first owner
	for _, o := range owners {
		t := timerOf[o.who]
		if f, taken := first[t]; taken {
			c.errorf("B414", o.at, "a", f.who, "b", o.who, "board", c.board.Display, "timer", t)
			continue
		}
		first[t] = o
	}
}

// noteShare records a global variable touched by a block: written or read on a secondary core (core>0) or on the
// main core. coreShareCheck turns core-write + main-use and core-read + main-write into W620.
func (c *checker) noteShare(sym *symbol, at Pos, core int, write bool) {
	if sym.t != nil && nonOpt(sym.t).Name == "queue" {
		return // a Queue is the sanctioned way to share across cores; using it isn't a hazard to warn about
	}
	if core > 0 {
		if write {
			if c.coreWrites == nil {
				c.coreWrites = map[*symbol]Pos{}
			}
			if _, ok := c.coreWrites[sym]; !ok {
				c.coreWrites[sym] = at
			}
			return
		}
		if c.coreReads == nil {
			c.coreReads = map[*symbol]Pos{}
		}
		if _, ok := c.coreReads[sym]; !ok {
			c.coreReads[sym] = at
		}
		return
	}
	if write {
		if c.mainWrites == nil {
			c.mainWrites = map[*symbol]bool{}
		}
		c.mainWrites[sym] = true
		return
	}
	if c.mainUses == nil {
		c.mainUses = map[*symbol]bool{}
	}
	c.mainUses[sym] = true
}

// queueCtor checks q = Queue(elem, capacity): the element is one of the value types (int, float, bool, text or a
// measurement) and capacity a literal from 1 to 64. It runs outside the normal arg machinery because the first
// argument names a type ('text' in Queue(text, 8)), not a value - and the result carries both into the var's type
// so q.push/q.pop specialize on it in 'call'.
func (c *checker) queueCtor(e *Call, sc *scope, cx ctx) *Type {
	if len(e.Args) != 2 {
		c.errorf("C307", e.At, "fn", "Queue", "n", "2", "m", strconv.Itoa(len(e.Args)))
		for _, a := range e.Args {
			c.expr(a.Value, sc, cx)
		}
		return tAny
	}
	el := tAny
	want := "a value type (int, float, bool, text or a measurement)"
	if id, ok := e.Args[0].Value.(*Ident); ok {
		if t, okT := typeNamed(id.Name); okT {
			el = t
		} else {
			c.errorf("C303", id.At, "want", want, "got", id.Name)
		}
	} else {
		c.errorf("C303", e.Args[0].Value.Start(), "want", want, "got", nodeName(e.Args[0].Value))
		c.expr(e.Args[0].Value, sc, cx)
	}
	cap := 8
	if n, ok := e.Args[1].Value.(*NumberLit); ok && !n.IsFloat && n.Unit == "" {
		if n.Int >= 1 && n.Int <= 64 {
			cap = int(n.Int)
		} else {
			c.errorf("C303", n.At, "want", "a queue capacity from 1 to 64", "got", n.Text)
		}
	} else {
		c.errorf("C303", e.Args[1].Value.Start(), "want", "a queue capacity from 1 to 64", "got", nodeName(e.Args[1].Value))
		c.expr(e.Args[1].Value, sc, cx)
	}
	return &Type{K: TObject, Name: "queue", Elem: el, Cap: cap}
}

// rpFamily reports the RP2040/RP2350 boards: the two-core family whose drivers and flash rules differ from
// the ESP32 family's FreeRTOS-managed ones.
func (c *checker) rpFamily() bool {
	return c.board != nil && (c.board.Family == "RP2040" || c.board.Family == "RP2350")
}

// freqArg returns the literal Hz value a tone()/freq call asked for (440 Hz -> 440, 2 kHz -> 2000); a computed or
// non-literal argument can't be compared statically and returns false - no warning on what can't be proven.
func (c *checker) freqArg(e Node) (float64, bool) {
	n, ok := e.(*NumberLit)
	if !ok || n.IsFloat && n.Float == 0 {
		return 0, false
	}
	v := n.Float
	if !n.IsFloat {
		v = float64(n.Int)
	}
	switch n.Unit {
	case "Hz":
	case "kHz":
		v *= 1e3
	case "MHz":
		v *= 1e6
	default:
		return 0, false
	}
	return v, true
}

// toneSlice records a tone() call's pin and frequency on RP-family boards: PWM slice (pin>>1)&7 has one frequency
// for both channels, so a second tone() on the same slice with a provably different pitch warns W623. Same pin or
// same frequency repeats fine.
func (c *checker) toneSlice(slice, pin int, hz float64, at Pos) {
	if c.sliceTones == nil {
		c.sliceTones = map[int]struct {
			pin int
			hz  float64
		}{}
	}
	prev, seen := c.sliceTones[slice]
	if !seen || prev.pin == pin || prev.hz == hz {
		c.sliceTones[slice] = struct {
			pin int
			hz  float64
		}{pin, hz}
		return
	}
	c.errorf("W623", at, "a", fmt.Sprintf("pin %d", prev.pin), "b", fmt.Sprintf("pin %d", pin),
		"slice", strconv.Itoa(slice), "board", c.board.Display,
		"af", fmt.Sprintf("%g Hz", prev.hz), "bf", fmt.Sprintf("%g Hz", hz))
}

// shareRecv marks a method-call receiver as a cross-core write when it's a global object: 'call' invokes it for
// member calls (t.stop(), u.write("x")); 'assign' invokes it for member targets (store.x = v). A read of the same
// name through ident() is already recorded separately, so reads and writes land in the right sets.
func (c *checker) shareRecv(e Node, sc *scope, at Pos, core int) {
	id, ok := e.(*Ident)
	if !ok {
		return
	}
	sym, _ := sc.lookup(id.Name)
	if sym != nil && sym.kind == symVar && c.global.syms[id.Name] == sym && nonOpt(sym.t).K == TObject {
		c.noteShare(sym, at, core, true)
	}
}

// coreShareCheck warns (W620) that a global a secondary core touches while the main core touches it too isn't
// synchronized: a read can see a half-written value, and compound writes (n = n + 1) race. Boards have at most one
// secondary core, so 'touched on a core, touched on main' is the whole hazard surface.
func (c *checker) coreShareCheck() {
	if c.board == nil || c.board.Cores() <= 1 {
		return
	}
	for sym, at := range c.coreWrites {
		if c.mainUses[sym] || c.mainWrites[sym] {
			c.errorf("W620", at, "name", sym.name)
		}
	}
	for sym, at := range c.coreReads {
		if c.mainWrites[sym] {
			c.errorf("W620", at, "name", sym.name)
		}
	}
}

// wifiADC warns (W615) about analog reads on the pins Wi-Fi takes over (the ESP32 family's ADC2): while the radio
// is on, those reads stop working. It runs after the whole program is checked, so 'use' lines and network calls
// anywhere in the file count.
func (c *checker) wifiADC() {
	b := c.board
	if b == nil || !c.usesWifi || len(b.ADCBlockedByWifi) == 0 {
		return
	}
	var free []int
	for _, p := range b.ADC {
		if !has(b.ADCBlockedByWifi, p) {
			free = append(free, p)
		}
	}
	for _, u := range c.analogPins {
		if has(b.ADCBlockedByWifi, u.pin) {
			c.errorf("W615", u.at, "p", u.label, "board", b.Display, "list", b.PinList(free))
		}
	}
}

// regsCall watches register method calls (spec 10.4): a literal too wide for the register is an error, and a
// write/set/clear on a GPIO port register is recorded for regsConflicts.
func (c *checker) regsCall(e *Call, reg, name string) {
	table, ok := regsOf(c.board)
	if !ok {
		return
	}
	r, ok := table[reg]
	if !ok {
		return
	}
	if name == "write" || name == "set" || name == "clear" {
		if len(e.Args) > 0 {
			if n, isLit := e.Args[0].Value.(*NumberLit); isLit && !n.IsFloat && n.Unit == "" && n.Int > regMax(r) {
				c.errorf("C327", e.Args[0].Value.Start(), "n", strconv.FormatUint(n.Int, 10), "reg", reg, "bits", strconv.Itoa(r.Bits))
			}
		}
		if len(r.Port) > 0 {
			c.regsWrites = append(c.regsWrites, regWrite{at: e.At, reg: reg, pins: r.Port})
		}
	}
}

// regsConflicts warns (W616) when a register write/set/clear reaches a GPIO port whose pins the program
// declared - the write changes the pin behind the device's back (spec 10.4).
func (c *checker) regsConflicts() {
	for _, w := range c.regsWrites {
		var hit []int
		for _, p := range w.pins {
			if _, declared := c.pins[p]; declared {
				hit = append(hit, p)
			}
		}
		if len(hit) > 0 && c.board != nil {
			sort.Ints(hit)
			c.errorf("W616", w.at, "reg", w.reg, "list", c.board.PinList(hit))
		}
	}
}

// declPins lists every pin a device declaration names, a 'pins a..b' range one by one.
func (c *checker) declPins(d *DeviceDecl) []PinRef {
	if d.Kind != "pins" || len(d.Pins) != 2 {
		return d.Pins
	}
	a, okA := c.board.ResolvePin(d.Pins[0].Text)
	b, okB := c.board.ResolvePin(d.Pins[1].Text)
	if !okA || !okB {
		return nil
	}
	if d.ToEnd {
		b++
	}
	var out []PinRef
	for n := a; n < b; n++ {
		out = append(out, PinRef{At: d.Pins[0].At, Text: strconv.Itoa(n)})
	}
	return out
}

func (c *checker) declare(it Node) {
	switch d := it.(type) {
	case *ImportDecl:
		c.define(c.global, &symbol{kind: symModule, name: d.Name, pos: d.At, used: true})
		c.importModule(d)
	case *VarDecl:
		k := symVar
		if d.Kind == "const" {
			k = symConst
		}
		c.define(c.global, &symbol{kind: k, name: d.Name, pos: d.At, decl: d, t: tAny})
	case *DeviceDecl:
		c.define(c.global, &symbol{kind: symDevice, name: d.Name, pos: d.At, t: c.device(d)})
	case *FnDecl:
		c.define(c.global, &symbol{kind: symFn, name: d.Name, pos: d.At, fn: d, waits: containsWait(d.Body)})
	case *StructDecl:
		c.define(c.global, &symbol{kind: symStruct, name: d.Name, pos: d.At, strukt: d, t: &Type{K: TStruct, Name: c.qual(d.Name)}})
	case *EnumDecl:
		c.define(c.global, &symbol{kind: symEnum, name: d.Name, pos: d.At, enum: d, t: &Type{K: TEnum, Name: c.qual(d.Name)}})
	case *TaskDecl:
		c.define(c.global, &symbol{kind: symBlock, name: d.Name, pos: d.At, t: object("block")})
	case *EventBlock:
		if d.As != "" && d.Kind != "on" {
			c.define(c.global, &symbol{kind: symBlock, name: d.As, pos: d.At, t: object("block")})
		}
	}
}

// coreCheck gates 'on core N' on a top-level every/after/when block: N is a secondary core the board has
// ('on core 0' is the main core - the default, not an error). The block running there is the program's
// multicore requirement; a board with too few cores refuses.
func (c *checker) coreCheck(d *EventBlock) {
	if c.board.Cores() <= d.Core {
		c.errorf("B415", d.At, "who", fmt.Sprintf("'%s' on core %d", d.Kind, d.Core), "cap", "multicore", "board", c.board.Display)
	}
}

// qualify gives a module's type its importing name: Point in motors.muj is 'motors.Point'.
func (c *checker) qual(name string) string {
	if c.inModule {
		return c.modName + "." + name
	}
	return name
}

// typeSym finds the symbol a named type belongs to - in the program, or in a module when the name is 'mod.name'.
func (c *checker) typeSym(name string) *symbol {
	if i := strings.Index(name, "."); i > 0 {
		if mc := c.mods[name[:i]]; mc != nil {
			return mc.global.syms[name[i+1:]]
		}
		if c.inModule && strings.HasPrefix(name, c.modName+".") {
			return c.global.syms[name[len(c.modName)+1:]]
		}
		return nil
	}
	return c.global.syms[name]
}

// importModule loads and checks <name>.muj: a module. Its declarations become reachable as name.member; the file
// itself is a program, so it is checked fully (its blocks don't run for the importer - spec 10.3).
func (c *checker) importModule(d *ImportDecl) {
	if d.Name == "" {
		return // the parser already reported the missing name
	}
	if c.inModule {
		c.errorf("C322", d.At, "what", "import")
		return
	}
	if c.load == nil {
		c.errorf("C320", d.At, "feature", "import", "phase", phaseLibs)
		return
	}
	mf, mdiags, file := c.load(d.Name)
	if file == "" {
		c.errorf("C321", d.At, "name", d.Name, "file", d.Name+SourceExts[0])
		return
	}
	for i := range mdiags {
		mdiags[i].File = file
		if len(c.errs) < maxDiagnostics {
			c.errs = append(c.errs, mdiags[i])
		}
	}
	if mf == nil {
		return
	}
	unbound := c.bindModulePins(d, mf)
	// A module that names its board keeps it; a boardless one is checked against the program's board so its pins
	// and the program's can't collide (B407).
	mb := c.board
	if mf.Board != "" {
		mb = LookupBoard(mf.Board)
	}
	pins := c.pins
	if mb != nil && c.board != nil && mb.Name != c.board.Name {
		pins = map[int]pinUse{}
	}
	mc := &checker{file: mf, board: mb, global: newScope(nil), libs: map[string]bool{}, pins: pins,
		types: map[Node]*Type{}, declared: map[*Ident]*Type{}, inModule: true, modName: d.Name, modFile: file,
		pendingEnum: map[*Ident]pendingEnum{}, env: c.env, unboundPins: unbound}
	mc.run()
	mc.failed = FirstError(mdiags) != nil || FirstError(mc.errs) != nil
	for i := range mc.errs {
		mc.errs[i].File = file
		if len(c.errs) < maxDiagnostics {
			c.errs = append(c.errs, mc.errs[i])
		}
	}
	for lib := range mc.libs {
		c.libs[lib] = true
	}
	c.unchecked = c.unchecked || mc.unchecked
	c.mods[d.Name] = mc
}

// modMember resolves name.member against a module's own top-level names.
func (c *checker) modMember(e *Member, mod string) *Type {
	mc := c.mods[mod]
	if mc == nil || mc.failed {
		return tAny
	}
	sym := mc.global.syms[e.Name]
	if sym == nil || sym.kind == symModule || sym.kind == symBlock || sym.kind == symParam {
		c.errorf("C308", e.At, "type", "module '"+mod+"'", "member", e.Name)
		return tAny
	}
	sym.used = true
	switch sym.kind {
	case symFn:
		return &Type{K: TObject, Name: "fn:" + mod + "." + sym.name}
	case symStruct:
		return &Type{K: TObject, Name: "struct:" + sym.t.Name}
	case symEnum:
		return &Type{K: TObject, Name: "enum:" + sym.t.Name}
	}
	return sym.t
}

func (c *checker) globalVar(v *VarDecl) {
	s := c.global.syms[v.Name]
	if s == nil || s.checked {
		return
	}
	s.checked = true
	s.t = c.varType(v, c.global, ctx{})
	if v.Kind == "saved" && !storable(s.t) {
		c.errorf("C312", v.At)
	}
}

// storable reports whether a value can be kept in EEPROM / flash by 'saved var' or store.save (C312).
func storable(t *Type) bool {
	return isNumeric(t) || t.K == TBool || t.K == TQty || t.K == TEnum || t.K == TText || t.K == TAny
}

// varType works out a declaration's type from its written type and its value.
func (c *checker) varType(v *VarDecl, sc *scope, cx ctx) *Type {
	var want *Type
	if v.Type != nil {
		want = c.resolveType(v.Type)
	}
	got := c.expr(v.Value, sc, cx)
	if want == nil {
		if got.K == TText && !got.Opt {
			c.warnText(v.Value, textMax)
		}
		if got.K == TNone {
			return opt(tAny)
		}
		if got.K == TVoid {
			c.errorf("C303", v.Value.Start(), "want", "a value", "got", "nothing")
			return tAny
		}
		return got
	}
	if lit, ok := v.Value.(*ListLit); ok && len(lit.Items) == 0 && want.K == TList {
		return want
	}
	c.expect(v.Value, got, want)
	return want
}

func (c *checker) resolveType(t *TypeExpr) *Type {
	var out *Type
	if bt, ok := typeNamed(t.Name); ok {
		out = bt
		if t.Name == "text" && len(t.Args) > 0 {
			n, err := strconv.Atoi(t.Args[0].Name)
			if len(t.Args) != 1 || err != nil || n < 1 || n > 255 {
				c.errorf("C303", t.Args[0].At, "want", "a text size from 1 to 255", "got", t.Args[0].Name)
			} else {
				out = textCap(n)
			}
		}
	} else if t.Name == "list" {
		el := tAny
		if len(t.Args) > 0 {
			el = c.resolveType(t.Args[0])
		}
		out = listOf(el)
	} else if s, ok := c.global.syms[t.Name]; ok && (s.kind == symStruct || s.kind == symEnum) {
		out = s.t
	} else if en, ok := builtinEnums[t.Name]; ok {
		out = en.t // var r: ResetReason = ...
	} else if _, err := strconv.Atoi(t.Name); err == nil {
		out = tAny
	} else {
		c.errorf("C301", t.At, "name", t.Name)
		out = tAny
	}
	if t.Optional {
		out = opt(out)
	}
	return out
}

func (c *checker) topItem(it Node) {
	cx := ctx{}
	switch d := it.(type) {
	case *FnDecl:
		c.fnDecl(d)
	case *EventBlock:
		if d.Core > 0 {
			c.coreCheck(d)
		}
		c.eventBlock(d, c.global, cx)
	case *TaskDecl:
		if d.Core > 0 && c.board.Cores() <= d.Core {
			c.errorf("B415", d.At, "who", fmt.Sprintf("task %s on core %d", d.Name, d.Core), "cap", "multicore", "board", c.board.Display)
		}
		cx.inTask = true
		cx.block = "task"
		cx.core = d.Core
		c.block(d.Body, newScope(c.global), cx)
	case *SpecialBlock:
		if d.Kind == "test" {
			cx.inTest, cx.inTask = true, true
		} else {
			cx.inSafe = true
		}
		cx.block = d.Kind
		c.block(d.Body, newScope(c.global), cx)
	case *LimitDecl:
		t := c.expr(d.Target, c.global, cx)
		r := c.expr(d.Range, c.global, cx)
		switch {
		case t.K == TDevice && t.Name == "pwm":
			c.expect(d.Range, r, rangeOf(tRatio))
		case t.K == TObject && t.Name == "servo":
			c.expect(d.Range, r, rangeOf(tAngle))
		}
	case *ControlDecl:
		for i, name := range []string{d.Var, d.Var2} {
			if i == 1 && name == "" {
				break
			}
			s, _ := c.global.lookup(name)
			if s == nil {
				c.errorf("C301", d.At, "name", name)
				continue
			}
			s.used = true
			if s.kind != symVar {
				c.errorf("C306", d.At, "name", name)
			}
			switch {
			case d.Hold && s.t.K != TBool: // a hold button is pressed or not: true/false
				c.errorf("C303", d.At, "want", "a true/false variable (a hold button)", "got", describe(s.t))
			case d.Var2 != "" && (s.t.K == TBool || !(isNumeric(s.t) || s.t.K == TQty)): // a stick's axes are numbers
				c.errorf("C303", d.At, "want", "a number variable (a joystick axis)", "got", describe(s.t))
			case d.Range != nil:
				c.expect(d.Range, c.expr(d.Range, c.global, cx), rangeOf(s.t))
			}
		}
	case *NeedsDecl:
		c.needsDecl(d)
	case *CppBlock:
		c.errorf("W604", d.At)
		c.unchecked = true
	case *StructDecl:
		for _, f := range d.Fields {
			c.resolveType(f.Type)
		}
	case *DeviceDecl:
		if d.Kind == "object" {
			s := c.global.syms[d.Name]
			if s != nil {
				c.pinOwner = d // lid = Servo(pin 6): pin 6 belongs to lid (B407 names it)
				s.t = c.expr(d.Object, c.global, cx)
				c.pinOwner = nil
			}
			c.uartSpeed(d)
		}
	}
}

// softSerialBoards have no spare hardware serial port: a Uart there is SoftwareSerial (mujard_uart.h).
var softSerialBoards = map[string]bool{"uno": true, "nano": true}

// uartSpeed warns about a fast Uart that can only be a software port (W612).
func (c *checker) uartSpeed(d *DeviceDecl) {
	call, ok := d.Object.(*Call)
	if !ok || c.board == nil || !softSerialBoards[c.board.Name] {
		return
	}
	if id, ok := call.Fn.(*Ident); !ok || id.Name != "Uart" {
		return
	}
	for i, a := range call.Args {
		if a.Name != "speed" && !(a.Name == "" && i == 2) {
			continue
		}
		if n, ok := a.Value.(*NumberLit); ok && ((n.IsFloat && n.Float > 9600) || (!n.IsFloat && n.Int > 9600)) {
			c.errorf("W612", a.Value.Start(), "board", c.board.Display)
		}
	}
}

// ---- devices and pins

func (c *checker) device(d *DeviceDecl) *Type {
	switch d.Kind {
	case "object":
		return tAny // typed in topItem, once every name is known
	case "analog":
		c.pin(d.Pins[0], "analog", d)
		return &Type{K: TDevice, Name: "analog"}
	case "pin":
		c.pin(d.Pins[0], d.Role, d)
		return &Type{K: TDevice, Name: d.Role}
	}
	// groups
	var refs []PinRef
	if d.Kind == "pins" && c.board != nil {
		a, okA := c.board.ResolvePin(d.Pins[0].Text)
		b, okB := c.board.ResolvePin(d.Pins[1].Text)
		if !okA {
			c.errorf("B402", d.Pins[0].At, "p", d.Pins[0].Text, "board", c.board.Display)
		}
		if !okB {
			c.errorf("B402", d.Pins[1].At, "p", d.Pins[1].Text, "board", c.board.Display)
		}
		if okA && okB {
			if d.ToEnd {
				b++
			}
			for n := a; n < b; n++ {
				refs = append(refs, PinRef{At: d.Pins[0].At, Text: strconv.Itoa(n)})
			}
		}
	} else {
		refs = d.Pins
	}
	for _, r := range refs {
		c.pin(r, d.Role, d)
	}
	return &Type{K: TDevice, Name: "group", Elem: &Type{K: TDevice, Name: d.Role}}
}

// pin checks one pin of a device against the board (spec 7) and records who uses it.
func (c *checker) pin(r PinRef, role string, d *DeviceDecl) {
	b := c.board
	if b == nil || c.unboundPins[r.Text] {
		return
	}
	n, ok := b.ResolvePin(r.Text)
	if !ok {
		c.errorf("B402", r.At, "p", r.Text, "board", b.Display)
		return
	}
	label := r.Text
	if b.IsESP() {
		label = b.PinLabel(n)
	}
	if role == "analog" {
		c.analogPins = append(c.analogPins, analogPin{at: r.At, pin: n, label: label})
	}
	switch {
	case has(b.FlashPins, n):
		c.errorf("B406", r.At, "p", label, "board", b.Display)
	case role == "analog" && !b.IsAnalogInput(n):
		c.errorf("B404", r.At, "p", label, "board", b.Display, "list", b.PinList(b.AnalogPins()))
	case role != "analog" && has(b.AnalogOnly, n):
		c.errorf("B408", r.At, "p", label, "board", b.Display)
	case (role == "output" || role == "pwm") && has(b.InputOnly, n):
		c.errorf("B405", r.At, "p", label, "board", b.Display)
	case role == "pwm" && !has(b.PWM, n):
		c.errorf("B403", r.At, "p", label, "board", b.Display, "list", b.PinList(b.PWM))
	}
	if d != nil {
		for _, o := range d.Options {
			if (o == "pullup" || o == "pulldown") && role != "input" {
				c.errorf("B410", d.RoleAt)
			}
		}
	}
	if u, taken := c.pins[n]; taken {
		c.errorf("B407", r.At, "p", label, "name", u.name, "line", strconv.Itoa(u.line))
	} else {
		name := "pin " + r.Text // a pin used as a bare value names itself
		if role == "analog" {
			name = "analog " + r.Text
		}
		if d != nil {
			name = d.Name
		}
		c.pins[n] = pinUse{name: name, line: r.At.Line}
	}
	if has(b.USBSerialPins, n) {
		c.errorf("W601", r.At, "board", b.Display)
	}
	if has(b.Strapping, n) {
		c.errorf("W602", r.At, "p", label, "board", b.Display)
	}
	if lib, taken := c.libPins[n]; taken {
		c.errorf("W609", r.At, "p", label, "lib", lib)
	}
}

// ---- functions and blocks

func (c *checker) fnDecl(d *FnDecl) {
	sc := newScope(c.global)
	sym := c.global.syms[d.Name]
	cx := ctx{fn: d, fnSym: sym, inTask: true} // wait is allowed here; calls from outside tasks are checked (C311)
	for _, prm := range d.Params {
		t := c.resolveType(prm.Type)
		if prm.Default != nil {
			c.expect(prm.Default, c.expr(prm.Default, c.global, ctx{}), t)
		}
		c.define(sc, &symbol{kind: symParam, name: prm.Name, pos: prm.At, t: t, used: true})
	}
	c.block(d.Body, sc, cx)
	if d.Result != nil && !alwaysReturns(d.Body) {
		c.errorf("C313", d.At, "fn", d.Name)
	}
}

func alwaysReturns(b *Block) bool {
	if b == nil || len(b.Stmts) == 0 {
		return false
	}
	switch s := b.Stmts[len(b.Stmts)-1].(type) {
	case *SimpleStmt:
		return s.Kind == "return"
	case *IfStmt:
		if s.Else == nil || !alwaysReturns(s.Else) {
			return false
		}
		for _, t := range s.Thens {
			if !alwaysReturns(t) {
				return false
			}
		}
		return true
	case *MatchStmt:
		hasElse := false
		for _, mc := range s.Cases {
			if !alwaysReturns(mc.Body) {
				return false
			}
			hasElse = hasElse || mc.IsElse
		}
		return hasElse
	}
	return false
}

// exits reports whether a block always leaves (return / break / continue / board.halt), for none narrowing after an if.
func exits(b *Block) bool {
	if b == nil || len(b.Stmts) == 0 {
		return false
	}
	switch s := b.Stmts[len(b.Stmts)-1].(type) {
	case *SimpleStmt:
		return s.Kind == "return" || s.Kind == "break" || s.Kind == "continue"
	case *ExprStmt:
		if call, ok := s.X.(*Call); ok {
			if mem, ok := call.Fn.(*Member); ok && mem.Name == "halt" {
				if id, ok := mem.X.(*Ident); ok && id.Name == "board" {
					return true
				}
			}
		}
	}
	return false
}

func containsWait(b *Block) bool {
	found := false
	var walk func(n Node)
	walk = func(n Node) {
		switch s := n.(type) {
		case *SimpleStmt:
			if s.Kind == "wait" {
				found = true
			}
		case *Block:
			if s != nil {
				for _, x := range s.Stmts {
					walk(x)
				}
			}
		case *IfStmt:
			for _, t := range s.Thens {
				walk(t)
			}
			if s.Else != nil {
				walk(s.Else)
			}
		case *WhileStmt:
			walk(s.Body)
		case *ForStmt:
			walk(s.Body)
		case *MatchStmt:
			for _, mc := range s.Cases {
				walk(mc.Body)
			}
		}
	}
	walk(b)
	return found
}

func (c *checker) eventBlock(d *EventBlock, sc *scope, cx ctx) {
	cx.block = d.Kind
	inner := newScope(sc)
	switch d.Kind {
	case "every":
		t := c.expr(d.Expr, sc, cx)
		if !(t.K == TQty && (t.Dim == "time" || t.Dim == "frequency")) {
			c.expect(d.Expr, t, tTime)
		}
	case "after":
		c.expect(d.Expr, c.expr(d.Expr, sc, cx), tTime)
		inner.barrier = sc != c.global
	case "when":
		c.expect(d.Expr, c.expr(d.Expr, sc, cx), tBool)
		f, _ := facts(d.Expr)
		inner = inner.narrow(f)
	case "time": // at 07:30: needs the clock, set over Wi-Fi or read from an RTC (use mujard.rtc)
		c.libs["clock"] = true
		if c.board != nil && !c.board.Wifi && !c.libs["rtc"] {
			c.errorf("B409", d.At, "feature", "at HH:MM", "what", "Wi-Fi to set the clock or use mujard.rtc", "board", c.board.Display)
		}
	case "on":
		val := c.eventRef(d, sc, cx)
		if d.As != "" {
			c.define(inner, &symbol{kind: symVar, name: d.As, pos: d.At, t: val, used: true})
		}
		cx.inFast = d.Fast
	}
	cx.core = d.Core // only every/after/when can carry 'on core N'; the rest stay 0
	c.block(d.Body, inner, cx)
}

// eventRef checks what follows 'on': a device or object event, maybe with an argument (held(2 s)).
func (c *checker) eventRef(d *EventBlock, sc *scope, cx ctx) *Type {
	ref := d.Expr
	var arg Node
	if call, ok := ref.(*Call); ok {
		ref = call.Fn
		if len(call.Args) > 0 {
			arg = call.Args[0].Value
		}
	}
	mem, ok := ref.(*Member)
	if !ok {
		c.errorf("C303", d.Expr.Start(), "want", "an event, like button.pressed", "got", "something else")
		return tAny
	}
	ot := c.expr(mem.X, sc, cx)
	cl := c.classOf(ot)
	if cl == nil {
		return tAny // a library object: trusted
	}
	mm, found := cl[mem.Name]
	if !found || !mm.event {
		c.errorf("C308", mem.At, "type", bare(ot), "member", mem.Name)
		return tAny
	}
	if mm.eventArg != nil {
		if arg == nil {
			c.errorf("C307", mem.At, "fn", mem.Name, "n", "1", "m", "0")
		} else {
			c.expect(arg, c.expr(arg, sc, cx), mm.eventArg)
		}
	}
	if d.Fast {
		if ot.K == TObject && ot.Name == "timer" {
			// 'on t.tick fast': a hardware timer ISR - the body's fast rules are checked as usual
		} else if ot.K != TDevice || ot.Name != "input" {
			c.errorf("C314", d.At, "what", "'fast' on "+describe(ot))
		} else if id, ok := mem.X.(*Ident); ok && c.board != nil {
			if pin, ok := c.devicePin(id.Name); ok && !c.board.CanInterrupt(pin) {
				var ints []int
				for _, p := range c.board.ValidPins {
					if c.board.CanInterrupt(p) {
						ints = append(ints, p)
					}
				}
				c.errorf("B411", d.At, "p", c.board.PinLabel(pin), "board", c.board.Display, "list", c.board.PinList(ints))
			}
		}
	}
	if mm.eventValue != nil {
		return mm.eventValue
	}
	return tAny
}

// devicePin finds the pin number of a single-pin device declared at the top level.
func (c *checker) devicePin(name string) (int, bool) {
	for _, it := range c.file.Items {
		if d, ok := it.(*DeviceDecl); ok && d.Name == name && d.Kind == "pin" {
			return c.board.ResolvePin(d.Pins[0].Text)
		}
	}
	return 0, false
}

func (c *checker) block(b *Block, sc *scope, cx ctx) {
	if b == nil {
		return
	}
	for _, s := range b.Stmts {
		c.stmt(s, sc, cx)
	}
	for name, sym := range sc.syms {
		if !sym.used && sym.kind == symVar {
			c.errorf("W603", sym.pos, "name", name)
		}
	}
}

func (c *checker) stmt(n Node, sc *scope, cx ctx) {
	switch s := n.(type) {
	case *VarDecl:
		k := symVar
		if s.Kind == "const" {
			k = symConst
		}
		t := c.varType(s, sc, cx)
		c.define(sc, &symbol{kind: k, name: s.Name, pos: s.At, t: t})
	case *AssignStmt:
		c.assign(s, sc, cx)
	case *ExprStmt:
		c.expr(s.X, sc, cx)
	case *SimpleStmt:
		c.simple(s, sc, cx)
	case *ExpectStmt:
		c.expect(s.Cond, c.expr(s.Cond, sc, cx), tBool)
		positional := 0
		for _, a := range s.Args {
			t := c.expr(a.Value, sc, cx)
			switch {
			case a.Name == "":
				positional++
				if positional > 1 {
					c.errorf("C307", a.Value.Start(), "fn", "expect", "n", "2", "m", strconv.Itoa(len(s.Args)+1))
				}
			case a.Name == "halt":
				if lit, ok := a.Value.(*Literal); !ok || (lit.Value != "true" && lit.Value != "false") {
					c.errorf("C303", a.Value.Start(), "want", "true or false", "got", describe(t))
				}
			default:
				c.errorf("C308", a.Value.Start(), "type", "expect", "member", a.Name)
			}
		}
	case *IfStmt:
		c.ifStmt(s, sc, cx)
	case *WhileStmt:
		c.expect(s.Cond, c.expr(s.Cond, sc, cx), tBool)
		if lit, ok := s.Cond.(*Literal); ok && lit.Value == "true" && !cx.inTask {
			c.errorf("W610", s.At)
		}
		cx.loops++
		c.block(s.Body, newScope(sc), cx)
	case *ForStmt:
		it := c.expr(s.Iter, sc, cx)
		el := tAny
		switch {
		case it.K == TRange, it.K == TList:
			el = elemOf(it)
		case it.K == TDevice && it.Name == "group":
			el = it.Elem
		case it.K != TAny:
			c.errorf("C303", s.Iter.Start(), "want", "a range or a list", "got", describe(it))
		}
		inner := newScope(sc)
		c.define(inner, &symbol{kind: symVar, name: s.Var, pos: s.At, t: el, used: true})
		cx.loops++
		c.block(s.Body, inner, cx)
	case *MatchStmt:
		c.match(s, sc, cx)
	case *EventBlock:
		c.eventBlock(s, sc, cx)
	case *CppBlock:
		c.errorf("W604", s.At)
	}
}

func (c *checker) simple(s *SimpleStmt, sc *scope, cx ctx) {
	switch s.Kind {
	case "print":
		c.expr(s.Value, sc, cx)
		c.warnText(s.Value, textMax) // a line prints through a default-capacity text
		c.secrets(s.Value)
	case "wait":
		if !cx.inTask {
			c.errorf("C305", s.At)
		}
		if cx.inSafe {
			c.errorf("C316", s.At)
		}
		if s.Until {
			c.expect(s.Value, c.expr(s.Value, sc, cx), tBool)
		} else {
			c.expect(s.Value, c.expr(s.Value, sc, cx), tTime)
		}
	case "pass":
	case "return":
		if cx.fn == nil {
			c.errorf("C309", s.At)
			return
		}
		if s.Value != nil {
			got := c.expr(s.Value, sc, cx)
			if cx.fn.Result != nil {
				c.expect(s.Value, got, c.resolveType(cx.fn.Result))
			}
		}
	case "break", "continue":
		if cx.loops == 0 {
			c.errorf("C310", s.At, "kw", s.Kind)
		}
	}
}

func (c *checker) assign(s *AssignStmt, sc *scope, cx ctx) {
	var want *Type
	switch t := s.Target.(type) {
	case *Ident:
		sym, narrowed := sc.lookup(t.Name)
		if sym == nil {
			if !c.unchecked {
				c.errorf("C301", t.At, "name", t.Name)
			}
			c.expr(s.Value, sc, cx)
			return
		}
		sym.used = true
		if sym.kind == symConst || sym.kind == symDevice || sym.kind == symFn || sym.kind == symBlock || sym.kind == symModule {
			c.errorf("C306", t.At, "name", t.Name)
		}
		if sym.kind == symVar && c.global.syms[t.Name] == sym {
			c.noteShare(sym, s.At, cx.core, true)
		}
		want = sym.t
		c.types[t] = sym.t // the generator needs the target's type (int compound assignments, spec 4)
		if narrowed {
			want = nonOpt(want)
		}
	default:
		if mem, ok := s.Target.(*Member); ok {
			c.shareRecv(mem.X, sc, s.At, cx.core)
			if mt := c.expr(mem.X, sc, cx); mt.K == TObject && strings.HasPrefix(mt.Name, "mod:") {
				// motors.x = ...: only a module's vars may change
				if mc := c.mods[strings.TrimPrefix(mt.Name, "mod:")]; mc != nil && !mc.failed {
					if sym := mc.global.syms[mem.Name]; sym != nil && sym.kind != symVar {
						c.errorf("C306", mem.At, "name", mem.Name)
					}
				}
			}
		}
		if ix, ok := s.Target.(*Index); ok && nonOpt(c.expr(ix.X, sc, cx)).K == TText {
			c.errorf("C306", ix.X.Start(), "name", "a letter of a text")
		}
		want = c.expr(s.Target, sc, cx)
	}
	got := c.expr(s.Value, sc, cx)
	if s.Op != "=" {
		got = c.arith(strings.TrimSuffix(s.Op, "="), want, got, s.Target, s.Value, s.At)
	}
	c.expect(s.Value, got, nonOpt(want))
	if want.Opt && got.K != TNone && s.Op == "=" {
		// x = 5 after x was optional: from here on it's a value (simple flow: until the end of this block)
		if id, ok := s.Target.(*Ident); ok {
			sc.narrowed[id.Name] = true
		}
	}
}

func (c *checker) ifStmt(s *IfStmt, sc *scope, cx ctx) {
	var allFalse []string
	cur := sc
	for i, cond := range s.Conds {
		c.expect(cond, c.expr(cond, cur, cx), tBool)
		whenTrue, whenFalse := facts(cond)
		c.block(s.Thens[i], newScope(cur).narrow(whenTrue), cx)
		allFalse = append(allFalse, whenFalse...)
		cur = cur.narrow(whenFalse)
	}
	if s.Else != nil {
		c.block(s.Else, newScope(cur), cx)
	}
	// "if t == none: return" - after it, t has a value for the rest of the block.
	if len(s.Conds) == 1 && s.Else == nil && exits(s.Thens[0]) {
		_, whenFalse := facts(s.Conds[0])
		for _, n := range whenFalse {
			sc.narrowed[n] = true
		}
	}
}

func (c *checker) match(s *MatchStmt, sc *scope, cx ctx) {
	vt := c.expr(s.Value, sc, cx)
	seen := map[string]bool{}
	hasElse := false
	for _, mc := range s.Cases {
		if mc.IsElse {
			hasElse = true
		}
		for _, pt := range mc.Patterns {
			if r, ok := pt.(*RangeExpr); ok {
				c.expect(r.Lo, c.expr(r.Lo, sc, cx), vt)
				c.expect(r.Hi, c.expr(r.Hi, sc, cx), vt)
				continue
			}
			c.expect(pt, c.expr(pt, sc, cx), nonOpt(vt))
			if mem, ok := pt.(*Member); ok {
				seen[mem.Name] = true
			}
			if id, ok := pt.(*Ident); ok && c.declared[id] == nil {
				if t := c.types[id]; t != nil && t.K == TEnum && t.Name == vt.Name {
					seen[id.Name] = true // a bare member pattern (spec 4.6)
				}
			}
		}
		c.block(mc.Body, newScope(sc), cx)
	}
	if vt.K == TEnum && !hasElse {
		if sym := c.typeSym(vt.Name); sym != nil && sym.enum != nil {
			var missing []string
			for _, v := range sym.enum.Values {
				if !seen[v] {
					missing = append(missing, vt.Name+"."+v)
				}
			}
			if len(missing) > 0 {
				c.errorf("C315", s.At, "missing", strings.Join(missing, ", "))
			}
		}
	}
}

// facts finds which optional names a condition proves not to be none when it is true / false.
func facts(n Node) (whenTrue, whenFalse []string) {
	switch e := n.(type) {
	case *Binary:
		switch e.Op {
		case "!=", "==":
			name := ""
			if id, ok := e.X.(*Ident); ok && isNone(e.Y) {
				name = id.Name
			} else if id, ok := e.Y.(*Ident); ok && isNone(e.X) {
				name = id.Name
			}
			if name != "" {
				if e.Op == "!=" {
					return []string{name}, nil
				}
				return nil, []string{name}
			}
		case "and":
			t1, _ := facts(e.X)
			t2, _ := facts(e.Y)
			return append(t1, t2...), nil
		case "or":
			_, f1 := facts(e.X)
			_, f2 := facts(e.Y)
			return nil, append(f1, f2...)
		}
	case *Unary:
		if e.Op == "not" {
			t, f := facts(e.X)
			return f, t
		}
	}
	return nil, nil
}

func isNone(n Node) bool {
	l, ok := n.(*Literal)
	return ok && l.Value == "none"
}

// ---- expressions

// enumsWith lists the enum types that have a member called name, built once from the program's enums.
func (c *checker) enumsWith(name string) []string {
	if c.enumMembers == nil {
		c.enumMembers = map[string][]string{}
		for _, sym := range c.global.syms {
			if sym.kind == symEnum && sym.enum != nil {
				for _, v := range sym.enum.Values {
					c.enumMembers[v] = append(c.enumMembers[v], sym.t.Name)
				}
			}
		}
	}
	return c.enumMembers[name]
}

// tryEnum resolves a bare enum member when the context wants an enum that has a member of that name
// (spec 4.6): the pending C301 is cancelled and the name takes the enum type.
func (c *checker) tryEnum(n Node, want *Type) bool {
	id, ok := n.(*Ident)
	if !ok || want == nil || want.K != TEnum {
		return false
	}
	p, ok := c.pendingEnum[id]
	if !ok {
		// a bare name that is also a built-in enum's member ('high'): the program's own enum wins where it is wanted
		if t := c.types[id]; isBuiltinEnum(t) && !isBuiltinEnum(want) {
			for _, en := range c.enumsWith(id.Name) {
				if en == want.Name {
					c.types[id] = want
					return true
				}
			}
		}
		return false
	}
	for _, en := range p.enums {
		if en == want.Name {
			c.types[id] = want
			c.errs[p.err].Code = ""
			delete(c.pendingEnum, id)
			return true
		}
	}
	return false
}

// warnText reports a literal that cannot fit in the text type it's being assigned to.
func (c *checker) warnText(n Node, cap int) {
	switch e := n.(type) {
	case *StringLit:
		plain := 0
		for _, part := range e.Parts {
			plain += len(part.Text)
		}
		if plain > cap {
			c.errorf("W608", e.At, "n", strconv.Itoa(cap))
		}
	case *Binary:
		if e.Op == "+" {
			c.warnText(e.X, cap)
			c.warnText(e.Y, cap)
		}
	case *CondExpr:
		c.warnText(e.Then, cap)
		c.warnText(e.Else, cap)
	case *ListLit:
		for _, it := range e.Items {
			c.warnText(it, cap)
		}
	}
}

// expect reports a value that doesn't fit where it is used: a missing unit (U501), the wrong dimension (U503), none (C304)
// or another type (C303).
func (c *checker) expect(n Node, got, want *Type) {
	if n == nil {
		return
	}
	if want.K == TText {
		c.warnText(n, textCapOf(want))
	}
	if want.K == TList {
		if el := nonOpt(elemOf(want)); el != nil && el.K == TText {
			c.warnText(n, textCapOf(el))
		}
	}
	if c.tryEnum(n, want) {
		return
	}
	if assignable(want, got) {
		return
	}
	if got.Opt && !want.Opt && assignable(want, nonOpt(got)) {
		c.errorf("C304", n.Start(), "name", nodeName(n))
		return
	}
	if want.K == TQty {
		if num, ok := n.(*NumberLit); ok && num.Unit == "" {
			c.errorf("U501", num.At, "value", num.Text, "suggestion", unitExample[want.Dim])
			return
		}
		if got.K == TQty {
			c.errorf("U503", n.Start(), "want", want.Dim, "example", "1 "+unitExample[want.Dim], "got", got.Dim)
			return
		}
	}
	c.errorf("C303", n.Start(), "want", describe(want), "got", describe(got))
}

func nodeName(n Node) string {
	switch e := n.(type) {
	case *Ident:
		return e.Name
	case *Member:
		return nodeName(e.X) + "." + e.Name
	case *Call:
		return nodeName(e.Fn) + "()"
	}
	return "this value"
}

// isSecret reports whether t is a value that must never leave the board (spec 11.1).
func isSecret(t *Type) bool {
	return t != nil && t.K == TObject && (t.Name == "secret" || t.Name == "credentials")
}

// secretConsumers are the only places a secret("…") may be written (spec 11.1): as the argument that needs it, so its
// value can't be copied into a variable, joined into text, printed or sent. ok reports whether an argument (by name or
// position) takes a secret.
var secretConsumers = map[string]func(name string, i int) bool{
	"wifi.connect": func(name string, i int) bool { return name == "network" || (name == "" && i == 0) },
	"mqtt.connect": func(name string, i int) bool {
		return name == "user" || name == "password" || (name == "" && (i == 1 || i == 2))
	},
}

// secretName is what secret("…") may name: it becomes a C macro (HOME_WIFI, or HOME_WIFI_SSID and _PASSWORD for
// Wi-Fi), which the build server only accepts in capitals, at most 39 characters.
var secretName = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]{0,29}$`)

func isSecretCall(n Node) bool {
	if c, ok := n.(*Call); ok {
		if id, ok := c.Fn.(*Ident); ok && id.Name == "secret" {
			return true
		}
	}
	return false
}

// secretCall checks one secret("name"): a fixed, valid name (C325), written where it is used (C323).
func (c *checker) secretCall(e *Call) {
	if len(e.Args) != 1 {
		return // C307 already
	}
	lit, ok := e.Args[0].Value.(*StringLit)
	name := ""
	if ok {
		for _, p := range lit.Parts {
			if p.Expr != nil {
				ok = false
			}
			name += p.Text
		}
	}
	if !ok || !secretName.MatchString(name) {
		c.errorf("C325", e.Args[0].Value.Start(), "name", sourceText(e.Args[0].Value))
		return
	}
	if !c.secretOK[e] {
		c.errorf("C323", e.Fn.Start())
	}
}

// secretSinks are the calls that would show or send their arguments; a secret can't go there (C323).
var secretSinks = map[string]bool{
	"notify": true, "show": true, "plot": true,
	"log.info": true, "log.warn": true, "log.error": true,
	"mqtt.publish": true, "http.post": true,
}

// secrets reports C323 when a secret reaches print/notify/show/plot/log, also inside a text
// interpolation: print "{secret("key")}" is the same leak.
func (c *checker) secrets(n Node) {
	if isSecret(c.types[n]) {
		pos := n.Start()
		if call, ok := n.(*Call); ok {
			pos = call.Fn.Start() // point at 'secret', not its '('
		}
		c.errorf("C323", pos)
		return
	}
	if s, ok := n.(*StringLit); ok {
		for _, part := range s.Parts {
			if part.Expr != nil {
				c.secrets(part.Expr)
			}
		}
	}
}

// value marks a use that needs an actual value: an optional here is C304.
func (c *checker) value(n Node, t *Type) *Type {
	if t.Opt {
		c.errorf("C304", n.Start(), "name", nodeName(n))
		return nonOpt(t)
	}
	return t
}

func (c *checker) expr(n Node, sc *scope, cx ctx) *Type {
	t := c.exprType(n, sc, cx)
	if n != nil {
		c.types[n] = t
	}
	return t
}

func (c *checker) exprType(n Node, sc *scope, cx ctx) *Type {
	switch e := n.(type) {
	case nil:
		return tAny
	case *Ident:
		return c.ident(e, sc, cx)
	case *NumberLit:
		if e.Unit != "" {
			d := dimOfUnit(e.Unit)
			if d == "" {
				c.errorf("U504", e.At, "u", e.Unit)
				return tAny
			}
			return qty(d)
		}
		if e.IsFloat {
			return tFloat
		}
		if e.Int > 2147483647 {
			c.errorf("W607", e.At)
		}
		return tInt
	case *StringLit:
		plain := 0
		for _, part := range e.Parts {
			if part.Expr != nil {
				c.expr(part.Expr, sc, cx) // printing an optional shows "none": allowed
			}
			plain += len(part.Text) // bytes, as the runtime counts them
		}
		if plain > 255 {
			c.errorf("W608", e.At, "n", "255")
		}
		return tText
	case *Literal:
		if e.Value == "none" {
			return tNone
		}
		return tBool
	case *ListLit:
		el := tAny
		for i, it := range e.Items {
			t := c.value(it, c.expr(it, sc, cx))
			if i == 0 {
				el = t
			}
		}
		return listOf(el)
	case *PinLit:
		if c.board != nil {
			role := ""
			if e.Analog {
				role = "analog"
			}
			c.pin(e.Ref, role, c.pinOwner)
		}
		return object("pin")
	case *Unary:
		t := c.value(e.X, c.expr(e.X, sc, cx))
		switch e.Op {
		case "not":
			c.expect(e.X, t, tBool)
			return tBool
		case "~":
			c.expect(e.X, t, tInt)
			return tInt
		}
		if !(isNumeric(t) || t.K == TQty || t.K == TAny) {
			c.errorf("C303", e.X.Start(), "want", "a number", "got", describe(t))
		}
		return t
	case *Binary:
		return c.binary(e, sc, cx)
	case *RangeExpr:
		lo := c.value(e.Lo, c.expr(e.Lo, sc, cx))
		hi := c.value(e.Hi, c.expr(e.Hi, sc, cx))
		c.sameKind(lo, hi, e.Lo, e.Hi, e.At)
		if lo.K == TAny {
			return rangeOf(hi)
		}
		return rangeOf(lo)
	case *Conversion:
		t := c.value(e.X, c.expr(e.X, sc, cx))
		d := dimOfUnit(e.Unit)
		switch {
		case isNumeric(t):
			return qty(d) // 'as unit' on a plain number makes a quantity of that unit (spec 5)
		case t.K == TQty && t.Dim != d:
			c.errorf("U502", e.At, "a", t.Dim, "b", e.Unit)
		}
		return t
	case *CondExpr:
		c.expect(e.Cond, c.expr(e.Cond, sc, cx), tBool)
		t := c.expr(e.Then, sc, cx)
		te := c.expr(e.Else, sc, cx)
		c.tryEnum(e.Then, te)
		c.tryEnum(e.Else, t)
		return t
	case *Call:
		return c.call(e, sc, cx)
	case *Member:
		return c.member(e, sc, cx, false)
	case *Index:
		t := c.value(e.X, c.expr(e.X, sc, cx))
		if t.K == TObject && t.Name == "json" {
			it := c.expr(e.Ix, sc, cx)
			if !(it.K == TInt || it.K == TText || it.K == TAny) {
				c.errorf("C303", e.Ix.Start(), "want", "a key or an index", "got", describe(it))
			}
			return object("json")
		}
		c.expect(e.Ix, c.expr(e.Ix, sc, cx), tInt)
		switch {
		case t.K == TList:
			return elemOf(t)
		case t.K == TDevice && t.Name == "group":
			return t.Elem
		case t.K == TText:
			return tText
		case t.K == TAny:
			return tAny
		}
		c.errorf("C303", e.X.Start(), "want", "a list", "got", describe(t))
		return tAny
	}
	return tAny
}

func (c *checker) ident(e *Ident, sc *scope, cx ctx) *Type {
	if e.Name == "" {
		return tAny // a parse error was already reported here
	}
	if sym, narrowed := sc.lookup(e.Name); sym != nil {
		sym.used = true
		if sym.kind == symVar && c.global.syms[e.Name] == sym {
			c.noteShare(sym, e.At, cx.core, false)
		}
		defer func() { c.declared[e] = sym.t }()
		if sc.crossesBarrier(e.Name) {
			c.errorf("C319", e.At, "name", e.Name)
		}
		if sym.kind == symVar || sym.kind == symConst {
			if sym.decl != nil && !sym.checked {
				c.globalVar(sym.decl) // a global used before its line (in a function, say)
			}
		}
		switch sym.kind {
		case symFn:
			return &Type{K: TObject, Name: "fn:" + sym.name}
		case symStruct:
			return &Type{K: TObject, Name: "struct:" + sym.name}
		case symEnum:
			return &Type{K: TObject, Name: "enum:" + sym.name}
		case symModule:
			return &Type{K: TObject, Name: "mod:" + sym.name}
		}
		if narrowed {
			return nonOpt(sym.t)
		}
		return sym.t
	}
	if g, ok := globals[e.Name]; ok {
		// A built-in from a Mujard library the program doesn't use gives way to a C++ library it does use that
		// names the same thing: 'use "Servo"' then Servo() is the Arduino library's, not mujard.servo's.
		_, fromLib := c.libGlobal(e.Name)
		builtinOff := g.lib != "" && !c.libs[g.lib] && !c.unchecked
		if builtinOff && !fromLib {
			c.errorf("C318", e.At, "name", e.Name, "lib", g.lib)
			return tAny
		}
		if !builtinOff {
			switch {
			case g.value != nil:
				return g.value
			case g.class != "":
				return object(g.class)
			}
			return &Type{K: TObject, Name: "builtin:" + e.Name}
		}
	}
	if g, ok := c.libGlobal(e.Name); ok {
		if g.emit != "" && g.emit != e.Name {
			c.libReal[e.Name] = g.emit
		}
		switch {
		case g.value != nil:
			return g.value
		case g.class != "":
			return object(g.class)
		}
		return &Type{K: TObject, Name: "builtin:" + e.Name}
	}
	if c.unchecked {
		return tAny // a name from a C++ library or a cpp: block
	}
	if enums := c.enumsWith(e.Name); len(enums) > 0 {
		// A bare enum member (spec 4.6): resolves if the context fixes the enum type; the
		// C301 stays pending until then and suggests the full form.
		c.errs = append(c.errs, diag("C301", e.At, "name", e.Name))
		c.errs[len(c.errs)-1].Message += " Did you mean '" + enums[0] + "." + e.Name + "'?"
		c.pendingEnum[e] = pendingEnum{err: len(c.errs) - 1, enums: enums}
		return tAny
	}
	c.errorf("C301", e.At, "name", e.Name)
	return tAny
}

func (c *checker) binary(e *Binary, sc *scope, cx ctx) *Type {
	switch e.Op {
	case "and", "or":
		l := c.expr(e.X, sc, cx)
		c.expect(e.X, l, tBool)
		inner := sc
		if e.Op == "and" {
			t, _ := facts(e.X)
			inner = sc.narrow(t)
		} else {
			_, f := facts(e.X)
			inner = sc.narrow(f)
		}
		c.expect(e.Y, c.expr(e.Y, inner, cx), tBool)
		return tBool
	case "==", "!=":
		l, r := c.expr(e.X, sc, cx), c.expr(e.Y, sc, cx)
		if isNone(e.X) || isNone(e.Y) {
			return tBool
		}
		c.compare(nonOpt(c.value(e.X, l)), nonOpt(c.value(e.Y, r)), e)
		return tBool
	case "<", "<=", ">", ">=":
		l := c.value(e.X, c.expr(e.X, sc, cx))
		r := c.value(e.Y, c.expr(e.Y, sc, cx))
		c.compare(l, r, e)
		return tBool
	case "in":
		l := c.value(e.X, c.expr(e.X, sc, cx))
		r := c.value(e.Y, c.expr(e.Y, sc, cx))
		switch {
		case r.K == TRange || r.K == TList:
			el := nonOpt(elemOf(r))
			if el.K == TQty || l.K == TQty {
				c.sameKind(el, nonOpt(l), e.Y, e.X, e.At)
			} else if l.K != TAny && el.K != TAny && l.K != el.K {
				c.errorf("C303", e.X.Start(), "want", describe(el), "got", describe(l))
			}
		case r.K == TAny:
		default:
			c.errorf("C303", e.Y.Start(), "want", "a range or a list", "got", describe(r))
		}
		return tBool
	}
	l := c.value(e.X, c.expr(e.X, sc, cx))
	r := c.value(e.Y, c.expr(e.Y, sc, cx))
	return c.arith(e.Op, l, r, e.X, e.Y, e.At)
}

// compare checks that two compared values measure the same thing.
func (c *checker) compare(l, r *Type, e *Binary) {
	// 'm == idle': the other side fixes the enum type (spec 4.6)
	if a, b := c.tryEnum(e.X, r), c.tryEnum(e.Y, l); a || b {
		return
	}
	if l.K == TAny || r.K == TAny {
		return
	}
	if l.K == TQty || r.K == TQty {
		c.sameKind(l, r, e.X, e.Y, e.At)
		return
	}
	if isNumeric(l) && isNumeric(r) {
		return
	}
	if l.K != r.K || (l.K == TEnum && l.Name != r.Name) {
		c.errorf("C303", e.Y.Start(), "want", describe(l), "got", describe(r))
	}
}

// sameKind checks two values that must have the same dimension (a range's ends, a comparison, + and -).
func (c *checker) sameKind(l, r *Type, ln, rn Node, at Pos) bool {
	if l.K == TAny || r.K == TAny {
		return true
	}
	if l.K == TQty && r.K == TQty {
		if l.Dim != r.Dim {
			c.errorf("U502", at, "a", l.Dim, "b", r.Dim)
			return false
		}
		return true
	}
	if l.K == TQty && isNumeric(r) {
		return c.unitless(rn, l)
	}
	if r.K == TQty && isNumeric(l) {
		return c.unitless(ln, r)
	}
	return true
}

// unitless reports a plain number used with a quantity, except 0 which fits any unit.
func (c *checker) unitless(n Node, q *Type) bool {
	if num, ok := n.(*NumberLit); ok {
		if num.Unit == "" && num.Text == "0" {
			return true
		}
		c.errorf("U501", num.At, "value", num.Text, "suggestion", unitExample[q.Dim])
		return false
	}
	c.errorf("U502", n.Start(), "a", q.Dim, "b", "plain number")
	return false
}

// arith types arithmetic (spec 5): units must match for + and -, and combine only in ways that mean something for * and /.
func (c *checker) arith(op string, l, r *Type, ln, rn Node, at Pos) *Type {
	if l.K == TAny || r.K == TAny {
		if l.K == TAny {
			return r
		}
		return l
	}
	switch op {
	case "+":
		if l.K == TText || r.K == TText {
			if l.K != TText || r.K != TText {
				c.errorf("C303", rn.Start(), "want", "text", "got", describe(r))
			}
			return tText
		}
		fallthrough
	case "-":
		if l.K == TQty || r.K == TQty {
			if c.sameKind(l, r, ln, rn, at) {
				if l.K == TQty {
					return l
				}
				return r
			}
			return tAny
		}
		return c.numeric(l, r, ln, rn)
	case "*":
		switch {
		case l.K == TQty && r.K == TQty:
			if l.Dim == "ratio" {
				return r
			}
			if r.Dim == "ratio" {
				return l
			}
			if d, ok := mulDims(l.Dim, r.Dim); ok {
				if d == "" {
					return tFloat
				}
				return qty(d)
			}
			c.errorf("U502", at, "a", l.Dim, "b", r.Dim)
			return tAny
		case l.K == TQty && isNumeric(r):
			if l.Dim == "ratio" {
				return r
			}
			return l
		case r.K == TQty && isNumeric(l):
			if r.Dim == "ratio" {
				return l
			}
			return r
		}
		return c.numeric(l, r, ln, rn)
	case "/":
		switch {
		case l.K == TQty && r.K == TQty:
			if d, ok := divDims(l.Dim, r.Dim); ok {
				if d == "" {
					return tFloat
				}
				return qty(d)
			}
			c.errorf("U502", at, "a", l.Dim, "b", r.Dim)
			return tAny
		case l.K == TQty && isNumeric(r):
			return l
		case r.K == TQty && isNumeric(l):
			if r.Dim == "time" {
				return tFreq
			}
			if r.Dim == "frequency" {
				return tTime
			}
			c.errorf("U502", at, "a", "plain number", "b", r.Dim)
			return tAny
		}
		return c.numeric(l, r, ln, rn)
	case "mod", "&", "|", "^", "<<", ">>":
		c.expect(ln, l, tInt)
		c.expect(rn, r, tInt)
		return tInt
	}
	return tAny
}

func (c *checker) numeric(l, r *Type, ln, rn Node) *Type {
	if !isNumeric(l) {
		c.errorf("C303", ln.Start(), "want", "a number", "got", describe(l))
		return tAny
	}
	if !isNumeric(r) {
		c.errorf("C303", rn.Start(), "want", "a number", "got", describe(r))
		return tAny
	}
	if l.K == TFloat || r.K == TFloat {
		return tFloat
	}
	return tInt
}

// classOf finds the members of a type, or nil when they aren't known (undescribed library objects).
func (c *checker) classOf(t *Type) class {
	switch t.K {
	case TDevice:
		return classes[t.Name]
	case TObject:
		if cl := classes[t.Name]; cl != nil {
			return cl // built-ins win over a library class of the same name
		}
		return c.libClasses[t.Name]
	case TText:
		return textMembers
	case TInt:
		return intMembers
	case TList:
		return listMembers
	}
	return nil
}

// member types x.name. called says whether it is being called (x.name(...)).
func (c *checker) member(e *Member, sc *scope, cx ctx, called bool) *Type {
	xt := c.expr(e.X, sc, cx)
	if strings.HasPrefix(xt.Name, "mod:") && xt.K == TObject {
		return c.modMember(e, strings.TrimPrefix(xt.Name, "mod:"))
	}
	if strings.HasPrefix(xt.Name, "enum:") && xt.K == TObject {
		en := c.typeSym(strings.TrimPrefix(xt.Name, "enum:"))
		if en == nil || en.enum == nil {
			return tAny
		}
		for _, v := range en.enum.Values {
			if v == e.Name {
				return en.t
			}
		}
		c.errorf("C308", e.At, "type", en.name, "member", e.Name)
		return tAny
	}
	xt = c.value(e.X, xt)
	if xt.K == TStruct {
		sym := c.typeSym(xt.Name)
		if sym == nil || sym.strukt == nil {
			return tAny
		}
		owner := c
		if i := strings.Index(xt.Name, "."); i > 0 {
			if mc := c.mods[xt.Name[:i]]; mc != nil {
				owner = mc // a module struct's field types resolve in the module's scope
			}
		}
		for _, f := range sym.strukt.Fields {
			if f.Name == e.Name {
				return owner.resolveType(f.Type)
			}
		}
		c.errorf("C308", e.At, "type", xt.Name, "member", e.Name)
		return tAny
	}
	if xt.K == TQty && e.Name == "number" {
		return tFloat
	}
	if xt.K == TObject && xt.Name == "regs" {
		// 'use mujard.regs': register names are per-chip, generated into regs.json (spec 10.4)
		if table, ok := regsOf(c.board); ok {
			if _, ok := table[e.Name]; ok {
				return regTypeName(e.Name)
			}
			c.errorf("C308", e.At, "type", "regs", "member", e.Name)
			return tAny
		}
		if !c.regsNoChip {
			c.regsNoChip = true
			chip, display := "", ""
			if c.board != nil {
				chip, display = c.board.Family, c.board.Display
			}
			c.errorf("B413", e.At, "board", display, "chip", chip)
		}
		return tAny
	}
	if xt.K == TObject && regName(xt) != "" {
		if mm, ok := classes["regs_reg"][e.Name]; ok {
			if mm.method != nil {
				if !called {
					return &Type{K: TObject, Name: "method"}
				}
				return tAny // typed by call()
			}
			return mm.field
		}
		if table, ok := regsOf(c.board); ok {
			if r, ok := table[regName(xt)]; ok {
				if _, ok := r.Fields[e.Name]; ok {
					return tInt
				}
			}
		}
		c.errorf("C308", e.At, "type", regName(xt), "member", e.Name)
		return tAny
	}
	cl := c.classOf(xt)
	if cl == nil {
		return tAny
	}
	mm, ok := cl[e.Name]
	if !ok && c.libClasses[xt.Name] != nil {
		// A described C++ class: snake_case spellings reach the real member ('set_pixel_color' -> 'setPixelColor').
		if rn := realName(e.Name, cl); rn != "" {
			if rn != e.Name {
				c.libReal[xt.Name+"."+e.Name] = rn
			}
			mm, ok = cl[rn], true
		}
	}
	if !ok {
		if xt.K == TDevice && (xt.Name == "output" || xt.Name == "pwm") && cl["write"] != nil && (e.Name == "on" || e.Name == "off") {
			return tVoid
		}
		c.errorf("C308", e.At, "type", bare(xt), "member", e.Name)
		return tAny
	}
	switch xt.Name {
	case "wifi", "clock", "http", "mqtt":
		c.libs[xt.Name] = true // the generator adds mujard_wifi.h / mujard_http.h / mujard_mqtt.h
	case "store":
		c.libs["store"] = true // store.pending/failed/committed/last_commit need store_db (mujard_store.h) too
	}
	if xt.K == TDevice && (e.Name == "tone" || e.Name == "melody") {
		if !c.usesTone {
			c.toneAt = e.At
		}
		c.usesTone = true // the core's tone() takes a timer on AVR boards (libraryClaims, timerOwners)
	}
	if xt.Name == "wifi" || xt.Name == "mqtt" || xt.Name == "http" || (mm.method != nil && mm.method.network) {
		c.usesWifi = true // the radio will be on: analog pins Wi-Fi blocks get W615 (wifiADC)
	}
	if c.board != nil && !c.board.Wifi && (xt.Name == "wifi" || xt.Name == "mqtt" || xt.Name == "http" || (mm.method != nil && mm.method.network)) {
		// clock.time/date/source don't need Wi-Fi when an RTC is the source (use mujard.rtc); clock.sync does.
		c.errorf("B409", e.At, "feature", nodeName(e), "what", "Wi-Fi", "board", c.board.Display)
	}
	switch {
	case mm.field != nil:
		return mm.field
	case mm.method != nil:
		if !called {
			return &Type{K: TObject, Name: "method"}
		}
		return tAny // typed by call()
	}
	return tAny // an event used as a value
}

func (c *checker) call(e *Call, sc *scope, cx ctx) *Type {
	var sig *fnSig
	var callee string
	var recv *Type
	var memberName string
	switch fn := e.Fn.(type) {
	case *Member:
		xt := c.expr(fn.X, sc, cx)
		recv = nonOpt(xt)
		memberName = fn.Name
		if xt.K == TObject && strings.HasPrefix(xt.Name, "mod:") {
			return c.modCall(e, strings.TrimPrefix(xt.Name, "mod:"), fn.Name, sc, cx)
		}
		if xt.K == TObject && strings.HasPrefix(xt.Name, "enum:") {
			c.member(fn, sc, cx, true)
			break
		}
		c.member(fn, sc, cx, true)
		if cl := c.classOf(nonOpt(xt)); cl != nil {
			name := fn.Name
			if c.libClasses[xt.Name] != nil {
				name = realName(fn.Name, cl)
			}
			if mm := cl[name]; mm != nil && mm.method != nil {
				sig = mm.method
			}
		}
		if cx.core > 0 && c.board != nil && (c.board.Family == "RP2040" || c.board.Family == "RP2350") &&
			nonOpt(xt).K == TObject && (nonOpt(xt).Name == "uart" || nonOpt(xt).Name == "serial" ||
			nonOpt(xt).Name == "i2c" || nonOpt(xt).Name == "spi") {
			c.errorf("W621", e.At, "what", nonOpt(xt).Name+"."+fn.Name, "board", c.board.Display)
		}
		// A method call on a shared object can mutate it: for cross-core tracking it counts as a write
		// (a Timer's stop/fire, a Uart's write, a Can bus send). Fields read through member() stay reads.
		if nonOpt(xt).K == TObject {
			c.shareRecv(fn.X, sc, e.At, cx.core)
		}
		// RP2040/RP2350: two pins on one PWM slice share its frequency. tone() with a different pitch on each
		// can't both hold - warn only when the frequencies provably differ, never when they match.
		if c.rpFamily() && nonOpt(xt).K == TDevice && fn.Name == "tone" && len(e.Args) >= 1 {
			if id, ok := fn.X.(*Ident); ok {
				if pin, okP := c.devicePin(id.Name); okP {
					slice := (pin >> 1) & 7
					if hz, okH := c.freqArg(e.Args[0].Value); okH {
						c.toneSlice(slice, pin, hz, e.At)
					}
				}
			}
		}
		// A queue specializes its members on its declared element type: push takes the element, pop gives element?.
		if nonOpt(xt).Name == "queue" && nonOpt(xt).Elem != nil {
			el := nonOpt(xt).Elem
			switch fn.Name {
			case "push":
				sig = &fnSig{params: []param{p("value", el)}, result: tBool}
			case "pop":
				sig = &fnSig{params: nil, result: opt(el)}
			}
		}
		if reg := regName(nonOpt(xt)); reg != "" && nonOpt(xt).K == TObject {
			if mm := classes["regs_reg"][fn.Name]; mm != nil && mm.method != nil {
				sig = mm.method
			}
			c.regsCall(e, reg, fn.Name)
		}
		callee = nodeName(fn)
	case *Ident:
		t := c.ident(fn, sc, cx)
		callee = fn.Name
		switch {
		case strings.HasPrefix(t.Name, "fn:"):
			return c.userCall(e, c.global.syms[fn.Name], c, sc, cx)
		case strings.HasPrefix(t.Name, "struct:"):
			return c.structCall(e, c.global.syms[fn.Name], c, sc, cx)
		}
		if g, ok := globals[fn.Name]; ok && g.fn != nil {
			if _, local := sc.lookup(fn.Name); !local {
				sig = g.fn
			}
		}
		if g, ok := c.libGlobal(fn.Name); ok && g.fn != nil {
			if _, local := sc.lookup(fn.Name); !local {
				sig = g.fn
			}
		}
		if fn.Name == "Timer" {
			if c.timerCount == 0 {
				c.timerAt = e.At // the first Timer() is the owner the second collides with
			}
			c.usesTimer = true
			c.timerCount++
		}
		if fn.Name == "Queue" {
			return c.queueCtor(e, sc, cx)
		}
	default:
		c.expr(e.Fn, sc, cx)
	}
	if sig == nil {
		for _, a := range e.Args {
			c.expr(a.Value, sc, cx)
		}
		return tAny
	}
	if sig.testOnly && !cx.inTest {
		c.errorf("C317", e.Start(), "name", callee)
	}
	if sig.blocking {
		if !cx.inTask {
			c.errorf("C311", e.Start(), "call", callee)
		}
		if cx.fnSym != nil {
			cx.fnSym.waits = true // a fn holding a suspending call is itself a fn that waits
		}
	}
	if cx.inSafe && (sig.blocking || sig.network) {
		c.errorf("C316", e.Start())
	}
	if ok := secretConsumers[callee]; ok != nil {
		// the one place a secret may go: straight into the call that uses it (spec 11.1)
		for i, a := range e.Args {
			if s, isCall := a.Value.(*Call); isCall && ok(a.Name, i) {
				if c.secretOK == nil {
					c.secretOK = map[*Call]bool{}
				}
				c.secretOK[s] = true
			}
		}
	}
	args := c.args(e, sig.params, sig.variadic, callee, sc, cx)
	if g := globals[callee]; g != nil && g.fn == sig {
		c.siteBytes += siteMemory[callee] // each call site of smooth, peak, ... keeps its own memory (spec 6.9)
	}
	if callee == "secret" {
		c.secretCall(e)
	}
	if callee == "clock.sync" {
		for _, a := range e.Args {
			if a.Name != "timezone" && a.Name != "" {
				continue
			}
			if tz, ok := a.Value.(*StringLit); !ok || len(tz.Parts) != 1 || tz.Parts[0].Expr != nil {
				c.errorf("C326", a.Value.Start(), "tz", sourceText(a.Value))
			} else if _, known := posixZone(tz.Parts[0].Text); !known {
				c.errorf("C326", a.Value.Start(), "tz", tz.Parts[0].Text)
			}
		}
	}
	if callee == "wifi.connect" && len(e.Args) > 0 && !isSecretCall(e.Args[0].Value) {
		c.errorf("C303", e.Args[0].Value.Start(), "want", "credentials, written secret(\"name\")", "got", describe(args[0]))
	}
	if strings.HasPrefix(callee, "store.") {
		c.libs["store"] = true // the generator adds the store (mujard_store.h)
		if callee == "store.save" && len(args) > 1 && (args[1].Opt || !storable(args[1])) {
			c.errorf("C312", e.Args[1].Value.Start())
		}
	}
	if secretSinks[callee] {
		for _, a := range e.Args {
			c.secrets(a.Value)
		}
	}
	if sig.resultOf != nil {
		return sig.resultOf(args)
	}
	if recv != nil && recv.K == TText {
		cap := textCapOf(recv)
		switch memberName {
		case "upper", "lower":
			return textCap(cap)
		case "split":
			return listCap(textCap(cap), 8)
		}
	}
	if sig.result == nil {
		return tAny
	}
	return sig.result
}

// args checks a call's arguments against parameters (positional first, then named) and returns their types.
func (c *checker) args(e *Call, params []param, variadic bool, callee string, sc *scope, cx ctx) []*Type {
	var types []*Type
	if params == nil {
		for _, a := range e.Args {
			types = append(types, c.expr(a.Value, sc, cx))
		}
		return types
	}
	filled := make([]bool, len(params))
	pos := 0
	for _, a := range e.Args {
		t := c.expr(a.Value, sc, cx)
		types = append(types, t)
		idx := -1
		if a.Name != "" {
			for i, prm := range params {
				if prm.name == a.Name {
					idx = i
				}
			}
			if idx < 0 {
				c.errorf("C308", a.Value.Start(), "type", callee, "member", a.Name)
				continue
			}
		} else {
			for pos < len(params) && filled[pos] {
				pos++
			}
			if pos >= len(params) {
				if variadic && len(params) > 0 {
					idx = len(params) - 1
				} else {
					c.errorf("C307", e.Start(), "fn", callee, "n", strconv.Itoa(len(params)), "m", strconv.Itoa(len(e.Args)))
					continue
				}
			} else {
				idx = pos
			}
		}
		filled[idx] = true
		if want := params[idx].t; want != nil {
			c.expect(a.Value, t, want)
		}
	}
	for i, prm := range params {
		if !filled[i] && !prm.optional {
			c.errorf("C307", e.Start(), "fn", callee, "n", strconv.Itoa(requiredCount(params)), "m", strconv.Itoa(len(e.Args)))
			break
		}
	}
	return types
}

func requiredCount(ps []param) int {
	n := 0
	for _, prm := range ps {
		if !prm.optional {
			n++
		}
	}
	return n
}

// modCall types name.fn(...) and name.Struct(...): the arguments are the program's, the signature the module's.
func (c *checker) modCall(e *Call, mod, name string, sc *scope, cx ctx) *Type {
	mc := c.mods[mod]
	if mc == nil || mc.failed {
		for _, a := range e.Args {
			c.expr(a.Value, sc, cx)
		}
		return tAny
	}
	sym := mc.global.syms[name]
	if sym != nil {
		switch sym.kind {
		case symFn:
			return c.userCall(e, sym, mc, sc, cx)
		case symStruct:
			return c.structCall(e, sym, mc, sc, cx)
		}
	}
	c.errorf("C308", e.Fn.Start(), "type", "module '"+mod+"'", "member", name)
	for _, a := range e.Args {
		c.expr(a.Value, sc, cx)
	}
	return tAny
}

// userCall checks a call of a user's fn; owner is the program or module the fn lives in (its types resolve there).
func (c *checker) userCall(e *Call, sym *symbol, owner *checker, sc *scope, cx ctx) *Type {
	d := sym.fn
	var ps []param
	for _, prm := range d.Params {
		ps = append(ps, param{name: prm.Name, t: owner.resolveType(prm.Type), optional: prm.Default != nil})
	}
	if ps == nil {
		ps = []param{}
	}
	callee := d.Name
	if owner.inModule {
		callee = owner.modName + "." + d.Name
	}
	c.args(e, ps, false, callee, sc, cx)
	if cx.fnSym != nil {
		if c.waitDeps == nil {
			c.waitDeps = map[*symbol][]*symbol{}
		}
		c.waitDeps[cx.fnSym] = append(c.waitDeps[cx.fnSym], sym) // a fn that calls a waiting fn waits too (order-free)
	}
	if !cx.inTask {
		// reported by propagateWaits once every fn's waits flag is final
		c.waitCalls = append(c.waitCalls, waitCall{at: e.Start(), callee: callee, sym: sym})
	}
	if d.Result == nil {
		return tVoid
	}
	return owner.resolveType(d.Result)
}

// structCall checks a struct literal call; owner is where the struct's field types resolve.
func (c *checker) structCall(e *Call, sym *symbol, owner *checker, sc *scope, cx ctx) *Type {
	var ps []param
	for _, f := range sym.strukt.Fields {
		ps = append(ps, param{name: f.Name, t: owner.resolveType(f.Type), optional: f.Default != nil})
	}
	c.args(e, ps, false, sym.t.Name, sc, cx)
	return sym.t
}

// String helps debugging: "C303 at 4:5".
func (d Diagnostic) Short() string { return fmt.Sprintf("%s at %s", d.Code, d.Pos) }
