package main

import (
	"archive/tar"
	"archive/zip"
	"compress/bzip2"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// The toolchain manager installs pinned, checksummed packages into Mujard-owned
// directories so builds stop borrowing an IDE's data folder. `install` downloads
// from the checked-in index; `import` registers files already on disk (the files
// are identical content — the receipts record provenance either way).

type TCSystem struct {
	URL    string `json:"url"`
	SHA256 string `json:"sha256"`
	Size   int64  `json:"size"`
}

type TCPackage struct {
	ID      string              `json:"id"`
	Version string              `json:"version"`
	Name    string              `json:"name"`
	License string              `json:"license"`
	Source  string              `json:"source"`
	Systems map[string]TCSystem `json:"systems"`
}

type TCIndex struct {
	Schema   string      `json:"schema"`
	Packages []TCPackage `json:"packages"`
}

type TCReceipt struct {
	Schema      string `json:"schema"`
	ID          string `json:"id"`
	Version     string `json:"version"`
	Name        string `json:"name"`
	License     string `json:"license"`
	Source      string `json:"source"`
	URL         string `json:"url,omitempty"`
	SHA256      string `json:"sha256,omitempty"`
	Origin      string `json:"origin,omitempty"`
	InstalledAt string `json:"installedAt"`
}

// toolchainRoot is where installed packages live: MUJARD_TOOLCHAINS or the user
// cache dir — never inside a job's folders and never an IDE directory.
func toolchainRoot() string {
	if v := os.Getenv("MUJARD_TOOLCHAINS"); v != "" {
		return v
	}
	d, err := os.UserCacheDir()
	if err != nil {
		return ""
	}
	return filepath.Join(d, "mujard", "toolchains")
}

// toolchainIndexPath locates the package index: flag, env, or the toolchain/
// subtree found by walking up from the working directory.
func toolchainIndexPath(flagPath string) string {
	if flagPath != "" {
		return flagPath
	}
	if v := os.Getenv("MUJARD_TOOLCHAIN_INDEX"); v != "" {
		return v
	}
	cwd, err := os.Getwd()
	if err != nil {
		return filepath.Join("toolchain", "index.json")
	}
	for {
		p := filepath.Join(cwd, "toolchain", "index.json")
		if st, e := os.Stat(p); e == nil && !st.IsDir() {
			return p
		}
		up := filepath.Dir(cwd)
		if up == cwd {
			return filepath.Join("toolchain", "index.json")
		}
		cwd = up
	}
}

func loadIndex(path string) (*TCIndex, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read toolchain index %s: %w", path, err)
	}
	var ix TCIndex
	if err := json.Unmarshal(data, &ix); err != nil {
		return nil, fmt.Errorf("%s: %w", filepath.Base(path), err)
	}
	if ix.Schema != "mujard-toolchain-index.v1" {
		return nil, fmt.Errorf("%s: bad schema %q", filepath.Base(path), ix.Schema)
	}
	return &ix, nil
}

func (ix *TCIndex) find(id string) *TCPackage {
	for i := range ix.Packages {
		if ix.Packages[i].ID == id {
			return &ix.Packages[i]
		}
	}
	return nil
}

func installDir(id, version string) string {
	return filepath.Join(toolchainRoot(), id+"-"+version)
}

func installedReceipt(id, version string) *TCReceipt {
	data, err := os.ReadFile(filepath.Join(installDir(id, version), "receipt.json"))
	if err != nil {
		return nil
	}
	var r TCReceipt
	if json.Unmarshal(data, &r) != nil {
		return nil
	}
	return &r
}

