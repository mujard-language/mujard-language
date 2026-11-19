package main

// mujard build <board> [sketch|folder] - compile on the server through the developer API
// (POST /v1/builds is asynchronous: queue, poll, then fetch the artifact). The file rules
// below are the same ones ApiJobs::validateFiles enforces server-side, so a bad sketch is
// refused here before it ever costs a coin.
import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	boardArgRx = regexp.MustCompile(`^[a-z0-9_]{1,40}$|^[A-Za-z0-9_.\-]+(:[A-Za-z0-9_.\-]+){2}$`) // short id or FQBN
	fileNameRx = regexp.MustCompile(`^[A-Za-z0-9_\-]{1,40}\.(ino|h|hpp|c|cpp|S)$`)
)

const (
	maxFiles      = 20
	maxFileBytes  = 200_000
	maxTotalBytes = 500_000
)

// gatherSketch reads a sketch file or folder into the API's files map. The API calls the
// entry file "sketch.ino" whatever it is called on disk, so the one .ino is renamed.
func gatherSketch(path string) (files map[string]string, baseName string, err error) {
	st, e := os.Stat(path)
	if e != nil {
		return nil, "", fmt.Errorf("cannot read %s", path)
	}
	if !st.IsDir() {
		if !strings.HasSuffix(strings.ToLower(filepath.Base(path)), ".ino") {
			return nil, "", fmt.Errorf("%s is not a sketch - give an .ino file or the folder holding it", filepath.Base(path))
		}
		b, e := os.ReadFile(path)
		if e != nil {
			return nil, "", fmt.Errorf("cannot read %s", path)
		}
		return map[string]string{"sketch.ino": string(b)}, strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)), nil
	}
	entries, e := os.ReadDir(path)
	if e != nil {
		return nil, "", fmt.Errorf("cannot list %s", path)
	}
	files = map[string]string{}
	var ino string
	for _, en := range entries {
		name := en.Name()
		if en.IsDir() || !fileNameRx.MatchString(name) {
			continue
		}
		if strings.HasSuffix(strings.ToLower(name), ".ino") {
			if ino != "" {
				return nil, "", fmt.Errorf("two .ino files (%s and %s): a build needs one entry file - pass it directly, e.g. mujard build <board> %s", ino, name, ino)
			}
			ino = name
			continue
		}
		b, e := os.ReadFile(filepath.Join(path, name))
		if e != nil {
			return nil, "", fmt.Errorf("cannot read %s", name)
		}
		files[name] = string(b)
	}
	if ino == "" {
		return nil, "", fmt.Errorf("no .ino sketch in %s", path)
	}
	b, e := os.ReadFile(filepath.Join(path, ino))
	if e != nil {
		return nil, "", fmt.Errorf("cannot read %s", ino)
	}
	files["sketch.ino"] = string(b)
	return files, strings.TrimSuffix(ino, filepath.Ext(ino)), nil
}

// checkFiles is the client half of ApiJobs::validateFiles - same limits, friendlier messages.
func checkFiles(files map[string]string) error {
	if len(files) == 0 {
		return fmt.Errorf("no source files to build")
	}
	if _, ok := files["sketch.ino"]; !ok {
		return fmt.Errorf("a build needs a file called sketch.ino")
	}
	if len(files) > maxFiles {
		return fmt.Errorf("at most %d files per build", maxFiles)
	}
	total := 0
	for name, src := range files {
		if !fileNameRx.MatchString(name) {
			return fmt.Errorf("%q is not a file name the build accepts (letters, digits, _ and - with a .ino .h .hpp .c .cpp or .S ending, no folders)", name)
		}
		if len(src) > maxFileBytes {
			return fmt.Errorf("%s is too large (limit %d KB)", name, maxFileBytes/1000)
		}
		if !utf8.ValidString(src) || strings.ContainsRune(src, 0) {
			return fmt.Errorf("%s must be UTF-8 text", name)
		}
		total += len(src)
	}
	if total > maxTotalBytes {
		return fmt.Errorf("the sources are too large together (limit %d KB)", maxTotalBytes/1000)
	}
	return nil
}

