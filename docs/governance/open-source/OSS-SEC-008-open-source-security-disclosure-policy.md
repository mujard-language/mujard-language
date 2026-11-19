# OSS-SEC-008 Open-Source Security Disclosure Policy

- Policy ID: OSS-SEC-008
- Classification: Internal
- Status: Draft pending approval
- Owner: Open Source Governance
- Version: 1.0

## 1. Purpose

This policy defines how security vulnerabilities in the open-source
components are reported, triaged, fixed and disclosed.

## 2. Scope

In scope (report here):

- The Mujard language: lexer, parser, checker, generator
  (`tools/mujard/lang`).
- The C++ runtime and standard library descriptions the generated programs
  use.
- The `mujard` CLI and language server, public build (`tools/mujard`).
- `mujard-build` and the board/package manifests and pinned toolchain index
  (`toolchain/`).
- The language specification and documentation where a documentation bug
  creates a security-relevant misunderstanding (rare; mention it).

Out of scope (different process):

- The proprietary platform: PHP backend, developer API, hosted web app,
  Android app, desktop Studio app, hosted AI, cloud services. Reports about
  those go through the platform's own channel — see `SECURITY.md` at the
  repository root; until it names a separate address, contact the maintainer
  at [CONTACT EMAIL] and say it concerns the platform.
- Vulnerabilities in third-party components we merely depend on: report
  upstream; we will ship an updated pin as a PATCH release once a fix exists
  (OSS-SBOM-007).

## 3. Reporting channel

- Email: [CONTACT EMAIL] — the private security address. Use it for anything
  you believe could harm users; do not open a public issue first.
- A `security.txt` file (RFC 9116) will be published with the public
  site/repository launch pointing at the same channel; until then this
  document and `SECURITY.md` are the reference.
- PGP encryption is not yet set up; if a report needs it, ask and keys will
  be arranged. Do not delay a report waiting for encryption.
- Include: affected component and version (or "current source"), reproduction
  steps or a proof of concept, impact assessment, and whether you want public
  credit.

## 4. What to expect

Targets (subject to approval of this draft; the maintainer confirms real dates
during triage):

- Acknowledgement of receipt: within 5 business days.
- Triage and severity assessment: within 10 business days of acknowledgement.
- Fix timeline: depends on severity — critical fixes target the next PATCH
  release; the maintainer communicates the expected date at triage.
- Disclosure: after the fix is released, or earlier if the reporter requests
  and the maintainer agrees. Coordinated disclosure is the default; we do not
  publish details before a fix exists unless the reporter chooses full public
  disclosure themselves.

## 5. Triage and severity

- Severity is judged by real-world impact on users' hardware, data and
  devices — a compiler bug that generates wrong firmware for a
  safety-relevant circuit is treated seriously even if it needs unusual
  input.
- The maintainer triages alone today; a security contact role is added when
  contributors take on review duties (OSS-GOV-001).
- Reports are tracked privately; nothing about an unpatched issue is public,
  including in commit messages ("fixed a security issue" without detail is
  fine and preferred pre-disclosure).

## 6. Credit

- Reporters are credited in the release's security notes and changelog with
  their consent (name, handle or anonymous).
- We do not pay bounties today; credit and a thank-you in the advisories is
  the current policy. If a bounty program is created it will be announced
  here.
- Credit is not contingent on the report being accepted as a vulnerability —
  good-faith reports that turn out to be non-issues are still thanked
  privately.

## 7. Safe harbor

The project will not pursue legal action against researchers who:

- Report in good faith through the channel above.
- Make a reasonable effort to avoid privacy violations, destruction of data,
  and interruption or degradation of service.
- Do not exploit a vulnerability beyond what is needed to demonstrate it.
- Give the project reasonable time to fix before any public disclosure.
- Stay within the scope in section 2.

Testing should use your own programs and hardware, or the simulated board
(`mujard test`, `mujard repl`) — never other people's devices, accounts or
the hosted services beyond what your own account permits. The platform's
services have their own rules of use.

## 8. Coordinated disclosure details

- When a fix ships, the disclosure note states: what was wrong, who reported
  it (with consent), which versions are affected, how to update.
- If a fix cannot ship quickly, the maintainer tells the reporter the revised
  timeline; reporters who disagree may disclose after the timeline lapses,
  and we will not treat good-faith disclosure as hostile.
- Embargoes longer than 90 days are granted only by explicit agreement.

## 9. Prohibited conduct

- Publicly disclosing an unpatched vulnerability.
- Accessing or modifying data that is not yours.
- Using a vulnerability to harm users, devices or the platform's services.
- Social-engineering project members for vulnerability details.

Violations of section 9 remove safe-harbor protection; everything else in
good faith is protected.
