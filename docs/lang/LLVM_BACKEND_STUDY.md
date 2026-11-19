# Should Mujard get an LLVM back end? (Phase 13 study, 2026-10-29)

**Answer: not now.** Nothing measured so far is limited by compiling through C++, and an LLVM back end would cost
the one thing that makes Mujard practical on real boards: the vendors' own cores and the Arduino library ecosystem.
This study names the conditions that would change the answer, so it is a decision with triggers, not a "no".

## How Mujard compiles today

`.muj` → the Mujard compiler (lexer, parser, checker, generator, all in Go) → readable C++ plus the `mujard_rt`
runtime → each board's own toolchain through arduino-cli (avr-gcc for AVR, Espressif's GCC for ESP32/ESP8266, the
Arm GCC of the RP2040 and nRF/SAMD/STM32 cores). `#line` markers point compiler errors back at `.muj` lines.

## What was measured (2026-10-29, on this PC)

| Question | Measurement | Where the cost is |
|---|---|---|
| Is the generated code too big? | Blink on a Nano was 16,058 bytes; after `mujard_config.h` switches off the unused runtime it is 8,448 (-47 %), RAM 1,005 → 442 bytes | the runtime linking in unused features - fixed in the runtime, no back end involved |
| Is pin access too slow? | `digitalWrite` on AVR looks up tables on every call; outputs now write their port register directly | the core's API - fixed in the runtime |
| Is building too slow? | the Mujard compiler itself: milliseconds; AVR builds: tens of seconds; a cold ESP32 build: over 30 minutes on this loaded PC | the vendor toolchain and core compiling (ESP-IDF), which an LLVM back end would not remove: the core and libraries still have to be compiled |
| Does C++ block a language feature? | every built-in of the spec, the app protocol, tests on the PC, the 52 library descriptions: all generated as C++; the build matrix is green on AVR, ESP32, ESP8266 and RP2040 | - |
| Do C++ libraries work? | 52 described Arduino libraries checked and built from Mujard, against the real installed headers | this works *because* the output is C++ |

## What an LLVM back end would cost

- **Library interop.** Mujard programs call Arduino C++ libraries (NeoPixel, U8g2, ArduinoJson...). Emitting LLVM IR
  directly still needs every library compiled by a C++ front end (clang) and linked with matching ABIs, per core.
  Today the board's own GCC does all of it.
- **Targets.** Upstream LLVM covers Arm and RISC-V well and has an AVR back end; Xtensa (the ESP32 and ESP8266) is
  mainly Espressif's own fork. The vendors ship and support GCC for their cores; each LLVM target would be a
  toolchain Mujard maintains itself.
- **Debugging and errors.** `#line` already maps C++ errors and debug information to `.muj` lines; a new back end
  would have to rebuild that path.
- **Two back ends to keep equal.** Every built-in would be implemented and tested twice.

## When to look again (the triggers)

1. **A target with no C++ toolchain** that Mujard should support (Phase 19 may bring one).
2. **A guarantee C++ can't give**: for example cycle-exact timing analysis or stack-depth proofs for Phase 18's
   deadlines, where the back end must know what instructions run.
3. **The source-level debugger (Phase 20)** needs debug information richer than `#line` can carry.
4. **Benchmarks** show a real program bound by code the C++ route makes worse, after the runtime work above.

If one is met, the plan's rule holds: source compatibility first. `tools/mujard/lang/testdata/compat/2026` (85 frozen
programs, `TestEdition2026StillCompiles`) is the gate any new back end must pass unchanged, and the alternative is
compared against the C++ route on the build matrix (size, speed, build time) before it is offered.
