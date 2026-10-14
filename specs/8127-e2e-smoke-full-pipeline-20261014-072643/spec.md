# Feature Specification: E2E smoke full-pipeline marker line (20261014-072643)

**Feature Branch**: `fabrik/issue-8127`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261014-072643 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. Its purpose is to confirm that a single-repo issue can travel through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR, with no human intervention beyond the `fabrik:yolo` label. The code change is intentionally minimal so that any failure points at the pipeline and not at the change. The unique timestamped marker lets this run be told apart from earlier and later smoke runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator runs the smoke test. The issue is processed through all pipeline stages. The result is a merged PR that adds the run-specific marker line at the end of the smoke marker file.

**Why this priority**: This is the whole purpose of the issue. It exercises the complete pipeline with the smallest possible change.

**Independent Test**: After the pipeline completes, check the default branch. The last line of `e2e/markers/smoke-full-pipeline.md` is `<!-- smoke-full-pipeline-20261014-072643 -->`, and the issue is closed and in Done.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is made, **Then** the marker line is added as the new final line and all existing content is unchanged.
2. **Given** the marker file does not exist, **When** the change is made, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** every stage completes, **Then** the issue advances automatically, the PR is merged, and the issue reaches Done.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker must still land on its own line, not be joined to the previous line.
- The marker line is already present (for example on a re-run of the same stage): it must not be duplicated.
- Other marker files in `e2e/markers/` must not be modified.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/smoke-full-pipeline.md` MUST end with exactly one occurrence of the line `<!-- smoke-full-pipeline-20261014-072643 -->`.
- **FR-002**: If the file does not exist, it MUST be created, containing the marker line.
- **FR-003**: If the file exists, its prior content MUST be preserved unchanged, and the marker line MUST be appended as a separate final line.
- **FR-004**: The change MUST touch only this one file and add only this one line (plus a line break if needed to put the marker on its own line).
- **FR-005**: The work MUST be confined to this repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-006**: The change MUST be delivered as a PR that closes this issue and is merged to the default branch (`main`).

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with a merged PR, with no manual intervention beyond the `fabrik:yolo` label.
- **SC-002**: On `main`, the final line of `e2e/markers/smoke-full-pipeline.md` is `<!-- smoke-full-pipeline-20261014-072643 -->`, and the marker appears exactly once.
- **SC-003**: The merged PR's diff contains exactly one changed file and adds only the marker line.
- **SC-004**: No child issues are created.

## Assumptions

- The `e2e/markers/` directory may or may not already exist; creating it is implied if the file has to be created.
- The marker file is plain text and the HTML comment syntax is intentional, so it does not affect rendering.
- `fabrik:yolo` is already on the issue and is what auto-advances the pipeline past this stage's default of waiting for approval.
- The default base branch is `main`.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configs, or documentation.
- Cross-repo work or decomposition into child issues.
- Cleanup or deduplication of marker lines from earlier smoke runs.
