package lang

import "testing"

// TestBle: use mujard.ble - ble.start advertises, ble.send / ble.read_line / 'on ble.line' exchange text lines;
// tests fake the link with ble.pretend. Boards without a BLE backend refuse.
func TestBle(t *testing.T) {
	cxx := testCxx(t)
	src := `board esp32
use mujard.ble

at start:
    ble.start("mujard-esp32")
    print ble.connected
    print ble.send("hello")

on ble.line as l:
    print "ble rx {l}"

test "BLE lines go both ways":
    ble.pretend("ping")
    advance(200 ms)
    expect printed "true"
    expect printed "true"
    expect printed "ble rx ping"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("BLE test: %+v", res)
	}
	// a board without a BLE backend refuses
	_, diags := Compile("board uno\nuse mujard.ble\nat start:\n    ble.start(\"x\")\n", "p.muj")
	if e := FirstError(diags); e == nil || e.Code != "C320" {
		t.Errorf("BLE on an Uno should be refused: %v", diags)
	}
	// the Pico W links its CYW43 Bluetooth stack only through a board option, and only for BLE programs
	sk, diags := Compile("board picow\nuse mujard.ble\nat start:\n    ble.start(\"x\")\n", "p.muj")
	if FirstError(diags) != nil {
		t.Fatalf("BLE on a Pico W should compile: %v", diags)
	}
	if sk.FQBN != "rp2040:rp2040:rpipicow:ipbtstack=ipv4btcble" {
		t.Errorf("a picow BLE build needs the Bluetooth stack option, got %q", sk.FQBN)
	}
	sk, _ = Compile("board picow\nat start:\n    pass\n", "p.muj")
	if sk.FQBN != "rp2040:rp2040:rpipicow" {
		t.Errorf("a picow build without BLE keeps the plain FQBN, got %q", sk.FQBN)
	}
}
