# Feature Specification: E2E cross-repo spawn (beta function, then alpha consumer)

**Feature Branch**: `fabrik/issue-10885`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description (verbatim):

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202611072104036462() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261107-210403-6462".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202611072104036462() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This issue is a live end-to-end regression scenario, not a product feature. It checks that a single issue whose work spans two repositories is decomposed correctly: the part that must land first (in `handarbeit/fabrik-test-beta`) is split into a child issue in that repo, and the dependent part (in `handarbeit/fabrik-test-alpha`) waits on it. The dependency is real. Alpha cannot call a function that does not yet exist in beta, so beta must land first.

It guards two earlier fixes:
- #803: on-demand initialisation of a spawn-target repo.
- #1308: a per-run-unique fixture token, so that a stale or invented value cannot satisfy the final assertion.

The function name `HelloE2E202611072104036462` and the string `e2e-cross-repo-spawn-20261107-210403-6462` carry that per-run token. They are fixed by this scenario and must be used exactly as written.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the run-unique greeting function (Priority: P1)

The beta repo gains an exported greeting function that returns the run-unique literal, with a test covering it. This is the work that must land first, and Plan is expected to decompose it into a child issue in the beta repo.

**Why this priority**: Alpha's change depends on it. Without it, alpha cannot compile against the new function, and the scenario's final assertion (the token text present in the beta repo) fails.

**Independent Test**: In the beta repo, run the package's tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261107-210403-6462`.

**Acceptance Scenarios**:

1. **Given** the beta repo's `pkg/greeting` package, **When** `HelloE2E202611072104036462()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261107-210403-6462`.
2. **Given** the beta repo's greeting test file, **When** the package's tests run, **Then** a test for `HelloE2E202611072104036462` exists and passes.
3. **Given** this issue enters Plan, **When** Plan decomposes the work, **Then** it emits a spawn-child block targeting the beta repo for the beta-side change.

---

### User Story 2 - Alpha consumes the beta function and prints the result (Priority: P2)

After beta has landed, the alpha repo's `main.go` calls `greeting.HelloE2E202611072104036462()` and prints the returned value, in addition to its existing behaviour. Module metadata is tidied if the new dependency requires it.

**Why this priority**: This proves the dependent half works once its prerequisite exists. It is blocked on User Story 1 by design.

**Independent Test**: With the beta change available to alpha, build and run alpha. The output includes `e2e-cross-repo-spawn-20261107-210403-6462`, and the existing behaviour is unchanged.

**Acceptance Scenarios**:

1. **Given** beta exposes `HelloE2E202611072104036462`, **When** alpha is built and run, **Then** it prints `e2e-cross-repo-spawn-20261107-210403-6462`.
2. **Given** alpha's existing `main.go` behaviour, **When** the new call is added, **Then** the existing behaviour is still produced.
3. **Given** the beta change has not yet landed, **When** the alpha item is evaluated, **Then** it is held behind the beta child issue and does not proceed ahead of it.

---

### Edge Cases

- The spawned child must be created in the beta repo, not on alpha's board by default. This depends on the spawn target repo being initialised on demand (#803).
- Using any name or literal other than the exact ones given (for example a regenerated token) must not satisfy the scenario. The token text landing in the beta repo is the assertion.
- If alpha's dependency metadata already resolves the new function without changes, tidying the module is a no-op and must not be forced.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export a function `HelloE2E202611072104036462() string` in `pkg/greeting/greeting.go` that returns exactly `e2e-cross-repo-spawn-20261107-210403-6462`.
- **FR-002**: The beta repo MUST include a test for `HelloE2E202611072104036462` in `pkg/greeting/greeting_test.go` that asserts the exact returned literal.
- **FR-003**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202611072104036462()` and print the result, while preserving its existing behaviour.
- **FR-004**: Alpha's module metadata MUST be tidied (`go mod tidy`) if, and only if, the new call requires it.
- **FR-005**: The beta-side change MUST land before the alpha-side change; the work MUST be decomposed so that alpha depends on the beta child issue.
- **FR-006**: Plan MUST decompose this work with a spawn-child block targeting the beta repo, rather than the beta change being made from within alpha.
- **FR-007**: The exact function name and string literal given in this issue MUST be used verbatim; they MUST NOT be altered or replaced by invented values.

### Key Entities *(if applicable)*

- **Parent issue (alpha)**: This issue, which owns the alpha-side change and depends on the child.
- **Child issue (beta)**: The spawned sub-issue in `handarbeit/fabrik-test-beta` that owns the beta-side change.
- **Run-unique token**: The per-run suffix embedded in the function name and literal that ties the final assertion to this specific run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue exists in `handarbeit/fabrik-test-beta`, created by decomposition of this issue, and the parent is blocked on it until it is resolved.
- **SC-002**: The beta repo's default branch contains the exact text `e2e-cross-repo-spawn-20261107-210403-6462` and the function name `HelloE2E202611072104036462` after the child lands.
- **SC-003**: Running the beta package's tests passes, including the new test.
- **SC-004**: Building and running alpha after both changes prints `e2e-cross-repo-spawn-20261107-210403-6462`.
- **SC-005**: The beta change lands before the alpha change.

## Assumptions

- The beta repo already has `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`, as the issue implies, and the alpha repo already imports (or can import) that greeting package.
- The alpha instance's board can serve the beta repo (multi-repo or covering both), so spawning there is permitted.
- "Print the result" means writing it to standard output, consistent with `main.go`'s existing output style.
- The exact function name and string in this issue are authoritative for this run and are not to be regenerated.

## Out of Scope *(optional)*

- Any change to the greeting package beyond the single new function and its test.
- Any other behavioural change to alpha's `main.go`.
- Changes to the engine's spawn or decomposition logic. This scenario only exercises it.
- Generating or rotating the run-unique token.

## Source References *(optional)*

- #803: on-demand spawn-target repo initialisation.
- #1308: per-run-unique fixture token.
