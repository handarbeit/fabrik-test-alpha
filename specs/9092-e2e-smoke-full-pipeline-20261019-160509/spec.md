# Feature Specification: E2E smoke full-pipeline (20261019-160509)

**Feature Branch**: `fabrik/issue-9092`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261019-160509 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to verify that an issue can travel from Specify through Research, Plan, Implement, Review and Validate to Done, ending in a merged PR, with no human intervention (the `fabrik:yolo` label enables auto-advance and auto-merge). The code change is deliberately trivial so that any failure points at the pipeline, not at the change. The marker file `e2e/markers/smoke-full-pipeline.md` already exists and already holds earlier runs' marker lines, one per line, each appended at the end.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the run's marker line (Priority: P1)

An operator running the smoke suite wants this run to leave a single, uniquely identifiable marker in the repository, proving that the change went through the pipeline and was merged.

**Why this priority**: It is the only change this issue asks for; the pipeline traversal is validated by its landing.

**Independent Test**: After merge, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch is `<!-- smoke-full-pipeline-20261019-160509 -->`, and the diff of the merged PR is exactly that one added line in that one file.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier run markers, **When** the pipeline implements this issue, **Then** exactly one line, `<!-- smoke-full-pipeline-20261019-160509 -->`, is appended at the end of the file and no existing line is modified or removed.
2. **Given** the marker file does not exist, **When** the pipeline implements this issue, **Then** the file is created containing that single line.

---

### User Story 2 - Full unattended pipeline traversal (Priority: P1)

An operator wants the issue to reach Done with a merged PR without manual action, with Plan not decomposing the work into child issues.

**Why this priority**: This is the actual smoke objective.

**Independent Test**: The issue ends in the Done column, closed, with a merged PR linked via `Closes #9092`, and no child issues were spawned.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances automatically through to Validate and the PR is merged.
2. **Given** the Plan stage runs, **When** it produces its plan, **Then** it does not decompose the work and spawns no child issues.

---

### Edge Cases

- The existing file's last line lacks a trailing newline: the marker must still land on its own line.
- The marker line is already present (e.g. a retry after a partial run): it must not be added a second time.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends the exact line `<!-- smoke-full-pipeline-20261019-160509 -->` at the very end of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: The file is created first if it does not already exist.
- **FR-003**: The change modifies only that one file, and only by adding that one line; existing content is untouched.
- **FR-004**: The work stays within this single repository; no cross-repo work is performed.
- **FR-005**: The Plan stage does not decompose the work and no child issues are spawned.
- **FR-006**: The resulting PR is linked to the issue via `Closes #9092` and is merged, after which the issue is Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The default branch's marker file ends with the line `<!-- smoke-full-pipeline-20261019-160509 -->`, which appears exactly once.
- **SC-002**: The merged PR's diff is exactly one added line in one file.
- **SC-003**: The issue reaches Done with its PR merged, without any human intervention after labelling.
- **SC-004**: Zero child issues are created for this issue.

## Assumptions

- The marker file already exists and contains prior run markers; FR-002 covers the case where it does not.
- `fabrik:yolo` is already applied, so no human approval of the spec is required.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Cross-repo work and decomposition into child issues.
- Cleaning up or reorganising earlier marker lines.
