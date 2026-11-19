package lang

import (
	"strings"
	"testing"
)

// TestBrownOut: brown_out() in a test (spec 12, plan Phase 20 fault injection) is a reset, not a power-off: the board
// starts again with only committed EEPROM/flash, and board.reset_reason tells the program why it started. On a real
// AVR the brown-out detector reports that reason after a supply dip.
func TestBrownOut(t *testing.T) {
	cxx := testCxx(t)
	src := `board uno
saved var presses = 0
var started = 0
button = pin 2 as input

at start:
    started += 1
    if board.reset_reason == brownout:
        print "boot after brown-out"
    else:
        print "boot after power"

on button.pressed:
    presses += 1

test "a brown-out reset keeps committed saves":
    button.pretend(true)
    advance(100 ms)
    button.pretend(false)
    advance(61 s)
    expect board.reset_reason == power_on, "first boot is power-on"
    brown_out()
    expect board.reset_reason == brownout, "the supply dipped, it did not go off"
    expect presses == 1, "written before the dip"
    expect started == 1, "the program started again from the beginning"
    expect printed "boot after brown-out"

test "a brown-out loses a change not yet saved":
    button.pretend(true)
    advance(100 ms)
    button.pretend(false)
    advance(100 ms)
    expect presses == 1, "counted in RAM"
    brown_out()
    expect presses == 0, "not written yet when the brown-out reset fired"

test "a power cut after a brown-out reports power-on":
    brown_out()
    expect board.reset_reason == brownout, "the last reset was the dip"
    power_cut()
    expect board.reset_reason == power_on, "the power came on again"
`
	res := runTests(t, cxx, src)
	if len(res) != 3 {
		t.Fatalf("want 3 results, got %v", res)
	}
	for _, r := range res {
		if !r.Passed {
			t.Errorf("%s: %s (line %d)\n%s", r.Name, r.Message, r.FailLine, r.Output)
		}
	}
	if !strings.Contains(res[0].Output, "--- brown out ---") || !strings.Contains(res[2].Output, "--- power cut ---") {
		t.Errorf("the output names each kind of restart:\n%s\n%s", res[0].Output, res[2].Output)
	}
	// only in tests (C317), and only on its own line in the test itself
	if _, diags := CheckSource("board uno\nat start:\n    brown_out()\n"); FirstError(diags) == nil || FirstError(diags).Code != "C317" {
		t.Errorf("brown_out in the program: %v", diags)
	}
	_, diags := CompileTests("board uno\nvar x = 0\ntest \"t\":\n    if x == 0:\n        brown_out()\n", "p.muj", "")
	if e := FirstError(diags); e == nil || e.Code != "C320" {
		t.Errorf("brown_out inside an if: %v", diags)
	}
	// reset_reason is observable in a PC test, where the restart can say why it happened. ESP targets read a reset
	// API, Optiboot AVR boards read the reset flags passed in r2, and the ARM chips read their reset-cause
	// registers. Other bootloaders/targets still refuse it.
	for _, board := range []string{"uno", "nano", "esp32", "esp8266", "bluepill", "zero", "nano33ble", "unor4"} {
		sk, diags := Compile("board "+board+"\nat start:\n    if board.reset_reason == wake:\n        print 1\n", "p.muj")
		if e := FirstError(diags); e != nil || sk == nil || !strings.Contains(sk.Files["mujard_program.cpp"], "mujard::board_reset_reason() == mujard::RESET_WAKE") {
			t.Errorf("reset_reason on %s should reach board_reset_reason(): %v", board, diags)
		}
		if board == "uno" || board == "nano" {
			rt := sk.Files["mujard_rt.cpp"]
			if !strings.Contains(rt, `section(".init0")`) || !strings.Contains(rt, "sts %0, r2") {
				t.Errorf("%s runtime has no Optiboot reset-flag capture", board)
			}
		}
		if board == "bluepill" && !strings.Contains(sk.Files["mujard_rt.cpp"], "RCC->CSR") {
			t.Error("bluepill runtime has no STM32 reset-flag path")
		}
		if board == "zero" && !strings.Contains(sk.Files["mujard_rt.cpp"], "PM->RCAUSE.reg") {
			t.Error("zero runtime has no SAMD reset-cause path")
		}
		if board == "nano33ble" && !strings.Contains(sk.Files["mujard_rt.cpp"], "NRF_POWER->RESETREAS") {
			t.Error("nano33ble runtime has no nRF reset-reason path")
		}
		if board == "unor4" && !strings.Contains(sk.Files["mujard_rt.cpp"], "R_SYSTEM->RSTSR0") {
			t.Error("unor4 runtime has no RA4M1 reset-status path")
		}
	}
	for _, board := range []string{"mega", "leonardo", "micro", "pico"} {
		_, diags = Compile("board "+board+"\nat start:\n    print board.reset_reason\n", "p.muj")
		if e := FirstError(diags); e == nil || e.Code != "C320" {
			t.Errorf("reset_reason on %s should be C320: %v", board, diags)
		}
	}
}
