package main

// Building on this PC only. "mujard build ... --offline" never sends anything to the server: it uses the direct backend
// (mujard-build with the SHA-pinned toolchain packages) and then arduino-cli, and says plainly when neither can build the
// board here. With --format json a local build answers in the same mujard-cli.v1 envelope as a server build, plus the
// compiler's diagnostics with file names relative to the project - this is what Mujard Studio in the desktop app reads.
import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"mujard/lang"
)

// fqbnRx is a full board name for arduino-cli ("arduino:avr:uno"); short ids need the board table of a .muj build.
var fqbnRx = regexp.MustCompile(`^[A-Za-z0-9_.\-]+(:[A-Za-z0-9_.\-]+){2}$`)

// localDiag is one compiler message, located in the project's own files.
type localDiag struct {
	File     string `json:"file"`
	Line     int    `json:"line"`
	Column   int    `json:"column"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
	Code     string `json:"code,omitempty"`
}

// localOK reports firmware built on this PC: where it was saved, its size and hash, and the sizes the compiler gave.
func localOK(w io.Writer, backend, board, format, path string, data []byte, sizes map[string]int) int {
	sum := sha256.Sum256(data)
	d := map[string]any{"backend": backend, "board": board, "format": format,
		"artifact": map[string]any{"path": path, "bytes": len(data), "sha256": hex.EncodeToString(sum[:])}}
	if len(sizes) > 0 {
		d["sizes"] = sizes
	}
	return envOK(w, "build", d)
}

// localFail is a failed local build: the reason, plus the compiler's messages and the end of its log when there are any.
func localFail(w io.Writer, code, msg string, diags []localDiag, log string) int {
	if diags == nil {
		diags = []localDiag{}
	}
	b, _ := json.Marshal(map[string]any{"schema": "mujard-cli.v1", "command": "build", "ok": false,
		"data":  map[string]any{"errors": diags, "log": log},
		"error": map[string]any{"code": code, "message": msg}})
	fmt.Fprintln(w, string(b))
	return 1
}

// localUnavailable: --offline, and nothing on this PC can build this board.
func localUnavailable(w io.Writer, asJSON bool, board string) int {
	msg := fmt.Sprintf("This PC can't build for %s without the internet yet: install its offline build pack (Mujard Studio, Hardware & build), or install arduino-cli - or build online.", board)
	if asJSON {
		return localFail(w, "no_local_toolchain", msg, nil, "")
	}
	fmt.Fprintln(w, msg)
	return 1
}

// gccDiagRx is a gcc/g++ message: "<path>:<line>:<col>: error: <text>" (the column is optional; Windows paths keep their drive).
var gccDiagRx = regexp.MustCompile(`(?m)^(.+?):(\d+):(?:(\d+):)? (fatal error|error|warning): (.+?)\r?$`)

// gccDiags picks the compiler's messages out of a build log. Paths become the file's own name (sketch.ino, motors.muj -
// the generated C++ carries #line markers back to the .muj lines), so the editor can open them.
func gccDiags(log string) []localDiag {
	var out []localDiag
	seen := map[string]bool{}
	for _, m := range gccDiagRx.FindAllStringSubmatch(log, 200) {
		line, _ := strconv.Atoi(m[2])
		col, _ := strconv.Atoi(m[3])
		sev := "error"
		if m[4] == "warning" {
			sev = "warning"
		}
		d := localDiag{File: filepath.Base(strings.ReplaceAll(m[1], "\\", "/")), Line: line, Column: col, Severity: sev, Message: strings.TrimSpace(m[5])}
		key := fmt.Sprintf("%s:%d:%d:%s", d.File, d.Line, d.Column, d.Message)
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, d)
	}
	return out
}

// langDiags turns the Mujard checker's messages into localDiag (file names as the project has them).
func langDiags(path string, diags []lang.Diagnostic) (list []localDiag, errors int) {
	for _, d := range diags {
		sev := "error"
		if d.IsWarning() {
			sev = "warning"
		} else {
			errors++
		}
		file := filepath.Base(path)
		if d.File != "" {
			file = d.File
		}
		list = append(list, localDiag{File: file, Line: d.Pos.Line, Column: d.Pos.Col, Severity: sev, Message: d.Message, Code: d.Code})
	}
	return
}

// arduinoBuild compiles a sketch with this machine's arduino-cli and saves the artifact to out. secretsHeader, when not nil,
// is written beside the sketch as mujard_secrets.h for this build only (a .muj program's secret() values).
func arduinoBuild(cli, fqbn, display string, files map[string]string, secretsHeader *string, out, baseName string, asJSON bool, w io.Writer) int {
	fail := func(code, msg string, diags []localDiag, log string) int {
		if asJSON {
			return localFail(w, code, msg, diags, log)
		}
		fmt.Fprintln(w, msg)
		if log != "" {
			fmt.Fprintln(w, log)
		}
		return 1
	}
	tmp, err := os.MkdirTemp("", "mujard-build")
	if err != nil {
		return fail("local_error", "Cannot make a temp folder: "+err.Error(), nil, "")
	}
	defer os.RemoveAll(tmp)
	sketchDir, err := writeSketchDir(tmp, files)
	if err != nil {
		return fail("local_error", "Cannot write the sketch for the build: "+err.Error(), nil, "")
	}
	if secretsHeader != nil {
		if err := os.WriteFile(filepath.Join(sketchDir, secretsHeaderName), []byte(*secretsHeader), 0o600); err != nil {
			return fail("local_error", "Cannot write the secrets for this build: "+err.Error(), nil, "")
		}
	}
	outDir := filepath.Join(tmp, "out")
	args := []string{"compile", "--fqbn", fqbn, "--output-dir", outDir}
	if cfg := strings.TrimSpace(envGet("MUJARD_ARDUINO_CONFIG")); cfg != "" {
		args = append(args, "--config-file", cfg)
	}
	args = append(args, sketchDir)
	if !asJSON {
		fmt.Fprintf(w, "Compiling the generated sketch for %s with arduino-cli...\n", display)
	}
	raw, err := exec.Command(cli, args...).CombinedOutput()
	sizes := buildSizes(string(raw))
	if err != nil {
		// a program that doesn't fit gets a plain reason instead of the linker's
		if why := tooBig(sizes, display); why != "" {
			if asJSON {
				return localFail(w, "too_big", why, nil, tail(string(raw), 20))
			}
			fmt.Fprintln(w, why)
			return 1
		}
		if asJSON {
			return localFail(w, "build_failed", "The build failed.", gccDiags(string(raw)), tail(string(raw), 40))
		}
		fmt.Fprintf(w, "arduino-cli could not build it (%s):\n%s\n", err, tail(string(raw), 30))
		return 1
	}
	art, ext := findArtifact(outDir)
	if art == "" {
		return fail("build_failed", "arduino-cli finished but produced no .hex/.bin/.uf2", nil, tail(string(raw), 20))
	}
	if out == "" {
		out = baseName + "." + ext
	}
	data, err := os.ReadFile(art)
	if err != nil || os.WriteFile(out, data, 0o644) != nil {
		return fail("file_error", "The build worked but the firmware could not be saved to "+out, nil, "")
	}
	if asJSON {
		return localOK(w, "arduino-cli", fqbn, ext, out, data, sizes)
	}
	fmt.Fprintf(w, "Built %s for %s -> %s\n", filepath.Base(art), display, out)
	for _, l := range memoryLines(sizes, display) {
		fmt.Fprintln(w, l)
	}
	return 0
}
