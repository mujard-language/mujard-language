# Mujard built-ins: how to use them

Mujard is a simple language for programming boards. **Bare code. Every board.** Light and fast hardware programming.
It is compiled to C++ and then by the same compilers Arduino uses, so it runs at the same speed, on every board the
toolchain supports - built locally for free with `mujard-build` or `arduino-cli`, or through the optional hosted
build service. Source files end in `.muj` (`.mjd` is accepted too).

> **Status: implemented, pre-1.0.** This page is the reference the Mujard compiler and libraries are built to.
> The compiler (lexer, parser, checker, C++ generator), the `mujard_rt` runtime and the CLI are implemented, and every
> board in `toolchain/boards` builds the sample programs (see `tools/mujard/lang/capabilities.json`); real-board runs
> are verified for AVR. Items marked **[app]** still need the phone/web/server features of the platform. The language rules are in `MUJARD_LANGUAGE_SPEC.md`, which wins if the two disagree;
> exact parameter and result types are in its Appendix A.

This page explains all 100 built-in names: what each one does, how to write it, what it gives back, and what you would
write in Arduino C++ to do the same.

## Contents

1. [Before you start](#before-you-start): declaring things, units, events, `none`
2. [Messages and the app](#1-messages-and-the-app) (1-10)
3. [The board itself](#2-the-board-itself) (11-18)
4. [Time](#3-time) (19-24)
5. [Maths](#4-maths) (25-32)
6. [Smart signals](#5-smart-signals) (33-39)
7. [Random](#6-random) (40-42)
8. [Digital pins](#7-digital-pins) (43-50)
9. [Analog and PWM](#8-analog-and-pwm) (51-57)
10. [Buttons and inputs](#9-buttons-and-inputs) (58-64)
11. [Sound](#10-sound) (65-67)
12. [Saving data](#11-saving-data) (68-70)
13. [Communication](#12-communication) (71-76)
14. [Network](#13-network) (77-81)
15. [Power](#14-power) (82-85)
16. [Safety](#15-safety) (86-89)
17. [Testing without hardware](#16-testing-without-hardware) (90-93)
18. [Text](#17-text) (94-97)
19. [Bits](#18-bits) (98-100)
20. [More links and more cores](#19-more-links-and-more-cores) (101-114): CAN, BLE, Ethernet, USB, second core, hardware timer, Queue, registers
21. [Mujard compared with Arduino](#mujard-compared-with-arduino)
22. [Appendix: keywords](#appendix-keywords)

---

## Before you start

### A complete program

```mujard
board uno                              # 1. which board

led    = pin 13 as output              # 2. name your pins
button = pin 2  as input pullup

at start:                              # 3. what happens once, at power-up
    serial 9600 baud
    print "Ready"

every 500 ms:                          # 4. what happens on a timer
    led.toggle()

on button.pressed:                     # 5. what happens on an event
    print "Pressed!"
```

There is no `setup()` or `loop()`. Blocks such as `every`, `on`, `when` and `after` all run side by side, and none of them
stops the others.

### Where built-ins live

| Kind | How you write it | Examples |
|---|---|---|
| **Global** (about 10, usable anywhere) | `name(...)` or `name ...` | `print`, `notify`, `show`, `plot`, `now()`, `random()`, `scale()`, `min()`, `max()`, `clamp()` |
| **On something you declared** | `thing.name` | `led.on()`, `button.pressed`, `light.volts` |
| **On a built-in object** | `object.name` | `board.uptime`, `log.warn`, `clock.time`, `wifi.connect()` |
| **In a library** (add a `use` line) | `use mujard.math` then `math.sqrt()` | `math`, `servo`, `ultrasonic`, `i2c`, `spi`, `shift`, `mqtt` |
| **In another `.muj` file** (add an `import` line) | `import motors` then `motors.forward()` | your own modules |

Typing `led.` in the editor lists everything an LED can do.

`import motors` loads `motors.muj` from the same folder. Its variables, functions, structs, enums and devices are
reached as `motors.name`; its event blocks and tasks don't run in the importing program.

### Units

Numbers that measure something carry a unit. Mujard checks them, so `wait 500` (500 what?) is an error.

| Measures | Units |
|---|---|
| Time | `us` `ms` `s` `min` `h` |
| Frequency | `Hz` `kHz` `MHz` |
| Voltage / current | `mV` `V` / `mA` `A` |
| Length | `mm` `cm` `m` |
| Angle | `deg` |
| Temperature | `C` |
| Share of the full range | `%` |
| Serial speed | `baud` |
| Memory | `bytes` |

Convert with `as`: `print "{distance as cm}"`. Combine units in maths: `t * 343 m/s / 2` gives a length. `%` always means percent; the remainder is the word `mod` (`x mod 3`).

### Events

`on <thing>.<event>:` runs a block each time the event happens. Events are checked many times a second without blocking
anything, and buttons are debounced for you.

### `none`

Something that can fail returns either a value or `none` (for example a pulse measurement that timed out). The compiler
makes you check: `if t == none:`. Arduino usually returns `0` in that case, which is easy to mistake for a real reading.

### Named options

Some built-ins take or give a fixed set of named values: `board.reset_reason` (`power_on`, `reset_button`, `watchdog`,
`brownout`, `software`, `wake`, `crash`), `level:` (`high`, `low`) and the battery `type:` (`lipo_1s`, `aa_x3`, ...).
Write the bare name where the compiler knows which set is meant: `if board.reset_reason == crash:`. If you have a
variable with that name, yours wins; the full form (`ResetReason.crash`) always works. Your own enums behave the same
way (`Mode.idle`, or just `idle` where the type is clear).

---

## 1. Messages and the app

### 1. `print`

Writes a line to the serial monitor (in the app, the website or the PC). Values go inside `{ }`.

```mujard
at start:
    serial 115200 baud
    print "Ready"
    print "Temperature: {t}  Fan: {fan_on}"
    print "{distance as cm}"                  # a value converted to a unit
```

- **Write:** `print <text>`
- **Arduino:** `Serial.print("Temperature: "); Serial.println(t);`
- If `serial` hasn't been started, `print` starts it at 115200 baud.

### 2. `notify` [app]

Sends a push notification to the owner's phone through the Mujard app.

```mujard
when bin_level > 90 %:
    notify "Dustbin almost full ({bin_level})"

on door.released:
    notify "Front door opened", urgent: true
```

- **Write:** `notify <text>` and optionally `urgent: true` (shown even in quiet hours)
- **Needs:** a connection to the app: Wi-Fi with the board linked to an account, or through Mujard.
- **Limit:** at most one notification per 10 s for each line of code, so a loop can't flood the phone.
- **Arduino:** not possible without your own server and push service.

### 3. `show` [app]

Puts a live tile on the project's dashboard in the app.

```mujard
every 2 s:
    show "Temperature", t                       # a number with its unit
    show "Fan", fan_on                          # true/false appears as on/off
    show "Bin level", fill, as: gauge(0..100 %) # a gauge
```

- **Write:** `show <label>, <value>` and optionally `as: number | gauge(range) | light`
- The tile updates whenever `show` runs; the app shows the last value and when it arrived.
- **Arduino:** not possible.

### 4. `plot`

Draws live graphs in the app's plotter. Each value becomes a line, named after its variable.

```mujard
every 100 ms:
    plot temperature, humidity
```

- **Write:** `plot <value>, <value>, ...`
- Works over USB with no setup; over Wi-Fi when the board is linked to the app.
- **Arduino:** write it yourself: `Serial.print("temperature:"); Serial.print(t); Serial.print(",humidity:"); Serial.println(h);`

### 5. `control` [app]

Adds a control to the app's dashboard that changes a variable on the board while it runs.

```mujard
var fan_speed = 40 %
var lights = false

control "Fan speed", fan_speed, 0..100 %       # a slider
control "Lights", lights                       # a switch (for true/false)

every 100 ms:
    fan.write(fan_speed)
    lamp.write(lights)
```

- **Write:** `control <label>, <variable>[, <range>]`; declare at the top, like a pin.
- Values from the app are clamped to the range, so the phone can never set something out of bounds.
- **Arduino:** not possible.

### 6. `ask` [app]

Asks a yes/no question on the phone and waits for the answer (inside a `task`, because it waits).

```mujard
task calibrate:
    var answer = ask "Calibrate the scale now? Remove everything from it first."
    if answer != none and answer:
        scale_sensor.tare()
        notify "Calibrated"
```

- **Write:** `ask <text>` gives `true` (yes), `false` (no) or `none` (no answer within `timeout:`, default 5 min), so
  check for `none` before using the answer. Use it as `var x = ask …`, `x = ask …` or directly in a statement inside
  a `task`; the task waits there until the person answers.
- **Arduino:** not possible.

### 7-9. `log.info`, `log.warn`, `log.error`

Messages with a level. The app's terminal can filter by level; errors are also flagged on the project.

```mujard
at start:
    log.info "Started on {board.name}"

when battery.percent < 20 %:
    log.warn "Battery low: {battery.percent}"

when sensor.value == none:
    log.error "Temperature sensor not answering"
```

- **Write:** `log.info <text>`, `log.warn <text>`, `log.error <text>`
- Each line carries the time and level, e.g. `[12.403 s] WARN Battery low: 18 %`.
- **Arduino:** write it yourself with `Serial.print`.

### 10. `secret` [app]

Reads a password or key from the app's secure store. It is never written in the code, so code can be shared safely.

```mujard
at start:
    wifi.connect(secret("home_wifi"))
    mqtt.connect("broker.example.com", password: secret("mqtt_password"))
```

- **Write:** `secret(<name>)`; you add the value once in the app (Project > Secrets). A secret is either text (a password, a
  key) or credentials (a Wi-Fi name and password together); `wifi.connect` takes credentials.
- The value goes into the board when it is flashed, encrypted where the board supports it, and never into project files,
  shared copies or build logs; files you can see (`--show-cpp`, `.gen/`) show a placeholder instead. On boards without
  flash encryption, anyone holding the board can read the flash: the value is safe in your code, not from someone who
  owns the board.
- **Arduino:** not possible; passwords are typed into the sketch.

---

## 2. The board itself

### 11. `board.name`

```mujard
print "Running on {board.name}"      # "Arduino Uno", "ESP32-S3 DevKit", ...
```
**Arduino:** write it yourself with `#ifdef ARDUINO_AVR_UNO` ... for each board.

### 12. `board.chip`

```mujard
print "Chip: {board.chip}"           # "ATmega328P", "ESP32-S3", "RP2040"
```
**Arduino:** library or core-specific (`ESP.getChipModel()` on ESP32 only).

### 13. `board.id`

A unique serial number for this physical board, the same after every restart.

```mujard
show "Board", board.id
mqtt.publish("bins/{board.id}/level", fill)    # each bin uses its own topic
```
Returns text, e.g. `"3C61-05A8-9F12"`. **Arduino:** core-specific.

### 14. `board.uptime`

Time since power-up, as a time value.

```mujard
every 1 min:
    show "Running for", board.uptime          # "2 h 14 min"
```
**Arduino:** `millis()` (a number, which overflows after 49 days; `board.uptime` doesn't).

### 15. `board.free_memory`

Free RAM, in bytes.

```mujard
when board.free_memory < 200 bytes:
    log.warn "Low memory: {board.free_memory}"
```
**Arduino:** write it yourself (a `freeMemory()` function differs per chip).

### 16. `board.voltage`

The board's supply voltage, measured by the chip where it can.

```mujard
every 10 s:
    var v = board.voltage
    if v != none and v < 4.6 V:
        log.warn "USB power is weak ({v}) - try another cable"
```
Returns `none` on boards that can't measure it. **Arduino:** write it yourself (register tricks).

### 17. `board.restart()`

```mujard
on button.held(5 s):
    log.info "Restarting"
    board.restart()

after 24 h:
    board.restart()                          # a daily fresh start for long-running projects
```
**Arduino:** core-specific (`ESP.restart()`) or a watchdog trick on AVR.

### 18. `board.reset_reason`

Why the board last started: `power_on`, `reset_button`, `watchdog`, `brownout`, `software`, `wake` (from deep sleep) or `crash`.

```mujard
at start:
    if board.reset_reason == crash or board.reset_reason == watchdog:
        notify "The controller restarted after a problem ({board.reset_reason})"
```
**Arduino:** core-specific.

---

## 3. Time

### 19. `now()`

The current time since start, as a time value. Mostly used with `since()`.

```mujard
var started = now()
```
**Arduino:** `millis()` / `micros()`.

### 20. `since(t)`

How long ago `t` was.

```mujard
var last_seen = now()

every 100 ms:
    if hand.distance() < 20 cm:
        last_seen = now()
    if since(last_seen) > 3 s:
        lid.move(0 deg)
```
**Arduino:** `millis() - lastSeen` (and remembering it is milliseconds).

### 21. `stopwatch()`

```mujard
var lap = stopwatch()

on start_gate.pressed:
    lap.start()
on finish_gate.pressed:
    lap.stop()
    print "Time: {lap.elapsed as s}"
on reset_button.pressed:
    lap.reset()
```
- **Methods:** `start()`, `stop()`, `reset()`, `elapsed`, `running`
- **Arduino:** write it yourself.

### 22-24. `clock.time`, `clock.date`, `clock.sync()`

The time of day and date, from an RTC module or the internet.

```mujard
board esp32

at start:
    wifi.connect(secret("home_wifi"))
    clock.sync(timezone: "Africa/Lagos")       # set from the internet

every 1 min:
    print "{clock.date} {clock.time}"          # 2026-10-18 14:05

at 07:30:                                      # the 'at' keyword with a time of day
    alarm.melody("C5 E5 G5")
```
- Boards without Wi-Fi use an RTC module: `use mujard.rtc` and `clock.source(rtc)`.
- `clock.time` is `none` until the clock has been set.
- **Arduino:** RTC or NTP libraries, plus time-zone code.

---

## 4. Maths

### 25-27. `abs`, `min`, `max`

```mujard
var error = abs(target - position)
var safe  = min(requested, 80 %)
var loud  = max(left_mic.percent, right_mic.percent)
```
`min` and `max` take two or more values. **Arduino:** `abs`, `min`, `max`.

### 28. `clamp(x, range)`

Keeps a value inside a range.

```mujard
motor.write(clamp(speed, 0..80 %))
```
**Arduino:** `constrain(speed, 0, 80)`.

### 29. `scale(x, from, to)`

Converts a value from one range to another. Works with decimals and units, and can run backwards.

```mujard
var percent = scale(light.value, 0..1023, 0..100 %)
var angle   = scale(knob.percent, 0..100 %, 0..180 deg)
var fill    = scale(level.distance(), 40 cm..5 cm, 0..100 %)   # an empty bin is far, a full bin is near
```
**Arduino:** `map()`, which works only with whole numbers (`map(512, 0, 1023, 0, 100)` gives 50, never 50.05).

### 30. `round(x)`

```mujard
print "Temperature: {round(t, 1)}"    # one decimal place, e.g. 23.4 C
```
`round(x)` rounds to a whole number, `round(x, n)` to n decimals. **Arduino:** `round()` (whole numbers only).

### 31-32. `math.sqrt`, `math.sin` (and the rest of `mujard.math`)

```mujard
use mujard.math

var tilt   = math.sqrt(ax * ax + ay * ay)
var height = 2 m * math.sin(30 deg)          # angles in degrees, no radians needed
```
`mujard.math` also contains `cos`, `tan`, `atan2`, `pow`, `log`, `PI` and `E`. **Arduino:** `sqrt`, `sin` (radians only).

---

## 5. Smart signals

These remember things between runs. Each one has its own memory for **each place in the code where you write it**, and
the memory is updated only when that line runs (a call skipped by `and`, `or` or an `if` doesn't update it). A function
that uses one shares the memory between everyone who calls the function. At the first run, `rising` (like `when`) counts a
condition that is already true as a rise; `falling` and `changed` take the first value as their starting point and
don't fire.

### 33. `smooth(source, n)`

The average of the last `n` readings: steady values from noisy sensors.

```mujard
every 100 ms:
    var t = smooth(temp_sensor.value, 10)     # average of the last 10 readings
    show "Temperature", t
```
**Arduino:** write it yourself (an array, an index and a running sum).

### 34. `between(x, range)`

```mujard
if between(t, 20 C..26 C):
    status.on()                               # comfortable
```
Range ends are included. **Arduino:** `t >= 20 && t <= 26`.

### 35. `changed(x)`

True once each time the value changes.

```mujard
every 50 ms:
    if changed(door.read()):
        log.info "Door is now {door.read()}"
```
**Arduino:** write it yourself (store the previous value).

### 36-37. `rising(x)`, `falling(x)`

True once when a true/false value switches on (rising) or off (falling).

```mujard
every 50 ms:
    if rising(t > 30 C):
        fan.on()                              # once, when it gets hot
    if falling(t > 30 C):
        log.info "Cooled down"
```
The `when` keyword is a shortcut for `rising`: `when t > 30 C:` does the same as the first `if`.
**Arduino:** write it yourself.

### 38. `hysteresis(x, on:, off:)`

Switches on above one level and off below another, so it never flickers around a single threshold.

```mujard
every 1 s:
    heater.write(hysteresis(t, on: 19 C, off: 21 C))   # heats below 19 C, stops above 21 C
```
If `on:` is below `off:`, it switches on when the value falls (heating); if above, when it rises (cooling).
**Arduino:** write it yourself.

### 39. `peak(x, window)`

The highest value seen during the last time window.

```mujard
every 100 ms:
    show "Loudest (10 s)", peak(mic.percent, 10 s)
```
**Arduino:** write it yourself.

---

## 6. Random

### 40. `random(range)`

```mujard
var dice  = random(1..=6)                  # 1 to 6
var delay = random(200 ms..800 ms)          # a random time
var coin  = random(bool)                    # true or false
```
**Arduino:** `random(1, 7)` (where the top is excluded and easy to get wrong).

### 41. `random.seed()`

Starts the random numbers from real noise, so each power-up gives a different sequence. Mujard calls it automatically at
start; call it yourself only to choose a fixed seed for repeatable tests: `random.seed(42)`.

**Arduino:** `randomSeed(analogRead(0))`, forgotten in most sketches.

### 42. `random.pick(list)`

```mujard
const COLOURS = ["red", "green", "blue"]
print "Today's colour: {random.pick(COLOURS)}"
```
**Arduino:** write it yourself.

---

## 7. Digital pins

Declare first:

```mujard
led    = pin 13 as output
relay  = pin 7  as output inverted     # "on" drives the pin LOW
button = pin 2  as input pullup
trig   = pin 9  as output
echo   = pin 10 as input
```

Pin names follow the board's labels: `13`, `A2`, `D5` (NodeMCU), `GP15` (Pico), `GPIO5` (ESP32) and `LED` (the built-in
LED). Wrong, dangerous or double-used pins are compile errors, with the reason.

### 43. `on()`

```mujard
led.on()
relay.on()        # inverted: drives LOW, which switches most relay boards on
```
Stops a running `blink` or `pattern` first. **Arduino:** `digitalWrite(13, HIGH)`.

### 44. `off()`

```mujard
led.off()
```
**Arduino:** `digitalWrite(13, LOW)`.

### 45. `toggle()`

```mujard
on button.pressed:
    led.toggle()
```
Works from the last value set, so it is reliable even where reading an output pin back doesn't work.
**Arduino:** `digitalWrite(13, !digitalRead(13))`.

### 46. `write(v)`

```mujard
led.write(button.read())               # the LED follows the button
lamp.write(light.percent < 30 %)       # on when it's dark
```
**Arduino:** `digitalWrite(13, value)`.

### 47. `read()`

```mujard
if button.read():
    print "Button is down"
```
Gives `true`/`false` for the pin's *active* level: HIGH normally, LOW for a `pullup` input, swapped by `inverted`. So a
button wired to ground on a pull-up pin reads `true` when it is pressed.
To react to presses, use the events in [section 9](#9-buttons-and-inputs), which never miss one.
**Arduino:** `digitalRead(2) == LOW` (and remembering why it's `LOW`).

### 48. `blink(rate, for:, on:)`

Blinks by itself in the background.

```mujard
led.blink(2 Hz)                       # forever
led.blink(10 Hz, for: 3 s)            # 3 seconds, then off
led.blink(1 Hz, on: 100 ms)           # a short flash once a second
```
Stop with `off()`, `on()` or `stop()`. **Arduino:** write it yourself with `millis()`.

### 49. `pattern(text, step:, repeat:)`

```mujard
led.pattern("SOS")                        # Morse code, repeated
led.pattern("x.x...", step: 200 ms)       # x = on, . = off
led.pattern("xxx", repeat: 1)             # three flashes, once
```
Default `step` is 150 ms; default `repeat` is forever. **Arduino:** write it yourself.

### 50. `measure_pulse(level:, timeout:)`

How long the pin stays in one state: used for ultrasonic sensors, IR remotes, frequency counting.

```mujard
every 100 ms:
    trig.pulse(10 us)                              # a single 10 µs pulse out
    var t = echo.measure_pulse(timeout: 30 ms)
    if t == none:
        print "Nothing in range"
    else:
        print "Distance: {t * 343 m/s / 2 as cm}"
```
- `level:` is `high` (default) or `low`; `timeout:` defaults to 1 s.
- Returns a time, or `none` on timeout.
- The output side is `pin.pulse(duration)`.
- **Arduino:** `pulseIn(10, HIGH, 30000)`, which returns `0` on timeout.

---

## 8. Analog and PWM

Declare first:

```mujard
light = analog A0
lamp  = pin 9 as pwm
fan   = pin 10 as pwm
```

### 51. `value`

The raw reading (0-1023 on an Uno, 0-4095 on an ESP32).

```mujard
print "Raw: {light.value}"
```
**Arduino:** `analogRead(A0)`.

### 52. `volts`

The reading in volts, using the board's real reference voltage.

```mujard
print "Battery: {battery_pin.volts * 2}"     # through a 1:2 voltage divider
```
**Arduino:** `analogRead(A0) * 5.0 / 1023` (wrong on 3.3 V boards).

### 53. `percent`

The reading as 0-100 %.

```mujard
lamp.write(100 % - light.percent)            # brighter when the room is darker
```
**Arduino:** `map(analogRead(A0), 0, 1023, 0, 100)`.

### 54. `write(level)`

PWM output, as a percentage.

```mujard
lamp.write(50 %)
fan.write(clamp(speed, 20 %..100 %))
```
**Arduino:** `analogWrite(9, 127)` (0-255 on most boards, different on others).

### 55. `fade(to:, over:)`

Changes the level smoothly in the background.

```mujard
lamp.fade(to: 100 %, over: 2 s)       # sunrise
lamp.fade(to: 0 %, over: 500 ms)
```
**Arduino:** write it yourself with `millis()` steps.

### 56. `frequency(f)`

Sets how fast the PWM switches: for quiet fans, motors and LED drivers.

```mujard
fan.frequency(25 kHz)                 # PC fans expect 25 kHz
```
A compile error on boards whose PWM frequency is fixed. **Arduino:** core-specific (`ledcSetup` on ESP32, timer
registers on AVR).

### 57. `reference(v)`

The analog reference voltage: for more precise readings of small voltages.

```mujard
at start:
    light.reference(1.1 V)            # the Uno's internal 1.1 V reference
```
Only the values the board supports are allowed. **Arduino:** `analogReference(INTERNAL)`.

---

## 9. Buttons and inputs

Events, used with `on`. All are debounced (20 ms by default).

```mujard
button = pin 2 as input pullup
```

### 58. `pressed`

```mujard
on button.pressed:
    led.toggle()
```

### 59. `released`

```mujard
on button.released:
    print "Let go"
```

### 60. `held(time)`

Fires once after the button has been held down for that long.

```mujard
on button.held(3 s):
    board.restart()
```

### 61. `clicked`

A short press (released before a `held` would fire), so one button can have both a click and a long press. If the same
button also has a `double_clicked` handler, `clicked` arrives after the 400 ms double-click window; otherwise it arrives
when the button is released.

```mujard
on button.clicked:
    mode = next_mode(mode)
on button.held(2 s):
    mode = off
```

### 62. `double_clicked`

Two clicks within 400 ms.

```mujard
on button.double_clicked:
    lamp.write(100 %)
```

### 63. `changed`

Any change, in either direction. Add `fast` to react within microseconds using a hardware interrupt. `rising` and
`falling` events work the same way when you only want one direction (`on encoder_a.rising fast:`). `fast` events are not
debounced, and a variable that a `fast` handler changes is read safely by the rest of the program (never half-updated).

```mujard
on door.changed:
    log.info "Door is {door.read()}"

on encoder_a.changed fast:            # a hardware interrupt: for rotary encoders and counters
    count += 1
```

### 64. `debounce(time)`

Changes the debounce time for noisy switches.

```mujard
at start:
    reed_switch.debounce(50 ms)
```

**Arduino for 58-64:** write it yourself: a state variable, a timer and edge detection for each button, plus
`attachInterrupt` for `fast`.

---

## 10. Sound

```mujard
buzzer = pin 8 as output
```

### 65. `tone(frequency, length)`

```mujard
buzzer.tone(2 kHz, 200 ms)            # a beep
buzzer.tone(440 Hz)                   # until stopped
```
**Arduino:** `tone(8, 2000, 200)`.

### 66. `stop()`

```mujard
buzzer.stop()
```
**Arduino:** `noTone(8)`.

### 67. `melody(notes, tempo:)`

Plays notes in the background. Notes are `C4`, `D#5` and so on, `-` is a rest, and `/2` or `*2` makes a note half or
double length.

```mujard
on doorbell.pressed:
    buzzer.melody("E5 C5 - E5 C5")
buzzer.melody("C4 E4 G4 C5*2", tempo: 140)   # beats per minute, default 120
```
**Arduino:** write it yourself (a table of frequencies and durations, plus timing).

---

## 11. Saving data

Kept through power cuts, in EEPROM or flash depending on the board.

### 68. `saved var`

A variable that keeps its value when the power goes off.

```mujard
saved var opens = 0
saved var target_temp = 21 C

on lid_opened:
    opens += 1                        # still counting after a restart
```
- Saved automatically when it changes: at most once every 5 s (60 s on boards without wear-levelled storage, such as
  the Uno), spread over the free EEPROM or flash so no single cell wears out. Anything still waiting is written just
  before `power.sleep`, `power.deep_sleep`, `board.restart()` and `board.halt`; a sudden power cut can lose the last few
  seconds of changes.
- Changing its type in the code starts it fresh, instead of loading garbage.
- **Arduino:** `EEPROM.put(address, value)` / `EEPROM.get(...)`, with addresses you manage yourself.

### 69. `store.save(name, value)`

Saves a named value when you choose to.

```mujard
on calibrate.pressed:
    store.save("zero", scale_sensor.value)
```

### 70. `store.load(name, default:)`

```mujard
at start:
    var zero = store.load("zero", default: 0)
```
Returns the default (or `none`) if nothing was saved under that name.

**Arduino for 69-70:** the EEPROM library, with fixed addresses and sizes.

### `store.pending` · `store.failed` · `store.committed` · `store.last_commit`

```mujard
every 1 s:
    if store.failed > 0:
        led.blink(4 Hz)
```

Writes are paced to protect the storage, so a save doesn't land the instant it runs. `store.pending` is `true`
while a `store.save` or a `saved var` change is still waiting to reach storage (powering off while it's true can
lose that change); `store.failed` counts commits that didn't land; `store.committed` and `store.last_commit` say
how the most recent commit went.

---

## 12. Communication

### 71. `serial <speed>`

Starts the serial port. Optional, since `print` starts it at 115200 baud if needed.

```mujard
at start:
    serial 9600 baud
```
**Arduino:** `Serial.begin(9600)`.

### 72. `serial.read_line()` and `on serial.line`

```mujard
on serial.line as cmd:                # an event for each line typed in the monitor
    if cmd == "on":
        led.on()
    elif cmd.starts_with("speed "):
        var n = cmd.split(" ")[1].to_number()
        if n != none:
            fan.write(clamp(scale(n, 0..100, 0..100 %), 0..100 %))
```
`serial.read_line()` is the same thing used inside a `task`; it waits for one line.
**Arduino:** `Serial.readStringUntil('\n')`, which blocks.

### 73. `i2c.scan()`

Lists the I2C devices connected, for checking wiring.

```mujard
use mujard.i2c

at start:
    print "Found: {i2c.scan()}"       # e.g. [0x27, 0x68]
```
**Arduino:** the "I2C scanner" example sketch.

### 74. `i2c.write(address, data)` / `i2c.read(address, count)`

```mujard
use mujard.i2c

i2c.write(0x68, [0x6B, 0x00])          # wake an MPU-6050
var raw = i2c.read(0x68, 14)           # 14 bytes, or none if the device didn't answer
```
**Arduino:** `Wire.beginTransmission`, `Wire.write`, `Wire.endTransmission`, `Wire.requestFrom`, `Wire.read`.

### 75. `spi.transfer(data)`

```mujard
use mujard.spi

chip = spi.device(cs: pin 10, speed: 1 MHz, mode: 0)
var reply = chip.transfer([0x9F, 0, 0, 0])
```
**Arduino:** `SPI.beginTransaction`, `digitalWrite(cs, LOW)`, `SPI.transfer`, `digitalWrite(cs, HIGH)`, `SPI.endTransaction`.

### 76. `shift.out(data)`

Drives shift registers such as the 74HC595, to get more outputs from three pins.

```mujard
use mujard.shift

leds = shift.out_register(data: pin 11, clock: pin 12, latch: pin 8)
leds.write(0b1010_1010)
```
**Arduino:** `shiftOut(11, 12, MSBFIRST, value)` plus the latch pin by hand.

### More serial ports: `mujard.uart`

For GPS, Bluetooth and GSM modules. It is a library (like `servo`), so it is not one of the 100 numbered built-ins.

```mujard
use mujard.uart

gps = Uart(rx: pin 4, tx: pin 3, speed: 9600 baud)

on gps.line as sentence:
    if sentence.starts_with("$GPGGA"):
        print sentence
```
`gps.write(text)` sends; `gps.read_line()` waits for one line inside a `task`. A hardware port is used where the board
has a spare one (Mega, ESP32, Pico) and a software port otherwise (a warning above 9600 baud on an Uno).
**Arduino:** `SoftwareSerial` or `Serial1`.

---

## 13. Network

For boards with Wi-Fi (ESP32, ESP8266, Pico W, Uno R4 WiFi). On other boards these are compile errors that say why.

### 77. `wifi.connect(credentials)`

```mujard
board esp32

at start:
    wifi.connect(secret("home_wifi"))  # the network name and password from the app's secrets
```
Connects in the background and reconnects by itself if the network drops.
**Arduino:** `WiFi.begin(ssid, password)`, a waiting loop, and your own reconnect code.

### 78. `wifi.connected`

```mujard
every 1 s:
    status_led.write(wifi.connected)

on wifi.lost:
    log.warn "Wi-Fi lost, reconnecting"
```
Events: `on wifi.joined`, `on wifi.lost`. **Arduino:** `WiFi.status() == WL_CONNECTED`.

### 79. `http.get(url)`

```mujard
task weather:
    var reply = http.get("https://api.example.com/weather?city=Lagos")
    if reply.ok:
        show "Outside", reply.json["temp"] as C
```
Used inside a `task` (it waits for the answer, at most 5 s, while the rest of the program keeps running). `reply` has `ok`, `status`, `text` and `json` - index `reply.json["temp"]` with a name or a number, `.ok` tells whether the entry exists, `.to_number()` gives the number or `none`, and `as C` makes a quantity straight away. `https://` addresses work too: the certificate is always checked - the ESP32 uses the core's CA bundle, the ESP8266 and the Pico W use roots built into the sketch. Where the check can't run, an `https://` address is refused with `status -1` rather than sent unchecked.
`http.post(url, data)` works the same way. **Arduino:** the `HTTPClient` library.

### 80-81. `mqtt.publish`, `mqtt.subscribe`

```mujard
use mujard.mqtt

at start:
    mqtt.connect("broker.example.com", user: "bin7", password: secret("mqtt"))
    mqtt.subscribe("bins/{board.id}/command")

every 1 min:
    mqtt.publish("bins/{board.id}/level", fill)

on mqtt.message as m:
    if m.text == "open":
        lid.move(90 deg)
```
`mqtt.connect` and `mqtt.publish` never wait: the link comes up in the background once Wi-Fi is up, up to 4 publishes are kept while it is down, and subscriptions come back after a reconnect.
**Arduino:** a library such as PubSubClient, plus reconnect handling.

---

## 14. Power

### 82. `power.sleep(time)`

Light sleep: pins and memory stay as they are, and `wake_on:` (an event such as a button press) can wake it early. **It
pauses the whole program** (every block, task and timer), not just the code that called it. Missed `every` runs are
skipped on waking. Use it when nothing else needs to keep running.

```mujard
task logger:
    while true:
        log_reading()
        power.sleep(50 s, wake_on: button.pressed)
        wait 10 s
```

### 83. `power.deep_sleep(time)`

Deep sleep: lowest power. The board restarts when it wakes, so `at start:` runs again (use `saved var` to remember
things; pending changes are written first).

```mujard
board esp32
saved var readings = 0
soil = analog 34

at start:
    wifi.connect(secret("home_wifi"))      # connects in the background

task report:
    var t0 = now()
    readings += 1
    wait until wifi.connected or since(t0) > 20 s
    if wifi.connected:
        mqtt.connect("broker.example.com", password: secret("mqtt"))
        mqtt.publish("garden/soil", soil.percent)    # in a task, waits until it is sent
    power.deep_sleep(30 min)
```
**Arduino for 82-83:** board-specific libraries (LowPower on AVR, `esp_deep_sleep` on ESP32).

### 84-85. `battery.percent`, `battery.volts`

```mujard
battery = battery_on(analog A1, divider: 2, type: lipo_1s)

every 5 min:
    show "Battery", battery.percent
    when battery.percent < 15 %:
        notify "Charge me!"
```
`type:` (`lipo_1s`, `aa_x3`, `aa_x4` ...) tells Mujard how voltage maps to charge, which isn't a straight line.
**Arduino:** write it yourself (divider maths plus a discharge curve).

---

## 15. Safety

### 86. `expect condition, message`

A check that must stay true. If it fails, Mujard logs the message, reports it to the app, runs `safe_state` and
**halts** the program (USB and uploads still work; reset or power-cycle to start again). The message is optional; without
one, the condition's own text is used. Add `halt: false` for a condition that can clear by itself: the program then
carries on and the check runs again next time.

```mujard
every 1 s:
    expect motor_temp < 85 C, "Motor overheating", halt: false   # it can cool down again
    expect since(last_heartbeat) < 5 s, "Lost contact with the remote"
```
**Arduino:** write it yourself (`if (...) { ... }` everywhere, plus reporting).

### 87. `limit thing to range`

Declares a limit once; every later write is clamped to it, so a bug or a phone command can't exceed it.

```mujard
motor = pin 5 as pwm
limit motor to 0..80 %          # the gearbox can't take full power

motor.write(100 %)              # actually sets 80 %, and logs a warning once
```
**Arduino:** `constrain` at every single write.

### 88. `safe_state:`

What to do whenever something goes wrong: a failed `expect`, a crash or a watchdog restart.

```mujard
safe_state:
    heater.off()
    motor.write(0 %)
    alarm.pattern("SOS")
```
It also runs first at every start-up, so outputs are safe before the rest of the program begins.
**Arduino:** no equivalent; scattered by hand.

### 89. `watchdog(time)`

Restarts the board if the program stops responding for that long.

```mujard
at start:
    watchdog(2 s)
```
Mujard keeps the watchdog fed by itself while the program is healthy; a stuck `cpp:` block or a crash lets it run out.
**Arduino:** `wdt_enable(WDTO_2S)` on AVR, different on every other chip.

---

## 16. Testing without hardware

`test` blocks run in the Mujard simulator with `mujard test`, or with the Test button in the app. They are never
included in the firmware.

```mujard
board uno
use mujard.servo
use mujard.ultrasonic

lid  = Servo(pin 6)
hand = Ultrasonic(trig: pin 9, echo: pin 10)

every 100 ms:
    if hand.distance() < 20 cm:
        lid.move(90 deg)
    ...
```

### 90. `pretend(value)`

Makes a sensor or input report a value of your choice.

```mujard
test "opens when a hand is near":
    hand.pretend(10 cm)
    advance(200 ms)
    expect lid.angle == 90 deg
```

### 91. `advance(time)`

Moves simulated time forward instantly.

```mujard
test "closes three seconds after the hand leaves":
    hand.pretend(10 cm)
    advance(200 ms)
    hand.pretend(80 cm)
    advance(3.5 s)
    expect lid.angle == 0 deg
```

### 92. `notified(text)`

Did the program send that notification? Partial text matches.

```mujard
test "warns when the bin is full":
    level.pretend(4 cm)
    advance(1 min)
    expect notified("almost full")
```

### 93. `printed(text)`

Did the program print that?

```mujard
test "says hello":
    advance(1 s)
    expect printed("Ready")
```

Teachers can attach tests to an LMS assignment, so the app grades a student's program automatically.
**Arduino for 90-93:** not possible.

---

## 17. Text

### 94. `length`

```mujard
var name = "Bridge"
print "{name.length}"                # 6
```
**Arduino:** `String.length()`.

### 95. `contains(part)`

```mujard
on serial.line as cmd:
    if cmd.contains("help"):
        print "Commands: on, off, speed <0-100>"
```
Also `starts_with` and `ends_with`. **Arduino:** `cmd.indexOf("help") >= 0`.

### 96. `split(separator)`

```mujard
var parts = "21.5,48,1013".split(",")   # ["21.5", "48", "1013"]
```
**Arduino:** write it yourself (`indexOf` and `substring` in a loop).

### 97. `to_number()`

```mujard
var t = parts[0].to_number()            # 21.5, or none if it isn't a number
```
**Arduino:** `toInt()` / `toFloat()`, which return `0` for bad input.

Mujard text never grows without limit: its maximum size is fixed when compiled, so there is no surprise memory use like
Arduino's `String` on small boards.

---

## 18. Bits

For registers, flags and protocols.

### 98. `bit(n)`

```mujard
var status = i2c.read(0x68, 1)[0]
if status.bit(3):
    print "Data ready"
```
**Arduino:** `bitRead(status, 3)`.

### 99. `set_bit(n)` (and `clear_bit(n)`, `toggle_bit(n)`)

```mujard
var flags = 0
flags = flags.set_bit(2)
```
**Arduino:** `bitSet(flags, 2)`.

### 100. `bytes`

```mujard
var reading = 0x1234
print "{reading.bytes}"                 # [0x12, 0x34] (highest byte first)
i2c.write(0x50, [0x00] + reading.bytes)
```
**Arduino:** `highByte(reading)`, `lowByte(reading)`.

The usual operators also work: `&` `|` `^` `<<` `>>` `~`.

---

## 19. More links and more cores

These came after the first hundred - the numbers continue the spec's own list. Each needs its `use` line at the top.

### 101-104. `Can(...)` · `bus.send` · `bus.read` · `on bus.message`

```mujard
board esp32
use mujard.can

bus = Can(pin 21, pin 22, 500000 baud)

every 1 s:
    bus.send(0x123, [1, 2, 3])

on bus.message as m:
    print "CAN {m.id}: {m.data}"
```
Opens the board's CAN/TWAI controller (the ESP32 family). `send` queues a frame of up to 8 bytes and is `false`
when the queue is full; `read` gives the next received frame or `none`. A message `m` has `id`, `data` and `rtr`.
Boards without a controller refuse (`C320`).

### 105-107. `ble.start` · `ble.send` / `ble.read_line` / `ble.connected` · `on ble.line`

```mujard
use mujard.ble

at start:
    ble.start("garage-door")

on ble.line as l:
    if l == "open":
        door.on()
```
The board advertises as a BLE peripheral a phone or another board can connect to and exchange text lines with.
`send` is `false` when nothing is connected. ESP32-family (except S2) and Nano 33 BLE.

### 108-109. `ethernet.start` / `ethernet.connected` / `ethernet.ip` · `on ethernet.joined` / `on ethernet.lost`

```mujard
use mujard.ethernet

at start:
    ethernet.start()

on ethernet.joined:
    print "Wired up, {ethernet.ip}"
```
A wired Ethernet link on the boards that have one (the classic ESP32's EMAC + external PHY). `http.get`, `mqtt`
and `clock` work over the wire too.

### 110-111. `usb.type` / `usb.key` · `usb.mouse_move` / `usb.click` / `usb.connected`

```mujard
use mujard.usb

on button.pressed:
    usb.type("password123")
    usb.key("enter")
```
The board acts as a keyboard+mouse to whatever it's plugged into; each call is `false` when the HID device isn't
up. Standard-HID boards only: Leonardo, Micro, Pico, Pico 2, Zero and Uno R4.

### 112. `every ... on core 1:`

```mujard
board esp32

task watch on core 1:
    while true:
        wait until alert
        bell.tone(2 kHz, 200 ms)
```
A top-level `every`/`after`/`when`/`task` can run on a second core (`board.cores` says how many the board has).
A plain `var` shared across cores races - the compiler warns **W620**; pass values through a `Queue` instead, and
peripherals that aren't safe to share warn **W621**.

### `Queue(type, capacity)` - `use mujard.queue`

```mujard
use mujard.queue

q = Queue(text, 8)

task sender on core 1:
    while true:
        q.push("ping")
        wait 500 ms

every 100 ms:
    if q.available:
        print q.pop()
```
A bounded queue - the safe way to hand values between cores or blocks. `push` copies a value in (`false` when
full; the drop is counted in `board.dropped_events`), `pop` gives the oldest or `none`, and `available`, `full`
and `n` report the state. Element types: `int`, `float`, `bool`, `text` or a measurement; capacity 1 to 64.

### 113-114. `Timer()` · `t.every` / `t.freq` / `t.once` / `t.start` / `t.stop` · `on t.tick` / `t.elapsed`

```mujard
use mujard.timer

t = Timer()

at start:
    t.every(1 ms)
    t.start()

on t.tick:
    step.toggle()
```
A hardware timer - a real timer interrupt at exact periods, not the cooperative `every`. `on t.tick` runs on the
scheduler; `on t.tick fast` runs inside the ISR itself (keep it to pin writes and numbers). Boards without a
timer backend refuse (`C320`).

### `regs.` - `use mujard.regs`

```mujard
use mujard.regs

at start:
    regs.PORTB.set(regs.PORTB.PORTB0)
```
Typed register and bit-field names for the board's chip, generated from the chip maker's own description files.
`.read()` reads, `.write(v)` `.set(bits)` `.clear(bits)` write, and `regs.PORTB.PORTB0` is that bit's mask.
Every board in the catalog is covered; the per-chip tables are under `docs/lang/reference`.

---

## Mujard compared with Arduino

For Mujard's first hundred built-ins, Arduino offers:

| Arduino has... | Count | Which |
|---|---|---|
| A built-in function | 32 | 1, 14, 19, 25-32, 40-41, 43-44, 46-47, 50, 51, 54, 57, 65-66, 71-72, 76, 94-95, 97-100 |
| A library or board-specific code | 21 | 12-13, 17-18, 22-24, 56, 68-70, 74-75, 77-83, 89 |
| Nothing: write it yourself | 38 | 4, 7-9, 11, 15-16, 20-21, 33-39, 42, 45, 48-49, 52-53, 55, 58-64, 67, 73, 84-88, 96 |
| Nothing possible (needs the app, cloud or simulator) | 9 | 2-3, 5-6, 10, 90-93 |

Nearly everything Arduino offers has a Mujard counterpart, through built-ins, keywords (`at start:` for `setup()`,
`pin ... as` for `pinMode`), units or the libraries (`mujard.uart` covers extra serial ports such as GPS and Bluetooth
modules; `mujard.timer`, `mujard.regs`, `mujard.usb` and `mujard.ble` cover hardware timers, register access, USB
keyboard/mouse and Bluetooth LE). What is left to `cpp:` or a library: I2C and SPI slave mode, DMA, and ESP32 extras
such as touch and DAC.

---

## Appendix: keywords

36 reserved words (they can't be used as names):

| Group | Keywords |
|---|---|
| Program and board | `board` `use` `import` |
| Declarations | `const` `var` `saved` `fn` `struct` `enum` |
| Hardware | `pin` `analog` `as` |
| Decisions and loops | `if` `elif` `else` `while` `for` `in` `match` `break` `continue` `return` |
| Events and time | `every` `on` `when` `after` `at` `task` `wait` |
| Logic | `and` `or` `not` |
| Values | `true` `false` `none` |
| Escape hatch | `cpp` |

Words with a special meaning only in one place: `input` `output` `pwm` `pullup` `pulldown` `inverted` (pin roles),
`start` (after `at`), `fast` (after an event), `until` (after `wait`), `pass` (a line on its own),
`test`, `limit`, `safe_state`, `control`, `print` and `expect` (at the start of a line), `pins` (starting a pin-group
declaration), `to` (in a `limit` line), `mod` (between two values), `version`, `idf` and `pio` (in a `use` line),
and the unit names (straight after a number).

A word right after a `.` is always a member name, never a keyword, so `led.on()` and `on button.pressed:` can both exist.

Names that are reserved in C++ but not in Mujard (`new`, `class`, `int`, `delete`, ...) can be used freely; the compiler
renames them in the C++ it generates.