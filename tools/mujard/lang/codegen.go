package lang

import (
	"embed"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math"
	"path"
	"sort"
	"strconv"
	"strings"
)

// The runtime the generated C++ links against; copied next to every generated sketch.
//
//go:embed runtime/mujard_rt.h runtime/mujard_rt.cpp runtime/mujard_servo.h runtime/mujard_ultrasonic.h runtime/mujard_i2c.h runtime/mujard_shift.h runtime/mujard_uart.h runtime/mujard_can.h runtime/mujard_ble.h runtime/mujard_ethernet.h runtime/mujard_usb.h runtime/mujard_timer.h runtime/mujard_queue.h runtime/mujard_store.h runtime/mujard_wifi.h runtime/mujard_clock.h runtime/mujard_http.h runtime/mujard_mqtt.h runtime/mujard_spi.h runtime/mujard_rtc.h runtime/mujard_test.h runtime/mujard_test.cpp
var runtimeFS embed.FS

// Sketch is a generated Arduino sketch: file name -> content, plus the board it is for.
type Sketch struct {
	Files     map[string]string // sketch.ino, mujard_program.cpp, mujard_rt.h, mujard_rt.cpp, and the mujard_* library headers in use
	FQBN      string
	Board     *Board
	Libraries []LibDep   // the libraries 'use' lines need installed (also written as mujard_libs.json)
	Tests     []TestInfo // a test build's tests, in source order (GenerateTests)
	Secrets   []string   // the secret macros the program needs (HOME_WIFI_SSID ...), sorted; their values are never in Files
}

// fqbns are the arduino-cli board ids of the Mujard boards.
var fqbns = map[string]string{
	"uno": "arduino:avr:uno", "nano": "arduino:avr:nano", "mega": "arduino:avr:mega", "leonardo": "arduino:avr:leonardo",
	"micro": "arduino:avr:micro", "esp32": "esp32:esp32:esp32", "esp8266": "esp8266:esp8266:nodemcuv2",
	"pico": "rp2040:rp2040:rpipico", "picow": "rp2040:rp2040:rpipicow", "esp32s3": "esp32:esp32:esp32s3",
	"rp2350usba": "rp2040:rp2040:generic_rp2350:flash=2097152_0",
	"esp32s2":    "esp32:esp32:esp32s2", "esp32c3": "esp32:esp32:esp32c3", "esp32c6": "esp32:esp32:esp32c6",
	"unor4": "arduino:renesas_uno:minima", "nano33ble": "arduino:mbed_nano:nano33ble",
	"zero": "arduino:samd:arduino_zero_native", "bluepill": "STMicroelectronics:stm32:GenF1:pnum=BLUEPILL_F103C8",
}

// Compile checks a program and, if it has no errors, generates its sketch. srcName is the .muj file's name, used in
// the #line markers so C++ compiler messages point at the Mujard source.
func Compile(src, srcName string) (*Sketch, []Diagnostic) {
	return CompileIn(src, srcName, "")
}

// CompileIn is Compile with a directory for resolving 'import name' to name.muj next to the program.
func CompileIn(src, srcName, dir string) (*Sketch, []Diagnostic) {
	f, diags := ParseSource(src)
	if FirstError(diags) != nil {
		return nil, sortDiags(diags)
	}
	info, cdiags := analyzeWith(f, moduleLoader(dir), dirEnv(dir))
	diags = append(diags, cdiags...)
	if FirstError(diags) != nil || info.Board == nil {
		return nil, sortDiags(diags)
	}
	sk, gdiags := Generate(f, info, srcName)
	diags = append(diags, gdiags...)
	if FirstError(diags) != nil {
		return nil, sortDiags(diags)
	}
	return sk, sortDiags(diags)
}

func sortDiags(d []Diagnostic) []Diagnostic {
	sort.SliceStable(d, func(i, j int) bool {
		a, b := d[i], d[j]
		if a.File != b.File {
			return a.File < b.File
		}
		return a.Pos.Line < b.Pos.Line || (a.Pos.Line == b.Pos.Line && a.Pos.Col < b.Pos.Col)
	})
	return d
}

// ---- output with source mapping

type outLine struct {
	text string
	src  int    // the .muj line this C++ line comes from; 0 = generated scaffolding
	file string // the file src is a line of ("" = the program's own file; a module's name while it is generated)
}

// curSrcFile is the .muj file that emitted lines map to: set while a module's code is generated.
var curSrcFile string

type section []outLine

func (s *section) add(src int, format string, args ...any) {
	*s = append(*s, outLine{text: fmt.Sprintf(format, args...), src: src, file: curSrcFile})
}

// assemble writes the sections, adding #line markers whenever the source mapping changes.
func assemble(srcName, outName string, parts ...section) string {
	var b strings.Builder
	phys := 0
	nextSrc := -1 // the .muj line the compiler currently thinks the next line is
	mapped := false
	curFile := ""
	write := func(s string) {
		b.WriteString(s)
		b.WriteByte('\n')
		phys++
	}
	for _, part := range parts {
		for _, l := range part {
			if l.src > 0 {
				file := l.file
				if file == "" {
					file = srcName
				}
				if !mapped || l.src != nextSrc || file != curFile {
					write(fmt.Sprintf("#line %d %s", l.src, cString(file)))
					mapped = true
					curFile = file
				}
				write(l.text)
				nextSrc = l.src + 1
			} else {
				if mapped {
					write(fmt.Sprintf("#line %d %s", phys+2, cString(outName)))
					mapped = false
					curFile = ""
				}
				write(l.text)
			}
		}
	}
	return b.String()
}

// ---- the generator

type gen struct {
	f        *File
	info     *Info
	board    *Board
	src      string
	errs     []Diagnostic
	n        int // unique ids
	head     section
	globals  section
	fns      section
	setup    section
	loop     section
	deferred section // hoisted nested 'after' bodies, checked in loop()
	coreInc  bool    // a core-N block already wrote the multicore include (pico/multicore.h or freertos/task.h)

	enums      map[string]*EnumDecl
	structs    map[string]*StructDecl
	fnDecls    map[string]*FnDecl
	devices    map[string]*DeviceDecl
	varTypes   map[string]*TypeExpr // declared types of globals (list sizes)
	volatiles  map[string]bool      // globals changed in fast handlers
	atomics    map[string]bool      // ... that are wider than a byte: the rest of the program must update them with interrupts off (spec 6.5)
	blockNames map[string]string    // named blocks and tasks -> their C++ objects
	inTask     *taskCtx
	inFast     bool
	inAtomic   bool // emitting inside a noInterrupts()/interrupts() pair

	types       map[Node]*Type      // expression types of the code being generated (switches inside a module)
	decl        map[*Ident]*Type    // declared types of the code being generated
	mods        map[string]*checker // 'import name' -> the module's check result
	modGen      map[string]*modGen  // ... and its declarations, for generation
	curMod      string              // the module being generated ("" = the program)
	curModScope *scope              // its top-level names
	modQual     bool                // emit module names qualified (defaults at a call site)
	savedVars   []savedVar
	secrets     map[string]int // secret macros in use -> the first .muj line that needs it

	// the app built-ins (appgen.go)
	usesApp       bool
	usesNotify    bool // a notify runs: with Wi-Fi the sketch also pulls mujard_http.h for the account path
	rawCpp        bool // a cpp: block: hand-written C++ may call any of the runtime, so none of it is switched off
	controlLabels []string
	askIDs        map[*Call]int
	askN          int

	// calls that can suspend a task: ask, serial/uart read_line, http.get/post, and calls of fns that wait.
	// The statement that holds one first emits its begin/send and a MUJARD_WAIT_UNTIL; the expression then reads
	// the result (waitResults). A waiting fn gets a coroutine object per call site (fnInsts), declared in the
	// globals area once all fn structs exist (scanWaits).
	waitResults   map[*Call]string
	fnInsts       map[*Call]string
	waitCallOrder []*Call          // waiting-fn call sites in program order (for the instance declarations)
	waiting       map[*FnDecl]bool // fns whose run() suspends
	cycFn         map[*FnDecl]bool // a waiting fn that (in)directly calls itself: C320 at the call
	wantType      map[Node]*Type   // the declared type a value must have (store.load without a default / fixed text[N])
	retType       *Type            // the current function's declared result type
	uartLines     []string         // uart objects that must buffer incoming lines for read_line
	serialLines   bool             // the program uses serial.read_line()
	httpN         int              // http call sites get their own numbers (http_begin/done/result)

	// a test build (testgen.go): the program plus its 'test' blocks, for the PC
	testMode   bool
	inTestBody bool
	usesBle    bool // 'use mujard.ble' actually called: the Pico W needs its Bluetooth stack in the FQBN
	tests      []TestInfo
}

// savedVar is a 'saved var' to register, with the module it came from ("" = the program's own).
type savedVar struct {
	d   *VarDecl
	mod string
}

// modGen is a module's declarations in the shape the generator uses.
type modGen struct {
	enums    map[string]*EnumDecl
	structs  map[string]*StructDecl
	fnDecls  map[string]*FnDecl
	devices  map[string]*DeviceDecl
	varTypes map[string]*TypeExpr
}

type taskCtx struct {
	fields   []string // hoisted locals: "int32_t i;"
	seen     map[string]bool
	waits    int
	alias    map[string]string // a for-over-pins variable -> the element expression it stands for
	fnResult bool              // inside a fn that waits: 'return x' stores into result_
}

func (g *gen) errorf(code string, pos Pos, args ...string) {
	for _, e := range g.errs {
		if e.Pos == pos && e.Code == code && e.File == curSrcFile {
			return
		}
	}
	d := diag(code, pos, args...)
	d.File = curSrcFile // "" for the program, the module's name while its code is generated
	g.errs = append(g.errs, d)
}

func (g *gen) id(prefix string) string {
	g.n++
	return fmt.Sprintf("mujard_%s_%d", prefix, g.n)
}

// later reports a feature that comes with a later phase (C320).
func (g *gen) later(pos Pos, feature, phase string) string {
	g.errorf("C320", pos, "feature", feature, "phase", phase)
	return "0"
}

const (
	phaseApp   = "Phase 9 (phone and cloud features)"
	phaseLibs  = "Phase 8 (libraries)"
	phaseTests = "Phase 10 (tests in the simulator)"
	phaseLater = "a later Phase 4 update"
)

// cppReserved are names that would clash in C++: keywords, Arduino macros and functions.
var cppReserved = map[string]bool{}

func init() {
	for _, w := range strings.Fields(`alignas alignof and_eq asm auto bitand bitor bool break case catch char char16_t char32_t
	class compl const constexpr const_cast continue decltype default delete do double dynamic_cast else enum explicit export
	extern false float for friend goto if inline int long mutable namespace new noexcept not_eq nullptr operator or_eq private
	protected public register reinterpret_cast return short signed sizeof static static_assert static_cast struct switch
	template this thread_local throw true try typedef typeid typename union unsigned using virtual void volatile wchar_t while
	xor xor_eq and or not min max abs round constrain map delay millis micros random setup loop sq sqrt sin cos tan pow exp
	log bit bitRead bitSet bitClear bitWrite highByte lowByte HIGH LOW INPUT OUTPUT INPUT_PULLUP LED_BUILTIN PI HALF_PI
	TWO_PI DEG_TO_RAD RAD_TO_DEG Serial String byte word boolean tone noTone pulseIn shiftIn shiftOut pinMode digitalRead
	digitalWrite analogRead analogWrite interrupts noInterrupts attachInterrupt detachInterrupt yield F A0 A1 A2 A3 A4 A5
	A6 A7 B0 B1 EEPROM Wire SPI time printf mujard mujard_user`) {
		cppReserved[w] = true
	}
}

func cname(name string) string {
	if cppReserved[name] {
		return "u_" + name
	}
	return name
}

// qname is a type's C++ name: 'motors.Point' is motors::Point (each part renamed like cname, so a module
// called 'setup' is u_setup::Point - matching its generated namespace).
func qname(name string) string {
	parts := strings.Split(name, ".")
	for i, p := range parts {
		parts[i] = cname(p)
	}
	return strings.Join(parts, "::")
}

// flatName is a type's name for generated helpers: 'motors.Mode' -> motors_Mode.
func flatName(name string) string { return strings.ReplaceAll(name, ".", "_") }

var dimEnum = map[string]string{
	"frequency": "mujard::D_FREQ", "length": "mujard::D_LENGTH", "voltage": "mujard::D_VOLTAGE", "current": "mujard::D_CURRENT",
	"angle": "mujard::D_ANGLE", "temperature": "mujard::D_TEMP", "ratio": "mujard::D_RATIO", "speed": "mujard::D_SPEED",
	"memory": "mujard::D_MEMORY", "baud": "mujard::D_BAUD",
}

// unitFactor converts a value in a unit to the runtime's base unit (time: microseconds;
// frequency: millihertz; others: SI).
func unitFactor(u string) float64 {
	if i := strings.Index(u, "/"); i > 0 {
		return unitFactor(u[:i]) / (unitFactor(u[i+1:]) / 1e6) // m/s: metres per second, time in seconds
	}
	switch u {
	case "us":
		return 1
	case "ms":
		return 1e3
	case "s":
		return 1e6
	case "min":
		return 6e7
	case "h":
		return 3.6e9
	case "Hz":
		return 1e3
	case "kHz":
		return 1e6
	case "MHz":
		return 1e9
	case "mm", "mV", "mA":
		return 1e-3
	case "cm":
		return 1e-2
	case "%":
		return 1e-2
	}
	return 1
}

func cfloat(v float64) string {
	s := strconv.FormatFloat(v, 'g', -1, 32) // the shortest text that is exactly this float: 0.2, not 0.200000003
	if !strings.ContainsAny(s, ".eE") {
		s += ".0"
	}
	return s + "f"
}

// ctype is the C++ type of a Mujard type.
func (g *gen) ctype(t *Type, size int) string {
	var s string
	switch t.K {
	case TInt:
		s = "int32_t"
	case TFloat:
		s = "float"
	case TBool:
		s = "bool"
	case TText:
		if size <= 0 {
			size = t.Cap
		}
		if size <= 0 || size == textMax {
			s = "mujard::Text"
		} else {
			s = fmt.Sprintf("mujard::TextN<%d>", size)
		}
	case TQty:
		switch t.Dim {
		case "time":
			s = "mujard::Time"
		case "frequency":
			s = "mujard::Freq"
		default:
			s = "float"
		}
	case TList:
		if size <= 0 {
			size = t.Cap
		}
		if size <= 0 {
			size = 8
		}
		s = fmt.Sprintf("mujard::List<%s, %d>", g.ctype(elemOf(t), 0), size)
	case TStruct, TEnum:
		switch {
		case isBuiltinEnum(t) && t.Name == "Level":
			s = "bool"
		case isBuiltinEnum(t) && t.Name == "BatteryType":
			s = "mujard::BatteryType"
		default:
			s = qname(t.Name)
		}
	case TObject:
		switch t.Name {
		case "stopwatch":
			s = "mujard::Stopwatch"
		case "time_of_day":
			s = "mujard::TimeOfDay"
		case "date":
			s = "mujard::Date"
		case "http_reply":
			s = "mujard::HttpReply"
		case "can":
			s = "mujard::CanBus"
		case "can_message":
			s = "mujard::CanMessage"
		case "timer":
			s = "mujard::Timer"
		case "queue":
			if t.Elem != nil && t.Cap > 0 {
				s = fmt.Sprintf("mujard::Queue<%s, %d>", g.ctype(t.Elem, 0), t.Cap)
			} else {
				s = "mujard::Queue<int32_t, 8>"
			}
		case "json":
			s = "mujard::Json"
		case "mqtt_message":
			s = "mujard::MqttMessage"
		default:
			if g.info != nil && g.info.LibClasses[t.Name] != nil {
				s = t.Name // a described library class is its own C++ type
			} else {
				s = "auto"
			}
		}
	default:
		s = "auto"
	}
	if t.Opt && s != "auto" {
		return "mujard::Opt<" + s + ">"
	}
	return s
}

func declCap(te *TypeExpr) int {
	if te == nil || len(te.Args) == 0 {
		return 0
	}
	arg := 0
	if te.Name == "text" {
		arg = 0
	} else if te.Name == "list" && len(te.Args) > 1 {
		arg = 1
	} else {
		return 0
	}
	n, _ := strconv.Atoi(te.Args[arg].Name)
	return n
}

func (g *gen) typeOf(n Node) *Type {
	if t, ok := g.types[n]; ok && t != nil {
		return t
	}
	return tAny
}

// setWant records the type a value expression must have (store.load without a default uses it for T?).
func (g *gen) setWant(n Node, t *Type) {
	if n == nil || t == nil {
		return
	}
	if g.wantType == nil {
		g.wantType = map[Node]*Type{}
	}
	g.wantType[n] = t
	if nonOpt(t).K == TText {
		switch e := n.(type) {
		case *Binary:
			if e.Op == "+" {
				g.setWant(e.X, t)
				g.setWant(e.Y, t)
			}
		case *CondExpr:
			g.setWant(e.Then, t)
			g.setWant(e.Else, t)
		}
	}
	if lt := nonOpt(t); lt.K == TList {
		if e, ok := n.(*ListLit); ok {
			for _, it := range e.Items {
				g.setWant(it, elemOf(lt))
			}
		}
	}
}

// wantText is the fixed text type an expression should build, if one is known.
func (g *gen) wantText(n Node) *Type {
	if t := g.wantType[n]; t != nil && nonOpt(t).K == TText {
		return t
	}
	if t := g.typeOf(n); nonOpt(t).K == TText {
		return t
	}
	return nil
}

// textCtor is Text or TextN<N> for the expression's expected Mujard text type.
func (g *gen) textCtor(n Node) string {
	if t := g.wantText(n); t != nil {
		if cap := textCapOf(t); cap != textMax {
			return fmt.Sprintf("mujard::TextN<%d>", cap)
		}
	}
	return "mujard::Text"
}

// Generate turns a checked program into an Arduino sketch.
func Generate(f *File, info *Info, srcName string) (*Sketch, []Diagnostic) {
	return generateSketch(f, info, srcName, false)
}

