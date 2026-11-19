package lang

import "testing"

// TestMulticore: 'every'/'after'/'when' take 'on core N' to run on a secondary core (ESP32-family's
// xTaskCreatePinnedToCore, RP2040/RP2350's multicore_launch_core1). The host test keeps them on the main
// scheduler - a fake can't prove real parallelism - and boards without a second core refuse.
func TestMulticore(t *testing.T) {
	cxx := testCxx(t)
	src := `board esp32

every 1 s on core 1:
    print "core1 tick"

after 500 ms on core 1:
    print "core1 once"

test "core-1 blocks keep running on the fake's same scheduler":
    advance(1600 ms)
    expect printed "core1 tick"
    expect printed "core1 once"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("multicore test: %+v", res)
	}
	// a single-core board refuses; a board with 2 cores refuses 'on core 2'
	_, diags := Compile("board uno\nevery 1 s on core 1:\n    print 1\n", "p.muj")
	if e := FirstError(diags); e == nil || e.Code != "B415" {
		t.Errorf("'on core 1' on an Uno should be refused: %v", diags)
	}
	_, diags = Compile("board pico\nevery 1 s on core 2:\n    print 1\n", "p.muj")
	if e := FirstError(diags); e == nil || e.Code != "B415" {
		t.Errorf("'on core 2' on a Pico should be refused: %v", diags)
	}
	// a task on another core works (the coroutine runs on that core); the single-core Uno refuses
	_, diags = Compile("board esp32\ntask blink on core 1:\n    print 1\n", "p.muj")
	if e := FirstError(diags); e != nil {
		t.Errorf("a task on core 1 should compile on a 2-core board: %v", diags)
	}
	_, diags = Compile("board uno\ntask blink on core 1:\n    print 1\n", "p.muj")
	if e := FirstError(diags); e == nil || e.Code != "B415" {
		t.Errorf("a task on core 1 on an Uno should be refused: %v", diags)
	}
}
