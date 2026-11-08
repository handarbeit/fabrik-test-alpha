# Feature Specification: e2e base-branch pipeline (20261108-062810)

**Feature Branch**: `fabrik/issue-11233`
**Created**: 2026-11-08
**Status**: Specified
**Input**: User description: 

> ## Goal
>
> End-to-end verification of the base:<branch> (non-default base branch)
> pipeline contract — regression coverage for handarbeit/fabrik#1046,
> validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
> fixes.
>
> ## Trivial change
>
> Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):
>
> `<!-- base-branch-pipeline-20261108-062810 -->`
>
> That is the entire change. One file, one line. Plan should NOT decompose.
>
> ## Scope
>
> Single repo only, targeting the non-default base branch e2e-base-branch-20261108-062810 (already set via
> the base:<branch> label applied at filing — do not add or remove labels).

## Background

This issue is a live end-to-end regression probe for the `base:<branch>` contract. It exercises the full pipeline for an item whose work targets a non-default base branch, in order to confirm that the fixes for handarbeit/fabrik#1047 (issue-to-PR linkage on a non-default base) and #1050 (the review-gate data feed) keep working (regression coverage for #1046). The file change itself is deliberately trivial so that any failure points at the pipeline, not at the content of the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

The pipeline takes this issue through to a pull request that targets `e2e-base-branch-20261108-062810` and appends one marker line to `e2e/markers/base-branch-pipeline.md`.

**Why this priority**: This is the whole change, and it is what proves the pipeline handles a non-default base end to end.

**Independent Test**: Inspect the pull request's diff against `e2e-base-branch-20261108-062810`. It adds exactly one line to `e2e/markers/base-branch-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is applied, **Then** `<!-- base-branch-pipeline-20261108-062810 -->` is the last line of the file and no other existing line is altered.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created containing the marker line.
3. **Given** the item carries the `base:e2e-base-branch-20261108-062810` label, **When** the pipeline opens a pull request, **Then** the pull request targets that branch and not the repository default.

---

### Edge Cases

- The marker file exists but its last line has no trailing newline: the marker must still go on its own line.
- The marker line is already present (for example on a re-run): it must not be duplicated.
- The labels (`fabrik:cruise`, `base:e2e-base-branch-20261108-062810`) must not be added, removed or changed by the work.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the line `<!-- base-branch-pipeline-20261108-062810 -->` as the final line of `e2e/markers/base-branch-pipeline.md`, creating the file first if it does not exist.
- **FR-002**: The change MUST touch exactly one file and add exactly one line (plus the file creation when needed); no other files or lines are modified.
- **FR-003**: The resulting pull request MUST target the base branch `e2e-base-branch-20261108-062810`.
- **FR-004**: The pull request MUST be linked to this issue so the pipeline's issue-to-PR linkage works for a non-default base (the #1047 contract).
- **FR-005**: The pipeline MUST NOT decompose this work into child issues.
- **FR-006**: The work MUST NOT add or remove labels on the issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The pull request diff contains exactly one added line and no removed lines in exactly one file.
- **SC-002**: The pull request's base branch equals `e2e-base-branch-20261108-062810`.
- **SC-003**: The issue progresses through the pipeline stages with the PR discovered and linked to it, and the review gate receives its data, without manual intervention.
- **SC-004**: No child issues are spawned for this issue.

## Assumptions

- The branch `e2e-base-branch-20261108-062810` already exists on the remote and is the effective base for this work.
- The marker line is exact and case-sensitive, with no leading or trailing whitespace.
- Appending the line is idempotent: if the exact line is already the last line of the file, nothing further is added.

## Out of Scope

- Any change to engine code, tests or documentation.
- Work in any repository other than this single one.
- Adding or removing issue labels.
- Decomposing the work into multiple issues.
