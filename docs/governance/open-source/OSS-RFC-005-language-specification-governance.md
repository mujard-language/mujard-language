# OSS-RFC-005 Language Specification Governance

- Policy ID: OSS-RFC-005
- Classification: Internal
- Status: Draft pending approval
- Owner: Open Source Governance
- Version: 1.0

## 1. Purpose

This policy defines how the Mujard language changes: how proposals are made,
discussed and decided, and the rule that keeps the specification, parser and
goldens from drifting apart.

It implements the "three homes" rule recorded in the project notes: a language
change has three homes — the parser (`tools/mujard/lang`), the normative EBNF
in `docs/lang/MUJARD_LANGUAGE_SPEC.md` section 3 (the grammar tests parse it
directly), and the goldens (`tools/mujard/lang/testdata/cpp` and
`testdata/trees`). All three must change together or the test suite fails.

## 2. What counts as a language change

A change to the language is anything that alters what programs mean or what
the compiler accepts or rejects:

- grammar, tokens, keywords, indentation rules
- types, units, `none` rules, built-in signatures (spec Appendix A)
- event/task/scheduler semantics, pin and board rules
- error codes and their conditions
- the `.muj`/`.mjd` file format rules

Not language changes (normal review, no RFC): CLI commands, board manifests,
library descriptions for existing syntax, documentation typos, performance
work that preserves semantics.

## 3. Proposal process

Until a public remote exists, RFCs are submitted as patches or documents to
the maintainer at [CONTACT EMAIL]; afterwards they are filed in the
repository's RFC area. Either way the proposal is a document with:

- Title and RFC number
- Motivation: the problem, with a concrete program that cannot be written
  today
- Proposed change: the spec text delta (EBNF or prose), new error codes,
  affected built-ins
- Compatibility: does every program that compiled before still compile? If
  not, it is a breaking change (section 5)
- Boards affected: which families in `toolchain/boards` gain or lose
  capability
- Test plan: new `examples/ok` and `examples/bad` samples, golden updates,
  conformance expectations
- Alternatives considered

A proposal is discussion, not commitment. The maintainer may ask for
revisions; acceptance means "specify and implement", not "already decided
forever".

## 4. Discussion and decision

- Discussion happens in the proposal thread (mailing list or issue tracker
  once the public remote exists).
- The maintainer decides accept, reject or defer (OSS-GOV-001 section 5.1).
  Defer is a real answer — the proposal stays open for the next edition.
- Accepted proposals are recorded in the specification's revision sections
  (the spec already keeps revisions, for example section 15) and, when a
  second language edition exists, in the edition's changelog.
- Speculative or one-off requests ("just for my project") are declined unless
  they generalize; the language stays small on purpose.

## 5. Versioning and breaking changes

- The specification names its language edition (currently "Mujard 2026",
  spec section 14). Programs do not state an edition until a second one
  exists.
- Within an edition, changes are backward compatible: programs that compile
  under the edition keep compiling. The compatibility suite
  (`tools/mujard/lang/testdata/compat/2026`) enforces this — a frozen set of
  programs must keep compiling, warnings allowed.
- A breaking change requires a new language edition. The edition's syntax for
  selecting it will be added when the second edition exists (spec section
  14); until then breaking changes are not accepted.
- Deprecations inside an edition are announced in the spec revision notes and
  the release changelog (OSS-REL-006) before the release that removes them,
  giving users at least one release of notice.
- The compiler's public version follows semantic versioning (spec section
  14); the internal language-version integer in
  `tools/mujard/lang/version.json` bumps automatically whenever the language
  sources change and is the number every copy of Mujard reports.

## 6. Implementation requirements

An accepted language change is not done until:

- Parser, checker and generator updated (`tools/mujard/lang`).
- Spec section 3 EBNF (or the relevant section) updated — the grammar tests
  parse it, so a stale spec fails the suite.
- Goldens regenerated and the diff reviewed:
  `go test ./lang -run 'TestGenerateSamples|TestGoldenTrees' -update`.
- `editor-info.json`, `boards.json` and capabilities evidence refreshed via
  the language-update script where user-visible.
- Docs (`docs/lang`) synced to the platform's served copy — never edit
  that generated copy.
- Learn examples or reference pages updated when the change is user-facing.
- `go test ./lang` and `go test ./...` green.

## 7. Relationship to the roadmap

Language editions track the implementation plan's milestones (M1 to M4 in
`docs/lang/MUJARD_IMPLEMENTATION_PLAN.md`). The specification's revision
history is the authoritative record of what changed and when; RFCs are the
proposals that led there.
