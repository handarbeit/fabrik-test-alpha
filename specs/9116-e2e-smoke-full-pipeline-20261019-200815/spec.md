# Feature Specification: e2e smoke full-pipeline (20261019-200815)

**Feature Branch**: `fabrik/issue-9116`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261019-200815 -->`. That's the entire change. One file, one line. Scope: Single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) and reach Done with its pull request merged. The code change is deliberately trivial so that any failure points at the pipeline and not at the change. The `fabrik:yolo` label is applied so the issue auto-advances through the stages and auto-merges at the end without human gating.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged via the full pipeline (Priority: P1)

An operator files this issue and applies `fabrik:yolo`. Fabrik carries it through all pipeline stages. The result is a merged pull request that adds one marker line to the end of the smoke-test marker file.

**Why this priority**: This is the only story. It is the whole smoke test.

**Independent Test**: After the issue reaches Done, check the default branch. The marker file must end with the marker line and the linked PR must show as merged.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist, **When** the Implement stage runs, **Then** the file is created containing the marker line.
2. **Given** the marker file already exists, **When** the Implement stage runs, **Then** the marker line is appended as the final line and existing content is unchanged.
3. **Given** the change is implemented and reviewed, **When** Validate completes under `fabrik:yolo`, **Then** the linked PR is merged and the issue moves to Done.
4. **Given** the issue is single-repo and trivial, **When** the Plan stage runs, **Then** it produces a plan without spawning child issues.

---

### Edge Cases

- The marker file exists but lacks a trailing newline: the marker line must still land on its own line at the end of the file.
- The marker file already contains an identical marker line (for example from a rerun): the line is not duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261019-200815 -->`.
- **FR-002**: If the file does not exist, it MUST be created, and its only required content is the marker line.
- **FR-003**: If the file exists, its existing content MUST be left unchanged and the marker line appended at the very end.
- **FR-004**: The change MUST consist of exactly one file and one added line (excluding any file creation overhead).
- **FR-005**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-006**: The issue MUST be able to pass through all pipeline stages to Done with the linked PR merged.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline finishes, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch is exactly `<!-- smoke-full-pipeline-20261019-200815 -->`.
- **SC-002**: The merged diff touches exactly one file and adds one line.
- **SC-003**: The issue reaches Done with its linked PR merged, and no child issues were created.

## Assumptions

- The `e2e/markers/` directory may or may not exist; creating it as needed is acceptable.
- `fabrik:yolo` handles auto-advance and the final merge, so no manual approval is needed.
- The marker line is an HTML comment and is not intended to render in the document.

## Out of Scope

- Any change beyond the single marker line.
- Cross-repo work.
- Splitting the work into child issues.
- Changes to pipeline behavior or configuration.
