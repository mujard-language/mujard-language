package lang

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const neopixelMLib = `{
  "name": "NeoPixel",
  "header": "Adafruit_NeoPixel.h",
  "architectures": ["avr", "esp32"],
  "pins": [6],
  "classes": {
    "Adafruit_NeoPixel": {
      "ctors": [{"params": [
        {"name": "count", "type": "int"},
        {"name": "pin", "type": "int"},
        {"name": "flags", "type": "int", "optional": true}
      ]}],
      "methods": {
        "begin": {"params": [], "result": "void"},
        "show": {"params": [], "result": "void"},
        "setPixelColor": {"params": [
          {"name": "n", "type": "int"}, {"name": "r", "type": "int"},
          {"name": "g", "type": "int"}, {"name": "b", "type": "int"}
        ], "result": "void"},
        "numPixels": {"params": [], "result": "int"}
      },
      "fields": {"brightness": "int"}
    }
  },
  "constants": {"NEO_GRB": "int", "NEO_KHZ800": "int"},
  "functions": {"gamma8": {"params": [{"name": "x", "type": "int"}], "result": "int"}}
}`

const neoProgram = `board uno
use "NeoPixel"
strip = Adafruit_NeoPixel(60, 6, NEO_GRB)
at start:
    strip.begin()
    strip.set_pixel_color(0, 255, 0, 0)
    print strip.num_pixels()
`

func hasDiag(diags []Diagnostic, code string) bool {
	for _, d := range diags {
		if d.Code == code {
			return true
		}
	}
	return false
}

func TestLibCheckedCalls(t *testing.T) {
	_, diags := CheckSourceModules(neoProgram, map[string]string{"neo.mlib": neopixelMLib})
	for _, d := range diags {
		t.Logf("diag: %s", d)
	}
	if FirstError(diags) != nil {
		t.Fatalf("described library calls should check clean: %v", FirstError(diags))
	}
	if hasDiag(diags, "W611") {
		t.Error("a described library must not warn W611")
	}
}

func TestLibChecksArgsAndMembers(t *testing.T) {
	bad := `board uno
use "NeoPixel"
strip = Adafruit_NeoPixel(60, 6, NEO_GRB)
at start:
    strip.set_pixel_color(0)
    strip.no_such_thing()
    strip.unknown_field = 3
`
	_, diags := CheckSourceModules(bad, map[string]string{"neo.mlib": neopixelMLib})
	if !hasDiag(diags, "C307") {
		t.Error("a wrong argument count should be C307")
	}
	if !hasDiag(diags, "C308") {
		t.Error("an unknown member should be C308")
	}
}

func TestLibWithoutDescriptionStaysUnchecked(t *testing.T) {
	src := `board uno
use "Mystery"
thing = WeirdObject(1)
at start:
    thing.anything()
`
	_, diags := CheckSourceModules(src, nil)
	if !hasDiag(diags, "W611") {
		t.Error("an undescribed library should warn W611")
	}
	if FirstError(diags) != nil {
		t.Errorf("undescribed library calls still compile: %v", FirstError(diags))
	}
}

func TestUseProjectFileIsNotALibrary(t *testing.T) {
	src := "board uno\nuse \"lib/crc.h\"\nat start:\n    pass\n"
	_, diags := CheckSourceModules(src, map[string]string{"lib/crc.h": "unsigned crc();"})
	if hasDiag(diags, "W611") {
		t.Error("a project file is not a library - no W611")
	}
	if FirstError(diags) != nil {
		t.Errorf("project-file use should check clean: %v", FirstError(diags))
	}
}

func TestLibArchitectures(t *testing.T) {
	_, diags := CheckSourceModules(neoProgram, map[string]string{"neo.mlib": neopixelMLib})
	if hasDiag(diags, "B412") {
		t.Error("uno is avr, the desc allows it - no B412")
	}
	espOnly := strings.Replace(neopixelMLib, `"avr", "esp32"`, `"esp32"`, 1)
	_, diags = CheckSourceModules(neoProgram, map[string]string{"neo.mlib": espOnly})
	if !hasDiag(diags, "B412") {
		t.Error("an esp32-only library on an avr board should be B412")
	}
}

