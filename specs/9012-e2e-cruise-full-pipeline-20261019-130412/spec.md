# Feature Specification: e2e cruise full pipeline (20261019-130412)

**Feature Branch**: `fabrik/issue-9012`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Trivial change: Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261019-130412 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` pipeline contract (#898). With `fabrik:cruise`, an item auto-advances through every stage up to Validate-complete. It does not auto-merge the PR and does not move the issue to Done by itself. A human merges the PR, and the issue should then close correctly.

The code change is deliberately trivial. It is a marker line in a single file, so the test exercises pipeline behaviour and not feature complexity. Each run uses a unique timestamped marker (`20261019-130412`) so the run can be told apart from earlier ones.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the e2e marker file (Priority: P1)

An e2e harness operator needs the pipeline to produce a one-line change to `e2e/markers/cruise-full-pipeline.md`. That change is the observable artifact of the run.

**Why this priority**: The marker line is the only deliverable. The rest of the verification depends on a PR existing with this change.

**Independent Test**: After Implement, the PR diff contains exactly one added line, `<!-- cruise-pipeline-20261019-130412 -->`, in `e2e/markers/cruise-full-pipeline.md`, positioned at the end of the file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is implemented, **Then** the file's last line is `<!-- cruise-pipeline-20261019-130412 -->` and no existing content is altered.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line as its last line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

An operator verifying the cruise contract needs the issue to auto-advance through Specify, Research, Plan, Implement, Review and Validate with no manual stage approvals. It must stop at Validate-complete with the PR open and unmerged.

**Why this priority**: This is the core contract under test.

**Independent Test**: With only `fabrik:cruise` set, observe that every stage completes without human intervention and that the PR is still open and unmerged once Validate completes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the PR is not merged and the issue has not been moved to Done by the engine.

---

### User Story 3 - Issue closes correctly after a human merges the PR (Priority: P2)

After a human merges the PR, the issue should close and reach its terminal state correctly.

**Why this priority**: It completes the end-to-end contract, but it depends on Stories 1 and 2.

**Independent Test**: Merge the PR by hand and confirm the issue closes and reaches Done.

**Acceptance Scenarios**:

1. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue is closed and moved to Done.

---

### Edge Cases

- The marker file does not exist yet: it is created with the marker line as its content.
- The marker file exists without a trailing newline: the marker still lands on its own line at the end.
- The marker file already contains markers from earlier runs: they are preserved, and only the new line is appended.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the exact line `<!-- cruise-pipeline-20261019-130412 -->` as the last line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created first and then given the marker line.
- **FR-003**: The change MUST touch only that one file and add only that one line.
- **FR-004**: Existing content of the marker file MUST NOT be modified or removed.
- **FR-005**: The work MUST NOT be decomposed into child issues; a single PR carries the whole change.
- **FR-006**: Under `fabrik:cruise`, the issue MUST auto-advance through all stages to Validate-complete without the PR being merged by the engine.
- **FR-007**: After a human merges the PR, the issue MUST close correctly.

### Key Entities

- **Marker file** (`e2e/markers/cruise-full-pipeline.md`): The file whose last line is the run's observable artifact.
- **Run marker** (`<!-- cruise-pipeline-20261019-130412 -->`): The unique per-run HTML comment line.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one file with exactly one added line, the run marker.
- **SC-002**: The issue reaches Validate-complete with zero manual stage approvals.
- **SC-003**: At Validate-complete, the PR is open and unmerged.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The `e2e/markers/` directory may or may not already exist; if it does not, creating it is part of creating the file.
- The marker line is plain text with no surrounding blank lines or extra whitespace.
- Single repo only; no cross-repo work is involved.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Child issues or decomposition of the work.
- Changes to the cruise pipeline behaviour itself; this issue only verifies it.
- Engine-initiated merging of the PR.
