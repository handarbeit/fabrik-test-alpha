# Feature Specification: Cross-repo spawn end-to-end regression (alpha + beta)

**Feature Branch**: `fabrik/issue-8983`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: 

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202610191207441871() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261019-120744-1871".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202610191207441871() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This issue is a live end-to-end regression scenario for Fabrik's cross-repo decomposition. It guards two earlier fixes: #803 (on-demand initialisation of a spawn-target repo) and #1308 (a per-run-unique fixture token, so each run's assertion can tell its own output from stale output of earlier runs).

The work is split across two repos with a hard ordering dependency. The alpha repo (this one) calls a function that does not exist until the beta repo (`handarbeit/fabrik-test-beta`) provides it. The pipeline therefore has to decompose the work: a child issue in beta lands first, then alpha's change lands. The function name and string literal contain a per-run token and are fixed by the scenario. The final assertion depends on that exact text appearing in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repo exposes a new exported greeting function that returns the run-unique token string, with a test covering it. This is delivered through a child issue spawned from this one.

**Why this priority**: Alpha depends on it, and the scenario's final assertion checks that this text lands in beta.

**Independent Test**: In the beta repo, the greeting package's tests pass, and calling `HelloE2E202610191207441871()` returns exactly `e2e-cross-repo-spawn-20261019-120744-1871`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610191207441871()` is called, **Then** it returns the literal `e2e-cross-repo-spawn-20261019-120744-1871`.
2. **Given** the new test in the beta greeting package's test file, **When** the beta test suite runs, **Then** the test passes and would fail if the returned literal changed.
3. **Given** this issue reaches Plan, **When** Plan decomposes the work, **Then** it emits a child-spawn declaration targeting the beta repo, and a child issue is created there.

---

### User Story 2 - Alpha consumes the beta function (Priority: P2)

The alpha repo's entry point calls the new beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It can only be done after User Story 1 has landed, because the function must exist in beta first.

**Independent Test**: Run alpha's program and see the token string in its output, alongside its prior output.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** alpha's `main.go` runs, **Then** its output includes `e2e-cross-repo-spawn-20261019-120744-1871`.
2. **Given** alpha's existing behaviour, **When** the change is applied, **Then** the prior output is unchanged.
3. **Given** the dependency change, **When** the module is tidied (if needed) and built, **Then** alpha builds and its tests pass.

---

### Edge Cases

- If the beta function does not exist yet when alpha is worked on, alpha must not land its change. Ordering is enforced by decomposition, with the beta child blocking alpha.
- The target repo for the child may not yet be initialised for Fabrik (the #803 case). Spawning must still succeed.
- The token is baked into the name and literal. A different or invented name or literal breaks the scenario's final assertion.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's `pkg/greeting/greeting.go` MUST export `HelloE2E202610191207441871() string`, returning exactly `"e2e-cross-repo-spawn-20261019-120744-1871"`.
- **FR-002**: The beta repo's `pkg/greeting/greeting_test.go` MUST include a test asserting that return value.
- **FR-003**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610191207441871()` and print the result, in addition to its existing behaviour.
- **FR-004**: The alpha repo's module files MUST be updated (`go mod tidy`) if the change requires it.
- **FR-005**: The function name and string literal MUST be used verbatim as given. They must not be altered or replaced with invented values.
- **FR-006**: The work MUST be decomposed so the beta change lands before the alpha change. Plan declares a child-spawn block targeting the beta repo for this.

### Key Entities *(if applicable)*

- **Parent issue (alpha)**: This issue. It owns the alpha-side change and is blocked by its child.
- **Child issue (beta)**: Spawned by Plan in `handarbeit/fabrik-test-beta`. It owns the beta-side function and test.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue is created in the beta repo as a result of Plan's decomposition, and it is linked as a blocker of this issue.
- **SC-002**: After the run, the beta repo contains the exact function name and exact literal given above, with a passing test.
- **SC-003**: After the run, the alpha program's output contains the exact literal, and alpha builds and passes its tests.
- **SC-004**: The beta change is merged before the alpha change is merged.

## Assumptions

- The beta repo already has `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` (or the greeting package can be extended there), and alpha already depends on that module.
- Alpha's `main.go` already imports or can import the beta greeting package.
- The fixture token in the issue is final for this run and is not to be regenerated.

## Out of Scope *(optional)*

- Any change to the greeting package beyond the single new function and its test.
- Changes to alpha's behaviour other than printing the new function's result.
- Changes to Fabrik's engine. This issue only exercises existing spawn behaviour.
