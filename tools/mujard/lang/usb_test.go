package lang

import "testing"

// TestUsb: use mujard.usb - the board types text, taps keys and moves the pointer on the host it's plugged
// into; the fake host is usb.pretend. Boards without a standard-HID backend refuse.
func TestUsb(t *testing.T) {
	cxx := testCxx(t)
	src := `board leonardo
use mujard.usb

at start:
    print usb.connected

test "USB typing goes out once the host is listening":
    usb.pretend(true)
    advance(200 ms)
    print usb.type("hello host")
    print usb.key("a")
    print usb.mouse_move(10, -5)
    print usb.click()
    expect printed "false"
    expect printed "true"
    expect printed "true"
    expect printed "true"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("USB test: %+v", res)
	}
	// a board without a standard-HID backend refuses
	_, diags := Compile("board uno\nuse mujard.usb\nat start:\n    print usb.connected\n", "p.muj")
	if e := FirstError(diags); e == nil || e.Code != "C320" {
		t.Errorf("USB on an Uno should be refused: %v", diags)
	}
}
