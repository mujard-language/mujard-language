# OSS-GOV-001 Open-Source Governance Policy

- Policy ID: OSS-GOV-001
- Classification: Internal
- Status: Draft pending approval
- Owner: Open Source Governance
- Version: 1.0

## 1. Purpose and scope

This policy defines who governs the open-source parts of Mujard, how decisions
are made, who may release software, who answers security reports, and what
happens if the current maintainer steps down.

It applies to the open-source components listed in OSS-LIC-002: the Mujard
language specification, compiler, runtime libraries, library descriptions,
command-line tool, language server, board and toolchain definitions, and the
language documentation.

It does not apply to the proprietary platform described in section 3.

## 2. Principles

- The language must stay usable without an account: compiling a `.muj` program
  on your own machine requires no registration, key or service.
- Decisions that change the language are made openly and recorded
  (OSS-RFC-005).
- Security reports are handled privately and promptly (OSS-SEC-008).
- The "Mujard" name is held for the community, not sold with the code
  (OSS-TM-004).

## 3. What is open and what is not

Open source under Apache-2.0 (OSS-LIC-002):

- `tools/mujard` — the Go command line (`mujard`, public build `MujardCLI.exe`)
  and the language toolchain.
- `tools/mujard/lang` — lexer, parser, checker, C++ code generator.
- `tools/mujard/lang/runtime` — the C++ runtime (`mujard_rt`) and feature
  headers the generated programs link.
- `tools/mujard/lang/descs` — `.mlib` library descriptions.
- `toolchain/` — board manifests, package descriptors and the pinned download
  index.
- `docs/lang` — the specification, reference and learn examples.

Proprietary, unchanged:

- the PHP backend and developer API
- the hosted web app and hosted AI services
- the Android app
- the desktop Studio app and `MujardServer.exe`
- cloud services, accounts and billing

The boundary is deliberate: the language and its toolchain are open so anyone
can inspect, port and extend them; the hosted platform is the product that
funds the project. Open-sourcing the platform is a separate future decision,
not implied by this policy.

## 4. Governance structure

### 4.1 Current model: maintainer (BDFL)

The project is maintained by a single maintainer today. In a BDFL (benevolent
dictator for life) model the maintainer holds final decision authority on:

- language specification changes and editions
- releases of the compiler, CLI and toolchain
- trademark and brand decisions
- security response priorities
- who may commit to the repository

This matches the project's size: one codebase, one primary author, no
foundation. The model is recorded here so contributors know how decisions are
made rather than discovering it by surprise.

### 4.2 Contributors

Anyone may contribute under OSS-CON-003. Contributors earn trust through review
history; there are no elected roles yet. The maintainer may invite frequent
contributors to review privileges and, later, co-maintainer status.

### 4.3 Future governance

When the project outgrows one maintainer — measured by active contributors,
downstream forks needing coordination, or a trademark/legal entity — governance
should move toward a foundation or a contributor-elected council. The
succession section (5.6) is the trigger document for that conversation.

## 5. Decision-making

### 5.1 Language changes

Language changes follow OSS-RFC-005: a written proposal, public discussion, and
the "three homes" atomic update rule (parser, normative EBNF in the
specification, and goldens). The maintainer decides accept/reject;
controversial decisions should be deferred to a broader discussion before
deciding.

### 5.2 Toolchain and tooling changes

Changes that do not alter the language (CLI commands, board manifests, package
descriptors, build tooling) need maintainer approval through normal code
review. Board manifests and package descriptors are data; adding a board is
routine, changing how a family builds is a design discussion.

### 5.3 Releases

The maintainer approves every release (OSS-REL-006). Release authority cannot
be delegated to automation; the language-update script prepares artifacts, it does not
approve them.

### 5.4 Disagreements

When a contributor and the maintainer disagree, the path is: discussion in the
proposal or review thread, then the maintainer's decision stands for now, and
the contributor may fork under Apache-2.0 (subject to OSS-TM-004 naming rules).
A recorded dissent in the proposal is welcome so future maintainers can
revisit.

### 5.5 Documentation and policy

This governance set is versioned like code. Changes to these policies are
proposed the same way as other changes and approved by the maintainer; until a
public remote exists they travel as patches (OSS-CON-003).

### 5.6 Succession plan

The project must not depend on one person's continued availability:

- The repository, specification, goldens and release scripts already live
  together; a successor can rebuild everything from source plus the
  documented toolchain table.
- The maintainer will name a successor in this document (or in a private
  succession note) before the open-source release, and review the naming
  annually.
- If the maintainer becomes unavailable without naming a successor, the most
  active contributors should form an interim council, freeze breaking changes,
  and resolve governance by consensus; trademark and domain control move only
  through the legal entity holding them (section 6), never informally.
- Critical-person risk today: compiler knowledge is concentrated in the
  maintainer. Mitigation: this documentation set, the specification, the
  implementation plan (`docs/lang/MUJARD_IMPLEMENTATION_PLAN.md`) and the
  release checklist (`docs/RELEASE_CHECKLIST.md`) are the written record a
  successor works from.

### 5.7 Funding and resources

No funding is assumed. If donations or sponsorships are accepted in future,
funds are for project infrastructure (build machines, CI, domain, legal) and
are not distributed as revenue share. The trademark (OSS-TM-004) is held by
the legal entity or the maintainer personally — to be decided when the entity
is formed — precisely so money and name cannot be separated in a sale.

## 6. Relationship to the proprietary platform

- The open language is a standalone product: `.muj` programs compile to C++
  and then through the same board toolchains Arduino uses. Nothing in the
  language requires the platform.
- The platform adds hosted builds, the phone/desktop apps, AI assistance,
  accounts and payments. Those conveniences are proprietary and may change or
  be withdrawn without affecting the language.
- Features that only make sense hosted (accounts, cloud builds, AI) are out of
  scope for the specification unless they also have an offline meaning.
- API keys, hosted services and app-store distribution are governed by the
  platform's own terms, not by Apache-2.0.
- Contributors should not assume platform roadmap items; the language roadmap
  lives in the specification revisions and RFCs.
