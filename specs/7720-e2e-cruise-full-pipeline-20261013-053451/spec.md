# Feature Specification: e2e cruise full pipeline marker (20261013-053451)

**Feature Branch**: `fabrik/issue-7720`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261013-053451 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This issue is a live end-to-end test fixture for the `fabrik:cruise` label (#898). Cruise should auto-advance an item through every pipeline stage up to Validate-complete, leave the PR unmerged, and let the issue close only after a human merges the PR. The code change is deliberately trivial: one marker line in one file. That keeps the test about pipeline behavior rather than feature content. The run-specific timestamp in the marker makes each run's change unique and identifiable.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the cruise marker file (Priority: P1)

A pipeline run for this issue produces a change that appends exactly one HTML comment line, `<!-- cruise-pipeline-20261013-053451 -->`, to the end of `e2e/markers/cruise-full-pipeline.md`. If the file does not exist, it is created first.

**Why this priority**: This is the whole deliverable. It gives the cruise pipeline something real to carry from Specify through Validate.

**Independent Test**: Inspect the PR diff. Apart from the engine-persisted spec artifact (`specs/7720-*/spec.md`), it touches only `e2e/markers/cruise-full-pipeline.md`, and the only added content in that file is the marker line at the end of the file.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is applied, **Then** the file's existing content is unchanged and the marker line is the final line.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created and contains the marker line as its last line.

---

### User Story 2 - Cruise carries the item to Validate-complete without merging (Priority: P1)

With `fabrik:cruise` on the issue, the item advances automatically through all stages to Validate-complete. The PR is left open and unmerged.

**Why this priority**: This is the pipeline contract the e2e test exists to verify.

**Independent Test**: After the run, the issue's stages are all complete through Validate and the linked PR is open and not merged.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances without waiting for human approval.
2. **Given** Validate has completed, **When** the engine reaches the end of the pipeline, **Then** the PR is not auto-merged and the issue is not moved to Done.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

After a human merges the PR, the issue closes and reaches its terminal state correctly.

**Why this priority**: It completes the contract, but it depends on the earlier stories and on human action.

**Independent Test**: Merge the PR by hand. The issue closes and moves to its terminal state.

**Acceptance Scenarios**:

1. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue is closed and reaches Done.

---

### Edge Cases

- The marker file already ends without a trailing newline. The marker must still land on its own line at the end of the file.
- The marker file already contains an identical marker line. The line is still appended once for this run; no deduplication is required.
- Plan is tempted to split the work into child issues. It must not, because the change is one file and one line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly the line `<!-- cruise-pipeline-20261013-053451 -->` at the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created before the line is appended.
- **FR-003**: The change MUST NOT modify or remove any pre-existing content of that file.
- **FR-004**: The change MUST touch no file other than `e2e/markers/cruise-full-pipeline.md`, apart from the spec artifact the engine's Specify stage persists under `specs/7720-*/spec.md`.
- **FR-005**: The Plan stage MUST NOT decompose the work into child issues.
- **FR-006**: Under `fabrik:cruise`, the item MUST advance through all stages to Validate-complete without human approval, and the PR MUST NOT be auto-merged.
- **FR-007**: After a human merges the PR, the issue MUST close correctly.

### Key Entities

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, which records one marker line per cruise e2e run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff changes exactly one file other than the engine-persisted spec artifact (`specs/7720-*/spec.md`): `e2e/markers/cruise-full-pipeline.md`, with exactly one added marker line, `<!-- cruise-pipeline-20261013-053451 -->`, at the end of the file.
- **SC-002**: Every pipeline stage through Validate is marked complete with no human intervention, and the PR is still open and unmerged at that point.
- **SC-003**: Within one poll cycle of a human merging the PR, the issue is closed.
- **SC-004**: No child issues are spawned.

## Assumptions

- The single repository is the one this issue is filed in, and the default base branch is `main`.
- The marker is plain text appended as a new line. It needs no trailing blank line.
- An already-existing marker file may hold markers from earlier runs. They are left untouched.

## Out of Scope

- Any change to engine behavior or to cruise semantics; this issue only exercises them.
- Multi-repo scenarios.
- Automating the human merge step.
- Decomposing the work into sub-issues.
