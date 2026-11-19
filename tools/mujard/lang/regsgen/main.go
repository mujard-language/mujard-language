// regsgen builds lang/regs.json from the AVR GCC io headers - the register and bit names the maker publishes
// (the same data ATDF files carry). Run from the lang directory:
//
//	go run ./regsgen -toolchain "%LOCALAPPDATA%\mujard\toolchains\avr-gcc-7.3.0-atmel3.6.1-arduino7"
//
// Espressif and other non-AVR chips come from their SVD files (-svd-dir). A board whose chip has no table yet
// gets B413 at compile time, never a crash.
package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// chip -> its io header. Families live in lang/regs.go's chipRegsByFamily.
var chips = map[string]string{
	"atmega328p": "iom328p.h",
	"atmega32u4": "iom32u4.h",
	"atmega2560": "iom2560.h",
}

// chip -> its SVD file name in the esp32-core tools/ide-debug/svd directory.
var svdChips = map[string]string{
	"esp32":   "esp32.svd",
	"esp32s2": "esp32s2.svd",
	"esp32s3": "esp32s3.svd",
	"esp32c3": "esp32c3.svd",
	"esp32c6": "esp32c6.svd",
	"esp32c2": "esp32c2.svd",
	"esp32h2": "esp32h2.svd",
	"esp8266": "esp8266.svd",
}

// chip -> its vendor SVD file in regsgen/svd (downloaded from the maker's SDK or the cmsis-svd-data aggregate).
var vendorChips = map[string]string{
	"rp2040":    "rp2040.svd",    // pico-sdk: src/rp2040/hardware_regs/rp2040.svd
	"nrf52840":  "nrf52840.svd",  // Nordic MDK (via AdaCore/svd2ada CMSIS-SVD)
	"samd21":    "samd21.svd",    // Microchip ATSAMD21G18A (cmsis-svd-data)
	"ra4m1":     "ra4m1.svd",     // Renesas R7FA4M1AB (cmsis-svd-data)
	"stm32f103": "stm32f103.svd", // ST STM32F103xx (cmsis-svd-data)
}

// gpioCount: how many GPIO pads each Espressif chip has (datasheet). The high bank (OUT1/ENABLE1) covers 32 up.
var gpioCount = map[string]int{
	"esp32": 40, "esp32s2": 47, "esp32s3": 49, "esp32c2": 21, "esp32c3": 22, "esp32c6": 31, "esp32h2": 28, "esp8266": 17,
}

// portPins: which board pins each GPIO port register drives (Arduino numbering, per chip's common boards).
var portPins = map[string]map[string][]int{
	"atmega328p": { // Uno/Nano
		"PINB": {8, 9, 10, 11, 12, 13}, "DDRB": {8, 9, 10, 11, 12, 13}, "PORTB": {8, 9, 10, 11, 12, 13},
		"PINC": {14, 15, 16, 17, 18, 19}, "DDRC": {14, 15, 16, 17, 18, 19}, "PORTC": {14, 15, 16, 17, 18, 19},
		"PIND": {0, 1, 2, 3, 4, 5, 6, 7}, "DDRD": {0, 1, 2, 3, 4, 5, 6, 7}, "PORTD": {0, 1, 2, 3, 4, 5, 6, 7},
	},
	"atmega32u4": { // Leonardo/Micro
		"PINB": {17, 15, 16, 14, 8, 9, 10, 11}, "DDRB": {17, 15, 16, 14, 8, 9, 10, 11}, "PORTB": {17, 15, 16, 14, 8, 9, 10, 11},
		"PINC": {5, 13}, "DDRC": {5, 13}, "PORTC": {5, 13},
		"PIND": {3, 2, 0, 1, 4, 12}, "DDRD": {3, 2, 0, 1, 4, 12}, "PORTD": {3, 2, 0, 1, 4, 12},
		"PINE": {7}, "DDRE": {7}, "PORTE": {7},
		"PINF": {23, 22, 21, 20, 19, 18}, "DDRF": {23, 22, 21, 20, 19, 18}, "PORTF": {23, 22, 21, 20, 19, 18},
	},
	"atmega2560": { // Mega 2560
		"PINA": {22, 23, 24, 25, 26, 27, 28, 29}, "DDRA": {22, 23, 24, 25, 26, 27, 28, 29}, "PORTA": {22, 23, 24, 25, 26, 27, 28, 29},
		"PINB": {53, 52, 51, 50, 10, 11, 12, 13}, "DDRB": {53, 52, 51, 50, 10, 11, 12, 13}, "PORTB": {53, 52, 51, 50, 10, 11, 12, 13},
		"PINC": {37, 36, 35, 34, 33, 32, 31, 30}, "DDRC": {37, 36, 35, 34, 33, 32, 31, 30}, "PORTC": {37, 36, 35, 34, 33, 32, 31, 30},
		"PIND": {21, 20, 18, 19}, "DDRD": {21, 20, 18, 19}, "PORTD": {21, 20, 18, 19},
		"PINE": {0, 1, 5, 2, 3}, "DDRE": {0, 1, 5, 2, 3}, "PORTE": {0, 1, 5, 2, 3},
		"PINF": {54, 55, 56, 57, 58, 59, 60, 61}, "DDRF": {54, 55, 56, 57, 58, 59, 60, 61}, "PORTF": {54, 55, 56, 57, 58, 59, 60, 61},
		"PING": {41, 40, 39}, "DDRG": {41, 40, 39}, "PORTG": {41, 40, 39},
		"PINH": {17, 16, 6, 7, 8, 9}, "DDRH": {17, 16, 6, 7, 8, 9}, "PORTH": {17, 16, 6, 7, 8, 9},
		"PINJ": {15, 14}, "DDRJ": {15, 14}, "PORTJ": {15, 14},
		"PINK": {62, 63, 64, 65, 66, 67, 68, 69}, "DDRK": {62, 63, 64, 65, 66, 67, 68, 69}, "PORTK": {62, 63, 64, 65, 66, 67, 68, 69},
		"PINL": {42, 43, 44, 45, 46, 47, 48, 49}, "DDRL": {42, 43, 44, 45, 46, 47, 48, 49}, "PORTL": {42, 43, 44, 45, 46, 47, 48, 49},
	},
}