// generateSketch is Generate, or with testMode the PC test build (testgen.go).
func generateSketch(f *File, info *Info, srcName string, testMode bool) (*Sketch, []Diagnostic) {
	g := &gen{f: f, info: info, board: info.Board, src: path.Base(srcName), testMode: testMode, enums: map[string]*EnumDecl{},
		structs: map[string]*StructDecl{}, fnDecls: map[string]*FnDecl{}, devices: map[string]*DeviceDecl{},
		varTypes: map[string]*TypeExpr{}, volatiles: map[string]bool{}, blockNames: map[string]string{},
		types: info.Types, decl: info.Declared, mods: info.Modules, modGen: map[string]*modGen{}}
	for name, mc := range g.mods {
		mg := &modGen{enums: map[string]*EnumDecl{}, structs: map[string]*StructDecl{}, fnDecls: map[string]*FnDecl{},
			devices: map[string]*DeviceDecl{}, varTypes: map[string]*TypeExpr{}}
		for _, it := range mc.file.Items {
			switch d := it.(type) {
			case *EnumDecl:
				mg.enums[d.Name] = d
			case *StructDecl:
				mg.structs[d.Name] = d
			case *FnDecl:
				mg.fnDecls[d.Name] = d
			case *DeviceDecl:
				mg.devices[d.Name] = d
			case *VarDecl:
				mg.varTypes[d.Name] = d.Type
			}
		}
		g.modGen[name] = mg
	}
	g.run()
	if len(g.secrets) > 0 {
		// The build writes the secrets header; a secret the project doesn't have stops the build here, on the
		// .muj line that needs it, with a message that says what to add (spec 11.1).
		g.head.add(0, "#include \"%s\"", secretsHeader)
		var names []string
		for n := range g.secrets {
			names = append(names, n)
		}
		sort.Strings(names)
		for _, n := range names {
			g.head.add(0, "#ifndef %s", n)
			g.head.add(g.secrets[n], "#error \"This program needs the secret %s: add it to the project's secrets (for a local build, set the environment variable MUJARD_SECRET_%s)\"", n, n)
			g.head.add(0, "#endif")
		}
	}
	// notify can reach the account through the server when the program is on Wi-Fi and the build carries
	// a notify key - that path lives in mujard_http.h (a program that already 'use's it has it already).
	if g.usesNotify && info.Libs["wifi"] && !info.Libs["http"] {
		g.head.add(0, "#include \"mujard_http.h\"")
	}
	// The program lives in its own namespace: inside it, the program's names win over anything a core or library
	// declares globally (an ESP32 sketch couldn't otherwise have a variable called 'open', 'time' or 'connect').
	var open, closeNS section
	open.add(0, "")
	open.add(0, "namespace mujard_user {")
	closeNS.add(0, "")
	closeNS.add(0, "}  // namespace mujard_user")
	closeNS.add(0, "")
	closeNS.add(0, "void setup() { mujard_user::setup(); }")
	closeNS.add(0, "void loop() { mujard_user::loop(); }")
	files := map[string]string{
		"sketch.ino": "// Generated by Mujard from " + commentSafe(g.src) + ". The program is in mujard_program.cpp; change the .muj file, not these.\n" +
			"#include \"mujard_rt.h\"\n",
		"mujard_program.cpp": assemble(g.src, "mujard_program.cpp", g.head, open, g.globals, g.fns, g.setup, g.loop, closeNS),
	}
	if len(info.LibDeps) > 0 {
		// The build server installs these before compiling ('version' pins it); mujard_libs.json sits next
		// to the sketch - arduino-cli ignores it.
		if b, err := json.MarshalIndent(info.LibDeps, "", "  "); err == nil {
			files["mujard_libs.json"] = string(b) + "\n"
		}
	}
	need := []string{"mujard_rt.h", "mujard_rt.cpp"}
	for _, lib := range runtimeLibs {
		if info.Libs[lib] {
			need = append(need, "mujard_"+lib+".h")
		}
	}
	if !info.Libs["wifi"] && (info.Libs["clock"] || info.Libs["http"] || info.Libs["mqtt"] || info.Libs["ethernet"]) {
		need = append(need, "mujard_wifi.h") // they include it (mujard_ethernet.h for net_up())
	}
	if g.usesNotify && info.Libs["wifi"] && !info.Libs["http"] {
		need = append(need, "mujard_http.h") // notify to the account (Phase 9)
	}
	if testMode {
		need = append(need, "mujard_test.h", "mujard_test.cpp")
	}
	for _, name := range need {
		b, err := runtimeFS.ReadFile("runtime/" + name)
		if err != nil {
			panic(err)
		}
		files[name] = string(b)
	}
	// mujard_rt.h always includes mujard_config.h. Hand-written C++ (a cpp: block, or a project's own .h/.cpp files) may
	// call any of the runtime: it keeps all of it (a config with nothing switched off).
	if !g.rawCpp && len(info.UseFiles) == 0 && !testMode {
		files[runtimeConfigFile] = runtimeConfig(files)
	} else {
		files[runtimeConfigFile] = runtimeConfigHeader(nil)
	}
	var secrets []string
	for n := range g.secrets {
		secrets = append(secrets, n)
	}
	sort.Strings(secrets)
	fqbn := fqbns[g.board.Name]
	if g.usesBle && g.board.Name == "picow" {
		fqbn += ":ipbtstack=ipv4btcble" // the CYW43's BLE stack only links in when the board option asks for it
	}
	return &Sketch{Files: files, FQBN: fqbn, Board: g.board, Libraries: info.LibDeps, Secrets: secrets, Tests: g.tests}, g.errs
}

func (g *gen) run() {
	f := g.f
	g.head.add(0, "// Generated by Mujard from %s for the %s - do not edit; change the .muj file instead.", commentSafe(g.src), g.board.Display)
	g.head.add(0, "#include \"mujard_rt.h\"")
	for _, lib := range runtimeLibs {
		if g.info.Libs[lib] {
			g.head.add(0, "#include \"mujard_%s.h\"", lib)
		}
	}
	if g.testMode {
		g.head.add(0, "#include \"mujard_test.h\"")
	}
	for _, it := range f.Items {
		switch d := it.(type) {
		case *UseDecl:
			switch d.From {
			case "arduino":
				switch {
				case g.info.UseFiles[d]:
					g.head.add(d.At.Line, "#include \"%s\"", d.Path) // a file in the project
				case g.info.UseLibs[d] != nil:
					hdr := g.info.UseLibs[d].Header
					if hdr == "" {
						hdr = headerFromName(d.Path)
					}
					g.head.add(d.At.Line, "#include <%s>", hdr)
					for _, more := range g.info.UseLibs[d].Headers {
						g.head.add(d.At.Line, "#include <%s>", more)
					}
				default:
					hdr := d.Path
					if !isFilePath(hdr) {
						hdr = headerFromName(hdr)
					}
					g.head.add(d.At.Line, "#include <%s>", hdr)
				}
			case "pio":
				hdr := headerFromName(d.Path)
				if desc := g.info.UseLibs[d]; desc != nil && desc.Header != "" {
					hdr = desc.Header
				}
				g.head.add(d.At.Line, "#include <%s>", hdr)
			case "idf":
				// An ESP-IDF component is a build dependency, not an include: it is recorded in
				// mujard_libs.json and cpp: blocks include its headers by name.
				if desc := g.info.UseLibs[d]; desc != nil && desc.Header != "" {
					g.head.add(d.At.Line, "#include <%s>", desc.Header)
				}
			case "module":
				g.later(d.At, "use "+d.From, phaseLibs)
			case "mujard":
				lib := strings.TrimPrefix(d.Path, "mujard.")
				switch lib {
				case "servo", "ultrasonic", "i2c", "spi", "rtc", "shift", "uart", "can", "ble", "ethernet", "usb", "timer", "queue", "mqtt", "math", "regs":
				default:
					g.later(d.At, "mujard."+lib, phaseLibs)
				}
			}
		case *ImportDecl:
			g.modUse(d)
		case *EnumDecl:
			g.enums[d.Name] = d
		case *StructDecl:
			g.structs[d.Name] = d
		case *FnDecl:
			g.fnDecls[d.Name] = d
		case *DeviceDecl:
			g.devices[d.Name] = d
		case *VarDecl:
			g.varTypes[d.Name] = d.Type
		case *TaskDecl:
			g.blockNames[d.Name] = "mujard_task_" + d.Name
		}
	}
	g.findVolatiles()
	g.head.add(0, "")
	g.head.add(0, "const char* mujard::board_name() { return %q; }", g.board.Display)

	// types
	for _, it := range f.Items {
		switch d := it.(type) {
		case *ImportDecl:
			g.modTypes(d)
		case *EnumDecl:
			var vals []string
			for _, v := range d.Values {
				vals = append(vals, cname(v))
			}
			g.globals.add(d.At.Line, "enum class %s : uint8_t { %s };", cname(d.Name), strings.Join(vals, ", "))
			g.globals.add(0, "static const char* mujard_name_%s(%s v) {", d.Name, cname(d.Name))
			g.globals.add(0, "    switch (v) {")
			for _, v := range d.Values {
				g.globals.add(0, "        case %s::%s: return %q;", cname(d.Name), cname(v), d.Name+"."+v)
			}
			g.globals.add(0, "    }")
			g.globals.add(0, "    return \"?\";")
			g.globals.add(0, "}")
		case *StructDecl:
			g.globals.add(d.At.Line, "struct %s {", cname(d.Name))
			for _, fl := range d.Fields {
				t := g.fieldType(fl)
				g.globals.add(fl.At.Line, "    %s %s;", g.ctype(t, declCap(fl.Type)), cname(fl.Name))
			}
			g.globals.add(0, "};")
		}
	}

	// devices, variables, block state
	g.setup.add(0, "")
	g.setup.add(0, "void setup() {")
	g.setup.add(0, "    mujard::safe_state_hook = mujard_safe_state;")
	g.setup.add(0, "    mujard_safe_state();  // outputs are safe before anything else runs (spec 11)")
	g.setup.add(0, "    mujard::begin();")
	// Core 0 opts into the peer core's flash lockout before anything can commit (RP-family; a no-op elsewhere).
	g.setup.add(0, "    mujard::coreInitFlashSafe();")
	hasSafe := false
	for _, it := range f.Items {
		switch d := it.(type) {
		case *ImportDecl:
			g.modGlobals(d)
		case *DeviceDecl:
			g.device(d)
		case *VarDecl:
			g.globalVar(d)
			if d.Kind == "saved" {
				g.savedVars = append(g.savedVars, savedVar{d: d})
			}
		case *SpecialBlock:
			if d.Kind == "safe_state" {
				hasSafe = true
				g.fns.add(d.At.Line, "static void mujard_safe_state() {")
				g.block(&g.fns, d.Body, 1)
				g.fns.add(0, "}")
			}
		}
	}
	if !hasSafe {
		g.fns.add(0, "static void mujard_safe_state() {}")
	}
	if len(g.savedVars) > 0 {
		h := fnv.New32a()
		for _, sv := range g.savedVars {
			name := sv.d.Name
			if sv.mod == "" {
				fmt.Fprintf(h, "%s:%s;", name, g.ctype(g.declVarType(sv.d), declCap(sv.d.Type)))
				continue
			}
			g.withModule(g.mods[sv.mod], sv.mod, false, func() {
				fmt.Fprintf(h, "%s:%s;", sv.mod+"."+name, g.ctype(g.declVarType(sv.d), declCap(sv.d.Type)))
			})
		}
		g.setup.add(0, "    mujard::saved_begin(0x%08Xu);", h.Sum32())
		for _, sv := range g.savedVars {
			name := cname(sv.d.Name)
			if sv.mod != "" {
				name = cname(sv.mod) + "::" + name
			}
			g.setup.add(sv.d.At.Line, "    mujard::saved_add(&%s, sizeof(%s));", name, name)
		}
		g.setup.add(0, "    mujard::saved_load();")
	}

	// app controls, then hello (filled in at the end, once it is known whether the program uses the app at all)
	for _, it := range f.Items {
		if d, ok := it.(*ControlDecl); ok {
			g.controlDecl(d)
		}
	}
	appBeginAt := len(g.setup)
	g.setup.add(0, "")

	// functions: waiting fns are Task coroutines, so their structs and one call object per call site must exist
	// before any body that calls them (scanWaits fills fnInsts; a caller's body comes later in this section).
	g.scanWaits()
	for _, it := range f.Items {
		switch d := it.(type) {
		case *FnDecl:
			if g.waitingFn(d) {
				g.fns.add(d.At.Line, "struct mujard_Fn_%s : mujard::Task {", cname(d.Name))
				for _, p := range d.Params {
					t := g.declTypeFromTypeExpr(p.Type)
					if t == nil {
						t = tAny
					}
					g.fns.add(p.At.Line, "    %s %s;", g.ctype(t, declCap(p.Type)), cname(p.Name))
				}
				if d.Result != nil {
					if t := g.declTypeFromTypeExpr(d.Result); t != nil {
						g.fns.add(d.Result.At.Line, "    %s result_;", g.ctype(t, declCap(d.Result)))
					}
				}
				g.fns.add(0, "    void run();")
				g.fns.add(0, "};")
			} else {
				g.fns.add(d.At.Line, "%s;", g.fnHeader(d))
			}
		case *ImportDecl:
			g.modProtos(d)
		}
	}
	for _, c := range g.waitCallOrder {
		if id, ok := c.Fn.(*Ident); ok {
			if d := g.fnDecls[id.Name]; d != nil {
				g.fns.add(c.At.Line, "static mujard_Fn_%s %s;", cname(d.Name), g.fnInsts[c])
			}
		}
	}
	for _, it := range f.Items {
		switch d := it.(type) {
		case *FnDecl:
			g.fnDecl(d)
		case *ImportDecl:
			g.modFns(d)
		}
	}

	// blocks: the scheduler phases stay fixed (spec 6); priority N only orders entries inside a phase.
	g.loop.add(0, "")
	g.loop.add(0, "void loop() {")
	g.loop.add(0, "    if (!mujard::tick()) return;")
	type schedPart struct {
		phase    int
		priority uint64
		order    int
		lines    section
	}
	var sched []schedPart
	order := 0
	schedule := func(phase int, priority uint64, emit func()) {
		old := g.loop
		g.loop = section{}
		emit()
		if len(g.loop) > 0 {
			sched = append(sched, schedPart{phase: phase, priority: priority, order: order, lines: g.loop})
		}
		g.loop = old
		order++
	}
	for _, it := range f.Items {
		switch d := it.(type) {
		case *EventBlock:
			phase := map[string]int{"on": 0, "when": 1, "every": 2, "after": 2, "time": 2}[d.Kind]
			schedule(phase, d.Priority, func() { g.topEvent(d) })
		case *TaskDecl:
			schedule(3, d.Priority, func() { g.task(d) })
		case *LimitDecl:
			g.limit(d)
		case *CppBlock:
			g.rawCpp = true
			for _, l := range strings.Split(strings.TrimRight(d.Code, "\n"), "\n") {
				g.fns.add(d.At.Line+1, "%s", l)
			}
		case *SpecialBlock:
			if d.Kind == "test" {
				if g.testMode {
					schedule(3, 0, func() { g.testBlock(d) })
				} else {
					g.fns.add(d.At.Line, "// test %s runs on the PC (mujard test), never on the board", cString(d.Name))
				}
			}
		}
	}
	sort.SliceStable(sched, func(i, j int) bool {
		if sched[i].phase != sched[j].phase {
			return sched[i].phase < sched[j].phase
		}
		if sched[i].priority != sched[j].priority {
			return sched[i].priority > sched[j].priority
		}
		return sched[i].order < sched[j].order
	})
	for _, part := range sched {
		g.loop = append(g.loop, part.lines...)
	}
	if g.testMode && len(g.tests) > 0 {
		// A halt freezes the program's blocks, but a test still runs so it can check stopped() (spec 6.7).
		var halt strings.Builder
		halt.WriteString("    if (!mujard::tick()) {\n")
		for k := range g.tests {
			fmt.Fprintf(&halt, "        if (mujard::test_selected == %d) mujard_test_%d.run();\n", k, k)
		}
		halt.WriteString("        return;\n    }")
		for i := range g.loop {
			if g.loop[i].text == "    if (!mujard::tick()) return;" {
				g.loop[i].text = halt.String()
				break
			}
		}
	}
	g.loop = append(g.loop, g.deferred...)
	g.loop.add(0, "}")
	if g.usesApp {
		g.setup[appBeginAt].text = fmt.Sprintf("    mujard::app_begin(%s, 0x%08Xu);  // hello to the app's dashboard", cString(g.src), g.appHash())
	} else {
		g.setup = append(g.setup[:appBeginAt], g.setup[appBeginAt+1:]...)
	}
	if g.serialLines {
		g.setup.add(0, "    mujard::serial_want_lines();")
	}
	seen := map[string]bool{}
	for _, obj := range g.uartLines {
		if !seen[obj] {
			seen[obj] = true
			g.setup.add(0, "    %s.wantLines();", obj)
		}
	}
	g.setup.add(0, "}")
}

// ---- modules (spec 10.3)
//
// A module's declarations are generated inside 'namespace <mod>' (in mujard_user), so 'import motors' becomes
// motors::speed, motors::forward(), motors::Mode::run. Its event blocks, tasks and other program behaviour are
// not generated: only declarations are imported.

// withModule runs fn with the generator switched to a module's world: its types, declarations, board and files
// (#line markers point at the module's own file). qual asks for qualified 'mod::name' on module globals - needed
// only when a module expression is emitted outside the namespace (call-site defaults).
func (g *gen) withModule(mc *checker, mod string, qual bool, fn func()) {
	if mc == nil {
		return
	}
	saveFile := curSrcFile
	saveTypes, saveDecl := g.types, g.decl
	saveMod, saveScope, saveQual := g.curMod, g.curModScope, g.modQual
	saveBoard := g.board
	saveEnums, saveStructs, saveFns, saveDevs, saveVars := g.enums, g.structs, g.fnDecls, g.devices, g.varTypes
	mg := g.modGen[mod]
	g.types, g.decl = mc.types, mc.declared
	g.curMod, g.curModScope, g.modQual = mod, mc.global, qual
	if mc.board != nil {
		g.board = mc.board
	}
	if mg != nil {
		g.enums, g.structs, g.fnDecls, g.devices, g.varTypes = mg.enums, mg.structs, mg.fnDecls, mg.devices, mg.varTypes
	}
	curSrcFile = mc.modFile
	fn()
	curSrcFile = saveFile
	g.types, g.decl = saveTypes, saveDecl
	g.curMod, g.curModScope, g.modQual = saveMod, saveScope, saveQual
	g.board = saveBoard
	g.enums, g.structs, g.fnDecls, g.devices, g.varTypes = saveEnums, saveStructs, saveFns, saveDevs, saveVars
}

// modUse pulls in the C++ headers a module's 'use' lines ask for.
func (g *gen) modUse(d *ImportDecl) {
	mc := g.mods[d.Name]
	if mc == nil {
		g.later(d.At, "import", phaseLibs) // no module loader (a bare Check/Generate)
		return
	}
	for _, it := range mc.file.Items {
		u, ok := it.(*UseDecl)
		if !ok || u.From != "arduino" {
			continue
		}
		hdr := u.Path
		if !strings.HasSuffix(hdr, ".h") {
			hdr = strings.ReplaceAll(hdr, " ", "_") + ".h"
		}
		g.head.add(u.At.Line, "#include <%s>", hdr)
	}
}

// modTypes emits a module's structs and enums in its namespace, and the enum-name helpers (outside it).
func (g *gen) modTypes(d *ImportDecl) {
	mc := g.mods[d.Name]
	if mc == nil || mc.failed {
		return
	}
	has := false
	for _, it := range mc.file.Items {
		switch it.(type) {
		case *EnumDecl, *StructDecl:
			has = true
		}
	}
	if !has {
		return
	}
	g.globals.add(d.At.Line, "namespace %s {", cname(d.Name))
	g.withModule(mc, d.Name, false, func() {
		for _, it := range mc.file.Items {
			switch x := it.(type) {
			case *EnumDecl:
				var vals []string
				for _, v := range x.Values {
					vals = append(vals, cname(v))
				}
				g.globals.add(x.At.Line, "enum class %s : uint8_t { %s };", cname(x.Name), strings.Join(vals, ", "))
			case *StructDecl:
				g.globals.add(x.At.Line, "struct %s {", cname(x.Name))
				for _, fl := range x.Fields {
					t := g.fieldType(fl)
					g.globals.add(fl.At.Line, "    %s %s;", g.ctype(t, declCap(fl.Type)), cname(fl.Name))
				}
				g.globals.add(0, "};")
			}
		}
	})
	g.globals.add(0, "%s", "}  // namespace "+d.Name)
	for _, it := range mc.file.Items {
		x, ok := it.(*EnumDecl)
		if !ok {
			continue
		}
		g.globals.add(x.At.Line, "static const char* mujard_name_%s(%s v) {", flatName(d.Name+"."+x.Name), qname(d.Name+"."+x.Name))
		g.globals.add(0, "    switch (v) {")
		for _, v := range x.Values {
			g.globals.add(0, "        case %s::%s: return %q;", qname(d.Name+"."+x.Name), cname(v), x.Name+"."+v)
		}
		g.globals.add(0, "    }")
		g.globals.add(0, "    return \"?\";")
		g.globals.add(0, "}")
	}
}

