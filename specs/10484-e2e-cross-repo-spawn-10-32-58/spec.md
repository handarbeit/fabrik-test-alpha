# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-10484`
**Created**: 2026-10-31
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610311432580141() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261031-143258-0141\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610311432580141() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It exercises on-demand initialisation of a spawn-target repository (#803) and the per-run-unique fixture token (#1308). The work genuinely spans two repositories with a hard ordering: the alpha repo's entry point must call a greeting function that only exists once the beta repo has gained it. Plan is therefore expected to decompose the work by spawning a child issue in the beta repo (`handarbeit/fabrik-test-beta`) that must land before the alpha change. The scenario's final assertion depends on the exact function name and string literal below appearing in the beta repo, so these values are fixed and must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the per-run greeting function (Priority: P1)

The beta repo gains a new exported greeting function returning the per-run-unique token, with a test, delivered via a child issue spawned from this one.

**Why this priority**: The alpha change depends on it, and the scenario's final assertion checks that this exact text lands in the beta repo.

**Independent Test**: In the beta repo, run the package tests for `pkg/greeting`; the new test passes and asserts the exact literal.

**Acceptance Scenarios**:

1. **Given** the beta repo at its default branch, **When** the child issue's work lands, **Then** `pkg/greeting/greeting.go` exports `HelloE2E202610311432580141() string` returning exactly `e2e-cross-repo-spawn-20261031-143258-0141`.
2. **Given** the new function, **When** the beta repo's tests run, **Then** a test in `pkg/greeting/greeting_test.go` covers it and passes.

---

### User Story 2 - Alpha repo calls the new function (Priority: P2)

After the beta change has landed, this repo's `main.go` calls `greeting.HelloE2E202610311432580141()` in addition to its existing behaviour and prints the result.

**Why this priority**: It consumes the beta change and so must follow it; it completes the cross-repo flow.

**Independent Test**: Build and run the alpha program; its output includes `e2e-cross-repo-spawn-20261031-143258-0141` and it still does everything it did before.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** alpha is built, **Then** the build succeeds with module dependencies tidy.
2. **Given** the built alpha program, **When** it runs, **Then** it prints the string returned by `greeting.HelloE2E202610311432580141()` and retains its prior output.

---

### User Story 3 - Plan decomposes across repos in the correct order (Priority: P1)

Plan emits a spawn block targeting the beta repo so that the beta work is a child issue that blocks the alpha work.

**Why this priority**: Verifying this decomposition is the purpose of the scenario (#803, #1308).

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-beta`, is on the board, and is linked as a blocker of this issue.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan runs, **Then** it emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta` containing the beta-side work.
2. **Given** the spawned child is open, **When** the alpha work would be dispatched, **Then** it waits until the child is resolved.

---

### Edge Cases

- The spawn-target repo has not been initialised locally yet; it must be set up on demand (#803).
- The function name and literal contain a per-run-unique token; they must not be altered, regenerated, or paraphrased.
- `go mod tidy` may or may not change anything in alpha; it is run only if needed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) exports `HelloE2E202610311432580141() string` in `pkg/greeting/greeting.go`.
- **FR-002**: That function returns exactly the literal `e2e-cross-repo-spawn-20261031-143258-0141`.
- **FR-003**: `pkg/greeting/greeting_test.go` in the beta repo contains a test of the function asserting the exact literal.
- **FR-004**: This repo's `main.go` calls `greeting.HelloE2E202610311432580141()` in addition to its existing behaviour and prints the result.
- **FR-005**: Module files in this repo are tidied (`go mod tidy`) if the change requires it.
- **FR-006**: The beta change lands before the alpha change; the alpha work is blocked on the beta child issue.
- **FR-007**: Plan decomposes the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta`.
- **FR-008**: The function name and string literal are used verbatim, exactly as given in this spec.

### Key Entities

- **Beta repo**: `handarbeit/fabrik-test-beta`, the spawn target that owns the greeting package.
- **Alpha repo**: `handarbeit/fabrik-test-alpha`, this repo, which consumes the greeting function from its entry point.
- **Spawned child issue**: the beta-side work item that must complete before the alpha work.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The exact text `HelloE2E202610311432580141` and `e2e-cross-repo-spawn-20261031-143258-0141` is present in the beta repo's default branch after the run.
- **SC-002**: A child issue exists in the beta repo, created by the spawn from this issue, and is merged before the alpha change.
- **SC-003**: The alpha program builds and its output contains the literal string.
- **SC-004**: Tests pass in both repos.

## Assumptions

- `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` already exist in the beta repo; the new function and test are additive.
- The alpha repo already depends on the beta repo's greeting package, or acquires that dependency as part of the change, so that `go mod tidy` can resolve it once the beta change has landed.
- The token in the function name and literal is final for this run and will not be regenerated.

## Out of Scope

- Any change to the greeting package beyond the one function and its test.
- Any behaviour change in alpha other than the added call and print.
- Changes to the spawn mechanism itself; this issue only exercises it.
