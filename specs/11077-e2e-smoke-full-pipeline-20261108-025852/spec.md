# Feature Specification: E2E smoke full-pipeline marker (20261108-025852)

**Feature Branch**: `fabrik/issue-11077`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261108-025852 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) and reach Done with its pull request merged. The change is deliberately trivial, so a failure points at the pipeline rather than at the work. The issue carries the `fabrik:yolo` label, so stages are expected to advance automatically and the linked PR is expected to be auto-merged once Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added through the full pipeline (Priority: P1)

An operator running the smoke test wants a unique marker line to appear at the end of the smoke-test marker file on the default branch. This shows that the pipeline carried the issue from spec to merged change without manual intervention.

**Why this priority**: This is the only deliverable. It is the proof that the pipeline works end to end.

**Independent Test**: After the PR merges, read `e2e/markers/smoke-full-pipeline.md` on `main` and confirm that its last line is `<!-- smoke-full-pipeline-20261108-025852 -->`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists, **When** the change is implemented, **Then** the file's final line is `<!-- smoke-full-pipeline-20261108-025852 -->` and all prior content is unchanged.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the change is merged, **When** the pipeline finishes, **Then** the issue is in Done, closed, with one merged PR that links to it.

---

### Edge Cases

- The file already ends without a trailing newline. The marker must still land on its own line and must not be joined to the previous line.
- The marker line is already present, for example on a re-run. It must not be duplicated.
- The Plan stage must not split this work into child issues. It is a single repo and a single one-line change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with exactly the line `<!-- smoke-full-pipeline-20261108-025852 -->`.
- **FR-002**: If the file does not exist, it MUST be created and contain the marker line.
- **FR-003**: Any existing content of the file MUST be preserved, with the marker appended after it.
- **FR-004**: The change MUST touch only that one file and add only that one line.
- **FR-005**: The work MUST stay within this single repository, with no cross-repo changes and no spawned child issues.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, an append-only log of smoke-run marker comments.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: On the default branch, the marker file's last line is `<!-- smoke-full-pipeline-20261108-025852 -->` and the marker appears exactly once.
- **SC-002**: The merged PR's diff consists of a single file with a single added line, or a single new file with that line if it did not exist before.
- **SC-003**: The issue reaches Done and is closed with its linked PR merged, with no manual intervention beyond the initial label.

## Assumptions

- The default base branch is `main`.
- `fabrik:yolo` handles auto-advance between stages and auto-merge at the end of Validate.
- The marker is an HTML comment, so it does not affect how the file renders.

## Out of Scope

- Any change to files other than `e2e/markers/smoke-full-pipeline.md`.
- Cross-repo work, decomposition into child issues, and changes to pipeline or engine behavior.
