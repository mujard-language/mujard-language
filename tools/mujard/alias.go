package main

// The friendly spellings: "mujard get boards" is "mujard boards", "mujard models" is "mujard ai model list", and a
// mistyped verb is answered with the nearest real one. Every alias is a rewrite of the argument list to
// the canonical command BEFORE it runs, so the output, the JSON and the exit code are exactly the
// canonical command's - there is no second implementation to drift.
import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
)

// dispatch runs a built-in by name. It is set in init so that the "get" verb can call back into the
// builtins table without the table and its own handlers forming an initialisation cycle.
var dispatch func(name string, args []string, w io.Writer) int

func init() {
	dispatch = func(name string, args []string, w io.Writer) int {
		c := findBuiltin(name)
		if c == nil {
			fmt.Fprintf(w, "mujard: no built-in command %q\n", name)
			return 2
		}
		return c.run(args, w)
	}
}

// getNouns maps "mujard get <noun>" to the canonical verb and the arguments that come before the user's own.
// Only nouns that already have a read command are here; nothing is invented.
var getNouns = map[string][]string{
	"me":       {"me"},
	"usage":    {"usage"},
	"boards":   {"boards"},
	"board":    {"board", "info"},
	"models":   {"models", "list"},
	"model":    {"models", "info"},
	"devices":  {"devices"},
	"adapters": {"adapters"},
	"webhooks": {"webhooks", "list"},
	"webhook":  {"webhooks", "info"},
	"builds":   {"builds", "list"},
	"build":    {"builds", "get"},
}

func getNounList() string {
	names := make([]string, 0, len(getNouns))
	for n := range getNouns {
		names = append(names, n)
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}

// rewriteGet turns the arguments after "get" into the canonical command line (verb first).
func rewriteGet(args []string) ([]string, bool) {
	if len(args) == 0 {
		return nil, false
	}
	canon, ok := getNouns[args[0]]
	if !ok {
		return nil, false
	}
	out := append([]string{}, canon...)
	return append(out, args[1:]...), true
}

func runGet(args []string, w io.Writer) int {
	full, ok := rewriteGet(args)
	if !ok {
		if len(args) == 0 {
			fmt.Fprintf(w, "Usage: mujard get <what> [options]   what = %s\n", getNounList())
		} else {
			fmt.Fprintf(w, "mujard get: %q is not something this can read. Choose one of: %s\n", args[0], getNounList())
		}
		return 2
	}
	return dispatch(full[0], full[1:], w)
}

// runModels: "mujard models [list]" and "mujard models info <id>" are the spellings of "mujard ai model list|info <id>".
// The desktop app answers them from the server's model catalog, exactly as before.
func runModels(args []string, w io.Writer) int {
	sub, rest := "list", args
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		sub, rest = args[0], args[1:]
	}
	return runDesktop(w, "ai", append([]string{"model", sub}, rest...))
}

// selfHelp lists the commands where --help / -h anywhere in the arguments prints "mujard help <command>".
// Only commands that are new, or entirely this program's own, are here: the USB verbs keep passing
// --help to the desktop app, which prints its own fuller usage, and "mujard ai" keeps treating every
// word as a prompt.
var selfHelp = map[string]bool{
	"auth": true, "builds": true, "check": true, "diagnose": true, "get": true, "models": true, "libraries": true,
	"fmt": true, "new": true, "gen": true, "test": true,
}

func wantsHelp(args []string) bool {
	for _, a := range args {
		if a == "--" {
			return false
		}
		if a == "-h" || a == "--help" {
			return true
		}
	}
	return false
}

