# Feature Specification: e2e cruise full pipeline marker line

**Feature Branch**: `fabrik/issue-11319`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261116-142024 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an end-to-end test fixture for the `fabrik:cruise` label contract (#898). Under cruise, an item auto-advances through every pipeline stage up to Validate-complete, but the engine does not merge the PR. A human merges it, and the issue then closes. The file change is deliberately minimal so the test exercises the pipeline mechanics rather than the content of the change. The `fabrik:cruise` label is already on the issue.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

An operator running the cruise e2e check needs a distinct, traceable one-line change on the issue branch. It gives the pipeline a real diff and PR to carry through every stage.

**Why this priority**: This is the only deliverable. Without the diff there is no PR to advance or merge.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the issue branch. Its final line should be exactly `<!-- cruise-pipeline-20261116-142024 -->`, and the PR diff should contain only that file with only that added line (plus file creation if it was absent).

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is implemented, **Then** the file's last line is `<!-- cruise-pipeline-20261116-142024 -->` and all prior content is unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created containing the marker line.

---

### User Story 2 - Cruise pipeline contract holds (Priority: P1)

An operator wants to confirm that, under `fabrik:cruise`, the item reaches Validate-complete without being auto-merged, and that the issue closes once a human merges the PR.

**Why this priority**: This is the purpose of the e2e run.

**Independent Test**: Observe the item progress through every stage with the PR left open at Validate-complete, then merge the PR by hand and confirm the issue closes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** Validate has completed, **When** no human has merged the PR, **Then** the PR remains unmerged and the issue is not advanced to Done.
3. **Given** a human merges the PR, **When** the engine processes the merge, **Then** the issue closes correctly.

---

### Edge Cases

- If the marker file already ends without a trailing newline, the marker must still land on its own final line.
- If the marker line is already present, it must not be appended a second time.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- cruise-pipeline-20261116-142024 -->`, at the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it is created first, so that the marker line is its content.
- **FR-003**: No other files are changed and no existing content in the marker file is modified.
- **FR-004**: The work is a single unit and is not decomposed into child issues.
- **FR-005**: Under `fabrik:cruise`, the item auto-advances through all stages to Validate-complete without the PR being merged by the engine.
- **FR-006**: After a human merges the PR, the issue closes correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff touches exactly one file and adds exactly one line (the marker comment).
- **SC-002**: The item reaches Validate-complete with no manual stage approvals, and the PR is still open and unmerged at that point.
- **SC-003**: After a human merge, the issue is closed.

## Assumptions

- The `e2e/markers/` directory may need to be created along with the file if it is absent.
- Single repo only; no cross-repo effects.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Child-issue spawning or plan decomposition.
- Changes to the cruise feature itself (#898).
- Automatic merging of the PR.
