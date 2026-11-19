package main

import (
	"strings"
	"testing"

	"mujard/lang"
)

// A REPL session runs on the real runtime: values persist, statements stick, int wraps like on the chip, and a line
// with an error is reported and forgotten. Each line is a real build on this PC (needs g++/clang++ or MUJARD_CXX).
func TestReplSession(t *testing.T) {
	cxx := lang.FindCxx()
	if cxx == "" {
		t.Skip("no C++ compiler on this PC (MUJARD_CXX)")
	}
	s := &replSession{board: "uno", cxx: cxx}
	step := func(in, wantOut string) {
		t.Helper()
		out, prob := s.eval(in)
		if prob != "" {
			t.Fatalf("%q: unexpected problem: %s", in, prob)
		}
		if strings.TrimSpace(out) != wantOut {
			t.Fatalf("%q printed %q, want %q", in, strings.TrimSpace(out), wantOut)
		}
	}
	step("var x = 5", "")
	step("x * 2", "10")
	step("x += 1", "")
	step("x", "6")
	step("2147483647 + 1", "-2147483648") // int is 32-bit and wraps, on every board (spec 5.1)

	out, prob := s.eval("y + 1")
	if prob == "" || !strings.Contains(prob, "C30") {
		t.Fatalf("an unknown name should be the compiler's error, got %q / %q", out, prob)
	}
	step("x", "6") // the bad line left the session as it was
	if len(s.decls) != 1 || len(s.steps) != 1 {
		t.Errorf("the session should hold one declaration and one step: %v %v", s.decls, s.steps)
	}
}
