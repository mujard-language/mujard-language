package lang

import "testing"

// TestTimer: use mujard.timer - a Timer() object for hardware-precision callbacks. t.every/t.freq/t.once arm it,
// 'on t.tick' runs on the scheduler and 'on t.tick fast' in the ISR; the fake ticks through advance/pretend.
// Boards without a timer backend refuse.
func TestTimer(t *testing.T) {
	cxx := testCxx(t)
	src := `board uno
use mujard.timer

var fastTicks = 0
t = Timer()

at start:
    t.every(1 s)

on t.tick:
    print "tick"

on t.tick fast:
    fastTicks = fastTicks + 1

test "the timer ticks on the fake's scheduler and ISR":
    advance(2600 ms)
    t.pretend()
    advance(100 ms)
    print fastTicks
    expect printed "tick"
    expect printed "3"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("Timer test: %+v", res)
	}
	// a board without a timer backend refuses
	_, diags := Compile("board bluepill\nuse mujard.timer\nt = Timer()\n", "p.muj")
	if e := FirstError(diags); e == nil || e.Code != "C320" {
		t.Errorf("Timer on a Blue Pill should be refused: %v", diags)
	}
}
