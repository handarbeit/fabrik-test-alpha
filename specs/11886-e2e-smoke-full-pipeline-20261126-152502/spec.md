# Feature Specification: E2E smoke full-pipeline marker (20261126-152502)

**Feature Branch**: `fabrik/issue-11886`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261126-152502 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can move through every stage (Specify → Research → Plan → Implement → Review → Validate) and end at Done with a merged PR. The change itself is intentionally trivial so that any failure points to the pipeline, not the work. The `fabrik:yolo` label is applied so the issue auto-advances and auto-merges.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended and merged (Priority: P1)

An operator running the smoke test wants a single marker line appended to the smoke-full-pipeline marker file, delivered through the full pipeline as a merged PR, so that a successful run proves the pipeline works end to end.

**Why this priority**: It is the only deliverable of the issue.

**Independent Test**: After the PR merges, the marker file on the default branch ends with the marker line, and the issue is in Done.

**Acceptance Scenarios**:

1. **Given** the marker file exists with prior marker lines, **When** the change is merged, **Then** the file's last line is `<!-- smoke-full-pipeline-20261126-152502 -->` and all prior lines are unchanged.
2. **Given** the issue carries `fabrik:yolo`, **When** every stage completes, **Then** the PR is merged and the issue reaches Done without manual intervention.

---

### Edge Cases

- If the marker file did not exist, it is created containing just the marker line (it currently exists, so the line is appended).
- If the file's last line lacks a trailing newline, the marker still lands on its own line.
- The marker line must not be duplicated if a stage is re-run.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The marker file `e2e/markers/smoke-full-pipeline.md` MUST end with the exact line `<!-- smoke-full-pipeline-20261126-152502 -->` after the change.
- **FR-002**: The change MUST modify only that one file by adding only that one line; existing content is untouched.
- **FR-003**: The work MUST be confined to this repository; no cross-repo work and no decomposition into child issues.
- **FR-004**: The change MUST be delivered through a PR that closes this issue and is merged by the pipeline.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff consists of exactly one added line in exactly one file.
- **SC-002**: The issue reaches Done with its PR merged, with every stage completed and no manual intervention.
- **SC-003**: The marker line appears exactly once in the file on the default branch.

## Assumptions

- The marker file already exists, so no creation is needed.
- `fabrik:yolo` handles auto-advance and auto-merge, so no human approvals are required.

## Out of Scope *(optional)*

- Any cross-repo changes or child issue spawning.
- Changes to any file other than the marker file.
- Changes to pipeline or engine behavior.
