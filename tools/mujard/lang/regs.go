package lang

import (
	_ "embed"
	"encoding/json"
	"strings"
)

// regs.json is generated from the chip makers' register descriptions (the AVR GCC io headers carry the same data
// as Microchip's ATDF files) by tools/mujard/lang/regsgen. It backs 'use mujard.regs': regs.PORTB.write(v) and
// friends (spec 10.4).

//go:embed regs.json
var regsJSON []byte

// RegDef is one chip register: where it sits, how wide, its named bits, and - for GPIO port registers - which
// board pins the register's bits drive (for the write-onto-a-declared-pin warning).
type RegDef struct {
	Addr   int              `json:"addr"`           // data-space address
	Bits   int              `json:"bits"`           // 8, 16, 32
	Fields map[string]int64 `json:"fields"`         // bit-field name -> its mask (single bit: 1<<pos)
	Port   []int            `json:"port,omitempty"` // board pins this register drives (GPIO ports only)
}

// chipRegsByFamily is board family ("ATmega328P") -> the regs.json chip key.
var chipRegsByFamily = map[string]string{
	"ATmega328P":    "atmega328p",
	"ATmega32U4":    "atmega32u4",
	"ATmega2560":    "atmega2560",
	"ESP32":         "esp32",
	"ESP8266":       "esp8266",
	"ESP32-S3":      "esp32s3",
	"ESP32-S2":      "esp32s2",
	"ESP32-C3":      "esp32c3",
	"ESP32-C6":      "esp32c6",
	"RP2040":        "rp2040",
	"nRF52840":      "nrf52840",
	"SAMD21G18":     "samd21",
	"Renesas RA4M1": "ra4m1",
	"STM32F103":     "stm32f103",
}

var chipRegs map[string]map[string]RegDef // chip -> register -> def

// ChipRegs is the whole register table, for generated documentation ('mujard doc').
func ChipRegs() map[string]map[string]RegDef { return chipRegs }

// RegChipKey is a board family's chip key in ChipRegs, for generated documentation.
func RegChipKey(family string) (string, bool) {
	k, ok := chipRegsByFamily[family]
	return k, ok
}

func regsOf(b *Board) (map[string]RegDef, bool) {
	if b == nil {
		return nil, false
	}
	key, ok := chipRegsByFamily[b.Family]
	if !ok {
		return nil, false
	}
	t, ok := chipRegs[key]
	return t, ok
}

func regMax(r RegDef) uint64 {
	if r.Bits >= 64 {
		return ^uint64(0)
	}
	return uint64(1)<<uint(r.Bits) - 1
}

func regTypeName(reg string) *Type { return &Type{K: TObject, Name: "regs:" + reg} }

func regName(t *Type) string {
	if t == nil || t.K != TObject || !strings.HasPrefix(t.Name, "regs:") {
		return ""
	}
	return strings.TrimPrefix(t.Name, "regs:")
}

func init() {
	chipRegs = map[string]map[string]RegDef{}
	if err := json.Unmarshal(regsJSON, &chipRegs); err != nil {
		panic("regs.json: " + err.Error())
	}
}
