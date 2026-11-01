# Feature Specification: E2E smoke full-pipeline (20261101-062657)

**Feature Branch**: `fabrik/issue-10685`
**Created**: 2026-11-01
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261101-062657 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. It checks that a single-repo issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with its PR merged. The change itself is deliberately trivial so that any failure points at the pipeline and not at the work. The `fabrik:yolo` label is applied so the issue auto-advances through the stages and auto-merges at the end of Validate.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the base branch via the full pipeline (Priority: P1)

An operator wants confidence that the pipeline works end to end. The issue goes from Specify to Done unattended, and the marker line ends up on the default branch.

**Why this priority**: This is the only purpose of the issue.

**Independent Test**: After the issue reaches Done, check that the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch is `<!-- smoke-full-pipeline-20261101-062657 -->` and that the PR closing the issue is merged.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists, **When** the change is implemented, **Then** the marker line is appended as the final line and no existing content is altered.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created containing the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** Validate completes, **Then** the PR is merged and the issue reaches Done with no human intervention.

---

### Edge Cases

- The file exists but does not end with a newline: the marker must start on its own line, not be appended onto the previous line.
- The marker line is already present (for example after a retried stage): it must not be duplicated.
- The Plan stage must not split the work into child issues. Single-repo, single-file scope.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261101-062657 -->`.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: The change MUST be limited to that one file and that one added line; existing lines in the file MUST NOT be modified or removed.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The work MUST be done in this repository only and MUST NOT spawn child issues.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, a markdown file whose trailing HTML comment records that a smoke run completed.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its PR merged into the default branch, without human intervention beyond applying the `fabrik:yolo` label.
- **SC-002**: The merged diff touches exactly one file and adds exactly one line (more only if the file had to be created).
- **SC-003**: No child issues are created.

## Assumptions

- The `fabrik:yolo` label stays on the issue throughout, so each stage auto-advances and Validate auto-merges.
- The marker file may not exist yet; creating it is part of the change.
- The file is plain markdown, so an HTML comment line is invisible when rendered and harmless.

## Out of Scope *(optional)*

- Cross-repo work.
- Any change beyond the single marker line.
- Decomposition into child issues.
- Changes to engine code, stage configuration or documentation.
