package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// OptionDef is one named board option: real choices, never FQBN menu suffixes.
type OptionDef struct {
	Choices []string                     `json:"choices"`
	Default string                       `json:"default"`
	Vars    map[string]map[string]string `json:"vars"`
}

// BoardManifest is schema mujard-board.v1 — one board's identity plus what the build needs.
type BoardManifest struct {
	Schema   string               `json:"schema"`
	ID       string               `json:"id"`
	Short    string               `json:"short"`
	Name     string               `json:"name"`
	Vendor   string               `json:"vendor"`
	Family   string               `json:"family"`
	Chip     string               `json:"chip"`
	Aliases  []string             `json:"aliases"`
	Package  string               `json:"package"`
	Core     string               `json:"core"`
	Variant  string               `json:"variant"`
	Options  map[string]OptionDef `json:"options"`
	Vars     map[string]string    `json:"vars"`
	Defines  []string             `json:"defines"`
	Output   string               `json:"output"`
	FlashMax int                  `json:"flashMax"`
	RamMax   int                  `json:"ramMax"`
	Upload   map[string]any       `json:"upload"`

	file string
}

// Package is schema mujard-package.v1 — one toolchain bundle: a core source tree plus
// the absolute tool paths the build steps run.
type Package struct {
	Schema  string `json:"schema"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Family  string `json:"family"`
	Root    string `json:"root"`
	LibRoot string `json:"libraryRoot"`
	// More folders of libraries, each its own pinned toolchain package (Servo is not part of the AVR core).
	// Searched after libraryRoot, which wins when both have a library of the same name.
	LibRoots []string          `json:"libraryRoots"`
	Tools    map[string]string `json:"tools"`
	SDKs     map[string]string `json:"sdks"` // per-chip prebuilt SDK trees (esp32: esp32s3-libs/...)

	file string
}

var envPat = regexp.MustCompile(`\$\{([A-Za-z_][A-Za-z0-9_]*)\}`)

// expandEnv replaces ${VAR} in manifest paths — keeps machine paths out of the schema.
func expandEnv(s string) string {
	return envPat.ReplaceAllStringFunc(s, func(m string) string {
		name := envPat.FindStringSubmatch(m)[1]
		if v := os.Getenv(name); v != "" {
			return v
		}
		if name == "MUJARD_TOOLCHAINS" {
			return toolchainRoot()
		}
		return ""
	})
}

func loadBoardFile(path string) (*BoardManifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var m BoardManifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("%s: %w", filepath.Base(path), err)
	}
	if m.Schema != "mujard-board.v1" {
		return nil, fmt.Errorf("%s: bad schema %q", filepath.Base(path), m.Schema)
	}
	if m.ID == "" || m.Short == "" || m.Family == "" {
		return nil, fmt.Errorf("%s: id, short and family are required", filepath.Base(path))
	}
	m.file = path
	return &m, nil
}

func loadPackageFile(path string) (*Package, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var p Package
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, fmt.Errorf("%s: %w", filepath.Base(path), err)
	}
	if p.Schema != "mujard-package.v1" {
		return nil, fmt.Errorf("%s: bad schema %q", filepath.Base(path), p.Schema)
	}
	if p.ID == "" {
		return nil, fmt.Errorf("%s: id is required", filepath.Base(path))
	}
	p.Root = expandEnv(p.Root)
	p.LibRoot = expandEnv(p.LibRoot)
	for i, r := range p.LibRoots {
		p.LibRoots[i] = expandEnv(r)
	}
	for k, v := range p.Tools {
		p.Tools[k] = expandEnv(v)
	}
	for k, v := range p.SDKs {
		p.SDKs[k] = expandEnv(v)
	}
	p.file = path
	return &p, nil
}

// tool resolves one package tool to an executable path; "" and false when absent.
func (p *Package) tool(name string) (string, bool) {
	raw, ok := p.Tools[name]
	if !ok || raw == "" {
		return "", false
	}
	for _, cand := range []string{raw, raw + ".exe"} {
		if st, err := os.Stat(cand); err == nil && !st.IsDir() {
			if abs, err := filepath.Abs(cand); err == nil {
				return abs, true
			}
			return cand, true
		}
	}
	return "", false
}

// subRoot joins a manifest-relative dir (cores/arduino) onto the package root.
func (p *Package) subRoot(rel string) string {
	rel = strings.TrimPrefix(filepath.ToSlash(rel), "/")
	return filepath.Join(p.Root, filepath.FromSlash(rel))
}
