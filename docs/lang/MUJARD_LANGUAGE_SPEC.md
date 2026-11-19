# Mujard language specification

Version 1.0-draft, revision 3 (2026-10-22); first written 2026-10-18 (Phase 0). Revisions are listed in section 15. This is the reference the Mujard compiler, runtime and tests are built against.
Built-ins with examples: `docs/lang/README.md`; built-in signatures: Appendix A below (this document wins if the two disagree).

**Decisions taken in Phase 0** (owner can still change them before Phase 1 code depends on them):

| Question | Decision |
|---|---|
| Blocks | indentation (spaces only), like Python |
| Functions | `fn` |
| Event words | `every`, `on`, `when`, `after`, `at start`, `task` / `wait` |
| File extension | **`.muj`**; **`.mjd`** is accepted everywhere `.muj` is (the compiler, `mujard fmt`, `import`). Files were `.usbp` before the rename |
| Language name | **Mujard**, after its author Mujarhid Addillah. Tagline: *Bare code. Every board.* A web search on 2026-10-19 found no product of that name; that is not trademark clearance |
| Remainder | the word `mod`, not `%`, so `%` always means percent (2.7) |
| Quantity sizes | `time` and `frequency` are 64-bit; other quantities 32-bit (5.1) |
| Named options | a bare name such as `crash` or `high` is accepted where the enum type is known (4.6) |
| `expect` | the message is optional; `halt: false` keeps the program running (6.8) |

Sections marked **(M1)** are needed for the first milestone (Blink, buttons, the dustbin on real boards); the rest
follow in later phases.

---

## 1. Source files (M1)

- A program is one or more `.muj` files in UTF-8. The main file starts with a `board` line (after optional comments and
  blank lines).
- Other files of the same project are modules, brought in with `import name` (the file `name.muj`). Its top-level
  names are reached as `name.member` (`motors.forward()`); nothing is copied into the importing file's names.
- Line endings: `\n` or `\r\n`.

## 2. Lexical structure (M1)

### 2.1 Comments
`#` to the end of the line. There are no block comments.

### 2.2 Indentation
- Only spaces may indent. A tab at the start of a line is error **L101**.
- A line that ends with `:` opens a block; the next non-blank line must be indented more than the opening line.
- A block ends at the first line indented less. The new indentation must equal an earlier level (else **L102**).
- Blank lines and comment-only lines never affect indentation.
- Inside `( )` and `[ ]`, newlines and indentation are ignored, so long calls and lists can span lines.
- The lexer turns indentation into `INDENT` and `DEDENT` tokens, and line ends into `NEWLINE` (one per logical line).

### 2.3 Names
`[A-Za-z_][A-Za-z0-9_]*`. Case matters. Keywords can't be names (**P211**), except straight after a `.`, where any word
is a member name (so `led.on()` is fine).

### 2.4 Keywords (36)

```
board   use     import
const   var     saved   fn      struct  enum
pin     analog  as
if      elif    else    while   for     in      match   break   continue  return
every   on      when    after   at      task    wait
and     or      not
true    false   none
cpp
```

### 2.5 Context words
Ordinary names everywhere except in the position shown.

| Words | Special only... |
|---|---|
| `input` `output` `pwm` | after `as` in a pin declaration |
| `pullup` `pulldown` `inverted` | after a pin role |
| `pins` | at the start of a pin-group declaration |
| `start` | after `at` |
| `fast` | after an event in an `on` header |
| `test` `limit` `safe_state` `control` | at the start of a top-level line |
| `to` | in a `limit` line |
| `print` `expect` | at the start of a statement (they take a value without parentheses) |
| `mod` | between two values (the remainder operator) |
| `until` | after `wait` |
| `core` | after `on` in a top-level event or task header (`on core N`) |
| `priority` | in a top-level `every`/`after`/`when`/`task` header (`priority N`) |
| `pass` | a statement on its own line |
| `version` `idf` `pio` | in a `use` line |
| unit names (2.7) | straight after a number, after `as`, and inside compound units |

### 2.6 Numbers

| Form | Examples | Type |
|---|---|---|
| decimal whole | `0`, `42`, `1_000_000` | `int` |
| decimal with point | `3.14`, `0.5`, `1e3`, `2.5e-3` | `float` |
| hexadecimal | `0xFF`, `0x1F_A0` | `int` |
| binary | `0b1010`, `0b1011_0001` | `int` |

- `_` may separate digits; not at the start, end, or twice in a row.
- A point needs digits on both sides: `.5` and `5.` are **L105**.
- A minus sign is an operator, not part of the number.

### 2.7 Units

A unit straight after a number (with or without a space: `500 ms`, `500ms`) makes a quantity.

| Dimension | Units | Base unit |
|---|---|---|
| time | `us` `ms` `s` `min` `h` | microsecond |
| frequency | `Hz` `kHz` `MHz` | millihertz |
| length | `mm` `cm` `m` | micrometre |
| voltage | `mV` `V` | microvolt |
| current | `mA` `A` | microamp |
| angle | `deg` | millidegree |
| temperature | `C` | millidegree Celsius |
| ratio | `%` | parts per million |
| baud rate | `baud` | baud |
| memory | `bytes` | byte |

- **Compound units** are written without spaces around `/`: `343 m/s`, `100 mm/s`. They are lexed as one unit token.
- `%` after a number is always the percent unit (`50 %`, `50%`). The remainder operator is the word `mod`: `10 mod 3`
  is 1, `x mod 3` is a remainder. `%` is never an operator.
- A unit word anywhere else is an ordinary name (`var m = 3` is allowed).

### 2.8 Time of day
`HH:MM` (two digits after the colon) straight after `at` is a time-of-day token: `at 07:30:`.

### 2.9 Text

- Double quotes only: `"Hello"`. Single quotes are **L106**.
- Escapes: `\n` `\t` `\"` `\\` `\{` `\}`. Anything else after `\` is **L108**.
- `{expression}` inside text inserts a value: `"Temp: {t}"`. A `{` without a matching `}` is **L107**.
- Text can't run over a line end (**L103**).

### 2.10 Operators and punctuation

```
+  -  *  /  ==  !=  <  <=  >  >=  =  +=  -=  *=  /=
&  |  ^  ~  <<  >>  ..  ..=  .  ,  :  ->  (  )  [  ]  ?
```

`?` only appears after a type (`time?`).

`%` only appears as a unit straight after a number. Remainder is the word `mod` (2.5).

Any other character outside text or comments is **L104**.

---

## 3. Grammar (M1 except where noted)

ISO EBNF (`,` sequence, `|` choice, `[ ]` optional, `{ }` any number of times, `(* *)` comment), plus two PEG
operators that look ahead without consuming anything: `& x` means "x must come next", `! x` means "x must not come next". This block is normative and machine-checked:
`tools/mujard/lang/grammar_test.go` reads it from this file, checks that every rule is defined and used and that every
keyword and operator of section 2 appears, accepts every sample in `examples/ok`, rejects every parser-error sample in
`examples/bad`, and runs the hand-written parser and this grammar side by side on thousands of altered programs: they
must accept and reject exactly the same ones.

Upper-case names are token classes from the lexer: `NAME`, `KEYWORD` (any of the 36 in 2.4), `INT`, `NUMBER` (`INT` or a
number with a point), `UNIT` (2.7), `STRING` (2.9; each `{…}` part inside it must be an `expr`), `TIMEOFDAY` (2.8),
`RAW_BLOCK` (the indented lines of a `cpp:` block), `NEWLINE`, `INDENT`, `DEDENT`, `EOF`. A quoted word matches that word:
a keyword, an operator, or a context word (2.5), which the lexer reads as a `NAME`.

```ebnf
program        = main_program | module ;
main_program   = { NEWLINE } , board_line , [ safety_line ] , { top_item } , EOF ;   (* the file given to the compiler; without the board line: P206 *)
module         = { NEWLINE } , { top_item } , EOF ;                (* a file brought in with 'import' (10.1) *)
board_line     = "board" , board_name , NEWLINE ;
safety_line    = { NEWLINE } , "safety" , NAME , NEWLINE ;      (* safety standard | safety strict (6.11) *)
board_name     = NAME | STRING ;                      (* uno, esp32, "arduino:avr:mega" *)

top_item       = use_line | import_line | const_decl | var_decl | saved_decl
               | device_decl | fn_decl | struct_decl | enum_decl
               | event_block | task_decl
               | safe_block | test_block | limit_line | control_line
               | needs_line | cpp_block | NEWLINE ;

use_line       = "use" , ( lib_path | STRING | "idf" , STRING | "pio" , STRING )
                 , [ "version" , STRING ] , [ "as" , NAME ] , NEWLINE ;
lib_path       = NAME , { "." , NAME } ;              (* mujard.servo *)
import_line    = "import" , NAME , [ "(" , import_pin , { "," , import_pin } , ")" ] , NEWLINE ;
import_pin     = NAME , ":" , "pin" , pin_ref ;               (* import sonar(trig: pin 7) *)
needs_line     = "needs" , need , { "," , need } , NEWLINE ;  (* needs i2c, pin trig *)
need           = [ "pin" ] , NAME ;

const_decl     = "const" , NAME , [ ":" , type ] , "=" , expr , NEWLINE ;
var_decl       = "var" , NAME , [ ":" , type ] , "=" , expr , NEWLINE ;
saved_decl     = "saved" , var_decl ;

device_decl    = NAME , "=" , device , NEWLINE ;
device         = "pin" , pin_ref , "as" , role , { pin_option }
               | "pins" , pin_ref , range_op , pin_ref , "as" , role , { pin_option }
               | "[" , "pin" , pin_ref , { "," , "pin" , pin_ref } , "]" , "as" , role , { pin_option }
               | "analog" , pin_ref
               | call_expr ;
call_expr      = primary , { suffix } , "(" , [ call_args ] , ")"   (* Servo(pin 6), Ultrasonic(trig: pin 9, echo: pin 10) *)
               | NAME , STRING ;                      (* a device made by a call written with text: x = ask "Ready?" *)
