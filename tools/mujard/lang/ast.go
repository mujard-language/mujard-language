package lang

// The syntax tree of a .muj program (spec section 3). Every node knows where it starts in the source.

// Node is any node of the tree.
type Node interface{ Start() Pos }

// File is a whole program.
type File struct {
	Board    string
	BoardPos Pos
	// Safety is the line after the board line, 'safety strict' (spec 6.10): "" when the program has none (standard).
	Safety    string
	SafetyPos Pos
	Items     []Node // top-level declarations and blocks, in source order
}

func (f *File) Start() Pos { return f.BoardPos }

// ---- top level

// UseDecl: use mujard.servo | use "Adafruit NeoPixel" version "1.12" | use idf "x" | use pio "x" | ... as name
type UseDecl struct {
	At      Pos
	From    string // "mujard" (lib path), "arduino" (a quoted name or file), "idf", "pio"
	Path    string // "mujard.servo", "Adafruit NeoPixel", "lib/crc.h"
	Version string
	Alias   string
}

// ImportDecl: import motors | import sonar(trig: pin 7, echo: pin 8)
type ImportDecl struct {
	At   Pos
	Name string
	Pins []ImportPin // the pins a driver module asks for with 'needs pin x' (spec 10.5)
}

// ImportPin is one 'name: pin N' of an import.
type ImportPin struct {
	At   Pos
	Name string
	Ref  PinRef
}

// NeedsDecl: needs i2c, pwm, pin trig - the hardware a driver module (or a program) asks of its board, by capability
// name, and the pins the importing program must give it (spec 10.5), checked before any C++ is made.
type NeedsDecl struct {
	At       Pos
	Names    []string
	Pos      []Pos // where each name is
	PinNames []string
	PinPos   []Pos
}

// VarDecl: const / var / saved var, at the top level or in a block.
type VarDecl struct {
	At    Pos
	Kind  string // "const", "var", "saved"
	Name  string
	Type  *TypeExpr
	Value Node
}

// DeviceDecl: led = pin 13 as output | light = analog A0 | digits = pins 2..9 as output | lid = Servo(pin 6)
type DeviceDecl struct {
	At      Pos
	Name    string
	Kind    string // "pin", "pins" (a range: Pins holds both ends), "pinlist", "analog", "object"
	Pins    []PinRef
	ToEnd   bool   // "pins": the range includes its last pin (..=)
	Role    string // "input", "output", "pwm"
	Options []string
	RoleAt  Pos
	Object  Node // Kind "object": the constructor call
}

// PinRef is a pin as written: 13, A0, GPIO5, LED.
type PinRef struct {
	At   Pos
	Text string
}

// FnDecl: fn name(params) -> type: block
type FnDecl struct {
	At     Pos
	Name   string
	Params []Param
	Result *TypeExpr
	Body   *Block
}

// Param is one function parameter.
type Param struct {
	At      Pos
	Name    string
	Type    *TypeExpr
	Default Node
}

// StructDecl: struct Name: fields
type StructDecl struct {
	At     Pos
	Name   string
	Fields []Param
}

// EnumDecl: enum Mode: idle, run, stop
type EnumDecl struct {
	At     Pos
	Name   string
	Values []string
}

// EventBlock: at start / at 07:30 / every / after / when / on ... : block (after and when may also be nested).
type EventBlock struct {
	At        Pos
	Kind      string // "start", "time", "every", "after", "when", "on"
	Expr      Node   // the period, delay, condition or event; nil for start/time
	Hour, Min int    // Kind "time"
	Fast      bool   // on ... fast
	As        string // the block's name (every/after/when) or the event value's name (on)
	Priority  uint64 // priority N on a top-level every/after/when block (0: the default)
	Core      int    // 'on core N': the core the block runs on (0: the main core, the default)
	Body      *Block
}

// TaskDecl: task name [priority N] [on core N]: block
type TaskDecl struct {
	At       Pos
	Name     string
	Priority uint64 // 0: the default
	Core     int    // 'on core N': the core the block runs on (0: the main core, the default)
	Body     *Block
}

// SpecialBlock: safe_state: / test "name": (Name is the test's name).
type SpecialBlock struct {
	At   Pos
	Kind string // "safe_state", "test"
	Name string
	Body *Block
}

// LimitDecl: limit motor to 0..80 %
type LimitDecl struct {
	At     Pos
	Target Node
	Range  Node
}

// ControlDecl: control "Speed", speed, 0..100 % - or a hold button (control "Horn", horn, hold) or a joystick
// (control "Drive", steer, throttle, -100 %..100 %: Var is the x axis, Var2 the y axis).
type ControlDecl struct {
	At    Pos
	Label string
	Var   string
	Var2  string
	Hold  bool
	Range Node
}

// CppBlock: cpp: followed by raw C++ lines.
type CppBlock struct {
	At   Pos
	Code string
}

// ---- statements

// Block is an indented list of statements.
type Block struct {
	At    Pos
	Stmts []Node
}

// IfStmt: if / elif ... / else
type IfStmt struct {
	At    Pos
	Conds []Node
	Thens []*Block
	Else  *Block
}

// WhileStmt: while cond: block
type WhileStmt struct {
	At   Pos
	Cond Node
	Body *Block
}

// ForStmt: for x in range-or-list: block
type ForStmt struct {
	At   Pos
	Var  string
	Iter Node
	Body *Block
}

// MatchStmt: match x: cases
type MatchStmt struct {
	At    Pos
	Value Node
	Cases []MatchCase
}

// MatchCase is one arm of a match; IsElse for "else:".
type MatchCase struct {
	At       Pos
	Patterns []Node
	IsElse   bool
	Body     *Block
}

