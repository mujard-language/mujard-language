package lang

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func TestOkSamplesCheck(t *testing.T) {
	for _, s := range loadSamples(t, "ok") {
		_, diags := CheckSourceIn(s.src, s.dir)
		for _, d := range diags {
			if d.IsWarning() {
				t.Logf("%s: warning %s", s.name, d)
			} else {
				t.Errorf("%s: unexpected %s", s.name, d)
			}
		}
	}
}

// A check tells the editors which secret names the build will need - names only, never values.
func TestCheckReportsSecretNames(t *testing.T) {
	src := "board esp32\nuse mujard.mqtt\nat start:\n    wifi.connect(secret(\"home_wifi\"))\n" +
		"    mqtt.connect(\"broker.example.com\", user: \"bin7\", password: secret(\"mqtt\"))\n"
	r := CheckProgramModules(src, nil)
	if e := FirstError(r.Diags); e != nil {
		t.Fatalf("unexpected %s", e)
	}
	if got := strings.Join(r.Secrets, ","); got != "HOME_WIFI_PASSWORD,HOME_WIFI_SSID,MQTT" {
		t.Errorf("secrets = %q", got)
	}
	if r := CheckProgramModules("board esp32\nat start:\n    print 1\n", nil); len(r.Secrets) != 0 {
		t.Errorf("a program without secret() needs none, got %v", r.Secrets)
	}
	// A program that doesn't check clean reports no names (the build would never get that far).
	if r := CheckProgramModules("board esp32\nat start:\n    wifi.connect(secret(\"home_wifi\"))\n    print nope\n", nil); r.Secrets != nil {
		t.Errorf("a failing check must report no secrets, got %v", r.Secrets)
	}
}

func TestBadSamplesChecker(t *testing.T) {
	n := 0
	for _, s := range loadSamples(t, "bad") {
		if strings.ContainsRune("LP", rune(s.code[0])) {
			continue
		}
		n++
		_, diags := CheckSourceIn(s.src, s.dir)
		e := FirstError(diags)
		if e == nil {
			t.Errorf("%s: expected %s at %s, got no error", s.name, s.code, s.pos)
			continue
		}
		if e.Code != s.code || e.Pos != s.pos {
			t.Errorf("%s: expected %s at %s, got %s (all: %v)", s.name, s.code, s.pos, e, diags)
		}
		if strings.Contains(e.Message, "{") {
			t.Errorf("%s: message not filled in: %q", s.name, e.Message)
		}
	}
	if n != 12 {
		t.Errorf("expected 12 checker samples (C, U, B), found %d", n)
	}
}

// The catalogue in the compiler and the table in the spec must list the same codes with the same words.
func TestCatalogueMatchesSpec(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("..", "..", "..", "docs", "lang", "MUJARD_LANGUAGE_SPEC.md"))
	if err != nil {
		t.Fatal(err)
	}
	rx := regexp.MustCompile(`(?m)^\| ([LPCBUW]\d{3}) \| (.+) \|$`)
	spec := map[string]string{}
	for _, m := range rx.FindAllStringSubmatch(string(b), -1) {
		spec[m[1]] = m[2]
	}
	var codes []string
	for code := range catalogue {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	for _, code := range codes {
		want, ok := spec[code]
		if !ok {
			t.Errorf("%s is in the compiler but not in the spec's catalogue", code)
			continue
		}
		if want != catalogue[code] {
			t.Errorf("%s differs:\n spec: %s\n code: %s", code, want, catalogue[code])
		}
	}
	for code := range spec {
		if _, ok := catalogue[code]; !ok {
			t.Errorf("%s is in the spec but the compiler never reports it", code)
		}
	}
}

func check(t *testing.T, src string) []Diagnostic {
	t.Helper()
	_, d := CheckSource(src)
	return d
}

