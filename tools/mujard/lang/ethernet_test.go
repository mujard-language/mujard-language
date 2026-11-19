package lang

import "testing"

// TestEthernet: use mujard.ethernet - ethernet.start() brings the wire up; connected/ip read it and
// 'on ethernet.joined'/'lost' fire on link changes; tests fake the wire with ethernet.pretend. Boards
// without an Ethernet MAC refuse.
func TestEthernet(t *testing.T) {
	cxx := testCxx(t)
	src := `board esp32
use mujard.ethernet

at start:
    ethernet.start()

on ethernet.joined:
    print "eth up {ethernet.ip}"

on ethernet.lost:
    print "eth down"

test "the wire comes up and drops":
    ethernet.pretend(true)
    advance(200 ms)
    print ethernet.connected
    expect printed "eth up 192.168.1.50"
    expect printed "true"
    ethernet.pretend(false)
    advance(200 ms)
    expect printed "eth down"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("Ethernet test: %+v", res)
	}
	// a board without an EMAC refuses
	_, diags := Compile("board uno\nuse mujard.ethernet\nat start:\n    ethernet.start()\n", "p.muj")
	if e := FirstError(diags); e == nil || e.Code != "C320" {
		t.Errorf("Ethernet on an Uno should be refused: %v", diags)
	}
}
