# Feature Specification: e2e cruise full pipeline marker (20261127-014741)

**Feature Branch**: `fabrik/issue-12130`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261127-014741 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` pipeline contract (#898). With `fabrik:cruise` set, an item should advance automatically through every stage up to Validate-complete. It should not merge its PR or move to Done on its own. After a human merges the PR, the issue should close correctly.

The code change is deliberately tiny, so that any failure comes from the pipeline and not from the work. It is a single marker line in a single file. The run-specific token in the line (`20261127-014741`) makes this run's change distinguishable from earlier runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

An operator running the cruise e2e check needs the repository to contain this run's marker line at the end of the cruise-pipeline marker file. The line is the observable proof that the pipeline produced and delivered a change.

**Why this priority**: The line is the only deliverable. The rest of the pipeline check depends on a real PR carrying it.

**Independent Test**: After the change, the last line of `e2e/markers/cruise-full-pipeline.md` is exactly `<!-- cruise-pipeline-20261127-014741 -->`, and no other file is modified.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is made, **Then** the file's original content is unchanged and the marker line is appended as the final line.
2. **Given** `e2e/markers/cruise-full-pipeline.md` does not exist, **When** the change is made, **Then** the file is created and contains the marker line.
3. **Given** the change is complete, **When** the diff is inspected, **Then** it touches exactly one file and adds exactly one line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

An operator wants to confirm that the issue, labelled `fabrik:cruise`, advances through every stage to Validate-complete without the PR being merged automatically. The issue must also not move to Done at that point.

**Why this priority**: This is the contract under test (#898). The marker line only serves to give the pipeline something to carry.

**Independent Test**: Observe the issue after Validate completes. The linked PR is open and unmerged, the issue has not moved to Done, and the Validate stage is complete.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the PR remains open and unmerged and the issue stays out of Done.
3. **Given** a human merges the PR after Validate-complete, **When** the merge lands, **Then** the issue closes correctly.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker line must still start on its own line, so a newline is added first.
- The marker file already contains marker lines from earlier runs: they are left untouched and this run's line is added after them.
- The Plan stage must not split this work into child issues, since it is a one-line change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/cruise-full-pipeline.md` ends with the line `<!-- cruise-pipeline-20261127-014741 -->` after the change.
- **FR-002**: If the file does not exist, it is created, with the marker line as its content.
- **FR-003**: Existing content of the file is preserved; the change only appends.
- **FR-004**: The change touches only this one file and adds only this one line.
- **FR-005**: The Plan stage does not decompose the work into child issues.
- **FR-006**: With `fabrik:cruise` set, the issue auto-advances through all stages to Validate-complete.
- **FR-007**: The pipeline does not auto-merge the PR and does not advance the issue to Done at Validate-complete.
- **FR-008**: After a human merges the PR, the issue closes correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one added line in exactly one file, and the line matches the marker text above.
- **SC-002**: The issue reaches Validate-complete with no human advancement between stages.
- **SC-003**: At Validate-complete the PR is unmerged and the issue is not in Done.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker directory `e2e/markers/` may or may not exist; creating it along with the file is acceptable.
- The marker text is used exactly as given in the issue, including the timestamp token `20261127-014741`.
- The marker file is plain text or markdown, so an HTML comment line is harmless to any consumer.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Changes to the cruise logic itself or to engine behaviour.
- Cross-repo scenarios.
- Cleanup of marker lines from earlier runs.