pin_ref        = INT | NAME ;                         (* 13, A0, D5, GP15, GPIO5, LED *)
role           = "input" | "output" | "pwm" ;
pin_option     = "pullup" | "pulldown" | "inverted" ;

fn_decl        = "fn" , NAME , "(" , [ param , { "," , param } , [ "," ] ] , ")" , [ "->" , type ] , ":" , block ;
param          = NAME , ":" , type , [ "=" , expr ] ;
struct_decl    = "struct" , NAME , ":" , NEWLINE , INDENT , field , { field } , DEDENT ;
field          = NAME , ":" , type , [ "=" , expr ] , NEWLINE ;
enum_decl      = "enum" , NAME , ":" , NAME , { "," , NAME } , NEWLINE ;

event_block    = ( "at" , "start"
                 | "at" , TIMEOFDAY
                 | "every" , expr , [ "as" , NAME ] , [ "priority" , INT ] , [ "on" , "core" , INT ]
                 | "after" , expr , [ "as" , NAME ] , [ "priority" , INT ] , [ "on" , "core" , INT ]
                 | "when" , expr , [ "as" , NAME ] , [ "priority" , INT ] , [ "on" , "core" , INT ]
                 | "on" , event_ref , [ "fast" ] , [ "as" , NAME ] ) , ":" , block ;
event_ref      = postfix ;                            (* button.pressed, button.held(2 s), serial.line *)
task_decl      = "task" , NAME , [ "priority" , INT ] , [ "on" , "core" , INT ] , ":" , block ;
safe_block     = "safe_state" , ":" , block ;
test_block     = "test" , STRING , ":" , block ;
limit_line     = "limit" , postfix , "to" , expr , NEWLINE ;      (* 'to' is a context word here *)
control_line   = "control" , STRING , "," , NAME , [ "," , expr , [ "," , expr ] ] , NEWLINE ;
cpp_block      = "cpp" , ":" , RAW_BLOCK ;           (* the indented lines, passed through as written *)

type           = NAME , [ "[" , type_args , "]" ] , [ "?" ] ;      (* int, text[32], list[int, 10], time? *)
type_args      = ( type | INT ) , { "," , ( type | INT ) } ;

block          = NEWLINE , INDENT , stmt , { stmt } , DEDENT ;
stmt           = var_decl | const_decl
               | if_stmt | while_stmt | for_stmt | match_stmt
               | "break" , NEWLINE | "continue" , NEWLINE
               | "return" , [ expr ] , NEWLINE
               | "wait" , ( expr | "until" , expr ) , NEWLINE   (* a bare 'wait until' waits on a variable called until *)
               | "pass" , NEWLINE
               | "print" , expr , NEWLINE
               | "expect" , expr , [ "," , args ] , NEWLINE   (* the first unnamed value is the message; halt: false *)
               | nested_event                         (* 'after' and 'when' inside a block *)
               | cpp_block
               | command_stmt
               | simple_stmt ;
nested_event   = ( "after" , expr | "when" , expr ) , [ "as" , NAME ] , ":" , block ;
command_stmt   = callee , & command_start , ! "mod" , args , NEWLINE ;   (* a call without parentheses: notify "Hi", show "T", t, log.info "x"; 'total mod i' is a remainder *)
callee         = ( NAME | "board" ) , { "." , NAME } ;
command_start  = STRING | NUMBER | NAME | "true" | "false" | "none" | "not" | "pin" | "analog" ;
simple_stmt    = expr , [ assign_op , expr ] , NEWLINE ;   (* the target of an assignment must be a variable, field or list item (P213) *)
assign_op      = "=" | "+=" | "-=" | "*=" | "/=" ;

if_stmt        = "if" , expr , ":" , block , { "elif" , expr , ":" , block } , [ "else" , ":" , block ] ;
while_stmt     = "while" , expr , ":" , block ;
for_stmt       = "for" , NAME , "in" , expr , ":" , block ;
match_stmt     = "match" , expr , ":" , NEWLINE , INDENT , case , { case } , DEDENT ;
case           = ( pattern , { "," , pattern } | "else" ) , ":" , block ;
pattern        = expr ;                               (* 3, "on", Mode.idle, idle, 1..5 *)

expr           = cond_expr ;
cond_expr      = or_expr , [ "if" , or_expr , "else" , cond_expr ] ;   (* a if c else b *)
or_expr        = and_expr , { "or" , and_expr } ;
and_expr       = not_expr , { "and" , not_expr } ;
not_expr       = "not" , not_expr | compare ;
compare        = range_expr , [ ( comp_op | "in" ) , range_expr ] ;
comp_op        = "==" | "!=" | "<" | "<=" | ">" | ">=" ;
range_expr     = conversion , [ range_op , conversion ] ;
range_op       = ".." | "..=" ;
conversion     = bit_or , [ "as" , unit_name ] ;
unit_name      = UNIT | "us" | "ms" | "s" | "min" | "h" | "Hz" | "kHz" | "MHz" | "mm" | "cm" | "m"
               | "mV" | "V" | "mA" | "A" | "deg" | "C" | "baud" | "bytes" ;   (* 2.7 *)
bit_or         = bit_xor , { "|" , bit_xor } ;
bit_xor        = bit_and , { "^" , bit_and } ;
bit_and        = shift , { "&" , shift } ;
shift          = sum , { ( "<<" | ">>" ) , sum } ;
sum            = product , { ( "+" | "-" ) , product } ;
product        = unary , { ( "*" | "/" | "mod" ) , unary } ;
unary          = ( "-" | "~" ) , unary | postfix ;
postfix        = primary , { suffix } ;
suffix         = "." , NAME | "(" , [ call_args ] , ")" | "[" , expr , "]" ;
call_args      = args , [ "," ] ;                     (* a trailing comma is allowed inside ( ) and [ ] *)
args           = arg , { "," , arg } ;
arg            = [ arg_name , ":" ] , expr ;          (* named arguments: over: 400 ms *)
arg_name       = NAME | KEYWORD ;                     (* any word, keywords too: on: 19 C *)
primary        = NUMBER , [ UNIT ] | STRING | "true" | "false" | "none"
               | NAME , STRING                        (* a call with one text value: if ask "Ready?": *)
               | NAME
               | "board" , & "."                      (* the built-in board object: board.uptime *)
               | "(" , expr , ")" | "[" , [ list_items ] , "]"
               | "pin" , pin_ref                      (* a pin value, e.g. Servo(pin 6) *)
               | "analog" , pin_ref ;                 (* an analog input value, e.g. battery_on(analog 35, ...) *)
list_items     = expr , { "," , expr } , [ "," ] ;
```

Notes (explanations; the productions above are the rule):
- **Where a word could be read two ways**, the compiler takes the statement form only when a value follows:
  `print x` prints, `print = 3` assigns to a variable called `print`; `wait until x` waits for x, a bare `wait until`
  waits on a variable called `until`; `notify "Hi"` is a call, `total mod i` is a remainder.
- **One comparison per expression:** `a < b < c` is **P212**; for a range write `x in a..b`, otherwise split it:
  `a < b and b < c`.
- **`=` inside a condition** (`if x = 3:`) is **P214**.
- **Statements at the top level** other than the `top_item`s (for example a bare `led.on()`) are **P216**.
- `as` after an expression is a conversion when the next word is a unit; in an event header it names the block. A block
  name can't be a unit name.
- **Commands:** a call that is a whole statement may leave out its parentheses: `notify "Door open"`,
  `show "Temperature", t`, `plot t, h`, `log.info "Started"`, `serial 9600 baud`. The arguments run to the end of the
  line (named ones too: `notify "Fire!", urgent: true`). `print` and `expect` are written this way as well.
- In an expression, a name followed directly by text is a call with that text: `if ask "Ready?":`.
- A named argument may be any word, keywords included: `hysteresis(t, on: 19 C, off: 21 C)`.
- `board` is a keyword, but `board.` followed by a name is the built-in board object (`board.uptime`,
  `board.halt("reason")`).

---

## 4. Types (M1 except lists/structs)

| Type | Values | Notes |
|---|---|---|
| `int` | -2 147 483 648 to 2 147 483 647 | `+ - *` wrap on overflow (two's complement: `2147483647 + 1` is `-2147483648`), on every board; constant overflow is warning **W607** |
| `byte` | 0 to 255 | used for raw data; mixes with `int` as `int` |
| `float` | 32-bit floating point | |
| `bool` | `true`, `false` | |
| `text` | up to 64 characters by default; `text[N]` fixes the size to N characters (1 to 255; other sizes are **C303**) | never grows past its size; longer results are cut. Text written in the program that is already longer than its size is warning **W608** |
| quantities | `time`, `frequency`, `length`, `voltage`, `current`, `angle`, `temperature`, `ratio`, `speed`, `memory` | a number with a dimension (section 5) |
| `list[T, N]` | up to N items of type T | fixed capacity; `.length`, indexing from 0 |
| structs, enums | user-defined | |
| devices | `Output`, `Input`, `Pwm`, `Analog`, `OutputGroup`, library classes | made by device declarations only |
| `T?` | a T or `none` | results that can fail |

### 4.1 Type inference
`var x = 5` is an `int`, `var t = 20 C` a `temperature`, `var on = false` a `bool`. A declared type wins:
`var x: float = 5`.

### 4.2 Mixing
- `int` with `float` gives `float`. `byte` with `int` gives `int`.
- `/` between two `int`s is whole-number division, truncated towards zero (as Arduino). Make one side a float for
  decimals: `x / 2.0`. `mod` takes the sign of the left side (`-7 mod 3` is `-1`). Dividing by zero (`/`, `mod`, `/=`,
  a `time` divided by 0) and shifting by less than 0 or more than 31 (`<<`, `>>`) are runtime errors (6.8).
  `-2147483648 / -1` wraps to `-2147483648`. `>>` keeps the sign.
- None of these depend on the C++ compiler: the generated code never uses C++'s own `int` operators, whose overflow
  and division by zero are undefined there.
- Comparing values of different dimensions is **U502**.

### 4.3 `none` and `T?`
- Functions that can fail return `T?`: `echo.measure_pulse()`, `store.load()`, `s.to_number()`, `i2c.read()`, `ask`.
- A `T?` must be checked before use as a `T`; otherwise **C304**. Checks narrow the type for the rest of the block:

```muj
var t = echo.measure_pulse(timeout: 30 ms)
if t == none:
    return
