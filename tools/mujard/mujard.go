package main

// The .muj side of the command line: the Mujard language (tools/mujard/lang compiles it to Arduino C++).
//   mujard check file.muj            lex, parse, check - no build, no charge, no key needed
//   mujard build file.muj            compile, then build: locally with arduino-cli when one is installed,
//                                  else on the server with the developer key (the generated C++ goes
//                                  through the same build endpoint a .ino sketch uses)
//   mujard run file.muj              generate the sketch and hand it to the desktop app's "run"
//   mujard fmt file.muj              the standard layout, in place (lang.Format is shared with the editors)
//   mujard new <name>                 a first project from a template
import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"mujard/lang"
)

func isMujardPath(p string) bool { return lang.IsSource(p) }

// printDiags lists diagnostics the way a compiler does: file:line:col code message, errors first in place.
// Diagnostis from imported modules carry their own file name ("motors.muj:4:1 ...").
func printDiags(w io.Writer, path string, diags []lang.Diagnostic) (errors, warnings int) {
	for _, d := range diags {
		name := filepath.Base(path)
		if d.File != "" {
			name = d.File
		}
		fmt.Fprintf(w, "%s:%d:%d  %s  %s\n", name, d.Pos.Line, d.Pos.Col, d.Code, d.Message)
		if d.IsWarning() {
			warnings++
		} else {
			errors++
		}
	}
	return
}

// checkMujardFile is "mujard check file.muj": the same pipeline the compiler runs, reported and done.
func checkMujardFile(path, format string, tree bool, w io.Writer) int {
	asJSON := format == "json"
	bad := func(code, msg string, exit int) int {
		if asJSON {
			return envFail(w, "check", code, msg, exit)
		}
		fmt.Fprintln(w, msg)
		return exit
	}
	src, err := os.ReadFile(path)
	if err != nil {
		return bad("usage_error", "Cannot read "+path, 2)
	}
	res := lang.CheckProgramIn(string(src), filepath.Dir(path))
	f, diags := res.File, res.Diags
	if tree && f != nil && !asJSON {
		fmt.Fprint(w, lang.Dump(f))
	}
	if asJSON {
		list := []map[string]any{}
		errs := 0
		for _, d := range diags {
			sev := "error"
			if d.IsWarning() {
				sev = "warning"
			} else {
				errs++
			}
			file := filepath.Base(path)
			if d.File != "" {
				file = d.File
			}
			list = append(list, map[string]any{"file": file, "line": d.Pos.Line, "column": d.Pos.Col,
				"code": d.Code, "severity": sev, "message": d.Message})
		}
		// secrets: the names the project's secrets must hold for a build (names only, never values).
		secrets := res.Secrets
		if secrets == nil {
			secrets = []string{}
		}
		data := map[string]any{"file": filepath.Base(path), "diagnostics": list, "secrets": secrets}
		if errs > 0 {
			b, _ := json.Marshal(map[string]any{"schema": "mujard-cli.v1", "command": "check", "ok": false, "data": data,
				"error": map[string]any{"code": "check_failed", "message": fmt.Sprintf("%d error(s) in %s", errs, filepath.Base(path))}})
			fmt.Fprintln(w, string(b))
			return 1
		}
		return envOK(w, "check", data)
	}
	errs, warns := printDiags(w, path, diags)
	switch {
	case errs > 0:
		fmt.Fprintf(w, "%d error(s), %d warning(s) - not checked clean.\n", errs, warns)
		return 1
	case warns > 0:
		fmt.Fprintf(w, "No errors, %d warning(s).\n", warns)
	default:
		fmt.Fprintln(w, filepath.Base(path)+" checks clean.")
	}
	return 0
}

// compileMujard reads a .muj file and compiles it. boardName, when given, must be the board the program
// declares - it is a guard against building the wrong file, not an override (the board is part of the source).
func compileMujard(path, boardName string, w io.Writer) (*lang.Sketch, int) {
	return compileMujardAs(path, boardName, false, w)
}

