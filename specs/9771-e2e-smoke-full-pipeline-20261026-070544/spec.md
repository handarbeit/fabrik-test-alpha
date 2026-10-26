# Feature Specification: E2E smoke full-pipeline marker (20261026-070544)

**Feature Branch**: `fabrik/issue-9771`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261026-070544 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to confirm that an issue can move from Specify through Research, Plan, Implement, Review and Validate to Done, ending in a merged PR. The code change is deliberately trivial, so that any failure points at the pipeline and not at the change. The `fabrik:yolo` label is present, so stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator running the smoke test wants a single, unique marker line added to the end of the smoke marker file. The change must go through every pipeline stage and be merged without human intervention.

**Why this priority**: This is the only story. It is the whole point of the smoke test.

**Independent Test**: After the PR is merged, check that `e2e/markers/smoke-full-pipeline.md` on the default branch ends with the marker line, that the PR changes only that file, and that the issue is in Done and closed.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is implemented, **Then** the file's last line is `<!-- smoke-full-pipeline-20261026-070544 -->` and all earlier content is unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** every stage completes, **Then** the PR is merged and the issue reaches Done with no manual action.

---

### Edge Cases

- The marker file exists but its last line has no trailing newline: the marker must still land on its own line, not appended to the previous line.
- The marker line is already present, for example after a retry: it must not be added a second time.
- The `e2e/markers/` directory does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261026-070544 -->`.
- **FR-002**: If the file does not exist, the change MUST create it (including any missing parent directory) containing the marker line.
- **FR-003**: If the file exists, the change MUST preserve all existing content and add the marker line on its own line at the very end.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The change MUST touch only that one file in a single repository, apart from the spec file the engine commits for this issue.
- **FR-006**: The Plan stage MUST NOT decompose this issue into child issues.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, a markdown file that accumulates one HTML-comment marker line per smoke run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done and is closed with no manual intervention after labelling.
- **SC-002**: The PR closing this issue is merged, and its diff shows exactly one added line in exactly one file, other than the engine-committed spec file.
- **SC-003**: The marker file on the default branch ends with `<!-- smoke-full-pipeline-20261026-070544 -->`, which appears once.
- **SC-004**: No child issues are created for this issue.

## Assumptions

- The default branch is `main`, and the PR targets it.
- The marker is an HTML comment, so it does not change how the markdown file renders.
- Other markers from earlier smoke runs may already be in the file. They are left untouched.

## Out of Scope *(optional)*

- Cross-repository changes.
- Any change to engine code, stage configuration, or other files.
- Decomposition into child issues.
- Changes to the CI configuration or test suite.
