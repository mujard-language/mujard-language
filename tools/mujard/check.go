package main

// mujard check sketch | server | desktop, and mujard diagnose - read-only checks a user can run.
//   check sketch <board> [path]  the same file rules "mujard build" enforces, plus "is this board buildable now";
//                                it compiles nothing and charges nothing
//   check server                 the public health page (no key): is the website up, in maintenance, able to build
//   check desktop                is the desktop app (which does the USB work) installed, and which version
//   diagnose                     all of the above plus the desktop app's own self-check
// None of them prints server internals: only what the public health page already tells every app.
import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

type checkResult struct {
	lvl          level
	name, detail string
}

func (r checkResult) status() string { return strings.ToLower(strings.TrimSpace(r.lvl.tag())) }

// printChecks writes the results in the doctor's layout and returns how many failed.
func printChecks(w io.Writer, rs []checkResult) int {
	fails := 0
	for _, r := range rs {
		if r.lvl == fail {
			fails++
		}
		line(w, r.lvl, r.name, r.detail)
	}
	return fails
}

func checksData(rs []checkResult) map[string]any {
	rows := []map[string]any{}
	sum := map[string]int{"ok": 0, "warn": 0, "fail": 0, "skip": 0}
	for _, r := range rs {
		rows = append(rows, map[string]any{"name": r.name, "status": r.status(), "detail": r.detail})
		sum[r.status()]++
	}
	return map[string]any{"checks": rows, "summary": sum}
}

// apiRoot is the server's API folder (where health.php lives): the --api value without a trailing /api/v1.
func apiRoot(base string) string {
	r := strings.TrimRight(base, "/")
	return strings.TrimSuffix(r, "/api/v1")
}

// serverChecks reads the public health page. Nothing here needs a key.
func serverChecks(root string) []checkResult {
	code, h, err := plainGet(apiRoot(root)+"/health.php", 8*time.Second)
	switch {
	case err != nil:
		return []checkResult{{fail, "website", short(err) + " (" + apiRoot(root) + ")"}}
	case code != 200 || h["ok"] != true:
		return []checkResult{{fail, "website", fmt.Sprintf("answered HTTP %d without a healthy result", code)}}
	}
	rs := []checkResult{{ok, "website", "answers"}}
	if h["maintenance"] == true {
		rs = append(rs, checkResult{fail, "maintenance", "the service is in maintenance mode right now"})
	} else {
		rs = append(rs, checkResult{ok, "maintenance", "not in maintenance"})
	}
	if h["compilerReady"] == true {
		rs = append(rs, checkResult{ok, "building", "builds are on"})
	} else {
		rs = append(rs, checkResult{fail, "building", "builds are switched off or the compiler is not ready"})
	}
	if h["aiReady"] == true {
		rs = append(rs, checkResult{ok, "AI assistant", "ready"})
	} else {
		rs = append(rs, checkResult{warn, "AI assistant", "not ready"})
	}
	return rs
}