print "{t}"                # fine: t is a time here
```

Narrowing rules:
- After an `if t == none:` block that ends with `return`, `break` or `continue`, `t` is narrowed for the rest of the block.
- Inside `if t != none:`, and inside the `else:` of `if t == none:`, `t` is narrowed.
- In `t != none and ...` and in `t == none or ...`, `t` is narrowed in the rest of the expression.
- Narrowing applies to variables and parameters, not to properties: `board.voltage` read twice may give different
  answers, so copy it first (`var v = board.voltage`).
- A `when`, `every` or `if` condition that uses a `T?` without one of these checks is **C304**.

### 4.4 Lists and pin groups
- `var r: list[int, 10] = []` holds up to 10 ints. Methods: `.length`, `.add(x)` (a runtime error when full),
  `.remove(i)`, `.clear()`; items by index from 0: `r[0]`. `for x in r:` visits the items in order.
- A pin group (`pins 2..9 as output`, `[pin 3, pin 5] as pwm`) is a fixed list of devices: `leds[i].on()`,
  `for l in leds:`, and for output groups `leds.write(byte)` sets them all from one value.

### 4.5 Text
- `+` joins text. `{…}` in text formats any value; quantities print with a readable unit (`1.5 s`, `250 ms`).
- Methods: see the README, section 17.

### 4.6 Named options and built-in enums
Built-ins that take or give a fixed set of named values use enums:

| Type | Members | Used by |
|---|---|---|
| `ResetReason` | `power_on` `reset_button` `watchdog` `brownout` `software` `wake` `crash` | `board.reset_reason` |
| `Level` | `high` `low` | `measure_pulse(level:)` |
| `BatteryType` | `lipo_1s` `aa_x3` `aa_x4` (more in Phase 8) | `battery_on(type:)` |
| `ShowStyle` | `number` `gauge(range)` `light` | `show(as:)` |

A bare member name (`crash`, `idle`) means that enum's member when all of these hold: the enum type is known from the
context (compared with `==` or `!=` against a value of that enum, a `match` over it, an assignment to a variable or field
of that type, or an argument for a parameter of that type), and no variable, const, function or device with that name is in
scope. A name that is in scope always wins. The full form (`ResetReason.crash`, `Mode.idle`) always works. Anything else
is **C301**, and the message suggests the full form.

---

## 5. Quantities and units (M1)

- A quantity is a number plus a dimension. `500 ms` and `0.5 s` are equal.
- `+`, `-` and comparisons need the same dimension (**U502**).
- `*` and `/` by a plain number keep the dimension (`2 * 500 ms` is `1 s`).
- `length / time` is a `speed`; `speed * time` is a `length`; `1 / time` is a `frequency`; `time * frequency` and
  `ratio` values are plain numbers. Other products are **U502**.
- Where a quantity is expected, a plain number is **U501** ("500 needs a unit here"), except `0`, which fits any
  dimension.
- **`x as unit`** keeps the value and fixes the unit it is shown in: `print "{d as cm}"` prints `12.4 cm`. To pass a plain
  number to C++ or a library, use `.number`: `(d as cm).number` is `12.4`.
- On a plain number, `as unit` makes a quantity of that unit. It is how a number from outside gets its unit:
  `reply.json["temp"] as C`.

### 5.1 Representation
- `time` is a signed 64-bit count of microseconds and `frequency` a signed 64-bit count of millihertz. A 32-bit count
  would overflow after 35 minutes (time) or at 2 MHz (frequency), which `every 1 h`, `deep_sleep(30 min)` and
  `SPI 8 MHz` need. All other quantities are 32-bit, and `board.uptime` and `now()` never wrap.
- The compiler may use 32-bit arithmetic where it can prove the values fit (constants, periods under 35 minutes). The
  speed and size gates are measured on an Uno with this rule.
- `ratio`: `50 %` is 0.5. A ratio times a number scales it: `200 * 50 %` is `100`.
- `scale(x, a..b, c..d)` maps between ranges and may change the dimension (`0.5 V..1.5 V` to `0 C..100 C`). It does not
  clamp: values outside `a..b` map outside `c..d`. Use `clamp` (or a `limit`) when that matters.

---

## 6. Programs and the scheduler (M1)

Mujard programs are event-driven. There is no `setup()`/`loop()`; the runtime runs this cycle forever:

1. **Start-up, once:** `safe_state` runs; global variables get their initial values in source order; saved variables
   are loaded; every `at start` block runs, in source order; every task starts.
2. **Each pass of the main loop:**
   1. inputs are sampled and debounced; line input is buffered when the program uses `read_line()`;
   2. matching `on` handlers run, in source order;
   3. every `when` condition is checked, in priority order;
   4. every due `every`, `after` and `at HH:MM` block runs, in priority order;
   5. every task resumes until its next `wait` or its end, in priority order;
   6. housekeeping: saved variables, the watchdog, messages to the app.

Blocks run to completion and are never interrupted by other blocks (except `fast` events, 6.5). That is why `wait` is
only allowed in tasks: a block that waited would stop everything else. In steps 3-5, `priority N` on a top-level
`every`, `after`, `when` or `task` runs it before lower numbers inside that step; `priority 0` is the default and
equal priorities keep source order. A priority never moves a block between steps or interrupts one that is running.

### 6.1 `at start:`
Runs once at start-up, in source order with other `at start` blocks.

### 6.2 `every P [as name] [priority N]:`
- First runs on the first pass after start-up, then every P measured from its schedule (not from when it finished), so
  it doesn't drift.
- If a run is late by a whole period or more, the missed runs are skipped (no burst of catch-up runs) and
  `board.overruns` counts them.
- The checker warns about the overrun it can prove (**W619**): on a board whose serial port goes through a USB bridge,
  a block that always prints more plain text each run (plus print's `\r\n`) than the port sends in one period - 10 bit
  times per character at the program's `serial(...)` speed, 115200 baud by default. Prints under an `if` or in a loop,
  and the text of `{...}` parts, don't count, so the real output can only be longer. Native-USB boards (Leonardo,
  Micro, Pico) have no such limit.
- `P` must be a time or a frequency (`every 2 Hz:` is every 500 ms), and constant.
- With a name: `name.stop()`, `name.start()`, `name.running`.
- `priority N` is an integer dispatch order for due timed blocks (higher first; 0 is the default).

### 6.3 `after D [as name] [priority N]:`
- At the top level: runs once, D after start-up.
- Inside a block: each time that statement runs, it (re)starts a one-shot timer; the block runs D later. Running the
  statement again before then restarts the timer (the latest wins). Each `after` statement has its own timer.

### 6.4 `when C [as name] [priority N]:`
- Runs once each time `C` goes from `false` to `true`. Its "previous" value starts as `false`, so it runs once at
  start-up if `C` is already true.
- At the top level it is checked every pass. Inside a block it is checked each time that statement runs, with its own
  memory.

### 6.5 `on event [fast] [as value]:`
- Runs once when its event is seen on a pass. Pin events are sampled state changes, not queues. `serial.read_line()`
  and `uart.read_line()` keep bounded input-line queues (serial: 2 lines on AVR, 4 elsewhere; UART: 4); a line that
  arrives while the queue is full is dropped and `board.dropped_events` counts it.
- `as value` binds the event's value, for events that carry one (`on serial.line as cmd:`).
- `fast` uses a hardware interrupt on a pin event: the handler runs immediately, even in the middle of another block. Not
  every pin can interrupt; others are **B411**.
- Pin events are `changed`, `rising` and `falling`; `fast` also works on `pressed` and `released`. **`rising` and
  `falling` follow the wire**: `rising` is the voltage going from LOW to HIGH, whatever `pullup` or `inverted` say.
  `pressed` and `released` follow the active level (7). Without `fast`, all of them are debounced like `changed`.
  Variables wider than a byte that a `fast` handler writes are read and written by the rest of the program with
  interrupts briefly off, so a value is never half-updated (a 32-bit counter on an 8-bit chip). `fast` events are not
  debounced.
- **What a `fast` handler may do (C314).** It may read and assign variables of numbers, `bool`, quantities and enums;
  use `if`, `match`, `for` over a range, `break`, `continue`, `return`; call `on`, `off`, `toggle`, `write` and `read`
  on pins and pin groups; call `abs`, `min`, `max`, `clamp`, `round`, `between`, `now`, `since` and the bit methods;
  call the program's own functions that follow these same rules (checked through every call they make); and call
  `board.halt` with its reason written out as plain text. Everything else is **C314**: `print`, `log`, `expect`,
  `notify` and the other app built-ins, anything that waits, `while`, library and device calls other than pins,
  properties of devices and objects (`light.value`, `board.uptime`), text values, list changes, assigning a
  `saved var`, `store`, `after`/`when` inside the handler, and starting or stopping named blocks and tasks. `cpp:`
  blocks stay unchecked, as everywhere (W604).
- **A halt inside a `fast` handler.** `board.halt` and runtime errors (6.8) in a `fast` handler lock every output at
  once; the log line and `safe_state` follow on the next pass of the main loop, outside the interrupt. From the moment
  a program starts halting, `fast` handlers return without doing anything.
- A button with a `double_clicked` handler delivers `clicked` only after the 400 ms double-click window has passed;
  without one, `clicked` arrives at release.

### 6.6 `task name [priority N]:`
- Starts at start-up and runs alongside everything else. `wait D` pauses only this task.
- When it reaches its end, it stops. `name.start()` / `name.restart()` run it again from the top; `name.stop()` stops it.
- `ask`, `http.get`, `http.post`, `serial.read_line()`, `uart.read_line()` and other calls that wait for an answer are
  only allowed in tasks (**C311**); Appendix A marks each one. `mqtt.connect` and `mqtt.publish` are allowed everywhere:
  in a task they wait until done, elsewhere they queue the work and return at once.
- `wait until cond` pauses the task until `cond` is true instead of for a fixed time. It is checked on every pass of
  the main loop, so the condition must be something cheap: a pin, a variable or the clock.

### 6.7 Stopping
- `return`, `break`, `continue` work as in other languages.
- `board.halt("reason")` runs `safe_state`, stops every block and task, blinks the status LED and reports the reason.
  USB and uploads keep working.
- A program never "ends" on its own: with nothing left to do, the board waits.

### 6.8 Runtime errors
An out-of-range list index, division by zero, a shift out of range (4.2), or a failed `expect` is a runtime error:
the runtime logs `log.error` with the `.muj` line, runs `safe_state`, and halts (as `board.halt`). The board never
runs on with garbage, and never simply crashes.
- From the halt on, every output stays as `safe_state` left it: pin writes, PWM, tones, servos, shift registers and
  I2C writes are ignored. The rest of the block that failed may still run to its end, but it can't switch anything
  back on. (`cpp:` blocks are not covered: they are unchecked C++.)
- A failed `expect` does the same unless it is written with `halt: false`: then it logs the message (at most once per
  10 s for each line), runs `safe_state` and the program carries on, and the check runs again at its next turn. Use
  `halt: false` for conditions that can clear by themselves (a hot motor that cools down) and the default for
  conditions that mean the program can't be trusted. After a halt the board stays halted until it is reset or
  power-cycled; USB and uploads keep working.

### 6.9 Built-ins with memory
`smooth`, `changed`, `rising`, `falling`, `hysteresis` and `peak` remember values between runs.
- Each place in the source where one is written has its own memory (like `when`, 6.4). The memory is updated only when
  that line runs: a call skipped by `and`, `or` or an `if` does not update it.
- A function that uses one shares that memory between all its callers.
- First run: `rising` (like `when`) starts from `false`, so a condition that is already true counts as a rise once;
  `falling` and `changed` take the first value they see as the starting point and don't fire; `hysteresis` starts
  `false`; `smooth` averages only the readings it has had so far.

### 6.10 Sleeping
- `power.sleep(t[, wake_on: event])` stops the **whole program** (every block, task and timer) for up to `t`, or until
  the event. Time asleep counts as time passed: `every` blocks that were due run once on waking and the skipped ones
  are counted in `board.overruns`. It is meant for programs where nothing else needs to run.
- `power.deep_sleep(t)` does not return. The board restarts on waking and `at start` blocks run again. On targets
  that distinguish sleep wake-ups `board.reset_reason` is `wake`; AVR wakes through its watchdog and reports
  `watchdog`.
- Before either sleep, and before `board.restart()` and `board.halt`, pending `saved var` changes are written (11).

### 6.11 Safety level
- A program may say how strict the compiler is, on the line after `board`: `safety standard` (the default, the same as
  writing nothing) or `safety strict`. Any other word is **C329**.
- Under `safety strict` the warnings that mark code the compiler can't check, or a resource the program may not really
  have, are errors: raw C++ (`cpp:`, W604), a library without a description (W611), a pin or timer a library also uses
  (W609, W613), RAM that is mostly used (W614), an analog pin Wi-Fi blocks (W615), a register written behind a
  device (W616) and a function that calls itself (W617: each call takes more stack, which the RAM estimate can't see).
  Each gets **C328**, which quotes the warning; the warning itself stays, so the reason is still shown.
- It is a property of the program (a school can ask for it, a safety-related project should): nothing is switched off
  in the runtime, and a strict program that compiles behaves exactly as it would without the line.

---

## 7. Pins and devices (M1)

```muj
led    = pin 13 as output
relay  = pin 7  as output inverted
button = pin 2  as input pullup
lamp   = pin 9  as pwm
light  = analog A0
digits = pins 2..9 as output
leds   = [pin 3, pin 5, pin 6] as pwm
lid    = Servo(pin 6)
```

- Device declarations are top-level only, and a device can't be reassigned.
- **Pin names** follow the board's printed labels, from the board database: numbers, `A0`..., `D0`... (NodeMCU),
  `GP0`... (Pico), `GPIO0`... (ESP32), and `LED` for the built-in LED. Unknown names are **B402**.
- **Roles:** `output`, `input`, `pwm`; `analog` has its own form.
- **Options:** `pullup`/`pulldown` (inputs only, else **B410**), `inverted` (any role).
- **Active level.** `on()`, `read()` and the button events all speak in terms of the pin's *active* level. An output is
  active HIGH. An input is active HIGH, except that a `pullup` input is active LOW (a button to ground reads `true`
  when pressed). **`inverted`** swaps the active level of any pin: `output inverted` drives LOW for `on()`;
  `input inverted` is active LOW; `input pullup inverted` is active HIGH (the rare button wired to the supply).
  The one exception is the `rising` and `falling` events, which always follow the voltage on the wire (6.5).
- **Checks** (board database): PWM-capable (**B403**), analog-capable (**B404**), input-only (**B405**), flash pins
  (**B406**), used twice (**B407**), analog-only (**B408**), serial pins (**W601**), strapping pins (**W602**), pins taken
  by a library (**W609**: the board's I2C pins with `mujard.i2c` or an I2C library, its SPI pins with `spi` or an SPI
  library), and PWM pins whose timer a Servo library or `tone`/`melody` takes on AVR boards (**W613**: Servo takes 9
  and 10 on an Uno, Nano, Leonardo or Micro and 44-46 on a Mega; `tone` takes 3 and 11 on an Uno or Nano, 5 on a
  Leonardo or Micro, 9 and 10 on a Mega). The pins are the Arduino cores' defaults, from the board database.
- **Ownership**: a pin has one owner - a device line, or the object whose constructor it is written in
  (`arm = Servo(pin 9)` owns pin 9) - and a second use is **B407**, naming the first owner. A hardware timer has one
  owner too: when two parts of a program each bring the same timer's interrupt, the build can't link, so the checker
  says so first (**B414**: `tone`/`melody` and the NewPing library both take Timer 2 on an Uno, Nano or Mega). Only
  conflicts shown with the real toolchain are listed; shared buses (I2C, SPI) may have many users.
- **Capabilities** (the compiler's evidence file): what a program asks of its board - outputs, inputs, PWM, analog
  inputs, fast events, tone, the `mujard.*` libraries, Wi-Fi, the clock, HTTP, MQTT, more than one CPU core, DMA,
  hardware timers, CAN, Bluetooth, native USB-device hardware, an RTOS runtime and an Ethernet MAC - each has a level per board:
  *language* (Mujard has it), *backend* (the generator produces it for that board), *build* (a real build of a program
  using it succeeded for that board) or *hardware* (it ran on a real board and was seen working). Only the build matrix
  writes *build*, and only a real board run writes *hardware*. A program that needs something below *build* on its
  board gets one warning at the `board` line listing those capabilities (**W618**); `mujard boards caps <board>` prints
  the table.
- **Memory** (board database): when what the program itself keeps in RAM - global variables, devices, tasks, timed
  blocks and the memory of `smooth`, `peak`, ... - is about more than 60 % of the board's RAM, the checker warns
  (**W614**). It is an estimate in the sizes of an 8-bit AVR; the runtime and libraries come on top. `mujard build`
  then prints what the linker really used, warns above 75 % of RAM, and explains a program that doesn't fit.
- Members of each device type: README sections 7-9.

---

## 8. Functions, structs, enums (functions M1)

```muj
fn average(a: float, b: float) -> float:
    return (a + b) / 2

