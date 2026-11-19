package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// mujard pkg: the Mujard package registry (the platform's package-registry endpoint serves it).
//
//	mujard pkg search [text]                         packages by name or description
//	mujard pkg info <name>                           its versions
//	mujard pkg install <name>[@version] [--dir d]    put its files next to your program: then 'import <name>'
//	mujard pkg publish <folder> --version 1.0.0      share a library (the server checks it with the compiler first)
//
// install writes only Mujard sources and descriptions, only inside the folder, and never replaces a different file of
// the same name unless --force; mujard_packages.json records which version each package is at.

var (
	pkgName    = regexp.MustCompile(`^[a-z][a-z0-9_]{1,39}$`)
	pkgVersion = regexp.MustCompile(`^(0|[1-9]\d{0,3})\.(0|[1-9]\d{0,3})\.(0|[1-9]\d{0,3})$`)
	pkgFile    = regexp.MustCompile(`^[a-z][a-z0-9_]{0,39}\.(muj|mjd|mlib)$`)
)

const pkgLockFile = "mujard_packages.json"

func runPkg(args []string, w io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(w, "mujard pkg search [text] | info <name> | install <name>[@version] [--dir d] [--force] | publish <folder> --version x.y.z [--name n] [--board b] [--description d]")
		return 2
	}
	sub, rest := args[0], args[1:]
	var positional []string
	for len(rest) > 0 && !strings.HasPrefix(rest[0], "-") {
		positional, rest = append(positional, rest[0]), rest[1:]
	}
	fs := flag.NewFlagSet("pkg "+sub, flag.ContinueOnError)
	fs.SetOutput(w)
	c := addAPIFlags(fs)
	dir := fs.String("dir", ".", "install: the program's folder")
	force := fs.Bool("force", false, "install: replace files of the same name that differ")
	ver := fs.String("version", "", "publish: the version, MAJOR.MINOR.PATCH")
	name := fs.String("name", "", "publish: the package name (default: the folder's name)")
	board := fs.String("board", "uno", "publish: the board the server checks it for")
	desc := fs.String("description", "", "publish: one line about it")
	if err := fs.Parse(rest); err != nil {
		return 2
	}
	positional = append(positional, fs.Args()...)
	c.finish()
	if c.needKey(w) {
		return 1
	}
	switch sub {
	case "search":
		return pkgSearch(c, strings.Join(positional, " "), w)
	case "info":
		if len(positional) != 1 {
			fmt.Fprintln(w, "mujard pkg info <name>")
			return 2
		}
		return pkgInfo(c, positional[0], w)
	case "install":
		if len(positional) != 1 {
			fmt.Fprintln(w, "mujard pkg install <name>[@version] [--dir d] [--force]")
			return 2
		}
		return pkgInstall(c, positional[0], *dir, *force, w)
	case "publish":
		if len(positional) != 1 {
			fmt.Fprintln(w, "mujard pkg publish <folder> --version x.y.z")
			return 2
		}
		return pkgPublish(c, positional[0], *name, *ver, *board, *desc, w)
	}
	fmt.Fprintf(w, "Unknown: mujard pkg %s. Use search, info, install or publish.\n", sub)
	return 2
}

func pkgSearch(c *apiClient, text string, w io.Writer) int {
	status, doc, err := c.call("GET", "/packages?q="+url.QueryEscape(text), nil, nil)
	if err != nil {
		fmt.Fprintln(w, "Couldn't reach the server:", err)
		return 1
	}
	if status != 200 {
		return failAPI(w, status, doc)
	}
	list, _ := doc["data"].([]any)
	if len(list) == 0 {
		fmt.Fprintln(w, "No packages found.")
		return 0
	}
	for _, it := range list {
		p, _ := it.(map[string]any)
		fmt.Fprintf(w, "%-24s %-9s %s\n", str(p, "name"), str(p, "latest"), str(p, "description"))
	}
	return 0
}

func pkgInfo(c *apiClient, name string, w io.Writer) int {
	if !pkgName.MatchString(name) {
		fmt.Fprintln(w, "A package name is 2-40 lowercase letters, digits and _.")
		return 2
	}
	status, doc, err := c.call("GET", "/packages/"+name, nil, nil)
	if err != nil {
		fmt.Fprintln(w, "Couldn't reach the server:", err)
		return 1
	}
	if status != 200 {
		return failAPI(w, status, doc)
	}
	fmt.Fprintf(w, "%s - %s\n", str(doc, "name"), str(doc, "description"))
	versions, _ := doc["versions"].([]any)
	for _, it := range versions {
		v, _ := it.(map[string]any)
		fmt.Fprintf(w, "  %-9s checked for the %-8s published %s\n", str(v, "version"), str(v, "board"), str(v, "published_at"))
	}
	return 0
}

