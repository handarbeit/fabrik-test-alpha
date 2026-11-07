# Feature Specification: e2e base-branch pipeline (20261107-195139)

**Feature Branch**: `fabrik/issue-10812`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261107-195139 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261107-195139 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is an end-to-end regression probe for the `base:<branch>` contract, in which an issue carries a label naming a non-default base branch and the whole pipeline (worktree, PR creation, issue-to-PR linkage, review-gate data) must operate against that branch instead of the repository default. It guards the behaviour covered by handarbeit/fabrik#1046 and validates the fixes for #1047 (issue↔PR linkage) and #1050 (review-gate data feed). The code change is deliberately trivial so that any failure observed is attributable to the pipeline's handling of the non-default base branch, not to the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator files an issue labelled `base:e2e-base-branch-20261107-195139`. The pipeline carries it through every stage, producing a single-line change to the marker file via a PR that targets that base branch.

**Why this priority**: This is the entire change and the artifact the e2e check inspects.

**Independent Test**: Inspect `e2e/markers/base-branch-pipeline.md` on the PR branch: its final line is exactly `<!-- base-branch-pipeline-20261107-195139 -->` and the diff against the base branch is that one added line (or the file creation with only that line).

**Acceptance Scenarios**:

1. **Given** the marker file already exists on the base branch, **When** the change is made, **Then** exactly one line is appended at the very end and no existing content is altered.
2. **Given** the marker file does not exist, **When** the change is made, **Then** the file is created containing only the marker line.

---

### User Story 2 - Pipeline honours the non-default base (Priority: P1)

The resulting PR targets `e2e-base-branch-20261107-195139`, not the repository default, and the engine can still discover it for the issue.

**Why this priority**: This is the regression the e2e exists to protect (#1046, #1047, #1050).

**Independent Test**: Check the PR's base ref and confirm the engine discovers the PR for the issue and surfaces its review state to the review gate.

**Acceptance Scenarios**:

1. **Given** the issue carries the `base:` label, **When** the PR is created, **Then** its base branch is `e2e-base-branch-20261107-195139`.
2. **Given** the PR targets a non-default base, **When** the pipeline evaluates linkage and review state, **Then** the issue↔PR link is found and review-gate data is available.

---

### Edge Cases

- The marker file exists but its last line lacks a trailing newline: the appended line must still sit on its own line.
- The marker file already contains the same marker line: the change must not add a duplicate.
- The labels on the issue must not be added to or removed during the run.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends the line `<!-- base-branch-pipeline-20261107-195139 -->` at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If the marker file does not exist, it is created and contains only that line.
- **FR-003**: The change touches exactly one file and adds exactly one line; no other files change.
- **FR-004**: The resulting PR targets the base branch `e2e-base-branch-20261107-195139`.
- **FR-005**: The work is a single unit and is not decomposed into child issues.
- **FR-006**: No labels are added to or removed from the issue as part of this work.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261107-195139` consists of exactly one added line in exactly one file.
- **SC-002**: The PR's base branch equals `e2e-base-branch-20261107-195139` and the issue is linked to it.
- **SC-003**: The issue completes the pipeline with no child issues spawned.

## Assumptions

- The base branch `e2e-base-branch-20261107-195139` already exists on the remote.
- The `fabrik:cruise` label is intentional and drives auto-advance; it is left untouched.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Modifying issue labels or the e2e harness itself.
- Decomposing the work into child issues.