func TestCheckerRules(t *testing.T) {
	cases := []struct {
		name, src string
		code      string // "" = no error expected
		pos       Pos
	}{
		{"none narrowed after an early return", "board uno\necho = pin 10 as input\nfn d() -> time:\n    var t = echo.measure_pulse()\n    if t == none:\n        return 0 ms\n    return t * 2\nat start:\n    print d()\n", "", Pos{}},
		{"const can't change", "board uno\nconst A = 1\nat start:\n    A = 2\n", "C306", Pos{4, 5}},
		{"break outside a loop", "board uno\nat start:\n    break\n", "C310", Pos{3, 5}},
		{"return outside a function", "board uno\nat start:\n    return\n", "C309", Pos{3, 5}},
		{"match must cover an enum", "board uno\nenum M: a, b\nvar m = M.a\nat start:\n    match m:\n        M.a:\n            print 1\n", "C315", Pos{5, 5}},
		{"missing return", "board uno\nfn f(x: int) -> int:\n    if x > 1:\n        return 1\nat start:\n    print f(2)\n", "C313", Pos{2, 1}},
		{"wrong argument count", "board uno\nfn f(x: int):\n    print x\nat start:\n    f(1, 2)\n", "C307", Pos{5, 6}},
		{"ask only in a task", "board esp32\nat start:\n    if ask \"Go?\" == true:\n        print 1\n", "C311", Pos{3, 8}},
		{"test-only outside a test", "board uno\nat start:\n    advance(1 s)\n", "C317", Pos{3, 12}},
		{"library not used", "board uno\nlid = Servo(pin 6)\n", "C318", Pos{2, 7}},
		// spec 6.11: 'safety strict' makes the unchecked/unsafe warnings errors; 'standard' (and no line) leaves them warnings
		{"safety strict refuses raw C++", "board uno\nsafety strict\ncpp:\n    int x;\n", "C328", Pos{3, 1}},
		{"safety standard keeps it a warning", "board uno\nsafety standard\ncpp:\n    int x;\n", "", Pos{}},
		{"a blank line before the safety line", "board uno\n\nsafety strict\ncpp:\n    int x;\n", "C328", Pos{4, 1}},
		{"an unknown safety level", "board uno\nsafety loose\n", "C329", Pos{2, 1}},
		{"a strict program without warnings is fine", "board uno\nsafety strict\nled = pin LED as output\nevery 1 s:\n    led.toggle()\n", "", Pos{}},
		{"a device called safety is still a device", "board uno\nsafety = pin 3 as output\nat start:\n    safety.on()\n", "", Pos{}},
		// W617: recursion takes stack RAM the estimate can't see; under safety strict it is an error
		{"recursion under safety strict", "board uno\nsafety strict\nfn f(n: int) -> int:\n    if n < 1:\n        return 0\n    return f(n - 1)\nat start:\n    print f(3)\n", "C328", Pos{3, 1}},
		{"wifi on an uno", "board uno\nat start:\n    wifi.connect(secret(\"home\"))\n", "B409", Pos{3, 10}},
		{"flash pin", "board esp32\nx = pin 6 as output\n", "B406", Pos{2, 9}},
		{"analog-only pin as digital", "board nano\nx = pin A6 as input\n", "B408", Pos{2, 9}},
		{"not an analog pin", "board uno\nx = analog 7\n", "B404", Pos{2, 12}},
		{"no such pin", "board uno\nx = pin 99 as output\n", "B402", Pos{2, 9}},
		{"pullup on an output", "board uno\nx = pin 5 as output pullup\n", "B410", Pos{2, 14}},
		{"fast on a pin without an interrupt", "board uno\nb = pin 5 as input pullup\non b.pressed fast:\n    print 1\n", "B411", Pos{3, 1}},
		{"unknown board", "board toaster\n", "B401", Pos{1, 1}},
		{"wrong dimension", "board uno\ntask t:\n    wait 5 cm\n", "U503", Pos{3, 10}},
		{"speed times time is a length", "board uno\nat start:\n    var d = 2 m/s * 3 s\n    var e: length = d\n    print e\n", "", Pos{}},
		{"every with a frequency", "board uno\nled = pin 13 as output\nevery 2 Hz:\n    led.toggle()\n", "", Pos{}},
		{"zero fits any unit", "board uno\nat start:\n    var t = 5 ms\n    if t > 0:\n        print t\n", "", Pos{}},
		{"text plus number", "board uno\nat start:\n    print \"a\" + 1\n", "C303", Pos{3, 17}},
		{"duplicate name", "board uno\nvar a = 1\nvar a = 2\nat start:\n    print a\n", "C302", Pos{3, 1}},
		{"unknown member", "board uno\nled = pin 13 as output\nat start:\n    led.blinky()\n", "C308", Pos{4, 9}},
		{"condition must be true/false", "board uno\nat start:\n    if 3:\n        print 1\n", "C303", Pos{3, 8}},
		{"bare enum member in context", "board uno\nenum Mode: idle, run\nvar m: Mode = idle\nat start:\n    if m == run:\n        print m\n", "", Pos{}},
		{"bare enum member in a match", "board uno\nenum Mode: idle, run\nvar m = Mode.idle\nat start:\n    match m:\n        idle:\n            print 1\n        run:\n            print 2\n", "", Pos{}},
		{"bare enum member as an argument", "board uno\nenum Mode: idle, run\nfn go(m: Mode):\n    pass\nat start:\n    go(idle)\n", "", Pos{}},
		{"bare enum member without context", "board uno\nenum Mode: idle, run\nat start:\n    print idle\n", "C301", Pos{4, 11}},
		{"a name in scope wins over an enum member", "board uno\nenum Mode: idle, run\nvar idle = 3\nvar m: Mode = idle\n", "C303", Pos{4, 15}},
		{"number gets its unit with as", "board uno\nat start:\n    var d = 5 as cm\n    var e: length = d\n    print e\n", "", Pos{}},
		{"secret can't be printed", "board esp32\nat start:\n    print secret(\"home\")\n", "C323", Pos{3, 11}},
		{"secret can't be notified", "board esp32\nat start:\n    notify secret(\"home\")\n", "C323", Pos{3, 12}},
		{"secret hidden in text", "board esp32\nat start:\n    print \"key: {secret(\"home\")}\"\n", "C323", Pos{3, 18}},
		{"secret can be passed on", "board esp32\nat start:\n    wifi.connect(secret(\"home\"))\n", "", Pos{}},
		{"remainder is the word mod", "board uno\nvar hue = 0\nevery 20 ms:\n    hue = (hue + 256) mod 65536\n", "", Pos{}},
		{"expect without a message", "board uno\nvar t = 1\nat start:\n    expect t > 0\n", "", Pos{}},
		{"expect that keeps running", "board uno\nvar t = 1\nat start:\n    expect t > 0, \"too hot\", halt: false\n", "", Pos{}},
		{"expect halt must be true or false", "board uno\nvar t = 1\nat start:\n    expect t > 0, halt: 5\n", "C303", Pos{4, 25}},
		{"expect unknown option", "board uno\nvar t = 1\nat start:\n    expect t > 0, retry: 3\n", "C308", Pos{4, 26}},
		{"expect takes one message", "board uno\nvar t = 1\nat start:\n    expect t > 0, \"a\", \"b\"\n", "C307", Pos{4, 24}},
		// fast handlers (spec 6.5)
		{"fast can't print", "board uno\nb = pin 2 as input\nvar n = 0\non b.changed fast:\n    n += 1\n    print n\n", "C314", Pos{6, 5}},
		{"fast can't call a function that prints", "board uno\nb = pin 2 as input\nfn f():\n    print 1\non b.rising fast:\n    f()\n", "C314", Pos{6, 5}},
		{"fast can't change a saved var", "board uno\nb = pin 2 as input\nsaved var n = 0\non b.falling fast:\n    n += 1\n", "C314", Pos{5, 5}},
		{"fast can't read an analog pin", "board uno\nb = pin 2 as input\nlight = analog A0\nvar v = 0\non b.changed fast:\n    v = light.value\n", "C314", Pos{6, 9}},
		{"fast can't build text", "board uno\nb = pin 2 as input\nvar s = \"\"\non b.changed fast:\n    s = \"hit\"\n", "C314", Pos{5, 5}},
		{"fast can't loop with while", "board uno\nb = pin 2 as input\nvar n = 0\non b.changed fast:\n    while n > 0:\n        n -= 1\n", "C314", Pos{5, 5}},
		{"fast halt needs plain text", "board uno\nb = pin 2 as input\nvar n = 0\non b.changed fast:\n    board.halt(\"n is {n}\")\n", "C314", Pos{5, 16}},
		{"what fast may do", "board uno\nb = pin 2 as input pullup\nled = pin 13 as output\nvar n = 0\nvar last = now()\nfn bump(x: int) -> int:\n    return clamp(x + 1, 0..10)\non b.rising fast:\n    n = bump(n)\n    last = now()\n    for i in 0..3:\n        if n.bit(i):\n            led.toggle()\n    if n > 9:\n        board.halt(\"too many\")\nevery 1 s:\n    print \"{n} {since(last)}\"\n", "", Pos{}},
		{"rising and falling without fast", "board uno\nb = pin 2 as input\non b.rising:\n    print 1\non b.falling:\n    print 0\n", "", Pos{}},
		// built-in enums (spec 4.6)
		{"reset reason isn't a number", "board esp32\nat start:\n    if board.reset_reason == 3:\n        print 1\n", "C303", Pos{3, 30}},
		{"level isn't a number", "board uno\ne = pin 7 as input\nat start:\n    print e.measure_pulse(level: 1)\n", "C303", Pos{4, 34}},
		{"level by its bare name and in full", "board uno\ne = pin 7 as input\nat start:\n    print e.measure_pulse(level: high)\n    print e.measure_pulse(level: Level.low)\n", "", Pos{}},
		{"a program enum wins over a built-in member", "board uno\nenum Speed: low, high\nvar s = Speed.low\nat start:\n    if s == high:\n        print 1\n", "", Pos{}},
		{"battery type in full", "board uno\ncell = battery_on(analog A1, divider: 2, type: BatteryType.aa_x3)\nevery 1 s:\n    print cell.percent\n", "", Pos{}},
		{"battery type isn't text", "board uno\ncell = battery_on(analog A1, divider: 2, type: \"aa\")\nevery 1 s:\n    print cell.percent\n", "C303", Pos{2, 48}},
		{"a smaller fixed text size", "board uno\nvar t: text[8] = \"hi\"\nat start:\n    print t\n", "", Pos{}},
		{"a bad text size", "board uno\nvar t: text[0] = \"hi\"\n", "C303", Pos{2, 13}},
		{"a text letter can't change", "board uno\nvar t = \"abc\"\nat start:\n    t[0] = \"x\"\n", "C306", Pos{4, 5}},
		// What the generator refuses is a check error too, so a program that checks clean also builds.
		{"the generator's refusals reach the checker", "board mega\nat start:\n    print board.reset_reason\n", "C320", Pos{3, 17}},
	}
	for _, c := range cases {
		d := check(t, c.src)
		e := FirstError(d)
		switch {
		case c.code == "" && e != nil:
			t.Errorf("%s: unexpected %s", c.name, e)
		case c.code != "" && e == nil:
			t.Errorf("%s: expected %s at %s, got nothing (%v)", c.name, c.code, c.pos, d)
		case c.code != "" && (e.Code != c.code || e.Pos != c.pos):
			t.Errorf("%s: expected %s at %s, got %s", c.name, c.code, c.pos, e)
		}
	}
}

