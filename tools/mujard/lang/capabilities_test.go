package lang

import (
	"reflect"
	"strings"
	"testing"
)

// TestCapabilityEvidenceIsConsistent: capabilities.json only names real boards and capabilities, and never claims a
// build or a run of something the board doesn't have.
func TestCapabilityEvidenceIsConsistent(t *testing.T) {
	known := map[string]bool{}
	for _, cp := range Capabilities {
		known[cp.Name] = true
	}
	for half, m := range map[string]map[string]map[string]string{"build": evidence.Build, "hardware": evidence.Hardware} {
		for board, caps := range m {
			b := boards[board]
			if b == nil {
				t.Errorf("%s: unknown board %q", half, board)
				continue
			}
			for cp, note := range caps {
				if !known[cp] {
					t.Errorf("%s %s: unknown capability %q", half, board, cp)
				}
				if CapabilityLevel(b, cp) == LevelNone {
					t.Errorf("%s %s: claims %q, which the board doesn't have", half, board, cp)
				}
				if len(note) < 10 || note[4] != '-' || note[7] != '-' {
					t.Errorf("%s %s %s: the evidence must start with its date: %q", half, board, cp, note)
				}
			}
		}
	}
}

// TestProgramCaps: what a program asks of its board, from its devices, events, calls and libraries.
func TestProgramCaps(t *testing.T) {
	src := "board uno\nuse mujard.servo\nled = pin 13 as output\nlamp = pin 5 as pwm\nknob = analog A0\nbutton = pin 2 as input\n" +
		"buzzer = pin 8 as output\narm = Servo(pin 9)\non button.changed fast:\n    led.toggle()\nat start:\n" +
		"    lamp.write(50 %)\n    print knob.value\n    buzzer.tone(440 Hz, 100 ms)\n    arm.move(90 deg)\n"
	f, diags := ParseSource(src)
	if e := FirstError(diags); e != nil {
		t.Fatal(e)
	}
	info, diags := Analyze(f)
	if e := FirstError(diags); e != nil {
		t.Fatal(e)
	}
	want := []string{"output", "input", "pwm", "analog", "fast", "tone", "servo"}
	if !reflect.DeepEqual(info.Caps, want) {
		t.Errorf("caps: got %v, want %v", info.Caps, want)
	}
}

// TestCapabilityLevels and W618: the level comes from the board's hardware and the evidence, never from code alone.
func TestCapabilityLevels(t *testing.T) {
	saved := evidence
	defer func() { evidence = saved }()
	evidence = &Evidence{
		Build:    map[string]map[string]string{"uno": {"output": "2026-10-30"}},
		Hardware: map[string]map[string]string{"nano": {"output": "2026-10-29: Blink, seen by the owner"}},
	}
	uno, nano, esp32 := boards["uno"], boards["nano"], boards["esp32"]
	for _, tc := range []struct {
		b      *Board
		cp, lv string
	}{
		{nano, "output", LevelHardware}, {uno, "output", LevelBuild}, {uno, "pwm", LevelBackend},
		{uno, "wifi", LevelNone}, {esp32, "wifi", LevelBackend}, {uno, "nonsense", ""},
	} {
		if got := CapabilityLevel(tc.b, tc.cp); got != tc.lv {
			t.Errorf("%s %s: got %q, want %q", tc.b.Name, tc.cp, got, tc.lv)
		}
	}
	w618 := func(src string) []Diagnostic {
		_, diags := CheckSource(src)
		var out []Diagnostic
		for _, d := range diags {
			if d.Code == "W618" {
				out = append(out, d)
			}
		}
		return out
	}
	if d := w618("board uno\nled = pin 13 as output\nevery 1 s:\n    led.toggle()\n"); len(d) != 0 {
		t.Errorf("a built capability doesn't warn: %v", d)
	}
	d := w618("board uno\nled = pin 13 as output\nlamp = pin 5 as pwm\nknob = analog A0\nat start:\n    led.on()\n    lamp.write(50 %)\n    print knob.value\n")
	want := "Mujard hasn't built analog, pwm for the Arduino Uno yet, so this program may not compile there."
	if len(d) != 1 || d[0].Pos != (Pos{1, 1}) || d[0].Message != want {
		t.Errorf("one W618 at the board line naming what isn't built: %v", d)
	}
	if d := w618("board uno\nsafety strict\nlamp = pin 5 as pwm\nat start:\n    lamp.write(50 %)\n"); len(d) != 1 {
		t.Errorf("W618 stays a warning under safety strict (unproven isn't unsafe): %v", d)
	}
	if !strings.Contains(catalogue["W618"], "{list}") {
		t.Error("W618 names the capabilities")
	}
}

