package main

// Board-over-USB commands (devices, drivers, adapters, flash, burn, run) need the desktop
// app's flash engine - the roadmap's rule is "reused, never duplicated". These verbs find
// MujardServer.exe and pass the arguments through unchanged, so `mujard flash firmware.hex
// --board nano --port COM5` is exactly the desktop CLI's own command with its exit code.
import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// desktopExeNames are the names the desktop program has had, newest first. Installed copies
// from before a rename still answer their own CLI the same way, so any of them is usable.
var desktopExeNames = []string{"MujardServer.exe", "MujardDesktop.exe"}

// desktopExe finds the desktop program: the MUJARD_DESKTOP variable, then the PATH, then the
// repository's build output and the usual install folders.
func desktopExe() string {
	if v := strings.TrimSpace(envGet("MUJARD_DESKTOP")); v != "" {
		if _, err := os.Stat(v); err == nil {
			return v
		}
	}
	for _, name := range desktopExeNames {
		if p, err := exec.LookPath(name); err == nil {
			return p
		}
	}
	dirs := []string{}
	if root := mujardRoot(); root != "" {
		dirs = append(dirs, filepath.Join(root, "desktop-app", "bin"))
	}
	if exe, err := os.Executable(); err == nil {
		dirs = append(dirs, filepath.Dir(exe))
	}
	if la := os.Getenv("LOCALAPPDATA"); la != "" {
		dirs = append(dirs, filepath.Join(la, "Programs", "Mujard"))
	}
	dirs = append(dirs, `C:\Program Files\Mujard`)
	for _, d := range dirs {
		for _, name := range desktopExeNames {
			c := filepath.Join(d, name)
			if _, err := os.Stat(c); err == nil {
				return c
			}
		}
	}
	return ""
}

// runDesktop hands the command line to the desktop exe: its console, its output, its exit code.
func runDesktop(w io.Writer, verb string, args []string) int {
	exe := desktopExe()
	if exe == "" {
		fmt.Fprintf(w, "\"mujard %s\" is done by the desktop app (Mujard) and it is not on this machine.\n", verb)
		fmt.Fprintln(w, "Install it, or set MUJARD_DESKTOP to where MujardServer.exe is. On a server with no USB boards, the API commands (mujard build ...) are the ones to use.")
		return 1
	}
	cmdArgs := append([]string{verb}, args...)
	c := exec.Command(exe, cmdArgs...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	if err := c.Run(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return ee.ExitCode()
		}
		fmt.Fprintf(w, "The desktop app could not be started: %s\n", err)
		return 1
	}
	return 0
}

// dropList lets "devices list" / "adapters list" mean the bare verb (the desktop takes no argument there).
func dropList(args []string) []string {
	if len(args) > 0 && args[0] == "list" {
		return args[1:]
	}
	return args
}

func runDeviceList(args []string, w io.Writer) int {
	if len(args) > 0 && args[0] == "info" {
		return runDeviceInfo(args[1:], w)
	}
	return runDesktop(w, "list", dropList(args))
}
func runDeviceDrivers(args []string, w io.Writer) int { return runDesktop(w, "drivers", args) }
func runDeviceAdapters(args []string, w io.Writer) int {
	if len(args) > 0 && args[0] == "info" {
		return runAdapterInfo(args[1:], w)
	}
	return runDesktop(w, "adapters", dropList(args))
}
func runDeviceFlash(args []string, w io.Writer) int { return runDesktop(w, "flash", args) }
func runDeviceBurn(args []string, w io.Writer) int  { return runDesktop(w, "burn", args) }

// A .muj program is compiled to an Arduino sketch here first; the desktop's "run" gets the generated
// sketch folder. Anything else goes through unchanged.
func runDeviceRun(args []string, w io.Writer) int { return runMujardSketch("run", args, w) }

// `mujard serial` watches a board's serial output (and types lines back) through the desktop app's
// own port - the same watch the phone's terminal and Studio's serial panel open. Not billed: it
// reads and writes the wire, nothing is built.
func runDeviceSerial(args []string, w io.Writer) int { return runDesktop(w, "serial", args) }

// `mujard logs` lists (or prints) the upload/burn traces the desktop app keeps on this PC.
func runDeviceLogs(args []string, w io.Writer) int { return runDesktop(w, "logs", args) }

// "upload" is the mujard-style name for the desktop's "run" (build, then write to the board). It is
// delegated as "run" - never "upload" - so it works with any installed MujardServer.exe, including
// one older than the alias.
func runDeviceUpload(args []string, w io.Writer) int { return runMujardSketch("run", args, w) }

// `mujard studio` opens Mujard Studio, the code editor inside the desktop app. The desktop's own
// `studio` command asks the running copy to open its Studio window over a named pipe, or starts the
// app when none is running - so the verb works from a shell whether or not Mujard is already up.
func runDeviceStudio(args []string, w io.Writer) int {
	if len(args) != 0 {
		fmt.Fprintln(w, "'studio' takes no arguments - just: mujard studio")
		return 2
	}
	return runDesktop(w, "studio", args)
}
