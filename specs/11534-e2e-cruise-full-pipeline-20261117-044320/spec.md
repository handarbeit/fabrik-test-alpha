# Feature Specification: e2e cruise full pipeline (20261117-044320)

**Feature Branch**: `fabrik/issue-11534`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261117-044320 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an end-to-end check of the `fabrik:cruise` pipeline contract (#898). Under that contract, an issue carrying `fabrik:cruise` auto-advances through every stage up to Validate-complete. It does not auto-merge the linked PR or move the issue to Done. A human merges the PR, and the issue then closes correctly.

The code change is deliberately trivial so that the pipeline behavior is the only thing under test. It is a single marker line appended to one tracking file. The line carries a unique timestamp, so each e2e run produces a distinguishable change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended to the tracking file (Priority: P1)

A reviewer opens the PR for this issue and sees that exactly one line was added at the end of `e2e/markers/cruise-full-pipeline.md`.

**Why this priority**: This is the whole deliverable. The pipeline run is only meaningful if the change itself is correct and minimal.

**Independent Test**: Diff the PR against the base branch. Confirm that the only change is one added line at the end of `e2e/markers/cruise-full-pipeline.md`, and that the line is exactly `<!-- cruise-pipeline-20261117-044320 -->`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is applied, **Then** the file's previous content is unchanged and a new final line `<!-- cruise-pipeline-20261117-044320 -->` is present.
2. **Given** `e2e/markers/cruise-full-pipeline.md` does not exist, **When** the change is applied, **Then** the file is created and contains the marker line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

An operator watching the issue sees it move through every pipeline stage automatically to Validate-complete, and the linked PR stays open and unmerged.

**Why this priority**: This is the contract being verified (#898).

**Independent Test**: Observe the issue's stage labels and board column through the run. Confirm that the PR is not merged and that the issue is not moved to Done by the engine at Validate completion.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the engine evaluates the landing decision, **Then** the PR is not auto-merged and the issue is not advanced to Done.
3. **Given** Plan runs, **When** it produces its output, **Then** it does not decompose the work into child issues.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

After the pipeline reaches Validate-complete, a human merges the PR, and the issue is closed and completed correctly.

**Why this priority**: It completes the end-to-end contract, but it depends on the first two stories.

**Independent Test**: Merge the PR by hand and confirm the issue closes and reaches Done.

**Acceptance Scenarios**:

1. **Given** the issue is at Validate-complete with an open PR, **When** a human merges the PR, **Then** the issue is closed and moved to Done.

---

### Edge Cases

- The marker file does not exist yet: it is created with the marker line as its content.
- The marker file exists but its last line has no trailing newline: the marker goes on its own new line, so the existing last line is not altered or joined to it.
- The marker file already contains marker lines from earlier runs: they are preserved and the new line is appended after them.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- cruise-pipeline-20261117-044320 -->`, to the end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created first, containing the marker line.
- **FR-003**: The change MUST NOT modify, reorder or remove any existing content in that file.
- **FR-004**: The change MUST touch only that one file.
- **FR-005**: The Plan stage MUST NOT decompose this work into child issues.
- **FR-006**: With `fabrik:cruise` present, the issue MUST auto-advance through all stages to Validate-complete without human stage approval.
- **FR-007**: The engine MUST NOT merge the linked PR, and MUST NOT advance the issue to Done, at Validate completion.
- **FR-008**: After a human merges the PR, the issue MUST close correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one file and exactly one added line, equal to `<!-- cruise-pipeline-20261117-044320 -->`.
- **SC-002**: The issue reaches Validate-complete with no human action between Specify and Validate.
- **SC-003**: At Validate-complete the linked PR is open and unmerged, and the issue is not in Done.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- `e2e/markers/` is the correct location. If the directory is missing, it is created together with the file.
- The marker is added as its own line, followed by a trailing newline, with no blank line before it.
- This is a single-repo change on the repository's default base branch (`main`).

## Out of Scope

- Any change to engine behavior, stage configuration or the cruise contract itself.
- Changes to any file other than `e2e/markers/cruise-full-pipeline.md`.
- Multi-repo or cross-repo behavior.
- Decomposing the work into child issues.
- Merging the PR as part of the pipeline run. A human does that afterward.

## Source References

- #898 (the `fabrik:cruise` pipeline contract)
