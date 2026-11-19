package lang

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// The compatibility promise (spec 14, plan Phase 12): a program that compiles under language edition Mujard 2026
// keeps compiling under every later 1.x compiler. testdata/compat/2026 is a frozen copy of the samples
// (docs/lang/examples/ok) and the learning examples (docs/lang/examples/learn) as they were on 2026-10-29. Never edit
// or delete these files to make this test pass: a failure here is a breaking change, which needs a new edition and
// an automatic `mujard fix`. Programs may be added (new 2026 programs that must keep working); warnings are allowed,
// since a later compiler may warn more.
var compatDir = filepath.Join("testdata", "compat", "2026")

var compatMain = regexp.MustCompile(`(?m)^board[ \t]+\S`)

func TestEdition2026StillCompiles(t *testing.T) {
	n := 0
	err := filepath.WalkDir(compatDir, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !(strings.HasSuffix(p, ".muj") || strings.HasSuffix(p, ".mjd")) {
			return err
		}
		raw, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		src := strings.ReplaceAll(string(raw), "\r\n", "\n")
		if !compatMain.MatchString(src) {
			return nil // a module: compiled through the program that imports it
		}
		n++
		sk, diags := CompileIn(src, filepath.Base(p), filepath.Dir(p))
		for _, d := range diags {
			if !d.IsWarning() {
				t.Errorf("%s: %s", filepath.ToSlash(p), d)
			}
		}
		if sk == nil && !t.Failed() {
			t.Errorf("%s: no sketch generated", filepath.ToSlash(p))
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if n < 80 {
		t.Fatalf("only %d programs in %s: the frozen set is missing files", n, compatDir)
	}
}