// modGlobals emits a module's devices and variables in its namespace; initializers run in setup().
func (g *gen) modGlobals(d *ImportDecl) {
	mc := g.mods[d.Name]
	if mc == nil || mc.failed {
		return
	}
	has := false
	for _, it := range mc.file.Items {
		switch it.(type) {
		case *DeviceDecl, *VarDecl:
			has = true
		}
	}
	if !has {
		return
	}
	g.globals.add(d.At.Line, "namespace %s {", cname(d.Name))
	g.withModule(mc, d.Name, false, func() {
		for _, it := range mc.file.Items {
			switch x := it.(type) {
			case *DeviceDecl:
				g.device(x)
			case *VarDecl:
				g.globalVar(x)
				if x.Kind == "saved" {
					g.savedVars = append(g.savedVars, savedVar{d: x, mod: d.Name})
				}
			}
		}
	})
	g.globals.add(0, "%s", "}  // namespace "+d.Name)
}

// modProtos / modFns emit a module's functions in its namespace.
func (g *gen) modProtos(d *ImportDecl) {
	mc := g.mods[d.Name]
	if mc == nil || mc.failed {
		return
	}
	has := false
	for _, it := range mc.file.Items {
		if _, ok := it.(*FnDecl); ok {
			has = true
		}
	}
	if !has {
		return
	}
	g.fns.add(d.At.Line, "namespace %s {", cname(d.Name))
	g.withModule(mc, d.Name, false, func() {
		for _, it := range mc.file.Items {
			if x, ok := it.(*FnDecl); ok {
				g.fns.add(x.At.Line, "%s;", g.fnHeader(x))
			}
		}
	})
	g.fns.add(0, "}")
}

func (g *gen) modFns(d *ImportDecl) {
	mc := g.mods[d.Name]
	if mc == nil || mc.failed {
		return
	}
	has := false
	for _, it := range mc.file.Items {
		if _, ok := it.(*FnDecl); ok {
			has = true
		}
	}
	if !has {
		return
	}
	g.fns.add(d.At.Line, "namespace %s {", cname(d.Name))
	g.withModule(mc, d.Name, false, func() {
		for _, it := range mc.file.Items {
			if x, ok := it.(*FnDecl); ok {
				g.fnDecl(x)
			}
		}
	})
	g.fns.add(0, "%s", "}  // namespace "+d.Name)
}

// qualName gives a declared type its name as the user knows it: 'Point' inside motors is 'motors.Point'.
func (g *gen) qualName(name string) string {
	if g.curMod != "" {
		return g.curMod + "." + name
	}
	return name
}

// cppQual qualifies a module name for emission outside its namespace (a default argument at a call site):
// 'faster' in motors becomes motors::faster. Inside the namespace (modQual false) names stay plain.
func (g *gen) cppQual(name string) string {
	if g.modQual && g.curMod != "" {
		return cname(g.curMod) + "::" + name
	}
	return name
}

func (g *gen) fieldType(p Param) *Type {
	if t := g.declTypeFromTypeExpr(p.Type); t != nil {
		return t
	}
	return tAny
}

// findVolatiles marks globals that 'fast' handlers change: they run as interrupts. A variable wider
// than a byte also needs interrupts-off writes and atomic reads in the rest of the program (spec 6.5).
func (g *gen) findVolatiles() {
	for _, it := range g.f.Items {
		ev, ok := it.(*EventBlock)
		if !ok || !ev.Fast {
			continue
		}
		walkStmts(ev.Body, func(n Node) {
			if a, ok := n.(*AssignStmt); ok {
				switch t := a.Target.(type) {
				case *Ident:
					g.volatiles[t.Name] = true
				case *Member:
					if id, ok := t.X.(*Ident); ok {
						g.volatiles[id.Name+"."+t.Name] = true // a module's variable
					}
				}
			}
		})
	}
	// Only plain values can take the interrupt-safe path (a bool already updates in one byte; an
	// optional, text, list or struct can't).
	size := func(t *Type) int {
		if t == nil || t.Opt {
			return 0
		}
		switch t.K {
		case TBool:
			return 1
		case TInt, TFloat, TQty, TEnum:
			return 4
		}
		return 0
	}
	g.atomics = map[string]bool{}
	for _, it := range g.f.Items {
		d, ok := it.(*VarDecl)
		if !ok || !g.volatiles[d.Name] {
			continue
		}
		var t *Type
		if d.Type != nil {
			t = g.declTypeFromTypeExpr(d.Type)
		}
		if t == nil {
			t = g.typeOf(d.Value)
		}
		if size(t) > 1 {
			g.atomics[d.Name] = true
		}
	}
	for mod, mc := range g.mods { // a program fast handler can write a module's variable (motors.speed)
		for name := range g.volatiles {
			parts := strings.SplitN(name, ".", 2)
			if len(parts) != 2 || parts[0] != mod {
				continue
			}
			if sym := mc.global.syms[parts[1]]; sym != nil && size(sym.t) > 1 {
				g.atomics[name] = true
			}
		}
	}
}

// atomicKey is the atomics key of an assignment target ("x", "mod.x"), or "".
func atomicKey(t Node) string {
	switch t := t.(type) {
	case *Ident:
		return t.Name
	case *Member:
		if id, ok := t.X.(*Ident); ok {
			return id.Name + "." + t.Name
		}
	}
	return ""
}

func walkStmts(b *Block, fn func(Node)) {
	if b == nil {
		return
	}
	for _, s := range b.Stmts {
		fn(s)
		switch x := s.(type) {
		case *IfStmt:
			for _, t := range x.Thens {
				walkStmts(t, fn)
			}
			walkStmts(x.Else, fn)
		case *WhileStmt:
			walkStmts(x.Body, fn)
		case *ForStmt:
			walkStmts(x.Body, fn)
		case *MatchStmt:
			for _, c := range x.Cases {
				walkStmts(c.Body, fn)
			}
		case *EventBlock:
			walkStmts(x.Body, fn)
		}
	}
}

// ---- declarations

func (g *gen) pinNumber(r PinRef) string {
	n, ok := g.board.ResolvePin(r.Text)
	if !ok {
		return "0"
	}
	return strconv.Itoa(n)
}

func hasOpt(opts []string, o string) bool {
	for _, x := range opts {
		if x == o {
			return true
		}
	}
	return false
}

func (g *gen) device(d *DeviceDecl) {
	name := cname(d.Name)
	line := d.At.Line
	inv := strconv.FormatBool(hasOpt(d.Options, "inverted"))
	pull := "0"
	if hasOpt(d.Options, "pullup") {
		pull = "1"
	} else if hasOpt(d.Options, "pulldown") {
		pull = "2"
	}
	// The constructor arguments of one pin device; devices are always built in place (they register themselves).
	ctorArgs := func(pin string) string {
		switch d.Role {
		case "input":
			return fmt.Sprintf("%s, %s, %s", pin, pull, inv)
		case "pwm":
			return pin
		}
		return fmt.Sprintf("%s, %s", pin, inv)
	}
	cls := map[string]string{"output": "mujard::Output", "input": "mujard::Input", "pwm": "mujard::Pwm"}[d.Role]
	switch d.Kind {
	case "pin":
		g.globals.add(line, "static %s %s(%s);", cls, name, ctorArgs(g.pinNumber(d.Pins[0])))
	case "analog":
		g.globals.add(line, "static mujard::Analog %s(%s);", name, g.pinNumber(d.Pins[0]))
	case "pins", "pinlist":
		var pins []string
		if d.Kind == "pins" {
			a, _ := g.board.ResolvePin(d.Pins[0].Text)
			b, _ := g.board.ResolvePin(d.Pins[1].Text)
			if d.ToEnd {
				b++
			}
			for p := a; p < b; p++ {
				pins = append(pins, strconv.Itoa(p))
			}
		} else {
			for _, r := range d.Pins {
				pins = append(pins, g.pinNumber(r))
			}
		}
		var items []string
		for _, p := range pins {
			items = append(items, "{"+ctorArgs(p)+"}")
		}
		g.globals.add(line, "static %s %s[%d] = {%s};", cls, name, len(pins), strings.Join(items, ", "))
	case "object":
		call, _ := d.Object.(*Call)
		if call == nil {
			return
		}
		g.objectDevice(d, call)
	}
}

// runtimeLibs are the 'use mujard.<lib>' libraries with a runtime header (mujard_<lib>.h).
var runtimeLibs = []string{"servo", "ultrasonic", "i2c", "spi", "rtc", "shift", "uart", "can", "ble", "ethernet", "usb", "timer", "queue", "store", "wifi", "clock", "http", "mqtt"}

// batteryTypes and resetReasons are the built-in enums' runtime constants (spec 4.6).
var batteryTypes = map[string]string{"lipo_1s": "mujard::BAT_LIPO_1S", "aa_x3": "mujard::BAT_AA_X3", "aa_x4": "mujard::BAT_AA_X4"}
var resetReasons = map[string]string{
	"power_on": "mujard::RESET_POWER_ON", "reset_button": "mujard::RESET_RESET_BUTTON", "watchdog": "mujard::RESET_WATCHDOG",
	"brownout": "mujard::RESET_BROWNOUT", "software": "mujard::RESET_SOFTWARE", "wake": "mujard::RESET_WAKE", "crash": "mujard::RESET_CRASH",
}

// builtinEnum is the C++ for a member of a built-in enum (spec 4.6): a Level is the pin level as a bool, a
// BatteryType or ResetReason the runtime's constant.
func (g *gen) builtinEnum(enum, member string, at Pos) string {
	switch enum {
	case "Level":
		return map[string]string{"high": "true", "low": "false"}[member]
	case "BatteryType":
		return batteryTypes[member]
	case "ResetReason":
		return resetReasons[member]
	}
	return g.later(at, enum+"."+member, phaseLater)
}

// objectDevice declares a library object: lid = Servo(pin 6), leds = shift.out_register(...), strip = Adafruit_NeoPixel(...)
func (g *gen) objectDevice(d *DeviceDecl, call *Call) {
	name := cname(d.Name)
	arg := func(i int, named string) string {
		for _, a := range call.Args {
			if a.Name == named {
				return g.expr(a.Value)
			}
		}
		pos := 0
		for _, a := range call.Args {
			if a.Name == "" {
				if pos == i {
					return g.expr(a.Value)
				}
				pos++
			}
		}
		return "0"
	}
	switch fn := call.Fn.(type) {
	case *Ident:
		// 'use "Servo"' without 'use mujard.servo': Servo() is the Arduino library's class (the checker resolved it so).
		name0 := fn.Name
		if gl, ok := globals[fn.Name]; ok && gl.lib != "" && !g.info.Libs[gl.lib] && g.info.LibClasses[fn.Name] != nil {
			name0 = ""
		}
		switch name0 {
		case "Servo":
			g.globals.add(d.At.Line, "static mujard::ServoMotor %s(%s);", name, arg(0, "pin"))
			return
		case "Ultrasonic":
			g.globals.add(d.At.Line, "static mujard::Ultrasonic %s(%s, %s);", name, arg(0, "trig"), arg(1, "echo"))
			return
		case "Uart":
			speed := "9600"
			for i, a := range call.Args {
				if a.Name == "speed" || (a.Name == "" && i == 2) {
					speed = "(uint32_t)(" + g.expr(a.Value) + ")"
				}
			}
			g.globals.add(d.At.Line, "static mujard::Uart %s(%s, %s, %s);", name, arg(0, "rx"), arg(1, "tx"), speed)
			return
		case "Can":
			if !g.board.HasCAN() {
				g.later(d.At, "CAN on this board", phaseLater)
				return
			}
			rate := "500000"
			for i, a := range call.Args {
				if a.Name == "rate" || (a.Name == "" && i == 2) {
					rate = "(int32_t)(" + g.expr(a.Value) + ")"
				}
			}
			g.globals.add(d.At.Line, "static mujard::CanBus %s(%s, %s, %s);", name, arg(0, "tx"), arg(1, "rx"), rate)
			return
		case "Timer":
			if !timerBackend(g.board) {
				g.later(d.At, "Timer on this board", phaseLater)
				return
			}
			g.globals.add(d.At.Line, "static mujard::Timer %s;", name)
			return
		case "Queue":
			// Queue(text, 8): the first arg names a value type, the second the capacity - the checker resolved both
			// into the call's type, and the C++ object is the bounded runtime Queue.
			elem, cap := "int32_t", "8"
			if id, ok := call.Args[0].Value.(*Ident); ok {
				if t, okT := typeNamed(id.Name); okT {
					elem = g.ctype(t, 0)
				}
			}
			if n, ok := call.Args[1].Value.(*NumberLit); ok {
				cap = strconv.Itoa(int(n.Int))
			}
			g.globals.add(d.At.Line, "static mujard::Queue<%s, %s> %s;", elem, cap, name)
			return
		case "battery_on":
			// battery = battery_on(analog A1, divider: 2, type: lipo_1s)
			divider, kind := "1.0f", "mujard::BAT_LIPO_1S"
			for i, a := range call.Args {
				switch {
				case a.Name == "divider" || (a.Name == "" && i == 1):
					divider = "(float)(" + g.expr(a.Value) + ")"
				case a.Name == "type" || (a.Name == "" && i == 2):
					kind = g.expr(a.Value) // lipo_1s or BatteryType.lipo_1s
				}
			}
			g.globals.add(d.At.Line, "static mujard::Battery %s(%s, %s, %s);", name, arg(0, "pin"), divider, kind)
			return
		}
		if g.info.LibClasses[fn.Name] != nil {
			var args []string
			for _, a := range call.Args {
				args = append(args, g.expr(a.Value))
			}
			g.globals.add(d.At.Line, "static %s;", libObject(g.libName(fn.Name), name, args))
			return
		}
		if g.info.Unchecked {
			var args []string
			for _, a := range call.Args {
				if a.Name != "" {
					g.later(a.Value.Start(), "named arguments to a C++ library", phaseLibs)
				}
				args = append(args, g.expr(a.Value))
			}
			g.globals.add(d.At.Line, "static %s;", libObject(fn.Name, name, args))
			return
		}
	case *Member:
		if id, ok := fn.X.(*Ident); ok && id.Name == "shift" && fn.Name == "out_register" {
			g.globals.add(d.At.Line, "static mujard::ShiftOut %s(%s, %s, %s);", name, arg(0, "data"), arg(1, "clock"), arg(2, "latch"))
			return
		}
		if id, ok := fn.X.(*Ident); ok && id.Name == "spi" && fn.Name == "device" {
			g.globals.add(d.At.Line, "static %s;", g.spiDeviceCtor(call, name))
			return
		}
	}
	g.later(d.At, "this kind of device", phaseLibs)
}

// libObject declares a C++ library object. With no arguments it is 'T name', never 'T name()': C++ reads that as a
// function declaration (the most vexing parse), and every 'name.begin()' after it fails to compile.
func libObject(class, name string, args []string) string {
	if len(args) == 0 {
		return class + " " + name
	}
	return fmt.Sprintf("%s %s(%s)", class, name, strings.Join(args, ", "))
}

// spiCall is spi.device(cs, speed, mode) used as a value; spiDeviceCtor is the same as a device declaration.
func (g *gen) spiCall(c *Call) string {
	a := argMap(c, "cs", "speed", "mode")
	speed := "1000000"
	if a["speed"] != nil {
		speed = "(uint32_t)((" + g.expr(a["speed"]) + ") / 1000)" // Mujard frequencies are in millihertz
	}
	mode := "0"
	if a["mode"] != nil {
		mode = g.expr(a["mode"])
	}
	cs := "0"
	if a["cs"] != nil {
		cs = g.expr(a["cs"])
	}
	return fmt.Sprintf("mujard::SpiDevice(%s, %s, %s)", cs, speed, mode)
}

func (g *gen) spiDeviceCtor(call *Call, name string) string {
	a := argMap(call, "cs", "speed", "mode")
	speed := "1000000"
	if a["speed"] != nil {
		speed = "(uint32_t)((" + g.expr(a["speed"]) + ") / 1000)"
	}
	mode := "0"
	if a["mode"] != nil {
		mode = g.expr(a["mode"])
	}
	cs := "0"
	if a["cs"] != nil {
		cs = g.expr(a["cs"])
	}
	return fmt.Sprintf("mujard::SpiDevice %s(%s, %s, %s);", name, cs, speed, mode)
}

func (g *gen) globalVar(d *VarDecl) {
	t := g.declVarType(d)
	g.setWant(d.Value, t)
	name := cname(d.Name)    // declared inside the module's namespace: plain
	ref, key := name, d.Name // ... but setup() and the volatiles key know it as mod::name / mod.name
	if g.curMod != "" {
		ref = cname(g.curMod) + "::" + name
		key = g.curMod + "." + d.Name
	}
	cv := ""
	if g.volatiles[key] {
		cv = "volatile "
	}
	ct := g.ctype(t, declCap(d.Type))
	if ct == "auto" {
		g.later(d.At, "a variable holding a library value", phaseLibs)
		return
	}
	g.globals.add(d.At.Line, "static %s%s %s;", cv, ct, name)
	if lit, ok := d.Value.(*ListLit); ok {
		for _, it := range lit.Items {
			g.setWant(it, elemOf(t))
			g.setup.add(d.At.Line, "    %s.add(%s);", ref, g.expr(it))
		}
		return
	}
	if isNone(d.Value) {
		return
	}
	if g.atomics[key] { // wider than a byte and written by a fast handler: even the init is atomic (6.5)
		g.setup.add(d.At.Line, "    noInterrupts();")
		g.inAtomic = true
		g.setup.add(d.At.Line, "    %s = %s;", ref, g.expr(d.Value))
		g.inAtomic = false
		g.setup.add(d.At.Line, "    interrupts();")
		return
	}
	g.setup.add(d.At.Line, "    %s = %s;", ref, g.expr(d.Value))
}

func (g *gen) declVarType(d *VarDecl) *Type {
	t := g.typeOf(d.Value)
	if d.Type != nil {
		if dt := g.declTypeFromTypeExpr(d.Type); dt != nil {
			t = dt
		}
	}
	return t
}

func (g *gen) declTypeFromTypeExpr(te *TypeExpr) *Type {
	var t *Type
	if bt, ok := typeNamed(te.Name); ok {
		t = bt
		if te.Name == "text" && len(te.Args) > 0 {
			t = textCap(declCap(te))
		}
	} else if te.Name == "list" {
		el := tAny
		if len(te.Args) > 0 {
			if e := g.declTypeFromTypeExpr(te.Args[0]); e != nil {
				el = e
			}
		}
		t = listOf(el)
	} else if _, ok := g.enums[te.Name]; ok {
		t = &Type{K: TEnum, Name: g.qualName(te.Name)}
	} else if _, ok := g.structs[te.Name]; ok {
		t = &Type{K: TStruct, Name: g.qualName(te.Name)}
	} else {
		return nil
	}
	if te.Optional {
		t = opt(t)
	}
	return t
}

func (g *gen) fnHeader(d *FnDecl) string {
	res := "void"
	if d.Result != nil {
		if t := g.declTypeFromTypeExpr(d.Result); t != nil {
			res = g.ctype(t, declCap(d.Result))
		}
	}
	var ps []string
	for _, p := range d.Params {
		t := g.declTypeFromTypeExpr(p.Type)
		if t == nil {
			t = tAny
		}
		ps = append(ps, g.ctype(t, declCap(p.Type))+" "+cname(p.Name))
	}
	return fmt.Sprintf("static %s %s(%s)", res, cname(d.Name), strings.Join(ps, ", "))
}

