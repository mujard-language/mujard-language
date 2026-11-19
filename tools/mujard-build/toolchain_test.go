package main

import (
	"archive/zip"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// An index entry without a pinned hash, or one served over plain http, is refused before anything is downloaded.
func TestInstallRefusesUnpinnedPackages(t *testing.T) {
	root := t.TempDir()
	t.Setenv("MUJARD_TOOLCHAINS", root)
	key := runtime.GOOS + "-" + runtime.GOARCH
	for name, sys := range map[string]TCSystem{
		"no hash":    {URL: "https://example.invalid/x.zip"},
		"short hash": {URL: "https://example.invalid/x.zip", SHA256: "abc"},
		"plain http": {URL: "http://example.invalid/x.zip", SHA256: strings.Repeat("a", 64)},
	} {
		p := &TCPackage{ID: "x", Version: "1", Systems: map[string]TCSystem{key: sys}}
		if code := tcInstall(p, "", ""); code != 1 {
			t.Errorf("%s: install must refuse (exit %d)", name, code)
		}
		if _, err := os.Stat(filepath.Join(root, "x-1")); err == nil {
			t.Errorf("%s: nothing may be installed", name)
		}
	}
}

// --archive installs from a copy already on this machine, but only when it matches the pinned hash.
func TestInstallFromALocalArchiveChecksThePin(t *testing.T) {
	root := t.TempDir()
	t.Setenv("MUJARD_TOOLCHAINS", root)
	zipPath := filepath.Join(t.TempDir(), "Lib-1.0.zip")
	f, err := os.Create(zipPath)
	if err != nil {
		t.Fatal(err)
	}
	zw := zip.NewWriter(f)
	w, _ := zw.Create("Lib/src/Lib.h")
	w.Write([]byte("// lib\n"))
	zw.Close()
	f.Close()
	sum := sha256File(zipPath)
	key := runtime.GOOS + "-" + runtime.GOARCH
	pkg := func(hash string) *TCPackage {
		return &TCPackage{ID: "lib", Version: "1.0", Systems: map[string]TCSystem{key: {URL: "https://example.invalid/Lib-1.0.zip", SHA256: hash}}}
	}
	if code := tcInstall(pkg(strings.Repeat("0", 64)), "", zipPath); code != 1 {
		t.Fatalf("a wrong hash must be refused (exit %d)", code)
	}
	if _, err := os.Stat(filepath.Join(root, "lib-1.0")); err == nil {
		t.Fatal("nothing may be installed after a hash mismatch")
	}
	if code := tcInstall(pkg(sum), "", zipPath); code != 0 {
		t.Fatalf("matching archive refused (exit %d)", code)
	}
	if _, err := os.Stat(filepath.Join(root, "lib-1.0", "Lib", "src", "Lib.h")); err != nil {
		t.Fatalf("not extracted: %v", err)
	}
	r := installedReceipt("lib", "1.0")
	if r == nil || r.Source != "archive" || r.SHA256 != sum {
		t.Fatalf("receipt %+v", r)
	}
}

// Every package in the shipped index is pinned and served over https.
func TestShippedIndexIsPinned(t *testing.T) {
	ix, err := loadIndex(filepath.Join("..", "..", "toolchain", "index.json"))
	if err != nil {
		t.Skipf("no index here: %v", err)
	}
	for _, p := range ix.Packages {
		for sysName, s := range p.Systems {
			if len(s.SHA256) != 64 || !strings.HasPrefix(s.URL, "https://") {
				t.Errorf("%s (%s): sha256 %q url %q", p.ID, sysName, s.SHA256, s.URL)
			}
		}
	}
}