// desktopChecks finds the desktop app and asks it for its version.
func desktopChecks() []checkResult {
	exe := desktopExe()
	if exe == "" {
		return []checkResult{{fail, "desktop app", "MujardServer.exe was not found. USB commands (devices, flash, run...) need it; install it or set MUJARD_DESKTOP."}}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	out, err := exec.CommandContext(ctx, exe, "--version").Output()
	ver := strings.TrimSpace(string(out))
	if err != nil || ver == "" {
		return []checkResult{{warn, "desktop app", "found at " + exe + " but it did not report a version"}}
	}
	return []checkResult{{ok, "desktop app", ver + " at " + exe}}
}

// authResults is "auth status" as a check line.
func authResults(info authInfo) []checkResult {
	if !info.SignedIn {
		return []checkResult{{skip, "developer key", "none saved - only needed for build, boards, ai, webhooks (mujard login --key ...)"}}
	}
	return []checkResult{{ok, "developer key", info.Mode + " key " + info.KeyMasked + " (from " + info.KeySource + ")"}}
}

func runCheck(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("check", flag.ContinueOnError)
	fs.SetOutput(w)
	cl := addAPIFlags(fs)
	format := fs.String("format", "", "\"json\" prints the mujard-cli.v1 document")
	tree := fs.Bool("tree", false, "for a .muj file: also print the syntax tree the parser made")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	asJSON := *format == "json"
	if *format != "" && !asJSON {
		fmt.Fprintln(w, "--format is \"json\" or nothing.")
		return 2
	}
	sub := ""
	if len(rest) > 0 {
		sub, rest = rest[0], rest[1:]
	}
	switch sub {
	case "server", "desktop":
		if len(rest) > 0 {
			fmt.Fprintf(w, "\"mujard check %s\" takes no arguments.\n", sub)
			return 2
		}
		var rs []checkResult
		if sub == "server" {
			rs = serverChecks(cl.base)
		} else {
			rs = desktopChecks()
		}
		return finishChecks(w, "check "+sub, rs, asJSON)
	case "sketch":
		if *tree {
			fmt.Fprintln(w, "--tree is for a .muj file (mujard check file.muj --tree).")
			return 2
		}
		return checkSketch(cl, rest, w, asJSON)
	}
	// A .muj program is checked here, by the real compiler - not sent anywhere.
	if isMujardPath(sub) {
		if len(rest) > 0 {
			fmt.Fprintln(w, "\"mujard check <file.muj>\" takes one file.")
			return 2
		}
		return checkMujardFile(sub, *format, *tree, w)
	}
	fmt.Fprintln(w, "Usage: mujard check sketch <board> [sketch|folder]   (the build's file rules and the board; compiles nothing)")
	fmt.Fprintln(w, "       mujard check <file.muj> [--tree]             (the Mujard language check - errors and warnings, no build)")
	fmt.Fprintln(w, "       mujard check server                           (is the service up, and able to build?)")
	fmt.Fprintln(w, "       mujard check desktop                          (is the desktop app installed?)")
	return 2
}

func finishChecks(w io.Writer, command string, rs []checkResult, asJSON bool) int {
	fails := 0
	for _, r := range rs {
		if r.lvl == fail {
			fails++
		}
	}
	if asJSON {
		d := checksData(rs)
		if fails > 0 {
			b, _ := json.Marshal(map[string]any{"schema": "mujard-cli.v1", "command": command, "ok": false, "data": d,
				"error": map[string]any{"code": "check_failed", "message": fmt.Sprintf("%d check(s) failed", fails)}})
			fmt.Fprintln(w, string(b))
			return 1
		}
		return envOK(w, command, d)
	}
	printChecks(w, rs)
	if fails > 0 {
		return 1
	}
	return 0
}

// checkSketch validates a sketch exactly as "mujard build" would before sending it, and (with a key) asks
// whether the board can be built right now. It never starts a build.
func checkSketch(cl *apiClient, rest []string, w io.Writer, asJSON bool) int {
	cmd := "check sketch"
	bad := func(msg string) int {
		if asJSON {
			return envFail(w, cmd, "usage_error", msg, 2)
		}
		fmt.Fprintln(w, msg)
		return 2
	}
	if len(rest) == 0 || len(rest) > 2 {
		return bad("Usage: mujard check sketch <board> [sketch|folder]   e.g. mujard check sketch nano blink.ino")
	}
	board, path := rest[0], "."
	if len(rest) == 2 {
		path = rest[1]
	}
	if !boardArgRx.MatchString(board) {
		return bad(fmt.Sprintf("%q is not a board id - \"mujard boards\" lists them (uno, nano, esp32... or a full FQBN).", board))
	}
	files, _, gerr := gatherSketch(path)
	if gerr != nil {
		return bad(gerr.Error())
	}
	if cerr := checkFiles(files); cerr != nil {
		return bad(cerr.Error())
	}
	total := 0
	for _, s := range files {
		total += len(s)
	}
	rs := []checkResult{{ok, "files", fmt.Sprintf("%d file(s), %d bytes - accepted by the build's rules", len(files), total)}}

	cl.finish()
	if cl.key == "" {
		rs = append(rs, checkResult{skip, "board " + board, "not checked - no developer key (mujard login --key ...)"})
	} else {
		status, doc, err := cl.call("GET", "/boards/"+board, nil, nil)
		row := doc
		if d, _ := doc["data"].(map[string]any); d != nil {
			row = d
		}
		switch {
		case err != nil:
			rs = append(rs, checkResult{warn, "board " + board, "could not ask the server: " + short(err)})
		case status == 404:
			rs = append(rs, checkResult{fail, "board " + board, "the server does not know this board (\"mujard boards\" lists them)"})
		case status != 200:
			rs = append(rs, checkResult{fail, "board " + board, apiErrText(doc)})
		case row["available"] == false:
			rs = append(rs, checkResult{fail, "board " + board, "known, but its package is not on the build server right now"})
		default:
			rs = append(rs, checkResult{ok, "board " + board, "can be built now"})
		}
	}
	if !asJSON {
		fmt.Fprintln(w, "Nothing is compiled and nothing is charged by this check.")
	}
	return finishChecks(w, cmd, rs, asJSON)
}

// runDiagnose is the user's version of "doctor": everything a person can check about their own setup,
// in one list. The operator's "mujard doctor" (the server's internals) is separate and unchanged.
func runDiagnose(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("diagnose", flag.ContinueOnError)
	fs.SetOutput(w)
	cl := addAPIFlags(fs)
	fix := fs.Bool("fix", false, "let the desktop app retry what can repair itself (registers its key, refreshes its board list)")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if fs.NArg() > 0 {
		fmt.Fprintln(w, "\"mujard diagnose\" takes no arguments (only --api, --key and --fix).")
		return 2
	}
	info := describeAuth(cl, visitedFlags(fs))
	cl.finish()
	fails := 0
	section := func(title string) { fmt.Fprintf(w, "\n%s\n", title) }

	fmt.Fprintf(w, "mujard %s - diagnose (%s; nothing is charged)\n", version, map[bool]string{true: "--fix lets the desktop retry repairs", false: "reads only; nothing is changed"}[*fix])
	section("Your developer key")
	fails += printChecks(w, authResults(info))
	if cl.key != "" {
		status, raw, _, err := cl.callRaw("GET", "/me", nil, nil)
		var doc map[string]any
		_ = json.Unmarshal(raw, &doc)
		switch {
		case err != nil:
			fails += printChecks(w, []checkResult{{fail, "server accepts key", short(err)}})
		case status != 200:
			fails += printChecks(w, []checkResult{{fail, "server accepts key", apiErrText(doc)}})
		default:
			fails += printChecks(w, []checkResult{{ok, "server accepts key", "yes"}})
		}
	}
	section("The service")
	fails += printChecks(w, serverChecks(cl.base))
	section("This PC")
	desk := desktopChecks()
	fails += printChecks(w, desk)
	if desk[0].lvl != fail {
		section("The desktop app's own self-check")
		var dargs []string
		if *fix {
			dargs = []string{"--fix"}
		}
		if runDesktop(w, "doctor", dargs) != 0 {
			fails++
		}
	}
	fmt.Fprintln(w)
	if fails > 0 {
		fmt.Fprintf(w, "%d problem(s) found - see the FAIL lines above.\n", fails)
		return 1
	}
	fmt.Fprintln(w, "Nothing wrong found.")
	return 0
}
