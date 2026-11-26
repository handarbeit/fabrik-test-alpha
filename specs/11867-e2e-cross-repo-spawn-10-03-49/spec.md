# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-11867`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202611261503499288() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261126-150349-9288\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202611261503499288() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for cross-repo decomposition. It exercises on-demand initialisation of a spawn-target repository (#803) and the per-run-unique fixture token (#1308). The work needs changes in two repositories in a strict order: the beta repo must gain a new greeting function before the alpha repo can call it. That ordering is what Plan is expected to decompose, by spawning a child issue in the beta repo that blocks this one.

The function name and the string literal both embed a per-run-unique token. The scenario's final assertion checks that exact text in the beta repo, so both must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repo (`handarbeit/fabrik-test-beta`) exposes a new exported greeting function that returns the run-unique string. A test covers it.

**Why this priority**: Alpha depends on it, and the scenario's final assertion inspects this change in the beta repo.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261126-150349-9288`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202611261503499288()` is called, **Then** it returns `e2e-cross-repo-spawn-20261126-150349-9288`.
2. **Given** the beta repo's greeting test file, **When** the tests run, **Then** a test for the new function exists and passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

The alpha repo (this repo) calls the beta function from its entry point and prints the result, in addition to its existing behaviour.

**Why this priority**: It depends on Story 1 landing first. It proves the cross-repo dependency order was respected.

**Independent Test**: With the beta change available, build and run alpha. Its output includes `e2e-cross-repo-spawn-20261126-150349-9288`, and its existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta change has landed, **When** alpha's entry point runs, **Then** it prints the string returned by `greeting.HelloE2E202611261503499288()`.
2. **Given** alpha's existing behaviour, **When** the new call is added, **Then** the existing behaviour is unchanged.
3. **Given** the module dependency changed, **When** the module files are tidied if needed, **Then** alpha builds cleanly.

---

### User Story 3 - Plan decomposes across repos in order (Priority: P1)

During Plan, the pipeline decomposes the work by spawning a child issue in the beta repo. The beta work completes before the alpha work proceeds.

**Why this priority**: This is the purpose of the scenario, namely verifying cross-repo spawn, including on-demand init of the spawn-target repo.

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-beta` and blocks this issue. Alpha's implementation does not start until the child is done.

**Acceptance Scenarios**:

1. **Given** the spec above, **When** Plan runs, **Then** it emits a spawn block targeting the beta repo for the beta-side work.
2. **Given** the child issue is created, **When** it is registered, **Then** it is linked as a blocker of this issue.
3. **Given** the beta child has landed, **When** alpha is implemented, **Then** the greeting function is importable.

---

### Edge Cases

- The function name or string literal is altered (for example a regenerated token). The final assertion fails, so both must be copied verbatim.
- Alpha is implemented before the beta function exists. The import fails to compile, so strict ordering is required.
- The beta repo is not yet initialised for spawning. The spawn must still succeed (#803).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's greeting package MUST export `HelloE2E202611261503499288() string`, returning exactly `e2e-cross-repo-spawn-20261126-150349-9288`.
- **FR-002**: The beta repo's greeting test file MUST contain a test for `HelloE2E202611261503499288`.
- **FR-003**: The alpha repo's entry point MUST call `greeting.HelloE2E202611261503499288()` and print the result, in addition to its existing behaviour.
- **FR-004**: The alpha repo's module files MUST be tidied if the new dependency requires it.
- **FR-005**: The beta-side change MUST land before the alpha-side change is implemented.
- **FR-006**: Plan MUST decompose the work by emitting a spawn block that targets the beta repo.
- **FR-007**: The function name and string literal MUST be used verbatim, with no alternative token invented.

### Key Entities

- **Alpha repo**: `handarbeit/fabrik-test-alpha`, the repo of this issue. It consumes the greeting function.
- **Beta repo**: `handarbeit/fabrik-test-beta`, the spawn target. It provides the greeting function.
- **Run-unique token**: `20261126-150349-9288`, embedded in the function name (as `202611261503499288`) and in the returned string.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The beta repo contains the exact function name and the exact returned string, and its greeting tests pass.
- **SC-002**: Alpha's run output includes `e2e-cross-repo-spawn-20261126-150349-9288`, and alpha builds and tests cleanly.
- **SC-003**: A child issue was spawned in the beta repo and completed before the alpha work was finished.
- **SC-004**: The scenario's final assertion finds the exact token text in the beta repo.

## Assumptions

- Alpha's existing entry point already has a way to print output, and the new print is simply added.
- The beta repo already has a greeting package, and the new function is added to it.
- The token is already unique per run, so no new token needs generating.
- Child issue creation in the beta repo is performed by the engine's existing spawn mechanism.

## Out of Scope *(optional)*

- Changing any other beta or alpha behaviour.
- Changes to the engine's spawn mechanism itself. This issue only exercises it.
- Any additional greeting functions or tokens.
