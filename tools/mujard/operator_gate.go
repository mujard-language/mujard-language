package main

// The operator tier: commands that work on, or reveal, the server's internals - maintenance scripts that
// can rewrite every table ("mujard admin migrate"), load testing, and the ops checks that print env-file
// paths, secret presence and worker/gateway addresses.
//
// They are NOT in the program users get. The public build (the default "go build") does not contain their code at all, so no environment variable, flag or
// saved key can switch them on - there is nothing to switch. They exist only in the operator build:
//
//	go build -tags operator -o mujard-operator .
//
// which is built and kept on the server / dev machine and never distributed. Even there each command
// still refuses unless MUJARD_ADMIN=1 is set (a guard against a stray run, not a security boundary: whoever
// can run the operator build on the server can already run the scripts directly).
import (
	"fmt"
	"io"
)

// operatorOnly names the operator commands. In the public build these names still exist so that typing
// one is answered with the same refusal as before (and so help, "did you mean" and "get" can be tested to
// never offer them); in that build the handlers are not compiled in.
var operatorOnly = map[string]bool{
	"admin":    true, // runs the platform's maintenance scripts against the database
	"loadtest": true, // drives load against a server
	"doctor":   true, // prints env-file paths, secret presence/lengths, systemd units
	"status":   true, // prints worker/gateway addresses and internal health
}

// operatorRefused prints the refusal for an operator command, without naming scripts, paths or the switch.
func operatorRefused(w io.Writer, name string) {
	if operatorBuild {
		fmt.Fprintf(w, "\"mujard %s\" works on the server's internals - it is for the person operating\n", name)
		fmt.Fprintln(w, "the server, not for users.")
		fmt.Fprintln(w, "Set MUJARD_ADMIN=1 in the environment of this operator build to run it.")
		return
	}
	fmt.Fprintf(w, "\"mujard %s\" is a server maintenance command. It is not part of this program and\n", name)
	fmt.Fprintln(w, "is not available to users.")
}
