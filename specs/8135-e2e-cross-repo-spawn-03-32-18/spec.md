# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-8135`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610140732181422() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261014-073218-1422\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610140732181422() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for cross-repo decomposition. It guards two earlier fixes: #803 (on-demand initialisation of the repo that a spawned sub-issue targets) and #1308 (a per-run-unique fixture token, so repeated runs cannot be confused with each other).

The feature needs changes in two repositories with a strict ordering dependency. The beta repo (`handarbeit/fabrik-test-beta`) must gain a new greeting function first. This repo (`handarbeit/fabrik-test-alpha`) can then import and call it. Alpha cannot compile against a function that does not exist yet, so the work cannot be done as a single change in one repo. It has to be decomposed into a beta sub-issue that lands first and an alpha change that depends on it.

The function name and the string literal both contain a per-run-unique token. The scenario's final assertion checks that this exact text lands in the beta repo, so both must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the unique greeting function (Priority: P1)

A sub-issue in the beta repo adds the per-run-unique greeting function and a test for it, and lands before any alpha change depends on it.

**Why this priority**: The alpha change cannot compile without it, and the scenario's final assertion checks that this text is in the beta repo.

**Independent Test**: In the beta repo, build and run the tests for the greeting package. The new function exists, returns the exact literal, and its test passes.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610140732181422()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261014-073218-1422`.
2. **Given** the new test in the beta greeting package, **When** the beta test suite runs, **Then** the test passes and covers the returned literal.

---

### User Story 2 - Alpha repo calls the beta function and prints the result (Priority: P2)

After the beta change has landed, the alpha repo's entry point calls the new greeting function and prints its result, in addition to its existing behaviour.

**Why this priority**: It completes the cross-repo feature, but it depends on User Story 1 landing first.

**Independent Test**: Build and run the alpha program with the beta change available. Its output includes `e2e-cross-repo-spawn-20261014-073218-1422` and still includes its prior output.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** the alpha program runs, **Then** it prints the literal `e2e-cross-repo-spawn-20261014-073218-1422`.
2. **Given** the alpha program's existing behaviour, **When** the change is applied, **Then** the existing output and behaviour are unchanged.
3. **Given** the alpha module's dependencies, **When** the change is applied, **Then** the module builds cleanly with its dependency metadata tidy (`go mod tidy` run if needed).

---

### User Story 3 - Plan decomposes the work across repos in strict order (Priority: P1)

The Plan stage recognises that the work spans two repos with a strict order and decomposes it with a spawn block targeting the beta repo. The beta sub-issue is created and lands before the alpha work proceeds.

**Why this priority**: This is the behaviour under test. It exercises the on-demand spawn-target repo initialisation (#803) and the per-run-unique token (#1308).

**Independent Test**: Run the issue through the pipeline. A sub-issue exists in the beta repo, the alpha work is blocked on it, and the beta sub-issue carries the exact function name and literal.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan finishes, **Then** its output contains a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta`.
2. **Given** the spawn block, **When** the engine processes it, **Then** a sub-issue is created in the beta repo and the alpha issue is blocked on it.
3. **Given** the beta sub-issue's description, **When** it is read, **Then** it contains the exact function name and exact string literal given in this spec.

---

### Edge Cases

- The spawn-target repo has not been initialised on this instance yet (the #803 case). The spawn must still succeed rather than failing for the missing repo.
- A previous run left a similar fixture behind. The per-run-unique token must keep the two runs distinguishable (the #1308 case).
- Alpha must not be changed to reference the beta function before the beta change has landed.
- The function name and literal must not be altered, abbreviated or regenerated at any stage. A different token breaks the final assertion.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST expose an exported function `HelloE2E202610140732181422() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202610140732181422()` MUST return exactly the string `e2e-cross-repo-spawn-20261014-073218-1422`.
- **FR-003**: The beta repo MUST include a test for `HelloE2E202610140732181422` in `pkg/greeting/greeting_test.go` that passes.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610140732181422()` and print the result, in addition to its existing behaviour.
- **FR-005**: The alpha module MUST build with its dependencies tidy (`go mod tidy` run if needed).
- **FR-006**: The beta change MUST land before the alpha change, because alpha depends on a function that exists only in beta.
- **FR-007**: The Plan stage MUST decompose the work with a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo.
- **FR-008**: The exact function name and string literal given in this spec MUST be used verbatim in every artifact (spawn block, beta change, alpha change). No substitute names or tokens may be introduced.

### Key Entities

- **Alpha repo**: `handarbeit/fabrik-test-alpha`, this repo, which consumes the greeting function.
- **Beta repo**: `handarbeit/fabrik-test-beta`, which provides the greeting package and receives the spawned sub-issue.
- **Per-run-unique token**: The text `E2E202610140732181422` in the function name and `e2e-cross-repo-spawn-20261014-073218-1422` in the literal, which ties the run to the final assertion.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains the function `HelloE2E202610140732181422` and the literal `e2e-cross-repo-spawn-20261014-073218-1422`.
- **SC-002**: Running the alpha program after both changes have landed prints `e2e-cross-repo-spawn-20261014-073218-1422`.
- **SC-003**: A sub-issue exists in the beta repo that was created by the Plan stage's spawn, and the alpha issue was blocked on it until it landed.
- **SC-004**: The beta test suite and the alpha build and test suite both pass.

## Assumptions

- The beta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`, and alpha already depends on it as a module.
- The alpha repo's `main.go` already exists and may already print other output, which must be preserved.
- The `fabrik:yolo` label on this issue means the pipeline auto-advances. The spec is still written to be complete without human review.
- Both repositories are reachable by the instance running this scenario.

## Out of Scope

- Any change to the engine's spawn, board or dependency logic. This issue only exercises it.
- Any behaviour change in the greeting package beyond the single new function and its test.
- Any change in alpha beyond calling and printing the new function.
- Cleanup of fixtures from earlier runs.

## Source References

- #803: on-demand spawn-target repo initialisation
- #1308: per-run-unique fixture token
