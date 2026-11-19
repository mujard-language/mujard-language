package lang

// Tests (spec 12): a program's 'test' blocks run on the PC, against a fake Arduino with simulated time
// (runtime/tests/shim), one fresh board per test. GenerateTests turns the program into a test build - the program as
// it would run on the board, plus each test as a task beside it - and RunTests compiles that with the PC's C++ compiler
// and runs every test. Nothing here ever reaches a board: firmware builds leave test blocks out.

import (
	"bufio"
	"context"
	"embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed runtime/tests/shim/*.h
var shimFS embed.FS

// TestInfo is one 'test' block of a program.
type TestInfo struct {
	Name string   `json:"name"`
	Line int      `json:"line"`
	Plan TestPlan `json:"plan,omitempty"` // the test's steps for the emulated-board driver (sim-web)
}

// TestResult is what one test did.
type TestResult struct {
	Name     string `json:"name"`
	Line     int    `json:"line"`
	Passed   bool   `json:"passed"`
	Message  string `json:"message,omitempty"`  // why it failed
	FailLine int    `json:"failLine,omitempty"` // the .muj line of the check that failed (0: not a line of the test)
	Output   string `json:"output,omitempty"`   // what the program printed during the test
}

// CompileTests checks a program and generates its test build (no test blocks: an empty Tests list).
func CompileTests(src, srcName, dir string) (*Sketch, []Diagnostic) {
	f, diags := ParseSource(src)
	if FirstError(diags) != nil {
		return nil, sortDiags(diags)
	}
	info, cdiags := analyzeWith(f, moduleLoader(dir), dirEnv(dir))
	diags = append(diags, cdiags...)
	if FirstError(diags) != nil || info.Board == nil {
		return nil, sortDiags(diags)
	}
	sk, gdiags := generateSketch(f, info, srcName, true)
	diags = append(diags, gdiags...)
	if FirstError(diags) != nil {
		return nil, sortDiags(diags)
	}
	return sk, sortDiags(diags)
}

// testBlock emits one test as a task that runs beside the program, when the run selects it.
func (g *gen) testBlock(d *SpecialBlock) {
	k := len(g.tests)
	g.tests = append(g.tests, TestInfo{Name: d.Name, Line: d.At.Line, Plan: g.testPlan(d)})
	ctx := &taskCtx{seen: map[string]bool{}}
	var body section
	g.inTask, g.inTestBody = ctx, true
	// power_cut() / brown_out() (spec 12) split the test: the board resets there and the run ends; RunTests starts the
	// same test on a fresh board that keeps only what was written to the EEPROM/flash, and the segments before the
	// restart are skipped, so the test goes on at the line after it.
	segs := [][]Node{nil}
	var cutAt []int
	var cutKind []string
	for _, s := range d.Body.Stmts {
		if name := restartCallName(s); name != "" {
			segs = append(segs, nil)
			cutAt = append(cutAt, s.Start().Line)
			cutKind = append(cutKind, name)
			continue
		}
		segs[len(segs)-1] = append(segs[len(segs)-1], s)
	}
	if len(segs) == 1 {
		g.block(&body, d.Body, 2)
	} else {
		for i, seg := range segs {
			body.add(0, "        if (mujard::test_resume_cut <= %d) {", i)
			if len(seg) > 0 {
				g.block(&body, &Block{At: d.Body.At, Stmts: seg}, 3)
			}
			if i < len(cutAt) {
				body.add(cutAt[i], "            mujard::test_%s(%d);", cutKind[i], i+1)
			}
			body.add(0, "        }")
		}
	}
	g.inTask, g.inTestBody = nil, false
	g.fns.add(d.At.Line, "// test %s", cString(d.Name))
	g.fns.add(d.At.Line, "struct mujard_Test_%d : mujard::Task {", k)
	for _, fl := range ctx.fields {
		g.fns.add(d.At.Line, "    %s", fl)
	}
	g.fns.add(0, "    void run() {")
	g.fns.add(0, "        MUJARD_TASK_BEGIN();")
	g.fns = append(g.fns, body...)
	g.fns.add(0, "        mujard::test_pass();")
	g.fns.add(0, "        MUJARD_TASK_END();")
	g.fns.add(0, "    }")
	g.fns.add(0, "} mujard_test_%d;", k)
	g.loop.add(d.At.Line, "    if (mujard::test_selected == %d) mujard_test_%d.run();", k, k)
}

// restartCallName is the function name of a test statement that resets the board ("" for anything else).
func restartCallName(n Node) string {
	s, ok := n.(*ExprStmt)
	if !ok {
		return ""
	}
	c, ok := s.X.(*Call)
	if !ok || len(c.Args) != 0 {
		return ""
	}
	id, ok := c.Fn.(*Ident)
	if !ok || (id.Name != "power_cut" && id.Name != "brown_out") {
		return ""
	}
	return id.Name
}

// advanceTime is the time of an 'advance(t)' statement in a test, as C++ ("" for anything else).
func (g *gen) advanceTime(n Node) string {
	c, ok := n.(*Call)
	if !ok || !g.inTestBody || g.inTask == nil || len(c.Args) != 1 {
		return ""
	}
	if id, ok := c.Fn.(*Ident); !ok || id.Name != "advance" {
		return ""
	}
	return g.expr(c.Args[0].Value)
}

// pretendAnalog sets what an analog pin reads: a raw reading, a voltage, or a share of the range.
func (g *gen) pretendAnalog(obj string, v Node) string {
	t := g.typeOf(v)
	switch {
	case t.K == TInt:
		return fmt.Sprintf("mujard::test_pretend_analog(%s.pin(), (int32_t)(%s))", obj, g.expr(v))
	case t.K == TQty && t.Dim == "voltage":
		return fmt.Sprintf("mujard::test_pretend_analog(%s.pin(), %s.rawFor((float)(%s)))", obj, obj, g.expr(v))
	case t.K == TQty && t.Dim == "ratio":
		return fmt.Sprintf("mujard::test_pretend_analog(%s.pin(), %s.rawForRatio((float)(%s)))", obj, obj, g.expr(v))
	}
	return g.later(v.Start(), "pretend with "+describe(t), phaseLater)
}

// FindCxx finds the PC's C++ compiler for tests: MUJARD_CXX first, then g++, c++ or clang++ on the PATH ("" if none).
func FindCxx() string {
	if v := strings.TrimSpace(os.Getenv("MUJARD_CXX")); v != "" {
		if _, err := os.Stat(v); err == nil {
			return v
		}
	}
	for _, name := range []string{"g++", "c++", "clang++"} {
		if p, err := exec.LookPath(name); err == nil {
			return p
		}
	}
	return ""
}

// adcDefines make the PC's analog pins read like the board's (the runtime's ADC range and reference), and give
// board.voltage the nominal supply of the boards whose real runtime can measure it.
func adcDefines(b *Board) []string {
	var out []string
	switch b.Arch {
	case "avr":
		out = append(out, "-DMUJARD_TEST_ADC_MAX=1023.0f", "-DMUJARD_TEST_VREF=5.0f")
	case "esp32":
		out = append(out, "-DMUJARD_TEST_ADC_MAX=4095.0f", "-DMUJARD_TEST_VREF=3.3f")
	}
	if b.CanMeasureVoltage() {
		voltage := "3.3f" // Pico-family, nRF52840, SAMD21, STM32 and RA4M1 test boards model their 3.3 V rail
		if b.IsAVR() {
			voltage = "5.0f"
		}
		out = append(out, "-DMUJARD_TEST_BOARD_VOLTAGE="+voltage)
	}
	if b.CanMeasureFreeMemory() {
		out = append(out, "-DMUJARD_TEST_FREE_MEMORY=2048")
	}
	return out
}

// RunTests compiles a test build with the C++ compiler cxx and runs each test in a fresh simulated board. Each run may
// take up to perTest of real time (an hour of simulated time is the test's own limit).
func RunTests(sk *Sketch, cxx string, perTest time.Duration) ([]TestResult, error) {
	if len(sk.Tests) == 0 {
		return nil, nil
	}
	if len(sk.Libraries) > 0 {
		var missing []string
		for _, l := range sk.Libraries {
			// A library runs on the PC only when the shim has a fake for its header.
			if _, err := shimFS.Open("runtime/tests/shim/" + l.Header); err != nil {
				missing = append(missing, l.Name)
			}
		}
		if len(missing) > 0 {
			return nil, fmt.Errorf("tests can't run on the PC yet: no PC fake for %s", strings.Join(missing, ", "))
		}
	}
	tmp, err := os.MkdirTemp("", "mujard-test")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmp)
	var sources []string
	for name, content := range sk.Files {
		switch {
		case name == "sketch.ino":
			name = "sketch.cpp"
		case strings.HasSuffix(name, ".json"):
			continue
		}
		if err := os.WriteFile(filepath.Join(tmp, name), []byte(content), 0o644); err != nil {
			return nil, err
		}
		if strings.HasSuffix(name, ".cpp") {
			sources = append(sources, name)
		}
	}
	sort.Strings(sources)
	if len(sk.Secrets) > 0 {
		// a test never sees a real secret: each one is a placeholder
		var h strings.Builder
		h.WriteString("#pragma once\n")
		for _, n := range sk.Secrets {
			fmt.Fprintf(&h, "#define %s \"test-%s\"\n", n, strings.ToLower(n))
		}
		if err := os.WriteFile(filepath.Join(tmp, secretsHeader), []byte(h.String()), 0o644); err != nil {
			return nil, err
		}
	}
	shimDir := filepath.Join(tmp, "shim")
	if err := os.MkdirAll(shimDir, 0o755); err != nil {
		return nil, err
	}
	entries, _ := shimFS.ReadDir("runtime/tests/shim")
	for _, e := range entries {
		b, _ := shimFS.ReadFile("runtime/tests/shim/" + e.Name())
		if err := os.WriteFile(filepath.Join(shimDir, e.Name()), b, 0o644); err != nil {
			return nil, err
		}
	}
	exe := filepath.Join(tmp, "tests")
	if runtime.GOOS == "windows" {
		exe += ".exe"
	}
	args := append([]string{"-std=c++11", "-O1", "-w", "-DMUJARD_HOST_TEST", "-Ishim", "-I.", "-o", exe}, adcDefines(sk.Board)...)
	args = append(args, sources...)
	build := exec.Command(cxx, args...)
	build.Dir = tmp
	if out, err := build.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("the tests couldn't be built on this PC (%v):\n%s", err, tailLines(string(out), 25))
	}
	var results []TestResult
	for k, t := range sk.Tests {
		r := TestResult{Name: t.Name, Line: t.Line}
		ctx, cancel := context.WithTimeout(context.Background(), perTest)
		// A power_cut() ends a run with "MUJARD-TEST restart <cut>": the same test starts again on a fresh board with
		// what survived (the state file), after that cut. What each run printed is kept, in order.
		state := filepath.Join(tmp, fmt.Sprintf("power%d.bin", k))
		var out []byte
		var err error
		before := ""
		for cut, runs := 0, 0; ; runs++ {
			args := []string{strconv.Itoa(k), state}
			if cut > 0 {
				args = append(args, strconv.Itoa(cut))
			}
			out, err = exec.CommandContext(ctx, exe, args...).Output()
			i := strings.LastIndex(string(out), "\nMUJARD-TEST restart ")
			if i < 0 || err != nil || runs >= 20 {
				break
			}
			mark := strings.Fields(strings.TrimSpace(string(out[i+len("\nMUJARD-TEST restart "):])))
			if len(mark) == 0 {
				break
			}
			next, _ := strconv.Atoi(mark[0])
			if next <= cut {
				break // a restart must move forward; anything else is a runner bug, not a loop
			}
			label := "power cut"
			if len(mark) > 1 {
				label = strings.ReplaceAll(mark[1], "-", " ")
			}
			before += strings.TrimRight(string(out[:i]), "\r\n") + "\n--- " + label + " ---\n"
			cut = next
		}
		timedOut := errors.Is(ctx.Err(), context.DeadlineExceeded)
		cancel()
		r.Output, r.Passed, r.FailLine, r.Message = parseTestOutput(string(out))
		r.Output = strings.TrimLeft(before+r.Output, "\n")
		switch {
		case timedOut:
			r.Passed, r.Message = false, fmt.Sprintf("the test took longer than %s of real time", perTest)
		case err != nil && r.Message == "" && !r.Passed:
			r.Message = "the test program stopped: " + err.Error()
		}
		results = append(results, r)
	}
	return results, nil
}

// parseTestOutput splits a test run's output into what the program printed and the result line.
func parseTestOutput(out string) (printed string, passed bool, line int, message string) {
	idx := strings.LastIndex(out, "\nMUJARD-TEST ")
	if idx < 0 {
		return strings.TrimSpace(out), false, 0, "the test didn't report a result"
	}
	printed = strings.TrimRight(out[:idx], "\r\n")
	sc := bufio.NewScanner(strings.NewReader(out[idx+1:]))
	sc.Scan()
	res := strings.TrimSpace(strings.TrimPrefix(sc.Text(), "MUJARD-TEST "))
	if res == "pass" {
		return printed, true, 0, ""
	}
	res = strings.TrimPrefix(res, "fail ")
	n, rest, _ := strings.Cut(res, " ")
	line, _ = strconv.Atoi(n)
	return printed, false, line, rest
}

func tailLines(s string, n int) string {
	ls := strings.Split(strings.TrimRight(s, "\r\n"), "\n")
	if len(ls) > n {
		ls = ls[len(ls)-n:]
	}
	return strings.Join(ls, "\n")
}
