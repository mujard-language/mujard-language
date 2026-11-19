# Security Policy

## Reporting a vulnerability

Report security issues privately to **mujardlang@gmail.com**.

Please include:

- The affected component: language/compiler, CLI, runtime, toolchain, or docs
- The version or commit you tested
- Reproduction steps or a proof of concept
- The impact you believe it has

Do not include other people's data, and do not access systems you do not own.

A `security.txt` file (RFC 9116) will be published with the public site and
repository launch, pointing at the same channel; until then this address is the
way in. PGP encryption is not set up yet — do not delay a report waiting for it;
ask and keys will be arranged if needed.

We aim to acknowledge reports within 5 business days and give a triage
assessment within 10. Critical fixes ship as patch releases. Disclosure is
coordinated with you, and you are credited in the release notes if you wish.

## Scope

In scope:

- The Mujard language implementation: lexer, parser, checker and C++ generator
  (`tools/mujard/lang`)
- The C++ runtime and library descriptions that generated programs use
- The `mujard` command line and language server, including the public
  `MujardCLI.exe` build
- `mujard-build`, board manifests, package descriptors and the pinned toolchain
  index (`toolchain/`)

Out of scope — handled by a separate process:

- The proprietary platform: Backend, developer API, hosted web app,
  Android app, desktop Studio app, hosted AI and cloud services. If your report
  concerns those, contact **mujardlang@gmail.com** and say it is a platform issue.
- Third-party dependencies themselves — report those upstream; we ship updated
  pins as patch releases once a fix exists.

## Safe harbor

We will not pursue legal action against anyone who reports a vulnerability in
good faith, makes a reasonable effort to avoid privacy violations, data
destruction and service disruption, does not exploit the issue beyond what is
needed to demonstrate it, and gives us reasonable time to fix before any public
disclosure.

Test only against your own programs, your own hardware, and the simulated board
(`mujard test`, `mujard repl`) — never other people's devices, accounts or the
hosted services beyond what your own account permits.

Good-faith security research is welcome and appreciated.
