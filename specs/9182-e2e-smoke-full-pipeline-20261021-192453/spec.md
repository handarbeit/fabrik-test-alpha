# Feature Specification: E2E smoke full-pipeline marker append (20261021-192453)

**Feature Branch**: `fabrik/issue-9182`
**Created**: 2026-10-21
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261021-192453 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to confirm that one issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) and end with a merged PR and the issue in Done. The code change is deliberately trivial so that any failure points at the pipeline rather than the change. The `fabrik:yolo` label is set, so each stage is expected to advance automatically and the PR is expected to be merged automatically when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended and merged through the full pipeline (Priority: P1)

An operator runs a pipeline smoke test. The issue goes from Specify to Done unattended. The result is one new marker line at the end of the smoke marker file, merged into the default branch.

**Why this priority**: This is the only story. It is the whole purpose of the issue.

**Independent Test**: After the issue reaches Done, read `e2e/markers/smoke-full-pipeline.md` on the default branch. Its last line is the marker comment, and the merged PR touches only that file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the marker line `<!-- smoke-full-pipeline-20261021-192453 -->` is appended at the very end of the file and no existing content is altered.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line `<!-- smoke-full-pipeline-20261021-192453 -->`.
3. **Given** the change is on the issue branch, **When** the pipeline completes under `fabrik:yolo`, **Then** a PR linked to this issue is merged into `main` and the issue ends in Done.
4. **Given** the issue is planned, **When** the Plan stage runs, **Then** it produces a single-repo plan and does not decompose the work into child issues.

---

### Edge Cases

- If the marker line is already present at the end of the file (for example on a re-run), it is not appended a second time.
- If the existing file does not end with a newline, the marker still lands on its own line, so the file's last line is exactly the marker.
- The marker must be written exactly as given, including the timestamp suffix `20261021-192453`.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` ends with the line `<!-- smoke-full-pipeline-20261021-192453 -->`.
- **FR-002**: If the file already exists, its prior content is preserved unchanged and the marker is appended after it.
- **FR-003**: If the file does not exist, it is created containing the marker line.
- **FR-004**: The marker line appears exactly once in the file.
- **FR-005**: The change modifies exactly one file and adds exactly one line.
- **FR-006**: The work is confined to this repository. No cross-repo changes are made and the Plan stage does not decompose it.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, the file the smoke test appends to.
- **Marker line**: the HTML comment `<!-- smoke-full-pipeline-20261021-192453 -->`, unique to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the issue is Done, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch equals the marker line exactly.
- **SC-002**: The merged PR's diff contains one changed file and one added line (excluding file creation overhead if the file was new).
- **SC-003**: The issue progresses from Specify to Done with a merged PR and no human intervention beyond applying the initial `fabrik:yolo` label.
- **SC-004**: No child issues are created for this issue.

## Assumptions

- The marker file may or may not exist at the start. The spec covers both cases (FR-002, FR-003).
- The marker is an HTML comment, so it does not change how the Markdown file renders.
- The `fabrik:yolo` label is applied, so stages auto-advance and the PR auto-merges once Validate passes.
- The default base branch is `main`.

## Out of Scope *(optional)*

- Any change to a file other than `e2e/markers/smoke-full-pipeline.md`.
- Cross-repo work.
- Decomposing the work into child issues.
- Changes to pipeline behaviour or test infrastructure. This issue only exercises them.