func TestLibPinsWarn(t *testing.T) {
	src := `board uno
use "NeoPixel"
strip = Adafruit_NeoPixel(60, 6, NEO_GRB)
led = pin 6 as output
at start:
    strip.begin()
`
	_, diags := CheckSourceModules(src, map[string]string{"neo.mlib": neopixelMLib})
	if !hasDiag(diags, "W609") {
		t.Error("a pin the library keeps should warn W609")
	}
}

func TestLibIdfBoards(t *testing.T) {
	_, diags := CheckSourceModules("board uno\nuse idf \"esp32-camera\"\nat start:\n    pass\n", nil)
	if !hasDiag(diags, "B412") {
		t.Error("an ESP-IDF component on an AVR board should be B412")
	}
	_, diags = CheckSourceModules("board esp32\nuse idf \"esp32-camera\"\nat start:\n    pass\n", nil)
	if hasDiag(diags, "B412") {
		t.Error("esp32-camera on an ESP32 board is fine")
	}
}

// -- generated C++ --

func generateWithLib(t *testing.T, src, lib string) *Sketch {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "neo.mlib"), []byte(lib), 0o644); err != nil {
		t.Fatal(err)
	}
	sk, diags := CompileIn(src, "sketch.muj", dir)
	for _, d := range diags {
		t.Logf("diag: %s", d)
	}
	if e := FirstError(diags); e != nil {
		t.Fatalf("compile: %v", e)
	}
	if sk == nil {
		t.Fatal("no sketch")
	}
	return sk
}

func TestLibCodegen(t *testing.T) {
	sk := generateWithLib(t, neoProgram, neopixelMLib)
	prog := sk.Files["mujard_program.cpp"]
	if !strings.Contains(prog, "#include <Adafruit_NeoPixel.h>") {
		t.Error("the described header should be included")
	}
	if !strings.Contains(prog, "static Adafruit_NeoPixel strip(60, 6, NEO_GRB);") {
		t.Errorf("the described constructor should be emitted, got:\n%s", prog)
	}
	if !strings.Contains(prog, "strip.setPixelColor(0, 255, 0, 0)") {
		t.Error("a snake_case call must emit the real camelCase member")
	}
	if !strings.Contains(prog, "strip.numPixels()") {
		t.Error("num_pixels -> numPixels")
	}
	libs, ok := sk.Files["mujard_libs.json"]
	if !ok {
		t.Fatal("mujard_libs.json should be generated")
	}
	if !strings.Contains(libs, `"NeoPixel"`) {
		t.Error("the manifest should name the library")
	}
	if len(sk.Libraries) != 1 || sk.Libraries[0].Name != "NeoPixel" {
		t.Errorf("Sketch.Libraries should carry the dep, got %+v", sk.Libraries)
	}
}

func TestLibVersionPin(t *testing.T) {
	src := "board uno\nuse \"NeoPixel\" version \"1.12.3\"\nat start:\n    pass\n"
	sk := generateWithLib(t, src, neopixelMLib)
	if len(sk.Libraries) != 1 || sk.Libraries[0].Version != "1.12.3" {
		t.Errorf("the version pin should land in the manifest, got %+v", sk.Libraries)
	}
}

func TestLibPio(t *testing.T) {
	dir := t.TempDir()
	src := "board uno\nuse pio \"bblanchon/ArduinoJson\"\nat start:\n    pass\n"
	sk, diags := CompileIn(src, "sketch.muj", dir)
	if e := FirstError(diags); e != nil {
		t.Fatalf("compile: %v", e)
	}
	if !strings.Contains(sk.Files["mujard_program.cpp"], "#include <ArduinoJson.h>") {
		t.Error("a pio library includes the last path segment as its header")
	}
	if len(sk.Libraries) != 1 || sk.Libraries[0].From != "pio" || sk.Libraries[0].Name != "bblanchon/ArduinoJson" {
		t.Errorf("pio dep: %+v", sk.Libraries)
	}
}

func TestLibFileInclude(t *testing.T) {
	dir := t.TempDir()
	sub := filepath.Join(dir, "lib")
	if err := os.MkdirAll(sub, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sub, "crc.h"), []byte("unsigned crc16();"), 0o644); err != nil {
		t.Fatal(err)
	}
	src := "board uno\nuse \"lib/crc.h\"\nat start:\n    pass\n"
	sk, diags := CompileIn(src, "sketch.muj", dir)
	if e := FirstError(diags); e != nil {
		t.Fatalf("compile: %v", e)
	}
	if !strings.Contains(sk.Files["mujard_program.cpp"], `#include "lib/crc.h"`) {
		t.Error("a project file is included with quotes")
	}
	if len(sk.Libraries) != 0 {
		t.Error("a project file is not a library dependency")
	}
}

