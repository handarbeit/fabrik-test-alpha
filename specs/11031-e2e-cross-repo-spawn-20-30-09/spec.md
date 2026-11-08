# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-11031`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202611080130099963() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261108-013009-9963".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202611080130099963() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This issue is a live end-to-end regression scenario for Fabrik's cross-repo decomposition. It guards two earlier fixes: #803 (on-demand initialisation of a spawn-target repository) and #1308 (a per-run-unique fixture token, so repeated runs cannot be confused with one another).

The scenario needs a piece of work that cannot be done in one repository. A new exported function must exist in the beta repo (`handarbeit/fabrik-test-beta`) before the alpha repo (`handarbeit/fabrik-test-alpha`, this repo) can import and call it. The work therefore has a strict order: beta first, then alpha. The pipeline's Plan stage is expected to recognise this and decompose the work by spawning a sub-issue in the beta repo, which blocks this issue until it is done.

The function name and string literal contain a per-run-unique token. The scenario's final assertion checks for that exact text in the beta repo, so they must be used exactly as given.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repo exposes a new exported greeting function that returns a fixed, per-run-unique string, and has a test covering it. This work is delivered through a sub-issue spawned in the beta repo, not directly from this issue.

**Why this priority**: Everything else depends on it. Alpha cannot compile against a function that does not exist, and the scenario's final assertion looks for this text in the beta repo.

**Independent Test**: In the beta repo, the greeting package's tests pass, and calling `HelloE2E202611080130099963()` returns exactly `e2e-cross-repo-spawn-20261108-013009-9963`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202611080130099963()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261108-013009-9963`.
2. **Given** the beta repo's greeting test file, **When** the tests run, **Then** a test for this function exists and passes.

---

### User Story 2 - Alpha repo uses the beta function (Priority: P2)

After the beta change has landed, the alpha repo's entry point calls the new beta function, in addition to its existing behaviour, and prints the result.

**Why this priority**: This is the dependent half of the change. It proves the ordering constraint was honoured, because it cannot build until beta has landed.

**Independent Test**: With the beta change available, build and run the alpha program. Its output includes `e2e-cross-repo-spawn-20261108-013009-9963`, and its existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** alpha's `main.go` runs, **Then** it prints the string returned by `greeting.HelloE2E202611080130099963()`.
2. **Given** alpha's existing behaviour, **When** the new call is added, **Then** the existing behaviour is preserved.
3. **Given** dependency changes are needed, **When** `go mod tidy` is run, **Then** the module builds cleanly.

---

### User Story 3 - Plan decomposes the work across repos (Priority: P1)

The Plan stage recognises that the work spans two repos with a strict ordering and decomposes it by spawning a sub-issue targeting the beta repo, rather than planning all changes into this single issue.

**Why this priority**: Exercising cross-repo spawn, and the on-demand init of the spawn target, is the whole point of this scenario.

**Independent Test**: After Plan runs, a sub-issue exists in `handarbeit/fabrik-test-beta` and blocks this issue. This issue's own plan covers only the alpha-side change.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan completes, **Then** it has emitted a spawn block targeting the beta repo.
2. **Given** the spawned beta sub-issue, **When** it is created, **Then** it carries the exact function name and string literal from this spec.
3. **Given** the alpha work depends on the beta work, **When** the pipeline proceeds, **Then** the alpha change does not land before the beta change.

---

### Edge Cases

- The spawn-target repo has not been initialised for the pipeline yet. The spawn must still succeed (#803).
- The function name or string literal is altered during decomposition. The final assertion would then fail, so both must be carried through unchanged.
- The beta change has not landed when the alpha change is ready. Alpha must not land first.
- A previous run left a similar function in the beta repo. The per-run-unique token must keep this run's text distinguishable (#1308).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST gain an exported function `HelloE2E202611080130099963() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202611080130099963()` MUST return exactly the literal `e2e-cross-repo-spawn-20261108-013009-9963`.
- **FR-003**: The beta repo MUST gain a test for this function in `pkg/greeting/greeting_test.go`, and it MUST pass.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202611080130099963()` and print the result, while preserving its existing behaviour.
- **FR-005**: The alpha repo MUST build cleanly after the change, running `go mod tidy` if needed.
- **FR-006**: The beta change MUST land before the alpha change.
- **FR-007**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo, so that the beta-side work is carried by a sub-issue spawned there.
- **FR-008**: The function name and string literal MUST be used exactly as given in this spec, in the spawned sub-issue and in both repos. No substitute names or values may be invented.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A sub-issue exists in `handarbeit/fabrik-test-beta` as a result of Plan's spawn, and blocks this issue.
- **SC-002**: The beta repo contains `HelloE2E202611080130099963` returning `e2e-cross-repo-spawn-20261108-013009-9963`, with a passing test.
- **SC-003**: The alpha program's output contains `e2e-cross-repo-spawn-20261108-013009-9963`.
- **SC-004**: The beta change merges before the alpha change.
- **SC-005**: The scenario's final assertion finds the exact per-run-unique text in the beta repo.

## Assumptions

- `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` already exist in the beta repo, and alpha already imports the beta greeting package or can add it as a dependency. If not, the work includes setting that up.
- The `fabrik:yolo` label means this issue auto-advances through stages. Spec approval therefore does not wait on a human.
- The per-run token in the function name and literal is final and not to be regenerated.
- "In addition to whatever it already does" means alpha's existing `main.go` behaviour stays as it is.

## Out of Scope

- Any change to the spawn mechanism itself. This issue only exercises it.
- Any greeting functions, or changes to existing functions, beyond the one named here.
- Cleanup of fixture artefacts from earlier runs.

## Source References

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
