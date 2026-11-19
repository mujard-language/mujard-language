package lang

import (
	"strings"
	"testing"
)

// TestCoreSharedWrites (W620): a global a secondary-core block touches while the main core touches it too has no
// synchronization - reads can tear and compound writes race. A global used by only one side doesn't warn.
func TestCoreSharedWrites(t *testing.T) {
	shared := `board pico
var n = 0

every 100 ms on core 1:
    n = n + 1

every 1 s:
    print n
`
	if d := check(t, shared); !hasDiag(d, "W620") {
		t.Errorf("a core-1 write used on the main core should warn W620: %v", d)
	}
	// the same hazard read-side: the core reads what the main core writes
	readSide := `board pico
var n = 0
var m = 0

every 100 ms on core 1:
    m = n

every 1 s:
    n = n + 1
`
	if d := check(t, readSide); !hasDiag(d, "W620") {
		t.Errorf("a core-1 read of a main-written global should warn W620: %v", d)
	}
	// core-only and main-only globals are fine
	separate := `board pico
var n = 0
var m = 0

every 100 ms on core 1:
    m = m + 1

every 1 s:
    n = n + 1
`
	if d := check(t, separate); hasDiag(d, "W620") {
		t.Errorf("globals kept to one core shouldn't warn: %v", d)
	}
	// a single-core board never warns (the block can't be on a second core anyway)
	_, diags := Compile("board uno\nvar n = 0\n\nevery 100 ms:\n    n = n + 1\n\nevery 1 s:\n    print n\n", "p.muj")
	if hasDiag(diags, "W620") {
		t.Errorf("one core means nothing is shared across cores: %v", diags)
	}
}

// TestCoreSerial (W621): peripheral calls from a secondary core warn only where the board's driver really isn't
// synchronized. On RP2040-family boards 'print' goes through SerialUSB, which arduino-pico guards with CoreMutex -
// so it doesn't warn; SerialUART (uart/serial), Wire (i2c) and SPI have no such guard and do.
func TestCoreSerial(t *testing.T) {
	src := `board pico

every 100 ms on core 1:
    print "tick"
`
	if d := check(t, src); hasDiag(d, "W621") {
		t.Errorf("print on a Pico is SerialUSB/CoreMutex-protected; it shouldn't warn: %v", d)
	}
	esp := `board esp32

every 100 ms on core 1:
    print "tick"
`
	if d := check(t, esp); hasDiag(d, "W621") {
		t.Errorf("print on an ESP32 core is FreeRTOS-managed; it shouldn't warn: %v", d)
	}
	uart := `board pico
use mujard.uart
u = Uart(pin 4, pin 5)

task talk on core 1:
    u.write("x")
`
	if d := check(t, uart); !hasDiag(d, "W621") {
		t.Errorf("a uart write on a secondary core on a Pico should warn W621: %v", d)
	}
}

// TestTimerPwmConflict (B414): on AVR a pwm pin sits on a named timer - pins 9/10 on an Uno are Timer 1, which a
// Timer() owns outright; a pwm pin on a different timer (5/6 are Timer 0) shares nothing. tone() takes Timer 2, so
// it fights pwm pins 3/11, not 9/10.
func TestTimerPwmConflict(t *testing.T) {
	conflict := `board uno
use mujard.timer
t = Timer()
p = pin 9 as pwm
`
	if e := FirstError(check(t, conflict)); e == nil || e.Code != "B414" {
		t.Errorf("Timer() + pwm 9 (both Timer 1 on an Uno) should refuse: %v", check(t, conflict))
	}
	ok := `board uno
use mujard.timer
t = Timer()
p = pin 5 as pwm
`
	if e := FirstError(check(t, ok)); e != nil {
		t.Errorf("Timer() + pwm 5 (Timer 1 vs Timer 0) shouldn't conflict: %v", e)
	}
	tone := `board uno
p = pin 3 as pwm
b = pin 5 as output

at start:
    b.tone(440 Hz)
`
	if e := FirstError(check(t, tone)); e == nil || e.Code != "B414" {
		t.Errorf("tone + pwm 3 (both Timer 2 on an Uno) should refuse: %v", check(t, tone))
	}
}

