# Feature Specification: e2e cruise full pipeline marker (20261107-214246)

**Feature Branch**: `fabrik/issue-10913`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261107-214246 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an end-to-end smoke test of the `fabrik:cruise` label contract (#898). Cruise should auto-advance an item through every pipeline stage up to Validate-complete, without auto-merging the linked PR or moving the issue to Done. The issue should then close properly once a human merges the PR. The code change is deliberately trivial, so that any failure observed comes from the pipeline and not from the change. The marker file already exists and holds earlier run markers, each one an HTML comment line.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

A single run-identifying HTML comment line is appended to the end of the cruise full-pipeline marker file, giving the pipeline a real but minimal diff to carry through every stage.

**Why this priority**: The marker line is the only deliverable; it gives the pipeline a PR to carry.

**Independent Test**: Inspect the marker file on the PR branch. Its last line is the new marker, and the rest of the file is unchanged.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier markers, **When** the change is applied, **Then** the file ends with exactly the line `<!-- cruise-pipeline-20261107-214246 -->` and all earlier lines are unchanged.
2. **Given** the change is complete, **When** the PR diff is reviewed, **Then** it touches only the marker file and adds exactly one line (the spec file the engine commits for this issue aside).

---

### User Story 2 - Cruise advances without merging (Priority: P1)

With `fabrik:cruise` set, the item moves through all stages to Validate-complete automatically, and the PR is neither auto-merged nor is the issue moved to Done.

**Why this priority**: This is the contract being verified.

**Independent Test**: Observe the item after Validate completes. The PR is still open and the issue is not in Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage with no human action.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** the PR remains open and unmerged, and the issue stays out of Done.
3. **Given** a human then merges the PR, **When** the engine next processes the item, **Then** the issue closes and the item moves to Done.

---

### Edge Cases

- The marker file ends without a trailing newline: the new line must still sit on its own line and must not be joined to the previous marker.
- The marker file is missing: it is created and holds only the new line.
- The marker line is already present from an earlier run: it is not duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The marker file `e2e/markers/cruise-full-pipeline.md` MUST end with the line `<!-- cruise-pipeline-20261107-214246 -->`, written on its own line at the very end of the file.
- **FR-002**: Existing content of the marker file MUST be left unchanged.
- **FR-003**: The change MUST consist of exactly one added line in exactly one file, apart from the spec file the engine commits.
- **FR-004**: Planning MUST NOT decompose the work into child issues.
- **FR-005**: With `fabrik:cruise` set, the item MUST auto-advance through all stages to Validate-complete, and the PR MUST NOT be auto-merged.
- **FR-006**: After a human merges the PR, the issue MUST close and move to Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The marker file's last line equals `<!-- cruise-pipeline-20261107-214246 -->` and the diff against the base branch is one added line in that file.
- **SC-002**: The item reaches Validate-complete with no manual stage advancement, and the PR is still open at that point.
- **SC-003**: The issue is closed and in Done after the human merge, with no further manual steps.

## Assumptions

- The marker file already exists, so only an append is needed. If it does not, it is created.
- The marker line is not already present in the file; the timestamp is unique to this run.
- The spec file the engine commits is an expected addition and does not count against the one-file rule.

## Out of Scope *(optional)*

- Any code, configuration or documentation change beyond the single marker line.
- Multi-repo behaviour.
- Auto-merge behaviour (covered by the yolo tests).
- Decomposition into child issues.
