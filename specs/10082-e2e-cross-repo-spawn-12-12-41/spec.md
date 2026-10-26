# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-10082`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202610261612416366() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261026-161241-6366".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202610261612416366() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. It checks that when one piece of work needs changes in two repositories, in a strict order, the pipeline splits it into a sub-issue in the second repository instead of attempting everything in one repo. It guards two earlier fixes:

- #803: on-demand initialization of the repository that a spawned sub-issue targets.
- #1308: the per-run-unique fixture token. The function name and string literal below carry that token, so the scenario's final assertion can confirm that this run's text, and no other run's, landed in the beta repo.

The ordering constraint drives the decomposition. The alpha repo (this one) will call a new greeting function that does not exist until the beta repo provides it, so the beta change must land first.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

A sub-issue in `handarbeit/fabrik-test-beta` adds a new exported greeting function and a test for it. This is the prerequisite for the alpha change.

**Why this priority**: Alpha cannot compile against a function that does not exist, so this work must land first. It is also the work the scenario's final assertion inspects.

**Independent Test**: In the beta repo, run the package tests for `pkg/greeting`. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261026-161241-6366`.

**Acceptance Scenarios**:

1. **Given** the beta repo without the new function, **When** the sub-issue's work is merged, **Then** `pkg/greeting/greeting.go` exports `HelloE2E202610261612416366() string` returning the literal `"e2e-cross-repo-spawn-20261026-161241-6366"`.
2. **Given** the merged beta change, **When** the `pkg/greeting` tests run, **Then** a test in `pkg/greeting/greeting_test.go` covers the new function and passes.

---

### User Story 2 - Alpha repo uses the new function (Priority: P2)

After the beta change has landed, the alpha repo's `main.go` calls `greeting.HelloE2E202610261612416366()` and prints the result, in addition to its existing behavior.

**Why this priority**: It depends on User Story 1 and completes the cross-repo change.

**Independent Test**: With the beta function available, build and run the alpha program. It compiles, still does what it did before, and also prints `e2e-cross-repo-spawn-20261026-161241-6366`.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** alpha is built, **Then** the build succeeds and the dependency files are consistent (`go mod tidy` has been run if it was needed).
2. **Given** the updated alpha program, **When** it is run, **Then** its prior output is unchanged and the additional line contains `e2e-cross-repo-spawn-20261026-161241-6366`.

---

### User Story 3 - Planning decomposes the work across repositories (Priority: P1)

During planning, the pipeline recognizes that the work spans both repositories in strict order and spawns a sub-issue targeting `handarbeit/fabrik-test-beta` through the spawn-child mechanism (`FABRIK_SPAWN_CHILD_BEGIN` block). The alpha parent is blocked on that sub-issue.

**Why this priority**: This is the behavior under test. Without the spawn, the scenario has not exercised cross-repo decomposition at all.

**Independent Test**: After Plan runs, a sub-issue exists in the beta repo with the beta requirements, and the alpha issue is linked as blocked by it.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan completes, **Then** a sub-issue exists in `handarbeit/fabrik-test-beta` whose requirements use the exact function name and literal given in this spec.
2. **Given** the spawned sub-issue is not yet done, **When** the alpha issue's implementation would otherwise begin, **Then** the alpha work waits until the beta sub-issue has landed.

---

### Edge Cases

- The spawn target repo has not been initialized on demand before (#803): the spawn must still succeed.
- A different run's token must never appear in the work: only the exact identifiers in this spec are valid (#1308).
- The alpha change must not land before the beta change, or alpha will not compile.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202610261612416366() string` in `pkg/greeting/greeting.go`.
- **FR-002**: That function MUST return exactly the string `e2e-cross-repo-spawn-20261026-161241-6366`.
- **FR-003**: The beta repo MUST include a test for that function in `pkg/greeting/greeting_test.go`, and it MUST pass.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610261612416366()` and print the result, in addition to its existing behavior.
- **FR-005**: The alpha repo MUST build successfully, with module dependency files tidied if needed.
- **FR-006**: The beta change MUST land before the alpha change.
- **FR-007**: Planning MUST decompose the work by spawning a sub-issue in `handarbeit/fabrik-test-beta` for the beta changes, rather than handling them inside the alpha issue.
- **FR-008**: The function name and string literal MUST be used exactly as given in this spec. They MUST NOT be altered, regenerated or replaced with other values.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A sub-issue targeting the beta repo is created by the pipeline from this issue, without manual creation.
- **SC-002**: The beta repo's default branch contains the text `e2e-cross-repo-spawn-20261026-161241-6366` and the function name `HelloE2E202610261612416366` after the beta sub-issue lands.
- **SC-003**: The alpha repo's program output includes `e2e-cross-repo-spawn-20261026-161241-6366` once both changes have landed.
- **SC-004**: Both repositories' tests and builds pass after both changes have landed.

## Assumptions

- The existing `pkg/greeting` package in the beta repo and the existing `main.go` in the alpha repo are the intended places for the changes, as the request states.
- The alpha repo already depends on the beta repo's `greeting` package, so only a `go mod tidy` or version refresh may be needed.
- The identifiers in this spec are already unique to this run, so no further uniqueness handling is required.

## Out of Scope *(optional)*

- Any change to the greeting package beyond the single new function and its test.
- Any change to alpha's existing behavior other than the one added call and print.
- Changes to the pipeline's spawn mechanism itself. This scenario only exercises it.

## Source References *(optional)*

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
