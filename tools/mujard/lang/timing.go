package lang

import (
	"fmt"
	"strconv"
)

// Timing analysis (plan Phase 18): the one overrun a program can be shown to have before it runs. A serial port
// through a USB bridge sends one character in 10 bit times (8 data bits, a start and a stop bit); an 'every' block
// that prints more each run than the port can send in its period fills the transmit buffer, print then waits, and
// the scheduler skips runs (spec 6.2, board.overruns). Only what is certain counts: the plain text of prints the
// block always reaches (not those under an if or in a loop), plus print's \r\n - the real output is longer. Serial.write
// waits while the core's 64-byte transmit buffer is full, which is what makes the block late.

// serialBaud is the program's serial speed: its serial(...) call, else the runtime's default (mujard_rt.cpp).
func (c *checker) serialBaud() float64 {
	baud := 115200.0
	walkTree(c.file, func(n Node) {
		call, ok := n.(*Call)
		if !ok {
			return
		}
		if id, ok := call.Fn.(*Ident); ok && id.Name == "serial" && len(call.Args) == 1 {
			if lit, ok := call.Args[0].Value.(*NumberLit); ok {
				baud = litValue(lit)
			}
		}
	})
	return baud
}

func litValue(n *NumberLit) float64 {
	if n.IsFloat {
		return n.Float
	}
	return float64(n.Int)
}

// serialBudget warns (W619) about 'every' blocks that print faster than the serial port can send.
func (c *checker) serialBudget() {
	if c.board == nil || len(c.board.USBSerialPins) == 0 { // native-USB serial (Leonardo, Pico...) has no baud limit
		return
	}
	baud := c.serialBaud()
	if baud <= 0 {
		return
	}
	charUS := 10e6 / baud
	for _, it := range c.file.Items {
		ev, ok := it.(*EventBlock)
		if !ok || ev.Kind != "every" || ev.Body == nil {
			continue
		}
		lit, ok := ev.Expr.(*NumberLit)
		if !ok {
			continue
		}
		var periodUS float64
		switch Units[lit.Unit] {
		case "time":
			periodUS = litValue(lit) * unitFactor(lit.Unit)
		case "frequency": // every 2 Hz: the runtime's frequencies are in millihertz
			if hz := litValue(lit) * unitFactor(lit.Unit) / 1000; hz > 0 {
				periodUS = 1e6 / hz
			}
		}
		if periodUS <= 0 {
			continue
		}
		chars := 0
		for _, s := range ev.Body.Stmts { // only the statements every run reaches
			p, ok := s.(*SimpleStmt)
			if !ok || p.Kind != "print" {
				continue
			}
			chars += 2 // print ends the line with \r\n (printClean in mujard_rt.cpp)
			if sl, ok := p.Value.(*StringLit); ok {
				for _, part := range sl.Parts {
					if part.Expr == nil {
						chars += len(part.Text)
					}
				}
			}
		}
		if needUS := float64(chars) * charUS; chars > 0 && needUS > periodUS {
			c.errorf("W619", ev.At, "n", strconv.Itoa(chars), "period", lit.Text+" "+lit.Unit,
				"baud", strconv.FormatFloat(baud, 'f', -1, 64), "need", fmtMicros(needUS))
		}
	}
}

// fmtMicros writes a duration for people: "1.7 ms", "430 us".
func fmtMicros(us float64) string {
	if us >= 1000 {
		return fmt.Sprintf("%.1f ms", us/1000)
	}
	return fmt.Sprintf("%.0f us", us)
}
