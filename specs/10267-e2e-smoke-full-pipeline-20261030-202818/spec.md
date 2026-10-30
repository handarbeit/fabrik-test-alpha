# Feature Specification: E2E smoke full-pipeline marker (20261030-202818)

**Feature Branch**: `fabrik/issue-10267`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261030-202818 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is a pipeline smoke test. Its purpose is to confirm that Fabrik can carry a single-repo issue through every stage (Specify → Research → Plan → Implement → Review → Validate) to a merged PR and a Done item, with `fabrik:yolo` auto-advancing between stages and auto-merging at the end. The code change is deliberately trivial so that any failure points at the pipeline, not at the change. Each smoke run appends a uniquely timestamped marker line to a shared marker file; the file already exists and holds markers from earlier runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator runs the smoke issue and expects the pipeline to produce a merged PR whose only effect is one new marker line at the end of the smoke marker file.

**Why this priority**: This is the only deliverable. It exercises the whole pipeline end to end.

**Independent Test**: After the run, inspect the default branch's `e2e/markers/smoke-full-pipeline.md`. Its last line is the marker for this run, and the file otherwise matches its prior content.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier markers, **When** the pipeline completes, **Then** the file's final line is `<!-- smoke-full-pipeline-20261030-202818 -->` and all previous lines are unchanged.
2. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances automatically, the PR merges after Validate, and the issue reaches Done.
3. **Given** the issue is single-repo, **When** the Plan stage runs, **Then** it does not decompose the work or spawn child issues.

---

### Edge Cases

- The marker file does not exist: it is created containing only the marker line.
- The file's last line has no trailing newline: the marker still lands on its own line, not joined to the previous one.
- The marker line is already present at the end of the file: no duplicate is added.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The marker line `<!-- smoke-full-pipeline-20261030-202818 -->` is appended at the very end of `e2e/markers/smoke-full-pipeline.md`, on its own line.
- **FR-002**: The file is created first if it does not already exist.
- **FR-003**: Existing content of the file is left unchanged.
- **FR-004**: The change touches only that one file and adds only that one line.
- **FR-005**: The work stays within this single repository; no cross-repo work and no child issues are created.
- **FR-006**: The PR is linked to this issue and merges to the default branch with the issue reaching Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff contains exactly one changed file and exactly one added line.
- **SC-002**: The last line of `e2e/markers/smoke-full-pipeline.md` on the default branch equals the marker text.
- **SC-003**: The issue reaches Done with its PR merged, with no manual intervention between stages.
- **SC-004**: No child or sub-issues are spawned.

## Assumptions

- The marker file already exists on the default branch (verified at Specify time) and ends with earlier smoke-run markers.
- Marker lines are order-sensitive only in that the newest run's line is last.
- `fabrik:yolo` drives automatic advancement and merge; no human approval step is expected.

## Out of Scope *(optional)*

- Any change to other files, including other marker files.
- Cross-repo work or decomposition into child issues.
- Cleaning up or reformatting earlier markers in the file.
