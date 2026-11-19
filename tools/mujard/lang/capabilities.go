package lang

import (
	_ "embed"
	"encoding/json"
	"sort"
	"strings"
)

// The capability model (plan Phase 16): what a program asks of the hardware, and how far each of those things is
// proven on each board. The four maturity levels are the plan's honest-claims rule - a capability is never called
// working because its code exists:
//
//	language  - Mujard has the syntax and meaning
//	backend   - the generator and runtime produce it for that board's family
//	build     - a real arduino-cli build of a program using it succeeded for that board (the build matrix wrote it)
//	hardware  - it ran on a real board and did what it should (written by hand, with who and when)
//
// "none" means the board doesn't have the hardware (Wi-Fi on an Uno); the checker refuses that (B409).

// Capability is one thing a program can ask of a board.
type Capability struct {
	Name string // "pwm", "wifi", ...
	What string // plain words for docs and 'mujard boards caps'
	Need func(b *Board) bool
}

// Capabilities in the order docs show them. Need is nil when every board has it.
var Capabilities = []Capability{
	{"output", "digital outputs (pin N as output)", nil},
	{"input", "digital inputs and their events (pressed, changed...)", nil},
	{"pwm", "PWM outputs (pin N as pwm)", func(b *Board) bool { return len(b.PWM) > 0 }},
	{"analog", "analog inputs (analog A0)", func(b *Board) bool { return len(b.AnalogPins()) > 0 }},
	{"fast", "fast pin events (interrupts)", nil},
	{"tone", "tone and melody", nil},
	{"servo", "servo motors (mujard.servo)", nil},
	{"ultrasonic", "ultrasonic distance sensors (mujard.ultrasonic)", nil},
	{"uart", "extra serial ports (mujard.uart)", nil},
	{"i2c", "I2C (mujard.i2c)", nil},
	{"spi", "SPI (mujard.spi)", nil},
	{"shift", "shift registers (mujard.shift)", nil},
	{"store", "saved values and store.save/load (EEPROM or flash)", nil},
	{"rtc", "a real-time clock chip (mujard.rtc)", nil},
	{"regs", "register access (mujard.regs)", func(b *Board) bool { _, ok := regsOf(b); return ok }},
	{"wifi", "Wi-Fi", func(b *Board) bool { return b.Wifi }},
	{"clock", "the clock, at HH:MM", nil},
	{"http", "HTTP and HTTPS requests", func(b *Board) bool { return b.Wifi }},
	{"mqtt", "MQTT (mujard.mqtt)", func(b *Board) bool { return b.Wifi }},
	{"multicore", "more than one CPU core (blocks can run `on core N`)", func(b *Board) bool { return b.Cores() > 1 }},
	{"dma", "a general-purpose DMA engine (the DMA API isn't built)", func(b *Board) bool { return b.DMA }},
	{"timers", "hardware timer units (Timer objects)", func(b *Board) bool { return b.Timers > 0 }},
	{"can", "a CAN/TWAI controller (mujard.can)", func(b *Board) bool { return b.HasCAN() }},
	{"bluetooth", "a Bluetooth/BLE radio (mujard.ble)", func(b *Board) bool { return b.HasBluetooth() }},
	{"usb_device", "a native USB-device peripheral (mujard.usb)", func(b *Board) bool { return b.HasUSBDevice() }},
	{"rtos", "an RTOS runtime (ESP32 targets run FreeRTOS; Mujard doesn't schedule through it)", func(b *Board) bool { return b.RTOS() != "none" }},
	{"ethernet", "an on-chip Ethernet MAC (mujard.ethernet)", func(b *Board) bool { return b.HasEthernet() }},
}

// languageOnly has a language-level capability name but no generated backend yet, so having an FQBN must not count
// as backend proof. Remove an entry once code generation and the runtime actually provide it.
var languageOnly = map[string]bool{
	"dma": true, "rtos": true,
}

// NotInMujardYet are the plan's hardware areas the language doesn't reach yet; docs list them so the table isn't read
// as covering them.
var NotInMujardYet = []string{"DMA", "displays as built-ins"}

// Maturity levels, lowest first.
const (
	LevelNone     = "none"
	LevelLanguage = "language"
	LevelBackend  = "backend"
	LevelBuild    = "build"
	LevelHardware = "hardware"
)

//go:embed capabilities.json
var evidenceJSON []byte

// Evidence is capabilities.json: per board, per capability, when it was proven and how.
type Evidence struct {
	Build    map[string]map[string]string `json:"build"`    // board -> capability -> date of the build matrix run
	Hardware map[string]map[string]string `json:"hardware"` // board -> capability -> "date: what ran, who saw it"
}

var evidence = loadEvidence()

func loadEvidence() *Evidence {
	var e Evidence
	if err := json.Unmarshal(evidenceJSON, &e); err != nil {
		panic("lang: capabilities.json: " + err.Error())
	}
	return &e
}

// CapabilityEvidence is the evidence the compiler was built with.
func CapabilityEvidence() *Evidence { return evidence }

