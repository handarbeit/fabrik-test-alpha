# Feature Specification: e2e base-branch pipeline (20260724-015821)

**Feature Branch**: `fabrik/issue-3608`
**Created**: 2026-07-23
**Status**: Draft
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Append a single HTML comment line to README.md at the very end of the file: `<!-- base-branch-pipeline-20260724-015821 -->`. That is the entire change. One file, one line. Plan should NOT decompose."

## Background

Fabrik supports overriding the default worktree base branch for a single issue via the `base:<branch>` label, so that forking, rebasing, and PR targeting all happen against a non-default branch instead of the repository default. A prior defect (handarbeit/fabrik#1046) affected this pipeline; #1047 fixed issue↔PR linkage under a non-default base branch, and #1050 fixed the review-gate data feed under the same condition.

This issue exists purely to exercise that pipeline end-to-end: run a real issue through Specify → Research → Plan → Implement → Review → Validate while targeting a non-default base branch (`e2e-base-branch-20260724-015821`, already set via the `base:<branch>` label applied at filing), and confirm the full flow — worktree forking, PR creation and linkage, review-gate behavior, and merge — completes correctly against that branch. The actual code change is intentionally trivial so the test isolates pipeline mechanics rather than implementation complexity.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Base-branch pipeline runs end-to-end (Priority: P1)

As a Fabrik maintainer validating the `base:<branch>` feature, I want a minimal issue to flow through the full SDLC pipeline while targeting a non-default base branch, so that I can confirm worktree forking, PR issue-linkage, and the review-gate data feed all function correctly under that condition.

**Why this priority**: This is the entire purpose of the issue — there is no other scope.

**Independent Test**: Run the issue through the pipeline and confirm the resulting PR is forked from and targets `e2e-base-branch-20260724-015821`, correctly links back to issue #3608, passes through the review gate, and merges successfully.

**Acceptance Scenarios**:

1. **Given** the `base:e2e-base-branch-20260724-015821` label is set on issue #3608, **When** the Implement stage creates the PR, **Then** the PR's base branch is `e2e-base-branch-20260724-015821` (not the repository's actual default branch) and its body contains `Closes #3608`.
2. **Given** the PR is open, **When** Fabrik looks up PR comments/reviews for the issue, **Then** it correctly discovers them via the issue↔PR linkage (regression check for #1047).
3. **Given** the PR reaches the review-gate stage, **When** review data is fetched, **Then** the review-gate data feed reflects the correct state (regression check for #1050).
4. **Given** all stages complete, **When** Validate finishes, **Then** the PR merges cleanly into `e2e-base-branch-20260724-015821`.

### Edge Cases

- None beyond the above — this is a minimal single-file regression check, not a feature with edge-case branching.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The implementation MUST append exactly one line, `<!-- base-branch-pipeline-20260724-015821 -->`, to the end of `README.md`, with no other file or content changes.
- **FR-002**: The Plan stage MUST NOT decompose this issue into sub-issues or multiple tasks — it is a single atomic change.
- **FR-003**: The pipeline MUST fork the working branch from, and the resulting PR MUST target, `e2e-base-branch-20260724-015821` rather than the repository's actual default branch, per the existing `base:<branch>` label already applied to this issue.
- **FR-004**: The resulting PR MUST correctly link to issue #3608 (e.g. via `Closes #3608` in the PR body) so that issue↔PR discovery works under the non-default base branch.

### Key Entities

- Not applicable — no data entities are introduced by this change.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `README.md` on the merged branch ends with the exact line `<!-- base-branch-pipeline-20260724-015821 -->` and no other content changed.
- **SC-002**: The PR created for this issue has `e2e-base-branch-20260724-015821` as its base and is merged into it.
- **SC-003**: The issue progresses through every pipeline stage (Specify → Research → Plan → Implement → Review → Validate) to Done without manual intervention beyond what `fabrik:cruise` already permits.

## Assumptions

- The `e2e-base-branch-20260724-015821` branch already exists on the remote (created when the `base:<branch>` label was applied at filing).
- No labels beyond `fabrik:cruise` and `base:e2e-base-branch-20260724-015821` should be added or removed as part of this work.

## Out of Scope

- Any change to README.md content beyond the single specified comment line.
- Any fix to the pipeline itself — this issue is verification only; if a regression is found, it should be filed and fixed separately.

## Source References

- handarbeit/fabrik#1046 (original defect)
- handarbeit/fabrik#1047 (issue↔PR linkage fix)
- handarbeit/fabrik#1050 (review-gate data feed fix)
