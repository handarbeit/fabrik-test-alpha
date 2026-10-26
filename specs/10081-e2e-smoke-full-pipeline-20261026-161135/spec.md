# Feature Specification: E2E smoke full-pipeline (20261026-161135)

**Feature Branch**: `fabrik/issue-10081`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261026-161135 -->`. That's the entire change. One file, one line. Scope: Single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to show that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with its pull request merged. The code change is deliberately trivial, so a failure points at the pipeline and not at the work.

The issue carries the `fabrik:yolo` label, so each stage is expected to advance automatically and the PR is expected to merge automatically once Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator files this issue and applies `fabrik:yolo`. Fabrik takes it through every stage with no manual intervention. The result is a merged PR that adds one marker line to `e2e/markers/smoke-full-pipeline.md`, and the issue ends in Done.

**Why this priority**: This is the only story. It is the whole purpose of the smoke test.

**Independent Test**: After the run, check that `e2e/markers/smoke-full-pipeline.md` on the default branch ends with the line `<!-- smoke-full-pipeline-20261026-161135 -->`. Check that the linked PR is merged and the issue is in Done.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the file's final line is `<!-- smoke-full-pipeline-20261026-161135 -->` and all prior content is unchanged.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** every stage completes, **Then** the PR is merged automatically and the issue reaches Done.
4. **Given** the issue is single-repo and trivial, **When** the Plan stage runs, **Then** it produces a plan without spawning child issues.

---

### Edge Cases

- The marker file exists but its last line has no trailing newline: the marker must still be on its own line, not joined to the previous line.
- The marker file exists and already contains the same marker line: the line is not duplicated, and the final state still ends with the marker.
- The marker file exists but is empty: the file contains just the marker line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the exact line `<!-- smoke-full-pipeline-20261026-161135 -->` after the change merges.
- **FR-002**: If the file does not exist, it MUST be created, and its content MUST be the marker line.
- **FR-003**: If the file exists, existing content MUST be preserved and the marker line appended at the very end on its own line.
- **FR-004**: The change MUST touch exactly one file and add exactly one line (plus a separating newline if the existing file lacks a trailing one).
- **FR-005**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-006**: The linked PR MUST be merged, and the issue MUST reach Done, with no manual intervention given the `fabrik:yolo` label.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR in the merged state.
- **SC-002**: The merged diff contains exactly one changed file, `e2e/markers/smoke-full-pipeline.md`, with one added marker line.
- **SC-003**: The pipeline completes without any human comment, label change or manual stage move after `fabrik:yolo` is applied.
- **SC-004**: No child issues are created during the run.

## Assumptions

- The marker line is a literal HTML comment and is added exactly as written, including the timestamp suffix `20261026-161135`.
- A trailing newline after the marker line is acceptable and is the normal end of a text file.
- The default base branch (`main`) is the PR target; no `base:<branch>` label applies.
- The `e2e/markers/` directory may need to be created along with the file if it does not exist.

## Out of Scope

- Any change to engine, stage, or plugin behavior.
- Cross-repo work or multi-repo pipeline coverage.
- Decomposition of the work into child issues.
- Modifying any file other than `e2e/markers/smoke-full-pipeline.md`.
