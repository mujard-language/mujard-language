package lang

import (
	"strings"
	"testing"
)

func w617(t *testing.T, src string) []Diagnostic {
	t.Helper()
	_, diags := CheckSourceModules(src, nil)
	if e := FirstError(diags); e != nil {
		t.Fatalf("should check without errors: %v", e)
	}
	var out []Diagnostic
	for _, d := range diags {
		if d.Code == "W617" {
			out = append(out, d)
		}
	}
	return out
}

func TestRecursionWarning(t *testing.T) {
	direct := w617(t, "board uno\nfn count(n: int) -> int:\n    if n < 1:\n        return 0\n    return count(n - 1) + 1\nat start:\n    print count(3)\n")
	if len(direct) != 1 || direct[0].Pos != (Pos{2, 1}) || !strings.Contains(direct[0].Message, "'count' calls itself:") || !strings.Contains(direct[0].Message, "2048 bytes") {
		t.Errorf("direct recursion: one W617 at the function, naming the Uno's RAM: %v", direct)
	}
	mutual := w617(t, "board uno\nfn ping(n: int):\n    if n > 0:\n        pong(n - 1)\nfn pong(n: int):\n    if n > 0:\n        ping(n - 1)\nat start:\n    ping(4)\n")
	if len(mutual) != 1 || !strings.Contains(mutual[0].Message, "'ping' calls itself (through pong)") {
		t.Errorf("recursion through another function: one W617 naming the chain: %v", mutual)
	}
	plain := w617(t, "board uno\nfn a() -> int:\n    return b() + 1\nfn b() -> int:\n    return 2\nat start:\n    print a()\n")
	if len(plain) != 0 {
		t.Errorf("a plain call chain isn't recursion: %v", plain)
	}
}
