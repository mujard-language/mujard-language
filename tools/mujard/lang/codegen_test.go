package lang

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"
)

// laterPhase are the ok samples that use features of later phases (phone/cloud built-ins): the generator must refuse
// them with C320, never produce something half-working.
var laterPhase = map[string]bool{
	// none left since Phase 9 (2026-10-22): show, plot, notify, control and ask are built
}

func TestGenerateSamples(t *testing.T) {
	dir := filepath.Join("testdata", "cpp")
	generated := 0
	for _, s := range loadSamples(t, "ok") {
		sk, diags := CompileIn(s.src, s.name, s.dir)
		e := FirstError(diags)
		if laterPhase[s.name] {
			if e == nil || e.Code != "C320" {
				t.Errorf("%s: expected C320 (a later phase), got %v", s.name, diags)
			}
			continue
		}
		if e != nil {
			t.Errorf("%s: %s", s.name, e)
			continue
		}
		generated++
		got := sk.Files["mujard_program.cpp"]
		path := filepath.Join(dir, strings.TrimSuffix(s.name, ".muj")+".cpp")
		if *update {
			os.MkdirAll(dir, 0o755)
			if err := os.WriteFile(path, []byte(got), 0o644); err != nil {
				t.Fatal(err)
			}
			continue
		}
		want, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("%s: no golden C++ (run with -update and review it): %v", s.name, err)
			continue
		}
		if got != strings.ReplaceAll(string(want), "\r\n", "\n") {
			t.Errorf("%s: generated C++ changed - run with -update and review the diff", s.name)
		}
		if sk.FQBN == "" || sk.Files["sketch.ino"] == "" || sk.Files["mujard_rt.cpp"] == "" {
			t.Errorf("%s: incomplete sketch", s.name)
		}
	}
	if generated != 42 {
		t.Errorf("expected 42 samples to generate, got %d", generated)
	}
}

// TestModuleCodegen: 'import motors' emits the module's declarations in 'namespace motors' and qualifies
// the program's references as motors::name (spec 10.1).
func TestModuleCodegen(t *testing.T) {
	dir := t.TempDir()
	err := os.WriteFile(filepath.Join(dir, "motors.muj"),
		[]byte("board uno\nconst top_speed = 800\nfn faster(v: int) -> int:\n    return v + 1\nevery 9 s:\n    print 0\n"), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	sk, diags := CompileIn("board uno\nimport motors\nat start:\n    print motors.faster(motors.top_speed)\n", "main.muj", dir)
	if e := FirstError(diags); e != nil {
		t.Fatal(e)
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{
		"namespace motors {", "static int32_t faster(", "motors::faster(motors::top_speed)",
		`#line 3 "motors.muj"`, // the module's own file in the markers
	} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %q in\n%s", want, cpp)
		}
	}
	if strings.Contains(cpp, "mujard_every_") {
		t.Errorf("the module's 'every' block must not be generated for the importer:\n%s", cpp)
	}
}

func TestLineMarkers(t *testing.T) {
	sk, diags := Compile("board uno\nled = pin 13 as output\nevery 500 ms:\n    led.toggle()\n", "blink.muj")
	if FirstError(diags) != nil {
		t.Fatal(diags)
	}
	cpp := sk.Files["mujard_program.cpp"]
	// "#line 3" before the every block's 'if'; its body (source line 4) follows on the next line, so needs no marker.
	for _, want := range []string{`#line 2 "blink.muj"`, "static mujard::Output led(13, false);", `#line 3 "blink.muj"`,
		"    if (mujard_every_1.due()) {\n        led.toggle();", "static mujard::Every mujard_every_1(((mujard::Time)500000LL));"} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %q in\n%s", want, cpp)
		}
	}
	// after a mapped region, the markers point back at the generated file with the right line
	lines := strings.Split(cpp, "\n")
	for i, l := range lines {
		if strings.HasPrefix(l, `#line `) && strings.HasSuffix(l, `"mujard_program.cpp"`) {
			var n int
			for _, f := range strings.Fields(l)[1:2] {
				for _, c := range f {
					n = n*10 + int(c-'0')
				}
			}
			if n != i+2 {
				t.Errorf("line %d: %q should say %d", i+1, l, i+2)
			}
		}
	}
}

