# Feature Specification: E2E smoke full-pipeline marker (20261019-141423)

**Feature Branch**: `fabrik/issue-9036`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261019-141423 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the single-repo Fabrik pipeline. It checks that an issue can move from Specify through Research, Plan, Implement, Review and Validate to Done, ending in a merged PR. The change is deliberately trivial so that any failure points at the pipeline and not at the work. Each smoke run appends one uniquely timestamped marker line to a shared marker file. The file already exists and holds the lines from earlier smoke runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the smoke marker line (Priority: P1)

An operator running the pipeline smoke needs this run's unique marker line recorded in the shared marker file. The recorded line shows the run's change was implemented and merged.

**Why this priority**: This is the only change in the issue. The pipeline smoke has nothing to verify without it.

**Independent Test**: After the change, check that the last line of `e2e/markers/smoke-full-pipeline.md` is `<!-- smoke-full-pipeline-20261019-141423 -->` and that all earlier lines are unchanged.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists with earlier marker lines, **When** the change is applied, **Then** the file's final line is `<!-- smoke-full-pipeline-20261019-141423 -->` and every earlier line is unchanged.
2. **Given** the change is applied, **When** the PR diff is inspected, **Then** it touches only `e2e/markers/smoke-full-pipeline.md` and adds exactly one line.
3. **Given** the PR is open and the issue carries `fabrik:yolo`, **When** the pipeline runs to completion, **Then** the PR is merged and the issue reaches Done.

---

### Edge Cases

- The marker file does not exist: create it with the marker line as its content. It is expected to exist, so this is a fallback only.
- The file's last line has no trailing newline: the marker still goes on its own line, with no earlier line altered or merged into it.
- The marker line is already present: do not add a duplicate.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- smoke-full-pipeline-20261019-141423 -->` is appended as the last line of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: All existing content of that file is preserved unchanged.
- **FR-003**: If the file does not exist, it is created containing the marker line.
- **FR-004**: The change touches only `e2e/markers/smoke-full-pipeline.md`, with no other files modified.
- **FR-005**: The work stays within this single repository. The Plan stage does not decompose it into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged file ends with exactly one occurrence of the marker line `<!-- smoke-full-pipeline-20261019-141423 -->`.
- **SC-002**: The merged PR's diff is one added line in one file.
- **SC-003**: The issue reaches Done with its linked PR merged, with no human intervention beyond the `fabrik:yolo` label.

## Assumptions

- The marker file already exists, as confirmed in the repository, so file creation is a fallback only.
- Markers are one HTML comment per line, appended chronologically, matching earlier runs.
- The `fabrik:yolo` label means the spec needs no human approval before the pipeline advances.

## Out of Scope

- Any change to files other than `e2e/markers/smoke-full-pipeline.md`.
- Cross-repo work or decomposition into child issues.
- Cleanup or deduplication of earlier marker lines.
- Any change to pipeline, engine or stage configuration.
