# Mujard sample programs

The test suite every compiler phase is checked against (`docs/lang/MUJARD_LANGUAGE_SPEC.md`, Phase 0 of the plan).

- `ok/` - 41 correct programs covering every keyword and about 65 of the 100 built-ins. Each starts with
  `# sample: <what it does>`. Every phase must accept all of them (the lexer tokenises them, the parser parses them, the
  checker finds no errors, and later the C++ builds for their board).
- `bad/` - 34 programs with one mistake each. The first line says what the compiler must report:
  `# expect: <CODE> <line>:<column>` (line 1 is that comment line; columns count from 1). The code's first letter says
  which phase reports it: `L` lexer, `P` parser, `C` checker, `U` units, `B` board and pins.

Changing a sample is a change to the language's tests: update the expected result in the same edit.
