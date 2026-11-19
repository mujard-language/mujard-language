# From Arduino to Mujard

For people who already write Arduino sketches: what changes, what maps to what, and what you get back.

## The big change: no `setup()`, no `loop()`, no `delay()`

An Arduino sketch is one loop you run as fast as possible, with `delay()` freezing everything. A Mujard
program is a set of **timed and event blocks** that share the board:

```cpp
// Arduino
void loop() {
    digitalWrite(LED_BUILTIN, !digitalRead(LED_BUILTIN));
    delay(500);
}
```

```muj
# Mujard
every 500 ms:
    led.toggle()
```

`every 500 ms:` schedules itself. Other blocks - a button handler, a sensor poll, a `task` - keep running on
their own schedule; nothing waits on anything else. `delay()`'s role is taken by `every`, `after`, `wait`
(inside tasks) and `when`.

## Side by side

```cpp
// Arduino
int led = 13, btn = 2;
bool on = false;
void setup() {
    pinMode(led, OUTPUT);
    pinMode(btn, INPUT_PULLUP);
    Serial.begin(9600);
}
void loop() {
    if (digitalRead(btn) == LOW) {
        on = !on;
        digitalWrite(led, on);
        Serial.println("Click!");
        delay(50); // debounce
    }
}
```

```muj
# Mujard
board uno

led    = pin 13 as output
button = pin 2 as input pullup

on button.pressed:
    led.toggle()
    print "Click!"
```

The debounce, the pull-up, the serial setup and the state variable are all handled - `button.pressed` fires
once per real press.

## Translation table

| Arduino | Mujard |
|---|---|
| `pinMode(13, OUTPUT)` | `led = pin 13 as output` |
| `pinMode(2, INPUT_PULLUP)` | `button = pin 2 as input pullup` |
| `digitalWrite(led, HIGH)` | `led.on()` / `led.write(true)` / `led.toggle()` |
| `digitalRead(btn)` | `button.read()` - or better, `on button.pressed:` / `released:` |
| `analogRead(A0)` | `pot = analog A0`, then `pot.value` / `pot.percent` |
| `analogWrite(9, v)` | `led = pin 9 as pwm`, then `led.write(50 %)` / `led.fade(100 %, 2 s)` |
| `delay(1000)` | `every 1 s:` for repeating work, `wait 1 s` inside a task, `after 1 s:` once |
| `millis()` | `board.uptime` |
| `Serial.begin(9600); Serial.print(x)` | `print x` |
| `attachInterrupt(...)` | `on pin.rising:` / `on pin.falling:` (add `fast` for a true ISR) |
| `Servo.attach(9); .write(a)` | `use mujard.servo` + `arm = Servo(pin 9)`, `arm.move(90 deg)` |
| `Wire.begin(); write/read` | `use mujard.i2c` + `i2c.write(...)`, `i2c.read(...)`, `i2c.scan()` |
| `SPI.begin(); transfer` | `use mujard.spi` + `dev = spi.device(cs)`, `dev.transfer(...)` |
| `EEPROM.write/read` | `saved var x = 0` (auto save/load) or `store.save("k", v)` / `store.load("k")` |
| `#include <Lib.h>` | `use "Lib Name"` - checked calls from the library description |
| register poking (`PORTB = ...`) | `use mujard.regs` + `regs.PORTB.write(v)` - typed, checked |
| bare `cpp:` escape | `cpp:` blocks when you need raw C++ (unchecked by design) |

## What the checker catches that the Arduino IDE doesn't

- **Wrong pin for the job**: PWM on a non-PWM pin, analog on a digital-only pin, a strapping or flash pin on
  ESP32 - errors and warnings at `mujard check` time, per board.
- **Pin collisions**: two devices on one pin, a library claiming the I2C/SPI/Servo-timer pins, a `regs` write
  onto a port a declared pin uses.
- **Board reality**: RAM estimate before the runtime's share, flash-pinned Wi-Fi ADC pins, serial-port limits
  on Uno.
- **Type and unit slips**: `500 ms` where `9600 baud` belongs, text in a number slot, a secret printed.
- **Names**: a misspelt register, field or method is an error, not a silent no-op.

## What Mujard adds on top

- **Units**: `500 ms`, `9600 baud`, `50 %`, `90 deg`, `3.3 V` are typed values - `scale(x, 0 %..100 %, 0..)` maps them.
- **Tasks**: `task name:` blocks run cooperatively next to your events - real multitasking without threads.
- **The app**: `show`, `plot`, `notify`, `ask`, `control` - the phone app / Studio shows values, graphs, controls and notifications from the same program.
- **Secrets**: `secret("home_wifi")` - credentials live in the build's secrets, never in source.
- **Tests**: `test` blocks + `mujard test` run the program on your PC against a simulated board.
- **Safety**: a runtime error or `board.halt` locks outputs into a safe state instead of leaving them live.
- **Checked everything**: the generated C++ is the output of a real language, not a macro layer - the file at `mujard gen`'s output is what actually compiles.

## Moving a sketch over, step by step

1. Keep the wiring; change `board` to your board.
2. Every `pinMode` becomes a device declaration (`= pin N as output` / `as input` / `as pwm` / `analog N`).
3. Everything in `setup()` goes to `at start:`.
4. Split `loop()` by *what* it does: periodic work → `every`, input reactions → `on ...`, long sequences → `task`.
5. `delay()` never survives translation - the block *is* the schedule.
6. Libraries: `use "Name"` for described Arduino libs, `use mujard.x` for the built-ins.
7. `mujard check` until it's quiet; the warnings are usually real.
8. `mujard run file.muj --port ...` - same board, same wiring.

## Worth knowing

- `.muj` and `.mjd` are both accepted.
- `print` is a statement (no parens needed); `var x = `, `const`, `saved var`, `fn`, `struct`, `enum`, `match`, `import` are all in the spec.
- `mujard doc docs` generates the full reference from the compiler.
- Real register access is `use mujard.regs` - `regs.PORTB.write(0xFF)` - generated per chip, checked against declared pins.
