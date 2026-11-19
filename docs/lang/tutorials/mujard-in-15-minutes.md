# Mujard in 15 minutes

Everything you need to write, check, build and run your first Mujard program - and the five things that make
it different from Arduino C++. Programs live in `.muj` files.

## 1. Get the tool

You need the `mujard` command. It checks your program, turns it into Arduino C++, builds firmware and flashes it.

```console
mujard --version
```

## 2. A first program (2 minutes)

Make a file `blink.muj`:

```muj
board uno

led = pin LED as output

every 500 ms:
    led.toggle()
```

Line by line:

- `board uno` - the target. The checker knows this board's pins, memory and peripherals; change it to `esp32`,
  `pico`, `mega` and the same program is checked against that board instead.
- `led = pin LED as output` - a device. `LED` is the built-in LED (pin 13 on an Uno); `pin 13` works too.
  The name `led` is yours.
- `every 500 ms:` - a timed block. It runs on its own schedule - there is no `loop()` and no `delay()`.
  `500 ms` is a real unit, not a bare number.
- `led.toggle()` - flips the LED each time.

Indentation is spaces, 4 per level. Comments start with `#`.

## 3. Check it (1 minute)

```console
mujard check blink.muj
```

Misspell `led.toggle()` as `led.togle()` and check again: you get `blink.muj:5:9  C308  A output has no 'togle'.
Did you mean 'toggle'?` - the compiler knows the board, the language and your devices. A pin that can't do
what you ask (PWM on a non-PWM pin, a strapping pin on ESP32, more RAM than the board has) is a *warning or
error before anything is built*, not a mystery on the bench.

## 4. See the C++ (optional)

```console
mujard gen blink.muj --out out
```

`out/` has `sketch.ino` plus the Mujard runtime files - readable Arduino C++ you can also hand to the IDE.

## 5. Build and run (3 minutes)

```console
mujard build blink.muj
mujard run blink.muj --port COM3
```

`build` compiles firmware (the direct backend first, arduino-cli as fallback). `run` builds and flashes a board
on the USB port. The LED should blink twice a second.

## 6. Make it react (3 minutes)

```muj
board uno

led    = pin LED as output
button = pin 2 as input pullup

on button.pressed:
    led.toggle()
    print "Click!"
```

`on button.pressed:` is an event: it runs once per press, debounced for you. No polling loop, no
`digitalRead()` bookkeeping. `print` goes to the serial monitor - and to the Mujard app when it's attached.

## 7. Read a knob (2 minutes)

```muj
knob = analog A0

every 100 ms:
    print(knob.percent)
```

`analog A0` reads the pin; `.percent` is the reading scaled to the pot. Units work both ways:

```muj
var angle = scale(knob.percent, 0 %..100 %, 0 deg..180 deg)
```

## 8. Drive a servo (2 minutes)

```muj
board uno
use mujard.servo

arm = Servo(pin 9)

every 20 ms:
    arm.move(scale(knob.percent, 0 %..100 %, 0 deg..180 deg))
```

`use mujard.servo` brings in the servo library - `Servo(pin 9)` is a device like the pin, and the checker
already knows the pin it claims.

## 9. Go wireless (2 minutes)

```muj
board esp32

at start:
    wifi.connect(secret("home_wifi"))

every 10 s:
    notify "still alive: {board.uptime}"
```

`board esp32` retargets the same language. `secret("home_wifi")` means the password lives in the build's
secrets, never in the file. `notify` pings your phone - over the app when it's attached, or straight to the
server when the board is on its own Wi-Fi.

## What just happened

You wrote real firmware with: no `setup()`/`loop()`, no `delay()`, no `digitalRead`/`digitalWrite`, no type
casts - and every mistake was caught before it cost you a flash.

## Next

- **45 worked examples** in `docs/lang/examples/learn/` (Basics to Advanced - the same files open in the web
  app, Studio and the phone app under Examples).
- **The reference** - run `mujard doc docs` for the generated pages: every built-in, board, error and
  register, straight from the compiler.
- **The spec** - `docs/lang/MUJARD_LANGUAGE_SPEC.md` for the whole language.
- **`mujard test file.muj`** - `test` blocks run your program on this PC with a simulated board.
