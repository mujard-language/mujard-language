package main

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestSplitLine(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"build ./Blink --board nano", []string{"build", "./Blink", "--board", "nano"}},
		{`run "C:\My Sketches\Blink" -b nano -p COM3`, []string{"run", `C:\My Sketches\Blink`, "-b", "nano", "-p", "COM3"}},
		{`ai 'why is my LED off'`, []string{"ai", "why is my LED off"}},
		{`ai "don't panic"`, []string{"ai", "don't panic"}},
		{`flash C:\fw\blink.hex --port COM5`, []string{"flash", `C:\fw\blink.hex`, "--port", "COM5"}},
		{"  boards\t list  ", []string{"boards", "list"}},
		{`build "" --board uno`, []string{"build", "", "--board", "uno"}},
		{`--out="a b.hex"`, []string{"--out=a b.hex"}},
		{"", nil},
	}
	for _, c := range cases {
		got, err := splitLine(c.in)
		if err != nil || !reflect.DeepEqual(got, c.want) {
			t.Errorf("splitLine(%q) = %q, %v; want %q", c.in, got, err, c.want)
		}
	}
	for _, bad := range []string{`run "C:\Blink`, `ai 'oops`} {
		if _, err := splitLine(bad); err == nil {
			t.Errorf("splitLine(%q) must refuse an unclosed quote", bad)
		}
	}
}

// shellRun drives the loop with scripted input and a recorder in place of the child process.
func shellRun(t *testing.T, input string, interactive bool, codes map[string]int, sig chan os.Signal) (int, [][]string, string) {
	t.Helper()
	var ran [][]string
	var out bytes.Buffer
	run := func(args []string) int {
		ran = append(ran, args)
		return codes[args[0]]
	}
	code := runShell(strings.NewReader(input), &out, interactive, run, sig)
	return code, ran, out.String()
}

func TestShellRunsEachLineThroughTheDispatcher(t *testing.T) {
	code, ran, out := shellRun(t, "boards list\r\nmujard board info nano\nupload ./Blink -b nano -p COM3 -y\nexit\nversion\n", false, nil, nil)
	want := [][]string{{"boards", "list"}, {"board", "info", "nano"}, {"upload", "./Blink", "-b", "nano", "-p", "COM3", "-y"}}
	if code != 0 || !reflect.DeepEqual(ran, want) {
		t.Fatalf("exit %d, ran %q; want 0 and %q (nothing after exit)", code, ran, want)
	}
	if strings.Contains(out, shellPrompt) {
		t.Error("no prompt or banner when the input is a script")
	}
}

func TestShellIgnoresAByteOrderMark(t *testing.T) {
	_, ran, out := shellRun(t, "\uFEFF# comment\nversion\n", false, nil, nil)
	if !reflect.DeepEqual(ran, [][]string{{"version"}}) {
		t.Errorf("a BOM before a comment on line 1 must still be a comment: ran %q, out %q", ran, out)
	}
	_, ran, _ = shellRun(t, "\uFEFFversion\nboards\n", false, nil, nil)
	if !reflect.DeepEqual(ran, [][]string{{"version"}, {"boards"}}) {
		t.Errorf("a BOM before a command on line 1 must not change the command: %q", ran)
	}
}

func TestShellRecoversFromErrors(t *testing.T) {
	codes := map[string]int{"flash": 1, "bogus": 2}
	code, ran, out := shellRun(t, "flash x.hex\nbogus\nrun \"C:\\Blink\nshell\nhelp\nquit\n", false, codes, nil)
	if code != 0 {
		t.Fatalf("quit leaves with 0, got %d", code)
	}
	if !reflect.DeepEqual(ran, [][]string{{"flash", "x.hex"}, {"bogus"}, {"help"}}) {
		t.Errorf("a failed line returns to the prompt; the open-quote line and nested shell never run: %q", ran)
	}
	for _, s := range []string{"(exit 1)", "(exit 2)", "never closed - nothing was run", "Already at the Mujard CLI> prompt"} {
		if !strings.Contains(out, s) {
			t.Errorf("output lacks %q:\n%s", s, out)
		}
	}
}

