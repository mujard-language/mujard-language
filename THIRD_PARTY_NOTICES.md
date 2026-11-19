# Third-party notices

This file is the human-readable entry point for third-party license obligations.
The machine-generated, complete inventory is maintained in the platform repository
(`docs/THIRD_PARTY_LICENSES.md` there; regenerate with its license-inventory
generator whenever dependencies change); the pinned
toolchain packages and their licenses are in [`toolchain/index.json`](toolchain/index.json).

## What must travel where

- **Open components** (compiler, runtime, CLI, toolchain definitions): Apache-2.0.
  The only bundled third-party code is the Go module set below; everything else
  is downloaded by the user's toolchain installer, not redistributed.
- **Platform apps** (Android APK, web bundle, desktop app): carry the full
  inventory — the Android app shows it under Settings → Open source licences,
  the web app at `/licenses`.
- **Server-side PHP packages**: not distributed to users; listed for completeness
  in the generated inventory.

## Go module dependencies (tools/mujard)

`golang.org/x/mobile` (with `x/mod`, `x/sync`, `x/tools`) — BSD-3-Clause —
used only by the Android binding build (`tools/mujard/androidbind`). The public
CLI binary does not link it.

## Toolchain packages (downloaded, sha256-pinned)

Full list with SPDX identifiers: `toolchain/index.json`. Summary by license:

| License | Components |
|---|---|
| GPL-3.0-or-later + GCC Runtime Exception | avr-gcc, xtensa-esp-elf, riscv32-esp-elf, arm-none-eabi (xpack & pico-quick), xtensa-lx106-elf |
| GPL-2.0-or-later | esptool, dfu-util, openocd |
| LGPL-2.1 / LGPL-3.0 | Arduino cores (AVR, ESP32, RP2040, SAMD, Renesas, mbed, ESP8266), Adafruit NeoPixel, Arduino Servo |
| MIT / Apache-2.0 / BSD-3-Clause | STM32 core, CMSIS headers, ESP-IDF prebuilt libs, bossac, picotool |
| PSF-2.0 | Python 3.7.2 embeddable (ESP8266 image tools) |

GPL-family toolchains are installed on the user's machine and invoked as separate
programs; Mujard does not copy their code into its own sources. See
`docs/OPEN_SOURCE_ADOPTION_PLAN.md` for the adoption rules these choices follow.
