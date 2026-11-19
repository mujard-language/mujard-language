# mujard - the Mujard command line

One small program (Go, standard library only) that runs the whole system from a terminal:

* **the cloud API** - build sketches, list boards, check your usage, manage webhooks (`mujard build nano blink.ino`)
* **this PC's USB ports** - flash and run boards through the desktop app's engine (`mujard flash firmware.hex --board uno --port COM5`)
* **the Mujard language** - check, format, test, run and inspect `.muj` programs (`mujard check blink.muj`, `mujard test blink.muj`)
* **your own commands** - any executable named `mujard-<name>` in the commands folder becomes `mujard <name>`

## Commands

| Command | What it does |
|---|---|
| `mujard login` | save your developer key (`mjd_live_...`/`mjd_test_...`) and API address in `~/.mujard/config.json` |
| `mujard logout` / `mujard config` | forget the key / show the saved settings (key masked) |
| `mujard config path` / `get api\|key` / `set api <url>` | where the settings live / read one (the key stays masked) / change the server address. The key itself is only set by `mujard login` |
| `mujard auth status` / `mujard auth test` | is a key set, and where it came from (no network) / does the server accept it. `--format json` gives the `mujard-cli.v1` document |
| `mujard me` | your key: scopes, limits, this month's spend |
| `mujard usage` | requests and coins this month |
| `mujard boards` / `mujard boards list` | the boards a build can target (`--format json` for the raw document) |
| `mujard board info <id>` / `mujard boards <id>` | one build target's detail |
| `mujard build <board> [sketch\|folder]` | compile on the server: `mujard build nano blink.ino`. No path = the `.ino` in the current folder. Polls until done, downloads the firmware to `<sketch>.hex`/`.bin` (`--out`), prints compiler errors `file:line`, shows cost. `--wait 10m`, `--format json`. `--backend direct` builds locally with `mujard-build` (no arduino-cli, no key); `--backend auto` tries it then falls back. `mujard build <file.muj>` compiles the Mujard program first (`--show-cpp` writes the generated C++ to `<name>.gen/`; with `arduino-cli` installed it builds locally, else on the server) |
| `mujard gen <file.muj>` | write the C++ a `.muj` program compiles to (`sketch.ino` + runtime) into `<name>.gen/` (`--out dir`, `--board id` to guard) - builds nothing |
| `mujard boards caps <board>` | how far each capability (outputs, PWM, analog, Wi-Fi, the `mujard.*` libraries...) is proven on that board: *hardware* (ran on a real board), *build* (a real build succeeded), *backend* (generated, never built for it), *none* (the board lacks it). Reads the compiler's own evidence: no key, no network (`--format json`) |
| `mujard boards search <text> [--available]` / `mujard boards info <id>` | find a build target by id, name or family / one board (same as `board info`) |
| `mujard builds list\|get\|artifact\|cancel` | your builds: `list [--state s] [--mode m] [--limit n] [--after id]`, `get <id>`, `artifact <id> [--out f] [--force]`, `cancel <id> [--yes]` (the server refunds a cancelled build). `mujard build get\|status\|cancel\|artifact bld_...` works too, but only for a real build id |
| `mujard check sketch <board> [sketch\|folder]` | the build's file rules plus "can this board be built now" - **compiles nothing, charges nothing** |
| `mujard check server` / `mujard check desktop` | is the service up, in maintenance, able to build (public health page, no key) / is the desktop app installed |
| `mujard check <file.muj>` | the Mujard language check on a `.muj` program: `file:line:col` errors and warnings, `--tree` also prints the syntax tree. Compiles nothing, needs no key |
| `mujard fmt <file.muj>` | rewrite a `.muj` program in the standard layout (in place; `--check` only reports, `--print` writes to stdout) |
| `mujard new <name>` | start a `.muj` project: `<name>\<name>.muj` from a template (`--board uno`) |
| `mujard test <file.muj>` | run the program's `test` blocks on this PC, each on a fresh simulated board with simulated time (`--verbose` shows what it printed, `--format json`). Needs a C++ compiler (`g++`/`clang++` on the PATH, or `MUJARD_CXX`); exit 1 when a test fails |
| `mujard repl [--board uno]` | type Mujard lines and see what they do: each runs on this PC's simulated board through the real compiler and runtime, so values, units and int wrap-around match a board. Declarations persist; a bad line is reported and forgotten. Needs the same C++ compiler as `mujard test` |
| `mujard lsp [--stdio]` | the Mujard language server for editors (VS Code, Neovim...): live errors from the same check as `mujard check`, completion, hover, formatting |
| `mujard pkg search\|info\|install\|publish` | Mujard packages in the registry: `search [text]`, `info <name>`, `install <name>[@version] [--dir d] [--force]`, `publish <folder> --version x.y.z` (needs a key with `packages:publish`; the server compiles every version before storing it) |
| `mujard describe <library folder>` | write a `.mlib` description of an installed Arduino library from its `library.properties` and headers (`--out file.mlib`, `--print` to stdout) - put it next to a `.muj` program, or in `MUJARD_LIB_DESCS`, and `use "Name"` checks its calls |
| `mujard diagnose` | your key, the service, the desktop app and its own self-check in one list. Reads only. (The operator's `doctor` is separate) |
| `mujard get <what>` | `get boards`, `board <id>`, `builds`, `build <id>`, `models`, `model <id>`, `webhooks`, `webhook <id>`, `devices`, `adapters`, `usage`, `me` - each is exactly the command it stands for |
| `mujard libraries search <text>` / `installed` / `examples <name> [--out folder]` | Arduino libraries on the **build server** (read-only; answered by the desktop app). `installed` is the server's shared set. `examples` lists an installed library's example sketches and, with `--out`, saves them as `folder/<example>/<example>.ino` (never overwrites). There is deliberately no `install`: it would change the server for every user |
| `mujard usage windows` | this desktop's session and week allowance, free use today and prices (answered by the desktop app). A desktop is a guest device: this is **not** the developer key's `mujard usage` and not your phone's account |
| `mujard board info <id> --desktop` / `mujard boards --desktop` | the desktop app's own board catalog: upload protocol, baud, USB ids, whether this PC can flash it. No developer key needed |
| `mujard devices info <port>` / `mujard adapters info <id>` | one plugged-in board / one upload adapter and the boards it writes |
| `mujard models` / `mujard models info <id>` | the AI models the server offers (same as `ai model list\|info`) |
| `mujard ai "prompt"` | ask the AI assistant through your **developer key** (fixed price, a prompt only) |
| `mujard ai --mode <m> --board <id> [--file sketch] [--model id] [--out folder] "prompt"` | the AI with a mode and your sketch, asked through the **desktop app**, so it uses *this desktop's* allowance (free use first, then coins - a guest desktop has none), not the developer key's coins. Modes: `auto write fix debug explain chat review optimize`. The desktop app shows what it will use and asks first (`--yes` skips; a script answers "no"). The answer is printed; your files are never changed, and `--out` saves its code as new files only |
| `mujard simulate <sketch\|folder> --board uno` | build the sketch and run it a few seconds in the simulator, printing what it sent on serial. Counts as a build against this desktop's allowance; asks first. The server decides which boards it can simulate (today: the Uno) |
| `mujard ai model list` / `mujard ai model info <id>` | the AI models the server offers (read-only, via the desktop app and its key; `info` takes the full `provider:model` id or the part after `:` when unique; `--format json`) |
| `mujard webhooks list\|info\|add\|del\|test\|enable\|rotate-secret\|deliveries` | build-finished callbacks. `enable` switches one back on after failed deliveries; `rotate-secret` prints a new signing secret once (asks first, `--yes` skips) |
| `mujard devices` / `mujard list` | boards plugged into this PC (`devices list` works too) |
| `mujard drivers` / `mujard adapters` | driver problems / upload protocols |
| `mujard flash <file> --board <id> --port <p>` | write firmware to a board. Add `--check` for a dry run: it finds the board and port, validates the file and runs the adapter's preflight, then stops - **no port is opened, nothing is written, nothing is asked**. It cannot tell whether the board will answer; that is only known while flashing. `--format json` works with it |
| `mujard burn <board-id> --port <p>` | burn a bootloader via ArduinoISP |
| `mujard run <sketch\|file.muj> --board <id> --port <p>` | build then write straight to the board (a `.muj` program is compiled to a sketch first; its `board` line supplies the board) |
| `mujard upload <sketch> --board <id> --port <p>` | same as `run` |
| `mujard studio` | open Mujard Studio: the code editor inside the desktop app (opens its window when Mujard is already running, starts the app when it is not) |
| `mujard setup path [--check\|--undo]` | add the folder this `mujard` is in to your user PATH (see Install) |
| `mujard help [command]` / `mujard commands` | the command list; `commands` also lists your plugins |
| `mujard` (in a console window) / `mujard shell` | the interactive `Mujard CLI>` prompt - see below |

Exit codes: 0 done, 1 failed, 2 bad command line. Read commands take `--format json`; `mujard build --format json` prints the `mujard-cli.v1` envelope the desktop CLI uses.

## The Mujard CLI> prompt

```
C:\> mujard
Mujard CLI> build ./Blink --board nano
Mujard CLI> run "C:\My Sketches\Blink" --board nano --port COM3
Mujard CLI> exit
```

Each line runs as `mujard <line>` - the same program, so every command, plugin and exit code is
exactly the one-shot one. Spaces separate arguments, `"double"` or `'single'` quotes keep spaces
together, backslashes are plain characters (Windows paths work as typed); an unclosed quote is
refused and nothing runs. A failed command prints `(exit N)` and returns to the prompt;
`exit`/`quit` or end of input leave with 0. `clear`/`cls` clears the screen, a leading `mujard` is
ignored, `#` lines are comments. Ctrl+C cancels the running command, not the prompt.

Bare `mujard` opens the prompt only in a console window; from a pipe or script it still prints help
and exits 2. `mujard shell` opens it anywhere - fed a script (`mujard shell < steps.txt`) it runs the
lines without a prompt, and a confirm question reads end-of-input, i.e. "no": nothing is flashed
unless the line says `--yes`. A UTF-8 byte-order mark at the start of a script (Notepad and
PowerShell 5.1 add one) is ignored.

**From the desktop app:** the Mujard window has a **Command line** button (top right), and the
tray icon's menu has **Open command line (mujard)**. Either opens a console window at the `Mujard CLI>`
prompt using the `MujardCLI.exe` that sits beside `MujardServer.exe` (never one found on the PATH; a
`mujard.exe` copy ships beside it so the short command keeps working). That
window has the app's folder on its PATH, so `mujard` keeps working after you type `exit`, and
`MUJARD_DESKTOP` points at the running install. If `MujardCLI.exe` is missing (it is built only when Go is
installed) the app says so instead of failing silently.

## Keys and addresses

The API commands use the **developer API** (`/api/v1` of the hosted platform):

```
mujard login --key mjd_live_xxxxxxxx --api https://your-host/mujard-api
```

or set `MUJARD_API` / `MUJARD_KEY` (CI). Flag beats env beats the saved file. A `mjd_test_` key costs
nothing and returns sample results - it never runs a compiler. Hosted builds cost coins; local
builds (`--backend direct`, or `arduino-cli` on the PATH) are free and need no key at all.

USB commands (`flash`, `run`, `devices`, ...) are run by the desktop app: `mujard` finds
`MujardServer.exe` (`MUJARD_DESKTOP`, the PATH, a full checkout's desktop bin folder, the install
folders - older `MujardDesktop.exe` installs are found the same way)
and passes the arguments through unchanged - the same command line the desktop CLI documents.

## Who can run what

Three tiers, all fail closed:

* **everyone** - the USB commands, `login`/`logout`/`config`, `auth status`, `check server|desktop`,
  `diagnose`, `models`, `shell`, `setup`, `help`, plugins: they only touch this machine (or read the
  service's public health page).
* **signed-in** - `me`, `usage`, `boards`, `build`, `builds`, `check sketch` (the board half), `auth test`,
  `ai`, `webhooks`, and `get` of any of them: they need the developer
  key saved by `mujard login` (or `--key`/`MUJARD_KEY`). Every answer is scoped to that key by the
  server - one developer can never see another's builds, webhooks or spend.
* **operator** - the platform also has server-maintenance commands, but their implementations are not part of
  this repository: this source tree contains only the refusal gate (`operator_gate.go`), so typing one prints
  "a server maintenance command ... not available to users" and exits 1, they are absent from
  `help`/`commands`/"did you mean", and a plugin of the same name is ignored. A developer key is not an
  operator role - `login` does not unlock anything.

On a shared or kiosk machine, `MUJARD_REQUIRE_LOGIN=1` makes the interactive shell refuse to open
until a key is saved; without it the shell opens for anyone and each command checks what IT
needs (the banner says so).

## Custom commands (plugins)

Anything the built-ins do not cover can be a plugin - no rebuild, any language:

1. Write an executable named `mujard-<verb>`: `mujard-deploy.ps1`, `mujard-backup.sh`, `mujard-report.php`, `mujard-whatever.exe`.
2. Put it in `%USERPROFILE%\.mujard\commands` (Linux/macOS `~/.mujard/commands`), the repo's
   `tools/mujard/commands.d/`, a `commands/` folder next to `mujard`, or anywhere on the PATH.
3. `mujard deploy prod` now runs it: arguments are passed through, it inherits `MUJARD_API`, `MUJARD_KEY`,
   `MUJARD_ROOT`, `MUJARD_HOME`, `MUJARD_VERSION`, its output goes to the console, and its exit code is
   mujard's exit code. `mujard commands` lists it.

Scripts run through their interpreter: `.ps1` -> powershell, `.bat/.cmd` -> cmd, `.php` -> PHP,
`.py` -> python, `.sh` -> sh. A plugin is trusted code running as you; a plugin can never shadow
a built-in (so a plugin named `get`, `auth`, `check`, `builds`, `models` or `diagnose` is no longer
reachable - rename it), and the folders are searched in the order above (first match wins). See
`commands.d/mujard-hello.ps1` for the example.

## Build

```
cd tools/mujard
go build -o mujard .              # the program users get
GOOS=linux GOARCH=amd64 go build -o mujard-linux .   # the same, for Linux
go test ./...                     # the public build + proof that users cannot reach the operator tier
```

The desktop app ships its own build of this program (`MujardCLI.exe`, plus a `mujard.exe` copy) next to
`MujardServer.exe`, so the two always ship together - there is one `mujard`: this program, which the desktop
exe delegates board commands to.

## Install

Build the binary for your platform (`go build -o mujard .`), put it wherever you like, then once:

```
mujard setup path
```

That appends the folder to **your** user PATH (never the machine's), keeping every other entry
exactly as it was, `%VARS%` included. Open a new terminal and type `mujard`. `mujard setup path --check`
only reports; `mujard setup path --undo` takes the entry off again. If a different `mujard` is already
on the PATH it refuses and names it, since adding a second one would never be used.
