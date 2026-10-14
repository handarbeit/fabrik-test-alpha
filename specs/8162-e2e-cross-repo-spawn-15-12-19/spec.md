# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-8162`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202610141912195737() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261014-191219-5737".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202610141912195737() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).
>
> ## Labels
>
> fabrik:yolo

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. Work that spans two repositories with a strict ordering dependency must be split so that the dependency (the beta repo) lands before the dependent change (the alpha repo, which imports from beta). The scenario exercises two previously fixed behaviors:

- **#803**: on-demand initialization of the spawn-target repo.
- **#1308**: a per-run-unique fixture token, so the final assertion can tell this run's output apart from any other run's.

The exact function name and string literal below carry that per-run-unique token and are the scenario's final assertion. They must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the greeting function (Priority: P1)

The beta repo gains an exported greeting function that returns the run-unique literal, with a test. It is delivered as a spawned sub-issue in the beta repo, because the alpha change depends on it.

**Why this priority**: The alpha change cannot compile or be merged until this exists, so it is the foundation for the whole scenario. It is also what the scenario's final assertion checks.

**Independent Test**: In the beta repo, run the package tests for `pkg/greeting`. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261014-191219-5737`.

**Acceptance Scenarios**:

1. **Given** the beta repo's `pkg/greeting/greeting.go`, **When** the change lands, **Then** it exports `HelloE2E202610141912195737() string` returning exactly `"e2e-cross-repo-spawn-20261014-191219-5737"`.
2. **Given** `pkg/greeting/greeting_test.go`, **When** the beta tests run, **Then** a test covering `HelloE2E202610141912195737` passes.

---

### User Story 2 - Alpha repo calls and prints the beta function (Priority: P2)

After the beta change has landed, the alpha repo's `main.go` calls `greeting.HelloE2E202610141912195737()` in addition to its existing behavior and prints the result.

**Why this priority**: It depends on User Story 1 and completes the cross-repo flow, but it delivers nothing until the beta function exists.

**Independent Test**: Build and run alpha once beta contains the function. The program still does what it did before and also prints `e2e-cross-repo-spawn-20261014-191219-5737`.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** alpha's `main.go` runs, **Then** its prior behavior is preserved and the returned literal is printed.
2. **Given** alpha's dependencies changed, **When** the change is complete, **Then** the module files are consistent (`go mod tidy` applied if needed) and the build succeeds.

---

### User Story 3 - Plan decomposes the work across repos in order (Priority: P1)

The Plan stage recognizes that the work spans two repositories with a strict order. It declares a spawn of a sub-issue targeting the beta repo, and the alpha work is blocked until that sub-issue is done.

**Why this priority**: This is the behavior the scenario exists to verify (cross-repo decomposition, #803 and #1308).

**Independent Test**: After Plan, a sub-issue exists in the beta repo describing the beta change with the exact name and literal. The alpha issue is blocked by it and does not implement until it is resolved.

**Acceptance Scenarios**:

1. **Given** the spec above, **When** Plan runs, **Then** it emits a spawn declaration targeting the beta repo, and a sub-issue is created there.
2. **Given** the spawned beta sub-issue is unresolved, **When** the alpha issue is evaluated, **Then** alpha implementation does not proceed ahead of it.
3. **Given** the beta repo had not yet been prepared as a spawn target, **When** the spawn occurs, **Then** the target is initialized on demand (#803).

---

### Edge Cases

- The function name and literal must be copied exactly, including the `E2E202610141912195737` and `20261014-191219-5737` tokens. Any invented or altered value fails the scenario's final assertion (#1308).
- Alpha must not be merged ahead of beta; if it were, alpha would reference a non-existent function and fail to build.
- Alpha's existing `main.go` behavior must not be removed or altered; the new call is additive.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) exports `HelloE2E202610141912195737() string` in `pkg/greeting/greeting.go`, returning exactly `"e2e-cross-repo-spawn-20261014-191219-5737"`.
- **FR-002**: The beta repo has a test for `HelloE2E202610141912195737` in `pkg/greeting/greeting_test.go` that passes.
- **FR-003**: The alpha repo's `main.go` calls `greeting.HelloE2E202610141912195737()` in addition to its existing behavior and prints the result.
- **FR-004**: If alpha's module dependencies need updating, `go mod tidy` is applied so the module files are consistent.
- **FR-005**: The work is decomposed so the beta change is delivered as a sub-issue in the beta repo, and the alpha change does not land before it.
- **FR-006**: The decomposition is declared by Plan through the engine's sanctioned spawn-child mechanism targeting the beta repo (not by creating issues directly).
- **FR-007**: The exact function name and string literal given in this spec are used verbatim in every repo and artifact; no substitute identifiers or values are introduced.

### Key Entities

- **Alpha repo** (`handarbeit/fabrik-test-alpha`): the repo this issue lives in; consumes the greeting function from its `main.go`.
- **Beta repo** (`handarbeit/fabrik-test-beta`): the spawn-target repo; owns `pkg/greeting`.
- **Spawned sub-issue**: the beta-repo issue carrying the beta change, which blocks the alpha work.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The beta repo contains `HelloE2E202610141912195737` returning exactly `e2e-cross-repo-spawn-20261014-191219-5737`, with a passing test.
- **SC-002**: Running the alpha program prints `e2e-cross-repo-spawn-20261014-191219-5737` and still performs all of its prior behavior.
- **SC-003**: The beta change merges before the alpha change, and no point in the history has alpha referencing a beta function that does not yet exist.
- **SC-004**: A sub-issue targeting the beta repo was created through the spawn mechanism, demonstrating on-demand spawn-target initialization (#803) and the per-run-unique token (#1308).

## Assumptions

- The beta repo already has `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` (or an equivalent package layout) that the new function and test can be added to.
- Alpha already imports the beta `greeting` package or can add that import.
- The `fabrik:yolo` label means the pipeline auto-advances; this does not change the requirements.
- "In addition to whatever it already does" means alpha's existing `main.go` behavior is preserved.

## Out of Scope

- Any change to Fabrik's engine, spawn mechanism or stage configuration.
- Any other changes to the greeting package or to alpha beyond the single function, its test and the single call.
- Releasing or versioning the beta repo beyond what is needed for alpha to consume the new function.

## Source References

- #803 (on-demand spawn-target repo init)
- #1308 (per-run-unique fixture token)
