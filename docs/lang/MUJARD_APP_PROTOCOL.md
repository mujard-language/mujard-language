# Mujard board ↔ app protocol (Phase 9)

Status: **built over USB serial, 2026-10-22** - board runtime (`tools/mujard/lang/runtime/mujard_rt.cpp`, `app_*`),
generator (`tools/mujard/lang/appgen.go`), the web app's protocol parser and dashboard, and
the Android app's link and dashboard screens, all tested against `app-protocol.golden.json`; not yet run on a
real board and phone. The server-side notify path and the internet relay below are still plans. The runtime's code
lives in `mujard_rt.cpp` rather than a separate `mujard_app.h`, because the print filter and the input router have to
sit in the core anyway.

**Over Bluetooth (2026-11-02, language v60):** a program with `use mujard.ble` also sends every app message over the
BLE link (Nordic-UART TX notify, in 20-byte packets the phone joins up to the LF), and the app's lines it receives
over BLE go to the same handler as serial ones. `ble.send` strips 0x1F like `print`, so a relayed line can't pose as a
message. Board side: `app_mirror` / `app_input_line` (`mujard_rt.h`, not on AVR) and `mujard_ble.h`; phone side:
the Android app's BLE packet and link layer, with the same dashboard behind its app link. Compiled for the ESP32
(esp32:esp32:esp32); not yet run on a real board and phone. Web Bluetooth is still the plain line console.

**Held controls (2026-11-02, language v61):** a `control` line's kind field may also be `button`
(`control "Horn", horn, hold`) or `stick-x` / `stick-y` (`control "Drive", x, y, range`: the y axis is the control
right after the x axis, with the same label; the app draws one joystick). Readers that don't know a kind ignore that
control. While a finger is on a held control the app sends its `set` every 200 ms; on release it sends the rest value
(0, or the range end nearest to 0). The board puts a held control back at rest by itself when 500 ms pass without a
`set` for it (`appTick`), and tells the app with `value` - so a car stops when the phone drops away. Tested on the
host (`runtime_test.cpp`: held controls) and in `app-protocol.golden.json`.

**Wi-Fi (contract 2026-11-02):** a board that serves the app protocol over Wi-Fi listens on TCP **port 8266** and
advertises the mDNS service type **`_mujard._tcp`** (its service name is what the app lists; `board:`/`program:` may
ride as TXT records). The byte stream carries the same M1 lines as serial/BLE, LF-ended both ways; the phone sends
`who` right after connecting, and `set`/`answer` come back on the socket. A held control's failsafe must also fire
when the socket closes or goes quiet, like on BLE. Board side: `wifi.serve("name")` + `wifi.client` (language v62,
ESP32/ESP8266/Pico W in `mujard_wifi.h`); phone side: `wifi/WifiLink.kt`, the Terminal screen's Wi-Fi section - NSD
discovery plus a typed address, with a trust prompt before an unknown board can be driven.

This is what the runtime, the web app, the Android app and the desktop
app implement for the Phase 9 built-ins `show`, `plot`, `notify`, `control` and `ask`
(`MUJARD_IMPLEMENTATION_PLAN.md`, Phase 9). Version **M1**.

## Goals

1. **Works on the serial port that is already there.** No second USB interface, no new driver: a Mujard program's
   app messages travel on the same serial line as `print`, so web serial, Android USB, the desktop app, the Mujard
   board's bridge and the simulator all carry them unchanged.
2. **Can't be forged by data.** A program that prints text it received (a GPS sentence, an MQTT message, a web reply)
   must never be able to fake a `notify` or set a `control`. App messages carry a marker byte that the runtime
   removes from everything a program prints.
3. **Plain serial monitors stay usable.** A message is one line; the marker is an invisible control character;
   everything after it is readable text.
4. **The existing plotters keep working.** `plot` lines also read as Arduino Serial Plotter lines once the marker
   is stripped, which is what Android's `Telemetry.parseSample` and the web plotter already parse.

## Framing

```
<US> M1 <kind> <field> <TAB> <field> <TAB> ... <LF>
```

- `<US>` is the byte 0x1F (ASCII "unit separator") at the **start of a line**. Nothing else in Mujard output starts
  with it: the runtime replaces every 0x1F in printed and logged text with `?` (also inside `uart`, `mqtt` and `http`
  text that a program passes to `print`). This is what makes app messages unforgeable from program data.