func TestCheckerWarnings(t *testing.T) {
	cases := []struct {
		name, src, code string
	}{
		{"serial pins", "board uno\nx = pin 1 as output\nat start:\n    x.on()\n", "W601"},
		{"strapping pin", "board esp32\nx = pin 0 as input\nat start:\n    print x.read()\n", "W602"},
		{"unused variable", "board uno\nat start:\n    var x = 1\n", "W603"},
		{"cpp block", "board uno\ncpp:\n    int a;\n", "W604"},
		{"while true outside a task", "board uno\nat start:\n    while true:\n        break\n", "W610"},
		{"text longer than a text holds", "board uno\nat start:\n    print \"" + strings.Repeat("x", 65) + "\"\n", "W608"},
		{"the board's I2C pins", "board uno\nuse mujard.i2c\nsda = pin A4 as input\nat start:\n    print sda.read()\n    print i2c.scan()\n", "W609"},
		{"an I2C library's pins", "board esp32\nuse \"LiquidCrystal I2C\"\nx = pin GPIO21 as output\nat start:\n    x.on()\n", "W609"},
		{"PWM on the Servo timer", "board uno\nuse mujard.servo\nlid = Servo(pin 9)\nlamp = pin 10 as pwm\nat start:\n    lid.move(90 deg)\n    lamp.write(50 %)\n", "W613"},
	}
	for _, c := range cases {
		found := false
		for _, d := range check(t, c.src) {
			if d.Code == c.code {
				found = true
			}
			if !d.IsWarning() {
				t.Errorf("%s: unexpected error %s", c.name, d)
			}
		}
		if !found {
			t.Errorf("%s: expected warning %s", c.name, c.code)
		}
	}
}