// helpExamples are shown by "mujard help <command>" for commands that benefit from a worked example.
var helpExamples = map[string]string{
	"get":       "  mujard get boards\n  mujard get board nano --format json\n  mujard get builds --state done\n  mujard get models\n\n  What it can read: boards, board, builds, build, models, model, webhooks, webhook, devices, adapters, usage, me.\n  Each is exactly the command it stands for, with the same output.\n",
	"auth":      "  mujard auth status          is a key set? (no network)\n  mujard auth test            does the server accept it?\n",
	"builds":    "  mujard builds list --state done --limit 5\n  mujard builds get bld_0123456789abcdef\n  mujard builds artifact bld_0123456789abcdef --out blink.hex\n  mujard builds cancel bld_0123456789abcdef\n",
	"build":     "  mujard build nano blink.ino\n  mujard build esp32 ./MySketch --out firmware.bin\n  mujard build blink.muj                   compile Mujard, then build: locally with arduino-cli, else on the server\n  mujard build blink.muj --show-cpp        also write the generated C++ to blink.gen/ for review\n  mujard build get bld_0123456789abcdef       (same as mujard builds get ...)\n",
	"check":     "  mujard check blink.muj              the Mujard language check: errors and warnings, compiles nothing\n  mujard check blink.muj --tree       ...plus the syntax tree the parser made\n  mujard check sketch nano blink.ino   checks the files and the board; compiles nothing, charges nothing\n  mujard check server                  is the service up and able to build?\n  mujard check desktop                 is the desktop app installed?\n",
	"fmt":       "  mujard fmt blink.muj                rewrite it in the standard layout (4-space indents, spacing)\n  mujard fmt ./sketches --check        list the files that are not formatted; change nothing\n",
	"new":       "  mujard new blink                     makes blink/blink.muj, a first program for the board\n  mujard new blink --board esp32\n",
	"lsp":       "  An editor starts it, not you: point your editor's language-server setting for .muj/.mjd files at\n    mujard lsp --stdio\n  It checks as you type with the same compiler as 'mujard check', completes keywords, built-ins and\n  members, shows what a name is on hover, and formats with 'mujard fmt''s rules. No network, no key.\n",
	"gen":       "  mujard gen blink.muj                  the generated Arduino sketch, in blink.gen/\n  mujard gen blink.muj --out ./out --board uno   choose the folder; --board checks the program's board line\n\n  Writes sketch.ino, mujard_program.cpp and the runtime files - what the build server runs before arduino-cli.\n",
	"test":      "  mujard test dustbin.muj              run its test blocks on a simulated board on this PC\n  mujard test dustbin.muj --verbose    also show what the program printed in each test\n\n  Each test starts a fresh board with simulated time, so 'advance(3 s)' takes milliseconds. Needs a C++\n  compiler (g++ or clang++ on the PATH, or MUJARD_CXX); nothing is uploaded and no board is involved.\n",
	"diagnose":  "  mujard diagnose           your key, the service, the desktop app and its self-check, in one list\n  mujard diagnose --fix     ...and let the desktop retry what can repair itself (key, board list)\n",
	"boards":    "  mujard boards\n  mujard boards search esp32 --available\n  mujard boards info nano\n",
	"webhooks":  "  mujard webhooks add https://example.com/hook --events build.completed\n  mujard webhooks info 3\n  mujard webhooks enable 3            switch it back on after failed deliveries\n  mujard webhooks rotate-secret 3     a new signing secret, shown once\n",
	"libraries": "  mujard libraries search servo\n  mujard libraries installed             what the build server has (shared by every user)\n  mujard libraries examples Servo --out ./examples\n\n  Read-only. Installing a library is not offered: it changes the build server everyone shares.\n",
	"usage":     "  mujard usage                   your developer key's requests and coins this month\n  mujard usage windows           this desktop's session and week allowance (a guest device: not your key, not your phone account)\n",
	"devices":   "  mujard devices\n  mujard devices info COM5\n",
	"adapters":  "  mujard adapters\n  mujard adapters info serial-flash\n",
	"board":     "  mujard board info nano                the build API's view: can it be built now?\n  mujard board info nano --desktop      the desktop app's view: protocol, baud, USB ids, can this PC flash it?\n",
	"ai":        "  mujard ai \"why is my LED off\"                                  one question through your developer key (fixed price, no files)\n  mujard ai --mode explain --file blink.ino --board uno             the AI with your sketch and a mode, through the desktop app\n  mujard ai --mode write --board uno \"blink pin 9\" --out ./ai      ...and save its code as NEW files in ./ai\n\n  The two are paid differently: plain \"mujard ai\" uses your developer key's coins; with --mode/--file/--model it uses THIS DESKTOP's\n  allowance (free use first, then coins - a guest desktop has none). The desktop app asks before it sends; --yes skips the question.\n  Modes: auto, write, fix, debug, explain, chat, review, optimize. Your files are never changed.\n",
	"simulate":  "  mujard simulate blink.ino --board uno\n\n  Builds the sketch and runs it for a few seconds in the simulator, printing what it sent on serial. It counts as a build against THIS\n  DESKTOP's allowance (free builds first); the desktop app asks before it starts. The server decides which boards it can simulate.\n",
	"flash":     "  mujard flash blink.hex --board uno --port COM5            writes it (asks first, --yes skips the question)\n  mujard flash blink.hex --board uno --port COM5 --check    dry run: validates the file, board and port; opens no port, writes nothing, asks nothing\n",
	"serial":    "  mujard serial COM5 [--baud 115200]\n\n  Watches the board's serial output and types lines back to it - the same port the phone's terminal and Studio's serial\n  panel watch. Ctrl+C stops. It is not billed: it reads and writes the wire, nothing is built.\n",
	"logs":      "  mujard logs                    the upload/burn traces this PC keeps, newest first\n  mujard logs latest             print the newest one (last 400 lines)\n  mujard logs <name>             print one by name\n",
	"config":    "  mujard config path\n  mujard config get api\n  mujard config set api https://your-host/mujard-api\n",
	"models":    "  mujard models                   the models the server offers this desktop\n  mujard models info <id>         one model (use an id from the list)\n",
}

