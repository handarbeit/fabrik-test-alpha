# Feature Specification: E2E smoke full-pipeline (20261108-064512)

**Feature Branch**: `fabrik/issue-11259`
**Created**: 2026-11-08
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261108-064512 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. Its purpose is to confirm that a single-repo issue can travel through every stage (Specify → Research → Plan → Implement → Review → Validate) and finish at Done with its pull request merged. The code change is deliberately trivial so that any failure points at the pipeline, not at the change. The run is identified by the unique marker `20261108-064512`, which lets this run be told apart from other smoke runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is recorded by a merged change (Priority: P1)

An operator running the pipeline smoke wants a single, uniquely identifiable line to appear at the end of the smoke marker file on the base branch, proving the full pipeline ran and merged.

**Why this priority**: This is the whole purpose of the issue. The marker line landing on the base branch via a merged PR is the evidence that the pipeline works end to end.

**Independent Test**: After the PR merges, read `e2e/markers/smoke-full-pipeline.md` on the base branch and check that its last line is the marker comment.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is applied, **Then** the file's existing content is unchanged and the marker line is appended as the new last line.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created and contains the marker line.
3. **Given** the PR carrying the change has passed all pipeline stages, **When** it is merged, **Then** the issue reaches Done.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker line must still occupy its own line, not be joined to the previous content.
- The marker line for this run is already present (for example after a re-run): it must not be appended a second time.
- Markers from earlier smoke runs already exist in the file: they are left untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261108-064512 -->` after the change is applied.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, the change MUST create it, containing the marker line.
- **FR-003**: If the file exists, its prior content MUST be preserved unchanged; the change only appends.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The change MUST touch only `e2e/markers/smoke-full-pipeline.md` (one file, one added line).
- **FR-006**: The work MUST be confined to this repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-007**: The pull request carrying the change MUST be merged and the issue MUST reach Done through the pipeline's normal progression, with the `fabrik:yolo` label governing auto-advance and auto-merge.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: On the base branch, the last line of `e2e/markers/smoke-full-pipeline.md` is `<!-- smoke-full-pipeline-20261108-064512 -->`.
- **SC-002**: The merged change's diff contains exactly one changed file and exactly one added line (or one added file with one line, if the file was created).
- **SC-003**: The issue reaches Done with its linked PR merged, without manual intervention.
- **SC-004**: No child or cross-repo issues are created during the run.

## Assumptions

- The marker line is written exactly as given, including the HTML comment delimiters and the run identifier.
- A trailing newline after the marker line is acceptable and does not count as an additional line.
- The `fabrik:yolo` label stays on the issue for the whole run, so stages advance and the PR merges automatically.

## Out of Scope *(optional)*

- Any change to files other than `e2e/markers/smoke-full-pipeline.md`.
- Cross-repo work or decomposition into child issues.
- Changes to Fabrik engine code, stage configuration or documentation.
- Cleanup or consolidation of marker lines from earlier smoke runs.