// TestBoardClaims: which pins a library takes over depends on the board (spec 7, boards.json claims).
func TestBoardClaims(t *testing.T) {
	servoAndPwm := func(board, servo, pwm string) string {
		return "board " + board + "\nuse mujard.servo\nlid = Servo(pin " + servo + ")\nlamp = pin " + pwm + " as pwm\n" +
			"at start:\n    lid.move(90 deg)\n    lamp.write(50 %)\n"
	}
	has := func(d []Diagnostic, code string) bool {
		for _, x := range d {
			if x.Code == code {
				return true
			}
		}
		return false
	}
	if !has(check(t, servoAndPwm("mega", "9", "45")), "W613") {
		t.Error("a Mega's Servo takes Timer5: PWM on 45 should warn")
	}
	if has(check(t, servoAndPwm("mega", "9", "10")), "W613") {
		t.Error("a Mega's Servo leaves 10's timer alone")
	}
	if has(check(t, servoAndPwm("esp32", "GPIO13", "GPIO25")), "W613") {
		t.Error("an ESP32's servos don't take a PWM pin")
	}
	if has(check(t, "board uno\nlamp = pin 10 as pwm\nat start:\n    lamp.write(50 %)\n"), "W613") {
		t.Error("no Servo, no warning")
	}
	// tone() takes Timer2 on the Uno (PWM 3 and 11) and Timer3 on the Leonardo (5)
	toneAndPwm := func(board, pwm string) string {
		return "board " + board + "\nbuzzer = pin 8 as output\nlamp = pin " + pwm + " as pwm\n" +
			"at start:\n    buzzer.tone(440 Hz, 100 ms)\n    lamp.write(50 %)\n"
	}
	if !has(check(t, toneAndPwm("uno", "3")), "W613") || !has(check(t, toneAndPwm("leonardo", "5")), "W613") {
		t.Error("tone takes a PWM pin's timer on the Uno and the Leonardo")
	}
	if has(check(t, toneAndPwm("uno", "5")), "W613") || has(check(t, toneAndPwm("esp32", "GPIO25")), "W613") {
		t.Error("tone leaves the Uno's pin 5 and the ESP32's pins alone")
	}
	melody := "board nano\nbuzzer = pin 8 as output\nfan = pin 11 as pwm\nat start:\n    buzzer.melody(\"C4 E4 G4\")\n    fan.write(20 %)\n"
	if !has(check(t, melody), "W613") {
		t.Error("melody plays through tone() too")
	}
	// the message names the type without a second article
	d := check(t, "board uno\nlamp = pin 3 as pwm\non lamp.changed:\n    pass\n")
	if e := FirstError(d); e == nil || e.Message != "A pin (pwm) has no 'changed'." {
		t.Errorf("C308 message: got %v", e)
	}
}

