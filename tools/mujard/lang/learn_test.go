package lang

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// The learning examples (docs/lang/examples/learn): graded programs, Basics to Advanced, that the web app, Mujard Studio
// and the Android app offer under "Examples". examples.json lists them in order; every one must compile all the way to
// C++ (not just check), so the apps never offer an example that can't be built. The apps read a bundle generated from
// them: MUJARD_GOLDEN=1 go test ./lang -run Learn rewrites it after an example changes.
var (
	learnDir          = filepath.Join("..", "..", "..", "docs", "lang", "examples", "learn")
	learnWebBundle    = filepath.Join("..", "..", "..", "web", "src", "examples.ts")
	learnAndroidAsset = filepath.Join("..", "..", "..", "app", "src", "main", "assets", "examples.json")
)

type learnManifest struct {
	Levels []struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Summary string `json:"summary"`
	} `json:"levels"`
	Examples []struct {
		ID      string            `json:"id"`
		Level   string            `json:"level"`
		Title   string            `json:"title"`
		Summary string            `json:"summary"`
		Files   map[string]string `json:"files"` // the name in the project -> the path under learn/
	} `json:"examples"`
}

type learnExample struct {
	ID      string            `json:"id"`
	Title   string            `json:"title"`
	Summary string            `json:"summary"`
	Board   string            `json:"board"`
	Files   map[string]string `json:"files"`
}

type learnLevel struct {
	ID       string         `json:"id"`
	Title    string         `json:"title"`
	Summary  string         `json:"summary"`
	Examples []learnExample `json:"examples"`
}

func loadLearn(t *testing.T) learnManifest {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(learnDir, "examples.json"))
	if err != nil {
		t.Fatalf("read examples.json: %v", err)
	}
	var m learnManifest
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatalf("examples.json: %v", err)
	}
	return m
}

func readLearnFile(t *testing.T, rel string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(learnDir, filepath.FromSlash(rel)))
	if err != nil {
		t.Fatalf("%v", err)
	}
	return strings.ReplaceAll(string(b), "\r\n", "\n")
}

// learnBundle compiles every example and returns the levels with their examples, in the manifest's order.
func learnBundle(t *testing.T) []learnLevel {
	t.Helper()
	m := loadLearn(t)
	levels := map[string]int{}
	var out []learnLevel
	for _, l := range m.Levels {
		if _, dup := levels[l.ID]; dup || l.ID == "" || l.Title == "" {
			t.Errorf("level %q: empty or listed twice", l.ID)
		}
		levels[l.ID] = len(out)
		out = append(out, learnLevel{ID: l.ID, Title: l.Title, Summary: l.Summary, Examples: []learnExample{}})
	}
	ids := map[string]bool{}
	used := map[string]bool{}
	levelLine := regexp.MustCompile(`^# Level (\d) - (.+)$`)
	for _, e := range m.Examples {
		if ids[e.ID] || e.ID == "" || e.Title == "" || e.Summary == "" {
			t.Errorf("example %q: empty id, title or summary, or listed twice", e.ID)
		}
		ids[e.ID] = true
		li, ok := levels[e.Level]
		if !ok {
			t.Errorf("%s: unknown level %q", e.ID, e.Level)
			continue
		}
		main, ok := e.Files["sketch.muj"]
		if !ok {
			t.Errorf("%s: no sketch.muj (the apps' main file)", e.ID)
			continue
		}
		files := map[string]string{}
		for name, rel := range e.Files {
			if used[rel] {
				t.Errorf("%s: %s is used by two examples", e.ID, rel)
			}
			used[rel] = true
			files[name] = readLearnFile(t, rel)
		}
		src := files["sketch.muj"]
		// The header every example starts with: its title, then its level, as the apps list it.
		lines := strings.SplitN(src, "\n", 3)
		if len(lines) < 3 || lines[0] != "# "+e.Title {
			t.Errorf("%s: the first line should be %q", main, "# "+e.Title)
		} else if lm := levelLine.FindStringSubmatch(lines[1]); lm == nil || !strings.HasPrefix(out[li].Title, lm[1]+". ") || !strings.HasSuffix(out[li].Title, lm[2]) {
			t.Errorf("%s: the second line should name its level, like %q", main, "# Level 1 - Basics")
		}
		sk, diags := CompileIn(src, "sketch.muj", filepath.Dir(filepath.Join(learnDir, filepath.FromSlash(main))))
		for _, d := range diags {
			if !d.IsWarning() || d.Code != "W604" { // W604: a cpp: block, shown on purpose in the Arduino-library example
				t.Errorf("%s: %s", main, d)
			}
		}
		if sk == nil {
			continue
		}
		out[li].Examples = append(out[li].Examples, learnExample{ID: e.ID, Title: e.Title, Summary: e.Summary, Board: sk.Board.Name, Files: files})
	}
	// Every program in the folder is in the list (a file nobody lists would never reach the apps).
	_ = filepath.WalkDir(learnDir, func(p string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() && (strings.HasSuffix(p, ".muj") || strings.HasSuffix(p, ".mjd")) {
			rel, _ := filepath.Rel(learnDir, p)
			if !used[filepath.ToSlash(rel)] {
				t.Errorf("%s is not listed in examples.json", filepath.ToSlash(rel))
			}
		}
		return nil
	})
	for _, l := range out {
		if len(l.Examples) == 0 {
			t.Errorf("level %s has no examples", l.ID)
		}
	}
	return out
}

func learnBundles(t *testing.T, levels []learnLevel) (web, android []byte) {
	t.Helper()
	j, err := json.MarshalIndent(levels, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	web = []byte("// Generated by tools/mujard/lang/learn_test.go from docs/lang/examples/learn - do not edit by hand.\n" +
		"// After changing an example: MUJARD_GOLDEN=1 go test ./lang -run Learn (in tools/mujard).\n" +
		"export interface MujardExample {\n  id: string;\n  title: string;\n  summary: string;\n  board: string;\n  files: Record<string, string>;\n}\n" +
		"export interface ExampleLevel {\n  id: string;\n  title: string;\n  summary: string;\n  examples: MujardExample[];\n}\n" +
		"export const EXAMPLE_LEVELS: ExampleLevel[] = " + string(j) + ";\n")
	return web, append(j, '\n')
}

func TestLearnExamples(t *testing.T) {
	levels := learnBundle(t)
	n := 0
	for _, l := range levels {
		n += len(l.Examples)
	}
	if n < 40 {
		t.Errorf("only %d examples compiled", n)
	}
	if t.Failed() {
		return
	}
	web, android := learnBundles(t, levels)
	want := map[string][]byte{}
	for p, b := range map[string][]byte{learnWebBundle: web, learnAndroidAsset: android} {
		if _, err := os.Stat(filepath.Dir(p)); err != nil {
			continue // the web//app checkouts are only in the monorepo
		}
		want[p] = b
	}
	for p, b := range want {
		if os.Getenv("MUJARD_GOLDEN") == "1" {
			if err := os.WriteFile(p, b, 0o644); err != nil {
				t.Fatalf("write %s: %v", p, err)
			}
			continue
		}
		got, err := os.ReadFile(p)
		if err != nil {
			t.Fatalf("read %s: %v (run with MUJARD_GOLDEN=1 to generate)", p, err)
		}
		if string(got) != string(b) {
			t.Errorf("%s is stale; regenerate with MUJARD_GOLDEN=1 go test ./lang -run Learn", p)
		}
	}
}
