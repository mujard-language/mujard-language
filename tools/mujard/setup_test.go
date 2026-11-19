package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

// memPath stands in for the registry: no test ever touches the real user PATH.
type memPath struct {
	value  string
	kind   uint32
	writes int
	getErr error
}

func (m *memPath) get() (string, uint32, error) { return m.value, m.kind, m.getErr }
func (m *memPath) set(v string, k uint32) error {
	m.value, m.kind = v, k
	m.writes++
	return nil
}

func noMujard(string) (string, error) { return "", errors.New("not found") }

const selfExe = `C:\Tools\Mujard\MujardCLI.exe`

func TestPathHelpers(t *testing.T) {
	t.Setenv("MUJARDTESTDIR", `C:\Tools\Mujard`)
	if !pathHas(`C:\a;c:\tools\mujard\;C:\b`, `C:\Tools\Mujard`) {
		t.Error("matching is case-blind and ignores a trailing backslash")
	}
	if !pathHas(`%MUJARDTESTDIR%;C:\b`, `C:\Tools\Mujard`) {
		t.Error("an entry written as %VAR% matches its expansion")
	}
	if !pathHas(`"C:\Tools\Mujard"`, `C:\Tools\Mujard`) {
		t.Error("a quoted entry matches")
	}
	if pathHas(`C:\Tools\Mujard2`, `C:\Tools\Mujard`) {
		t.Error("a longer folder name is a different folder")
	}
	if got := pathAdd(`C:\a;C:\b;`, `C:\x`); got != `C:\a;C:\b;C:\x` {
		t.Errorf("pathAdd = %q", got)
	}
	if got := pathAdd("", `C:\x`); got != `C:\x` {
		t.Errorf("pathAdd to empty = %q", got)
	}
	got, removed := pathRemove(`%SystemRoot%\x;C:\x\;C:\b;C:\X`, `C:\x`)
	if !removed || got != `%SystemRoot%\x;C:\b` {
		t.Errorf("pathRemove = %q, %v (every copy goes, %%VAR%% entries stay as written)", got, removed)
	}
	if expandPercent("%NO_SUCH_MUJARD_VAR%\\bin") != "%NO_SUCH_MUJARD_VAR%\\bin" {
		t.Error("an unknown %VAR% stays as written")
	}
}

func TestSetupPathAddsOnceAndUndoes(t *testing.T) {
	store := &memPath{value: `%USERPROFILE%\AppData\Local\Microsoft\WindowsApps;C:\b`, kind: 2}
	var out bytes.Buffer
	if code := setupPath(nil, &out, store, selfExe, noMujard); code != 0 || !strings.Contains(out.String(), "Added") {
		t.Fatalf("add: exit %d\n%s", code, out.String())
	}
	want := `%USERPROFILE%\AppData\Local\Microsoft\WindowsApps;C:\b;C:\Tools\Mujard`
	if store.value != want || store.kind != 2 {
		t.Fatalf("appended last, %%VAR%% kept raw, type kept: %q (kind %d)", store.value, store.kind)
	}
	out.Reset()
	if code := setupPath(nil, &out, store, selfExe, noMujard); code != 0 || store.writes != 1 || !strings.Contains(out.String(), "already") {
		t.Errorf("a second run changes nothing (exit %d, writes %d)\n%s", code, store.writes, out.String())
	}
	out.Reset()
	if code := setupPath([]string{"--undo"}, &out, store, selfExe, noMujard); code != 0 || store.value != `%USERPROFILE%\AppData\Local\Microsoft\WindowsApps;C:\b` {
		t.Errorf("undo restores the original value (exit %d): %q", code, store.value)
	}
	out.Reset()
	writes := store.writes
	if code := setupPath([]string{"--undo"}, &out, store, selfExe, noMujard); code != 0 || store.writes != writes {
		t.Errorf("undo when absent writes nothing (exit %d)", code)
	}
}

func TestSetupPathNeverShadowsAnotherMujard(t *testing.T) {
	store := &memPath{value: `C:\b`, kind: 2}
	other := func(string) (string, error) { return `C:\Other\MujardCLI.exe`, nil }
	var out bytes.Buffer
	if code := setupPath(nil, &out, store, selfExe, other); code != 1 || store.writes != 0 || !strings.Contains(out.String(), `C:\Other\MujardCLI.exe`) {
		t.Errorf("a different mujard on the PATH: refuse, name it, change nothing (exit %d, writes %d)\n%s", code, store.writes, out.String())
	}
	self := func(string) (string, error) { return selfExe, nil }
	out.Reset()
	if code := setupPath(nil, &out, store, selfExe, self); code != 0 || store.writes != 1 {
		t.Errorf("finding this same mujard is not a conflict (exit %d)", code)
	}
}

func TestSetupPathCheckAndUsage(t *testing.T) {
	store := &memPath{value: `C:\b`, kind: 2}
	var out bytes.Buffer
	if code := setupPath([]string{"--check"}, &out, store, selfExe, noMujard); code != 0 || store.writes != 0 || !strings.Contains(out.String(), "on user PATH:    false") {
		t.Errorf("--check only reports (exit %d)\n%s", code, out.String())
	}
	for _, args := range [][]string{{"--check", "--undo"}, {"extra"}, {"--nope"}} {
		out.Reset()
		if code := setupPath(args, &out, store, selfExe, noMujard); code != 2 || store.writes != 0 {
			t.Errorf("%q is a usage error that changes nothing (exit %d)", args, code)
		}
	}
	out.Reset()
	if code := runSetup(nil, &out); code != 2 {
		t.Errorf("bare \"mujard setup\" is a usage error, got %d", code)
	}
	out.Reset()
	if code := setupPath(nil, &out, store, `C:\Users\x\AppData\Local\Temp\go-build123\b001\exe\MujardCLI.exe`, noMujard); code != 1 || store.writes != 0 {
		t.Errorf("a go-run temp binary is refused (exit %d)", code)
	}
	out.Reset()
	if code := setupPath(nil, &out, &memPath{getErr: errSetupUnsupported}, `/opt/mujard/mujard`, noMujard); code != 1 || !strings.Contains(out.String(), "export PATH") {
		t.Errorf("without a user PATH store: print the profile line instead (exit %d)\n%s", code, out.String())
	}
}
