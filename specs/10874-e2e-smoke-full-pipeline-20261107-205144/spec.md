# Feature Specification: e2e smoke full-pipeline (20261107-205144)

**Feature Branch**: `fabrik/issue-10874`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261107-205144 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to verify that an issue can move through every stage (Specify → Research → Plan → Implement → Review → Validate) to Done with a merged PR. The code change is deliberately trivial so that any failure points to the pipeline rather than to the change itself. The `fabrik:yolo` label is set, so stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the base branch via the full pipeline (Priority: P1)

An operator runs the smoke issue through the pipeline. A PR is opened that appends one marker comment line to the smoke-full-pipeline marker file, and the PR merges.

**Why this priority**: It is the only story; it is the whole purpose of the smoke test.

**Independent Test**: After the issue reaches Done, check the marker file on the base branch. It ends with the marker line, and the linked PR is merged.

**Acceptance Scenarios**:

1. **Given** the marker file `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the file's final line is `<!-- smoke-full-pipeline-20261107-205144 -->` and all prior content is unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line as its last line.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the item advances automatically and ends in Done with a merged PR that closes this issue.

---

### Edge Cases

- The marker file already ends without a trailing newline: the marker must still occupy its own line.
- The marker line is already present (for example from a re-run): it must not be duplicated.
- Plan must not split the work into child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261107-205144 -->`, appended after any existing content.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: Existing content of the file, if any, MUST be left unchanged.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The change MUST touch only that one file, in this single repository, and MUST NOT spawn child issues or involve other repositories.
- **FR-006**: The PR for this change MUST link this issue with `Closes #10874`.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged, with no human intervention beyond the initial labels.
- **SC-002**: The merged diff changes exactly one file and adds exactly one line.
- **SC-003**: The marker file on the base branch ends with the specified marker line.

## Assumptions

- The marker file, if absent, is created with the marker line as its only content.
- The `fabrik:yolo` label drives auto-advance and auto-merge, so no manual approvals are expected.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Cross-repo work and decomposition into child issues.
- Changes to engine behavior, documentation or tests.
