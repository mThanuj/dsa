# Pull Request

> 📖 **Before you submit**, please make sure you've read [`CONTRIBUTING.md`](../blob/main/CONTRIBUTING.md) — especially the **no-AI policy** and the **use-existing-packages** rule.

## 📋 Summary

<!-- One or two sentences describing what this PR does and why. -->

> e.g. "Add Insertion Sort to `io.github.mthanuj.sorting` with full JUnit 4 test coverage."

## 🔗 Related Issue

<!-- Link the issue this PR closes, if any. Use `Closes #<n>` so GitHub auto-links and auto-closes the issue on merge. -->

- Closes #
- Related: #

## 🔄 Type of Change

<!-- Pick the option(s) that apply. -->

- [ ] 🧮 New algorithm / data structure
- [ ] 🐛 Bug fix
- [ ] ⚡ Performance improvement
- [ ] 📖 Documentation only
- [ ] 🧪 Tests only
- [ ] ♻️ Refactor (no behavior change)
- [ ] 🛠️ Build / tooling / CI

## 🧠 No-AI Confirmation

<!-- This is a learning project. The code in this PR was written by me. -->

- [ ] I confirm that **all code in this PR was written by me** (no AI generation), in line with the project's [no-AI policy](../blob/main/CONTRIBUTING.md#-a-note-on-ai-assisted-contributions).
- [ ] I am able to **explain every part of this PR in my own words** if asked by a reviewer.

## 📦 Package / Location

<!-- Confirm your changes go into a package that already exists in the repo. Do NOT create new topic packages. -->

- Files added/modified under: `src/main/java/io/github/mthanuj/<topic>/` and `src/test/java/io/github/mthanuj/<topic>/`
- Topic package: `<topic>` (must already exist — e.g. `sorting`)
- New class(es) introduced (if any): `<Algorithm>`, `<Algorithm>Test>`

> ⚠️ If you think a brand-new topic package is warranted, **open a separate issue first** and don't create it in this PR.

## 🛠 What Changed

<!-- Bullet list of the actual changes. Be specific. -->

-
-
-

## 🧪 How I Tested It

<!-- Describe how you verified the change works. Include the commands you ran. -->

- [ ] `mvn clean test` passes locally
- [ ] Added/updated JUnit 4 tests covering:
  - [ ] Typical (random, mixed-order) input
  - [ ] Already-sorted input
  - [ ] Reverse-sorted input
  - [ ] All-duplicates input
  - [ ] Single element
  - [ ] Empty input
  - [ ] Two elements
  - [ ] Large input (≥ 10 000 elements)
- Manual testing notes (optional):

```text
# Example
$ mvn test
Tests run: 12, Failures: 0, Errors: 0, Skipped: 0
```

## ⚡ Performance / Benchmarks (optional)

<!-- Only fill this out for PRs that claim a performance improvement. -->

- Before: …
- After: …
- How measured: …

## 💥 Breaking Changes

<!-- If this PR changes a public API or the behavior of an existing algorithm, describe it here. Otherwise, write "None". -->

- None / Describe:

## 📸 Screenshots / Output (optional)

<!-- For PRs that change behavior visibly (e.g. CLI output, docs rendering), paste a screenshot or short output here. -->

## 📝 Checklist

- [ ] My code follows the project's [coding conventions](../blob/main/CONTRIBUTING.md#-coding-conventions).
- [ ] I have added Javadoc to any non-obvious public method or helper.
- [ ] I have updated the [README](../blob/main/README.md) if the change is user-visible (e.g. new algorithm).
- [ ] My commit messages follow the [Conventional Commits](../blob/main/CONTRIBUTING.md#-commit-message-conventions) format.
- [ ] I have run `mvn clean test` and all tests pass.
- [ ] I have read and agree to the [license terms](../blob/main/CONTRIBUTING.md#-license) (contributions are MIT-licensed).
