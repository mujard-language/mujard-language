package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// The warm-core cache keeps compiled core objects + core.a between jobs, so a rebuild
// pays only for the sketch. It lives in Mujard-owned storage (MUJARD_BUILD_CACHE or the
// user cache dir), never inside the job folder - a job's outDir stays its own.

// coreCacheRoot is the directory holding per-key core archives; "" means caching is off
// because no cache location could be determined.
func coreCacheRoot() string {
	if v := os.Getenv("MUJARD_BUILD_CACHE"); v != "" {
		return filepath.Join(v, "core")
	}
	d, err := os.UserCacheDir()
	if err != nil {
		return ""
	}
	return filepath.Join(d, "mujard", "core")
}

// coreCacheDir is the cache slot for this plan's core build: everything that can change
// the objects feeds the key - the manifest and package files, the flag sets, the tools
// and the core source list (name, size, mtime).
func (p *Plan) coreCacheDir(flagSets ...[]string) string {
	root := coreCacheRoot()
	if root == "" {
		return ""
	}
	h := sha256.New()
	// The object naming scheme: v2 keeps the source extension (wiring_pulse.c.o / wiring_pulse.S.o). Archives
	// cached under v1 lost one of the two, so they must never be reused.
	h.Write([]byte("objects:v2\x00"))
	for _, f := range []string{p.Board.file, p.Pkg.file} {
		if b, err := os.ReadFile(f); err == nil {
			h.Write(b)
		}
	}
	for _, flags := range flagSets {
		for _, f := range flags {
			// job-local paths (the temp sketch dir, the out dir) differ every run and
			// must not poison the key; package paths (core, variant, libs) stay in.
			if strings.Contains(f, p.Job.SketchDir) || strings.Contains(f, p.Job.OutDir) {
				continue
			}
			h.Write([]byte(f))
			h.Write([]byte{0})
		}
	}
	for _, src := range p.CoreSrcs {
		if st, err := os.Stat(src); err == nil {
			fmt.Fprintf(h, "%s:%d:%d;", src, st.Size(), st.ModTime().Unix())
		}
	}
	return filepath.Join(root, hex.EncodeToString(h.Sum(nil))[:24])
}

// coreStaging is where a cold build writes the cacheable objects: a sibling of the final
// slot, renamed into place only after the link succeeds. The name is deterministic so a
// timed-out build resumes - finished objects are skipped and .part files recompiled -
// while the final slot can still never hold a half-written cache. A parallel build
// publishing the slot mid-run may orphan this staging; leftover dirs are pruned on use.
func (p *Plan) coreStaging(final string) string {
	return final + ".stage"
}

// publishCore moves a finished staging dir into its final slot. If another job won the
// race, the staging copy is dropped - the firmware is already linked either way.
func publishCore(p *Plan) {
	if p.CoreFinal == "" || p.CoreStaging == "" {
		return
	}
	if _, err := os.Stat(filepath.Join(p.CoreFinal, "core.a")); err == nil {
		os.RemoveAll(p.CoreStaging)
		return
	}
	if err := os.MkdirAll(filepath.Dir(p.CoreFinal), 0o755); err != nil {
		return
	}
	if err := os.Rename(p.CoreStaging, p.CoreFinal); err != nil {
		os.RemoveAll(p.CoreStaging)
	}
}
