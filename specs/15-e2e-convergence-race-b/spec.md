# Feature Specification: e2e convergence-race B (20260525-225522)

**Feature Branch**: `fabrik/issue-15`
**Created**: 2026-05-25
**Status**: Draft
**Input**: User description: "Deterministic provocation of the post-Validate auto-merge race covered by handarbeit/fabrik#829. This issue is one of a deliberately-conflicting pair filed by the e2e harness."

## Background

Fabrik's auto-merge path (post-Validate, with `fabrik:yolo`) can race against a concurrent merge landing on `main` while the PR is waiting for CI. Issue handarbeit/fabrik#829 tracks this race condition. To produce a deterministic test window, the e2e harness files two sibling issues that each insert a different HTML comment at the **same position** in `README.md` (immediately after the `# fabrik-test-alpha` heading). Because both inserts target the same line, one PR's rebase onto the other's merged result produces a textual conflict — exactly the scenario that exposes the race.

This issue is the "B" member of the pair. Its partner ("A") inserts a different discriminator at the same position. The CI slow-gate (`slow-ci-required` in the PR body) creates a deterministic 6-minute window during which `main` can move under the waiting PR.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Insert convergence-race-B marker at the required position (Priority: P1)

The e2e harness expects that, after this issue completes, `README.md` on `main` contains the line `<!-- convergence-race-B-20260525-225522 -->` immediately after `# fabrik-test-alpha` (possibly alongside the pair-partner's line if both are present due to conflict resolution).

**Why this priority**: This is the sole purpose of the issue. The exact position and exact text are load-bearing for the e2e conflict detection logic.

**Independent Test**: Check out `main` after the PR merges and confirm that line 2 of `README.md` is `<!-- convergence-race-B-20260525-225522 -->` (or that both pair-member lines are present if the partner merged first).

**Acceptance Scenarios**:

1. **Given** `README.md` with `# fabrik-test-alpha` as its first line, **When** Implement runs, **Then** the line `<!-- convergence-race-B-20260525-225522 -->` is inserted immediately after it (as line 2)
2. **Given** the PR created by Implement, **When** its body is read, **Then** it contains the literal string `slow-ci-required` so the CI slow-gate fires
3. **Given** a rebase conflict where the pair-partner's marker is already present at line 2, **When** the conflict is resolved, **Then** BOTH marker lines are kept (in either order) and the rebase completes cleanly
4. **Given** `main` after the PR merges, **When** `README.md` is read, **Then** no files other than `README.md` (and this spec) differ from the pre-issue state

---

### Edge Cases

- If the pair-partner PR merges first, a rebase conflict will occur at the insert position — resolution must keep both lines, not discard either
- The insert position is **immediately after** the `# fabrik-test-alpha` line, not at the end of the file
- No other file may be modified; the change is strictly single-line in `README.md`

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Implement MUST insert exactly `<!-- convergence-race-B-20260525-225522 -->` as a new line immediately after the line `# fabrik-test-alpha` in `README.md`
- **FR-002**: No other source or configuration files MUST be modified
- **FR-003**: The PR body MUST contain the literal string `slow-ci-required` so the test repo's CI slow-gate activates
- **FR-004**: The Plan stage MUST NOT spawn sub-issues or decompose — this is a single-repo, single-file, single-line change
- **FR-005**: If a rebase conflict is encountered at the insert position (pair-partner marker present), BOTH lines MUST be retained in either order

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches the Done column on the project board
- **SC-002**: The linked PR is merged into `main`
- **SC-003**: `README.md` on `main` contains `<!-- convergence-race-B-20260525-225522 -->` on the line immediately after `# fabrik-test-alpha`
- **SC-004**: No other files (excluding this specification) are changed in the merged PR diff
- **SC-005**: The PR body contains the literal string `slow-ci-required`

## Assumptions

- The `fabrik:yolo` label is set on the issue, enabling auto-advance and auto-merge through all stages
- This is a single-repo operation — `handarbeit/fabrik-test-alpha` only; no cross-repo work
- `README.md` begins with `# fabrik-test-alpha` as its first line; the insert target is unambiguous
- The pair-partner issue (convergence-race-A) may or may not have merged before this PR rebases; both orderings are valid

## Out of Scope

- Any changes beyond inserting the single comment line into `README.md`
- Cross-repo sub-issue spawning or decomposition
- Updating any Go source files, tests, or configuration

## Source References

- `README.md` — the only file modified by this issue
- handarbeit/fabrik#829 — the upstream race condition this e2e test exercises
