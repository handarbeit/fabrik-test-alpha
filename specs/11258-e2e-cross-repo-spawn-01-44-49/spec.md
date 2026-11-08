# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-11258`
**Created**: 2026-11-08
**Status**: Specified
**Input**: User description: the original request, verbatim:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-beta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-beta (the beta repo)
> - Add an exported function HelloE2E202611080644494414() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261108-064449-4414".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-alpha (this repo)
> - Update main.go to call greeting.HelloE2E202611080644494414() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This is an end-to-end regression scenario for cross-repo decomposition. It covers two earlier fixes:

- **#803**: on-demand initialisation of a spawn-target repo.
- **#1308**: a per-run-unique fixture token, so each run's final assertion can tell its own output from stale output of earlier runs.

The feature needs coordinated changes in two repositories with a hard ordering dependency:

- The **beta** repo (`handarbeit/fabrik-test-beta`) must gain a new greeting function.
- The **alpha** repo (`handarbeit/fabrik-test-alpha`, this one) must call that function.

Alpha cannot compile against a function that does not exist yet, so beta's change has to land first. The pipeline is expected to split the work into a sub-issue in the beta repo, with this issue blocked on it, rather than treat it as a single-repo change.

The function name and string literal contain a per-run-unique token. They must be used exactly as given. The scenario's final assertion checks that this exact text appears in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

The beta repo's greeting package gains an exported function that returns the per-run-unique literal, with a test covering it.

**Why this priority**: Alpha depends on it, and the scenario's final assertion checks for this text in the beta repo.

**Independent Test**: Run the beta repo's tests for the greeting package. The new test passes and asserts the exact literal.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202611080644494414()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261108-064449-4414`.
2. **Given** the beta repo's greeting test file, **When** the package tests run, **Then** a test for the new function is present and passes.

---

### User Story 2 - Alpha calls the new function and prints the result (Priority: P2)

The alpha repo's `main.go` calls the beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It is the consumer side of the dependency. It can only be delivered after User Story 1 has landed.

**Independent Test**: Build and run alpha once beta's function is available. The output includes `e2e-cross-repo-spawn-20261108-064449-4414`, and the existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** beta's function has landed, **When** alpha is built, **Then** it compiles, with module dependencies tidied if needed.
2. **Given** alpha is run, **When** `main` executes, **Then** it prints the value returned by `greeting.HelloE2E202611080644494414()` and still does everything it did before.

---

### User Story 3 - Work is decomposed across repos in dependency order (Priority: P1)

Plan decomposes the work into a beta-repo sub-issue, and alpha's work is sequenced after it.

**Why this priority**: This is the behaviour under test (#803). Without it the scenario does not exercise cross-repo spawn.

**Independent Test**: After Plan, a child issue exists in the beta repo covering User Story 1. This issue is blocked by it, and the child lands before alpha's change does.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** its output contains a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-beta` for the beta-side work.
2. **Given** the child issue was spawned, **When** the pipeline proceeds, **Then** alpha's change is not implemented or landed until the beta child has landed.

---

### Edge Cases

- The function name and literal contain a per-run token. Any altered, regenerated or shortened form breaks the final assertion (#1308), so both must be copied verbatim into the child issue and the code.
- The beta repo may not yet be initialised for spawn use; on-demand initialisation of the spawn target is part of what is being verified (#803).
- If alpha's module tooling needs no change, `go mod tidy` is a no-op. It is run only if needed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-beta`) MUST export `HelloE2E202611080644494414() string` from `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202611080644494414()` MUST return exactly the string `e2e-cross-repo-spawn-20261108-064449-4414`.
- **FR-003**: The beta repo MUST include a test for that function in `pkg/greeting/greeting_test.go` that asserts the exact returned string.
- **FR-004**: The alpha repo's `main.go` MUST call `greeting.HelloE2E202611080644494414()` and print the result, while keeping all of its existing behaviour.
- **FR-005**: The alpha repo MUST build successfully with the new import. `go mod tidy` MUST be run if module files need updating.
- **FR-006**: The beta-repo change MUST land before the alpha-repo change.
- **FR-007**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets `handarbeit/fabrik-test-beta` and covers FR-001 to FR-003.
- **FR-008**: The function name and string literal MUST be used exactly as given in this spec, with no renamed or regenerated token, in the spawned child's description and in the code.

### Key Entities *(if applicable)*

- **Beta repo** (`handarbeit/fabrik-test-beta`): the repo that hosts the greeting package and the new function; it is the spawn target.
- **Alpha repo** (`handarbeit/fabrik-test-alpha`): this repo; it consumes the beta function from `main.go`.
- **Per-run-unique token**: the `20261108-064449-4414` / `202611080644494414` fragments embedded in the function name and literal.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains `HelloE2E202611080644494414` and the literal `e2e-cross-repo-spawn-20261108-064449-4414`.
- **SC-002**: The beta repo's greeting package tests pass, including the new test.
- **SC-003**: The alpha repo builds and its tests pass, and running it prints `e2e-cross-repo-spawn-20261108-064449-4414`.
- **SC-004**: A sub-issue targeting the beta repo was created from this issue's Plan, and it landed before alpha's change.

## Assumptions

- The alpha repo already imports the beta repo's `pkg/greeting` package or can add the import. This is a Plan and Research concern.
- The per-run token in the issue body is authoritative. Nothing needs to regenerate or validate it.
- "Print the result" means writing the returned string to standard output, alongside existing output.
- The existing greeting behaviour in both repos stays unchanged.

## Out of Scope *(optional)*

- Any change to beta or alpha behaviour beyond the single new function, its test and the call in `main.go`.
- Changes to the spawn, decomposition or fixture-token mechanisms themselves. This issue only exercises them.
- Releasing or versioning the beta repo beyond what landing the change requires.

## Source References *(optional)*

- #803: on-demand spawn-target repo init.
- #1308: per-run-unique fixture token.
