//go:build !windows

package main

import "os"

// stdinIsConsole: a terminal, not a pipe or a file.
func stdinIsConsole() bool {
	st, err := os.Stdin.Stat()
	return err == nil && st.Mode()&os.ModeCharDevice != 0
}
