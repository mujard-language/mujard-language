// mujard is the one command-line tool for Mujard. It talks to the developer API with your
// key (build, boards, usage, webhooks, ai), runs the system checks (status, doctor), drives the
// desktop app for anything that touches a board over USB (flash, run, list), runs the server's
// own admin scripts (admin), and loads custom commands named mujard-<verb> from your commands folder
// or the PATH - so `mujard anything` can be taught without rebuilding this program.
package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"mujard/lang"
)

const version = "1.0.0"

// command is one built-in verb. run takes the arguments after the verb and writes everything -
// output and errors alike - to w, so tests capture it all in one buffer.
type command struct {
	name    string
	group   string // the heading it is listed under in help
	summary string
	usage   string // "" when the verb takes no arguments of its own
	run     func(args []string, w io.Writer) int
}

var builtins = []command{
	// Your account and the cloud
	{"login", "Your key", "save your developer key and API address on this machine", "[--api url] [--key mujard_...]", runLogin},
	{"logout", "Your key", "forget the saved key", "", runLogout},
	{"config", "Your key", "show the saved API address and key (masked); path | get api|key | set api <url>", "[path | get api|key | set api <url>]", runConfig},
	{"auth", "Your key", "are you signed in? status (this machine) | test (asks the server)", "status | test [--format json]", runAuth},
	{"me", "Your key", "your key: scopes, limits, this month's spend", "", runMe},
	{"usage", "Your key", "requests and coins this month (usage windows: this desktop's session/week allowance)", "[windows] [--format json]", runUsage},
	// Building
	{"boards", "Build", "the boards a build can target (or one board's detail)", "[list] | [board-id] | info <id> | search <text> [--available] [--desktop] | caps <board> [--format json]", runBoards},
	{"board", "Build", "one build target's detail: mujard board info nano (--desktop: what this PC can flash)", "info <board-id> | <board-id> [--desktop]", runBoardInfo},
	{"build", "Build", "compile a sketch for a board: mujard build nano blink.ino", "<board> [sketch|folder] [--out f] [--wait 10m]", runBuild},
	{"builds", "Build", "your builds: list | get <id> | artifact <id> | cancel <id>", "list [--state s] [--mode m] [--limit n] [--after id] | get <id> | artifact <id> [--out f] | cancel <id> [--yes]", runBuilds},
	{"ai", "Build", "ask the AI assistant: mujard ai \"why is my LED off\" (mujard ai model list|info <id>: the models)", "<prompt> | model list | model info <id> | --mode <m> --board <id> [--file sketch] [--model id] [--out folder] \"<prompt>\"", runAI},
	{"simulate", "Build", "run a sketch in the simulator and print its serial output (uses this desktop's allowance)", "<sketch|folder> --board uno [--yes] [--format json]", runSimulate},
	{"libraries", "Build", "Arduino libraries on the build server: search | installed | examples <name>", "search <text> | installed | examples <name> [--out folder] [--format json]", runLibraries},
	{"models", "Build", "the AI models the server offers (same as mujard ai model list|info)", "[list] | info <id> [--format json]", runModels},
	{"webhooks", "Build", "build-finished callbacks: mujard webhooks list|info|add|del|test|enable|rotate-secret|deliveries", "list | info <id> | add <url> [--events a,b] | del <id> | test <id> | enable <id> | rotate-secret <id> [--yes] | deliveries <id>", runWebhooks},
	{"get", "Build", "read anything by name: mujard get boards | models | usage | builds ...", "<what> [options]", runGet},
	// This machine's USB ports - the desktop app's engine does the work
	{"devices", "Boards on USB", "the boards plugged into this PC right now (devices info <port>: one of them)", "[list | info <port>] [--format json]", runDeviceList},
	{"list", "Boards on USB", "same as devices", "", runDeviceList},
	{"drivers", "Boards on USB", "USB/serial devices with driver problems", "", runDeviceDrivers},
	{"adapters", "Boards on USB", "the upload protocols built into the desktop app (adapters info <id>: one adapter)", "[list | info <id>] [--format json]", runDeviceAdapters},
	{"flash", "Boards on USB", "write a firmware file to a board (--check: a dry run that writes nothing)", "<file> --board <id> --port <p> [--check]", runDeviceFlash},
	{"burn", "Boards on USB", "burn a bootloader through an ArduinoISP on a port", "<board-id> --port <p>", runDeviceBurn},
	{"run", "Boards on USB", "build a sketch or .muj program and write it straight to the board", "<sketch|file.muj> [--board <id>] --port <p>", runDeviceRun},
	{"upload", "Boards on USB", "same as run (the mujard-style name)", "<sketch|file.muj> [--board <id>] --port <p>", runDeviceUpload},
	{"serial", "Boards on USB", "watch a board's serial output; type a line + Enter to send (Ctrl+C stops; not billed)", "<port> [--baud N]", runDeviceSerial},
	{"logs", "Boards on USB", "the upload/burn traces this PC keeps (logs latest|<name> prints one)", "[latest|<name>]", runDeviceLogs},
	{"studio", "Boards on USB", "open Mujard Studio - the code editor inside the desktop app", "", runDeviceStudio},
	// Checks anyone can run
	{"check", "Checks", "check a sketch or .muj program (compiles nothing), the service, or the desktop app", "<file.muj> [--tree] | sketch <board> [path] | server | desktop [--format json]", runCheck},
	{"diagnose", "Checks", "is my setup working? key, service, desktop app - reads only", "", runDiagnose},
	// The Mujard language (.muj files)
	{"fmt", "Mujard language", "format a .muj program the standard way", "<file.muj> [more | folder] [--check | --print]", runMujardFmt},
	{"new", "Mujard language", "start a .muj project from a template", "<name> [--board uno]", runMujardNew},
	{"gen", "Mujard language", "write the C++ a .muj program compiles to (sketch.ino + runtime), without building", "<file.muj> [--out dir] [--board id]", runMujardGen},
	{"test", "Mujard language", "run a .muj program's test blocks on this PC, on a simulated board", "<file.muj> [--verbose] [--format json]", runMujardTest},
	{"describe", "Mujard language", "write a .mlib library description from an Arduino library's headers", "<library folder> [--out file.mlib | --print]", runMujardDescribe},
	{"doc", "Mujard language", "write the language reference pages (built-ins, boards, registers, errors)", "<out dir>", runMujardDoc},
	{"pkg", "Mujard language", "Mujard packages: search | info <name> | install <name>[@version] | publish <folder> --version x.y.z", "search [text] | info <name> | install <name>[@version] [--dir d] [--force] | publish <folder> --version x.y.z [--name n] [--board b] [--description d]", runPkg},
	{"repl", "Mujard language", "type Mujard lines and see what they do, on this PC's simulated board (same runtime as a build)", "[--board uno]", runRepl},
	{"lsp", "Mujard language", "the Mujard language server for editors (VS Code, Neovim...): errors, completion, hover, formatting", "[--stdio]", runLsp},
	// This machine
	{"setup", "The system", "put this mujard on your user PATH: mujard setup path (--check, --undo)", "path [--check | --undo]", runSetup},
	// The prompt
	{"shell", "Interactive", "the Mujard CLI> prompt: type commands without \"mujard\" (also: mujard alone)", "", runShellCommand},
}

