package lang

import (
	"strings"
	"testing"
)

// TestPowerCut: power_cut() in a test (spec 12, plan Phase 20 fault injection) - the board loses power and starts
// again with only what was written to its EEPROM, so a test can show what a saved variable keeps and what a cut
// right after a change loses (saved values are written once a minute on an Uno, spec 11).
func TestPowerCut(t *testing.T) {
	cxx := testCxx(t)
	src := `board uno
saved var presses = 0
var started = 0
button = pin 2 as input

at start:
    started += 1
    print "power on, presses = {presses}"

on button.pressed:
    presses += 1

test "a saved press survives a power cut":
    button.pretend(true)
    advance(100 ms)
    button.pretend(false)
    advance(61 s)
    power_cut()
    expect presses == 1, "written before the cut"
    expect started == 1, "the program started again from the beginning"
    expect printed "power on, presses = 1"

test "a press just before a cut is lost":
    button.pretend(true)
    advance(100 ms)
    button.pretend(false)
    advance(100 ms)
    expect presses == 1, "counted in RAM"
    power_cut()
    expect presses == 0, "not written yet when the power went"

test "two cuts in a row":
    power_cut()
    button.pretend(true)
    advance(100 ms)
    button.pretend(false)
    advance(61 s)
    power_cut()
    expect presses == 1, "kept across the second cut"
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
	if !strings.Contains(res[0].Output, "--- power cut ---") || !strings.Contains(res[0].Output, "power on, presses = 0") {
		t.Errorf("the output shows both runs with the cut between them:\n%s", res[0].Output)
	}
	// only in tests (C317), and only on its own line in the test itself
	if _, diags := CheckSource("board uno\nat start:\n    power_cut()\n"); FirstError(diags) == nil || FirstError(diags).Code != "C317" {
		t.Errorf("power_cut in the program: %v", diags)
	}
	_, diags := CompileTests("board uno\nvar x = 0\ntest \"t\":\n    if x == 0:\n        power_cut()\n", "p.muj", "")
	if e := FirstError(diags); e == nil || e.Code != "C320" {
		t.Errorf("power_cut inside an if: %v", diags)
	}
}

// TestPowerCutBigSaved: a saved region over 255 bytes - a text[255] is a 257-byte slot, and the slot after it sits
// past the old one-byte sizes, which lost or corrupted values until the layout fields grew (spec 11).
func TestPowerCutBigSaved(t *testing.T) {
	cxx := testCxx(t)
	long := strings.Repeat("a", 250)
	src := `board uno
saved var big: text[255] = "empty"
saved var tail = 0
var started = 0

at start:
    started += 1
    print "started {started}"

test "a saved var over 255 bytes survives a power cut":
    big = "` + long + `"
    tail = 42
    advance(61 s)
    power_cut()
    expect big == "` + long + `", "the big saved var kept its 250 characters"
    expect tail == 42, "the slot after it kept its value"
    expect started == 1, "the program restarted"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("big saved var test: %+v", res)
	}
}
