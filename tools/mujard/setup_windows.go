package main

import (
	"syscall"
	"unsafe"
)

const (
	regExpandSz     = 2 // REG_EXPAND_SZ: the usual type of a user PATH
	hwndBroadcast   = 0xFFFF
	wmSettingChange = 0x001A
	smtoAbortIfHung = 0x0002
)

var (
	advapi32           = syscall.NewLazyDLL("advapi32.dll")
	procRegSetValueExW = advapi32.NewProc("RegSetValueExW")
	user32             = syscall.NewLazyDLL("user32.dll")
	procSendMessageTOW = user32.NewProc("SendMessageTimeoutW")
)

// regUserPath is HKEY_CURRENT_USER\Environment\Path, read and written raw (no %VAR% expansion).
type regUserPath struct{}

func userPathStore() pathStore { return regUserPath{} }

func openEnvKey(access uint32) (syscall.Handle, error) {
	var h syscall.Handle
	err := syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, syscall.StringToUTF16Ptr("Environment"), 0, access, &h)
	return h, err
}

func (regUserPath) get() (string, uint32, error) {
	h, err := openEnvKey(syscall.KEY_QUERY_VALUE)
	if err != nil {
		return "", 0, err
	}
	defer syscall.RegCloseKey(h)
	name := syscall.StringToUTF16Ptr("Path")
	var kind, n uint32
	if err := syscall.RegQueryValueEx(h, name, nil, &kind, nil, &n); err != nil {
		if err == syscall.ERROR_FILE_NOT_FOUND {
			return "", regExpandSz, nil // no user PATH yet
		}
		return "", 0, err
	}
	if n == 0 {
		return "", kind, nil
	}
	buf := make([]uint16, (n+1)/2+1)
	if err := syscall.RegQueryValueEx(h, name, nil, &kind, (*byte)(unsafe.Pointer(&buf[0])), &n); err != nil {
		return "", 0, err
	}
	if kind != syscall.REG_SZ && kind != regExpandSz {
		kind = regExpandSz
	}
	return syscall.UTF16ToString(buf), kind, nil
}

func (regUserPath) set(value string, kind uint32) error {
	h, err := openEnvKey(syscall.KEY_SET_VALUE)
	if err != nil {
		return err
	}
	defer syscall.RegCloseKey(h)
	data, err := syscall.UTF16FromString(value)
	if err != nil {
		return err
	}
	if kind != syscall.REG_SZ {
		kind = regExpandSz
	}
	r, _, _ := procRegSetValueExW.Call(uintptr(h), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Path"))), 0,
		uintptr(kind), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)*2))
	if r != 0 {
		return syscall.Errno(r)
	}
	// Tell Explorer (and so every terminal opened from now on) that the environment changed.
	var result uintptr
	procSendMessageTOW.Call(hwndBroadcast, wmSettingChange, 0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Environment"))), smtoAbortIfHung, 5000, uintptr(unsafe.Pointer(&result)))
	return nil
}
