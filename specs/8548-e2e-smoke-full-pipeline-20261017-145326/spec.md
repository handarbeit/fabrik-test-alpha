# Feature Specification: E2E smoke full-pipeline marker (20261017-145326)

**Feature Branch**: `fabrik/issue-8548`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261017-145326 -->`. That's the entire change. One file, one line. Single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) and reach Done with a merged PR. The content change is deliberately trivial so that any failure points at the pipeline rather than the change. The issue carries the `fabrik:yolo` label, so stages are expected to auto-advance and the PR is expected to be merged automatically at the end of Validate.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator files this issue and lets Fabrik process it unattended. At the end, the repository's default branch contains the marker line at the end of `e2e/markers/smoke-full-pipeline.md`, delivered by a merged PR, and the issue is in Done.

**Why this priority**: This is the only deliverable. It both makes the change and proves the pipeline works end to end.

**Independent Test**: After the pipeline finishes, read `e2e/markers/smoke-full-pipeline.md` on the default branch. Its last line is `<!-- smoke-full-pipeline-20261017-145326 -->`, and the PR that introduced it is merged and linked to this issue.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is made, **Then** the marker line is appended as the new last line and all pre-existing content is unchanged.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is made, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances automatically, the PR is merged at the end of Validate, and the issue reaches Done.
4. **Given** the scope is a single repository and a single one-line change, **When** the Plan stage runs, **Then** it produces a single plan and does not decompose the work into child issues.

---

### Edge Cases

- The file exists but its last line has no trailing newline: the marker must still land on its own line rather than being joined to the previous line.
- The marker line is already present (for example, on a retry after a partial run): it must not be added a second time.
- The file exists with other marker lines from earlier smoke runs: those lines must be preserved.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` must end with the exact line `<!-- smoke-full-pipeline-20261017-145326 -->`.
- **FR-002**: If the file does not exist, it must be created containing the marker line.
- **FR-003**: If the file exists, the marker line is appended at the very end and existing content is preserved unchanged.
- **FR-004**: The change must touch only that one file and add only that one line (plus a file-creation if the file was absent), apart from the engine-committed spec artifact `specs/<issue>-<slug>/spec.md`.
- **FR-005**: The change must be delivered through the normal pipeline PR flow, with the PR linked to this issue so that it merges and the issue reaches Done.
- **FR-006**: The work must stay within this repository, and the Plan stage must not decompose it into child issues.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, a file that accumulates one HTML comment line per smoke run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The default branch's `e2e/markers/smoke-full-pipeline.md` ends with the marker line `<!-- smoke-full-pipeline-20261017-145326 -->`, appearing exactly once.
- **SC-002**: The PR diff for this issue consists of exactly one file with exactly one added line, apart from the new-file case and the engine-committed spec artifact `specs/<issue>-<slug>/spec.md`.
- **SC-003**: The issue reaches Done with its PR merged, with no human intervention beyond the initial `fabrik:yolo` label.
- **SC-004**: No child issues are created.

## Assumptions

- The directory `e2e/markers/` either already exists or may be created along with the file.
- The marker is an HTML comment, so it does not alter how the markdown renders.
- `fabrik:yolo` is expected to auto-advance every stage and trigger the merge, so no manual approval of this spec is needed.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Cross-repository work.
- Decomposition of the work into child issues.
- Changes to pipeline behavior or to stage configuration.