// runToolchain is the `mujard-build toolchain <cmd>` entry point.
func runToolchain(args []string) int {
	indexPath := ""
	var cmd, idArg, fromDir, verArg, archiveArg string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--archive":
			i++
			if i < len(args) {
				archiveArg = args[i]
			}
		case "--index":
			i++
			if i < len(args) {
				indexPath = args[i]
			}
		case "--from":
			i++
			if i < len(args) {
				fromDir = args[i]
			}
		case "--version":
			i++
			if i < len(args) {
				verArg = args[i]
			}
		default:
			if cmd == "" {
				cmd = args[i]
			} else if idArg == "" {
				idArg = args[i]
			}
		}
	}
	root := toolchainRoot()
	if root == "" {
		fmt.Fprintln(os.Stderr, "cannot determine toolchain dir (set MUJARD_TOOLCHAINS)")
		return 2
	}
	switch cmd {
	case "dir":
		fmt.Println(root)
		return 0
	case "list":
		ix, err := loadIndex(toolchainIndexPath(indexPath))
		if err != nil {
			fmt.Fprintln(os.Stderr, "mujard-build: "+err.Error())
			return 1
		}
		for _, p := range ix.Packages {
			state := "  "
			if r := installedReceipt(p.ID, p.Version); r != nil {
				state = "* "
				fmt.Printf("%s %-16s %-28s %s  (%s)\n", state, p.ID, p.Name, p.Version, r.Source)
			} else {
				fmt.Printf("%s %-16s %-28s %s\n", state, p.ID, p.Name, p.Version)
			}
		}
		return 0
	case "install":
		if idArg == "" {
			fmt.Fprintln(os.Stderr, "usage: mujard-build toolchain install <id> [--version v] [--archive file]")
			return 2
		}
		ix, err := loadIndex(toolchainIndexPath(indexPath))
		if err != nil {
			fmt.Fprintln(os.Stderr, "mujard-build: "+err.Error())
			return 1
		}
		p := ix.find(idArg)
		if p == nil {
			fmt.Fprintf(os.Stderr, "no package %q in index\n", idArg)
			return 1
		}
		return tcInstall(p, verArg, archiveArg)
	case "import":
		if idArg == "" || fromDir == "" {
			fmt.Fprintln(os.Stderr, "usage: mujard-build toolchain import <id> --from <dir> [--version v]")
			return 2
		}
		ix, ixErr := loadIndex(toolchainIndexPath(indexPath))
		var p *TCPackage
		if ixErr == nil {
			p = ix.find(idArg)
		}
		if p == nil {
			p = &TCPackage{ID: idArg, Name: idArg}
		}
		return tcImport(p, fromDir, verArg)
	default:
		fmt.Fprintln(os.Stderr, "usage: mujard-build toolchain dir|list|install <id>|import <id> --from <dir>")
		return 2
	}
}

// tcInstall downloads a package archive for this OS/arch, verifies the SHA-256
// and extracts it into the toolchain dir. With localArchive (a copy of the same file
// already on this machine, e.g. a download cache) nothing is downloaded, but the copy
// must match the same pinned hash.
func tcInstall(p *TCPackage, verArg, localArchive string) int {
	version := p.Version
	if verArg != "" {
		version = verArg
	}
	sys, ok := p.Systems[runtime.GOOS+"-"+runtime.GOARCH]
	if !ok || sys.URL == "" {
		fmt.Fprintf(os.Stderr, "package %s has no %s-%s download\n", p.ID, runtime.GOOS, runtime.GOARCH)
		return 1
	}
	// Nothing is installed unchecked: an index entry without its pinned hash (or not over https) is refused, not trusted.
	if len(sys.SHA256) != 64 || !strings.HasPrefix(strings.ToLower(sys.URL), "https://") {
		fmt.Fprintf(os.Stderr, "package %s has no pinned sha256 or https URL in the index - refusing to install it\n", p.ID)
		return 1
	}
	dest := installDir(p.ID, version)
	if r := installedReceipt(p.ID, version); r != nil {
		fmt.Printf("%s %s already installed at %s\n", p.ID, version, dest)
		return 0
	}
	tmp, err := os.MkdirTemp("", "mujard-tc")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer os.RemoveAll(tmp)

	// The archive keeps the URL's file name either way: extract picks the format by extension.
	archive := filepath.Join(tmp, filepath.Base(sys.URL))
	source := "download"
	if localArchive != "" {
		fmt.Printf("installing %s %s from %s\n", p.ID, version, localArchive)
		if err := copyOr(localArchive, archive); err != nil {
			fmt.Fprintln(os.Stderr, "cannot read the archive: "+err.Error())
			return 1
		}
		source = "archive"
	} else {
		fmt.Printf("downloading %s %s\n  %s\n", p.ID, version, sys.URL)
		if err := download(sys.URL, archive); err != nil {
			fmt.Fprintln(os.Stderr, "download failed: "+err.Error())
			return 1
		}
	}
	sum := sha256File(archive)
	if !strings.EqualFold(sum, sys.SHA256) {
		fmt.Fprintf(os.Stderr, "sha256 mismatch: got %s want %s\n", sum, sys.SHA256)
		return 1
	}
	fmt.Printf("sha256 ok (%s), extracting to %s\n", sum[:12], dest)
	if err := extract(archive, dest); err != nil {
		os.RemoveAll(dest)
		fmt.Fprintln(os.Stderr, "extract failed: "+err.Error())
		return 1
	}
	r := &TCReceipt{Schema: "mujard-toolchain-receipt.v1", ID: p.ID, Version: version,
		Name: p.Name, License: p.License, Source: source, URL: sys.URL, SHA256: sum,
		InstalledAt: time.Now().UTC().Format(time.RFC3339)}
	writeReceipt(dest, r)
	fmt.Printf("installed %s %s\n", p.ID, version)
	return 0
}

