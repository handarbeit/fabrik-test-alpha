# Feature Specification: e2e base-branch pipeline (20261117-072418)

**Feature Branch**: `fabrik/issue-11621`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261117-072418 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: Single repo only, targeting the non-default base branch e2e-base-branch-20261117-072418 (already set via the base:<branch> label applied at filing — do not add or remove labels). Labels: fabrik:cruise, base:e2e-base-branch-20261117-072418"

## Background

This issue is a live end-to-end test fixture. It checks that an issue labelled `base:<branch>` moves through the whole Fabrik pipeline correctly when its target is a non-default base branch. It is regression coverage for handarbeit/fabrik#1046 and validates the fixes in #1047 (issue↔PR linkage) and #1050 (review-gate data feed). The code change is deliberately trivial, so that any failure points at the pipeline's base-branch handling and not at the work itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

A single marker line is added to a marker file. The work flows through the pipeline against the issue's declared base branch and not the repository's default branch.

**Why this priority**: This is the whole change and the whole purpose of the issue.

**Independent Test**: After the pipeline completes, read `e2e/markers/base-branch-pipeline.md` on branch `e2e-base-branch-20261117-072418` and confirm its last line is the marker comment.

**Acceptance Scenarios**:

1. **Given** the issue carries the `base:e2e-base-branch-20261117-072418` label, **When** the pipeline runs through Implement, **Then** the issue's PR targets `e2e-base-branch-20261117-072418` as its base.
2. **Given** the marker file does not exist on the base branch, **When** the change is made, **Then** the file is created and contains the marker line as its final line.
3. **Given** the marker file already exists, **When** the change is made, **Then** the marker line is appended at the very end and existing content is unchanged.

---

### Edge Cases

- The marker file may already exist or may not. Both cases must end with the marker line as the final line.
- The marker line must not be duplicated if the stage is re-run.
- The file should end cleanly, with the marker on its own line and no blank trailing garbage added.
- The labels `fabrik:cruise` and `base:e2e-base-branch-20261117-072418` must be left unchanged.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- base-branch-pipeline-20261117-072418 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created, and it contains the marker line.
- **FR-003**: The change touches only that one file, and the diff is one added line (plus the file creation if it was missing).
- **FR-004**: The work is done on, and its PR targets, the base branch `e2e-base-branch-20261117-072418`.
- **FR-005**: The issue's labels are not added to or removed from by the work.
- **FR-006**: The work is not decomposed into child issues.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a text file whose last line is the run-specific marker comment.
- **Base branch**: `e2e-base-branch-20261117-072418`, the non-default branch the issue and its PR target.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The final line of the marker file on the base branch is exactly `<!-- base-branch-pipeline-20261117-072418 -->`.
- **SC-002**: The change set consists of exactly one file and one added line (plus file creation if it was missing).
- **SC-003**: The issue's PR is linked to the issue and targets the non-default base branch, so the pipeline can discover the PR and its review state.
- **SC-004**: The item proceeds through every stage without a pause, a base-branch fallback comment or a linkage failure.

## Assumptions

- The base branch `e2e-base-branch-20261117-072418` already exists on the remote.
- The `base:<branch>` and `fabrik:cruise` labels were applied at filing and are correct as they stand.
- A missing marker file means it is created with the marker line as its only content.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Splitting the work into child issues.
- Adding, removing or changing labels.
- Changes to Fabrik engine code or the base-branch handling itself (this issue only verifies it).

## Source References *(optional)*

- handarbeit/fabrik#1046 (base-branch pipeline regression), #1047 (issue↔PR linkage fix), #1050 (review-gate data feed fix)
