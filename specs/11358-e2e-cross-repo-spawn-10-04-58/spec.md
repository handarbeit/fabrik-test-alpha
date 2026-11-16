# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-11358`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202611161504582916() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261116-150458-2916\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202611161504582916() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. It guards two behaviours: on-demand initialisation of a spawn-target repository (#803) and the per-run-unique fixture token (#1308). The work spans two repositories with a strict ordering dependency. The alpha repo (this one) cannot use a greeting function until the beta repo (`handarbeit/fabrik-test-beta`) provides it. So the pipeline must decompose the work: the beta-side change becomes a spawned child issue, and the alpha-side change waits on it. The scenario's final assertion checks that the exact per-run token text landed in the beta repo, so the identifiers and literal below must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the per-run greeting function (Priority: P1)

The beta repo's greeting package gains a new exported function that returns the per-run-unique token string, with a test covering it. This work is carried out by a child issue spawned in the beta repo.

**Why this priority**: Alpha depends on it, and the scenario's final assertion inspects the beta repo for this exact text.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261116-150458-2916`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202611161504582916()` is called, **Then** it returns the literal `e2e-cross-repo-spawn-20261116-150458-2916`.
2. **Given** the new function, **When** the beta repo's greeting tests run, **Then** a test for that function exists and passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

After the beta change has landed, this repo's entry point calls the new greeting function and prints its result, in addition to its existing behaviour.

**Why this priority**: It depends on Story 1 and proves the ordered cross-repo dependency resolved correctly.

**Independent Test**: Build and run the alpha program. Its output includes `e2e-cross-repo-spawn-20261116-150458-2916` alongside its previous output.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** alpha's `main.go` calls `greeting.HelloE2E202611161504582916()` and prints the result, **Then** the program builds and the token appears in its output.
2. **Given** the module dependencies change, **When** `go mod tidy` is needed, **Then** it is run so the module files are consistent.

---

### User Story 3 - Plan decomposes the work across repos in order (Priority: P1)

The Plan stage recognises that the work crosses two repos with a strict ordering constraint. It decomposes it by spawning a child issue targeting the beta repo, so that beta lands before alpha.

**Why this priority**: This is the behaviour under test, namely cross-repo decomposition, and the spawn-target repo initialisation from #803.

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-beta` describing the beta change, and the alpha issue is blocked on it.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan runs, **Then** it emits a spawn block targeting the beta repo for the beta-side work.
2. **Given** the spawned child is not yet complete, **When** the alpha issue would proceed to its own implementation, **Then** it waits until the child has landed.

---

### Edge Cases

- The spawn-target repo has not been initialised for the engine yet (the #803 case). It must be initialised on demand so the child can be created and placed.
- The function name or literal is altered or "improved" by a later stage. This fails the scenario, because the final assertion matches the exact text.
- Alpha implementation starts before beta has landed. The import would fail to compile, which is the ordering failure this scenario guards against.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202611161504582916() string` in `pkg/greeting/greeting.go`.
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261116-150458-2916`.
- **FR-003**: The beta repo MUST contain a test for that function in `pkg/greeting/greeting_test.go`.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202611161504582916()` and print the result, while keeping its existing behaviour.
- **FR-005**: `go mod tidy` MUST be run in the alpha repo if the module files require it.
- **FR-006**: The beta-side work MUST be decomposed into a child issue in the beta repo via the Plan stage's spawn mechanism. It MUST NOT be done by creating the issue out of band.
- **FR-007**: The beta change MUST land before the alpha change is implemented or merged.
- **FR-008**: The function name and the string literal MUST be used verbatim in every stage and repo; none may be invented, altered or re-derived.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains `HelloE2E202611161504582916` returning `e2e-cross-repo-spawn-20261116-150458-2916` and a passing test for it.
- **SC-002**: After the run, the alpha program's output includes `e2e-cross-repo-spawn-20261116-150458-2916`.
- **SC-003**: Exactly one child issue was spawned in the beta repo for this decomposition, and the alpha issue did not proceed to implementation until that child was complete.
- **SC-004**: The scenario's final assertion for the exact token text in the beta repo passes.

## Assumptions

- The per-run token embedded in the function name and literal is already unique to this run and must not be changed.
- The beta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go` (the function and its test are added alongside existing code).
- The `fabrik:yolo` label governs auto-advancement; no human approval step is expected for this spec.

## Out of Scope

- Changes to any other function or file in either repo beyond what is listed.
- Altering how the engine implements cross-repo spawning or repo initialisation; this scenario only verifies it.
- Designing the implementation, which belongs to the Plan stage.

## Source References

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
