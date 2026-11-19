# History of Mujard

How the language got here — from first sketches to the compiler, runtime and toolchain
that exist today. Dates are the project's own records; the full design rationale lives in
the specification and the reference docs.

## 2025 — planning

* **January** — Mujard is conceived: a small language for microcontrollers, "bare code,
  every board", compiled to C++ and built by the same toolchains Arduino uses. The goal
  is hardware programming without the ceremony — no IDE lock-in, no hidden magic, real
  speed.
* **Through the year** — the design is iterated privately: the surface syntax, the unit
  system, the event model, the safety rules, and how a `.muj` program maps onto real
  chips. The decision is made early that the language must be *checkable* — errors caught
  before a board ever sees the code — and that its semantics must be identical on a
  simulator, a desktop build and real silicon.

## 2026 — the specification and the first compiler

* **18 Oct** — the formal specification lands (`docs/lang/MUJARD_LANGUAGE_SPEC.md`):
  grammar, types, units, pins, events, built-ins and the runtime contract, plus thirty
  sample programs that define what "correct" means. The first implementation phase is
  done the same day, along with the first two compiler pieces — the lexer and the parser —
  and the `mujard` command line skeleton.
* **20 Oct** — spec revision 2 fixes contradictions found in review (unit widths, `mod`,
  `expect`, named options).
* **21 Oct** — server and build integration is verified end to end: a `.muj` program
  builds through every build path and produces the same firmware as a local build.
* **22 Oct** — the checker is finished: types, units, pins and events are all verified,
  closing the exceptions from the spec review. The board↔app protocol (M1) is built over
  serial the same day. A formal-grammar and security pass is run at the owner's request:
  "formal, secure, powerful, all hardware".
* **24–25 Oct** — live checking lands in the editors; documentation generation starts.
* **27 Oct** — the testing framework (`test`, `pretend`, `advance`) is done: all forty
  sample programs have passing tests on the host simulator.
* **29 Oct** — the AI assistant learns Mujard (guide + compiler-verified examples), and
  the first real-board milestone runs: Blink and the button sample execute on a physical
  Nano. A "gap pass" makes the live checker report everything the generator refuses, and
  surfaces the secret names a program needs.
* **30 Oct** — a big landing: standard libraries and Arduino bindings (every description
  built against its real library), reliability work (generated code validated against the
  build rules, timing and memory verified for AVR), safety (profiles, recursion warning,
  pin and timer ownership), the hardware capability model, drivers (`needs`, driver pin
  parameters), real-time semantics for one core, and simulator debugging with timing and
  overrun measurement.

## November 2026 — breadth

* **2 Nov** — the app protocol grows: Bluetooth LE transport (language v60), held
  controls with on-board failsafes (v61), and Wi-Fi serving with mDNS discovery (v62).
* **7–8 Nov** — the board matrix expands to seventeen board families and every `.muj`
  language board builds through the direct backend — all eighteen targets. A lesson is
  recorded the hard way: `on core N` shipped without its grammar update and five tests
  failed until the spec caught up — the "three homes" rule (parser, spec EBNF, goldens
  change together) is reinforced.

## Where it stands

The language is **pre-1.0 and implemented**: lexer, parser, checker, C++ generator,
`mujard_rt` runtime, CLI, language server, WASM build, testing framework, app protocol —
all in this repository, all tested. Capability is proven honestly per board (`mujard
boards caps <board>`): *hardware* means it ran on a real board, *build* means a real build
succeeded, *backend* means generated but not yet built for it, *none* means the board
lacks it.

The platform around it — accounts, cloud builds, the web/Android/desktop apps — is
proprietary and not part of this repository. See [ROADMAP.md](ROADMAP.md) for what comes
next and [CONTRIBUTING.md](CONTRIBUTING.md) for how to help.
