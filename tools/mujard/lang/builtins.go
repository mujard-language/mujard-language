package lang

// The built-in names the checker knows (docs/lang/README.md). Anything not described here that comes from a C++ library
// is TAny: called and used freely, checked by the C++ compiler instead.

type param struct {
	name     string
	t        *Type // nil: any value
	optional bool
}

type fnSig struct {
	params   []param // nil: arguments aren't checked
	variadic bool    // more arguments like the last one may follow
	result   *Type
	resultOf func(args []*Type) *Type // a result that depends on the arguments (abs, clamp, scale, ...)
	blocking bool                     // waits for an answer: only in tasks (and tests)
	testOnly bool
	network  bool // needs Wi-Fi on the board
}

type member struct {
	field      *Type  // a property
	method     *fnSig // a method
	event      bool   // usable after 'on'
	eventArg   *Type  // the event takes an argument: held(2 s)
	eventValue *Type  // what 'on ... as x' binds
}

type class map[string]*member

func m(params []param, result *Type) *member {
	return &member{method: &fnSig{params: params, result: result}}
}
func field(t *Type) *member { return &member{field: t} }
func event() *member        { return &member{event: true} }
func p(name string, t *Type) param {
	return param{name: name, t: t}
}
func optp(name string, t *Type) param { return param{name: name, t: t, optional: true} }

var (
	tTime    = qty("time")
	tFreq    = qty("frequency")
	tLength  = qty("length")
	tVoltage = qty("voltage")
	tAngle   = qty("angle")
	tRatio   = qty("ratio")
	tMemory  = qty("memory")
)

// sameAsFirst returns the (non-optional) type of the first argument.
func sameAsFirst(args []*Type) *Type {
	if len(args) == 0 {
		return tAny
	}
	return nonOpt(args[0])
}

func elemOf(t *Type) *Type {
	if t != nil && (t.K == TRange || t.K == TList) && t.Elem != nil {
		return t.Elem
	}
	return tAny
}

