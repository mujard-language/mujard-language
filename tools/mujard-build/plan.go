package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Step is one external command: an executable plus an argument list. Never a shell line.
// Parallel steps produce independent outputs and may run concurrently; sequential
// steps (archive, link, images) run in order. When Out is set the step writes
// <Out>.part which the runner renames to Out on success — a killed compile can
// never leave a half-written object to be mistaken for a finished one.
type Step struct {
	Name     string   `json:"name"`
	Exe      string   `json:"exe"`
	Args     []string `json:"args"`
	Parallel bool     `json:"parallel,omitempty"`
	Out      string   `json:"out,omitempty"`
}

// Plan is the full recipe instance for one job: resolved board+package, the source sets
// and every step in order.
type Plan struct {
	Job        *Job
	Board      *BoardManifest
	Pkg        *Package
	Options    map[string]string // job options after manifest defaults+validation
	Vars       map[string]string // board vars with option overrides applied
	CoreSrcs   []string
	SketchSrcs []string          // post-preprocessing compile units (.c/.cpp/.S paths)
	PreprocIno map[string]string // generated .cpp path -> original .ino
	// Defined is the macros the recipe's flags define. When set, the library scan skips includes in #ifdef blocks
	// that are off for this build (rp2); nil keeps the plain scan of every include.
	Defined map[string]bool
	Steps   []Step

	// Warm-core cache: CoreFinal is the published slot; CoreStaging is where a cold build
	// writes before publishCore renames it into place.
	CoreFinal   string
	CoreStaging string
}

// unsupported marks a job that cannot run under this backend; callers fall back.
type unsupported struct{ why string }

func (e *unsupported) Error() string { return e.why }

// buildPlan resolves board, options, package and recipe. Recipe failures return
// *unsupported so runBuild reports them distinctly from compile errors.
func buildPlan(reg *Registry, job *Job) (*Plan, error) {
	b, err := reg.Resolve(job.Board)
	if err != nil {
		// An unknown board is a coverage gap, not a build error: callers fall back.
		return nil, &unsupported{err.Error()}
	}
	opts, err := resolveOptions(b, job.Options)
	if err != nil {
		return nil, err
	}
	p := &Plan{Job: job, Board: b, Options: opts, Vars: effectiveVars(b, opts), PreprocIno: map[string]string{}}
	pkg := reg.Package(b)
	if pkg == nil {
		return p, &unsupported{fmt.Sprintf("board %s (%s) has no toolchain package yet", b.Short, b.ID)}
	}
	p.Pkg = pkg
	switch b.Family {
	case "avr":
		if err := p.avrSteps(); err != nil {
			return p, err
		}
	case "esp32":
		if err := p.esp32Steps(); err != nil {
			return p, err
		}
	case "rp2":
		if err := p.rp2Steps(); err != nil {
			return p, err
		}
	case "esp8266":
		if err := p.esp8266Steps(); err != nil {
			return p, err
		}
	case "stm32":
		if err := p.stm32Steps(); err != nil {
			return p, err
		}
	case "samd":
		if err := p.samdSteps(); err != nil {
			return p, err
		}
	case "renesas":
		if err := p.renesasSteps(); err != nil {
			return p, err
		}
	case "mbed":
		if err := p.mbedSteps(); err != nil {
			return p, err
		}
	default:
		return p, &unsupported{fmt.Sprintf("family %q has no direct-build recipe yet", b.Family)}
	}
	return p, nil
}

// resolveOptions applies manifest defaults and validates each job option. Unknown keys
// or values are job errors — options are typed choices, never raw menu suffixes.
func resolveOptions(b *BoardManifest, given map[string]string) (map[string]string, error) {
	out := map[string]string{}
	for k, def := range b.Options {
		out[k] = def.Default
	}
	for k, v := range given {
		def, ok := b.Options[k]
		if !ok {
			return nil, fmt.Errorf("board %s has no option %q", b.Short, k)
		}
		found := false
		for _, c := range def.Choices {
			if c == v {
				found = true
			}
		}
		if !found {
			return nil, fmt.Errorf("option %q for %s must be one of %s (got %q)", k, b.Short, strings.Join(def.Choices, ", "), v)
		}
		out[k] = v
	}
	return out, nil
}

// effectiveVars clones the board's vars and applies each resolved option: an
// explicit per-choice var map wins, otherwise the option value lands on the var
// of the same name.
func effectiveVars(b *BoardManifest, opts map[string]string) map[string]string {
	out := map[string]string{}
	for k, v := range b.Vars {
		out[k] = v
	}
	for k, v := range opts {
		if def, ok := b.Options[k]; ok {
			if m, ok := def.Vars[v]; ok {
				for vn, vv := range m {
					out[vn] = vv
				}
				continue
			}
		}
		out[k] = v
	}
	return out
}

var srcExts = map[string]bool{".c": true, ".cpp": true, ".s": true}

// sources lists the compile sources in dir (top level only). .ino files go to inoFiles
// for preprocessing.
func sources(dir string) (srcs, inoFiles []string, err error) {
	ents, err := os.ReadDir(dir)
	if err != nil {
		return nil, nil, err
	}
	for _, e := range ents {
		if e.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(e.Name()))
		p := filepath.Join(dir, e.Name())
		switch {
		case ext == ".ino":
			inoFiles = append(inoFiles, p)
		case srcExts[ext]:
			srcs = append(srcs, p)
		}
	}
	sort.Strings(srcs)
	sort.Strings(inoFiles)
	return srcs, inoFiles, nil
}

// preprocessIno writes each .ino as C++ (Arduino.h first) into outDir/preproc and returns
// the generated file list. Prototype generation for hand-written .ino sketches is a
// later phase; Mujard-generated sketches carry their real code in .cpp files already.
func (p *Plan) preprocessIno(files []string) ([]string, error) {
	if len(files) == 0 {
		return nil, nil
	}
	dir := filepath.Join(p.Job.OutDir, "preproc")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	var out []string
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		name := filepath.Base(f) + ".cpp"
		dst := filepath.Join(dir, name)
		body := "#include <Arduino.h>\n#line 1 \"" + filepath.ToSlash(f) + "\"\n" + string(data)
		if err := os.WriteFile(dst, []byte(body), 0o644); err != nil {
			return nil, err
		}
		p.PreprocIno[dst] = f
		out = append(out, dst)
	}
	return out, nil
}

// Explain renders the resolution --explain prints.
func (p *Plan) Explain() string {
	var s strings.Builder
	fmt.Fprintf(&s, "board:    %s (%s) — %s\n", p.Board.Short, p.Board.ID, p.Board.Name)
	fmt.Fprintf(&s, "family:   %s   chip: %s\n", p.Board.Family, p.Board.Chip)
	if p.Pkg != nil {
		fmt.Fprintf(&s, "package:  %s %s\n", p.Pkg.ID, p.Pkg.Version)
		fmt.Fprintf(&s, "root:     %s\n", p.Pkg.Root)
	}
	if len(p.Options) > 0 {
		var kv []string
		for k, v := range p.Options {
			kv = append(kv, k+"="+v)
		}
		sort.Strings(kv)
		fmt.Fprintf(&s, "options:  %s\n", strings.Join(kv, " "))
	}
	fmt.Fprintf(&s, "vars:     %v\n", p.Board.Vars)
	fmt.Fprintf(&s, "core:     %d sources   sketch: %d sources   steps: %d\n", len(p.CoreSrcs), len(p.SketchSrcs), len(p.Steps))
	fmt.Fprintf(&s, "output:   %s\n", p.Board.Output)
	return s.String()
}
