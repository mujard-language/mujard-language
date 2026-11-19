package main

// Phase 2: read-only views that the desktop app answers, plus two Go-side lookups over its JSON.
//   mujard libraries search|installed|examples   (desktop verb "libraries")
//   mujard usage windows                         (desktop verb "usage windows")
//   mujard board info <id> --desktop             (desktop verb "board"; also "boards --desktop" for the list)
//   mujard devices info <port>                   (the desktop's "list --format json", one board)
//   mujard adapters info <id>                    (the desktop's "adapters --format json", one adapter)
// None of them changes anything. The library and usage verbs only exist in a desktop app built with
// Phase 2; an older MujardServer.exe answers 'Unknown command', which is turned into one plain sentence.
import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// runDesktopNew is runDesktop for a verb a desktop app may not have yet. Output is passed through
// unchanged (stdout to w, stderr to stderr), the desktop's exit code is the exit code, and a desktop that
// does not know the verb gets one clear message instead of its generic usage dump.
func runDesktopNew(w io.Writer, verb string, args []string) int {
	exe := desktopExe()
	if exe == "" {
		return runDesktop(w, verb, args) // prints the "desktop app is not on this machine" message
	}
	var stdout, stderr bytes.Buffer
	c := exec.Command(exe, append([]string{verb}, args...)...)
	c.Stdout, c.Stderr, c.Stdin = &stdout, &stderr, os.Stdin
	code := 0
	if err := c.Run(); err != nil {
		ee, ok := err.(*exec.ExitError)
		if !ok {
			fmt.Fprintf(w, "The desktop app could not be started: %s\n", err)
			return 1
		}
		code = ee.ExitCode()
	}
	if code == 2 && isUnknownDesktopVerb(stdout.String()+stderr.String(), verb) {
		msg := fmt.Sprintf("\"mujard %s\" needs a newer Mujard desktop app than the one installed (%s). Update the desktop app, or set MUJARD_DESKTOP to a newer MujardServer.exe.", verb, exe)
		if wantsJSONFormat(args) {
			return envFail(w, verb, "desktop_too_old", msg, 2)
		}
		fmt.Fprintln(w, msg)
		return 2
	}
	_, _ = w.Write(stdout.Bytes())
	_, _ = os.Stderr.Write(stderr.Bytes())
	return code
}

var unknownVerbRx = map[string]*regexp.Regexp{}

// isUnknownDesktopVerb: the desktop CLI's wording for a verb it has never heard of, in text or inside its JSON envelope.
func isUnknownDesktopVerb(out, verb string) bool {
	rx := unknownVerbRx[verb]
	if rx == nil {
		rx = regexp.MustCompile(`Unknown command \\?"` + regexp.QuoteMeta(verb) + `\\?"`)
		unknownVerbRx[verb] = rx
	}
	return rx.MatchString(out)
}

func wantsJSONFormat(args []string) bool {
	for i, a := range args {
		if a == "--format=json" || (a == "--format" && i+1 < len(args) && args[i+1] == "json") {
			return true
		}
	}
	return false
}

// formatArgs picks "--format x" / "--format=x" out of a command line - the only Go-side option the desktop
// app understands, so the Go flags (--api, --key) are never forwarded to it.
func formatArgs(args []string) []string {
	for i, a := range args {
		if strings.HasPrefix(a, "--format=") {
			return []string{a}
		}
		if a == "--format" && i+1 < len(args) {
			return []string{"--format", args[i+1]}
		}
	}
	return nil
}

// desktopBoards is "mujard boards --desktop" and "mujard board info <id> --desktop": the desktop app's own board
// catalog. It needs no developer key, and it is the answer for "can this PC flash that board, and how".
func desktopBoards(rest []string, format string, w io.Writer) int {
	if len(rest) > 0 && rest[0] == "list" {
		rest = rest[1:]
	}
	if len(rest) > 0 && rest[0] == "info" {
		rest = rest[1:]
	}
	if len(rest) > 1 || (format != "" && format != "json" && format != "text") {
		fmt.Fprintln(w, "Usage: mujard boards --desktop | mujard board info <id> --desktop   [--format json]")
		return 2
	}
	var fmtArgs []string
	if format != "" {
		fmtArgs = []string{"--format", format}
	}
	if len(rest) == 1 {
		return runDesktop(w, "board", append([]string{rest[0]}, fmtArgs...))
	}
	return runDesktop(w, "boards", fmtArgs)
}

// desktopHelp is the installed desktop app's own usage text (cached): the cheapest honest way to learn
// whether it has a command, for the verbs that ask a question before spending - those must stream their
// prompt, so they cannot be run through runDesktopNew's capture to find out.
var desktopHelpCache = map[string]string{}

