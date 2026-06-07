# DSA — Data Structures and Algorithms in Java

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Java](https://img.shields.io/badge/Java-26-orange.svg)](https://www.java.com)
[![Maven](https://img.shields.io/badge/Maven-3.6%2B-blue.svg)](https://maven.apache.org/)
[![JUnit Jupiter](https://img.shields.io/badge/JUnit%20Jupiter-6.1.0-green.svg)](https://junit.org/junit5/)
[![CI](https://github.com/mThanuj/dsa/actions/workflows/ci.yml/badge.svg)](https://github.com/mThanuj/dsa/actions/workflows/ci.yml)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.1-ff69b4.svg)](CODE_OF_CONDUCT.md)

A personal collection of classic **Data Structures and Algorithms** implemented in **Java**, organized as a Maven project. Each algorithm is written from scratch (no third-party libraries) and paired with **JUnit** tests to verify correctness.

> 🎯 The goal of this repository is to build a clean, well-tested, and easy-to-read reference library of fundamental algorithms, one topic at a time.

> 🧠 **Note to contributors:** This is a **learning project**. Please read [`CONTRIBUTING.md`](CONTRIBUTING.md) — especially the **no-AI policy** and the **use-existing-packages** rule — before opening a PR.

---

## 📑 Table of Contents

- [DSA — Data Structures and Algorithms in Java](#dsa--data-structures-and-algorithms-in-java)
  - [📑 Table of Contents](#-table-of-contents)
  - [🔍 Overview](#-overview)
  - [✨ Features](#-features)
  - [📂 Project Structure](#-project-structure)
  - [✅ Prerequisites](#-prerequisites)
  - [🚀 Getting Started](#-getting-started)
    - [1. Clone the repository](#1-clone-the-repository)
    - [2. Verify the build](#2-verify-the-build)
  - [🛠 Building \& Testing](#-building--testing)
    - [Compile the project](#compile-the-project)
    - [Run all unit tests](#run-all-unit-tests)
    - [Package as a JAR](#package-as-a-jar)
    - [Clean build artifacts](#clean-build-artifacts)
  - [🧰 Shared Utilities](#-shared-utilities)
  - [📚 Implemented Algorithms](#-implemented-algorithms)
    - [🔢 Sorting](#-sorting)
  - [🤝 Contributing](#-contributing)
  - [📄 License](#-license)
  - [🌟 Acknowledgements](#-acknowledgements)

---

## 🔍 Overview

This repository contains implementations of common algorithms, starting with **sorting algorithms**. Every implementation:

- Lives in its own class under a topic-specific package (e.g., `io.github.mthanuj.sorting`).
- Follows a simple, consistent API (`solve(...)` for the main operation).
- Reuses a small set of **shared utilities** (e.g., array `swap` / `reverse`) from `io.github.mthanuj.Utils`.
- Ships with unit tests using **JUnit Jupiter (JUnit 5+)**.

The project is intentionally lightweight — just **Java + Maven** — so you can clone, build, and run the tests in seconds.

---

## ✨ Features

- ✅ **Pure Java** implementations (no external algorithm libraries).
- ✅ **Maven** build system for easy compilation and dependency management.
- ✅ **JUnit Jupiter (JUnit 5+)** unit tests for every algorithm.
- ✅ **Topic-based packaging** (e.g., `sorting`, `searching`, `graph`, …) for scalability.
- ✅ **Shared utilities** (`Utils.swap`, `Utils.reverse`, …) to keep implementations DRY.
- ✅ **Java 26** — uses the latest language features.
- ✅ **Clean & readable** code, ideal for learning and interview prep.

---

## 📂 Project Structure

```
dsa/
├── pom.xml
├── .gitignore
├── LICENSE
├── README.md
├── CONTRIBUTING.md
└── src/
    ├── main/
    │   └── java/
    │       └── io/github/mthanuj/
    │           ├── Utils.java
    │           └── <topic>/          # One package per existing topic
    │               └── <Algorithm>.java
    └── test/
        └── java/
            └── io/github/mthanuj/
                └── <topic>/          # Mirror of the main package structure
                    └── <Algorithm>Test.java
```

| Path | Purpose |
| --- | --- |
| `pom.xml` | Maven project descriptor (dependencies, plugins, Java version). |
| `src/main/java/io/github/mthanuj/Utils.java` | Shared static utilities (`swap`, `reverse`, …) used by the algorithms. |
| `src/main/java/io/github/mthanuj/<topic>/` | Production source code (algorithm implementations), grouped by topic. |
| `src/test/java/io/github/mthanuj/<topic>/` | Unit tests for the corresponding implementations. |
| `/target/` | Build output (gitignored). |

---

## ✅ Prerequisites

Make sure you have the following installed:

- **Java Development Kit (JDK) 26** or newer
  - Check with: `java -version`
- **Apache Maven 3.6+**
  - Check with: `mvn -v`
- **Git** (to clone the repository)

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/mThanuj/dsa.git
cd dsa
```

### 2. Verify the build

```bash
mvn clean compile
```

If everything compiles, you're ready to go! 🎉

---

## 🛠 Building & Testing

### Compile the project

```bash
mvn compile
```

### Run all unit tests

```bash
mvn test
```

### Package as a JAR

```bash
mvn package
```

The resulting JAR will be placed in the `target/` directory.

### Clean build artifacts

```bash
mvn clean
```

---

## 🧰 Shared Utilities

The class [`Utils`](src/main/java/io/github/mthanuj/Utils.java) provides small, **static**, dependency-free helpers reused across algorithms:

| Method | Description |
| --- | --- |
| `Utils.swap(int[] arr, int i, int j)` | Swaps the elements at indices `i` and `j` in `arr`. |
| `Utils.reverse(int[] arr)` | Reverses the entire array in place. |
| `Utils.reverse(int[] arr, int start, int end)` | Reverses the sub-array from `start` (inclusive) to `end` (exclusive) in place. |

> 🔒 The constructor is private — `Utils` is **not** meant to be instantiated.

Example:

```java
int[] arr = { 1, 2, 3 };
Utils.swap(arr, 0, 2);   // arr -> { 3, 2, 1 }
Utils.reverse(arr);      // arr -> { 1, 2, 3 }
```

---

## 📚 Implemented Algorithms

Every algorithm lives in its own class under a topic-specific package (e.g., `io.github.mthanuj.sorting`) and follows the same uniform API: instantiate the class, call `solve(...)`, and the input is processed in place (and also returned for convenience).

```java
import io.github.mthanuj.sorting.<Algorithm>;

int[] arr = { /* some input */ };
<Algorithm> algorithm = new <Algorithm>();
int[] sorted = algorithm.solve(arr);
```

The tables below are the **source of truth** for what's implemented. New topics and entries are added here as they're solved — keep the columns consistent: a **name → file link**, a short **description**, and **pointers** on how the algorithm is meant to be reasoned about (key invariant, core idea, common pitfalls).

> 💡 **Adding a new entry?** Append a new row to the relevant topic's table, or create a new `### <Topic>` section with the same three columns. Keep the package link right under the topic heading.

### 🔢 Sorting

> 📖 Package: [`io.github.mthanuj.sorting`](src/main/java/io/github/mthanuj/sorting/)

| Name | Description | Pointers |
| --- | --- | --- |
| [BubbleSort](src/main/java/io/github/mthanuj/sorting/BubbleSort.java) | Repeatedly walks the array and swaps adjacent elements that are out of order, "bubbling" the largest unsorted element to the end on each pass. | Outer loop = pass index (grows the sorted suffix); inner loop = pairwise comparisons up to `arr.length - i - 1`; swap with `Utils.swap`. |
| [SelectionSort](src/main/java/io/github/mthanuj/sorting/SelectionSort.java) | Finds the minimum element in the unsorted prefix and swaps it into the next sorted position at the front. | Outer loop = boundary of the sorted prefix; inner loop scans for the index of the smallest remaining element; one `Utils.swap` per pass. |
| [InsertionSort](src/main/java/io/github/mthanuj/sorting/InsertionSort.java) | Builds the sorted array one element at a time, taking the next unsorted element and shifting larger sorted elements right to make room for it. | Outer loop = the element being inserted (starting at `i = 1`); inner `while` shifts sorted elements greater than the held `value` one slot to the right; drop the value into the gap at `arr[j + 1]`. |
| [MergeSort](src/main/java/io/github/mthanuj/sorting/MergeSort.java) | A divide-and-conquer sort that recursively splits the array in half, sorts each half, then merges the two sorted halves back together. | Recursion base = `low >= high` (single element is sorted); split range `[low..high]` at `mid = (low + high) / 2`, sort `[low..mid]` and `[mid+1..high]`, then `merge`; merge uses a temporary buffer to walk two pointers and copies results back into the original range. |

---

## 🤝 Contributing

Contributions are welcome! But this is a **learning project**, so please take a moment to read **[`CONTRIBUTING.md`](CONTRIBUTING.md)** before opening an issue or PR. The guide covers:

- 🧠 **The no-AI policy** — write the code yourself; AI is fine as a tutor but not an author.
- 📦 **Use existing packages** — add your `.java` files to the packages already in the repo (e.g., `io.github.mthanuj.sorting`). Don't create new topic packages; if you think one is needed, open an issue first.
- 🛠 **Development setup** (JDK 26, Maven, IDE)
- 📏 **Coding conventions** (Java style, naming, `solve(...)` API, `Utils` reuse, Javadoc)
- ✅ **Testing guidelines** (typical, edge, and adversarial cases)
- 💬 **Commit message conventions** (Conventional Commits)
- 🔁 **Pull request process** (branching, review, merge)

Quick start:

1. **Fork** the repository.
2. **Create a feature branch** from `main`:

   ```bash
   git checkout -b feat/your-algorithm-name
   ```

3. **Add your new algorithm** as a `.java` file inside one of the **existing** topic packages (and a matching `<Algorithm>Test` in the parallel test package).
4. **Run the full test suite** locally:

   ```bash
   mvn clean test
   ```

5. **Open a Pull Request** against `main`.

> 📖 The full guide ([`CONTRIBUTING.md`](CONTRIBUTING.md)) is the source of truth — please give it a read before opening your first PR.

---

## 📄 License

This project is licensed under the **MIT License** — a short, permissive license that lets people do almost anything they want with the code, including using it commercially, as long as the copyright notice and license terms are preserved.

See the full license text in the [`LICENSE`](LICENSE) file.

```
MIT License

Copyright (c) 2026 mThanuj

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

---

## 🌟 Acknowledgements

- Inspired by classic algorithm textbooks (CLRS, _Introduction to Algorithms_).
- Powered by [Maven](https://maven.apache.org/) and [JUnit Jupiter](https://junit.org/junit5/).

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/mThanuj">@mThanuj</a>
</p>
