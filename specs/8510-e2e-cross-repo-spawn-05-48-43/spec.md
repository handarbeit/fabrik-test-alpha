# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-8510`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610150948432786() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261015-094843-2786\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610150948432786() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. It exercises two earlier fixes: #803 (on-demand initialisation of a spawn-target repo) and #1308 (a per-run-unique fixture token).

The work spans two repositories with a strict ordering dependency. The beta repo (`handarbeit/fabrik-test-beta`) must gain a new greeting function first. The alpha repo (this one, `handarbeit/fabrik-test-alpha`) can only call it once it exists in beta. Because of that ordering, the pipeline is expected to split the work: the Plan stage should decompose out the beta-repo change as a spawned sub-issue, and the alpha-repo change should depend on it.

The function name and string literal embed a per-run-unique token. The scenario's final assertion checks that this exact text lands in the beta repo, so both must be used exactly as given.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

A consumer of the beta repo's greeting package can call a new exported function and receive the run-unique literal.

**Why this priority**: The alpha change depends on this one, and the scenario's final assertion checks that this text lands in the beta repo.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes and asserts the function returns the literal string.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610150948432786()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261015-094843-2786`.
2. **Given** the beta repo's greeting test file, **When** the package tests run, **Then** a test covering the new function exists and passes.

---

### User Story 2 - Alpha repo uses the beta function (Priority: P2)

The alpha program calls the new beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It depends on User Story 1 and completes the cross-repo flow.

**Independent Test**: Build and run the alpha program after the beta change has landed. The output includes the literal string, and existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** the alpha program runs, **Then** it prints `e2e-cross-repo-spawn-20261015-094843-2786` in addition to its prior output.
2. **Given** the alpha module's dependencies, **When** the module is built, **Then** the build succeeds with dependency metadata tidy.

---

### User Story 3 - Plan decomposes the work across repos (Priority: P1)

The pipeline's Plan stage recognises the strict ordering and spawns a sub-issue targeting the beta repo rather than doing everything in the alpha repo.

**Why this priority**: Verifying this decomposition is the purpose of the scenario (regression coverage for #803 and #1308).

**Independent Test**: After Plan completes, a child issue exists in the beta repo describing the beta change, and this issue is blocked on it.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan runs, **Then** it emits a child-spawn block targeting `handarbeit/fabrik-test-beta` for the beta change.
2. **Given** the spawned beta sub-issue, **When** it is not yet complete, **Then** the alpha change does not proceed ahead of it.

---

### Edge Cases

- The beta repo may need on-demand initialisation as a spawn target (#803); the spawn must still succeed.
- The function name and literal must match exactly. A different token, or any rewording, breaks the final assertion.
- Alpha must not reference the beta function before it exists in the version of beta that alpha depends on.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's `pkg/greeting/greeting.go` MUST contain an exported function `HelloE2E202610150948432786() string` that returns exactly `"e2e-cross-repo-spawn-20261015-094843-2786"`.
- **FR-002**: The beta repo's `pkg/greeting/greeting_test.go` MUST contain a test that asserts the return value of `HelloE2E202610150948432786()`.
- **FR-003**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610150948432786()` and print the result, while keeping its existing behaviour.
- **FR-004**: The alpha module's dependency metadata MUST be tidy (`go mod tidy` run if needed) so the alpha program builds.
- **FR-005**: The beta change MUST land before the alpha change is completed.
- **FR-006**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo.
- **FR-007**: The function name and string literal MUST be used verbatim as given in this spec. No substitute token may be used.

### Key Entities *(if applicable)*

- **Beta repo**: `handarbeit/fabrik-test-beta`, which owns the greeting package.
- **Alpha repo**: `handarbeit/fabrik-test-alpha` (this repo), whose `main.go` consumes the beta function.
- **Spawned sub-issue**: the beta-repo child issue created by Plan for the beta change.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The beta repo contains the literal `e2e-cross-repo-spawn-20261015-094843-2786` returned by `HelloE2E202610150948432786()`, which is the scenario's final assertion.
- **SC-002**: Running the alpha program prints that literal, and its prior output is unchanged.
- **SC-003**: Exactly one sub-issue is created in the beta repo for the beta change, and this issue is ordered after it.
- **SC-004**: The beta package tests and the alpha build both succeed.

## Assumptions

- The beta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`.
- The alpha repo's `main.go` already imports the beta greeting package, or can import it once the function exists.
- The `fabrik:yolo` label means the pipeline advances automatically, so no human approval gate is expected between stages.
- "Print the result" means writing it to standard output, in addition to existing output.

## Out of Scope *(optional)*

- Any change to the beta repo beyond the one function and its test.
- Any change to the alpha repo beyond the `main.go` call and dependency tidying.
- Changes to the Fabrik engine's spawn logic itself. This scenario only verifies it.
