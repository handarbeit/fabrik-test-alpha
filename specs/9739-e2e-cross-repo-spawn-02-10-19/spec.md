# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-9739`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610260610197278() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261026-061019-7278\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610260610197278() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. It guards two earlier behaviors: on-demand initialization of a spawn-target repository (#803) and the use of a per-run-unique fixture token (#1308).

The work spans two repositories with a strict ordering dependency. The beta repo (`handarbeit/fabrik-test-beta`) must gain a new greeting function first. The alpha repo (`handarbeit/fabrik-test-alpha`, this one) can then consume it. Alpha cannot reference a function that does not yet exist in beta, so the work has to be decomposed. The Plan stage is expected to spawn a child issue in the beta repo for the beta-side change, and the alpha-side change follows once beta has landed.

The function name and string literal embed a per-run-unique token. The scenario's final assertion checks that this exact text landed in the beta repo, so both must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

The beta repo provides an exported greeting function that returns the run-unique string, with a test covering it. This is the prerequisite for everything on the alpha side.

**Why this priority**: Alpha depends on it, and the scenario's final assertion checks that this exact text landed in beta.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes. Calling `HelloE2E202610260610197278()` returns `"e2e-cross-repo-spawn-20261026-061019-7278"`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610260610197278()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261026-061019-7278`.
2. **Given** the beta repo's greeting tests, **When** the suite runs, **Then** a test for the new function exists and passes.

---

### User Story 2 - Alpha repo calls and prints the beta function's result (Priority: P2)

After beta has landed, the alpha repo's main program calls the new beta function and prints the result, in addition to its existing behavior.

**Why this priority**: It consumes the beta change, so it can only be completed and verified after User Story 1 has landed.

**Independent Test**: Build and run the alpha program with the beta change available. Its output includes `e2e-cross-repo-spawn-20261026-061019-7278`, and its prior output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** the alpha program runs, **Then** it prints `e2e-cross-repo-spawn-20261026-061019-7278`.
2. **Given** the alpha program's existing behavior, **When** the change is applied, **Then** that behavior is preserved and the new output is added alongside it.
3. **Given** the alpha module dependencies, **When** the change is complete, **Then** the module builds cleanly, with dependency tidying applied if it was needed.

---

### User Story 3 - Work is decomposed across repos in the correct order (Priority: P1)

The planning stage decomposes the work so that the beta-repo change is a separately spawned sub-issue in the beta repo, and it is completed before the alpha-repo change.

**Why this priority**: Verifying this decomposition is the purpose of the scenario, as a regression check for #803 and #1308.

**Independent Test**: After planning, a child issue exists in `handarbeit/fabrik-test-beta` carrying the beta-side work. The alpha issue is blocked on it.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** a spawn block targeting the beta repo is emitted and a child issue is created in the beta repo.
2. **Given** the beta repo has not yet been initialized for spawning, **When** the child is spawned, **Then** the target repo is initialized on demand (#803).
3. **Given** the child issue and the alpha issue, **When** their order is inspected, **Then** the alpha work does not complete before the beta work has landed.

---

### Edge Cases

- The function name and string literal must be used verbatim. Any variation, such as a different suffix or a regenerated token, breaks the final assertion.
- If alpha were built before beta lands, the missing function would fail the build. The ordering constraint exists to prevent this.
- If the beta target repo has not yet been initialized for spawning, on-demand initialization (#803) must handle it rather than failing the spawn.
- If `go mod tidy` produces no changes, no dependency file changes are required.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202610260610197278() string` in its `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202610260610197278()` MUST return exactly the literal string `e2e-cross-repo-spawn-20261026-061019-7278`.
- **FR-003**: The beta repo MUST include a test for `HelloE2E202610260610197278()` in `pkg/greeting/greeting_test.go`, and it MUST pass.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610260610197278()` and print the result, in addition to its existing behavior.
- **FR-005**: The alpha repo MUST build successfully after the change, with `go mod tidy` run if needed.
- **FR-006**: The beta-repo change MUST be delivered through a child issue spawned in `handarbeit/fabrik-test-beta`, and MUST land before the alpha-repo change.
- **FR-007**: Planning MUST decompose the work with a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo.
- **FR-008**: The function name and string literal MUST be used exactly as given in this spec, with no substitution.

### Key Entities *(if applicable)*

- **Alpha repo**: `handarbeit/fabrik-test-alpha`, this repo, which consumes the new function.
- **Beta repo**: `handarbeit/fabrik-test-beta`, which provides the new function and receives the spawned sub-issue.
- **Run-unique token**: The identifier embedded in the function name and string literal, used by the scenario's final assertion.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The exact text `e2e-cross-repo-spawn-20261026-061019-7278` is present in the beta repo's greeting package, satisfying the scenario's final assertion.
- **SC-002**: A child issue exists in the beta repo that was spawned from this issue's planning.
- **SC-003**: The beta repo's greeting tests pass, including the new test.
- **SC-004**: Running the alpha program prints `e2e-cross-repo-spawn-20261026-061019-7278`, and its prior output is unchanged.
- **SC-005**: The beta change merged before the alpha change.

## Assumptions

- The alpha repo already imports the beta repo's greeting package, or can do so as part of this change. The issue says to update `main.go` to call the function "in addition to whatever it already does".
- The beta repo's `pkg/greeting` package already exists, since the issue says to add to `pkg/greeting/greeting.go`.
- Spawning the beta child issue and ordering the work are handled by the Plan stage's existing decomposition mechanism, not by this spec.
- This is a test fixture, so the new function is not meant for production use.

## Out of Scope *(optional)*

- Changes to the beta repo other than the one new function and its test.
- Changes to alpha's existing behavior beyond adding the new call and its printed output.
- Redesigning or modifying the cross-repo spawn mechanism itself.
- Choosing a different function name or string literal.
