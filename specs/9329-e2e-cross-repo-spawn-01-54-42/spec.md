# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-9329`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610250554429282() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261025-055442-9282\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610250554429282() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It exercises two earlier fixes:

- #803: on-demand initialisation of a spawn-target repo.
- #1308: a per-run-unique fixture token, so that one run's result cannot be mistaken for another's.

The work spans two repositories with a strict ordering dependency. The beta repo (`handarbeit/fabrik-test-beta`) must gain a new greeting function before the alpha repo (`handarbeit/fabrik-test-alpha`, this repo) can call it. Planning is therefore expected to decompose the work into a sub-issue in the beta repo that lands first, and an alpha-side change that depends on it.

The scenario's final assertion checks that the exact function name and string literal, which carry the per-run-unique token, landed in the beta repo. They must be used verbatim and not replaced.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the per-run greeting function (Priority: P1)

The beta repo gains an exported greeting function that returns the per-run-unique literal, with a test covering it. This is the part the final assertion inspects, and the alpha repo depends on it.

**Why this priority**: The alpha change cannot compile without it, and the scenario's final assertion depends on the exact text landing here.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes, and calling `HelloE2E202610250554429282()` returns exactly `e2e-cross-repo-spawn-20261025-055442-9282`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610250554429282()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261025-055442-9282`.
2. **Given** the new test in the beta greeting package's test file, **When** the package tests run, **Then** the test passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

The alpha repo's `main.go` calls `greeting.HelloE2E202610250554429282()` from the beta repo and prints the result, in addition to its existing behaviour.

**Why this priority**: It completes the cross-repo change. It can only land after User Story 1 because the function must already exist in beta.

**Independent Test**: With the beta change available, build and run the alpha program. It still does what it did before and also prints `e2e-cross-repo-spawn-20261025-055442-9282`.

**Acceptance Scenarios**:

1. **Given** the beta function is merged, **When** the alpha program runs, **Then** its output includes `e2e-cross-repo-spawn-20261025-055442-9282` in addition to its prior output.
2. **Given** the alpha change, **When** the module is built, **Then** it builds cleanly with dependencies resolved (a dependency tidy step is run if needed).

---

### User Story 3 - Plan decomposes the work across repos in order (Priority: P1)

Planning decomposes this issue into a sub-issue targeting the beta repo, and the beta work lands before the alpha work.

**Why this priority**: This is the behaviour under test. It verifies cross-repo spawn, including on-demand initialisation of the spawn-target repo.

**Independent Test**: After Plan, a sub-issue exists in `handarbeit/fabrik-test-beta` for the beta-side change, and the alpha issue is blocked on it.

**Acceptance Scenarios**:

1. **Given** work that needs ordered changes in two repos, **When** Plan completes, **Then** it declares a spawned child targeting the beta repo.
2. **Given** the child was spawned, **When** the pipeline proceeds, **Then** the beta change lands before the alpha change.

---

### Edge Cases

- The function name and string literal must be used exactly as given. Substituting a different name or literal fails the scenario's final assertion.
- The alpha repo must not be changed to import a function that has not yet landed in beta.
- The spawn-target repo may not yet be initialised for Fabrik. On-demand initialisation (#803) must make the spawn succeed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202610250554429282() string` in its greeting package that returns exactly `e2e-cross-repo-spawn-20261025-055442-9282`.
- **FR-002**: The beta repo MUST include a test for that function in the greeting package's test file.
- **FR-003**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610250554429282()` and print the result, in addition to its existing behaviour.
- **FR-004**: The alpha repo's module dependencies MUST be tidied if the change requires it.
- **FR-005**: The beta-side change MUST land before the alpha-side change.
- **FR-006**: Plan MUST decompose the work by emitting a spawn declaration targeting the beta repo.
- **FR-007**: The function name and string literal MUST be used verbatim as given in this issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains `HelloE2E202610250554429282` and the literal `e2e-cross-repo-spawn-20261025-055442-9282`.
- **SC-002**: The beta greeting package tests pass, including the new test.
- **SC-003**: Running the alpha program prints `e2e-cross-repo-spawn-20261025-055442-9282`.
- **SC-004**: A sub-issue targeting the beta repo was created by Plan's spawn, and the alpha issue was blocked on it until the beta change landed.

## Assumptions

- The beta repo already has a greeting package at the stated location, and the alpha repo already depends on that package.
- The spawn mechanism's repo-servability rules permit spawning into the beta repo from this instance.
- The per-run-unique token in the names and literal is already final and must not be regenerated.

## Out of Scope

- Any change to Fabrik's engine or spawn mechanics. This issue only exercises them.
- Other changes to the greeting packages or `main.go` beyond those listed.
- Cleanup of the fixture repos after the run.
