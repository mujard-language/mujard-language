//go:build !operator

package main

// The public build: no operator code is compiled in, so nothing the user types, sets or saves can run it.

const operatorBuild = false

// adminEnabled is always false here, whatever MUJARD_ADMIN says: the commands it would unlock are not in this program.
func adminEnabled() bool { return false }

// operatorCommands is empty: admin, doctor, status and loadtest have no handlers in the public build.
var operatorCommands []command
