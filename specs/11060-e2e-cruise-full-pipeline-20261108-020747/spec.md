# Feature Specification: e2e cruise full pipeline (20261108-020747)

**Feature Branch**: `fabrik/issue-11060`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261108-020747 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` contract (#898). Under cruise, an item auto-advances through every pipeline stage up to Validate-complete. Cruise does not auto-merge the PR and does not move the issue to Done by itself. The issue should only close after a human merges the PR.

A deliberately tiny, low-risk change exercises the whole pipeline without adding unrelated variables. The change is one appended marker line in a single file. It must stay a single unit of work, so Plan should not decompose it into child issues.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended to the cruise marker file (Priority: P1)

The repository's `e2e/markers/cruise-full-pipeline.md` file ends with the line `<!-- cruise-pipeline-20261108-020747 -->`. This is the one concrete change the pipeline delivers.

**Why this priority**: It is the only deliverable. Without it there is nothing for the pipeline to carry through its stages.

**Independent Test**: Check out the PR branch and read the last line of `e2e/markers/cruise-full-pipeline.md`. Confirm that the diff against the base branch adds only that one line (plus the file itself if it was newly created).

**Acceptance Scenarios**:

1. **Given** the file already exists on the base branch, **When** the change is applied, **Then** the file's final line is `<!-- cruise-pipeline-20261108-020747 -->` and all prior content is unchanged.
2. **Given** the file does not exist on the base branch, **When** the change is applied, **Then** the file is created containing the marker line as its final line.

---

### User Story 2 - Cruise auto-advances to Validate-complete without merging (Priority: P1)

With `fabrik:cruise` on the issue, the item advances through all stages to Validate-complete with no human approval between stages. The PR is neither merged nor enqueued, and the issue is not moved to Done.

**Why this priority**: Verifying this cruise behavior is the purpose of the issue.

**Independent Test**: Observe the issue's stage labels and board column after the pipeline settles. Confirm that Validate is complete, the PR is open and unmerged, and the issue is open.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item moves to the next stage without a human comment or approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the linked PR remains open and unmerged and the issue remains open.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

After a human merges the PR, the issue closes correctly and reaches Done.

**Why this priority**: It completes the contract. It depends on the first two stories having succeeded.

**Independent Test**: Merge the PR manually and confirm that the issue closes and the item reaches Done.

**Acceptance Scenarios**:

1. **Given** the PR is open and Validate is complete, **When** a human merges the PR, **Then** the issue is closed and the item reaches Done.

---

### Edge Cases

- The marker file does not exist yet: it is created with the marker line as its content.
- The existing file lacks a trailing newline: the marker still lands on its own line at the end of the file.
- The marker line is already present at the end of the file: it is not duplicated.
- Plan must not decompose this into child issues. It is a single-file, single-line change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly the line `<!-- cruise-pipeline-20261108-020747 -->` as the final line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created, containing the marker line.
- **FR-003**: The change MUST NOT modify any other file or any existing content in the marker file, apart from the possible file creation in FR-002.
- **FR-004**: The work MUST be delivered as a single change with no decomposition into child issues.
- **FR-005**: With `fabrik:cruise` present, the item MUST advance through all stages to Validate-complete without requiring human approval between stages.
- **FR-006**: The pipeline MUST NOT merge the PR and MUST NOT advance the issue to Done on its own at Validate-complete.
- **FR-007**: After a human merges the PR, the issue MUST close and the item MUST reach Done.

### Key Entities

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, a single-purpose file whose last line records this run's marker.
- **Marker line**: the HTML comment `<!-- cruise-pipeline-20261108-020747 -->`, unique to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against the base branch contains exactly one added content line, the marker line, in one file.
- **SC-002**: The item reaches Validate-complete with no human intervention between stages.
- **SC-003**: At Validate-complete the linked PR is open and unmerged, and the issue is open.
- **SC-004**: After a human merges the PR, the issue is closed and the item is in Done.

## Assumptions

- The marker line is taken verbatim from the issue, including the `20261108-020747` identifier, even though it differs from the spec's creation date. It is a run identifier, not a date to be reconciled.
- A missing trailing newline in an existing file is handled so the marker sits on its own line.
- The `e2e/markers/` directory may need to be created along with the file.

## Out of Scope *(optional)*

- Any change to Fabrik engine behavior or to the cruise contract itself.
- Changes to files other than the marker file.
- Multi-repo behavior.
- Automatically merging the PR. Cruise intentionally leaves the merge to a human.

## Source References *(optional)*

- #898: the `fabrik:cruise` pipeline contract this issue verifies.
