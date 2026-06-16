# Feature Specification: E2E Cross-Repo Spawn (Issue 31)

**Feature Branch**: `fabrik/issue-31`
**Created**: 2026-06-16
**Status**: Draft
**Input**: User description: "Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order."

## Background

Fabrik's cross-repo decomposition feature allows a Plan stage to spawn child issues in other repositories, execute them in strict dependency order, and coordinate the result. Issue #8 previously exercised this path by adding `HelloE2E()` to `fabrik-test-beta` and having `fabrik-test-alpha` call it. That code is now merged into both repos.

Issue #31 is a fresh regression run targeting the same Fabrik capability — on-demand spawn-target repo init (Fabrik issue #803). To serve as a meaningful and non-conflicting test, it introduces a **new** exported function (`HelloIssue31()`) in `fabrik-test-beta` alongside the existing `HelloE2E()`, and updates `fabrik-test-alpha` to also call `HelloIssue31()`. The Plan stage must emit a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta` so that the beta change lands before alpha's Implement stage begins.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta exports HelloIssue31 (Priority: P1)

A developer working on `fabrik-test-beta` after this change has a new exported function `HelloIssue31() string` in `pkg/greeting` returning the fixed string `"e2e-cross-repo-spawn-31"`. Tests exist alongside it. The existing `HelloE2E()` and `GreetingFor()` functions are unchanged.

**Why this priority**: Alpha cannot compile or pass tests until this function exists in beta at a resolvable module version. It must land and merge first.

**Independent Test**: Check out `fabrik-test-beta` after the sub-issue merges, run `go test ./pkg/greeting/...` — all tests pass including at least one case covering `HelloIssue31`.

**Acceptance Scenarios**:

1. **Given** `fabrik-test-beta` at HEAD after this change, **When** `greeting.HelloIssue31()` is called, **Then** it returns exactly `"e2e-cross-repo-spawn-31"`
2. **Given** `fabrik-test-beta` at HEAD after this change, **When** `go test ./...` is run, **Then** all tests pass including the `HelloIssue31` test case
3. **Given** `fabrik-test-beta` at HEAD after this change, **When** `go build ./...` is run, **Then** it exits 0
4. **Given** the beta change merged, **When** Fabrik's sub-issue for beta reaches Done, **Then** alpha's Implement stage may proceed

---

### User Story 2 - Alpha prints HelloIssue31 output (Priority: P2)

A developer running the `fabrik-test-alpha` binary after this change sees an additional line sourced from `greeting.HelloIssue31()` printed to stdout, alongside the existing `GreetingFor` and `HelloE2E` lines. All existing behaviour is preserved.

**Why this priority**: Depends on Story 1 (beta must merge first). This is the primary observable outcome validating the cross-repo spawn path end-to-end.

**Independent Test**: Build and run the alpha binary with no arguments; assert stdout contains `"Hello, world, from fabrik-test-beta"`, `"e2e-cross-repo-spawn"`, and `"e2e-cross-repo-spawn-31"`.

**Acceptance Scenarios**:

1. **Given** the alpha binary built at HEAD after this change, **When** run with no arguments, **Then** stdout contains `"Hello, world, from fabrik-test-beta"`, `"e2e-cross-repo-spawn"`, and `"e2e-cross-repo-spawn-31"` (all three lines present, order unspecified)
2. **Given** the alpha binary built at HEAD after this change, **When** run with argument `"Alice"`, **Then** stdout contains `"Hello, Alice, from fabrik-test-beta"`, `"e2e-cross-repo-spawn"`, and `"e2e-cross-repo-spawn-31"`
3. **Given** `fabrik-test-alpha` at HEAD, **When** `go test ./...` is run, **Then** all tests pass
4. **Given** `fabrik-test-alpha` at HEAD, **When** `go build ./...` is run, **Then** it exits 0

---

### Edge Cases

- `HelloIssue31()` returns a fixed string with no parameters — there are no boundary conditions on inputs
- `go.mod` in alpha must be updated to a version of `fabrik-test-beta` that includes `HelloIssue31`; `go.sum` must be committed; `go mod tidy` must be run if module graph changes
- The existing `GreetingFor(name)` call and `HelloE2E()` call in `main.go` must not be removed or altered
- `main_test.go` must be updated to assert `"e2e-cross-repo-spawn-31"` appears in output (in addition to existing assertions)

## Requirements *(mandatory)*

### Functional Requirements

#### fabrik-test-beta

- **FR-001**: `pkg/greeting/greeting.go` MUST export `HelloIssue31() string` returning the fixed string `"e2e-cross-repo-spawn-31"`
- **FR-002**: The existing `GreetingFor(name string) string` and `HelloE2E() string` functions MUST remain unchanged
- **FR-003**: `pkg/greeting/greeting_test.go` MUST include at least one test case verifying `HelloIssue31()` returns `"e2e-cross-repo-spawn-31"`
- **FR-004**: `go build ./...` and `go test ./...` MUST exit 0 in `fabrik-test-beta`
- **FR-005**: `go.mod` MUST remain pinned to `go 1.22`