func runBuild(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("build", flag.ContinueOnError)
	fs.SetOutput(w)
	cl := addAPIFlags(fs)
	out := fs.String("out", "", "where the firmware goes (default: <sketch>.<format> here)")
	format := fs.String("format", "", "\"json\" for machine-readable output")
	wait := fs.Duration("wait", 10*time.Minute, "how long to wait for the build before giving up")
	boardFlag := fs.String("board", "", "the board id (for a .muj program it must match its 'board' line; with it the positional board is the sketch path)")
	showCpp := fs.Bool("show-cpp", false, "for a .muj program: also write the generated C++ next to the source (<name>.gen/)")
	backend := fs.String("backend", "", "build backend: direct (mujard-build only) | auto (direct, then the normal path) | arduino (the existing path)")
	offline := fs.Bool("offline", false, "build on this PC only (the direct backend, then arduino-cli) - never on the server; needs no key")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	// "mujard build get|status|cancel|artifact bld_..." is "mujard builds ...": only when the second word is
	// really a build id, so a board id can never be taken for a sub-command.
	if len(rest) == 2 && buildIDRx.MatchString(rest[1]) {
		if canon, ok := buildSubAlias(rest[0]); ok {
			forwarded := append([]string{}, args...)
			for i, a := range forwarded {
				if a == rest[0] {
					forwarded[i] = canon
					break
				}
			}
			return runBuilds(forwarded, w)
		}
	}
	cl.finish()
	asJSON := *format == "json"
	if *format != "" && !asJSON {
		fmt.Fprintln(w, "--format is \"json\" or nothing.")
		return 2
	}
	switch *backend {
	case "", "arduino", "direct", "auto":
	default:
		fmt.Fprintln(w, "--backend is direct, auto or arduino.")
		return 2
	}
	// A .muj file compiles here first; the generated C++ is what gets built (locally or on the server).
	if len(rest) > 0 && isMujardPath(rest[0]) {
		return buildMujard(cl, rest[0], *boardFlag, *out, *showCpp, rest[1:], *wait, asJSON, *backend, *offline, w)
	}
	if *showCpp {
		fmt.Fprintln(w, "--show-cpp is for a .muj program (mujard build file.muj --show-cpp).")
		return 2
	}
	board := *boardFlag
	sketchPath := "."
	switch {
	case board != "":
		// mujard build --board uno [sketch]: the positional is the sketch, not the board.
		if len(rest) > 1 {
			fmt.Fprintln(w, "Usage: mujard build --board <id> [sketch|folder]")
			return 2
		}
		if len(rest) == 1 {
			sketchPath = rest[0]
		}
	default:
		if len(rest) > 0 {
			board = rest[0]
		}
		if len(rest) > 1 {
			sketchPath = rest[1]
		}
	}
	if board == "" || len(rest) > 2 {
		fmt.Fprintln(w, "Usage: mujard build <board> [sketch|folder]   e.g. mujard build nano blink.ino")
		fmt.Fprintln(w, "Or:    mujard build <file.muj> [--board id] [--show-cpp]   (a Mujard program is compiled here first)")
		fmt.Fprintln(w, "With no sketch, the .ino in the current folder is built. \"mujard boards\" lists the board ids.")
		return 2
	}
	if !boardArgRx.MatchString(board) {
		fmt.Fprintf(w, "%q is not a board id - \"mujard boards\" lists them (uno, nano, esp32... or a full FQBN).\n", board)
		return 2
	}
	if *offline {
		return buildInoOffline(board, sketchPath, *backend, *out, asJSON, w)
	}
	// The direct backend (mujard-build) needs no key and no server: gather the files, try it
	// first, and only fall through to the server when the board is unsupported. "direct" is a
	// hard request - unsupported is reported, not retried. A compile error is handled and
	// never retried elsewhere either.
	var files map[string]string
	var baseName string
	if *backend == "direct" || *backend == "auto" {
		f, b, gerr := gatherSketch(sketchPath)
		if gerr != nil {
			fmt.Fprintf(w, "%s\n", gerr)
			return 2
		}
		if cerr := checkFiles(f); cerr != nil {
			fmt.Fprintf(w, "%s\n", cerr)
			return 2
		}
		files, baseName = f, b
		if code, ok := buildDirectFiles(board, files, baseName, *out, false, w); ok {
			return code
		}
		if *backend == "direct" {
			fmt.Fprintf(w, "The direct backend cannot build for %s yet (no board manifest or toolchain package). Use --backend auto or arduino.\n", board)
			return 1
		}
	}
	if cl.needKey(w) {
		return 2
	}
	if files == nil {
		var gerr error
		files, baseName, gerr = gatherSketch(sketchPath)
		if gerr != nil {
			fmt.Fprintf(w, "%s\n", gerr)
			return 2
		}
		if cerr := checkFiles(files); cerr != nil {
			fmt.Fprintf(w, "%s\n", cerr)
			return 2
		}
	}
	return serverBuild(cl, board, files, baseName, *out, *wait, asJSON, w)
}

