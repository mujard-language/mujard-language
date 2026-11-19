package lang

import "testing"

// TestSerialBudget: W619 only where the serial port provably can't keep up with an 'every' block's prints.
func TestSerialBudget(t *testing.T) {
	w619 := func(src string) []Diagnostic {
		t.Helper()
		_, diags := CheckSource(src)
		var out []Diagnostic
		for _, d := range diags {
			if d.Code == "W619" {
				out = append(out, d)
			}
		}
		return out
	}
	// 38 characters + \r\n = 40; at 115200 baud each takes 86.8 us: 3.5 ms a run, every 1 ms
	busy := "board uno\nevery 1 ms:\n    print \"the sensor reads quite a lot right now\"\n"
	d := w619(busy)
	want := "This block prints at least 40 characters every 1 ms; at 115200 baud that takes 3.5 ms, so the serial port falls behind and runs are skipped (board.overruns). Print less, or less often."
	if len(d) != 1 || d[0].Pos != (Pos{2, 1}) || d[0].Message != want {
		t.Errorf("a fast block printing a lot: %v", d)
	}
	if d := w619("board uno\nevery 100 ms:\n    print \"the sensor reads quite a lot right now\"\n"); len(d) != 0 {
		t.Errorf("3.5 ms of printing every 100 ms is fine: %v", d)
	}
	// a slower port makes the same print too much
	if d := w619("board uno\nat start:\n    serial(9600 baud)\nevery 20 ms:\n    print \"twenty-five characters !!\"\n"); len(d) != 1 {
		t.Errorf("26 characters at 9600 baud (27 ms) every 20 ms: %v", d)
	}
	// a frequency is a period too
	if d := w619("board uno\nevery 1 kHz:\n    print \"the sensor reads quite a lot right now\"\n"); len(d) != 1 {
		t.Errorf("every 1 kHz is every 1 ms: %v", d)
	}
	// only the plain text that every run reaches counts
	if d := w619("board uno\nvar x = 0\nevery 1 ms:\n    x += 1\n    if x > 1000:\n        print \"the sensor reads quite a lot right now\"\n"); len(d) != 0 {
		t.Errorf("a print under an if doesn't always run: %v", d)
	}
	// native USB serial has no baud limit
	if d := w619("board leonardo\nevery 1 ms:\n    print \"the sensor reads quite a lot right now\"\n"); len(d) != 0 {
		t.Errorf("the Leonardo's serial is native USB: %v", d)
	}
}
