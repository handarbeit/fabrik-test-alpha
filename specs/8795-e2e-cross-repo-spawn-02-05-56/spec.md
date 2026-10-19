# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-8795`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610190605566209() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261019-060556-6209\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610190605566209() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It exercises two behaviours:

- **#803**: on-demand initialisation of the repository a spawned child issue targets.
- **#1308**: a per-run-unique fixture token, so each run's final assertion can tell its own output from output left by earlier runs.

The feature needs a real dependency between two repositories. The alpha repo's entry point must use a new greeting function that lives in the beta repo. Alpha cannot compile against that function until beta has it, so the work cannot be done in one repo or in either order. Beta must land first. That ordering is what should drive the Plan stage to decompose the work into a sub-issue in the beta repo.

The function name and string literal embed a per-run-unique token. They must be used exactly as written. The scenario's final assertion checks that this exact text landed in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the run-unique greeting (Priority: P1)

The beta repo's greeting package gains an exported function that returns the run-unique string, together with a test for it. This is the foundation, and alpha depends on it.

**Why this priority**: Alpha cannot use the function until it exists in beta. The scenario's final assertion also looks for this text in the beta repo.

**Independent Test**: In the beta repo, run the greeting package's tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261019-060556-6209`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610190605566209()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261019-060556-6209`.
2. **Given** the beta repo's greeting test file, **When** the package tests run, **Then** a test for `HelloE2E202610190605566209` is present and passes.

---

### User Story 2 - Alpha calls and prints the beta greeting (Priority: P2)

After beta has landed, the alpha repo's entry point calls the new function and prints its result, in addition to its existing behaviour.

**Why this priority**: This is the consuming half of the change. It can only be done after User Story 1 has landed.

**Independent Test**: Build and run the alpha program once the beta change is available. Its output includes `e2e-cross-repo-spawn-20261019-060556-6209`, and its previous output is unchanged.

**Acceptance Scenarios**:

1. **Given** beta contains `HelloE2E202610190605566209`, **When** alpha's entry point runs, **Then** it prints the returned string `e2e-cross-repo-spawn-20261019-060556-6209`.
2. **Given** the alpha change, **When** the module is built, **Then** the build succeeds and dependency metadata is consistent (`go mod tidy` leaves no diff).
3. **Given** the alpha entry point's existing behaviour, **When** the change is applied, **Then** that behaviour is preserved and the new output is added to it.

---

### User Story 3 - Work is decomposed across repos in landing order (Priority: P1)

The Plan stage splits the work so the beta change is a sub-issue in the beta repo that blocks this alpha issue. Beta therefore lands before alpha.

**Why this priority**: Verifying cross-repo decomposition end to end is the purpose of this issue.

**Independent Test**: After Plan, a child issue exists in the beta repo, is registered as blocking this issue, and carries the beta-side requirements. The alpha work does not start before the child is done.

**Acceptance Scenarios**:

1. **Given** the Plan stage runs for this issue, **When** it decomposes the work, **Then** it emits a spawn block targeting the beta repo.
2. **Given** the child issue exists, **When** the pipeline continues, **Then** this issue stays blocked until the child is resolved.
3. **Given** the child has landed, **When** this issue proceeds, **Then** the alpha change is implemented against the landed beta function.

---

### Edge Cases

- The function name or string must not be altered, re-derived or replaced with a different token. Any difference breaks the final assertion.
- If the beta repo needs on-demand initialisation as a spawn target (#803), the spawn must still succeed.
- If alpha is built before beta has landed, the build fails. The ordering constraint exists to prevent this.
- A leftover artifact from an earlier run must not satisfy this run's assertion, because the token is unique to this run (#1308).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202610190605566209() string` in its greeting package (`pkg/greeting/greeting.go`).
- **FR-002**: `HelloE2E202610190605566209()` MUST return exactly the literal `e2e-cross-repo-spawn-20261019-060556-6209`.
- **FR-003**: The beta repo MUST include a test for `HelloE2E202610190605566209` in `pkg/greeting/greeting_test.go`, and it MUST pass.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610190605566209()` and print the result, in addition to its existing behaviour.
- **FR-005**: The alpha module's dependency metadata MUST be consistent after the change (`go mod tidy` run if needed).
- **FR-006**: The beta change MUST land before the alpha change.
- **FR-007**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo, with the beta-side requirements.
- **FR-008**: The function name and string literal MUST be used verbatim as given in this issue, with no substitute token.

### Key Entities

- **Alpha repo** (`handarbeit/fabrik-test-alpha`): the repo this issue lives in. It consumes the greeting function.
- **Beta repo** (`handarbeit/fabrik-test-beta`): the spawn-target repo. It provides the greeting function.
- **Run-unique token**: the value baked into the function name and returned string, which identifies this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The beta repo contains `HelloE2E202610190605566209` returning `e2e-cross-repo-spawn-20261019-060556-6209`, and the final assertion finds that exact text in the beta repo.
- **SC-002**: The beta repo's greeting package tests pass, including the new test.
- **SC-003**: The alpha program builds and, when run, prints `e2e-cross-repo-spawn-20261019-060556-6209` in addition to its prior output.
- **SC-004**: A child issue targeting the beta repo was created by the Plan stage and resolved before the alpha change merged.

## Assumptions

- The beta repo already has a greeting package at the stated path (or it will be initialised on demand per #803), and alpha already depends on the beta module.
- Alpha's `main.go` already exists and can import the greeting package.
- "In addition to whatever it already does" means existing output is kept as is and the new line is appended.
- The label `fabrik:yolo` is set, so the pipeline advances through stages without waiting for human approval.

## Out of Scope

- Any change to the beta repo beyond the single function and its test.
- Any change to the alpha repo beyond calling and printing the new function, and tidying dependencies.
- Changes to Fabrik's engine, spawn logic or e2e harness. This issue only provides the scenario they are tested against.
- Choosing a different function name or string value.