// TestMulticoreFoundation: the checker knows which chips really have a second usable core. 'on core N' gives it a
// backend on the 2-core boards (ESP32's pinned tasks, RP2040/RP2350's multicore_launch_core1); the build matrix
// proved it on ESP32, the rest stay unproven until their build matrix run.
func TestMulticoreFoundation(t *testing.T) {
	for name, want := range map[string]int{
		"uno": 1, "esp32": 2, "esp8266": 1, "pico": 2, "rp2350usba": 2,
		"esp32s3": 2, "esp32s2": 1, "esp32c3": 1, "esp32c6": 1,
		"unor4": 1, "nano33ble": 1, "zero": 1, "bluepill": 1,
	} {
		if got := boards[name].Cores(); got != want {
			t.Errorf("%s cores: got %d, want %d", name, got, want)
		}
	}
	for name, want := range map[string]string{"uno": LevelNone, "esp32": LevelBuild, "pico": LevelBackend, "esp32s2": LevelNone} {
		if got := CapabilityLevel(boards[name], "multicore"); got != want {
			t.Errorf("%s multicore: got %q, want %q", name, got, want)
		}
	}
	sk, diags := Compile("board pico\nat start:\n    print board.cores\n", "p.muj")
	if e := FirstError(diags); e != nil || sk == nil {
		t.Fatalf("board.cores should compile: %v", diags)
	}
	if !strings.Contains(sk.Files["mujard_program.cpp"], ".add((int32_t)(2))") {
		t.Errorf("the Pico count should be a compile-time 2:\n%s", sk.Files["mujard_program.cpp"])
	}
	if _, diags := CheckSource("board uno\nneeds multicore\n"); FirstError(diags) == nil || FirstError(diags).Code != "B415" {
		t.Errorf("an Uno doesn't have a second core: %v", diags)
	}
	if _, diags := CheckSource("board pico\nneeds multicore\n"); FirstError(diags) != nil || !hasDiag(diags, "W618") {
		t.Errorf("a Pico has the cores, but no backend proof yet: %v", diags)
	}
}