fn beep(times: int = 1):
    for i in 0..times:
        buzzer.tone(2 kHz, 100 ms)
```

- Parameters need types; defaults allowed. Return type after `->`; a function without `->` returns nothing.
- Every path of a function with `->` must return a value (**C313**).
- Calls use positional and named arguments: `beep(times: 3)`.
- Functions can't contain `wait` unless only called from tasks (**C305**).
- `struct` groups fields with defaults; values are made with `Reading(t: 20 C)`; fields with `.`.
- `enum Mode: idle, run, stop`; values are `Mode.idle`; `match` over an enum must cover every value or have `else`
  (**C315**).

---

## 9. Control flow (M1)

- `if / elif / else`, `while`, `for x in <range or list>`, `match`, `break`, `continue`, `return` as usual.
- Ranges: `a..b` excludes `b`, `a..=b` includes it. If `a > b`, the range counts down (`10..=1` is 10, 9, ..., 1).
- Membership: `x in a..b` is true when `x` is inside the range; `x in list` is true when an item equals `x`. The right
  side must be a range or a list.
- `pass` is a statement that does nothing: the placeholder for a block that doesn't need to do anything yet.
- `while` loops with no `wait` in a task, and loops in event blocks, must finish quickly: a loop that runs longer than
  the watchdog time is a runtime error (6.8). The checker warns about `while true` outside tasks (**W610**).

---

## 10. Libraries, modules and C++ (external libraries Phase 8)

```muj
use mujard.servo                              # a Mujard library
use "Adafruit NeoPixel" version "1.12.3"   # an Arduino library from the library index
use idf "esp32-camera"                     # an ESP-IDF component (ESP32 boards)
use pio "bblanchon/ArduinoJson"            # a PlatformIO registry library
use "lib/crc.h"                            # a C/C++ file in the project
import motors                              # motors.muj in the project
```

- Arduino and C++ libraries are described to the checker by description files generated from their headers (Phase 8).
  Without one, calls still compile but are checked less (**W611**).
- A library's `architectures=` must include the board's, else **B412**.
- Library names keep their own spelling; snake_case aliases also work (`read_temperature` for `readTemperature`).
- A description (`.mlib`, JSON) can say that a class `extends` another of the same library (it has the base's methods
  too, as a C++ subclass does), that a class is `abstract` (only extended - `Adafruit_ADS1X15()` is an unknown name,
  `Adafruit_ADS1115()` is the one to make), and can describe a `families` of generated classes too many to list:
  U8g2's displays are `U8G2_<display>_<F|1|2>_<bus>`, each extending `U8G2` with the constructor its bus has. Only
  displays the library really has are known names, so a misspelt one is an error before the build, not a C++ error.
- What `use "…"` names becomes an `#include`, so it may only be a library name or a file inside the project: letters,
  digits, spaces and `_ . - + /`, starting with a letter, digit or `_`, with no `..` or `.` folder and no drive or
  leading `/`. Anything else is **C324** - the build never reads a file outside the project.

