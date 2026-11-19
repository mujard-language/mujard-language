package lang

import (
	"strings"
	"testing"
)

// TestI2CPretend: tests can put a device on the fake I2C bus and take it away again (spec 12, plan Phase 20 fault
// injection), so an I2C program's working path and its failure path are both testable - before, the bus was always
// silent and only "no answer" could be tested.
func TestI2CPretend(t *testing.T) {
	cxx := testCxx(t)
	src := `board uno
use mujard.i2c

const SENSOR = 0x48
var request: list[int, 1] = [0]
var reply: list[int, 2] = [0x19, 0x40]
var last = -1

every 1 s:
    i2c.write(SENSOR, request)
    var data = i2c.read(SENSOR, 2)
    if data == none:
        print "No answer from the sensor"
        last = -1
    else:
        last = (data[0] << 8) | data[1]
        print "Raw reading: {last}"

test "the sensor answers":
    i2c.pretend(SENSOR, reply)
    advance(1500 ms)
    expect last == 6464, "0x19 0x40 is 6464"
    expect printed "Raw reading: 6464"

test "the sensor drops off the bus mid-run":
    i2c.pretend(SENSOR, reply)
    advance(1500 ms)
    expect last == 6464, "answering at first"
    i2c.pretend(SENSOR, none)
    advance(1000 ms)
    expect last == -1, "gone: the program sees no answer"
    expect printed "No answer from the sensor"

test "a scan finds what is on the bus":
    i2c.pretend(SENSOR, reply)
    i2c.pretend(0x3C, request)
    var found = i2c.scan()
    expect found.length == 2, "two devices"
    expect found[0] == 0x3C and found[1] == SENSOR, "in address order"
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
	// outside a test, i2c.pretend is refused like every pretend (C317)
	_, diags := CheckSource("board uno\nuse mujard.i2c\nvar r: list[int, 1] = [1]\nat start:\n    i2c.pretend(0x48, r)\n")
	if e := FirstError(diags); e == nil || e.Code != "C317" || !strings.Contains(e.Message, "pretend") {
		t.Errorf("i2c.pretend in the program: %v", diags)
	}
}
