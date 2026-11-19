package lang

import (
	"bytes"
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	goscanner "go/scanner"
	gotoken "go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// The language version: one number every copy of Mujard reports (the CLI, the in-editor checker, editor-info.json for the
// phone, the server's compiler), so a program written for a newer language than an app or server knows gets a plain "update"
// message instead of a confusing error. version.json holds it with a fingerprint of the language's sources; when they change,
// TestLanguageVersionIsCurrent fails until it is refreshed (MUJARD_GOLDEN=1, which update-mujard.ps1 runs), and that refresh
// bumps the number - so nobody has to remember to.
//
//go:embed version.json
var versionJSON []byte

type versionFile struct {
	Version     int    `json:"version"`
	Fingerprint string `json:"fingerprint"`
}

// LanguageVersion is this build's language version (0 if version.json is unreadable, which the tests refuse).
func LanguageVersion() int {
	var v versionFile
	if json.Unmarshal(versionJSON, &v) != nil {
		return 0
	}
	return v.Version
}

// languageFingerprint hashes what makes the language what it is, in dir (tools/mujard/lang): the Go sources without their
// comments and formatting (a reworded comment is not a new language), the runtime the generated programs run on, the library
// descriptions and the data files (boards, capabilities). Tests, testdata and version.json itself are left out.
func languageFingerprint(dir string) (string, error) {
	var files []string
	for _, pattern := range []string{"*.go", "*.json", "runtime/*", "descs/*.mlib"} {
		m, err := filepath.Glob(filepath.Join(dir, pattern))
		if err != nil {
			return "", err
		}
		files = append(files, m...)
	}
	sort.Strings(files)
	h := sha256.New()
	fset := gotoken.NewFileSet()
	for _, f := range files {
		name := filepath.ToSlash(strings.TrimPrefix(f, dir))
		if strings.HasSuffix(f, "_test.go") || filepath.Base(f) == "version.json" {
			continue
		}
		if st, err := os.Stat(f); err != nil || st.IsDir() {
			continue
		}
		var content []byte
		if filepath.Base(f) == "capabilities.json" {
			raw, err := os.ReadFile(f)
			if err != nil {
				return "", err
			}
			content, err = capStructure(raw)
			if err != nil {
				return "", err
			}
		} else if strings.HasSuffix(f, ".go") {
			src, err := os.ReadFile(f)
			if err != nil {
				return "", err
			}
			content = goTokens(fset, src)
		} else {
			raw, err := os.ReadFile(f)
			if err != nil {
				return "", err
			}
			content = bytes.ReplaceAll(raw, []byte("\r\n"), []byte("\n"))
		}
		h.Write([]byte(name + "\x00"))
		h.Write(content)
		h.Write([]byte{0})
	}
	return hex.EncodeToString(h.Sum(nil))[:16], nil
}

// capStructure canonicalizes capabilities.json to only the part the language reads: which capabilities each board
// has proven. W618 checks presence (evidence.Build[board][cap] != ""), never the date - so a matrix re-dating a
// proven capability must not bump the language version, but a capability being added or lost still does.
func capStructure(raw []byte) ([]byte, error) {
	var ev Evidence
	if err := json.Unmarshal(raw, &ev); err != nil {
		return nil, err
	}
	var keys []string
	for board, caps := range ev.Build {
		for cp := range caps {
			keys = append(keys, "build\x00"+board+"\x00"+cp)
		}
	}
	for board, caps := range ev.Hardware {
		for cp := range caps {
			keys = append(keys, "hardware\x00"+board+"\x00"+cp)
		}
	}
	sort.Strings(keys)
	return []byte(strings.Join(keys, "\n")), nil
}

// goTokens is a Go source as its tokens only: comments, spacing, line breaks and the semicolons Go inserts at line ends are
// left out, so a reworded comment or a reformatted function is the same language and a changed word or symbol is not.
func goTokens(fset *gotoken.FileSet, src []byte) []byte {
	var s goscanner.Scanner
	s.Init(fset.AddFile("", fset.Base(), len(src)), src, nil, 0) // mode 0: no comments
	var b bytes.Buffer
	for {
		_, tok, lit := s.Scan()
		if tok == gotoken.EOF {
			return b.Bytes()
		}
		if tok == gotoken.SEMICOLON && lit == "\n" {
			continue
		}
		b.WriteString(tok.String())
		b.WriteByte(' ')
		b.WriteString(lit)
		b.WriteByte('\n')
	}
}
