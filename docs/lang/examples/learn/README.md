# Mujard examples: Basics to Advanced

The programs people learn from. The web app, Mujard Studio and the Android app all show them under **Examples**
(Project menu; in Studio also the command palette, "Open an example"). Each opens as a new project of its own.

| Level | Folder | What it covers |
|---|---|---|
| 1. Basics | `1-basics` | print, blink, fade, buttons, knobs |
| 2. Next steps | `2-next-steps` | if/elif/else, when, tasks, counting, button gestures, pin groups, random, sound |
| 3. Sensors and motors | `3-sensors-and-motors` | servo, ultrasonic, TMP36, thermostat, units, encoder, 74HC595, I2C |
| 4. Bigger programs | `4-bigger-programs` | serial commands, functions, lists, enum/struct/match, stopwatch, task control, smart signals, import, bits |
| 5. Advanced | `5-advanced` | safety, test blocks, saved var, interrupts, multitasking, ESP32, Pico, Arduino libraries, a full project |

## Adding or changing an example

1. Write the `.muj` file. Line 1 is `# <title>`, line 2 is `# Level N - <level name>`, then what it does, what you
   learn and what you need (wiring). The main file of a project is `sketch.muj`; a module sits next to it.
2. List it in `examples.json` (same title as line 1).
3. In `tools/mujard`: `MUJARD_GOLDEN=1 go test ./lang -run Learn` - it compiles every example all the way to C++
   (an example that only passes `mujard check` is not enough) and refreshes the bundled example lists the apps
   ship. Without `MUJARD_GOLDEN` the same test fails when those are stale.

Only use features the code generator supports today: `notify`, `show`, `control`, `ask`, MQTT, `store.save`,
`battery_on`, `Uart`, `i2c.read`, `.bytes`, `board.reset_reason`, Wi-Fi events and `at 07:30` still say
"isn't available yet" (C320). Test blocks run with `mujard test file.muj` (needs a host C++ compiler on the PATH,
or `MUJARD_CXX` pointing at one).