// TestExpectWarn: 'expect ..., halt: false' emits the non-halting expect_warn; with no message the
// condition's own text is reported (spec 6.8).
func TestExpectWarn(t *testing.T) {
	src := "board uno\nvar t = 1\nat start:\n    expect t > 0, halt: false\n    expect t < 5, \"fine\"\n    expect t != 2\n"
	sk, diags := Compile(src, "x.muj")
	if e := FirstError(diags); e != nil {
		t.Fatal(e)
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{
		`if (!((t > 0))) mujard::expect_warn("t > 0", 4);`,
		`if (!((t < 5))) mujard::fail("fine", 5);`,
		`if (!((t != 2))) mujard::fail("t != 2", 6);`,
	} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %q in\n%s", want, cpp)
		}
	}
}

// TestFastAtomics: a variable wider than a byte that a fast handler writes is updated by the rest
// of the program with interrupts off, and read through mujard::atomic_read (spec 6.5).
func TestFastAtomics(t *testing.T) {
	src := "board uno\n" +
		"enc = pin 2 as input pullup\n" +
		"var ticks = 0\n" +
		"var flag = false\n" +
		"on enc.changed fast:\n" +
		"    ticks += 1\n" +
		"    flag = true\n" +
		"at start:\n" +
		"    ticks = 0\n" +
		"    flag = false\n" +
		"every 1 s:\n" +
		"    print ticks\n" +
		"    if flag:\n" +
		"        print \"seen\"\n"
	sk, diags := Compile(src, "x.muj")
	if e := FirstError(diags); e != nil {
		t.Fatal(e)
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{
		"static volatile int32_t ticks;",
		"static volatile bool flag;",
		"atomic_read(ticks)", // the every-block read
		"if (flag)",          // a bool fits in one byte: reads stay plain
	} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %q in\n%s", want, cpp)
		}
	}
	for _, not := range []string{
		"atomic_read(flag)",
		"noInterrupts();\n        ticks += 1", // inside the interrupt itself the write stays plain
	} {
		if strings.Contains(cpp, not) {
			t.Errorf("didn't want %q in\n%s", not, cpp)
		}
	}
	// both plain-code writes of the 4-byte var are wrapped (the init in setup and 'at start')
	if n := strings.Count(cpp, "noInterrupts();"); n != 2 {
		t.Errorf("expected 2 noInterrupts() writes of ticks, got %d\n%s", n, cpp)
	}
}

// TestFastEdges: rising/falling follow the wire, pressed/released the active level (spec 6.5, 7), and every
// fast handler stops once the program halts and marks itself as an interrupt for the runtime.
func TestFastEdges(t *testing.T) {
	src := "board mega\n" +
		"a = pin 2 as input\n" +
		"b = pin 3 as input pullup\n" +
		"c = pin 18 as input pullup inverted\n" +
		"d = pin 19 as input pullup\n" +
		"var n = 0\n" +
		"on a.rising fast:\n    n += 1\n" +
		"on b.pressed fast:\n    n += 1\n" +
		"on c.pressed fast:\n    n += 1\n" +
		"on d.falling fast:\n    n += 1\n" +
		"every 1 s:\n    print n\n"
	sk, diags := Compile(src, "x.muj")
	if e := FirstError(diags); e != nil {
		t.Fatal(e)
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{
		"attachInterrupt(digitalPinToInterrupt(a.pin()), mujard_isr_1, RISING);",
		"attachInterrupt(digitalPinToInterrupt(b.pin()), mujard_isr_2, FALLING);", // pullup: pressed pulls the wire low
		"attachInterrupt(digitalPinToInterrupt(c.pin()), mujard_isr_3, RISING);",  // pullup inverted: pressed is high
		"attachInterrupt(digitalPinToInterrupt(d.pin()), mujard_isr_4, FALLING);", // falling is the wire, pullup or not
		"if (mujard::stopping()) return;",
		"mujard::IsrScope mujard_isr;",
	} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %q in\n%s", want, cpp)
		}
	}
}

