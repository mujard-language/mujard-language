package lang

import (
	"regexp"
	"strings"
	"testing"
)

// mujard_config.h switches off exactly the runtime parts a program doesn't call, and nothing when hand-written C++
// could call any of it.
func TestRuntimeConfig(t *testing.T) {
	cases := []struct {
		name    string
		src     string
		off, on []string // defines that must / must not be in mujard_config.h
	}{
		{"blink uses none of it", "board uno\nled = pin LED as output\nevery 500 ms:\n    led.toggle()\n",
			[]string{"MUJARD_NO_APP", "MUJARD_NO_SERIAL_IN", "MUJARD_NO_BLINK", "MUJARD_NO_PATTERN", "MUJARD_NO_MELODY", "MUJARD_NO_TONE"}, nil},
		{"tone stays when a tone plays", "board uno\nbuzzer = pin 8 as output\nat start:\n    buzzer.tone(440 Hz)\n",
			[]string{"MUJARD_NO_MELODY"}, []string{"MUJARD_NO_TONE"}},
		{"serial input stays for serial.line", "board uno\non serial.line as cmd:\n    print cmd\n",
			[]string{"MUJARD_NO_APP"}, []string{"MUJARD_NO_SERIAL_IN"}},
		{"serial input stays for read_line in a task", "board uno\ntask t:\n    var l = serial.read_line()\n    print l\n",
			[]string{"MUJARD_NO_APP"}, []string{"MUJARD_NO_SERIAL_IN"}},
		{"the app protocol (and its serial input) stay for show", "board uno\nvar n = 0\nevery 1 s:\n    n += 1\n    show \"Count\", n\n",
			nil, []string{"MUJARD_NO_APP", "MUJARD_NO_SERIAL_IN"}},
		{"notify is an app built-in", "board uno\nat start:\n    notify \"hi\"\n",
			nil, []string{"MUJARD_NO_APP", "MUJARD_NO_SERIAL_IN"}},
		{"blink stays when blink is used", "board uno\nled = pin LED as output\nat start:\n    led.blink(2 Hz)\n",
			[]string{"MUJARD_NO_PATTERN", "MUJARD_NO_MELODY"}, []string{"MUJARD_NO_BLINK"}},
		{"pattern stays when a pattern is used", "board uno\nled = pin LED as output\nat start:\n    led.pattern(\"SOS\")\n",
			[]string{"MUJARD_NO_BLINK", "MUJARD_NO_MELODY"}, []string{"MUJARD_NO_PATTERN"}},
		{"melody stays when a melody is used", "board uno\nbuzzer = pin 8 as output\nat start:\n    buzzer.melody(\"C4 E4 G4\")\n",
			[]string{"MUJARD_NO_BLINK", "MUJARD_NO_PATTERN"}, []string{"MUJARD_NO_MELODY", "MUJARD_NO_TONE"}},
	}
	for _, c := range cases {
		sk, diags := Compile(c.src, "sketch.muj")
		if sk == nil {
			t.Fatalf("%s: %v", c.name, diags)
		}
		cfg := sk.Files[runtimeConfigFile]
		for _, d := range c.off {
			if !strings.Contains(cfg, "#define "+d+"\n") {
				t.Errorf("%s: %s should be switched off:\n%s", c.name, d, cfg)
			}
		}
		for _, d := range c.on {
			if strings.Contains(cfg, d+"\n") {
				t.Errorf("%s: %s is used and must stay on:\n%s", c.name, d, cfg)
			}
		}
	}

	// Hand-written C++ may call any of the runtime: no part of it is switched off.
	sk, diags := Compile("board uno\nled = pin LED as output\ncpp:\n    // anything\nevery 500 ms:\n    led.toggle()\n", "sketch.muj")
	if sk == nil {
		t.Fatalf("cpp block: %v", diags)
	}
	if cfg, has := sk.Files[runtimeConfigFile]; !has || strings.Contains(cfg, "#define") {
		t.Errorf("a program with a cpp: block gets a config that keeps the whole runtime (mujard_rt.h always includes it):\n%s", cfg)
	}
}

// The build server refuses sketches with these (the platform compiler's checkCpp, run over every generated file too):
// the runtime once looked for mujard_config.h with __has_include, which would have made the server refuse every
// Mujard build. Every sample's generated files must pass the same rules.
var serverRefuses = regexp.MustCompile(`(?im)__has_include|__has_embed|^\s*#\s*embed\b|incbin|\.\s*include\b|^\s*#\s*(?:include|include_next|import)\s*[^<"\s]`)
var serverIncludePath = regexp.MustCompile(`(?m)^\s*#\s*(?:include|include_next|import)\s*[<"]([^>"]+)[>"]`)

func TestGeneratedFilesPassTheServerChecks(t *testing.T) {
	samples := loadSamples(t, "ok")
	dir := ""
	for _, s := range samples {
		dir = s.dir
	}
	for _, s := range samples {
		sk, diags := CompileIn(s.src, s.name, dir)
		if sk == nil {
			if FirstError(diags) == nil {
				t.Errorf("%s: no sketch", s.name)
			}
			continue
		}
		if _, has := sk.Files[runtimeConfigFile]; !has {
			t.Errorf("%s: no %s, which mujard_rt.h includes", s.name, runtimeConfigFile)
		}
		for name, text := range sk.Files {
			if !(strings.HasSuffix(name, ".h") || strings.HasSuffix(name, ".cpp") || strings.HasSuffix(name, ".ino")) {
				continue
			}
			if m := serverRefuses.FindString(text); m != "" {
				t.Errorf("%s: %s has %q, which the build server refuses", s.name, name, m)
			}
			for _, inc := range serverIncludePath.FindAllStringSubmatch(text, -1) {
				if p := inc[1]; strings.Contains(p, "..") || strings.Contains(p, ":") || strings.HasPrefix(p, "/") || strings.HasPrefix(p, `\`) {
					t.Errorf("%s: %s includes %q, which the build server refuses", s.name, name, p)
				}
			}
		}
	}
}

// The runtime itself is never scanned: its own definitions would otherwise count as "used".
func TestRuntimeConfigIgnoresTheRuntime(t *testing.T) {
	cfg := runtimeConfig(map[string]string{
		"mujard_rt.h": "void app_begin(); bool serial_line(); x.blink(1); x.pattern(1); x.melody(1);", "mujard_rt.cpp": "app_on();",
		"mujard_program.cpp": "int main() {}", "sketch.ino": "#include \"mujard_rt.h\"", "mujard_libs.json": "app_begin(",
	})
	for _, d := range []string{"MUJARD_NO_APP", "MUJARD_NO_SERIAL_IN", "MUJARD_NO_BLINK", "MUJARD_NO_PATTERN", "MUJARD_NO_MELODY", "MUJARD_NO_TONE"} {
		if !strings.Contains(cfg, "#define "+d+"\n") {
			t.Errorf("%s should be off when only the runtime mentions it:\n%s", d, cfg)
		}
	}
}
