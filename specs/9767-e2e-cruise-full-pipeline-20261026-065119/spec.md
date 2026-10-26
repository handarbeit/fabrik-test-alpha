# Feature Specification: E2E cruise full pipeline marker (20261026-065119)

**Feature Branch**: `fabrik/issue-9767`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261026-065119 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898). Under cruise, an item advances automatically through every pipeline stage (Specify → Research → Plan → Implement → Review → Validate) without waiting for human approval at each stage. At Validate completion it stops: the linked PR is not auto-merged and the issue does not move to Done. A human then merges the PR, and the issue should close correctly as a result.

The code change is deliberately trivial so the pipeline's behavior, not the change itself, is what gets verified: one HTML comment line appended to a marker file. The change is intended to be a single, non-decomposable unit of work.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the cruise marker file (Priority: P1)

The marker file `e2e/markers/cruise-full-pipeline.md` ends with the line `<!-- cruise-pipeline-20261026-065119 -->`, delivered through a PR linked to this issue.

**Why this priority**: This is the only deliverable. It gives the pipeline a real, reviewable change to carry through every stage.

**Independent Test**: Inspect the PR diff. It contains exactly one file change, and that change adds the marker line at the end of `e2e/markers/cruise-full-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is implemented, **Then** the exact line `<!-- cruise-pipeline-20261026-065119 -->` is appended as the last line and all existing content is unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created containing the marker line as its last (and only) line.
3. **Given** the change is complete, **When** the PR diff is reviewed, **Then** it touches only that one file and adds only that one line.

---

### User Story 2 - Cruise carries the issue through all stages without merging (Priority: P1)

With `fabrik:cruise` present, the issue moves through all stages up to Validate-complete automatically, and the PR stays unmerged.

**Why this priority**: This is the contract under verification, and the reason the issue exists.

**Independent Test**: Watch the issue from Specify through Validate with no human intervention between stages, then confirm the PR is still open and the issue is not in Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the next stage starts without human action.
2. **Given** Validate completes, **When** the engine evaluates the landing decision, **Then** the linked PR is not merged and the issue is not moved to Done.
3. **Given** Plan receives this spec, **When** it produces the plan, **Then** it does not decompose the work or spawn child issues.

---

### User Story 3 - Issue closes correctly after a human merges the PR (Priority: P2)

After Validate completes, a human merges the PR, and the issue then closes and reaches Done.

**Why this priority**: It completes the contract. The human-merge step is needed to confirm that cruise's non-merging behavior still ends in a correct close.

**Independent Test**: Merge the PR manually after Validate-complete and confirm the issue closes and moves to Done.

**Acceptance Scenarios**:

1. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue closes and is moved to Done.

---

### Edge Cases

- The marker file already exists with or without a trailing newline: the marker line must still end up on its own final line, with prior content intact.
- The marker file already contains an earlier `cruise-pipeline-*` comment from a previous run: it is left untouched, and only the new line is appended.
- The marker's timestamp (`20261026-065119`) must be reproduced exactly as written.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/cruise-full-pipeline.md` MUST end with the line `<!-- cruise-pipeline-20261026-065119 -->`.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: If the file exists, its existing content MUST be preserved and the line appended after it, on its own line.
- **FR-004**: The change MUST be limited to this one file and this one added line.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: With `fabrik:cruise` present, the issue MUST advance through all stages to Validate-complete without human approval between stages.
- **FR-007**: At Validate completion the linked PR MUST NOT be auto-merged, and the issue MUST NOT be advanced to Done by the engine.
- **FR-008**: After a human merges the PR, the issue MUST close and reach Done.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, a Markdown file whose trailing HTML comments record e2e runs. Comments do not render.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file with exactly one added line, equal to `<!-- cruise-pipeline-20261026-065119 -->`.
- **SC-002**: The issue reaches Validate-complete with no human action between Specify and Validate.
- **SC-003**: At Validate-complete the PR is open and unmerged, and the issue is not in Done.
- **SC-004**: After a human merges the PR, the issue is closed and in Done.
- **SC-005**: No child issues are created.

## Assumptions

- The marker file may or may not exist yet. Either case is handled by the requirements above.
- The `e2e/markers/` directory is the established location for e2e marker files.
- The cruise contract (#898) defines the expected behavior (auto-advance, no auto-merge, no auto-Done), and this issue verifies it rather than redefining it.

## Out of Scope *(optional)*

- Any change to engine behavior, stage configuration or cruise semantics.
- Any file change beyond the single marker line.
- Cleaning up or reorganizing earlier markers in the file.
- Verifying `fabrik:yolo` behavior.