func (g *gen) fnDecl(d *FnDecl) {
	oldRet := g.retType
	g.retType = nil
	if d.Result != nil {
		g.retType = g.declTypeFromTypeExpr(d.Result)
	}
	if !g.waitingFn(d) {
		g.fns.add(d.At.Line, "%s {", g.fnHeader(d))
		g.block(&g.fns, d.Body, 1)
		g.fns.add(0, "}")
		g.retType = oldRet
		return
	}
	// A fn that waits is a coroutine: 'return x' stores into result_ and stops it; the caller restarts it,
	// steps it until it stops, and reads result_ (spec 4.4: a fn's own waits are part of the caller's wait).
	ctx := &taskCtx{seen: map[string]bool{}, fnResult: d.Result != nil}
	var body section
	g.inTask = ctx
	g.block(&body, d.Body, 2)
	g.inTask = nil
	g.retType = oldRet
	g.fns.add(d.At.Line, "void mujard_Fn_%s::run() {", cname(d.Name))
	g.fns.add(0, "    MUJARD_TASK_BEGIN();")
	g.fns = append(g.fns, body...)
	g.fns.add(0, "    MUJARD_TASK_END();")
	g.fns.add(0, "}")
}

// eachNode visits every node under n: statements and the expressions inside them. Typed nils (an IfStmt's Else,
// say) arrive as a Node holding a nil pointer, so every pointer case nil-checks.
func eachNode(n Node, fn func(Node)) {
	if n == nil {
		return
	}
	fn(n)
	switch e := n.(type) {
	case *Block:
		if e == nil {
			return
		}
		for _, s := range e.Stmts {
			eachNode(s, fn)
		}
	case *VarDecl:
		eachNode(e.Value, fn)
	case *AssignStmt:
		eachNode(e.Target, fn)
		eachNode(e.Value, fn)
	case *ExprStmt:
		eachNode(e.X, fn)
	case *SimpleStmt:
		eachNode(e.Value, fn)
	case *ExpectStmt:
		eachNode(e.Cond, fn)
		for _, a := range e.Args {
			eachNode(a.Value, fn)
		}
	case *IfStmt:
		for _, c := range e.Conds {
			eachNode(c, fn)
		}
		for _, t := range e.Thens {
			eachNode(t, fn)
		}
		eachNode(e.Else, fn)
	case *WhileStmt:
		eachNode(e.Cond, fn)
		eachNode(e.Body, fn)
	case *ForStmt:
		eachNode(e.Iter, fn)
		eachNode(e.Body, fn)
	case *MatchStmt:
		eachNode(e.Value, fn)
		for _, mc := range e.Cases {
			for _, pt := range mc.Patterns {
				eachNode(pt, fn)
			}
			eachNode(mc.Body, fn)
		}
	case *EventBlock:
		eachNode(e.Expr, fn)
		eachNode(e.Body, fn)
	case *Call:
		eachNode(e.Fn, fn)
		for _, a := range e.Args {
			eachNode(a.Value, fn)
		}
	case *Binary:
		eachNode(e.X, fn)
		eachNode(e.Y, fn)
	case *Unary:
		eachNode(e.X, fn)
	case *Member:
		eachNode(e.X, fn)
	case *Index:
		eachNode(e.X, fn)
		eachNode(e.Ix, fn)
	case *Conversion:
		eachNode(e.X, fn)
	case *CondExpr:
		eachNode(e.Then, fn)
		eachNode(e.Cond, fn)
		eachNode(e.Else, fn)
	case *StringLit:
		if e == nil {
			return
		}
		for _, p := range e.Parts {
			if p.Expr != nil {
				eachNode(p.Expr, fn)
			}
		}
	case *ListLit:
		if e == nil {
			return
		}
		for _, it := range e.Items {
			eachNode(it, fn)
		}
	case *RangeExpr:
		if e == nil {
			return
		}
		eachNode(e.Lo, fn)
		eachNode(e.Hi, fn)
	}
}

// needsWait reports whether evaluating this call suspends a task: a blocking built-in (ask, a line read, a network
// request) or a call of a fn that waits.
func (g *gen) needsWait(c *Call) bool {
	switch fn := c.Fn.(type) {
	case *Ident:
		if fn.Name == "ask" {
			return true
		}
		if d := g.fnDecls[fn.Name]; d != nil {
			return g.waitingFn(d)
		}
	case *Member:
		xt := nonOpt(g.typeOf(fn.X))
		if xt.K == TObject {
			switch {
			case xt.Name == "serial" && fn.Name == "read_line":
				return true
			case xt.Name == "uart" && fn.Name == "read_line":
				return true
			case xt.Name == "http" && (fn.Name == "get" || fn.Name == "post"):
				return true
			}
		}
	}
	return false
}

// waitingFn reports whether d's run() suspends: it has a 'wait', calls a blocking built-in, or calls a fn that
// waits. The set is a fixed point: fn f waits when it calls fn g that waits.
func (g *gen) waitingFn(d *FnDecl) bool {
	if g.waiting == nil {
		g.computeWaiting()
	}
	return g.waiting[d]
}

func (g *gen) computeWaiting() {
	g.waiting = map[*FnDecl]bool{}
	g.cycFn = map[*FnDecl]bool{}
	changed := true
	for changed {
		changed = false
		for _, d := range g.fnDecls {
			if g.waiting[d] {
				continue
			}
			suspends := false
			eachNode(d.Body, func(n Node) {
				if s, ok := n.(*SimpleStmt); ok && s.Kind == "wait" {
					suspends = true
				}
				if c, ok := n.(*Call); ok && g.needsWait(c) {
					suspends = true
				}
			})
			if suspends {
				g.waiting[d] = true
				changed = true
			}
		}
	}
	// A waiting fn that calls itself - directly or through other fns - can't work: each call site is one coroutine
	// object, so the second round would overwrite the first's paused state (spec 4.4 keeps recursion out of waits).
	edges := map[*FnDecl][]*FnDecl{}
	for _, d := range g.fnDecls {
		if !g.waiting[d] {
			continue
		}
		eachNode(d.Body, func(n Node) {
			c, ok := n.(*Call)
			if !ok {
				return
			}
			if id, ok := c.Fn.(*Ident); ok {
				if t := g.fnDecls[id.Name]; t != nil && g.waiting[t] {
					edges[d] = append(edges[d], t)
				}
			}
		})
	}
	var reach func(from, to *FnDecl, seen map[*FnDecl]bool) bool
	reach = func(from, to *FnDecl, seen map[*FnDecl]bool) bool {
		if from == to {
			return true
		}
		if seen[from] {
			return false
		}
		seen[from] = true
		for _, nx := range edges[from] {
			if reach(nx, to, seen) {
				return true
			}
		}
		return false
	}
	for _, d := range g.fnDecls {
		if g.waiting[d] && reach(d, d, map[*FnDecl]bool{}) {
			g.cycFn[d] = true
		}
	}
}

// scanWaits finds every call of a waiting fn in a place that can suspend (tasks, test blocks, waiting fns) and
// gives each site its own coroutine object, in program order. Called before any body is generated; the objects are
// declared between the fn structs and the fn bodies.
func (g *gen) scanWaits() {
	if g.fnDecls == nil {
		return
	}
	g.fnInsts = map[*Call]string{}
	g.waitResults = map[*Call]string{}
	scan := func(b *Block) {
		if b == nil {
			return
		}
		eachNode(b, func(n Node) {
			c, ok := n.(*Call)
			if !ok {
				return
			}
			id, ok := c.Fn.(*Ident)
			if !ok || g.fnDecls[id.Name] == nil || !g.waitingFn(g.fnDecls[id.Name]) {
				return
			}
			if g.fnInsts[c] == "" {
				g.fnInsts[c] = g.id("fn")
				g.waitCallOrder = append(g.waitCallOrder, c)
			}
		})
	}
	for _, it := range g.f.Items {
		switch d := it.(type) {
		case *TaskDecl:
			scan(d.Body)
		case *SpecialBlock:
			if d.Kind == "test" {
				scan(d.Body)
			}
		case *FnDecl:
			if g.waitingFn(d) {
				scan(d.Body)
			}
		}
	}
}

func (g *gen) limit(d *LimitDecl) {
	r, ok := d.Range.(*RangeExpr)
	name := g.typeOf(d.Target).Name
	if !ok || (name != "pwm" && name != "servo") {
		g.later(d.At, "limit on this kind of device", phaseLater)
		return
	}
	g.setup.add(d.At.Line, "    %s.limit(%s, %s);", g.expr(d.Target), g.expr(r.Lo), g.expr(r.Hi))
}

// ---- blocks

// coreLaunch emits the secondary-core launch for an 'on core N' block or task: the multicore include once,
// then RP-family's multicore_launch_core1 or ESP32-family's xTaskCreatePinnedToCore. 'priority N' maps to the
// FreeRTOS priority on ESP32 (1+N); RP-family has no task priority and says so.
func (g *gen) coreLaunch(at Pos, fn string, priority uint64, core int) {
	if !g.coreInc {
		g.coreInc = true
		if g.board.Family == "RP2040" || g.board.Family == "RP2350" {
			g.head.add(at.Line, "#include \"pico/multicore.h\"")
		} else {
			g.head.add(at.Line, "#include \"freertos/FreeRTOS.h\"")
			g.head.add(at.Line, "#include \"freertos/task.h\"")
		}
	}
	if g.board.Family == "RP2040" || g.board.Family == "RP2350" {
		if priority > 0 {
			g.later(at, "priority on an RP2040/RP2350 core block", phaseLater)
		}
		// multicore_launch_core1 wants void(*)() but the generated entry is void(void*) (FreeRTOS' signature) - a
		// no-arg thunk adapts it, so the RP2040/RP2350 core-1 launch actually compiles.
		g.fns.add(at.Line, "static void %s_launch() { %s(nullptr); }", fn, fn)
		g.setup.add(at.Line, "    multicore_launch_core1(%s_launch);", fn)
		return
	}
	// A failed task creation (out of heap for the stack) must not look successful: the block would silently never
	// run. Fail like a runtime error, at setup, naming the core. RP-family's launch can't fail observably.
	g.setup.add(at.Line, "    if (xTaskCreatePinnedToCore(%s, %q, 8192, nullptr, %d, nullptr, %d) != pdPASS) mujard::fail(%q, %d);",
		fn, fn, 1+priority, core, fmt.Sprintf("A block on core %d could not start: no memory for its task", core), at.Line)
}

// coreBlock routes an 'on core N' every/after/when block to a secondary core: on ESP32-family a FreeRTOS task
// pinned to the core, on RP2040/RP2350 multicore_launch_core1 (they have a second core, no more). In a host test
// the block keeps running on the main scheduler - the fake can't prove real parallelism.
func (g *gen) coreBlock(d *EventBlock) {
	name := "mujard_" + g.id("core")
	if d.As != "" {
		name = "mujard_block_" + d.As
		g.blockNames[d.As] = name
	}
	var gate string
	switch d.Kind {
	case "every":
		g.globals.add(d.At.Line, "static mujard::Every %s(%s);", name, g.period(d.Expr))
		gate = name + ".due()"
	case "after":
		g.globals.add(d.At.Line, "static mujard::After %s;", name)
		g.setup.add(d.At.Line, "    %s.arm(%s);", name, g.expr(d.Expr))
		gate = name + ".fire()"
	case "when":
		g.globals.add(d.At.Line, "static mujard::When %s;", name)
		gate = fmt.Sprintf("%s.edge(%s)", name, g.expr(d.Expr))
	}
	if g.testMode {
		g.loop.add(d.At.Line, "    if (%s) {", gate)
		g.block(&g.loop, d.Body, 2)
		g.loop.add(0, "    }")
		return
	}
	fn := g.id("coreFn")
	g.fns.add(d.At.Line, "static void %s(void*) {", fn)
	// Before the loop runs any user code: let the peer core lock this one out during a flash commit (RP).
	g.fns.add(1, "    mujard::coreInitFlashSafe();")
	g.fns.add(1, "for (;;) {")
	g.fns.add(2, "    if (mujard::stopping()) break;  // halted: secondary cores stop too (spec 6.7)")
	g.fns.add(2, "if (%s) {", gate)
	g.block(&g.fns, d.Body, 3)
	g.fns.add(2, "}")
	g.fns.add(1, "}")
	g.fns.add(0, "}")
	g.coreLaunch(d.At, fn, d.Priority, d.Core)
}

func (g *gen) topEvent(d *EventBlock) {
	if d.Core > 0 {
		g.coreBlock(d)
		return
	}
	line := d.At.Line
	switch d.Kind {
	case "start":
		g.setup.add(line, "    {  // at start")
		g.block(&g.setup, d.Body, 2)
		g.setup.add(0, "    }")
	case "time":
		// at 07:30: once a day, in that minute of the local time (mujard_clock.h); the block remembers the day it ran
		day := g.id("at_day")
		g.globals.add(line, "static int32_t %s = -1;", day)
		g.loop.add(line, "    if (mujard::clock_rt.at(%d, %d, %s)) {", d.Hour, d.Min, day)
		g.block(&g.loop, d.Body, 2)
		g.loop.add(0, "    }")
	case "every":
		name := g.id("every")
		if d.As != "" {
			name = "mujard_block_" + d.As
			g.blockNames[d.As] = name
		}
		g.globals.add(line, "static mujard::Every %s(%s);", name, g.period(d.Expr))
		g.loop.add(line, "    if (%s.due()) {", name)
		g.block(&g.loop, d.Body, 2)
		g.loop.add(0, "    }")
	case "after":
		name := g.id("after")
		if d.As != "" {
			name = "mujard_block_" + d.As
			g.blockNames[d.As] = name
		}
		g.globals.add(line, "static mujard::After %s;", name)
		g.setup.add(line, "    %s.arm(%s);", name, g.expr(d.Expr))
		g.loop.add(line, "    if (%s.fire()) {", name)
		g.block(&g.loop, d.Body, 2)
		g.loop.add(0, "    }")
	case "when":
		name := g.id("when")
		if d.As != "" {
			name = "mujard_block_" + d.As
			g.blockNames[d.As] = name
		}
		g.globals.add(line, "static mujard::When %s;", name)
		g.loop.add(line, "    if (%s.edge(%s)) {", name, g.expr(d.Expr))
		g.block(&g.loop, d.Body, 2)
		g.loop.add(0, "    }")
	case "on":
		g.onEvent(d)
	}
}

// period is an every block's period as a mujard::Time (a frequency becomes its period).
func (g *gen) period(n Node) string {
	t := g.typeOf(n)
	if t.K == TQty && t.Dim == "frequency" {
		return fmt.Sprintf("mujard::period_us(%s)", g.expr(n))
	}
	return g.expr(n)
}

func (g *gen) onEvent(d *EventBlock) {
	line := d.At.Line
	ref := d.Expr
	var arg Node
	if c, ok := ref.(*Call); ok {
		ref = c.Fn
		if len(c.Args) > 0 {
			arg = c.Args[0].Value
		}
	}
	mem, ok := ref.(*Member)
	if !ok {
		g.later(d.At, "this event", phaseLater)
		return
	}
	xt := g.typeOf(mem.X)
	obj := g.expr(mem.X)
	switch {
	case xt.K == TDevice && xt.Name == "input":
		if d.Fast {
			g.fastHandler(d, mem)
			return
		}
		if mem.Name == "double_clicked" {
			// then a plain click waits out the 400 ms double-click window (spec 6.5)
			g.setup.add(line, "    %s.wantDouble();", obj)
		}
		cond := fmt.Sprintf("%s.%s()", obj, mem.Name)
		if mem.Name == "held" {
			cond = fmt.Sprintf("%s.held(%s)", obj, g.expr(arg))
		}
		g.loop.add(line, "    if (%s) {", cond)
		g.block(&g.loop, d.Body, 2)
		g.loop.add(0, "    }")
	case xt.K == TObject && xt.Name == "serial" && mem.Name == "line":
		v := "mujard_line"
		if d.As != "" {
			v = cname(d.As)
		}
		g.loop.add(line, "    {")
		g.loop.add(line, "        mujard::Text %s;", v)
		g.loop.add(line, "        if (mujard::serial_line(%s)) {", v)
		g.block(&g.loop, d.Body, 3)
		g.loop.add(0, "        }")
		g.loop.add(0, "    }")
	case xt.K == TObject && xt.Name == "uart" && mem.Name == "line":
		v := "mujard_line"
		if d.As != "" {
			v = cname(d.As)
		}
		g.loop.add(line, "    {")
		g.loop.add(line, "        mujard::Text %s;", v)
		g.loop.add(line, "        if (%s.line(%s)) {", obj, v)
		g.block(&g.loop, d.Body, 3)
		g.loop.add(0, "        }")
		g.loop.add(0, "    }")
	case xt.K == TObject && xt.Name == "wifi" && (mem.Name == "joined" || mem.Name == "lost"):
		g.loop.add(line, "    if (mujard::wifi_link.%s()) {", mem.Name)
		g.block(&g.loop, d.Body, 2)
		g.loop.add(0, "    }")
	case xt.K == TObject && xt.Name == "ethernet" && (mem.Name == "joined" || mem.Name == "lost"):
		if !g.board.HasEthernet() {
			g.later(d.At, "Ethernet on this board", phaseLater)
			return
		}
		g.loop.add(line, "    if (mujard::eth_rt.%s()) {", mem.Name)
		g.block(&g.loop, d.Body, 2)
		g.loop.add(0, "    }")
	case xt.K == TObject && xt.Name == "timer" && mem.Name == "tick":
		if !timerBackend(g.board) {
			g.later(d.At, "Timer on this board", phaseLater)
			return
		}
		if d.Fast {
			g.timerIsr(d, mem)
			return
		}
		g.loop.add(line, "    if (%s.due()) {", obj)
		g.block(&g.loop, d.Body, 1)
		g.loop.add(0, "    }")
	case xt.K == TObject && xt.Name == "ble" && mem.Name == "line":
		if !bleBackend(g.board) {
			g.later(d.At, "BLE on this board", phaseLater)
			return
		}
		v := "mujard_ble_line"
		if d.As != "" {
			v = cname(d.As)
		}
		g.loop.add(line, "    {")
		g.loop.add(line, "        mujard::Text %s;", v)
		g.loop.add(line, "        if (mujard::ble_rt.line(%s)) {", v)
		g.block(&g.loop, d.Body, 3)
		g.loop.add(0, "        }")
		g.loop.add(0, "    }")
	case xt.K == TObject && xt.Name == "mqtt" && mem.Name == "message":
		v := "mujard_message"
		if d.As != "" {
			v = cname(d.As)
		}
		g.loop.add(line, "    {")
		g.loop.add(line, "        mujard::MqttMessage %s;", v)
		g.loop.add(line, "        if (mujard::mqtt_rt.message(%s)) {", v)
		g.block(&g.loop, d.Body, 3)
		g.loop.add(0, "        }")
		g.loop.add(0, "    }")
	case xt.K == TObject && xt.Name == "can" && mem.Name == "message":
		v := "mujard_can_message"
		if d.As != "" {
			v = cname(d.As)
		}
		g.loop.add(line, "    {")
		g.loop.add(line, "        mujard::CanMessage %s;", v)
		g.loop.add(line, "        if (%s.message(%s)) {", obj, v)
		g.block(&g.loop, d.Body, 3)
		g.loop.add(0, "        }")
		g.loop.add(0, "    }")
	default:
		g.later(d.At, "this event", phaseLater)
	}
}

