# Feature Specification: e2e cruise full pipeline (20261126-055620)

**Feature Branch**: `fabrik/issue-11859`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261126-055620 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This issue is a live end-to-end probe of the `fabrik:cruise` pipeline contract (#898). Under cruise, an issue auto-advances through every stage up to and including Validate completion, but the engine must not merge the linked PR on its own. Once a human merges the PR, the issue must close correctly. The code change is deliberately trivial, so that any failure of the run points at the pipeline's behavior and not at the work item.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

A single, uniquely identified marker line is appended to the cruise-pipeline marker file, producing a minimal one-file, one-line PR.

**Why this priority**: This is the only deliverable. It gives the pipeline a real, reviewable change to carry from Specify through Validate.

**Independent Test**: After Implement, `e2e/markers/cruise-full-pipeline.md` exists and its last line is exactly `<!-- cruise-pipeline-20261126-055620 -->`. The PR diff touches only that file and only adds that one line (plus the file's creation, if it did not exist).

**Acceptance Scenarios**:

1. **Given** the file already exists, **When** the change is implemented, **Then** the marker line is appended at the very end and existing content is unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

With `fabrik:cruise` applied, the issue moves automatically through Specify, Research, Plan, Implement, Review and Validate to Validate-complete, and the PR is left unmerged.

**Why this priority**: This is the contract under test (#898).

**Independent Test**: Observe the issue and PR after Validate completes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the engine finishes its processing, **Then** the PR remains open and unmerged, and the issue is not advanced to Done.
3. **Given** Plan runs on this issue, **When** it produces its output, **Then** it does not decompose the work into child issues.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

After a human merges the PR, the issue closes and reaches its terminal state correctly.

**Why this priority**: It completes the contract but only happens after the cruise-controlled portion.

**Independent Test**: Merge the PR by hand and observe the issue closing and moving to Done.

**Acceptance Scenarios**:

1. **Given** Validate-complete with an open PR, **When** a human merges the PR, **Then** the issue closes and moves to Done.

---

### Edge Cases

- The file exists but its last line has no trailing newline: the marker must still land on its own line at the end of the file.
- The marker file's parent directory does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- cruise-pipeline-20261126-055620 -->`, as the final line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it is created first, and then the marker line is appended.
- **FR-003**: The change touches no file other than `e2e/markers/cruise-full-pipeline.md`, and existing content in that file is not modified.
- **FR-004**: Plan does not decompose this issue into child issues.
- **FR-005**: With `fabrik:cruise` applied, the issue auto-advances through all stages to Validate-complete without human approval at each stage.
- **FR-006**: The linked PR is not merged by the engine at Validate completion, and the issue is not moved to Done by the engine at that point.
- **FR-007**: After a human merges the PR, the issue closes correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one file with exactly one added line (not counting file creation), and that line equals the specified marker.
- **SC-002**: The issue reaches Validate-complete with no manual stage approvals.
- **SC-003**: At Validate-complete, the PR is open and unmerged.
- **SC-004**: After a human merge, the issue is closed.

## Assumptions

- The marker file may or may not exist at implementation time; either case is handled per FR-002.
- The marker line is an HTML comment, so it does not alter rendered Markdown.
- The scope is a single repository.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Decomposing the work into child issues.
- Merging the PR as part of the pipeline (merging is a human action).
- Changes to the cruise implementation itself.