// The operator commands (status, doctor, admin, loadtest) are added only by the operator build
// (operator_on.go); in the public build operatorCommands is empty, so they are not in this program at all.
func init() { builtins = append(builtins, operatorCommands...) }

func findBuiltin(name string) *command {
	for i := range builtins {
		if builtins[i].name == name {
			return &builtins[i]
		}
	}
	return nil
}

const usageHead = `mujard - the Mujard command line

  mujard <command> [arguments]     run a command (see the list below)
  mujard help <command>            a command's full help and options
  mujard commands                  every command, including your own plugins
  mujard                           the interactive Mujard CLI> prompt (in a console window;
                                "mujard shell" anywhere) - "exit" leaves it
  mujard --version

A command that is not built in is looked up as "mujard-<command>" in your commands
folder (%s) and on the PATH - write a script named mujard-weather and "mujard weather"
runs it. See README.md for the plugin contract.

Defaults come from the environment: MUJARD_API (the server's API folder), MUJARD_KEY
(your developer key), MUJARD_DESKTOP, MUJARD_PHP, MUJARD_ROOT, MUJARD_REQUIRE_LOGIN (1 =
the shell needs a saved key).

`

func main() {
	if len(os.Args) < 2 {
		// A person at a console gets the prompt; a script or pipe keeps the old answer (help, exit 2).
		if stdinIsConsole() {
			os.Exit(runShellCommand(nil, os.Stdout))
		}
		printHelp(os.Stdout, "")
		os.Exit(2)
	}
	verb := os.Args[1]
	switch verb {
	case "--version", "version", "-v":
		// --json: the tool's version and the language version, for the build server's self-check and the apps.
		if len(os.Args) > 2 && (os.Args[2] == "--json" || os.Args[2] == "--format=json") {
			fmt.Printf("{\"version\":%q,\"language\":%d}\n", version, lang.LanguageVersion())
			return
		}
		fmt.Println(version)
		return
	case "-h", "--help":
		printHelp(os.Stdout, "")
		return
	case "help":
		printHelp(os.Stdout, firstArg(os.Args[2:]))
		return
	case "commands":
		printHelp(os.Stdout, "")
		return
	}
	// An operator command name is answered before anything else is looked up - a plugin or a typo can never stand in for it.
	if operatorOnly[verb] && !adminEnabled() {
		operatorRefused(os.Stderr, verb)
		os.Exit(1)
	}
	if c := findBuiltin(verb); c != nil {
		if selfHelp[c.name] && wantsHelp(os.Args[2:]) {
			printHelp(os.Stdout, c.name)
			return
		}
		os.Exit(c.run(os.Args[2:], os.Stdout))
	}
	if p := findPlugin(verb); p != "" {
		os.Exit(runPlugin(p, os.Args[2:]))
	}
	fmt.Fprintf(os.Stderr, "mujard: unknown command %q (not built in, and no mujard-%s found in the commands folder or on the PATH)\n", verb, verb)
	if s := suggest(verb, visibleVerbs()); len(s) > 0 {
		fmt.Fprintf(os.Stderr, "Did you mean \"%s\"?  (\"mujard help <command>\" explains one)\n", strings.Join(s, "\" or \""))
	}
	fmt.Fprintln(os.Stderr)
	printHelp(os.Stderr, "")
	os.Exit(2)
}

