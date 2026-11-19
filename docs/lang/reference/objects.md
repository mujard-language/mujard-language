# Objects and their members

## `BatteryType`

| Member | Kind | Signature |
|---|---|---|
| `aa_x3` | field | `aa_x3: BatteryType` |
| `aa_x4` | field | `aa_x4: BatteryType` |
| `lipo_1s` | field | `lipo_1s: BatteryType` |

## `Level`

| Member | Kind | Signature |
|---|---|---|
| `high` | field | `high: Level` |
| `low` | field | `low: Level` |

## `ResetReason`

| Member | Kind | Signature |
|---|---|---|
| `brownout` | field | `brownout: ResetReason` |
| `crash` | field | `crash: ResetReason` |
| `power_on` | field | `power_on: ResetReason` |
| `reset_button` | field | `reset_button: ResetReason` |
| `software` | field | `software: ResetReason` |
| `wake` | field | `wake: ResetReason` |
| `watchdog` | field | `watchdog: ResetReason` |

## `analog`

| Member | Kind | Signature |
|---|---|---|
| `percent` | field | `percent: ratio` |
| `pretend` | method | `pretend(value)` |
| `reference` | method | `reference(voltage)` |
| `value` | field | `value: int` |
| `volts` | field | `volts: voltage` |

## `battery`

| Member | Kind | Signature |
|---|---|---|
| `percent` | field | `percent: ratio` |
| `pretend` | method | `pretend(value)` |
| `volts` | field | `volts: voltage` |

## `ble`

| Member | Kind | Signature |
|---|---|---|
| `connected` | field | `connected: bool` |
| `line` | event | `line` |
| `pretend` | method | `pretend(line)` |
| `read_line` | method | `read_line(...)` |
| `send` | method | `send(text) -> bool` |
| `start` | method | `start(name)` |

## `block`

| Member | Kind | Signature |
|---|---|---|
| `restart` | method | `restart(...)` |
| `running` | field | `running: bool` |
| `start` | method | `start(...)` |
| `stop` | method | `stop(...)` |

## `board`

| Member | Kind | Signature |
|---|---|---|
| `chip` | field | `chip: text` |
| `cores` | field | `cores: int` |
| `dma` | field | `dma: bool` |
| `dropped_events` | field | `dropped_events: int` |
| `ethernet` | field | `ethernet: bool` |
| `free_memory` | field | `free_memory: memory` |
| `halt` | method | `halt(reason)` |
| `id` | field | `id: text` |
| `name` | field | `name: text` |
| `overruns` | field | `overruns: int` |
| `reset_reason` | field | `reset_reason: ResetReason` |
| `restart` | method | `restart(...)` |
| `rtos` | field | `rtos: text` |
| `timers` | field | `timers: int` |
| `uptime` | field | `uptime: time` |
| `voltage` | field | `voltage: voltage?` |

## `can`

| Member | Kind | Signature |
|---|---|---|
| `message` | event | `message` |
| `pretend` | method | `pretend(id, data)` |
| `read` | method | `read(...)` |
| `send` | method | `send(id, data) -> bool` |

## `can_message`

| Member | Kind | Signature |
|---|---|---|
| `data` | field | `data: list` |
| `id` | field | `id: int` |
| `rtr` | field | `rtr: bool` |

## `clock`

| Member | Kind | Signature |
|---|---|---|
| `date` | field | `date: date?` |
| `source` | method | `source(source)` |
| `sync` | method | `sync(timezone?)` |
| `time` | field | `time: time_of_day?` |

## `date`

| Member | Kind | Signature |
|---|---|---|
| `day` | field | `day: int` |
| `month` | field | `month: int` |
| `weekday` | field | `weekday: int` |
| `year` | field | `year: int` |

## `ethernet`

| Member | Kind | Signature |
|---|---|---|
| `connected` | field | `connected: bool` |
| `ip` | field | `ip: text` |
| `joined` | event | `joined` |
| `lost` | event | `lost` |
| `pretend` | method | `pretend(up)` |
| `start` | method | `start(...)` |

## `group`

| Member | Kind | Signature |
|---|---|---|
| `length` | field | `length: int` |
| `write` | method | `write(bits)` |

## `http`

| Member | Kind | Signature |
|---|---|---|
| `get` | method | `get(url) -> http_reply` |
| `post` | method | `post(url, data) -> http_reply` |

