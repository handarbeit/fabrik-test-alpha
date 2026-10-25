# Feature Specification: E2E cruise full pipeline marker (20261025-112932)

**Feature Branch**: `fabrik/issue-9505`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal — End-to-end verification of the fabrik:cruise pipeline contract (#898). ## Trivial change — Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261025-112932 -->`. That is the entire change. One file, one line. Plan should NOT decompose. ## Scope — Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` pipeline contract (#898). With cruise, an item should auto-advance through every stage up to Validate-complete without the engine merging the PR. After a human merges the PR, the issue should close correctly.

The code change is deliberately trivial so that the pipeline behavior is what gets exercised, not the feature. The change is one marker line appended to a single file. It should not be decomposed into child issues.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

An operator running the cruise e2e check needs a unique, harmless, per-run marker committed to the repository. This gives the pipeline a real but minimal diff to carry through Specify, Research, Plan, Implement, Review and Validate.

**Why this priority**: The marker is the only deliverable. Without it there is no PR for the pipeline to carry.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the PR branch. Its final line should be `<!-- cruise-pipeline-20261025-112932 -->`, and the PR diff should touch only that file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is made, **Then** the marker line is the last line of the file and all prior content is unchanged.
2. **Given** the file does not exist, **When** the change is made, **Then** the file is created and contains the marker line as its last line.
3. **Given** the change is complete, **When** the PR diff is inspected, **Then** exactly one file is changed and exactly one line is added.

---

### User Story 2 - Cruise pipeline contract holds (Priority: P1)

An operator wants to confirm that an issue labelled `fabrik:cruise` auto-advances through all stages to Validate-complete, that the engine does not merge the PR, and that the issue closes once a human merges it.

**Why this priority**: This is the stated goal of the issue (verifying #898).

**Independent Test**: Observe the issue through the pipeline. It reaches Validate-complete with the PR still open and unmerged. After a human merges the PR, the issue closes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** the PR is not merged by the engine and the issue is not moved to Done.
3. **Given** a human merges the PR, **When** the merge lands, **Then** the issue closes correctly.

---

### Edge Cases

- If the target file exists but does not end with a newline, the marker still goes on its own final line.
- If the marker line is already present, it is not duplicated.
- If the `e2e/markers/` directory does not exist, it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- cruise-pipeline-20261025-112932 -->` MUST be appended as the last line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created first, and the marker line MUST be its last line.
- **FR-003**: The change MUST be limited to that one file and that one added line, with no other edits.
- **FR-004**: Plan MUST NOT decompose the work into child issues.
- **FR-005**: Under `fabrik:cruise`, the issue MUST advance through all stages to Validate-complete without human approval at each stage.
- **FR-006**: The engine MUST NOT merge the PR or advance the issue to Done at Validate-complete.
- **FR-007**: After a human merges the PR, the issue MUST close correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file and exactly one added line, equal to `<!-- cruise-pipeline-20261025-112932 -->`.
- **SC-002**: The issue reaches Validate-complete with no human intervention between stages.
- **SC-003**: At Validate-complete the PR is open and unmerged, and the issue is not in Done.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The path `e2e/markers/cruise-full-pipeline.md` is relative to the repository root.
- Single repo only; no cross-repo work is involved.
- The marker is a plain HTML comment, which is inert in Markdown rendering.

## Out of Scope

- Any change other than the single marker line.
- Splitting the work into child issues.
- Changes to the cruise pipeline behavior itself.
- Merging the PR, which is done by a human after Validate-complete.