// TestFrequency64: frequency is a signed 64-bit millihertz quantity (spec 5.1): literals, variables,
// every periods, tone(), 'n / time', 'n / freq', 'time * freq' and .number all keep their units.
func TestFrequency64(t *testing.T) {
	src := "board uno\n" +
		"buzzer = pin 8 as output\n" +
		"var f: frequency = 8 MHz\n" +
		"every 4 Hz:\n" +
		"    buzzer.toggle()\n" +
		"at start:\n" +
		"    buzzer.tone(440 Hz, 100 ms)\n" +
		"    var g = 1 / 500 ms\n" +
		"    var n = f / 2 Hz\n" +
		"    var c = 2 s * f\n" +
		"    print \"f={f} hz={f.number} g={g} c={c} n={n}\"\n"
	sk, diags := Compile(src, "x.muj")
	if e := FirstError(diags); e != nil {
		t.Fatal(e)
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{
		"static mujard::Freq f;",
		"f = ((mujard::Freq)8000000000LL);",
		"mujard::period_us(((mujard::Freq)4000LL))",
		"buzzer.tone(((mujard::Freq)440000LL), ((mujard::Time)100000LL))",
		"mujard::Freq g = ((mujard::Freq)((1) * 1000.0f / mujard::secs(((mujard::Time)500000LL))))",
		"float n = ((float)(f) / (float)(((mujard::Freq)2000LL)))",
		"float c = (mujard::secs(((mujard::Time)2000000LL)) * ((float)(f) / 1000.0f))",
		".add_freq(f)",
		".add_freq(g)",
		"((float)(f) / 1000.0f)",
	} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %q in\n%s", want, cpp)
		}
	}
}

// TestStoreOptionalLoad: store.load without a default takes its optional payload type from the declared target.
func TestStoreOptionalLoad(t *testing.T) {
	src := "board uno\n" +
		"at start:\n" +
		"    var n: int? = store.load(\"n\")\n" +
		"    var s: text? = store.load(\"s\")\n" +
		"    var ok = store.load(\"n\", 7)\n"
	sk, diags := Compile(src, "p.muj")
	if e := FirstError(diags); e != nil || sk == nil {
		t.Fatalf("optional store.load should compile: %v", diags)
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{
		`store_db.load("n", (int32_t*)nullptr)`,
		`store_db.load("s", (mujard::Text*)nullptr)`,
		`store_db.load("n", ((int32_t)(7)))`,
	} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %q in\n%s", want, cpp)
		}
	}
	_, diags = Compile("board uno\nat start:\n    var n = store.load(\"n\")\n", "p.muj")
	if e := FirstError(diags); e == nil || e.Code != "C320" {
		t.Errorf("untyped optional store.load should be C320: %v", diags)
	}
}

// TestBoardFreeMemory: board.free_memory reaches a real heap/stack backend on every catalog board, and future
// unsupported boards get C320 instead of a bogus -1.
func TestBoardFreeMemory(t *testing.T) {
	if (&Board{Name: "future", Arch: "unknown"}).CanMeasureFreeMemory() {
		t.Error("unknown board architectures must not claim a free-memory backend")
	}
	for name, b := range boards {
		sk, diags := Compile("board "+name+"\nat start:\n    print board.free_memory\n", "p.muj")
		if !b.CanMeasureFreeMemory() {
			if e := FirstError(diags); e == nil || e.Code != "C320" {
				t.Errorf("%s board.free_memory should be C320: %v", name, diags)
			}
			continue
		}
		if e := FirstError(diags); e != nil || sk == nil || !strings.Contains(sk.Files["mujard_program.cpp"], "mujard::free_memory()") {
			t.Errorf("%s board.free_memory: %v", name, diags)
		}
	}
	sk, _ := Compile("board uno\nat start:\n    print board.free_memory\n", "p.muj")
	rt := sk.Files["mujard_rt.cpp"]
	for _, want := range []string{"MUJARD_TEST_FREE_MEMORY", "ESP.getFreeHeap()", "rp2040.getFreeHeap()", "_sbrk(0)", "mbed_stats_heap_get(&s)", "sbrk(0)"} {
		if !strings.Contains(rt, want) {
			t.Errorf("runtime missing %q", want)
		}
	}
	if !strings.Contains(strings.Join(adcDefines(boards["uno"]), " "), "MUJARD_TEST_FREE_MEMORY=2048") {
		t.Error("host tests should get a simulated free-memory value")
	}
}

