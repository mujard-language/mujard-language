package lang

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// boards.json is generated from the platform's shared pin tables (the ones the website and the phone app use);
// a drift test on the platform side fails when the two fall out of step.
//
//go:embed boards.json
var boardsJSON []byte

// Board is what the checker knows about one board.
type Board struct {
	Name    string // Mujard name: uno, esp32, ...
	Display string // "Arduino Uno"
	Family  string // "ATmega328P", "ESP32", ...
	Arch    string // the core's architecture: "avr", "esp32", "esp8266", "rp2040" (library.properties 'architectures=')

	ValidPins        []int          `json:"validPins"`
	PWM              []int          `json:"pwm"`
	InputOnly        []int          `json:"inputOnly"`
	NoPullup         []int          `json:"noPullup"`
	FlashPins        []int          `json:"flashPins"`
	USBSerialPins    []int          `json:"usbSerialPins"`
	AnalogOnly       []int          `json:"analogOnly"`
	ADC              []int          `json:"adc"`
	ADCBlockedByWifi []int          `json:"adcBlockedByWifi"`
	ADCBits          int            `json:"adcBits"`
	DMA              bool           `json:"dma"`
	Timers           int            `json:"timers"`
	Names            map[string]int `json:"names"`
	// Claims are the pins a bus or library takes over: "i2c" [SDA, SCL], "spi" [MOSI, MISO, SCK], "servo" the PWM
	// pins whose timer the Servo library takes.
	Claims map[string][]int `json:"claims"`
	// Memory is what a program can use, in bytes: flash for the program, RAM, and EEPROM (or the flash Mujard uses as one).
	Memory struct {
		Flash  int `json:"flash"`
		RAM    int `json:"ram"`
		EEPROM int `json:"eeprom"`
	} `json:"memory"`

	// Mujard's own facts (not in hardware.ts)
	Strapping  []int // boot (strapping) pins
	Interrupts []int // pins that can use 'fast' events; nil = every valid pin
	Wifi       bool
}

// boardExtras are the facts hardware.ts doesn't carry. Strapping pins: the Espressif datasheets; interrupt pins: the
// Arduino attachInterrupt() reference.
var boardExtras = map[string]struct {
	display    string
	strapping  []int
	interrupts []int
	wifi       bool
}{
	"uno":        {"Arduino Uno", nil, []int{2, 3}, false},
	"nano":       {"Arduino Nano", nil, []int{2, 3}, false},
	"leonardo":   {"Arduino Leonardo", nil, []int{0, 1, 2, 3, 7}, false},
	"micro":      {"Arduino Micro", nil, []int{0, 1, 2, 3, 7}, false},
	"mega":       {"Arduino Mega 2560", nil, []int{2, 3, 18, 19, 20, 21}, false},
	"esp32":      {"ESP32", []int{0, 2, 5, 12, 15}, nil, true},
	"esp8266":    {"ESP8266", []int{0, 2, 15}, nil, true},
	"pico":       {"Raspberry Pi Pico", nil, nil, false},
	"picow":      {"Raspberry Pi Pico W", nil, nil, true},
	"rp2350usba": {"Waveshare RP2350-USB-A", nil, nil, false},
	"esp32s3":    {"ESP32-S3", []int{0, 3, 45, 46}, nil, true},
	"esp32s2":    {"ESP32-S2", []int{0, 45, 46}, nil, true},
	"esp32c3":    {"ESP32-C3", []int{2, 8, 9}, nil, true},
	"esp32c6":    {"ESP32-C6", []int{8, 9, 15}, nil, true},
	"unor4":      {"Arduino Uno R4 Minima", nil, nil, false},
	"nano33ble":  {"Arduino Nano 33 BLE", nil, nil, false},
	"zero":       {"Arduino Zero", nil, nil, false},
	"bluepill":   {"STM32 BluePill", nil, nil, false},
}

