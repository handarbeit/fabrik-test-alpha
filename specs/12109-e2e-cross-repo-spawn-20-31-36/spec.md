# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-12109`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202611270131363503() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261127-013136-3503".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202611270131363503() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. Some work spans two repositories with a strict ordering dependency: the change in the beta repository (`handarbeit/fabrik-test-beta`) must land before the change in the alpha repository (`handarbeit/fabrik-test-alpha`, this repo), because alpha consumes something beta must first provide. The scenario verifies that the pipeline decomposes such work by spawning a child issue in the beta repo and then completes the dependent alpha work once the child is done.

It regression-tests two earlier behaviors:

- #803 — on-demand initialization of the repo targeted by a spawned child.
- #1308 — a per-run-unique fixture token, so each run's final assertion checks text that could only have come from that run.

The function name and string literal below contain a per-run-unique token and must be used exactly as written. The scenario's final assertion depends on that exact text appearing in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the unique greeting function (Priority: P1)

The beta repo gains an exported function that returns a per-run-unique string, with a test, delivered through a child issue spawned from this one.

**Why this priority**: It is the prerequisite for everything in alpha and the content the scenario's final assertion checks. It also exercises the cross-repo spawn path.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261127-013136-3503`.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan decomposes the work, **Then** it declares a spawned child issue targeting `handarbeit/fabrik-test-beta` for the beta-side change.
2. **Given** the child issue is processed, **When** its work lands, **Then** the beta repo's greeting package contains an exported `HelloE2E202611270131363503() string` returning the literal `e2e-cross-repo-spawn-20261127-013136-3503`.
3. **Given** the function exists in beta, **When** the greeting package's tests run, **Then** a test for that function is present and passes.

---

### User Story 2 - Alpha uses the beta function (Priority: P2)

After the beta change has landed, the alpha repo's program calls the new beta function and prints the result, in addition to whatever it already does.

**Why this priority**: It completes the cross-repo change, but it can only succeed after User Story 1.

**Independent Test**: With the beta change available, build and run the alpha program. Its output includes `e2e-cross-repo-spawn-20261127-013136-3503`, and its previous behavior is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function exists and is available to alpha, **When** the alpha program runs, **Then** it prints the string returned by `greeting.HelloE2E202611270131363503()`.
2. **Given** the alpha change is applied, **When** the module is built, **Then** the module dependency state is consistent (tidied if needed) and the build succeeds.
3. **Given** the alpha program previously produced output, **When** the change is applied, **Then** that existing output is still produced.

---

### Edge Cases

- The alpha change must not be completed before the beta change has landed. Alpha cannot reference a function that does not yet exist.
- The function name and literal must match the given text exactly. A re-typed, reformatted or "similar" token fails the scenario's final assertion.
- The target beta repo may not yet be initialized for spawned work. This is the #803 behavior under test, and spawning must still succeed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The Plan stage MUST decompose this work by declaring a spawned child issue that targets the `handarbeit/fabrik-test-beta` repository (via a `FABRIK_SPAWN_CHILD_BEGIN` block).
- **FR-002**: The beta repo's `pkg/greeting/greeting.go` MUST gain an exported function named exactly `HelloE2E202611270131363503`, taking no arguments and returning a `string`.
- **FR-003**: `HelloE2E202611270131363503()` MUST return exactly the literal `e2e-cross-repo-spawn-20261127-013136-3503`.
- **FR-004**: The beta repo's `pkg/greeting/greeting_test.go` MUST contain a test for `HelloE2E202611270131363503` that passes.
- **FR-005**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202611270131363503()` and print the result, while retaining its existing behavior.
- **FR-006**: The alpha module's dependency state MUST be tidied (`go mod tidy`) if the change requires it, so the module builds.
- **FR-007**: The beta change MUST land before the alpha change that depends on it; alpha work MUST NOT be completed against a function that does not yet exist.
- **FR-008**: The function name and string literal MUST be used verbatim as given in this spec. No substitute names or tokens may be introduced.

### Key Entities *(if applicable)*

- **Parent issue (alpha)**: This issue, #12109 in `handarbeit/fabrik-test-alpha`, owning the alpha-side change.
- **Spawned child issue (beta)**: The issue created in `handarbeit/fabrik-test-beta`, owning the beta-side change and blocking the parent.
- **Per-run-unique token**: The `20261127-013136-3503` / `E2E202611270131363503` identifiers that tie this run to its final assertion.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline completes, the beta repo contains the exact text `HelloE2E202611270131363503` and `e2e-cross-repo-spawn-20261127-013136-3503` in its greeting package, with a passing test.
- **SC-002**: A child issue targeting the beta repo exists as a result of this issue's Plan stage, and it is a blocker of this issue.
- **SC-003**: The alpha program, built and run after the beta change is available, prints `e2e-cross-repo-spawn-20261127-013136-3503`.
- **SC-004**: The beta change is merged before the alpha change that references it.

## Assumptions

- The beta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`, as the spec implies by "add to" those files.
- Alpha's `main.go` already imports the beta greeting package or can be given that import. Wiring the dependency is part of this work and need not be separately specified.
- The `fabrik:yolo` label means the stages advance without waiting for human approval of this spec.
- The `20261127-013136-3503` timestamp-like token is a fixture token, not a real date that constrains behavior.

## Out of Scope

- Any change to the beta repo beyond the single function and its test.
- Any change to the alpha repo beyond calling and printing the new function (and tidying module files as needed).
- Changing Fabrik's spawn, board-servability or repo-initialization behavior. This scenario only verifies it.
- Follow-up refactoring of the greeting package or the alpha program.

## Source References

- #803 — on-demand spawn-target repo init (behavior under test)
- #1308 — per-run-unique fixture token (behavior under test)
