# Feature Specification: e2e cruise full pipeline (20261031-023606)

**Feature Branch**: `fabrik/issue-10356`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261031-023606 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898). Cruise should carry an item automatically through every pipeline stage up to Validate-complete, without merging the linked PR. After that, a human merges the PR, and the issue should close correctly. The code change is deliberately trivial: one marker line appended to one file. That keeps the pipeline mechanics, not the change, as the thing being tested. The unique timestamped marker lets the run be told apart from other e2e runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the marker file (Priority: P1)

An e2e operator wants a single, uniquely identifiable change to flow through the pipeline. The change is one HTML comment line appended to the cruise-pipeline marker file.

**Why this priority**: This is the only code change. Without it the pipeline has nothing to carry and the PR cannot be produced.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the issue branch. The last line must be exactly `<!-- cruise-pipeline-20261031-023606 -->`, and the PR diff must contain only that file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is implemented, **Then** the line `<!-- cruise-pipeline-20261031-023606 -->` is the final line of the file and all prior content is unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created containing the marker line as its last line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

An e2e operator wants to confirm that an item labelled `fabrik:cruise` advances through Specify, Research, Plan, Implement, Review and Validate on its own, and stops at Validate-complete with the PR left unmerged.

**Why this priority**: This is the contract under test (#898), the actual purpose of the issue.

**Independent Test**: Observe the item's stage-complete labels and board column across the run, and check the linked PR state at Validate-complete.

**Acceptance Scenarios**:

1. **Given** the item carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the linked PR is still open and unmerged, and the issue has not been advanced to Done.
3. **Given** the pipeline is at Validate-complete, **When** a human merges the PR, **Then** the issue is closed and moved to Done correctly.

---

### Edge Cases

- The marker file already ends without a trailing newline. The appended marker must still land on its own line, not be joined to the previous content.
- The marker file already contains the marker line, for example after a re-run. The line should not be duplicated.
- Plan must not decompose or spawn child issues, since the change is a single line in a single file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/cruise-full-pipeline.md` MUST end with the exact line `<!-- cruise-pipeline-20261031-023606 -->`.
- **FR-002**: If the file does not exist, it MUST be created, containing the marker line.
- **FR-003**: Existing content of the file MUST NOT be altered, and no other file MUST be modified by the implementation PR.
- **FR-004**: The change MUST NOT be decomposed into child issues.
- **FR-005**: With `fabrik:cruise` present, the item MUST auto-advance through all stages to Validate-complete without requiring human approval at any stage.
- **FR-006**: At Validate-complete the linked PR MUST remain unmerged and the issue MUST NOT be advanced to Done by the engine.
- **FR-007**: After a human merges the PR, the issue MUST end up closed.

### Key Entities

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, a text file whose last line identifies this e2e run.
- **Cruise label**: `fabrik:cruise`, the label that enables auto-advance without auto-merge.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file with exactly one added line, the specified marker.
- **SC-002**: The item reaches Validate-complete with zero manual stage approvals.
- **SC-003**: At Validate-complete the PR state is open (not merged) and the issue is open and not in Done.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The `fabrik:cruise` label is already on the issue and stays on it throughout.
- The marker value `20261031-023606` is a fixed literal from the request and is not regenerated.
- The `e2e/markers/` directory may need to be created if it does not exist.
- The marker file's other contents, if any, are irrelevant to the test and are left as they are.

## Out of Scope *(optional)*

- Any change to engine behaviour, stage configuration or documentation.
- Decomposition into multiple issues or cross-repo work.
- Merging the PR (a human action that is part of the verification, not of the implementation).
