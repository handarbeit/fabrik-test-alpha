# Feature Specification: Bootstrap: wire fabrik-test-alpha to fabrik-test-beta end-to-end

**Feature Branch**: `fabrik/issue-1`
**Created**: 2026-05-24
**Status**: Draft
**Input**: User description: "Wire fabrik-test-alpha and fabrik-test-beta together so they form a working two-repo system. This is the canonical multi-repo smoke test for Fabrik — the very issue that exercises the cross-repo decomposition feature shipped in v0.0.66."

## Background

Fabrik v0.0.66 shipped cross-repo decomposition: the ability for a single issue filed against one repo to spawn sub-issues in other repos, execute them in order, and coordinate the result. This issue is the bootstrap smoke test for that feature, using two dedicated test repositories (`fabrik-test-alpha` and `fabrik-test-beta`) under the `handarbeit` org.

Currently `fabrik-test-alpha/main.go` is an empty scaffold that prints a "nothing wired" message to stderr. `fabrik-test-beta` exports a trivial `Greeting() string` function. Neither repo is connected to the other.

Once this issue is resolved:
1. The cross-repo decomposition feature is verified end-to-end (the v0.0.66 regression scenario from issue #800)
2. A working two-repo substrate exists for filing further Fabrik test scenarios
3. A baseline reproducer exists for any future cross-repo bug

This work **must** be decomposed by the Plan stage: the beta repo changes must land and CI must pass there before alpha can import the new function.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta exports GreetingFor (Priority: P1)

A developer running `fabrik-test-beta` after this change has a richer greeting function `GreetingFor(name string) string` available in `pkg/greeting`, replacing the old zero-argument `Greeting()`. The function returns `"Hello, <name>, from fabrik-test-beta"`.

**Why this priority**: Alpha cannot compile until this function exists in beta. It must land first.

**Independent Test**: Check out `fabrik-test-beta`, run `go test ./...` — the greeting package tests must pass with at least two cases (empty name and normal name).

**Acceptance Scenarios**:

1. **Given** `fabrik-test-beta` at HEAD after this change, **When** `greeting.GreetingFor("Alice")` is called, **Then** it returns `"Hello, Alice, from fabrik-test-beta"`
2. **Given** `fabrik-test-beta` at HEAD after this change, **When** `greeting.GreetingFor("")` is called, **Then** it returns `"Hello, , from fabrik-test-beta"` (empty name is passed through literally)
3. **Given** `fabrik-test-beta` at HEAD after this change, **When** `go build ./...` is run, **Then** it exits 0
4. **Given** `fabrik-test-beta` at HEAD after this change, **When** `go test ./...` is run, **Then** all tests pass

---

### User Story 2 - Alpha prints greeting via beta (Priority: P2)

A developer running the `fabrik-test-alpha` binary after this change sees a greeting sourced from `fabrik-test-beta`. The binary accepts an optional name argument; if omitted it defaults to `"world"`.

**Why this priority**: Depends on Story 1 landing first. This is the primary observable outcome of the wiring.

**Independent Test**: Build and run the alpha binary with no arguments; assert stdout contains `"from fabrik-test-beta"`.

**Acceptance Scenarios**:

1. **Given** the alpha binary built at HEAD, **When** run with no arguments, **Then** stdout is `"Hello, world, from fabrik-test-beta\n"`
2. **Given** the alpha binary built at HEAD, **When** run with argument `"Alice"`, **Then** stdout is `"Hello, Alice, from fabrik-test-beta\n"`
3. **Given** the alpha binary built at HEAD, **When** run with no arguments in the automated integration test suite, **Then** `go test ./...` exits 0 and the test asserts output contains `"from fabrik-test-beta"`

---

### Edge Cases

- `GreetingFor("")` in beta: empty name is passed through literally, yielding `"Hello, , from fabrik-test-beta"` — no default substitution in beta
- Alpha with empty string argument `""` explicitly passed on CLI: passed to `GreetingFor("")` as-is (not treated as "no argument")
- Alpha with no arguments: `os.Args` has length 1, name defaults to `"world"`
- `go.mod` in both repos is pinned to `go 1.22` — must not be bumped

## Requirements *(mandatory)*

### Functional Requirements

#### fabrik-test-beta

- **FR-001**: `pkg/greeting` MUST export `GreetingFor(name string) string` returning `"Hello, " + name + ", from fabrik-test-beta"`
- **FR-002**: The existing `Greeting() string` function MUST be replaced by `GreetingFor` (it is not depended on by any external repo yet, so removal is safe)
- **FR-003**: `pkg/greeting/greeting_test.go` MUST include at least two test cases: one for empty name, one for a non-empty name
- **FR-004**: The package MUST remain fully exported and importable as `github.com/handarbeit/fabrik-test-beta/pkg/greeting`
- **FR-005**: `go.mod` MUST remain pinned to `go 1.22`
- **FR-006**: `go build ./...` and `go test ./...` MUST exit 0 in `fabrik-test-beta`

#### fabrik-test-alpha

- **FR-007**: `main.go` MUST accept a single optional positional command-line argument as the name, defaulting to `"world"` when omitted
- **FR-008**: `main.go` MUST import `github.com/handarbeit/fabrik-test-beta/pkg/greeting`
- **FR-009**: `main.go` MUST print the result of `greeting.GreetingFor(name)` to stdout (not stderr)
- **FR-010**: `go.mod` MUST declare `github.com/handarbeit/fabrik-test-beta` as a dependency at the version where `GreetingFor` exists; `go.sum` MUST be committed
- **FR-011**: `main_test.go` MUST include an integration test that: compiles the binary, runs it with no arguments, and asserts the output contains `"from fabrik-test-beta"`
- **FR-012**: `go.mod` MUST remain pinned to `go 1.22`
- **FR-013**: `go build ./...` and `go test ./...` MUST exit 0 in `fabrik-test-alpha`

### Key Entities

- **`GreetingFor(name string) string`**: The single exported function in `fabrik-test-beta/pkg/greeting`. Takes a name string, returns a greeting string of the form `"Hello, <name>, from fabrik-test-beta"`.
- **alpha binary**: The compiled `fabrik-test-alpha` executable. Accepts 0 or 1 positional args; sources its greeting from beta.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Running the alpha binary with no arguments prints exactly `"Hello, world, from fabrik-test-beta"` to stdout and exits 0
- **SC-002**: Running the alpha binary with argument `"Alice"` prints exactly `"Hello, Alice, from fabrik-test-beta"` to stdout and exits 0
- **SC-003**: `go test ./...` exits 0 in `fabrik-test-beta` with at least two greeting test cases
- **SC-004**: `go test ./...` exits 0 in `fabrik-test-alpha` with an integration test that asserts output contains `"from fabrik-test-beta"`
- **SC-005**: Both repos' CI workflows (`.github/workflows/ci.yml`) pass green after the respective changes merge
- **SC-006**: Fabrik's cross-repo decomposition feature is exercised: the Plan stage spawns a sub-issue in `fabrik-test-beta`, that sub-issue completes, and then alpha's implementation proceeds

## Assumptions

- The `arbeithand` PAT has write access to both `handarbeit/fabrik-test-alpha` and `handarbeit/fabrik-test-beta`
- `fabrik-test-beta` currently has no external dependents importing `Greeting()`, so removing it is non-breaking
- The beta sub-issue will be merged and a tagged/commit-pinned version available before alpha's `go.mod` is updated
- Both repos' CI runs `go build ./...` and `go test ./...` — no additional CI steps need to be satisfied
- `go 1.22` is sufficient for all code patterns used here; no language features requiring a newer version are needed

## Out of Scope

- Adding flags beyond a single positional name argument to the alpha binary
- Caching, streaming, or retry logic for the greeting call
- Any changes to `fabrik-test-beta` beyond replacing `Greeting()` with `GreetingFor(name string)`
- Resetting the test bed state, file-smoke runner, or further test scenarios (separate future issues)
- Publishing `fabrik-test-beta` to a Go module proxy; direct GitHub module references are sufficient

## Source References

- `fabrik-test-alpha/main.go` — current scaffold (prints to stderr, no beta dependency)
- `fabrik-test-alpha/main_test.go` — current placeholder test (no cross-repo assertion)
- `handarbeit/fabrik-test-beta/pkg/greeting/greeting.go` — target of beta change
- `handarbeit/fabrik-test-beta/pkg/greeting/greeting_test.go` — target of beta test update
- Fabrik issue #800 — the v0.0.66 cross-repo regression scenario this issue validates