// compileMujardAs is compileMujard; with asJSON its failures are one mujard-cli.v1 envelope with the checker's messages.
func compileMujardAs(path, boardName string, asJSON bool, w io.Writer) (*lang.Sketch, int) {
	fail := func(code, msg string, diags []localDiag, exit int) (*lang.Sketch, int) {
		if asJSON {
			localFail(w, code, msg, diags, "")
		} else {
			fmt.Fprintln(w, msg)
		}
		return nil, exit
	}
	src, err := os.ReadFile(path)
	if err != nil {
		return fail("usage_error", "Cannot read "+path, nil, 2)
	}
	sk, diags := lang.CompileIn(string(src), filepath.Base(path), filepath.Dir(path))
	if asJSON {
		if list, errs := langDiags(path, diags); errs > 0 {
			return fail("compile_failed", fmt.Sprintf("%s does not compile: %d error(s).", filepath.Base(path), errs), list, 1)
		}
	} else if errs, _ := printDiags(w, path, diags); errs > 0 {
		fmt.Fprintf(w, "%s does not compile - fix the error(s) above (\"mujard check %s\" explains them).\n", filepath.Base(path), path)
		return nil, 1
	}
	if sk == nil || sk.Board == nil {
		return fail("compile_failed", filepath.Base(path)+" does not compile (no board declared?).", nil, 1)
	}
	// The --board name is normalized through the board table first, so an upload profile
	// ("nano_old" = a Nano at a slower upload baud) matches the program's "board nano".
	want := boardName
	if lb := lang.LookupBoard(boardName); lb != nil {
		want = lb.Name
	}
	if want != "" && !strings.EqualFold(want, sk.Board.Name) {
		return fail("usage_error", fmt.Sprintf("%s is for the %s, not %q - fix the 'board' line or drop --board.", path, sk.Board.Display, boardName), nil, 2)
	}
	return sk, 0
}

// writeSketchDir puts a generated sketch on disk as <dir>/sketch/sketch.ino (+ its helper files): the folder
// and the .ino share a name because that is what arduino-cli and the desktop's gather both ask for.
// secretsHeaderName is the header a build gets its secrets from (the server's ProjectConfig::SECRETS_HEADER).
const secretsHeaderName = "mujard_secrets.h"

// localSecretsHeader is the secrets header for a local build: each name the program needs, from the environment
// variable MUJARD_SECRET_<NAME>, escaped like the server does it (no character can end the C string). Names without a
// value are left out, so the program's #error says which one is missing.
func localSecretsHeader(names []string) string {
	var b strings.Builder
	b.WriteString("// Made by mujard build for this build only, from MUJARD_SECRET_* - never saved with the project.\n#pragma once\n")
	for _, name := range names {
		v := os.Getenv("MUJARD_SECRET_" + name)
		if v == "" {
			continue
		}
		appendSecretDefine(&b, name, v)
	}
	return b.String()
}

// appendSecretDefine writes one `#define NAME "value"` escaped like the server does it (no character
// can end the C string).
func appendSecretDefine(b *strings.Builder, name, v string) {
	b.WriteString("#define " + name + " \"")
	for i := 0; i < len(v); i++ {
		c := v[i]
		switch {
		case c == '\\' || c == '"' || c == '?':
			b.WriteByte('\\')
			b.WriteByte(c)
		case c < 0x20 || c >= 0x7f:
			fmt.Fprintf(b, "\\%03o", c)
		default:
			b.WriteByte(c)
		}
	}
	b.WriteString("\"\n")
}

// notifyDefines asks the server for this account's board-notify key and returns the
// MUJARD_NOTIFY_URL/KEY defines a notify-over-Wi-Fi program gets in its secrets header (Phase 9).
// "" when the program doesn't notify or has no Wi-Fi, or when no key can be fetched - the serial
// line still reaches a connected app.
func notifyDefines(sk *lang.Sketch, cl *apiClient, label string) string {
	if cl == nil || cl.key == "" {
		return ""
	}
	prog, hasNotify := sk.Files["mujard_program.cpp"]
	if !hasNotify || !strings.Contains(prog, "app_notify(") || sk.Files["mujard_wifi.h"] == "" {
		return ""
	}
	root := strings.TrimSuffix(cl.base, "/api/v1") // the API folder itself (notify.php lives there)
	req, err := http.NewRequest("POST", root+"/notify.php",
		bytes.NewReader([]byte(`{"action":"issue","label":`+strconv.Quote(label)+`}`)))
	if err != nil {
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cl.key)
	resp, err := (&http.Client{Timeout: 20 * time.Second}).Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	var doc struct {
		Key string `json:"key"`
		URL string `json:"url"`
	}
	if json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(&doc) != nil || doc.Key == "" || doc.URL == "" {
		return ""
	}
	var b strings.Builder
	appendSecretDefine(&b, "MUJARD_NOTIFY_URL", doc.URL)
	appendSecretDefine(&b, "MUJARD_NOTIFY_KEY", doc.Key)
	return b.String()
}

