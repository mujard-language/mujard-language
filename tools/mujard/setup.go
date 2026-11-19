package main

// "mujard setup path" puts the folder this MujardCLI.exe is in on the user's own PATH, so typing "mujard" works
// in any new terminal. It changes nothing until the user runs it, touches only the current user's
// PATH (never the machine's), appends (so it can never shadow a program already on the PATH),
// refuses when a different mujard would win anyway, and "--undo" takes the entry out again.
import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// pathStore is the user's saved PATH: the raw value (unexpanded %VARS% kept as they are) and its
// registry type, so a write puts back exactly what was there plus or minus one entry.
type pathStore interface {
	get() (value string, kind uint32, err error)
	set(value string, kind uint32) error
}

var errSetupUnsupported = errors.New("only Windows keeps a per-user PATH this command can edit")

func runSetup(args []string, w io.Writer) int {
	if len(args) == 0 || (args[0] != "path" && args[0] != "toolchain") {
		fmt.Fprintln(w, "Usage: mujard setup path [--check | --undo]      (adds the folder this mujard is in to your user PATH)")
		fmt.Fprintln(w, "       mujard setup toolchain [pack-id ...]     (install the board toolchain into Mujard's own folder)")
		return 2
	}
	if args[0] == "toolchain" {
		return setupToolchain(args[1:], w)
	}
	self, err := os.Executable()
	if err == nil {
		self, err = filepath.EvalSymlinks(self)
	}
	if err != nil {
		fmt.Fprintf(w, "Cannot tell where this mujard is: %s\n", err)
		return 1
	}
	return setupPath(args[1:], w, userPathStore(), self, exec.LookPath)
}

// toolchainBasePacks are the packs a fresh install needs for the AVR boards (Uno, Nano, Mega) - the small,
// default set. ESP32 and other families are installed by naming their pack ids (mujard-build toolchain list
// shows the index), so nobody downloads an SDK they never asked for.
var toolchainBasePacks = []string{"avr-gcc", "arduino-avr-core", "arduino-servo"}

// setupToolchain installs the pinned, checksummed toolchain packages into Mujard's own toolchain folder by
// delegating to mujard-build's package manager - so `mujard build --backend direct` works without the server
// and nothing is borrowed from another IDE's folders.
func setupToolchain(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("setup toolchain", flag.ContinueOnError)
	fs.SetOutput(w)
	listOnly := fs.Bool("list", false, "show the toolchain index instead of installing")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	mb := mujardBuildExe()
	if mb == "" {
		fmt.Fprintln(w, "mujard-build is not installed next to this mujard - install the desktop app or build tools/mujard-build first.")
		return 1
	}
	if *listOnly {
		c := exec.Command(mb, "toolchain", "list")
		c.Env = directEnv()
		c.Stdout, c.Stderr = w, w
		if err := c.Run(); err != nil {
			return 1
		}
		return 0
	}
	ids := rest
	if len(ids) == 0 {
		ids = toolchainBasePacks
	}
	fmt.Fprintf(w, "Installing %d toolchain package(s) into Mujard's folder...\n", len(ids))
	for _, id := range ids {
		c := exec.Command(mb, "toolchain", "install", id)
		c.Env = directEnv()
		c.Stdout, c.Stderr = w, w
		if err := c.Run(); err != nil {
			fmt.Fprintf(w, "%s failed to install (%s). The rest may still install - retry with: mujard setup toolchain %s\n", id, err, id)
			return 1
		}
	}
	fmt.Fprintln(w, "Toolchain ready - `mujard build <file.muj> --backend direct` now works without the server.")
	return 0
}

