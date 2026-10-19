# Feature Specification: e2e cruise full pipeline (20261019-200310)

**Feature Branch**: `fabrik/issue-9109`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261019-200310 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898). A cruise-labelled item should advance automatically through every pipeline stage up to Validate-complete. It should not auto-merge its pull request and should not move to Done on its own. After a human merges the PR, the issue should close correctly.

The code change is deliberately trivial: one marker line in one file. That keeps the pipeline mechanics, not the content, as the thing under test. The timestamped marker value ties this run to its own change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the cruise marker file (Priority: P1)

The marker file `e2e/markers/cruise-full-pipeline.md` ends with the line `<!-- cruise-pipeline-20261019-200310 -->`. If the file did not exist, it is created containing that line.

**Why this priority**: This is the only code change. It gives the pipeline a real, reviewable diff to carry through every stage.

**Independent Test**: Inspect the PR diff. It should touch exactly one file and add exactly one line, the marker, at the end of that file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is implemented, **Then** the file's last line is `<!-- cruise-pipeline-20261019-200310 -->` and all prior content is unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

With `fabrik:cruise` on the issue, the item moves through Specify, Research, Plan, Implement, Review and Validate without human approval at each stage. It stops at Validate-complete with the PR open and unmerged.

**Why this priority**: This is the cruise contract being verified.

**Independent Test**: Observe the board and PR after Validate completes. The PR is open and unmerged, and the issue has not moved to Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage automatically.
2. **Given** Validate has completed, **When** the engine evaluates landing, **Then** the PR is not auto-merged and the issue is not advanced to Done.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

After a human merges the PR, the issue is closed and reaches Done correctly.

**Why this priority**: It completes the contract's end state, but it depends on the earlier stories.

**Independent Test**: Merge the PR by hand and confirm the issue closes and reaches Done.

**Acceptance Scenarios**:

1. **Given** the PR is open at Validate-complete, **When** a human merges it, **Then** the linked issue is closed and the item reaches Done.

---

### Edge Cases

- The marker file already contains the marker line: do not add a duplicate. This is unlikely, since the timestamp is unique to this run.
- The marker file exists without a trailing newline: the marker still ends up on its own final line.
- Plan must not decompose the work or spawn child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/cruise-full-pipeline.md` MUST end with the line `<!-- cruise-pipeline-20261019-200310 -->`, creating the file if it does not exist.
- **FR-002**: The change MUST modify only that one file, adding only that one line, with no other edits.
- **FR-003**: The work MUST NOT be decomposed into child issues.
- **FR-004**: With `fabrik:cruise` present, the item MUST auto-advance through all stages to Validate-complete.
- **FR-005**: At Validate-complete, the engine MUST NOT auto-merge the PR or move the issue to Done.
- **FR-006**: After a human merges the PR, the issue MUST close and reach Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file and exactly one added line, equal to the marker.
- **SC-002**: The item reaches Validate-complete with no human approval between stages.
- **SC-003**: At Validate-complete, the PR is open and unmerged and the issue is not in Done.
- **SC-004**: After a human merge, the issue is closed and in Done.

## Assumptions

- The `e2e/markers/` directory may need to be created if absent.
- Cruise precedence over `fabrik:yolo` is already specified elsewhere and is not re-tested here.
- The issue's own `fabrik:cruise` label is the only automation setting needed.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Testing yolo behaviour, multi-repo behaviour, or merge-train landing.
- Changes to engine or pipeline code.