// tcImport copies an on-disk toolchain tree into the toolchain dir — the way
// machines that already have the files (e.g. an IDE data folder) stop borrowing
// them. The receipt records the origin directory.
func tcImport(p *TCPackage, fromDir, verArg string) int {
	version := p.Version
	if verArg != "" {
		version = verArg
	}
	src, err := filepath.Abs(fromDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if st, err := os.Stat(src); err != nil || !st.IsDir() {
		fmt.Fprintf(os.Stderr, "%s is not a directory\n", src)
		return 1
	}
	dest := installDir(p.ID, version)
	if installedReceipt(p.ID, version) != nil {
		fmt.Printf("%s %s already installed at %s\n", p.ID, version, dest)
		return 0
	}
	os.RemoveAll(dest)
	if err := os.MkdirAll(dest, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Printf("importing %s\n  from %s\n", p.ID, src)
	if err := copyTree(src, dest); err != nil {
		os.RemoveAll(dest)
		fmt.Fprintln(os.Stderr, "import failed: "+err.Error())
		return 1
	}
	r := &TCReceipt{Schema: "mujard-toolchain-receipt.v1", ID: p.ID, Version: version,
		Name: p.Name, License: p.License, Source: "import", Origin: src,
		InstalledAt: time.Now().UTC().Format(time.RFC3339)}
	writeReceipt(dest, r)
	fmt.Printf("installed %s %s -> %s\n", p.ID, version, dest)
	return 0
}

func writeReceipt(dest string, r *TCReceipt) {
	data, _ := json.MarshalIndent(r, "", "  ")
	os.WriteFile(filepath.Join(dest, "receipt.json"), data, 0o644)
}

func download(url, dst string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}

func sha256File(path string) string {
	f, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer f.Close()
	h := sha256.New()
	io.Copy(h, f)
	return hex.EncodeToString(h.Sum(nil))
}

// extract unpacks .zip, .tar, .tar.gz and .tar.bz2 archives. Entries escaping
// the destination (.., absolute paths) are rejected.
func extract(archive, dest string) error {
	switch {
	case strings.HasSuffix(archive, ".zip"):
		return extractZip(archive, dest)
	case strings.HasSuffix(archive, ".tar.gz"), strings.HasSuffix(archive, ".tgz"):
		f, err := os.Open(archive)
		if err != nil {
			return err
		}
		defer f.Close()
		gz, err := gzip.NewReader(f)
		if err != nil {
			return err
		}
		return extractTar(tar.NewReader(gz), dest)
	case strings.HasSuffix(archive, ".tar.bz2"):
		f, err := os.Open(archive)
		if err != nil {
			return err
		}
		defer f.Close()
		return extractTar(tar.NewReader(bzip2.NewReader(f)), dest)
	default:
		return fmt.Errorf("unsupported archive %s", filepath.Base(archive))
	}
}

// safeDest rejects zip-slip style entries.
func safeDest(dest, name string) (string, error) {
	p := filepath.Join(dest, filepath.FromSlash(name))
	if !strings.HasPrefix(p, filepath.Clean(dest)+string(filepath.Separator)) {
		return "", fmt.Errorf("archive entry escapes destination: %s", name)
	}
	return p, nil
}

func extractZip(archive, dest string) error {
	zr, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer zr.Close()
	for _, f := range zr.File {
		p, err := safeDest(dest, f.Name)
		if err != nil {
			return err
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(p, 0o755)
			continue
		}
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		out, err := os.OpenFile(p, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o755)
		if err == nil {
			_, err = io.Copy(out, rc)
			out.Close()
		}
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func extractTar(tr *tar.Reader, dest string) error {
	for {
		h, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		p, err := safeDest(dest, h.Name)
		if err != nil {
			return err
		}
		switch h.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(p, 0o755)
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
				return err
			}
			out, err := os.OpenFile(p, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(h.Mode&0o777))
			if err != nil {
				return err
			}
			if _, err := io.Copy(out, tr); err != nil {
				out.Close()
				return err
			}
			out.Close()
		}
	}
}

func copyTree(src, dest string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil || rel == "." {
			return err
		}
		p := filepath.Join(dest, rel)
		if info.IsDir() {
			return os.MkdirAll(p, 0o755)
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		out, err := os.OpenFile(p, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			return err
		}
		_, err = io.Copy(out, in)
		out.Close()
		return err
	})
}
