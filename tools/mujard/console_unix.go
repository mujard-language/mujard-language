//go:build unix

package main

import "os"

// stdinIsConsole: a terminal, not a pipe, a file or /dev/null (a char device too, but not a console).
func stdinIsConsole() bool {
	st, err := os.Stdin.Stat()
	if err != nil || st.Mode()&os.ModeCharDevice == 0 {
		return false
	}
	dn, err := os.Open(os.DevNull)
	if err != nil {
		return true
	}
	defer dn.Close()
	dst, err := dn.Stat()
	return err != nil || !os.SameFile(st, dst)
}
