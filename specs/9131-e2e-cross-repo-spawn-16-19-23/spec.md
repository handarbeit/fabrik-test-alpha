# Feature Specification: E2E cross-repo spawn (beta-first decomposition)

**Feature Branch**: `fabrik/issue-9131`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610192019235408() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261019-201923-5408\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610192019235408() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. The work genuinely spans two repositories with a strict ordering dependency: the alpha repository (`handarbeit/fabrik-test-alpha`) cannot consume a greeting function until the beta repository (`handarbeit/fabrik-test-beta`) has it. The pipeline is expected to recognise that dependency and decompose the work: Plan should spawn a child issue in the beta repo, and the alpha work should wait for the beta work to land.

The scenario guards two earlier fixes: #803 (on-demand initialisation of a spawn-target repo) and #1308 (a per-run-unique fixture token, so a run's final assertion cannot be satisfied by leftovers from an earlier run). The function name and string literal below contain that per-run token. They must be used exactly as given, because the scenario's final assertion checks that this exact text appears in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the new greeting function (Priority: P1)

The beta repo's greeting package exposes a new exported function that returns the per-run-unique literal, with a test covering it. This work is carried out in a child issue spawned into the beta repo.

**Why this priority**: It is the prerequisite for the alpha change and carries the text the scenario's final assertion looks for.

**Independent Test**: In the beta repo, call the new function and confirm it returns the expected literal. Run the beta test suite and confirm the new test passes.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610192019235408()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261019-201923-5408`.
2. **Given** the beta repo's greeting tests, **When** the suite is run, **Then** a test for the new function exists and passes.

---

### User Story 2 - Alpha program uses the beta function (Priority: P2)

After the beta change has landed, the alpha repo's entry point calls the new beta function in addition to its existing behaviour and prints the result.

**Why this priority**: It completes the cross-repo feature, but it depends on Story 1.

**Independent Test**: With the beta change available to alpha, build and run the alpha program and confirm its output includes the literal string.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** the alpha program is run, **Then** its output includes `e2e-cross-repo-spawn-20261019-201923-5408` and its previous output is unchanged.
2. **Given** alpha's dependency manifest, **When** the alpha module is built, **Then** the build succeeds (dependency tidying applied if needed).

---

### User Story 3 - Pipeline decomposes the work across repos in order (Priority: P1)

The Plan stage splits the work into a beta-repo child issue and the alpha-repo remainder, and the alpha work does not proceed to completion before the beta work lands.

**Why this priority**: This is the behaviour under test; Stories 1 and 2 are the observable artifacts of it.

**Independent Test**: Run the issue through the pipeline and confirm a child issue exists in the beta repo, linked as a blocker of this issue, and that the beta change merged before the alpha change.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** a child issue targeting `handarbeit/fabrik-test-beta` is created, with its repo initialised on demand if needed (#803).
2. **Given** the child issue exists, **When** the alpha work would otherwise start, **Then** it is held until the beta child has landed.
3. **Given** the run is complete, **When** the final assertion runs, **Then** the exact function name and literal from this issue are found in the beta repo.

---

### Edge Cases

- The function name or literal is altered or "improved" by an agent: the final assertion fails, so both must be used verbatim.
- The beta repo is not yet prepared as a spawn target: it must be initialised on demand rather than failing the spawn (#803).
- Leftovers from an earlier run exist in the beta repo: the per-run token must make them unable to satisfy this run's assertion (#1308).
- Alpha's `main.go` already prints output: the new call is added alongside it, not instead of it.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's `pkg/greeting/greeting.go` MUST define the exported function `HelloE2E202610192019235408() string` returning exactly `e2e-cross-repo-spawn-20261019-201923-5408`.
- **FR-002**: The beta repo's `pkg/greeting/greeting_test.go` MUST contain a test that asserts the return value of `HelloE2E202610192019235408()`.
- **FR-003**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610192019235408()` in addition to its existing behaviour and print the result.
- **FR-004**: The alpha module's dependency manifest MUST be left consistent (tidied if required) so the alpha repo builds with the new import.
- **FR-005**: The Plan stage MUST decompose the work by emitting a spawn block that targets `handarbeit/fabrik-test-beta` for the beta changes.
- **FR-006**: The beta change MUST land before the alpha change is completed.
- **FR-007**: The function name and string literal MUST be used verbatim as given in this issue; no substitute name or token may be invented.

### Key Entities

- **Parent issue**: This issue (#9131) in the alpha repo; owns the alpha-side change.
- **Child issue**: The issue spawned in the beta repo; owns the beta-side change and blocks the parent.
- **Per-run-unique token**: The `20261019-201923-5408` fragment embedded in the function name and literal, which ties the final assertion to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Exactly one child issue is created in `handarbeit/fabrik-test-beta` for the beta-side work, and it is registered as blocking this issue.
- **SC-002**: After the run, the beta repo's default branch contains the text `HelloE2E202610192019235408` and `e2e-cross-repo-spawn-20261019-201923-5408` verbatim.
- **SC-003**: The beta repo's tests pass, including the new test.
- **SC-004**: The alpha program builds and, when run, prints `e2e-cross-repo-spawn-20261019-201923-5408` alongside its prior output.
- **SC-005**: The beta change merged before the alpha change.

## Assumptions

- The beta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`, as the issue states.
- The alpha repo's `main.go` already imports or can import the beta greeting package; wiring that module dependency is part of the alpha-side work.
- Output of the new call is printed to standard output.
- This is a test-fixture scenario, so no documentation or user-facing behaviour changes are required beyond the code changes described.

## Out of Scope

- Any change to the beta repo other than the new function and its test.
- Any change to the alpha repo other than `main.go` and dependency-manifest tidying.
- Changes to Fabrik's engine, spawn logic, or the e2e harness itself.
- Renaming, refactoring or removing existing greeting functions.
