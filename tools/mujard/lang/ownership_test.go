package lang

import "testing"

// TestPinOwnership: a pin has one owner, whether a device line or an object's constructor took it (B407), and the
// message names that owner.
func TestPinOwnership(t *testing.T) {
	first := func(src string) *Diagnostic {
		t.Helper()
		_, diags := CheckSourceModules(src, nil)
		for i := range diags {
			if diags[i].Code == "B407" {
				return &diags[i]
			}
		}
		return nil
	}
	cases := []struct{ name, src, msg string }{
		{"a device, then an object on its pin",
			"board uno\nuse mujard.servo\nled = pin 9 as output\narm = Servo(pin 9)\nat start:\n    led.on()\n    arm.move(90 deg)\n",
			"Pin 9 is already used by 'led' (line 3)."},
		{"two objects on one pin",
			"board uno\nuse mujard.servo\nuse mujard.ultrasonic\narm = Servo(pin 9)\nhand = Ultrasonic(trig: pin 9, echo: pin 10)\nat start:\n    arm.move(90 deg)\n    print hand.distance()\n",
			"Pin 9 is already used by 'arm' (line 4)."},
		{"a bare pin value names itself",
			"board uno\nat start:\n    print pin 9\n    print pin 9\n",
			"Pin 9 is already used by 'pin 9' (line 3)."},
	}
	for _, tc := range cases {
		d := first(tc.src)
		if d == nil || d.Message != tc.msg {
			t.Errorf("%s: want B407 %q, got %v", tc.name, tc.msg, d)
		}
	}
	if d := first("board uno\nuse mujard.servo\nuse mujard.ultrasonic\narm = Servo(pin 9)\nhand = Ultrasonic(trig: pin 7, echo: pin 8)\nat start:\n    arm.move(90 deg)\n    print hand.distance()\n"); d != nil {
		t.Errorf("separate pins have separate owners: %v", d)
	}
}

// TestTimerOwnership: tone() and NewPing both define Timer 2's interrupt on an Uno, which the linker refuses
// (shown with arduino-cli 2026-10-30), so the checker reports it first (B414) - and only where it is true.
func TestTimerOwnership(t *testing.T) {
	b414 := func(src string) []Diagnostic {
		t.Helper()
		_, diags := CheckSourceModules(src, nil)
		var out []Diagnostic
		for _, d := range diags {
			if d.Code == "B414" {
				out = append(out, d)
			}
		}
		return out
	}
	prog := func(board, buzzer string) string {
		return "board " + board + "\nuse \"NewPing\"\nsonar = NewPing(7, 6, 200)\nbuzzer = pin " + buzzer + " as output\n" +
			"at start:\n    print sonar.ping_cm()\n    buzzer.tone(440 Hz, 100 ms)\n"
	}
	d := b414(prog("uno", "8"))
	want := "'NewPing' and 'tone' both need the Arduino Uno's Timer 2 interrupt, so the program can't be built. Use only one of them on this board."
	if len(d) != 1 || d[0].Pos != (Pos{7, 12}) || d[0].Message != want {
		t.Errorf("tone after NewPing on an Uno: one B414 at the tone call, got %v", d)
	}
	if d := b414(prog("nano", "8")); len(d) != 1 {
		t.Errorf("the Nano is the same chip: %v", d)
	}
	if d := b414(prog("mega", "8")); len(d) != 1 {
		t.Errorf("the Mega's tone and NewPing share Timer 2 too: %v", d)
	}
	if d := b414(prog("leonardo", "8")); len(d) != 0 {
		t.Errorf("the Leonardo's tone uses Timer 3: %v", d)
	}
	if d := b414(prog("esp32", "GPIO13")); len(d) != 0 {
		t.Errorf("an ESP32 has no such conflict: %v", d)
	}
	if d := b414("board uno\nuse \"NewPing\"\nsonar = NewPing(7, 6, 200)\nat start:\n    print sonar.ping_cm()\n"); len(d) != 0 {
		t.Errorf("NewPing alone owns its timer: %v", d)
	}
	if d := b414("board uno\nbuzzer = pin 8 as output\nat start:\n    buzzer.melody(\"C4 E4 G4\")\n"); len(d) != 0 {
		t.Errorf("melody alone owns its timer: %v", d)
	}
}