## `http_reply`

| Member | Kind | Signature |
|---|---|---|
| `json` | field | `json: json` |
| `ok` | field | `ok: bool` |
| `status` | field | `status: int` |
| `text` | field | `text: text` |

## `i2c`

| Member | Kind | Signature |
|---|---|---|
| `pretend` | method | `pretend(address, reply)` |
| `read` | method | `read(address, count) -> list?` |
| `scan` | method | `scan(...)` |
| `write` | method | `write(address, data)` |

## `input`

| Member | Kind | Signature |
|---|---|---|
| `changed` | event | `changed` |
| `clicked` | event | `clicked` |
| `debounce` | method | `debounce(time)` |
| `double_clicked` | event | `double_clicked` |
| `falling` | event | `falling` |
| `held` | event | `held(time)` |
| `measure_pulse` | method | `measure_pulse(level?, timeout?) -> time?` |
| `pressed` | event | `pressed` |
| `pretend` | method | `pretend(value)` |
| `read` | method | `read(...)` |
| `released` | event | `released` |
| `rising` | event | `rising` |

## `json`

| Member | Kind | Signature |
|---|---|---|
| `ok` | field | `ok: bool` |
| `to_number` | method | `to_number(...)` |

## `list`

| Member | Kind | Signature |
|---|---|---|
| `add` | method | `add(item)` |
| `clear` | method | `clear(...)` |
| `length` | field | `length: int` |
| `remove` | method | `remove(index)` |

## `log`

| Member | Kind | Signature |
|---|---|---|
| `error` | method | `error(message)` |
| `info` | method | `info(message)` |
| `warn` | method | `warn(message)` |

## `math`

| Member | Kind | Signature |
|---|---|---|
| `E` | field | `E: float` |
| `PI` | field | `PI: float` |
| `atan2` | method | `atan2(y, x) -> angle` |
| `cos` | method | `cos(angle) -> float` |
| `log` | method | `log(x) -> float` |
| `pow` | method | `pow(x, y) -> float` |
| `sin` | method | `sin(angle) -> float` |
| `sqrt` | method | `sqrt(x) -> float` |
| `tan` | method | `tan(angle) -> float` |

## `mqtt`

| Member | Kind | Signature |
|---|---|---|
| `connect` | method | `connect(broker, user?, password?)` |
| `message` | event | `message` |
| `publish` | method | `publish(topic, value)` |
| `subscribe` | method | `subscribe(topic)` |

## `mqtt_message`

| Member | Kind | Signature |
|---|---|---|
| `text` | field | `text: text` |
| `topic` | field | `topic: text` |

## `number`

| Member | Kind | Signature |
|---|---|---|
| `bit` | method | `bit(n) -> bool` |
| `bytes` | field | `bytes: list` |
| `clear_bit` | method | `clear_bit(n) -> int` |
| `set_bit` | method | `set_bit(n) -> int` |
| `toggle_bit` | method | `toggle_bit(n) -> int` |

## `output`

| Member | Kind | Signature |
|---|---|---|
| `blink` | method | `blink(rate, for?, on?)` |
| `melody` | method | `melody(notes, tempo?)` |
| `off` | method | `off(...)` |
| `on` | method | `on(...)` |
| `pattern` | method | `pattern(pattern, step?, repeat?)` |
| `pretend` | method | `pretend(value)` |
| `pulse` | method | `pulse(length)` |
| `read` | method | `read(...)` |
| `stop` | method | `stop(...)` |
| `toggle` | method | `toggle(...)` |
| `tone` | method | `tone(pitch, length?)` |
| `write` | method | `write(value)` |

## `power`

| Member | Kind | Signature |
|---|---|---|
| `deep_sleep` | method | `deep_sleep(time)` |
| `sleep` | method | `sleep(time, wake_on?)` |

## `pwm`

| Member | Kind | Signature |
|---|---|---|
| `fade` | method | `fade(to, over)` |
| `frequency` | method | `frequency(frequency)` |
| `off` | method | `off(...)` |
| `on` | method | `on(...)` |
| `pretend` | method | `pretend(value)` |
| `stop` | method | `stop(...)` |
| `write` | method | `write(level)` |

## `queue`

| Member | Kind | Signature |
|---|---|---|
| `available` | field | `available: bool` |
| `full` | field | `full: bool` |
| `n` | field | `n: int` |
| `pop` | method | `pop(...)` |
| `push` | method | `push(value) -> bool` |