var (
	reReg   = regexp.MustCompile(`^#define\s+(\w+)\s+_SFR_(IO|MEM)(8|16|32)\((0x[0-9A-Fa-f]+|\d+)\)`)
	reField = regexp.MustCompile(`^#define\s+(\w+)\s+(-?\d+)\s*$`)
	reBlank = regexp.MustCompile(`^\s*$`)
)

type regDef struct {
	Addr   int              `json:"addr"`
	Bits   int              `json:"bits"`
	Fields map[string]int64 `json:"fields"` // bit-field name -> its mask (single bit: 1<<pos)
	Port   []int            `json:"port,omitempty"`
}

var reInc = regexp.MustCompile(`^#include\s+<avr/(\w+\.h)>`)

func parse(header string) (map[string]regDef, error) {
	regs := map[string]regDef{}
	if err := parseInto(header, regs); err != nil {
		return nil, err
	}
	return regs, nil
}

func parseInto(header string, regs map[string]regDef) error {
	data, err := os.ReadFile(header)
	if err != nil {
		return err
	}
	lastReg := "" // scalar #defines after a register line are its bit names
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if m := reReg.FindStringSubmatch(line); m != nil {
			addr, _ := strconv.ParseUint(m[4], 0, 32)
			if m[2] == "IO" {
				addr += 0x20 // IO space sits at offset 0x20 in the data map
			}
			bits, _ := strconv.Atoi(m[3])
			r := regs[m[1]]
			r.Addr = int(addr)
			r.Bits = bits
			if r.Fields == nil {
				r.Fields = map[string]int64{}
			}
			regs[m[1]] = r
			lastReg = m[1]
			continue
		}
		if m := reField.FindStringSubmatch(line); m != nil && lastReg != "" {
			v, _ := strconv.Atoi(m[2])
			if v >= 0 && v < 32 {
				r := regs[lastReg]
				r.Fields[m[1]] = int64(1) << uint(v) // bit-field names carry their mask
				regs[lastReg] = r
			}
			continue
		}
		if m := reInc.FindStringSubmatch(line); m != nil {
			if err := parseInto(filepath.Join(filepath.Dir(header), m[1]), regs); err != nil {
				return err
			}
		}
		if !reBlank.MatchString(line) {
			lastReg = "" // a non-field line ends the run
		}
	}
	return nil
}

// ---- CMSIS-SVD (Espressif chips) ----

type svdField struct {
	Name        string `xml:"name"`
	BitOffset   string `xml:"bitOffset"`
	BitWidth    string `xml:"bitWidth"`
	BitRange    string `xml:"bitRange"`
	Lsb         string `xml:"lsb"`
	Msb         string `xml:"msb"`
	DerivedFrom string `xml:"derivedFrom,attr"`
}

type svdReg struct {
	Name          string     `xml:"name"`
	AddressOffset string     `xml:"addressOffset"`
	Size          string     `xml:"size"`
	DerivedFrom   string     `xml:"derivedFrom,attr"`
	Dim           string     `xml:"dim"`
	DimIncrement  string     `xml:"dimIncrement"`
	Fields        []svdField `xml:"fields>field"`
}

type svdCluster struct {
	Name          string   `xml:"name"`
	AddressOffset string   `xml:"addressOffset"`
	Dim           string   `xml:"dim"`
	DimIncrement  string   `xml:"dimIncrement"`
	Registers     []svdReg `xml:"register"`
}

