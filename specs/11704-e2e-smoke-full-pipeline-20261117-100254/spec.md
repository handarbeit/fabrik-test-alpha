# Feature Specification: E2E smoke full-pipeline (20261117-100254)

**Feature Branch**: `fabrik/issue-11704`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261117-100254 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This is an end-to-end smoke test of the Fabrik pipeline in a single repository. The goal is to confirm that an issue can travel from Specify through to Done and end with a merged PR. The code change is deliberately trivial, so a failure points at the pipeline rather than at the change. The issue carries the `fabrik:yolo` label, so stages are expected to auto-advance and the PR is expected to be merged automatically once Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator runs a pipeline smoke test. The issue goes through every stage and ends with one marker line appended to the smoke marker file, merged to the default branch.

**Why this priority**: This is the only purpose of the issue. It proves the full pipeline works end to end.

**Independent Test**: After the issue reaches Done, inspect the default branch. The marker file must end with the marker line, and the merged PR must contain only that one-line change.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists, **When** the change is implemented, **Then** the file's final line is `<!-- smoke-full-pipeline-20261117-100254 -->` and all prior content is unchanged.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances automatically and the PR is merged, with the issue ending in Done.

---

### Edge Cases

- The marker file already ends without a trailing newline: the marker must still land on its own final line.
- The marker line is already present at the end of the file (for example on a re-run): it must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261117-100254 -->`.
- **FR-002**: If the file already exists, the marker line MUST be appended at the very end and existing content MUST be left unchanged.
- **FR-003**: If the file does not exist, it MUST be created containing the marker line.
- **FR-004**: The change MUST consist of exactly one file and one added line (apart from creating the file if absent).
- **FR-005**: The work MUST stay in this single repository. The Plan stage MUST NOT decompose it into child issues.
- **FR-006**: The PR MUST link to the issue so that merging it takes the issue to Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged, without manual intervention beyond the initial label.
- **SC-002**: The merged diff touches exactly one file and adds exactly one line.
- **SC-003**: The default branch's copy of the marker file ends with the marker line, and the line appears exactly once.

## Assumptions

- The `fabrik:yolo` label causes auto-advance through all stages and an automatic merge of the PR at Validate completion.
- The marker file's other content, if any, is irrelevant to this test and is not inspected.
- The marker string contains the timestamp `20261117-100254`, which makes it unique to this test run.

## Out of Scope *(optional)*

- Cross-repository work.
- Child-issue decomposition in Plan.
- Any change to engine behavior, stage configs or documentation.
- Any change to files other than the one marker file.
