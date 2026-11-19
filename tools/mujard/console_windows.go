package main

import (
	"os"
	"syscall"
)

// stdinIsConsole: a real console window, not a pipe, a file or NUL.
func stdinIsConsole() bool {
	var mode uint32
	return syscall.GetConsoleMode(syscall.Handle(os.Stdin.Fd()), &mode) == nil
}
