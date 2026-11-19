package lang

// The security pass (2026-10-22): what a program can make the generated C++ do, and what it can make the compiler
// read. Each test names the hole it closes.

import (
	"strings"
	"testing"
)

func generate(t *testing.T, src string) string {
	t.Helper()
	sk, diags := Compile(src, "prog.muj")
	if e := FirstError(diags); e != nil {
		t.Fatalf("doesn't compile: %v", e)
	}
	return sk.Files["mujard_program.cpp"]
}

// Signed overflow and division by zero are undefined behaviour in C++; the spec promises wrapping ints and a safe halt.
func TestIntArithmeticNeverUsesUndefinedOperators(t *testing.T) {
	cpp := generate(t, `board uno
var a = 7
var b = 0
var t = 2 s
at start:
    a = a + b
    a = a - 1
    a = a * 3
    a = a / b
    a = a mod b
    a = a << 2
    a = a >> b
    a = -a
    a += 1
    a -= b
    a *= 2
    a /= b
    t = t / b
    t = t / 0.5
`)
	for _, want := range []string{
		"mujard::iadd(a, b)", "mujard::isub(a, 1)", "mujard::imul(a, 3)", "mujard::idiv(a, b, 9)", "mujard::imod(a, b, 10)",
		"mujard::ishl(a, 2, 11)", "mujard::ishr(a, b, 12)", "mujard::ineg(a)", "mujard::iadd_to(a, 1);", "mujard::isub_to(a, b);",
		"mujard::imul_to(a, 2);", "mujard::idiv_to(a, b, 17);", "mujard::tdiv(t, b, 18)", "mujard::tdivf(t, ",
	} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %s in\n%s", want, cpp)
		}
	}
	for _, bad := range []string{"(a / b)", "(a % b)", "a += 1;", "a /= b;", "(a + b)"} {
		if strings.Contains(cpp, bad) {
			t.Errorf("found a plain C++ operator %q", bad)
		}
	}
}

// Text goes into C++ string literals: no escape may swallow the next character, and no trigraph may form.
func TestTextBecomesSafeCppLiterals(t *testing.T) {
	cases := map[string]string{
		"a\"b":        `"a\"b"`,
		`back\slash`:  `"back\\slash"`,
		"??/":         `"\?\?/"`,
		"\x01" + "ab": `"\001ab"`, // a hex escape (\x01ab) would read as one character
		"°C":          `"\302\260C"`,
		"line\nnext":  `"line\nnext"`,
	}
	for in, want := range cases {
		if got := cString(in); got != want {
			t.Errorf("cString(%q) = %s, want %s", in, got, want)
		}
	}
	cpp := generate(t, `board uno
at start:
    print "Ready?? \"yes\" \\ done"
`)
	if !strings.Contains(cpp, `"Ready\?\? \"yes\" \\ done"`) {
		t.Errorf("print text not escaped:\n%s", cpp)
	}
}

