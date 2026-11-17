# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-11468`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202611170345251556() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261117-034525-1556\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202611170345251556() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. It protects two earlier behaviours: on-demand initialisation of a spawn-target repository (#803) and the per-run-unique fixture token (#1308).

The scenario needs a piece of work that spans two repositories with a hard ordering dependency. The beta repo (`handarbeit/fabrik-test-beta`) must gain a new greeting function first. Only then can the alpha repo (`handarbeit/fabrik-test-alpha`, this repo) call it. Because of that dependency, the planning step is expected to split the work and spawn a sub-issue in the beta repo, rather than handling everything in this repo.

The function name and string literal contain a per-run-unique token. The scenario's final assertion checks that this exact text landed in the beta repo, so it must be used exactly as given.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

The beta repo's greeting package gains an exported function that returns the per-run-unique literal, together with a test covering it. This is delivered through a sub-issue spawned in the beta repo.

**Why this priority**: Alpha cannot compile against a function that does not exist. This is the prerequisite for everything else, and the scenario's final assertion checks for this text in the beta repo.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes. Calling `HelloE2E202611170345251556()` returns exactly `e2e-cross-repo-spawn-20261117-034525-1556`.

**Acceptance Scenarios**:

1. **Given** the beta repo without the function, **When** the beta work lands, **Then** `pkg/greeting/greeting.go` exports `HelloE2E202611170345251556() string` returning `"e2e-cross-repo-spawn-20261117-034525-1556"`.
2. **Given** the function exists, **When** the beta repo's tests run, **Then** a test in `pkg/greeting/greeting_test.go` asserts that exact return value and passes.

---

### User Story 2 - Alpha repo consumes the new function (Priority: P2)

After the beta change has landed, the alpha repo's `main.go` calls the new function and prints its result. Its existing behaviour is retained.

**Why this priority**: This completes the cross-repo flow, but it depends on User Story 1 landing first.

**Independent Test**: Build and run the alpha program once the beta function is available. Its output includes `e2e-cross-repo-spawn-20261117-034525-1556`, and the module dependencies resolve cleanly.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** the alpha program runs, **Then** it prints the string returned by `greeting.HelloE2E202611170345251556()`, in addition to any output it already produced.
2. **Given** the dependency change in alpha, **When** module tidying is needed, **Then** it is run so the build succeeds.

---

### User Story 3 - Planning decomposes the work across repos in order (Priority: P1)

The planning step recognises that the work spans two repos with a strict ordering. It decomposes the work by spawning a sub-issue targeting the beta repo, and the alpha work is sequenced after it.

**Why this priority**: This is the behaviour the e2e scenario exists to verify (#803, #1308).

**Independent Test**: After Plan, a sub-issue exists in `handarbeit/fabrik-test-beta` carrying the beta requirements, and the alpha issue is blocked by it.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan runs, **Then** it emits a spawn block targeting the beta repo describing the beta change.
2. **Given** the spawned beta sub-issue, **When** the alpha work is sequenced, **Then** alpha does not proceed to completion before beta has landed.

---

### Edge Cases

- The function name and literal must match the issue text character for character. The token is already baked in, so a regenerated or "improved" name or value must not be substituted.
- The beta repo may not yet be initialised for spawn use. On-demand initialisation of the spawn-target repo (#803) is part of what this scenario exercises.
- If the alpha change is attempted before the beta function exists, the alpha build fails. This is why the ordering constraint exists.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's `pkg/greeting/greeting.go` MUST export a function `HelloE2E202611170345251556() string` that returns exactly `"e2e-cross-repo-spawn-20261117-034525-1556"`.
- **FR-002**: The beta repo's `pkg/greeting/greeting_test.go` MUST contain a test verifying the return value of `HelloE2E202611170345251556()`.
- **FR-003**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202611170345251556()` and print the result, while retaining its existing behaviour.
- **FR-004**: The alpha repo's module dependencies MUST be tidied (`go mod tidy`) if needed so that it builds with the new call.
- **FR-005**: The beta work MUST land before the alpha work completes. The alpha work MUST NOT be considered done until the beta function is available to import.
- **FR-006**: Planning MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta` for the beta changes (FR-001, FR-002).
- **FR-007**: The exact function name and string literal given in this spec MUST be used verbatim; no alternative name or value may be invented.

### Key Entities *(if applicable)*

- **Alpha repo** (`handarbeit/fabrik-test-alpha`): the repo this issue lives in. It consumes the new function.
- **Beta repo** (`handarbeit/fabrik-test-beta`): the spawn-target repo that provides the new function.
- **Per-run-unique token**: the `20261117-034525-1556` identifier embedded in the function name and literal, checked by the scenario's final assertion.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A sub-issue targeting the beta repo is created by the decomposition, and the alpha issue is sequenced after it.
- **SC-002**: The exact text `e2e-cross-repo-spawn-20261117-034525-1556` is present in the beta repo's `pkg/greeting/greeting.go` once the beta work lands.
- **SC-003**: The beta repo's tests pass, including the new test, and the alpha program builds and prints the returned string.
- **SC-004**: The alpha change is never merged before the beta change has landed.

## Assumptions

- The beta repo already has a `pkg/greeting` package; this spec names the existing file paths in that repo.
- The alpha repo already imports the beta greeting package, or the import is added as part of FR-003/FR-004. Whichever applies is for later stages to determine.
- "Landed" means merged into the beta repo's default branch.
- The date in the token (`20261117`) is part of the opaque per-run-unique value, not a date to be recomputed.

## Out of Scope *(optional)*

- Any other changes to either repo's greeting behaviour or `main.go` output beyond what is listed above.
- Changes to the Fabrik engine itself. This issue only exercises the existing cross-repo spawn behaviour.
- Defining the implementation or design of the function beyond its signature and return value.

## Source References *(optional)*

- #803 — on-demand spawn-target repo init
- #1308 — per-run-unique fixture token