func TestLibAliasFields(t *testing.T) {
	src := `board uno
use "NeoPixel"
strip = Adafruit_NeoPixel(60, 6, NEO_GRB)
at start:
    print strip.num_pixels
`
	sk := generateWithLib(t, src, neopixelMLib)
	if !strings.Contains(sk.Files["mujard_program.cpp"], "strip.numPixels") {
		t.Error("a snake_case field resolves to the real member")
	}
}

// -- the auto-wrapper: one pasted .h -> the .mlib a 'use "Name"' binds --

func TestDescribeHeaderSource(t *testing.T) {
	d, err := DescribeHeaderSource("MAX30102", "MAX30102.h", `class MAX30102 {
public:
    bool begin();
    int getHeartRate();
    void setLEDIntensity(int level);
};`)
	if err != nil {
		t.Fatal(err)
	}
	c := d.Classes["MAX30102"]
	if c == nil {
		t.Fatal("the class the header declares")
	}
	// The description keeps the C++ spellings (snakeName/libAlias gives the program the snake_case alias).
	if c.Methods["begin"].Result != "bool" || c.Methods["getHeartRate"].Result != "int" || c.Methods["setLEDIntensity"].Result != "void" {
		t.Errorf("methods + types: %v", c.Methods)
	}
	if c.Methods["setLEDIntensity"].Params[0].Name != "level" {
		t.Error("the parameter name survives")
	}
}

func TestDescribeHeaderSourceBinds(t *testing.T) {
	lib, err := DescribeHeaderSourceJSON("MAX30102", "MAX30102.h", `class MAX30102 {
public:
    bool begin();
    int getHeartRate();
    void setLEDIntensity(int level);
};`)
	if err != nil {
		t.Fatal(err)
	}
	// The generated .mlib makes 'use "MAX30102"' check the class's calls exactly like a folder-described library.
	src := `board esp32
use "MAX30102"
sensor = MAX30102()
at start:
    sensor.begin()
    print sensor.get_heart_rate()
    sensor.set_led_intensity(50)
`
	sk := generateWithLib(t, src, lib)
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{"sensor.begin()", "sensor.getHeartRate()", "sensor.setLEDIntensity"} {
		if !strings.Contains(cpp, want) {
			t.Errorf("the wrapped call emits real C++: %s", want)
		}
	}
}

func TestDescribeHeaderSourceNothing(t *testing.T) {
	if _, err := DescribeHeaderSource("empty", "empty.h", `// no classes or functions in here`); err == nil {
		t.Error("a header with nothing describable is an error, not a silent .mlib")
	}
}

func TestDescribePioLibrary(t *testing.T) {
	dir := t.TempDir()
	inc := filepath.Join(dir, "include")
	if err := os.MkdirAll(inc, 0o755); err != nil {
		t.Fatal(err)
	}
	libJSON := `{"name": "PioSensor", "platforms": ["espressif32", "atmelavr"], "headers": ["PioSensor.h"]}`
	if err := os.WriteFile(filepath.Join(dir, "library.json"), []byte(libJSON), 0o644); err != nil {
		t.Fatal(err)
	}
	hdr := `class PioSensor {
public:
    bool begin();
    float readLux();
};`
	if err := os.WriteFile(filepath.Join(inc, "PioSensor.h"), []byte(hdr), 0o644); err != nil {
		t.Fatal(err)
	}
	d, err := DescribeLibrary(dir)
	if err != nil {
		t.Fatal(err)
	}
	if d.Name != "PioSensor" || d.Header != "PioSensor.h" {
		t.Errorf("pio name/header: %+v", d)
	}
	if len(d.Architectures) != 2 || d.Architectures[0] != "esp32" || d.Architectures[1] != "avr" {
		t.Errorf("pio architectures mapped to Mujard families: %v", d.Architectures)
	}
	if d.Classes["PioSensor"] == nil || d.Classes["PioSensor"].Methods["readLux"] == nil {
		t.Errorf("pio include/ header parsed: %+v", d.Classes)
	}
	b, err := DescribeLibraryJSON(dir)
	if err != nil {
		t.Fatal(err)
	}
	src := `board esp32
use pio "acme/PioSensor"
s = PioSensor()
at start:
    s.begin()
    print s.read_lux()
`
	_, diags := CheckSourceModules(src, map[string]string{"piosensor.mlib": b})
	if e := FirstError(diags); e != nil {
		t.Errorf("use pio with generated .mlib should check clean: %v", e)
	}
}

