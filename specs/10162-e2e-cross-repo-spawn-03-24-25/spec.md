# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-10162`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202610300724254399() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261030-072425-4399".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202610300724254399() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).
>
> ## Labels
>
> fabrik:yolo

## Background

This is an end-to-end regression scenario for the pipeline's cross-repo decomposition. It covers two earlier fixes:

- #803: on-demand initialisation of a spawn-target repository.
- #1308: a per-run-unique fixture token, so each run's assertion can only be satisfied by that run's own output.

The work spans two repositories with a strict ordering dependency. The beta repo (`handarbeit/fabrik-test-beta`) must gain a new greeting function before the alpha repo (`handarbeit/fabrik-test-alpha`, this one) can call it. A correct pipeline run therefore decomposes the work: Plan spawns a child issue in the beta repo, and the alpha work waits until the beta work has landed.

The function name and string literal embed a per-run-unique token. The scenario's final assertion checks that this exact text appears in the beta repo, so both must be used exactly as given.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the run-unique greeting function (Priority: P1)

A new exported function exists in the beta repo's greeting package. It returns the run-unique string, and a test covers it.

**Why this priority**: Alpha depends on it, and the scenario's final assertion inspects the beta repo for this exact text. Nothing else can succeed without it.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261030-072425-4399`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610300724254399()` is called, **Then** it returns the string `e2e-cross-repo-spawn-20261030-072425-4399`.
2. **Given** the beta repo's greeting test file, **When** the package tests run, **Then** a test for `HelloE2E202610300724254399` exists and passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

The alpha repo's main program calls the new beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It completes the cross-repo change. It can only be done after User Story 1 has landed.

**Independent Test**: Build and run the alpha program once the beta change is available. Its output includes `e2e-cross-repo-spawn-20261030-072425-4399`, and its previous output is still present.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** the alpha program runs, **Then** it prints the string returned by `greeting.HelloE2E202610300724254399()`.
2. **Given** the alpha program's pre-existing behaviour, **When** the change is applied, **Then** that behaviour is unchanged and the new output is added to it.
3. **Given** the dependency on the beta repo changed, **When** the module is built, **Then** dependency metadata is consistent (`go mod tidy` has been run if needed).

---

### User Story 3 - Work is decomposed across repos in the required order (Priority: P1)

The pipeline treats the beta work as a separate child issue in the beta repo. The alpha work is sequenced after it.

**Why this priority**: Exercising this decomposition is the purpose of the scenario (#803).

**Independent Test**: After Plan runs, a child issue exists in `handarbeit/fabrik-test-beta` describing the beta-side change, and the alpha issue is blocked on it until it lands.

**Acceptance Scenarios**:

1. **Given** the alpha issue reaches Plan, **When** Plan decomposes the work, **Then** it emits a spawn block targeting the beta repo for the beta-side change.
2. **Given** the child issue in the beta repo is open, **When** the alpha issue is evaluated, **Then** alpha implementation does not proceed ahead of the beta change.
3. **Given** the beta change has landed, **When** the alpha issue resumes, **Then** it can import and call the new function.

---

### Edge Cases

- The child issue's description must carry the exact function name and string literal. Paraphrasing or inventing a different token would defeat the scenario's final assertion.
- The beta repo may have had no prior pipeline activity (#803), so on-demand initialisation of the spawn target must succeed.
- If the alpha change were built before the beta function exists, the build would fail on the missing symbol. This is the ordering constraint the decomposition enforces.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's `pkg/greeting/greeting.go` MUST export a function named exactly `HelloE2E202610300724254399` with signature `() string`.
- **FR-002**: `HelloE2E202610300724254399` MUST return exactly the literal `e2e-cross-repo-spawn-20261030-072425-4399`.
- **FR-003**: The beta repo's `pkg/greeting/greeting_test.go` MUST contain a test that verifies the return value of `HelloE2E202610300724254399`.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202610300724254399()` and print the result, while preserving its existing behaviour.
- **FR-005**: The alpha repo's module metadata MUST be consistent after the change (`go mod tidy` run if needed).
- **FR-006**: The beta-side change MUST be delivered through a child issue spawned in `handarbeit/fabrik-test-beta` by Plan, not authored directly from this issue's own branch.
- **FR-007**: The beta-side change MUST land before the alpha-side change is implemented or merged.
- **FR-008**: The exact function name and string literal given above MUST be used verbatim in all artifacts (child issue, code, tests). No alternative token may be invented.

### Key Entities

- **Alpha repo** (`handarbeit/fabrik-test-alpha`): the consumer, containing the main program that prints the greeting.
- **Beta repo** (`handarbeit/fabrik-test-beta`): the provider, containing the greeting package.
- **Child issue**: the spawned issue in the beta repo that carries the beta-side work.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains `HelloE2E202610300724254399` returning `e2e-cross-repo-spawn-20261030-072425-4399`, with a passing test.
- **SC-002**: After the run, the alpha repo's program output includes `e2e-cross-repo-spawn-20261030-072425-4399`, and its previous output is unchanged.
- **SC-003**: A child issue was created in the beta repo by Plan, and the beta change merged before the alpha change.
- **SC-004**: The scenario's final assertion finds the exact run-unique token in the beta repo.

## Assumptions

- The beta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`, as the request states.
- The alpha repo's `main.go` can import the beta greeting package, and the module dependency wiring needed is part of the alpha-side work.
- The `fabrik:yolo` label means stages auto-advance once their own gates are satisfied. It does not change the ordering requirement.
- The token embedded in the function name and literal is already unique to this run and needs no further generation.

## Out of Scope

- Any greeting functionality beyond the single run-unique function and its test.
- Changes to the pipeline's decomposition or spawn mechanisms themselves; this issue only exercises them.
- Cleanup of fixtures after the run.

## Source References

- #803: on-demand spawn-target repo init.
- #1308: per-run-unique fixture token.
