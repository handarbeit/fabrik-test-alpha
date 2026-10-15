# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-8364`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "

## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610150633121095() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261015-063312-1095".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610150633121095() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. Delivering the feature requires coordinated changes in two repositories, `handarbeit/fabrik-test-beta` (the library that provides the greeting package) and `handarbeit/fabrik-test-alpha` (the application that consumes it). The alpha change depends on a function that does not exist until the beta change has landed, so the work has a strict order: beta first, then alpha.

The scenario exercises two earlier fixes:

- **#803**: on-demand initialisation of the spawn-target repository.
- **#1308**: a per-run-unique fixture token, so a run's final assertion cannot be satisfied by leftovers from an earlier run.

The function name `HelloE2E202610150633121095` and the string `e2e-cross-repo-spawn-20261015-063312-1095` embed that per-run token. They must be used exactly as written. The scenario's final assertion checks that this exact text landed in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the per-run greeting function (Priority: P1)

The beta repository gains an exported greeting function that returns the per-run-unique literal, together with a test. This work is decomposed into its own sub-issue in the beta repo and lands before any alpha change.

**Why this priority**: Alpha cannot compile against a function that does not exist, so this is the prerequisite for everything else. It is also what the scenario's final assertion inspects.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes, and calling `HelloE2E202610150633121095()` returns exactly `e2e-cross-repo-spawn-20261015-063312-1095`.

**Acceptance Scenarios**:

1. **Given** the beta repo without the function, **When** the beta change is merged, **Then** `pkg/greeting/greeting.go` exports `HelloE2E202610150633121095() string` returning `"e2e-cross-repo-spawn-20261015-063312-1095"`.
2. **Given** the beta change is merged, **When** the beta test suite runs, **Then** a test in `pkg/greeting/greeting_test.go` asserting that return value passes.

---

### User Story 2 - Alpha consumes the new function (Priority: P2)

After the beta change has landed, the alpha application calls the new greeting function and prints the result, in addition to whatever it already does.

**Why this priority**: This is the dependent half of the decomposition. It cannot be built or validated until User Story 1 has landed.

**Independent Test**: With the beta change available, build and run the alpha application. The output includes `e2e-cross-repo-spawn-20261015-063312-1095`, and the module's dependencies are consistent (`go mod tidy` leaves no unexpected diff).

**Acceptance Scenarios**:

1. **Given** the beta function exists, **When** the alpha application runs, **Then** it prints the string returned by `greeting.HelloE2E202610150633121095()`.
2. **Given** the alpha change, **When** the project is built, **Then** the build succeeds and previously existing output of `main.go` is unchanged.

---

### User Story 3 - Cross-repo decomposition with enforced ordering (Priority: P1)

Planning for this issue recognises that the work spans two repositories and decomposes it by spawning a sub-issue targeting `handarbeit/fabrik-test-beta`. The alpha work is held until that sub-issue completes.

**Why this priority**: This is what the e2e scenario verifies. The code changes are only the vehicle for exercising cross-repo spawn.

**Independent Test**: After Plan completes, a child issue exists in the beta repo, is linked as a blocker of this issue, and carries the beta-side requirements. This issue does not proceed to implement the alpha change until the child is resolved.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan identifies the beta-first dependency, **Then** it declares a spawned child targeting `handarbeit/fabrik-test-beta`.
2. **Given** the child issue is created, **When** it is still open, **Then** this issue remains blocked and the alpha change is not implemented.
3. **Given** the child has landed in beta, **When** this issue resumes, **Then** the alpha change is implemented against the merged beta function.

---

### Edge Cases

- The spawn-target repo has not been initialised for use yet (the #803 scenario). Spawning must still succeed on demand.
- A previous run's artifacts exist in either repo. Only the exact per-run token text counts as satisfying this run; the name and literal must not be altered or substituted.
- Alpha is built before beta has landed. This must not happen, because the ordering constraint prevents it.
- `go mod tidy` changes nothing. That is acceptable, since it is only required "if needed".

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202610150633121095() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202610150633121095()` MUST return exactly the literal `"e2e-cross-repo-spawn-20261015-063312-1095"`.
- **FR-003**: The beta repo MUST include a test for `HelloE2E202610150633121095` in `pkg/greeting/greeting_test.go` that verifies the returned literal.
- **FR-004**: The alpha repo (`handarbeit/fabrik-test-alpha`) MUST update `main.go` to call `greeting.HelloE2E202610150633121095()` and print the result, while preserving its existing behaviour.
- **FR-005**: The alpha change MUST NOT land before the beta change has landed.
- **FR-006**: The work MUST be decomposed across repositories. Planning MUST emit a spawn declaration (`FABRIK_SPAWN_CHILD_BEGIN` block) targeting the beta repo for the beta-side work.
- **FR-007**: The function name and string literal MUST be used verbatim as given in this spec. No alternative identifiers or literals may be introduced.
- **FR-008**: Alpha's module dependencies MUST be left consistent (`go mod tidy` run if needed).

### Key Entities

- **Alpha repo** (`handarbeit/fabrik-test-alpha`): the application repo that consumes the greeting function. It is this issue's home repo.
- **Beta repo** (`handarbeit/fabrik-test-beta`): the library repo that owns `pkg/greeting` and is the spawn target.
- **Spawned child issue**: the beta-side work item created by decomposition and linked as a blocker of this issue.
- **Per-run token**: the unique suffix baked into the function name and string literal, which lets the final assertion tell this run's output from any other's.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The beta repo's default branch contains the text `HelloE2E202610150633121095` and `e2e-cross-repo-spawn-20261015-063312-1095`, which is what the scenario's final assertion checks.
- **SC-002**: A child issue targeting the beta repo exists for this issue and is resolved before the alpha change merges.
- **SC-003**: The beta test suite passes, including the new greeting test.
- **SC-004**: The alpha application builds and its output includes `e2e-cross-repo-spawn-20261015-063312-1095`.
- **SC-005**: The sequence completes with no manual intervention.

## Assumptions

- `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` already exist in the beta repo, and alpha already imports the beta greeting package, or can import it once the function exists.
- "In addition to whatever it already does" means existing output and behaviour of alpha's `main.go` are retained. The new line is added alongside it.
- The format of the printed line is not constrained beyond containing the returned string.
- The mechanics of spawning, ordering and cross-repo servability are provided by the existing engine and are not defined here.

## Out of Scope

- Any changes to the engine's spawn logic itself. The scenario only exercises it.
- Additional greeting functions or refactoring of the existing greeting package.
- Changes to repos other than alpha and beta.
- Release or versioning work in the beta repo beyond what alpha needs in order to import the function.

## Source References

- #803: on-demand spawn-target repo initialisation.
- #1308: per-run-unique fixture token.
