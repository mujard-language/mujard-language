package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"mujard/lang"
)

// mujard repl: type Mujard lines and see what they do (plan Phase 13, "on-board interpreter / REPL" - the PC half).
// Every line runs through the real compiler and the real runtime on this PC's simulated board - the same path as
// 'mujard test' - so `250 ms * 4`, int wrap-around, units and text behave exactly as on a board, never as an
// approximation. Declarations (var, const, fn, enum, struct, devices) stay for the session; a statement is replayed
// before each new line so variables keep their values; an expression's value is printed. A line with an error is
// reported with the compiler's message and forgotten.

type replSession struct {
	board string
	decls []string // top-level lines, in order
	steps []string // statements already run, replayed before each new line
	cxx   string
}

// replMarker separates what earlier steps print (replayed, not shown again) from what the new line prints.
const replMarker = "--mujard-repl-line--"

var replDecl = regexp.MustCompile(`^(var|const|saved|fn|enum|struct|use|import)\b|^[a-z_][a-z0-9_]*\s*=\s*(pin|pins|analog|Servo|Ultrasonic|Uart|battery_on)\b`)

func runRepl(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("repl", flag.ContinueOnError)
	fs.SetOutput(w)
	board := fs.String("board", "uno", "the board whose rules and pins apply")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	s := &replSession{board: *board, cxx: lang.FindCxx()}
	if s.cxx == "" {
		fmt.Fprintln(w, "mujard repl runs each line on this PC's simulated board and needs a C++ compiler: g++ or clang++ on the PATH, or MUJARD_CXX.")
		return 1
	}
	fmt.Fprintf(w, "Mujard on a simulated %s - type a line, a declaration, or a block ending in ':' (an empty line ends it). exit leaves.\n", s.board)
	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Fprint(w, "mujard> ")
		if !in.Scan() {
			fmt.Fprintln(w)
			return 0
		}
		line := in.Text()
		if strings.TrimSpace(line) == "exit" || strings.TrimSpace(line) == "quit" {
			return 0
		}
		if strings.HasSuffix(strings.TrimSpace(line), ":") { // a block: read until an empty line
			for {
				fmt.Fprint(w, "   ... ")
				if !in.Scan() || strings.TrimSpace(in.Text()) == "" {
					break
				}
				line += "\n" + in.Text()
			}
		}
		out, problem := s.eval(line)
		if problem != "" {
			fmt.Fprintln(w, problem)
			continue
		}
		if out != "" {
			fmt.Fprint(w, out)
			if !strings.HasSuffix(out, "\n") {
				fmt.Fprintln(w)
			}
		}
	}
}

// eval runs one input and returns what it printed, or the compiler's (or runtime's) complaint. Only an input that ran
// becomes part of the session.
func (s *replSession) eval(input string) (output, problem string) {
	text := strings.TrimRight(input, " \t\r\n")
	if strings.TrimSpace(text) == "" {
		return "", ""
	}
	if replDecl.MatchString(strings.TrimSpace(text)) {
		out, prob := s.run(append(append([]string{}, s.decls...), text), s.steps, "")
		if prob == "" {
			s.decls = append(s.decls, text)
		}
		return out, prob
	}
	// An expression is printed; anything else is a statement that becomes part of the session.
	if !strings.Contains(text, "\n") {
		if out, prob := s.run(s.decls, s.steps, "print "+strings.TrimSpace(text)); prob == "" {
			return out, ""
		}
	}
	out, prob := s.run(s.decls, s.steps, text)
	if prob == "" {
		s.steps = append(s.steps, text)
	}
	return out, prob
}

// run builds the session as a program with one test - the replayed steps, the marker, then the new line - and runs it.
func (s *replSession) run(decls, steps []string, line string) (string, string) {
	var b strings.Builder
	fmt.Fprintf(&b, "board %s\n", s.board)
	for _, d := range decls {
		b.WriteString(d + "\n")
	}
	b.WriteString("test \"repl\":\n")
	indent := func(code string) {
		for _, l := range strings.Split(code, "\n") {
			b.WriteString("    " + l + "\n")
		}
	}
	for _, st := range steps {
		indent(st)
	}
	indent(fmt.Sprintf("print %q", replMarker))
	newAt := strings.Count(b.String(), "\n") + 1
	if line != "" {
		indent(line)
	}
	src := b.String()
	sk, diags := lang.CompileTests(src, "repl.muj", "")
	if e := lang.FirstError(diags); e != nil || sk == nil {
		if e == nil {
			return "", "That can't run on its own."
		}
		if line != "" && e.Pos.Line >= newAt && strings.Contains(line, "\n") {
			return "", fmt.Sprintf("%s  %s (line %d of what you typed)", e.Code, e.Message, e.Pos.Line-newAt+1)
		}
		return "", fmt.Sprintf("%s  %s", e.Code, e.Message)
	}
	results, err := lang.RunTests(sk, s.cxx, 60*time.Second) // a fresh test program can start slowly on a busy PC
	if err != nil {
		return "", "Couldn't run it on the simulated board: " + err.Error()
	}
	if len(results) != 1 {
		return "", "Couldn't run it on the simulated board."
	}
	r := results[0]
	out := r.Output
	if i := strings.Index(out, replMarker); i >= 0 {
		out = strings.TrimLeft(out[i+len(replMarker):], "\r\n")
	}
	if !r.Passed && r.Message != "" {
		return out, "The board stopped: " + r.Message
	}
	return out, ""
}