func desktopHelp(exe string) string {
	if s, ok := desktopHelpCache[exe]; ok {
		return s
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	out, _ := exec.CommandContext(ctx, exe, "--help").Output()
	desktopHelpCache[exe] = string(out)
	return string(out)
}

// requireDesktopFeature runs a desktop verb that bills only if the installed desktop app knows it.
// marker is a phrase from that app's usage text. A desktop that is missing (or older) is explained in one
// sentence; nothing is sent to it.
func requireDesktopFeature(w io.Writer, verb, marker string, args []string) (bool, int) {
	exe := desktopExe()
	if exe == "" {
		return false, runDesktop(w, verb, args) // prints the "desktop app is not on this machine" message
	}
	if strings.Contains(desktopHelp(exe), marker) {
		return true, 0
	}
	msg := fmt.Sprintf("\"mujard %s\" needs a newer Mujard desktop app than the one installed (%s). Update the desktop app, or set MUJARD_DESKTOP to a newer MujardServer.exe.", verb, exe)
	if wantsJSONFormat(args) {
		return false, envFail(w, verb, "desktop_too_old", msg, 2)
	}
	fmt.Fprintln(w, msg)
	return false, 2
}

// runSimulate: "mujard simulate <sketch> --board uno" - the desktop app asks first, then spends this desktop's allowance.
func runSimulate(args []string, w io.Writer) int {
	if firstArg(withoutFlags(args, "board", "format")) == "" {
		fmt.Fprintln(w, "Usage: mujard simulate <sketch|folder> --board uno [--yes] [--format json]")
		fmt.Fprintln(w, "Runs the sketch in the simulator and prints what it sent on serial. It counts as a build against THIS DESKTOP's")
		fmt.Fprintln(w, "allowance (not your developer key); the desktop app asks before it starts.")
		return 2
	}
	if ok, code := requireDesktopFeature(w, "simulate", "MujardServer.exe simulate", args); !ok {
		return code
	}
	return runDesktop(w, "simulate", args)
}

// withoutFlags drops "--name value" / "--name=value" pairs, so the developer key and API address of the Go
// side can never be passed on to the desktop app.
func withoutFlags(args []string, names ...string) []string {
	skip := map[string]bool{}
	for _, n := range names {
		skip["--"+n] = true
	}
	out := []string{}
	for i := 0; i < len(args); i++ {
		a := args[i]
		if skip[a] {
			i++
			continue
		}
		if k, _, ok := strings.Cut(a, "="); ok && skip[k] {
			continue
		}
		out = append(out, a)
	}
	return out
}

// isDesktopAIRequest: "--mode", "--file" or "--model" turns "mujard ai" into a question the desktop app sends with
// THIS DESKTOP's allowance. Before this existed those flags were a usage error, so nothing that worked changes.
func isDesktopAIRequest(args []string) bool {
	for _, a := range args {
		if a == "--" {
			return false
		}
		for _, f := range []string{"--mode", "--file", "--model"} {
			if a == f || strings.HasPrefix(a, f+"=") {
				return true
			}
		}
	}
	return false
}

func runDesktopAI(args []string, w io.Writer) int {
	pass := withoutFlags(args, "api", "key")
	if ok, code := requireDesktopFeature(w, "ai", "MujardServer.exe ai --mode", pass); !ok {
		return code
	}
	return runDesktop(w, "ai", pass)
}

func runLibraries(args []string, w io.Writer) int {
	if firstArg(withoutFlags(args, "format")) == "" {
		fmt.Fprintln(w, "Usage: mujard libraries search <text>            find libraries on the build server's index")
		fmt.Fprintln(w, "       mujard libraries installed               what is installed on the build server (shared by every user)")
		fmt.Fprintln(w, "       mujard libraries examples <name> [--out folder]   an installed library's example sketches")
		return 2
	}
	return runDesktopNew(w, "libraries", args)
}

// desktopJSON runs a desktop verb with --format json and returns its "data" object.
func desktopJSON(verb string) (map[string]any, string, int) {
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()
	exe := desktopExe()
	cmd := exec.CommandContext(ctx, exe, verb, "--format", "json")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil && out.Len() == 0 {
		return nil, "The desktop app could not answer: " + err.Error(), 1
	}
	var env struct {
		OK    bool           `json:"ok"`
		Data  map[string]any `json:"data"`
		Error map[string]any `json:"error"`
	}
	if err := json.Unmarshal(bytes.TrimSpace(out.Bytes()), &env); err != nil {
		return nil, "The desktop app's answer was not understood.", 1
	}
	if !env.OK {
		return nil, str(env.Error, "message"), 1
	}
	return env.Data, "", 0
}

var devicePortRx = regexp.MustCompile(`(?i)^(COM\d{1,3}|[A-Z]:)$`)

// runDeviceInfo: "mujard devices info COM5" - one plugged-in board, from the desktop's own scan.
func runDeviceInfo(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("devices info", flag.ContinueOnError)
	fs.SetOutput(w)
	format := fs.String("format", "", "\"json\" prints the mujard-cli.v1 document")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	asJSON := *format == "json"
	if len(rest) != 1 || !devicePortRx.MatchString(rest[0]) || (*format != "" && !asJSON) {
		fmt.Fprintln(w, "Usage: mujard devices info <port> [--format json]   e.g. mujard devices info COM5   (\"mujard devices\" lists the ports)")
		return 2
	}
	if desktopExe() == "" {
		return runDesktop(w, "list", nil)
	}
	data, why, code := desktopJSON("list")
	if code != 0 {
		if asJSON {
			return envFail(w, "devices info", "desktop_error", why, 1)
		}
		fmt.Fprintln(w, why)
		return 1
	}
	list, _ := data["boards"].([]any)
	for _, e := range list {
		b, _ := e.(map[string]any)
		if b == nil || !strings.EqualFold(str(b, "port"), rest[0]) {
			continue
		}
		if asJSON {
			return envOK(w, "devices info", b)
		}
		fmt.Fprintf(w, "%s  %s\n", str(b, "port"), str(b, "name"))
		fmt.Fprintf(w, "  chip      %s\n  USB       VID %s  PID %s\n", str(b, "chip"), str(b, "vid"), str(b, "pid"))
		if id := str(b, "boardId"); id != "" {
			fmt.Fprintf(w, "  board id  %s   (flash it with: mujard flash <file> --board %s --port %s)\n", id, id, str(b, "port"))
		} else {
			fmt.Fprintln(w, "  board id  not named by its USB ids - pass --board yourself when flashing")
		}
		return 0
	}
	msg := "Nothing is plugged into " + strings.ToUpper(rest[0]) + " on this PC. \"mujard devices\" lists what is."
	if asJSON {
		return envFail(w, "devices info", "not_found", msg, 1)
	}
	fmt.Fprintln(w, msg)
	return 1
}

// runAdapterInfo: "mujard adapters info serial-flash" - one upload adapter and the boards it can write.
func runAdapterInfo(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("adapters info", flag.ContinueOnError)
	fs.SetOutput(w)
	format := fs.String("format", "", "\"json\" prints the mujard-cli.v1 document")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	asJSON := *format == "json"
	if len(rest) != 1 || (*format != "" && !asJSON) {
		fmt.Fprintln(w, "Usage: mujard adapters info <adapter id> [--format json]   (\"mujard adapters\" lists them)")
		return 2
	}
	if desktopExe() == "" {
		return runDesktop(w, "adapters", nil)
	}
	data, why, code := desktopJSON("adapters")
	if code != 0 {
		if asJSON {
			return envFail(w, "adapters info", "desktop_error", why, 1)
		}
		fmt.Fprintln(w, why)
		return 1
	}
	list, _ := data["adapters"].([]any)
	for _, e := range list {
		a, _ := e.(map[string]any)
		if a == nil || !strings.EqualFold(str(a, "id"), rest[0]) {
			continue
		}
		if asJSON {
			return envOK(w, "adapters info", a)
		}
		boards := strList2(a["boards"])
		fmt.Fprintf(w, "%s  version %s  (runs on %s)\n", str(a, "id"), str(a, "version"), str(a, "runsOn"))
		fmt.Fprintf(w, "  protocols  %s\n", strings.Join(strList2(a["protocols"]), ", "))
		show := boards
		if len(show) > 20 {
			show = show[:20]
		}
		fmt.Fprintf(w, "  boards     %d: %s", len(boards), strings.Join(show, ", "))
		if len(boards) > len(show) {
			fmt.Fprintf(w, ", ... (\"mujard adapters --format json\" lists all %d)", len(boards))
		}
		fmt.Fprintln(w)
		return 0
	}
	msg := fmt.Sprintf("This desktop has no adapter called %q.", rest[0])
	if phone := strList2(data["phoneOnly"]); len(phone) > 0 {
		for _, p := range phone {
			if strings.EqualFold(p, rest[0]) {
				msg = p + " uploads run on the phone app, not on this desktop."
			}
		}
	}
	msg += " \"mujard adapters\" lists the ones it has."
	if asJSON {
		return envFail(w, "adapters info", "not_found", msg, 1)
	}
	fmt.Fprintln(w, msg)
	return 1
}

// strList2 reads a JSON array of strings in the order given (strList sorts; the desktop's order is meaningful).
func strList2(v any) []string {
	out := []string{}
	if l, _ := v.([]any); l != nil {
		for _, e := range l {
			if s, ok := e.(string); ok {
				out = append(out, s)
			}
		}
	}
	return out
}
