package lang

import "strconv"

// An estimate of the RAM a program keeps for its whole life (W614), so the editor can warn before anything is built:
// a list of 30 texts is about 2 KB, the whole RAM of an Uno. It counts what the program itself declares - global
// variables, devices, the state of tasks and timed blocks, and the per-call-site memory of smooth, peak, ... (spec 6.9) -
// in the sizes the generated C++ has on an 8-bit AVR. The runtime, the serial port's buffers and C++ libraries come on
// top, so the estimate is low rather than high; the real figure comes from the linker when the program is built
// (mujard build prints it).

// ramWarnPercent: W614 when the estimate passes this share of the board's RAM.
const ramWarnPercent = 60

// siteMemory is the RAM each call site of a built-in with memory keeps (the runtime's Smooth, Peak, Changed,
// Hysteresis and Edge objects).
var siteMemory = map[string]int{"smooth": 130, "peak": 13, "changed": 5, "hysteresis": 1, "rising": 1, "falling": 1}

// deviceRAM is the size of the runtime's device objects on AVR (vtable and list pointers included; outputs and inputs
// also keep their port register and bit for direct pin access). TestRAMEstimateMatchesTheBuild checks these against a
// real build, so they follow the runtime.
var deviceRAM = map[string]int{"output": 72, "input": 70, "pwm": 43, "analog": 9}

// ramOf is the size of a value of type t; declCap is a list's declared capacity (0: its type's or 8).
func (c *checker) ramOf(t *Type, declCap int) int {
	if t == nil {
		return 0
	}
	n := 0
	switch t.K {
	case TInt, TFloat:
		n = 4
	case TBool, TEnum:
		n = 1
	case TText:
		capacity := declCap
		if capacity <= 0 {
			capacity = t.Cap
		}
		if capacity <= 0 {
			capacity = textMax
		}
		n = capacity + 2
	case TQty:
		n = 4
		if t.Dim == "time" || t.Dim == "frequency" {
			n = 8
		}
	case TList:
		capacity := declCap
		if capacity <= 0 {
			capacity = t.Cap
		}
		if capacity <= 0 {
			capacity = 8
		}
		n = capacity*c.ramOf(t.Elem, 0) + 4
	case TStruct:
		if sym := c.typeSym(t.Name); sym != nil && sym.strukt != nil {
			for _, f := range sym.strukt.Fields {
				n += c.ramOf(c.resolveType(f.Type), declaredCap(f.Type))
			}
		}
	case TDevice:
		n = deviceRAM[t.Name]
	case TObject:
		if t.Name == "stopwatch" {
			n = 17
		}
		if t.Name == "queue" && t.Elem != nil {
			n = t.Cap*c.ramOf(t.Elem, 0) + 16 // the ring buffer plus its head/count (and mutex on RP-family)
		}
	}
	if t.Opt {
		n++
	}
	return n
}

// declaredCap is the capacity written in a fixed-size type (text[8] or list[int, 10]), else 0.
func declaredCap(te *TypeExpr) int {
	if te == nil || len(te.Args) == 0 {
		return 0
	}
	arg := 0
	if te.Name == "list" {
		if len(te.Args) < 2 {
			return 0
		}
		arg = 1
	} else if te.Name != "text" {
		return 0
	}
	n, _ := strconv.Atoi(te.Args[arg].Name)
	return n
}

// ramEstimate adds up what the program keeps in RAM and warns (W614) when it is most of the board's.
func (c *checker) ramEstimate() {
	if c.board == nil || c.board.Memory.RAM <= 0 || c.inModule {
		return
	}
	total := c.siteBytes
	countLocals := func(b *Block) {
		walkStmts(b, func(n Node) {
			if v, ok := n.(*VarDecl); ok && v.Kind != "const" {
				total += c.ramOf(c.types[v.Value], declaredCap(v.Type)) // task locals live as long as the task
			}
		})
	}
	for _, it := range c.file.Items {
		switch d := it.(type) {
		case *VarDecl:
			if d.Kind == "const" {
				continue
			}
			if sym := c.global.syms[d.Name]; sym != nil {
				total += c.ramOf(sym.t, declaredCap(d.Type))
			}
		case *DeviceDecl:
			sym := c.global.syms[d.Name]
			if sym == nil {
				continue
			}
			if sym.t.K == TDevice && sym.t.Name == "group" {
				total += len(c.declPins(d)) * deviceRAM[d.Role]
			} else {
				total += c.ramOf(sym.t, 0)
			}
		case *TaskDecl:
			total += 11 // the runtime's Task: where it waits and until when
			if c.coreStack(d.Core) > 0 {
				total += c.coreStack(d.Core)
			}
			countLocals(d.Body)
		case *EventBlock:
			total += map[string]int{"every": 17, "after": 9, "when": 2}[d.Kind]
			if c.coreStack(d.Core) > 0 {
				total += c.coreStack(d.Core)
			}
			walkStmts(d.Body, func(n Node) {
				if e, ok := n.(*EventBlock); ok {
					total += map[string]int{"after": 9, "when": 2}[e.Kind]
				}
			})
		}
	}
	ram := c.board.Memory.RAM
	if total*100 > ram*ramWarnPercent {
		c.errorf("W614", c.file.BoardPos, "n", strconv.Itoa(total), "board", c.board.Display, "ram", strconv.Itoa(ram))
	}
}

// coreStack is the RAM an 'on core N' block costs beyond its gate object: on ESP32-family the codegen emits
// xTaskCreatePinnedToCore with an 8192-byte stack per core block (coreLaunch). RP2040/RP2350's multicore_launch_core1
// runs on the second core's fixed SDK stack - already part of the linked image, not heap this estimate can see.
func (c *checker) coreStack(core int) int {
	if core <= 0 || c.board == nil {
		return 0
	}
	if c.board.Family == "RP2040" || c.board.Family == "RP2350" {
		return 0
	}
	return 8192
}
