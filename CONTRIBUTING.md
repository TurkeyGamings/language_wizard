
# Contributing

Thanks for your interest in improving this project!

## Quick Start

* **Fork** the repo and create a feature branch: `git checkout -b feat/short-descriptor`.
* **Make changes** in small, focused commits.
* **Format** code with `gofmt` and keep imports tidy (`go fmt ./...`).
* **Test** everything: `go test ./...`. Add/adjust tests for any behavior you change.
* **Document** public APIs and update README examples if behavior changes.

## Style & Scope

* Prefer small PRs that solve one problem.
* Keep dependencies minimal; stick to the standard library when possible.
* Log messages and errors should be clear and actionable.

## Commit Messages

* Use clear, imperative subjects, e.g., `fix: handle empty language input`.
* If a change is breaking, include `BREAKING CHANGE:` in the body and explain the migration.

## Pull Requests

* Describe **what** and **why** (link related issues).
* Include usage notes and test coverage.
* Be ready to address review feedback; we aim for constructive, concise reviews.

## Reporting Issues

* Provide steps to reproduce, expected vs. actual behavior, environment details (Go version, OS), and logs if relevant.

## Licensing

By submitting a contribution, you agree that it will be licensed under the project’s **LGPL-2.1** license.