func pkgInstall(c *apiClient, spec, dir string, force bool, w io.Writer) int {
	name, version, _ := strings.Cut(spec, "@")
	if version == "" {
		version = "latest"
	}
	if !pkgName.MatchString(name) || (version != "latest" && !pkgVersion.MatchString(version)) {
		fmt.Fprintln(w, "Write it as name or name@1.2.3.")
		return 2
	}
	status, doc, err := c.call("GET", "/packages/"+name+"/"+version, nil, nil)
	if err != nil {
		fmt.Fprintln(w, "Couldn't reach the server:", err)
		return 1
	}
	if status != 200 {
		return failAPI(w, status, doc)
	}
	files, _ := doc["files"].(map[string]any)
	if len(files) == 0 {
		fmt.Fprintln(w, "The server sent no files for that version.")
		return 1
	}
	// Everything is checked before anything is written: one bad name or a conflict writes nothing.
	names := make([]string, 0, len(files))
	for f, v := range files {
		src, ok := v.(string)
		if !ok || !pkgFile.MatchString(f) || f == "sketch.muj" || f == "sketch.mjd" {
			fmt.Fprintf(w, "Refused: the package names a file %q that isn't a Mujard source or description.\n", f)
			return 1
		}
		if old, err := os.ReadFile(filepath.Join(dir, f)); err == nil && string(old) != src && !force {
			fmt.Fprintf(w, "%s already exists here and differs from %s's. Nothing was written (--force replaces it).\n", f, name)
			return 1
		}
		names = append(names, f)
	}
	sort.Strings(names)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		fmt.Fprintln(w, "Couldn't make the folder:", err)
		return 1
	}
	for _, f := range names {
		if err := os.WriteFile(filepath.Join(dir, f), []byte(files[f].(string)), 0o644); err != nil {
			fmt.Fprintln(w, "Couldn't write", f+":", err)
			return 1
		}
	}
	lock := map[string]string{}
	if b, err := os.ReadFile(filepath.Join(dir, pkgLockFile)); err == nil {
		_ = json.Unmarshal(b, &lock)
	}
	got := str(doc, "version")
	lock[name] = got
	b, _ := json.MarshalIndent(lock, "", "  ")
	if err := os.WriteFile(filepath.Join(dir, pkgLockFile), append(b, '\n'), 0o644); err != nil {
		fmt.Fprintln(w, "Couldn't record it in", pkgLockFile+":", err)
		return 1
	}
	fmt.Fprintf(w, "Installed %s %s (%s). In your program: import %s\n", name, got, strings.Join(names, ", "), name)
	return 0
}

func pkgPublish(c *apiClient, folder, name, version, board, desc string, w io.Writer) int {
	if name == "" {
		name = strings.ToLower(filepath.Base(filepath.Clean(folder)))
	}
	if !pkgName.MatchString(name) {
		fmt.Fprintf(w, "%q isn't a package name (2-40 lowercase letters, digits and _): pass --name.\n", name)
		return 2
	}
	if !pkgVersion.MatchString(version) {
		fmt.Fprintln(w, "Pass --version MAJOR.MINOR.PATCH, for example --version 1.0.0.")
		return 2
	}
	entries, err := os.ReadDir(folder)
	if err != nil {
		fmt.Fprintln(w, "Couldn't read the folder:", err)
		return 1
	}
	files := map[string]string{}
	for _, e := range entries {
		f := e.Name()
		if e.IsDir() || !pkgFile.MatchString(f) || f == "sketch.muj" || f == "sketch.mjd" {
			continue // a package is the library's sources and descriptions; the folder's program and anything else stay home
		}
		b, err := os.ReadFile(filepath.Join(folder, f))
		if err != nil {
			fmt.Fprintln(w, "Couldn't read", f+":", err)
			return 1
		}
		files[f] = string(b)
	}
	if files[name+".muj"] == "" && files[name+".mjd"] == "" {
		fmt.Fprintf(w, "The folder needs the package's module, %s.muj.\n", name)
		return 1
	}
	body := map[string]any{"name": name, "version": version, "board": board, "description": desc, "files": files}
	status, doc, err := c.call("POST", "/packages", body, map[string]string{"Idempotency-Key": idemKey()})
	if err != nil {
		fmt.Fprintln(w, "Couldn't reach the server:", err)
		return 1
	}
	if status != 201 {
		code := failAPI(w, status, doc)
		if e, _ := doc["error"].(map[string]any); e != nil {
			if d, _ := e["details"].(map[string]any); d != nil {
				if errs, _ := d["errors"].([]any); len(errs) > 0 {
					for _, x := range errs {
						fmt.Fprintf(w, "  %v\n", x)
					}
				}
			}
		}
		return code
	}
	fmt.Fprintf(w, "Published %s %s (%d files), checked for the %s. Others install it with: mujard pkg install %s@%s\n",
		name, version, len(files), board, name, version)
	return 0
}