// TestBoardChip: board.chip is a compile-time catalog fact on every board. This keeps RP2350 and the ARM boards
// from falling through to a guessed AVR name.
func TestBoardChip(t *testing.T) {
	for name, b := range boards {
		sk, diags := Compile("board "+name+"\nat start:\n    print board.chip\n", "p.muj")
		want := fmt.Sprintf("mujard::Text(%q)", b.Family)
		if e := FirstError(diags); e != nil || sk == nil || !strings.Contains(sk.Files["mujard_program.cpp"], want) {
			t.Errorf("%s board.chip should emit %q: %v", name, want, diags)
		}
	}
}

// TestBoardID: boards with a hardware serial number get a real board.id; chips without an implemented backend
// keep the honest "unknown" result.
func TestBoardID(t *testing.T) {
	for _, tc := range []struct {
		board string
		paths []string
	}{
		{"pico", []string{"#elif defined(ARDUINO_ARCH_RP2040)", "pico_get_unique_board_id(&id)"}},
		{"picow", []string{"#elif defined(ARDUINO_ARCH_RP2040)", "pico_get_unique_board_id(&id)"}},
		{"rp2350usba", []string{"#elif defined(ARDUINO_ARCH_RP2040)", "pico_get_unique_board_id(&id)"}},
		{"bluepill", []string{"#elif defined(ARDUINO_ARCH_STM32)", "HAL_GetUIDw0()"}},
		{"zero", []string{"#elif defined(ARDUINO_ARCH_SAMD)", "0x0080A00C"}},
		{"nano33ble", []string{"ARDUINO_ARCH_MBED", "NRF_FICR->DEVICEID"}},
		{"unor4", []string{"ARDUINO_ARCH_RENESAS_UNO", "0x407FB1B0"}},
	} {
		sk, diags := Compile("board "+tc.board+"\nat start:\n    print board.id\n", "p.muj")
		if e := FirstError(diags); e != nil || sk == nil || !strings.Contains(sk.Files["mujard_program.cpp"], "mujard::board_id()") {
			t.Errorf("%s board.id: %v", tc.board, diags)
			continue
		}
		rt := sk.Files["mujard_rt.cpp"]
		for _, want := range tc.paths {
			if !strings.Contains(rt, want) {
				t.Errorf("%s runtime should contain %q", tc.board, want)
			}
		}
	}
}

// TestBoardRestart: board.restart() reaches the runtime on every catalog board, and the runtime has a reset path
// for the ARM cores too (not just the first four architectures).
func TestBoardRestart(t *testing.T) {
	for _, board := range []string{"unor4", "nano33ble", "zero", "bluepill"} {
		sk, diags := Compile("board "+board+"\nat start:\n    board.restart()\n", "p.muj")
		if e := FirstError(diags); e != nil || sk == nil || !strings.Contains(sk.Files["mujard_program.cpp"], "mujard::restart()") {
			t.Errorf("%s board.restart(): %v", board, diags)
			continue
		}
		if !strings.Contains(sk.Files["mujard_rt.cpp"], "NVIC_SystemReset()") {
			t.Errorf("%s runtime has no ARM reset path", board)
		}
	}
}

// TestPowerSleep: power.sleep/deep_sleep emit the runtime calls; a wake_on input event picks the
// edge matching the input's polarity (spec 6.10).
func TestPowerSleep(t *testing.T) {
	src := "board uno\n" +
		"btn = pin 2 as input pullup\n" +
		"every 1 s:\n" +
		"    power.sleep(5 s, wake_on: btn.pressed)\n" +
		"every 1 min:\n" +
		"    power.deep_sleep(30 min)\n"
	sk, diags := Compile(src, "x.muj")
	if e := FirstError(diags); e != nil {
		t.Fatal(e)
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{
		"mujard::power_sleep(((mujard::Time)5000000LL), btn.pin(), FALLING);",
		"mujard::power_deep_sleep(((mujard::Time)1800000000LL), -1, 0);",
	} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %q in\n%s", want, cpp)
		}
	}
}

