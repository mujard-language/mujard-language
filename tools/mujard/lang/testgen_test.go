package lang

import (
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

// testCxx is the C++ compiler for tests on this PC, or skips the test.
func testCxx(t *testing.T) string {
	t.Helper()
	if _, err := os.Stat(`C:\tools\w64devkit\bin\g++.exe`); err == nil {
		return `C:\tools\w64devkit\bin\g++.exe`
	}
	if p := FindCxx(); p != "" {
		return p
	}
	if _, err := exec.LookPath("g++"); err != nil {
		t.Skip("no C++ compiler on this PC")
	}
	return "g++"
}

func runTests(t *testing.T, cxx, src string) []TestResult {
	t.Helper()
	sk, diags := CompileTests(src, "prog.muj", "")
	if e := FirstError(diags); e != nil {
		t.Fatalf("doesn't compile: %v", e)
	}
	res, err := RunTests(sk, cxx, 60*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

// TestProgramTests: 'test' blocks run on the PC against a simulated board (spec 12) - the sample's own tests pass, and
// a failing check, printed(), pretend on a button and an analog pin, and a program that halts all report what happened.
func TestProgramTests(t *testing.T) {
	cxx := testCxx(t)
	src, err := os.ReadFile(examplesDir + "/ok/26_tests.muj")
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range runTests(t, cxx, string(src)) {
		if !r.Passed {
			t.Errorf("26_tests.muj %q (line %d) failed: %s (line %d)\n%s", r.Name, r.Line, r.Message, r.FailLine, r.Output)
		}
	}

	prog := `board uno
lamp   = pin 12 as output
button = pin 2 as input pullup
light  = analog A0
var presses = 0

on button.pressed:
    lamp.toggle()
    presses += 1
    print "pressed {presses}"

every 100 ms:
    if light.volts > 4 V:
        print "bright"

at start:
    print "ready"

test "a press turns the lamp on":
    expect printed("ready")
    button.pretend(true)
    advance(50 ms)
    button.pretend(false)
    advance(50 ms)
    expect lamp.read(), "the lamp should be on"
    expect printed("pressed 1")

test "a wrong expectation fails with its message":
    advance(10 ms)
    expect lamp.read(), "the lamp is off, so this fails"

test "analog readings":
    light.pretend(4.5 V)
    advance(150 ms)
    expect printed("bright")
    expect not printed("dark")

test "nothing pressed":
    advance(1 s)
    expect presses == 0
`
	res := runTests(t, cxx, prog)
	if len(res) != 4 {
		t.Fatalf("4 tests expected, got %d", len(res))
	}
	if !res[0].Passed || !res[2].Passed || !res[3].Passed {
		t.Errorf("passing tests failed: %+v", res)
	}
	if r := res[1]; r.Passed || r.FailLine != 30 || r.Message != "the lamp is off, so this fails" {
		t.Errorf("failing test: %+v", r)
	}
	if !strings.Contains(res[0].Output, "pressed 1") {
		t.Errorf("the program's output comes back: %q", res[0].Output)
	}

	halts := "board uno\nvar d = 0\nevery 100 ms:\n    print 10 / d\ntest \"divides\":\n    advance(1 s)\n"
	if r := runTests(t, cxx, halts)[0]; r.Passed || !strings.Contains(r.Message, "the program halted") || !strings.Contains(r.Message, "Division by zero") {
		t.Errorf("a halt fails the test with the error: %+v", r)
	}

	// a firmware build leaves the tests out
	sk, _ := Compile(prog, "prog.muj")
	if cpp := sk.Files["mujard_program.cpp"]; strings.Contains(cpp, "mujard_Test_") || strings.Contains(cpp, "mujard_test.h") || sk.Files["mujard_test.cpp"] != "" {
		t.Error("tests must never reach the firmware")
	}
}

// TestDroppedEvents: the serial/UART read_line queues are bounded, so board.dropped_events counts a line that
// arrives when the queue is full (spec 6.5). It is not a pretend general event queue the runtime doesn't have.
func TestDroppedEvents(t *testing.T) {
	cxx := testCxx(t)
	src := `board uno
use mujard.uart
gps = Uart(rx: pin 4, tx: pin 3, speed: 9600 baud)

task serial_reader:
    wait until board.dropped_events > 0
    var line = serial.read_line()
    print line

task uart_reader:
    wait until board.dropped_events > 1
    var line = gps.read_line()
    print line

test "full input-line queues count their drops":
    serial.pretend("one")
    serial.pretend("two")
    serial.pretend("three")
    serial.pretend("four")
    serial.pretend("five")
    gps.pretend("u1")
    gps.pretend("u2")
    gps.pretend("u3")
    gps.pretend("u4")
    gps.pretend("u5")
    advance(100 ms)
    expect board.dropped_events == 2, "one line from each queue could not fit"
    expect printed "one", "the oldest serial line was kept"
    expect printed "u1", "the oldest UART line was kept"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("drop-count test: %+v", res)
	}
	sk, diags := Compile("board uno\nat start:\n    print board.dropped_events\n", "p.muj")
	if e := FirstError(diags); e != nil || sk == nil {
		t.Fatalf("board.dropped_events should compile: %v", diags)
	}
	if !strings.Contains(sk.Files["mujard_program.cpp"], "mujard::dropped_events") {
		t.Fatalf("board.dropped_events should reach the generated code:\n%s", sk.Files["mujard_program.cpp"])
	}
}

// TestBoardVoltage: board.voltage is optional. AVR, Pico/Pico W, nRF52840, SAMD21, STM32F103 and RA4M1 have a
// real supply-rail read; the rest honestly get none rather than a made-up number.
func TestBoardVoltage(t *testing.T) {
	if !boards["uno"].CanMeasureVoltage() || !boards["pico"].CanMeasureVoltage() || !boards["picow"].CanMeasureVoltage() ||
		!boards["nano33ble"].CanMeasureVoltage() || !boards["zero"].CanMeasureVoltage() || !boards["bluepill"].CanMeasureVoltage() ||
		!boards["unor4"].CanMeasureVoltage() {
		t.Error("AVR, Pico/Pico W, Nano 33 BLE, Zero, Blue Pill and Uno R4 can measure their supply")
	}
	for _, name := range []string{"esp32", "rp2350usba"} {
		if boards[name].CanMeasureVoltage() {
			t.Errorf("%s has no implemented supply-voltage path", name)
		}
	}
	cxx := testCxx(t)
	src := `board uno
var supply = 0 V

at start:
    var v = board.voltage
    if v != none:
        supply = v

test "board voltage reports the simulated supply":
    advance(10 ms)
    expect supply > 4.9 V and supply < 5.1 V, "the fake Uno rail is 5 V"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("board-voltage test: %+v", res)
	}
	for _, tc := range []struct{ board, want string }{
		{"uno", "mujard::board_voltage()"}, {"pico", "mujard::board_voltage()"},
		{"picow", "mujard::board_voltage()"}, {"nano33ble", "mujard::board_voltage()"},
		{"zero", "mujard::board_voltage()"}, {"bluepill", "mujard::board_voltage()"},
		{"unor4", "mujard::board_voltage()"}, {"esp32", "mujard::Opt<float>()"},
		{"rp2350usba", "mujard::Opt<float>()"},
	} {
		sk, diags := Compile("board "+tc.board+"\nat start:\n    var v = board.voltage\n", "p.muj")
		if e := FirstError(diags); e != nil || sk == nil || !strings.Contains(sk.Files["mujard_program.cpp"], tc.want) {
			t.Errorf("%s board.voltage should generate %q: %v", tc.board, tc.want, diags)
		}
		if tc.board == "nano33ble" && !strings.Contains(sk.Files["mujard_rt.cpp"], "SAADC_CH_PSELP_PSELP_VDD") {
			t.Error("nano33ble runtime has no SAADC VDD path")
		}
		if tc.board == "zero" && !strings.Contains(sk.Files["mujard_rt.cpp"], "ADC_INPUTCTRL_MUXPOS_SCALEDIOVCC") {
			t.Error("zero runtime has no SAMD scaled-IOVCC path")
		}
		if tc.board == "bluepill" && !strings.Contains(sk.Files["mujard_rt.cpp"], "analogRead(AVREF)") {
			t.Error("bluepill runtime has no STM32 VREFINT path")
		}
		if tc.board == "unor4" && !strings.Contains(sk.Files["mujard_rt.cpp"], "AVCC_MEASURE_PIN") {
			t.Error("unor4 runtime has no RA4M1 AVCC path")
		}
	}
	for _, board := range []string{"nano33ble", "zero", "bluepill", "unor4"} {
		if !strings.Contains(strings.Join(adcDefines(boards[board]), " "), "-DMUJARD_TEST_BOARD_VOLTAGE=3.3f") {
			t.Errorf("%s host test should simulate a 3.3 V rail", board)
		}
	}
}
