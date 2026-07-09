# Feature Specification: e2e convergence-race A (20260709-142728)

**Feature Branch**: `fabrik/issue-3374`
**Created**: 2026-07-09
**Status**: Draft
**Input**: User description: "Deterministic provocation of the post-Validate auto-merge race covered by handarbeit/fabrik#829. This issue is one of a deliberately-conflicting pair filed by the e2e harness."

## Background

Fabrik's post-Validate auto-merge path can race when two issues both target the same position in `main` and one merges while the other is in CI. Issue handarbeit/fabrik#829 tracks this race condition. This issue is the A-side of a deliberately-conflicting pair that the e2e harness files to provoke and verify the race deterministically.

Both pair members insert a single HTML comment immediately after the `# fabrik-test-alpha` heading in `README.md`. Because they target the identical line position, whichever merges second will encounter a textual conflict on rebase. The CI slow-gate (`slow-ci-required` in the PR body) holds the first-merged PR in CI for ~10 minutes, creating a deterministic window for `main` to advance under the second PR before it completes Validate.

The `fabrik:yolo` label is present, so Fabrik will auto-advance through all stages and auto-merge the PR on Validate completion.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Convergence-race A side executes and resolves conflict (Priority: P1)

The e2e harness confirms that when both pair members run concurrently, the A-side PR either merges cleanly (if it wins the race) or rebases successfully after keeping both markers (if B-side merged first), and both reach Done with their respective markers present on `main`.

**Why this priority**: This is the sole purpose of the issue — acting as the A-side of the convergence race. There is no secondary story.

**Independent Test**: After the issue moves to Done, check out `main` and confirm `README.md` contains the line `<!-- convergence-race-A-20260709-142728 -->` immediately after `# fabrik-test-alpha` (possibly preceded or followed by the B-side marker on adjacent lines).

**Acceptance Scenarios**:

1. **Given** the issue at Specify, **When** Fabrik processes it with `fabrik:yolo`, **Then** it advances through all stages without human input
2. **Given** the Implement stage, **When** it runs, **Then** `README.md` has `<!-- convergence-race-A-20260709-142728 -->` inserted on the line immediately after `# fabrik-test-alpha`, and no other file is modified
3. **Given** the PR created by Implement, **When** the PR body is read, **Then** it contains the literal text `slow-ci-required` so the slow-gate CI job fires
4. **Given** a rebase conflict where B-side's marker is already present at the same position, **When** the conflict is resolved, **Then** both markers are kept (in either order), and the rebase completes cleanly
5. **Given** Validate completing, **When** auto-merge runs, **Then** the PR merges into `main` and the issue moves to Done

---

### Edge Cases

- If the B-side marker is already on `main` when Implement runs (B-side merged very quickly), there is no conflict — just insert the A-side marker on the immediately following line
- If a rebase conflict occurs with both markers in the conflict block, keep both; do not discard either
- The `slow-ci-required` text in the PR body must appear as a standalone token (not inside a code fence or HTML comment) so the CI matcher recognises it

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Implement MUST insert `<!-- convergence-race-A-20260709-142728 -->` as a new line immediately after the line `# fabrik-test-alpha` in `README.md`
- **FR-002**: No other lines in `README.md` and no other files (except this specification) MUST be modified
- **FR-003**: The PR body MUST contain the literal text `slow-ci-required` so the test repo's slow-gate CI job fires
- **FR-004**: The Plan stage MUST NOT spawn sub-issues or decompose the work
- **FR-005**: If a rebase conflict is encountered at the insertion position (because the B-side marker is already there), the conflict MUST be resolved by keeping both markers in either order

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches the Done column on the project board
- **SC-002**: The linked PR is merged into `main`
- **SC-003**: `README.md` on `main` contains the line `<!-- convergence-race-A-20260709-142728 -->` immediately after `# fabrik-test-alpha` (possibly with the B-side marker adjacent)
- **SC-004**: No files other than `README.md` (and this specification) are changed in the merged PR diff

## Assumptions

- The `fabrik:yolo` label is set, enabling auto-advance and auto-merge through all stages
- This is a single-repo operation — `handarbeit/fabrik-test-alpha` only
- The pair partner (B-side) is filed as a separate issue targeting the same insertion position with a different discriminator but the same timestamp `20260709-142728`
- CI on this repo includes a slow-gate job that is triggered by the presence of `slow-ci-required` in the PR body and delays completion by ~10 minutes

## Out of Scope

- Any changes beyond inserting the single comment line into `README.md`
- Cross-repo sub-issue spawning or decomposition
- Updating any Go source files, tests, or configuration

## Source References

- `README.md` — the only file modified by this issue
- handarbeit/fabrik#829 — the upstream race condition this test is designed to provoke
