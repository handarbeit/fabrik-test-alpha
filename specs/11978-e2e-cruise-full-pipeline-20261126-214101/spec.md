# Feature Specification: e2e cruise full pipeline (20261126-214101)

**Feature Branch**: `fabrik/issue-11978`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261126-214101 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue exists to exercise the `fabrik:cruise` pipeline contract (#898) end to end. Under cruise, an item should advance automatically through every stage up to Validate-complete, but the linked PR must not be auto-merged and the item must not move to Done until a human merges the PR. The code change itself is deliberately trivial, so that any failure observed is attributable to the pipeline's behaviour and not to the work. The issue carries the `fabrik:cruise` label.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added by the pipeline (Priority: P1)

The pipeline produces a PR that appends one marker line, unique to this run, to the cruise marker file.

**Why this priority**: The marker line is the only deliverable; it gives the pipeline a real, reviewable change to carry through every stage.

**Independent Test**: Inspect the PR diff and the final contents of `e2e/markers/cruise-full-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` already exists, **When** the change is implemented, **Then** the file's last line is `<!-- cruise-pipeline-20261126-214101 -->` and every prior line is unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line as its final line.
3. **Given** the PR is open, **When** its diff is viewed, **Then** apart from the engine-generated spec file under `specs/`, it touches exactly one file and adds exactly one line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

With `fabrik:cruise` applied, the item advances automatically through every stage to Validate-complete and stops there with the PR unmerged.

**Why this priority**: This is the contract under test (#898).

**Independent Test**: Observe the board column and PR state once Validate completes.

**Acceptance Scenarios**:

1. **Given** the item carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage with no human action.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the linked PR is open and unmerged and the item is not in Done.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P1)

After a human merges the PR, the item reaches Done and the issue closes correctly.

**Why this priority**: It verifies the second half of the cruise contract.

**Independent Test**: Merge the PR by hand and observe the issue state and board column.

**Acceptance Scenarios**:

1. **Given** the item is at Validate-complete with an open PR, **When** a human merges the PR, **Then** the issue is closed and the item moves to Done.

---

### Edge Cases

- The marker file does not exist: it is created, and the marker is its only line.
- The file's existing last line has no trailing newline: the marker still lands on its own line, with prior content unchanged.
- Plan must not decompose the work or spawn child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends the single line `<!-- cruise-pipeline-20261126-214101 -->` at the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it is created first, so that the marker line is its final line.
- **FR-003**: Apart from the engine-generated spec file under `specs/11978-*/spec.md`, the change modifies exactly one file and adds exactly one line; no existing content is altered or removed.
- **FR-004**: The work is not decomposed into child issues.
- **FR-005**: Under `fabrik:cruise`, the item auto-advances through all stages to Validate-complete without a human advancing it.
- **FR-006**: Under `fabrik:cruise`, the linked PR is not auto-merged and the item does not move to Done at Validate completion.
- **FR-007**: After a human merges the PR, the issue is closed and the item reaches Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Apart from the engine-generated spec file under `specs/`, the PR diff consists of exactly one added line in exactly one file, and that line is the specified marker.
- **SC-002**: The item reaches Validate-complete with no manual stage advancement.
- **SC-003**: At Validate-complete, the PR is unmerged and the issue is open.
- **SC-004**: After a human merge, the issue is closed and the item is in Done.

## Assumptions

- The marker line text is fixed by the request and must be reproduced exactly, including the run-specific timestamp.
- Appending means the marker is the last line of the file, followed by a trailing newline.
- The issue is scoped to the single repository where it lives.

## Out of Scope

- Any change to the cruise feature itself or to the engine.
- Any file other than `e2e/markers/cruise-full-pipeline.md`.
- Multi-repo behaviour.
- Decomposition into child issues.
