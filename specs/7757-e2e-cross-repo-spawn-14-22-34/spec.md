# Feature Specification: E2E cross-repo spawn of a beta-repo sub-issue

**Feature Branch**: `fabrik/issue-7757`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202610131822346893() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261013-182234-6893".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202610131822346893() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This issue is a live end-to-end regression scenario, not a product feature. It guards two previously fixed behaviours:

- **#803**: a spawn-target repository (here `handarbeit/fabrik-test-beta`) is initialised on demand when a sub-issue is spawned into it.
- **#1308**: each run uses a unique fixture token, so the final assertion cannot be satisfied by leftovers from an earlier run.

The scenario needs a change that spans two repositories with a strict ordering. The beta repo must gain a new greeting function first, because the alpha repo (this one) cannot import a function that does not yet exist. That ordering is what forces the pipeline to decompose the work into a sub-issue in the beta repo instead of doing it all in one place.

The function name and string literal are generated per run. The run's final assertion checks that exactly this text lands in the beta repo, so it must be used as given.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function via a spawned sub-issue (Priority: P1)

The pipeline decomposes the work, so the beta-repo change is carried out as a sub-issue in `handarbeit/fabrik-test-beta` rather than as part of this alpha-repo issue.

**Why this priority**: This is the behaviour under test. Without a cross-repo spawn that lands the beta change, the scenario has no value and the final assertion fails.

**Independent Test**: After the pipeline runs, the beta repo's default branch contains the function `HelloE2E202610131822346893` in `pkg/greeting/greeting.go`, returning `"e2e-cross-repo-spawn-20261013-182234-6893"`. A test for it exists in `pkg/greeting/greeting_test.go`, and a sub-issue in the beta repo exists that is linked to this issue.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan decomposes the work, **Then** it declares a spawned sub-issue targeting `handarbeit/fabrik-test-beta` that covers the beta-repo change.
2. **Given** the beta repo has not yet been prepared for spawning, **When** the sub-issue is spawned, **Then** the beta repo is initialised on demand and the sub-issue is created there without manual setup (#803).
3. **Given** the sub-issue is processed, **When** it completes, **Then** the beta repo contains the exported function with the exact name and exact returned string literal, plus a passing test for it.

---

### User Story 2 - Alpha repo consumes the new function only after beta has landed (Priority: P1)

This repo's `main.go` calls the new beta function and prints its result, in addition to its existing behaviour. The alpha change must not land or build before the beta function exists.

**Why this priority**: This is the other half of the cross-repo change, and it checks the strict ordering the scenario is built around.

**Independent Test**: After the beta change has landed, building and running this repo's program prints the string `e2e-cross-repo-spawn-20261013-182234-6893`, alongside whatever it printed before.

**Acceptance Scenarios**:

1. **Given** the beta sub-issue has not completed, **When** this issue would otherwise proceed to implementation, **Then** it is held (blocked on the sub-issue) so the alpha change cannot be built against a missing function.
2. **Given** the beta function exists, **When** this repo's `main.go` is updated, **Then** it calls `greeting.HelloE2E202610131822346893()` and prints the result, and the module's dependency metadata is consistent (tidied if needed).

---

### Edge Cases

- The per-run token is already baked into the function name and the string literal. They must be used verbatim and must not be regenerated, shortened or varied, because the final assertion matches this exact text in the beta repo.
- If the beta repo has not been set up for spawning, on-demand initialisation must cover it (#803) rather than failing the spawn.
- The existing behaviour of this repo's program must be preserved. The new output is additive.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST expose an exported function named exactly `HelloE2E202610131822346893` in `pkg/greeting/greeting.go`. It takes no arguments and returns a `string`.
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261013-182234-6893`.
- **FR-003**: The beta repo MUST include a test for that function in `pkg/greeting/greeting_test.go` that passes.
- **FR-004**: This repo's `main.go` MUST call `greeting.HelloE2E202610131822346893()`, in addition to its existing behaviour, and print the returned value.
- **FR-005**: This repo's module dependency metadata MUST be tidied (`go mod tidy`) if the new import requires it.
- **FR-006**: The beta-repo work MUST be decomposed into a spawned sub-issue targeting `handarbeit/fabrik-test-beta`, declared in Plan's output with a spawn block, rather than done within this alpha-repo issue.
- **FR-007**: The beta change MUST land before the alpha change. The alpha work MUST NOT be built or landed against a beta repo that lacks the function.
- **FR-008**: The exact function name and string literal given in this spec MUST be used verbatim in every artifact (sub-issue description, code, tests). No substitute name or token may be invented.

### Key Entities *(if applicable)*

- **Alpha repo** (`handarbeit/fabrik-test-alpha`): this repo. It consumes the greeting function from `main.go`.
- **Beta repo** (`handarbeit/fabrik-test-beta`): the spawn target. It owns `pkg/greeting` and defines the new function.
- **Per-run fixture token**: the unique identifier `E2E202610131822346893` / `e2e-cross-repo-spawn-20261013-182234-6893`, which ties the run's final assertion to this run's output.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A sub-issue targeting `handarbeit/fabrik-test-beta` exists and is linked to this issue as a blocker, created by Plan's decomposition.
- **SC-002**: The beta repo's default branch contains the string `e2e-cross-repo-spawn-20261013-182234-6893` returned by `HelloE2E202610131822346893` in `pkg/greeting/greeting.go`, and the beta test suite passes.
- **SC-003**: Running this repo's program after both changes prints `e2e-cross-repo-spawn-20261013-182234-6893`, and its previous output is unchanged.
- **SC-004**: The beta change lands strictly before the alpha change. The alpha change never builds against a missing function.
- **SC-005**: The on-demand initialisation of the spawn-target repo (#803) and the per-run-unique token (#1308) both work, shown by the scenario's final assertion passing with no manual intervention.

## Assumptions

- The beta repo already has `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`, or on-demand initialisation provides an equivalent package layout. The issue treats them as the places to edit.
- This repo already imports the beta repo's `greeting` package, or adding that import is part of the `main.go` change. Either way, dependency metadata is tidied as needed.
- The `fabrik:yolo` label is present on this issue, so stages auto-advance. This does not change the required ordering of the two repos' changes.
- The function name and string literal in this spec are final for this run and must not change.

## Out of Scope *(optional)*

- Any change to the engine, the spawn mechanism or the on-demand repo initialisation. This issue only exercises them.
- Any change to the beta repo beyond the one function and its test.
- Any change to this repo's behaviour other than the additive call and print in `main.go` (plus `go mod tidy` if needed).
- Designing how Plan expresses the spawn. That is a Plan-stage decision.

## Source References *(optional)*

- #803: on-demand spawn-target repo initialisation (behaviour under regression test)
- #1308: per-run-unique fixture token (behaviour under regression test)