// CapabilityLevel is how far a capability is proven on a board.
func CapabilityLevel(b *Board, name string) string {
	var cp *Capability
	for i := range Capabilities {
		if Capabilities[i].Name == name {
			cp = &Capabilities[i]
		}
	}
	switch {
	case cp == nil:
		return ""
	case cp.Need != nil && !cp.Need(b):
		return LevelNone
	case languageOnly[name]:
		return LevelLanguage
	case evidence.Hardware[b.Name][name] != "":
		return LevelHardware
	case evidence.Build[b.Name][name] != "":
		return LevelBuild
	case fqbns[b.Name] != "":
		return LevelBackend
	}
	return LevelLanguage
}

// programCaps is what a checked program asks of its board, sorted (modules' devices and libraries included).
func (c *checker) programCaps() []string {
	set := map[string]bool{}
	var walk func(f *File)
	walk = func(f *File) {
		for _, it := range f.Items {
			switch d := it.(type) {
			case *DeviceDecl:
				switch d.Kind {
				case "analog":
					set["analog"] = true
				case "pin", "pins", "pinlist":
					set[d.Role] = true
				}
			case *EventBlock:
				if d.Fast {
					set["fast"] = true
				}
				if d.Core > 0 {
					set["multicore"] = true
				}
			case *TaskDecl:
				if d.Core > 0 {
					set["multicore"] = true
				}
			case *NeedsDecl:
				for _, n := range d.Names {
					set[n] = true
				}
			}
		}
	}
	walk(c.file)
	for _, mc := range c.mods {
		if mc != nil && mc.file != nil {
			walk(mc.file)
		}
	}
	if c.usesTone {
		set["tone"] = true
	}
	for lib := range c.libs {
		set[lib] = true
	}
	var out []string
	for _, cp := range Capabilities {
		if set[cp.Name] {
			out = append(out, cp.Name)
		}
	}
	return out
}

// needsDecl checks 'needs i2c, pwm' (spec 10.5): every name is a capability (C330) the board has (B415). A driver
// module says this once, so a program importing it onto the wrong board hears it in plain words, from the module.
func (c *checker) needsDecl(d *NeedsDecl) {
	if !c.inModule {
		for i := range d.PinNames {
			c.errorf("C333", d.PinPos[i], "pin", d.PinNames[i])
		}
	}
	for i, n := range d.Names {
		if !isCapability(n) {
			var names []string
			for _, cp := range Capabilities {
				names = append(names, cp.Name)
			}
			c.errorf("C330", d.Pos[i], "name", n, "list", strings.Join(names, ", "))
			continue
		}
		if c.board != nil && CapabilityLevel(c.board, n) == LevelNone {
			who := "This program"
			if c.inModule {
				who = "The module '" + c.modName + "'"
			}
			c.errorf("B415", d.Pos[i], "who", who, "cap", n, "board", c.board.Display)
		}
	}
}

// bindModulePins gives a driver module the pins its importer wired it to (spec 10.5): the module says
// 'needs pin trig', the program 'import sonar(trig: pin 7)', and every 'pin trig' in the module becomes pin 7 before
// the module is checked - so the board's pin rules and pin ownership (B407) apply to the real pin, and the
// generated C++ uses it. Missing pins are C331, pins the module doesn't ask for C332. It returns the pins left
// unbound, which the module's check then skips (one C331, not a B402 at every use as well).
func (c *checker) bindModulePins(d *ImportDecl, mf *File) map[string]bool {
	params := map[string]bool{}
	var order []string
	for _, it := range mf.Items {
		if nd, ok := it.(*NeedsDecl); ok {
			for _, n := range nd.PinNames {
				if !params[n] {
					params[n] = true
					order = append(order, n)
				}
			}
		}
	}
	given := map[string]PinRef{}
	for _, ip := range d.Pins {
		if !params[ip.Name] {
			list := "none"
			if len(order) > 0 {
				list = strings.Join(order, ", ")
			}
			c.errorf("C332", ip.At, "mod", d.Name, "pin", ip.Name, "list", list)
			continue
		}
		given[ip.Name] = ip.Ref
	}
	unbound := map[string]bool{}
	for _, n := range order {
		if _, ok := given[n]; !ok {
			c.errorf("C331", d.At, "mod", d.Name, "pin", n)
			unbound[n] = true
		}
	}
	bind := func(r *PinRef) {
		if v, ok := given[r.Text]; ok && params[r.Text] {
			r.Text = v.Text // the position stays in the module, where the device is
		}
	}
	for _, it := range mf.Items {
		if dd, ok := it.(*DeviceDecl); ok {
			for i := range dd.Pins {
				bind(&dd.Pins[i])
			}
		}
	}
	walkTree(mf, func(n Node) {
		if pl, ok := n.(*PinLit); ok {
			bind(&pl.Ref)
		}
	})
	return unbound
}

func isCapability(name string) bool {
	for _, cp := range Capabilities {
		if cp.Name == name {
			return true
		}
	}
	return false
}

// capabilityCheck warns (W618) when the program needs something its board hasn't been shown to build: the build
// may still work, but nobody has seen it do so. One warning at the board line, naming everything unproven.
func (c *checker) capabilityCheck() {
	if c.board == nil || c.inModule {
		return
	}
	var unproven []string
	for _, cp := range c.programCaps() {
		switch CapabilityLevel(c.board, cp) {
		case LevelLanguage, LevelBackend:
			unproven = append(unproven, cp)
		}
	}
	if len(unproven) == 0 {
		return
	}
	sort.Strings(unproven)
	c.errorf("W618", c.file.BoardPos, "board", c.board.Display, "list", strings.Join(unproven, ", "))
}
