# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-7929`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202610140208027520() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261014-020802-7520".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202610140208027520() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This is an end-to-end regression scenario for cross-repo decomposition. It exercises two behaviours: on-demand initialisation of a spawn-target repository (#803) and a per-run-unique fixture token (#1308).

The work needs changes in two repositories, and they have a strict order. The beta repository (`handarbeit/fabrik-test-beta`) must gain a new greeting function first. The alpha repository (`handarbeit/fabrik-test-alpha`, this one) can then call it. Alpha cannot compile against a function that does not yet exist in beta. The pipeline is therefore expected to split the work: a sub-issue in the beta repo lands first, and the alpha change follows once it is done.

The function name and string literal contain a per-run-unique token. The scenario's final assertion checks that this exact text appears in the beta repo, so the values must be used verbatim and never re-derived or altered.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

A sub-issue targeting `handarbeit/fabrik-test-beta` adds an exported greeting function and a test for it, and that work lands before any alpha change.

**Why this priority**: The alpha change depends on it. The scenario's final assertion also inspects the beta repo for the exact token.

**Independent Test**: In the beta repo, the greeting package exposes `HelloE2E202610140208027520()` returning exactly `e2e-cross-repo-spawn-20261014-020802-7520`, and its test passes.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610140208027520()` is called, **Then** it returns the string `e2e-cross-repo-spawn-20261014-020802-7520`.
2. **Given** the beta repo's greeting tests, **When** the test suite runs, **Then** a test covers the new function and passes.

---

### User Story 2 - Plan decomposes the work across repos (Priority: P1)

The Plan stage for this alpha issue recognises that the beta change must come first. It spawns a child issue in the beta repo, using a spawn block that targets that repo, instead of planning everything in alpha.

**Why this priority**: Cross-repo decomposition is the behaviour under test.

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-beta` and is a blocking dependency of this alpha issue.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** a spawn block targeting `handarbeit/fabrik-test-beta` has been emitted and a child issue has been created there.
2. **Given** the child issue exists, **When** the alpha work is scheduled, **Then** alpha implementation does not proceed until the beta child has landed.

---

### User Story 3 - Alpha calls the new function (Priority: P1)

After the beta change has landed, the alpha repo's entry point calls the new greeting function and prints its result, in addition to its existing behaviour.

**Why this priority**: It completes the end-to-end flow and proves the dependency ordering worked.

**Independent Test**: Run the alpha program and confirm its output includes `e2e-cross-repo-spawn-20261014-020802-7520` alongside its prior output.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** the alpha program runs, **Then** it prints `e2e-cross-repo-spawn-20261014-020802-7520`.
2. **Given** the alpha program's previous output, **When** the change is applied, **Then** that previous output is still produced.
3. **Given** alpha's dependency metadata may need updating, **When** the change is applied, **Then** `go mod tidy` has been run if needed and the build succeeds.

---

### Edge Cases

- If the alpha change were attempted before beta lands, alpha would fail to build. The ordering constraint exists to prevent this.
- The spawn target repo may not yet be initialised for spawning. Fabrik's on-demand init (#803) must handle this without manual setup.
- The token must be identical in the function name and the literal everywhere it appears. Any alteration breaks the final assertion.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202610140208027520() string` in `pkg/greeting/greeting.go`.
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261014-020802-7520`.
- **FR-003**: The beta repo MUST include a test for the function in `pkg/greeting/greeting_test.go`.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610140208027520()` and print the result, while keeping its existing behaviour.
- **FR-005**: `go mod tidy` MUST be run in alpha if needed so the alpha module builds.
- **FR-006**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta` for the beta change.
- **FR-007**: The beta change MUST land before the alpha change is implemented or merged.
- **FR-008**: The function name and string literal MUST be used verbatim, with no alternative names or values invented, in the spawned child issue and in all code.

### Key Entities

- **Alpha repo** (`handarbeit/fabrik-test-alpha`): the repo of this issue, containing the consumer entry point.
- **Beta repo** (`handarbeit/fabrik-test-beta`): the spawn target, containing the greeting package that gains the new function.
- **Spawned child issue**: the beta-repo sub-issue that carries the beta change and blocks this issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The beta repo contains the exact text `HelloE2E202610140208027520` and `e2e-cross-repo-spawn-20261014-020802-7520` on its default branch, so the scenario's final assertion passes.
- **SC-002**: A child issue is created in the beta repo by Plan's spawn, with no manual repo setup, exercising on-demand spawn-target init.
- **SC-003**: The beta change lands strictly before the alpha change.
- **SC-004**: The alpha program's output includes the beta-provided string, and both repos' tests and builds pass.

## Assumptions

- The per-run-unique token is already fixed in the given function name and literal. It is not regenerated at any stage.
- "Whatever it already does" in `main.go` means its current output is preserved unchanged.
- The test in beta only needs to assert the returned literal.
- Only one child issue is needed, for the single beta change.

## Out of Scope

- Any change to the beta repo beyond the single function and its test.
- Any change to the alpha repo beyond the `main.go` call and the dependency tidy.
- Changes to Fabrik's spawn engine itself. This scenario only verifies existing behaviour (#803, #1308).
- Additional greeting functions or generalising the fixture.
