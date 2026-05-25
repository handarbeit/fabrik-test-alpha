# Feature Specification: E2E Cross-Repo Spawn (HelloE2E)

**Feature Branch**: `fabrik/issue-8`
**Created**: 2026-05-24
**Status**: Draft
**Input**: User description: "Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order."

## Background

Fabrik's cross-repo decomposition feature allows a Plan stage to spawn child issues in other repositories, execute them in strict dependency order, and coordinate the result. Issue #1 bootstrapped the two-repo test bed (`fabrik-test-alpha` + `fabrik-test-beta`) by wiring them together via `GreetingFor`. Issue #8 is a targeted regression test for Fabrik issue #803 — on-demand spawn-target repo init — which introduced the ability to initialize a target repo's Fabrik worktree on-demand at spawn time (i.e., the target repo need not have been previously known to Fabrik).

This issue exercises that code path by adding a second exported function in beta (`HelloE2E()`) that alpha then calls. The Plan stage must emit a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta` so that the beta change lands and CI passes before alpha's implementation begins. Without `HelloE2E` merged in beta, the alpha build would fail at import time.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta exports HelloE2E (Priority: P1)

A developer working on `fabrik-test-beta` after this change has a new exported function `HelloE2E() string` in `pkg/greeting` returning the fixed string `"e2e-cross-repo-spawn"`. Tests for the function exist alongside it.

**Why this priority**: Alpha cannot compile until this function exists in beta and is available at a resolvable module version. It must land first.

**Independent Test**: Check out `fabrik-test-beta` after the sub-issue merges, run `go test ./pkg/greeting/...` — the greeting package tests must pass including at least one case covering `HelloE2E`.

**Acceptance Scenarios**:

1. **Given** `fabrik-test-beta` at HEAD after this change, **When** `greeting.HelloE2E()` is called, **Then** it returns exactly `"e2e-cross-repo-spawn"`
2. **Given** `fabrik-test-beta` at HEAD after this change, **When** `go test ./...` is run, **Then** all tests pass including the `HelloE2E` test case
3. **Given** `fabrik-test-beta` at HEAD after this change, **When** `go build ./...` is run, **Then** it exits 0
4. **Given** the beta change, **When** Fabrik's sub-issue for beta reaches the Done stage, **Then** alpha's implementation stage may proceed

---

### User Story 2 - Alpha prints HelloE2E output (Priority: P2)

A developer running the `fabrik-test-alpha` binary after this change sees an additional line sourced from `greeting.HelloE2E()` printed to stdout, alongside the existing greeting line. The binary's existing behaviour (accepting an optional name argument and printing `GreetingFor(name)`) is preserved.

**Why this priority**: Depends on Story 1 (beta must be merged first). This is the primary observable outcome that validates the cross-repo spawn path end-to-end.

**Independent Test**: Build and run the alpha binary with no arguments; assert stdout contains both `"from fabrik-test-beta"` (from `GreetingFor`) and `"e2e-cross-repo-spawn"` (from `HelloE2E`).

**Acceptance Scenarios**:

1. **Given** the alpha binary built at HEAD after this change, **When** run with no arguments, **Then** stdout contains `"Hello, world, from fabrik-test-beta"` on one line and `"e2e-cross-repo-spawn"` on another line (both printed, order unspecified)
2. **Given** the alpha binary built at HEAD after this change, **When** run with argument `"Alice"`, **Then** stdout contains `"Hello, Alice, from fabrik-test-beta"` and `"e2e-cross-repo-spawn"`
3. **Given** `fabrik-test-alpha` at HEAD, **When** `go test ./...` is run, **Then** all tests pass
4. **Given** `fabrik-test-alpha` at HEAD, **When** `go build ./...` is run, **Then** it exits 0

---

### Edge Cases

- `HelloE2E()` returns a fixed string with no parameters — there are no boundary conditions on inputs
- `go.mod` in alpha already depends on `fabrik-test-beta v0.1.0`; it must be updated to a version that includes `HelloE2E`; `go mod tidy` must be run and `go.sum` committed
- The existing `GreetingFor(name)` call in `main.go` must not be removed or altered

## Requirements *(mandatory)*

### Functional Requirements

#### fabrik-test-beta

- **FR-001**: `pkg/greeting/greeting.go` MUST export `HelloE2E() string` returning the fixed string `"e2e-cross-repo-spawn"`
- **FR-002**: The existing `GreetingFor(name string) string` function MUST remain unchanged
- **FR-003**: `pkg/greeting/greeting_test.go` MUST include at least one test case verifying `HelloE2E()` returns `"e2e-cross-repo-spawn"`
- **FR-004**: `go build ./...` and `go test ./...` MUST exit 0 in `fabrik-test-beta`
- **FR-005**: `go.mod` MUST remain pinned to `go 1.22`

#### fabrik-test-alpha

- **FR-006**: `main.go` MUST call `greeting.HelloE2E()` and print the returned string to stdout (via `fmt.Println` or equivalent)
- **FR-007**: `main.go` MUST retain the existing `greeting.GreetingFor(name)` call and its print
- **FR-008**: `go.mod` MUST be updated to depend on the version of `fabrik-test-beta` that includes `HelloE2E`; `go.sum` MUST be committed
- **FR-009**: `go mod tidy` MUST be run if module graph changes
- **FR-010**: `go build ./...` and `go test ./...` MUST exit 0 in `fabrik-test-alpha`
- **FR-011**: `go.mod` MUST remain pinned to `go 1.22`

#### Plan stage (Fabrik pipeline constraint)

- **FR-012**: The Plan stage for this issue MUST emit a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta`, so that the beta sub-issue is created and executed before alpha's Implement stage begins
- **FR-013**: The beta sub-issue MUST be marked as a blocking dependency of this issue so that alpha's Implement stage does not start until beta's sub-issue reaches Done