// TestFlashSafety (W622): every persistence commit on RP-family now runs inside the SDK's flash_safe_execute (the
// multicore lockout), with both cores opting in via coreInitFlashSafe. So a flash-writing call is protected in
// both directions - 'sleep'/'restart'/'halt' and store.save no longer warn on either core, and the generated
// sketch carries the peer-lockout init on core 0 (setup) and core 1 (its entry). ESP32 (NVS is driver-managed)
// and single-core boards were never at risk.
func TestFlashSafety(t *testing.T) {
	for _, src := range []string{
		"board pico\n\nvar n = 0\nsaved var s = 1\n\ntask nap on core 1:\n    sleep 1 s\n",
		"board pico\n\nvar n = 0\nsaved var s = 1\n\nevery 1 s on core 1:\n    store.save(\"x\", 1)\n",
		"board pico\n\ntask stopper on core 1:\n    restart()\n",
		// both directions: the main core's write while core 1 is active is protected too (was incorrectly implied safe before)
		"board pico\n\nsaved var s = 1\n\nevery 1 s on core 1:\n    store.save(\"x\", 1)\n\nevery 5 s:\n    store.save(\"y\", 2)\n",
		"board esp32\n\nsaved var s = 1\n\nevery 1 s on core 1:\n    store.save(\"x\", 1)\n",
	} {
		if d := check(t, src); hasDiag(d, "W622") {
			t.Errorf("flash commits are lockout-protected on both cores now; no W622 should fire:\n%s\n%v", src, d)
		}
	}
	// The protection must actually be in the generated sketch: core 0's setup() and every core-1 entry call
	// coreInitFlashSafe so flash_safe_execute can lock out the peer during a commit.
	sk, diags := Compile("board pico\n\nevery 1 s on core 1:\n    store.save(\"y\", 1)\n\nevery 5 s:\n    store.save(\"x\", 1)\n", "t.muj")
	if e := FirstError(diags); e != nil {
		t.Fatalf("compile: %v", e)
	}
	var cpp string
	for _, content := range sk.Files {
		cpp += content + "\n" // setup() and the core-1 entry live in the program file, not the sketch.ino stub
	}
	if c := strings.Count(cpp, "mujard::coreInitFlashSafe()"); c < 2 {
		t.Errorf("coreInitFlashSafe should appear in setup AND the core-1 entry (got %d):\n%s", c, cpp)
	}
}

// TestPwmSlice (W623): RP-family PWM pins sit on slices of two - pins n and n+1 share a slice and its single
// frequency. Two tone() calls on same-slice pins with different pitches can't both hold; the same pitch is fine.
func TestPwmSlice(t *testing.T) {
	conflict := `board pico
a = pin 0 as output
b = pin 1 as output

at start:
    a.tone(440 Hz)
    b.tone(880 Hz)
`
	if d := check(t, conflict); !hasDiag(d, "W623") {
		t.Errorf("different-tone pins on one PWM slice should warn W623: %v", d)
	}
	// same frequency on one slice is exactly what the hardware does
	compatible := `board pico
a = pin 0 as output
b = pin 1 as output

at start:
    a.tone(440 Hz)
    b.tone(440 Hz)
`
	if d := check(t, compatible); hasDiag(d, "W623") {
		t.Errorf("same frequency on one slice shouldn't warn: %v", d)
	}
	// different slices are independent
	separate := `board pico
a = pin 0 as output
b = pin 2 as output

at start:
    a.tone(440 Hz)
    b.tone(880 Hz)
`
	if d := check(t, separate); hasDiag(d, "W623") {
		t.Errorf("pins on different slices shouldn't warn: %v", d)
	}
	// AVR has its own timer model, not PWM slices - no W623 there
	avr := `board uno
a = pin 3 as output
b = pin 5 as output

at start:
    a.tone(440 Hz)
    b.tone(880 Hz)
`
	if d := check(t, avr); hasDiag(d, "W623") {
		t.Errorf("AVR doesn't have RP-style PWM slices: %v", d)
	}
}

// TestCoreLaunchFails: a secondary-core launch on ESP32-family checks xTaskCreatePinnedToCore's return - a task the
// heap can't fund must not look like it started.
func TestCoreLaunchFails(t *testing.T) {
	sk, diags := Compile("board esp32\nvar n = 0\n\nevery 100 ms on core 1:\n    n = n + 1\n", "p.muj")
	if FirstError(diags) != nil || sk == nil {
		t.Fatalf("a core block on an ESP32 should compile: %v", diags)
	}
	if !strings.Contains(sk.Files["mujard_program.cpp"], "!= pdPASS") {
		t.Error("the core launch should check xTaskCreatePinnedToCore's return")
	}
}

// TestCommsOverflow: receive queues drop + count instead of failing when full. The host fakes exercise the same
// push path the callbacks use, so pretend over the cap lands in board.dropped_events.
func TestCommsOverflow(t *testing.T) {
	cxx := testCxx(t)
	src := `board esp32
use mujard.ble

at start:
    ble.start("x")

test "a full BLE queue drops and counts":
    ble.pretend("a")
    ble.pretend("b")
    ble.pretend("c")
    ble.pretend("d")
    ble.pretend("e")
    print board.dropped_events
    expect printed "1"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("BLE overflow test: %+v", res)
	}
	canSrc := `board esp32
use mujard.can
bus = Can(pin 21, pin 22, 500000 baud)

test "a full CAN queue drops and counts":
    bus.pretend(1, [1])
    bus.pretend(2, [1])
    bus.pretend(3, [1])
    bus.pretend(4, [1])
    bus.pretend(5, [1])
    bus.pretend(6, [1])
    bus.pretend(7, [1])
    bus.pretend(8, [1])
    bus.pretend(9, [1])
    print board.dropped_events
    expect printed "1"
`
	res = runTests(t, cxx, canSrc)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("CAN overflow test: %+v", res)
	}
}
