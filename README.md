# mini-systems

**mini-systems** is a collection of small, focused implementations of foundational software systems.

Each project recreates a well-known concept - a language interpreter, a version control system, a shell, a compression algorithm - with the goal of understanding how it works from first principles. The emphasis is on clarity, minimalism, and learning by doing rather than completeness or production-readiness.

Most projects are written in Go and are intentionally scoped to remain readable and self-contained.

## Projects

| Project           | Description                                                                                         | Status            |
| ----------------- | --------------------------------------------------------------------------------------------------- | ----------------- |
| **[mini-lang](mini-lang/)**     | A small interpreted programming language, built to explore parsing, evaluation, and language design | Building          |
| **mini-git**      | A minimal implementation of Git's core data model (objects, trees, commits)                         | Planned          |
| **mini-bash**     | A tiny Unix shell supporting command execution, pipes, and redirection                              | Planned          |
| **mini-zip**      | A simple compression tool implementing classic compression algorithms                               | Planned          |
| **mini-terminal** | A minimal terminal emulator exploring PTYs, ANSI escape sequences, and screen buffers               | Planned          |

> Each project lives in its own directory and is intentionally scoped to remain readable and focused.

## Philosophy

This repository is not about production-ready tools.

It is about:

* learning through implementation rather than theory alone
* understanding core ideas deeply
* recreating systems without relying on heavy abstractions
* using minimal dependencies and standard libraries when possible
* building from scratch with minimal AI assistance to maximize learning
* making trade-offs explicit

Many features are deliberately omitted if they don’t serve the learning goal.

## Status

This repository is a living collection. Projects may be incomplete, experimental, or evolve over time as understanding deepens.

Feedback and discussion are welcome.

## License

[MIT](LICENSE)
