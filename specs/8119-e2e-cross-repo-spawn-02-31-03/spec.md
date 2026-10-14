# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-8119`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-beta (the beta repo)
- Add an exported function HelloE2E202610140631034181() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261014-063103-4181".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-alpha (this repo)
- Update main.go to call greeting.HelloE2E202610140631034181() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. The change itself is small, but it is split across two repositories with a hard ordering dependency: the alpha repo (this one) cannot import a function that does not yet exist in the beta repo. That forces the pipeline's Plan stage to decompose the work and spawn a child issue in `handarbeit/fabrik-test-beta` before the alpha change can be implemented.

The scenario guards two earlier fixes:
- **#803** — on-demand initialisation of a spawn-target repo (the beta repo must be usable as a spawn target when first referenced).
- **#1308** — a per-run-unique fixture token. The function name and string literal already contain that token, and the scenario's final assertion checks that exact text in the beta repo. Any invented or altered name or literal would fail the test.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the greeting function (Priority: P1)

The beta repo gains a new exported greeting function returning a fixed, per-run-unique string, with a test covering it. This is the prerequisite for the alpha change and is delivered through a child issue spawned in the beta repo.

**Why this priority**: Alpha cannot compile against a function that does not exist, so this must land first. It is also the text the scenario's final assertion looks for.

**Independent Test**: In the beta repo, run the package tests for `pkg/greeting`; the new test passes, and the function returns the exact literal.

**Acceptance Scenarios**:

1. **Given** the beta repo's `pkg/greeting/greeting.go`, **When** `HelloE2E202610140631034181()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261014-063103-4181`.
2. **Given** `pkg/greeting/greeting_test.go`, **When** the beta repo's tests are run, **Then** a test for `HelloE2E202610140631034181` exists and passes.

---

### User Story 2 - Alpha consumes the beta function (Priority: P2)

After the beta change has landed, the alpha repo's `main.go` calls the new beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It depends on User Story 1 and demonstrates the ordered cross-repo flow end to end.

**Independent Test**: Build and run alpha after the beta change has landed; the output includes the literal returned by the beta function, and existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** alpha's `main.go` runs, **Then** it prints the result of `greeting.HelloE2E202610140631034181()` alongside its existing output.
2. **Given** the alpha module's dependencies, **When** `go mod tidy` is needed to resolve the beta dependency, **Then** it is run so the module builds cleanly.

---

### User Story 3 - Plan decomposes the work across repos (Priority: P1)

The Plan stage recognises that the work spans two repos in strict order and emits a spawn block targeting the beta repo, causing a child issue to be created there and the parent to wait on it.

**Why this priority**: This is the behaviour under test (#803 / #1308 regression); without it the scenario does not exercise cross-repo decomposition at all.

**Independent Test**: After Plan completes, a child issue exists in `handarbeit/fabrik-test-beta` describing the beta change, and the alpha issue is blocked by it.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan runs, **Then** its output includes a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta`.
2. **Given** the spawned child, **When** its description is read, **Then** it carries the exact function name and string literal given in this spec.
3. **Given** the child is unresolved, **When** the alpha issue is evaluated, **Then** alpha implementation does not start until the child completes.

---

### Edge Cases

- The function name and literal must be used verbatim; any normalisation, renaming or re-generated token breaks the final assertion.
- The beta repo may not yet be initialised as a spawn target (the #803 case); the spawn must still succeed.
- Alpha must not be implemented or merged ahead of the beta change.
- Alpha's existing `main.go` behaviour must be preserved; the new call is additive.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202610140631034181() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202610140631034181()` MUST return exactly the string `e2e-cross-repo-spawn-20261014-063103-4181`.
- **FR-003**: The beta repo MUST include a test for `HelloE2E202610140631034181` in `pkg/greeting/greeting_test.go`.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610140631034181()` and print the result, in addition to its existing behaviour.
- **FR-005**: The alpha module MUST build cleanly; `go mod tidy` is run if needed to achieve this.
- **FR-006**: The beta change MUST land before the alpha change is implemented.
- **FR-007**: The Plan stage MUST emit a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo to decompose the work.
- **FR-008**: The exact function name and string literal above MUST be used verbatim in both repos and in the spawned child's description.

### Key Entities

- **Beta repo**: `handarbeit/fabrik-test-beta`, owner of `pkg/greeting`, the spawn target.
- **Alpha repo**: `handarbeit/fabrik-test-alpha`, this repo, the consumer of the beta function.
- **Fixture token**: the per-run-unique value (`20261014-063103-4181`) baked into the function name and literal.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue exists in the beta repo as a result of the Plan stage's spawn, and the parent is blocked on it until it completes.
- **SC-002**: The exact text `HelloE2E202610140631034181` and `e2e-cross-repo-spawn-20261014-063103-4181` are present in the beta repo's default branch after the scenario completes.
- **SC-003**: The beta repo's tests pass, including the new greeting test.
- **SC-004**: The alpha repo builds and its `main.go` output includes the beta function's returned string.
- **SC-005**: The beta change lands strictly before the alpha change.

## Assumptions

- The function name and string literal in the issue are final and must not be altered; the per-run-unique token is already embedded.
- The beta repo already has a `pkg/greeting` package (the issue says to add to existing files); if not, creating it is within the child issue's scope.
- Alpha obtains the beta function through the normal dependency mechanism for the project; how that is wired is a Plan/Implement concern.
- The `fabrik:yolo` label means stages auto-advance without human approval of the spec.

## Out of Scope

- Any greeting functionality beyond the single named function.
- Changes to Fabrik's engine, spawn logic, or e2e harness.
- Other changes to either repo beyond those listed.

## Source References

- #803 — on-demand spawn-target repo init
- #1308 — per-run-unique fixture token