// ---- "did you mean" ---------------------------------------------------------------------------------

func editDistance(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	prev := make([]int, len(rb)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(ra); i++ {
		cur := make([]int, len(rb)+1)
		cur[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			cur[j] = min(prev[j]+1, cur[j-1]+1, prev[j-1]+cost)
		}
		prev = cur
	}
	return prev[len(rb)]
}

// suggest returns up to three names close to the mistyped verb. It only ever offers what the person may
// run: operator commands are left out unless the operator gate is on.
func suggest(verb string, names []string) []string {
	verb = strings.ToLower(verb)
	type cand struct {
		name string
		d    int
	}
	var found []cand
	for _, n := range names {
		d := editDistance(verb, n)
		limit := 1
		if len(n) >= 5 {
			limit = 2
		}
		if d <= limit || (len(verb) >= 3 && (strings.HasPrefix(n, verb) || strings.HasPrefix(verb, n))) {
			found = append(found, cand{n, d})
		}
	}
	sort.SliceStable(found, func(i, j int) bool { return found[i].d < found[j].d })
	out := []string{}
	for _, c := range found {
		if len(out) == 3 {
			break
		}
		out = append(out, c.name)
	}
	return out
}

// visibleVerbs is every verb the person can run: built-ins they may use, then their plugins.
func visibleVerbs() []string {
	var names []string
	for _, c := range builtins {
		if operatorOnly[c.name] && !adminEnabled() {
			continue
		}
		names = append(names, c.name)
	}
	for _, p := range listPlugins() {
		names = append(names, p.name)
	}
	return names
}

// ---- the machine-readable envelope for the NEW commands -------------------------------------------

// envOK/envFail print the mujard-cli.v1 document the desktop CLI uses: {schema, command, ok, data, error}.
// The older Go commands keep their own output (raw API documents; build's slim envelope) unchanged.
func envOK(w io.Writer, command string, data any) int {
	b, _ := json.Marshal(map[string]any{"schema": "mujard-cli.v1", "command": command, "ok": true, "data": data, "error": nil})
	fmt.Fprintln(w, string(b))
	return 0
}

func envFail(w io.Writer, command, code, msg string, exit int) int {
	b, _ := json.Marshal(map[string]any{"schema": "mujard-cli.v1", "command": command, "ok": false, "data": nil,
		"error": map[string]any{"code": code, "message": msg}})
	fmt.Fprintln(w, string(b))
	return exit
}