// TestWifiADC: the ESP32 family's ADC2 pins stop reading while Wi-Fi is on (spec 7, boards.json
// adcBlockedByWifi): W615 fires only when the program actually uses Wi-Fi.
func TestWifiADC(t *testing.T) {
	has := func(d []Diagnostic, code string) bool {
		for _, x := range d {
			if x.Code == code {
				return true
			}
		}
		return false
	}
	prog := func(board, pin, body string) string {
		return "board " + board + "\npot = analog " + pin + "\nat start:\n    print(pot.value)\n" + body
	}
	wifi := "every 10 s:\n    wifi.connect(secret(\"home\"))\n"
	if !has(check(t, prog("esp32", "GPIO25", wifi)), "W615") {
		t.Error("GPIO 25 is ADC2 on the ESP32: Wi-Fi blocks it")
	}
	if has(check(t, prog("esp32", "GPIO25", "")), "W615") {
		t.Error("no Wi-Fi use, no W615")
	}
	if has(check(t, prog("esp32", "GPIO32", wifi)), "W615") {
		t.Error("GPIO 32 is ADC1: Wi-Fi doesn't touch it")
	}
	if !has(check(t, prog("esp32s3", "GPIO15", wifi)), "W615") {
		t.Error("the ESP32-S3's ADC2 pins block the same way")
	}
	if has(check(t, prog("esp32c6", "GPIO3", wifi)), "W615") {
		t.Error("the C6 has no ADC2: nothing is blocked")
	}
	// mqtt and http turn the radio on too, and clock.sync needs it
	if !has(check(t, prog("esp32", "GPIO25", "every 10 s:\n    clock.sync()\n")), "W615") {
		t.Error("clock.sync brings Wi-Fi up: the ADC2 pin warns")
	}
}

