//go:build unix

package main

import (
	"os"
	"syscall"
	"unsafe"
)

// stdinIsConsole: a terminal, not a pipe, a file or /dev/null - only a tty answers TCGETS.
func stdinIsConsole() bool {
	var t syscall.Termios
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&t)))
	return errno == 0
}
