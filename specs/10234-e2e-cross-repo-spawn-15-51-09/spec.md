# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-10234`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610301951096404() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261030-195109-6404\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610301951096404() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is a live end-to-end regression scenario for Fabrik's cross-repo decomposition. It guards two earlier fixes: on-demand initialisation of a spawn-target repository (#803) and the per-run-unique fixture token (#1308).

The work genuinely spans two repositories with a hard ordering. The beta repo (`handarbeit/fabrik-test-beta`) must gain a new greeting function first. The alpha repo (`handarbeit/fabrik-test-alpha`, where this issue lives) can only call that function after it exists. This forces the pipeline to decompose the work into a sub-issue in beta that blocks the parent issue in alpha.

The function name and string literal contain a per-run-unique token. The scenario's final assertion checks for that exact text in the beta repo, so they must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repo exposes a new exported greeting function that returns the per-run-unique string, with a test, delivered through a sub-issue spawned from this issue.

**Why this priority**: Alpha depends on it, and the scenario's final assertion checks that this text lands in beta. Without it the scenario fails.

**Independent Test**: In the beta repo, the new function exists in the greeting package, returns the exact literal, and its test passes.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan decomposes the work, **Then** a sub-issue targeting `handarbeit/fabrik-test-beta` is created and this issue is blocked by it.
2. **Given** the beta sub-issue is implemented and merged, **When** the greeting package is inspected, **Then** it contains an exported `HelloE2E202610301951096404() string` returning exactly `e2e-cross-repo-spawn-20261030-195109-6404`.
3. **Given** the new function, **When** the beta repo's tests run, **Then** the added test in the greeting package's test file passes.

---

### User Story 2 - Alpha consumes the new function (Priority: P2)

After beta has landed, the alpha repo's entry point calls the new function and prints its result, in addition to its existing behaviour.

**Why this priority**: It proves the ordered cross-repo dependency resolved. It cannot be done until User Story 1 is merged.

**Independent Test**: Build and run alpha's entry point and check that its output includes `e2e-cross-repo-spawn-20261030-195109-6404`, with earlier behaviour intact.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** alpha's `main.go` runs, **Then** it calls `greeting.HelloE2E202610301951096404()` and prints the result alongside its previous output.
2. **Given** the change, **When** alpha is built, **Then** the build succeeds and module files are tidy (`go mod tidy` run if needed).

---

### Edge Cases

- The beta sub-issue must not be created in the alpha repo; it must target the beta repo.
- Alpha must not be worked on before the beta change has landed.
- The function name and literal must match the given text exactly, including the embedded run token. Any substitution fails the final assertion.
- The beta repo may need on-demand initialisation as a spawn target (#803).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Plan decomposes this issue by spawning a sub-issue that targets `handarbeit/fabrik-test-beta`, using the engine's child-spawn mechanism.
- **FR-002**: The beta repo's `pkg/greeting/greeting.go` exports `HelloE2E202610301951096404() string` returning exactly `"e2e-cross-repo-spawn-20261030-195109-6404"`.
- **FR-003**: The beta repo's `pkg/greeting/greeting_test.go` contains a test for that function asserting the exact returned literal.
- **FR-004**: The beta change lands before any alpha change that depends on it (this issue is blocked by the beta sub-issue).
- **FR-005**: The alpha repo's `main.go` calls `greeting.HelloE2E202610301951096404()` and prints the result, in addition to its existing behaviour.
- **FR-006**: The alpha repo's module files are tidy after the change (`go mod tidy` run if needed).
- **FR-007**: The exact function name and string literal given in this spec are used verbatim; no alternative names or tokens are introduced.

### Key Entities

- **Parent issue (alpha)**: This issue in `handarbeit/fabrik-test-alpha`, which consumes the new function.
- **Child issue (beta)**: The spawned sub-issue in `handarbeit/fabrik-test-beta`, which adds the function and its test.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A sub-issue exists in `handarbeit/fabrik-test-beta` spawned from this issue, and this issue is blocked by it until it completes.
- **SC-002**: The beta repo's default branch contains the exact text `HelloE2E202610301951096404` and `e2e-cross-repo-spawn-20261030-195109-6404` in the greeting package after the run.
- **SC-003**: The beta repo's tests pass, including the new one.
- **SC-004**: Alpha builds and its entry point output includes `e2e-cross-repo-spawn-20261030-195109-6404`, with its prior output preserved.
- **SC-005**: Beta's change is merged before alpha's change is merged.

## Assumptions

- The beta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`, or the spawn-target initialisation makes them available.
- Alpha already depends on the beta module, or gains the dependency as part of this change.
- The existing behaviour of alpha's `main.go` is kept as is; the new call is purely additive.
- This issue is run in yolo mode (`fabrik:yolo`), so stages advance without waiting for human approval.

## Out of Scope *(optional)*

- Any change to Fabrik engine behaviour. This issue only exercises existing cross-repo spawn behaviour.
- Other greeting functions or refactoring of the greeting package.
- Work in repositories other than alpha and beta.
