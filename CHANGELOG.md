# Changelog

All notable changes to the Mujard language, compiler and toolchain are documented here.
The format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/). The
language is pre-1.0; version numbers track the language version reported by
`mujard --version` and `version.json`.

## [Unreleased]

### Added

- Initial public import of the open-source components:
  - `docs/lang/` — the normative language specification, built-ins reference,
    per-topic reference pages, tutorials and compile-tested examples
  - `tools/mujard/` — the `mujard` command line: compiler front-end (`check`,
    `gen`, `fmt`, `test`, `repl`, `lsp`), language server, WASM build and the
    hosted-service client
  - `tools/mujard/lang/` — lexer, parser, checker and C++ generator, the
    `mujard_rt` runtime, `.mlib` library descriptions and capability evidence
  - `tools/mujard-build/` — the direct build backend for local firmware builds
  - `toolchain/` — board manifests, package descriptors and the pinned download
    index
- `HISTORY.md` — the project timeline from planning (January 2025) through the
  2026 implementation milestones
- `ROADMAP.md` — public milestones and capability evidence model
- `CONTRIBUTING.md` — the three-homes rule, DCO sign-off and build/test guide