type svdPeripheral struct {
	Name        string       `xml:"name"`
	BaseAddress string       `xml:"baseAddress"`
	DerivedFrom string       `xml:"derivedFrom,attr"`
	Registers   []svdReg     `xml:"registers>register"`
	Clusters    []svdCluster `xml:"registers>cluster"`
}

type svdDevice struct {
	Peripherals []svdPeripheral `xml:"peripherals>peripheral"`
}

func svdNum(s string) int64 {
	s = strings.TrimSpace(s)
	l := strings.ToLower(s)
	if strings.HasPrefix(l, "0x") {
		v, _ := strconv.ParseInt(l[2:], 16, 64)
		return v
	}
	v, _ := strconv.ParseInt(l, 0, 64)
	return v
}

// fieldBit resolves a field's bitOffset / bitRange ("[msb:lsb]") / lsb into a bit position; fieldWidth its
// bitWidth / bitRange / msb-lsb span into a width.
func fieldBit(f svdField) int {
	if f.BitOffset != "" {
		return int(svdNum(f.BitOffset))
	}
	if f.BitRange != "" {
		r := strings.Trim(f.BitRange, "[]")
		if i := strings.Index(r, ":"); i >= 0 {
			return int(svdNum(r[i+1:]))
		}
	}
	if f.Lsb != "" {
		return int(svdNum(f.Lsb))
	}
	return 0
}

func fieldWidth(f svdField) int {
	if f.BitWidth != "" {
		return int(svdNum(f.BitWidth))
	}
	if f.BitRange != "" {
		r := strings.Trim(f.BitRange, "[]")
		if i := strings.Index(r, ":"); i >= 0 {
			return int(svdNum(r[:i]) - svdNum(r[i+1:]) + 1)
		}
	}
	if f.Msb != "" && f.Lsb != "" {
		return int(svdNum(f.Msb) - svdNum(f.Lsb) + 1)
	}
	return 1
}

// fieldMask is what regs.X.FIELD compiles to: the field's mask, so set/clear see the bits it covers.
func fieldMask(f svdField) int64 {
	off, w := fieldBit(f), fieldWidth(f)
	if w >= 64 {
		return -1
	}
	return ((int64(1) << uint(w)) - 1) << uint(off)
}

// expandArray turns a register/cluster array ("KEY_%s", dim=8) into per-index copies.
func expandArray(r svdReg) []svdReg {
	dim := int(svdNum(r.Dim))
	if !strings.Contains(r.Name, "%s") || dim == 0 {
		return []svdReg{r}
	}
	var out []svdReg
	inc := svdNum(r.DimIncrement)
	base := svdNum(r.AddressOffset)
	for i := 0; i < dim; i++ {
		c := r
		c.Name = strings.ReplaceAll(r.Name, "%s", strconv.Itoa(i))
		c.AddressOffset = strconv.FormatInt(base+int64(i)*inc, 10)
		out = append(out, c)
	}
	return out
}

func parseSVD(file string) (map[string]regDef, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var dev svdDevice
	if err := xml.Unmarshal(data, &dev); err != nil {
		return nil, err
	}
	byName := map[string]svdPeripheral{}
	for _, p := range dev.Peripherals {
		byName[p.Name] = p
	}
	resolve := func(p svdPeripheral) svdPeripheral {
		if p.DerivedFrom != "" {
			if base, ok := byName[p.DerivedFrom]; ok {
				p.Registers = base.Registers
				p.Clusters = base.Clusters
			}
		}
		return p
	}
	regs := map[string]regDef{}
	add := func(periph string, r svdReg, offBase int64) {
		for _, rr := range expandArray(r) {
			addr := offBase + svdNum(rr.AddressOffset)
			bits := 32
			if rr.Size != "" {
				bits = int(svdNum(rr.Size))
			}
			name := periph + "_" + rr.Name
			d := regs[name]
			d.Addr = int(addr)
			d.Bits = bits
			if d.Fields == nil {
				d.Fields = map[string]int64{}
			}
			for _, f := range rr.Fields {
				d.Fields[f.Name] = fieldMask(f)
			}
			regs[name] = d
		}
	}
	for _, p := range dev.Peripherals {
		p = resolve(p)
		base := svdNum(p.BaseAddress)
		for _, r := range p.Registers {
			// register derivedFrom: inherit the twin's fields when it has none
			if r.DerivedFrom != "" && len(r.Fields) == 0 {
				for _, q := range p.Registers {
					if q.Name == r.DerivedFrom {
						r.Fields = q.Fields
						break
					}
				}
			}
			add(p.Name, r, base)
		}
		for _, cl := range p.Clusters {
			cinc := svdNum(cl.DimIncrement)
			count := int(svdNum(cl.Dim))
			if count == 0 {
				count = 1
			}
			for i := 0; i < count; i++ {
				coff := base + svdNum(cl.AddressOffset) + int64(i)*cinc
				for _, r := range cl.Registers {
					add(p.Name, r, coff)
				}
			}
		}
	}
	return regs, nil
}