### 10.1 Modules: `import name`
- `import motors` loads `motors.muj` from the same folder as the program. No such file: **C321**.
- The module's top-level names - variables, `const`s, functions, structs, enums and devices - are reached as
  `motors.name` (`motors.forward()`, `motors.speed`). Nothing is copied into the program's names, and the
  program's names aren't visible inside the module.
- A module is itself a program: it is checked completely, and its messages point at lines in its own file. A
  `board` line there is optional; a module without one is checked against the importing program's board.
- The module's event blocks and tasks don't run in the importing program: only its declarations are shared. In
  the generated C++ they live in `namespace motors` inside `mujard_user`.
- A module can't `import` another module (**C322**). From outside, only a module's `var`s can be changed
  (`motors.speed = 3`); a `const`, device, function or type can't (**C306**).

### 10.2 `cpp:` blocks
- At the top level: global C++ declarations. Inside a block: C++ statements.
- Mujard variables, devices and functions are visible by their Mujard names (renamed ones as listed in the generated header).
- The checker doesn't check the C++ (**W604**, once per block); the C++ compiler does, and its errors point at the
  `.muj` line.

### 10.3 Names in the generated C++
User names that are C++ keywords or Arduino names (`new`, `class`, `delay`, `min`) are renamed `u_<name>` in the
generated C++. Users never see this unless they read the C++.

### 10.4 What 1.0 leaves out
I2C and SPI slave mode, direct timer or register access, USB keyboard/mouse, and ESP32 extras (touch, DAC) are not
built in; use `cpp:` or a library. Extra serial ports are in the library `mujard.uart`
(`Uart(rx: pin 4, tx: pin 3, speed: 9600 baud)`): `write(text)` sends the text as it is, `on x.line as l:` runs for
each line that arrives. It is a hardware port where the board has a spare one (ESP32: any pins; Pico: UART0/UART1
pins, else a PIO port on any pins; Mega: 19/18, 17/16, 15/14; Leonardo/Micro: 0/1) and a software port otherwise
(Uno and Nano: **W612** above 9600 baud). More `Uart`s than the board has ports is a runtime error. A halted program
sends nothing. `x.read_line()` waits inside a task until a full line arrives.

**CAN bus.** `use mujard.can` opens the board's CAN controller (`bus = Can(pin 21, pin 22, 500000 baud)`; the ESP32
  family's TWAI peripheral): `bus.send(id, bytes)` queues a frame of up to 8 data bytes and gives `false` when the
  queue is full, `bus.read()` gives the next received frame or `none`, and `on bus.message as m:` runs for each frame
  that arrives, `m.id` `m.data` `m.rtr`. Boards without a CAN controller refuse (`Can(...)` is **C320** there;
  `needs can` is **B415**).

**Bluetooth.** `use mujard.ble` turns the board into a BLE peripheral a phone or another board can connect to:
  `ble.start("name")` advertises, `ble.send(text)` notifies whatever is connected (false when nothing is),
  `ble.read_line()` gives the next line or `none`, `ble.connected` is `true` while something is, and
  `on ble.line as l:` runs for each line that arrives. Boards without a BLE backend refuse (**C320**; the ESP32
  family except S2, and the Nano 33 BLE).

**Ethernet.** `use mujard.ethernet` brings up a wired link on the boards that have one (the classic ESP32's EMAC +
  external PHY): `ethernet.start()` powers it and takes an address by DHCP, `ethernet.connected` is `true` while the
  wire is up, `ethernet.ip` is the address as text, and `on ethernet.joined`/`on ethernet.lost` fire when it comes up
  or drops. `http.get`/`mqtt`/`clock` work over the wire too. Boards without an EMAC refuse (**C320**).

**USB device.** `use mujard.usb` makes the board a keyboard+mouse to whatever it's plugged into: `usb.type("text")`
  types, `usb.key("a")` taps a key, `usb.mouse_move(x, y)` and `usb.click()` drive the pointer, each `false` when the
  HID device isn't up, and `usb.connected` says whether it is. Standard-HID boards only - Leonardo, Micro, Pico,
  Pico 2, Zero and Uno R4; the others refuse (**C320**).

**Multicore.** A top-level `every`/`after`/`when`/`task` can run on a second core: `every 100 ms on core 1:`,
  `task watch on core 1:`. The main core keeps the default scheduler; the block or task runs on the named core
  (ESP32-family's FreeRTOS task pinning, RP2040/RP2350's `multicore_launch_core1`). A board with too few cores
  refuses (**B415**).

**Hardware timers.** `use mujard.timer` makes a `Timer` object that triggers at hardware precision - a real
  timer interrupt, not the cooperative `every`: `t.every(1 s)`/`t.freq(38 khz)`/`t.once(500 ms)` arm it,
  `t.start()`/`t.stop()` control it, `t.elapsed` reads it, `on t.tick` runs on the scheduler and `on t.tick fast`
  runs in the ISR itself. Boards without a timer backend refuse (`Timer()` is **C320**).

**Queues.** `use mujard.queue` makes a bounded queue for handing values between blocks - and it is the safe way to
  pass data between a secondary core and the main one, where a plain `var` races (**W620**): `q = Queue(text, 8)`
  declares a queue of up to 8 texts (also `int`, `float`, `bool` or a measurement type; capacity 1 to 64).
  `q.push(x)` copies a value in and is `false` when the queue is full - the item is dropped and counted in
  `board.dropped_events`; `q.pop()` gives the oldest value, or `none` when empty; `q.available`, `q.full` and `q.n`
  report the state (a consumer typically drains with `every` + `if q.available:`; `when q.available:` fires once on
  the edge, like every `when`). On ESP32-family boards the queue is a real FreeRTOS
  queue; on RP2040/RP2350 a fixed ring guarded by the SDK's `mutex_t`; on single-core boards it is the same bounded
  ring. A queue object is the one global the cross-core warning doesn't flag: sharing it is the point.

**Register access.** `use mujard.regs` adds the `regs` object: typed register and bit-field names for the board's
chip, generated from the chip maker's own description files (the ATDF data in the AVR io headers for
ATmega328P/32U4/2560; the CMSIS-SVD files for ESP32/S2/S3/C2/C3/C6/H2, ESP8266, RP2040, nRF52840, SAMD21, RA4M1
and STM32F103 - every board in the catalog is covered; a future board whose chip isn't gets **B413**). `regs.PORTB.read()` reads, `regs.PORTB.write(v)`/`set(bits)`/`clear(bits)` write, and `regs.PORTB.PORTB0`
is that bit-field's mask as an int (so `regs.PORTB.set(regs.PORTB.PORTB0)` sets just that bit). A misspelt register or field (**C308**) and a literal that doesn't fit the register
(**C327**) are compile errors. Register writes are not covered by the output lock after a halt (6.8), the same as
`cpp:` blocks, and a write/set/clear on a GPIO port register that drives a declared pin warns (**W616**).

### 10.5 Drivers: `needs`
A driver is a module that speaks to one part (a sensor, a display, a motor board). It says once what it asks of the
board, by capability name (the Capabilities list in section 7, `mujard boards caps <board>`):

```muj
# light.muj - a BH1750 light sensor on I2C
needs i2c
use mujard.i2c

var measure: list[int, 1] = [0x10]

fn lux() -> int:
    i2c.write(0x23, measure)
    var raw = i2c.read(0x23, 2)
    if raw == none:
        return 0
    return (raw[0] * 256 + raw[1]) * 10 / 12
```

- Each name must be a capability Mujard knows (**C330**); a capability the board doesn't have is **B415**, reported
  in the module's own file, naming the module - so importing a Wi-Fi driver into an Uno program fails in plain words
  before any C++ is made.
- What a module needs counts as what the program needs: a capability not yet built for the board is in the
  program's **W618**.
- `needs` is a word only at the start of such a line; a device or variable called `needs` is still allowed.
- A program may say `needs` too (it documents and checks the board choice), but its own devices and calls already
  count.

**Pins a driver is given.** A driver doesn't know where the program wired its part, so it asks for named pins and
the program gives them when it imports the driver:

```muj
# sonar.muj
needs pin trig, pin echo
t = pin trig as output
e = pin echo as input
```

```muj
board uno
import sonar(trig: pin 7, echo: pin 8)
```

