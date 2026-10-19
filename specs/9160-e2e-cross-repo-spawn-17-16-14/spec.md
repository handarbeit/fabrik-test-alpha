# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-9160`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "e2e cross-repo spawn — Goal: Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order. What needs to change — In handarbeit/fabrik-test-beta (the beta repo): Add an exported function HelloE2E202610192116141137() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261019-211614-1137". Add a test for it in pkg/greeting/greeting_test.go. In handarbeit/fabrik-test-alpha (this repo): Update main.go to call greeting.HelloE2E202610192116141137() (in addition to whatever it already does) and print the result. Run go mod tidy if needed. Constraint: Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo. This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It guards two earlier fixes:

- **#803**: on-demand initialisation of the repo that a spawned sub-issue targets.
- **#1308**: a per-run-unique fixture token, so a run's final assertion cannot be satisfied by leftovers from an earlier run.

The work needs changes in two repositories, and they must land in a fixed order. The beta repo (`handarbeit/fabrik-test-beta`) must gain a new exported greeting function first. Only then can the alpha repo (`handarbeit/fabrik-test-alpha`, this repo) call it. Alpha cannot import a function that does not exist yet. That ordering is why the work has to be split across repos, with the beta part as a spawned, blocking sub-issue of this one.

The function name and string literal contain a per-run-unique token. They are fixed by this spec and must be used exactly as written. The scenario's final assertion checks that exactly this text lands in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the unique greeting function (Priority: P1)

The beta repo's greeting package exposes a new exported function that returns the run-unique literal, and a test covers it. The work is carried out through a sub-issue that Plan spawns in the beta repo.

**Why this priority**: The beta change is the prerequisite for everything else. It is also what the scenario's final assertion inspects.

**Independent Test**: In the beta repo, call `HelloE2E202610192116141137()` and confirm it returns `e2e-cross-repo-spawn-20261019-211614-1137`. Run the beta package's tests and confirm they pass.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610192116141137()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261019-211614-1137`.
2. **Given** the beta repo's test suite, **When** it is run, **Then** a test for this function exists and passes.
3. **Given** this issue's Plan stage, **When** it decomposes the work, **Then** it emits a spawn block targeting the beta repo, and a sub-issue is created there and registered as a blocker of this issue.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

After the beta change has landed, the alpha repo's entry point calls the new beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It depends on User Story 1 and completes the cross-repo flow. Without it the scenario proves nothing about ordered, dependent landing.

**Independent Test**: Build and run alpha's entry point and confirm that its output contains `e2e-cross-repo-spawn-20261019-211614-1137` and that its prior output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** alpha's entry point runs, **Then** it calls `greeting.HelloE2E202610192116141137()` and prints the returned string.
2. **Given** the alpha change, **When** the module is built, **Then** it compiles, and dependency metadata is tidied if the build requires it.
3. **Given** the beta sub-issue has not landed, **When** the pipeline reaches alpha's implementation, **Then** alpha's work does not proceed ahead of it.

---

### Edge Cases

- The beta repo may not yet be initialised for spawned work. The spawn must still succeed, which is the #803 behaviour under test.
- A previous run may have left similar changes in the beta repo. The per-run token must make any stale content unable to satisfy this run's assertion (#1308).
- The function name and literal must match exactly. An altered name or literal, even a near miss, fails the scenario.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's `pkg/greeting/greeting.go` MUST define an exported function `HelloE2E202610192116141137() string` that returns exactly `"e2e-cross-repo-spawn-20261019-211614-1137"`.
- **FR-002**: The beta repo's `pkg/greeting/greeting_test.go` MUST contain a test that asserts the return value of `HelloE2E202610192116141137()`.
- **FR-003**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610192116141137()` and print the result, while keeping its existing behaviour.
- **FR-004**: The alpha repo's dependency metadata MUST be tidied if the change requires it, and the module MUST build.
- **FR-005**: The beta change MUST land before the alpha change. The alpha work MUST be blocked on the beta work.
- **FR-006**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets `handarbeit/fabrik-test-beta`, so the beta change is carried out in a sub-issue in that repo.
- **FR-007**: The function name and string literal given in FR-001 MUST be used verbatim. No substitute name or token may be invented.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains `HelloE2E202610192116141137` and the literal `e2e-cross-repo-spawn-20261019-211614-1137`, and its test suite passes.
- **SC-002**: A sub-issue targeting the beta repo exists and is registered as a blocker of this issue. It reached completion before the alpha change landed.
- **SC-003**: The alpha repo builds, and its entry point's output includes `e2e-cross-repo-spawn-20261019-211614-1137`.
- **SC-004**: The scenario's final assertion finds the exact token text in the beta repo, and no other run's token satisfies it.

## Assumptions

- `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` already exist in the beta repo. The alpha repo already imports the beta greeting package, or can add the import with a dependency tidy.
- The alpha entry point already prints output, and that output stays as is.
- The Fabrik instance serving this issue's board covers the beta repo, so a cross-repo spawn is servable.
- The `fabrik:yolo` label governs auto-advance for this issue, so no human approval step is expected between stages.

## Out of Scope *(optional)*

- Any change to the beta or alpha repos beyond the function, its test and the call site described here.
- Changes to the greeting package's existing functions.
- Changes to Fabrik's engine, spawn logic or the e2e harness. This issue only exercises them.
- Registry, versioning or release work for the beta module.
