package main

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// Minimal bundled-library resolver: scan the sketch's #includes, map each header to a
// library under the package's libraryRoot, and contribute that library's include dirs +
// sources to the build. Libraries can include other libraries, so resolution iterates
// to a fixpoint (bounded). Unknown headers are left for the compiler to diagnose - the
// resolver never invents names.
var includePat = regexp.MustCompile(`#include\s*[<"]([^>"]+)[>"]`)

// lib is one resolved library directory under the package's library root.
type lib struct {
	dir      string   // <libRoot>/<Name>
	src      string   // <dir>/src when present, else dir
	incs     []string // include dirs: src + immediate source subdirs (e.g. utility)
	srcs     []string // .c/.cpp/.S under src, recursive
	hdrFiles []string // .h/.hpp under src - headers include other libraries too
	hdr      map[string]bool
}

// resolveLibs walks the sketch sources' includes against the package's bundled
// libraries and returns include dirs plus the libraries' sources (by library name).
func (p *Plan) resolveLibs() (incs []string, libSrcs map[string][]string, err error) {
	all := map[string]*lib{}
	for _, root := range append([]string{p.Pkg.LibRoot}, p.Pkg.LibRoots...) {
		if root == "" {
			continue
		}
		if !filepath.IsAbs(root) {
			root = p.Pkg.subRoot(root)
		}
		found, err := indexLibs(root)
		if err != nil {
			return nil, nil, err
		}
		for name, l := range found { // a missing root (a pack not installed) adds nothing
			if all[name] == nil {
				all[name] = l
			}
		}
	}
	if len(all) == 0 {
		return nil, nil, nil
	}
	used := map[string]*lib{}
	scan := p.SketchSrcs
	scanned := map[string]bool{}
	for _, f := range scan {
		scanned[filepath.Clean(f)] = true
	}
	for pass := 0; pass < 8; pass++ {
		var next []string
		for _, h := range gatherIncludes(scan, p.Defined) {
			matched := false
			for name, l := range all {
				if used[name] != nil || !l.provides(h) {
					continue
				}
				used[name] = l
				matched = true
				next = append(next, l.srcs...)
				next = append(next, l.hdrFiles...)
			}
			if matched {
				continue
			}
			// A sketch-local header (mujard_spi.h, a user header) can name a bundled
			// library header - open it so its includes are scanned too.
			f := filepath.Join(p.Job.SketchDir, filepath.FromSlash(h))
			if !scanned[f] {
				if st, serr := os.Stat(f); serr == nil && !st.IsDir() {
					scanned[f] = true
					next = append(next, f)
				}
			}
		}
		if len(next) == 0 {
			break
		}
		scan = next // library sources and local headers may pull in more libraries
	}
	var names []string
	for n := range used {
		names = append(names, n)
	}
	sort.Strings(names)
	libSrcs = map[string][]string{}
	for _, n := range names {
		l := used[n]
		incs = append(incs, l.incs...)
		libSrcs[n] = l.srcs
	}
	return incs, libSrcs, nil
}

// provides reports whether #include <h> finds a file in this library the way the compiler would: h relative to one
// of the library's include dirs. A file name alone is not enough - "pico/time.h" is not any library's time.h, and
// the Pico core's TinyUSB library has a tusb.h the core must never pick up instead of the Pico SDK's.
func (l *lib) provides(h string) bool {
	if !l.hdr[filepath.Base(h)] {
		return false
	}
	for _, inc := range l.incs {
		if st, err := os.Stat(filepath.Join(inc, filepath.FromSlash(h))); err == nil && !st.IsDir() {
			return true
		}
	}
	return false
}

// indexLibs maps every bundled library's headers. Returns nil when the root is absent.
func indexLibs(root string) (map[string]*lib, error) {
	ents, err := os.ReadDir(root)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	out := map[string]*lib{}
	for _, e := range ents {
		if !e.IsDir() {
			continue
		}
		l := &lib{dir: filepath.Join(root, e.Name()), hdr: map[string]bool{}}
		l.src = l.dir
		if st, err := os.Stat(filepath.Join(l.dir, "src")); err == nil && st.IsDir() {
			l.src = filepath.Join(l.dir, "src")
		}
		l.incs = append(l.incs, l.src)
		flat := l.src == l.dir
		filepath.Walk(l.src, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() {
				// A flat (1.0-format) library is its top folder plus utility/, as the Arduino library spec says:
				// examples/ and extras/ hold sketches and tools that are not part of the library (NeoPixel's
				// examples include a BLE sketch that does not compile anywhere else).
				if flat && path != l.src && !(filepath.Dir(path) == l.src && strings.EqualFold(info.Name(), "utility")) {
					return filepath.SkipDir
				}
				// one level of source subdirs gets an include dir (utility/, ...)
				if filepath.Dir(path) == l.src {
					l.incs = append(l.incs, path)
				}
				return nil
			}
			ext := strings.ToLower(filepath.Ext(path))
			switch ext {
			case ".h", ".hpp":
				l.hdr[info.Name()] = true
				l.hdrFiles = append(l.hdrFiles, path)
			case ".c", ".cpp", ".s":
				l.srcs = append(l.srcs, path)
			}
			return nil
		})
		if len(l.hdr) > 0 {
			out[e.Name()] = l
		}
	}
	return out, nil
}