- `M1` is the protocol version. A reader ignores lines with a version it doesn't know.
- Fields are separated by TAB (0x09). Inside a field, `\` is written `\\`, TAB `\t`, LF `\n`, CR `\r`, and 0x1F
  `\u`; nothing else is escaped. Text is UTF-8.
- One message per line, at most 240 bytes after escaping; the runtime cuts a longer text field and ends it with `…`.
- Lines are ended with LF (a CR before it is ignored), as `print` already does.

Board → app messages are written by the runtime only (never by generated code directly), through the same output
ring buffer as `print`, so they can't stall the scheduler (spec 6.1).

## Board → app

| Kind | Fields | Sent by | App does |
|---|---|---|---|
| `hello` | program name, board display name, protocol version, program hash (8 hex) | at start-up, and again when the app asks (`who`) | resets the dashboard for this program; a changed hash means the controls changed |
| `control` | id, label, kind (`switch` / `slider`), min, max, step, unit, current value | right after `hello`, one per `control` line | draws the control |
| `show` | label, value text, number (or empty), unit, style (`number` / `gauge` / `light`), gauge min, gauge max | each `show` that runs | updates the tile and its time |
| `plot` | `name=value` pairs, one per field | each `plot` that runs | adds a sample to the plotter (same as `name:value` lines) |
| `notify` | message, `urgent` (`0`/`1`), source line | each `notify` that runs (rate-limited, below) | a phone/desktop notification through the notification centre |
| `ask` | id, question, timeout in seconds | `ask` inside a task | shows a yes/no prompt; answers with `answer` |
| `value` | control id, value | when the program itself changes a controlled variable | moves the control |

Notes:
- `show` and `control` labels are the program's text; the app shows them as text, never as markup.
- `show` sends both the display text ("21.5 °C", "on") and, when there is one, the number in display units, so tiles
  and the plotter can use the number.
- `notify` is limited to one per 10 s **per source line** on the board (spec Appendix A, 2), and again by the app and
  the server (per device and per hour), so a loop can't spam a phone.

## App → board

| Kind | Fields | Board does |
|---|---|---|
| `who` | | sends `hello` and the `control` lines again |
| `set` | control id, value | clamps the value to the control's range (spec 11), stores it in the variable at the next pass of the program, answers with `value` |
| `answer` | ask id, `yes` / `no` | wakes the task waiting in `ask`; a late or unknown id is ignored |

App → board messages use the same framing. The runtime's serial reader takes lines that start with 0x1F out of the
input before `on serial.line` / `serial.read_line()` see them, so a program never receives them as text.

Values: a switch's value is `0`/`1`; a slider's is a decimal number in the control's display unit ("40" for 40 % ;
the runtime converts to the base unit). A value that doesn't parse, or a control id that doesn't exist, is ignored.

## Safety rules

- **Clamping is on the board.** Whatever the app sends, a `control` variable stays inside its declared range; a
  switch only takes true/false. The app's own limits are a convenience, not the protection.
- **`set` never runs code.** It only writes the variable; the program reacts on its next pass like with any other
  variable change (`when`, `changed()`).
- **No secrets.** The checker already refuses secrets in `show`, `plot`, `notify` (C323); the protocol has no field
  that could carry one.
- **A halted program** (runtime error, `board.halt`) still answers `who` and sends `hello` with a `halted` flag, but
  ignores `set` and `answer`.
- **Over the internet** (later): the desktop app and the Mujard board forward these lines through the platform's
  sealed serial relay (AES-CTR + HMAC per chunk), never in the clear; `notify` from a board that isn't
  connected to an app goes through the server only when the board has Wi-Fi and HTTPS with certificate checks
  (Phase 8 open item) - never plain HTTP.

## Where each part lives (to build)

| Part | Place |
|---|---|
| Framing, escaping, marker filter in `print`/`log`, the input router, `show`/`plot`/`notify`/`ask`/`control` | `tools/mujard/lang/runtime/mujard_app.h` (+ hooks in `mujard_rt.cpp` print/serial) |
| Generating calls and the `hello`/`control` table | `tools/mujard/lang/codegen.go` (today these give C320) |
| Parser shared by the web app | the web app's app-protocol parser (+ tests with recorded lines) |
| Web dashboard tiles, controls, prompts | the web app's serial page / studio panel |
| Android | the app's telemetry link learns the frames; its dashboard gets show tiles, controls and ask prompts; `notify` goes to the notification centre |
| Desktop | the phone-hub / serial bridge pass the lines through unchanged |
| Simulator | the web simulator shows the same dashboard for a simulated board |

## Tests (to build)

- Runtime host tests: escaping round trip, the marker can't appear in printed text, `set` clamping, `ask` timeout
  and late answers, `notify` rate limit, `hello` after `who`.
- A golden file of recorded lines that the Go, TypeScript and Kotlin parsers must all read the same way.
- Fuzzing the parsers with random lines (no crash, nothing outside the documented fields accepted).
