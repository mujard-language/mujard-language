package main

// The "direct" build backend: mujard-build (tools/mujard-build) compiles a sketch with
// the toolchain packages in toolchain/, no arduino-cli involved. Handled=false means the
// backend is absent or the board is unsupported - the caller then falls back (auto) or
// reports it (direct).
import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// mujardBuildExe finds the mujard-build binary: MUJARD_BUILD, the PATH, beside this
// executable, then the repo's tools/mujard-build build output.
func mujardBuildExe() string {
	if v := strings.TrimSpace(envGet("MUJARD_BUILD")); v != "" {
		if _, err := os.Stat(v); err == nil {
			return v
		}
	}
	if p, err := exec.LookPath("mujard-build"); err == nil {
		return p
	}
	for _, c := range []string{filepath.Join(exeDir(), "mujard-build.exe"), filepath.Join(exeDir(), "mujard-build")} {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	if r := mujardRoot(); r != "" {
		if p := filepath.Join(r, "tools", "mujard-build", "mujard-build.exe"); fileExists(p) {
			return p
		}
	}
	return ""
}

func fileExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}

// directResult mirrors the mujard-build.v1 result schema (the fields the CLI reports).
type directResult struct {
	Success     bool   `json:"success"`
	Unsupported bool   `json:"unsupported"`
	Board       string `json:"board"`
	Format      string `json:"format"`
	Artifact    string `json:"artifact"`
	Sizes       struct {
		Flash, FlashMax, Ram, RamMax int
	} `json:"sizes"`
	Errors []struct {
		File     string `json:"file"`
		Line     int    `json:"line"`
		Col      int    `json:"col"`
		Severity string `json:"severity"`
		Message  string `json:"message"`
	} `json:"errors"`
	Log string `json:"log"`
}

// buildDirectFiles writes files to a temp sketch dir and runs mujard-build on them.
// handled=false asks the caller to use the next backend; a real compile error is handled
// (reported, never silently retried elsewhere).
func buildDirectFiles(board string, files map[string]string, baseName, out string, asJSON bool, w io.Writer) (code int, handled bool) {
	mb := mujardBuildExe()
	if mb == "" {
		return 0, false
	}
	tmp, err := os.MkdirTemp("", "mujard-direct")
	if err != nil {
		return 0, false
	}
	defer os.RemoveAll(tmp)
	sketchDir, err := writeSketchDir(tmp, files)
	if err != nil {
		return 0, false
	}
	jobFile := filepath.Join(tmp, "job.json")
	resFile := filepath.Join(tmp, "result.json")
	job := map[string]any{
		"schema":     "mujard-build.v1",
		"board":      board,
		"sketchDir":  sketchDir,
		"outDir":     filepath.Join(tmp, "out"),
		"timeoutSec": 1800, // a cold ESP32 core build needs minutes even when parallel
	}
	data, _ := json.Marshal(job)
	if err := os.WriteFile(jobFile, data, 0o644); err != nil {
		return 0, false
	}

	if !asJSON {
		fmt.Fprintf(w, "Compiling for %s with mujard-build (direct backend)...\n", board)
	}
	cmd := exec.Command(mb, "--job", jobFile, "--result", resFile)
	cmd.Env = directEnv()
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	if err := cmd.Run(); err != nil {
		if raw, rerr := os.ReadFile(resFile); rerr == nil {
			var res directResult
			if json.Unmarshal(raw, &res) == nil {
				return reportDirect(&res, board, baseName, out, asJSON, w)
			}
		}
		// The tool failed without a result - treat as absent so auto can fall back.
		return 0, false
	}
	raw, err := os.ReadFile(resFile)
	if err != nil {
		return 0, false
	}
	var res directResult
	if err := json.Unmarshal(raw, &res); err != nil {
		return 0, false
	}
	return reportDirect(&res, board, baseName, out, asJSON, w)
}

// reportDirect turns a mujard-build result into CLI output. handled=false only for
// "unsupported" (unknown board, no package, missing recipe/tool).
func reportDirect(res *directResult, board, baseName, out string, asJSON bool, w io.Writer) (int, bool) {
	if res.Unsupported {
		return 0, false
	}
	if asJSON {
		return reportDirectJSON(res, board, baseName, out, w), true
	}
	if !res.Success {
		for _, e := range res.Errors {
			loc := e.File
			if e.Line > 0 {
				loc += fmt.Sprintf(":%d", e.Line)
				if e.Col > 0 {
					loc += fmt.Sprintf(":%d", e.Col)
				}
			}
			fmt.Fprintf(w, "  %s  %s\n", loc, e.Message)
		}
		if len(res.Errors) == 0 {
			fmt.Fprintln(w, tail(res.Log, 20))
		}
		fmt.Fprintln(w, "The direct build failed.")
		return 1, true
	}
	ext := res.Format
	if ext == "" {
		ext = "hex"
	}
	if out == "" {
		out = baseName + "." + ext
	}
	data, err := os.ReadFile(res.Artifact)
	if err != nil || os.WriteFile(out, data, 0o644) != nil {
		fmt.Fprintf(w, "The build worked but the firmware could not be saved to %s\n", out)
		return 1, true
	}
	fmt.Fprintf(w, "Built %s for %s -> %s\n", filepath.Base(res.Artifact), board, out)
	if res.Sizes.Flash > 0 || res.Sizes.Ram > 0 {
		fmt.Fprintf(w, "Sizes: flash %d of %d, ram %d of %d\n", res.Sizes.Flash, res.Sizes.FlashMax, res.Sizes.Ram, res.Sizes.RamMax)
	}
	return 0, true
}

