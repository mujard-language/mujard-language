package main

// Helpers every build uses. (The worker/gateway addresses, their secrets and the signed health calls live in
// the operator-only sources: they belong to the operator build only.)
import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// envGet reads a MUJARD_* variable.
func envGet(name string) string {
	return os.Getenv(name)
}

func envOr(name, def string) string {
	if v := strings.TrimSpace(envGet(name)); v != "" {
		return v
	}
	return def
}

func firstNonEmpty(a ...string) string {
	for _, s := range a {
		if strings.TrimSpace(s) != "" {
			return strings.TrimSpace(s)
		}
	}
	return ""
}

// parseArgs parses a flag.FlagSet but lets flags and positional arguments mix: the standard
// flag package stops at the first positional, yet people type "mujard build nano blink.ino --out x".
// It returns the positional arguments. "--" ends the flags; a flag's value is the next argument
// unless the flag is a boolean or already carries "=".
func parseArgs(fs *flag.FlagSet, args []string) ([]string, error) {
	var flags, rest []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		if a == "--" {
			rest = append(rest, args[i+1:]...)
			break
		}
		if !strings.HasPrefix(a, "-") || a == "-" {
			rest = append(rest, a)
			continue
		}
		name := strings.TrimLeft(a, "-")
		flags = append(flags, a)
		if strings.Contains(name, "=") {
			continue // --x=v carries its value
		}
		if f := fs.Lookup(name); f != nil {
			if _, isBool := f.Value.(interface{ IsBoolFlag() bool }); !isBool && i+1 < len(args) {
				i++
				flags = append(flags, args[i])
			}
		}
	}
	if err := fs.Parse(flags); err != nil {
		return nil, err
	}
	return rest, nil
}

func plainGet(rawURL string, timeout time.Duration) (int, map[string]any, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return 0, nil, err
	}
	return doJSON(req, timeout)
}

func doJSON(req *http.Request, timeout time.Duration) (int, map[string]any, error) {
	resp, err := (&http.Client{Timeout: timeout}).Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	var m map[string]any
	_ = json.Unmarshal(raw, &m)
	return resp.StatusCode, m, nil
}

// short turns a network error into one plain sentence (no addresses of other machines, no stack).
func short(err error) string {
	if err == nil {
		return ""
	}
	s := err.Error()
	switch {
	case strings.Contains(s, "refused"):
		return "nothing is listening there"
	case strings.Contains(s, "timeout") || strings.Contains(s, "deadline"):
		return "no answer (timed out)"
	case strings.Contains(s, "no such host"):
		return "the address does not exist"
	}
	if len(s) > 100 {
		s = s[:100]
	}
	return s
}

func num(m map[string]any, key string) int {
	if f, ok := m[key].(float64); ok {
		return int(f)
	}
	return 0
}

func str(m map[string]any, key string) string {
	if s, ok := m[key].(string); ok {
		return s
	}
	return ""
}

// level of one check line.
type level int

const (
	ok level = iota
	warn
	fail
	skip
)

func (l level) tag() string {
	return [...]string{" OK ", "WARN", "FAIL", "skip"}[l]
}

func line(w io.Writer, l level, what, detail string) {
	if detail != "" {
		fmt.Fprintf(w, "  [%s] %-28s %s\n", l.tag(), what, detail)
		return
	}
	fmt.Fprintf(w, "  [%s] %s\n", l.tag(), what)
}
