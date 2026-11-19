// Command descgen writes the library descriptions that are too big to write by hand from the library's own
// headers. Today that is U8g2 (descs/u8g2.mlib): its ~4,900 display classes become families (one per bus, with the
// constructor read from the header and checked to be the same for every display of that bus) and its ~2,000 fonts
// become constants. The U8G2 base class's methods are written below, checked against U8g2lib.h by hand.
//
//	go run ./lang/descgen -u8g2 <U8g2 library folder> -out lang/descs/u8g2.mlib
//
// The description says which library version it was made from; rerun it when the pinned version changes.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"mujard/lang"
)

func main() {
	dir := flag.String("u8g2", "", "the U8g2 library folder (with library.properties and src/)")
	out := flag.String("out", "", "where to write the description")
	flag.Parse()
	if *dir == "" || *out == "" {
		flag.Usage()
		os.Exit(2)
	}
	b, err := u8g2(*dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "descgen:", err)
		os.Exit(1)
	}
	if err := os.WriteFile(*out, b, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, "descgen:", err)
		os.Exit(1)
	}
}

// u8g2Base is the U8G2 class every display extends (U8g2lib.h, class U8G2 : public Print), checked by hand.
// Pointers the program can't make (tile and buffer pointers, Print&, U8G2LOG&, kerning tables) are left out.
const u8g2Base = `{
  "methods": {
    "begin": {"params": [{"name": "menu_select_pin", "type": "int", "optional": true}, {"name": "menu_next_pin", "type": "int", "optional": true}, {"name": "menu_prev_pin", "type": "int", "optional": true}, {"name": "menu_up_pin", "type": "int", "optional": true}, {"name": "menu_down_pin", "type": "int", "optional": true}, {"name": "menu_home_pin", "type": "int", "optional": true}], "result": "bool"},
    "beginSimple": {"params": [], "result": "void"},
    "setBusClock": {"params": [{"name": "clock_speed", "type": "int"}], "result": "void"},
    "setI2CAddress": {"params": [{"name": "adr", "type": "int"}], "result": "void"},
    "enableUTF8Print": {"params": [], "result": "void"},
    "disableUTF8Print": {"params": [], "result": "void"},
    "getCols": {"params": [], "result": "int"},
    "getRows": {"params": [], "result": "int"},
    "initDisplay": {"params": [], "result": "void"},
    "clearDisplay": {"params": [], "result": "void"},
    "setPowerSave": {"params": [{"name": "is_enable", "type": "int"}], "result": "void"},
    "setFlipMode": {"params": [{"name": "mode", "type": "int"}], "result": "void"},
    "setContrast": {"params": [{"name": "value", "type": "int"}], "result": "void"},
    "setDisplayRotation": {"params": [{"name": "rotation", "type": "U8G2Rotation"}], "result": "void"},
    "setMaxClipWindow": {"params": [], "result": "void"},
    "setClipWindow": {"params": [{"name": "x0", "type": "int"}, {"name": "y0", "type": "int"}, {"name": "x1", "type": "int"}, {"name": "y1", "type": "int"}], "result": "void"},
    "getDisplayHeight": {"params": [], "result": "int"},
    "getDisplayWidth": {"params": [], "result": "int"},
    "sendBuffer": {"params": [], "result": "void"},
    "clearBuffer": {"params": [], "result": "void"},
    "firstPage": {"params": [], "result": "void"},
    "nextPage": {"params": [], "result": "int"},
    "getBufferSize": {"params": [], "result": "int"},
    "setAutoPageClear": {"params": [{"name": "mode", "type": "int"}], "result": "void"},
    "updateDisplayArea": {"params": [{"name": "tx", "type": "int"}, {"name": "ty", "type": "int"}, {"name": "tw", "type": "int"}, {"name": "th", "type": "int"}], "result": "void"},
    "updateDisplay": {"params": [], "result": "void"},
    "refreshDisplay": {"params": [], "result": "void"},
    "setDrawColor": {"params": [{"name": "color_index", "type": "int"}], "result": "void"},
    "getDrawColor": {"params": [], "result": "int"},
    "drawPixel": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}], "result": "void"},
    "drawHLine": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "w", "type": "int"}], "result": "void"},
    "drawVLine": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "h", "type": "int"}], "result": "void"},
    "drawFrame": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "w", "type": "int"}, {"name": "h", "type": "int"}], "result": "void"},
    "drawRFrame": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "w", "type": "int"}, {"name": "h", "type": "int"}, {"name": "r", "type": "int"}], "result": "void"},
    "drawBox": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "w", "type": "int"}, {"name": "h", "type": "int"}], "result": "void"},
    "drawRBox": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "w", "type": "int"}, {"name": "h", "type": "int"}, {"name": "r", "type": "int"}], "result": "void"},
    "drawButtonUTF8": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "flags", "type": "int"}, {"name": "width", "type": "int"}, {"name": "padding_h", "type": "int"}, {"name": "padding_v", "type": "int"}, {"name": "text", "type": "text"}], "result": "void"},
    "drawCircle": {"params": [{"name": "x0", "type": "int"}, {"name": "y0", "type": "int"}, {"name": "rad", "type": "int"}, {"name": "opt", "type": "int", "optional": true}], "result": "void"},
    "drawDisc": {"params": [{"name": "x0", "type": "int"}, {"name": "y0", "type": "int"}, {"name": "rad", "type": "int"}, {"name": "opt", "type": "int", "optional": true}], "result": "void"},
    "drawEllipse": {"params": [{"name": "x0", "type": "int"}, {"name": "y0", "type": "int"}, {"name": "rx", "type": "int"}, {"name": "ry", "type": "int"}, {"name": "opt", "type": "int", "optional": true}], "result": "void"},
    "drawFilledEllipse": {"params": [{"name": "x0", "type": "int"}, {"name": "y0", "type": "int"}, {"name": "rx", "type": "int"}, {"name": "ry", "type": "int"}, {"name": "opt", "type": "int", "optional": true}], "result": "void"},
    "drawArc": {"params": [{"name": "x0", "type": "int"}, {"name": "y0", "type": "int"}, {"name": "rad", "type": "int"}, {"name": "start", "type": "int"}, {"name": "end", "type": "int"}], "result": "void"},
    "drawLine": {"params": [{"name": "x1", "type": "int"}, {"name": "y1", "type": "int"}, {"name": "x2", "type": "int"}, {"name": "y2", "type": "int"}], "result": "void"},
    "drawTriangle": {"params": [{"name": "x0", "type": "int"}, {"name": "y0", "type": "int"}, {"name": "x1", "type": "int"}, {"name": "y1", "type": "int"}, {"name": "x2", "type": "int"}, {"name": "y2", "type": "int"}], "result": "void"},
    "setBitmapMode": {"params": [{"name": "is_transparent", "type": "int"}], "result": "void"},
    "setFont": {"params": [{"name": "font", "type": "U8G2Font"}], "result": "void"},
    "setFontMode": {"params": [{"name": "is_transparent", "type": "int"}], "result": "void"},
    "setFontDirection": {"params": [{"name": "dir", "type": "int"}], "result": "void"},
    "getAscent": {"params": [], "result": "int"},
    "getDescent": {"params": [], "result": "int"},
    "setFontPosBaseline": {"params": [], "result": "void"},
    "setFontPosBottom": {"params": [], "result": "void"},
    "setFontPosTop": {"params": [], "result": "void"},
    "setFontPosCenter": {"params": [], "result": "void"},
    "setFontRefHeightText": {"params": [], "result": "void"},
    "setFontRefHeightExtendedText": {"params": [], "result": "void"},
    "setFontRefHeightAll": {"params": [], "result": "void"},
    "drawGlyph": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "encoding", "type": "int"}], "result": "int"},
    "drawGlyphX2": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "encoding", "type": "int"}], "result": "int"},
    "drawStr": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "s", "type": "text"}], "result": "int"},
    "drawStrX2": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "s", "type": "text"}], "result": "int"},
    "drawUTF8": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "s", "type": "text"}], "result": "int"},
    "drawUTF8X2": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}, {"name": "s", "type": "text"}], "result": "int"},
    "getStrWidth": {"params": [{"name": "s", "type": "text"}], "result": "int"},
    "getUTF8Width": {"params": [{"name": "s", "type": "text"}], "result": "int"},
    "getXOffsetGlyph": {"params": [{"name": "encoding", "type": "int"}], "result": "int"},
    "getXOffsetUTF8": {"params": [{"name": "utf8", "type": "text"}], "result": "int"},
    "userInterfaceSelectionList": {"params": [{"name": "title", "type": "text"}, {"name": "start_pos", "type": "int"}, {"name": "sl", "type": "text"}], "result": "int"},
    "userInterfaceMessage": {"params": [{"name": "title1", "type": "text"}, {"name": "title2", "type": "text"}, {"name": "title3", "type": "text"}, {"name": "buttons", "type": "text"}], "result": "int"},
    "home": {"params": [], "result": "void"},
    "clear": {"params": [], "result": "void"},
    "noDisplay": {"params": [], "result": "void"},
    "display": {"params": [], "result": "void"},
    "setCursor": {"params": [{"name": "x", "type": "int"}, {"name": "y", "type": "int"}], "result": "void"},
    "getCursorX": {"params": [], "result": "int"},
    "getCursorY": {"params": [], "result": "int"},
    "sleepOn": {"params": [], "result": "void"},
    "sleepOff": {"params": [], "result": "void"},
    "setColorIndex": {"params": [{"name": "color_index", "type": "int"}], "result": "void"},
    "getColorIndex": {"params": [], "result": "int"},
    "getFontAscent": {"params": [], "result": "int"},
    "getFontDescent": {"params": [], "result": "int"},
    "getMaxCharHeight": {"params": [], "result": "int"},
    "getMaxCharWidth": {"params": [], "result": "int"},
    "getHeight": {"params": [], "result": "int"},
    "getWidth": {"params": [], "result": "int"},
    "print": {"params": [{"name": "value", "type": "any"}, {"name": "base", "type": "any", "optional": true}], "result": "int"},
    "println": {"params": [{"name": "value", "type": "any", "optional": true}], "result": "int"}
  }
}`

