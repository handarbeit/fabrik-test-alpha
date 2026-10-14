# Feature Specification: E2E smoke full-pipeline marker (20261014-193827)

**Feature Branch**: `fabrik/issue-8196`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261014-193827 -->`. That's the entire change. One file, one line. Single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can move from Specify through Research, Plan, Implement, Review and Validate to Done, ending in a merged PR. The code change is deliberately trivial so that any failure points to the pipeline, not to the change. The issue carries the `fabrik:yolo` label, so stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on main via the full pipeline (Priority: P1)

An operator runs a smoke issue through every pipeline stage. The result is one new marker line in the smoke marker file, merged to the default branch, with the issue in Done.

**Why this priority**: This is the whole purpose of the issue. It confirms the pipeline works end to end.

**Independent Test**: After the issue reaches Done, check that `e2e/markers/smoke-full-pipeline.md` on `main` ends with the marker line and that the merged PR touched only that file.

**Acceptance Scenarios**:

1. **Given** the marker file exists, **When** the Implement stage runs, **Then** exactly one line, `<!-- smoke-full-pipeline-20261014-193827 -->`, is appended at the end of the file and no other content changes.
2. **Given** the marker file does not exist, **When** the Implement stage runs, **Then** the file is created and contains the marker line.
3. **Given** the PR is open with `Closes #8196` and the `fabrik:yolo` label is set, **When** Validate completes, **Then** the PR is merged and the issue moves to Done.

---

### Edge Cases

- The marker file already exists with earlier smoke lines. The new line is appended after them, and none are modified or reordered.
- The existing file's last line has no trailing newline. The appended marker must still sit on its own line.
- The marker line is already present from a previous run. The change is then a no-op and must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261014-193827 -->` after the change.
- **FR-002**: The change MUST add exactly one line and modify no other line or file.
- **FR-003**: The marker line MUST appear exactly once in the file.
- **FR-004**: The work MUST stay within this repository and the Plan stage MUST NOT decompose it into child issues.
- **FR-005**: The PR MUST include `Closes #8196` in its body.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches the Done column with its PR merged, with no manual intervention beyond the initial labels.
- **SC-002**: The merged PR diff consists of one file and one added line.
- **SC-003**: No child issues are spawned.

## Assumptions

- The marker file already exists on `main`. If it does not, creating it with the marker line satisfies the requirement.
- `fabrik:yolo` handles auto-advance through all stages and the auto-merge at Validate.

## Out of Scope *(optional)*

- Cross-repo work.
- Any change to engine code, stage configuration or documentation.
- Edits to other marker files.
