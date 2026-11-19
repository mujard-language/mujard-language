# Roadmap

Public milestones for the Mujard language. "Done" means implemented and tested in this
repository; board-level proof is tracked per board in `tools/mujard/lang/capabilities.json`
(`mujard boards caps <board>` shows the evidence level).

## Done

- [x] Formal specification (`docs/lang/MUJARD_LANGUAGE_SPEC.md`) with grammar, types,
      units, pins, events and built-ins
- [x] Compiler: lexer, parser, syntax tree, checker (types, units, pins, events)
- [x] C++ generator + `mujard_rt` runtime
- [x] `mujard` CLI: `check`, `fmt`, `gen`, `new`, `test`, `repl`, `lsp`, `build`, `boards`,
      `builds`, `packages`, `webhooks`, plugins
- [x] Language server (WASM build for editors, `mujard lsp` for desktop editors)
- [x] Testing framework: `test` blocks, `pretend`, `advance` on the host simulator
- [x] Board↔app protocol M1: `show`, `plot`, `notify`, `control`, `ask` over serial, BLE
      and Wi-Fi
- [x] Standard libraries and Arduino library bindings (`.mlib` descriptions)
- [x] Safety: profiles, recursion warning, pin and timer ownership
- [x] Real-time semantics (deterministic order, run-to-completion, interrupt-safe vars)
- [x] Capability evidence model: hardware / build / backend / none per board
- [x] 18 build targets across AVR, ESP32, ESP8266, RP2040, SAMD, megaAVR, Renesas, mbed
      and STM32 families

## Next

- [ ] **M1 completion** — the full sample set running on real boards (Uno, ESP32) with
      owner-verified runs; capability evidence promoted from *build* to *hardware*
- [ ] **M2** — private beta of the language inside the platform apps
- [ ] **M3** — every documented built-in working end to end (`notify`, `show`, `control`,
      `ask`, secrets) on real Wi-Fi hardware and phone
- [ ] **M4** — Mujard 1.0: documentation freeze, courses, public beta, release
- [ ] More targets: additional board families and chips as the matrix grows
- [ ] Post-1.0: on-board REPL, source debugger on hardware, package registry growth,
      compiler speed work

## Under consideration

- Alternative compiler back ends (an LLVM study exists; decided "not now" until a
  trigger — code size, a target C++ can't reach, or compile-time pain — actually fires)

## How priorities are set

Language changes follow the RFC process: propose, discuss, then land the change in all
three homes together — parser, specification EBNF and golden tests. See
[CONTRIBUTING.md](CONTRIBUTING.md).
