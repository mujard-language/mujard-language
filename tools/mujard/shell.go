package main

// The interactive prompt: "mujard" alone (on a console) or "mujard shell" reads command lines and runs
// each one as "mujard <line>" - the same program re-run as a child, so a line goes through exactly
// the dispatcher, built-ins, plugins and exit codes a one-shot call does. There is no second
// command engine here: this file only splits lines and loops.
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"time"
)

const shellPrompt = "Mujard CLI> "

// signedIn reports whether the API commands would find a key: the MUJARD_KEY environment variable
// or the one saved by "mujard login". It answers for the shell's banner and the optional
// sign-in requirement; each command still checks the key itself.
func signedIn() bool {
	return envGet("MUJARD_KEY") != "" || loadConfig().Key != ""
}

// runShellCommand is the "shell" verb.
func runShellCommand(args []string, w io.Writer) int {
	if len(args) > 0 {
		fmt.Fprintln(w, "\"mujard shell\" takes no arguments: start it, then type commands at the Mujard CLI> prompt.")
		return 2
	}
	interactive := stdinIsConsole()
	sigc := make(chan os.Signal, 4)
	// Ctrl+C stops the running command (the child shares the console), not the shell.
	signal.Notify(sigc, os.Interrupt)
	defer signal.Stop(sigc)
	return runShell(os.Stdin, w, interactive, childRunner(interactive), sigc)
}

// childRunner runs one shell line as this same executable. Only an interactive console is handed
// to the child as stdin: from a script or pipe a confirm prompt reads end-of-input and says no,
// so nothing is ever flashed unprompted.
func childRunner(interactive bool) func(args []string) int {
	return func(args []string) int {
		exe, err := os.Executable()
		if err != nil {
			fmt.Fprintf(os.Stderr, "mujard: cannot find this program to run the command: %s\n", err)
			return 1
		}
		c := exec.Command(exe, args...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if interactive {
			c.Stdin = os.Stdin
		}
		if err := c.Run(); err != nil {
			if ee, ok := err.(*exec.ExitError); ok {
				return ee.ExitCode()
			}
			fmt.Fprintf(os.Stderr, "mujard: could not run the command: %s\n", err)
			return 1
		}
		return 0
	}
}

// runShell is the loop: prompt, read, split, run, repeat. A failed or malformed line prints and
// returns to the prompt; only "exit"/"quit" or the end of input leave, always with 0.
func runShell(in io.Reader, out io.Writer, interactive bool, run func(args []string) int, interrupts <-chan os.Signal) int {
	// On a shared or kiosk machine the owner can demand sign-in before the prompt opens:
	// MUJARD_REQUIRE_LOGIN=1 refuses the shell with no saved key. Per-command checks still apply
	// either way, so a key here is a door policy, not the authorization itself.
	if envGet("MUJARD_REQUIRE_LOGIN") == "1" && !signedIn() {
		fmt.Fprintln(out, "This command line needs a sign-in first: \"mujard login --key mjd_live_...\".")
		return 1
	}
	cancelled := func() bool {
		got := false
		for {
			select {
			case <-interrupts:
				got = true
			default:
				return got
			}
		}
	}
	if interactive {
		fmt.Fprintf(out, "Mujard command line %s - type a command (\"help\" lists them), \"exit\" to leave.\n", version)
		if !signedIn() {
			fmt.Fprintln(out, "Not signed in: cloud commands (me, usage, boards, build, ai, webhooks) need \"login\";")
			fmt.Fprintln(out, "USB commands (devices, flash, run...) work with the desktop app without it.")
		}
		fmt.Fprintln(out)
	}
	r := bufio.NewReader(in)
	first := true
	for {
		if interactive {
			fmt.Fprint(out, shellPrompt)
		}
		line, err := r.ReadString('\n')
		if err != nil && line == "" {
			// Ctrl+C at the prompt can end the console read; that is not a request to leave.
			if interactive && interrupts != nil {
				select {
				case <-interrupts:
					fmt.Fprintln(out)
					r = bufio.NewReader(in)
					continue
				case <-time.After(150 * time.Millisecond):
				}
			}
			if interactive {
				fmt.Fprintln(out)
			}
			return 0
		}
		if first {
			// A script saved by Notepad or PowerShell 5.1 starts with a UTF-8 BOM; it is not part of the first command.
			line = strings.TrimPrefix(line, "\uFEFF")
			first = false
		}
		args, perr := splitLine(strings.TrimRight(line, "\r\n"))
		switch {
		case perr != nil:
			fmt.Fprintf(out, "mujard: %s - nothing was run.\n", perr)
			continue
		case len(args) > 0 && args[0] == "mujard":
			args = args[1:] // "mujard build ..." typed out of habit
		}
		if len(args) == 0 || strings.HasPrefix(args[0], "#") {
			continue
		}
		switch strings.ToLower(args[0]) {
		case "exit", "quit":
			return 0
		case "clear", "cls":
			clearScreen(out)
			continue
		case "shell":
			fmt.Fprintln(out, "Already at the Mujard CLI> prompt.")
			continue
		}
		cancelled() // a stale Ctrl+C from the prompt doesn't belong to this command
		code := run(args)
		if cancelled() {
			fmt.Fprintln(out, "(cancelled)")
		} else if code != 0 {
			fmt.Fprintf(out, "(exit %d)\n", code)
		}
		if err != nil {
			return 0 // last line had no newline: it ran, and that was the end of the input
		}
	}
}

// splitLine breaks a line into arguments the way a user expects at a Windows prompt: spaces
// separate, "double" or 'single' quotes keep spaces together, and backslashes are ordinary
// characters so C:\My Sketches\Blink stays a path. An open quote is an error, not a guess.
func splitLine(s string) ([]string, error) {
	var args []string
	var cur strings.Builder
	var quote rune
	has := false
	for _, r := range s {
		switch {
		case quote != 0:
			if r == quote {
				quote = 0
			} else {
				cur.WriteRune(r)
			}
		case r == '"' || r == '\'':
			quote, has = r, true
		case r == ' ' || r == '\t':
			if has {
				args = append(args, cur.String())
				cur.Reset()
				has = false
			}
		default:
			cur.WriteRune(r)
			has = true
		}
	}
	if quote != 0 {
		return nil, errors.New("a " + string(quote) + " quote was opened and never closed")
	}
	if has {
		args = append(args, cur.String())
	}
	return args, nil
}

func clearScreen(out io.Writer) {
	if runtime.GOOS == "windows" && out == os.Stdout {
		c := exec.Command("cmd", "/c", "cls")
		c.Stdout = os.Stdout
		_ = c.Run()
		return
	}
	fmt.Fprint(out, "\033[H\033[2J")
}
