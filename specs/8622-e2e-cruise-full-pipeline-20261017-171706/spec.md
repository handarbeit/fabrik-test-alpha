# Feature Specification: e2e cruise full pipeline marker (20261017-171706)

**Feature Branch**: `fabrik/issue-8622`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261017-171706 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end test fixture for the `fabrik:cruise` label contract (#898). Under cruise, an item auto-advances through every pipeline stage up to Validate-complete, but the engine does not merge the PR or move the issue to Done. A human merges the PR, and the issue is then expected to close correctly.

The code change is deliberately trivial: one marker line appended to one file. The point is to exercise the pipeline, not the change. The change must stay a single-file, single-line edit so that Plan does not decompose it into child issues.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is recorded in the marker file (Priority: P1)

A maintainer running the cruise e2e check needs a unique, verifiable change produced by the pipeline. The pipeline appends a run-specific HTML comment to the cruise marker file, giving the PR a real diff to merge.

**Why this priority**: Without a diff there is no PR to carry through the stages, so no other part of the cruise contract can be verified.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the issue branch. Its last line is exactly `<!-- cruise-pipeline-20261017-171706 -->`, and the diff against the base branch adds only that line (plus the file itself if it did not exist).

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is implemented, **Then** the file's final line is `<!-- cruise-pipeline-20261017-171706 -->` and all prior content is unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line as its final line.

---

### User Story 2 - Cruise advances to Validate-complete without merging (Priority: P1)

The item carries `fabrik:cruise`. It moves through Specify, Research, Plan, Implement, Review and Validate without human approval at each stage. When Validate completes, the engine leaves the PR unmerged and the issue not Done.

**Why this priority**: This is the cruise contract the test exists to verify.

**Independent Test**: After Validate completes, the linked PR is open and unmerged, the issue is open, and the issue has not moved to Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the engine advances to the next stage without waiting for human approval.
2. **Given** Validate has completed, **When** the engine evaluates the landing decision, **Then** it does not merge the PR and does not advance the issue to Done.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

After Validate-complete, a human merges the PR, and the issue closes as a result.

**Why this priority**: It completes the cruise contract, but it depends on the earlier stories.

**Independent Test**: Merge the PR by hand and confirm the issue ends up closed and in Done.

**Acceptance Scenarios**:

1. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue is closed and reaches Done.

---

### Edge Cases

- The marker file exists but its last line has no trailing newline: the marker must still land on its own line.
- The marker file does not exist: it is created with the marker line.
- Plan must not decompose the work or spawn child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/cruise-full-pipeline.md` contains the line `<!-- cruise-pipeline-20261017-171706 -->` as its final line.
- **FR-002**: If the file already exists, all of its prior content is preserved unchanged; if it does not exist, it is created.
- **FR-003**: The change touches exactly one file and adds exactly one line, apart from creating the file if it was absent.
- **FR-004**: The work is not decomposed into child issues.
- **FR-005**: With `fabrik:cruise` present, the item advances through every stage to Validate-complete without human approval, and the PR is not auto-merged.
- **FR-006**: After a human merges the PR, the issue closes.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists solely of the single marker line in `e2e/markers/cruise-full-pipeline.md`.
- **SC-002**: The item reaches Validate-complete with no human stage approvals, and the PR is still open and unmerged at that point.
- **SC-003**: After the human merge, the issue is closed within the normal post-merge settle window.
- **SC-004**: No child issues are created for this item.

## Assumptions

- The marker text is exactly as given, including the run-specific timestamp `20261017-171706`.
- The file lives at `e2e/markers/cruise-full-pipeline.md` in the repository the issue belongs to (single repo).
- The e2e harness, not this change, is responsible for performing the human merge.

## Out of Scope

- Any change other than the single marker line.
- Changes to the cruise label's behaviour or to engine code.
- Multi-repo scenarios.
- Decomposing the work into child issues.
