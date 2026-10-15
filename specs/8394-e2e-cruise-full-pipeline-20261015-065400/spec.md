# Feature Specification: e2e cruise full pipeline marker (20261015-065400)

**Feature Branch**: `fabrik/issue-8394`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261015-065400 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` autonomy contract (#898). Under cruise, an item should advance automatically through every pipeline stage up to Validate-complete, but the engine must not merge the linked PR. After a human merges the PR, the issue should close correctly.

To exercise that pipeline, the issue needs a real, minimal code change that produces a PR. The change is deliberately trivial: appending one uniquely stamped HTML comment line to a marker file. The work is not meant to be decomposed into sub-issues.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended to the cruise marker file (Priority: P1)

An operator running the cruise e2e check needs the repository to carry a uniquely stamped marker line in the cruise marker file, so that the pipeline has a concrete, reviewable diff to carry through every stage.

**Why this priority**: This is the entire deliverable. Without the change there is no PR, and the cruise pipeline contract cannot be verified.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the PR branch. Its final line must be exactly `<!-- cruise-pipeline-20261015-065400 -->`, and the PR diff must touch only that file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` already exists, **When** the change is implemented, **Then** the file's last line is `<!-- cruise-pipeline-20261015-065400 -->` and all prior content is unchanged.
2. **Given** `e2e/markers/cruise-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line as its last line.

---

### User Story 2 - Cruise pipeline runs to Validate-complete without merging (Priority: P1)

The operator needs the item to move through all stages on its own and stop at Validate-complete, leaving the PR open for a human to merge.

**Why this priority**: This is the contract being verified. It is the reason the issue exists.

**Independent Test**: Observe the item with the `fabrik:cruise` label. It reaches Validate-complete with the PR still open and unmerged, and it is closed only after a human merges the PR.

**Acceptance Scenarios**:

1. **Given** the item carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** the linked PR is not auto-merged and the item does not advance to Done.
3. **Given** a human merges the linked PR, **When** the engine next processes the item, **Then** the issue closes and the item reaches Done.

---

### Edge Cases

- The marker file already ends without a trailing newline: the appended marker must still land on its own line, not be joined to the previous content.
- The marker file already contains an earlier marker line from a different run: it is left untouched and the new line is added after it.
- The marker line is already the last line (a re-run): it is not duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/cruise-full-pipeline.md` MUST end with the exact line `<!-- cruise-pipeline-20261015-065400 -->`.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: Existing content in the file MUST be preserved unchanged, and the new line appended after it.
- **FR-004**: The change MUST touch only that one file and add only that one line.
- **FR-005**: The work MUST be planned and implemented as a single unit, with no decomposition into child issues.

### Key Entities

- **Cruise marker file**: `e2e/markers/cruise-full-pipeline.md`, which accumulates one uniquely stamped HTML comment line per cruise e2e run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff consists of exactly one added line in exactly one file.
- **SC-002**: The item reaches Validate-complete with no human action between stages.
- **SC-003**: The linked PR is still open and unmerged when Validate completes.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker file's location is fixed at `e2e/markers/cruise-full-pipeline.md`; the parent directory is created if absent.
- The timestamp `20261015-065400` is an opaque unique run identifier and is used verbatim.
- The marker is an HTML comment so it has no visible effect on rendered Markdown.

## Out of Scope

- Any change to engine, stage, or cruise behavior.
- Any additional files, documentation, or tests.
- Decomposing the work into sub-issues.
- Multi-repo scenarios.