// buildInoOffline is "mujard build <board> [sketch] --offline": the direct backend, then arduino-cli for a full FQBN
// (a short id needs the board table only a .muj build has), and nothing ever goes to the server.
func buildInoOffline(board, sketchPath, backend, out string, asJSON bool, w io.Writer) int {
	bad := func(msg string) int {
		if asJSON {
			return localFail(w, "usage_error", msg, nil, "")
		}
		fmt.Fprintln(w, msg)
		return 2
	}
	files, baseName, err := gatherSketch(sketchPath)
	if err != nil {
		return bad(err.Error())
	}
	if err := checkFiles(files); err != nil {
		return bad(err.Error())
	}
	if backend != "arduino" {
		if code, ok := buildDirectFiles(board, files, baseName, out, asJSON, w); ok {
			return code
		}
	}
	if cli := arduinoCli(); cli != "" && backend != "direct" && fqbnRx.MatchString(board) {
		return arduinoBuild(cli, board, board, files, nil, out, baseName, asJSON, w)
	}
	return localUnavailable(w, asJSON, board)
}

// serverBuild is the part of "mujard build" that runs on the server: post the files, poll, fetch the artifact.
// A .ino sketch and the C++ a .muj program was compiled into both go through here unchanged.
func serverBuild(cl *apiClient, board string, files map[string]string, baseName, outPath string, wait time.Duration, asJSON bool, w io.Writer) int {
	if !asJSON {
		fmt.Fprintf(w, "Building %d file(s) for %s...\n", len(files), board)
	}
	status, job, err := cl.call("POST", "/builds",
		map[string]any{"board": board, "files": files},
		map[string]string{"Idempotency-Key": idemKey()})
	if err != nil {
		return buildFail(w, asJSON, "server_error", "The build could not be sent: "+short(err))
	}
	if status != 200 && status != 202 {
		return buildFail(w, asJSON, str(errDoc(job), "code"), apiErrText(job))
	}
	id := str(job, "id")

	// Poll until the build reaches a terminal state.
	deadline := time.Now().Add(wait)
	last := ""
	for str(job, "state") == "queued" || str(job, "state") == "running" {
		if str(job, "state") != last && !asJSON {
			last = str(job, "state")
			fmt.Fprintf(w, "  %s...\n", last)
		}
		if time.Now().After(deadline) {
			return buildFail(w, asJSON, "timeout", fmt.Sprintf("still %s after %s - the build may finish later (id %s)", last, wait, id))
		}
		time.Sleep(1500 * time.Millisecond)
		status, job, err = cl.call("GET", "/builds/"+id, nil, nil)
		if err != nil {
			return buildFail(w, asJSON, "server_error", "Lost the build while waiting: "+short(err))
		}
		if status != 200 {
			return buildFail(w, asJSON, str(errDoc(job), "code"), apiErrText(job))
		}
	}

	state := str(job, "state")
	result, _ := job["result"].(map[string]any)
	if state != "done" || result == nil || result["success"] != true {
		msg := "the build did not produce firmware"
		if result != nil {
			if m := str(result, "message"); m != "" {
				msg = m
			}
		}
		if !asJSON {
			printBuildErrors(w, result)
			if state != "done" {
				fmt.Fprintf(w, "Build %s.\n", state)
			}
		}
		return buildFail(w, asJSON, "build_failed", msg)
	}

	// Artifact: download and save it.
	ext := "hex"
	if f := str(result, "format"); f == "bin" || f == "uf2" {
		ext = f
	}
	if outPath == "" {
		outPath = filepath.Join(".", baseName+"."+ext)
	}
	if job["has_artifact"] == true {
		status, raw, _, err := cl.callRaw("GET", "/builds/"+id+"/artifact", nil, nil)
		if err != nil {
			return buildFail(w, asJSON, "server_error", "The firmware could not be downloaded: "+short(err))
		}
		if status != 200 {
			var doc map[string]any
			_ = json.Unmarshal(raw, &doc)
			return buildFail(w, asJSON, str(errDoc(doc), "code"), apiErrText(doc))
		}
		// A test key's build carries no real firmware: the "artifact" is the HEX file's
		// end-of-file line only. Say what it is instead of writing a stub that looks like firmware.
		if strings.TrimSpace(string(raw)) == ":00000001FF" {
			if asJSON {
				return buildOK(w, map[string]any{"build": job, "artifact": nil})
			}
			fmt.Fprintf(w, "Build finished for %s but the firmware is a placeholder (a test key builds nothing for real).\n", board)
			printSizes(w, result)
			return 0
		}
		sum := sha256.Sum256(raw)
		if err := os.WriteFile(outPath, raw, 0o644); err != nil {
			return buildFail(w, asJSON, "file_error", "The build worked but the firmware could not be saved to "+outPath+": "+err.Error())
		}
		if asJSON {
			return buildOK(w, map[string]any{"build": job, "artifact": map[string]any{
				"path": outPath, "bytes": len(raw), "sha256": hex.EncodeToString(sum[:])}})
		}
		fmt.Fprintf(w, "Built %d bytes of %s for %s -> %s (sha256 %s)\n", len(raw), ext, board, outPath, hex.EncodeToString(sum[:])[:12])
	} else {
		if asJSON {
			return buildOK(w, map[string]any{"build": job})
		}
		fmt.Fprintf(w, "Build finished for %s but has no artifact (a test key builds nothing for real).\n", board)
	}
	printSizes(w, result)
	if c := num(job, "cost"); c > 0 {
		fmt.Fprintf(w, "Cost: %d coin(s).\n", c)
	}
	return 0
}

