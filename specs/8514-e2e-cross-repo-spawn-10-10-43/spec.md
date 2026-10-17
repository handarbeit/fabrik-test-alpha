# Feature Specification: E2E cross-repo spawn (10:10:43)

**Feature Branch**: `fabrik/issue-8514`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202610171410432373() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261017-141043-2373".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202610171410432373() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This is an end-to-end regression scenario for cross-repo decomposition. It exercises two behaviours: on-demand initialisation of a spawn-target repository (#803) and a per-run-unique fixture token (#1308). The work spans two repositories with a strict ordering dependency. The alpha repo (this one) consumes a new greeting function that must first exist in the beta repo (`handarbeit/fabrik-test-beta`). Beta's change therefore has to land first, and the pipeline is expected to decompose the work by spawning a child issue in beta from the Plan stage. The function name and string literal embed a per-run-unique token. The scenario's final assertion checks that this exact text lands in the beta repo, so both must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the greeting function (Priority: P1)

The beta repo gains a new exported greeting function carrying the run-unique token, with a test, so that dependants can import it.

**Why this priority**: Alpha's change cannot compile until this exists, and the scenario's final assertion inspects this text in beta.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes and asserts the exact literal returned by the function.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610171410432373()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261017-141043-2373`.
2. **Given** the beta repo's greeting tests, **When** the test suite runs, **Then** a test covering this function exists and passes.

---

### User Story 2 - Alpha consumes the beta function (Priority: P2)

The alpha repo's entry point calls the new beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It demonstrates the ordered dependency across repos. It can only be completed after User Story 1 has landed.

**Independent Test**: With beta's change available, build and run alpha. Its output includes the literal `e2e-cross-repo-spawn-20261017-141043-2373`, and existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** beta contains `HelloE2E202610171410432373()`, **When** alpha's `main.go` runs, **Then** it prints the value returned by `greeting.HelloE2E202610171410432373()`.
2. **Given** alpha's existing behaviour, **When** the change is applied, **Then** the existing behaviour and output remain intact.
3. **Given** the dependency graph is updated, **When** the module is tidied, **Then** `go mod tidy` leaves no stale or missing entries.

---

### User Story 3 - Plan decomposes work across repos (Priority: P1)

The Plan stage decomposes the work by spawning a child issue targeting the beta repo, ordered so the child lands before alpha's change.

**Why this priority**: This is the behaviour under regression test (#803, #1308).

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-beta` describing the beta change with the exact function name and literal, and the alpha item is blocked by it until it is resolved.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan runs, **Then** it emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta`.
2. **Given** the spawned child, **When** it is created, **Then** its description uses the function name and string literal verbatim.
3. **Given** the child is unresolved, **When** the alpha item is evaluated, **Then** alpha's implementation does not proceed ahead of beta.

---

### Edge Cases

- The beta repo has not been initialised for spawning yet (#803): the child must still be created on demand.
- A re-run of the scenario uses a different token (#1308): each run's token is baked into the issue text, so the exact text given here must be used and no other run's text.
- The function name or literal is altered, shortened or "corrected" by any stage: the final assertion fails. Verbatim use is required.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202610171410432373() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202610171410432373()` MUST return exactly the string `e2e-cross-repo-spawn-20261017-141043-2373`.
- **FR-003**: The beta repo MUST have a test for this function in `pkg/greeting/greeting_test.go` that asserts the exact returned literal.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610171410432373()` and print the result, while preserving its existing behaviour.
- **FR-005**: `go mod tidy` MUST be run in alpha if needed, so that module files are consistent with the new import.
- **FR-006**: The beta change MUST land before the alpha change that depends on it.
- **FR-007**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo.
- **FR-008**: The function name and string literal given in this spec MUST be used verbatim in every artifact (child issue, code, tests); no substitute token may be invented.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue is created in `handarbeit/fabrik-test-beta` as a result of Plan's spawn block.
- **SC-002**: The beta repo's default branch contains `HelloE2E202610171410432373` returning `e2e-cross-repo-spawn-20261017-141043-2373`, with a passing test.
- **SC-003**: Alpha builds and its run output includes `e2e-cross-repo-spawn-20261017-141043-2373`.
- **SC-004**: The scenario's final assertion that the exact token text exists in the beta repo passes.

## Assumptions

- The beta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`.
- Alpha already imports the beta module's `greeting` package, or adding the import is part of the `go mod tidy` step.
- The exact function name and literal in this issue are the per-run-unique token for this run and are authoritative.

## Out of Scope *(optional)*

- Changes to the engine's spawn or on-demand repo-init logic (#803, #1308); this issue only exercises them.
- Any changes in either repo beyond the function, its test and alpha's call site.
- Changing the format of the per-run-unique token.
