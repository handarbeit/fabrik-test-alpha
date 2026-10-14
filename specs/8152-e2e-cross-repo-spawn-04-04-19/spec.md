# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-8152`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610140804199005() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261014-080419-9005\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610140804199005() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It covers two earlier fixes: on-demand initialisation of a spawn-target repo (#803) and the per-run-unique fixture token (#1308).

The work spans two repositories with a strict ordering dependency. The alpha repo (this one) will call a new greeting function. That function must exist in the beta repo first, because alpha cannot import something that does not yet exist. This ordering forces the Plan stage to split the work. The beta change becomes a child issue in the beta repo, and the alpha change waits on it.

The function name and string literal embed a per-run-unique token. The scenario's final assertion checks that this exact text landed in the beta repo, so both must be used exactly as given.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repo exposes a new exported greeting function that returns the run-unique string, covered by a test. This is the prerequisite for everything in alpha.

**Why this priority**: Alpha depends on it, and the scenario's final assertion looks for this exact text in the beta repo.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes and asserts the function returns the exact literal.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610140804199005()` is called, **Then** it returns `e2e-cross-repo-spawn-20261014-080419-9005`.
2. **Given** the beta repo's greeting test file, **When** the package tests run, **Then** a test covering the new function exists and passes.

---

### User Story 2 - Alpha consumes the beta function (Priority: P2)

After the beta change has landed, the alpha repo's entry point calls the new function and prints the result, in addition to its existing behaviour.

**Why this priority**: It depends on User Story 1 and demonstrates the ordered, cross-repo result.

**Independent Test**: Build and run alpha against a beta version that contains the function. The output includes the exact literal, and the existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** beta contains the new function, **When** alpha's entry point runs, **Then** it prints `e2e-cross-repo-spawn-20261014-080419-9005` in addition to its prior output.
2. **Given** alpha's dependencies, **When** the module is built, **Then** it compiles, with dependency metadata tidied if required.

---

### User Story 3 - Plan decomposes across repos in order (Priority: P1)

The Plan stage for this issue decomposes the work into a child issue in the beta repo. Alpha's work is blocked until that child has landed.

**Why this priority**: This is the behaviour under test (#803, #1308). Without the decomposition the scenario proves nothing.

**Independent Test**: After Plan, a child issue exists in the beta repo for the beta change. This issue is blocked by that child until it is done.

**Acceptance Scenarios**:

1. **Given** the Plan stage runs for this issue, **When** it completes, **Then** it has emitted a spawn declaration targeting the beta repo for the beta-side change.
2. **Given** the child exists in the beta repo, **When** alpha's implementation is considered, **Then** it does not proceed until the beta child is done.
3. **Given** the child's issue text, **When** it is read, **Then** it carries the exact function name and string literal from this spec, with no substitutions.

---

### Edge Cases

- The beta repo may not yet be initialised for spawn targets. On-demand initialisation (#803) must handle this.
- A different function name or literal anywhere in the pipeline breaks the final assertion. Both must be carried verbatim into the child issue, code, and tests.
- If alpha is implemented before beta lands, the build fails on the missing symbol. The ordering must prevent this.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202610140804199005() string` in its `pkg/greeting/greeting.go`.
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261014-080419-9005`.
- **FR-003**: The beta repo MUST include a test for that function in `pkg/greeting/greeting_test.go`.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610140804199005()` and print the result, in addition to its existing behaviour.
- **FR-005**: Alpha's module dependency metadata MUST be tidied (`go mod tidy`) if the change requires it.
- **FR-006**: The beta change MUST land before the alpha change is built or merged.
- **FR-007**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo for the beta-side change.
- **FR-008**: The function name and string literal MUST be used exactly as given in this spec, with no renaming and no regenerated token.

### Key Entities *(if applicable)*

- **Alpha repo**: `handarbeit/fabrik-test-alpha`, the repo that owns this issue and consumes the greeting function.
- **Beta repo**: `handarbeit/fabrik-test-beta`, the repo that provides the greeting function and receives the spawned child issue.
- **Run-unique token**: The identifier embedded in the function name and literal, used by the final assertion.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The beta repo's default branch contains the exact text `e2e-cross-repo-spawn-20261014-080419-9005` and the function `HelloE2E202610140804199005`.
- **SC-002**: The beta repo's greeting tests pass with the new test included.
- **SC-003**: The alpha repo builds successfully and its program output includes the exact literal.
- **SC-004**: A child issue exists in the beta repo, created by the Plan stage's spawn, and it completed before alpha's work did.

## Assumptions

- The beta repo is served by the same Fabrik instance, or can be initialised on demand as spawn-target repo init (#803) provides.
- Alpha's `main.go` already imports or can import the beta greeting package. The Research and Plan stages determine the exact dependency wiring.
- The run-unique token in the given name and literal is fixed for this run and must not be changed.
- Alpha's existing output stays as it is. The new line is purely additive.

## Out of Scope *(optional)*

- Any change to Fabrik's engine or spawn logic. This issue only exercises it.
- Any greeting behaviour in either repo beyond the single new function and its use.
- Release or versioning of the beta repo beyond what alpha needs to consume the function.
