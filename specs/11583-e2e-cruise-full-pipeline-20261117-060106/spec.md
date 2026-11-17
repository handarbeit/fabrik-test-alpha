# Feature Specification: E2E cruise full-pipeline marker (20261117-060106)

**Feature Branch**: `fabrik/issue-11583`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line `<!-- cruise-pipeline-20261117-060106 -->` to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist). That is the entire change. One file, one line. Plan should NOT decompose. Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label's pipeline contract (#898). Under cruise, an item auto-advances through every stage up to Validate-complete but the linked PR is not auto-merged; the issue is expected to close only after a human merges the PR. The code change is deliberately trivial (one marker line) so that any failure observed is attributable to the pipeline, not to the change itself. The marker file already exists and holds earlier cruise-run markers, one per line.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the run marker (Priority: P1)

A pipeline operator needs a unique, traceable marker for this cruise run committed to the repository, so the run's PR has a real, minimal diff to carry through the pipeline.

**Why this priority**: It is the only deliverable; the pipeline run has nothing to move through without it.

**Independent Test**: Inspect the last line of `e2e/markers/cruise-full-pipeline.md` on the PR branch and the diff against the base branch.

**Acceptance Scenarios**:

1. **Given** the marker file exists with prior markers, **When** the change is applied, **Then** its final line is exactly `<!-- cruise-pipeline-20261117-060106 -->` and all earlier lines are unchanged.
2. **Given** the change is applied, **When** the PR diff is reviewed, **Then** exactly one file is changed and exactly one line is added.

---

### User Story 2 - Cruise reaches Validate-complete without merging (Priority: P1)

An operator wants to confirm that cruise carries the issue through every stage to Validate-complete and stops there, leaving the PR open for a human.

**Why this priority**: This is the contract under test (#898).

**Independent Test**: Observe the issue's stage labels and the PR state after Validate completes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** Validate completes, **When** the engine evaluates landing, **Then** the PR remains open and unmerged and the issue is not advanced to Done.

---

### User Story 3 - Issue closes after a human merge (Priority: P2)

After a human merges the PR, the issue should close correctly.

**Why this priority**: Verifies the tail of the contract, but depends on the human step outside the pipeline.

**Independent Test**: Merge the PR by hand and observe the issue state.

**Acceptance Scenarios**:

1. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue ends up closed and at Done.

---

### Edge Cases

- If the marker file's last line has no trailing newline, the new marker must still be on its own line.
- Plan must not decompose or spawn child issues; the work is a single one-line change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- cruise-pipeline-20261117-060106 -->` MUST be appended as the final line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: The change MUST NOT modify, remove or reorder any existing line in that file.
- **FR-003**: The change MUST touch only that one file and add only that one line.
- **FR-004**: Plan MUST NOT decompose the work into child issues.
- **FR-005**: The linked PR MUST NOT be auto-merged by the engine; it stays open at Validate-complete.

### Key Entities

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, an append-only list of HTML-comment run markers.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against the base branch is exactly one added line in exactly one file, and that line matches the specified marker.
- **SC-002**: The issue reaches Validate-complete without any human stage approval.
- **SC-003**: At Validate-complete the PR is open and unmerged.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker file already exists on the base branch, so no creation is needed.
- The marker value is used verbatim, including its timestamp.

## Out of Scope *(optional)*

- Any change to engine code, stage configuration or documentation.
- Multi-repo behaviour.
- Cleaning up or deduplicating earlier markers in the file.