// gatherIncludes lists the header names every file in srcs includes. With defined (the macros the build's flags
// define) it skips the includes inside simple #ifdef / #ifndef / #if defined(X) blocks that are off for this build:
// Adafruit_NeoPixel.h includes <Adafruit_TinyUSB.h> only under #ifdef USE_TINYUSB, and pulling that library in
// regardless put its tusb.h ahead of the Pico SDK's. Any other #if counts as on. A nil defined scans every include.
func gatherIncludes(srcs []string, defined map[string]bool) []string {
	seen := map[string]bool{}
	var out []string
	for _, f := range srcs {
		data, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		var found []string
		if defined == nil {
			for _, m := range includePat.FindAllSubmatch(data, -1) {
				found = append(found, string(m[1]))
			}
		} else {
			found = activeIncludes(string(data), defined)
		}
		for _, h := range found {
			if !seen[h] {
				seen[h] = true
				out = append(out, h)
			}
		}
	}
	return out
}

var (
	ifdefPat   = regexp.MustCompile(`^#\s*(ifdef|ifndef)\s+([A-Za-z_]\w*)`)
	ifDefPat   = regexp.MustCompile(`^#\s*if\s+(!\s*)?defined\s*\(?\s*([A-Za-z_]\w*)\s*\)?\s*(?://.*|/\*.*)?$`)
	ifPat      = regexp.MustCompile(`^#\s*if\b`)
	elsePat    = regexp.MustCompile(`^#\s*(else|elif)\b`)
	endifPat   = regexp.MustCompile(`^#\s*endif\b`)
	definePat  = regexp.MustCompile(`^#\s*define\s+([A-Za-z_]\w*)`)
	includeDir = regexp.MustCompile(`^#\s*include\s*[<"]([^>"]+)[>"]`)
)

// activeIncludes is the include list of one file with the switched-off conditional blocks left out. A block's state
// is on, off, or unknown (any #if or #elif this doesn't evaluate, which counts as on). After a known branch was
// taken the rest of the chain is off; when none was, #else is on and an #elif is unknown.
func activeIncludes(src string, defined map[string]bool) []string {
	type frame struct{ on, known, taken, parentOn bool }
	var stack []frame
	on := true
	local := map[string]bool{}
	isDef := func(name string) bool { return defined[name] || local[name] }
	var out []string
	for _, raw := range strings.Split(src, "\n") {
		line := strings.TrimSpace(raw)
		if !strings.HasPrefix(line, "#") {
			continue
		}
		switch {
		case ifdefPat.MatchString(line):
			m := ifdefPat.FindStringSubmatch(line)
			val := isDef(m[2]) == (m[1] == "ifdef")
			stack = append(stack, frame{on: val, known: true, taken: val, parentOn: on})
			on = on && val
		case ifDefPat.MatchString(line):
			m := ifDefPat.FindStringSubmatch(line)
			val := isDef(m[2]) == (m[1] == "")
			stack = append(stack, frame{on: val, known: true, taken: val, parentOn: on})
			on = on && val
		case ifPat.MatchString(line):
			stack = append(stack, frame{on: true, parentOn: on})
		case elsePat.MatchString(line):
			if n := len(stack); n > 0 {
				f := &stack[n-1]
				switch {
				case !f.known:
					f.on = true
				case f.taken:
					f.on = false
				case elsePat.FindStringSubmatch(line)[1] == "else":
					f.on, f.taken = true, true
				default: // an #elif this doesn't evaluate
					f.on, f.known = true, false
				}
				on = f.parentOn && f.on
			}
		case endifPat.MatchString(line):
			if n := len(stack); n > 0 {
				on = stack[n-1].parentOn
				stack = stack[:n-1]
			}
		case on && definePat.MatchString(line):
			local[definePat.FindStringSubmatch(line)[1]] = true
		case on && includeDir.MatchString(line):
			out = append(out, includeDir.FindStringSubmatch(line)[1])
		}
	}
	return out
}
