// Package lang is the Mujard compiler: lexer, parser, checker and C++ generator for .muj programs.
// The language is defined in docs/lang/MUJARD_LANGUAGE_SPEC.md; the test programs are in docs/lang/examples.
package lang

import "fmt"

// Kind is the kind of a token.
type Kind int

// Token kinds.
const (
	EOF       Kind = iota
	NEWLINE        // the end of a logical line
	INDENT         // the start of a more-indented block
	DEDENT         // the end of an indented block
	NAME           // led, button, Servo (and any word straight after '.')
	KEYWORD        // one of the 36 keywords
	INT            // 42, 0xFF, 0b1010, 1_000
	FLOAT          // 3.14, 1e3
	UNIT           // ms, kHz, m/s, % - only straight after a number
	STRING         // "text {with} parts"
	TIMEOFDAY      // 07:30 after 'at'
	OP             // operators and punctuation
	RAW            // the indented lines of a cpp: block, as written
)

var kindNames = [...]string{"EOF", "NEWLINE", "INDENT", "DEDENT", "NAME", "KEYWORD", "INT", "FLOAT", "UNIT", "STRING", "TIMEOFDAY", "OP", "RAW"}

func (k Kind) String() string {
	if int(k) < len(kindNames) {
		return kindNames[k]
	}
	return fmt.Sprintf("Kind(%d)", int(k))
}

// Pos is a place in a source file. Lines and columns count from 1; columns count characters (runes).
type Pos struct {
	Line int
	Col  int
}

func (p Pos) String() string { return fmt.Sprintf("%d:%d", p.Line, p.Col) }

// StringPart is one piece of a text literal: plain text, or an interpolated expression ("{t}").
type StringPart struct {
	Text   string  // plain text (escapes already applied), when Tokens is nil
	Tokens []Token // the tokens of an interpolated expression, without NEWLINE/EOF
	Pos    Pos     // where the part starts (for an expression: the '{')
	IsExpr bool
}

// Token is one token of a .muj file.
type Token struct {
	Kind  Kind
	Text  string // the source text (for STRING: the whole literal; for RAW: the block's lines, dedented)
	Pos   Pos
	Int   uint64       // INT value
	Float float64      // FLOAT value
	Parts []StringPart // STRING parts
	Hour  int          // TIMEOFDAY
	Min   int          // TIMEOFDAY
}

func (t Token) String() string {
	switch t.Kind {
	case NEWLINE, INDENT, DEDENT, EOF:
		return fmt.Sprintf("%s %s", t.Kind, t.Pos)
	}
	return fmt.Sprintf("%s %q %s", t.Kind, t.Text, t.Pos)
}

// Is reports whether the token is the keyword or operator s.
func (t Token) Is(s string) bool { return (t.Kind == KEYWORD || t.Kind == OP) && t.Text == s }

// Keywords are the 36 reserved words (spec 2.4).
var Keywords = map[string]bool{
	"board": true, "use": true, "import": true,
	"const": true, "var": true, "saved": true, "fn": true, "struct": true, "enum": true,
	"pin": true, "analog": true, "as": true,
	"if": true, "elif": true, "else": true, "while": true, "for": true, "in": true, "match": true, "break": true, "continue": true, "return": true,
	"every": true, "on": true, "when": true, "after": true, "at": true, "task": true, "wait": true,
	"and": true, "or": true, "not": true,
	"true": true, "false": true, "none": true,
	"cpp": true,
}

// Units are the unit names that may follow a number (spec 2.7). Compound units (m/s) are two of these joined by '/'.
var Units = map[string]string{
	"us": "time", "ms": "time", "s": "time", "min": "time", "h": "time",
	"Hz": "frequency", "kHz": "frequency", "MHz": "frequency",
	"mm": "length", "cm": "length", "m": "length",
	"mV": "voltage", "V": "voltage",
	"mA": "current", "A": "current",
	"deg":   "angle",
	"C":     "temperature",
	"%":     "ratio",
	"baud":  "baud",
	"bytes": "memory",
}

// The operators and punctuation, longest first so that "..=" wins over "..".
var operators = []string{
	"..=", "<<", ">>", "==", "!=", "<=", ">=", "+=", "-=", "*=", "/=", "..", "->",
	"+", "-", "*", "/", "<", ">", "=", "&", "|", "^", "~", ".", ",", ":", "(", ")", "[", "]", "?",
}
