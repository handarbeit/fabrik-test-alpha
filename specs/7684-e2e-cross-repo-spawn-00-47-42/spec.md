# Feature Specification: Cross-repo spawn e2e regression scenario

**Feature Branch**: `fabrik/issue-7684`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610130447429594() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261013-044742-9594\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610130447429594() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is a live end-to-end regression scenario for cross-repo decomposition. It guards two earlier fixes: #803 (on-demand initialisation of a spawn-target repo) and #1308 (a per-run-unique fixture token, so runs cannot be confused with one another).

The work is deliberately split across two repositories with a hard ordering dependency. The beta repo (`handarbeit/fabrik-test-beta`) must gain a new greeting function first. The alpha repo (`handarbeit/fabrik-test-alpha`, where this issue lives) can only consume it afterwards, because alpha cannot import a function that does not yet exist. This forces the Plan stage to decompose the work and spawn a child issue in beta, which is the behaviour under test.

The function name and the returned string both embed a per-run-unique token (`202610130447429594` / `20261013-044742-9594`). The scenario's final assertion checks for that exact text in the beta repo, so both must be used exactly as given.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the unique greeting function (Priority: P1)

A child issue in the beta repo adds an exported greeting function whose name and return value carry the per-run-unique token, with a test.

**Why this priority**: The scenario's final assertion looks for this exact text in the beta repo, and alpha cannot build without it.

**Independent Test**: In the beta repo, run the greeting package's tests; the new test passes and asserts the function returns the exact literal.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610130447429594()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261013-044742-9594`.
2. **Given** the beta repo's test suite, **When** it is run, **Then** a test covering the new function passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

After beta's change has landed, the alpha repo's entry point calls the new beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It demonstrates the dependent half of the cross-repo change, and it can only be done after Story 1.

**Independent Test**: Build and run the alpha program; its output includes `e2e-cross-repo-spawn-20261013-044742-9594` alongside its previous output.

**Acceptance Scenarios**:

1. **Given** beta's change has landed, **When** the alpha program runs, **Then** it prints the string returned by `greeting.HelloE2E202610130447429594()` and its prior output is still present.
2. **Given** the alpha module's dependencies, **When** the module is built, **Then** it builds cleanly, with dependency metadata tidied if that was needed.

---

### User Story 3 - Plan decomposes the work into a beta-targeted child (Priority: P1)

The Plan stage recognises the cross-repo ordering constraint and emits a spawn block targeting the beta repo, so the beta change is its own issue that lands before alpha's.

**Why this priority**: This is the behaviour the regression test exists to verify (#803).

**Independent Test**: After Plan completes, a child issue exists in the beta repo, is linked as a blocker of this issue, and alpha's implementation does not proceed until it is resolved.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan runs, **Then** it emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta` that carries the exact function name and string literal.
2. **Given** the child issue in beta is unresolved, **When** this issue would otherwise proceed to implementation, **Then** it stays blocked until the child is done.

---

### Edge Cases

- The beta repo may not yet be initialised on the engine side. On-demand initialisation of a spawn-target repo (#803) must make it work without manual setup.
- The function name and literal must not be altered, abbreviated or regenerated. Any other value breaks the final assertion.
- The existing behaviour of alpha's entry point must be preserved. The new call and print are additive.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo MUST expose an exported function named exactly `HelloE2E202610130447429594`, taking no arguments and returning a string, in its greeting package.
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261013-044742-9594`.
- **FR-003**: The beta repo MUST include a test in the greeting package's test file that verifies FR-002.
- **FR-004**: The alpha repo's entry point MUST call `greeting.HelloE2E202610130447429594()` from the beta repo's greeting package and print the result, in addition to its existing behaviour.
- **FR-005**: The alpha module's dependency metadata MUST be tidied (`go mod tidy`) if the new import requires it.
- **FR-006**: The beta change MUST land before the alpha change that depends on it.
- **FR-007**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo, using the function name and literal verbatim.
- **FR-008**: The function name and literal given in this issue MUST be used verbatim. They MUST NOT be replaced with other values.

### Key Entities

- **Alpha repo** (`handarbeit/fabrik-test-alpha`): the consumer, which hosts this issue.
- **Beta repo** (`handarbeit/fabrik-test-beta`): the provider of the greeting function, and the target of the spawned child issue.
- **Per-run-unique token**: the value embedded in the function name and the literal that identifies this run (#1308).

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo contains `HelloE2E202610130447429594` and the literal `e2e-cross-repo-spawn-20261013-044742-9594`, and its tests pass.
- **SC-002**: A child issue targeting the beta repo was spawned by Plan and completed before alpha's change was implemented.
- **SC-003**: The alpha program builds and, when run, prints the literal above in addition to its previous output.
- **SC-004**: The scenario completes without manual repo initialisation or operator intervention.

## Assumptions

- The beta repo already has a greeting package with a source file and a test file, as named in the request.
- The alpha repo already depends on, or can be given a dependency on, the beta module.
- The `fabrik:yolo` label means stages advance automatically once complete, so no human approval is expected between them.

## Out of Scope

- Any change to the engine's spawn logic itself. This issue only exercises it.
- Changes to other functions in the greeting package.
- Any generalisation of the token scheme beyond the single value given.

## Source References

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
