# Feature Specification: e2e cruise full pipeline (20261014-193634)

**Feature Branch**: `fabrik/issue-8194`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261014-193634 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` contract (#898). Under cruise, an item auto-advances through every pipeline stage up to Validate-complete, but the linked PR is not auto-merged and the issue does not advance to Done until a human merges the PR. The change itself is deliberately trivial, a one-line marker append to a tracking file, so that any failure observed comes from pipeline behavior and not from the content of the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

An operator running the cruise e2e check needs the pipeline to produce a PR whose only change is one new marker line at the end of the marker file.

**Why this priority**: The marker line is the only artifact that proves this particular run produced the change.

**Independent Test**: Inspect the PR diff: exactly one file changed, exactly one line added, at the end of the file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the pipeline implements this issue, **Then** the file ends with the line `<!-- cruise-pipeline-20261014-193634 -->` and all earlier lines are unchanged.
2. **Given** the file did not exist, **When** the pipeline implements this issue, **Then** the file is created containing that line.

---

### User Story 2 - Cruise advances without merging (Priority: P1)

An operator needs confirmation that, with `fabrik:cruise` set, the item moves through all stages to Validate-complete and then stops, leaving the PR unmerged.

**Why this priority**: This is the cruise contract being verified.

**Independent Test**: Observe the item reach Validate-complete with the PR still open and the issue still open.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** the PR is not merged and the issue is not moved to Done.

---

### User Story 3 - Issue closes after human merge (Priority: P2)

An operator needs the issue to close correctly once a human merges the PR.

**Why this priority**: It completes the cruise lifecycle, but depends on the earlier stories.

**Independent Test**: Merge the PR by hand and confirm the issue moves to Done and closes.

**Acceptance Scenarios**:

1. **Given** the PR is open at Validate-complete, **When** a human merges it, **Then** the issue is closed and moved to Done.

---

### Edge Cases

- The marker file lacks a trailing newline: the new line must still sit on its own line.
- The file already contains the same marker line (for example on a retry): it is not duplicated.
- Plan must not decompose the work into child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- cruise-pipeline-20261014-193634 -->`, at the very end of `e2e/markers/cruise-full-pipeline.md`, creating the file first if it does not exist.
- **FR-002**: The change MUST touch only that one file and MUST NOT alter any existing line in it.
- **FR-003**: The work MUST NOT be decomposed into child issues.
- **FR-004**: Under `fabrik:cruise`, the item MUST auto-advance through all stages to Validate-complete.
- **FR-005**: At Validate-complete the PR MUST NOT be auto-merged and the issue MUST NOT be moved to Done by the engine.
- **FR-006**: After a human merges the PR, the issue MUST be closed and reach Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly 1 changed file and exactly 1 added line, with no deletions.
- **SC-002**: The item reaches Validate-complete with no manual stage approval.
- **SC-003**: The PR is unmerged at Validate-complete, and the issue is closed once it is merged by a human.

## Assumptions

- The marker file already exists with earlier run markers, and the new line goes after them.
- The repository is a single repo and no cross-repo work is involved.

## Out of Scope

- Any change other than the single marker line.
- Splitting the work into child issues.
- Auto-merging the PR.
- Changes to the cruise implementation itself.
