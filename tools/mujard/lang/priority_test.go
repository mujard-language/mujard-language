package lang

import (
	"strings"
	"testing"
)

// priority N orders blocks inside their existing scheduler phase; it never makes one phase preempt another (spec 6).
func TestBlockPriority(t *testing.T) {
	sk, diags := Compile(`board uno
when true:
    print "when-low"
when true priority 2:
    print "when-high"
every 1 s:
    print "timed-low"
every 1 s priority 3:
    print "timed-high"
task low:
    print "task-low"
task high priority 4:
    print "task-high"
`, "priority.muj")
	if FirstError(diags) != nil {
		t.Fatal(diags)
	}
	cpp := sk.Files["mujard_program.cpp"]
	at := func(what string) int {
		i := strings.Index(cpp, what)
		if i < 0 {
			t.Fatalf("missing %q in\n%s", what, cpp)
		}
		return i
	}
	got := []int{
		at(`"when-high"`), at(`"when-low"`),
		at(`"timed-high"`), at(`"timed-low"`),
		at(`mujard_task_high.run();`), at(`mujard_task_low.run();`),
	}
	for i := 1; i < len(got); i++ {
		if got[i] < got[i-1] {
			t.Fatalf("priorities/phases emitted out of order %v in\n%s", got, cpp)
		}
	}
}

func TestBlockPriorityPlacement(t *testing.T) {
	for _, tc := range []struct {
		name, src, code string
	}{
		{"an on event has no priority", "board uno\nb = pin 2 as input\non b.changed priority 2:\n    pass\n", "P201"},
		{"a nested after has no priority", "board uno\nat start:\n    after 1 s priority 2:\n        pass\n", "P201"},
		{"the number is required", "board uno\nevery 1 s priority high:\n    pass\n", "P205"},
	} {
		e := FirstError(check(t, tc.src))
		if e == nil || e.Code != tc.code {
			t.Errorf("%s: expected %s, got %v", tc.name, tc.code, e)
		}
	}
}
