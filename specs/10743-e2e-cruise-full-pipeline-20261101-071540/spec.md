# Feature Specification: e2e cruise full pipeline (20261101-071540)

**Feature Branch**: `fabrik/issue-10743`
**Created**: 2026-11-01
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261101-071540 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` pipeline contract (#898). With `fabrik:cruise` set, the item should advance automatically through every stage up to Validate-complete. It must not auto-merge its PR or advance to Done on its own. After a human merges the PR, the issue should close correctly.

The code change is deliberately trivial: one marker line appended to one file. It exists to give the pipeline a real, minimal diff to carry through the stages. The behaviour under test is the pipeline's, not the change's.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is recorded in the cruise marker file (Priority: P1)

The marker file `e2e/markers/cruise-full-pipeline.md` ends with the unique marker comment for this run, so the run produces a real, reviewable one-line diff.

**Why this priority**: Without a PR carrying the change, there is nothing for the cruise pipeline to advance or for a human to merge.

**Independent Test**: Inspect the PR diff and the file contents. The last line of the file is `<!-- cruise-pipeline-20261101-071540 -->`, and no other file or line changed.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is applied, **Then** the file's final line is `<!-- cruise-pipeline-20261101-071540 -->` and all earlier content is unchanged.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created and contains the marker line.

---

### User Story 2 - Cruise carries the issue through all stages without merging (Priority: P1)

With `fabrik:cruise` present, the issue auto-advances through every stage to Validate-complete, and the PR is left open for a human to merge.

**Why this priority**: This is the contract under test (#898).

**Independent Test**: Observe the issue's board column and labels across the run. It reaches Validate with the stage complete, and the PR is still open and unmerged.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the PR is not merged and the issue is not moved to Done by the engine.
3. **Given** a human merges the PR after Validate-complete, **When** the merge lands, **Then** the issue closes and reaches Done.

---

### Edge Cases

- The marker file already ends without a trailing newline: the marker must still occupy its own final line.
- The marker file already contains markers from earlier runs: they are preserved, and only this run's line is added at the end.
- The plan must remain a single unit of work, with no child issues spawned.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The final line of `e2e/markers/cruise-full-pipeline.md` MUST be exactly `<!-- cruise-pipeline-20261101-071540 -->`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created containing that line.
- **FR-003**: If the file exists, its prior content MUST be left unchanged. The marker is appended, not substituted.
- **FR-004**: The change MUST touch only that one file and add only that one line.
- **FR-005**: Planning MUST NOT decompose the work into child issues.
- **FR-006**: With `fabrik:cruise` set, the issue MUST auto-advance through all stages to Validate-complete without human approval at intermediate stages.
- **FR-007**: The engine MUST NOT merge the PR or advance the issue to Done at Validate-complete.
- **FR-008**: After a human merges the PR, the issue MUST close.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one file with exactly one added line, equal to the specified marker.
- **SC-002**: The issue reaches Validate-complete with no human action between stages.
- **SC-003**: At Validate-complete the PR is open and unmerged, and the issue is not in Done.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The directory `e2e/markers/` may or may not already exist; it is created if needed.
- A trailing newline after the marker line is acceptable and does not affect "last line" semantics.
- The engine's existing cruise behaviour (#898) is the reference contract.

## Out of Scope *(optional)*

- Any change to engine behaviour or to cruise itself.
- Any change to files other than `e2e/markers/cruise-full-pipeline.md`.
- Multi-repo scenarios.
- Decomposing the work into child issues.
