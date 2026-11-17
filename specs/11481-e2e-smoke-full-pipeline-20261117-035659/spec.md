# Feature Specification: E2E smoke full-pipeline (20261117-035659)

**Feature Branch**: `fabrik/issue-11481`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261117-035659 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is a live end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to show that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR, without human intervention (the `fabrik:yolo` label enables auto-advance and auto-merge). The code change is deliberately trivial so that any failure points at the pipeline rather than at the work.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended via the full pipeline (Priority: P1)

An operator runs the smoke issue through the pipeline. The pipeline produces a PR that appends one marker comment line to the smoke marker file, and that PR is merged, leaving the issue in Done.

**Why this priority**: This is the whole purpose of the issue. It exercises every stage and the merge path.

**Independent Test**: After the run, inspect the merged PR's diff and the default-branch version of the marker file.

**Acceptance Scenarios**:

1. **Given** the marker file `e2e/markers/smoke-full-pipeline.md` already exists, **When** the pipeline completes, **Then** the file's last line is `<!-- smoke-full-pipeline-20261117-035659 -->` and all prior content is unchanged.
2. **Given** the marker file does not exist, **When** the pipeline completes, **Then** the file exists and its last line is `<!-- smoke-full-pipeline-20261117-035659 -->`.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances automatically through Validate, its PR is merged, and the issue reaches Done.

---

### Edge Cases

- If the existing file does not end with a newline, the marker must still appear on its own line at the end.
- Re-running the stage must not add the marker line twice.
- The Plan stage must not spawn child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- smoke-full-pipeline-20261117-035659 -->`, at the very end of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: Existing content in the file, if any, MUST be left unchanged.
- **FR-004**: The change MUST touch only that one file, in this one repository.
- **FR-005**: The Plan stage MUST NOT decompose the work or spawn child issues.
- **FR-006**: The work MUST be delivered through a PR that carries `Closes #11481` and is merged, ending with the issue in Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged PR's diff consists of exactly one added line in exactly one file.
- **SC-002**: The marker line appears exactly once in the marker file on the default branch after the merge.
- **SC-003**: The issue reaches Done with no human interaction beyond applying the `fabrik:yolo` label.
- **SC-004**: No child issues are created.

## Assumptions

- The marker is a Markdown/HTML comment, so it does not alter how the file renders.
- A trailing newline after the marker line is acceptable.
- The file lives at the repository-relative path `e2e/markers/smoke-full-pipeline.md`.

## Out of Scope *(optional)*

- Any cross-repo work.
- Any change to engine, stage configuration or other files.
- Decomposition into child issues.