- Every `pin trig` in the module becomes the pin the importer gave, before the module is checked: the board's pin
  rules (B402-B408) and pin ownership (B407, with the program's own pins) apply to the real pin, and the generated
  C++ uses it.
- A pin the module asks for and the import doesn't give is **C331** (one error, at the `import`); a pin the import
  gives that the module doesn't ask for is **C332**; `needs pin` in a program, which nothing imports, is **C333**.

---

## 11. Saving, safety and the app (Phase 9 for app features)

- `saved var x = v` (top level only): loaded at start-up, written when changed, at most once per 5 s on boards with
  wear-levelled storage (ESP32) and once per 60 s on the others. The runtime spreads the writes over the free EEPROM
  or flash area, so one cell is not rewritten every time. Pending changes are written immediately before
  `power.sleep`, `power.deep_sleep`, `board.restart()` and `board.halt`; a power cut or crash can lose changes made
  in the last interval. A change of type or a new program version with a different layout starts it fresh. Allowed types: numbers, `bool`, quantities, enums,
  `text` up to 64 characters (**C312**).
- `store.save(name, value)` / `store.load(name, default:)`: named values the program saves when it chooses. The
  store keeps 8 names (a 9th logs a warning and isn't kept) in its own area at the end of the EEPROM / flash, so it
  never overlaps `saved var`. A save is in memory at once and written with the same pacing as `saved var` (and
  before sleep, restart and halt), so saving in a loop can't wear the storage out. Text keeps 18 characters.
  `store.load` gives the default's type; loading a name saved with another type gives the default. Without a
  default it gives `T?`, with `T` taken from the declared target (`var x: int? = store.load("n")`); without a
  declared target it is **C320**. Values: as `saved var` (**C312**; an optional value can't be saved).
  Because writes are paced, the program can watch them: `store.pending` is true while a `store.save` write or a
  `saved var` change hasn't reached storage yet, `store.failed` counts commits that didn't land, and
  `store.committed` / `store.last_commit` report the most recent commit's result.
- `safe_state:` (one per program, top level): runs at start-up before anything else, on runtime errors, and on
  `board.halt`. It must not wait, ask or use the network (**C316**).
- `limit device to range` (top level): every later write to the device is clamped; the first clamp logs a warning.
- `expect cond[, "message"][, halt: false]`: a runtime check (6.8). Without a message, the text of the condition is used.
- `control "Label", variable[, range]` (top level): an app control; values from the app are clamped to the range.
  `control "Label", flag, hold` is a hold button: `flag` is true only while it's pressed. `control "Label", x, y, range`
  is a joystick: two number variables (x and y axis) with one range. Both are **held controls**: the app refreshes them
  while a finger is on them, and when 500 ms pass without a refresh (the phone left, the link dropped) the board puts
  them back at rest itself - false, or 0 (the range end nearest to 0 when 0 is outside it). A hold on a variable that
  isn't true/false, or a joystick axis that isn't a number, is **C303**.
- `notify`, `show`, `plot`, `ask`, `secret`: README section 1.

### 11.1 Secrets
- `secret(name)` has type `text` or `credentials`, as chosen when it is created in the app (a Wi-Fi entry holds a network
  name and a password, so it is `credentials`). `wifi.connect` takes `credentials`; `mqtt.connect(password:)` takes
  `text`. A `credentials` value can only be passed on, never printed or sent.
- The value is put into the firmware only in the temporary build copy. It never appears in `.muj` files,
  `--show-cpp` output, `.gen/` folders, golden test files or build logs. A secret is written only where it is used:
  as the network of `wifi.connect`, or `user:` / `password:` of `mqtt.connect`. Anywhere else - a variable, joined
  text, `print`, `log`, `notify`, `show`, `plot`, a publish or a post - is **C323**, so there is no path by which the
  value could leave the board. The name is plain text: letters, digits and `_`, starting with a letter, at most 30
  characters (**C325**).
- How it is built: `secret("mqtt")` is the C macro `MQTT`; Wi-Fi credentials `secret("home_wifi")` are `HOME_WIFI_SSID`
  and `HOME_WIFI_PASSWORD`. The build defines them from the project's secrets (the server writes `mujard_secrets.h` into
  the build copy only); `mujard build` on a PC takes them from the environment variables `MUJARD_SECRET_<NAME>`. A
  missing one stops the build on the `.muj` line that needs it, naming what to add.
- Firmware is not source: on boards without flash encryption, anyone holding the board can read the secrets from its
  flash. Sharing the `.muj` file is safe; handing over the board is not.

## 12. Tests (Phase 10)

```muj
test "lid opens when a hand is near":
    hand.pretend(10 cm)
    advance(200 ms)
    expect lid.angle == 90 deg, "lid should be open"
```

- Test blocks are top level only, run with `mujard test`, and are never compiled into firmware.
- `pretend`, `advance`, `answer`, `notified`, `printed`, `stopped`, `halted`, `power_cut` and `brown_out` are only
  allowed in tests (**C317**).
- **Where tests run.** On the PC, on a simulated board: the program and the Mujard runtime are compiled for the PC
  against a fake Arduino with simulated time, so `advance(3 s)` takes milliseconds and a run is the same every time.
  Analog readings use the board's range and reference (1023 and 5 V on AVR boards, 4095 and 3.3 V on the ESP32).
  The PC needs a C++ compiler (`g++` or `clang++`, or `MUJARD_CXX`). Programs that use C++ libraries (`use "…"`)
  can't be tested this way yet. The same tests also run on an emulated chip: `mujard test --plan` writes each test
  as data, and the simulator's test driver (sim-web, Pico) carries it out against the real firmware (Phase 10).
- **Each test** starts a fresh board: the program starts up as it would on the board (6, step 1), then the test runs
  as a task beside it, so a test may also `wait`. It passes when it reaches its end.
- `advance(t)` lets the whole program run on for `t` of simulated time (`wait t` does the same).
- `x.pretend(v)` sets what a device reads from then on, as if the wire changed: an input `true`/`false` (its active
  level, 7; debouncing and events work as on the board), an analog pin a reading, a voltage or a `%` of the range,
  an `Ultrasonic` a distance, or `none` when nothing answers. Outputs, PWM pins and servos have no `pretend`: read
  them (`lamp.read()`, `lid.angle`).
- `i2c.pretend(address, reply)` puts a device on the simulated I2C bus: it acknowledges its address (a scan finds
  it) and answers every `i2c.read` with the bytes of `reply`, a list, from the first. `i2c.pretend(address, none)`
  takes it off the bus again, so a test can check what the program does when a sensor fails while it runs. Without a
  pretend device the bus is silent.
- `power_cut()` cuts the board's power and starts it again: the program starts from the beginning (6, step 1) with
  only what had been written to the EEPROM/flash - a `saved var` keeps its value if it was saved before the cut
  (spec 11 says how often) and loses a change made since. The outside world stays as it was: pretended inputs,
  analog readings, distances and I2C devices (an input with `pullup` reads released until it is pretended again).
  The test goes on at the next line; its own `var`s start again too, and `printed` sees what was printed since
  power came back. It is written on its own line directly in the test (inside an `if` or a loop it is **C320**).
- `brown_out()` dips the supply enough for the brown-out detector to reset the board, rather than taking the power
  away. The program starts again as after `power_cut()`, except `board.reset_reason` is `brownout` in the PC test
  (a power cut leaves it `power_on`). It is also written on its own line directly in the test (**C320** inside a block).
- `printed("x")` is true when the program printed a line containing `x` since the start; `notified("x")` likewise for
  a notification (`notify`).
- A failed `expect` ends the test with its message and line; the program itself isn't stopped. A runtime error or
  `board.halt` in the program fails the test with the error. A test that hasn't ended after an hour of simulated time
  fails.

---

## 13. Error and warning catalogue

The compiler reports `file:line:col  CODE  message`. Messages are in plain words and may add a suggestion. The tests
match the code and the message text; `{…}` parts are filled in.

### Lexer (L1xx)
| Code | Message |
|---|---|
| L101 | Tabs aren't allowed for indentation. Use spaces (4 per level). |
| L102 | This line's indentation doesn't match any block above it. |
| L103 | This text has no closing ". |
| L104 | The character '{c}' isn't used in Mujard. |
| L105 | '{text}' isn't a valid number. |
| L106 | Use double quotes for text: "...". |
| L107 | A '{' in this text has no matching '}'. Write \{ for a plain brace. |
| L108 | '\{c}' isn't a valid escape. Use \n, \t, \", \\, \{ or \}. |
| L109 | The first line can't be indented. |
| L110 | This '{open}' is never closed. |
| L111 | There's no '{open}' to match this '{close}'. |
| L112 | '{text}' isn't a valid time of day. Write it as HH:MM, for example 07:30. |

### Parser (P2xx)
| Code | Message |
|---|---|
| P201 | Add ':' at the end of this line. |
| P202 | This line needs an indented block below it. |
| P203 | This line is indented, but nothing above it opens a block. |
| P204 | Something is missing here: a value or expression. |
| P205 | Didn't expect '{tok}' here. |
| P206 | Start the program with the board it's for, for example: board uno |
| P207 | '{kw}' needs an 'if' just above it. |
| P208 | Write a pin like this: led = pin 13 as output |
| P209 | Unknown pin role '{x}'. Use input, output or pwm. |
| P211 | '{kw}' is a Mujard keyword and can't be used as a name. |
| P212 | Comparing needs one operator. For a range write 'x in a..b'; otherwise split it: 'a < b and b < c'. |
| P213 | You can only assign to a variable, a field or a list item. |
| P214 | Use '==' to compare; '=' changes a variable. |
| P215 | '{kw}' is only allowed at the top level of the program, not inside a block. |
| P216 | Code at the top level must be inside a block such as 'at start:' or 'every 1 s:'. |

### Checker (C3xx)
| Code | Message |
|---|---|
| C301 | '{name}' isn't defined. |
| C302 | '{name}' is already defined (line {line}). |
| C303 | Expected {want} here, but this is {got}. |
| C304 | '{name}' may be none here. Check it first: if {name} == none: ... |
| C305 | 'wait' only works inside a task. In this block, use 'after' or another 'every' instead. |
| C306 | '{name}' is a const and can't be changed. |
| C307 | '{fn}' takes {n} values, but {m} were given. |
| C308 | A {type} has no '{member}'. |
| C309 | 'return' only works inside a function. |
| C310 | '{kw}' only works inside a loop. |
| C311 | '{call}' waits for an answer, so it can only be used inside a task. |
| C312 | Only numbers, true/false, units, enums and short text can be saved. |
| C313 | '{fn}' must return a value on every path. |
| C314 | A fast handler may only change variables and pins: {what} isn't allowed here. |
| C315 | This match doesn't cover {missing}. Add them or an 'else:'. |
| C316 | safe_state must not wait, ask or use the network. |
| C317 | '{name}' only works inside a test. |
| C318 | '{name}' needs 'use mujard.{lib}' at the top of the program. |
| C319 | '{name}' belongs to the block around this 'after', which has finished by the time it runs. Use a var at the top level instead. |
| C320 | '{feature}' isn't available yet: it comes with {phase}. |
| C321 | Can't find the module '{name}' - there's no '{file}' next to this program. |
| C322 | '{what}' isn't allowed in a module. |
| C323 | A secret must never be printed or sent. |
| C324 | '{path}' can't be used: name a library, or a file inside the project (no '..', no drive or leading '/', only letters, digits, spaces and _ . - + /). |
| C325 | A secret's name is written as plain text in quotes - letters, digits and _, starting with a letter, at most 30 characters - not {name}. |
| C326 | '{tz}' isn't a time zone Mujard knows. Write a name such as "Africa/Lagos" or "Europe/London", or a POSIX rule such as "WAT-1". |
| C327 | '{n}' doesn't fit in {reg}: the register is {bits} bits wide. |
| C328 | This program is 'safety strict', so this is an error: {warning} |
| C329 | '{level}' isn't a safety level. Write 'safety standard' or 'safety strict'. |
| C330 | '{name}' isn't a capability Mujard knows. 'needs' takes: {list}. |
| C331 | The module '{mod}' needs pin '{pin}'. Say which: import {mod}({pin}: pin N). |
| C332 | The module '{mod}' doesn't ask for a pin '{pin}'. The pins it needs: {list}. |
| C333 | 'needs pin {pin}' is for driver modules: the program that imports one gives it its pins. |

Unknown names and members add "Did you mean '{suggestion}'?" when a close match exists.

### Board and pins (B4xx)
| Code | Message |
|---|---|
| B401 | Unknown board '{b}'. Try uno, nano, mega, esp32, esp8266 or pico. |
| B402 | There's no pin {p} on the {board}. |
| B403 | Pin {p} can't do PWM on the {board}. PWM pins: {list}. |
| B404 | Pin {p} isn't an analog input on the {board}. Analog pins: {list}. |
| B405 | {p} is input-only on the {board} - it can't be an output. |
| B406 | {p} is wired to the {board}'s flash chip. Using it will crash the board. |
| B407 | Pin {p} is already used by '{name}' (line {line}). |
| B408 | {p} on the {board} is analog only - it can't be a digital pin. |
| B409 | '{feature}' needs {what}, and the {board} doesn't have it. |
| B410 | Only inputs can use pullup or pulldown. |
| B411 | Pin {p} can't use 'fast' on the {board}. Interrupt pins: {list}. |
| B412 | The library '{lib}' doesn't support the {board}. |
| B413 | mujard.regs doesn't describe the {board}'s registers yet (chip {chip}). |
| B414 | '{a}' and '{b}' both need the {board}'s {timer} interrupt, so the program can't be built. Use only one of them on this board. |
| B415 | {who} needs {cap}, and the {board} doesn't have it. |

### Units (U5xx)
| Code | Message |
|---|---|
| U501 | {value} needs a unit here. Did you mean {value} {suggestion}? |
| U502 | Can't combine {a} and {b}: they measure different things. |
| U503 | Expected a {want} (like {example}), but this is a {got}. |
| U504 | '{u}' isn't a unit Mujard knows. |

### Warnings (W6xx)
| Code | Message |
|---|---|
| W601 | Pins 0 and 1 are the USB serial port on the {board}. Uploads and 'print' stop working while something is connected there. |
| W602 | {p} is a boot (strapping) pin on the {board}. What's wired to it can stop the board from starting. |
| W603 | '{name}' is never used. |
| W604 | This cpp: block isn't checked by Mujard. |
| W607 | This number doesn't fit in an int. |
| W608 | This text may be longer than {n} characters and will be cut. |
| W609 | Pin {p} is used by the library '{lib}'. |
| W610 | 'while true' outside a task stops everything else. Use 'every' or put it in a task. |
| W611 | The library '{lib}' has no description yet, so its calls aren't checked. |
| W612 | The {board} has no spare serial port, so this Uart is a software port: above 9600 baud it may lose characters. |
| W613 | Pin {p} can't do PWM on the {board} while '{lib}' is used: it takes the pin's timer. |
| W614 | This program's own values need about {n} bytes of the {board}'s {ram} bytes of RAM, before the runtime and libraries take theirs. Shorter lists and fewer texts (64 characters each) free RAM. |
| W615 | While Wi-Fi is on, {p} can't be read: on the {board} it sits on the ADC that Wi-Fi blocks. Free analog pins: {list}. |
| W616 | Writing {reg} also drives {list}, which the program declared: the write reaches the pin behind the device's back. |
| W617 | '{fn}' calls itself{through}: every call takes more stack RAM, and the {board} has {ram} bytes, so a deep recursion can overwrite the program's own variables. |
| W618 | Mujard hasn't built {list} for the {board} yet, so this program may not compile there. |
| W619 | This block prints at least {n} characters every {period}; at {baud} baud that takes {need}, so the serial port falls behind and runs are skipped (board.overruns). Print less, or less often. |
| W620 | '{name}' is shared between a secondary core and the main one. The cores don't synchronize variables: a read can see a half-written value, and compound writes (n = n + 1) race. Keep {name} to one core. |
| W621 | '{what}' from a secondary core isn't synchronized on the {board}: the board's driver for it isn't safe to share between cores. Keep it on the main core. |
| W622 | '{what}' writes to flash on the {board}, and the other core keeps running while flash is programmed - that can crash the board. Do it on the main core. |
| W623 | Pins {a} and {b} share PWM slice {slice} on the {board} - a slice has one frequency, so tone({af}) and tone({bf}) can't both hold. Use the same frequency or different pins. |

---

## 14. Versioning

- This document is language edition **Mujard 2026**. Programs don't state an edition until a second one exists; the syntax
  for choosing one will be added then.
- The compiler follows semantic versions. Programs that compile under 1.x keep compiling under every later 1.x. The
  compiler's test suite holds it to this: a frozen set of Mujard 2026 programs (`tools/mujard/lang/testdata/compat/2026`)
  must keep compiling, warnings allowed.

---

## 15. Revision 2 (2026-10-20): what changed and what the code must follow

Found in review; the compiler, runtime and samples built against revision 1 need these changes.

| Area | Change | Touches |
|---|---|---|
| Lexer | `%` is only a unit; `mod` is the remainder word | `lexer.go`, token tests, any sample using `%` as remainder |
| Parser | `board` line optional (modules); `expect` takes optional message and named args; `match` patterns accept `Mode.idle` | `parser.go`, golden trees |
| Checker | narrowing rules (4.3); bare enum members (4.6); secrets can't be printed (C323); `as unit` on a plain number | `check.go`, `types.go`, `builtins.go` |
| Types and runtime | 64-bit `time` and `frequency` (5.1); interrupt-safe variables for `fast`; `clicked` delay; memory per call site (6.9); `power.sleep` pauses everything (6.10) | codegen, `mujard_rt`, size and speed gates |
| Saving | write interval, wear levelling, flush before sleep, restart and halt (11) | `mujard_rt` storage |
| Pins | one rule for active level and `inverted` (7) | checker, runtime, README |
| Library | `mujard.uart` (extra serial ports) | Phase 8 |
| Docs | Appendix A (this file) is the normative list of signatures | README |

### 15.1 Revision 3 (2026-10-22): decisions after the audit

Taken after the 2026-10-22 implementation audit. Spec text is updated; the code follows in Milestone A
unless marked otherwise.

| Decision | Rule | Status |
|---|---|---|
| `rising` / `falling` | events on inputs that follow the voltage on the wire; `pressed` / `released` follow the active level (6.5, 7) | built 2026-10-22 |
| `fast` handlers | the exact allowed set and C314 for the rest, checked through called functions; a halt or runtime error locks the outputs at once and finishes outside the interrupt (6.5) | built 2026-10-22 |
| Built-in enums | `ResetReason`, `Level`, `BatteryType` are real enum types: `board.reset_reason == 42` is C303 (4.6 already said so) | built 2026-10-22 (`board.reset_reason` now has ESP and Optiboot-AVR backends; unsupported targets are still C320) |
| Text sizes | `text[N]` fixes the size to N characters (1 to 255; 64 is the default); written text longer than its target's size is W608 | built 2026-10-22 |
| Where tests run | on the PC, on a simulated board (the program + runtime against a fake Arduino, simulated time): fast and repeatable enough to grade many students; the emulated chip (sim-web) can run the same test build later (12) | built 2026-10-23 |
| Low-level access | a generated, typed `mujard.regs` library, not new syntax (10.4) | built: `regs.PORTB` reads/writes/sets/clears the chip's registers, `regs.PORTB.PORTB5` is a bit mask (10.4) |
| Board facts | peripheral pins (I2C, SPI, UART), timers that libraries take, and RAM / flash sizes join the board database, generated from the same source as the pins, so library pin claims (W609) can be per board | I2C, SPI and Servo-timer pins built 2026-10-22 (W609, W613); the rest Milestone C |

---

## Appendix A. Built-in signatures

**Takes** lists the arguments (`name:` ones are named). **Gives** is the result (`T?` can be `none`). **Where:** *any* =
any block; *task* = waits for an answer, tasks only (C311); *test* = tests only (C317); *top* = top level only;
*fast-ok* = allowed in a `fast` handler (otherwise C314). ◆ = has its own memory (6.9). Numbers match the README.
Capacities in lists (`list[int, 16]`) are defaults chosen here; confirm them in Phase 4.

| # | Built-in | Takes | Gives | Where / notes |
|---|---|---|---|---|
| 1 | `print` | text | nothing | any, not fast |
| 2 | `notify` | text; `urgent:` bool | nothing | any, not fast; one per 10 s per line |
| 3 | `show` | label text, value; `as:` ShowStyle | nothing | any, not fast |
| 4 | `plot` | one or more values | nothing | any, not fast |
| 5 | `control` | label text, variable, optional range | nothing | top |
| 6 | `ask` | text; `timeout:` time | `bool?` | task |
| 7-9 | `log.info` `log.warn` `log.error` | text | nothing | any, not fast |
| 10 | `secret` | name text | `text` or `credentials` | any (11.1) |
| 11-13 | `board.name` `board.chip` `board.id` | | text | any; `board.chip` is the catalog chip family, `board.id` is the chip's unique ID on ESP32, ESP8266, RP2040-family, STM32F103, SAMD21G18, nRF52840 and Renesas RA4M1 boards, `"unknown"` where the target has no implemented ID backend |
| 14 | `board.uptime` | | time | any |
| 15 | `board.free_memory` | | memory | all catalog boards; heap/free-headroom measurement, not an exact fragmentation report |
| | `board.cores` `board.dma` `board.timers` `board.rtos` `board.ethernet` | | int, bool, int, text, bool | any; compile-time board facts (a DMA API and RTOS scheduling are not built; `on core N`, `mujard.timer` and `mujard.ethernet` cover the rest) |
| 16 | `board.voltage` | | `voltage?` | AVR via the bandgap reference, Pico/Pico W via VSYS/3, Nano 33 BLE via nRF52840 SAADC VDD, Zero via SAMD21 scaled IOVCC, Blue Pill via STM32F103 VREFINT, Uno R4 via RA4M1 AVCC; `none` elsewhere |
| 17 | `board.restart()` | | does not return | any; writes saved vars first |
| 18 | `board.reset_reason` | | ResetReason | ESP32/ESP8266, Optiboot AVR boards (Uno/Nano), Blue Pill (STM32F103), Zero (SAMD21G18), Nano 33 BLE (nRF52840), Uno R4 (RA4M1) and PC tests; **C320** elsewhere |
| | `board.halt(reason)`, `board.overruns`, `board.dropped_events` | text; none | does not return; int | any (not among the numbered 100) |
| 19 | `now()` | | time | any |
| 20 | `since(t)` | time | time | any |
| 21 | `stopwatch()` | | Stopwatch: `start() stop() reset() elapsed running` | variable declaration |
| 22 | `clock.time` | | `time of day?` (`hour` `minute` `second`) | any; `none` until set; Wi-Fi boards |
| 23 | `clock.date` | | `date?` (`year` `month` `day` `weekday`, 1 = Monday) | any; `none` until set; Wi-Fi boards |
| 24 | `clock.sync(timezone:)` | text: an IANA name ("Africa/Lagos") or a POSIX rule, fixed in the code (C326); none = UTC | nothing | any; Wi-Fi boards (NTP); `clock.source(rtc)` and `at HH:MM` without Wi-Fi come later |
| 25-27 | `abs` `min` `max` | one value / two or more of one dimension | same | any |
| 28 | `clamp` | value, range | same | any |
| 29 | `scale` | value, two ranges | dimension of the second range | any; does not clamp |
| 30 | `round` | number or quantity; optional int | same | any |
| 31-32 | `math.sqrt` `math.sin` ... | float (angles for trig) | float (`atan2` gives an angle) | any; `use mujard.math` |
| 33 | `smooth` ◆ | value, int | same as value | any |
| 34 | `between` | value, range | bool | any |
| 35 | `changed` ◆ | value | bool | any |
| 36-37 | `rising` ◆ `falling` ◆ | bool | bool | any |
| 38 | `hysteresis` ◆ | value; `on:` `off:` values | bool | any |
| 39 | `peak` ◆ | value, constant time | same as value | any |
| 40 | `random` | range, or the word `bool` | same type | any; `bool` is a special form |
| 41 | `random.seed` | optional int | nothing | any |
| 42 | `random.pick` | list | item | any |
| 43-46 | `on()` `off()` `toggle()` `write(bool)` | | nothing | any; fast-ok |
| 47 | `read()` | | bool | any; fast-ok |
| 48 | `blink` | frequency; `for:` time, `on:` time | nothing | any |
| 49 | `pattern` | text; `step:` time, `repeat:` int | nothing | any |
| 50 | `measure_pulse` | `level:` Level, `timeout:` time | `time?` | any; blocks up to the timeout. `pin.pulse(time)` gives nothing |
| 51-53 | `value` `volts` `percent` | | int, voltage, ratio | any |
| 54 | `write(level)` | ratio | nothing | any; fast-ok |
| 55 | `fade` | `to:` ratio, `over:` time | nothing | any |
| 56 | `frequency` | frequency | nothing | any |
| 57 | `reference` | voltage | nothing | any |
| 58-63 | events `pressed` `released` `held(time)` `clicked` `double_clicked` `changed`, also `rising` `falling` | | event | `on` header; `fast` on `changed`, `rising`, `falling`, `pressed`, `released` (6.5) |
| 64 | `debounce` | time | nothing | any |
| 65 | `tone` | frequency; optional time | nothing | any |
| 66 | `stop()` | | nothing | any |
| 67 | `melody` | text; `tempo:` int | nothing | any |
| 68 | `saved var` | | | top |
| 69 | `store.save` | name text, value | nothing | any |
| 70 | `store.load` | name text; `default:` value | value, or `T?` without a default | any |
| 71 | `serial` | baud | nothing | any (command) |
| 72 | `on serial.line` / `serial.read_line()` | | text / `text?` | event / task |
| 73 | `i2c.scan()` | | `list[int, 16]` | any; `use mujard.i2c` |
| 74 | `i2c.write` / `i2c.read` | int, `list[byte]` / int, int | nothing / `list[byte, 32]?` | any |
| 75 | `spi.device(...)`, then `.transfer(data)` | `cs:` pin, `speed:` frequency, `mode:` int; `list[byte]` | device; `list[byte, 32]` | top; any |
| 76 | `shift.out_register(...)`, then `.write(byte)` | `data:` `clock:` `latch:` pins | device | top; any |
| | `Uart(rx:, tx:, speed:)`: `write(text)`, `read_line()`, `on x.line` | pins, baud | device; nothing; `text?`; event | top; any; task; event |
| 77 | `wifi.connect` | credentials: `secret("name")` written right there (11.1) | nothing | any; works in the background, reconnects by itself |
| 78 | `wifi.connected`, events `wifi.joined` `wifi.lost` | | bool | any |
| 79 | `http.get` / `http.post` | url text; data | Reply (`ok status text json`) | task; waits up to 5 s and pauses the program meanwhile; `http://` and `https://` - https runs only with the certificate checked (the ESP32 core's CA bundle; built-in roots on ESP8266/RP2040) and gives `status -1` where the check can't run |
| 80-81 | `mqtt.connect` `mqtt.publish` `mqtt.subscribe`, event `mqtt.message as m` (`m.topic`, `m.text`) | host, `user:`, `password:`; topic, value; topic | nothing | any; `use mujard.mqtt`; MQTT 3.1.1, QoS 0, port 1883; never waits: connects in the background, keeps 4 publishes while down, re-subscribes after a reconnect |
| 82 | `power.sleep` | time; `wake_on:` event | nothing | any; pauses the whole program (6.10) |
| 83 | `power.deep_sleep` | time | does not return | any |
| 84-85 | `battery_on(analog pin, divider:, type:)`: `percent`, `volts` | pin, number, BatteryType | device; ratio, voltage | top; any |
| 86 | `expect` | condition; optional message; `halt:` bool | nothing | statement |
| 87 | `limit` | device, range | | top |
| 88 | `safe_state:` | | | top, one per program |
| 89 | `watchdog` | time | nothing | any |
| 101 | `Can(tx, rx, rate)` | tx, rx: pin; rate: baud (optional, 500000) | `can` object | `use mujard.can`; ESP32 family (TWAI), C320 elsewhere |
| 102 | `bus.send` | id: int; data: `list[int]` | bool | queues a frame (up to 8 bytes); false when the queue is full |
| 103 | `bus.read` | | `can_message?` | the next received frame or none (`id`, `data`, `rtr`) |
| 104 | `on bus.message` | | event | runs for each received frame |
| 105 | `ble.start` / `ble.send` | name: text; text | nothing; bool | `use mujard.ble`; advertises, notifies the connected central |
| 106 | `ble.read_line` / `ble.connected` | | `text?`; bool | the next received line or none; whether a central is connected |
| 107 | `on ble.line` | | event | runs for each received line |
| 108 | `ethernet.start` / `ethernet.connected` / `ethernet.ip` | | nothing; bool; text | `use mujard.ethernet`; classic ESP32's EMAC |
| 109 | `on ethernet.joined` / `on ethernet.lost` | | event | wire came up / went down |
| 110 | `usb.type` / `usb.key` | text; text | bool; bool | `use mujard.usb`; type / tap a key on the host |
| 111 | `usb.mouse_move` / `usb.click` / `usb.connected` | x, y: int | bool; bool; bool | move the pointer; click; HID device is up |
| 112 | `every ... on core N` | | block | run the block on a secondary core (every/after/when) |
| 113 | `t.every` / `t.freq` / `t.once` / `t.start` / `t.stop` | period; hz; period | nothing | `use mujard.timer`; arm a hardware timer |
| 114 | `on t.tick` / `t.elapsed` | | event; time | each timer fire; µs running |
| 90-93 | `pretend` `advance` `notified` `printed` | value; time; text; text | nothing; nothing; bool; bool | test |
| 94-97 | `length` `contains` `starts_with` `ends_with` `split` `to_number` | text; text; text; text; text; | int; bool; bool; bool; `list[text, 8]`; `number?` | any |
| 98-100 | `bit` `set_bit` `clear_bit` `toggle_bit` `bytes` | int | bool; int; int; int; `list[byte, 4]` | any |