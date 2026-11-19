package lang

import "testing"

// TestListLiteralInExpression: a list literal is a value anywhere a list goes (spec 4.5) - call arguments, a for
// range, methods on the temporary - not only a variable's initializer.
func TestListLiteralInExpression(t *testing.T) {
	cxx := testCxx(t)
	src := `board uno

fn total(l: list[int, 8]) -> int:
    var s = 0
    for x in l:
        s += x
    return s

fn second(l: list[text[4], 8]) -> text[4]:
    return l[1]

at start:
    print total([1, 2, 3, 4])
    for x in [7, 8]:
        print x
    print second(["abcde", "fghij"])
    var pair = [9, 11]
    print pair[1]
    print ["x", "y"].length
    print ["keep", "me"][1]

test "a list literal is a value":
    advance(100 ms)
    expect printed "10"
    expect printed "7"
    expect printed "8"
    expect printed "fghi"
    expect printed "11"
    expect printed "2"
    expect printed "me"
`
	res := runTests(t, cxx, src)
	if len(res) != 1 || !res[0].Passed {
		t.Fatalf("list literal test: %+v", res)
	}
}