// boardTimers names the hardware timer whose interrupt a claim's code defines on a board: the AVR core's tone() and
// NewPing (its timer functions are compiled in whenever the library is). Two claims on one timer each define that
// interrupt, so the program can't be linked (B414). Only what the real toolchain showed is listed (2026-10-30,
// arduino-cli with the installed cores and libraries); IRremote and Servo build next to tone, so they are absent.
var boardTimers = map[string]map[string]string{
	"uno":      {"tone": "Timer 2", "newping": "Timer 2", "timer": "Timer 1", "servo": "Timer 1"},
	"nano":     {"tone": "Timer 2", "newping": "Timer 2", "timer": "Timer 1", "servo": "Timer 1"},
	"mega":     {"tone": "Timer 2", "newping": "Timer 2", "timer": "Timer 1", "servo": "Timer 1"}, // the Leonardo's tone uses Timer 3: no conflict there
	"leonardo": {"timer": "Timer 1", "servo": "Timer 1"},
	"micro":    {"timer": "Timer 1", "servo": "Timer 1"},
	// ESP32-family: Timer() takes a hw_timer channel; Servo uses ledc PWM - different names, no conflict.
	"esp32":   {"timer": "hw_timer_1"},
	"esp32s2": {"timer": "hw_timer_1"},
	"esp32s3": {"timer": "hw_timer_1"},
	"esp32c3": {"timer": "hw_timer_1"},
	"esp32c6": {"timer": "hw_timer_1"},
	"esp32h2": {"timer": "hw_timer_1"},
	// RP2040/RP2350: Timer() takes the Pico SDK repeating_timer; Servo uses PWM - no shared interrupt.
	"pico":       {"timer": "repeating_timer"},
	"picow":      {"timer": "repeating_timer"},
	"rp2350usba": {"timer": "repeating_timer"},
	"rp2350usb":  {"timer": "repeating_timer"},
}

// TimerOf is the hardware timer a claim's code takes on this board, or "".
func (b *Board) TimerOf(claim string) string { return boardTimers[b.Name][claim] }

// boardPwmTimers names the hardware timer each AVR pwm pin is attached to (the Arduino core's analogWrite wiring).
// It exists only for the AVR boards where 'Timer()' claims Timer 1 outright: a pwm pin on the same timer is a real
// conflict (Timer 1 on Uno pins 9/10), a pin on a different timer is not. ESP32 ledc channels and RP2040 PWM slices
// are allocated per attach at runtime, so there is no static name to claim for them here.
var boardPwmTimers = map[string]map[int]string{
	// ATmega328P: Timer 0 = 5,6; Timer 1 = 9,10; Timer 2 = 3,11.
	"uno":  {3: "Timer 2", 5: "Timer 0", 6: "Timer 0", 9: "Timer 1", 10: "Timer 1", 11: "Timer 2"},
	"nano": {3: "Timer 2", 5: "Timer 0", 6: "Timer 0", 9: "Timer 1", 10: "Timer 1", 11: "Timer 2"},
	// ATmega2560: Timer 0 = 4,13; Timer 1 = 11,12; Timer 2 = 9,10; Timer 3 = 2,3,5; Timer 4 = 6,7,8; Timer 5 = 44,45,46.
	"mega": {4: "Timer 0", 13: "Timer 0", 11: "Timer 1", 12: "Timer 1", 9: "Timer 2", 10: "Timer 2",
		2: "Timer 3", 3: "Timer 3", 5: "Timer 3", 6: "Timer 4", 7: "Timer 4", 8: "Timer 4",
		44: "Timer 5", 45: "Timer 5", 46: "Timer 5"},
	// ATmega32U4 (Leonardo/Micro): Timer 0 = 3,11; Timer 1 = 9,10; Timer 3 = 5; Timer 4 = 6,13.
	"leonardo": {3: "Timer 0", 5: "Timer 3", 6: "Timer 4", 9: "Timer 1", 10: "Timer 1", 11: "Timer 0", 13: "Timer 4"},
	"micro":    {3: "Timer 0", 5: "Timer 3", 6: "Timer 4", 9: "Timer 1", 10: "Timer 1", 11: "Timer 0", 13: "Timer 4"},
}

// PwmTimerOf names the hardware timer a pwm pin is attached to on this board ("" when it isn't an AVR table or the
// pin isn't pwm): B414 conflicts against 'timer'/'servo'/'tone' claims in timerOwners.
func (b *Board) PwmTimerOf(pin int) string { return boardPwmTimers[b.Name][pin] }