// A 'use "…"' path ends up in an #include: it must not reach files outside the project (a server's config, say),
// nor carry characters the preprocessor would act on.
func TestUsePathsStayInsideTheProject(t *testing.T) {
	refused := []string{"../secret.h", "lib/../../config.php", "/etc/passwd", "C:/xampp/private/config.php", `C:\x.h`,
		"./crc.h", "a//b.h", "x.h> /*", "ok.h\" bad", ""}
	for _, p := range refused {
		src := "board uno\nuse \"" + strings.ReplaceAll(strings.ReplaceAll(p, `\`, `\\`), `"`, `\"`) + "\"\nat start:\n    pass\n"
		_, diags := CheckSource(src)
		if !hasDiag(diags, "C324") {
			t.Errorf("use %q: expected C324, got %v", p, diags)
		}
	}
	for _, p := range []string{"lib/crc.h", "Adafruit NeoPixel", "bblanchon/ArduinoJson", "DHT sensor library", "my-lib+extra.h"} {
		_, diags := CheckSource("board uno\nuse \"" + p + "\"\nat start:\n    pass\n")
		if hasDiag(diags, "C324") {
			t.Errorf("use %q: refused, but it is a fine name", p)
		}
	}
	for _, p := range []string{"../x", "C:/x"} {
		_, diags := CheckSource("board esp32\nuse idf \"" + p + "\"\nat start:\n    pass\n")
		if !hasDiag(diags, "C324") {
			t.Errorf("use idf %q: expected C324", p)
		}
	}
}

// A secret may only be written where it is used (spec 11.1): never copied into a variable, joined into text, printed,
// logged or sent - so its value can't leak by any route the checker doesn't see.
func TestSecretsOnlyWhereTheyAreUsed(t *testing.T) {
	ok := `board esp32
use mujard.mqtt
at start:
    wifi.connect(secret("home_wifi"))
    mqtt.connect("broker", user: "bin7", password: secret("mqtt"))
`
	if _, diags := CheckSource(ok); FirstError(diags) != nil {
		t.Fatalf("the allowed uses were refused: %v", diags)
	}
	leaks := map[string]string{
		"print":       `print secret("mqtt")`,
		"in text":     `print "pw {secret("mqtt")}"`,
		"joined":      `var s = "pw: " + secret("mqtt")`,
		"kept in var": `var s = secret("mqtt")`,
		"log":         `log.info secret("mqtt")`,
		"published":   `mqtt.publish("t", secret("mqtt"))`,
		"as a broker": `mqtt.connect(secret("mqtt"))`,
	}
	for name, line := range leaks {
		src := "board esp32\nuse mujard.mqtt\nat start:\n    " + line + "\n"
		if _, diags := CheckSource(src); !hasDiag(diags, "C323") {
			t.Errorf("%s: expected C323, got %v", name, diags)
		}
	}
	for _, bad := range []string{`secret(name)`, `secret("wifi {x}")`, `secret("1st")`, `secret("has space")`} {
		src := "board esp32\nvar name = \"a\"\nvar x = 1\nat start:\n    wifi.connect(" + bad + ")\n"
		if _, diags := CheckSource(src); !hasDiag(diags, "C325") {
			t.Errorf("%s: expected C325, got %v", bad, diags)
		}
	}
	if _, diags := CheckSource("board esp32\nat start:\n    wifi.connect(\"MyNet\")\n"); !hasDiag(diags, "C303") {
		t.Errorf("wifi.connect with a plain text: expected C303 (credentials are secrets)")
	}
}

// The generated C++ names the secrets but never holds their values; a missing one stops the build on its .muj line.
func TestSecretsBecomeMacros(t *testing.T) {
	sk, diags := Compile(`board esp32
at start:
    wifi.connect(secret("home_wifi"))
`, "p.muj")
	if e := FirstError(diags); e != nil {
		t.Fatal(e)
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{`#include "mujard_secrets.h"`, "#ifndef HOME_WIFI_SSID", "#ifndef HOME_WIFI_PASSWORD",
		"mujard::wifi_link.connect(HOME_WIFI_SSID, HOME_WIFI_PASSWORD)"} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %q in\n%s", want, cpp)
		}
	}
	if len(sk.Secrets) != 2 || sk.Secrets[0] != "HOME_WIFI_PASSWORD" || sk.Secrets[1] != "HOME_WIFI_SSID" {
		t.Errorf("Sketch.Secrets = %v", sk.Secrets)
	}
	if _, has := sk.Files[secretsHeader]; has {
		t.Error("the generated files must not contain the secrets header (only the build writes it)")
	}
}

// mqtt.publish and http.post send any value as text (sample 37 publishes a ratio).
func TestNetworkValuesBecomeText(t *testing.T) {
	cpp := generate(t, `board esp32
use mujard.mqtt
level = analog 34
at start:
    mqtt.publish("level", level.percent)
    mqtt.publish("count", 42)
    mqtt.publish("state", "on")
`)
	for _, want := range []string{`mujard::mqtt_rt.publish(mujard::Text("level"), mujard::Text().add_qty(`,
		`mujard::mqtt_rt.publish(mujard::Text("count"), mujard::Text().add((int32_t)(42)))`,
		`mujard::mqtt_rt.publish(mujard::Text("state"), "on")`} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %s in\n%s", want, cpp)
		}
	}
}

// ask inside a task: the question is sent and the task waits at that point; elsewhere it is refused (C320), and
// it is never asked where the program wouldn't evaluate it (the right side of and/or, the branches of a if c else b).
func TestAskWaitsInTheTask(t *testing.T) {
	cpp := generate(t, `board uno
led = pin 13 as output
task check:
    var ready = ask "Ready?"
    if ready != none:
        led.write(ready)
    var again = ask("Again?", timeout: 30 s)
`)
	for _, want := range []string{`mujard::app_ask_begin(1, mujard::Text("Ready\?"), 0);`, "MUJARD_WAIT_UNTIL(mujard::app_ask_done(1), 1);",
		"mujard::app_ask_result(1)", `mujard::app_ask_begin(2, mujard::Text("Again\?"), ((mujard::Time)30000000LL));`,
		"mujard::app_begin(\"prog.muj\""} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %s in\n%s", want, cpp)
		}
	}
	_, diags := Compile("board uno\nvar busy = true\ntask t:\n    if busy or ask \"x\" != none:\n        pass\n", "p.muj")
	if !hasDiag(diags, "C320") {
		t.Errorf("ask on the right of 'or': expected C320 (it would ask even when 'busy' is true), got %v", diags)
	}
}