func firstArg(args []string) string {
	for _, a := range args {
		if !strings.HasPrefix(a, "-") {
			return a
		}
	}
	return ""
}

// printHelp writes the command list grouped like the reference card, or one command's detail.
func printHelp(w io.Writer, only string) {
	if only != "" {
		if operatorOnly[only] && !adminEnabled() {
			operatorRefused(w, only)
			return
		}
		if c := findBuiltin(only); c != nil {
			fmt.Fprintf(w, "mujard %s %s\n\n  %s\n", c.name, c.usage, c.summary)
			if ex := helpExamples[c.name]; ex != "" {
				fmt.Fprintf(w, "\nExamples\n%s", ex)
			}
			fmt.Fprint(w, "\nEvery command takes --help for its own options.\n")
		} else {
			fmt.Fprintf(w, "no built-in command %q; if it is a plugin, run \"mujard %s --help\".\n", only, only)
		}
		return
	}
	fmt.Fprintf(w, usageHead, commandsDir())
	if operatorBuild && adminEnabled() {
		fmt.Fprintln(w, "This is the operator build (MUJARD_ADMIN=1 is set): status, doctor, admin and loadtest are enabled.")
		fmt.Fprintln(w)
	}
	last := ""
	for _, c := range builtins {
		if operatorOnly[c.name] && !adminEnabled() {
			continue // hidden unless the operator enabled it - users never see it exists
		}
		if c.group != last {
			fmt.Fprintf(w, "\n%s\n", c.group)
			last = c.group
		}
		fmt.Fprintf(w, "  %-9s %s\n", c.name, c.summary)
	}
	if plugins := listPlugins(); len(plugins) > 0 {
		sort.Slice(plugins, func(i, j int) bool { return plugins[i].name < plugins[j].name })
		fmt.Fprintf(w, "\nYour commands (plugins)\n")
		for _, p := range plugins {
			fmt.Fprintf(w, "  %-9s %s\n", p.name, p.path)
		}
	}
	fmt.Fprintln(w, "\nExit codes: 0 done, 1 failed, 2 bad command line.")
}
