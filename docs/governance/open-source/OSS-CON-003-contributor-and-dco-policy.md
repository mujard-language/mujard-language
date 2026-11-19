# OSS-CON-003 Contributor and DCO Policy

- Policy ID: OSS-CON-003
- Classification: Internal
- Status: Draft pending approval
- Owner: Open Source Governance
- Version: 1.0

## 1. Purpose

This policy explains how to contribute to the open-source components, what
sign-off is required, how contributions are reviewed, and what will be
rejected.

## 2. Ground rules before you start

- Read the license policy (OSS-LIC-002): your contribution will be Apache-2.0.
- Read the governance policy (OSS-GOV-001): the maintainer holds final
  decision authority today.
- For language changes, read OSS-RFC-005 first — a proposal comes before code.
- Follow `CODE_OF_CONDUCT.md` at the repository root.

## 3. How contributions arrive today

The repository is local-only with no public remote yet. Until a remote exists:

- Prepare your change as a patch: `git format-patch` for commit-based work, or
  `git diff` output for uncommitted work.
- Send the patch to the maintainer at [CONTACT EMAIL] with a short description.
- Every commit message in the patch must end with a Signed-off-by line
  (section 4).
- Do not include secrets, keys or build output — `.gitignore` covers the
  usual spots; review your diff before pushing.
- The repository is shared by several working sessions; keep patches scoped to
  your change so they apply cleanly.

When a public remote exists this section is replaced by the normal
fork-and-pull-request flow; the DCO and review rules do not change.

## 4. Developer Certificate of Origin (DCO)

Every contribution must be signed off under the Developer Certificate of
Origin, Version 1.1. Sign-off certifies the contribution is yours (or you have
the right to submit it) and may be relicensed under Apache-2.0.

Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.

Everyone is permitted to copy and distribute verbatim copies of this license
document, but changing it is not allowed.

Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I have the
    right to submit it under the open source license indicated in the file; or

(b) The contribution is based upon previous work that, to the best of my
    knowledge, is covered under an appropriate open source license and I have
    the right under that license to submit that work with modifications,
    whether created in whole or in part by me, under the same open source
    license (unless I am permitted to submit under a different license), as
    indicated in the file; or

(c) The contribution was provided directly to me by some other person who
    certified (a), (b) or (c) and I have not modified it.

(d) I understand and agree that this project and the contribution are public
    and that a record of the contribution (including all personal information
    I submit with it, including my sign-off) is maintained indefinitely and
    may be redistributed consistent with this project or the open source
    license(s) involved.

How to sign off:

- `git commit -s` adds `Signed-off-by: Your Name <you@example.com>`
  automatically.
- In a patch without git, add the same line at the end of the commit message.
- The name and email must be real enough to identify you; they become part of
  the permanent public record (DCO clause d).

## 5. CLA decision: DCO, not CLA

The project uses the DCO rather than a Contributor License Agreement. Reasons:

- Lighter: a sign-off line versus a legal document to sign per contribution or
  per company.
- Standard: the DCO is what the Linux kernel and many foundations use;
  companies' legal teams already know it.
- Copyright stays with contributors, matching Apache-2.0's model
  (OSS-LIC-002 section 5).

A CLA may be reconsidered if a major corporate contributor requires one; that
is a governance decision (OSS-GOV-001).

## 6. What a good contribution looks like

- Scoped: one change, one purpose. Language changes go through OSS-RFC-005
  first.
- Tested: see section 7.
- Documented: spec, README, reference and `editor-info.json` updated where the
  change is user-visible.
- Idiomatic: Go code is gofmt-clean and passes `go vet`; C++ runtime changes
  keep the existing style; JSON manifests match their schema.

## 7. Tests required

Minimum bar before a patch is reviewed:

- `cd tools/mujard; go test ./...` passes (use a standard Go toolchain if it
  is not on PATH).
- Language changes (`tools/mujard/lang`): `go test ./lang` passes, and the
  "three homes" are updated together — parser, normative EBNF in
  `docs/lang/MUJARD_LANGUAGE_SPEC.md` section 3, and goldens. Regenerate
  goldens with `go test ./lang -run 'TestGenerateSamples|TestGoldenTrees'
  -update` and review the diff before submitting.
- CLI changes: `go test ./...` plus the operator-tier proof (`go test -tags
  operator ./...`) must keep passing — operator commands must stay compiled
  out of the public build.
- Build-tool changes (`tools/mujard-build`): `go test .` passes; new board
  families add a manifest and sha256-pinned index entries only.
- Docs-only changes: the docs-sync check reports clean after `docs/lang`
  changes reach the platform's served copy (that copy is generated — never
  edit it).

## 8. Code review

- The maintainer reviews every patch today; there is no required second
  reviewer yet.
- Review focuses on correctness against the specification, test coverage, the
  three-homes rule for language changes, and license clarity.
- A patch may be accepted with requested changes, deferred, or rejected with a
  reason. Rejection reasons are recorded so the same argument is not repeated.

## 9. What gets rejected

- Code copied from sources without a compatible license, or with copyright
  headers stripped (adoption plan rule 4).
- Anything that weakens the operator-command boundary: no path that lets a
  user reach `admin`, `status`, `doctor` or `loadtest` in the public build, and
  no environment variable that is the only guard.
- Secrets, keys, credentials or `.gitignore`d content in any form.
- Language changes merged without the spec EBNF and goldens (the grammar
  tests enforce this — a parser/spec mismatch fails the suite).
- Trademark misuse (OSS-TM-004): patches that rename the project or imply
  official status for forks.
- Changes that break the four-way release contract (OSS-REL-006): a language
  change that cannot ship server-first with the documented checklist is not
  mergeable as-is.

## 10. Intellectual property warranty

By submitting a patch with a Signed-off-by line, the contributor certifies the
DCO terms: the work is theirs or properly licensed, and it is contributed
under Apache-2.0. The project relies on this warranty; a contribution later
found to be misappropriated will be removed and the contributor's trust
revoked. Willful violations may have legal consequences — this is stated
plainly so nobody signs carelessly.

## 11. Copyright in the record

Contributions are recorded with the sign-off name and email. Copyright of
contributed files remains with their authors; the project's NOTICE/copyright
lines aggregate attribution (OSS-LIC-002 section 5).
