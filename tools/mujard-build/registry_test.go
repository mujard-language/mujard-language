package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// writeManifests makes a scratch registry: boards/ and packages/ under t.TempDir().
func writeManifests(t *testing.T, boards, pkgs map[string]string) (string, string) {
	t.Helper()
	root := t.TempDir()
	bdir := filepath.Join(root, "boards")
	pdir := filepath.Join(root, "packages")
	for dir, files := range map[string]map[string]string{bdir: boards, pdir: pkgs} {
		if len(files) == 0 {
			continue
		}
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
		for name, body := range files {
			if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
				t.Fatal(err)
			}
		}
	}
	return bdir, pdir
}

const testBoard = `{
  "schema": "mujard-board.v1", "id": "microchip:avr:m328p", "short": "m328p",
  "name": "ATmega328P board", "vendor": "microchip", "family": "avr",
  "aliases": ["uno", "arduino:avr:uno"], "package": "pkg1",
  "core": "cores/arduino", "variant": "variants/std",
  "vars": {"mcu": "atmega328p", "fCpu": "16000000L", "boardDefine": "ARDUINO_AVR_UNO"},
  "output": "hex"
}`

func TestResolveNames(t *testing.T) {
	bdir, pdir := writeManifests(t,
		map[string]string{"m328p.json": testBoard},
		map[string]string{"p.json": `{"schema":"mujard-package.v1","id":"pkg1","family":"avr","root":"/tmp","tools":{}}`})
	reg, err := LoadRegistry(bdir, pdir)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"m328p", "M328P", "microchip:avr:m328p", "uno", "ARDUINO:AVR:UNO"} {
		m, err := reg.Resolve(name)
		if err != nil {
			t.Fatalf("resolve %q: %v", name, err)
		}
		if m.ID != "microchip:avr:m328p" {
			t.Fatalf("resolve %q gave %s", name, m.ID)
		}
	}
	if _, err := reg.Resolve("esp9000"); err == nil {
		t.Fatal("unknown board resolved")
	}
}

func TestDuplicateNamesRejected(t *testing.T) {
	dupShort := strings.Replace(testBoard, `"short": "m328p"`, `"short": "m328p"`, 1)
	other := `{"schema":"mujard-board.v1","id":"x:y:z","short":"m328p","name":"x","family":"avr"}`
	bdir, pdir := writeManifests(t, map[string]string{"a.json": dupShort, "b.json": other}, nil)
	if _, err := LoadRegistry(bdir, pdir); err == nil || !strings.Contains(err.Error(), "duplicate") {
		t.Fatalf("duplicate short name not rejected: %v", err)
	}

	aliasClash := `{"schema":"mujard-board.v1","id":"x:y:z","short":"other","name":"x","family":"avr","aliases":["m328p"]}`
	bdir2, pdir2 := writeManifests(t, map[string]string{"a.json": testBoard, "b.json": aliasClash}, nil)
	if _, err := LoadRegistry(bdir2, pdir2); err == nil || !strings.Contains(err.Error(), "collides") {
		t.Fatalf("alias colliding with a short name not rejected: %v", err)
	}

	dupAlias := `{"schema":"mujard-board.v1","id":"x:y:z","short":"other","name":"x","family":"avr","aliases":["uno"]}`
	bdir3, pdir3 := writeManifests(t, map[string]string{"a.json": testBoard, "b.json": dupAlias}, nil)
	if _, err := LoadRegistry(bdir3, pdir3); err == nil {
		t.Fatal("duplicate alias not rejected")
	}
}

func TestBadSchemaRejected(t *testing.T) {
	bdir, pdir := writeManifests(t, map[string]string{"x.json": `{"schema":"nope","id":"a","short":"b","family":"c"}`}, nil)
	if _, err := LoadRegistry(bdir, pdir); err == nil {
		t.Fatal("bad schema accepted")
	}
}

func TestExpandEnv(t *testing.T) {
	t.Setenv("MJ_TEST_ROOT", "/x/y")
	if got := expandEnv("${MJ_TEST_ROOT}/bin/tool"); got != "/x/y/bin/tool" {
		t.Fatalf("expandEnv gave %q", got)
	}
	if got := expandEnv("${MJ_UNSET_VAR}/bin"); got != "/bin" {
		t.Fatalf("unset var gave %q", got)
	}
}
