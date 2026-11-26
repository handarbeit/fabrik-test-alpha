# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-11794`
**Created**: 2026-11-25
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202611260445593091() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261126-044559-3091\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202611260445593091() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. A single piece of work spans two repositories with a strict ordering dependency: the beta repository (`handarbeit/fabrik-test-beta`) must gain a new exported greeting function before the alpha repository (`handarbeit/fabrik-test-alpha`, this repo) can call it. The pipeline's Plan stage is expected to recognise this and decompose the work by spawning a child issue in the beta repository, rather than trying to do everything in one repo.

The scenario regression-tests two earlier fixes:
- #803 — on-demand initialisation of a spawn-target repository.
- #1308 — a per-run-unique fixture token, so the scenario's final assertion can identify this run's output unambiguously.

The function name and string literal below embed that per-run-unique token. They are fixed inputs, not design choices, and the scenario's final assertion depends on that exact text landing in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

The beta repo gains an exported function that returns a fixed, run-unique string, with a test covering it.

**Why this priority**: Alpha depends on it. It must land first, and the scenario's final assertion checks for this exact text in the beta repo.

**Independent Test**: In the beta repo, run the package tests for the greeting package; the new test passes and the function returns the exact literal.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202611260445593091()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261126-044559-3091`.
2. **Given** the beta repo, **When** its test suite runs, **Then** the new test for this function passes.

---

### User Story 2 - Alpha repo calls the beta function and prints the result (Priority: P2)

The alpha repo's program calls the new beta function in addition to its existing behaviour and prints the returned value.

**Why this priority**: It completes the cross-repo change, but cannot be built until User Story 1 has landed.

**Independent Test**: After the beta change has landed, build and run the alpha program; its output contains `e2e-cross-repo-spawn-20261126-044559-3091` and its previous output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** alpha's program runs, **Then** it prints the string returned by `greeting.HelloE2E202611260445593091()`.
2. **Given** the alpha change, **When** the module is built, **Then** it compiles and module dependencies are consistent (tidied if needed).

---

### User Story 3 - Plan decomposes the work across repos in order (Priority: P1)

The Plan stage splits the work into a beta child issue and an alpha parent change, with beta ordered first, by emitting a spawn block that targets the beta repo.

**Why this priority**: This is the behaviour under test — cross-repo decomposition with correct ordering.

**Independent Test**: After Plan completes, a child issue exists in the beta repo describing User Story 1, the alpha issue is blocked by it, and the alpha work proceeds only after the child is done.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan runs, **Then** it emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta`.
2. **Given** the child issue is created, **When** the pipeline continues, **Then** the alpha issue does not proceed to implementation until the beta child has landed.
3. **Given** the child's description, **When** it is read, **Then** it names the function `HelloE2E202611260445593091` and the string `e2e-cross-repo-spawn-20261126-044559-3091` exactly.

---

### Edge Cases

- The beta repo may not yet be initialised as a spawn target (the #803 scenario); on-demand initialisation must make the spawn succeed.
- The function name and string literal must be used exactly as given. Any alternative token breaks the scenario's final assertion.
- The alpha change must not be attempted before the beta function exists, since alpha cannot import a function that does not yet exist.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The work MUST be decomposed so the beta repo change lands before the alpha repo change, with Plan emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta`.
- **FR-002**: In `handarbeit/fabrik-test-beta`, an exported function named exactly `HelloE2E202611260445593091` taking no arguments and returning a `string` MUST be added to the greeting package (`pkg/greeting/greeting.go`).
- **FR-003**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261126-044559-3091`.
- **FR-004**: A test for that function MUST be added to the beta repo's greeting package tests (`pkg/greeting/greeting_test.go`) and pass.
- **FR-005**: In `handarbeit/fabrik-test-alpha`, `main.go` MUST call `greeting.HelloE2E202611260445593091()` and print the result, in addition to its existing behaviour.
- **FR-006**: Alpha's module dependencies MUST be tidied (`go mod tidy`) if the change requires it.
- **FR-007**: The function name and string literal MUST be used verbatim as given in this issue; no substitute token may be invented.

### Key Entities *(if applicable)*

- **Beta repo** (`handarbeit/fabrik-test-beta`): the dependency repository that must receive the new function first.
- **Alpha repo** (`handarbeit/fabrik-test-alpha`): this repository, the consumer that calls the new function.
- **Spawned child issue**: the beta-repo issue created by Plan to carry the beta change.
- **Run-unique token**: the value embedded in the function name and string literal that the scenario's final assertion checks for.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue targeting the beta repo is created by Plan's spawn block, and the alpha issue is blocked on it.
- **SC-002**: The beta repo's default branch contains the function `HelloE2E202611260445593091` returning exactly `e2e-cross-repo-spawn-20261126-044559-3091`, verified by its passing test.
- **SC-003**: Running the alpha program after both changes prints `e2e-cross-repo-spawn-20261126-044559-3091`.
- **SC-004**: The alpha change lands only after the beta change has landed.

## Assumptions

- The beta repo already has a greeting package with a source file and a test file at the stated locations, and is reachable by the pipeline's identity.
- The alpha repo already depends on the beta module, so no new module relationship has to be invented.
- "In addition to whatever it already does" means alpha's existing output is preserved and the new line is added.
- The exact output format of the print (for example, a trailing newline) is at the implementer's discretion, provided the literal appears in the output.

## Out of Scope *(optional)*

- Any change to the beta repo other than the new function and its test.
- Any change to alpha other than the call and print in `main.go` (plus module tidying if needed).
- Changes to the spawn, board or pipeline machinery itself; this issue only exercises it.

## Source References *(optional)*

- #803 — on-demand spawn-target repo initialisation (regression covered).
- #1308 — per-run-unique fixture token (regression covered).
