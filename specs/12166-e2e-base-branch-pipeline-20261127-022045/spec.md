# Feature Specification: e2e base-branch pipeline (20261127-022045)

**Feature Branch**: `fabrik/issue-12166`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261127-022045 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261127-022045 (already set via
the base:<branch> label applied at filing — do not add or remove labels).


## Labels

fabrik:cruise, base:e2e-base-branch-20261127-022045"

## Background

This issue is a live end-to-end exercise of the pipeline's handling of issues that target a non-default base branch (the `base:<branch>` label). It is regression coverage for handarbeit/fabrik#1046, validating the fixes in #1047 (issue-to-PR linkage for non-default-base PRs) and #1050 (the review gate's data feed for such PRs). The code change itself is deliberately trivial, so that any failure observed is attributable to the pipeline's base-branch handling and not to the work being done. The issue carries `fabrik:cruise` (auto-advance through all stages without auto-merge) and `base:e2e-base-branch-20261127-022045`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

A maintainer running the e2e suite wants a single, trivially verifiable change to travel through every pipeline stage against the non-default base branch, so that the base-branch contract (branching, PR targeting, issue-PR linkage, review gate) can be checked end to end.

**Why this priority**: It is the only deliverable; the whole purpose of the issue is that this change flows through the pipeline correctly.

**Independent Test**: Inspect the resulting PR: its base is `e2e-base-branch-20261127-022045`, its diff is exactly one added line in `e2e/markers/base-branch-pipeline.md`, and the file's last line is `<!-- base-branch-pipeline-20261127-022045 -->`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the change is implemented, **Then** the line `<!-- base-branch-pipeline-20261127-022045 -->` is appended as the last line and no other line in the file is modified.
2. **Given** the file does not exist on the base branch, **When** the change is implemented, **Then** the file is created containing the marker line.
3. **Given** the issue carries `base:e2e-base-branch-20261127-022045`, **When** the pipeline opens a PR, **Then** the PR targets that branch and is linked to this issue.

---

### Edge Cases

- The marker line is already present at the end of the file (e.g. after a retry): it must not be appended a second time.
- The existing file does not end in a trailing newline: the marker still appears on its own line, as the final line.
- Labels on the issue must not be added or removed by the work.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- base-branch-pipeline-20261127-022045 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created first, containing the marker line.
- **FR-003**: The change touches no file other than `e2e/markers/base-branch-pipeline.md` (the Specify stage's own `specs/` projection excepted).
- **FR-004**: Work is branched from, and the PR targets, the non-default base branch `e2e-base-branch-20261127-022045`.
- **FR-005**: The work is not decomposed into child issues.
- **FR-006**: No labels are added to or removed from the issue by the change.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR's diff against its base consists of exactly one added line in exactly one file (the marker file), apart from the spec projection.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261127-022045` and the PR is linked to issue #12166.
- **SC-003**: The issue passes through all pipeline stages with no pause or manual intervention caused by base-branch handling.

## Assumptions

- The default branch named in the stage context and the `base:` label name the same branch, so there is no conflict between them.
- "Very end of the file" means the marker is the final line, with a single trailing newline.
- Idempotency (not duplicating the marker on retry) is an acceptable reading of "single line".

## Out of Scope

- Any change to Fabrik engine code, stage configuration or labels.
- Auto-merging the PR (`fabrik:cruise` does not merge).
- Any additional marker files or content.
- Decomposition into sub-issues.

## Source References

- handarbeit/fabrik#1046 (base-branch pipeline regression), #1047 (issue-PR linkage fix), #1050 (review-gate data feed fix)
