package lang

import (
	"strings"
	"testing"
)

// TestTextCapacity: text[N] is a fixed-capacity text (spec 4.5). A declared size keeps its own buffer through
// declarations, assignments, interpolation, concatenation, methods, lists, optionals, functions, structs,
// saved variables and the store: values cut to the target's capacity, and text[128] holds what text can't.
func TestTextCapacity(t *testing.T) {
	cxx := testCxx(t)
	long := strings.Repeat("a", 70)
	src := `board uno
var t8: text[8] = "01234567"
var big: text[128] = ""
var s: text[8] = ""
var msgs: list[text[8], 4] = ["alpha", "betabetabeta"]
saved var keep: text[16] = "keep me"

fn shout(x: text[8]) -> text[8]:
    return x.upper()

struct Msg:
    body: text[8] = ""

at start:
    print t8
    s = "ab" + "cdefghi"
    print s
    s += "ij"
    print s
    var m = Msg(body: "hi there friend")
    print m.body
    print shout("hello world")
    print msgs[1]
    print msgs[0].upper()
    big = "` + long + `" + "TAIL"
    print big
    var u8 = t8.upper()
    print u8
    var sp = t8.split("5")
    print sp[1]
    print t8[3]
    if t8[7] == "7":
        print "last is 7"
    store.save("k", s)
    var o: text[8]? = store.load("k")
    print o
    print keep

test "text[N] keeps its own capacity":
    advance(100 ms)
    expect printed "01234567"
    expect printed "abcdefgh"
    expect printed "hi there"
    expect printed "HELLO WO"
    expect printed "betabeta"
    expect printed "ALPHA"
    expect printed "` + long + `TAIL"
    expect printed "67"
    expect printed "3"
    expect printed "last is 7"
    expect printed "keep me"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("text[N] test: %+v", res)
	}
}

// TestTextCapacityWarnings: a literal that can't fit the declared size is a W608, and the size is checked.
func TestTextCapacityWarnings(t *testing.T) {
	for _, tc := range []struct {
		name, src, code string
	}{
		{"over a small declared size", "board uno\nvar t: text[8] = \"012345678\"\n", "W608"},
		{"over the default size", "board uno\nvar t = \"" + strings.Repeat("x", 65) + "\"\n", "W608"},
		{"over a list element's size", "board uno\nvar l: list[text[4], 2] = [\"12345\"]\n", "W608"},
		{"over a big declared size is fine", "board uno\nvar t: text[128] = \"" + strings.Repeat("x", 70) + "\"\nat start:\n    print t\n", ""},
		{"a bad size", "board uno\nvar t: text[0] = \"hi\"\n", "C303"},
		{"too big", "board uno\nvar t: text[256] = \"hi\"\n", "C303"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var diags []Diagnostic
			if tc.code == "" {
				_, diags = Compile(tc.src, "p.muj")
				if e := FirstError(diags); e != nil {
					t.Fatalf("should compile: %v", diags)
				}
				return
			}
			_, diags = CheckSource(tc.src)
			found := false
			for _, d := range diags {
				if d.Code == tc.code {
					found = true
				}
			}
			if !found {
				t.Fatalf("want %s: %v", tc.code, diags)
			}
		})
	}
}
