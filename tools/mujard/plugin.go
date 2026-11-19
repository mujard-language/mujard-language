package main

// Custom commands, git-style: an executable named mujard-<verb> becomes "mujard <verb>". Looked up in
// the user's commands folder (~/.mujard/commands), a "commands" folder next to the program, the
// repository's tools/mujard/commands.d, then every PATH folder - first match wins, built-ins always
// win over plugins. The plugin gets the MUJARD_* environment and the command line after the verb,
// and its exit code becomes mujard's. A plugin is trusted code: it runs as you.
import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

type plugin struct {
	name string // the verb - "mujard-weather.ps1" is the "weather" command
	path string
}

// pluginDirs are the folders searched in order; PATH comes last in findPlugin.
func pluginDirs() []string {
	dirs := []string{}
	if d := commandsDir(); d != "" {
		dirs = append(dirs, d)
	}
	if e := exeDir(); e != "" {
		dirs = append(dirs, filepath.Join(e, "commands"))
	}
	if root := mujardRoot(); root != "" {
		dirs = append(dirs, filepath.Join(root, "tools", "mujard", "commands.d"))
	}
	return dirs
}

// pluginNames are the file names a verb may have, most native first.
func pluginNames(verb string) []string {
	if runtime.GOOS == "windows" {
		return []string{"mujard-" + verb + ".exe", "mujard-" + verb + ".cmd", "mujard-" + verb + ".bat",
			"mujard-" + verb + ".ps1", "mujard-" + verb + ".php", "mujard-" + verb + ".py", "mujard-" + verb}
	}
	return []string{"mujard-" + verb, "mujard-" + verb + ".sh", "mujard-" + verb + ".php", "mujard-" + verb + ".py"}
}

// runnable: on unix a plain "mujard-x" needs the executable bit; scripts going through an
// interpreter (.sh, .php, .py) do not. Windows has no bit - an existing file is runnable.
func runnable(p string, st os.FileInfo) bool {
	if runtime.GOOS == "windows" || st.Mode()&0o111 != 0 {
		return true
	}
	switch strings.ToLower(filepath.Ext(p)) {
	case ".sh", ".php", ".py":
		return true
	}
	return false
}

// findPlugin resolves an unknown verb to a program, or "" when nothing matches.
// Names that look like paths or flags are never looked up.
func findPlugin(verb string) string {
	if verb == "" || strings.HasPrefix(verb, "-") || strings.ContainsAny(verb, `/\`) {
		return ""
	}
	for _, dir := range pluginDirs() {
		for _, name := range pluginNames(verb) {
			p := filepath.Join(dir, name)
			if st, err := os.Stat(p); err == nil && !st.IsDir() && runnable(p, st) {
				return p
			}
		}
	}
	for _, dir := range filepath.SplitList(os.Getenv("PATH")) {
		if dir == "" {
			continue
		}
		for _, name := range pluginNames(verb) {
			p := filepath.Join(dir, name)
			if st, err := os.Stat(p); err == nil && !st.IsDir() && runnable(p, st) {
				return p
			}
		}
	}
	return ""
}

// pluginFromFile: "mujard-deploy.ps1" is the "deploy" command; other names are not plugins.
func pluginFromFile(dir, file string) (plugin, bool) {
	if !strings.HasPrefix(file, "mujard-") {
		return plugin{}, false
	}
	base := strings.TrimPrefix(file, "mujard-")
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)
	if name == "" {
		return plugin{}, false
	}
	switch strings.ToLower(ext) {
	case ".exe", ".cmd", ".bat", ".ps1", ".php", ".py", ".sh", "":
		return plugin{name: name, path: filepath.Join(dir, file)}, true
	}
	return plugin{}, false
}

// listPlugins finds every plugin in the dedicated folders (the PATH is searched at run time,
// not listed - it holds other people's programs).
func listPlugins() []plugin {
	seen := map[string]bool{}
	out := []plugin{}
	for _, dir := range pluginDirs() {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			if p, ok := pluginFromFile(dir, e.Name()); ok && !seen[p.name] && findBuiltin(p.name) == nil && !operatorOnly[p.name] {
				seen[p.name] = true
				out = append(out, p)
			}
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].name < out[j].name })
	return out
}

// runPlugin executes a found plugin: interpreter by extension, our MUJARD_* environment passed on,
// stdio and exit code straight through.
func runPlugin(path string, args []string) int {
	var cmd *exec.Cmd
	switch strings.ToLower(filepath.Ext(path)) {
	case ".ps1":
		cmd = exec.Command("powershell", append([]string{"-NoProfile", "-ExecutionPolicy", "Bypass", "-File", path}, args...)...)
	case ".bat", ".cmd":
		cmd = exec.Command("cmd", append([]string{"/c", path}, args...)...)
	case ".php":
		cmd = exec.Command(phpExe(), append([]string{path}, args...)...)
	case ".py":
		cmd = exec.Command("python", append([]string{path}, args...)...)
	case ".sh":
		cmd = exec.Command("sh", append([]string{path}, args...)...)
	default:
		cmd = exec.Command(path, args...)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = append(os.Environ(), pluginEnv()...)
	if err := cmd.Run(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return ee.ExitCode()
		}
		fmt.Fprintf(os.Stderr, "mujard: could not run %s: %s\n", path, err)
		return 1
	}
	return 0
}

// pluginEnv gives a plugin the same context the built-ins have.
func pluginEnv() []string {
	out := []string{"MUJARD_VERSION=" + version}
	if root := mujardRoot(); root != "" && envGet("MUJARD_ROOT") == "" {
		out = append(out, "MUJARD_ROOT="+root)
	}
	if h := mujardHome(); h != "" && envGet("MUJARD_HOME") == "" {
		out = append(out, "MUJARD_HOME="+h)
	}
	return out
}