### Key Entities

- **`HelloE2E() string`**: New exported function in `fabrik-test-beta/pkg/greeting`. No parameters; returns the fixed string `"e2e-cross-repo-spawn"`. Serves as a minimal new API surface for the cross-repo regression test.
- **`fabrik-test-beta` sub-issue**: The child issue spawned by the Plan stage in `handarbeit/fabrik-test-beta` covering FR-001 through FR-005.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Running the alpha binary with no arguments prints both `"Hello, world, from fabrik-test-beta"` and `"e2e-cross-repo-spawn"` to stdout and exits 0
- **SC-002**: Running the alpha binary with argument `"Alice"` prints both `"Hello, Alice, from fabrik-test-beta"` and `"e2e-cross-repo-spawn"` to stdout and exits 0
- **SC-003**: `go test ./...` exits 0 in `fabrik-test-beta` with at least one test covering `HelloE2E`
- **SC-004**: `go test ./...` exits 0 in `fabrik-test-alpha`
- **SC-005**: Both repos' CI workflows pass green after their respective changes merge
- **SC-006**: Fabrik's cross-repo decomposition and on-demand repo init code path (regression test for #803) is exercised: the Plan stage spawns a sub-issue in `fabrik-test-beta`, that sub-issue completes, and then alpha's implementation proceeds

## Assumptions

- The `arbeithand` PAT has write access to both `handarbeit/fabrik-test-alpha` and `handarbeit/fabrik-test-beta`
- `fabrik-test-beta` already exports `GreetingFor`; adding `HelloE2E` alongside it is additive and non-breaking
- Fabrik v0.0.66+ (which includes the on-demand spawn-target repo init feature for #803) is deployed for this run
- The beta sub-issue will be merged and a tagged/commit-pinned version available before alpha's `go.mod` is updated
- Both repos' CI runs `go build ./...` and `go test ./...` — no additional CI steps need to be satisfied

## Out of Scope

- Any changes to `GreetingFor` or other existing functions in beta
- Adding flags or additional behaviour to the alpha binary beyond the `HelloE2E` print
- Removing or replacing the existing `GreetingFor` call in `main.go`
- Adding benchmarks, fuzzing, or additional test infrastructure beyond the `HelloE2E` unit test
- Publishing `fabrik-test-beta` to a Go module proxy; direct GitHub module references are sufficient

## Source References

- `fabrik-test-alpha/main.go` — current alpha entry point (calls `GreetingFor`, must also call `HelloE2E`)
- `fabrik-test-alpha/main_test.go` — current integration tests
- `handarbeit/fabrik-test-beta/pkg/greeting/greeting.go` — target of beta change (add `HelloE2E`)
- `handarbeit/fabrik-test-beta/pkg/greeting/greeting_test.go` — target of beta test update
- Fabrik issue #803 — on-demand spawn-target repo init regression scenario this issue validates
- Fabrik issue #1 (`specs/1-bootstrap-wire-fabrik-test/spec.md`) — prior bootstrap spec establishing the two-repo test bed