func writeSketchDir(dir string, files map[string]string) (string, error) {
	sk := filepath.Join(dir, "sketch")
	if err := os.MkdirAll(sk, 0o755); err != nil {
		return "", err
	}
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(sk, name), []byte(content), 0o644); err != nil {
			return "", err
		}
	}
	return sk, nil
}

// arduinoCli finds a local arduino-cli (MUJARD_ARDUINO_CLI first, then the PATH); "" when there is none.
func arduinoCli() string {
	if v := strings.TrimSpace(envGet("MUJARD_ARDUINO_CLI")); v != "" {
		if _, err := os.Stat(v); err == nil {
			return v
		}
	}
	if p, err := exec.LookPath("arduino-cli"); err == nil {
		return p
	}
	return ""
}

// buildLocal compiles the generated sketch with the machine's arduino-cli and copies the artifact to out.
func buildLocal(cli string, sk *lang.Sketch, out, baseName, extra string, asJSON bool, w io.Writer) int {
	// The program's secrets come from MUJARD_SECRET_<NAME>, into this temporary build copy only (never the .gen
	// folder); a missing one stops the build with the program's own message saying which. extra carries build-time
	// defines like MUJARD_NOTIFY_* (already escaped; "" when there are none).
	header := localSecretsHeader(sk.Secrets) + extra
	return arduinoBuild(cli, sk.FQBN, sk.Board.Display, sk.Files, &header, out, baseName, asJSON, w)
}

// withSecretsHeader is the generated sketch plus the build's secrets header: the program's secret() values from
// MUJARD_SECRET_<NAME> and build-time defines (extra: MUJARD_NOTIFY_*), empty when there are none - every build
// writes it, because the runtime includes it plainly (the server refuses sketches that probe for files) - for a
// local direct build (arduinoBuild writes the same header itself).
func withSecretsHeader(sk *lang.Sketch, extra string) map[string]string {
	files := make(map[string]string, len(sk.Files)+1)
	for k, v := range sk.Files {
		files[k] = v
	}
	files[secretsHeaderName] = localSecretsHeader(sk.Secrets) + extra
	return files
}

// findArtifact picks the firmware file arduino-cli leaves in an --output-dir: the sketch's .hex, .uf2 or
// .bin (the merged .bin, which carries the bootloader and partitions too, wins over the bare app .bin).
func findArtifact(dir string) (path, ext string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", ""
	}
	best, rank := "", -1
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		n := e.Name()
		r := -1
		switch {
		case strings.HasSuffix(n, ".ino.hex"):
			r = 4
		case strings.HasSuffix(n, ".ino.uf2"):
			r = 3
		case strings.HasSuffix(n, ".ino.merged.bin"), strings.HasSuffix(n, ".ino.merged-fw.bin"):
			r = 2
		case strings.HasSuffix(n, ".ino.bin"):
			r = 1
		}
		if r > rank {
			best, rank = filepath.Join(dir, n), r
		}
	}
	if best == "" {
		return "", ""
	}
	for _, x := range []string{".hex", ".uf2", ".bin"} {
		if strings.HasSuffix(best, x) {
			return best, x[1:]
		}
	}
	return best, "bin"
}

func tail(s string, lines int) string {
	ls := strings.Split(strings.TrimRight(s, "\n"), "\n")
	if len(ls) > lines {
		ls = ls[len(ls)-lines:]
	}
	return strings.Join(ls, "\n")
}

