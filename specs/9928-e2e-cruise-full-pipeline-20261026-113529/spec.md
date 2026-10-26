# Feature Specification: e2e cruise full pipeline marker (20261026-113529)

**Feature Branch**: `fabrik/issue-9928`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261026-113529 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` contract (#898). Cruise should carry an item through every pipeline stage automatically up to Validate-complete, stop without merging the linked PR, and let the issue close once a human merges that PR. The code change is deliberately trivial so the pipeline's behavior, not the change itself, is what gets verified. The marker file `e2e/markers/cruise-full-pipeline.md` already exists in the repository and holds one HTML comment line per earlier run of this test.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

An operator running the cruise e2e check needs a unique, run-specific marker line added to the shared marker file. That gives the run a real, reviewable one-line diff to carry through the pipeline.

**Why this priority**: Without the change there is no PR, so nothing else in the pipeline contract can be checked.

**Independent Test**: Inspect the PR diff. It should add exactly one line to `e2e/markers/cruise-full-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier run lines, **When** the change is implemented, **Then** the file ends with `<!-- cruise-pipeline-20261026-113529 -->` on its own final line and all earlier content is unchanged.
2. **Given** the PR is opened, **When** its diff is reviewed, **Then** it touches only `e2e/markers/cruise-full-pipeline.md` and adds only that one line.

---

### User Story 2 - Cruise carries the item to Validate-complete without merging (Priority: P1)

An operator who labels an issue `fabrik:cruise` expects every stage to advance automatically through Validate. The PR must be left open for a human to merge.

**Why this priority**: This is the contract under test (#898).

**Independent Test**: Watch the issue run from Specify to Validate-complete, then confirm the linked PR is still open and unmerged.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage with no human action.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the linked PR is not merged by the engine and the issue is not yet closed.
3. **Given** Validate is complete, **When** a human merges the PR, **Then** the issue closes and moves to Done.

---

### Edge Cases

- The marker file is missing at implementation time: it is created containing just the marker line.
- The file's last line has no trailing newline: the marker still goes on its own line rather than being joined to the previous one.
- The same marker line already exists: this run's timestamp is unique, so a duplicate is not expected. If it does appear, the file must not end up with two copies.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- cruise-pipeline-20261026-113529 -->` is appended as the last line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: The file is created, containing only that line, if it does not exist.
- **FR-003**: No other lines in the file are modified, removed or reordered.
- **FR-004**: No file other than `e2e/markers/cruise-full-pipeline.md` is changed.
- **FR-005**: The work is done as a single unit and is not decomposed into child issues.
- **FR-006**: Under `fabrik:cruise`, the engine does not merge the linked PR when Validate completes.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff is exactly one added line in exactly one file.
- **SC-002**: The item reaches Validate-complete with no human action between stages.
- **SC-003**: At Validate-complete, the linked PR is open and unmerged.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker file already exists, as the repository shows. The create-if-missing instruction is a safeguard only.
- The date-time token `20261026-113529` is copied exactly as given and is not regenerated.
- Specify's output is only this spec; the marker line is written during Implement.

## Out of Scope

- Any change other than the single marker line.
- Changes to the cruise implementation, engine or stage configuration.
- Cleanup or deduplication of earlier marker lines.
- Spawning child issues or decomposing the work.