// classes of pin devices and built-in / library objects.
var classes = map[string]class{
	"output": {
		"on": m(nil, tVoid), "off": m(nil, tVoid), "toggle": m(nil, tVoid), "stop": m(nil, tVoid),
		"write":   m([]param{p("value", tBool)}, tVoid),
		"read":    m(nil, tBool),
		"blink":   m([]param{p("rate", nil), optp("for", tTime), optp("on", tTime)}, tVoid),
		"pattern": m([]param{p("pattern", tText), optp("step", tTime), optp("repeat", tInt)}, tVoid),
		"pulse":   m([]param{p("length", tTime)}, tVoid),
		"tone":    m([]param{p("pitch", tFreq), optp("length", tTime)}, tVoid),
		"melody":  m([]param{p("notes", tText), optp("tempo", tInt)}, tVoid),
		"pretend": {method: &fnSig{params: []param{p("value", nil)}, result: tVoid, testOnly: true}},
	},
	"input": {
		"read":    m(nil, tBool),
		"pressed": event(), "released": event(), "clicked": event(), "double_clicked": event(), "changed": event(),
		"rising": event(), "falling": event(), // the voltage on the wire, not the active level (spec 6.5)
		"held":          {event: true, eventArg: tTime},
		"debounce":      m([]param{p("time", tTime)}, tVoid),
		"measure_pulse": m([]param{optp("level", tLevel), optp("timeout", tTime)}, opt(tTime)),
		"pretend":       {method: &fnSig{params: []param{p("value", nil)}, result: tVoid, testOnly: true}},
	},
	"pwm": {
		"on": m(nil, tVoid), "off": m(nil, tVoid), "stop": m(nil, tVoid),
		"write":     m([]param{p("level", tRatio)}, tVoid),
		"fade":      m([]param{p("to", tRatio), p("over", tTime)}, tVoid),
		"frequency": m([]param{p("frequency", tFreq)}, tVoid),
		"pretend":   {method: &fnSig{params: []param{p("value", nil)}, result: tVoid, testOnly: true}},
	},
	"analog": {
		"value": field(tInt), "volts": field(tVoltage), "percent": field(tRatio),
		"reference": m([]param{p("voltage", tVoltage)}, tVoid),
		"pretend":   {method: &fnSig{params: []param{p("value", nil)}, result: tVoid, testOnly: true}},
	},
	"group": {
		"write":  m([]param{p("bits", tInt)}, tVoid),
		"length": field(tInt),
	},
	"board": {
		"name": field(tText), "chip": field(tText), "id": field(tText), "uptime": field(tTime),
		"cores": field(tInt), "dma": field(tBool), "timers": field(tInt), "rtos": field(tText), "ethernet": field(tBool),
		"free_memory": field(tMemory), "voltage": field(opt(tVoltage)), "reset_reason": field(tResetReason),
		"overruns": field(tInt), "dropped_events": field(tInt),
		"restart": m(nil, tVoid),
		"halt":    m([]param{p("reason", tText)}, tVoid),
	},
	"log": {
		"info": m([]param{p("message", nil)}, tVoid), "warn": m([]param{p("message", nil)}, tVoid), "error": m([]param{p("message", nil)}, tVoid),
	},
	"time_of_day": {"hour": field(tInt), "minute": field(tInt), "second": field(tInt)},
	"date":        {"year": field(tInt), "month": field(tInt), "day": field(tInt), "weekday": field(tInt)},
	"clock": {
		"time": field(opt(object("time_of_day"))), "date": field(opt(object("date"))),
		"sync":   {method: &fnSig{params: []param{optp("timezone", tText)}, result: tVoid, network: true}},
		"source": m([]param{p("source", nil)}, tVoid),
	},
	"wifi": {
		"connect": {method: &fnSig{params: []param{p("network", nil)}, result: tVoid, network: true}},
		// serve("name"): the app-protocol socket on _mujard._tcp / TCP 8266 (docs/lang/MUJARD_APP_PROTOCOL.md);
		// serve("name", "1234") also asks the phone for that PIN before it may drive controls (pairing).
		"serve":     {method: &fnSig{params: []param{p("name", tText), optp("pin", tText)}, result: tVoid, network: true}},
		"connected": {field: tBool},
		// client: a phone is on the serve() socket right now
		"client": {field: tBool},
		"joined": event(), "lost": event(),
	},
	"serial": {
		"line":      {event: true, eventValue: tText},
		"read_line": {method: &fnSig{result: opt(tText), blocking: true}},
		"pretend":   {method: &fnSig{params: []param{p("line", tText)}, result: tVoid, testOnly: true}},
	},
	"random": {
		"seed": m([]param{optp("seed", tInt)}, tVoid),
		"pick": {method: &fnSig{params: []param{p("list", nil)}, resultOf: func(a []*Type) *Type {
			if len(a) > 0 {
				return elemOf(a[0])
			}
			return tAny
		}}},
	},
	"stopwatch": {
		"start": m(nil, tVoid), "stop": m(nil, tVoid), "reset": m(nil, tVoid),
		"elapsed": field(tTime), "running": field(tBool),
	},
	"block": { // a named every/after/when block, or a task
		"start": m(nil, tVoid), "stop": m(nil, tVoid), "restart": m(nil, tVoid), "running": field(tBool),
	},
	"battery": {"percent": field(tRatio), "volts": field(tVoltage),
		"pretend": {method: &fnSig{params: []param{p("value", nil)}, result: tVoid, testOnly: true}}},
	"servo": {
		"move":    m([]param{p("angle", tAngle), optp("over", tTime)}, tVoid),
		"angle":   field(tAngle),
		"pretend": {method: &fnSig{params: []param{p("value", nil)}, result: tVoid, testOnly: true}},
	},
	"ultrasonic": {
		"distance": m(nil, opt(tLength)),
		"pretend":  {method: &fnSig{params: []param{p("distance", tLength)}, result: tVoid, testOnly: true}},
	},
	"math": {
		"sqrt": m([]param{p("x", tFloat)}, tFloat), "pow": m([]param{p("x", tFloat), p("y", tFloat)}, tFloat),
		"log": m([]param{p("x", tFloat)}, tFloat),
		"sin": m([]param{p("angle", tAngle)}, tFloat), "cos": m([]param{p("angle", tAngle)}, tFloat), "tan": m([]param{p("angle", tAngle)}, tFloat),
		"atan2": m([]param{p("y", tFloat), p("x", tFloat)}, tAngle),
		"PI":    field(tFloat), "E": field(tFloat),
	},
	"mqtt": {
		"connect":   {method: &fnSig{params: []param{p("broker", tText), optp("user", nil), optp("password", nil)}, result: tVoid, network: true}},
		"publish":   {method: &fnSig{params: []param{p("topic", tText), p("value", nil)}, result: tVoid, network: true}},
		"subscribe": {method: &fnSig{params: []param{p("topic", tText)}, result: tVoid, network: true}},
		"message":   {event: true, eventValue: object("mqtt_message")},
	},
	"mqtt_message": {"text": field(tText), "topic": field(tText)},
	"uart": {
		"write":     m([]param{p("text", tText)}, tVoid),
		"read_line": {method: &fnSig{result: opt(tText), blocking: true}},
		"line":      {event: true, eventValue: tText},
		"pretend":   {method: &fnSig{params: []param{p("line", tText)}, result: tVoid, testOnly: true}},
	},
	"can": {
		"send":    m([]param{p("id", tInt), p("data", listOf(tInt))}, tBool),
		"read":    m(nil, opt(object("can_message"))),
		"message": {event: true, eventValue: object("can_message")},
		// tests only: a frame the fake bus "received"
		"pretend": {method: &fnSig{params: []param{p("id", tInt), p("data", listOf(tInt))}, result: tVoid, testOnly: true}},
	},
	"can_message": {"id": field(tInt), "data": field(listCap(tInt, 8)), "rtr": field(tBool)},
	"spi":         {"device": m([]param{p("cs", nil), optp("speed", tFreq), optp("mode", tInt)}, object("spi_device"))},
	"spi_device":  {"transfer": m([]param{p("data", listOf(tInt))}, listCap(tInt, 32))},
	"rtc":         {"set": m([]param{p("year", tInt), p("month", tInt), p("day", tInt), p("hour", tInt), p("minute", tInt), optp("second", tInt)}, tVoid)},
	"i2c": {
		"scan":  m(nil, listCap(tInt, 16)),
		"write": m([]param{p("address", tInt), p("data", listOf(tInt))}, tVoid),
		"read":  m([]param{p("address", tInt), p("count", tInt)}, opt(listCap(tInt, 32))),
		// tests only: a device on the fake bus answers with reply's bytes; none takes it off the bus (spec 12)
		"pretend": {method: &fnSig{params: []param{p("address", tInt), p("reply", nil)}, result: tVoid, testOnly: true}},
	},
	"shift":          {"out_register": m([]param{p("data", nil), p("clock", nil), p("latch", nil)}, object("shift_register"))},
	"shift_register": {"write": m([]param{p("bits", tInt)}, tVoid)},
	"http": {
		"get":  {method: &fnSig{params: []param{p("url", tText)}, result: object("http_reply"), blocking: true, network: true}},
		"post": {method: &fnSig{params: []param{p("url", tText), p("data", nil)}, result: object("http_reply"), blocking: true, network: true}},
	},
	"http_reply": {"ok": field(tBool), "status": field(tInt), "text": field(tText), "json": field(object("json"))},
	// json: a value read out of reply.json[...]. Index with text or int; .ok is false for a missing key;
	// to_number() gives a number or none; 'as unit' makes a quantity straight away.
	"json": {
		"ok":        field(tBool),
		"to_number": m(nil, opt(tFloat)),
	},
	"store": {
		"save": m([]param{p("name", tText), p("value", nil)}, tVoid),
		// durability observability: store.save is async (writes go to RAM, commit on the next tick/flush).
		// pending = a store.save write accepted but not yet committed, or a saved var whose live value
		// isn't in storage yet; failed = commits that didn't land (shared with saved vars);
		// committed/last_commit = the most recent commit's result.
		"pending":     field(tBool),
		"failed":      field(tInt),
		"committed":   field(tBool),
		"last_commit": field(tBool),
		// the default gives the type; without one the result is T?, inferred from its declared target
		"load": {method: &fnSig{params: []param{p("name", tText), optp("default", nil)}, resultOf: func(a []*Type) *Type {
			if len(a) > 1 && a[1] != nil {
				return nonOpt(a[1])
			}
			return opt(tAny)
		}}},
	},
	// BLE peripheral: the board advertises and exchanges text lines with a phone or another board (spec 12,
	// mujard.ble). Nordic-UART shape: send/read_line/connected plus the 'on ble.line' event.
	"ble": {
		"start":     m([]param{p("name", tText)}, tVoid),
		"send":      m([]param{p("text", tText)}, tBool),
		"read_line": m(nil, opt(tText)),
		"connected": field(tBool),
		"line":      {event: true, eventValue: tText},
		// tests only: a line the link "received"
		"pretend": {method: &fnSig{params: []param{p("line", tText)}, result: tVoid, testOnly: true}},
	},
	// A wired Ethernet link (the classic ESP32's EMAC, mujard.ethernet): ethernet.start() brings it up,
	// connected/ip read its state, 'on ethernet.joined'/'lost' fire on link changes, pretend fakes it in tests.
	"ethernet": {
		"start":     m(nil, tVoid),
		"connected": field(tBool),
		"ip":        field(tText),
		"joined":    {event: true},
		"lost":      {event: true},
		"pretend":   {method: &fnSig{params: []param{p("up", tBool)}, result: tVoid, testOnly: true}},
	},
	// USB device: the board acts as a keyboard+mouse to whatever it's plugged into (spec 12, mujard.usb).
	"usb": {
		"type":       m([]param{p("text", tText)}, tBool),
		"key":        m([]param{p("key", tText)}, tBool),
		"mouse_move": m([]param{p("x", tInt), p("y", tInt)}, tBool),
		"click":      m(nil, tBool),
		"connected":  field(tBool),
		// tests only: the fake host is listening
		"pretend": {method: &fnSig{params: []param{p("up", tBool)}, result: tVoid, testOnly: true}},
	},
	// A hardware timer (mujard.timer): t.every/t.freq/t.once set a precise trigger - a real timer interrupt,
	// not the cooperative 'every'. 'on t.tick' runs on the scheduler; 'on t.tick fast' runs in the ISR.
	"timer": {
		"every":   m([]param{p("period", tTime)}, tVoid),
		"freq":    m([]param{p("hz", tFreq)}, tVoid),
		"once":    m([]param{p("period", tTime)}, tVoid),
		"start":   m(nil, tVoid),
		"stop":    m(nil, tVoid),
		"elapsed": m(nil, tTime),
		"tick":    {event: true},
		"pretend": {method: &fnSig{params: nil, result: tVoid, testOnly: true}},
	},
	"power": {
		"sleep":      m([]param{p("time", tTime), optp("wake_on", nil)}, tVoid),
		"deep_sleep": m([]param{p("time", nil)}, tVoid),
	},
	// A bounded queue (mujard.queue): q = Queue(text, 8) passes values between blocks - and is the safe way to hand
	// data between cores, where a plain global races (W620). push is false and counts a drop when full; pop gives
	// none when empty; 'when q.available:' waits for items. push/pop take the declared element type.
	"queue": {
		"push":      m([]param{p("value", nil)}, tBool),
		"pop":       m(nil, tAny),
		"available": field(tBool),
		"full":      field(tBool),
		"n":         field(tInt),
	},
	// regs_reg: the member methods every register object (regs.PORTB and friends) carries. The register names
	// themselves aren't here - they come from the chip's generated table (regs.json) via 'use mujard.regs'.
	"regs_reg": {
		"read":  m(nil, tInt),
		"write": m([]param{p("value", tInt)}, tVoid),
		"set":   m([]param{p("bits", tInt)}, tVoid),
		"clear": m([]param{p("bits", tInt)}, tVoid),
	},
}

