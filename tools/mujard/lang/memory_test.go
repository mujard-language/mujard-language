package lang

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// Plan Phase 14, "verify memory behaviour": the RAM estimate behind W614 counts each thing a program declares with a
// size from the tables in memory.go. This builds a program with one of each for the Uno and reads the real size of
// every object from the ELF (avr-nm), so the tables can't drift from the generated C++ and the runtime (as they did
// when direct pin access added a register pointer and a mask to every output and input). Gated like TestBuildMatrix.
func TestRAMEstimateMatchesTheBuild(t *testing.T) {
	cli := os.Getenv("MUJARD_ARDUINO_CLI")
	if cli == "" {
		t.Skip("set MUJARD_ARDUINO_CLI to compare the RAM estimate with a real build")
	}
	cfg := os.Getenv("MUJARD_ARDUINO_CONFIG")
	nm := avrNm(cfg)
	if nm == "" {
		t.Skip("no avr-nm next to the installed AVR core")
	}
	src := `board uno
led = pin 13 as output
button = pin 2 as input pullup
lamp = pin 9 as pwm
knob = analog A0
var n = 0
var level = 1.5
var lit = false
var label = "hi"
var pause = 2 s
every 100 ms:
    n += 1
after 1 s:
    led.on()
when n > 50:
    print knob.value
task blinker:
    wait pause
    print n
on button.pressed:
    lamp.write(50%)
    lit = not lit
    print label
    print lit
    print level
`
	sk, diags := Compile(src, "sketch.muj")
	if e := FirstError(diags); e != nil || sk == nil {
		t.Fatalf("the program should compile: %v", diags)
	}
	dir := filepath.Join(t.TempDir(), "sketch")
	build := filepath.Join(t.TempDir(), "build")
	os.MkdirAll(dir, 0o755)
	for f, content := range sk.Files {
		os.WriteFile(filepath.Join(dir, f), []byte(content), 0o644)
	}
	writeTestSecretsHeader(dir, sk.Secrets)
	args := []string{"compile", "--fqbn", sk.FQBN, "--build-path", build, dir}
	if cfg != "" {
		args = append([]string{"--config-file", cfg}, args...)
	}
	if out, err := exec.Command(cli, args...).CombinedOutput(); err != nil {
		t.Fatalf("build failed: %v\n%s", err, tail(string(out), 30))
	}
	out, err := exec.Command(nm, "-S", "--defined-only", filepath.Join(build, "sketch.ino.elf")).Output()
	if err != nil {
		t.Fatalf("avr-nm: %v", err)
	}
	// _ZN11mujard_userL5countE: the program's own objects, by name.
	sym := regexp.MustCompile(`(?m)^[0-9a-fA-F]{8} ([0-9a-fA-F]{8}) [BbDd] _ZN11mujard_userL?(\d{1,2})([A-Za-z_][A-Za-z0-9_]*)E\s*$`)
	sizes := map[string]int{}
	for _, m := range sym.FindAllStringSubmatch(string(out), -1) {
		if n, _ := strconv.Atoi(m[2]); n == len(m[3]) {
			size, _ := strconv.ParseInt(m[1], 16, 32)
			sizes[m[3]] = int(size)
		}
	}
	var names []string
	for n := range sizes {
		names = append(names, n)
	}
	sort.Strings(names)
	t.Logf("objects in the build: %v", names)

	want := map[string]int{
		"led": deviceRAM["output"], "button": deviceRAM["input"], "lamp": deviceRAM["pwm"], "knob": deviceRAM["analog"],
		"n": 4, "level": 4, "lit": 1, "label": textMax + 2, "pause": 8,
	}
	for name, w := range want {
		got, ok := sizes[name]
		if !ok {
			t.Errorf("%s isn't in the build's symbols (renamed or optimised away?)", name)
			continue
		}
		if got != w {
			t.Errorf("%s: the estimate counts %d bytes, the build has %d", name, w, got)
		}
	}
	// The scheduler's own objects for each block kind.
	for prefix, w := range map[string]int{"mujard_every_": 17, "mujard_after_": 9, "mujard_when_": 2, "mujard_task_": 11} {
		found := false
		for name, got := range sizes {
			if strings.HasPrefix(name, prefix) {
				found = true
				if got != w {
					t.Errorf("%s: the estimate counts %d bytes, the build has %d", name, w, got)
				}
			}
		}
		if !found {
			t.Errorf("no %s* object in the build: %v", prefix, names)
		}
	}
}

// avrNm is the avr-nm of the AVR toolchain arduino-cli installed (its config's data folder), or "".
func avrNm(cfg string) string {
	data := ""
	if b, err := os.ReadFile(cfg); err == nil {
		if m := regexp.MustCompile(`(?m)^\s*data:\s*(.+?)\s*$`).FindStringSubmatch(string(b)); m != nil {
			data = strings.Trim(m[1], `"'`)
		}
	}
	if data == "" {
		return ""
	}
	matches, _ := filepath.Glob(filepath.Join(data, "packages", "arduino", "tools", "avr-gcc", "*", "bin", "avr-nm*"))
	sort.Strings(matches)
	if len(matches) == 0 {
		return ""
	}
	return matches[len(matches)-1]
}
