# Mujard documentation

Everything you need to write, check and understand Mujard programs — in this order:

## Start here

| Doc | What it gives you |
|---|---|
| [lang/tutorials/mujard-in-15-minutes.md](lang/tutorials/mujard-in-15-minutes.md) | The fastest path from zero to a blinking LED |
| [lang/tutorials/from-arduino-to-mujard.md](lang/tutorials/from-arduino-to-mujard.md) | Coming from Arduino/C++? This maps the concepts over |
| [lang/examples/README.md](lang/examples/README.md) + [examples/learn/](lang/examples/learn/) | 49 guided examples in five levels — every one compiles |

## The language

| Doc | What it gives you |
|---|---|
| [lang/MUJARD_LANGUAGE_SPEC.md](lang/MUJARD_LANGUAGE_SPEC.md) | **The normative specification** — grammar, types, units, pins, events. When anything disagrees, this wins |
| [lang/README.md](lang/README.md) | All 100 built-ins: what each does, examples, and the Arduino equivalent |
| [lang/reference/](lang/reference/) | Deep reference: boards, capabilities, errors, keywords & units, libraries, objects, registers |
| [lang/reference/registers.md](lang/reference/registers.md) | Per-chip register maps (RP2040, ESP32 family, SAMD, STM32, AVR, Renesas, nRF52) |

## The board ↔ app protocol

| Doc | What it gives you |
|---|---|
| [lang/MUJARD_APP_PROTOCOL.md](lang/MUJARD_APP_PROTOCOL.md) | The M1 wire protocol: `show`, `plot`, `notify`, `control`, `ask` over serial, BLE and Wi-Fi |
| [lang/app-protocol.golden.json](lang/app-protocol.golden.json) | Recorded protocol lines every reader is tested against |

## For tool builders

| File | What it gives you |
|---|---|
| [lang/editor-info.json](lang/editor-info.json) | Keywords, built-ins, completions and hover data — what editors and the WASM checker consume |
| [lang/capabilities.json](../tools/mujard/lang/capabilities.json) | Per-board, per-feature evidence: *hardware* / *build* / *backend* / *none* |
| [lang/LLVM_BACKEND_STUDY.md](lang/LLVM_BACKEND_STUDY.md) | Why an LLVM backend is "not now" — the triggers that would change that |

## Trying it

```sh
cd tools/mujard
go build -o mujard .
./mujard check path/to/program.muj   # parse + check, compiles nothing
./mujard test path/to/program.muj    # run its test blocks on a simulated board
./mujard fmt path/to/program.muj     # canonical formatting
```

See [../CONTRIBUTING.md](../CONTRIBUTING.md) for how changes land.