// buildMujard is "mujard build file.muj": compile, optionally show the C++, then build - directly with
// mujard-build when asked (--backend direct|auto), locally when an arduino-cli is installed, else on the
// server with the developer key.
func buildMujard(cl *apiClient, path, boardName, out string, showCpp bool, rest []string, wait time.Duration, asJSON bool, backend string, offline bool, w io.Writer) int {
	if len(rest) > 0 {
		fmt.Fprintln(w, "Usage: mujard build <file.muj> [--board id] [--show-cpp] [--out f] [--offline]")
		return 2
	}
	// Offline (this PC only, e.g. Mujard Studio with no internet): its own backends, never the server - and the
	// answer is JSON all the way when asked, compile errors included.
	if offline {
		if showCpp {
			fmt.Fprintln(w, "--show-cpp and --offline don't go together (mujard gen shows the C++).")
			return 2
		}
		sk, code := compileMujardAs(path, boardName, asJSON, w)
		if sk == nil {
			return code
		}
		baseName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		if backend != "arduino" {
			if code, ok := buildDirectFiles(sk.Board.Name, withSecretsHeader(sk, ""), baseName, out, asJSON, w); ok {
				return code
			}
		}
		if cli := arduinoCli(); cli != "" && backend != "direct" {
			return buildLocal(cli, sk, out, baseName, "", asJSON, w)
		}
		return localUnavailable(w, asJSON, sk.Board.Display)
	}
	sk, code := compileMujard(path, boardName, w)
	if sk == nil {
		return code
	}
	baseName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	if showCpp {
		dir := filepath.Join(filepath.Dir(path), baseName+".gen")
		if err := os.MkdirAll(dir, 0o755); err != nil {
			fmt.Fprintf(w, "Cannot write %s: %s\n", dir, err)
			return 1
		}
		for name, content := range sk.Files {
			if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
				fmt.Fprintf(w, "Cannot write %s: %s\n", name, err)
				return 1
			}
		}
		fmt.Fprintf(w, "Generated C++ for review: %s\n", dir)
	}
	// A notify over Wi-Fi reaches the account even with no app connected: the server hands this build a
	// per-user key, baked in as a secret like the program's own (notify.php; "" when it can't be fetched).
	extra := notifyDefines(sk, cl, baseName)
	// Backend selection: "direct" is a hard request (unsupported is an answer, not a retry);
	// "auto" uses mujard-build when the board resolves, then the normal path.
	switch backend {
	case "direct":
		if code, ok := buildDirectFiles(sk.Board.Name, withSecretsHeader(sk, extra), baseName, out, false, w); ok {
			return code
		}
		fmt.Fprintf(w, "The direct backend cannot build for the %s yet (no board manifest or toolchain package). Use --backend auto or arduino.\n", sk.Board.Display)
		return 1
	case "auto":
		if code, ok := buildDirectFiles(sk.Board.Name, withSecretsHeader(sk, extra), baseName, out, false, w); ok {
			return code
		}
	}
	if cli := arduinoCli(); cli != "" {
		return buildLocal(cli, sk, out, baseName, extra, false, w)
	}
	if cl.needKey(w) {
		return 2
	}
	return serverBuild(cl, sk.Board.Name, sk.Files, baseName, out, wait, asJSON, w)
}

// runMujardSketch is "mujard run file.muj": the generated sketch goes to a temp folder and the desktop app's own
// "run" (build on the server, then upload) takes it from there - the same path a .ino takes.
func runMujardSketch(verb string, args []string, w io.Writer) int {
	idx := -1
	for i, a := range args {
		if isMujardPath(a) {
			idx = i
			break
		}
	}
	if idx < 0 {
		return runDesktop(w, verb, args)
	}
	sk, code := compileMujard(args[idx], "", w)
	if sk == nil {
		return code
	}
	tmp, err := os.MkdirTemp("", "mujard-run")
	if err != nil {
		fmt.Fprintf(w, "Cannot make a temp folder: %s\n", err)
		return 1
	}
	defer os.RemoveAll(tmp)
	dir, err := writeSketchDir(tmp, sk.Files)
	if err != nil {
		fmt.Fprintf(w, "Cannot write the generated sketch: %s\n", err)
		return 1
	}
	fwd := append([]string{}, args...)
	fwd[idx] = dir
	if !hasFlagValue(fwd, "--board", "-b") {
		fwd = append(fwd, "--board", sk.Board.Name)
	}
	fmt.Fprintf(w, "Compiled %s; handing the sketch to Mujard...\n", filepath.Base(args[idx]))
	return runDesktop(w, verb, fwd)
}