func TestAfterKeepsLocalsOut(t *testing.T) {
	_, diags := CheckSource("board uno\nled = pin 13 as output\nevery 1 s:\n    var n = 3\n    after 2 s:\n        print n\n")
	e := FirstError(diags)
	if e == nil || e.Code != "C319" || e.Pos != (Pos{6, 15}) {
		t.Errorf("want C319 at 6:15, got %v", diags)
	}
}

// TestRuntimeOnHost builds and runs the runtime's own tests (runtime/tests) with a g++ on this PC, if there is one.
func TestRuntimeOnHost(t *testing.T) {
	gxx := "g++"
	if _, err := os.Stat(`C:\tools\w64devkit\bin\g++.exe`); err == nil {
		gxx = `C:\tools\w64devkit\bin\g++.exe`
	} else if _, err := exec.LookPath("g++"); err != nil {
		t.Skip("no g++ on this PC")
	}
	exe := filepath.Join(t.TempDir(), "runtime_test.exe")
	build := exec.Command(gxx, "-std=c++11", "-Wall", "-Wextra", "-O1", "-DMUJARD_HOST_TEST", "-Iruntime/tests/shim", "-Iruntime",
		"-o", exe, "runtime/tests/runtime_test.cpp", "runtime/mujard_rt.cpp")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}
	out, err := exec.Command(exe).CombinedOutput()
	if err != nil || !strings.Contains(string(out), " 0 failed") {
		t.Fatalf("runtime tests failed: %v\n%s", err, out)
	}
	t.Log(strings.TrimSpace(string(out[strings.LastIndex(string(out), "runtime_test"):])))
}

