# Built-ins

Every name a program can use. `[lib]` names need `use mujard.<name>` first; `[test]` names only work inside `test` blocks.

| Name | Kind | Signature |
|---|---|---|
| `BatteryType` | object | `BatteryType` |
| `Can` | function `use mujard.can` | `Can(tx, rx, rate?) -> can` |
| `Level` | object | `Level` |
| `Queue` | function `use mujard.queue` | `Queue(of, capacity) -> queue` |
| `ResetReason` | object | `ResetReason` |
| `Servo` | function `use mujard.servo` | `Servo(pin) -> servo` |
| `Timer` | function `use mujard.timer` | `Timer(...)` |
| `Uart` | function `use mujard.uart` | `Uart(rx, tx, speed?) -> uart` |
| `Ultrasonic` | function `use mujard.ultrasonic` | `Ultrasonic(trig, echo) -> ultrasonic` |
| `aa_x3` | constant | `BatteryType` |
| `aa_x4` | constant | `BatteryType` |
| `abs` | function | `abs(x)` |
| `advance` | function `[test]` | `advance(time)` |
| `answer` | function `[test]` | `answer(value)` |
| `ask` | function | `ask(question, timeout?) -> bool?` |
| `battery_on` | function | `battery_on(pin, divider?, type?) -> battery` |
| `between` | function | `between(x, range) -> bool` |
| `ble` | object `use mujard.ble` | `ble` |
| `board` | object | `board` |
| `brown_out` | function `[test]` | `brown_out()` |
| `brownout` | constant | `ResetReason` |
| `changed` | function | `changed(value) -> bool` |
| `clamp` | function | `clamp(x, range)` |
| `clock` | object | `clock` |
| `crash` | constant | `ResetReason` |
| `ethernet` | object `use mujard.ethernet` | `ethernet` |
| `falling` | function | `falling(value) -> bool` |
| `forever` | constant | `time` |
| `gauge` | function | `gauge(range) -> any` |
| `halted` | function `[test]` | `halted() -> bool` |
| `high` | constant | `Level` |
| `http` | object | `http` |
| `hysteresis` | function | `hysteresis(value, on, off) -> bool` |
| `i2c` | object `use mujard.i2c` | `i2c` |
| `lipo_1s` | constant | `BatteryType` |
| `log` | object | `log` |
| `low` | constant | `Level` |
| `math` | object `use mujard.math` | `math` |
| `max` | function | `max(a, b)` |
| `min` | function | `min(a, b)` |
| `mqtt` | object `use mujard.mqtt` | `mqtt` |
| `notified` | function `[test]` | `notified(text) -> bool` |
| `notify` | function | `notify(message, urgent?)` |
| `now` | function | `now() -> time` |
| `peak` | function | `peak(value, window)` |
| `plot` | function | `plot(value)` |
| `power` | object | `power` |
| `power_cut` | function `[test]` | `power_cut()` |
| `power_on` | constant | `ResetReason` |
| `printed` | function `[test]` | `printed(text) -> bool` |
| `random` | object | `random(range) - the random object` |
| `regs` | object `use mujard.regs` | `regs` |
| `reset_button` | constant | `ResetReason` |
| `rising` | function | `rising(value) -> bool` |
| `round` | function | `round(x, decimals?)` |
| `rtc` | object `use mujard.rtc` | `rtc` |
| `scale` | function | `scale(x, from, to)` |
| `secret` | function | `secret(name) -> secret` |
| `serial` | object | `serial(speed) - the serial object` |
| `shift` | object `use mujard.shift` | `shift` |
| `show` | function | `show(label, value, as?)` |
| `since` | function | `since(time) -> time` |
| `smooth` | function | `smooth(value, count)` |
| `software` | constant | `ResetReason` |
| `spi` | object `use mujard.spi` | `spi` |
| `stopped` | function `[test]` | `stopped() -> bool` |
| `stopwatch` | function | `stopwatch() -> stopwatch` |
| `store` | object | `store` |
| `timezone` | constant | `any` |
| `usb` | object `use mujard.usb` | `usb` |
| `wake` | constant | `ResetReason` |
| `watchdog` | function | `watchdog(time)` |
| `wifi` | object | `wifi` |
