# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-9664`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610260456182349() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261026-045618-2349\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610260456182349() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for cross-repo decomposition. It exercises two behaviours: on-demand initialisation of a spawn-target repository (#803) and a per-run-unique fixture token (#1308).

The feature spans two repositories with a strict ordering dependency. The alpha repo (`handarbeit/fabrik-test-alpha`, this repo) consumes a new greeting function that must first exist in the beta repo (`handarbeit/fabrik-test-beta`). The work therefore has to be decomposed across repos, with the beta change landing before the alpha change. The scenario passes only if the exact function name and string literal below end up in the beta repo, so they must be used verbatim. A per-run-unique token is already baked into both.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the unique greeting function (Priority: P1)

A new exported greeting function carrying the per-run-unique token is added to the beta repo, with a test, as a separate unit of work spawned from this issue.

**Why this priority**: Alpha depends on it. Without it, nothing else can be built or verified, and the scenario's final assertion checks this text in the beta repo.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261026-045618-2349`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610260456182349()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261026-045618-2349`.
2. **Given** the beta repo's greeting tests, **When** the suite runs, **Then** a test covering the new function passes.

---

### User Story 2 - Alpha repo prints the beta greeting (Priority: P2)

After the beta change has landed, the alpha program calls the new beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It completes the end-to-end flow, but it can only be built after Story 1 has landed.

**Independent Test**: Build and run the alpha program. Its output includes `e2e-cross-repo-spawn-20261026-045618-2349`, and its previous output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** the alpha program runs, **Then** it prints `e2e-cross-repo-spawn-20261026-045618-2349` alongside its existing output.
2. **Given** the alpha module's dependencies, **When** it is built, **Then** the build succeeds with dependencies tidy.

---

### User Story 3 - Decomposition across repos is exercised (Priority: P1)

Planning decomposes the work into a spawned sub-issue targeting the beta repo, ordered before the alpha change. This demonstrates cross-repo spawn, including on-demand initialisation of the target repo.

**Why this priority**: Verifying this decomposition is the purpose of the scenario.

**Independent Test**: After planning, a sub-issue exists in the beta repo for the beta change. The alpha work is blocked on it until it lands.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan decomposes the work, **Then** a child issue is created in `handarbeit/fabrik-test-beta` covering the beta change.
2. **Given** the child exists, **When** pipeline ordering is applied, **Then** the alpha change does not proceed to merge before the beta change has landed.

---

### Edge Cases

- The function name and string literal must be used exactly as given; any variation fails the scenario's final assertion.
- The alpha change must not be attempted before the beta function exists, because the import would not compile.
- The beta repo may need on-demand initialisation as a spawn target (#803); this must not block the spawn.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's greeting package MUST export `HelloE2E202610260456182349() string`.
- **FR-002**: `HelloE2E202610260456182349()` MUST return exactly `e2e-cross-repo-spawn-20261026-045618-2349`.
- **FR-003**: The beta repo MUST include a test for `HelloE2E202610260456182349()` in the greeting package's tests.
- **FR-004**: The alpha program MUST call `greeting.HelloE2E202610260456182349()` and print the result, in addition to its existing behaviour.
- **FR-005**: The alpha module's dependencies MUST be left tidy so the program builds.
- **FR-006**: The beta change MUST land before the alpha change.
- **FR-007**: The work MUST be decomposed so that the beta change is carried out as a spawned child issue in `handarbeit/fabrik-test-beta`.
- **FR-008**: The function name and string literal MUST be used verbatim as given, with no substitution.

### Key Entities

- **Alpha repo**: `handarbeit/fabrik-test-alpha`, the repo of this issue, which consumes the greeting function.
- **Beta repo**: `handarbeit/fabrik-test-beta`, the repo that owns the greeting package and defines the function.
- **Per-run-unique token**: `20261026-045618-2349`, embedded in both the function name and the returned string.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The beta repo contains the exact text `e2e-cross-repo-spawn-20261026-045618-2349` in the greeting package, returned by the specified function.
- **SC-002**: The beta greeting tests pass, including the new test.
- **SC-003**: The alpha program builds and, when run, prints the exact string above in addition to its previous output.
- **SC-004**: A child issue was created in the beta repo by decomposition, and the beta change merged before the alpha change.

## Assumptions

- The existing greeting package exists in the beta repo at the stated location, and the alpha repo already depends on it.
- The per-run-unique token in the names was generated for this run and should not be changed.
- "In addition to whatever it already does" means existing alpha output is preserved and the new line is added.
- Where the printed result appears relative to existing output is unspecified and not significant.

## Out of Scope

- Changes to any other function in the greeting package.
- Changes to the engine's spawn or initialisation logic; this scenario only verifies existing behaviour (#803, #1308).
- Any additional repos beyond alpha and beta.

## Source References

- Fabrik issue #803: on-demand spawn-target repo initialisation.
- Fabrik issue #1308: per-run-unique fixture token.
