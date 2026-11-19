package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"mujard/lang"
)

// runMujardTest is "mujard test file.muj": run the program's 'test' blocks on this PC against a simulated board (spec
// 12) - each test on a fresh board, with simulated time - and say which passed. It needs a C++ compiler (MUJARD_CXX,
// or g++ / clang++ on the PATH); nothing is uploaded and no board is involved. Exit 0 when every test passes.
func runMujardTest(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.SetOutput(w)
	format := fs.String("format", "", "\"json\" prints the mujard-cli.v1 document")
	verbose := fs.Bool("verbose", false, "also show what the program printed during each test")
	plan := fs.Bool("plan", false, "emit the tests' plans for the emulated-board driver (sim-web), without running them")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	asJSON := *format == "json"
	if *format != "" && !asJSON {
		fmt.Fprintln(w, "--format is \"json\" or nothing.")
		return 2
	}
	if len(rest) != 1 || !isMujardPath(rest[0]) {
		fmt.Fprintln(w, "Usage: mujard test <file.muj> [--verbose] [--format json]")
		return 2
	}
	path := rest[0]
	fail := func(code, msg string, diags []localDiag, exit int) int {
		if asJSON {
			if diags != nil {
				localFail(w, code, msg, diags, "")
				return exit
			}
			return envFail(w, "test", code, msg, exit)
		}
		fmt.Fprintln(w, msg)
		return exit
	}
	src, err := os.ReadFile(path)
	if err != nil {
		return fail("usage_error", "Cannot read "+path, nil, 2)
	}
	sk, diags := lang.CompileTests(string(src), filepath.Base(path), filepath.Dir(path))
	if asJSON {
		if list, errs := langDiags(path, diags); errs > 0 {
			return fail("compile_failed", fmt.Sprintf("%s does not compile: %d error(s).", filepath.Base(path), errs), list, 1)
		}
	} else if errs, _ := printDiags(w, path, diags); errs > 0 {
		fmt.Fprintf(w, "%s does not compile - fix the error(s) above first.\n", filepath.Base(path))
		return 1
	}
	if sk == nil {
		return fail("compile_failed", filepath.Base(path)+" does not compile (no board declared?).", nil, 1)
	}
	if len(sk.Tests) == 0 {
		return fail("no_tests", filepath.Base(path)+" has no test blocks (test \"name\": ...).", nil, 1)
	}
	if *plan {
		// The plans are compile-only output for the sim-web driver; no C++ build needed.
		if asJSON {
			envOK(w, "testplan", map[string]any{"board": sk.Board.Name, "file": filepath.Base(path), "tests": sk.Tests})
		} else {
			for _, t := range sk.Tests {
				fmt.Fprintf(w, "test %q: %d step(s)%s\n", t.Name, len(t.Plan.Steps), map[bool]string{true: "", false: " (not emulatable)"}[t.Plan.Emulatable])
			}
		}
		return 0
	}
	cxx := lang.FindCxx()
	if cxx == "" {
		return fail("no_compiler", "Tests run on this PC and need a C++ compiler: install g++ (on Windows, for example w64devkit) "+
			"or set MUJARD_CXX to its path.", nil, 1)
	}
	if !asJSON {
		fmt.Fprintf(w, "Running %d test(s) of %s on a simulated %s...\n", len(sk.Tests), filepath.Base(path), sk.Board.Display)
	}
	results, err := lang.RunTests(sk, cxx, 60*time.Second)
	if err != nil {
		return fail("test_build_failed", err.Error(), nil, 1)
	}
	failed := 0
	for _, r := range results {
		if !r.Passed {
			failed++
		}
	}
	if asJSON {
		b := map[string]any{"file": filepath.Base(path), "board": sk.Board.Name, "passed": len(results) - failed, "failed": failed, "tests": results}
		if failed > 0 {
			envOK(w, "test", b)
			return 1
		}
		return envOK(w, "test", b)
	}
	for _, r := range results {
		if r.Passed {
			fmt.Fprintf(w, "  PASS  %s\n", r.Name)
		} else {
			where := ""
			if r.FailLine > 0 {
				where = fmt.Sprintf(" (%s:%d)", filepath.Base(path), r.FailLine)
			}
			fmt.Fprintf(w, "  FAIL  %s%s: %s\n", r.Name, where, r.Message)
		}
		if (*verbose || !r.Passed) && strings.TrimSpace(r.Output) != "" {
			for _, l := range strings.Split(strings.TrimRight(r.Output, "\n"), "\n") {
				if !strings.HasPrefix(l, "\x1f") { // the app protocol's frames aren't printed text
					fmt.Fprintf(w, "        | %s\n", l)
				}
			}
		}
	}
	fmt.Fprintf(w, "%d passed, %d failed.\n", len(results)-failed, failed)
	if failed > 0 {
		return 1
	}
	return 0
}