// fastHandler: 'on pin.changed fast' runs as an interrupt.
func (g *gen) fastHandler(d *EventBlock, mem *Member) {
	name := g.id("isr")
	id, ok := mem.X.(*Ident)
	if !ok {
		g.later(d.At, "fast on this event", phaseLater)
		return
	}
	dev := g.devices[id.Name]
	mode := map[string]string{"changed": "CHANGE", "rising": "RISING", "falling": "FALLING"}[mem.Name]
	if mem.Name == "pressed" || mem.Name == "released" {
		// pressed / released follow the active level (spec 7): active LOW means pressed is the falling edge
		activeLow := dev != nil && dev.Kind == "pin" && hasOpt(dev.Options, "pullup") != hasOpt(dev.Options, "inverted")
		if (mem.Name == "pressed") == activeLow {
			mode = "FALLING"
		} else {
			mode = "RISING"
		}
	}
	if mode == "" {
		g.later(d.At, "fast "+mem.Name, phaseLater)
		return
	}
	g.fns.add(d.At.Line, "#if defined(ARDUINO_ARCH_ESP32) || defined(ARDUINO_ARCH_ESP8266)")
	g.fns.add(0, "IRAM_ATTR")
	g.fns.add(0, "#endif")
	g.fns.add(d.At.Line, "static void %s() {", name)
	g.fns.add(0, "    if (mujard::stopping()) return;  // halted: fast handlers stop too (spec 6.7)")
	g.fns.add(0, "    mujard::IsrScope mujard_isr;")
	g.inFast = true
	g.block(&g.fns, d.Body, 1)
	g.inFast = false
	g.fns.add(0, "}")
	g.setup.add(d.At.Line, "    attachInterrupt(digitalPinToInterrupt(%s.pin()), %s, %s);", g.expr(mem.X), name, mode)
}

func (g *gen) task(d *TaskDecl) {
	obj := "mujard_task_" + d.Name
	ctx := &taskCtx{seen: map[string]bool{}}
	var body section
	g.inTask = ctx
	g.block(&body, d.Body, 2)
	g.inTask = nil
	g.fns.add(d.At.Line, "struct mujard_Task_%s : mujard::Task {", d.Name)
	for _, fl := range ctx.fields {
		g.fns.add(d.At.Line, "    %s", fl)
	}
	g.fns.add(0, "    void run() {")
	g.fns.add(0, "        MUJARD_TASK_BEGIN();")
	g.fns = append(g.fns, body...)
	g.fns.add(0, "        MUJARD_TASK_END();")
	g.fns.add(0, "    }")
	g.fns.add(0, "} %s;", obj)
	if d.Core == 0 {
		g.loop.add(d.At.Line, "    %s.run();", obj)
		return
	}
	if g.testMode { // a host test keeps the task on the main scheduler - the fake can't prove real parallelism
		g.loop.add(d.At.Line, "    %s.run();", obj)
		return
	}
	fn := g.id("coreTask")
	g.fns.add(d.At.Line, "static void %s(void*) {", fn)
	// Before the loop runs any user code: let the peer core lock this one out during a flash commit (RP).
	g.fns.add(1, "    mujard::coreInitFlashSafe();")
	g.fns.add(1, "for (;;) {")
	g.fns.add(2, "    if (mujard::stopping()) break;  // halted: secondary cores stop too (spec 6.7)")
	g.fns.add(2, "%s.run();", obj)
	g.fns.add(1, "}")
	g.fns.add(0, "}")
	g.coreLaunch(d.At, fn, d.Priority, d.Core)
}

func (g *gen) block(out *section, b *Block, depth int) {
	if b == nil {
		return
	}
	for _, s := range b.Stmts {
		g.stmt(out, s, depth)
	}
}

func ind(depth int) string { return strings.Repeat("    ", depth) }

// declare emits a local variable; inside a task it becomes a field of the task (it must survive a wait).
func (g *gen) declare(out *section, line, depth int, ctype, name, value string) {
	if g.inTask != nil {
		if !g.inTask.seen[name] {
			g.inTask.seen[name] = true
			g.inTask.fields = append(g.inTask.fields, fmt.Sprintf("%s %s;", ctype, name))
		}
		if value != "" {
			out.add(line, "%s%s = %s;", ind(depth), name, value)
		}
		return
	}
	if value == "" {
		out.add(line, "%s%s %s;", ind(depth), ctype, name)
		return
	}
	out.add(line, "%s%s %s = %s;", ind(depth), ctype, name, value)
}

func (g *gen) stmt(out *section, n Node, depth int) {
	p := ind(depth)
	line := n.Start().Line
	g.hoistWaits(out, n, depth) // a call that suspends inside a task starts, and waits, before this statement
	switch s := n.(type) {
	case *VarDecl:
		t := g.typeOf(s.Value)
		if s.Type != nil {
			if dt := g.declTypeFromTypeExpr(s.Type); dt != nil {
				t = dt
				g.setWant(s.Value, t)
			}
		}
		ct := g.ctype(t, declCap(s.Type))
		if lit, ok := s.Value.(*ListLit); ok {
			g.declare(out, line, depth, ct, cname(s.Name), "")
			if g.inTask != nil {
				out.add(line, "%s%s.clear();", p, cname(s.Name))
			}
			for _, it := range lit.Items {
				g.setWant(it, elemOf(t))
				out.add(line, "%s%s.add(%s);", p, cname(s.Name), g.expr(it))
			}
			return
		}
		if ct == "auto" && g.inTask != nil && !isNone(s.Value) {
			// a library value in a task: the field takes the initializer's own type (spec 4.4)
			val := g.expr(s.Value)
			g.declare(out, line, depth, "decltype("+val+")", cname(s.Name), "")
			out.add(line, "%s%s = %s;", p, cname(s.Name), val)
			return
		}
		if ct == "auto" && g.inTask != nil {
			g.later(s.At, "a variable of unknown type in a task", phaseLibs)
			return
		}
		val := g.expr(s.Value)
		if isNone(s.Value) {
			val = ct + "()"
		}
		g.declare(out, line, depth, ct, cname(s.Name), val)
	case *AssignStmt:
		g.assign(out, s, depth)
	case *ExprStmt:
		if d := g.advanceTime(s.X); d != "" {
			g.inTask.waits++ // advance(t) in a test: the whole program runs on for t (spec 12)
			out.add(line, "%sMUJARD_WAIT(%s, %d);", p, d, g.inTask.waits)
			return
		}
		out.add(line, "%s%s;", p, g.expr(s.X))
	case *SimpleStmt:
		switch s.Kind {
		case "print":
			out.add(line, "%smujard::print(%s);", p, g.text(s.Value))
		case "wait":
			if g.inTask == nil {
				g.later(s.At, "wait outside a task", phaseLater)
				return
			}
			g.inTask.waits++
			if s.Until {
				out.add(line, "%sMUJARD_WAIT_UNTIL(%s, %d);", p, g.expr(s.Value), g.inTask.waits)
			} else {
				out.add(line, "%sMUJARD_WAIT(%s, %d);", p, g.expr(s.Value), g.inTask.waits)
			}
		case "pass":
			out.add(line, "%s;", p)
		case "return":
			if s.Value != nil {
				g.setWant(s.Value, g.retType)
			}
			if g.inTask != nil && g.inTask.fnResult {
				if s.Value != nil {
					out.add(line, "%s{ result_ = %s; running_ = false; return; }", p, g.expr(s.Value))
				} else {
					out.add(line, "%s{ running_ = false; return; }", p)
				}
			} else if s.Value != nil {
				out.add(line, "%sreturn %s;", p, g.expr(s.Value))
			} else {
				out.add(line, "%sreturn;", p)
			}
		default:
			out.add(line, "%s%s;", p, s.Kind)
		}
	case *ExpectStmt:
		if g.inTestBody {
			// in a test, a failed check ends the test with its message; the program carries on unharmed
			out.add(line, "%sif (!(%s)) { mujard::test_fail(%s, %d); return; }", p, g.expr(s.Cond), g.expectText(s), line)
			return
		}
		call := "fail"
		if lit, ok := s.HaltArg().(*Literal); ok && lit.Value == "false" {
			call = "expect_warn"
		}
		out.add(line, "%sif (!(%s)) mujard::%s(%s, %d);", p, g.expr(s.Cond), call, g.expectText(s), line)
	case *IfStmt:
		for i, c := range s.Conds {
			kw := "if"
			if i > 0 {
				kw = "} else if"
			}
			out.add(c.Start().Line, "%s%s (%s) {", p, kw, g.expr(c))
			g.block(out, s.Thens[i], depth+1)
		}
		if s.Else != nil {
			out.add(s.Else.At.Line, "%s} else {", p)
			g.block(out, s.Else, depth+1)
		}
		out.add(0, "%s}", p)
	case *WhileStmt:
		out.add(line, "%swhile (%s) {", p, g.expr(s.Cond))
		g.block(out, s.Body, depth+1)
		out.add(0, "%s}", p)
	case *ForStmt:
		g.forStmt(out, s, depth)
	case *MatchStmt:
		g.match(out, s, depth)
	case *EventBlock:
		g.nestedEvent(out, s, depth)
	case *CppBlock:
		g.rawCpp = true
		for i, l := range strings.Split(strings.TrimRight(s.Code, "\n"), "\n") {
			out.add(line+1+i, "%s%s", p, l)
		}
	}
}

func (g *gen) assign(out *section, s *AssignStmt, depth int) {
	p := ind(depth)
	line := s.At.Line
	key := atomicKey(s.Target)
	atomic := !g.inFast && key != "" && (g.atomics[key] || (g.curMod != "" && !strings.Contains(key, ".") && g.atomics[g.curMod+"."+key]))
	tt := g.typeOf(s.Target)
	if id, ok := s.Target.(*Ident); ok && g.decl[id] != nil {
		tt = g.decl[id]
	}
	if s.Op == "=" {
		g.setWant(s.Value, tt)
	}
	if atomic {
		// wider than a byte, and a fast handler writes it too: the update runs with interrupts off (6.5)
		g.inAtomic = true
		out.add(line, "%snoInterrupts();", p)
	}
	target := g.expr(s.Target)
	if id, ok := s.Target.(*Ident); ok {
		target = cname(id.Name) // the variable itself, not its value
	}
	if isNone(s.Value) {
		out.add(line, "%s%s = %s();", p, target, g.ctype(tt, 0))
	} else if s.Op == "+=" && tt.K == TText {
		g.setWant(s.Value, tt)
		out.add(line, "%s%s.add(%s);", p, target, g.text(s.Value))
	} else if op := intAssignOps[s.Op]; op != "" && nonOpt(tt).K == TInt && nonOpt(g.typeOf(s.Value)).K != TFloat {
		// x += y on ints wraps and x /= 0 is a runtime error, like the operators (the target is evaluated once)
		if s.Op == "/=" {
			out.add(line, "%smujard::%s(%s, %s, %d);", p, op, target, g.expr(s.Value), line)
		} else {
			out.add(line, "%smujard::%s(%s, %s);", p, op, target, g.expr(s.Value))
		}
	} else {
		out.add(line, "%s%s %s %s;", p, target, s.Op, g.expr(s.Value))
	}
	if atomic {
		g.inAtomic = false
		out.add(line, "%sinterrupts();", p)
	}
}

func (g *gen) forStmt(out *section, s *ForStmt, depth int) {
	p := ind(depth)
	line := s.At.Line
	v := cname(s.Var)
	it := g.typeOf(s.Iter)
	if r, ok := s.Iter.(*RangeExpr); ok {
		ct := g.ctype(elemOf(it), 0)
		end, step := g.id("end"), g.id("step")
		incl := "0"
		if r.Inclusive {
			incl = step
		}
		if g.inTask != nil {
			g.declare(out, line, depth, ct, v, "")
			g.declare(out, line, depth, ct, end, "")
			g.declare(out, line, depth, ct, step, "")
			out.add(line, "%sfor (%s = %s, %s = %s, %s = (%s <= %s ? 1 : -1); %s != %s + %s; %s += %s) {", p, v, g.expr(r.Lo), end, g.expr(r.Hi),
				step, v, end, v, end, incl, v, step)
		} else {
			out.add(line, "%sfor (%s %s = %s, %s = %s, %s = (%s <= %s ? 1 : -1); %s != %s + %s; %s += %s) {", p, ct, v, g.expr(r.Lo), end, g.expr(r.Hi),
				step, v, end, v, end, incl, v, step)
		}
		g.block(out, s.Body, depth+1)
		out.add(0, "%s}", p)
		return
	}
	k := g.id("i")
	coll := g.expr(s.Iter)
	switch {
	case it.K == TList:
		if g.inTask != nil {
			// The index and the loop variable are task fields: they must survive a wait inside the body.
			g.declare(out, line, depth, "int32_t", k, "")
			g.declare(out, line, depth, g.ctype(elemOf(it), 0), v, "")
			out.add(line, "%sfor (%s = 0; %s < %s.n; %s++) {", p, k, k, coll, k)
			out.add(line, "%s    %s = %s.items[%s];", p, v, coll, k)
			break
		}
		out.add(line, "%sfor (int32_t %s = 0; %s < %s.n; %s++) {", p, k, k, coll, k)
		out.add(line, "%s    %s %s = %s.items[%s];", p, g.ctype(elemOf(it), 0), v, coll, k)
	case it.K == TDevice && it.Name == "group":
		n := g.groupSize(s.Iter)
		if g.inTask != nil {
			// The device can't be a field (pin objects don't copy), so the variable name is an alias for the
			// group element while the body is generated.
			g.declare(out, line, depth, "int32_t", k, "")
			out.add(line, "%sfor (%s = 0; %s < %d; %s++) {", p, k, k, n, k)
			if g.inTask.alias == nil {
				g.inTask.alias = map[string]string{}
			}
			g.inTask.alias[s.Var] = coll + "[" + k + "]"
			g.block(out, s.Body, depth+1)
			delete(g.inTask.alias, s.Var)
			out.add(0, "%s}", p)
			return
		}
		out.add(line, "%sfor (int32_t %s = 0; %s < %d; %s++) {", p, k, k, n, k)
		out.add(line, "%s    auto& %s = %s[%s];", p, v, coll, k)
	default:
		g.later(s.At, "this kind of for loop", phaseLater)
		return
	}
	g.block(out, s.Body, depth+1)
	out.add(0, "%s}", p)
}

func (g *gen) groupSize(n Node) int {
	if m, ok := n.(*Member); ok {
		if id, ok := m.X.(*Ident); ok && g.modGen[id.Name] != nil {
			board := g.board
			if mc := g.mods[id.Name]; mc != nil && mc.board != nil {
				board = mc.board
			}
			return groupSizeOf(g.modGen[id.Name].devices[m.Name], board)
		}
	}
	if id, ok := n.(*Ident); ok {
		return groupSizeOf(g.devices[id.Name], g.board)
	}
	return 0
}

func groupSizeOf(d *DeviceDecl, board *Board) int {
	if d == nil {
		return 0
	}
	if d.Kind == "pinlist" {
		return len(d.Pins)
	}
	a, _ := board.ResolvePin(d.Pins[0].Text)
	b, _ := board.ResolvePin(d.Pins[1].Text)
	if d.ToEnd {
		b++
	}
	return b - a
}

func (g *gen) match(out *section, s *MatchStmt, depth int) {
	p := ind(depth)
	vt := nonOpt(g.typeOf(s.Value))
	subject := g.expr(s.Value)
	first := true
	for _, c := range s.Cases {
		var cond string
		if c.IsElse {
			cond = ""
		} else {
			var parts []string
			for _, pt := range c.Patterns {
				if r, ok := pt.(*RangeExpr); ok {
					op := "<"
					if r.Inclusive {
						op = "<="
					}
					parts = append(parts, fmt.Sprintf("(%s >= %s && %s %s %s)", subject, g.expr(r.Lo), subject, op, g.expr(r.Hi)))
					continue
				}
				if vt.K == TText {
					parts = append(parts, fmt.Sprintf("%s.eq(%s)", subject, g.expr(pt)))
				} else {
					parts = append(parts, fmt.Sprintf("%s == %s", subject, g.expr(pt)))
				}
			}
			cond = strings.Join(parts, " || ")
		}
		switch {
		case first && cond != "":
			out.add(c.At.Line, "%sif (%s) {", p, cond)
		case cond != "":
			out.add(c.At.Line, "%s} else if (%s) {", p, cond)
		case first:
			out.add(c.At.Line, "%s{", p)
		default:
			out.add(c.At.Line, "%s} else {", p)
		}
		first = false
		g.block(out, c.Body, depth+1)
	}
	if !first {
		out.add(0, "%s}", p)
	}
}

func (g *gen) nestedEvent(out *section, s *EventBlock, depth int) {
	p := ind(depth)
	line := s.At.Line
	switch s.Kind {
	case "when":
		name := g.id("when")
		g.globals.add(line, "static mujard::When %s;", name)
		out.add(line, "%sif (%s.edge(%s)) {", p, name, g.expr(s.Expr))
		g.block(out, s.Body, depth+1)
		out.add(0, "%s}", p)
	case "after":
		// The block runs later, from loop(), when the code around it has finished (C319 keeps its locals out).
		name := g.id("after")
		g.globals.add(line, "static mujard::After %s;", name)
		out.add(line, "%s%s.arm(%s);", p, name, g.expr(s.Expr))
		saveTask := g.inTask
		g.inTask = nil
		g.deferred = append(g.deferred, outLine{text: fmt.Sprintf("    if (%s.fire()) {", name), src: line})
		g.block(&g.deferred, s.Body, 2)
		g.deferred = append(g.deferred, outLine{text: "    }"})
		g.inTask = saveTask
	}
}

// ---- expressions

// expr is the C++ for an expression (its value; a narrowed optional is unwrapped).
func (g *gen) expr(n Node) string {
	switch e := n.(type) {
	case nil:
		return ""
	case *Ident:
		return g.ident(e)
	case *NumberLit:
		return g.number(e)
	case *StringLit:
		return g.text(e)
	case *Literal:
		switch e.Value {
		case "none":
			return "{}"
		}
		return e.Value
	case *ListLit:
		t := g.wantType[e]
		if t == nil {
			t = g.typeOf(e)
		} else {
			t = nonOpt(t)
		}
		el := elemOf(t)
		if t.K != TList || el == nil || el == tAny {
			return g.later(e.At, "a list written inside an expression", phaseLater)
		}
		var b strings.Builder
		b.WriteString(g.ctype(t, 0))
		b.WriteString("({ ")
		for i, it := range e.Items {
			if i > 0 {
				b.WriteString(", ")
			}
			g.setWant(it, el)
			b.WriteString(g.expr(it))
		}
		b.WriteString(" })")
		return b.String()
	case *PinLit:
		return g.pinNumber(e.Ref)
	case *Unary:
		switch e.Op {
		case "not":
			return "!(" + g.expr(e.X) + ")"
		case "-":
			if nonOpt(g.typeOf(e.X)).K == TInt {
				return "mujard::ineg(" + g.expr(e.X) + ")" // -(-2147483648) wraps instead of being undefined
			}
		}
		return e.Op + "(" + g.expr(e.X) + ")"
	case *Binary:
		return g.binary(e)
	case *RangeExpr:
		return g.later(e.At, "a range used as a value", phaseLater)
	case *Conversion:
		// values stay in base units; 'as' only changes how they print - except on a plain
		// number, where it makes a quantity (spec 5): the base-unit value is n * factor.
		xt := nonOpt(g.typeOf(e.X))
		if xt.K == TObject && xt.Name == "json" {
			switch dimOfUnit(e.Unit) {
			case "time":
				return fmt.Sprintf("((mujard::Time)((float)(%s) * %s))", g.expr(e.X), cfloat(unitFactor(e.Unit)))
			case "frequency":
				return fmt.Sprintf("((mujard::Freq)((float)(%s) * %dLL))", g.expr(e.X), int64(unitFactor(e.Unit)))
			}
			return fmt.Sprintf("((float)(%s) * %s)", g.expr(e.X), cfloat(unitFactor(e.Unit)))
		}
		if isNumeric(xt) {
			switch dimOfUnit(e.Unit) {
			case "time":
				return fmt.Sprintf("((mujard::Time)((%s) * %s))", g.expr(e.X), cfloat(unitFactor(e.Unit)))
			case "frequency":
				return fmt.Sprintf("((mujard::Freq)((%s) * %dLL))", g.expr(e.X), int64(unitFactor(e.Unit)))
			}
			return fmt.Sprintf("((%s) * %s)", g.expr(e.X), cfloat(unitFactor(e.Unit)))
		}
		return g.expr(e.X)
	case *CondExpr:
		if nonOpt(g.typeOf(e)).K == TText {
			ct := g.textCtor(e)
			return fmt.Sprintf("((%s) ? (%s(%s)) : (%s(%s)))", g.expr(e.Cond), ct, g.expr(e.Then), ct, g.expr(e.Else))
		}
		return fmt.Sprintf("((%s) ? (%s) : (%s))", g.expr(e.Cond), g.expr(e.Then), g.expr(e.Else))
	case *Call:
		return g.call(e)
	case *Member:
		return g.member(e)
	case *Index:
		xt := nonOpt(g.typeOf(e.X))
		if xt.K == TList {
			return fmt.Sprintf("%s.at(%s, %d)", g.expr(e.X), g.expr(e.Ix), e.At.Line)
		}
		if xt.K == TObject && xt.Name == "json" {
			if it := nonOpt(g.typeOf(e.Ix)); it.K == TText {
				return fmt.Sprintf("%s.get(%s)", g.expr(e.X), g.cstr(e.Ix))
			}
			return fmt.Sprintf("%s.at((int32_t)(%s))", g.expr(e.X), g.expr(e.Ix))
		}
		if xt.K == TText {
			return fmt.Sprintf("%s.at(%s, %d)", g.expr(e.X), g.expr(e.Ix), e.At.Line)
		}
		return fmt.Sprintf("%s[%s]", g.expr(e.X), g.expr(e.Ix))
	}
	return "0"
}

