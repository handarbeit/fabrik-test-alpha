# Feature Specification: e2e cruise full pipeline (20261030-082309)

**Feature Branch**: `fabrik/issue-10227`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261030-082309 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898). Under cruise, an item auto-advances through every pipeline stage up to Validate-complete, but the PR is not auto-merged and the issue does not advance to Done on its own. A human merges the PR, and the issue should then close correctly.

The code change is deliberately trivial: one marker line appended to a marker file. The change is only there so the pipeline has a real, reviewable diff to carry through each stage. The timestamped marker line makes this run's change unique and easy to identify.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended to the cruise marker file (Priority: P1)

The pipeline produces a PR whose only change is one HTML comment line added at the end of `e2e/markers/cruise-full-pipeline.md`.

**Why this priority**: This is the entire deliverable. Without a diff, nothing else can be verified.

**Independent Test**: Inspect the PR diff and confirm that the only change is one added line, `<!-- cruise-pipeline-20261030-082309 -->`, as the last line of `e2e/markers/cruise-full-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is implemented, **Then** the file's existing content is unchanged and the marker line is its final line.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line as its last line.
3. **Given** the change is implemented, **When** the PR diff is reviewed, **Then** exactly one file is modified, with exactly one line added.

---

### User Story 2 - Cruise carries the item through all stages without merging (Priority: P1)

With `fabrik:cruise` applied, the item advances automatically from Specify through Validate-complete. The linked PR is left open and unmerged.

**Why this priority**: This is the cruise contract the issue exists to verify.

**Independent Test**: Observe the item's stage-complete labels and board column after the pipeline runs. Confirm the PR is still open.

**Acceptance Scenarios**:

1. **Given** the item carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** Validate completes, **When** the engine reaches its landing decision, **Then** the PR is not auto-merged and the issue is not moved to Done.

---

### User Story 3 - Issue closes correctly after a human merges the PR (Priority: P2)

After a human merges the PR, the issue ends in the correct closed state.

**Why this priority**: It completes the cruise contract, but it depends on a human action after the pipeline finishes.

**Independent Test**: Merge the PR by hand and confirm the issue closes and reaches Done.

**Acceptance Scenarios**:

1. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue is closed and moved to Done.

---

### Edge Cases

- The marker file already exists with or without a trailing newline: the marker must still end up on its own final line, and the existing content must not be altered.
- The marker line is already present from an earlier run: it is not duplicated. The timestamp makes this unlikely.
- Plan must not decompose the work or spawn child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly the line `<!-- cruise-pipeline-20261030-082309 -->` at the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created, containing the marker line.
- **FR-003**: If the file exists, its prior content MUST be left unchanged.
- **FR-004**: The PR MUST modify only that one file, adding only that one line.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: The PR MUST NOT be auto-merged by the pipeline, and the issue MUST NOT be advanced to Done automatically at Validate completion, because `fabrik:cruise` is present.
- **FR-007**: After a human merges the PR, the issue MUST close correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file and exactly one added line, equal to the marker line.
- **SC-002**: The item reaches Validate-complete through automatic advancement, with no manual stage approvals.
- **SC-003**: At Validate-complete, the PR is still open and unmerged.
- **SC-004**: Within one poll cycle of a human merging the PR, the issue is closed.

## Assumptions

- The `e2e/markers/` directory may or may not already exist. If not, it is created together with the file.
- The marker line ends with a newline like any other line in the file.
- The change is made in this repository only.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Child-issue spawning or multi-repo behavior.
- Changes to the cruise logic or engine behavior itself. This issue only verifies the existing contract.
