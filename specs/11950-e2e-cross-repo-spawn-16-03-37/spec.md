# Feature Specification: E2E cross-repo spawn (beta-first decomposition)

**Feature Branch**: `fabrik/issue-11950`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202611262103375137() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261126-210337-5137\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202611262103375137() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for cross-repo decomposition. It exercises two earlier behaviours: on-demand initialisation of a spawn-target repository (#803) and the use of a per-run-unique fixture token (#1308).

The work genuinely spans two repositories with a strict ordering dependency. The beta repository (`handarbeit/fabrik-test-beta`) must gain a new greeting function first, because the alpha repository (`handarbeit/fabrik-test-alpha`, this one) cannot call a function that does not yet exist. The pipeline is therefore expected to decompose the work: the Plan stage should spawn a child issue targeting the beta repository, and this issue's own change should depend on that child landing first.

The function name and string literal below contain a per-run-unique token. The scenario's final assertion checks for that exact text in the beta repository, so both must be used verbatim and never substituted with invented values.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repository exposes a new exported greeting function that returns the run-unique literal, covered by a test. This is delivered through a child issue spawned from this one.

**Why this priority**: It is the prerequisite for the alpha change and the target of the scenario's final assertion.

**Independent Test**: In the beta repository, run the greeting package tests and confirm the new function returns exactly `e2e-cross-repo-spawn-20261126-210337-5137`.

**Acceptance Scenarios**:

1. **Given** the beta repository's greeting package, **When** `HelloE2E202611262103375137()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261126-210337-5137`.
2. **Given** the beta repository's greeting tests, **When** the test suite runs, **Then** a test for the new function exists and passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

The alpha repository's entry point calls the new beta function and prints its result, in addition to its existing behaviour, only after the beta change has landed.

**Why this priority**: It is the dependent half of the decomposition and proves the ordering constraint holds.

**Independent Test**: Build and run the alpha program after beta has landed, and confirm its output includes `e2e-cross-repo-spawn-20261126-210337-5137` alongside its prior output.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** the alpha program runs, **Then** it prints the value returned by `greeting.HelloE2E202611262103375137()` and still performs its previous behaviour.
2. **Given** the beta function has not yet landed, **When** alpha's work would begin, **Then** it is held back until beta is complete (it is blocked by the beta child issue).

---

### User Story 3 - Plan decomposes across repos (Priority: P1)

The Plan stage recognises that the work crosses repositories and emits a spawn block targeting the beta repository, so a child issue is created there and linked as a blocker of this issue.

**Why this priority**: Verifying this decomposition behaviour is the purpose of the scenario.

**Independent Test**: After Plan completes, confirm a child issue exists in the beta repository, is on the board, and is linked as blocking this issue.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** a spawn block targeting `handarbeit/fabrik-test-beta` is emitted and a child issue is created there.
2. **Given** the child issue exists, **When** the pipeline continues, **Then** this issue does not proceed to implementation until the child is resolved.

---

### Edge Cases

- The beta repository has not been initialised for spawn targets yet; on-demand initialisation (#803) must make it usable without manual setup.
- The function name or literal is altered or regenerated by a stage; this would break the final assertion, so the exact given text must be preserved in every stage and in the child issue.
- `go mod tidy` is unnecessary; it is run only if needed and must not introduce unrelated changes.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repository's `pkg/greeting/greeting.go` MUST export a function `HelloE2E202611262103375137() string` that returns exactly the literal `e2e-cross-repo-spawn-20261126-210337-5137`.
- **FR-002**: The beta repository's `pkg/greeting/greeting_test.go` MUST contain a test for `HelloE2E202611262103375137` that passes.
- **FR-003**: The alpha repository's `main.go` MUST call `greeting.HelloE2E202611262103375137()` in addition to its existing behaviour and print the result.
- **FR-004**: The alpha repository's dependencies MUST be tidied (`go mod tidy`) if, and only if, needed for the build to succeed.
- **FR-005**: The beta change MUST land before the alpha change; the alpha work MUST NOT be implemented or merged ahead of it.
- **FR-006**: The Plan stage MUST emit a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta` to carry the beta work as a separate child issue.
- **FR-007**: The exact function name `HelloE2E202611262103375137` and literal `e2e-cross-repo-spawn-20261126-210337-5137` MUST be used verbatim everywhere they appear, including in the spawned child issue; no substitute name or value may be invented.

### Key Entities

- **Alpha repository**: `handarbeit/fabrik-test-alpha`, the repository this issue lives in; consumes the new function.
- **Beta repository**: `handarbeit/fabrik-test-beta`, the spawn target; provides the new function.
- **Spawned child issue**: The beta-repo issue created by Plan that blocks this issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repository contains the exact text `HelloE2E202611262103375137` and `e2e-cross-repo-spawn-20261126-210337-5137` in its greeting package.
- **SC-002**: The beta greeting package's tests pass, including the new one.
- **SC-003**: The alpha program, once built and run, prints the exact literal `e2e-cross-repo-spawn-20261126-210337-5137` in addition to its prior output.
- **SC-004**: The beta change is merged before the alpha change, with no moment where alpha's main branch references a missing function.
- **SC-005**: Exactly one child issue is created in the beta repository for this decomposition, with no manual repository setup required.

## Assumptions

- The beta repository already contains a `pkg/greeting` package with `greeting.go` and `greeting_test.go`, per the issue text.
- The alpha repository's `main.go` can import the beta greeting package under the import path the beta module defines, and the dependency wiring is handled in the alpha change.
- The single child issue covers both the function and its test in beta.
- The scenario's per-run token has already been generated and is correct as written.

## Out of Scope *(optional)*

- Any other behaviour changes in either repository beyond the function, its test, and the alpha call-and-print.
- Changes to the cross-repo spawn mechanism itself (#803, #1308); this issue only exercises them.
- Release or versioning of either repository.

## Source References *(optional)*

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
