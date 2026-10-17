# Feature Specification: E2E cross-repo spawn

**Feature Branch**: `fabrik/issue-8594`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202610171620438711() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261017-162043-8711".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202610171620438711() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This is an end-to-end regression scenario for cross-repo decomposition. Some work spans two repositories with a strict ordering dependency: the alpha repo (this one) cannot consume a function until the beta repo (`handarbeit/fabrik-test-beta`) provides it. The pipeline must therefore split the work, creating a sub-issue in the beta repo that lands first, and only then complete the alpha-side change.

The scenario regression-tests two earlier issues:
- **#803**: on-demand initialization of the spawn-target repository.
- **#1308**: a per-run-unique fixture token. The function name and string literal below embed such a token so the scenario's final assertion can confirm the exact text landed in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

A sub-issue in the beta repo adds an exported greeting function that returns a fixed, run-unique string, with a unit test.

**Why this priority**: Alpha depends on it, and the scenario's final assertion checks that this exact text lands in the beta repo.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261017-162043-8711`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610171620438711()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261017-162043-8711`.
2. **Given** the beta repo's greeting tests, **When** the test suite runs, **Then** a test covering this function exists and passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

After the beta change has landed, the alpha application calls the new greeting function and prints its result, in addition to its existing behavior.

**Why this priority**: It completes the cross-repo change, but it can only be built once User Story 1 is available.

**Independent Test**: Build and run the alpha application. Its output includes `e2e-cross-repo-spawn-20261017-162043-8711`, and its existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** the alpha application runs, **Then** it prints the string returned by `greeting.HelloE2E202610171620438711()` alongside its prior output.
2. **Given** the alpha module's dependencies, **When** the build runs, **Then** it succeeds, with module metadata tidied if needed.

---

### User Story 3 - Pipeline decomposes the work across repos in order (Priority: P1)

The Plan stage splits the work into a beta-repo sub-issue that must complete first and the alpha-repo change that depends on it.

**Why this priority**: Exercising this decomposition is the purpose of the scenario (regression coverage for #803 and #1308).

**Independent Test**: Run the scenario. A sub-issue targeting `handarbeit/fabrik-test-beta` is created from Plan's spawn block. The alpha issue is blocked until that sub-issue completes.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan runs, **Then** it emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo.
2. **Given** the spawned beta sub-issue is still open, **When** the pipeline proceeds, **Then** the alpha work does not complete before the beta work lands.

---

### Edge Cases

- The function name and string literal must not be altered, regenerated or re-derived. They embed a per-run-unique token, and any variation breaks the scenario's final assertion.
- The alpha change must not remove or alter existing behavior of `main.go`. It only adds a call and prints its result.
- The beta repo may need on-demand initialization as a spawn target (#803). The scenario must still succeed in that case.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST expose an exported function `HelloE2E202610171620438711() string` in `pkg/greeting/greeting.go`.
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261017-162043-8711`.
- **FR-003**: The beta repo MUST include a test for that function in `pkg/greeting/greeting_test.go`.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610171620438711()` and print the result, in addition to its existing behavior.
- **FR-005**: The alpha module's dependency metadata MUST be tidied (`go mod tidy`) if the change requires it.
- **FR-006**: The beta change MUST land before the alpha change is completed.
- **FR-007**: The Plan stage MUST decompose the work with a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo.
- **FR-008**: The exact function name and string literal from FR-001 and FR-002 MUST be used verbatim everywhere they appear. No substitute name or literal may be introduced.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains `HelloE2E202610171620438711` returning `e2e-cross-repo-spawn-20261017-162043-8711`, with a passing test.
- **SC-002**: After the run, the alpha repo's application output includes `e2e-cross-repo-spawn-20261017-162043-8711` and its pre-existing output is unchanged.
- **SC-003**: A sub-issue in the beta repo was created by the pipeline's spawn mechanism, and the alpha work completed only after that sub-issue was done.
- **SC-004**: The scenario's final assertion, that the exact token text exists in the beta repo, passes.

## Assumptions

- `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` already exist in the beta repo, or the spawn-target initialization (#803) makes them available.
- The alpha module can import the beta repo's `greeting` package, and `go mod tidy` is sufficient to wire the dependency.
- The `fabrik:yolo` label means stages auto-advance without manual spec approval.

## Out of Scope *(optional)*

- Any changes to the engine's spawn, decomposition or board logic. This issue verifies existing behavior and does not alter it.
- Greeting functionality beyond the single function and its test.
- Changes in repositories other than alpha and beta.

## Source References *(optional)*

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
