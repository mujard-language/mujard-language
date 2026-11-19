package lang

import (
	"strings"
	"testing"
)

// TestQueue: use mujard.queue - q = Queue(text, 8) passes values between blocks, and is the safe way to hand data
// between cores (where a plain global warns W620). push is false + counts a drop when full; pop gives none when
// empty; 'when q.available:' fires on items. The host fake runs the same ring unlocked.
func TestQueue(t *testing.T) {
	cxx := testCxx(t)
	src := `board esp32
use mujard.queue

q = Queue(text, 4)

task producer on core 1:
    q.push("one")
    q.push("two")
    q.push("three")

every 10 ms:
    if q.available:
        var m = q.pop()
        if m != none:
            print "got {m}"

test "queue values cross to the consumer in order":
    advance(400 ms)
    expect printed "got one"
    expect printed "got two"
    expect printed "got three"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("queue test: %+v", res)
	}
	// a full queue drops and counts instead of blocking or corrupting
	full := `board esp32
use mujard.queue

q = Queue(int, 2)

test "a full queue drops and counts":
    print q.push(1)
    print q.push(2)
    print q.push(3)
    print board.dropped_events
    expect printed "true"
    expect printed "true"
    expect printed "false"
    expect printed "1"
`
	res = runTests(t, cxx, full)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("queue overflow test: %+v", res)
	}
	// empty pops are none, not a hang or a garbage value
	empty := `board esp32
use mujard.queue

q = Queue(int, 2)

test "an empty queue pops none":
    var m = q.pop()
    if m == none:
        print "empty"
    print q.n
    expect printed "empty"
    expect printed "0"
`
	res = runTests(t, cxx, empty)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("queue empty test: %+v", res)
	}
	// the queue itself is the share mechanism: a queue touched from both cores doesn't warn W620
	shared := `board pico
use mujard.queue

q = Queue(int, 4)

task producer on core 1:
    q.push(1)

every 100 ms:
    var m = q.pop()
`
	if d := check(t, shared); hasDiag(d, "W620") {
		t.Errorf("a Queue is the sanctioned cross-core channel; it shouldn't warn W620: %v", d)
	}
	// bad element types and capacities refuse
	for src, want := range map[string]string{
		"board uno\nuse mujard.queue\nq = Queue(servo, 4)\n":        "C303",
		"board uno\nuse mujard.queue\nq = Queue(text, 0)\n":         "C303",
		"board uno\nuse mujard.queue\nq = Queue(text, 200)\n":       "C303",
		"board uno\nuse mujard.queue\nq = Queue(text)\n":            "C307",
		"board uno\nuse mujard.queue\nq = Queue(unknown_type, 4)\n": "C303",
	} {
		if e := FirstError(check(t, src)); e == nil || e.Code != want {
			t.Errorf("%q should be %s: %v", src, want, check(t, src))
		}
	}
	// Queue without 'use mujard.queue' refuses C318 like every lib name
	if e := FirstError(check(t, "board uno\nq = Queue(text, 4)\n")); e == nil || e.Code != "C318" {
		t.Errorf("Queue without use mujard.queue should be C318: %v", check(t, "board uno\nq = Queue(text, 4)\n"))
	}
	// the generated C++ is the bounded runtime queue with the declared element and capacity
	sk, diags := Compile("board esp32\nuse mujard.queue\nq = Queue(text, 8)\n", "p.muj")
	if FirstError(diags) != nil || sk == nil {
		t.Fatalf("a queue on an ESP32 should compile: %v", diags)
	}
	cpp := sk.Files["mujard_program.cpp"]
	if !strings.Contains(cpp, "mujard::Queue<") || !strings.Contains(cpp, ", 8>") {
		t.Errorf("the generated code should emit a bounded mujard::Queue<T, 8>: %s", cpp[:min(len(cpp), 400)])
	}
}