var boards = loadBoards()

func loadBoards() map[string]*Board {
	raw := map[string]*Board{}
	if err := json.Unmarshal(boardsJSON, &raw); err != nil {
		panic("lang: boards.json: " + err.Error())
	}
	for name, b := range raw {
		b.Name = name
		x, ok := boardExtras[name]
		if !ok {
			panic("lang: no extras for board " + name)
		}
		b.Display, b.Strapping, b.Interrupts, b.Wifi = x.display, x.strapping, x.interrupts, x.wifi
		if fq := fqbns[name]; fq != "" {
			b.Arch = strings.Split(fq, ":")[1] // "arduino:avr:uno" -> "avr"
		}
	}
	return raw
}

// Upload-profile names that share a compiler board - the same chip and pins, only the
// bootloader's upload speed differs ("nano_old" is a Nano at 57600 instead of 115200).
var boardAliases = map[string]string{"nano_old": "nano"}

// LookupBoard finds a board by its Mujard name (case-insensitive); upload variants resolve to their board.
func LookupBoard(name string) *Board {
	n := strings.ToLower(name)
	if a, ok := boardAliases[n]; ok {
		n = a
	}
	return boards[n]
}

// BoardNames lists the known boards, sorted.
func BoardNames() []string {
	var n []string
	for k := range boards {
		n = append(n, k)
	}
	sort.Strings(n)
	return n
}

func has(list []int, v int) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}

// IsESP reports whether the board is one of the Espressif chips (GPIO numbering).
func (b *Board) IsESP() bool { return strings.HasPrefix(b.Family, "ESP") }

// IsAVR reports whether the board is an Arduino AVR board.
func (b *Board) IsAVR() bool { return strings.HasPrefix(b.Family, "ATmega") }

// IsRP reports whether the board is a Raspberry Pi chip (RP2040, RP2350: GP numbering).
func (b *Board) IsRP() bool { return b.Family == "RP2040" || b.Family == "RP2350" }

// Cores is the chip's usable CPU-core count (what a program could ask for, not what Mujard schedules yet).
func (b *Board) Cores() int {
	if b.IsRP() || b.Family == "ESP32" || b.Family == "ESP32-S3" {
		return 2
	}
	return 1
}

// HasCAN is a dedicated CAN/TWAI controller, not an SPI add-on board.
func (b *Board) HasCAN() bool { return strings.HasPrefix(b.Family, "ESP32") }

// HasBluetooth is a radio on the board/chip itself (BLE on the ESP32 family except S2; CYW43 on Pico W).
func (b *Board) HasBluetooth() bool {
	return b.Name == "picow" || b.Family == "ESP32" || b.Family == "ESP32-S3" || b.Family == "ESP32-C3" || b.Family == "ESP32-C6" || b.Family == "nRF52840"
}

// HasUSBDevice is a general native USB-device peripheral, not just a UART-to-USB bridge. ESP32-C3/C6's
// USB Serial/JTAG block is deliberately excluded; it isn't a general device endpoint.
func (b *Board) HasUSBDevice() bool {
	switch b.Family {
	case "ATmega32U4", "RP2040", "RP2350", "ESP32-S2", "ESP32-S3", "Renesas RA4M1", "nRF52840", "SAMD21G18", "STM32F103":
		return true
	}
	return false
}

// HasEthernet is an on-chip Ethernet MAC. It does not promise a PHY, magnetics or a connector on the board.
func (b *Board) HasEthernet() bool { return b.Family == "ESP32" }

// CanMeasureVoltage reports whether the generated runtime can read the board's own supply rail. Pico W's GP29
// still reads VSYS/3 even though the pin is owned by the CYW43 chip; RP2350-USB-A exposes GP29 as a normal pin,
// so it doesn't get a built-in supply reading. nRF52840, SAMD21, STM32F103 and RA4M1 expose internal
// supply/reference ADC inputs.
func (b *Board) CanMeasureVoltage() bool {
	return b.IsAVR() || b.Name == "pico" || b.Name == "picow" || b.Family == "nRF52840" ||
		b.Family == "SAMD21G18" || b.Family == "STM32F103" || b.Family == "Renesas RA4M1"
}

