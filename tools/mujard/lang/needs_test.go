package lang

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestNeeds: a driver module says what it asks of the board (spec 10.5); the board is checked against it before any
// C++ is made, in the module's own words.
func TestNeeds(t *testing.T) {
	light := "needs i2c\nuse mujard.i2c\n\nvar measure: list[int, 1] = [0x10]\n\nfn lux() -> int:\n    i2c.write(0x23, measure)\n    var raw = i2c.read(0x23, 2)\n" +
		"    if raw == none:\n        return 0\n    return (raw[0] * 256 + raw[1]) * 10 / 12\n"
	weather := "needs wifi, http\n\nfn ping() -> bool:\n    return wifi.connected\n"
	find := func(diags []Diagnostic, code string) *Diagnostic {
		for i := range diags {
			if diags[i].Code == code {
				return &diags[i]
			}
		}
		return nil
	}
	// the spec's own example, imported onto an Uno: checks clean
	_, diags := CheckSourceModules("board uno\nimport light\nat start:\n    print light.lux()\n", map[string]string{"light.muj": light})
	if e := FirstError(diags); e != nil {
		t.Errorf("the I2C driver on an Uno: %v", e)
	}
	// a Wi-Fi driver onto an Uno: B415 in the module's file, naming the module
	_, diags = CheckSourceModules("board uno\nimport weather\nat start:\n    print weather.ping()\n", map[string]string{"weather.muj": weather})
	d := find(diags, "B415")
	if d == nil || d.File != "weather.muj" || d.Pos != (Pos{1, 7}) || d.Message != "The module 'weather' needs wifi, and the Arduino Uno doesn't have it." {
		t.Errorf("a Wi-Fi driver on an Uno: %v", diags)
	}
	// ... and onto an ESP32 it's fine
	_, diags = CheckSourceModules("board esp32\nimport weather\nat start:\n    print weather.ping()\n", map[string]string{"weather.muj": weather})
	if d := find(diags, "B415"); d != nil {
		t.Errorf("a Wi-Fi driver on an ESP32: %v", d)
	}
	// unknown capability names
	_, diags = CheckSource("board uno\nneeds i2c, warp_drive\n")
	if d := find(diags, "C330"); d == nil || d.Pos != (Pos{2, 12}) || !strings.Contains(d.Message, "'warp_drive' isn't a capability") || !strings.Contains(d.Message, "i2c, spi") {
		t.Errorf("unknown capability: %v", diags)
	}
	// a program's own needs line names the program
	_, diags = CheckSource("board uno\nneeds mqtt\n")
	if d := find(diags, "B415"); d == nil || !strings.HasPrefix(d.Message, "This program needs mqtt") {
		t.Errorf("program needs: %v", diags)
	}
	// a device called needs is still a device
	if _, diags := CheckSource("board uno\nneeds = pin 3 as output\nat start:\n    needs.on()\n"); FirstError(diags) != nil {
		t.Errorf("a device called needs: %v", FirstError(diags))
	}
	// what a module needs counts as what the program asks of the board (W618 lists it)
	f, _ := ParseSource("board uno\nimport light\nat start:\n    print light.lux()\n")
	info, _ := analyzeWith(f, moduleLoaderMap(map[string]string{"light.muj": light}), nil)
	found := false
	for _, cp := range info.Caps {
		found = found || cp == "i2c"
	}
	if !found {
		t.Errorf("the module's needs are the program's capabilities: %v", info.Caps)
	}
	// the tree shows it
	pf, _ := ParseSource("board uno\nneeds i2c, spi\n")
	if s := Dump(pf); !strings.Contains(s, "(needs i2c spi)") {
		t.Errorf("tree: %s", s)
	}
}

// TestDriverPins: a driver module asks for pins ('needs pin trig'), the program gives them
// ('import sonar(trig: pin 7, echo: pin 8)'), and the module is checked and generated with the real pins.
func TestDriverPins(t *testing.T) {
	sonar := "needs pin trig, pin echo\n\nt = pin trig as output\ne = pin echo as input\n\nfn ping():\n    t.on()\n    t.off()\n"
	mods := map[string]string{"sonar.muj": sonar}
	codes := func(src string) []Diagnostic {
		_, diags := CheckSourceModules(src, mods)
		var out []Diagnostic
		for _, d := range diags {
			if !d.IsWarning() {
				out = append(out, d)
			}
		}
		return out
	}
	if e := codes("board uno\nimport sonar(trig: pin 7, echo: pin 8)\nat start:\n    sonar.ping()\n"); len(e) != 0 {
		t.Errorf("both pins given: %v", e)
	}
	e := codes("board uno\nimport sonar(trig: pin 7)\nat start:\n    sonar.ping()\n")
	if len(e) != 1 || e[0].Code != "C331" || e[0].Message != "The module 'sonar' needs pin 'echo'. Say which: import sonar(echo: pin N)." {
		t.Errorf("a missing pin: %v", e)
	}
	e = codes("board uno\nimport sonar(trig: pin 7, echo: pin 8, led: pin 13)\nat start:\n    sonar.ping()\n")
	if len(e) != 1 || e[0].Code != "C332" || e[0].Pos != (Pos{2, 40}) || !strings.HasSuffix(e[0].Message, "The pins it needs: trig, echo.") {
		t.Errorf("a pin the module doesn't ask for: %v", e)
	}
	// the real pin meets the board's rules and the program's other pins
	e = codes("board uno\nled = pin 7 as output\nimport sonar(trig: pin 7, echo: pin 8)\nat start:\n    sonar.ping()\n    led.on()\n")
	if len(e) != 1 || e[0].Code != "B407" || e[0].File != "sonar.muj" {
		t.Errorf("the driver's pin 7 is the program's led: %v", e)
	}
	e = codes("board uno\nimport sonar(trig: pin 99, echo: pin 8)\nat start:\n    sonar.ping()\n")
	if len(e) != 1 || e[0].Code != "B402" {
		t.Errorf("an Uno has no pin 99: %v", e)
	}
	// 'needs pin' belongs in a driver, not a program
	if _, diags := CheckSource("board uno\nneeds pin x\n"); FirstError(diags) == nil || FirstError(diags).Code != "C333" {
		t.Errorf("needs pin in a program: %v", diags)
	}
	// the generated C++ drives pin 7 and reads pin 8
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "sonar.muj"), []byte(sonar), 0o644); err != nil {
		t.Fatal(err)
	}
	sk, diags := CompileIn("board uno\nimport sonar(trig: pin 7, echo: pin 8)\nat start:\n    sonar.ping()\n", "main.muj", dir)
	if e := FirstError(diags); e != nil {
		t.Fatalf("compile: %v", e)
	}
	prog := sk.Files["mujard_program.cpp"]
	if !strings.Contains(prog, "Output t(7,") || !strings.Contains(prog, "Input e(8,") || strings.Contains(prog, "trig") {
		t.Errorf("generated pins: want t on 7, e on 8 and no 'trig' left:\n%s", prog)
	}
}
