package lang

import "testing"

// TestCanBus: use mujard.can - Can(tx, rx, rate) opens the TWAI bus on the ESP32 family; bus.send queues a frame,
// bus.read() and 'on bus.message' give received frames. Tests fake the bus with bus.pretend, like i2c.
func TestCanBus(t *testing.T) {
	cxx := testCxx(t)
	src := `board esp32
use mujard.can

bus = Can(pin 21, pin 22, 500000 baud)

at start:
    print bus.send(0x123, [1, 2, 3])

on bus.message as m:
    print "rx {m.id} {m.data.length}"

test "CAN frames go both ways":
    bus.pretend(1110, [9, 8])
    advance(200 ms)
    expect printed "true"
    expect printed "rx 1110 2"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("CAN test: %+v", res)
	}
	// a board without a TWAI controller refuses
	_, diags := Compile("board uno\nuse mujard.can\nbus = Can(pin 3, pin 4, 500000 baud)\n", "p.muj")
	if e := FirstError(diags); e == nil || e.Code != "C320" {
		t.Errorf("CAN on an Uno should be refused: %v", diags)
	}
}
