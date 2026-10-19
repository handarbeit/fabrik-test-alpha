# Feature Specification: Cross-repo spawn end-to-end regression (alpha → beta)

**Feature Branch**: `fabrik/issue-9076`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610191543299409() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261019-154329-9409\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610191543299409() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. It exercises two previously fixed behaviors: on-demand initialization of a spawn-target repository (#803) and the per-run-unique fixture token (#1308).

The work spans two repositories with a hard ordering dependency. The beta repository (`handarbeit/fabrik-test-beta`) must gain a new greeting function first. The alpha repository (`handarbeit/fabrik-test-alpha`, this one) can then call it. Alpha cannot compile against a function that does not yet exist, so the change must be decomposed into a beta sub-issue that lands before the alpha change. The pipeline's Plan stage is expected to perform that decomposition by spawning a child issue targeting the beta repo.

The function name and string literal embed a per-run-unique token. The scenario's final assertion checks for that exact text in the beta repo, so both must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

A spawned child issue in the beta repo adds an exported greeting function that returns the per-run-unique token string, with a test, and lands first.

**Why this priority**: Alpha depends on it, and the scenario's final assertion checks that this exact text lands in the beta repo.

**Independent Test**: In the beta repo, the new function exists with the exact name, returns the exact literal, and its test passes.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610191543299409()` is called, **Then** it returns `e2e-cross-repo-spawn-20261019-154329-9409`.
2. **Given** the beta repo's greeting tests, **When** the test suite runs, **Then** a test covering the new function passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

After the beta change has landed, the alpha repo's `main.go` calls the new beta function and prints its result, in addition to its existing behavior.

**Why this priority**: It depends on User Story 1, and it proves the cross-repo ordering was respected.

**Independent Test**: Build and run the alpha program. Its output includes the token string, and it still does everything it did before.

**Acceptance Scenarios**:

1. **Given** the beta change has landed, **When** alpha is built, **Then** the build succeeds with `greeting.HelloE2E202610191543299409()` resolved.
2. **Given** the alpha program runs, **When** it executes, **Then** it prints `e2e-cross-repo-spawn-20261019-154329-9409` and its prior output is unchanged.

---

### User Story 3 - Plan decomposes work across repos (Priority: P1)

Plan recognizes that the work requires ordered changes in two repos and spawns a child issue in the beta repo, with alpha's change blocked until beta lands.

**Why this priority**: This is the behavior under test (#803 and #1308).

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-beta`, is linked as a blocker of this issue, and carries the exact function name and literal.

**Acceptance Scenarios**:

1. **Given** this spec reaches Plan, **When** Plan completes, **Then** a spawn block targeting the beta repo is emitted and a child issue is created there.
2. **Given** the child issue is open, **When** alpha's Implement would start, **Then** it waits until the child is resolved.

---

### Edge Cases

- The beta repo may need on-demand initialization before the child issue can be created or worked (#803).
- The token must not be paraphrased, regenerated or altered. Any divergence breaks the final assertion (#1308).
- If dependency tidying is a no-op, no `go.mod` or `go.sum` change is required in alpha.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202610191543299409() string` in its greeting package that returns exactly `e2e-cross-repo-spawn-20261019-154329-9409`.
- **FR-002**: The beta repo MUST include a test for that function in its greeting package tests.
- **FR-003**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610191543299409()` and print the returned value, in addition to its existing behavior.
- **FR-004**: The alpha repo MUST run a dependency tidy if the module files require it after the change.
- **FR-005**: The beta change MUST land before the alpha change, and the alpha work MUST be blocked on the beta work.
- **FR-006**: Decomposition MUST be performed by Plan emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo.
- **FR-007**: The function name and string literal MUST be used verbatim as given, with no substitution or invented alternatives.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue targeting `handarbeit/fabrik-test-beta` is created by Plan and linked as a blocker of this issue.
- **SC-002**: The exact text `HelloE2E202610191543299409` and `e2e-cross-repo-spawn-20261019-154329-9409` is present in the beta repo's default branch after landing.
- **SC-003**: Alpha builds successfully and its output contains the token string, only after the beta change has landed.
- **SC-004**: The beta repo's tests, including the new one, pass.

## Assumptions

- The beta repo already has a greeting package; the issue names its function file and test file.
- The alpha repo's `main.go` already imports or can import the beta module's greeting package.
- The pipeline's existing spawn and `blockedBy` mechanisms provide the ordering, so no new ordering mechanism is needed.

## Out of Scope *(optional)*

- Changes to pipeline or engine behavior (this issue only verifies existing behavior).
- Any change to the beta repo beyond the single function and its test.
- Any change to alpha's behavior beyond printing the new value.

## Source References *(optional)*

- #803 — on-demand spawn-target repo init
- #1308 — per-run-unique fixture token
