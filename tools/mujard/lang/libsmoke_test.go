package lang

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// Phase 8's exit: every shipped library description has a program in the build matrix, so its checked calls are proven
// against the real library's C++, not just against the description. libSmoke writes that program from the description
// itself: each class made through its described constructor (or used as the library's ready-made object), each method
// and function called with arguments of the described types, each plain constant read. The fast test below checks every
// one; TestEveryLibraryBuilds builds every one with arduino-cli against the installed libraries. A library added later is
// covered with no new test code.

// libSmokeBoard is where a library's program is built: a Mega for anything that runs on AVR (room for the big ones),
// otherwise the first family the library names.
func libSmokeBoard(d *LibDesc) string {
	if len(d.Architectures) == 0 {
		return "mega"
	}
	for _, a := range d.Architectures {
		if a == "*" || a == "avr" {
			return "mega"
		}
	}
	for _, want := range []struct{ arch, board string }{{"esp32", "esp32"}, {"esp8266", "esp8266"}, {"rp2040", "pico"}} {
		for _, a := range d.Architectures {
			if a == want.arch {
				return want.board
			}
		}
	}
	return "mega"
}

// smokeArg is a value of a described type, or "" when the type can't be written as a plain literal. "any" is left out
// too: it can be a pointer in C++ (a bitmap, a font), which no plain value fits.
func smokeArg(typ string) string {
	switch typ {
	case "int":
		return "2"
	case "float":
		return "1.5"
	case "bool":
		return "false"
	case "text":
		return `"x"`
	}
	return ""
}

func smokeArgs(fn *LibFn) (string, bool) {
	var args []string
	for _, p := range fn.Params {
		if p.Optional {
			break // the described form every caller can write: required arguments only
		}
		a := smokeArg(p.Type)
		if a == "" {
			return "", false
		}
		args = append(args, a)
	}
	return strings.Join(args, ", "), true
}

var smokePrintable = map[string]bool{"int": true, "float": true, "bool": true, "text": true}

// libSmoke is the program for d, and how many described calls it had to leave out (arguments of a class or list type).
func libSmoke(d *LibDesc) (src string, skipped int) {
	var top, body strings.Builder
	fmt.Fprintf(&top, "board %s\nuse %q\n", libSmokeBoard(d), d.Name)
	call := func(target, name string, fn *LibFn) {
		args, ok := smokeArgs(fn)
		if !ok {
			skipped++
			return
		}
		expr := fmt.Sprintf("%s%s(%s)", target, name, args)
		if smokePrintable[fn.Result] {
			fmt.Fprintf(&body, "    print %s\n", expr)
		} else {
			fmt.Fprintf(&body, "    %s\n", expr)
		}
	}
	classes := make([]string, 0, len(d.Classes))
	for name := range d.Classes {
		classes = append(classes, name)
	}
	sort.Strings(classes)
	n := 0
	for _, cname := range classes {
		c := d.Classes[cname]
		if c.Abstract || c.Value || d.Families[cname] != nil {
			continue // only extended, or only handed out by the library; a family's members are tested by TestU8g2Family
		}
		target := cname + "." // no constructor: the library's ready-made object
		if len(c.Ctors) > 0 {
			args, ok := smokeArgs(&c.Ctors[0])
			if !ok {
				skipped++
				continue
			}
			n++
			obj := fmt.Sprintf("o%d", n)
			fmt.Fprintf(&top, "%s = %s(%s)\n", obj, cname, args)
			target = obj + "."
		}
		// its own methods and the ones it inherits through 'extends' (its own win, as in C++)
		methods := map[string]*LibFn{}
		seen := map[string]bool{}
		for at := cname; d.Classes[at] != nil && !seen[at]; at = d.Classes[at].Extends {
			seen[at] = true
			for m, f := range d.Classes[at].Methods {
				if methods[m] == nil {
					methods[m] = f
				}
			}
		}
		names := make([]string, 0, len(methods))
		for m := range methods {
			names = append(names, m)
		}
		sort.Strings(names)
		for _, m := range names {
			call(target, m, methods[m])
		}
	}
	fns := make([]string, 0, len(d.Functions))
	for f := range d.Functions {
		fns = append(fns, f)
	}
	sort.Strings(fns)
	for _, f := range fns {
		call("", f, d.Functions[f])
	}
	consts := make([]string, 0, len(d.Constants))
	for k, typ := range d.Constants {
		if smokePrintable[typ] {
			consts = append(consts, k)
		}
	}
	sort.Strings(consts)
	for _, k := range consts {
		if i := strings.LastIndex(k, "::"); i >= 0 {
			k = k[i+2:] // a class's constant: the program writes its own name (AccelStepper::DRIVER is DRIVER)
		}
		fmt.Fprintf(&body, "    print %s\n", k)
	}
	if body.Len() == 0 {
		body.WriteString("    print \"no plain calls\"\n")
	}
	return top.String() + "at start:\n" + body.String(), skipped
}