// textMembers, numberMembers and listMembers are the methods of plain values.
var textMembers = class{
	"length": field(tInt), "contains": m([]param{p("part", tText)}, tBool),
	"starts_with": m([]param{p("part", tText)}, tBool), "ends_with": m([]param{p("part", tText)}, tBool),
	"split": m([]param{p("separator", tText)}, listOf(tText)), "to_number": m(nil, opt(tFloat)),
	"upper": m(nil, tText), "lower": m(nil, tText),
}
var intMembers = class{
	"bit": m([]param{p("n", tInt)}, tBool), "set_bit": m([]param{p("n", tInt)}, tInt),
	"clear_bit": m([]param{p("n", tInt)}, tInt), "toggle_bit": m([]param{p("n", tInt)}, tInt),
	"bytes": field(listOf(tInt)),
}
var listMembers = class{
	"length": field(tInt), "add": m([]param{p("item", nil)}, tVoid), "remove": m([]param{p("index", tInt)}, tVoid),
	"clear": m(nil, tVoid),
}

// global is a built-in name: a function, an object, or both (random, serial).
type global struct {
	fn    *fnSig
	class string // object members
	value *Type  // a constant
	lib   string // only after 'use mujard.<lib>'
	emit  string // a described library name's real C++ spelling (for snake_case aliases)
}

