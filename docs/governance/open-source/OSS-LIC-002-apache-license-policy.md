# OSS-LIC-002 Apache License Policy

- Policy ID: OSS-LIC-002
- Classification: Internal
- Status: Draft pending approval
- Owner: Open Source Governance
- Version: 1.0

## 1. Purpose

This policy states which Mujard components are open source, under which
license, and how licenses apply to contributions, generated code, notices and
third-party dependencies.

It implements the dependency rules already recorded in
`docs/OPEN_SOURCE_ADOPTION_PLAN.md`.

## 2. License selection

The open-source components are licensed under the Apache License, Version 2.0
(Apache-2.0). The `LICENSE` file at the repository root carries the full text.

Why Apache-2.0:

- Permissive: commercial use, modification and redistribution are allowed,
  including inside closed products.
- Explicit patent grant from each contributor, with a retaliation clause —
  important for a compiler and runtime that generate code for third-party
  cores.
- NOTICE-file mechanism gives a place for attribution that survives
  redistribution (section 6).
- Compatible with the project's dependency rule: MIT/BSD/Apache-2.0 preferred;
  GPL-family code is never copied into Mujard sources (OSS-SBOM-007).

Apache-2.0 was chosen over GPL-family licenses so that companies, schools and
OEMs can embed the generated runtime and toolchain in products without
copyleft reaching their own code.

## 3. What the license covers

Apache-2.0 applies to:

- The Mujard language specification and companion docs
  (`docs/lang/MUJARD_LANGUAGE_SPEC.md`, `docs/lang/README.md`, reference and
  learn examples) — see section 8 for the documentation note.
- The compiler front end and C++ generator (`tools/mujard/lang`).
- The C++ runtime and feature headers (`tools/mujard/lang/runtime`).
- The `.mlib` library descriptions (`tools/mujard/lang/descs`).
- The `mujard` command line and language server (`tools/mujard`, public build
  `MujardCLI.exe`); operator-only commands compiled behind `-tags operator` are
  part of the same source and license but are never distributed.
- Board manifests, package descriptors and the toolchain index (`toolchain/`).
- The learn examples and testdata under `docs/lang/examples` and
  `tools/mujard/lang/testdata`.
- Build and release tooling that ships or supports the above
  (`tools/mujard-build` and the release tooling).

## 4. What the license does not cover

Apache-2.0 does not apply to the proprietary platform (OSS-GOV-001 section 3):
PHP backend, developer API implementation details, hosted web app, Android
app, desktop Studio app, hosted AI, cloud services. Those components carry
their own terms and remain closed.

A file or folder not listed in section 3 is proprietary unless a LICENSE or
header says otherwise. When in doubt, ask before assuming a component is open.

## 5. Contributions

- Every contribution is licensed under Apache-2.0 (OSS-CON-003 for the DCO
  process).
- Contributors retain their own copyright; the project does not require
  copyright assignment. Apache-2.0 section 5 means a submitted contribution is
  licensed under the same terms as the Work.
- The project copyright line is "Copyright 2026 Mujard contributors". A
  contributor may add their own copyright line to files they substantially
  create; individual lines are added alongside, never replacing the project
  line.

## 6. NOTICE file

Apache-2.0 section 4(d) requires redistributors to carry a NOTICE file's
attribution notices when one exists. Policy:

- A NOTICE file at the repository root will list third-party components whose
  notices must travel with copies of the Work, starting from the pinned
  toolchain entries in `toolchain/index.json` that carry license obligations.
- Until the first NOTICE file is added, no section-4(d) obligation exists; the
  absence is recorded here so nobody assumes one is already required.
- When the NOTICE file is created it is regenerated alongside
  `docs/THIRD_PARTY_LICENSES.md` and reviewed at each release (OSS-REL-006).
- Generated programs (section 7) embed the runtime; the generated output must
  carry a pointer to the runtime's license so firmware authors can comply.

## 7. Generated code

- Programs written in Mujard (`.muj`/`.mjd`) and the firmware built from them
  belong to their authors. The Apache-2.0 license does not reach into user
  programs; it only covers the Mujard sources.
- The C++ the compiler generates embeds `mujard_rt` runtime code. That embedded
  portion remains Apache-2.0 licensed: a user may compile, modify and
  distribute firmware containing it, including commercially, under Apache-2.0
  terms for the runtime portion. The user's own program logic is theirs.
- Generated artifacts must therefore carry an attribution notice naming the
  Mujard runtime and pointing to the LICENSE; the generator is required to add
  this. It is a compliance feature, not a restriction on the user's program.
- Patents: the Apache-2.0 patent grant covers the Work (compiler, runtime,
  CLI). It does not extend to a user's program or to the third-party Arduino
  cores the generated code links against; those cores carry their own licenses
  (for example LGPL-2.1 Arduino cores, GPL toolchains — downloaded, not
  redistributed, per OSS-SBOM-007).

## 8. Documentation licensing

Code, specification and toolchain data are Apache-2.0. Prose documentation
(README, reference pages, learn-example explanations) is proposed under
CC BY 4.0 so it can be quoted and taught from freely — pending legal review
before that change is made. Until then all docs are Apache-2.0 like the rest
of the repository. The specification itself stays Apache-2.0 because the
grammar tests parse it as a build input (the "three homes" rule,
OSS-RFC-005); treating it as code-adjacent avoids a split license on a file
the toolchain consumes.

## 9. License headers

- New Go source files carry `// SPDX-License-Identifier: Apache-2.0` near the
  top.
- New C/C++ runtime files carry the same SPDX identifier in the file's comment
  style.
- Data files (JSON manifests, `.mlib`, `.muj` examples) do not carry headers;
  the LICENSE at the root covers them.
- Existing headers are never removed or reworded to weaken attribution
  (consistent with rule 4 of the adoption plan).
- Files copied from third-party projects keep their original headers and are
  recorded in `docs/THIRD_PARTY_LICENSES.md` and, once it exists, NOTICE.

## 10. Trademark boundary

Apache-2.0 section 6 grants no trademark rights. The "Mujard" name and any logo
are governed by OSS-TM-004; a fork may use the code freely but must follow the
naming rules there.