// TestHardwareFactsFoundation: DMA and timer counts are real board facts a driver can require, while the APIs that
// would use them stay language-level until they exist (the board data comes from hardware.ts -> boards.json).
func TestHardwareFactsFoundation(t *testing.T) {
	for name, want := range map[string][2]int{
		"uno": {0, 3}, "esp32": {1, 4}, "esp8266": {0, 2}, "pico": {1, 4},
		"esp32s3": {1, 4}, "esp32s2": {1, 4}, "esp32c3": {1, 2}, "esp32c6": {1, 2},
		"unor4": {1, 8}, "nano33ble": {1, 5}, "zero": {1, 6}, "bluepill": {1, 4},
	} {
		b := boards[name]
		if (b.DMA != (want[0] == 1)) || b.Timers != want[1] {
			t.Errorf("%s hardware facts: got dma=%v timers=%d, want %d/%d", name, b.DMA, b.Timers, want[0], want[1])
		}
	}
	for _, tc := range []struct{ name, cap, want string }{
		{"uno", "dma", LevelNone}, {"pico", "dma", LevelLanguage}, {"uno", "timers", LevelBackend},
	} {
		if got := CapabilityLevel(boards[tc.name], tc.cap); got != tc.want {
			t.Errorf("%s %s: got %q, want %q", tc.name, tc.cap, got, tc.want)
		}
	}
	sk, diags := Compile("board uno\nat start:\n    print board.dma\n    print board.timers\n", "p.muj")
	if e := FirstError(diags); e != nil || sk == nil {
		t.Fatalf("board facts should compile: %v", diags)
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{".add((bool)(false))", ".add((int32_t)(3))"} {
		if !strings.Contains(cpp, want) {
			t.Errorf("missing compile-time fact %q:\n%s", want, cpp)
		}
	}
	if _, diags := CheckSource("board uno\nneeds dma\n"); FirstError(diags) == nil || FirstError(diags).Code != "B415" {
		t.Errorf("an Uno has no general-purpose DMA: %v", diags)
	}
	if _, diags := CheckSource("board pico\nneeds dma, timers\n"); FirstError(diags) != nil || !hasDiag(diags, "W618") {
		t.Errorf("the hardware exists, but DMA/timer APIs have no backend proof yet: %v", diags)
	}
}

// TestPeripheralRequirementFoundation: CAN, Bluetooth and USB-device requirements name real hardware. CAN has a
// generated backend (mujard.can, unproven until the matrix runs); Bluetooth and USB-device stay
// language-level until Mujard has APIs and generated backends for them.
func TestPeripheralRequirementFoundation(t *testing.T) {
	for _, tc := range []struct{ board, cap, want string }{
		{"uno", "can", LevelNone}, {"esp32", "can", LevelBackend}, {"esp32s2", "can", LevelBackend},
		{"esp8266", "can", LevelNone}, {"pico", "can", LevelNone},
		{"esp32", "bluetooth", LevelBackend}, {"esp32s2", "bluetooth", LevelNone},
		{"esp32s3", "bluetooth", LevelBackend}, {"esp32c3", "bluetooth", LevelBackend}, {"esp32c6", "bluetooth", LevelBackend},
		{"picow", "bluetooth", LevelBackend}, {"nano33ble", "bluetooth", LevelBackend},
		{"uno", "usb_device", LevelNone}, {"leonardo", "usb_device", LevelBackend}, {"esp32", "usb_device", LevelNone},
		{"esp32s2", "usb_device", LevelBackend}, {"esp32s3", "usb_device", LevelBackend},
		{"esp32c3", "usb_device", LevelNone}, {"esp32c6", "usb_device", LevelNone}, {"pico", "usb_device", LevelBackend},
		{"rp2350usba", "usb_device", LevelBackend}, {"unor4", "usb_device", LevelBackend},
		{"nano33ble", "usb_device", LevelBackend}, {"zero", "usb_device", LevelBackend}, {"bluepill", "usb_device", LevelBackend},
	} {
		if got := CapabilityLevel(boards[tc.board], tc.cap); got != tc.want {
			t.Errorf("%s %s: got %q, want %q", tc.board, tc.cap, got, tc.want)
		}
	}
	if _, diags := CheckSource("board uno\nneeds can, bluetooth, usb_device\n"); FirstError(diags) == nil || FirstError(diags).Code != "B415" {
		t.Errorf("an Uno has none of these peripherals: %v", diags)
	}
	if _, diags := CheckSource("board esp32\nneeds can, bluetooth\n"); FirstError(diags) != nil || !hasDiag(diags, "W618") {
		t.Errorf("an ESP32 has CAN and BLE hardware, but the APIs are still unproven: %v", diags)
	}
	if _, diags := CheckSource("board esp32\nneeds usb_device\n"); FirstError(diags) == nil || FirstError(diags).Code != "B415" {
		t.Errorf("a classic ESP32 has no general USB-device peripheral: %v", diags)
	}
	if _, diags := CheckSource("board pico\nneeds usb_device\n"); FirstError(diags) != nil || !hasDiag(diags, "W618") {
		t.Errorf("the RP2040 has USB-device hardware, but the API is still unproven: %v", diags)
	}
}

// TestRTOSFoundation: ESP32 Arduino targets already run under FreeRTOS; other catalog boards are bare-metal in the
// runtime Mujard generates today. The requirement stays language-level until an RTOS scheduling backend exists.
func TestRTOSFoundation(t *testing.T) {
	for name, want := range map[string]string{
		"uno": "none", "esp32": "freertos", "esp8266": "none", "pico": "none",
		"esp32s3": "freertos", "esp32s2": "freertos", "esp32c3": "freertos", "esp32c6": "freertos",
		"unor4": "none", "nano33ble": "none", "zero": "none", "bluepill": "none",
	} {
		if got := boards[name].RTOS(); got != want {
			t.Errorf("%s RTOS: got %q, want %q", name, got, want)
		}
	}
	for name, want := range map[string]string{"uno": LevelNone, "esp32": LevelLanguage, "pico": LevelNone} {
		if got := CapabilityLevel(boards[name], "rtos"); got != want {
			t.Errorf("%s rtos: got %q, want %q", name, got, want)
		}
	}
	sk, diags := Compile("board uno\nat start:\n    print board.rtos\n", "p.muj")
	if e := FirstError(diags); e != nil || sk == nil {
		t.Fatalf("board.rtos should compile: %v", diags)
	}
	if !strings.Contains(sk.Files["mujard_program.cpp"], `mujard::Text("none")`) {
		t.Errorf("the Uno runtime fact should be compile-time 'none':\n%s", sk.Files["mujard_program.cpp"])
	}
	if _, diags := CheckSource("board uno\nneeds rtos\n"); FirstError(diags) == nil || FirstError(diags).Code != "B415" {
		t.Errorf("the generated Uno runtime is bare-metal: %v", diags)
	}
	if _, diags := CheckSource("board esp32\nneeds rtos\n"); FirstError(diags) != nil || !hasDiag(diags, "W618") {
		t.Errorf("ESP32 runs FreeRTOS, but Mujard has no RTOS backend proof yet: %v", diags)
	}
}

// TestEthernetFoundation: only the classic ESP32 in the catalog has an Ethernet MAC. mujard.ethernet gives it a
// backend (unproven until a build matrix run); the MAC alone still doesn't mean the board has PHY or a jack.
func TestEthernetFoundation(t *testing.T) {
	for name, want := range map[string]bool{
		"esp32": true, "esp32s2": false, "esp32s3": false, "esp32c3": false, "esp32c6": false,
		"uno": false, "pico": false, "unor4": false, "nano33ble": false, "zero": false, "bluepill": false,
	} {
		if got := boards[name].HasEthernet(); got != want {
			t.Errorf("%s Ethernet MAC: got %v, want %v", name, got, want)
		}
	}
	for name, want := range map[string]string{"esp32": LevelBackend, "esp32s3": LevelNone, "pico": LevelNone} {
		if got := CapabilityLevel(boards[name], "ethernet"); got != want {
			t.Errorf("%s ethernet: got %q, want %q", name, got, want)
		}
	}
	sk, diags := Compile("board esp32\nat start:\n    print board.ethernet\n", "p.muj")
	if e := FirstError(diags); e != nil || sk == nil {
		t.Fatalf("board.ethernet should compile: %v", diags)
	}
	if !strings.Contains(sk.Files["mujard_program.cpp"], ".add((bool)(true))") {
		t.Errorf("the ESP32 MAC fact should be a compile-time true:\n%s", sk.Files["mujard_program.cpp"])
	}
	if _, diags := CheckSource("board pico\nneeds ethernet\n"); FirstError(diags) == nil || FirstError(diags).Code != "B415" {
		t.Errorf("the RP2040 has no Ethernet MAC: %v", diags)
	}
	if _, diags := CheckSource("board esp32\nneeds ethernet\n"); FirstError(diags) != nil || !hasDiag(diags, "W618") {
		t.Errorf("ESP32 has the MAC, but no Ethernet backend proof yet: %v", diags)
	}
}
