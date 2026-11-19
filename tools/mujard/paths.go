package main

// Where things are on this machine: the repository root (when mujard runs inside a checkout), the program's
// own folder, and the PHP interpreter that "mujard-<name>.php" plugins run through. These are used by the
// plugin loader and the desktop finder in EVERY build; they give no access to anything by themselves.
import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// mujardRoot finds the repository root: MUJARD_ROOT first, then the first folder above the working
// directory (or the program's own folder) that looks like a full platform checkout.
func mujardRoot() string {
	if v := strings.TrimSpace(envGet("MUJARD_ROOT")); v != "" {
		return v
	}
	for _, start := range []string{mustGetwd(), exeDir()} {
		for dir := start; dir != "" && dir != filepath.Dir(dir); dir = filepath.Dir(dir) {
			if st, err := os.Stat(filepath.Join(dir, "backend", "bin")); err == nil && st.IsDir() {
				return dir
			}
		}
	}
	return ""
}

func mustGetwd() string {
	d, _ := os.Getwd()
	return d
}

func exeDir() string {
	e, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Dir(e)
}

// phpExe finds the PHP interpreter: MUJARD_PHP, the PATH, then XAMPP's usual place on Windows.
func phpExe() string {
	if v := strings.TrimSpace(envGet("MUJARD_PHP")); v != "" {
		return v
	}
	if p, err := exec.LookPath("php"); err == nil {
		return p
	}
	for _, c := range []string{`C:\xampp\php\php.exe`, `C:\php\php.exe`} {
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	return "php"
}
