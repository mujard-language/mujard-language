# RFC process

The canonical process is
`docs/governance/open-source/OSS-RFC-005-language-specification-governance.md`.

Summary:

- Any change to what programs mean or what the compiler accepts is a **language
  change** and requires a written proposal first: motivation, spec delta,
  compatibility impact, boards affected, test plan, alternatives.
- The "three homes" rule: parser (`tools/mujard/lang`), normative EBNF in
  `docs/lang/MUJARD_LANGUAGE_SPEC.md` §3, and goldens — all change together or
  the test suite fails.
- Statuses: draft → under review → accepted / rejected / withdrawn → implemented.
- Accepted changes land in the spec's revision sections and the release
  changelog (`CHANGELOG.md`, `docs/governance/open-source/OSS-REL-006`).
