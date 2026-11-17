# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-11555`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202611170521188784() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261117-052118-8784\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202611170521188784() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for cross-repo decomposition. Work that spans two repositories must be split so each repo's part is delivered in the right order. Here the alpha repo (`handarbeit/fabrik-test-alpha`) consumes a new function that must first exist in the beta repo (`handarbeit/fabrik-test-beta`), so the beta change has to land first.

The scenario guards two behaviours:
- **#803**: on-demand initialisation of the spawn-target repo.
- **#1308**: a per-run-unique fixture token. The function name `HelloE2E202611170521188784` and the string `e2e-cross-repo-spawn-20261117-052118-8784` are unique to this run, and the scenario's final assertion checks that exact text in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repo exposes a new exported function that returns the run-unique literal, with a test covering it.

**Why this priority**: Alpha's change depends on it, and the scenario's final assertion checks this text in the beta repo.

**Independent Test**: In the beta repo, run the tests for the greeting package. The new test passes, and the function returns the exact literal.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202611170521188784()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261117-052118-8784`.
2. **Given** the beta repo's greeting test file, **When** the package tests run, **Then** a test for the new function exists and passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

The alpha repo's entry point calls the new beta function, in addition to its existing behaviour, and prints the result.

**Why this priority**: It is the consumer of the beta change and can only be built once beta's function exists.

**Independent Test**: Build and run alpha's main program with the beta change available. Its output includes `e2e-cross-repo-spawn-20261117-052118-8784` and its existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** alpha's main program runs, **Then** it prints the value returned by `greeting.HelloE2E202611170521188784()` alongside its prior output.
2. **Given** alpha's module dependencies, **When** the build runs, **Then** it succeeds, with `go mod tidy` run if needed.

---

### User Story 3 - Plan decomposes the work across repos in order (Priority: P1)

Planning recognises that the work spans two repos with a strict ordering and decomposes it by spawning a beta-repo sub-issue rather than treating it as single-repo work.

**Why this priority**: Verifying cross-repo decomposition is the stated goal of the issue.

**Independent Test**: After Plan runs, a sub-issue exists in the beta repo describing the beta change, and the alpha issue is blocked on it.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan runs, **Then** it emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta` for the beta work.
2. **Given** the spawned beta sub-issue, **When** both issues progress, **Then** the beta change lands before the alpha change is completed.

---

### Edge Cases

- The beta repo may need on-demand initialisation when the sub-issue is spawned (#803). That must succeed without manual setup.
- The function name and literal must be used verbatim. Any alteration breaks the scenario's final assertion.
- If alpha is built before the beta change is available, the import fails. The ordering constraint exists to prevent this.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's `pkg/greeting/greeting.go` MUST export `HelloE2E202611170521188784() string`, returning exactly `"e2e-cross-repo-spawn-20261117-052118-8784"`.
- **FR-002**: The beta repo's `pkg/greeting/greeting_test.go` MUST contain a test for `HelloE2E202611170521188784` that asserts the exact returned literal.
- **FR-003**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202611170521188784()` in addition to its existing behaviour and print the result.
- **FR-004**: The alpha repo's module files MUST be tidy (`go mod tidy` run if needed) so the program builds.
- **FR-005**: The beta change MUST land before the alpha change completes, and Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets `handarbeit/fabrik-test-beta`.
- **FR-006**: The function name and string literal MUST be used exactly as given, with no substitution or invented alternatives.

### Key Entities *(if applicable)*

- **Spawned beta sub-issue**: The child issue in `handarbeit/fabrik-test-beta` carrying the beta-side work. The alpha issue is blocked on it.
- **Run-unique token**: The identifier embedded in the function name and literal, so the final assertion can tell this run's change from any other.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo contains the function `HelloE2E202611170521188784` returning exactly `e2e-cross-repo-spawn-20261117-052118-8784`.
- **SC-002**: The beta repo's greeting package tests pass, including the new test.
- **SC-003**: The alpha repo builds and its program output includes the literal `e2e-cross-repo-spawn-20261117-052118-8784` in addition to its prior output.
- **SC-004**: A sub-issue targeting the beta repo was created by decomposition, and the beta change landed before the alpha change.

## Assumptions

- The beta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`; if the files are missing, the beta work creates them.
- Alpha's `main.go` already imports or can import the beta greeting package, and `go mod tidy` handles any dependency update.
- The existing behaviour of alpha's `main.go` is preserved; the new call is purely additive.

## Out of Scope *(optional)*

- Any change to the beta or alpha repos beyond the function, its test, and the alpha call-and-print.
- Changes to the engine's spawn or decomposition mechanism itself; this issue only exercises it.

## Source References *(optional)*

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
