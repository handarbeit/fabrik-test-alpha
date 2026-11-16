# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-11303`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202611161405545760() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261116-140554-5760\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202611161405545760() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. It exercises two behaviours: on-demand initialisation of a spawn-target repository (#803) and a per-run-unique fixture token (#1308).

The work is deliberately split across two repositories with a strict ordering. The beta repo (`handarbeit/fabrik-test-beta`) must gain a new greeting function first. The alpha repo (`handarbeit/fabrik-test-alpha`, this one) can only consume that function once it exists in beta. Because of that dependency, the pipeline is expected to decompose the work by spawning a sub-issue in the beta repo rather than doing everything in this repo's issue.

A per-run-unique token (`20261116-140554-5760`) is already baked into the function name and the returned string. The scenario's final assertion checks for that exact text in the beta repo, so it must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the new greeting function (Priority: P1)

The beta repo exposes a new exported greeting function that returns the run-unique literal. It is covered by a test.

**Why this priority**: The alpha change depends on it, and the scenario's final assertion inspects the beta repo for this exact text.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes and asserts the exact returned literal.

**Acceptance Scenarios**:

1. **Given** the beta repo at its default branch, **When** the change lands, **Then** `pkg/greeting/greeting.go` contains an exported `HelloE2E202611161405545760() string` that returns exactly `e2e-cross-repo-spawn-20261116-140554-5760`.
2. **Given** the change has landed, **When** the greeting package tests run, **Then** a test in `pkg/greeting/greeting_test.go` for the new function passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

The alpha repo's program calls the new beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It completes the cross-repo flow, but it can only land after User Story 1.

**Independent Test**: With the beta change available, build and run alpha. Its output includes `e2e-cross-repo-spawn-20261116-140554-5760` alongside its previous output.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** alpha's `main.go` runs, **Then** it prints the result of `greeting.HelloE2E202611161405545760()`.
2. **Given** alpha's existing output, **When** the change lands, **Then** that existing output is still produced.
3. **Given** the dependency on beta changed, **When** module dependencies are tidied where needed, **Then** alpha builds cleanly.

---

### User Story 3 - Cross-repo decomposition into an ordered sub-issue (Priority: P1)

The Plan stage decomposes the work by spawning a sub-issue in the beta repo, and that sub-issue lands before the alpha change.

**Why this priority**: This is the behaviour the scenario exists to verify (#803, #1308).

**Independent Test**: After the Plan stage runs, a sub-issue exists in the beta repo describing the beta-side work, and the alpha issue is ordered after it.

**Acceptance Scenarios**:

1. **Given** the work needs changes in both repos, **When** Plan completes, **Then** a sub-issue targeting `handarbeit/fabrik-test-beta` has been spawned for the beta-side changes.
2. **Given** the sub-issue exists, **When** the pipeline proceeds, **Then** the beta change lands before the alpha change is implemented against it.
3. **Given** the spawn target repo may not yet be initialised for Fabrik, **When** the spawn happens, **Then** the target repo is initialised on demand (#803).

---

### Edge Cases

- The function name and string literal must be used exactly as given. Any invented or altered token breaks the final assertion.
- If alpha is implemented before beta has landed, alpha cannot compile. The ordering constraint exists to prevent this.
- Running `go mod tidy` is required only if the dependency change needs it.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo MUST export `HelloE2E202611161405545760() string` from `pkg/greeting/greeting.go`.
- **FR-002**: That function MUST return exactly the string `e2e-cross-repo-spawn-20261116-140554-5760`.
- **FR-003**: The beta repo MUST include a test for the function in `pkg/greeting/greeting_test.go`.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202611161405545760()` and print the result, in addition to its existing behaviour.
- **FR-005**: Alpha's dependencies MUST be tidied (`go mod tidy`) if needed so the module builds.
- **FR-006**: The beta-side work MUST land before the alpha-side work.
- **FR-007**: Plan MUST decompose the work by emitting a spawn block that targets the beta repo, so the beta-side work becomes a sub-issue there.
- **FR-008**: The function name and string literal MUST be used verbatim as given in this issue; none may be invented or substituted.

### Key Entities *(if applicable)*

- **Alpha repo**: `handarbeit/fabrik-test-alpha`, the repo this issue lives in and the consumer of the new function.
- **Beta repo**: `handarbeit/fabrik-test-beta`, the spawn target that owns the new function.
- **Run-unique token**: `20261116-140554-5760`, baked into the function name and the returned literal, used by the scenario's final assertion.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains the exact literal `e2e-cross-repo-spawn-20261116-140554-5760` and the function `HelloE2E202611161405545760`.
- **SC-002**: The beta greeting package tests pass, including the new test.
- **SC-003**: The alpha program builds and its output contains the literal `e2e-cross-repo-spawn-20261116-140554-5760`.
- **SC-004**: A sub-issue targeting the beta repo was created by decomposition, and the beta change merged before the alpha change.

## Assumptions

- `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` already exist in the beta repo, and the alpha repo already imports (or can import) the beta greeting package.
- "Print the result" means writing the returned string to standard output on its own line, in addition to existing output.
- The fixture token is final and must not be regenerated.

## Out of Scope

- Any change to the beta repo beyond the one function and its test.
- Any change to alpha beyond `main.go` and dependency tidying.
- Changes to the engine's spawn or on-demand initialisation behaviour. This scenario only verifies it.

## Source References

- #803: on-demand spawn-target repo init (mentioned for context).
- #1308: per-run-unique fixture token (mentioned for context).
