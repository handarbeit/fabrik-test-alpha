# Feature Specification: e2e cruise full pipeline (20261126-160816)

**Feature Branch**: `fabrik/issue-11936`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Trivial change: append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261126-160816 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898). Cruise should carry an item through every pipeline stage automatically, stop at Validate-complete without merging the linked PR, and leave the issue to close only once a human merges that PR. The code change is deliberately trivial, a single marker line in a single file, so the pipeline behaviour is what is under test and not the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

An operator running the e2e suite needs a unique, harmless marker line added to the cruise-pipeline marker file, so the work produces a real PR for the pipeline to carry.

**Why this priority**: Without the change there is no PR, and the pipeline contract cannot be exercised.

**Independent Test**: Check the PR diff. It touches only `e2e/markers/cruise-full-pipeline.md` and adds exactly one line, which is the last line of the file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is applied, **Then** the file's final line is `<!-- cruise-pipeline-20261126-160816 -->` and all prior content is unchanged.
2. **Given** the file does not exist, **When** the change is applied, **Then** the file is created and contains the marker line as its last line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

An operator wants to confirm that an item labelled `fabrik:cruise` moves through every stage on its own and stops at Validate-complete with the PR still unmerged.

**Why this priority**: This is the contract under verification.

**Independent Test**: Watch the issue from Specify to Validate. The stage-complete labels appear without human approval, and the linked PR is still open and unmerged when Validate completes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage with no human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the linked PR remains open, unmerged, and the issue is not moved to Done by the engine.

---

### User Story 3 - Issue closes after a human merge (Priority: P2)

After a human merges the PR, the issue should close correctly.

**Why this priority**: It completes the end-to-end contract but depends on a manual step.

**Independent Test**: Merge the PR by hand and confirm the issue closes and reaches Done.

**Acceptance Scenarios**:

1. **Given** Validate is complete and the PR is unmerged, **When** a human merges the PR, **Then** the issue is closed and moves to Done.

---

### Edge Cases

- The marker file already exists with or without a trailing newline: the new line must still land on its own line at the very end.
- Plan must not decompose or spawn child issues for this one-line change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- cruise-pipeline-20261126-160816 -->`, at the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it is created first and then contains the marker line.
- **FR-003**: The change modifies no file other than `e2e/markers/cruise-full-pipeline.md`.
- **FR-004**: The work is a single unit; Plan does not decompose it into sub-issues.
- **FR-005**: With `fabrik:cruise` applied, the item auto-advances through all stages to Validate-complete.
- **FR-006**: The linked PR is not merged by the engine at Validate-complete.
- **FR-007**: After a human merges the PR, the issue closes correctly.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, which holds one marker line per e2e run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff is exactly one added line in one file.
- **SC-002**: The item reaches Validate-complete with zero human interventions before that point.
- **SC-003**: At Validate-complete the linked PR is open and unmerged.
- **SC-004**: After a human merge, the issue is closed with no further manual steps.

## Assumptions

- The marker line is added verbatim, including the timestamp `20261126-160816`.
- Single repository only; no cross-repo work is involved.

## Out of Scope *(optional)*

- Any change beyond the one marker line.
- Child-issue spawning or decomposition.
- Merging the PR as part of the pipeline; the human merge is the verification step.
