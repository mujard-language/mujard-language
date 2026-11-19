//go:build !windows

package main

type noUserPath struct{}

func userPathStore() pathStore { return noUserPath{} }

func (noUserPath) get() (string, uint32, error) { return "", 0, errSetupUnsupported }
func (noUserPath) set(string, uint32) error     { return errSetupUnsupported }
