# Feature Specification: E2E cruise full pipeline marker (20261026-051255)

**Feature Branch**: `fabrik/issue-9683`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261026-051255 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue exists to exercise the `fabrik:cruise` label end to end (contract #898). Under cruise, the issue should advance automatically through every pipeline stage up to Validate-complete. The engine should not merge the linked PR on its own. After a human merges the PR, the issue should close correctly.

The code change is deliberately trivial so that any failure points at pipeline behaviour, not at the content of the change. The marker line carries a unique timestamp (`20261026-051255`), so this run's change can be told apart from the changes of other runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is recorded in the cruise marker file (Priority: P1)

A maintainer running the cruise end-to-end check needs a minimal, uniquely identifiable change to flow through the pipeline. The change is one HTML comment line appended to the end of `e2e/markers/cruise-full-pipeline.md`.

**Why this priority**: This is the only code change in the issue. The pipeline run has nothing to carry without it.

**Independent Test**: Inspect the final content of `e2e/markers/cruise-full-pipeline.md` on the issue's branch. The last line is the marker comment, it appears exactly once, and all earlier content is unchanged.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` already exists, **When** the change is applied, **Then** the file's last line is `<!-- cruise-pipeline-20261026-051255 -->` and all previous content is unchanged.
2. **Given** `e2e/markers/cruise-full-pipeline.md` does not exist, **When** the change is applied, **Then** the file is created and contains the marker line as its content.

---

### User Story 2 - Cruise carries the issue through the pipeline without merging (Priority: P1)

A maintainer wants to confirm that an issue labelled `fabrik:cruise` auto-advances through all stages to Validate-complete, and that the linked PR is left unmerged for a human.

**Why this priority**: Verifying this contract is the purpose of the issue.

**Independent Test**: Observe the issue and its PR after Validate completes. The issue has advanced stage by stage with no manual stage approvals, and the PR is open and unmerged.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the linked PR is not merged by the engine and the issue is not moved to Done.

---

### User Story 3 - Issue closes correctly after a human merges the PR (Priority: P2)

After the pipeline stops at Validate-complete, a human merges the PR. The issue should then close and reach its terminal state.

**Why this priority**: This verifies the post-merge half of the cruise contract. It depends on the earlier stories.

**Independent Test**: Merge the PR by hand and confirm the issue closes and reaches Done.

**Acceptance Scenarios**:

1. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue is closed and moved to Done.

---

### Edge Cases

- The marker file already exists with no trailing newline: the marker must still end up on its own final line.
- The marker file already contains the marker line, for example after a retried stage: the line must not be duplicated.
- Plan must not split this issue into child issues. The work is one file and one line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- cruise-pipeline-20261026-051255 -->`, at the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created first, and the marker line MUST be its final line.
- **FR-003**: The implementation MUST touch only that one file; no other file is added or modified by it. The only other file in the PR is the spec at `specs/9683-e2e-cruise-full-pipeline-20261026-051255/spec.md`, which the engine writes during Specify.
- **FR-004**: Existing content of the marker file MUST be preserved unchanged.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: Under `fabrik:cruise`, the issue MUST auto-advance through all stages to Validate-complete without the engine merging the linked PR or moving the issue to Done.
- **FR-007**: After a human merges the PR, the issue MUST close correctly.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, a text file whose lines are HTML comments, one per run of this check.
- **Marker line**: `<!-- cruise-pipeline-20261026-051255 -->`, the unique identifier of this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Apart from the engine-written spec file, the PR diff contains exactly one changed file with one added line, the marker line.
- **SC-002**: The marker line appears exactly once in the marker file, as its last line.
- **SC-003**: The issue reaches Validate-complete with no manual stage approvals, and the PR is still open and unmerged at that point.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker file lives at `e2e/markers/cruise-full-pipeline.md` relative to the repository root, as the issue states.
- The marker line is an HTML comment, so it does not change how the file renders.
- The cruise contract (#898) is as described in the issue: auto-advance to Validate-complete, no auto-merge, and no advance to Done until a human merges.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Splitting the work into child issues.
- Changing cruise, yolo or other pipeline behaviour.
- Multi-repo scenarios.
