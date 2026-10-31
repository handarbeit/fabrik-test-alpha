# Feature Specification: E2E cruise full-pipeline marker

**Feature Branch**: `fabrik/issue-10512`
**Created**: 2026-10-31
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261031-151136 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898). Under cruise, an item auto-advances through every pipeline stage up to Validate-complete, but the engine does not merge the PR or move the issue to Done. A human merges the PR, and the issue then closes. The code change is deliberately minimal: its only purpose is to give the pipeline a real, reviewable diff to carry through each stage. The change must not be decomposed into child issues.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is recorded in the repository (Priority: P1)

An operator running the e2e check needs a unique, run-specific marker line committed to the shared marker file, so the pipeline has a real one-line change to carry end to end.

**Why this priority**: The marker line is the only deliverable. Without it the pipeline has no diff to carry and the check proves nothing.

**Independent Test**: After the change lands, open `e2e/markers/cruise-full-pipeline.md` and confirm its final line is exactly `<!-- cruise-pipeline-20261031-151136 -->` and the diff against the base is only that one added line (plus file creation if the file was absent).

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` already exists, **When** the change is implemented, **Then** exactly one new line `<!-- cruise-pipeline-20261031-151136 -->` is appended at the end of the file and all existing content is unchanged.
2. **Given** `e2e/markers/cruise-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created containing the marker line.

---

### User Story 2 - Cruise carries the item to Validate-complete without merging (Priority: P1)

An operator wants to confirm that a cruise-labelled item advances automatically through every stage and stops at Validate-complete with the PR unmerged.

**Why this priority**: This is the cruise contract the issue exists to verify.

**Independent Test**: Observe the issue's stage labels and the linked PR after the pipeline settles.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without waiting for human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the linked PR is still open and unmerged, and the issue has not been moved to Done.
3. **Given** a human then merges the PR, **When** the merge completes, **Then** the issue closes correctly.

---

### Edge Cases

- If the file exists but its last line has no trailing newline, the marker must still land on its own line and not be joined onto the previous line.
- If the marker line is already present (for example on a re-run), it must not be duplicated.
- Planning must not split this one-line change into child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/cruise-full-pipeline.md` ends with the single line `<!-- cruise-pipeline-20261031-151136 -->` after the change.
- **FR-002**: If the file does not exist, it is created before the marker line is added.
- **FR-003**: Existing content in the file is left unchanged; the marker line is appended at the very end.
- **FR-004**: The marker line appears exactly once in the file.
- **FR-005**: The change touches no file other than `e2e/markers/cruise-full-pipeline.md`.
- **FR-006**: The work is delivered as a single issue and PR, with no child issues spawned.
- **FR-007**: Under `fabrik:cruise`, the item advances through all stages to Validate-complete without the engine merging the PR or moving the issue to Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one added line (the marker), plus the file's creation if it was absent, in exactly one file.
- **SC-002**: The item reaches Validate-complete with no human approval between stages.
- **SC-003**: At Validate-complete, the PR is open and unmerged.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker comment is HTML and is harmless in a Markdown file, so no rendering or content concerns apply.
- The timestamped marker value is unique to this run, so it will not collide with markers from other runs.
- The repository's default branch is `main`, and the PR targets it.

## Out of Scope

- Any change beyond the single marker line.
- Modifying engine, cruise, or pipeline behavior; this issue only verifies it.
- Decomposing the work into child issues.
- Multi-repo scenarios.