#### fabrik-test-alpha

- **FR-006**: `main.go` MUST call `greeting.HelloIssue31()` and print the returned string to stdout (via `fmt.Println` or equivalent)
- **FR-007**: `main.go` MUST retain the existing `greeting.GreetingFor(name)` call and `greeting.HelloE2E()` call and their prints
- **FR-008**: `go.mod` MUST be updated to depend on the version of `fabrik-test-beta` that includes `HelloIssue31`; `go.sum` MUST be committed
- **FR-009**: `go mod tidy` MUST be run if module graph changes
- **FR-010**: `go build ./...` and `go test ./...` MUST exit 0 in `fabrik-test-alpha`
- **FR-011**: `go.mod` MUST remain pinned to `go 1.22`
- **FR-012**: `main_test.go` MUST assert `"e2e-cross-repo-spawn-31"` appears in the binary output for both the default and named-argument test cases

#### Plan stage (Fabrik pipeline constraint)

- **FR-013**: The Plan stage for this issue MUST emit a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta`, so that the beta sub-issue is created and executed before alpha's Implement stage begins
- **FR-014**: The beta sub-issue MUST be marked as a blocking dependency of this issue so that alpha's Implement stage does not start until beta's sub-issue reaches Done

### Key Entities

- **`HelloIssue31() string`**: New exported function in `fabrik-test-beta/pkg/greeting`. No parameters; returns the fixed string `"e2e-cross-repo-spawn-31"`. Serves as the new API surface for this cross-repo regression run, distinct from `HelloE2E()` from issue #8.
- **`fabrik-test-beta` sub-issue**: The child issue spawned by the Plan stage in `handarbeit/fabrik-test-beta` covering FR-001 through FR-005.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Running the alpha binary with no arguments prints `"Hello, world, from fabrik-test-beta"`, `"e2e-cross-repo-spawn"`, and `"e2e-cross-repo-spawn-31"` to stdout and exits 0
- **SC-002**: Running the alpha binary with argument `"Alice"` prints `"Hello, Alice, from fabrik-test-beta"`, `"e2e-cross-repo-spawn"`, and `"e2e-cross-repo-spawn-31"` to stdout and exits 0
- **SC-003**: `go test ./...` exits 0 in `fabrik-test-beta` with at least one test covering `HelloIssue31`
- **SC-004**: `go test ./...` exits 0 in `fabrik-test-alpha` with assertions covering `"e2e-cross-repo-spawn-31"`
- **SC-005**: Both repos' CI workflows pass green after their respective changes merge
- **SC-006**: Fabrik's cross-repo decomposition and on-demand repo init code path (regression test for #803) is exercised: the Plan stage spawns a sub-issue in `fabrik-test-beta`, that sub-issue completes, and then alpha's Implement stage proceeds

## Assumptions

- The `arbeithand` PAT has write access to both `handarbeit/fabrik-test-alpha` and `handarbeit/fabrik-test-beta`
- `fabrik-test-beta` already exports `GreetingFor` and `HelloE2E`; adding `HelloIssue31` alongside them is additive and non-breaking
- `fabrik-test-alpha` already calls `GreetingFor` and `HelloE2E`; adding a call to `HelloIssue31` alongside them is additive and non-breaking
- Fabrik v0.0.66+ (on-demand spawn-target repo init, issue #803) is deployed
- The beta sub-issue will be merged and a commit-pinned version available before alpha's `go.mod` is updated
- Both repos' CI runs `go build ./...` and `go test ./...` — no additional CI steps need to be satisfied

## Out of Scope

- Any changes to `GreetingFor` or `HelloE2E` in beta
- Removing or replacing existing calls in `main.go`
- Adding benchmarks, fuzzing, or additional test infrastructure beyond the `HelloIssue31` unit test
- Publishing `fabrik-test-beta` to a Go module proxy; direct GitHub module references are sufficient

## Source References

- `fabrik-test-alpha/main.go` — current alpha entry point (calls `GreetingFor` and `HelloE2E`, must also call `HelloIssue31`)
- `fabrik-test-alpha/main_test.go` — current integration tests (must be updated to assert `"e2e-cross-repo-spawn-31"`)
- `handarbeit/fabrik-test-beta/pkg/greeting/greeting.go` — target of beta change (add `HelloIssue31`)
- `handarbeit/fabrik-test-beta/pkg/greeting/greeting_test.go` — target of beta test update
- Fabrik issue #803 — on-demand spawn-target repo init regression scenario this issue validates
- Issue #8 spec (`specs/8-e2e-cross-repo-spawn/spec.md`) — prior run of the same regression test pattern
