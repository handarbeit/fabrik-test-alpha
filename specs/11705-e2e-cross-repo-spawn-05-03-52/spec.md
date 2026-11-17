# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-11705`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202611171003520754() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261117-100352-0754".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202611171003520754() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This is an end-to-end regression scenario for cross-repo decomposition. One piece of work spans two repositories with a strict ordering dependency: the beta repository (`handarbeit/fabrik-test-beta`) must gain a new greeting function before the alpha repository (`handarbeit/fabrik-test-alpha`, this repo) can call it.

The scenario checks that:
- Plan decomposes the work by spawning a sub-issue in the beta repo, rather than trying to do everything in one repo.
- The spawned beta work lands before the dependent alpha work.
- Two earlier fixes hold: on-demand initialisation of a spawn-target repo (#803) and a per-run-unique fixture token (#1308).

The function name and string literal embed a per-run-unique token. The scenario's final assertion depends on that exact text appearing in the beta repo, so both must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repo exposes a new exported greeting function that returns the run-unique literal, with a test covering it.

**Why this priority**: The alpha change depends on it, and the scenario's final assertion checks for this exact text in the beta repo.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261117-100352-0754`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202611171003520754()` is called, **Then** it returns the string `e2e-cross-repo-spawn-20261117-100352-0754`.
2. **Given** the beta repo's greeting test file, **When** the tests run, **Then** a test for this function exists and passes.

---

### User Story 2 - Alpha repo uses the beta function (Priority: P1)

The alpha repo's main program calls the new beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It shows the cross-repo dependency is satisfied in the correct order.

**Independent Test**: With the beta change available, build and run the alpha program. It still does what it did before and also prints `e2e-cross-repo-spawn-20261117-100352-0754`.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** the alpha program runs, **Then** its output includes `e2e-cross-repo-spawn-20261117-100352-0754` and its previous behaviour is unchanged.
2. **Given** the alpha change, **When** the module builds, **Then** dependencies are tidy and the build succeeds.

---

### User Story 3 - Plan decomposes across repos in strict order (Priority: P1)

Plan emits a spawn block targeting the beta repo, so the beta work is a separate sub-issue that must complete before the alpha work.

**Why this priority**: Verifying cross-repo decomposition is the purpose of this scenario (#803, #1308).

**Independent Test**: After Plan, a sub-issue exists in the beta repo describing the beta change, and the alpha issue is blocked on it.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** it has emitted a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta`.
2. **Given** the spawned beta sub-issue is open, **When** the alpha work is scheduled, **Then** it does not proceed until the beta work has landed.

---

### Edge Cases

- The function name and string literal must be used exactly as given. Substituting a different name or token breaks the scenario's final assertion.
- If the alpha module needs dependency changes after the import, `go mod tidy` must leave the module consistent.
- If the beta repo has not been initialised for spawning yet, the spawn path must initialise it on demand (#803).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's greeting package MUST export a function named exactly `HelloE2E202611171003520754` that takes no arguments and returns a string.
- **FR-002**: That function MUST return exactly the string `e2e-cross-repo-spawn-20261117-100352-0754`.
- **FR-003**: The beta repo MUST include a test for that function in its greeting package tests, and the test MUST pass.
- **FR-004**: The alpha repo's main program MUST call `greeting.HelloE2E202611171003520754()` and print the result, while keeping its existing behaviour.
- **FR-005**: The alpha module's dependencies MUST be tidy after the change (run `go mod tidy` if needed).
- **FR-006**: The beta change MUST land before the alpha change that depends on it.
- **FR-007**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo (`handarbeit/fabrik-test-beta`).
- **FR-008**: The function name and string literal MUST be used verbatim, with no substitution.

### Key Entities

- **Alpha repo** (`handarbeit/fabrik-test-alpha`): the repo this issue lives in. It consumes the new greeting function.
- **Beta repo** (`handarbeit/fabrik-test-beta`): the repo that provides the new greeting function and receives the spawned sub-issue.
- **Run-unique token**: the per-run-unique text embedded in the function name and literal, checked by the scenario's final assertion.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The beta repo contains the exported function `HelloE2E202611171003520754` returning exactly `e2e-cross-repo-spawn-20261117-100352-0754`, which the scenario's final assertion finds.
- **SC-002**: The beta repo's greeting package tests pass, including the new test.
- **SC-003**: The alpha program builds and, when run, prints `e2e-cross-repo-spawn-20261117-100352-0754` while keeping its previous output.
- **SC-004**: A sub-issue targeting the beta repo was created by Plan's spawn, and the beta change merged before the alpha change.

## Assumptions

- The beta repo already has a greeting package with a source file and a test file, as the issue implies.
- Printing the result means writing it to standard output alongside the existing output.
- Exactly one spawned sub-issue in the beta repo is sufficient to deliver the beta-side changes.

## Out of Scope *(optional)*

- Any change to the beta or alpha repos beyond the function, its test, and the alpha call and print.
- Changing the engine's spawn, board, or pipeline logic. This scenario only exercises existing behaviour.

## Source References *(optional)*

- #803: on-demand spawn-target repo initialisation
- #1308: per-run-unique fixture token