// -- the descriptions the compiler ships --

func TestBuiltinDescsChecked(t *testing.T) {
	if len(builtinDescs) == 0 {
		t.Fatal("no embedded library descriptions")
	}
	// A program using a shipped description needs no .mlib around and gets checked calls.
	src := `board uno
use "Adafruit NeoPixel"
strip = Adafruit_NeoPixel(60, 6, NEO_GRB)
at start:
    strip.begin()
    strip.set_pixel_color(0, 255, 0, 0)
    print strip.num_pixels()
`
	_, diags := CheckSourceModules(src, nil)
	if hasDiag(diags, "W611") {
		t.Error("a shipped description should not warn W611")
	}
	if e := FirstError(diags); e != nil {
		t.Errorf("shipped-description calls should check clean: %v", e)
	}
	// And it really checks them: a wrong call is a diagnostic, not silently let through.
	bad := `board uno
use "DHT sensor library"
d = DHT(4, DHT22)
at start:
    d.read_temperature(1, 2, 3, 4)
    d.no_such()
`
	_, diags = CheckSourceModules(bad, nil)
	if !hasDiag(diags, "C307") {
		t.Error("too many arguments against a shipped desc should be C307")
	}
	if !hasDiag(diags, "C308") {
		t.Error("an unknown member on a shipped desc should be C308")
	}
	// A project .mlib still wins over the shipped one.
	project := `{"name":"Adafruit NeoPixel","header":"X.h","classes":{"Adafruit_NeoPixel":{"ctors":[{"params":[],"result":"Adafruit_NeoPixel"}],"methods":{"only_this":{"params":[],"result":"void"}}}}}`
	_, diags = CheckSourceModules(`board uno
use "Adafruit NeoPixel"
s = Adafruit_NeoPixel()
at start:
    s.only_this()
`, map[string]string{"neo.mlib": project})
	if e := FirstError(diags); e != nil {
		t.Errorf("a project description should override the shipped one: %v", e)
	}
}

// u8g2Program uses two U8g2 displays on different buses.
var u8g2Program = `board uno
use "U8g2"
# A page buffer (_2_, 256 bytes): the full-frame _F_ one (1 KB) plus the runtime is more RAM than an Uno has.
oled = U8G2_SSD1306_128X64_NONAME_2_HW_I2C(U8G2_R0, U8X8_PIN_NONE)
spi = U8G2_ST7920_128X64_1_SW_SPI(U8G2_R2, 13, 11, 10)
at start:
    oled.begin()
    oled.clear_buffer()
    oled.setFont(u8g2_font_ncenB08_tr)
    oled.drawStr(0, 10, "Hello")
    oled.send_buffer()
    print oled.getDisplayWidth()
    spi.begin()
`