func (g *gen) ident(e *Ident) string {
	if g.inTask != nil {
		if a, ok := g.inTask.alias[e.Name]; ok {
			return a
		}
	}
	if obj, ok := g.blockNames[e.Name]; ok {
		return obj
	}
	if e.Name == "board" {
		return "mujard_board"
	}
	if gl, ok := globals[e.Name]; ok && gl.value != nil {
		if _, local := g.devices[e.Name]; !local {
			if e.Name == "forever" {
				return "((mujard::Time)0)"
			}
			if t := g.typeOf(e); isBuiltinEnum(t) {
				return g.builtinEnum(t.Name, e.Name, e.At)
			}
		}
	}
	if g.isLibGlobal(e) {
		// A library's own name (OUTPUT, HIGH, u8g2_font_â€¦) is what the library's C++ calls it: never renamed like
		// the program's names, which only stay clear of Arduino's.
		return g.libName(e.Name)
	}
	name := cname(e.Name)
	if g.modQual && g.curModScope != nil {
		if _, ok := g.curModScope.syms[e.Name]; ok {
			name = cname(g.curMod) + "::" + name // a module expression emitted outside its namespace
		}
	}
	declared := g.decl[e]
	used := g.typeOf(e)
	if used != nil && used.K == TEnum && declared == nil {
		return qname(used.Name) + "::" + name // a bare enum member (spec 4.6)
	}
	if declared != nil && declared.Opt && !used.Opt {
		return name + ".v" // checked not to be none here
	}
	if g.atomicRead(e.Name) {
		return "mujard::atomic_read(" + name + ")"
	}
	return name
}

// atomicRead reports whether a value-read of this name must go through mujard::atomic_read
// (a variable wider than a byte, changed by a fast handler - spec 6.5).
func (g *gen) atomicRead(name string) bool {
	if g.inFast || g.inAtomic {
		return false
	}
	key := name
	if g.curMod != "" {
		key = g.curMod + "." + name
	}
	return g.atomics[key] || g.atomics[name]
}

func (g *gen) number(e *NumberLit) string {
	v := e.Float
	if !e.IsFloat {
		v = float64(e.Int)
	}
	if e.Unit == "" {
		if e.IsFloat {
			return cfloat(v)
		}
		if e.Int > 32767 {
			return strconv.FormatUint(e.Int, 10) + "L"
		}
		return strconv.FormatUint(e.Int, 10)
	}
	base := v * unitFactor(e.Unit)
	switch dimOfUnit(e.Unit) {
	case "time":
		return fmt.Sprintf("((mujard::Time)%dLL)", int64(math.Round(base)))
	case "frequency":
		return fmt.Sprintf("((mujard::Freq)%dLL)", int64(math.Round(base)))
	}
	return cfloat(base)
}

func isTime(t *Type) bool { return t.K == TQty && t.Dim == "time" }

func (g *gen) binary(e *Binary) string {
	l, r := nonOpt(g.typeOf(e.X)), nonOpt(g.typeOf(e.Y))
	switch e.Op {
	case "and":
		return "(" + g.expr(e.X) + " && " + g.expr(e.Y) + ")"
	case "or":
		return "(" + g.expr(e.X) + " || " + g.expr(e.Y) + ")"
	case "in":
		return g.membership(e)
	case "==", "!=":
		if isNone(e.Y) || isNone(e.X) {
			other := e.X
			if isNone(e.X) {
				other = e.Y
			}
			name := cname(nodeName(other))
			if id, ok := other.(*Ident); ok {
				name = cname(id.Name)
			} else {
				name = g.expr(other)
			}
			if e.Op == "==" {
				return "!" + name + ".has"
			}
			return name + ".has"
		}
		if l.K == TText || r.K == TText {
			s := fmt.Sprintf("%s.eq(%s)", g.textValue(e.X), g.expr(e.Y))
			if e.Op == "!=" {
				return "!" + s
			}
			return s
		}
	}
	a, b := g.expr(e.X), g.expr(e.Y)
	if l.K == TInt && r.K == TInt {
		// ints wrap on overflow and a zero divisor is a runtime error (spec 4, 6.8): never C++'s plain operators,
		// whose signed overflow and division by zero are undefined behaviour.
		switch e.Op {
		case "+", "-", "*":
			return fmt.Sprintf("mujard::%s(%s, %s)", intOps[e.Op], a, b)
		case "/", "mod", "<<", ">>":
			return fmt.Sprintf("mujard::%s(%s, %s, %d)", intOps[e.Op], a, b, e.At.Line)
		}
	}
	switch e.Op {
	case "*":
		switch {
		case isTime(l) && isTime(r):
			return g.later(e.At, "time times time", phaseLater)
		case isTime(l) && r.K == TQty && r.Dim == "frequency":
			return fmt.Sprintf("(mujard::secs(%s) * ((float)(%s) / 1000.0f))", a, b)
		case isTime(r) && l.K == TQty && l.Dim == "frequency":
			return fmt.Sprintf("(((float)(%s) / 1000.0f) * mujard::secs(%s))", a, b)
		case isTime(l) && r.K == TQty && r.Dim != "ratio":
			return fmt.Sprintf("(mujard::secs(%s) * (%s))", a, b)
		case isTime(r) && l.K == TQty && l.Dim != "ratio":
			return fmt.Sprintf("((%s) * mujard::secs(%s))", a, b)
		case isTime(l) || isTime(r):
			return fmt.Sprintf("((mujard::Time)((%s) * (%s)))", a, b)
		}
	case "/":
		switch {
		case isTime(l) && isTime(r):
			return fmt.Sprintf("((float)(%s) / (float)(%s))", a, b)
		case r.K == TQty && r.Dim == "time" && l.K == TQty:
			return fmt.Sprintf("((%s) / mujard::secs(%s))", a, b)
		case isTime(r):
			return fmt.Sprintf("((mujard::Freq)((%s) * 1000.0f / mujard::secs(%s)))", a, b)
		case l.K == TQty && l.Dim == "frequency" && r.K == TQty && r.Dim == "frequency":
			return fmt.Sprintf("((float)(%s) / (float)(%s))", a, b)
		case l.K == TQty && l.Dim == "length" && r.K == TQty && r.Dim == "speed":
			return fmt.Sprintf("mujard::from_secs((%s) / (%s))", a, b)
		case r.K == TQty && r.Dim == "frequency" && isNumeric(l):
			return fmt.Sprintf("mujard::from_secs((float)(%s) * 1000.0f / (float)(%s))", a, b)
		case isTime(l) && r.K == TInt:
			return fmt.Sprintf("mujard::tdiv(%s, %s, %d)", a, b, e.At.Line)
		case isTime(l):
			return fmt.Sprintf("mujard::tdivf(%s, %s, %d)", a, b, e.At.Line)
		}
	case "+":
		if l.K == TText || r.K == TText {
			return fmt.Sprintf("%s(%s).add(%s)", g.textCtor(e), g.textValue(e.X), g.textValue(e.Y))
		}
	case "mod":
		return "(" + a + " % " + b + ")"
	}
	return "(" + a + " " + e.Op + " " + b + ")"
}

var intAssignOps = map[string]string{"+=": "iadd_to", "-=": "isub_to", "*=": "imul_to", "/=": "idiv_to"}

// intOps are the runtime helpers for int arithmetic (mujard_rt.h).
var intOps = map[string]string{"+": "iadd", "-": "isub", "*": "imul", "/": "idiv", "mod": "imod", "<<": "ishl", ">>": "ishr"}

// cString is a C++ string literal for s. Every byte outside printable ASCII is a 3-digit octal escape (a hex escape
// would swallow following hex digits), and '?' is escaped so no trigraph can form.
func cString(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '"' || c == '\\' || c == '?':
			b.WriteByte('\\')
			b.WriteByte(c)
		case c == '\n':
			b.WriteString(`\n`)
		case c == '\t':
			b.WriteString(`\t`)
		case c < 0x20 || c >= 0x7f:
			fmt.Fprintf(&b, "\\%03o", c)
		default:
			b.WriteByte(c)
		}
	}
	b.WriteByte('"')
	return b.String()
}

// commentSafe is text for a // comment: control characters (a line break would end the comment) and a trailing
// backslash (it would continue the comment onto the next line) become '?'.
func commentSafe(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c < 0x20 || c == 0x7f || (c == '\\' && i == len(b)-1) {
			b[i] = '?'
		}
	}
	return string(b)
}

// membership is the C++ for `x in a..b` (bounds test) and `x in list` (List::contains).
func (g *gen) membership(e *Binary) string {
	if r, ok := e.Y.(*RangeExpr); ok {
		op := "<"
		if r.Inclusive {
			op = "<="
		}
		x := g.expr(e.X)
		return fmt.Sprintf("((%s) >= (%s) && (%s) %s (%s))", x, g.expr(r.Lo), x, op, g.expr(r.Hi))
	}
	if nonOpt(g.typeOf(e.Y)).K == TList {
		return fmt.Sprintf("(%s.contains(%s))", g.expr(e.Y), g.expr(e.X))
	}
	return g.later(e.At, "this kind of 'in'", phaseLater)
}

// text is the C++ for a value used as text (print, log, halt, ...): a string literal or a mujard::Text.
func (g *gen) text(n Node) string {
	if s, ok := n.(*StringLit); ok {
		allText := true
		for _, p := range s.Parts {
			if p.Expr != nil {
				allText = false
			}
		}
		if allText {
			lit := ""
			for _, p := range s.Parts {
				lit += p.Text
			}
			return cString(lit)
		}
		var b strings.Builder
		b.WriteString(g.textCtor(s) + "()")
		for _, p := range s.Parts {
			if p.Expr == nil {
				if p.Text != "" {
					fmt.Fprintf(&b, ".add(%s)", cString(p.Text))
				}
				continue
			}
			b.WriteString(g.addPart(p.Expr))
		}
		return b.String()
	}
	if nonOpt(g.typeOf(n)).K == TText && !g.typeOf(n).Opt {
		return g.expr(n)
	}
	return g.textCtor(n) + "()" + g.addPart(n)
}

func (g *gen) textValue(n Node) string {
	if _, ok := n.(*StringLit); ok {
		return g.textCtor(n) + "(" + g.text(n) + ")"
	}
	return g.expr(n)
}

// expectText is the expect() message: the given one, or the condition's own text (spec 6.8).
func (g *gen) expectText(s *ExpectStmt) string {
	if m := s.Message(); m != nil {
		return g.text(m)
	}
	return cString(sourceText(s.Cond))
}

// sourceText renders an expression roughly as it was written, for the default expect message.
func sourceText(n Node) string {
	switch e := n.(type) {
	case nil:
		return ""
	case *NumberLit:
		if e.Unit != "" {
			return e.Text + " " + e.Unit
		}
		return e.Text
	case *StringLit:
		var b strings.Builder
		b.WriteByte('"')
		for _, p := range e.Parts {
			if p.Expr == nil {
				b.WriteString(p.Text)
			} else {
				b.WriteString("{" + sourceText(p.Expr) + "}")
			}
		}
		return b.String() + `"`
	case *Literal:
		return e.Value
	case *Ident:
		return e.Name
	case *Member:
		return sourceText(e.X) + "." + e.Name
	case *Index:
		return sourceText(e.X) + "[" + sourceText(e.Ix) + "]"
	case *Unary:
		if e.Op == "not" {
			return "not " + sourceText(e.X)
		}
		return e.Op + sourceText(e.X)
	case *Binary:
		return sourceText(e.X) + " " + e.Op + " " + sourceText(e.Y)
	case *RangeExpr:
		op := ".."
		if e.Inclusive {
			op = "..="
		}
		return sourceText(e.Lo) + op + sourceText(e.Hi)
	case *Conversion:
		return sourceText(e.X) + " as " + e.Unit
	case *CondExpr:
		return sourceText(e.Then) + " if " + sourceText(e.Cond) + " else " + sourceText(e.Else)
	case *Call:
		var b strings.Builder
		b.WriteString(sourceText(e.Fn) + "(")
		for i, a := range e.Args {
			if i > 0 {
				b.WriteString(", ")
			}
			if a.Name != "" {
				b.WriteString(a.Name + ": ")
			}
			b.WriteString(sourceText(a.Value))
		}
		return b.String() + ")"
	case *ListLit:
		var b strings.Builder
		b.WriteByte('[')
		for i, it := range e.Items {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(sourceText(it))
		}
		return b.String() + "]"
	case *PinLit:
		if e.Analog {
			return "analog " + e.Ref.Text
		}
		return "pin " + e.Ref.Text
	}
	return "expression"
}

// addPart is the .add...(x) call that puts one value into a text, in a readable form.
func (g *gen) addPart(n Node) string {
	t := g.typeOf(n)
	if c, ok := n.(*Conversion); ok {
		per := unitFactor(c.Unit)
		if dimOfUnit(c.Unit) == "time" || dimOfUnit(c.Unit) == "frequency" {
			return fmt.Sprintf(".add_unit((float)(%s), %s, %q)", g.expr(c.X), cfloat(per), c.Unit)
		}
		return fmt.Sprintf(".add_unit(%s, %s, %q)", g.expr(c.X), cfloat(per), c.Unit)
	}
	x := g.expr(n)
	if t.Opt {
		inner := nonOpt(t)
		switch {
		case isTime(inner):
			return fmt.Sprintf(".add_opt_time(%s)", x)
		case inner.K == TQty && inner.Dim == "frequency":
			return fmt.Sprintf(".add_opt_freq(%s)", x)
		case inner.K == TQty:
			return fmt.Sprintf(".add_opt_qty(%s, %s)", x, dimEnum[inner.Dim])
		}
		return fmt.Sprintf(".add_opt(%s)", x)
	}
	switch t.K {
	case TQty:
		if isTime(t) {
			return fmt.Sprintf(".add_time(%s)", x)
		}
		if t.Dim == "frequency" {
			return fmt.Sprintf(".add_freq(%s)", x)
		}
		return fmt.Sprintf(".add_qty(%s, %s)", x, dimEnum[t.Dim])
	case TObject:
		if t.Name == "json" {
			return fmt.Sprintf(".add((%s).to_text())", x)
		}
		return fmt.Sprintf(".add(%s)", x)
	case TInt:
		return fmt.Sprintf(".add((int32_t)(%s))", x)
	case TFloat:
		return fmt.Sprintf(".add((float)(%s))", x)
	case TBool:
		return fmt.Sprintf(".add((bool)(%s))", x)
	case TList:
		return fmt.Sprintf(".add_list(%s)", x)
	case TEnum:
		return fmt.Sprintf(".add(mujard_name_%s(%s))", flatName(t.Name), x)
	}
	return fmt.Sprintf(".add(%s)", x)
}

// ---- calls and members

// argMap matches a call's arguments to parameter names (positional first, then named).
func argMap(c *Call, names ...string) map[string]Node {
	out := map[string]Node{}
	pos := 0
	for _, a := range c.Args {
		if a.Name != "" {
			out[a.Name] = a.Value
			continue
		}
		if pos < len(names) {
			out[names[pos]] = a.Value
		}
		pos++
	}
	return out
}

func (g *gen) argOr(m map[string]Node, name, def string) string {
	if n, ok := m[name]; ok {
		return g.expr(n)
	}
	return def
}

// literalText requires a plain "text" (the runtime keeps a pointer to it: patterns, melodies).
func (g *gen) literalText(n Node) string {
	if s, ok := n.(*StringLit); ok {
		if t := g.text(s); strings.HasPrefix(t, "\"") {
			return t
		}
	}
	return g.later(n.Start(), "a pattern or melody that isn't written out as plain text", phaseLater)
}

func (g *gen) call(c *Call) string {
	if r, ok := g.waitResults[c]; ok {
		return r // a suspending call that was started and waited for before its statement
	}
	switch fn := c.Fn.(type) {
	case *Ident:
		if d, ok := g.fnDecls[fn.Name]; ok {
			if g.waitingFn(d) {
				return g.later(c.At, "a function that waits, here (it works as 'var x = â€¦' or 'x = â€¦' inside a task)", phaseLater)
			}
			return g.userCall(c, d)
		}
		if s, ok := g.structs[fn.Name]; ok {
			var fields []string
			m := map[string]Node{}
			for i, a := range c.Args {
				if a.Name != "" {
					m[a.Name] = a.Value
				} else if i < len(s.Fields) {
					m[s.Fields[i].Name] = a.Value
				}
			}
			for _, f := range s.Fields {
				ft := g.fieldType(f)
				if v, ok := m[f.Name]; ok {
					g.setWant(v, ft)
					fields = append(fields, g.expr(v))
				} else if f.Default != nil {
					g.setWant(f.Default, ft)
					fields = append(fields, g.expr(f.Default))
				} else {
					fields = append(fields, "{}")
				}
			}
			return fmt.Sprintf("%s{%s}", g.cppQual(cname(s.Name)), strings.Join(fields, ", "))
		}
		return g.builtinCall(c, fn)
	case *Member:
		return g.methodCall(c, fn)
	}
	var args []string
	for _, a := range c.Args {
		args = append(args, g.expr(a.Value))
	}
	return fmt.Sprintf("%s(%s)", g.expr(c.Fn), strings.Join(args, ", "))
}

func (g *gen) userCall(c *Call, d *FnDecl) string {
	var names []string
	for _, p := range d.Params {
		names = append(names, p.Name)
	}
	m := argMap(c, names...)
	var args []string
	for _, p := range d.Params {
		pt := g.declTypeFromTypeExpr(p.Type)
		if v, ok := m[p.Name]; ok {
			g.setWant(v, pt)
			args = append(args, g.expr(v))
		} else if p.Default != nil {
			g.setWant(p.Default, pt)
			args = append(args, g.expr(p.Default))
		}
	}
	return fmt.Sprintf("%s(%s)", g.cppQual(cname(d.Name)), strings.Join(args, ", "))
}

// castTo converts a float result back to the C++ type of t.
func (g *gen) castTo(t *Type, x string) string {
	switch {
	case isTime(t):
		return "((mujard::Time)(" + x + "))"
	case t.K == TQty && t.Dim == "frequency":
		return "((mujard::Freq)(" + x + "))"
	case t.K == TInt:
		return "((int32_t)(" + x + "))"
	}
	return x
}

// secretsHeader is the header the build server writes next to the sketch with the build's secrets as C macros
// (the platform's SECRETS_HEADER convention, and the mujard CLI's secretsHeaderName for local builds).
const secretsHeader = "mujard_secrets.h"