var (
	classRx = regexp.MustCompile(`(?m)^class (U8G2_\w+) : public U8G2 \{\s*\n\s*public: U8G2_\w+\(([^)]*)\)`)
	fontRx  = regexp.MustCompile(`(?m)^extern const uint8_t (u8g2_font_\w+)\[\]`)
	paramRx = regexp.MustCompile(`^(?:const\s+)?(\w+)\s*(\*?)\s*(\w+)\s*(=.*)?$`)
	mids    = []string{"F", "1", "2"}
)

// variants are the bus suffixes, longest first so "2ND_4W_HW_SPI" isn't read as "4W_HW_SPI".
var variants = []string{"2ND_4W_HW_SPI", "2ND_HW_I2C", "2ND_HW_SPI", "4W_SW_SPI", "4W_HW_SPI", "3W_SW_SPI", "3W_HW_SPI",
	"HW_I2C", "SW_I2C", "HW_SPI", "SW_SPI", "8080", "6800"}

func u8g2(dir string) ([]byte, error) {
	version := ""
	if props, err := os.ReadFile(filepath.Join(dir, "library.properties")); err == nil {
		for _, l := range strings.Split(string(props), "\n") {
			if v, ok := strings.CutPrefix(strings.TrimSpace(l), "version="); ok {
				version = v
			}
		}
	}
	lib, err := os.ReadFile(filepath.Join(dir, "src", "U8g2lib.h"))
	if err != nil {
		return nil, err
	}
	clib, err := os.ReadFile(filepath.Join(dir, "src", "clib", "u8g2.h"))
	if err != nil {
		return nil, err
	}

	fam := &lang.LibFamily{Prefix: "U8G2_", Mids: mids, Variants: map[string]*lang.LibVariant{}}
	type key struct{ variant, stem string }
	seen := map[key]map[string]bool{} // which mids each display has
	ctors := map[string]string{}      // variant -> the constructor parameter list as written
	skipped := 0
	for _, m := range classRx.FindAllStringSubmatch(string(lib), -1) {
		name, params := m[1], strings.Join(strings.Fields(m[2]), " ")
		variant, stem, mid := split(name)
		if variant == "" {
			skipped++ // parallel KS0108-style displays, U8G2_BITMAP, U8G2_NULL: no bus suffix
			continue
		}
		if prev, ok := ctors[variant]; ok && prev != params {
			return nil, fmt.Errorf("%s: constructor (%s) differs from the rest of %s (%s)", name, params, variant, prev)
		}
		if _, ok := ctors[variant]; !ok {
			ctor, err := ctorOf(params)
			if err != nil {
				return nil, fmt.Errorf("%s: %v", name, err)
			}
			ctors[variant] = params
			fam.Variants[variant] = &lang.LibVariant{Ctor: ctor}
		}
		k := key{variant, stem}
		if seen[k] == nil {
			seen[k] = map[string]bool{}
		}
		seen[k][mid] = true
	}
	n := 0
	for k, have := range seen {
		if len(have) != len(mids) {
			return nil, fmt.Errorf("U8G2_%s_?_%s does not come with every buffer size (%v): the family can't describe it", k.stem, k.variant, have)
		}
		fam.Variants[k.variant].Stems = append(fam.Variants[k.variant].Stems, k.stem)
		n += len(have)
	}
	for _, v := range fam.Variants {
		sort.Strings(v.Stems)
	}

	var base lang.LibClass
	if err := json.Unmarshal([]byte(u8g2Base), &base); err != nil {
		return nil, fmt.Errorf("the U8G2 base class: %v", err)
	}
	consts := map[string]string{
		"U8G2_R0": "U8G2Rotation", "U8G2_R1": "U8G2Rotation", "U8G2_R2": "U8G2Rotation", "U8G2_R3": "U8G2Rotation",
		"U8G2_MIRROR": "U8G2Rotation", "U8G2_MIRROR_VERTICAL": "U8G2Rotation",
		"U8X8_PIN_NONE": "int", "U8G2_DRAW_ALL": "int", "U8G2_DRAW_UPPER_RIGHT": "int", "U8G2_DRAW_UPPER_LEFT": "int",
		"U8G2_DRAW_LOWER_LEFT": "int", "U8G2_DRAW_LOWER_RIGHT": "int",
		"U8G2_BTN_BW0": "int", "U8G2_BTN_BW1": "int", "U8G2_BTN_BW2": "int", "U8G2_BTN_BW3": "int", "U8G2_BTN_SHADOW0": "int",
		"U8G2_BTN_SHADOW1": "int", "U8G2_BTN_SHADOW2": "int", "U8G2_BTN_INV": "int", "U8G2_BTN_HCENTER": "int", "U8G2_BTN_XFRAME": "int",
	}
	fonts := 0
	for _, m := range fontRx.FindAllStringSubmatch(string(clib), -1) {
		consts[m[1]] = "U8G2Font"
		fonts++
	}
	d := lang.LibDesc{
		Name:     "U8g2",
		Header:   "U8g2lib.h",
		Classes:  map[string]*lang.LibClass{"U8G2": &base},
		Families: map[string]*lang.LibFamily{"U8G2": fam},
		// The font names are pointers into flash and the rotations pointers to callbacks: objects of their own
		// type, so a font can't be passed where a rotation goes.
		Constants: consts,
	}
	body, err := json.MarshalIndent(d, "", " ")
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(os.Stderr, "U8g2 %s: %d display classes in %d bus families (%d without a bus left out), %d fonts\n",
		version, n, len(fam.Variants), skipped, fonts)
	return append(body, '\n'), nil
}