func errDoc(doc map[string]any) map[string]any {
	e, _ := doc["error"].(map[string]any)
	if e == nil {
		return map[string]any{"code": "server_error"}
	}
	return e
}

// buildFail/buildOK print the same mujard-cli.v1 envelope the desktop CLI uses.
func buildFail(w io.Writer, asJSON bool, code, msg string) int {
	if asJSON {
		fmt.Fprintf(w, `{"schema":"mujard-cli.v1","command":"build","ok":false,"error":{"code":%q,"message":%q}}`+"\n", code, msg)
	} else {
		fmt.Fprintf(w, "%s\n", msg)
	}
	return 1
}

func buildOK(w io.Writer, data map[string]any) int {
	payload := map[string]any{"schema": "mujard-cli.v1", "command": "build", "ok": true, "data": data}
	b, _ := json.Marshal(payload)
	fmt.Fprintln(w, string(b))
	return 0
}

// printBuildErrors lists the compiler's per-line diagnostics the way an IDE does.
func printBuildErrors(w io.Writer, result map[string]any) {
	if result == nil {
		return
	}
	if m := str(result, "message"); m != "" {
		fmt.Fprintln(w, m)
	}
	if errs, _ := result["errors"].([]any); len(errs) > 0 {
		for _, e := range errs {
			m, _ := e.(map[string]any)
			if m == nil {
				continue
			}
			loc := str(m, "file")
			if n := num(m, "line"); n > 0 {
				loc += fmt.Sprintf(":%d", n)
				if c := num(m, "column"); c > 0 {
					loc += fmt.Sprintf(":%d", c)
				}
			}
			fmt.Fprintf(w, "  %s  %s\n", loc, str(m, "message"))
		}
	}
}

func printSizes(w io.Writer, result map[string]any) {
	sizes, _ := result["sizes"].(map[string]any)
	if sizes == nil {
		return
	}
	parts := []string{}
	for _, k := range []string{"text", "data", "bss", "flash", "ram"} {
		if v, ok := sizes[k]; ok {
			parts = append(parts, fmt.Sprintf("%s %v", k, v))
		}
	}
	if len(parts) > 0 {
		fmt.Fprintln(w, "Sizes: "+strings.Join(parts, ", "))
	}
}
