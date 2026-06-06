# DSA — Data Structures and Algorithms in Java

A personal collection of classic **Data Structures and Algorithms** implemented in **Java**, organized as a Maven project. Each algorithm is written from scratch (no third-party libraries) and paired with **JUnit** tests to verify correctness.

> 🎯 The goal of this repository is to build a clean, well-tested, and easy-to-read reference library of fundamental algorithms, one topic at a time.

---

## 📑 Table of Contents

- [Overview](#-overview)
- [Features](#-features)
- [Project Structure](#-project-structure)
- [Prerequisites](#-prerequisites)
- [Getting Started](#-getting-started)
- [Building & Testing](#-building--testing)
- [Implemented Algorithms](#-implemented-algorithms)
- [Contributing](#-contributing)
- [License](#-license)

---

## 🔍 Overview

This repository contains implementations of common algorithms, starting with **sorting algorithms**. Every implementation:

- Lives in its own class under a topic-specific package (e.g., `io.github.mthanuj.sorting`).
- Follows a simple, consistent API (`solve(...)` for the main operation).
- Ships with unit tests using **JUnit 4**.

The project is intentionally lightweight — just **Java + Maven** — so you can clone, build, and run the tests in seconds.

---

## ✨ Features

- ✅ **Pure Java** implementations (no external algorithm libraries).
- ✅ **Maven** build system for easy compilation and dependency management.
- ✅ **JUnit 4** unit tests for every algorithm.
- ✅ **Topic-based packaging** (e.g., `sorting`, `searching`, `graph`, …) for scalability.
- ✅ **Java 26** — uses the latest language features.
- ✅ **Clean & readable** code, ideal for learning and interview prep.

---

## 📂 Project Structure

```
dsa/
├── pom.xml
├── .gitignore
└── src/
    ├── main/
    │   └── java/
    │       └── io/github/mthanuj/
    │           └── sorting/
    │               └── SelectionSort.java
    └── test/
        └── java/
            └── io/github/mthanuj/
                └── sorting/
                    └── SelectionSortTest.java
```

| Path | Purpose |
| --- | --- |
| `pom.xml` | Maven project descriptor (dependencies, plugins, Java version). |
| `src/main/java/` | Production source code (algorithm implementations). |
| `src/test/java/` | Unit tests for the corresponding implementations. |
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

## 📚 Implemented Algorithms

### 🔢 Sorting

| Algorithm | Class | Time Complexity (Best / Avg / Worst) | Space | Stable |
| --- | --- | --- | --- | --- |
| [Selection Sort](src/main/java/io/github/mthanuj/sorting/SelectionSort.java) | `SelectionSort` | `O(n²)` / `O(n²)` / `O(n²)` | `O(1)` | ❌ |

**Selection Sort** — repeatedly finds the minimum element from the unsorted portion of the array and places it at the beginning. Simple but inefficient on large datasets; useful as a teaching example.

```java
int[] arr = { 5, 3, 4, 1, 2 };
SelectionSort sorter = new SelectionSort();
int[] sorted = sorter.solve(arr); // [1, 2, 3, 4, 5]
```

> 💡 More algorithms are on the way! Topics planned: **Searching**, **Recursion**, **Dynamic Programming**, **Graphs**, **Trees**, and more.

---

## 🤝 Contributing

Contributions are welcome! Whether it's fixing a typo, adding a new algorithm, or improving tests — every bit helps.

1. **Fork** the repository.
2. Create a new branch for your feature:

   ```bash
   git checkout -b feature/new-algorithm
   ```

3. **Add your implementation** under the appropriate topic package
   (e.g., `src/main/java/io/github/mthanuj/<topic>/`).
4. **Write unit tests** in the matching test package
   (e.g., `src/test/java/io/github/mthanuj/<topic>/`).
5. Make sure all tests pass:

   ```bash
   mvn test
   ```

6. **Commit** your changes and **push** to your fork.
7. Open a **Pull Request** describing what you added and why.

### Coding Conventions

- One public class per file, named after the algorithm (e.g., `QuickSort`).
- Expose the main behavior via a `solve(...)` method (or a similarly named entry point).
- Keep code self-documenting; add Javadoc only where the intent isn't obvious.
- Always include JUnit tests covering the **typical**, **edge**, and **adversarial** cases.

---

## 📄 License

This project is open source. If a license is not yet specified, the source is provided **as-is** for educational purposes. Please add a `LICENSE` file (e.g., **MIT**, **Apache-2.0**) if you intend to share it publicly.

---

## 🌟 Acknowledgements

- Inspired by classic algorithm textbooks (CLRS, *Introduction to Algorithms*).
- Powered by [Maven](https://maven.apache.org/) and [JUnit](https://junit.org/junit4/).

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/mThanuj">@mThanuj</a>
</p>
