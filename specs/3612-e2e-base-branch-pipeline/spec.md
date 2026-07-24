# Feature Specification: e2e base-branch pipeline (20260724-054436)

**Feature Branch**: `fabrik/issue-3612`
**Created**: 2026-07-24
**Status**: Draft
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: append a single HTML comment line to README.md at the very end of the file: `<!-- base-branch-pipeline-20260724-054436 -->`. Single repo only, targeting the non-default base branch e2e-base-branch-20260724-054436 (already set via the base:<branch> label applied at filing)."

## Background

Fabrik supports overriding the worktree base branch for an issue via the `base:<branch>` label, so that work forks from, rebases onto, and targets PRs at a non-default branch instead of the repository's default branch (`main`). Two prior fixes changed how this path behaves:

- handarbeit/fabrik#1047 — issue↔PR linkage when the PR targets a non-default base branch
- handarbeit/fabrik#1050 — review-gate data feed correctness under a non-default base branch

This issue is a synthetic, trivial-payload end-to-end run whose purpose is to exercise the full Fabrik pipeline (Specify → Research → Plan → Implement → Review → Validate) against a non-default base branch, confirming issue↔PR linkage and the review-gate data feed continue to work correctly on this path. The `base:e2e-base-branch-20260724-054436` label was applied at filing and must not be added or removed by any stage. The `fabrik:cruise` label drives auto-advance through the pipeline without auto-merge.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Pipeline completes correctly against a non-default base branch (Priority: P1)

As the Fabrik maintainer, I need confidence that an issue filed with a `base:<branch>` label flows through the entire SDLC pipeline — forking, rebasing, and opening a PR against that branch rather than `main` — with correct issue↔PR linkage and review-gate behavior throughout.

**Why this priority**: This is the sole purpose of the issue; there is no secondary functionality.

**Independent Test**: File an issue with `base:<branch>` and `fabrik:cruise` labels, let it run unattended through all stages, and confirm the resulting PR targets the named branch, links back to this issue, and the review-gate correctly reflects PR state.

**Acceptance Scenarios**:

1. **Given** the issue carries `base:e2e-base-branch-20260724-054436`, **When** the worktree is created, **Then** it forks from and is based on `e2e-base-branch-20260724-054436`, not `main`.
2. **Given** the Implement stage has committed the trivial change, **When** the draft PR is created, **Then** the PR targets `e2e-base-branch-20260724-054436` and its body contains `Closes #3612`.
3. **Given** the PR is open, **When** Fabrik looks up PR comments and review state via `closedByPullRequestsReferences`, **Then** it correctly resolves to this issue's linked PR (issue↔PR linkage per #1047).
4. **Given** the Review stage completes, **When** the review-gate evaluates outstanding reviewer state, **Then** it reflects the correct, current PR review data (per #1050).
5. **Given** all stages complete, **When** Validate finishes, **Then** the PR is mergeable against `e2e-base-branch-20260724-054436` and the pipeline reaches Done without any stage having added or removed the `base:` label.

### Edge Cases

- None beyond the standard pipeline flow — this is a minimal-surface-area regression check, not a test of edge-case handling.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20260724-054436 -->`, to the end of `README.md`, with no other file or content changes.
- **FR-002**: The Plan stage MUST NOT decompose this issue into sub-tasks or child issues — it is a single trivial edit.
- **FR-003**: The worktree and PR MUST use `e2e-base-branch-20260724-054436` as the base/target branch throughout the pipeline, not the repository default (`main`).
- **FR-004**: The `base:e2e-base-branch-20260724-054436` and `fabrik:cruise` labels MUST NOT be added to or removed from the issue by any stage.
- **FR-005**: The PR body MUST contain `Closes #3612` so Fabrik can discover PR comments per the standard PR lifecycle convention.

### Key Entities

- Not applicable — no data entities are introduced by this change.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The pipeline reaches Done with a merged (or merge-ready) PR targeting `e2e-base-branch-20260724-054436`, containing only the single appended README.md line.
- **SC-002**: No stage reports issue↔PR linkage failures or review-gate data errors attributable to the non-default base branch.

## Assumptions

- The branch `e2e-base-branch-20260724-054436` already exists on the remote (implied by the label having been set successfully at filing).
- `fabrik:cruise` auto-advances the pipeline through all stages without requiring manual approval between stages, and without triggering auto-merge (which is reserved for `fabrik:yolo`).

## Out of Scope

- Any change to Fabrik's `base:<branch>` handling logic itself — this issue only exercises the existing, already-fixed behavior.
- Verification of cross-repo or multi-issue base-branch scenarios.

## Source References

- handarbeit/fabrik#1046, #1047, #1050
