package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadJobValidation(t *testing.T) {
	dir := t.TempDir()
	sketch := filepath.Join(dir, "sketch")
	if err := os.MkdirAll(sketch, 0o755); err != nil {
		t.Fatal(err)
	}
	write := func(body string) string {
		f := filepath.Join(dir, "job.json")
		if err := os.WriteFile(f, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
		return f
	}
	if _, err := loadJob(write(`{"schema":"nope"}`)); err == nil {
		t.Fatal("bad schema accepted")
	}
	if _, err := loadJob(write(`{"schema":"mujard-build.v1","board":"m328p"}`)); err == nil {
		t.Fatal("missing sketchDir accepted")
	}
	good := `{"schema":"mujard-build.v1","board":"m328p","sketchDir":"` + strings.ReplaceAll(sketch, `\`, `\\`) + `","outDir":"` + strings.ReplaceAll(filepath.Join(dir, "out"), `\`, `\\`) + `"}`
	j, err := loadJob(write(good))
	if err != nil {
		t.Fatal(err)
	}
	if j.TimeoutSec != 600 {
		t.Fatalf("default timeout %d", j.TimeoutSec)
	}
	if _, err := loadJob(write(`{"schema":"mujard-build.v1","board":"m328p","sketchDir":"` + strings.ReplaceAll(filepath.Join(dir, "nope"), `\`, `\\`) + `","outDir":"/tmp/o"}`)); err == nil {
		t.Fatal("missing sketch dir accepted")
	}
}

func TestResolveOptions(t *testing.T) {
	b := &BoardManifest{
		Short: "m328pe",
		Options: map[string]OptionDef{
			"bootloader": {Choices: []string{"new", "old"}, Default: "new"},
		},
	}
	got, err := resolveOptions(b, nil)
	if err != nil || got["bootloader"] != "new" {
		t.Fatalf("defaults: %v %v", got, err)
	}
	if _, err := resolveOptions(b, map[string]string{"bootloader": "ancient"}); err == nil {
		t.Fatal("bad option value accepted")
	}
	if _, err := resolveOptions(b, map[string]string{"nope": "x"}); err == nil {
		t.Fatal("unknown option accepted")
	}
}

// fakePackage builds a package dir: a core with one .c file, a variant dir and dummy tools.
func fakePackage(t *testing.T) (bdir, pdir string) {
	t.Helper()
	root := t.TempDir()
	core := filepath.Join(root, "core", "cores", "arduino")
	if err := os.MkdirAll(filepath.Join(root, "core", "variants", "std"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(core, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(core, "wiring.c"), []byte("int x;\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	tools := filepath.Join(root, "tools", "bin")
	if err := os.MkdirAll(tools, 0o755); err != nil {
		t.Fatal(err)
	}
	var toolJSON []string
	for _, n := range []string{"cc", "cxx", "ar", "objcopy", "size"} {
		p := filepath.Join(tools, n+".exe")
		if err := os.WriteFile(p, []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
		toolJSON = append(toolJSON, `"`+n+`": "`+strings.ReplaceAll(p, `\`, `/`)+`"`)
	}
	board := `{"schema":"mujard-board.v1","id":"microchip:avr:t","short":"t","name":"t",
	  "vendor":"microchip","family":"avr","package":"pkg","core":"cores/arduino",
	  "variant":"variants/std","vars":{"mcu":"atmega328p","fCpu":"16000000L","boardDefine":"ARDUINO_AVR_UNO"},
	  "output":"hex","flashMax":32256,"ramMax":2048}`
	pkg := `{"schema":"mujard-package.v1","id":"pkg","family":"avr","root":"` +
		strings.ReplaceAll(filepath.Join(root, "core"), `\`, `/`) + `","tools":{` + strings.Join(toolJSON, ",") + `}}`
	return writeManifests(t, map[string]string{"t.json": board}, map[string]string{"p.json": pkg})
}

func TestAvrPlan(t *testing.T) {
	bdir, pdir := fakePackage(t)
	reg, err := LoadRegistry(bdir, pdir)
	if err != nil {
		t.Fatal(err)
	}
	sketch := t.TempDir()
	if err := os.WriteFile(filepath.Join(sketch, "blink.ino"), []byte("void setup(){}\nvoid loop(){}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(t.TempDir(), "out")
	j := &Job{Schema: SchemaV1, Board: "t", SketchDir: sketch, OutDir: out, TimeoutSec: 60}
	plan, err := buildPlan(reg, j)
	if err != nil {
		t.Fatal(err)
	}
	if len(plan.SketchSrcs) != 1 || !strings.HasSuffix(plan.SketchSrcs[0], "blink.ino.cpp") {
		t.Fatalf("ino not preprocessed: %v", plan.SketchSrcs)
	}
	data, _ := os.ReadFile(plan.SketchSrcs[0])
	if !strings.Contains(string(data), "#include <Arduino.h>") {
		t.Fatal("preprocessed ino missing Arduino.h")
	}
	var names []string
	for _, s := range plan.Steps {
		names = append(names, s.Name)
	}
	for _, want := range []string{"archive core", "link", "hex", "size"} {
		found := false
		for _, n := range names {
			if n == want {
				found = true
			}
		}
		if !found {
			t.Fatalf("missing step %q in %v", want, names)
		}
	}
}

// The AVR core has wiring_pulse.c and wiring_pulse.S (pulseIn's countPulseASM). Both must reach core.a as
// their own objects; when they shared wiring_pulse.o, pulseIn() failed to link.
func TestSameNameCAndAsmGetTheirOwnObjects(t *testing.T) {
	bdir, pdir := fakePackage(t)
	reg, err := LoadRegistry(bdir, pdir)
	if err != nil {
		t.Fatal(err)
	}
	pkg := reg.Package(mustResolve(t, reg, "t"))
	core := pkg.subRoot("cores/arduino")
	for name, body := range map[string]string{"wiring_pulse.c": "int p;\n", "wiring_pulse.S": "countPulseASM:\n"} {
		if err := os.WriteFile(filepath.Join(core, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	sketch := t.TempDir()
	if err := os.WriteFile(filepath.Join(sketch, "sketch.ino"), []byte("void setup(){}\nvoid loop(){}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	j := &Job{Schema: SchemaV1, Board: "t", SketchDir: sketch, OutDir: filepath.Join(t.TempDir(), "out"), TimeoutSec: 60, NoCache: true}
	plan, err := buildPlan(reg, j)
	if err != nil {
		t.Fatal(err)
	}
	outs := map[string]string{} // object -> source
	var asm *Step
	for i, s := range plan.Steps {
		if !strings.HasPrefix(s.Name, "compile ") {
			continue
		}
		src := strings.TrimPrefix(s.Name, "compile ")
		if prev, dup := outs[strings.ToLower(s.Out)]; dup {
			t.Fatalf("%s and %s both compile to %s", prev, src, s.Out)
		}
		outs[strings.ToLower(s.Out)] = src
		if src == "wiring_pulse.S" {
			asm = &plan.Steps[i]
		}
	}
	if asm == nil || !strings.HasSuffix(asm.Out, "wiring_pulse.S.o") || !strings.Contains(strings.Join(asm.Args, " "), "assembler-with-cpp") {
		t.Fatalf("wiring_pulse.S is not assembled into its own object: %+v", asm)
	}
	for _, s := range plan.Steps {
		if s.Name == "archive core" {
			args := strings.Join(s.Args, " ")
			if !strings.Contains(args, "wiring_pulse.c.o") || !strings.Contains(args, "wiring_pulse.S.o") {
				t.Fatalf("core.a misses one of the wiring_pulse objects: %v", s.Args)
			}
		}
	}
}

func TestObjNamerNumbersRepeats(t *testing.T) {
	n := objNamer{}
	root := filepath.Join("obj", "libs", "Wire")
	a := n.path(root, filepath.Join("src", "twi.c"))
	b := n.path(root, filepath.Join("src", "utility", "twi.c"))
	c := n.path(root, filepath.Join("src", "TWI.C"))
	if a != filepath.Join(root, "twi.c.o") || b != filepath.Join(root, "twi.c.2.o") || c != filepath.Join(root, "TWI.C.3.o") {
		t.Fatalf("got %s, %s, %s", a, b, c)
	}
	if d := n.path(filepath.Join("obj", "core"), "twi.c"); d != filepath.Join("obj", "core", "twi.c.o") {
		t.Fatalf("another folder should start fresh: %s", d)
	}
}

// A library that isn't part of the core (Servo) comes from one of the package's libraryRoots; a library the
// core has itself (Wire) still comes from the core, even when an extra root has one of the same name.
func TestLibrariesFromExtraRoots(t *testing.T) {
	mk := func(path, body string) {
		t.Helper()
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	coreLibs, extra := t.TempDir(), t.TempDir()
	mk(filepath.Join(coreLibs, "Wire", "src", "Wire.h"), "// core Wire\n")
	mk(filepath.Join(coreLibs, "Wire", "src", "utility", "twi.c"), "int twi;\n")
	mk(filepath.Join(extra, "Servo", "src", "Servo.h"), "#include \"avr/ServoTimers.h\"\n")
	mk(filepath.Join(extra, "Servo", "src", "avr", "Servo.cpp"), "#include <Servo.h>\n")
	mk(filepath.Join(extra, "Wire", "src", "Wire.h"), "// another Wire\n")
	mk(filepath.Join(extra, "receipt.json"), "{}") // a package's receipt sits beside its libraries
	sketch := t.TempDir()
	mk(filepath.Join(sketch, "mujard_program.cpp"), "#include \"mujard_servo.h\"\n#include <Wire.h>\n")
	mk(filepath.Join(sketch, "mujard_servo.h"), "#include <Servo.h>\n")
	p := &Plan{
		Pkg:        &Package{LibRoot: coreLibs, LibRoots: []string{extra, filepath.Join(extra, "not-installed")}},
		Job:        &Job{SketchDir: sketch},
		SketchSrcs: []string{filepath.Join(sketch, "mujard_program.cpp")},
	}
	incs, srcs, err := p.resolveLibs()
	if err != nil {
		t.Fatal(err)
	}
	if len(srcs["Servo"]) != 1 || !strings.HasSuffix(srcs["Servo"][0], filepath.Join("avr", "Servo.cpp")) {
		t.Fatalf("Servo not resolved from the extra root: %v", srcs)
	}
	if len(srcs["Wire"]) != 1 || !strings.HasPrefix(srcs["Wire"][0], coreLibs) {
		t.Fatalf("Wire should come from the core: %v", srcs["Wire"])
	}
	joined := strings.Join(incs, "|")
	if !strings.Contains(joined, filepath.Join(extra, "Servo", "src")) || strings.Contains(joined, filepath.Join(extra, "Wire")) {
		t.Fatalf("include dirs %v", incs)
	}
}

// A library is used for #include <h> only when it has h at that path: the Pico core's TinyUSB library has a tusb.h
// and a time.h-like set of names, and matching on the file name alone pulled it into every build, where its tusb.h
// replaced the Pico SDK's and the core's USB.cpp failed to compile.
func TestLibraryMatchesTheIncludePath(t *testing.T) {
	mk := func(path string) {
		t.Helper()
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte("\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	libs, sketch := t.TempDir(), t.TempDir()
	mk(filepath.Join(libs, "TinyUSB", "src", "tusb.h"))
	mk(filepath.Join(libs, "TinyUSB", "src", "class", "time.h"))
	mk(filepath.Join(libs, "EEPROM", "src", "EEPROM.h"))
	prog := filepath.Join(sketch, "mujard_program.cpp")
	if err := os.WriteFile(prog, []byte("#include <pico/time.h>\n#include <EEPROM.h>\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	p := &Plan{Pkg: &Package{LibRoot: libs}, Job: &Job{SketchDir: sketch}, SketchSrcs: []string{prog}}
	_, srcs, err := p.resolveLibs()
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := srcs["TinyUSB"]; ok {
		t.Fatal("<pico/time.h> pulled in a library that only has a class/time.h")
	}
	if _, ok := srcs["EEPROM"]; !ok {
		t.Fatalf("EEPROM.h not resolved: %v", srcs)
	}
}

// esptool's PyInstaller unpack failing (an antivirus lock on Windows) is a reason to run it again; esptool's own
// errors and compile errors are not.
func TestToolFailedToStart(t *testing.T) {
	for _, c := range []struct {
		out  string
		want bool
	}{
		{"[PYI-1108:ERROR] Failed to extract ucrtbase.dll: failed to open target file!\nfopen: Permission denied\n", true},
		{"[PYI-12056:ERROR] Failed to extract entry: ucrtbase.dll.\n", true},
		{"A fatal error occurred: Invalid ELF file\n", false},
		{"sketch.ino:3:1: error: 'x' was not declared in this scope\n", false},
	} {
		if got := toolFailedToStart([]byte(c.out)); got != c.want {
			t.Errorf("%q: got %v", c.out, got)
		}
	}
}

// A flat library (no src/) is its top folder and utility/: its examples are never compiled into the firmware.
func TestFlatLibrarySkipsExamples(t *testing.T) {
	root := t.TempDir()
	for _, f := range []string{"Neo/Neo.h", "Neo/Neo.cpp", "Neo/utility/u.c", "Neo/examples/Ble/BLESerial.h", "Neo/examples/Ble/BLESerial.cpp"} {
		p := filepath.Join(root, filepath.FromSlash(f))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte("\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	libs, err := indexLibs(root)
	if err != nil {
		t.Fatal(err)
	}
	l := libs["Neo"]
	if l == nil || len(l.srcs) != 2 || l.hdr["BLESerial.h"] {
		t.Fatalf("flat library sources %v headers %v", l.srcs, l.hdr)
	}
}

// With the build's macros known (rp2), includes in switched-off #ifdef blocks are not followed: Adafruit_NeoPixel.h
// includes <Adafruit_TinyUSB.h> only under #ifdef USE_TINYUSB.
func TestActiveIncludesFollowsIfdef(t *testing.T) {
	src := "#ifndef NEO_H\n#define NEO_H\n#include <Arduino.h>\n#ifdef USE_TINYUSB\n#include <Adafruit_TinyUSB.h>\n#endif\n" +
		"#if defined(ARDUINO_ARCH_RP2040)\n#include \"hardware/pio.h\"\n#elif defined(ESP32)\n#include <esp.h>\n#else\n#include <other.h>\n#endif\n" +
		"#if FOO > 2\n#include <maybe.h>\n#endif\n#ifdef NEO_H\n#include <self.h>\n#endif\n#endif\n"
	got := strings.Join(activeIncludes(src, map[string]bool{"ARDUINO_ARCH_RP2040": true}), ",")
	if got != "Arduino.h,hardware/pio.h,maybe.h,self.h" {
		t.Fatalf("got %s", got)
	}
	if all := strings.Join(activeIncludes(src, map[string]bool{"USE_TINYUSB": true}), ","); !strings.Contains(all, "Adafruit_TinyUSB.h") || !strings.Contains(all, "other.h") {
		t.Fatalf("with USE_TINYUSB: %s", all)
	}
}

// The rp2 recipe (arduino-pico): the linker script gets this board's layout, the core's unsigned
// Updater header is written without Python, and the ELF becomes a UF2 for the board's chip family.
func TestRp2Plan(t *testing.T) {
	root := t.TempDir()
	mk := func(rel, body string) string {
		t.Helper()
		p := filepath.Join(root, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
		return strings.ReplaceAll(p, `\`, `/`)
	}
	mk("core/cores/rp2040/main.cpp", "int main(){}\n")
	mk("core/variants/generic_rp2350/pins_arduino.h", "")
	mk("core/lib/rp2350/memmap_default.ld", "FLASH LENGTH = __FLASH_LENGTH__; RAM LENGTH = __RAM_LENGTH__; EE = __EEPROM_START__\n")
	gcc := filepath.Dir(mk("gcc/bin/arm-none-eabi-gcc.exe", "x"))
	for _, n := range []string{"g++", "ar", "size"} {
		mk("gcc/bin/arm-none-eabi-"+n+".exe", "x")
	}
	picotool := mk("picotool/picotool.exe", "x")
	board := `{"schema":"mujard-board.v1","id":"raspberrypi:rp2:t","short":"t","name":"t","vendor":"raspberrypi","family":"rp2",
	  "aliases":["rp2350usba"],"package":"pkg","core":"cores/rp2040","variant":"variants/generic_rp2350",
	  "vars":{"chip":"rp2350","toolchainOpts":"-mcpu=cortex-m33 -mthumb","boardDefine":"GENERIC_RP2350","fCpu":"150000000L",
	    "usbVid":"0x2e8a","usbPid":"0xf00f","usbManufacturer":"Generic","usbProduct":"RP2350","flashTotal":"2097152",
	    "flashLength":"2088960","eepromStart":"270524416","fsStart":"270524416","fsEnd":"270524416","ramLength":"512k",
	    "psramLength":"0x000000","boot2":"none","uf2Family":"--family rp2350-arm-s --abs-block"},
	  "defines":["__PICO_RP2350A=1"],"output":"uf2","flashMax":2088960,"ramMax":524288}`
	pkg := `{"schema":"mujard-package.v1","id":"pkg","family":"rp2","root":"` + strings.ReplaceAll(filepath.Join(root, "core"), `\`, `/`) +
		`","tools":{"gccBin":"` + strings.ReplaceAll(gcc, `\`, `/`) + `","picotool":"` + strings.TrimSuffix(picotool, ".exe") + `"}}`
	bdir, pdir := writeManifests(t, map[string]string{"t.json": board}, map[string]string{"p.json": pkg})
	reg, err := LoadRegistry(bdir, pdir)
	if err != nil {
		t.Fatal(err)
	}
	sketch := t.TempDir()
	if err := os.WriteFile(filepath.Join(sketch, "mujard_program.cpp"), []byte("void setup(){}\nvoid loop(){}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(t.TempDir(), "out")
	plan, err := buildPlan(reg, &Job{Schema: SchemaV1, Board: "rp2350usba", SketchDir: sketch, OutDir: out, TimeoutSec: 60, NoCache: true})
	if err != nil {
		t.Fatal(err)
	}
	ld, _ := os.ReadFile(filepath.Join(out, "memmap_default.ld"))
	if string(ld) != "FLASH LENGTH = 2088960; RAM LENGTH = 512k; EE = 270524416\n" {
		t.Fatalf("linker script not filled in: %q", ld)
	}
	if h, _ := os.ReadFile(filepath.Join(out, "core", "Updater_Signing.h")); string(h) != "#define ARDUINO_SIGNING 0\n" {
		t.Fatalf("Updater_Signing.h: %q", h)
	}
	steps := map[string]Step{}
	for _, s := range plan.Steps {
		steps[s.Name] = s
	}
	for _, want := range []string{"archive core", "boot2", "link", "uf2", "size"} {
		if _, ok := steps[want]; !ok {
			t.Fatalf("missing step %q", want)
		}
	}
	if uf2 := strings.Join(steps["uf2"].Args, " "); !strings.HasSuffix(uf2, "firmware.uf2 --family rp2350-arm-s --abs-block") || !strings.HasPrefix(uf2, "uf2 convert ") {
		t.Fatalf("uf2 step: %s", uf2)
	}
	link := strings.Join(steps["link"].Args, " ")
	for _, want := range []string{"libpico.a", "boot2.o", "-Wl,--script=" + filepath.Join(out, "memmap_default.ld"), "-D__PICO_RP2350A=1", "-mcpu=cortex-m33"} {
		if !strings.Contains(link, want) {
			t.Fatalf("link misses %q: %s", want, link)
		}
	}
	if b := strings.Join(steps["boot2"].Args, " "); !strings.Contains(b, filepath.Join("boot2", "rp2350", "none.S")) {
		t.Fatalf("boot2 step: %s", b)
	}

	// Without the core's chip files (pack not installed) the board is unsupported, so the CLI falls back.
	os.Remove(filepath.Join(root, "core", "lib", "rp2350", "memmap_default.ld"))
	_, err = buildPlan(reg, &Job{Schema: SchemaV1, Board: "rp2350usba", SketchDir: sketch, OutDir: out, TimeoutSec: 60, NoCache: true})
	var u *unsupported
	if !errors.As(err, &u) {
		t.Fatalf("want unsupported without the core files, got %v", err)
	}
}

func TestRp2SizeSections(t *testing.T) {
	out := ".boot2 0 0\n.text 50000 0\n.rodata 7000 0\n.ARM.exidx 8 0\n.data 1200 0\n.bss 10000 0\n.ram_vector_table 400 0\n"
	if flash, ram := parseSize(out, "rp2"); flash != 57008 || ram != 11600 {
		t.Fatalf("flash %d ram %d", flash, ram)
	}
}

func mustResolve(t *testing.T, reg *Registry, name string) *BoardManifest {
	t.Helper()
	b, err := reg.Resolve(name)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestUnsupportedBoard(t *testing.T) {
	board := `{"schema":"mujard-board.v1","id":"espressif:esp32:esp32s3","short":"esp32s3","name":"s3","vendor":"espressif","family":"esp32"}`
	bdir, pdir := writeManifests(t, map[string]string{"s.json": board}, nil)
	reg, err := LoadRegistry(bdir, pdir)
	if err != nil {
		t.Fatal(err)
	}
	_, err = buildPlan(reg, &Job{Schema: SchemaV1, Board: "esp32s3", SketchDir: t.TempDir(), OutDir: t.TempDir()})
	var u *unsupported
	if !errors.As(err, &u) {
		t.Fatalf("want unsupported, got %v", err)
	}
}