// gpioPorts marks a chip's GPIO output registers with the board pins they drive (W616). Espressif: OUT/ENABLE
// banks. RP2040: the SIO GPIO bank (pins are the GPIO numbers on Pico-class boards). STM32: per-letter ports,
// numbered like the BluePill names (PA0-15=0-15, PB0-15=16-31, PC13=45). The other chips' register tables work
// without port marking - a variant's pin->pad map isn't in the board data.
func gpioPorts(chip string, regs map[string]regDef) {
	mark := func(match func(string) bool, pins []int) {
		for name, r := range regs {
			if match(name) {
				r.Port = pins
				regs[name] = r
			}
		}
	}
	if n, ok := gpioCount[chip]; ok { // Espressif
		low := make([]int, 0, 32)
		for i := 0; i < n && i < 32; i++ {
			low = append(low, i)
		}
		high := make([]int, 0, n)
		for i := 32; i < n; i++ {
			high = append(high, i)
		}
		mark(func(name string) bool {
			return strings.HasPrefix(name, "GPIO_OUT1") || strings.HasPrefix(name, "GPIO_ENABLE1")
		}, high)
		mark(func(name string) bool {
			return strings.HasPrefix(name, "GPIO_OUT") || strings.HasPrefix(name, "GPIO_ENABLE") || strings.HasPrefix(name, "GPIO_IN")
		}, low)
		return
	}
	switch chip {
	case "rp2040":
		all := make([]int, 30)
		for i := range all {
			all[i] = i
		}
		mark(func(name string) bool {
			return strings.HasPrefix(name, "SIO_GPIO_OUT") || strings.HasPrefix(name, "SIO_GPIO_OE") || name == "SIO_GPIO_IN"
		}, all)
	case "stm32f103": // BluePill numbering: PA0-15 = 0-15, PB0-15 = 16-31, PC13 = 45
		pa := make([]int, 16)
		pb := make([]int, 16)
		for i := 0; i < 16; i++ {
			pa[i] = i
			pb[i] = 16 + i
		}
		mark(func(name string) bool { return strings.HasPrefix(name, "GPIOA_") && name != "GPIOA_IDR" }, pa)
		mark(func(name string) bool { return strings.HasPrefix(name, "GPIOB_") && name != "GPIOB_IDR" }, pb)
		mark(func(name string) bool { return strings.HasPrefix(name, "GPIOC_") && name != "GPIOC_IDR" }, []int{45})
	}
}

func main() {
	toolchain := flag.String("toolchain", "", "avr-gcc toolchain root (contains avr/include/avr)")
	svdDir := flag.String("svd-dir", "", "directory of Espressif SVD files (esp32-core tools/ide-debug/svd)")
	vendorDir := flag.String("vendor-dir", "", "directory of vendor SVD files (regsgen/svd)")
	out := flag.String("out", "regs.json", "output file")
	flag.Parse()
	if *toolchain == "" {
		fmt.Fprintln(os.Stderr, "-toolchain is required")
		os.Exit(2)
	}
	all := map[string]map[string]regDef{}
	for chip, header := range chips {
		regs, err := parse(filepath.Join(*toolchain, "avr", "include", "avr", header))
		if err != nil {
			fmt.Fprintln(os.Stderr, chip+":", err)
			os.Exit(1)
		}
		for name, pins := range portPins[chip] {
			if r, ok := regs[name]; ok {
				r.Port = pins
				regs[name] = r
			}
		}
		all[chip] = regs
		fmt.Printf("%s: %d registers\n", chip, len(regs))
	}
	if *svdDir != "" {
		for chip, file := range svdChips {
			regs, err := parseSVD(filepath.Join(*svdDir, file))
			if err != nil {
				fmt.Fprintln(os.Stderr, chip+":", err)
				continue // a missing SVD leaves the chip on B413, not a broken generator
			}
			gpioPorts(chip, regs)
			all[chip] = regs
			fmt.Printf("%s: %d registers\n", chip, len(regs))
		}
	}
	if *vendorDir != "" {
		for chip, file := range vendorChips {
			regs, err := parseSVD(filepath.Join(*vendorDir, file))
			if err != nil {
				fmt.Fprintln(os.Stderr, chip+":", err)
				continue
			}
			gpioPorts(chip, regs)
			all[chip] = regs
			fmt.Printf("%s: %d registers\n", chip, len(regs))
		}
	}
	buf, err := json.MarshalIndent(all, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := os.WriteFile(*out, append(buf, '\n'), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("wrote", *out)
}