// TestRegs: 'use mujard.regs' types register and bit names per chip (spec 10.4): a misspelt name is C308, a
// literal too wide for the register is C327, a write onto a declared pin's port is W616, and a board whose chip
// has no table is B413.
func TestRegs(t *testing.T) {
	has := func(d []Diagnostic, code string) bool {
		for _, x := range d {
			if x.Code == code {
				return true
			}
		}
		return false
	}
	ok := check(t, "board uno\nuse mujard.regs\nat start:\n    regs.DDRB.write(0x20)\n    regs.PORTB.set(regs.PORTB.PORTB5)\n    var p = regs.PINB.read()\n    print(p)\n")
	if has(ok, "C308") || has(ok, "C327") {
		t.Error("a correct regs program should check clean")
	}
	if !has(check(t, "board uno\nuse mujard.regs\nat start:\n    regs.PORTX.write(1)\n"), "C308") {
		t.Error("an unknown register should be C308")
	}
	if !has(check(t, "board uno\nuse mujard.regs\nat start:\n    regs.PORTB.write(999)\n"), "C327") {
		t.Error("999 doesn't fit an 8-bit register: C327")
	}
	if !has(check(t, "board uno\nuse mujard.regs\nat start:\n    print(regs.PORTB.NOPE)\n"), "C308") {
		t.Error("an unknown bit field should be C308")
	}
	if !has(check(t, "board uno\nuse mujard.regs\nled = pin 8 as output\nat start:\n    regs.PORTB.write(0xFF)\n    led.on()\n"), "W616") {
		t.Error("writing PORTB reaches declared pin 8: W616")
	}
	if has(check(t, "board uno\nuse mujard.regs\nled = pin 8 as output\nat start:\n    regs.TCCR0A.write(0)\n    led.on()\n"), "W616") {
		t.Error("TCCR0A drives no declared pin: no W616")
	}
	if !has(check(t, "board uno\nat start:\n    regs.PORTB.write(1)\n"), "C318") {
		t.Error("regs without 'use mujard.regs' should be C318")
	}
	// every shipped chip has a table: pico (SIO) and bluepill (GPIOA_ODR) check clean
	if has(check(t, "board pico\nuse mujard.regs\nat start:\n    regs.SIO_GPIO_OUT_SET.write(0x20)\n"), "B413") {
		t.Error("the RP2040 has an SVD table now")
	}
	if has(check(t, "board bluepill\nuse mujard.regs\nat start:\n    regs.GPIOA_ODR.write(1)\n"), "B413") {
		t.Error("the STM32F103 has an SVD table now")
	}
	if has(check(t, "board esp32\nuse mujard.regs\nat start:\n    regs.GPIO_OUT.write(1)\n"), "B413") {
		t.Error("the ESP32 has an SVD table now")
	}
}