// U8g2's displays are a family: any real display class works, with its bus's constructor and the U8G2 methods.
func TestU8g2Family(t *testing.T) {
	sk, diags := Compile(u8g2Program, "u8.muj")
	if e := FirstError(diags); e != nil || sk == nil {
		t.Fatalf("a U8g2 program should compile: %v", diags)
	}
	if hasDiag(diags, "W611") {
		t.Error("U8g2 is described: no W611")
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{"static U8G2_SSD1306_128X64_NONAME_2_HW_I2C oled(U8G2_R0, U8X8_PIN_NONE);", "oled.clearBuffer()", "oled.setFont(u8g2_font_ncenB08_tr)", "oled.sendBuffer()"} {
		if !strings.Contains(cpp, want) {
			t.Errorf("generated C++ lacks %q", want)
		}
	}
	bad := map[string]string{
		"a display U8g2 doesn't have":  "oled = U8G2_SSD9999_128X64_NONAME_F_HW_I2C(U8G2_R0)\n",
		"a buffer size that isn't one": "oled = U8G2_SSD1306_128X64_NONAME_X_HW_I2C(U8G2_R0)\n",
		"software I2C needs its pins":  "oled = U8G2_SSD1306_128X64_NONAME_F_SW_I2C(U8G2_R0)\n",
		"a font is not a rotation":     "oled = U8G2_SSD1306_128X64_NONAME_F_HW_I2C(u8g2_font_ncenB08_tr)\n",
		"no such method":               "oled = U8G2_SSD1306_128X64_NONAME_F_HW_I2C(U8G2_R0)\nat start:\n    oled.draw_star(1, 2)\n",
		"the base isn't a display":     "oled = U8G2(U8G2_R0)\n",
	}
	for name, body := range bad {
		if _, d := CheckSourceModules("board uno\nuse \"U8g2\"\n"+body, nil); FirstError(d) == nil {
			t.Errorf("%s: should not check clean", name)
		}
	}
}

// Every shipped description must parse (loadDescs skips a broken one silently), and the plan's top 50 are there.
func TestEveryShippedDescParses(t *testing.T) {
	entries, err := descsFS.ReadDir("descs")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != len(builtinDescs) {
		for _, e := range entries {
			b, _ := descsFS.ReadFile("descs/" + e.Name())
			if _, err := ParseLibDesc(string(b)); err != nil {
				t.Errorf("%s: %v", e.Name(), err)
			}
		}
		t.Fatalf("%d description files, %d parsed", len(entries), len(builtinDescs))
	}
	if len(builtinDescs) < 50 {
		t.Errorf("%d shipped descriptions; the plan's goal is the top 50", len(builtinDescs))
	}
	for _, d := range builtinDescs {
		for name, cl := range d.Classes {
			if cl.Extends != "" && d.Classes[cl.Extends] == nil {
				t.Errorf("%s: %s extends %s, which the description doesn't have", d.Name, name, cl.Extends)
			}
		}
	}
}

// top50Smoke uses every library added for the top 50 with its real calls, base-class methods included.
var top50Smoke = `board esp32
use "HX711"
use "Adafruit BMP280 Library"
use "BH1750"
use "Adafruit PWM Servo Driver Library"
use "ezButton"
use "Adafruit ADS1X15"
use "Adafruit INA219"
use "Adafruit MLX90614 Library"
use "Adafruit SHT31 Library"
use "Adafruit Si7021 Library"
use "Adafruit MCP23017 Arduino Library"
use "VL53L1X"
use "LoRa"
use "QMC5883LCompass"
use "WiFiManager"
scale = HX711()
bmp = Adafruit_BMP280()
light = BH1750()
pwm = Adafruit_PWMServoDriver(0x40)
button = ezButton(7)
adc = Adafruit_ADS1115()
power_meter = Adafruit_INA219()
ir = Adafruit_MLX90614()
sht = Adafruit_SHT31()
si = Adafruit_Si7021()
io = Adafruit_MCP23X17()
tof = VL53L1X()
compass = QMC5883LCompass()
wm = WiFiManager()
at start:
    scale.begin(4, 5)
    scale.set_scale(420.0)
    scale.tare()
    print scale.get_units(5)
    if bmp.begin(BMP280_ADDRESS_ALT):
        print bmp.read_pressure()
    light.begin()
    print light.read_light_level()
    pwm.begin()
    pwm.set_pwm_freq(50.0)
    pwm.set_pwm(0, 0, 300)
    button.set_debounce_time(50)
    adc.begin()
    adc.set_gain(GAIN_ONE)
    print adc.compute_volts(adc.read_adc_single_ended(0))
    power_meter.begin()
    print power_meter.get_current_mA()
    ir.begin()
    print ir.read_object_temp_c()
    sht.begin(SHT31_DEFAULT_ADDR)
    print sht.read_humidity()
    si.begin()
    print si.read_temperature()
    io.begin_I2C()
    io.pin_mode(0, OUTPUT)
    io.digital_write(0, HIGH)
    print io.read_gpioab()
    tof.set_timeout(500)
    if tof.init():
        tof.start_continuous(50)
    print tof.read()
    if LoRa.begin(868000000) == 1:
        LoRa.begin_packet()
        LoRa.print("hi")
        LoRa.end_packet()
    compass.init()
    compass.read()
    print compass.get_azimuth()
    if wm.auto_connect("Mujard-Setup"):
        print wm.get_wifi_ssid()
every 10 ms:
    button.loop()
    if button.is_pressed():
        print "pressed"
`

func TestTop50Smoke(t *testing.T) {
	sk, diags := Compile(top50Smoke, "smoke.muj")
	if e := FirstError(diags); e != nil || sk == nil {
		t.Fatalf("the top-50 smoke program should compile: %v", diags)
	}
	if hasDiag(diags, "W611") {
		t.Errorf("every library here is described: %v", diags)
	}
	cpp := sk.Files["mujard_program.cpp"]
	for _, want := range []string{"static Adafruit_ADS1115 adc;", "adc.setGain(GAIN_ONE)", "io.begin_I2C()", "io.readGPIOAB()", "LoRa.beginPacket()", "wm.getWiFiSSID()",
		// A library's own names that Arduino also has (pinMode, OUTPUT, HIGH, loop) are never renamed.
		"io.pinMode(0, OUTPUT)", "io.digitalWrite(0, HIGH)", "button.loop()"} {
		if !strings.Contains(cpp, want) {
			t.Errorf("generated C++ lacks %q", want)
		}
	}
	bad := map[string]string{
		"an abstract base isn't a class to make": "use \"Adafruit ADS1X15\"\nadc = Adafruit_ADS1X15()\n",
		"a gain is not a plain number":           "use \"Adafruit ADS1X15\"\nadc = Adafruit_ADS1115()\nat start:\n    adc.set_gain(2)\n",
		"WiFiManager is ESP only":                "use \"WiFiManager\"\nwm = WiFiManager()\n",
	}
	for name, body := range bad {
		board := "esp32"
		if strings.Contains(body, "WiFiManager") {
			board = "uno"
		}
		if _, d := CheckSourceModules("board "+board+"\n"+body, nil); FirstError(d) == nil {
			t.Errorf("%s: should not check clean", name)
		}
	}
}

// TestLibraryBuilds compiles the library programs above with arduino-cli against the real libraries - the proof
// that a description matches the library's C++, not just itself. Gated like TestBuildMatrix:
//
//	$env:MUJARD_ARDUINO_CLI="C:\arduino-toolchain\bin\arduino-cli.exe"; $env:MUJARD_ARDUINO_CONFIG="C:\arduino-toolchain\arduino-cli.yaml"
//	go test ./lang -run TestLibraryBuilds -v -timeout 90m   (a cold ESP32 build of 15 libraries took over 30 min here)
func TestLibraryBuilds(t *testing.T) {
	cli := os.Getenv("MUJARD_ARDUINO_CLI")
	if cli == "" {
		t.Skip("set MUJARD_ARDUINO_CLI to build the library programs with arduino-cli")
	}
	cfg := os.Getenv("MUJARD_ARDUINO_CONFIG")
	for name, src := range map[string]string{"u8g2 on uno": u8g2Program, "top-50 libraries on esp32": top50Smoke} {
		sk, diags := Compile(src, "lib.muj")
		if e := FirstError(diags); e != nil {
			t.Errorf("%s: %s", name, e)
			continue
		}
		dir := filepath.Join(t.TempDir(), "sketch")
		os.MkdirAll(dir, 0o755)
		for f, content := range sk.Files {
			os.WriteFile(filepath.Join(dir, f), []byte(content), 0o644)
		}
		writeTestSecretsHeader(dir, sk.Secrets)
		args := []string{"compile", "--fqbn", sk.FQBN, "--warnings", "default", dir}
		if cfg != "" {
			args = append([]string{"--config-file", cfg}, args...)
		}
		out, err := exec.Command(cli, args...).CombinedOutput()
		if err != nil {
			t.Errorf("%s (%s): build failed: %v\n%s", name, sk.FQBN, err, tail(string(out), 40))
			continue
		}
		t.Logf("%s (%s): %s", name, sk.FQBN, lastLine(string(out), "Sketch uses"))
	}
}

func TestBuiltinDescArch(t *testing.T) {
	// WiFi is esp32-only in the shipped descriptions.
	src := "board uno\nuse \"WiFi\"\nat start:\n    WiFi.begin(\"net\", \"pw\")\n"
	_, diags := CheckSourceModules(src, nil)
	if !hasDiag(diags, "B412") {
		t.Error("WiFi on an AVR board should be B412")
	}
	src = "board esp32\nuse \"WiFi\"\nat start:\n    WiFi.begin(\"net\", \"pw\")\n"
	_, diags = CheckSourceModules(src, nil)
	if hasDiag(diags, "B412") {
		t.Error("WiFi on esp32 is fine")
	}
	if e := FirstError(diags); e != nil {
		t.Errorf("WiFi.begin should check clean: %v", e)
	}
}

func TestLibDescsEnvDir(t *testing.T) {
	dir := t.TempDir()
	cache := t.TempDir()
	if err := os.WriteFile(filepath.Join(cache, "neo.mlib"), []byte(neopixelMLib), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("MUJARD_LIB_DESCS", cache)
	_, diags := CheckSourceIn(neoProgram, dir)
	if hasDiag(diags, "W611") {
		t.Error("MUJARD_LIB_DESCS descriptions should apply")
	}
	if e := FirstError(diags); e != nil {
		t.Errorf("env-dir description calls should check clean: %v", e)
	}
}

// -- the description generator --

func TestDescribeLibrary(t *testing.T) {
	dir := t.TempDir()
	props := "name=Widget Sensor\nversion=2.1.0\narchitectures=avr,esp32\nincludes=Widget.h\n"
	if err := os.WriteFile(filepath.Join(dir, "library.properties"), []byte(props), 0o644); err != nil {
		t.Fatal(err)
	}
	header := `// Widget Sensor
#ifndef WIDGET_H
#define WIDGET_H
#define WIDGET_ADDR 0x40
#define WIDGET_NAME "widget"
class Widget {
public:
    Widget(uint8_t addr);
    void begin();
    float readTemperature(bool fahrenheit = false);
    uint8_t status() const;
private:
    int secret() { return 1; }
    uint8_t _addr;
};
int widgetVersion();
`
	if err := os.WriteFile(filepath.Join(dir, "Widget.h"), []byte(header), 0o644); err != nil {
		t.Fatal(err)
	}
	d, err := DescribeLibrary(dir)
	if err != nil {
		t.Fatalf("DescribeLibrary: %v", err)
	}
	if d.Name != "Widget Sensor" || d.Header != "Widget.h" {
		t.Errorf("name/header: %+v", d)
	}
	if len(d.Architectures) != 2 || d.Architectures[0] != "avr" {
		t.Errorf("architectures: %v", d.Architectures)
	}
	cl := d.Classes["Widget"]
	if cl == nil {
		t.Fatalf("class Widget: %+v", d.Classes)
	}
	if len(cl.Ctors) != 1 || len(cl.Ctors[0].Params) != 1 {
		t.Errorf("ctor: %+v", cl.Ctors)
	}
	if m := cl.Methods["readTemperature"]; m == nil || m.Result != "float" || len(m.Params) != 1 || !m.Params[0].Optional {
		t.Errorf("readTemperature: %+v", m)
	}
	if cl.Methods["secret"] != nil {
		t.Error("a private method must not be described")
	}
	if cl.Fields["_addr"] != "" {
		t.Error("a private field must not be described")
	}
	if f := d.Functions["widgetVersion"]; f == nil || f.Result != "int" {
		t.Errorf("widgetVersion: %+v", f)
	}
	if d.Constants["WIDGET_ADDR"] != "int" || d.Constants["WIDGET_NAME"] != "text" {
		t.Errorf("constants: %+v", d.Constants)
	}
	// The generated description must round-trip through the checker.
	b, err := DescribeLibraryJSON(dir)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := ParseLibDesc(b); err != nil {
		t.Fatalf("generated desc does not parse: %v", err)
	}
	src := `board uno
use "Widget Sensor"
w = Widget(0x40)
at start:
    w.begin()
    print w.read_temperature()
`
	_, diags := CheckSourceModules(src, map[string]string{"widget.mlib": b})
	if e := FirstError(diags); e != nil {
		t.Errorf("calls against a generated desc should check clean: %v", e)
	}
	if hasDiag(diags, "W611") {
		t.Error("a generated desc should not warn W611")
	}
}