// CanReadResetReason reports whether generated code can ask the chip why it last reset. ESP targets expose a
// reason API; the ARM chips in the catalog expose reset-cause registers. On AVR this only works on Optiboot
// boards: Optiboot passes MCUSR to the app in r2, while the Mega and 32u4 bootloaders do not use that contract.
func (b *Board) CanReadResetReason() bool {
	return b.IsESP() || b.Family == "STM32F103" || b.Family == "SAMD21G18" || b.Family == "nRF52840" ||
		b.Family == "Renesas RA4M1" || b.Name == "uno" || b.Name == "nano"
}

// CanMeasureFreeMemory reports whether the generated runtime has a real free-memory backend for this board. New
// board targets must be listed only after their heap/stack measurement has been verified; unknown future targets
// get C320 instead of a bogus -1.
func (b *Board) CanMeasureFreeMemory() bool {
	switch b.Arch {
	case "avr", "esp32", "esp8266", "rp2040", "renesas_uno", "mbed_nano", "samd", "stm32":
		return true
	}
	return false
}

// RTOS names the runtime the board target already runs on ("none" for the bare-metal cores). It is a fact a driver
// can require, not a claim that Mujard schedules through that RTOS yet.
func (b *Board) RTOS() string {
	if b.Arch == "esp32" {
		return "freertos"
	}
	return "none"
}

// ResolvePin turns a pin as written (13, A0, D5, GP15, GPIO34, LED) into the chip's pin number.
func (b *Board) ResolvePin(ref string) (int, bool) {
	if ref == "LED" {
		n, ok := b.Names["LED_BUILTIN"]
		return n, ok
	}
	if n, ok := b.Names[ref]; ok {
		return n, true
	}
	num := ref
	switch {
	case strings.HasPrefix(ref, "GPIO") && b.IsESP():
		num = ref[4:]
	case strings.HasPrefix(ref, "GP") && b.IsRP():
		num = ref[2:]
	}
	n, err := strconv.Atoi(num)
	if err != nil || n < 0 || !has(b.ValidPins, n) {
		return 0, false
	}
	return n, true
}

// PinLabel is how a pin is named in messages: "13" on an Arduino, "GPIO 34" on an ESP32, "GP15" on a Pico.
func (b *Board) PinLabel(n int) string {
	switch {
	case b.IsESP():
		return fmt.Sprintf("GPIO %d", n)
	case b.IsRP():
		return fmt.Sprintf("GP%d", n)
	}
	for name, v := range b.Names {
		if v == n && b.IsAVR() && len(name) >= 2 && name[0] == 'A' && name[1] >= '0' && name[1] <= '9' {
			return name
		}
	}
	return strconv.Itoa(n)
}

// PinList renders pins for messages: "3, 5, 6, 9, 10, 11".
func (b *Board) PinList(pins []int) string {
	var s []string
	for _, p := range pins {
		s = append(s, b.PinLabel(p))
	}
	return strings.Join(s, ", ")
}

// IsAnalogInput reports whether the pin can be read by an 'analog' device.
func (b *Board) IsAnalogInput(n int) bool {
	if b.ADC != nil {
		return has(b.ADC, n)
	}
	if has(b.AnalogOnly, n) {
		return true
	}
	for name, v := range b.Names {
		if v == n && len(name) >= 2 && name[0] == 'A' && name[1] >= '0' && name[1] <= '9' {
			return true
		}
	}
	return false
}

// AnalogPins lists the analog-capable pins for messages.
func (b *Board) AnalogPins() []int {
	var out []int
	for _, p := range b.ValidPins {
		if b.IsAnalogInput(p) {
			out = append(out, p)
		}
	}
	return out
}

// CanInterrupt reports whether a pin can run 'fast' events.
func (b *Board) CanInterrupt(n int) bool {
	if b.Interrupts == nil {
		return has(b.ValidPins, n) && !has(b.FlashPins, n)
	}
	return has(b.Interrupts, n)
}