// control values are registered with their range in the display unit, and only the dashboard's own programs say hello.
func TestControlsAndHello(t *testing.T) {
	cpp := generate(t, `board uno
var speed = 50 %
var lamp = false
control "Speed", speed, 0 %..100 %
control "Lamp", lamp
`)
	for _, want := range []string{`mujard::app_control("Speed", false,`, `"%", 0.01f, &speed, mujard::APP_FLOAT);`,
		`mujard::app_control("Lamp", true, 0, 1, 1, "", 1, &lamp, mujard::APP_BOOL);`, "mujard::app_begin("} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing %s in\n%s", want, cpp)
		}
	}
	if strings.Contains(generate(t, "board uno\nat start:\n    print \"x\"\n"), "app_begin") {
		t.Error("a program without app built-ins must not talk to the app")
	}
}

// A hold button and a joystick are held controls: registered like a switch / two sliders, then marked held so the
// board puts them back at rest when the app stops refreshing them (the failsafe for anything that drives).
func TestHeldControls(t *testing.T) {
	cpp := generate(t, `board esp32
var horn = false
var steer = 0 %
var throttle = 0 %
control "Horn", horn, hold
control "Drive", steer, throttle, -100 %..100 %
`)
	// each mark follows its own control (the #line lines between them don't matter)
	at := 0
	for _, want := range []string{
		`mujard::app_control("Horn", true, 0, 1, 1, "", 1, &horn, mujard::APP_BOOL);`, "mujard::app_control_hold(mujard::APP_HOLD_BUTTON);",
		"&steer, mujard::APP_FLOAT);", "mujard::app_control_hold(mujard::APP_HOLD_STICK_X);",
		"&throttle, mujard::APP_FLOAT);", "mujard::app_control_hold(mujard::APP_HOLD_STICK_Y);",
	} {
		i := strings.Index(cpp[at:], want)
		if i < 0 {
			t.Fatalf("missing %s (in order) in\n%s", want, cpp)
		}
		at += i + len(want)
	}
	for _, bad := range []string{
		"board esp32\nvar speed = 0 %\ncontrol \"Speed\", speed, hold\n",                          // hold is for true/false
		"board esp32\nvar lamp = false\nvar x = 0 %\ncontrol \"Drive\", lamp, x, -100 %..100 %\n", // a stick's axes are numbers
	} {
		if _, diags := Compile(bad, "p.muj"); !hasDiag(diags, "C303") {
			t.Errorf("expected C303 for\n%s\ngot %v", bad, diags)
		}
	}
}

// Time zones are checked and converted at compile time (C326).
func TestTimeZones(t *testing.T) {
	for zone, want := range map[string]string{"Africa/Lagos": `"WAT-1"`, "Europe/London": `"GMT0BST,M3.5.0/1,M10.5.0"`,
		"WAT-1": `"WAT-1"`, "<+0545>-5:45": `"<+0545>-5:45"`} {
		cpp := generate(t, "board esp32\nat start:\n    clock.sync(timezone: \""+zone+"\")\n")
		if !strings.Contains(cpp, "mujard::clock_rt.sync("+want+")") {
			t.Errorf("%s: want sync(%s) in\n%s", zone, want, cpp)
		}
	}
	for _, bad := range []string{"Africa/Lagoss", "Mars/Olympus", "+1"} {
		if _, diags := CheckSource("board esp32\nat start:\n    clock.sync(timezone: \"" + bad + "\")\n"); !hasDiag(diags, "C326") {
			t.Errorf("%s: expected C326, got %v", bad, diags)
		}
	}
	if _, diags := CheckSource("board uno\nat 07:30:\n    pass\n"); !hasDiag(diags, "B409") {
		t.Error("at HH:MM on a board without Wi-Fi: expected B409")
	}
}

// #line markers carry the file name: a name with a quote must not end the marker early.
func TestLineMarkerFileNameIsEscaped(t *testing.T) {
	sk, diags := Compile("board uno\nat start:\n    print \"x\"\n", `we"ird.muj`)
	if e := FirstError(diags); e != nil {
		t.Fatal(e)
	}
	if !strings.Contains(sk.Files["mujard_program.cpp"], `#line 2 "we\"ird.muj"`) {
		t.Errorf("file name not escaped in #line:\n%s", sk.Files["mujard_program.cpp"])
	}
}