// split reads U8G2_<stem>_<mid>_<variant>.
func split(name string) (variant, stem, mid string) {
	rest := strings.TrimPrefix(name, "U8G2_")
	for _, v := range variants {
		r, ok := strings.CutSuffix(rest, "_"+v)
		if !ok {
			continue
		}
		for _, m := range mids {
			if s, ok := strings.CutSuffix(r, "_"+m); ok && s != "" {
				return v, s, m
			}
		}
	}
	return "", "", ""
}

// ctorOf turns "const u8g2_cb_t *rotation, uint8_t clock, uint8_t reset = U8X8_PIN_NONE" into a signature.
func ctorOf(params string) (lang.LibFn, error) {
	var fn lang.LibFn
	for _, p := range strings.Split(params, ",") {
		m := paramRx.FindStringSubmatch(strings.TrimSpace(p))
		if m == nil {
			return fn, fmt.Errorf("can't read the parameter %q", p)
		}
		typ := "int"
		switch {
		case m[1] == "u8g2_cb_t" && m[2] == "*":
			typ = "U8G2Rotation"
		case m[2] == "*" || (m[1] != "uint8_t" && m[1] != "uint16_t" && m[1] != "int"):
			return fn, fmt.Errorf("unexpected parameter type in %q", p)
		}
		fn.Params = append(fn.Params, lang.LibParam{Name: m[3], Type: typ, Optional: m[4] != ""})
	}
	fn.Result = "U8G2"
	return fn, nil
}