var globals = map[string]*global{
	"notify":     {fn: &fnSig{params: []param{p("message", nil), optp("urgent", tBool)}, result: tVoid}},
	"show":       {fn: &fnSig{params: []param{p("label", tText), p("value", nil), optp("as", nil)}, result: tVoid}},
	"plot":       {fn: &fnSig{params: []param{p("value", nil)}, variadic: true, result: tVoid}},
	"ask":        {fn: &fnSig{params: []param{p("question", tText), optp("timeout", tTime)}, result: opt(tBool), blocking: true}},
	"secret":     {fn: &fnSig{params: []param{p("name", tText)}, result: object("secret")}},
	"now":        {fn: &fnSig{params: []param{}, result: tTime}},
	"since":      {fn: &fnSig{params: []param{p("time", tTime)}, result: tTime}},
	"stopwatch":  {fn: &fnSig{params: []param{}, result: object("stopwatch")}},
	"abs":        {fn: &fnSig{params: []param{p("x", nil)}, resultOf: sameAsFirst}},
	"min":        {fn: &fnSig{params: []param{p("a", nil), p("b", nil)}, variadic: true, resultOf: sameAsFirst}},
	"max":        {fn: &fnSig{params: []param{p("a", nil), p("b", nil)}, variadic: true, resultOf: sameAsFirst}},
	"clamp":      {fn: &fnSig{params: []param{p("x", nil), p("range", nil)}, resultOf: sameAsFirst}},
	"round":      {fn: &fnSig{params: []param{p("x", nil), optp("decimals", tInt)}, resultOf: sameAsFirst}},
	"smooth":     {fn: &fnSig{params: []param{p("value", nil), p("count", tInt)}, resultOf: sameAsFirst}},
	"peak":       {fn: &fnSig{params: []param{p("value", nil), p("window", tTime)}, resultOf: sameAsFirst}},
	"between":    {fn: &fnSig{params: []param{p("x", nil), p("range", nil)}, result: tBool}},
	"changed":    {fn: &fnSig{params: []param{p("value", nil)}, result: tBool}},
	"rising":     {fn: &fnSig{params: []param{p("value", tBool)}, result: tBool}},
	"falling":    {fn: &fnSig{params: []param{p("value", tBool)}, result: tBool}},
	"hysteresis": {fn: &fnSig{params: []param{p("value", nil), p("on", nil), p("off", nil)}, result: tBool}},
	"scale":      {fn: &fnSig{params: []param{p("x", nil), p("from", nil), p("to", nil)}, resultOf: func(a []*Type) *Type { return elemOfArg(a, 2) }}},
	"random":     {fn: &fnSig{params: []param{p("range", nil)}, resultOf: func(a []*Type) *Type { return elemOfArg(a, 0) }}, class: "random"},
	"watchdog":   {fn: &fnSig{params: []param{p("time", tTime)}, result: tVoid}},
	"battery_on": {fn: &fnSig{params: []param{p("pin", nil), optp("divider", nil), optp("type", tBatteryType)}, result: object("battery")}},
	"serial":     {fn: &fnSig{params: []param{p("speed", qty("baud"))}, result: tVoid}, class: "serial"},
	"advance":    {fn: &fnSig{params: []param{p("time", tTime)}, result: tVoid, testOnly: true}},
	"answer":     {fn: &fnSig{params: []param{p("value", nil)}, result: tVoid, testOnly: true}},
	"notified":   {fn: &fnSig{params: []param{p("text", tText)}, result: tBool, testOnly: true}},
	"printed":    {fn: &fnSig{params: []param{p("text", tText)}, result: tBool, testOnly: true}},
	"stopped":    {fn: &fnSig{params: []param{}, result: tBool, testOnly: true}},
	"halted":     {fn: &fnSig{params: []param{}, result: tBool, testOnly: true}},
	"power_cut":  {fn: &fnSig{params: []param{}, result: tVoid, testOnly: true}}, // the board loses power and starts again (spec 12)
	"brown_out":  {fn: &fnSig{params: []param{}, result: tVoid, testOnly: true}}, // the supply dips enough for a brown-out reset (spec 12)
	"board":      {class: "board"},
	"log":        {class: "log"},
	"clock":      {class: "clock"},
	"wifi":       {class: "wifi"},
	"store":      {class: "store"},
	"power":      {class: "power"},
	"ble":        {class: "ble", lib: "ble"},
	"ethernet":   {class: "ethernet", lib: "ethernet"},
	"usb":        {class: "usb", lib: "usb"},
	"Timer":      {fn: &fnSig{params: nil, result: object("timer")}, lib: "timer"},
	"Queue":      {fn: &fnSig{params: []param{p("of", nil), p("capacity", nil)}, result: object("queue")}, lib: "queue"},
	"http":       {class: "http"},
	"rtc":        {class: "rtc", lib: "rtc"},
	"math":       {class: "math", lib: "math"},
	"mqtt":       {class: "mqtt", lib: "mqtt"},
	"i2c":        {class: "i2c", lib: "i2c"},
	"shift":      {class: "shift", lib: "shift"},
	"Servo":      {fn: &fnSig{params: []param{p("pin", nil)}, result: object("servo")}, lib: "servo"},
	"Ultrasonic": {fn: &fnSig{params: []param{p("trig", nil), p("echo", nil)}, result: object("ultrasonic")}, lib: "ultrasonic"},
	"Uart":       {fn: &fnSig{params: []param{p("rx", nil), p("tx", nil), optp("speed", qty("baud"))}, result: object("uart")}, lib: "uart"},
	"Can":        {fn: &fnSig{params: []param{p("tx", nil), p("rx", nil), optp("rate", qty("baud"))}, result: object("can")}, lib: "can"},
	"spi":        {class: "spi", lib: "spi"},
	"regs":       {class: "regs", lib: "regs"},
	"timezone":   {value: tAny},
	"gauge":      {fn: &fnSig{params: []param{p("range", nil)}, result: tAny}},
	"forever":    {value: tTime},
}

