# Feature Specification: Cross-repo spawn e2e regression scenario

**Feature Branch**: `fabrik/issue-10328`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order. In handarbeit/fabrik-test-beta: add an exported function HelloE2E202610310159365899() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261031-015936-5899", and add a test for it in pkg/greeting/greeting_test.go. In handarbeit/fabrik-test-alpha (this repo): update main.go to call greeting.HelloE2E202610310159365899() (in addition to whatever it already does) and print the result; run go mod tidy if needed. Constraint: beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo. This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is a live end-to-end regression scenario for Fabrik's cross-repo decomposition. The work spans two repositories with a strict ordering dependency. The alpha repo (this one) can only consume a greeting function once the beta repo (`handarbeit/fabrik-test-beta`) provides it. The scenario checks that the pipeline splits the work across repos instead of attempting everything in one repo. It guards two earlier fixes:

- **#803**: on-demand initialisation of a spawn-target repo.
- **#1308**: a per-run-unique fixture token. The function name and string literal embed that token so the scenario's final assertion can confirm the exact text landed in the beta repo.

The labels include `fabrik:yolo`, so the pipeline is expected to advance without manual approval.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the unique greeting function (Priority: P1)

The beta repo gains an exported greeting function that returns the run-unique string, with a test covering it. This is the prerequisite for everything in alpha.

**Why this priority**: Alpha cannot compile against a function that does not exist, and the scenario's final assertion checks this exact text in the beta repo.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes and asserts the exact literal.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610310159365899()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261031-015936-5899`.
2. **Given** the beta repo's greeting tests, **When** the suite runs, **Then** a test for the new function exists and passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

After the beta change has landed, the alpha repo's entry point calls the new greeting function and prints its result, in addition to its existing behaviour.

**Why this priority**: It depends on User Story 1 and demonstrates that the ordering constraint was honoured.

**Independent Test**: Once the beta change is available, build and run alpha. The output contains `e2e-cross-repo-spawn-20261031-015936-5899`, and previous output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** alpha's main program runs, **Then** it prints the returned literal in addition to its prior output.
2. **Given** alpha's module dependencies, **When** the change is made, **Then** the module builds cleanly (dependencies tidied if needed).

---

### User Story 3 - Pipeline decomposes the work across repos (Priority: P1)

The Plan stage splits the work by spawning a child issue targeting the beta repo, so the beta work is done and landed before alpha's work proceeds.

**Why this priority**: This is the behaviour under test (cross-repo spawn and spawn-target repo init).

**Independent Test**: After Plan, a child issue exists in the beta repo and blocks the alpha parent. The parent does not implement against the beta repo's function until the child has landed.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** it has emitted a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta`.
2. **Given** the child issue is created, **When** the beta change is not yet landed, **Then** the alpha parent remains blocked.
3. **Given** the beta change has landed, **When** the alpha parent resumes, **Then** its change proceeds and completes.

---

### Edge Cases

- The function name and string literal must be used verbatim. Any variation breaks the final assertion.
- The beta repo may need on-demand initialisation as a spawn target (#803); the scenario must succeed in that case.
- If alpha has no pending dependency change, no module tidying is needed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's `pkg/greeting/greeting.go` MUST export `HelloE2E202610310159365899() string`, returning exactly `"e2e-cross-repo-spawn-20261031-015936-5899"`.
- **FR-002**: The beta repo's `pkg/greeting/greeting_test.go` MUST include a test asserting the return value of `HelloE2E202610310159365899()`.
- **FR-003**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610310159365899()` and print the result, while retaining its existing behaviour.
- **FR-004**: The alpha module MUST build cleanly with the new dependency usage (run `go mod tidy` if needed).
- **FR-005**: The work MUST be decomposed so the beta change lands before the alpha change. Plan MUST emit a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta`.
- **FR-006**: The function name and string literal MUST be used exactly as given in this spec, never regenerated or altered.

### Key Entities

- **Parent issue (alpha)**: This issue; owns the `main.go` change and is blocked by the child.
- **Child issue (beta)**: Spawned by Plan in `handarbeit/fabrik-test-beta`; owns the greeting function and its test.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains `HelloE2E202610310159365899` returning `e2e-cross-repo-spawn-20261031-015936-5899`, with a passing test.
- **SC-002**: A child issue targeting the beta repo was created by Plan, and the alpha change was not merged before the beta change.
- **SC-003**: Running alpha's program after both changes prints the literal and still produces its prior output.

## Assumptions

- The beta repo already has `pkg/greeting/greeting.go` and `greeting_test.go` (or can create them on demand per #803).
- Alpha's `main.go` already imports or can import the beta greeting package; `go mod tidy` handles dependency resolution if required.
- The alpha change is not required to be merged only after the beta *release*; the beta change merely needs to be available to alpha's build.

## Out of Scope

- Any other change to greeting behaviour in either repo.
- Changes to Fabrik engine code; this issue only exercises it.
- Cleanup of the per-run fixture after the scenario.
