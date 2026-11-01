# Feature Specification: E2E cross-repo spawn regression scenario

**Feature Branch**: `fabrik/issue-10674`
**Created**: 2026-11-01
**Status**: Specified
**Input**: User description:

```text
## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202611010617064675() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261101-061706-4675".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202611010617064675() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).
```

## Background

This issue is a live end-to-end scenario that verifies cross-repo decomposition works. The requested change spans two repositories with a strict ordering dependency: the beta repository (`handarbeit/fabrik-test-beta`) must gain a new greeting function first, because the alpha repository (`handarbeit/fabrik-test-alpha`, where this issue lives) cannot import a function that does not yet exist.

Because the work cannot be done in a single repository, the pipeline is expected to decompose it, creating a sub-issue in the beta repo for the beta-side change and leaving the alpha-side change in this issue, blocked until beta lands.

The scenario is a regression test for two earlier pieces of work:
- #803 — on-demand initialisation of a spawn-target repository.
- #1308 — a per-run-unique fixture token, so each run's final assertion can tell its own landing apart from earlier runs. The function name and string literal below already contain that token and must be used exactly as given.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the new greeting function (Priority: P1)

The beta repository exposes a new exported greeting function that returns the run-unique literal, with a test covering it. This is the change the final assertion of the scenario checks for.

**Why this priority**: The alpha-side change depends on it, and the scenario's final assertion looks for this exact text in the beta repo.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261101-061706-4675`.

**Acceptance Scenarios**:

1. **Given** the beta repo without the new function, **When** the beta-side work completes and lands, **Then** the greeting package exports `HelloE2E202611010617064675() string`, which returns `"e2e-cross-repo-spawn-20261101-061706-4675"`.
2. **Given** the new function exists, **When** the greeting package tests run, **Then** a test for it exists and passes.

---

### User Story 2 - Alpha repo consumes the new function (Priority: P2)

After the beta change has landed, the alpha application calls the new greeting function and prints its result alongside its existing behaviour.

**Why this priority**: It depends on User Story 1 and cannot be built or verified before it lands.

**Independent Test**: With beta's change available to alpha, build and run alpha. The output includes `e2e-cross-repo-spawn-20261101-061706-4675`, and the existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** beta's new function has landed, **When** alpha's `main.go` is updated, **Then** it calls `greeting.HelloE2E202611010617064675()` and prints the returned value in addition to its existing output.
2. **Given** the dependency on beta must be resolved, **When** the module files are checked, **Then** `go mod tidy` has been run if needed and the build succeeds.

---

### User Story 3 - Pipeline decomposes the work across repos in order (Priority: P1)

The pipeline recognises that the work spans two repos and splits it. A sub-issue is created in the beta repo for the beta-side change, and the alpha-side work waits for it.

**Why this priority**: Exercising this decomposition is the purpose of the scenario, and it covers the #803 and #1308 regressions.

**Independent Test**: After the Plan stage, a sub-issue exists in `handarbeit/fabrik-test-beta` describing the beta-side change, and this issue is blocked on it.

**Acceptance Scenarios**:

1. **Given** this issue's spec names changes in two repos with a strict order, **When** the Plan stage runs, **Then** it emits a child-spawn block targeting the beta repo, which results in a sub-issue there.
2. **Given** the spawn target repo may not yet be initialised, **When** the sub-issue is created, **Then** the repo is initialised on demand (#803) and the sub-issue is created successfully.
3. **Given** the sub-issue and this issue both exist, **When** the beta sub-issue has not landed, **Then** the alpha-side work does not proceed past the dependency.

---

### Edge Cases

- The function name and string literal must be used verbatim. Any variation (different token, casing, or spelling) breaks the scenario's final assertion.
- The alpha change must not land before the beta change, since alpha would fail to build.
- Reusing a token from an earlier run would make the final assertion pass falsely (#1308). The token in this issue is unique to this run.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202611010617064675() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202611010617064675()` MUST return exactly the string `e2e-cross-repo-spawn-20261101-061706-4675`.
- **FR-003**: The beta repo MUST include a test for that function in `pkg/greeting/greeting_test.go`, and the test MUST pass.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202611010617064675()` and print the result, in addition to the output it already produces.
- **FR-005**: The alpha repo's module files MUST be tidied (`go mod tidy`) if the change requires it, and the alpha repo MUST build.
- **FR-006**: The beta-side change MUST land before the alpha-side change.
- **FR-007**: The Plan stage MUST decompose the work by emitting a child-spawn block that targets the beta repo, so that the beta-side change is tracked as a sub-issue in that repo.
- **FR-008**: The function name and string literal given in this spec MUST be used verbatim, with no substitute token.

### Key Entities

- **Alpha repo** (`handarbeit/fabrik-test-alpha`): the repository holding this issue and the consuming application (`main.go`).
- **Beta repo** (`handarbeit/fabrik-test-beta`): the repository holding the greeting package that must gain the new function.
- **Run-unique token** (`20261101-061706-4675`): the identifier baked into the function name and literal, which lets the scenario's final assertion recognise this run's landing.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the scenario completes, the beta repo contains `HelloE2E202611010617064675` returning `e2e-cross-repo-spawn-20261101-061706-4675`, and its greeting tests pass.
- **SC-002**: A sub-issue targeting the beta repo was created by the pipeline from this issue's decomposition, with no manual creation.
- **SC-003**: The beta-side change merged before the alpha-side change.
- **SC-004**: The alpha application builds and its output includes `e2e-cross-repo-spawn-20261101-061706-4675` in addition to its prior output.

## Assumptions

- The beta repo already has `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`, or on-demand initialisation (#803) provides what is needed.
- The alpha repo's `main.go` already imports or can import the beta greeting package, and `go mod tidy` is enough to resolve the dependency once beta's change is available.
- The `fabrik:yolo` label on this issue means the stages advance automatically, so no human approval of the spec is awaited.
- The text in this issue is correct as given; the per-run-unique token was generated by the scenario harness.

## Out of Scope *(optional)*

- Changes to the engine's spawn, decomposition or repo-initialisation behaviour. This scenario only verifies the existing behaviour.
- Changes to the beta repo beyond the single function and its test.
- Changes to the alpha repo's existing behaviour beyond printing the additional value.

## Source References *(optional)*

- #803 — on-demand spawn-target repo init.
- #1308 — per-run-unique fixture token.