// SimpleStmt: break, continue, return [x], wait d, wait until c, print x, pass.
type SimpleStmt struct {
	At    Pos
	Kind  string // "break", "continue", "return", "wait", "print", "pass"
	Until bool   // Kind "wait": pause until the condition is true, not for a duration
	Value Node
}

// ExpectStmt: expect cond, "message", halt: false
type ExpectStmt struct {
	At   Pos
	Cond Node
	Args []Arg // the first unnamed value is the message; 'halt:' is the only named option (spec 6.8)
}

// Message is the first unnamed argument, or nil (then the condition's own text is reported).
func (n *ExpectStmt) Message() Node {
	for _, a := range n.Args {
		if a.Name == "" {
			return a.Value
		}
	}
	return nil
}

// HaltArg is the 'halt:' option's value, or nil.
func (n *ExpectStmt) HaltArg() Node {
	for _, a := range n.Args {
		if a.Name == "halt" {
			return a.Value
		}
	}
	return nil
}

// AssignStmt: target op value (op is "=", "+=", ...)
type AssignStmt struct {
	At     Pos
	Target Node
	Op     string
	Value  Node
}

// ExprStmt: an expression used as a statement (a call).
type ExprStmt struct {
	At Pos
	X  Node
}

// ---- expressions

// Ident is a name.
type Ident struct {
	At   Pos
	Name string
}

// NumberLit is a number, maybe with a unit: 42, 3.14, 500 ms, 50 %.
type NumberLit struct {
	At      Pos
	Text    string
	IsFloat bool
	Int     uint64
	Float   float64
	Unit    string
}

// StringLit is text with optional {expression} parts.
type StringLit struct {
	At    Pos
	Parts []StringLitPart
}

// StringLitPart is plain text or an interpolated expression.
type StringLitPart struct {
	Text string
	Expr Node
}

// Literal: true, false, none.
type Literal struct {
	At    Pos
	Value string
}

// ListLit: [a, b, c]
type ListLit struct {
	At    Pos
	Items []Node
}

// PinLit: pin 6 / analog 35 used as a value.
type PinLit struct {
	At     Pos
	Analog bool
	Ref    PinRef
}

// Unary: -x, ~x, not x
type Unary struct {
	At Pos
	Op string
	X  Node
}

// Binary: x op y (arithmetic, bit, comparison, and, or)
type Binary struct {
	At   Pos // the operator's position
	Op   string
	X, Y Node
}

// RangeExpr: a..b or a..=b
type RangeExpr struct {
	At        Pos
	Lo, Hi    Node
	Inclusive bool
}

// Conversion: x as unit
type Conversion struct {
	At   Pos
	X    Node
	Unit string
}

// CondExpr: then if cond else other
type CondExpr struct {
	At               Pos
	Then, Cond, Else Node
}

// Call: f(args)
type Call struct {
	At   Pos
	Fn   Node
	Args []Arg
}

// Arg is a call argument, maybe named (over: 400 ms).
type Arg struct {
	Name  string
	Value Node
}

// Member: x.name
type Member struct {
	At   Pos
	X    Node
	Name string
}

// Index: x[i]
type Index struct {
	At    Pos
	X, Ix Node
}

// TypeExpr: int, text[32], list[int, 10], time?
type TypeExpr struct {
	At       Pos
	Name     string
	Args     []*TypeExpr // a numeric argument is a TypeExpr whose Name is the number
	Optional bool
}

func (n *UseDecl) Start() Pos      { return n.At }
func (n *ImportDecl) Start() Pos   { return n.At }
func (n *NeedsDecl) Start() Pos    { return n.At }
func (n *VarDecl) Start() Pos      { return n.At }
func (n *DeviceDecl) Start() Pos   { return n.At }
func (n *FnDecl) Start() Pos       { return n.At }
func (n *StructDecl) Start() Pos   { return n.At }
func (n *EnumDecl) Start() Pos     { return n.At }
func (n *EventBlock) Start() Pos   { return n.At }
func (n *TaskDecl) Start() Pos     { return n.At }
func (n *SpecialBlock) Start() Pos { return n.At }
func (n *LimitDecl) Start() Pos    { return n.At }
func (n *ControlDecl) Start() Pos  { return n.At }
func (n *CppBlock) Start() Pos     { return n.At }
func (n *Block) Start() Pos        { return n.At }
func (n *IfStmt) Start() Pos       { return n.At }
func (n *WhileStmt) Start() Pos    { return n.At }
func (n *ForStmt) Start() Pos      { return n.At }
func (n *MatchStmt) Start() Pos    { return n.At }
func (n *SimpleStmt) Start() Pos   { return n.At }
func (n *ExpectStmt) Start() Pos   { return n.At }
func (n *AssignStmt) Start() Pos   { return n.At }
func (n *ExprStmt) Start() Pos     { return n.At }
func (n *Ident) Start() Pos        { return n.At }
func (n *NumberLit) Start() Pos    { return n.At }
func (n *StringLit) Start() Pos    { return n.At }
func (n *Literal) Start() Pos      { return n.At }
func (n *ListLit) Start() Pos      { return n.At }
func (n *PinLit) Start() Pos       { return n.At }
func (n *Unary) Start() Pos        { return n.At }
func (n *Binary) Start() Pos       { return n.At }
func (n *RangeExpr) Start() Pos    { return n.At }
func (n *Conversion) Start() Pos   { return n.At }
func (n *CondExpr) Start() Pos     { return n.At }
func (n *Call) Start() Pos         { return n.At }
func (n *Member) Start() Pos       { return n.At }
func (n *Index) Start() Pos        { return n.At }
func (n *TypeExpr) Start() Pos     { return n.At }
