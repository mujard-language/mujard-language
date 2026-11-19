# OSS-TM-004 Trademark and Branding Policy

- Policy ID: OSS-TM-004
- Classification: Internal
- Status: Draft pending approval
- Owner: Open Source Governance
- Version: 1.0

## 1. Purpose and status warning

This policy governs use of the "Mujard" name, the tagline and any logo.

Status note: the name was chosen for the language (after its author, per the
specification's Phase 0 decisions). A web search on 2026-10-19 found no product
of that name, but a formal trademark search in the target countries is still
open (tracked in `docs/lang/MUJARD_IMPLEMENTATION_PLAN.md`). Until registration
or clearance is confirmed, treat every mark below as claimed-but-unverified and
keep this policy's protections conservative.

## 2. The marks

- "Mujard" — the project and language name.
- "Bare code. Every board." — the tagline.
- The Mujard logo and app icons — to be produced; until then the wordmark is
  the mark.
- `.muj` / `.mjd` — community-recognized file extensions. Extensions are not
  marks but are covered by the same courtesy rules.

## 3. Permitted uses (no permission needed)

- Describing compatibility accurately: "this program is written in the Mujard
  language", "Mujard-compatible firmware", "runs on boards the Mujard
  toolchain supports".
- Linking to the official project and documentation.
- Academic, journalistic and educational commentary, including criticism.
- Using the name to identify the unmodified upstream in a fork's README,
  provided the fork is clearly marked unofficial (section 4).
- Package and file names inside a Mujard project (`.muj` files,
  `mujard_*` runtime headers) as produced by the toolchain itself.

## 4. Forks and derivatives

- A fork of the code must either rename itself or prominently mark itself
  "unofficial fork", "not affiliated with or endorsed by the Mujard project",
  with a link to upstream.
- A fork may state compatibility ("compatible with Mujard language version N")
  — that is nominative use, allowed.
- A fork must not use the Mujard logo, app icons, or trade dress, and must not
  use "Mujard" as its product name.
- Registries, app stores and domains for forks must not be confusingly similar
  (for example "MujardPro" or "MujardIDE" sold as official).

## 5. Prohibited uses

- Implying endorsement, certification or official status: "official Mujard",
  "certified by Mujard", "Mujard approved" for anything not produced by the
  project.
- Using the marks on merchandise, services or products without written
  permission.
- Domain squatting or registering mujard-style names to resell or impersonate.
- Using the name in a way that suggests the platform's proprietary services
  (accounts, hosted builds, the apps) are included when they are not.
- Removing or altering notices that attribute the work to Mujard.

## 6. Relation to Apache-2.0

Apache-2.0 section 6 grants no trademark rights. The license covers the code;
this policy covers the name. A licensee may exercise all Apache-2.0 rights —
copy, modify, sublicense, sell — while still obeying sections 3-5 on naming.

## 7. Enforcement

- First contact for good-faith misuse is a request to fix, not a lawsuit.
- Enforcement is handled by the maintainer today and by the legal entity
  holding the mark once formed (OSS-GOV-001 section 5.7).
- Escalation path: polite request, then formal notice, then takedown requests
  to hosts and registries, then legal action for willful infringement or
  passing-off.
- The project will not use trademark claims to restrict the Apache-2.0 rights
  themselves; enforcement targets confusion and impersonation, not use of the
  code.

## 8. Open items

- Complete the formal trademark search in the target countries before the
  public launch (tracked in the implementation plan's risk table).
- Decide the holding entity for the mark when the legal entity is formed.
- Produce official logo and app-icon assets and record their usage rules here.
