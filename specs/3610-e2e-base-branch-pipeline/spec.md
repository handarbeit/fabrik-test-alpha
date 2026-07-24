# Feature Specification: e2e base-branch pipeline (20260724-051545)

**Feature Branch**: `fabrik/issue-3610`
**Created**: 2026-07-24
**Status**: Draft
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes."

## Background

Fabrik supports overriding an issue's base branch via the `base:<branch>` label, so that the issue's worktree is forked from, rebased onto, and its PR targeted at a branch other than the repository default. handarbeit/fabrik#1046 tracked a regression in this contract; #1047 fixed issue↔PR linkage discovery when the PR targets a non-default base, and #1050 fixed the review-gate data feed (reviewer/CI status lookups) for PRs against a non-default base branch.

This issue is a recurrent single-repo end-to-end smoke test for the `base:<branch>` contract: it verifies that Fabrik correctly forks, rebases, and targets a PR at the non-default branch `e2e-base-branch-20260724-051545` throughout the full pipeline, that the engine still discovers the linked PR and its review/CI state via `closedByPullRequestsReferences` (the #1047/#1050 fix path), and that the pipeline auto-advances end-to-end under `fabrik:cruise`. The change is the simplest possible — one line appended to `README.md` — to keep noise minimal and isolate the base-branch mechanics as the thing under test.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Pipeline forks, rebases, and targets the non-default base branch throughout (Priority: P1)

A developer monitoring the Fabrik project board sees this issue advance through Research → Plan → Implement → Review → Validate, with the worktree forked from `e2e-base-branch-20260724-051545` (not the repository's `main`), rebases performed against that branch at Review/Validate, and the created PR targeting that branch rather than `main`.

**Why this priority**: This is the core `base:<branch>` contract under regression coverage per handarbeit/fabrik#1046. If the fork/rebase/PR-target base is wrong, the feature is broken regardless of anything else working.

**Independent Test**: After Implement runs, confirm on GitHub that the PR's base branch is `e2e-base-branch-20260724-051545`, not `main`.

**Acceptance Scenarios**:

1. **Given** the issue carries `base:e2e-base-branch-20260724-051545` before Research starts, **When** Fabrik creates the worktree, **Then** it is forked from `e2e-base-branch-20260724-051545`
2. **Given** the worktree is on the base-branch fork, **When** Implement creates the PR, **Then** the PR's base ref is `e2e-base-branch-20260724-051545` and its body contains `Closes #3610`
3. **Given** the PR is open against the non-default base, **When** Review or Validate runs, **Then** the stage rebases onto `e2e-base-branch-20260724-051545`, not `main`

---

### User Story 2 - Engine discovers PR linkage and review/CI state on the non-default base (Priority: P1)

The engine's issue→PR discovery (`closedByPullRequestsReferences`) and review-gate data feed (reviewer requests, CI check status) continue to work correctly when the linked PR targets a non-default base branch, exercising the #1047 and #1050 fixes.

**Why this priority**: #1047 and #1050 were fixes for regressions specific to non-default-base PRs; this is the regression test that must keep passing.

**Independent Test**: Confirm the engine transitions the issue through stages driven by PR state (e.g. Validate merge / CI gating) without falling back to polling errors or losing the PR link, despite the PR's base not being `main`.

**Acceptance Scenarios**:

1. **Given** the linked PR targets `e2e-base-branch-20260724-051545`, **When** the engine polls the board, **Then** it correctly discovers the PR via `closedByPullRequestsReferences` and reads its comments/reviews/CI status
2. **Given** `fabrik:cruise` is set (not `fabrik:yolo`), **When** Validate completes, **Then** the pipeline stops auto-advancing without merging the PR, and the PR remains open against the non-default base

---

### User Story 3 - Single-line README change only, no decomposition (Priority: P2)

The Implement stage appends exactly one HTML comment line to `README.md` and no other file is modified; the Plan stage does not decompose the issue into sub-issues.

**Why this priority**: A minimal, single-file diff isolates the base-branch mechanics as the only thing under test and keeps the smoke test noise-free.

**Independent Test**: Inspect the PR diff — it must show exactly one added line in `README.md` (excluding this spec file committed in Specify) and no other file changes.

**Acceptance Scenarios**:

1. **Given** the PR created by Implement, **When** the diff is inspected, **Then** the only changed file (excluding spec) is `README.md` with one line appended
2. **Given** `e2e-base-branch-20260724-051545` after the PR is merged, **When** the last line of `README.md` is read, **Then** it is exactly `<!-- base-branch-pipeline-20260724-051545 -->`

---

### Edge Cases

- `README.md` may or may not end with a trailing newline; the appended line must appear on its own line regardless
- The Plan stage must not decompose into sub-issues — this is a single-repo, single-file change
- The `base:<branch>` label must be read before Research and must not be added or removed during the run

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The worktree MUST be forked from `e2e-base-branch-20260724-051545`, not the repository's `main`
- **FR-002**: Implement MUST create the PR with base ref `e2e-base-branch-20260724-051545` and body containing `Closes #3610`
- **FR-003**: Review and Validate MUST rebase onto `e2e-base-branch-20260724-051545`, not `main`
- **FR-004**: The engine MUST correctly discover the linked PR and its review/CI state via `closedByPullRequestsReferences` while the PR targets the non-default base
- **FR-005**: Implement MUST append exactly `<!-- base-branch-pipeline-20260724-051545 -->` as the final line of `README.md`
- **FR-006**: No other source or configuration files MUST be modified
- **FR-007**: The Plan stage MUST NOT spawn sub-issues or decompose across repos
- **FR-008**: With `fabrik:cruise` set, the pipeline MUST auto-advance through all stages without human intervention, and MUST NOT auto-merge the PR after Validate completes
- **FR-009**: The `base:e2e-base-branch-20260724-051545` and `fabrik:cruise` labels MUST NOT be added or removed by any stage during the run

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Validate-complete on the project board without any human stage-advance action
- **SC-002**: The linked PR's base branch is `e2e-base-branch-20260724-051545` throughout the pipeline, never `main`
- **SC-003**: The engine correctly reads PR comments, reviews, and CI status for the non-default-base PR at every stage that needs them
- **SC-004**: The linked PR is still open (not merged) after Validate completes
- **SC-005**: `README.md` on `e2e-base-branch-20260724-051545` ends with the line `<!-- base-branch-pipeline-20260724-051545 -->` after a human merges the PR
- **SC-006**: No other files (excluding this specification) are changed in the merged PR diff

## Assumptions

- The `base:e2e-base-branch-20260724-051545` label was applied at filing and is already present; this run does not add or remove it
- `e2e-base-branch-20260724-051545` already exists as a branch on the remote before this issue is processed
- The `fabrik:cruise` label drives auto-advance through all stages but not auto-merge, per the existing cruise contract
- This is a single-repo operation — `handarbeit/fabrik-test-alpha` only; no cross-repo work
- CI on this repo does not block on a README-only change
- A human reviewer will manually merge the PR after Validate completes to close the issue

## Out of Scope

- Any changes beyond appending the single comment line to `README.md`
- Cross-repo sub-issue spawning or decomposition
- Updating any Go source files, tests, or configuration
- Auto-merge behavior (that is the `fabrik:yolo` contract, not cruise)
- Testing `base:<branch>` fallback behavior when the branch does not exist on the remote (out of scope — this run assumes the branch already exists)

## Source References

- `README.md` — the only file modified by this issue
- handarbeit/fabrik#1046 — regression tracking issue for the `base:<branch>` pipeline contract
- handarbeit/fabrik#1047 — issue↔PR linkage fix for non-default-base PRs
- handarbeit/fabrik#1050 — review-gate data feed fix for non-default-base PRs