// secretMacro is the C macro a secret("name") compiles to: NAME, or NAME_SSID / NAME_PASSWORD for Wi-Fi credentials.
// The value is never in the generated code - only in the build's own copy of secretsHeader.
func (g *gen) secretMacro(c *Call, suffix string) string {
	name := ""
	if len(c.Args) == 1 {
		if s, ok := c.Args[0].Value.(*StringLit); ok {
			for _, p := range s.Parts {
				name += p.Text
			}
		}
	}
	macro := strings.ToUpper(name) + suffix
	if g.secrets == nil {
		g.secrets = map[string]int{}
	}
	if _, seen := g.secrets[macro]; !seen {
		g.secrets[macro] = c.At.Line
	}
	return macro
}

// cstr is a text argument as a const char*: a string literal as it is, anything else through a Text (valid until
// the end of the statement, so only for calls that copy it).
func (g *gen) cstr(n Node) string {
	if isSecretCall(n) {
		return g.secretMacro(n.(*Call), "")
	}
	s := g.text(n)
	if _, lit := n.(*StringLit); lit && !strings.HasPrefix(s, "mujard::Text") {
		return s
	}
	return "mujard::Text(" + s + ").s"
}

// bleBackend reports whether 'use mujard.ble' has a runtime on this board: the BLE radio on the ESP32 family
// (not S2), the Nano 33 BLE's nRF52840 and the Pico W's CYW43 (the core's BLEServiceUART is the same Nordic
// shape). Other boards get C320.
func bleBackend(b *Board) bool { return b.HasBluetooth() }

// bleCall is the BLE peripheral's members (mujard_ble.h): ble.start(name), send(text), read_line(), connected,
// 'on ble.line' and ble.pretend(line). Boards without a BLE backend refuse with C320.
func (g *gen) bleCall(c *Call, mem *Member, args func() string) string {
	if !bleBackend(g.board) {
		return g.later(c.At, "BLE on this board", phaseLater)
	}
	g.usesBle = true
	switch mem.Name {
	case "start":
		return fmt.Sprintf("mujard::ble_rt.start(%s)", g.cstr(c.Args[0].Value))
	case "send":
		return fmt.Sprintf("mujard::ble_rt.send(%s)", g.text(c.Args[0].Value))
	case "read_line":
		return "mujard::ble_rt.read_line()"
	case "pretend":
		if !g.inTestBody {
			return g.later(c.At, "pretend", phaseTests)
		}
		return fmt.Sprintf("mujard::ble_rt.pretend(%s)", g.text(c.Args[0].Value))
	}
	return g.later(c.At, "ble."+mem.Name, phaseLater)
}

// usbBackend reports whether 'use mujard.usb' has a runtime on this board: the cores that ship the standard
// Keyboard.h/Mouse.h HID API - 32U4 (Leonardo/Micro), RP2040/RP2350, SAMD21 (Zero) and RA4M1 (Uno R4). ESP32-S2/S3,
// nRF52840 and STM32F103 use different USB stacks and get C320 until their backends are written.
func usbBackend(b *Board) bool {
	switch b.Family {
	case "ATmega32U4", "RP2040", "RP2350", "SAMD21G18", "Renesas RA4M1":
		return true
	}
	return false
}

// usbCall is the USB HID device's members (mujard_usb.h): usb.type(text), key(text), mouse_move(x, y), click(),
// connected and usb.pretend(up). Boards without a standard-HID backend refuse with C320.
func (g *gen) usbCall(c *Call, mem *Member, args func() string) string {
	if !usbBackend(g.board) {
		return g.later(c.At, "USB on this board", phaseLater)
	}
	if mem.Name == "pretend" {
		if !g.inTestBody {
			return g.later(c.At, "pretend", phaseTests)
		}
		return fmt.Sprintf("mujard::usb_rt.pretend(%s)", args())
	}
	return fmt.Sprintf("mujard::usb_rt.%s(%s)", mem.Name, args())
}

// timerBackend reports whether 'use mujard.timer' has a runtime on this board: the standard timer-interrupt
// idioms on AVR's Timer1, ESP32's hw_timer and RP2040/RP2350's repeating_timer. Other cores refuse for now.
func timerBackend(b *Board) bool {
	switch b.Family {
	case "ATmega328P", "ATmega32U4", "ATmega2560", "ATmega168", "ATmega32", "ATmega8",
		"ESP32", "ESP32-S2", "ESP32-S3", "ESP32-C3", "ESP32-C6", "ESP32-H2",
		"RP2040", "RP2350":
		return true
	}
	return false
}

// timerCall is the hardware timer's members (mujard_timer.h): t.every(period), freq(hz), once(period),
// start(), stop(), elapsed() and the test-only pretend(). Boards without a timer backend refuse.
func (g *gen) timerCall(c *Call, mem *Member, args func() string) string {
	if !timerBackend(g.board) {
		return g.later(c.At, "Timer on this board", phaseLater)
	}
	obj := g.expr(mem.X)
	switch mem.Name {
	case "every":
		return fmt.Sprintf("%s.every(%s)", obj, g.period(c.Args[0].Value))
	case "freq":
		return fmt.Sprintf("%s.freq(%s)", obj, g.expr(c.Args[0].Value))
	case "once":
		return fmt.Sprintf("%s.once(%s)", obj, g.period(c.Args[0].Value))
	case "start", "stop", "elapsed":
		return fmt.Sprintf("%s.%s()", obj, mem.Name)
	case "pretend":
		if !g.inTestBody {
			return g.later(c.At, "pretend", phaseTests)
		}
		return fmt.Sprintf("%s.pretend()", obj)
	}
	return g.later(c.At, "timer."+mem.Name, phaseLater)
}

// timerIsr: 'on t.tick fast' runs the body in the timer ISR - 'on t.tick' without 'fast' polls 'due()' on the
// scheduler. On a board the ISR calls the handler at ISR speed; a PC test fires it through pretend/advance.
func (g *gen) timerIsr(d *EventBlock, mem *Member) {
	name := g.id("timer_isr")
	g.fns.add(d.At.Line, "#if defined(ARDUINO_ARCH_ESP32) || defined(ARDUINO_ARCH_ESP8266)")
	g.fns.add(0, "IRAM_ATTR")
	g.fns.add(0, "#endif")
	g.fns.add(d.At.Line, "static void %s() {", name)
	g.fns.add(0, "    if (mujard::stopping()) return;  // halted: fast handlers stop too (spec 6.7)")
	g.fns.add(0, "    mujard::IsrScope mujard_isr;")
	g.inFast = true
	g.block(&g.fns, d.Body, 1)
	g.inFast = false
	g.fns.add(0, "}")
	g.setup.add(d.At.Line, "    %s.onTick(%s);", g.expr(mem.X), name)
}

// netCall is wifi / clock / http / mqtt (mujard_wifi.h, mujard_clock.h, mujard_http.h, mujard_mqtt.h).
func (g *gen) netCall(c *Call, class, method string) string {
	switch class + "." + method {
	case "wifi.connect":
		if len(c.Args) > 0 && isSecretCall(c.Args[0].Value) {
			s := c.Args[0].Value.(*Call)
			return fmt.Sprintf("mujard::wifi_link.connect(%s, %s)", g.secretMacro(s, "_SSID"), g.secretMacro(s, "_PASSWORD"))
		}
	case "wifi.serve":
		pin := `""`
		if len(c.Args) > 1 {
			pin = g.cstr(c.Args[1].Value)
		}
		return fmt.Sprintf("mujard::wifi_link.serve(%s, %s)", g.cstr(c.Args[0].Value), pin)
	case "clock.sync":
		tz := "UTC0"
		for _, a := range c.Args {
			if s, ok := a.Value.(*StringLit); ok && len(s.Parts) == 1 {
				if p, ok := posixZone(s.Parts[0].Text); ok {
					tz = p
				}
			}
		}
		return fmt.Sprintf("mujard::clock_rt.sync(%s)", cString(tz))
	case "clock.source":
		return "mujard::clock_rt.source_rtc()"
	case "http.get", "http.post":
		// a suspending call outside the start-wait-result spot (waitFor): there is no synchronous http
		return g.later(c.At, class+"."+method+" here (use it as 'var r = http.get(\"â€¦\")' inside a task)", phaseLater)
	case "mqtt.connect":
		a := argMap(c, "broker", "user", "password")
		user, pass := "nullptr", "nullptr"
		if a["user"] != nil {
			user = g.cstr(a["user"])
		}
		if a["password"] != nil {
			pass = g.cstr(a["password"])
		}
		return fmt.Sprintf("mujard::mqtt_rt.connect(%s, %s, %s)", g.cstr(a["broker"]), user, pass)
	case "mqtt.publish":
		a := argMap(c, "topic", "value")
		return fmt.Sprintf("mujard::mqtt_rt.publish(%s, %s)", g.textValue(a["topic"]), g.text(a["value"])) // any value, as text: 42, 0.5 (a ratio), "on"
	case "mqtt.subscribe":
		a := argMap(c, "topic")
		return fmt.Sprintf("mujard::mqtt_rt.subscribe(%s)", g.textValue(a["topic"]))
	}
	return g.later(c.At, class+"."+method, phaseLater)
}

// storeCall is store.save(name, value) / store.load(name, default:) (mujard_store.h). The value goes in as the C++
// type of its Mujard type, so the same name always lands in the same slot type.
func (g *gen) storeCall(c *Call, mem *Member) string {
	keys := []string{"name", "value"}
	if mem.Name == "load" {
		keys = []string{"name", "default"}
	}
	a := argMap(c, keys...)
	name := `""`
	if n := a["name"]; n != nil {
		name = g.cstr(n)
	}
	v := a[keys[1]]
	if v == nil {
		// No default: the result is T? and the type comes from where the value goes - 'var x: int? = store.load("k")'.
		wt := g.wantType[c]
		if wt == nil {
			return g.later(c.At, "store.load without a default and without a declared type (write e.g. 'var x: int? = store.load(\"k\")')", phaseLater)
		}
		et := nonOpt(wt)
		if et.K == TAny {
			return g.later(c.At, "store.load without a default and without a declared type (write e.g. 'var x: int? = store.load(\"k\")')", phaseLater)
		}
		ec := g.ctype(et, 0)
		if ec == "auto" {
			return g.later(c.At, "store.load of a "+describe(et), phaseLater)
		}
		return fmt.Sprintf("mujard::store_db.load(%s, (%s*)nullptr)", name, ec)
	}
	t := nonOpt(g.typeOf(v))
	val := g.expr(v)
	switch {
	case t.K == TText:
		val = g.textValue(v)
	case t.K == TBool:
		val = "(bool)(" + val + ")"
	case t.K == TFloat || (t.K == TQty && !isTime(t) && t.Dim != "frequency"):
		val = "(float)(" + val + ")"
	default:
		val = g.castTo(t, val)
	}
	return fmt.Sprintf("mujard::store_db.%s(%s, %s)", mem.Name, name, val)
}

func (g *gen) site(class string) string {
	name := g.id(strings.ToLower(class))
	g.globals.add(0, "static mujard::%s %s;", class, name)
	return name
}

// powerCall: power.sleep(t[, wake_on: pin-event]) and power.deep_sleep(t) (spec 6.10). The wake
// pin's edge is worked out here: a pullup input is active-low, so 'pressed' is FALLING on it.
func (g *gen) powerCall(c *Call, mem *Member) string {
	t := "0"
	var wake Node
	for _, a := range c.Args {
		switch {
		case a.Name == "wake_on":
			wake = a.Value
		case a.Name == "" || a.Name == "time":
			t = g.expr(a.Value)
		}
	}
	pin, mode := "-1", "0"
	if wake != nil {
		m, ok := wake.(*Member)
		if !ok {
			return g.later(wake.Start(), "wake_on for this event", phaseLater)
		}
		xt := g.typeOf(m.X)
		if xt.K != TDevice || xt.Name != "input" {
			return g.later(wake.Start(), "wake_on "+m.Name, phaseLater)
		}
		pin = g.expr(m.X) + ".pin()"
		switch m.Name {
		case "changed":
			mode = "CHANGE"
		case "pressed", "released":
			low := false
			if id, ok := m.X.(*Ident); ok {
				if dev := g.devices[id.Name]; dev != nil {
					low = hasOpt(dev.Options, "pullup") != hasOpt(dev.Options, "inverted")
				}
			}
			if (m.Name == "pressed") == low {
				mode = "FALLING"
			} else {
				mode = "RISING"
			}
		default:
			return g.later(wake.Start(), "wake_on "+m.Name, phaseLater)
		}
	}
	return fmt.Sprintf("mujard::power_%s(%s, %s, %s)", cname(mem.Name), t, pin, mode)
}

func rangeEnds(n Node) (lo, hi Node, incl bool, ok bool) {
	r, ok := n.(*RangeExpr)
	if !ok {
		return nil, nil, false, false
	}
	return r.Lo, r.Hi, r.Inclusive, true
}

func (g *gen) builtinCall(c *Call, fn *Ident) string {
	t := nonOpt(g.typeOf(c))
	arg := func(i int) Node {
		if i < len(c.Args) {
			return c.Args[i].Value
		}
		return nil
	}
	ct := g.ctype(t, 0)
	switch fn.Name {
	case "now":
		return "mujard::now()"
	case "since":
		return "(mujard::now() - (" + g.expr(arg(0)) + "))"
	case "stopwatch":
		return "mujard::Stopwatch()"
	case "abs":
		return "mujard::abs_<" + ct + ">(" + g.expr(arg(0)) + ")"
	case "min", "max":
		x := g.expr(arg(0))
		for i := 1; i < len(c.Args); i++ {
			x = fmt.Sprintf("mujard::%s_<%s>(%s, %s)", fn.Name, ct, x, g.expr(arg(i)))
		}
		return x
	case "clamp":
		lo, hi, _, ok := rangeEnds(arg(1))
		if !ok {
			return g.later(c.At, "clamp without a range", phaseLater)
		}
		return fmt.Sprintf("mujard::clamp_<%s>(%s, %s, %s)", ct, g.expr(arg(0)), g.expr(lo), g.expr(hi))
	case "scale":
		a, b, _, ok1 := rangeEnds(arg(1))
		cc, d, _, ok2 := rangeEnds(arg(2))
		if !ok1 || !ok2 {
			return g.later(c.At, "scale without ranges", phaseLater)
		}
		return g.castTo(t, fmt.Sprintf("mujard::scale_((float)(%s), (float)(%s), (float)(%s), (float)(%s), (float)(%s))",
			g.expr(arg(0)), g.expr(a), g.expr(b), g.expr(cc), g.expr(d)))
	case "round":
		n := "0"
		if arg(1) != nil {
			n = g.expr(arg(1))
		}
		if t.K == TInt {
			return g.expr(arg(0))
		}
		return g.castTo(t, fmt.Sprintf("mujard::round_((float)(%s), %s)", g.expr(arg(0)), n))
	case "smooth":
		return g.castTo(t, fmt.Sprintf("%s.add((float)(%s), %s)", g.site("Smooth"), g.expr(arg(0)), g.expr(arg(1))))
	case "peak":
		return g.castTo(t, fmt.Sprintf("%s.update((float)(%s), %s)", g.site("Peak"), g.expr(arg(0)), g.expr(arg(1))))
	case "hysteresis":
		m := argMap(c, "value", "on", "off")
		return fmt.Sprintf("%s.update((float)(%s), (float)(%s), (float)(%s))", g.site("Hysteresis"), g.argOr(m, "value", "0"), g.argOr(m, "on", "0"), g.argOr(m, "off", "0"))
	case "between":
		lo, hi, _, ok := rangeEnds(arg(1))
		if !ok {
			return g.later(c.At, "between without a range", phaseLater)
		}
		x := g.expr(arg(0))
		return fmt.Sprintf("((%s) >= (%s) && (%s) <= (%s))", x, g.expr(lo), x, g.expr(hi))
	case "changed":
		return fmt.Sprintf("%s.update((float)(%s))", g.site("Changed"), g.expr(arg(0)))
	case "rising", "falling":
		return fmt.Sprintf("%s.%s(%s)", g.site("Edge"), fn.Name, g.expr(arg(0)))
	case "random":
		lo, hi, incl, ok := rangeEnds(arg(0))
		if !ok {
			return g.later(c.At, "random without a range", phaseLater)
		}
		if t.K == TInt {
			h := g.expr(hi)
			if !incl {
				h = "(" + h + ") - 1"
			}
			return fmt.Sprintf("mujard::random_range((int32_t)(%s), (int32_t)(%s))", g.expr(lo), h)
		}
		return g.castTo(t, fmt.Sprintf("mujard::random_range((float)(%s), (float)(%s))", g.expr(lo), g.expr(hi)))
	case "watchdog":
		return "mujard::watchdog(" + g.expr(arg(0)) + ")"
	case "serial":
		return "mujard::serial_begin((uint32_t)(" + g.expr(arg(0)) + "))"
	case "secret":
		return g.secretMacro(c, "") // a text secret (mqtt password); Wi-Fi credentials are handled in wifi.connect
	case "notify", "show", "plot", "ask":
		s, _ := g.appCall(c, fn.Name)
		return s
	case "notified", "printed":
		if g.inTestBody && len(c.Args) == 1 {
			return fmt.Sprintf("mujard::test_%s(%s)", fn.Name, g.text(c.Args[0].Value))
		}
		return g.later(c.At, fn.Name, phaseTests)
	case "power_cut", "brown_out":
		// testBlock handles these written directly in a test body; anywhere else they would have to resume mid-block
		return g.later(c.At, fn.Name+"() inside a block (write it on its own line in the test itself)", phaseLater)
	case "stopped", "halted":
		if g.inTestBody && len(c.Args) == 0 {
			return "mujard::test_stopped()" // records that the halt was checked - an unobserved halt fails the test
		}
		return g.later(c.At, fn.Name, phaseTests)
	case "answer":
		if g.inTestBody && len(c.Args) == 1 {
			if isNone(c.Args[0].Value) {
				return "mujard::test_answer(-1)"
			}
			return fmt.Sprintf("mujard::test_answer((%s) ? 1 : 0)", g.expr(c.Args[0].Value))
		}
		return g.later(c.At, "answer", phaseTests)
	case "advance":
		return g.later(c.At, "advance inside an expression", phaseLater) // as a statement it is a wait (advanceTime)
	case "battery_on", "gauge":
		return g.later(c.At, fn.Name, phaseLibs)
	}
	// a C++ library function
	var args []string
	for _, a := range c.Args {
		args = append(args, g.expr(a.Value))
	}
	return fmt.Sprintf("%s(%s)", g.libName(fn.Name), strings.Join(args, ", "))
}