// hasFlagValue reports whether args contain --board (or -b) in either spelling.
func hasFlagValue(args []string, names ...string) bool {
	for _, a := range args {
		for _, n := range names {
			if a == n || strings.HasPrefix(a, n+"=") {
				return true
			}
		}
	}
	return false
}

// runMujardGen is "mujard gen file.muj [--out dir]": compile, then write the generated Arduino sketch
// (sketch.ino, mujard_program.cpp, mujard_rt.*) flat into a folder. Nothing is built - 'mujard build' does
// that. This is also the step a build server runs before arduino-cli.
func runMujardGen(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("gen", flag.ContinueOnError)
	fs.SetOutput(w)
	out := fs.String("out", "", "folder the generated sketch is written to (default <name>.gen)")
	board := fs.String("board", "", "the board the program must be for (a check, not an override)")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	if len(rest) != 1 || !isMujardPath(rest[0]) {
		fmt.Fprintln(w, "Usage: mujard gen <file.muj> [--out dir] [--board id]")
		return 2
	}
	sk, code := compileMujard(rest[0], *board, w)
	if sk == nil {
		return code
	}
	dir := *out
	if dir == "" {
		base := strings.TrimSuffix(filepath.Base(rest[0]), filepath.Ext(rest[0]))
		dir = filepath.Join(filepath.Dir(rest[0]), base+".gen")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		fmt.Fprintf(w, "Cannot write %s: %s\n", dir, err)
		return 1
	}
	for name, content := range sk.Files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			fmt.Fprintf(w, "Cannot write %s: %s\n", name, err)
			return 1
		}
	}
	// The runtime includes the build's secrets header plainly; a build writes the real one over this placeholder (the
	// server does, after this very step). A placeholder never holds a value, so the folder compiles as it is.
	if p := filepath.Join(dir, secretsHeaderName); !fileExists(p) {
		placeholder := "// The build writes this header with the program's secret() values; none are kept in generated files.\n#pragma once\n"
		if err := os.WriteFile(p, []byte(placeholder), 0o644); err != nil {
			fmt.Fprintf(w, "Cannot write %s: %s\n", secretsHeaderName, err)
			return 1
		}
	}
	fmt.Fprintf(w, "Generated %d file(s) in %s for the %s.\n", len(sk.Files), dir, sk.Board.Display)
	return 0
}

// ---- fmt ---------------------------------------------------------------------------------------------

// runMujardFmt is "mujard fmt": rewrite .muj files in the standard layout. --check only reports which files
// would change (exit 1 if any); --print writes the formatted text to stdout instead of the file.
func runMujardFmt(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("fmt", flag.ContinueOnError)
	fs.SetOutput(w)
	check := fs.Bool("check", false, "report files that are not in the standard layout; changes nothing")
	printIt := fs.Bool("print", false, "print the formatted source instead of writing it back")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	var files []string
	for _, p := range rest {
		st, err := os.Stat(p)
		switch {
		case err != nil:
			fmt.Fprintf(w, "Cannot read %s\n", p)
			return 2
		case st.IsDir():
			es, _ := os.ReadDir(p)
			for _, e := range es {
				if !e.IsDir() && isMujardPath(e.Name()) {
					files = append(files, filepath.Join(p, e.Name()))
				}
			}
			if len(files) == 0 {
				fmt.Fprintf(w, "No .muj or .mjd files in %s\n", p)
				return 2
			}
		case isMujardPath(p):
			files = append(files, p)
		default:
			fmt.Fprintf(w, "%s is not a .muj or .mjd file.\n", p)
			return 2
		}
	}
	if len(files) == 0 {
		fmt.Fprintln(w, "Usage: mujard fmt <file.muj> [more.muj | folder] [--check | --print]")
		return 2
	}
	if *check && *printIt {
		fmt.Fprintln(w, "--check and --print don't combine: check reads, print writes.")
		return 2
	}
	rc := 0
	for _, p := range files {
		b, err := os.ReadFile(p)
		if err != nil {
			fmt.Fprintf(w, "Cannot read %s\n", p)
			return 2
		}
		src := string(b)
		out, diags := lang.Format(src)
		if len(diags) > 0 {
			printDiags(w, p, diags)
			fmt.Fprintf(w, "%s can't be read cleanly - fix the error(s); fmt refuses to guess.\n", p)
			rc = 1
			continue
		}
		switch {
		case *check:
			if out != src {
				fmt.Fprintf(w, "%s is not in the standard layout.\n", p)
				rc = 1
			}
		case *printIt:
			fmt.Fprint(w, out)
		case out == src:
			fmt.Fprintf(w, "%s already is.\n", p)
		default:
			if err := os.WriteFile(p, []byte(out), 0o644); err != nil {
				fmt.Fprintf(w, "Cannot write %s: %s\n", p, err)
				rc = 1
				continue
			}
			fmt.Fprintf(w, "%s formatted.\n", p)
		}
	}
	return rc
}

