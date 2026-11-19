package main

import (
	"strings"
	"testing"
)

// TestMemoryReport: the build's flash and RAM use come from arduino-cli's own size report, with a warning when RAM is
// nearly full and a plain reason when the program doesn't fit.
func TestMemoryReport(t *testing.T) {
	uno := "Sketch uses 4862 bytes (15%) of program storage space. Maximum is 32256 bytes.\n" +
		"Global variables use 412 bytes (20%) of dynamic memory, leaving 1636 bytes for local variables. Maximum is 2048 bytes.\n"
	s := buildSizes(uno)
	if s["flash"] != 4862 || s["flashMax"] != 32256 || s["ram"] != 412 || s["ramMax"] != 2048 {
		t.Fatalf("sizes = %v", s)
	}
	lines := memoryLines(s, "Arduino Uno")
	if len(lines) != 1 || lines[0] != "Memory: program 4862 of 32256 bytes (15%), RAM 412 of 2048 bytes (20%)." {
		t.Errorf("lines = %q", lines)
	}
	if tooBig(s, "Arduino Uno") != "" {
		t.Error("a program that fits isn't too big")
	}

	full := buildSizes(strings.Replace(uno, "412 bytes (20%)", "1700 bytes (83%)", 1))
	lines = memoryLines(full, "Arduino Uno")
	if len(lines) != 2 || !strings.HasPrefix(lines[1], "Warning: 83% of the Arduino Uno's RAM") || !strings.Contains(lines[1], "348 bytes") {
		t.Errorf("nearly full RAM: %q", lines)
	}

	over := buildSizes(strings.Replace(uno, "412 bytes (20%)", "2500 bytes (122%)", 1))
	if why := tooBig(over, "Arduino Uno"); !strings.Contains(why, "need 2500 bytes and the board has 2048") {
		t.Errorf("RAM overflow: %q", why)
	}
	big := buildSizes(strings.Replace(uno, "4862 bytes (15%)", "40000 bytes (124%)", 1))
	if why := tooBig(big, "Arduino Uno"); !strings.Contains(why, "needs 40000 bytes of program memory and the board has 32256") {
		t.Errorf("flash overflow: %q", why)
	}

	// boards whose core reports only the program size, and logs with no report
	esp := buildSizes("Sketch uses 280000 bytes (21%) of program storage space. Maximum is 1310720 bytes.\n")
	if l := memoryLines(esp, "ESP32"); len(l) != 1 || !strings.HasPrefix(l[0], "Memory: program 280000 of 1310720 bytes (21%).") {
		t.Errorf("flash only: %q", l)
	}
	if buildSizes("nothing to see") != nil || memoryLines(nil, "x") != nil {
		t.Error("no report: no sizes, no lines")
	}
}