// modCallExpr is motors.fn(args) and motors.Struct(args): the call is the program's, the declaration the module's.
func (g *gen) modCallExpr(c *Call, mod, name string) string {
	mc := g.mods[mod]
	if mc == nil || mc.failed {
		return "0"
	}
	sym := mc.global.syms[name]
	if sym == nil {
		return "0"
	}
	if sym.fn != nil {
		d := sym.fn
		var names []string
		for _, p := range d.Params {
			names = append(names, p.Name)
		}
		m := argMap(c, names...)
		var args []string
		for _, p := range d.Params {
			var pt *Type
			g.withModule(mc, mod, false, func() { pt = g.declTypeFromTypeExpr(p.Type) })
			if v, ok := m[p.Name]; ok {
				g.setWant(v, pt)
				args = append(args, g.expr(v))
			} else if p.Default != nil {
				var x string
				g.withModule(mc, mod, true, func() {
					g.setWant(p.Default, pt)
					x = g.expr(p.Default)
				})
				args = append(args, x)
			}
		}
		return fmt.Sprintf("%s::%s(%s)", cname(mod), cname(d.Name), strings.Join(args, ", "))
	}
	if sym.strukt != nil {
		s := sym.strukt
		m := map[string]Node{}
		for i, a := range c.Args {
			if a.Name != "" {
				m[a.Name] = a.Value
			} else if i < len(s.Fields) {
				m[s.Fields[i].Name] = a.Value
			}
		}
		var fields []string
		for _, f := range s.Fields {
			var ft *Type
			g.withModule(mc, mod, false, func() { ft = g.fieldType(f) })
			if v, ok := m[f.Name]; ok {
				g.setWant(v, ft)
				fields = append(fields, g.expr(v))
			} else if f.Default != nil {
				var x string
				g.withModule(mc, mod, true, func() {
					g.setWant(f.Default, ft)
					x = g.expr(f.Default)
				})
				fields = append(fields, x)
			} else {
				fields = append(fields, "{}")
			}
		}
		return fmt.Sprintf("%s::%s{%s}", cname(mod), cname(s.Name), strings.Join(fields, ", "))
	}
	var args []string
	for _, a := range c.Args {
		args = append(args, g.expr(a.Value))
	}
	return fmt.Sprintf("%s::%s(%s)", cname(mod), cname(name), strings.Join(args, ", "))
}

func (g *gen) methodCall(c *Call, mem *Member) string {
	xt := nonOpt(g.typeOf(mem.X))
	obj := g.expr(mem.X)
	m := func(names ...string) map[string]Node { return argMap(c, names...) }
	args := func() string {
		var a []string
		for _, x := range c.Args {
			a = append(a, g.expr(x.Value))
		}
		return strings.Join(a, ", ")
	}
	switch {
	case xt.K == TObject && strings.HasPrefix(xt.Name, "mod:"):
		return g.modCallExpr(c, strings.TrimPrefix(xt.Name, "mod:"), mem.Name)
	case xt.K == TDevice && (xt.Name == "output" || xt.Name == "input" || xt.Name == "pwm"):
		switch mem.Name {
		case "blink":
			a := m("rate", "for", "on")
			return fmt.Sprintf("%s.blink(%s, %s, %s)", obj, g.period(a["rate"]), g.argOr(a, "for", "0"), g.argOr(a, "on", "-1"))
		case "pattern":
			a := m("pattern", "step", "repeat")
			return fmt.Sprintf("%s.pattern(%s, %s, %s)", obj, g.literalText(a["pattern"]), g.argOr(a, "step", "150000"), g.argOr(a, "repeat", "-1"))
		case "melody":
			a := m("notes", "tempo")
			return fmt.Sprintf("%s.melody(%s, %s)", obj, g.literalText(a["notes"]), g.argOr(a, "tempo", "120"))
		case "tone":
			a := m("pitch", "length")
			return fmt.Sprintf("%s.tone(%s, %s)", obj, g.argOr(a, "pitch", "0"), g.argOr(a, "length", "0"))
		case "fade":
			a := m("to", "over")
			return fmt.Sprintf("%s.fade(%s, %s)", obj, g.argOr(a, "to", "0"), g.argOr(a, "over", "0"))
		case "measure_pulse":
			a := m("level", "timeout")
			return fmt.Sprintf("%s.measure_pulse(%s, %s)", obj, g.argOr(a, "level", "true"), g.argOr(a, "timeout", "1000000"))
		case "read":
			if g.inFast && xt.Name == "input" {
				return obj + ".raw()" // an interrupt sees the wire as it is, not the debounced state
			}
		case "pretend":
			if g.inTestBody && xt.Name == "input" && len(c.Args) == 1 {
				// a length pretends a pulse (measure_pulse sees test_pretend_echo); a level pretends the wire
				if t := g.typeOf(c.Args[0].Value); t.K == TQty && t.Dim == "length" {
					return fmt.Sprintf("mujard::test_pretend_echo((float)(%s))", g.expr(c.Args[0].Value))
				}
				return fmt.Sprintf("mujard::test_pretend_pin(%s.pin(), %s.activeLow(), (bool)(%s))", obj, obj, g.expr(c.Args[0].Value))
			}
			return g.later(c.At, "pretend on an "+xt.Name, phaseLater)
		}
		return fmt.Sprintf("%s.%s(%s)", obj, mem.Name, args())
	case xt.K == TDevice && xt.Name == "analog":
		if mem.Name == "pretend" {
			if g.inTestBody && len(c.Args) == 1 {
				return g.pretendAnalog(obj, c.Args[0].Value)
			}
			return g.later(c.At, "pretend", phaseTests)
		}
		return fmt.Sprintf("%s.%s(%s)", obj, mem.Name, args())
	case xt.K == TObject && xt.Name == "battery":
		if mem.Name == "pretend" {
			if g.inTestBody && len(c.Args) == 1 {
				return g.pretendAnalog(obj, c.Args[0].Value) // battery reads its pin - pretend sets that pin
			}
			return g.later(c.At, "pretend", phaseTests)
		}
		return fmt.Sprintf("%s.%s(%s)", obj, mem.Name, args())
	case xt.K == TObject && xt.Name == "can":
		if mem.Name == "pretend" {
			if !g.inTestBody {
				return g.later(c.At, "pretend", phaseTests)
			}
			return fmt.Sprintf("%s.pretend(%s)", obj, args())
		}
		return fmt.Sprintf("%s.%s(%s)", obj, mem.Name, args())
	case xt.K == TDevice && xt.Name == "group":
		if mem.Name == "write" {
			return fmt.Sprintf("mujard::group_write(%s, %d, %s)", obj, g.groupSize(mem.X), args())
		}
	case xt.K == TObject:
		switch xt.Name {
		case "board":
			switch mem.Name {
			case "restart":
				return "mujard::restart()"
			case "halt":
				return "mujard::halt(" + g.text(c.Args[0].Value) + ")"
			}
		case "log":
			level := map[string]string{"info": "INFO", "warn": "WARN", "error": "ERROR"}[mem.Name]
			return fmt.Sprintf("mujard::log_line(%q, %s)", level, g.text(c.Args[0].Value))
		case "math":
			fnName := map[string]string{"sqrt": "sqrtf", "pow": "powf", "log": "logf", "sin": "mujard::sin_deg", "cos": "mujard::cos_deg",
				"tan": "mujard::tan_deg", "atan2": "mujard::atan2_deg"}[mem.Name]
			return fmt.Sprintf("%s(%s)", fnName, args())
		case "i2c":
			switch mem.Name {
			case "scan":
				return "mujard::i2c::scan()"
			case "write":
				a := m("address", "data")
				return fmt.Sprintf("mujard::i2c::write((int32_t)(%s), %s)", g.argOr(a, "address", "0"), g.argOr(a, "data", "mujard::List<int32_t, 1>()"))
			case "read":
				a := m("address", "count")
				return fmt.Sprintf("mujard::i2c::read((int32_t)(%s), (int32_t)(%s))", g.argOr(a, "address", "0"), g.argOr(a, "count", "1"))
			case "pretend":
				if !g.inTestBody || len(c.Args) != 2 {
					return g.later(c.At, "pretend", phaseTests)
				}
				a := m("address", "reply")
				if isNone(c.Args[1].Value) {
					return fmt.Sprintf("mujard::test_i2c_gone((int32_t)(%s))", g.argOr(a, "address", "0"))
				}
				return fmt.Sprintf("mujard::test_i2c_pretend((int32_t)(%s), %s)", g.argOr(a, "address", "0"), g.argOr(a, "reply", "mujard::List<int32_t, 1>()"))
			}
			return g.later(c.At, "i2c."+mem.Name, phaseLater)
		case "random":
			switch mem.Name {
			case "seed":
				if len(c.Args) > 0 {
					return "randomSeed(" + args() + ")"
				}
				return "(void)0"
			case "pick":
				l := g.expr(c.Args[0].Value)
				return fmt.Sprintf("%s.items[mujard::random_range((int32_t)0, %s.n - 1)]", l, l)
			}
		case "servo":
			if mem.Name == "move" {
				a := m("angle", "over")
				return fmt.Sprintf("%s.move(%s, %s)", obj, g.argOr(a, "angle", "0"), g.argOr(a, "over", "0"))
			}
			if mem.Name == "pretend" {
				return g.later(c.At, "pretend on a servo", phaseLater) // read its angle instead
			}
		case "ultrasonic":
			if mem.Name == "pretend" {
				if g.inTestBody && len(c.Args) == 1 {
					if isNone(c.Args[0].Value) {
						return "mujard::test_pretend_echo(-1.0f)" // nothing answers: distance() is none
					}
					return fmt.Sprintf("mujard::test_pretend_echo((float)(%s))", g.expr(c.Args[0].Value))
				}
				return g.later(c.At, "pretend", phaseTests)
			}
		case "serial":
			switch mem.Name {
			case "pretend":
				if g.inTestBody && len(c.Args) == 1 {
					return fmt.Sprintf("mujard::test_pretend_serial(%s)", g.text(c.Args[0].Value))
				}
				return g.later(c.At, "pretend on the serial port", phaseTests)
			}
			return g.later(c.At, "serial."+mem.Name+" here (read_line waits; use it as 'var x = serial.read_line()' in a task)", phaseLater)
		case "uart":
			switch mem.Name {
			case "write":
				return fmt.Sprintf("%s.write(%s)", obj, g.text(c.Args[0].Value))
			case "pretend":
				return fmt.Sprintf("%s.pretend(%s)", obj, g.text(c.Args[0].Value))
			}
			return g.later(c.At, "uart."+mem.Name+" here (read_line waits; use it as 'var x = "+mem.Name+"' in a task)", phaseLater)
		case "spi":
			if mem.Name == "device" {
				return g.spiCall(c)
			}
		case "rtc":
			if mem.Name == "set" {
				return fmt.Sprintf("mujard::rtc_dev.set(%s)", args())
			}
		case "clock", "wifi", "mqtt", "http":
			return g.netCall(c, xt.Name, mem.Name)
		case "ble":
			return g.bleCall(c, mem, args)
		case "timer":
			return g.timerCall(c, mem, args)
		case "usb":
			return g.usbCall(c, mem, args)
		case "ethernet":
			if !g.board.HasEthernet() {
				return g.later(c.At, "Ethernet on this board", phaseLater)
			}
			if mem.Name == "pretend" {
				if !g.inTestBody {
					return g.later(c.At, "pretend", phaseTests)
				}
				return fmt.Sprintf("mujard::eth_rt.pretend(%s)", args())
			}
			return fmt.Sprintf("mujard::eth_rt.%s(%s)", mem.Name, args())
		case "store":
			return g.storeCall(c, mem)
		case "power":
			return g.powerCall(c, mem)
		}
		return fmt.Sprintf("%s.%s(%s)", obj, g.libMemberName(xt.Name, mem.Name), args())
	case xt.K == TInt:
		n := args()
		switch mem.Name {
		case "bit":
			return fmt.Sprintf("(((%s) >> (%s)) & 1)", obj, n)
		case "set_bit":
			return fmt.Sprintf("((%s) | (1L << (%s)))", obj, n)
		case "clear_bit":
			return fmt.Sprintf("((%s) & ~(1L << (%s)))", obj, n)
		case "toggle_bit":
			return fmt.Sprintf("((%s) ^ (1L << (%s)))", obj, n)
		}
	case xt.K == TText:
		switch mem.Name {
		case "contains", "starts_with", "ends_with", "split":
			return fmt.Sprintf("%s.%s(%s)", g.textValue(mem.X), mem.Name, g.text(c.Args[0].Value))
		}
		return fmt.Sprintf("%s.%s(%s)", g.textValue(mem.X), mem.Name, args())
	}
	return fmt.Sprintf("%s.%s(%s)", obj, mem.Name, args())
}

// member is a property: light.percent, board.uptime, watch.elapsed, Mode.idle, reading.value, ...
func (g *gen) member(e *Member) string {
	xt := g.typeOf(e.X)
	if xt.K == TObject && strings.HasPrefix(xt.Name, "enum:") {
		return qname(strings.TrimPrefix(xt.Name, "enum:")) + "::" + cname(e.Name)
	}
	if xt.K == TObject && builtinEnums[xt.Name].t != nil {
		return g.builtinEnum(xt.Name, e.Name, e.At) // Level.high, BatteryType.lipo_1s
	}
	if xt.K == TObject && strings.HasPrefix(xt.Name, "mod:") {
		// a module member: motors.speed -> motors::speed (an optional unwraps to .v like a plain name)
		mod := strings.TrimPrefix(xt.Name, "mod:")
		s := cname(mod) + "::" + cname(e.Name)
		if mc := g.mods[mod]; mc != nil {
			if sym := mc.global.syms[e.Name]; sym != nil && sym.t != nil && sym.t.Opt && !g.typeOf(e).Opt {
				s += ".v"
			}
		}
		if !g.inFast && !g.inAtomic && g.atomics[mod+"."+e.Name] {
			return "mujard::atomic_read(" + s + ")"
		}
		return s
	}
	xt = nonOpt(xt)
	obj := g.expr(e.X)
	switch {
	case xt.K == TDevice && xt.Name == "analog":
		return fmt.Sprintf("%s.%s()", obj, e.Name)
	case xt.K == TDevice && xt.Name == "group" && e.Name == "length":
		return strconv.Itoa(g.groupSize(e.X))
	case xt.K == TQty && e.Name == "number":
		if c, ok := e.X.(*Conversion); ok {
			return fmt.Sprintf("((float)(%s) / %s)", g.expr(c.X), cfloat(unitFactor(c.Unit)))
		}
		if isTime(xt) {
			return fmt.Sprintf("mujard::secs(%s)", obj)
		}
		if xt.Dim == "frequency" {
			return fmt.Sprintf("((float)(%s) / 1000.0f)", obj)
		}
		return obj
	case xt.K == TObject:
		switch xt.Name {
		case "regs":
			// regs.PORTB -> mujard::reg8(0x25) - a small value object poking the chip register
			if table, ok := regsOf(g.board); ok {
				if r, ok := table[e.Name]; ok {
					return fmt.Sprintf("mujard::reg%d(%d)", r.Bits, r.Addr)
				}
			}
			return "mujard::reg8(0)" // unreachable on a checked program
		}
		if regName(xt) != "" {
			// regs.PORTB.PORTB5 -> the bit-field's mask as a literal int
			if table, ok := regsOf(g.board); ok {
				if r, ok := table[regName(xt)]; ok {
					if b, ok := r.Fields[e.Name]; ok {
						return strconv.FormatInt(b, 10)
					}
				}
			}
			return "0" // unreachable on a checked program
		}
		switch xt.Name {
		case "board":
			switch e.Name {
			case "name":
				return "mujard::board_name()"
			case "chip":
				return fmt.Sprintf("mujard::Text(%q)", g.board.Family)
			case "id":
				return "mujard::board_id()"
			case "uptime":
				return "mujard::now()"
			case "cores":
				return strconv.Itoa(g.board.Cores())
			case "dma":
				return strconv.FormatBool(g.board.DMA)
			case "timers":
				return strconv.Itoa(g.board.Timers)
			case "rtos":
				return fmt.Sprintf("mujard::Text(%q)", g.board.RTOS())
			case "ethernet":
				return strconv.FormatBool(g.board.HasEthernet())
			case "free_memory":
				if !g.board.CanMeasureFreeMemory() {
					return g.later(e.At, "board.free_memory on this board", phaseLater)
				}
				return "((float)mujard::free_memory())"
			case "voltage":
				if !g.board.CanMeasureVoltage() {
					return "mujard::Opt<float>()"
				}
				return "mujard::board_voltage()"
			case "reset_reason":
				if g.testMode {
					return "mujard::test_reset_reason()"
				}
				if g.board.CanReadResetReason() {
					return "mujard::board_reset_reason()"
				}
			case "overruns":
				return "((int32_t)mujard::overruns)"
			case "dropped_events":
				return "((int32_t)mujard::dropped_events)"
			}
			return g.later(e.At, "board."+e.Name, phaseLater)
		case "math":
			switch e.Name {
			case "PI":
				return "3.14159265f"
			case "E":
				return "2.71828183f"
			}
		case "stopwatch", "block", "servo", "battery", "queue":
			return fmt.Sprintf("%s.%s()", obj, e.Name)
		case "wifi":
			switch e.Name {
			case "connected":
				return "mujard::wifi_link.connected()"
			case "client":
				return "mujard::wifi_link.client()"
			}
		case "store":
			switch e.Name {
			case "pending":
				// store.save's queue plus a saved var whose live value isn't in storage yet (spec 11).
				return "(mujard::store_db.pending() || mujard::saved_pending())"
			case "failed", "committed", "last_commit":
				return fmt.Sprintf("mujard::store_db.%s()", e.Name)
			}
		case "ble":
			if e.Name == "connected" {
				if !bleBackend(g.board) {
					return g.later(e.At, "BLE on this board", phaseLater)
				}
				return "mujard::ble_rt.connected()"
			}
		case "usb":
			if e.Name == "connected" {
				if !usbBackend(g.board) {
					return g.later(e.At, "USB on this board", phaseLater)
				}
				return "mujard::usb_rt.connected()"
			}
		case "ethernet":
			if e.Name == "connected" || e.Name == "ip" {
				if !g.board.HasEthernet() {
					return g.later(e.At, "Ethernet on this board", phaseLater)
				}
				return fmt.Sprintf("mujard::eth_rt.%s()", e.Name)
			}
		case "clock":
			switch e.Name {
			case "time":
				return "mujard::clock_rt.time()"
			case "date":
				return "mujard::clock_rt.date()"
			}
		case "time_of_day", "date", "mqtt_message":
			return obj + "." + e.Name // plain fields
		case "http_reply":
			if e.Name == "json" {
				return fmt.Sprintf("mujard::Json(%s.text)", obj)
			}
			return obj + "." + e.Name
		}
	case xt.K == TObject && xt.Name == "json" && e.Name == "ok":
		return obj + ".ok"
	case xt.K == TText || xt.K == TList:
		if e.Name == "length" {
			return obj + ".length()"
		}
	case xt.K == TStruct:
		return obj + "." + cname(e.Name)
	case xt.K == TInt && e.Name == "bytes":
		return fmt.Sprintf("mujard::int_bytes(%s)", obj)
	}
	return fmt.Sprintf("%s.%s", obj, g.libMemberName(xt.Name, e.Name))
}

// libName is a described library's top-level name as it exists in C++ (a snake_case spelling resolves to
// its camelCase entry); names not from a described library come back unchanged.
func (g *gen) libName(written string) string {
	if g.info != nil && g.info.LibReal != nil {
		if r, ok := g.info.LibReal[written]; ok {
			return r
		}
	}
	return written
}

// isLibGlobal tells a described library's top-level name from the program's own: the checker records what each
// program name was declared as, and never a library's.
func (g *gen) isLibGlobal(e *Ident) bool {
	if g.info == nil || g.info.LibGlobals[e.Name] == nil || g.decl[e] != nil {
		return false
	}
	gl, builtin := globals[e.Name]
	if builtin && gl.lib != "" && !g.info.Libs[gl.lib] {
		builtin = false // its Mujard library isn't used: the C++ library's name (the checker resolved it so)
	}
	_, device := g.devices[e.Name]
	return !builtin && !device
}

// libMemberName is a member of an object as C++ spells it. A described library class's members are the library's
// own (pinMode, loop, print ...) and keep their names; anything else is renamed clear of Arduino's names.
func (g *gen) libMemberName(class, written string) string {
	real := g.libMember(class, written)
	if g.info != nil && g.info.LibClasses[class] != nil {
		return real
	}
	return cname(real)
}

// libMember is libName for members: 'strip.set_pixel_color' -> 'strip.setPixelColor'.
func (g *gen) libMember(class, written string) string {
	if g.info != nil && g.info.LibReal != nil {
		if r, ok := g.info.LibReal[class+"."+written]; ok {
			return r
		}
	}
	return written
}