// Found by the smoke programs: Mujard's own Servo (mujard.servo) shadowed the Arduino Servo library's class even in a
// program that never used mujard.servo, so 'use "Servo"' then Servo() was C318 and generated Mujard's servo device.
func TestCppLibraryWinsOverAnUnusedBuiltin(t *testing.T) {
	sk, diags := Compile("board mega\nuse \"Servo\"\narm = Servo()\nat start:\n    print arm.attach(9)\n    arm.write(90)\n", "s.muj")
	if e := FirstError(diags); e != nil || sk == nil {
		t.Fatalf("use \"Servo\" then Servo() should check clean: %v", diags)
	}
	if cpp := sk.Files["mujard_program.cpp"]; !strings.Contains(cpp, "static Servo arm;") || strings.Contains(cpp, "ServoMotor") {
		t.Errorf("Servo() should be the Arduino library's object:\n%s", cpp)
	}
	// With 'use mujard.servo' it is Mujard's own, and with neither it is still C318.
	sk, _ = Compile("board uno\nuse mujard.servo\nlid = Servo(pin 6)\n", "s.muj")
	if sk == nil || !strings.Contains(sk.Files["mujard_program.cpp"], "mujard::ServoMotor lid") {
		t.Error("use mujard.servo: Servo() stays Mujard's servo")
	}
	if _, d := CheckSourceModules("board uno\nlid = Servo(pin 6)\n", nil); !hasDiag(d, "C318") {
		t.Error("Servo() with no use line is C318")
	}
}

