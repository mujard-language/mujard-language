package main

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// Registry is every board manifest and package descriptor, indexed by canonical id,
// short name and alias. One name always resolves to exactly one manifest.
type Registry struct {
	boards   map[string]*BoardManifest // canonical id
	short    map[string]*BoardManifest // short name -> board
	alias    map[string]*BoardManifest // alias -> board
	packages map[string]*Package       // package id
}

// LoadRegistry reads all *.json manifests in boardsDir and packagesDir. A duplicate id,
// short name or alias anywhere in the set is a hard error — the index is the authority
// that keeps names unique.
func LoadRegistry(boardsDir, pkgsDir string) (*Registry, error) {
	r := &Registry{
		boards:   map[string]*BoardManifest{},
		short:    map[string]*BoardManifest{},
		alias:    map[string]*BoardManifest{},
		packages: map[string]*Package{},
	}
	files, err := filepath.Glob(filepath.Join(boardsDir, "*.json"))
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		m, err := loadBoardFile(f)
		if err != nil {
			return nil, err
		}
		id := strings.ToLower(m.ID)
		if _, dup := r.boards[id]; dup {
			return nil, fmt.Errorf("duplicate board id %q (%s)", m.ID, filepath.Base(f))
		}
		r.boards[id] = m
		short := strings.ToLower(m.Short)
		if _, dup := r.short[short]; dup {
			return nil, fmt.Errorf("duplicate board short name %q (%s)", m.Short, filepath.Base(f))
		}
		r.short[short] = m
		for _, a := range m.Aliases {
			a = strings.ToLower(a)
			if _, dup := r.alias[a]; dup {
				return nil, fmt.Errorf("duplicate board alias %q (%s)", a, filepath.Base(f))
			}
			if _, dup := r.short[a]; dup {
				return nil, fmt.Errorf("alias %q collides with a short name (%s)", a, filepath.Base(f))
			}
			r.alias[a] = m
		}
	}
	pfiles, err := filepath.Glob(filepath.Join(pkgsDir, "*.json"))
	if err != nil {
		return nil, err
	}
	for _, f := range pfiles {
		p, err := loadPackageFile(f)
		if err != nil {
			return nil, err
		}
		id := strings.ToLower(p.ID)
		if _, dup := r.packages[id]; dup {
			return nil, fmt.Errorf("duplicate package id %q (%s)", p.ID, filepath.Base(f))
		}
		r.packages[id] = p
	}
	return r, nil
}

// Resolve turns a user-typed board name — short, canonical id or alias — into the one
// manifest it means. Case-insensitive.
func (r *Registry) Resolve(name string) (*BoardManifest, error) {
	n := strings.ToLower(strings.TrimSpace(name))
	if m, ok := r.short[n]; ok {
		return m, nil
	}
	if m, ok := r.boards[n]; ok {
		return m, nil
	}
	if m, ok := r.alias[n]; ok {
		return m, nil
	}
	return nil, fmt.Errorf("unknown board %q (mujard-build --list shows every name)", name)
}

// Package returns the package descriptor a board needs, or nil when the board is
// declared but not yet buildable.
func (r *Registry) Package(m *BoardManifest) *Package {
	if m.Package == "" {
		return nil
	}
	return r.packages[strings.ToLower(m.Package)]
}

// List renders every resolvable name, canonical ids first.
func (r *Registry) List() []string {
	var out []string
	var ids []string
	for id := range r.boards {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		m := r.boards[id]
		line := fmt.Sprintf("%-14s %-28s %s", m.Short, m.ID, m.Name)
		if len(m.Aliases) > 0 {
			line += "  (aliases: " + strings.Join(m.Aliases, ", ") + ")"
		}
		if m.Package == "" {
			line += "  [not yet buildable]"
		}
		out = append(out, line)
	}
	return out
}
