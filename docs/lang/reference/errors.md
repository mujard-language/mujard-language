# Error and warning catalogue

Every diagnostic the compiler can report, from its message table.

## Lexer (L1xx)

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

## Parser (P2xx)

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

## Compile (C3xx)

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

## Board and pins (B4xx)

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

## Units (U5xx)

| Code | Message |
|---|---|
| U501 | {value} needs a unit here. Did you mean {value} {suggestion}? |
| U502 | Can't combine {a} and {b}: they measure different things. |
| U503 | Expected a {want} (like {example}), but this is a {got}. |
| U504 | '{u}' isn't a unit Mujard knows. |

## Warnings (W6xx)

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

