# Feature Specification: E2E smoke full-pipeline (20261013-201102)

**Feature Branch**: `fabrik/issue-7796`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261013-201102 -->`. That's the entire change. One file, one line. Scope: Single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. The goal is to confirm that an issue can move through every stage (Specify → Research → Plan → Implement → Review → Validate) to Done, with its PR merged. The code change is deliberately trivial so that any failure points at the pipeline and not at the change. The issue carries the `fabrik:yolo` label, so stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator files this issue and lets the pipeline run unattended. A single marker line ends up at the end of `e2e/markers/smoke-full-pipeline.md` on the default branch, delivered by a merged PR, and the issue reaches Done.

**Why this priority**: This is the only purpose of the issue. It proves the pipeline works end to end.

**Independent Test**: After the pipeline finishes, inspect the default branch. The last line of `e2e/markers/smoke-full-pipeline.md` is the marker, and the issue is closed and in Done.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line at its end.
2. **Given** the marker file already exists, **When** the change is implemented, **Then** the marker line is appended after all existing content and existing content is unchanged.
3. **Given** the change is implemented and passes Review and Validate, **When** Validate completes under `fabrik:yolo`, **Then** the linked PR is merged and the issue moves to Done.
4. **Given** the issue is single-repo and trivial, **When** the Plan stage runs, **Then** it produces a single-task plan and spawns no child issues.

---

### Edge Cases

- The marker file does not exist: it is created, containing the marker line.
- The marker file exists but its last line has no trailing newline: the marker still lands on its own line and does not merge into the previous line.
- The marker line is already present, for example on a re-run: it is not duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` ends with the exact line `<!-- smoke-full-pipeline-20261013-201102 -->`.
- **FR-002**: If the file does not exist, it is created. If it exists, its prior content is preserved and the marker is appended after it.
- **FR-003**: The marker appears exactly once in the file.
- **FR-004**: The change touches exactly one file and adds exactly one line (plus a newline separator if the prior content lacks a trailing newline).
- **FR-005**: The work is delivered through a single PR in this repository whose body contains `Closes #7796`.
- **FR-006**: The Plan stage does not decompose the work and spawns no child issues.
- **FR-007**: No other repository is modified.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue passes Specify, Research, Plan, Implement, Review and Validate and reaches Done with no manual intervention beyond the initial `fabrik:yolo` label.
- **SC-002**: The PR diff against the default branch contains exactly one changed file and one added line.
- **SC-003**: The final line of `e2e/markers/smoke-full-pipeline.md` on the default branch equals the marker string, which occurs exactly once.
- **SC-004**: The linked PR is merged and the issue is closed.

## Assumptions

- The `fabrik:yolo` label is present, so stages auto-advance and the PR is auto-merged at Validate, without human approval of this spec.
- The `e2e/markers/` directory may not exist yet. Creating it is part of creating the file and does not count as an extra change.
- The marker line is an HTML comment, so it does not affect how the markdown renders.

## Out of Scope

- Any change other than the single marker line.
- Cross-repo work and decomposition into child issues.
- Tests, documentation updates or code changes unrelated to the marker file.
