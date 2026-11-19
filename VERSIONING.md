# Versioning

The canonical policy is
`docs/governance/open-source/OSS-REL-006-compiler-release-governance-policy.md`;
the language specification §14 is the normative source.

Summary:

- Public releases of the compiler, CLI and toolchain follow **semantic
  versioning**: MAJOR = breaking language edition or CLI contract change,
  MINOR = backward-compatible additions, PATCH = fixes.
- The language's internal compatibility integer lives in
  `tools/mujard/lang/version.json`; it is fingerprinted from the language sources
  and bumps automatically when they change — every copy of Mujard reports it.
- The specification carries its edition name ("Mujard 2026") plus revision
  history; within an edition changes are backward compatible (the frozen
  compatibility suite enforces this).
- Board manifests and package descriptors version by their own schema fields;
  `toolchain/index.json` pins every download by version + sha256.
