package lang

import "testing"

// TestServoLimit: 'limit s to lo..hi' works on a servo, not only pwm (spec 11.3): every later move clamps into the
// range and the first clamp logs one warning, exactly like a pwm limit.
func TestServoLimit(t *testing.T) {
	cxx := testCxx(t)
	src := `board uno
use mujard.servo

arm = Servo(pin 9)
limit arm to 10 deg..90 deg

at start:
    arm.move(45 deg)
    print arm.angle
    arm.move(120 deg)
    print arm.angle
    arm.move(-20 deg)
    print arm.angle

test "a servo limit clamps every later move":
    advance(2 s)
    expect printed "45 deg"
    expect printed "90 deg"
    expect printed "10 deg"
    expect printed "WARN"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("servo limit test: %+v", res)
	}
}
