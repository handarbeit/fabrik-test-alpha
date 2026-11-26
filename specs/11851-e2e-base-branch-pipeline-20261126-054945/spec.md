# Feature Specification: E2E base-branch pipeline marker line

**Feature Branch**: `fabrik/issue-11851`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261126-054945 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: Single repo only, targeting the non-default base branch e2e-base-branch-20261126-054945 (already set via the base:<branch> label applied at filing — do not add or remove labels). Labels: fabrik:cruise, base:e2e-base-branch-20261126-054945"

## Background

This issue is an end-to-end regression check for the `base:<branch>` pipeline contract (handarbeit/fabrik#1046). Work targeting a non-default base branch must flow through every pipeline stage correctly. That means the issue and its PR stay linked (the #1047 fix) and the review gate receives its data (the #1050 fix). The code change is deliberately trivial so the pipeline behavior is the only thing being exercised. The `base:e2e-base-branch-20261126-054945` label is already set and must not be altered.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

The pipeline appends one marker line to the marker file, opens a PR against the non-default base branch, and carries the item through to completion.

**Why this priority**: This is the only deliverable. It exercises the whole `base:<branch>` contract end to end.

**Independent Test**: Inspect the PR diff. It touches exactly one file, `e2e/markers/base-branch-pipeline.md`, and adds exactly one line, `<!-- base-branch-pipeline-20261126-054945 -->`, at the end of the file. The PR's base is `e2e-base-branch-20261126-054945`.

**Acceptance Scenarios**:

1. **Given** the marker file exists, **When** Implement completes, **Then** the file's last line is `<!-- base-branch-pipeline-20261126-054945 -->` and no other content has changed.
2. **Given** the marker file does not exist, **When** Implement completes, **Then** the file is created and contains the marker line.
3. **Given** the item carries `base:e2e-base-branch-20261126-054945`, **When** the PR is created, **Then** the PR targets that branch and is linked to the issue.

---

### Edge Cases

- The marker file is absent: create it with just the marker line.
- The marker file lacks a trailing newline: the marker still goes on its own final line.
- The marker line is already present (a re-run): do not add a duplicate.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- base-branch-pipeline-20261126-054945 -->` as the last line of `e2e/markers/base-branch-pipeline.md`, creating the file if it does not exist.
- **FR-002**: The change MUST modify no other file and MUST NOT add any other content.
- **FR-003**: The PR MUST target the base branch `e2e-base-branch-20261126-054945` and be linked to this issue.
- **FR-004**: The pipeline MUST NOT add or remove labels on the issue as part of this work.
- **FR-005**: The work MUST be delivered as a single unit and MUST NOT be decomposed into child issues.

### Key Entities

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a plain text file holding HTML comment marker lines.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one added line in exactly one file.
- **SC-002**: The PR's base branch equals `e2e-base-branch-20261126-054945`, and the issue and PR are linked.
- **SC-003**: The item progresses through the pipeline stages without the review gate or the issue-to-PR linkage failing for lack of data.

## Assumptions

- The marker file's existing content, if any, is left untouched.
- Because `fabrik:cruise` is present, stages auto-advance without a separate human approval of this spec.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Splitting the work into child issues.
- Modifying issue labels.
- Changes to engine behavior; this issue only verifies it.