## `random`

| Member | Kind | Signature |
|---|---|---|
| `pick` | method | `pick(list)` |
| `seed` | method | `seed(seed?)` |

## `regs_reg`

| Member | Kind | Signature |
|---|---|---|
| `clear` | method | `clear(bits)` |
| `read` | method | `read(...)` |
| `set` | method | `set(bits)` |
| `write` | method | `write(value)` |

## `rtc`

| Member | Kind | Signature |
|---|---|---|
| `set` | method | `set(year, month, day, hour, minute, second?)` |

## `serial`

| Member | Kind | Signature |
|---|---|---|
| `line` | event | `line` |
| `pretend` | method | `pretend(line)` |
| `read_line` | method | `read_line(...)` |

## `servo`

| Member | Kind | Signature |
|---|---|---|
| `angle` | field | `angle: angle` |
| `move` | method | `move(angle, over?)` |
| `pretend` | method | `pretend(value)` |

## `shift`

| Member | Kind | Signature |
|---|---|---|
| `out_register` | method | `out_register(data, clock, latch) -> shift_register` |

## `shift_register`

| Member | Kind | Signature |
|---|---|---|
| `write` | method | `write(bits)` |

## `spi`

| Member | Kind | Signature |
|---|---|---|
| `device` | method | `device(cs, speed?, mode?) -> spi_device` |

## `spi_device`

| Member | Kind | Signature |
|---|---|---|
| `transfer` | method | `transfer(data) -> list` |

## `stopwatch`

| Member | Kind | Signature |
|---|---|---|
| `elapsed` | field | `elapsed: time` |
| `reset` | method | `reset(...)` |
| `running` | field | `running: bool` |
| `start` | method | `start(...)` |
| `stop` | method | `stop(...)` |

## `store`

| Member | Kind | Signature |
|---|---|---|
| `committed` | field | `committed: bool` |
| `failed` | field | `failed: int` |
| `last_commit` | field | `last_commit: bool` |
| `load` | method | `load(name, default?)` |
| `pending` | field | `pending: bool` |
| `save` | method | `save(name, value)` |

## `text`

| Member | Kind | Signature |
|---|---|---|
| `contains` | method | `contains(part) -> bool` |
| `ends_with` | method | `ends_with(part) -> bool` |
| `length` | field | `length: int` |
| `lower` | method | `lower(...)` |
| `split` | method | `split(separator) -> list` |
| `starts_with` | method | `starts_with(part) -> bool` |
| `to_number` | method | `to_number(...)` |
| `upper` | method | `upper(...)` |

## `time_of_day`

| Member | Kind | Signature |
|---|---|---|
| `hour` | field | `hour: int` |
| `minute` | field | `minute: int` |
| `second` | field | `second: int` |

## `timer`

| Member | Kind | Signature |
|---|---|---|
| `elapsed` | method | `elapsed(...)` |
| `every` | method | `every(period)` |
| `freq` | method | `freq(hz)` |
| `once` | method | `once(period)` |
| `pretend` | method | `pretend(...)` |
| `start` | method | `start(...)` |
| `stop` | method | `stop(...)` |
| `tick` | event | `tick` |

## `uart`

| Member | Kind | Signature |
|---|---|---|
| `line` | event | `line` |
| `pretend` | method | `pretend(line)` |
| `read_line` | method | `read_line(...)` |
| `write` | method | `write(text)` |

## `ultrasonic`

| Member | Kind | Signature |
|---|---|---|
| `distance` | method | `distance(...)` |
| `pretend` | method | `pretend(distance)` |

## `usb`

| Member | Kind | Signature |
|---|---|---|
| `click` | method | `click(...)` |
| `connected` | field | `connected: bool` |
| `key` | method | `key(key) -> bool` |
| `mouse_move` | method | `mouse_move(x, y) -> bool` |
| `pretend` | method | `pretend(up)` |
| `type` | method | `type(text) -> bool` |

## `wifi`

| Member | Kind | Signature |
|---|---|---|
| `client` | field | `client: bool` |
| `connect` | method | `connect(network)` |
| `connected` | field | `connected: bool` |
| `joined` | event | `joined` |
| `lost` | event | `lost` |
| `serve` | method | `serve(name, pin?)` |

