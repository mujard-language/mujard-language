package lang

import "strings"

// TKind is the kind of a type the checker works with (spec section 4).
type TKind int

// Type kinds.
const (
	TAny    TKind = iota // unknown (C++ libraries, cpp: interop): never reported, to avoid false errors
	TVoid                // no value (a function that returns nothing)
	TInt                 // int and byte
	TFloat               // float
	TBool                // true / false
	TText                // text
	TQty                 // a number with a dimension (time, length, ...)
	TList                // list[T, N] and pin groups
	TRange               // a..b
	TStruct              // a struct value, or the struct itself when called
	TEnum                // an enum value
	TDevice              // a pin device: Name is output / input / pwm / analog / group
	TObject              // a built-in or library object: Name is its class
	TNone                // the none literal
)

// Type is a checker type. Types are compared by value via same().
type Type struct {
	K    TKind
	Dim  string // TQty: time, frequency, length, voltage, current, angle, temperature, ratio, speed, memory, baud
	Elem *Type  // TList / TRange items; TDevice "group" items
	Name string // TStruct / TEnum / TDevice role / TObject class
	Opt  bool   // may be none (T?)
	Cap  int    // TText / TList: fixed capacity; 0 = the default
}

var (
	tAny   = &Type{K: TAny}
	tVoid  = &Type{K: TVoid}
	tInt   = &Type{K: TInt}
	tFloat = &Type{K: TFloat}
	tBool  = &Type{K: TBool}
	tText  = &Type{K: TText, Cap: textMax}
	tNone  = &Type{K: TNone}
)

func qty(dim string) *Type { return &Type{K: TQty, Dim: dim} }
func opt(t *Type) *Type {
	c := *t
	c.Opt = true
	return &c
}
func nonOpt(t *Type) *Type {
	if !t.Opt {
		return t
	}
	c := *t
	c.Opt = false
	return &c
}
func listOf(t *Type) *Type         { return &Type{K: TList, Elem: t} }
func listCap(t *Type, n int) *Type { return &Type{K: TList, Elem: t, Cap: n} }
func textCap(n int) *Type          { return &Type{K: TText, Cap: n} }
func textCapOf(t *Type) int {
	if t != nil && t.Cap > 0 {
		return t.Cap
	}
	return textMax
}
func rangeOf(t *Type) *Type { return &Type{K: TRange, Elem: t} }
func object(class string) *Type {
	return &Type{K: TObject, Name: class}
}

// dimOfUnit gives the dimension a unit measures ("m/s" -> speed).
func dimOfUnit(u string) string {
	if strings.Contains(u, "/") {
		parts := strings.SplitN(u, "/", 2)
		if Units[parts[0]] == "length" && Units[parts[1]] == "time" {
			return "speed"
		}
		return ""
	}
	return Units[u]
}

// typeNamed turns a written type into a checker type (unknown names are user structs/enums, resolved by the caller).
func typeNamed(name string) (*Type, bool) {
	switch name {
	case "int", "byte":
		return tInt, true
	case "float":
		return tFloat, true
	case "bool":
		return tBool, true
	case "text":
		return tText, true
	case "time", "frequency", "length", "voltage", "current", "angle", "temperature", "ratio", "speed", "memory":
		return qty(name), true
	}
	return nil, false
}

// describe names a type in messages, the way a beginner would say it.
func describe(t *Type) string {
	s := ""
	switch t.K {
	case TAny:
		s = "a value"
	case TVoid:
		s = "nothing"
	case TInt:
		s = "a whole number"
	case TFloat:
		s = "a number"
	case TBool:
		s = "true/false"
	case TText:
		s = "text"
	case TQty:
		s = "a " + t.Dim
	case TList:
		s = "a list"
	case TRange:
		s = "a range"
	case TStruct:
		s = "a " + t.Name
	case TEnum:
		s = "a " + t.Name + " value"
	case TDevice:
		s = "a pin (" + t.Name + ")"
	case TObject:
		switch {
		case strings.HasPrefix(t.Name, "mod:"):
			s = "the module '" + strings.TrimPrefix(t.Name, "mod:") + "'"
		case strings.HasPrefix(t.Name, "fn:"):
			s = "the function '" + strings.TrimPrefix(t.Name, "fn:") + "'"
		case strings.HasPrefix(t.Name, "struct:"):
			s = "the struct '" + strings.TrimPrefix(t.Name, "struct:") + "'"
		case strings.HasPrefix(t.Name, "enum:"):
			s = "the enum '" + strings.TrimPrefix(t.Name, "enum:") + "'"
		default:
			s = "a " + t.Name
		}
	case TNone:
		s = "none"
	}
	if t.Opt {
		s += " (or none)"
	}
	return s
}

// textMax is how many characters (bytes) a text holds: the runtime's MUJARD_TEXT_MAX (spec 4).
const textMax = 64

// bare is describe without its leading article, for messages that bring their own ("A {type} has no ...").
func bare(t *Type) string {
	s := describe(t)
	if strings.HasPrefix(s, "an ") {
		return s[3:]
	}
	return strings.TrimPrefix(s, "a ")
}

// unitExample is how U501/U503 suggest a value of a dimension.
var unitExample = map[string]string{
	"time": "ms", "frequency": "Hz", "length": "cm", "voltage": "V", "current": "mA", "angle": "deg",
	"temperature": "C", "ratio": "%", "speed": "m/s", "memory": "bytes", "baud": "baud",
}

func isNumeric(t *Type) bool { return t.K == TInt || t.K == TFloat }

// assignable reports whether a value of type got may go where want is expected.
func assignable(want, got *Type) bool {
	if want.K == TAny || got.K == TAny {
		return true
	}
	if got.K == TNone {
		return want.Opt
	}
	if got.Opt && !want.Opt {
		return false
	}
	switch want.K {
	case TFloat:
		return isNumeric(got)
	case TInt:
		return got.K == TInt || (got.K == TObject && got.Name == "pin")
	case TQty:
		return got.K == TQty && got.Dim == want.Dim
	case TList, TRange:
		return got.K == want.K && (want.Elem == nil || got.Elem == nil || assignable(want.Elem, got.Elem))
	case TStruct, TEnum, TObject:
		return got.K == want.K && got.Name == want.Name
	case TDevice:
		return got.K == TDevice
	}
	return want.K == got.K
}

// mulDims and divDims are the products and quotients of dimensions that mean something (spec 5).
func mulDims(a, b string) (string, bool) {
	switch {
	case a == "speed" && b == "time", a == "time" && b == "speed":
		return "length", true
	case a == "time" && b == "frequency", a == "frequency" && b == "time":
		return "", true // a plain number
	}
	return "", false
}

func divDims(a, b string) (string, bool) {
	switch {
	case a == b:
		return "", true
	case a == "length" && b == "time":
		return "speed", true
	case a == "length" && b == "speed":
		return "time", true
	}
	return "", false
}