func TestShellEndOfInputAndBlanks(t *testing.T) {
	code, ran, _ := shellRun(t, "\n   \n# a comment\nmujard\nEXIT\n", false, nil, nil)
	if code != 0 || len(ran) != 0 {
		t.Errorf("blank, comment and bare-mujard lines run nothing; EXIT is case-blind (exit %d, ran %q)", code, ran)
	}
	code, ran, _ = shellRun(t, "doctor", false, nil, nil)
	if code != 0 || !reflect.DeepEqual(ran, [][]string{{"doctor"}}) {
		t.Errorf("a last line without a newline still runs, then EOF exits 0 (exit %d, ran %q)", code, ran)
	}
	code, _, _ = shellRun(t, "", false, nil, nil)
	if code != 0 {
		t.Errorf("empty input exits 0, got %d", code)
	}
}

func TestShellInteractivePromptAndCancel(t *testing.T) {
	sig := make(chan os.Signal, 4)
	var out bytes.Buffer
	run := func(args []string) int {
		sig <- os.Interrupt // the user pressed Ctrl+C while the command ran
		return 1
	}
	code := runShell(strings.NewReader("run sk -b nano -p COM3\nexit\n"), &out, true, run, sig)
	s := out.String()
	if code != 0 || !strings.Contains(s, shellPrompt) || !strings.Contains(s, "Mujard command line") {
		t.Fatalf("interactive shell shows banner and prompt (exit %d):\n%s", code, s)
	}
	if !strings.Contains(s, "(cancelled)") || strings.Contains(s, "(exit 1)") {
		t.Errorf("Ctrl+C during a command reports it cancelled and stays in the shell:\n%s", s)
	}
}

// MUJARD_REQUIRE_LOGIN=1 makes the prompt refuse to open without a saved key - the owner's door
// policy on a shared machine. The per-command key checks still apply inside.
func TestShellRequireLogin(t *testing.T) {
	t.Setenv("MUJARD_REQUIRE_LOGIN", "1")
	t.Setenv("MUJARD_KEY", "")
	t.Setenv("MUJARD_CONFIG", filepath.Join(t.TempDir(), "none.json"))

	code, ran, out := shellRun(t, "version\n", false, nil, nil)
	if code != 1 || len(ran) != 0 {
		t.Fatalf("shell must refuse without a key (exit 1, nothing run); got %d, ran %q", code, ran)
	}
	if !strings.Contains(out, "sign-in") {
		t.Fatalf("refusal should say to sign in:\n%s", out)
	}

	t.Setenv("MUJARD_KEY", "mjd_test_x")
	code, ran, _ = shellRun(t, "version\n", false, map[string]int{"version": 0}, nil)
	if code != 0 || len(ran) != 1 {
		t.Fatalf("a key opens the shell; got %d, ran %q", code, ran)
	}
}

// The interactive banner tells an unsigned-in user what works and what needs login - and says
// nothing when a key is present.
func TestShellSignInBanner(t *testing.T) {
	t.Setenv("MUJARD_REQUIRE_LOGIN", "")
	t.Setenv("MUJARD_KEY", "")
	t.Setenv("MUJARD_CONFIG", filepath.Join(t.TempDir(), "none.json"))

	_, _, out := shellRun(t, "exit\n", true, nil, nil)
	if !strings.Contains(out, "Not signed in") {
		t.Fatalf("unsigned shell should explain what needs login:\n%s", out)
	}

	t.Setenv("MUJARD_KEY", "mjd_test_x")
	_, _, out = shellRun(t, "exit\n", true, nil, nil)
	if strings.Contains(out, "Not signed in") {
		t.Fatalf("signed-in shell should not nag:\n%s", out)
	}
}

func TestShellVerbIsBuiltIn(t *testing.T) {
	if findBuiltin("shell") == nil {
		t.Fatal("shell must be built in so a plugin can't shadow it")
	}
	var out bytes.Buffer
	if code := runShellCommand([]string{"extra"}, &out); code != 2 {
		t.Errorf("mujard shell <args> is a usage error, got %d", code)
	}
}
