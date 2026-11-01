# Feature Specification: e2e base-branch pipeline (20261101-065402)

**Feature Branch**: `fabrik/issue-10713`
**Created**: 2026-11-01
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261101-065402 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: Single repo only, targeting the non-default base branch e2e-base-branch-20261101-065402 (already set via the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end probe of the `base:<branch>` pipeline. The change itself is deliberately trivial. The point is to exercise the full pipeline against a non-default base branch and confirm two earlier fixes still hold:

- **#1047**: the issue and PR are linked correctly even though GitHub creates no Development-panel link for a PR that targets a non-default base.
- **#1050**: the review gate receives its data correctly for such a PR.

This is regression coverage for #1046. The content of the change does not matter. What matters is that the item moves through the pipeline on the `e2e-base-branch-20261101-065402` base, using the `base:` label already on the issue.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

A single marker line is appended to a tracked marker file. The work is branched from, and the PR targets, the non-default base branch named by the issue's `base:` label.

**Why this priority**: It is the only deliverable, and it is the vehicle that exercises the pipeline.

**Independent Test**: After the change merges, `e2e/markers/base-branch-pipeline.md` on `e2e-base-branch-20261101-065402` ends with the line `<!-- base-branch-pipeline-20261101-065402 -->`.

**Acceptance Scenarios**:

1. **Given** the marker file exists, **When** the change is made, **Then** the file's final line is exactly `<!-- base-branch-pipeline-20261101-065402 -->` and nothing else in the file changed.
2. **Given** the marker file does not exist, **When** the change is made, **Then** the file is created and contains the marker line as its last line.
3. **Given** the issue carries `base:e2e-base-branch-20261101-065402`, **When** the PR is opened, **Then** the PR targets that branch rather than the repository's other branches.

---

### User Story 2 - Pipeline contract holds for a non-default base (Priority: P1)

The item moves through the stages with the issue and PR correctly linked and the review gate seeing the PR's review data.

**Why this priority**: This is the regression the probe exists to detect (#1047, #1050).

**Independent Test**: Observe that the issue is discovered as linked to its PR, the review gate evaluates that PR's data, and the item completes the pipeline.

**Acceptance Scenarios**:

1. **Given** a PR targeting the non-default base, **When** the pipeline runs, **Then** the engine associates the PR with issue #10713 and uses it for the stages that need a linked PR.
2. **Given** that PR, **When** the review gate is evaluated, **Then** it uses the PR's review data and does not treat the PR as missing.

---

### Edge Cases

- The marker file already exists with or without a trailing newline: the new line must be appended at the end, as its own line, without altering existing content.
- The issue's labels must not be added or removed as part of this work.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/base-branch-pipeline.md` ends with the single line `<!-- base-branch-pipeline-20261101-065402 -->`.
- **FR-002**: If the file does not exist, it is created, and the marker line is its last line.
- **FR-003**: The change touches only that one file and adds only that one line (plus the file's creation if it was absent).
- **FR-004**: The work is based on, and its PR targets, `e2e-base-branch-20261101-065402`, as set by the existing `base:` label.
- **FR-005**: No labels are added to or removed from the issue as part of the change.
- **FR-006**: The work is not decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The diff of the resulting PR is exactly one file with one added line (plus the file's creation if it was absent).
- **SC-002**: The PR's base branch is `e2e-base-branch-20261101-065402`.
- **SC-003**: The item completes the pipeline with the issue and PR linked and the review gate satisfied, without manual intervention to repair linkage or review data.

## Assumptions

- The repository's default base branch for this run is `e2e-base-branch-20261101-065402`, as stated in the stage context. The `base:` label already names it.
- The marker line's timestamp (`20261101-065402`) is fixed and is not regenerated.
- If the file must be created, no header or other content is required beyond the marker line.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Child issues or decomposition.
- Changes to engine behaviour, labels, or other marker files.
- Re-fixing #1047 or #1050. This run only verifies them.
