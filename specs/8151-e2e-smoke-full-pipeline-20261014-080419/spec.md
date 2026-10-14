# Feature Specification: E2E smoke full-pipeline marker (20261014-080419)

**Feature Branch**: `fabrik/issue-8151`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261014-080419 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. It checks that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR. The code change is deliberately trivial so that any failure points at the pipeline and not at the work: one marker line is appended to one file. The issue carries `fabrik:yolo`, so stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged (Priority: P1)

An operator running the smoke test wants a single, unique marker line to appear at the end of the smoke-test marker file on the default branch. That confirms the pipeline carried the issue from spec to merge.

**Why this priority**: This is the whole purpose of the issue. The marker line landing on `main` is the proof that the pipeline worked.

**Independent Test**: After the PR merges, read `e2e/markers/smoke-full-pipeline.md` on `main` and confirm its last line is the marker comment.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist on `main`, **When** the pipeline implements this issue, **Then** the file exists and contains the marker line `<!-- smoke-full-pipeline-20261014-080419 -->`.
2. **Given** the marker file already exists on `main`, **When** the pipeline implements this issue, **Then** the marker line is appended as the new last line and all existing content is unchanged.
3. **Given** the PR contains the change, **When** Validate completes under `fabrik:yolo`, **Then** the PR is merged and the issue reaches Done.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker must still go on its own line at the very end.
- The marker line is already present (for example on a re-run): it must not be duplicated.
- Plan could be tempted to split the work into child issues: it must not, since the change is one line in one file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261014-080419 -->`.
- **FR-002**: If the file does not exist, it MUST be created containing the marker line.
- **FR-003**: If the file exists, its prior content MUST be preserved and the marker line appended after it on its own line.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The change MUST touch only `e2e/markers/smoke-full-pipeline.md`, in only this repository.
- **FR-006**: The pipeline MUST NOT decompose the work into child issues.
- **FR-007**: The change MUST be delivered through a PR that closes this issue and merges to the default branch.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After merge, the last line of `e2e/markers/smoke-full-pipeline.md` on `main` is exactly `<!-- smoke-full-pipeline-20261014-080419 -->`.
- **SC-002**: The merged PR's diff contains one file and one added content line (plus the file creation if the file was new).
- **SC-003**: The issue reaches Done with its PR merged, with no manual intervention beyond the initial `fabrik:yolo` label.
- **SC-004**: No child issues are created for this issue.

## Assumptions

- The marker directory `e2e/markers/` may or may not already exist; creating it as part of creating the file is acceptable.
- The marker line is a Markdown/HTML comment, so it does not change how the file renders.
- The spec file the engine commits under `specs/` is the engine's own artifact and does not count against the single-file scope of the implementation change.

## Out of Scope *(optional)*

- Any change to engine, stage or plugin behavior.
- Cross-repo work.
- Changes to any file other than the marker file.
- Decomposing the work into sub-issues.
