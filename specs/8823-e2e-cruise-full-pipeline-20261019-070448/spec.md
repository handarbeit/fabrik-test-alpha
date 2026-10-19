# Feature Specification: e2e cruise full pipeline (20261019-070448)

**Feature Branch**: `fabrik/issue-8823`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261019-070448 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an end-to-end test fixture for the `fabrik:cruise` label. It exists to check the cruise pipeline contract (#898). Cruise auto-advances an issue through every stage up to Validate-complete. It does not auto-merge the PR and does not move the issue to Done. The issue should close correctly only after a human merges the PR. The code change is deliberately trivial, a single marker line in one file, so the run exercises the pipeline rather than any feature logic. The timestamp in the marker makes this run's change distinguishable from earlier runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added by the pipeline (Priority: P1)

The pipeline produces a PR whose only change is one HTML comment line appended at the end of `e2e/markers/cruise-full-pipeline.md`. The file is created first if it does not exist.

**Why this priority**: It is the only deliverable. Without it there is nothing for the stages to carry through the pipeline.

**Independent Test**: Inspect the PR diff. It touches `e2e/markers/cruise-full-pipeline.md`, apart from the engine's own spec file under `specs/`, and adds exactly the line `<!-- cruise-pipeline-20261019-070448 -->` as the last line of that file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** Implement completes, **Then** the file ends with `<!-- cruise-pipeline-20261019-070448 -->` and earlier content is unchanged.
2. **Given** the file does not exist, **When** Implement completes, **Then** the file exists and contains the marker line.

---

### User Story 2 - Cruise advances to Validate-complete without merging (Priority: P1)

With `fabrik:cruise` present, the issue moves through Specify, Research, Plan, Implement, Review and Validate on its own. It stops at Validate-complete with the PR still open.

**Why this priority**: This is the cruise contract the test exists to verify.

**Independent Test**: After Validate completes, the issue carries `stage:Validate:complete`, is not in Done, and the linked PR is open and unmerged.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the engine evaluates the landing decision, **Then** the PR is not merged and the issue does not move to Done.

---

### User Story 3 - Issue closes after a human merges (Priority: P2)

After a human merges the PR, the issue closes and reaches Done.

**Why this priority**: It completes the contract but depends on a human step outside the automated run.

**Independent Test**: Merge the PR by hand and confirm the issue closes and moves to Done.

**Acceptance Scenarios**:

1. **Given** Validate-complete and an open PR, **When** a human merges the PR, **Then** the issue closes and reaches Done.

---

### Edge Cases

- The marker file already exists with or without a trailing newline. The new line must still land on its own line at the end of the file.
- The marker file already contains markers from earlier runs. They must be preserved untouched.
- Plan must not decompose or spawn child issues for this one-line change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- cruise-pipeline-20261019-070448 -->` at the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created containing the marker line.
- **FR-003**: Apart from the engine-generated spec file under `specs/`, the change MUST touch no file other than `e2e/markers/cruise-full-pipeline.md`, and MUST add no line other than the marker.
- **FR-004**: The work MUST NOT be decomposed into child issues.
- **FR-005**: With `fabrik:cruise` present, the issue MUST auto-advance through all stages to Validate-complete.
- **FR-006**: The PR MUST NOT be auto-merged and the issue MUST NOT auto-advance to Done at Validate-complete.
- **FR-007**: After a human merges the PR, the issue MUST close correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one added line in `e2e/markers/cruise-full-pipeline.md`, equal to `<!-- cruise-pipeline-20261019-070448 -->`, and no other change apart from the engine's spec file.
- **SC-002**: The issue reaches Validate-complete with no human action between stages.
- **SC-003**: At Validate-complete the PR is open and the issue is not in Done.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The `e2e/markers/` directory is created along with the file if it does not exist.
- The marker line is plain text with a trailing newline, and no other formatting is required.
- The engine's own spec file under `specs/` is expected in the PR and does not count against the "one file, one line" constraint.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Changes to cruise or pipeline behavior. This issue only verifies it.
- Merging the PR, which is a human step.
- Decomposing the work into child issues.
- Single repo only; no cross-repo work.

## Source References *(optional)*

- #898 — the `fabrik:cruise` pipeline contract this test verifies.