func setupPath(args []string, w io.Writer, store pathStore, self string, lookPath func(string) (string, error)) int {
	fs := flag.NewFlagSet("setup path", flag.ContinueOnError)
	fs.SetOutput(w)
	check := fs.Bool("check", false, "only report; change nothing")
	undo := fs.Bool("undo", false, "take this folder off your user PATH")
	if rest, err := parseArgs(fs, args); err != nil || len(rest) > 0 || (*check && *undo) {
		fmt.Fprintln(w, "Usage: mujard setup path [--check | --undo]")
		return 2
	}
	dir := filepath.Dir(self)
	lower := strings.ToLower(dir)
	if strings.Contains(lower, "go-build") || strings.HasPrefix(lower, strings.ToLower(filepath.Clean(os.TempDir()))) {
		fmt.Fprintf(w, "This mujard is running from a temporary folder (%s) - build it (go build -o MujardCLI.exe .) and run that one.\n", dir)
		return 1
	}
	value, kind, err := store.get()
	if err != nil {
		if errors.Is(err, errSetupUnsupported) {
			fmt.Fprintf(w, "Add this to your shell profile instead:\n  export PATH=\"$PATH:%s\"\n", dir)
			return 1
		}
		fmt.Fprintf(w, "Cannot read your user PATH: %s\n", err)
		return 1
	}
	listed := pathHas(value, dir)

	if *undo {
		if !listed {
			fmt.Fprintf(w, "%s is not on your user PATH - nothing to undo.\n", dir)
			return 0
		}
		next, _ := pathRemove(value, dir)
		if err := store.set(next, kind); err != nil {
			fmt.Fprintf(w, "Cannot write your user PATH: %s\n", err)
			return 1
		}
		fmt.Fprintf(w, "Took %s off your user PATH. Open a new terminal for it to apply.\n", dir)
		return 0
	}

	// A different mujard already found first would keep winning: say so instead of adding a dead entry.
	other := ""
	if found, err := lookPath("mujard"); err == nil {
		if abs, err := filepath.Abs(found); err == nil && !samePath(filepath.Dir(abs), dir) {
			other = abs
		}
	}
	if *check {
		fmt.Fprintf(w, "this mujard:        %s\n", self)
		fmt.Fprintf(w, "on user PATH:    %v\n", listed)
		if other != "" {
			fmt.Fprintf(w, "\"mujard\" runs now:  %s (a different one)\n", other)
		}
		return 0
	}
	if other != "" {
		fmt.Fprintf(w, "Another mujard comes first on your PATH: %s\nRemove or rename it first - adding %s would never be used. Nothing was changed.\n", other, dir)
		return 1
	}
	if listed {
		fmt.Fprintf(w, "%s is already on your user PATH. In a new terminal, \"mujard\" runs it.\n", dir)
		return 0
	}
	if err := store.set(pathAdd(value, dir), kind); err != nil {
		fmt.Fprintf(w, "Cannot write your user PATH: %s\n", err)
		return 1
	}
	fmt.Fprintf(w, "Added %s to your user PATH. Open a new terminal and type \"mujard\". (\"mujard setup path --undo\" takes it off.)\n", dir)
	return 0
}

// pathEntries splits a PATH value; empty entries are dropped.
func pathEntries(value string) []string {
	var out []string
	for _, e := range strings.Split(value, ";") {
		if strings.TrimSpace(e) != "" {
			out = append(out, e)
		}
	}
	return out
}

// expandPercent resolves %NAME% the way Windows does for a PATH entry (unknown names stay as written).
func expandPercent(s string) string {
	var b strings.Builder
	for {
		i := strings.IndexByte(s, '%')
		if i < 0 {
			break
		}
		j := strings.IndexByte(s[i+1:], '%')
		if j < 0 {
			break
		}
		name := s[i+1 : i+1+j]
		if v, ok := os.LookupEnv(name); ok && name != "" {
			b.WriteString(s[:i] + v)
		} else {
			b.WriteString(s[:i+2+j])
		}
		s = s[i+2+j:]
	}
	b.WriteString(s)
	return b.String()
}

func samePath(a, b string) bool {
	clean := func(p string) string {
		p = strings.Trim(strings.TrimSpace(expandPercent(p)), `"`)
		return strings.ToLower(strings.TrimRight(filepath.Clean(p), `\/`))
	}
	return clean(a) == clean(b)
}

func pathHas(value, dir string) bool {
	for _, e := range pathEntries(value) {
		if samePath(e, dir) {
			return true
		}
	}
	return false
}

func pathAdd(value, dir string) string {
	v := strings.TrimRight(value, "; ")
	if v == "" {
		return dir
	}
	return v + ";" + dir
}

func pathRemove(value, dir string) (string, bool) {
	var keep []string
	removed := false
	for _, e := range pathEntries(value) {
		if samePath(e, dir) {
			removed = true
			continue
		}
		keep = append(keep, e)
	}
	return strings.Join(keep, ";"), removed
}