// ---- describe ----------------------------------------------------------------------------------------

// runMujardDescribe is "mujard describe <library folder> [--out file.mlib]": read an installed Arduino
// library's headers and write the '.mlib' description that makes 'use "Name"' check its calls. Put the
// file next to the program (or in MUJARD_LIB_DESCS on a build server) to use it.
func runMujardDescribe(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("describe", flag.ContinueOnError)
	fs.SetOutput(w)
	out := fs.String("out", "", "where the .mlib goes (default <library name>.mlib here)")
	printIt := fs.Bool("print", false, "print the description instead of writing it")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	if len(rest) != 1 {
		fmt.Fprintln(w, "Usage: mujard describe <library folder> [--out file.mlib | --print]")
		return 2
	}
	d, err := lang.DescribeLibrary(rest[0])
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return 1
	}
	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return 1
	}
	if *printIt {
		fmt.Fprintln(w, string(b))
		return 0
	}
	name := *out
	if name == "" {
		name = strings.ToLower(strings.Map(func(r rune) rune {
			if r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
				return r
			}
			return '_'
		}, d.Name)) + ".mlib"
	}
	if err := os.WriteFile(name, append(b, '\n'), 0o644); err != nil {
		fmt.Fprintf(w, "Cannot write %s: %s\n", name, err)
		return 1
	}
	fmt.Fprintf(w, "Wrote %s: %d class(es), %d function(s), %d constant(s) from %s.\n",
		name, len(d.Classes), len(d.Functions), len(d.Constants), d.Name)
	return 0
}

// ---- new ---------------------------------------------------------------------------------------------

var mujardNameRx = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_\-]{0,39}$`)

// runMujardNew is "mujard new <name>": a folder holding <name>.muj with a blink program for the board.
func runMujardNew(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("new", flag.ContinueOnError)
	fs.SetOutput(w)
	board := fs.String("board", "uno", "the board the program is for ("+strings.Join(lang.BoardNames(), ", ")+")")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	if len(rest) != 1 || !mujardNameRx.MatchString(rest[0]) {
		fmt.Fprintln(w, "Usage: mujard new <name> [--board uno]   e.g. mujard new blink --board esp32")
		fmt.Fprintln(w, "The name starts with a letter; letters, digits, _ and - after it.")
		return 2
	}
	b := lang.LookupBoard(*board)
	if b == nil {
		fmt.Fprintf(w, "Unknown board %q - try one of: %s.\n", *board, strings.Join(lang.BoardNames(), ", "))
		return 2
	}
	name := rest[0]
	dir := filepath.Join(".", name)
	if _, err := os.Stat(dir); err == nil {
		fmt.Fprintf(w, "%s already exists - pick another name or remove it first.\n", dir)
		return 1
	}
	src := fmt.Sprintf(`# %s - a Mujard program (the language guide is docs/lang/README.md)
board %s

led = pin LED as output

every 500 ms:
    led.toggle()
`, name, b.Name)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		fmt.Fprintf(w, "Cannot create %s: %s\n", dir, err)
		return 1
	}
	file := filepath.Join(dir, name+".muj")
	if err := os.WriteFile(file, []byte(src), 0o644); err != nil {
		fmt.Fprintf(w, "Cannot write %s: %s\n", file, err)
		return 1
	}
	fmt.Fprintf(w, "Created %s for the %s.\n", file, b.Display)
	fmt.Fprintf(w, "Next: mujard check %s     or     mujard run %s --port <port>\n", file, file)
	return 0
}
