package main

// Where the developer key and the API address live between runs. Saved as JSON in
// ~/.mujard/config.json (private to the user); every flag and the MUJARD_* environment
// variables still win over the file, so CI never needs it.
import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type cliConfig struct {
	API string `json:"api"` // the server's API folder, e.g. https://mujard.com/mujard-api
	Key string `json:"key"` // mjd_live_... or mjd_test_...
}

// mujardHome is ~/.mujard - the folder for config.json and the commands/ plugin folder.
// MUJARD_HOME moves it (tests, portable installs).
func mujardHome() string {
	if v := strings.TrimSpace(envGet("MUJARD_HOME")); v != "" {
		return v
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".mujard")
}

func configPath() string {
	if v := strings.TrimSpace(envGet("MUJARD_CONFIG")); v != "" {
		return v
	}
	if h := mujardHome(); h != "" {
		return filepath.Join(h, "config.json")
	}
	return ""
}

func commandsDir() string {
	if v := strings.TrimSpace(envGet("MUJARD_COMMANDS")); v != "" {
		return v
	}
	if h := mujardHome(); h != "" {
		return filepath.Join(h, "commands")
	}
	return ""
}

func dirExists(p string) bool {
	st, err := os.Stat(p)
	return err == nil && st.IsDir()
}

func loadConfig() cliConfig {
	var c cliConfig
	p := configPath()
	if p == "" {
		return c
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return c
	}
	_ = json.Unmarshal(b, &c)
	return c
}

// saveConfig writes the file readable only by its owner.
func saveConfig(c cliConfig) error {
	p := configPath()
	if p == "" {
		return fmt.Errorf("no home folder to save into; set MUJARD_CONFIG")
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return err
	}
	b, _ := json.MarshalIndent(c, "", "  ")
	return os.WriteFile(p, append(b, '\n'), 0o600)
}

func maskKey(k string) string {
	if len(k) <= 12 {
		return "(saved, hidden)"
	}
	return k[:8] + strings.Repeat("*", 4) + k[len(k)-4:]
}

func runLogin(args []string, w io.Writer) int {
	fs := flag.NewFlagSet("login", flag.ContinueOnError)
	fs.SetOutput(w)
	api := fs.String("api", envOr("MUJARD_API", ""), "the server's API folder, e.g. https://mujard.com/mujard-api")
	key := fs.String("key", envOr("MUJARD_KEY", ""), "your developer key (mjd_live_... or mjd_test_...) - create one on the developer page")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	c := loadConfig()
	if *api != "" {
		c.API = strings.TrimRight(*api, "/")
	}
	if *key != "" {
		c.Key = *key
	}
	if c.API == "" {
		c.API = "https://mujard.com/mujard-api"
	}
	if c.Key == "" {
		fmt.Fprintln(w, "Give the key: mujard login --key mjd_live_... [--api https://host/mujard-api]")
		fmt.Fprintln(w, "Keys are created on the site's developer page; a mjd_test_ key costs nothing and never builds for real.")
		return 2
	}
	// Prove the key works before saving it: a typo saved is a failure discovered later.
	cl := &apiClient{base: apiV1(c.API), key: c.Key}
	status, doc, err := cl.call("GET", "/me", nil, nil)
	if err != nil {
		fmt.Fprintf(w, "Cannot reach the API at %s: %s\n", cl.base, short(err))
		return 1
	}
	if status != 200 {
		fmt.Fprintf(w, "The server refused the key: %s\n", apiErrText(doc))
		return 1
	}
	if err := saveConfig(c); err != nil {
		fmt.Fprintf(w, "The key works, but it could not be saved: %s\n", err)
		return 1
	}
	name := "?"
	if k, _ := doc["key"].(map[string]any); k != nil {
		name = str(k, "name")
	}
	fmt.Fprintf(w, "Signed in - key %q (%s) saved to %s\n", name, maskKey(c.Key), configPath())
	return 0
}

func runLogout(args []string, w io.Writer) int {
	p := configPath()
	if p == "" {
		fmt.Fprintln(w, "No config file.")
		return 0
	}
	if err := os.Remove(p); err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(w, "Could not remove %s: %s\n", p, err)
		return 1
	}
	fmt.Fprintln(w, "Signed out - the saved key is gone from this machine.")
	return 0
}

// runConfigSub handles "config path", "config get api|key" and "config set api <url>". The key is only
// ever shown masked, and it is set by "mujard login" (which proves it works first), never from here.
func runConfigSub(args []string, w io.Writer) int {
	c := loadConfig()
	switch args[0] {
	case "path":
		if len(args) != 1 || configPath() == "" {
			fmt.Fprintln(w, "Usage: mujard config path   (the file where the key and address are saved)")
			return 2
		}
		fmt.Fprintln(w, configPath())
		return 0
	case "get":
		if len(args) != 2 || (args[1] != "api" && args[1] != "key") {
			fmt.Fprintln(w, "Usage: mujard config get api | key   (the key is shown masked)")
			return 2
		}
		if args[1] == "api" {
			fmt.Fprintln(w, firstNonEmpty(envGet("MUJARD_API"), c.API, "https://mujard.com/mujard-api"))
			return 0
		}
		if k := firstNonEmpty(envGet("MUJARD_KEY"), c.Key); k != "" {
			fmt.Fprintln(w, maskKey(k))
		} else {
			fmt.Fprintln(w, "(not set)")
		}
		return 0
	default: // set
		if len(args) == 3 && args[1] == "api" {
			u, err := url.Parse(args[2])
			if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Hostname() == "" {
				fmt.Fprintf(w, "%q is not an address: use something like https://your-host/mujard-api\n", args[2])
				return 2
			}
			c.API = strings.TrimRight(args[2], "/")
			if err := saveConfig(c); err != nil {
				fmt.Fprintf(w, "Could not save: %s\n", err)
				return 1
			}
			fmt.Fprintf(w, "Server address saved: %s\n", c.API)
			return 0
		}
		if len(args) >= 2 && args[1] == "key" {
			fmt.Fprintln(w, "The key is set with: mujard login --key mjd_live_...   (it checks the key with the server before saving it)")
			return 2
		}
		fmt.Fprintln(w, "Usage: mujard config set api <https://host/mujard-api>")
		return 2
	}
}

func runConfig(args []string, w io.Writer) int {
	if len(args) > 0 && (args[0] == "path" || args[0] == "get" || args[0] == "set") {
		return runConfigSub(args, w)
	}
	c := loadConfig()
	if c.API == "" && c.Key == "" {
		fmt.Fprintf(w, "Nothing saved. Defaults: --api %s, key from MUJARD_KEY or \"mujard login\".\n", envOr("MUJARD_API", "https://mujard.com/mujard-api"))
		return 0
	}
	fmt.Fprintf(w, "file  %s\napi   %s\nkey   %s\n", configPath(), c.API, maskKey(c.Key))
	return 0
}
