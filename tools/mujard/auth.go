package main

// mujard auth status | test - is a developer key set, and does the server accept it?
// "status" is local and never touches the network; "test" makes the one call "login" already makes
// (GET /v1/me). Neither prints the key: only the masked form the rest of mujard uses.
import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"strings"
)

// authInfo describes where the key and the address came from, without revealing the key.
type authInfo struct {
	SignedIn  bool   `json:"signedIn"`
	KeyMasked string `json:"key"`
	KeySource string `json:"keySource"`
	Mode      string `json:"mode"`
	API       string `json:"api"`
	APISource string `json:"apiSource"`
}

// describeAuth works out the effective key and address the same way addAPIFlags does: flag, then the
// environment, then the saved file, then the default. explicit says which flags were typed.
func describeAuth(c *apiClient, explicit map[string]bool) authInfo {
	saved := loadConfig()
	info := authInfo{SignedIn: c.key != "", KeySource: "none", APISource: "default"}
	switch {
	case explicit["key"]:
		info.KeySource = "--key"
	case envGet("MUJARD_KEY") != "":
		info.KeySource = "MUJARD_KEY"
	case saved.Key != "":
		info.KeySource = "saved file"
	}
	switch {
	case explicit["api"]:
		info.APISource = "--api"
	case envGet("MUJARD_API") != "":
		info.APISource = "MUJARD_API"
	case saved.API != "":
		info.APISource = "saved file"
	}
	if c.key != "" {
		info.KeyMasked = maskKey(c.key)
		switch {
		case strings.HasPrefix(c.key, "mjd_test_"):
			info.Mode = "test"
		case strings.HasPrefix(c.key, "mjd_live_"):
			info.Mode = "live"
		}
	}
	info.API = strings.TrimRight(c.base, "/")
	return info
}

func visitedFlags(fs *flag.FlagSet) map[string]bool {
	set := map[string]bool{}
	fs.Visit(func(f *flag.Flag) { set[f.Name] = true })
	return set
}

func runAuth(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("auth", flag.ContinueOnError)
	fs.SetOutput(w)
	cl := addAPIFlags(fs)
	format := fs.String("format", "", "\"json\" prints the mujard-cli.v1 document")
	rest, err := parseArgs(fs, args)
	if err != nil {
		return 2
	}
	sub := firstArg(rest)
	if (sub != "status" && sub != "test") || len(rest) != 1 {
		fmt.Fprintln(w, "Usage: mujard auth status   (is a key saved? no network)\n       mujard auth test     (does the server accept it?)")
		return 2
	}
	asJSON := *format == "json"
	if *format != "" && !asJSON {
		fmt.Fprintln(w, "--format is \"json\" or nothing.")
		return 2
	}
	info := describeAuth(cl, visitedFlags(fs))
	cmd := "auth " + sub
	if sub == "status" {
		if asJSON {
			code := envOK(w, cmd, info)
			if !info.SignedIn {
				return 1
			}
			return code
		}
		if !info.SignedIn {
			fmt.Fprintln(w, "Not signed in: no developer key is set.")
			fmt.Fprintln(w, "Run \"mujard login --key mjd_live_...\" once, or set MUJARD_KEY. A mjd_test_ key costs nothing.")
			return 1
		}
		mode := info.Mode
		if mode == "" {
			mode = "unknown mode"
		}
		fmt.Fprintf(w, "Signed in with a %s key %s (from %s).\nServer address: %s (from %s).\n", mode, info.KeyMasked, info.KeySource, info.API, info.APISource)
		fmt.Fprintln(w, "This checked only what is saved on this machine. \"mujard auth test\" asks the server.")
		return 0
	}
	// test
	cl.finish()
	info.API = strings.TrimRight(cl.base, "/")
	if cl.key == "" {
		if asJSON {
			return envFail(w, cmd, "no_key", "No developer key is set.", 2)
		}
		cl.needKey(w)
		return 2
	}
	status, raw, _, cerr := cl.callRaw("GET", "/me", nil, nil)
	if cerr != nil {
		if asJSON {
			return envFail(w, cmd, "unreachable", "Cannot reach the API: "+short(cerr), 1)
		}
		fmt.Fprintf(w, "Cannot reach the API: %s\n", short(cerr))
		return 1
	}
	var doc map[string]any
	_ = json.Unmarshal(raw, &doc)
	if status != 200 {
		if asJSON {
			return envFail(w, cmd, "rejected", apiErrText(doc), 1)
		}
		return failAPI(w, status, doc)
	}
	k, _ := doc["key"].(map[string]any)
	data := map[string]any{"accepted": true, "keyName": str(k, "name"), "mode": str(k, "mode"), "scopes": strList(k["scopes"]), "expiresAt": str(k, "expires_at")}
	if asJSON {
		return envOK(w, cmd, data)
	}
	fmt.Fprintf(w, "The server accepts this key: %q (%s), scopes: %s, expires %s.\n", str(k, "name"), str(k, "mode"), strings.Join(strList(k["scopes"]), ", "), str(k, "expires_at"))
	return 0
}
