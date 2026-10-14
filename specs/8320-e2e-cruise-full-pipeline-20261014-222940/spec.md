# Feature Specification: E2E cruise full pipeline marker (20261014-222940)

**Feature Branch**: `fabrik/issue-8320`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261014-222940 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This issue is a live end-to-end test fixture for the `fabrik:cruise` label (#898). Cruise should carry an item through every pipeline stage automatically, stop at Validate-complete without merging the linked PR, and leave the merge to a human. The code change is deliberately trivial, so that any failure points at the pipeline rather than at the change. Each run appends a uniquely timestamped marker line to a shared file, so runs do not collide.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

A maintainer needs the shared cruise-pipeline marker file to gain exactly one new line identifying this run, so the pipeline run leaves a small, reviewable, verifiable diff.

**Why this priority**: The marker line is the only deliverable. The pipeline behavior under test depends on there being a real, minimal change to carry through the stages.

**Independent Test**: Inspect the diff of the resulting PR. It touches only `e2e/markers/cruise-full-pipeline.md` and adds only the line `<!-- cruise-pipeline-20261014-222940 -->` at the end of the file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists with prior marker lines, **When** the change is made, **Then** the final line of the file is `<!-- cruise-pipeline-20261014-222940 -->` and all prior content is unchanged.
2. **Given** the file does not exist, **When** the change is made, **Then** the file is created containing the marker line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

An operator running the e2e suite needs the issue, carrying `fabrik:cruise`, to auto-advance through every stage up to Validate-complete without the PR being merged automatically.

**Why this priority**: This is the contract under test (#898).

**Independent Test**: Observe the issue move through the stages with no human approval until Validate is complete, with the linked PR still open at that point.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the linked PR remains open and unmerged.
3. **Given** a human merges the PR, **When** the merge completes, **Then** the issue is closed and reaches Done.

---

### Edge Cases

- The marker file already contains many prior run lines. Only the new line is added, at the very end.
- The file's last line lacks a trailing newline. The new line must still land on its own line, without altering prior lines.
- The same marker value already exists. Not expected, because the timestamp is unique per run.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- cruise-pipeline-20261014-222940 -->` at the end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: The change MUST create `e2e/markers/cruise-full-pipeline.md` first if it does not exist.
- **FR-003**: The change MUST modify no other file and no existing line of the marker file.
- **FR-004**: The work MUST NOT be decomposed into child issues.
- **FR-005**: With `fabrik:cruise` present, the pipeline MUST auto-advance through all stages to Validate-complete without merging the linked PR.
- **FR-006**: After a human merges the PR, the issue MUST close.

### Key Entities

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, an append-only log of run identifiers written as HTML comments.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one added line and no deletions, in exactly one file.
- **SC-002**: The issue reaches Validate-complete with no manual stage approvals and with the PR unmerged.
- **SC-003**: After a human merge, the issue is closed.

## Assumptions

- The marker file already exists (it does on the current base), so the create-if-missing path is not exercised.
- The timestamp `20261014-222940` is the unique run identifier and is used as given.
- HTML comments do not affect rendered Markdown, so the added line has no visible effect.

## Out of Scope

- Any change beyond the single marker line.
- Splitting the work into child issues.
- Multi-repo behavior.
- Changes to the cruise mechanism itself.
