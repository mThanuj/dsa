# Contributing to DSA

First off, thank you for taking the time to contribute! 🎉
This project is a personal collection of classic **Data Structures and Algorithms** in Java, and every contribution — whether it's fixing a typo, adding a new algorithm, or improving tests — helps make it a better learning resource for everyone.

> Please read this guide in full **before** opening an issue or submitting a pull request. It exists to make the contribution process smooth, predictable, and pleasant for everyone involved.

---

## 📑 Table of Contents

- [Contributing to DSA](#contributing-to-dsa)
  - [📑 Table of Contents](#-table-of-contents)
  - [🧠 A Note on AI-Assisted Contributions](#-a-note-on-ai-assisted-contributions)
  - [📜 Code of Conduct](#-code-of-conduct)
  - [🎯 What We're Looking For](#-what-were-looking-for)
  - [🤝 How You Can Contribute](#-how-you-can-contribute)
    - [🐛 Reporting Bugs](#-reporting-bugs)
    - [✨ Suggesting Enhancements](#-suggesting-enhancements)
    - [🌱 Your First Code Contribution](#-your-first-code-contribution)
    - [📖 Improving Documentation](#-improving-documentation)
  - [🛠 Development Setup](#-development-setup)
    - [Prerequisites](#prerequisites)
    - [Fork \& Clone](#fork--clone)
    - [Workflow](#workflow)
  - [📂 Project Layout](#-project-layout)
  - [📏 Coding Conventions](#-coding-conventions)
    - [☕ Java Style](#-java-style)
    - [📦 Package \& Naming Rules](#-package--naming-rules)
    - [♻️ Reuse the Shared Utils](#️-reuse-the-shared-utils)
    - [📝 Documentation](#-documentation)
  - [✅ Testing Guidelines](#-testing-guidelines)
    - [Test Class Template](#test-class-template)
    - [Running the Tests](#running-the-tests)
  - [💬 Commit Message Conventions](#-commit-message-conventions)
    - [Format](#format)
    - [Types](#types)
    - [Examples](#examples)
    - [Rules of Thumb](#rules-of-thumb)
  - [🔁 Pull Request Process](#-pull-request-process)
  - [👀 Review Process](#-review-process)
  - [📄 License](#-license)

---

## 🧠 A Note on AI-Assisted Contributions

**Please do not use AI tools (ChatGPT, Copilot, Claude, Cursor, etc.) to generate your contributions to this repository.**

This isn't a code dump — it's a **learning project**. The whole point is the journey: struggling with a problem, looking up references, making mistakes, refactoring, and slowly building deep, lasting understanding of how these algorithms actually work under the hood.

When you let an AI write the code for you:

- 🚫 **You skip the thinking.** The reasoning, the dead ends, and the "aha!" moments are where the learning happens.
- 🚫 **You don't build intuition.** Future-you, debugging a similar problem at work, won't have the mental model.
- 🚫 **You can't fully defend your PR.** If you can't explain *why* a line is there, the review becomes a guessing game.
- 🚫 **It dilutes the repo.** This collection is meant to be authentic, hand-crafted reference code — not AI output.

> 💡 **AI is a great tutor, not a great author.** It's totally fine (and even encouraged) to use AI to *explain* a concept, *discuss* trade-offs, or *debug* a specific error message. But the code you submit should be **yours** — written, reasoned about, and tested by you.

Maintainers reserve the right to ask you to explain any part of your PR in your own words. If you can't, the PR will be closed.

---

## 📜 Code of Conduct

This project follows the [**Contributor Covenant**](https://www.contributor-covenant.org/) (v2.1). The full text lives in [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md) — please read it before participating.

By participating, you agree to:

- **Be respectful and inclusive.** Harassment of any kind is not tolerated.
- **Assume good faith.** Disagreements happen; keep them focused on the work.
- **Be constructive.** Critique ideas, not people.
- **Welcome newcomers.** We were all beginners once.

If you experience or witness unacceptable behavior, contact the maintainer through the channel listed in [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md#enforcement). All complaints are reviewed promptly, fairly, and confidentially.

---

## 🎯 What We're Looking For

Great contributions for this repo typically look like:

- 🧮 **A new algorithm** in an **existing** package (e.g., adding a new sort to `io.github.mthanuj.sorting`) with full JUnit tests.
- 🐛 **A bug fix** in an existing implementation or its tests.
- ⚡ **A performance improvement** with benchmarks that justify the change.
- 📚 **Improved documentation** (Javadoc, README, complexity notes).
- ✅ **More test cases** (edge cases: empty arrays, single element, duplicates, already sorted, reverse sorted, large inputs).

> 💡 If you'd like to add a **new topic** (e.g., a brand-new `searching/` or `graph/` package), please open an issue first to discuss — the maintainer curates the set of topic packages that exist.

> 📋 **When opening an issue, please use one of the provided templates** under [`.github/ISSUE_TEMPLATE/`](.github/ISSUE_TEMPLATE/):
>
> - 🐛 **Bug Report** — for incorrect behavior in an existing algorithm or its tests.
> - 🧮 **New Algorithm / Data Structure** — for proposing a new addition to an existing package.
> - 📖 **Documentation** — for typos, unclear explanations, broken links, or doc improvements.
>
> Blank issues are disabled, so the templates help us help you faster.

---

## 🤝 How You Can Contribute

### 🐛 Reporting Bugs

A great bug report is small, reproducible, and unambiguous. **Before opening an issue**, please:

1. **Search existing issues** to avoid duplicates.
2. **Reproduce the bug** against the latest `main` branch.
3. **Collect the facts**: Java version, Maven version, OS, exact command run, expected vs. actual output, full stack trace.

Open an issue with the following structure:

```markdown
**Describe the bug**
A clear, one-sentence summary.

**To Reproduce**
Steps to reproduce the behavior:
1. Run command `...`
2. Call method `...` with input `...`
3. See error

**Expected behavior**
What you expected to happen.

**Actual behavior**
What actually happened.

**Environment**
- OS: Windows 11 / macOS 15 / Ubuntu 24.04
- JDK: 26
- Maven: 3.9.x

**Additional context**
Stack trace, screenshots, related issues, etc.
```

### ✨ Suggesting Enhancements

We love proposals for improvements to existing algorithms or tests! When opening an enhancement issue, please include:

- **The package / algorithm it belongs to** (e.g., `io.github.mthanuj.sorting.MySort`).
- **The problem it solves** — Why is this useful?
- **The proposed API** — Sketch of the class/method signature.
- **Alternatives considered** — Other approaches and why this one is better.
- **References** — Textbook chapter, Wikipedia link, or paper.

> ⚠️ **Suggestions for entirely new topic packages** (e.g., introducing a `searching/` or `graph/` package) are welcome as proposals, but the maintainer curates which topic packages exist. Please open an issue first to discuss before sending a PR.

### 🌱 Your First Code Contribution

Not sure where to start? Look for issues labeled:

- [`good first issue`](https://github.com/mThanuj/dsa/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22) — small, well-scoped tasks perfect for newcomers.
- [`help wanted`](https://github.com/mThanuj/dsa/issues?q=is%3Aissue+is%3Aopen+label%3A%22help+wanted%22) — issues where we especially welcome outside help.

A great first contribution is adding **test cases** to an existing algorithm — e.g., empty array, single element, all duplicates, already sorted.

### 📖 Improving Documentation

Typos, broken links, unclear explanations, and missing Javadoc are all valid pull requests. Documentation PRs are often the fastest to review and merge.

---

## 🛠 Development Setup

### Prerequisites

- **JDK 26** or newer — `java -version`
- **Maven 3.6+** — `mvn -v`
- **Git** — `git --version`
- An IDE — IntelliJ IDEA, Eclipse, or VS Code with the *Extension Pack for Java*.

### Fork & Clone

```bash
# 1. Fork the repo on GitHub, then:
git clone https://github.com/<your-username>/dsa.git
cd dsa

# 2. Add the upstream remote
git remote add upstream https://github.com/mThanuj/dsa.git

# 3. Verify Maven can compile the project
mvn clean compile

# 4. Run the existing test suite to make sure everything passes
mvn test
```

### Workflow

```bash
# Always start from a fresh main
git checkout main
git pull upstream main

# Create a feature branch
git checkout -b feature/your-algorithm-name

# Make your changes, then run the checks locally
mvn clean test

# Push and open a PR
git push origin feature/your-algorithm-name
```

---

## 📂 Project Layout

```
src/
├── main/java/io/github/mthanuj/
│   ├── Utils.java              # Shared static helpers (swap, reverse, …)
│   └── <topic>/               # One package per existing topic
│       └── <Algorithm>.java   # One public class per algorithm
└── test/java/io/github/mthanuj/
    └── <topic>/               # Mirror of the main package structure
        └── <Algorithm>Test.java
```

- The `io.github.mthanuj` group is the **top-level namespace**.
- Each **topic** (sorting, searching, graph, …) has its own sub-package. The maintainer decides which topic packages exist.
- Each **algorithm** lives in its own file and class.
- Each algorithm must have a **matching `<Algorithm>Test`** in the parallel test package.

---

## 📏 Coding Conventions

### ☕ Java Style

- **Java version:** Source/target is **26** (set in `pom.xml`). Use modern language features where they improve clarity.
- **Indentation:** **4 spaces** — no tabs.
- **Line length:** Aim for **≤ 120 characters**; hard cap at 140.
- **Braces:** K&R style — opening brace on the **same line**.
- **Naming:**
  - `PascalCase` for classes and interfaces
  - `camelCase` for methods, fields, and local variables
  - `UPPER_SNAKE_CASE` for constants
  - `lower.snake.case` for package names
- **`final` by default:** Mark local variables, parameters, and fields `final` whenever they are not reassigned.
- **No wildcard imports.** Use explicit imports, alphabetized.
- **No unused code.** Don't leave commented-out blocks or `@SuppressWarnings` without justification.
- **Prefer immutability** for value-like types and return types where practical.

### 📦 Package & Naming Rules

- One **public class per file**, named exactly after the algorithm.
- The main entry point must be a method named **`solve(...)`** (e.g., `public int[] solve(int[] arr)`). This keeps the API uniform across all algorithms.
- **Add your new `.java` files into one of the packages that already exists** in `src/main/java/io/github/mthanuj/<topic>/` and its mirror under `src/test/java/io/github/mthanuj/<topic>/`.
- **Do not create new topic packages** (e.g., don't introduce a `searching/` or `graph/` folder) — the maintainer curates which topic packages exist. If you think a new topic package is warranted, open an issue first.
- Use **meaningful, descriptive names** — avoid abbreviations like `arr` only when a more specific name is possible (e.g., `numbers`, `input`).

### ♻️ Reuse the Shared Utils

Before inlining low-level operations, check [`io.github.mthanuj.Utils`](src/main/java/io/github/mthanuj/Utils.java). It currently provides:

- `Utils.swap(int[] arr, int i, int j)`
- `Utils.reverse(int[] arr)`
- `Utils.reverse(int[] arr, int start, int end)`

If you need a helper that **two or more** algorithms would benefit from, **add it to `Utils`** instead of duplicating it. New helpers must:

- Be `public static`.
- Be **pure** (no side effects beyond the documented behavior).
- Have clear Javadoc describing preconditions (e.g., valid indices).

### 📝 Documentation

- Add a **class-level Javadoc** explaining what the algorithm does, its **time and space complexity** (best / average / worst), and whether it is **stable** / **in-place** when relevant.
- Add Javadoc to any **non-obvious public method** or helper.
- Keep the description accurate — if the complexity changes, update the docs in the same PR.

Example:

```java
/**
 * Sorts an array of integers in ascending order using the <Algorithm Name> algorithm.
 *
 * <p><Brief description of how the algorithm works, in plain English.>
 *
 * <p><b>Time complexity:</b> O(?) best / O(?) average / O(?) worst.<br>
 * <b>Space complexity:</b> O(?) — in place or extra.<br>
 * <b>Stable:</b> yes / no.
 *
 * @param arr the array to sort; may be {@code null} or empty (no-op in both cases)
 * @return the same array, sorted in ascending order
 */
public int[] solve(int[] arr) { ... }
```

---

## ✅ Testing Guidelines

Tests are **not optional** in this project. Every algorithm must ship with JUnit 4 tests covering:

| Case | Why it matters |
| --- | --- |
| **Typical** (random, mixed-order input) | The happy path. |
| **Already sorted** | Verifies best-case behavior and stability. |
| **Reverse sorted** | Classic worst-case for many algorithms. |
| **All duplicates** | Stresses the `≤` / `<` comparison logic and stability. |
| **Single element** | Boundary case — the loop should be a no-op. |
| **Empty array** | Boundary case — must not throw. |
| **Two elements** | Smallest non-trivial case. |
| **Large input** (e.g., 10 000 elements) | Sanity-checks scalability. |

### Test Class Template

```java
package io.github.mthanuj.<topic>;

import static org.junit.Assert.assertArrayEquals;

import org.junit.Test;

public class <Algorithm>Test {

    private final <Algorithm> algorithm = new <Algorithm>();

    @Test
    public void sortsTypicalInput() {
        int[] arr = { /* some input */ };

        int[] expected = { /* sorted output */ };
        int[] actual = algorithm.solve(arr);

        assertArrayEquals(expected, actual);
    }

    @Test
    public void sortsReverseSortedInput() {
        int[] arr = { /* descending input */ };

        int[] expected = { /* ascending output */ };
        int[] actual = algorithm.solve(arr);

        assertArrayEquals(expected, actual);
    }

    // Add more @Test methods for edge cases …
}
```

### Running the Tests

```bash
mvn test                              # Run all tests
mvn test -Dtest=<Algorithm>Test       # Run a single test class
mvn test -Dtest=<Algorithm>Test#methodName   # Run a single test method
```

> ⚠️ **All tests must pass** before you open a pull request. CI will re-run them; local green saves everyone time.

---

## 💬 Commit Message Conventions

This project loosely follows [**Conventional Commits**](https://www.conventionalcommits.org/). A well-formed commit message makes `git log` and `git blame` dramatically more useful.

### Format

```
<type>(<scope>): <short summary>

[optional body — explain the *what* and *why*, not the *how*]

[optional footer — reference issues, breaking changes]
```

### Types

| Type | Use for |
| --- | --- |
| `feat` | A new algorithm or user-visible feature |
| `fix` | A bug fix |
| `docs` | Documentation only changes (README, Javadoc, …) |
| `test` | Adding or correcting tests |
| `refactor` | A code change that neither fixes a bug nor adds a feature |
| `perf` | A performance improvement |
| `chore` | Tooling, CI, build, or repo maintenance |

### Examples

```
feat(sorting): add <Algorithm Name> with <partition/strategy description>

Implements the <Algorithm Name> algorithm with the <strategy> approach
to improve performance on <scenario>.

- Time: O(?) avg / O(?) worst
- Space: O(?)
- Stable: yes / no
```

```
fix(sorting): <Algorithm> no longer mutates input when already sorted

Previously, the inner loop's `<=` check would cause an unnecessary
self-swap on equal elements. Replaced with `<` to preserve the
in-place guarantee.
```

```
docs: add new <Algorithm> to README
```

### Rules of Thumb

- Use the **imperative mood** in the summary: *"add"*, not *"added"* or *"adds"*.
- Keep the summary **under 72 characters**.
- **Capitalize** the first letter of the summary.
- Do **not** end the summary with a period.
- Reference issues in the footer: `Closes #12`, `Refs #34`.

---

## 🔁 Pull Request Process

1. **Create a feature branch** off `main`:

   ```bash
   git checkout -b feat/your-algorithm-name
   ```

2. **Make focused commits.** One logical change per commit; avoid squashing unrelated work.
3. **Run the full check locally** before pushing:

   ```bash
   mvn clean test
   ```

4. **Update documentation** if you changed public APIs, added a new algorithm, or changed complexity characteristics. At minimum, update the relevant section of the `README.md`.
5. **Push your branch** and **open a Pull Request** against `main`.
6. **Fill in the PR template** ([`.github/PULL_REQUEST_TEMPLATE.md`](.github/PULL_REQUEST_TEMPLATE.md)) completely. It enforces the no-AI policy, the use-existing-packages rule, and the test-coverage checklist so you don't forget anything.
7. **Link the related issue** with `Closes #<issue-number>` in the PR description.
8. **Be responsive to review feedback.** Push follow-up commits to the same branch — they will be squashed or merged as appropriate.

> 🧹 **Tip:** If your PR is still in progress, open it as a **draft** to signal that it's not ready for review yet.

---

## 👀 Review Process

- All PRs require **at least one review** from a maintainer before merge.
- Reviews focus on **correctness, tests, style, and clarity** — not personal preference.
- Expect **constructive comments**. If you disagree, explain your reasoning with data or references.
- Maintainers may (and probably will) ask you to **explain parts of your PR in your own words** to confirm understanding. This is part of the no-AI policy and helps us all learn together.
- After approval, the maintainer will **squash-merge** (default) or rebase-merge, depending on the project's history policy.
- Be patient — reviews can take a few days, especially for larger changes.

---

## 📄 License

By contributing to this project, you agree that your contributions will be licensed under the same **MIT License** that covers the project. See [`LICENSE`](LICENSE) for the full text.

> In short: you keep copyright on your own work, but you grant everyone the same permissive rights to use it that the rest of the project enjoys.

---

<p align="center">
  Thank you again for contributing — and remember: the struggle is the point. 🚀
</p>