// Built-in enums (spec 4.6). Each is a real enum type: board.reset_reason == 42 is C303. A member works bare
// (crash, high, lipo_1s) or in full (ResetReason.crash); a bare name that is something else already
// (the watchdog function) keeps that meaning, so only the full form reaches that member.
var (
	tResetReason = &Type{K: TEnum, Name: "ResetReason", Dim: "builtin"}
	tLevel       = &Type{K: TEnum, Name: "Level", Dim: "builtin"}
	tBatteryType = &Type{K: TEnum, Name: "BatteryType", Dim: "builtin"}
)

// isBuiltinEnum tells a built-in enum's value from one of a program enum that happens to share its name.
func isBuiltinEnum(t *Type) bool { return t != nil && t.K == TEnum && t.Dim == "builtin" }

var builtinEnums = map[string]struct {
	t       *Type
	members []string
}{
	"ResetReason": {tResetReason, []string{"power_on", "reset_button", "watchdog", "brownout", "software", "wake", "crash"}},
	"Level":       {tLevel, []string{"high", "low"}},
	"BatteryType": {tBatteryType, []string{"lipo_1s", "aa_x3", "aa_x4"}},
}

func init() {
	for name, en := range builtinEnums {
		cl := class{}
		for _, v := range en.members {
			cl[v] = field(en.t)
			if globals[v] == nil {
				globals[v] = &global{value: en.t}
			}
		}
		classes[name] = cl
		globals[name] = &global{class: name}
	}
}

func elemOfArg(a []*Type, i int) *Type {
	if i < len(a) {
		if a[i].K == TBool {
			return tBool
		}
		return nonOpt(elemOf(a[i]))
	}
	return tAny
}
