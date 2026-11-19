# Release policy

The canonical policy is
`docs/governance/open-source/OSS-REL-006-compiler-release-governance-policy.md`
(the step-by-step checklist lives in the platform repository).

Summary:

- The maintainer approves every release; automation (the language-update
  script) prepares artifacts but does not approve them.
- A release is cut only from a green tree: `go test ./...`, `go vet`, `gofmt`
  clean, goldens regenerated where the language changed, docs synced,
  capability evidence refreshed.
- Language releases ship in checklist order: server compiler first, then website,
  desktop, phone — an older server cannot build programs that use what it does
  not know.
- Every published binary ships with a SHA-256 checksum; artifact signing is a
  planned upgrade (OSS-REL-006 §7). Toolchain downloads are already sha256-pinned.
- Security fixes follow `SECURITY.md` / OSS-SEC-008: private report, fix,
  coordinated disclosure, PATCH release with changelog credit.
- The operator build (`-tags operator`) is never released.
