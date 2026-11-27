# Feature Specification: E2E smoke full-pipeline (20261127-015432)

**Feature Branch**: `fabrik/issue-12135`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261127-015432 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to show that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR. The code change is deliberately trivial so that any failure is attributable to the pipeline and not to the change. The issue carries the `fabrik:yolo` label, so stages advance automatically and the PR is merged automatically once Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator runs the smoke issue through the pipeline. The result is a merged PR that adds one marker comment line at the end of the smoke marker file, and the issue is Done.

**Why this priority**: This is the only deliverable of the issue. It is the observable evidence that the pipeline works end to end.

**Independent Test**: After the PR merges, read `e2e/markers/smoke-full-pipeline.md` on the default branch and confirm its final line is the marker comment.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line `<!-- smoke-full-pipeline-20261127-015432 -->`.
2. **Given** the file already exists, **When** the change is implemented, **Then** the marker line is appended as the last line and all prior content is unchanged.
3. **Given** the PR is merged, **When** the issue is inspected, **Then** it is at Done and the PR is linked and merged.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker must still go on its own line, not be joined to the previous line.
- The marker line is already present (for example after a re-run): it must not be duplicated.
- The Plan stage must not spawn child issues, since there is no cross-repo or multi-part work.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the single line `<!-- smoke-full-pipeline-20261127-015432 -->`, created first if it does not exist.
- **FR-002**: If the file already exists, its existing content MUST be preserved and the marker line appended at the very end, on its own line.
- **FR-003**: The marker line MUST appear exactly once in the file.
- **FR-004**: The change MUST touch only `e2e/markers/smoke-full-pipeline.md` (one file, one line added) in this repository, apart from the engine-generated spec file for this issue.
- **FR-005**: The Plan stage MUST NOT decompose the work or spawn child issues.
- **FR-006**: The change MUST be delivered through a PR that references this issue with `Closes #12135` and is merged to the default branch, leaving the issue at Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with no manual intervention beyond the initial labels.
- **SC-002**: A merged PR exists whose diff, excluding the spec file, adds exactly one line to one file.
- **SC-003**: On the default branch, the last line of `e2e/markers/smoke-full-pipeline.md` equals `<!-- smoke-full-pipeline-20261127-015432 -->` and that string appears once.
- **SC-004**: No child issues are created for this issue.

## Assumptions

- The `fabrik:yolo` label drives automatic advancement and merge, so no human approval is needed between stages.
- The marker is an HTML comment so it does not affect rendered Markdown output.
- If the `e2e/markers/` directory does not exist, it is created along with the file.

## Out of Scope *(optional)*

- Any cross-repo work or changes to other repositories.
- Any change beyond the single marker line.
- Decomposition of the work into child issues.
- Changes to the engine, stage configs, or tests.
