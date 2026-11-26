# Feature Specification: E2E smoke full-pipeline (20261126-215543)

**Feature Branch**: `fabrik/issue-11982`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261126-215543 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR. The change itself is deliberately trivial so that any failure points at the pipeline, not at the work. The issue carries the `fabrik:yolo` label, so stages advance automatically and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator runs a smoke test and expects Fabrik to append a unique marker comment line to the smoke marker file and land it on the default branch, with the issue reaching Done.

**Why this priority**: It is the only story. It exercises the whole pipeline end to end.

**Independent Test**: After the pipeline finishes, inspect the default branch. The marker file must end with the unique marker line, the PR must be merged, and the issue must be in Done.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists, **When** the pipeline completes, **Then** the file's last line is `<!-- smoke-full-pipeline-20261126-215543 -->` and its prior content is unchanged.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the pipeline completes, **Then** the file exists and contains the marker line.
3. **Given** the change is implemented, **When** Validate completes under `fabrik:yolo`, **Then** the PR is merged into `main` and the issue moves to Done.

---

### Edge Cases

- The marker file exists but its last line has no trailing newline: the marker must still land on its own line, not be joined to the preceding text.
- The marker line is already present at the end of the file (re-run): it must not be duplicated.
- The Plan stage must not decompose or spawn child issues, since the work is one file and one line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261126-215543 -->`.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: If the file exists, its existing content MUST be preserved and the marker line appended at the very end, on its own line.
- **FR-004**: The change MUST touch exactly one file and add exactly one line (plus a separating newline if the existing content lacks a trailing one).
- **FR-005**: The work MUST be done within this single repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-006**: The PR MUST include `Closes #11982` so Fabrik can link it to this issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its PR merged into `main`, with no human intervention beyond the `fabrik:yolo` label.
- **SC-002**: The merged diff is exactly one file with one added content line (the marker).
- **SC-003**: No child issues are created.

## Assumptions

- The marker file lives under the repository root at `e2e/markers/smoke-full-pipeline.md`.
- The marker line is an HTML comment, so it does not affect the rendered Markdown.
- A single trailing newline after the marker line is acceptable.

## Out of Scope

- Any cross-repo work.
- Any change other than the one marker line.
- Decomposing the work into child issues.
