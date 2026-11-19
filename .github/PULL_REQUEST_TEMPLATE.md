## What this changes

<!-- One or two sentences. Link the issue it fixes, if any. -->

## Checklist

- [ ] `gofmt -l .` clean and `go vet ./...` clean in every module I touched
- [ ] `go test ./...` passes in every module I touched
- [ ] Language changes update all three homes: the parser (`tools/mujard/lang`), the normative grammar in `docs/lang/MUJARD_LANGUAGE_SPEC.md` §3, and the goldens (`go test ./lang -run 'TestGenerateSamples|TestGoldenTrees' -update`, with the diff reviewed)
- [ ] Docs updated where behavior or syntax changed
- [ ] Every commit carries a `Signed-off-by:` trailer (`git commit -s`) per CONTRIBUTING.md