// TestBuildMatrix compiles generated sketches for real boards with arduino-cli. It is slow (minutes), so it only runs
// when MUJARD_ARDUINO_CLI points at arduino-cli (and MUJARD_ARDUINO_CONFIG at its config file):
//
//	$env:MUJARD_ARDUINO_CLI="C:\arduino-toolchain\bin\arduino-cli.exe"; $env:MUJARD_ARDUINO_CONFIG="C:\arduino-toolchain\arduino-cli.yaml"
//	go test ./lang -run TestBuildMatrix -v -timeout 60m
func TestBuildMatrix(t *testing.T) {
	cli := os.Getenv("MUJARD_ARDUINO_CLI")
	if cli == "" {
		t.Skip("set MUJARD_ARDUINO_CLI to run the arduino-cli build matrix")
	}
	cfg := os.Getenv("MUJARD_ARDUINO_CONFIG")
	only := os.Getenv("MUJARD_MATRIX_ONLY")          // optional: a sample name to build alone
	boardsOnly := os.Getenv("MUJARD_MATRIX_BOARDS")  // optional: "mega,esp8266" - only Blink and the dustbin on these boards
	samplesEnv := os.Getenv("MUJARD_MATRIX_SAMPLES") // optional with BOARDS: "03_dimmer.muj,..." instead of Blink+dustbin
	type job struct{ sample, board string }
	var jobs []job
	samples := loadSamples(t, "ok")
	for _, s := range samples {
		if boardsOnly != "" || laterPhase[s.name] || (only != "" && !strings.Contains(s.name, only)) {
			continue
		}
		jobs = append(jobs, job{s.name, ""})
	}
	// Blink and the dustbin on every board in the catalogue (the M1 milestone, then Phase 16's capability evidence).
	// A board whose core isn't installed is named as not proven, never skipped silently.
	installed := installedCores(t, cli, cfg)
	if only == "" {
		var list []string
		for name := range fqbns {
			if name != "uno" { // the samples themselves are Uno programs
				list = append(list, name)
			}
		}
		sort.Strings(list)
		if boardsOnly != "" {
			list = strings.Split(boardsOnly, ",")
		}
		for _, b := range list {
			if core := coreOf(fqbns[b]); !installed[core] {
				t.Logf("NOT PROVEN HERE - %s: the %s core isn't installed (arduino-cli core install %s)", b, core, core)
				continue
			}
			names := []string{"01_blink.muj", "06_dustbin.muj"}
			if samplesEnv != "" {
				names = strings.Split(samplesEnv, ",")
			}
			for _, n := range names {
				jobs = append(jobs, job{n, b})
			}
		}
	}
	evOut := os.Getenv("MUJARD_EVIDENCE_OUT")          // write the build half of capabilities.json from this run
	resume := os.Getenv("MUJARD_MATRIX_RESUME") == "1" // skip builds whose capabilities were all proven today
	built := map[string]map[string]bool{}
	failed := map[string]map[string]bool{} // attempted but the board build failed - its claim gets invalidated
	src := map[string]string{}
	srcDir := ""
	for _, s := range samples {
		src[s.name] = s.src
		srcDir = s.dir // every sample lives in the same folder (where 'import' looks for modules)
	}
	// The samples use Uno pins. On the ESP boards those are flash pins or missing (the checker refuses them, rightly:
	// B406 / B402), so the matrix uses the pins a person would pick there.
	pinsFor := map[string][][2]string{
		"esp32":   {{"pin LED", "pin 2"}, {"Servo(pin 6)", "Servo(pin 13)"}, {"trig: pin 9", "trig: pin 18"}, {"echo: pin 10", "echo: pin 19"}},
		"esp8266": {{"Servo(pin 6)", "Servo(pin D5)"}, {"trig: pin 9", "trig: pin D6"}, {"echo: pin 10", "echo: pin D7"}},
		// boards without an LED_BUILTIN name in the database
		"esp32s3": {{"pin LED", "pin 2"}}, "esp32s2": {{"pin LED", "pin 2"}}, "esp32c3": {{"pin LED", "pin 2"}},
		"esp32c6": {{"pin LED", "pin 2"}}, "rp2350usba": {{"pin LED", "pin 2"}},
	}
	for _, j := range jobs {
		code := src[j.sample]
		if j.board != "" {
			code = strings.Replace(code, "board uno", "board "+j.board, 1)
			for _, r := range pinsFor[j.board] {
				code = strings.ReplaceAll(code, r[0], r[1])
			}
		}
		name := j.sample + " on " + j.board
		sk, diags := CompileIn(code, j.sample, srcDir)
		if e := FirstError(diags); e != nil {
			t.Errorf("%s: %s", name, e)
			continue
		}
		var caps []string
		if f, _ := ParseSource(code); f != nil {
			if info, _ := AnalyzeIn(f, srcDir); info != nil {
				caps = info.Caps
			}
		}
		if resume && provenToday(evOut, sk.Board.Name, caps) {
			t.Logf("%s: already built today (MUJARD_MATRIX_RESUME)", name)
			markBuilt(built, sk.Board.Name, caps)
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
		if err != nil {
			t.Errorf("%s (%s): build failed: %v\n%s", name, sk.FQBN, err, tail(string(out), 40))
			markBuilt(failed, sk.Board.Name, caps) // attempted and failed - invalidate only these caps, not the board
			continue
		}
		t.Logf("%s (%s): %s", name, sk.FQBN, lastLine(string(out), "Sketch uses"))
		markBuilt(built, sk.Board.Name, caps)
		if evOut != "" { // saved after every build, so an interrupted run keeps what it proved
			writeBuildEvidence(t, evOut, map[string]map[string]bool{sk.Board.Name: set(caps)}, nil, false)
		}
	}
	if evOut != "" && only == "" && boardsOnly == "" {
		writeBuildEvidence(t, evOut, built, failed, true) // a whole run merges; attempted+failed loses its claim
	}
}

func set(list []string) map[string]bool {
	m := map[string]bool{}
	for _, s := range list {
		m[s] = true
	}
	return m
}

func markBuilt(built map[string]map[string]bool, board string, caps []string) {
	if built[board] == nil {
		built[board] = map[string]bool{}
	}
	for _, cp := range caps {
		built[board][cp] = true
	}
}

// provenToday: every capability of this build already has today's date for the board in the evidence file.
func provenToday(path, board string, caps []string) bool {
	b, err := os.ReadFile(path)
	if err != nil || path == "" {
		return false
	}
	var ev Evidence
	if json.Unmarshal(b, &ev) != nil {
		return false
	}
	today := time.Now().Format("2006-01-02")
	for _, cp := range caps {
		if ev.Build[board][cp] != today {
			return false
		}
	}
	return len(caps) > 0
}

// writeBuildEvidence puts this matrix run's builds into capabilities.json. A passed build dates the capabilities it
// actually exercised; a failed build removes only the capabilities it was meant to prove (an attempted build that
// failed can't keep its old claim). A capability that was never exercised this run keeps whatever date it already
// had - "not attempted" is not "failed". The hardware half is never touched: it is written by hand from real boards.
func writeBuildEvidence(t *testing.T, path string, built, failed map[string]map[string]bool, whole bool) {
	t.Helper()
	var ev Evidence
	if b, err := os.ReadFile(path); err == nil {
		if err := json.Unmarshal(b, &ev); err != nil {
			t.Fatalf("%s: %v", path, err)
		}
	}
	if ev.Build == nil {
		ev.Build = map[string]map[string]string{}
	}
	today := time.Now().Format("2006-01-02")
	for board, caps := range built {
		if ev.Build[board] == nil {
			ev.Build[board] = map[string]string{}
		}
		for cp := range caps {
			ev.Build[board][cp] = today
		}
	}
	for board, caps := range failed {
		for cp := range caps {
			delete(ev.Build[board], cp)
		}
	}
	out, _ := json.MarshalIndent(ev, "", "  ")
	if err := os.WriteFile(path, append(out, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Logf("wrote the build evidence of %d boards to %s (whole run: %v)", len(built), path, whole)
}

// installedCores asks arduino-cli which cores (vendor:arch) are installed.
func installedCores(t *testing.T, cli, cfg string) map[string]bool {
	args := []string{"core", "list", "--format", "json"}
	if cfg != "" {
		args = append([]string{"--config-file", cfg}, args...)
	}
	out, err := exec.Command(cli, args...).Output()
	if err != nil {
		t.Fatalf("arduino-cli core list: %v", err)
	}
	var list struct {
		Platforms []struct {
			ID string `json:"id"`
		} `json:"platforms"`
	}
	if err := json.Unmarshal(out, &list); err != nil {
		t.Fatalf("arduino-cli core list: %v", err)
	}
	have := map[string]bool{}
	for _, p := range list.Platforms {
		have[p.ID] = true
	}
	return have
}

// coreOf is the core of an FQBN: "esp32:esp32:esp32s3" -> "esp32:esp32".
func coreOf(fqbn string) string {
	p := strings.SplitN(fqbn, ":", 3)
	if len(p) < 2 {
		return fqbn
	}
	return p[0] + ":" + p[1]
}

// writeTestSecretsHeader writes the header every real build gets (the server always writes one, empty when the
// program has no secrets; the runtime includes it plainly), with placeholders for the secrets, like testgen.
func writeTestSecretsHeader(dir string, secrets []string) {
	var h strings.Builder
	h.WriteString("#pragma once\n")
	for _, n := range secrets {
		fmt.Fprintf(&h, "#define %s \"test-%s\"\n", n, strings.ToLower(n))
	}
	os.WriteFile(filepath.Join(dir, secretsHeader), []byte(h.String()), 0o644)
}

func tail(s string, n int) string {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	if len(lines) > n {
		lines = lines[len(lines)-n:]
	}
	return strings.Join(lines, "\n")
}

func lastLine(s, prefix string) string {
	for _, l := range strings.Split(s, "\n") {
		if strings.HasPrefix(strings.TrimSpace(l), prefix) {
			return strings.TrimSpace(l)
		}
	}
	return "built"
}

// TestStoreDurabilityStatus: store.save is async - the program reads store.pending / store.failed /
// store.committed / store.last_commit to tell a real commit from a silent retry (v58). store.pending
// also covers a saved var whose live value isn't in storage yet (v66).
func TestStoreDurabilityStatus(t *testing.T) {
	src := "board uno\nat start:\n    store.save(\"k\", 7)\n    print \"{store.pending} {store.failed} {store.committed} {store.last_commit}\"\n"
	sk, diags := Compile(src, "x.muj")
	if FirstError(diags) != nil {
		t.Fatalf("store durability fields should compile: %v", diags)
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{"store_db.pending()", "saved_pending()", "store_db.failed()", "store_db.committed()", "store_db.last_commit()"} {
		if !strings.Contains(cpp, want) {
			t.Errorf("generated C++ should call %s", want)
		}
	}
}
