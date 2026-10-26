# Feature Specification: E2E cruise full pipeline marker (20261026-162833)

**Feature Branch**: `fabrik/issue-10100`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261026-162833 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue exists to exercise the `fabrik:cruise` pipeline contract (#898) end to end. The code change is deliberately trivial so that the verification is about the pipeline's behavior, not the change itself: with `fabrik:cruise` set, the issue should auto-advance through every stage up to Validate-complete, the linked PR should not be merged automatically, and once a human merges the PR the issue should close correctly.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended to the cruise marker file (Priority: P1)

A single marker line is added at the end of `e2e/markers/cruise-full-pipeline.md` so the pipeline has a real, minimal change to carry through to a PR.

**Why this priority**: Without a change there is no PR, and the pipeline contract cannot be verified.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the PR branch and confirm its last line is `<!-- cruise-pipeline-20261026-162833 -->` and that this is the only change in the PR.

**Acceptance Scenarios**:

1. **Given** the file `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is implemented, **Then** the line `<!-- cruise-pipeline-20261026-162833 -->` is appended at the very end and existing content is unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created containing the line `<!-- cruise-pipeline-20261026-162833 -->`.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

With `fabrik:cruise` on the issue, the pipeline advances automatically through all stages to Validate-complete, leaving the PR unmerged for a human.

**Why this priority**: This is the contract under verification.

**Independent Test**: Observe the issue progress from Specify to Validate-complete with no manual stage advancement, and confirm the PR remains open and unmerged.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human action.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the PR is not auto-merged and the issue is not advanced to Done.
3. **Given** a human merges the PR, **When** the merge completes, **Then** the issue closes correctly.

---

### Edge Cases

- The marker file already exists with or without a trailing newline: the new line is still appended at the end so it is the last line of the file.
- The change must not be decomposed into child issues; Plan should produce a single task.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- cruise-pipeline-20261026-162833 -->` MUST be the last line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created containing that line.
- **FR-003**: Existing content of the marker file, if any, MUST be left unchanged.
- **FR-004**: The change MUST touch only that one file and add exactly one line.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: The PR MUST NOT be merged automatically by the pipeline; merging is performed by a human.

### Key Entities

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, a text file whose trailing HTML comment lines record e2e runs.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one file with exactly one added line, matching the specified marker.
- **SC-002**: The issue reaches Validate-complete with no manual stage advancement.
- **SC-003**: The PR is still unmerged when Validate completes.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The issue is single-repo and `fabrik:cruise` is already applied.
- The marker is an HTML comment so it does not render in Markdown output.
- The timestamp in the marker is a unique run identifier and is used verbatim.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Decomposing the work into child issues.
- Automatic merging of the PR.
- Multi-repo scenarios.