// Found by the library builds (2026-10-29): AccelStepper's interface constants live inside the class, MultiStepper in
// its own header. A description can now say both, and what it makes the sketch #include is checked.
func TestDescriptionClassConstantsAndHeaders(t *testing.T) {
	sk, diags := Compile("board mega\nuse \"AccelStepper\"\nm = AccelStepper(DRIVER, 2, 3)\nat start:\n    m.setMaxSpeed(500.0)\n", "s.muj")
	if e := FirstError(diags); e != nil || sk == nil {
		t.Fatalf("AccelStepper(DRIVER, ...) should check clean: %v", diags)
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{"AccelStepper::DRIVER", "#include <AccelStepper.h>", "#include <MultiStepper.h>"} {
		if !strings.Contains(cpp, want) {
			t.Errorf("generated C++ lacks %q:\n%s", want, cpp)
		}
	}
	for _, bad := range []string{"../../mujard-private/config.php", "C:/x.h", "/etc/x.h", "a/../b.h", "x.cpp", "sub\\x.h"} {
		if _, err := ParseLibDesc(`{"name": "X", "header": ` + strconv.Quote(bad) + `}`); err == nil {
			t.Errorf("header %q must be refused", bad)
		}
		if _, err := ParseLibDesc(`{"name": "X", "header": "x.h", "headers": [` + strconv.Quote(bad) + `]}`); err == nil {
			t.Errorf("extra header %q must be refused", bad)
		}
	}
	if _, err := ParseLibDesc(`{"name": "X", "header": "utility/twi.h", "headers": ["Adafruit_GFX.h"]}`); err != nil {
		t.Errorf("plain library headers are fine: %v", err)
	}
}

func TestEveryLibrarySmokeChecks(t *testing.T) {
	for _, d := range builtinDescs {
		src, skipped := libSmoke(d)
		_, diags := CheckSourceModules(src, nil)
		if e := FirstError(diags); e != nil {
			t.Errorf("%s: %s\n%s", d.Name, e, src)
			continue
		}
		if skipped > 0 {
			t.Logf("%s: %d described calls take a class or list argument and are left to the hand-written programs", d.Name, skipped)
		}
	}
}

// TestEveryLibraryBuilds builds each library's program with arduino-cli. Gated like TestBuildMatrix:
//
//	$env:MUJARD_ARDUINO_CLI="C:\arduino-toolchain\bin\arduino-cli.exe"; $env:MUJARD_ARDUINO_CONFIG="C:\arduino-toolchain\arduino-cli.yaml"
//	go test ./lang -run TestEveryLibraryBuilds -v -timeout 300m   (MUJARD_LIB_ONLY=<name part>[,<name part>...] builds those)
func TestEveryLibraryBuilds(t *testing.T) {
	cli := os.Getenv("MUJARD_ARDUINO_CLI")
	if cli == "" {
		t.Skip("set MUJARD_ARDUINO_CLI to build every library's program with arduino-cli")
	}
	cfg := os.Getenv("MUJARD_ARDUINO_CONFIG")
	only := strings.ToLower(os.Getenv("MUJARD_LIB_ONLY"))
	var notInstalled []string
	defer func() {
		if len(notInstalled) > 0 {
			t.Logf("NOT PROVEN HERE - not installed on this machine: %s (arduino-cli lib install them, then run again)", strings.Join(notInstalled, ", "))
		}
	}()
	wanted := func(name string) bool {
		if only == "" {
			return true
		}
		for _, part := range strings.Split(only, ",") {
			if part = strings.TrimSpace(part); part != "" && strings.Contains(strings.ToLower(name), part) {
				return true
			}
		}
		return false
	}
	for _, d := range builtinDescs {
		if !wanted(d.Name) {
			continue
		}
		src, _ := libSmoke(d)
		sk, diags := Compile(src, "lib.muj")
		if e := FirstError(diags); e != nil || sk == nil {
			t.Errorf("%s: %v", d.Name, e)
			continue
		}
		dir := filepath.Join(t.TempDir(), "sketch")
		os.MkdirAll(dir, 0o755)
		for f, content := range sk.Files {
			os.WriteFile(filepath.Join(dir, f), []byte(content), 0o644)
		}
		writeTestSecretsHeader(dir, sk.Secrets)
		args := []string{"compile", "--fqbn", sk.FQBN, "--warnings", "default", dir}
		if cfg != "" {
			args = append([]string{"--config-file", cfg}, args...)
		}
		out, err := exec.Command(cli, args...).CombinedOutput()
		if err != nil && d.Header != "" && strings.Contains(string(out), d.Header+": No such file or directory") {
			// The library isn't installed on this machine (the build server installs declared libraries itself): its
			// description is unproven here, which is said, not counted as a pass or as a wrong description.
			notInstalled = append(notInstalled, d.Name)
			continue
		}
		if err != nil {
			// Every compiler error, wherever it is in the output (a long build log pushes the first one out of a tail).
			var errs []string
			for _, l := range strings.Split(string(out), "\n") {
				if strings.Contains(l, "error:") || strings.Contains(l, "undefined reference") || strings.Contains(l, "multiple definition") {
					errs = append(errs, strings.TrimSpace(l))
				}
			}
			if len(errs) > 20 {
				errs = errs[:20]
			}
			t.Errorf("%s (%s): build failed: %v\n%s\nerrors:\n%s\n%s", d.Name, sk.FQBN, err, src, strings.Join(errs, "\n"), tail(string(out), 10))
			continue
		}
		t.Logf("%s (%s): %s", d.Name, sk.FQBN, lastLine(string(out), "Sketch uses"))
	}
}