// TestRamEstimate: W614 when what the program keeps in RAM is most of the board's (spec 7).
func TestRamEstimate(t *testing.T) {
	has := func(d []Diagnostic, code string) *Diagnostic {
		for i := range d {
			if d[i].Code == code {
				return &d[i]
			}
		}
		return nil
	}
	big := "board uno\nvar names: list[text, 20] = []\nat start:\n    names.add(\"a\")\n    print names.length\n"
	d := has(check(t, big), "W614")
	if d == nil || !strings.Contains(d.Message, "about 1324 bytes of the Arduino Uno's 2048 bytes") {
		t.Errorf("20 texts on an Uno: %v", d) // 20 x 66 + 4
	}
	if has(check(t, strings.Replace(big, "uno", "mega", 1)), "W614") != nil {
		t.Error("the same list fits a Mega's 8 KB")
	}
	if has(check(t, strings.Replace(big, "20]", "4]", 1)), "W614") != nil {
		t.Error("4 texts are fine on an Uno")
	}
	smooth := "board uno\nlight = analog A0\nvar a = 0.0\nevery 10 ms:\n" + strings.Repeat("    a = smooth(light.value, 32)\n", 10)
	if has(check(t, smooth), "W614") == nil {
		t.Error("ten smooth() call sites keep 1300 bytes on an Uno")
	}
	for _, s := range loadSamples(t, "ok") {
		if d := has(check(t, s.src), "W614"); d != nil {
			t.Errorf("%s: a sample shouldn't need most of the RAM: %s", s.name, d.Message)
		}
	}
}

// TestModules: 'import name' loads name.muj; its top-level names are reached as name.member.
func TestModules(t *testing.T) {
	dir := t.TempDir()
	write := func(name, src string) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(src), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("motors.muj", "board uno\nconst top_speed = 800\nvar armed = true\nfan = pin 9 as pwm\nenum Run: stopped, going\nfn faster(v: int) -> int:\n    return v + 1\nfn ready() -> bool:\n    return armed\n")
	prog := "board uno\nimport motors\nat start:\n    print motors.faster(3) + motors.top_speed\n    motors.fan.write(50 %)\n    match motors.Run.going:\n        motors.Run.going:\n            print 1\n        else:\n            pass\n"
	if d := checkIn(t, prog, dir); FirstError(d) != nil {
		t.Errorf("a module's names must be reachable as motors.<name>: %v", d)
	}
	// assigning a module's variable works; so does reading it
	if d := checkIn(t, "board uno\nimport motors\nat start:\n    motors.armed = false\n    print motors.armed\n", dir); FirstError(d) != nil {
		t.Errorf("motors.armed assign/read: %v", d)
	}

	cases := []struct {
		name, src string
		code      string
		pos       Pos
	}{
		{"missing module", "board uno\nimport ghost\n", "C321", Pos{2, 1}},
		{"unknown member", "board uno\nimport motors\nat start:\n    print motors.nope\n", "C308", Pos{4, 18}},
		{"unknown call member", "board uno\nimport motors\nat start:\n    motors.nope()\n", "C308", Pos{4, 12}},
		{"module call args", "board uno\nimport motors\nat start:\n    motors.faster(1, 2)\n", "C307", Pos{4, 18}},
		{"module is not a value", "board uno\nimport motors\nat start:\n    print motors + 1\n", "C303", Pos{4, 11}},
	}
	for _, c := range cases {
		d := checkIn(t, c.src, dir)
		e := FirstError(d)
		if e == nil || e.Code != c.code || e.Pos != c.pos {
			t.Errorf("%s: expected %s at %s, got %v", c.name, c.code, c.pos, d)
		}
	}

	// A module can't import a module.
	write("layer.muj", "import motors\n")
	if d := checkIn(t, "board uno\nimport layer\n", dir); FirstError(d) == nil {
		t.Errorf("a module importing a module must fail (C322)")
	} else {
		found := false
		for _, x := range d {
			if x.Code == "C322" && x.File == "layer.muj" {
				found = true
			}
		}
		if !found {
			t.Errorf("expected C322 tagged to layer.muj, got %v", d)
		}
	}

	// A module's own errors are reported with its file.
	write("broken.muj", "var x = nosuchthing\n")
	d := checkIn(t, "board uno\nimport broken\n", dir)
	found := false
	for _, x := range d {
		if x.File == "broken.muj" {
			found = true
		}
	}
	if !found {
		t.Errorf("the module's own error should carry its file name: %v", d)
	}

	// .mjd is the alternative extension: found when there is no .muj, and its errors carry that name.
	write("pumps.mjd", "const rate = 5\nvar y = nosuchthing\n")
	d = checkIn(t, "board uno\nimport pumps\n", dir)
	found = false
	for _, x := range d {
		if x.File == "pumps.mjd" {
			found = true
		}
	}
	if !found {
		t.Errorf("import must load pumps.mjd and tag its errors with that name: %v", d)
	}
	if !IsSource("a.muj") || !IsSource("B.MJD") || IsSource("c.usbp") || IsSource("d.ino") {
		t.Errorf("IsSource must accept .muj and .mjd only")
	}

	// Without a directory (bare CheckSource) 'import' stays a later-phase feature.
	d = check(t, prog)
	if e := FirstError(d); e == nil || e.Code != "C320" {
		t.Errorf("import with no directory must still report C320, got %v", d)
	}
}