// reportDirectJSON is reportDirect for --format json: the envelope, with the compiler's messages located in the
// project's own files.
func reportDirectJSON(res *directResult, board, baseName, out string, w io.Writer) int {
	if !res.Success {
		diags := make([]localDiag, 0, len(res.Errors))
		for _, e := range res.Errors {
			sev := e.Severity
			if sev != "warning" {
				sev = "error"
			}
			diags = append(diags, localDiag{File: filepath.Base(strings.ReplaceAll(e.File, "\\", "/")), Line: e.Line, Column: e.Col, Severity: sev, Message: e.Message})
		}
		if len(diags) == 0 {
			diags = gccDiags(res.Log)
		}
		msg := "The build failed."
		if len(diags) == 0 {
			// No compiler message to point at (a tool that failed, a timeout): say what the log says went wrong.
			if line := lastProblemLine(res.Log); line != "" {
				msg = "The build failed: " + line
			}
		}
		return localFail(w, "build_failed", msg, diags, tail(res.Log, 40))
	}
	ext := res.Format
	if ext == "" {
		ext = "hex"
	}
	if out == "" {
		out = baseName + "." + ext
	}
	data, err := os.ReadFile(res.Artifact)
	if err != nil || os.WriteFile(out, data, 0o644) != nil {
		return localFail(w, "file_error", "The build worked but the firmware could not be saved to "+out, nil, "")
	}
	sizes := map[string]int{}
	if res.Sizes.Flash > 0 || res.Sizes.Ram > 0 {
		sizes = map[string]int{"flash": res.Sizes.Flash, "flashMax": res.Sizes.FlashMax, "ram": res.Sizes.Ram, "ramMax": res.Sizes.RamMax}
	}
	return localOK(w, "direct", board, ext, out, data, sizes)
}

// problemRx is a log line that says what went wrong: a tool's error, a missing file, a timeout. Command lines are not
// ("-Werror=return-type" is a flag), so lines starting with a path to an .exe or with "===" are skipped.
var problemRx = regexp.MustCompile(`(?i)\b(error|failed|fatal|cannot|can't|not found|no such file|permission denied|timed out)\b`)

// lastProblemLine is the last log line naming a problem, trimmed to fit one notice; "" when there is none.
func lastProblemLine(log string) string {
	lines := strings.Split(strings.ReplaceAll(log, "\r", ""), "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		l := strings.TrimSpace(lines[i])
		if l == "" || strings.HasPrefix(l, "===") || strings.Contains(strings.ToLower(l), ".exe ") || !problemRx.MatchString(l) {
			continue
		}
		if len(l) > 200 {
			l = l[:200] + "…"
		}
		return l
	}
	return ""
}

// directEnv points mujard-build at the toolchain dir when the env vars are not set: the
// repo's (running from a checkout works with zero setup), else the one installed beside
// this program (the desktop app ships toolchain\ next to MujardCLI.exe).
func directEnv() []string {
	env := os.Environ()
	set := func(k, v string) {
		if os.Getenv(k) == "" && v != "" {
			env = append(env, k+"="+v)
		}
	}
	tc := ""
	if r := mujardRoot(); r != "" {
		tc = filepath.Join(r, "toolchain")
	} else if d := exeDir(); d != "" && fileExists(filepath.Join(d, "toolchain", "index.json")) {
		tc = filepath.Join(d, "toolchain")
	}
	if tc != "" {
		set("MUJARD_BOARDS", filepath.Join(tc, "boards"))
		set("MUJARD_PACKAGES", filepath.Join(tc, "packages"))
		set("MUJARD_TOOLCHAIN_INDEX", filepath.Join(tc, "index.json"))
	}
	// This dev machine's borrowed AVR toolchain (see toolchain/packages/arduino-avr.json).
	set("MUJARD_ARDUINO_DATA", firstExisting(`C:\arduino-toolchain\data`, filepath.Join(os.Getenv("LOCALAPPDATA"), "Arduino15")))
	return env
}

func firstExisting(paths ...string) string {
	for _, p := range paths {
		if fileExists(p) {
			return p
		}
	}
	return ""
}
