# OSS-REL-006 Compiler Release Governance Policy

- Policy ID: OSS-REL-006
- Classification: Internal
- Status: Draft pending approval
- Owner: Open Source Governance
- Version: 1.0

## 1. Purpose

This policy governs how the open-source toolchain is versioned and released:
the compiler, the `mujard` CLI, the runtime, board definitions and the
language specification.

It operationalizes `docs/RELEASE_CHECKLIST.md`, which remains the
authoritative step-by-step for shipping a language change to the four places
users meet it (build server, website, desktop app, phone app).

## 2. Versioning

- Public releases of the compiler, CLI and toolchain follow semantic
  versioning (spec section 14 commits the compiler to semver). MAJOR =
  breaking language edition or CLI contract change; MINOR =
  backward-compatible language or toolchain additions; PATCH = fixes.
- The language version integer in `tools/mujard/lang/version.json` is the
  compatibility signal every copy of Mujard reports (CLI, in-editor checker,
  phone app, server compiler). It is fingerprinted from the language sources
  and bumps automatically when they change; releases record it alongside the
  semver tag.
- The specification carries its edition name ("Mujard 2026") plus document
  revisions in its revision sections.
- Board manifests and package descriptors are versioned by their own schema
  fields; the toolchain index (`toolchain/index.json`) pins every download by
  version and sha256.
- The operator build (`go build -tags operator`) is never released; it has no
  version and is never distributed (OSS-CON-003 section 9,
  `tools/mujard/README.md`).

## 3. Release authority

- The maintainer approves every release (OSS-GOV-001). Automation prepares;
  it does not approve.
- A release is cut only from a green tree: `go test ./...` (and `go test -tags
  operator ./...` for operator-side verification), `go vet`, `gofmt` clean,
  plus the checks in section 4.
- Language releases follow `docs/RELEASE_CHECKLIST.md` order: server first,
  then website, desktop, phone — an older server cannot build programs that
  use what it does not know.

## 4. Release checklist (open-source artifacts)

For each public release, before tagging:

- Language change complete; spec updated (OSS-RFC-005 three homes).
- The language-update script run without skipping tests, ending "All done ...
  at language version N". This regenerates `boards.json`, `editor-info.json`,
  goldens and `version.json`, runs the tests, and rebuilds the web wasm bundle,
  Studio, the desktop app and the server compiler.
- Goldens regenerated where the language changed and the diff reviewed.
- The docs sync run so the platform's served docs copy matches `docs/lang`;
  the sync check is clean.
- Build-matrix evidence refreshed where board capability claims changed:
  `MUJARD_EVIDENCE_OUT=lang/capabilities.json go test ./lang -run
  TestBuildMatrix` — `capabilities.json` is the evidence file the language
  fingerprint reads.
- `docs/THIRD_PARTY_LICENSES.md` regenerated if dependencies changed; NOTICE
  file updated if present (OSS-LIC-002 section 6).
- SBOM produced for the release (OSS-SBOM-007).
- Changelog entry written (section 6).
- Checksums computed for every published artifact (section 7).

## 5. Release artifacts

- Source: the repository (the public remote).
- Binary: `MujardCLI.exe` (public build; `mujard.exe` copy included), Linux
  build (`GOOS=linux go build`), and the compiler `mujard` binary used
  server-side.
- Never shipped: the operator build, signing keys, private config and
  private hosting exports.
- The wasm build is a build artifact of the website, rebuilt by the
  language-update script, not a separately versioned release.

## 6. Changelog

- Every public release has a changelog entry: version, language version N,
  date, summary, breaking changes, deprecations, security fixes (credited per
  OSS-SEC-008 where the reporter consents).
- The changelog lives with the releases (repository CHANGELOG or release
  notes) and is written at release time, not reconstructed later.

## 7. Integrity: checksums and signing

- Every published binary artifact ships with a SHA-256 checksum file. This is
  the minimum today.
- Signing (minisign/GPG or Sigstore) is a planned upgrade once keys and
  infrastructure exist; until then checksums plus the DCO-signed source
  history are the integrity story, and this policy is revised when signing
  lands.
- Toolchain downloads are already integrity-pinned: `toolchain/index.json`
  records sha256 for every package and `mujard-build` verifies before install
  (OSS-SBOM-007).

## 8. Hotfixes and security releases

- Security fixes follow OSS-SEC-008: private report, fix, coordinated
  disclosure, then a PATCH (or MINOR if it changes behavior) release with a
  changelog credit.
- A hotfix that cannot wait for the full checklist still runs section 3's
  green-tree requirement; skipped documentation steps are completed before the
  next regular release and noted in the changelog.