func checkIn(t *testing.T, src, dir string) []Diagnostic {
	t.Helper()
	_, d := CheckSourceIn(src, dir)
	return d
}

func FuzzCheck(f *testing.F) {
	for _, kind := range []string{"ok", "bad"} {
		files, _ := filepath.Glob(filepath.Join(examplesDir, kind, "*.muj"))
		for _, file := range files {
			if b, err := os.ReadFile(file); err == nil {
				f.Add(string(b))
			}
		}
	}
	f.Fuzz(func(t *testing.T, src string) {
		// The checker runs only on programs that parse; it must never crash on any of them.
		file, diags := ParseSource(src)
		if len(diags) == 0 {
			Check(file)
		}
	})
}

// Every board in the database resolves its own LED and analog names, and the generated JSON has the extras.
func TestBoards(t *testing.T) {
	for _, name := range BoardNames() {
		b := LookupBoard(name)
		if b.Display == "" || len(b.ValidPins) == 0 {
			t.Errorf("%s: incomplete", name)
		}
		for pinName := range b.Names {
			if pinName == "LED_BUILTIN" {
				continue
			}
			if _, ok := b.ResolvePin(pinName); !ok {
				t.Errorf("%s: %s doesn't resolve", name, pinName)
			}
		}
	}
	esp := LookupBoard("esp32")
	if n, ok := esp.ResolvePin("GPIO34"); !ok || n != 34 {
		t.Errorf("GPIO34 -> %d %v", n, ok)
	}
	if n, ok := LookupBoard("esp8266").ResolvePin("D5"); !ok || n != 14 {
		t.Errorf("D5 -> %d %v", n, ok)
	}
	if n, ok := LookupBoard("pico").ResolvePin("GP15"); !ok || n != 15 {
		t.Errorf("GP15 -> %d %v", n, ok)
	}
	if n, ok := LookupBoard("rp2350usba").ResolvePin("GP16"); !ok || n != 16 {
		t.Errorf("rp2350usba GP16 -> %d %v", n, ok)
	}
	if _, ok := LookupBoard("rp2350usba").ResolvePin("LED"); ok {
		t.Error("the RP2350-USB-A has no plain LED (its light is the WS2812 on GP16)")
	}
	if n, ok := LookupBoard("uno").ResolvePin("LED"); !ok || n != 13 {
		t.Errorf("LED -> %d %v", n, ok)
	}
}
