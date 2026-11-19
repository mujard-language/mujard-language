# Contributing to Mujard

Thanks for your interest. Mujard is pre-1.0 and moving fast — small, focused changes
land best.

## Before you start

* Read the [specification](docs/lang/MUJARD_LANGUAGE_SPEC.md) — it is **normative**. If
  code and spec disagree, the spec wins and the code is wrong.
* Check [ROADMAP.md](ROADMAP.md) for where the project is headed, and
  [HISTORY.md](HISTORY.md) for how it got here.
* Language questions and proposals: open a discussion or issue first. Syntax changes are
  the most expensive kind — they must be right before they are written.

## The three-homes rule

New or changed syntax has **three homes, not one** — they must change together in the
same commit:

1. **The parser/checker/generator** — `tools/mujard/lang`
2. **The normative EBNF** — section 3 of `MUJARD_LANGUAGE_SPEC.md` (the grammar tests
   parse it straight out of the spec: `parser accepts / grammar rejects` means the spec
   is behind)
3. **The goldens** — `tools/mujard/lang/testdata` (`cpp` + `trees`)

A change that lands in only some of them fails tests on purpose. Regenerate goldens with:

```sh
cd tools/mujard
go test ./lang -run 'TestGenerateSamples|TestGoldenTrees' -update
```

…then **review the diff** — golden changes are generated, not authored.

## Build and test

You need [Go](https://go.dev) 1.23+ and a C++ compiler (`g++`/`clang++`) for the host
tests.

```sh
cd tools/mujard
go build -o mujard .        # the CLI
go vet ./... && go test ./...   # everything
cd ../mujard-build && go vet . && go test . && go build -o mujard-build .
```

Language changes also refresh the version stamp — the test tells you when:

```sh
MUJARD_GOLDEN=1 go test ./lang -run TestLanguageVersionIsCurrent
```

## House rules

* **Sign off your commits** (DCO): end commit messages with
  `Signed-off-by: Your Name <you@example.com>` (`git commit -s`).
* Keep commits focused; one logical change per commit where practical.
* Match the surrounding code style — the codebase is deliberately plain and commented
  where the *why* isn't obvious.
* No secrets, keys, tokens or credentials — ever. `.gitignore` covers the usual spots;
  review your diff before pushing.
* This repository holds the **open-source parts** of the platform. Do not reference or
  depend on the proprietary apps/backend in code or docs — describe roles ("the desktop
  app", "the platform API"), not internal paths.
* Never flash or touch real hardware in tests; the suites simulate everything.

## Reporting security issues

Do **not** open a public issue. See [SECURITY.md](SECURITY.md) for the private channel.

## Releasing (maintainers)

A release is a version tag — `.github/workflows/release.yml` builds `mujard` and
`mujard-build` for Windows, Linux and macOS and publishes them on the GitHub Release
automatically.

1. Bump `const version` in `tools/mujard/main.go` (e.g. `"2.0.0"` → `"2.1.0"`).
2. Move `[Unreleased]` in `CHANGELOG.md` to a dated `[X.Y.Z]` section.
3. Commit, then tag and push both:

   ```sh
   git commit -am "Release X.Y.Z"
   git tag vX.Y.Z
   git push origin main --tags
   ```

4. Watch **Actions → Release** — when it finishes, the binaries and `checksums.txt`
   appear under [Releases](https://github.com/mujard-language/mujard-language/releases).

Notes:

- The tag name (`vX.Y.Z`) should match the CLI version in `main.go`.
- The *language* version is separate: `tools/mujard/lang/version.json` bumps itself via
  `MUJARD_GOLDEN=1 go test ./lang -run TestLanguageVersionIsCurrent` whenever the
  language sources change — don't hand-edit it.
- To test the workflow without a release, run it by hand: **Actions → Release → Run
  workflow** (it builds the artifacts but skips publishing, since no tag is pushed).

## License

By contributing you agree your contributions are licensed under
[Apache-2.0](LICENSE), the same as the project.
