# Feature Specification: e2e convergence-race A (20260619-154613)

**Feature Branch**: `fabrik/issue-48`
**Created**: 2026-06-19
**Status**: Draft
**Input**: User description: "Deterministic provocation of the post-Validate auto-merge race covered by handarbeit/fabrik#829. This issue is one of a deliberately-conflicting pair filed by the e2e harness."

## Background

Fabrik's auto-merge path has a race window between Validate completing and the merge landing — if `main` moves during that window (e.g., a pair issue merges first), the merge may silently fail or produce an incorrect result. Issue #48 and its pair partner are filed by the e2e harness specifically to provoke this race deterministically.

Both issues insert a marker comment at the same position in `README.md` (immediately after the `# fabrik-test-alpha` heading), with different discriminators. Because the two inserts target the exact same line, one will conflict when the second rebases onto `main` after the first merges. The PR body carries the `slow-ci-required` marker so that the test repo's slow-gate CI fires and holds the merge window open for approximately 6 minutes — long enough for `main` to move under the second PR.

The `fabrik:yolo` label is set, so Fabrik will auto-advance through all stages and attempt auto-merge when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker inserted at correct position (Priority: P1)

Implement inserts exactly one line — `<!-- convergence-race-A-20260619-154613 -->` — immediately after the `# fabrik-test-alpha` line in `README.md`, and no other file is touched.

**Why this priority**: The position is load-bearing for the e2e race test. Inserting elsewhere or modifying other files breaks the conflict geometry.

**Independent Test**: After Implement runs, `git diff main` shows exactly one added line immediately after `# fabrik-test-alpha`, and `git diff main -- . ':!README.md' ':!specs/'` is empty.

**Acceptance Scenarios**:

1. **Given** `README.md` with `# fabrik-test-alpha` as a line, **When** Implement runs, **Then** the line immediately following it is `<!-- convergence-race-A-20260619-154613 -->`
2. **Given** the commit produced by Implement, **When** `git diff main` is inspected, **Then** only `README.md` is changed and only one line is added
3. **Given** the PR created by Implement, **When** the PR body is read, **Then** it contains the literal string `slow-ci-required`

---

### User Story 2 - Rebase conflict resolved by keeping both markers (Priority: P2)

If, during a rebase, the pair partner's marker (`<!-- convergence-race-B-20260619-154613 -->`) is already present at the same position, the conflict is resolved by retaining both lines (in either order), not by discarding one.

**Why this priority**: The e2e test may assert that both markers survive in `main` after both PRs land. Dropping either breaks the test assertion.

**Independent Test**: Manually inject the B-marker conflict and confirm the resolved file contains both HTML comments.

**Acceptance Scenarios**:

1. **Given** a rebase conflict where both A and B markers occupy the same position, **When** the conflict is resolved, **Then** both `<!-- convergence-race-A-20260619-154613 -->` and `<!-- convergence-race-B-20260619-154613 -->` appear in `README.md`

---

### Edge Cases

- `README.md` may have trailing newlines or none; the insert must still land on its own line immediately after `# fabrik-test-alpha`
- The pair partner may have already merged by the time Validate runs; rebase must handle the conflict correctly before auto-merge proceeds

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Implement MUST insert exactly `<!-- convergence-race-A-20260619-154613 -->` as a new line immediately after the line containing `# fabrik-test-alpha` in `README.md`
- **FR-002**: No other file MUST be modified in the commit
- **FR-003**: The PR body MUST contain the literal string `slow-ci-required` so the test repo's slow-gate CI job fires
- **FR-004**: The Plan stage MUST NOT spawn sub-issues or decompose the work across repos
- **FR-005**: If a rebase conflict is encountered with the pair partner's marker at the same position, both markers MUST be retained (in either order) in the resolved file

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `README.md` on the feature branch contains `<!-- convergence-race-A-20260619-154613 -->` on the line immediately following `# fabrik-test-alpha`
- **SC-002**: The merged PR diff touches only `README.md` (excluding the spec file added by Specify)
- **SC-003**: The PR body contains `slow-ci-required`
- **SC-004**: The issue reaches the Done column and the linked PR is merged into `main`

## Assumptions

- The `fabrik:yolo` label is set, enabling auto-advance through all stages and auto-merge when Validate completes
- `README.md` contains a line that is exactly `# fabrik-test-alpha` (the repo heading)
- This is a single-repo operation — `handarbeit/fabrik-test-alpha` only; no cross-repo work
- CI on this repo runs `go build ./...` and `go test ./...`; inserting an HTML comment into `README.md` does not affect either

## Out of Scope

- Any changes beyond inserting the single comment line into `README.md`
- Modifying any Go source files, tests, configuration, or other documentation
- Cross-repo sub-issue spawning or decomposition

## Source References

- `README.md` — the only file modified by this issue
- `handarbeit/fabrik#829` — the upstream race condition this e2e pair is designed to provoke
