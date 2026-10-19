# Feature Specification: E2E smoke full-pipeline (20261019-054252)

**Feature Branch**: `fabrik/issue-8776`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261019-054252 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is a pipeline smoke test. Its purpose is to confirm that Fabrik can carry a single issue through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR, without human intervention (the issue carries `fabrik:yolo`). The code change is intentionally trivial so that any failure points at the pipeline, not at the change. The unique timestamped marker line lets the run be identified and verified afterwards.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the base branch via the full pipeline (Priority: P1)

An operator runs the smoke test to confirm the pipeline works end to end. The issue moves through all stages unattended, and a PR adding one marker line is merged to the base branch.

**Why this priority**: This is the whole purpose of the issue. A merged PR containing the marker proves the pipeline works.

**Independent Test**: After the run, check that the base branch's `e2e/markers/smoke-full-pipeline.md` ends with the marker line and that the issue is closed and in Done.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the pipeline completes, **Then** the file's final line is `<!-- smoke-full-pipeline-20261019-054252 -->` and all prior content is unchanged.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the pipeline completes, **Then** the file exists and contains the marker line.
3. **Given** the issue has the `fabrik:yolo` label, **When** each stage completes, **Then** the issue advances automatically through every stage to Done and its PR is merged.
4. **Given** the issue is single-repo, **When** the Plan stage runs, **Then** it does not decompose or spawn child issues.

---

### Edge Cases

- The marker file exists but its last line lacks a trailing newline: the marker must still land on its own line, not appended to the previous line.
- The marker file already contains this exact marker line (for example, a re-run): the line is not duplicated.
- The `e2e/markers/` directory does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST add the line `<!-- smoke-full-pipeline-20261019-054252 -->` as the last line of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it MUST be created containing the marker line.
- **FR-003**: The change MUST NOT alter any existing content of that file, and MUST NOT touch any other file in the repository, apart from the spec file the engine commits automatically for the Specify stage.
- **FR-004**: The marker line MUST appear on its own line, even if the existing file's last line has no trailing newline.
- **FR-005**: The marker line MUST appear exactly once in the file.
- **FR-006**: The work MUST be confined to this single repository; the Plan stage MUST NOT decompose the issue or spawn child issues.
- **FR-007**: The change MUST be delivered as a PR that closes this issue and is merged to the base branch.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its PR merged and no human intervention, with `fabrik:yolo` as the only autonomy label.
- **SC-002**: The merged diff consists of exactly one added line in `e2e/markers/smoke-full-pipeline.md` (or that file's creation containing only that line), plus the engine-committed spec file.
- **SC-003**: The base branch's `e2e/markers/smoke-full-pipeline.md` ends with `<!-- smoke-full-pipeline-20261019-054252 -->`.
- **SC-004**: No child issues are created.

## Assumptions

- The marker file path `e2e/markers/smoke-full-pipeline.md` is relative to the repository root.
- The marker line is an HTML comment, so it is inert in rendered Markdown.
- The default base branch is `main`.
- Specify-stage output (the engine-committed `specs/<issue>-<slug>/spec.md`) is an expected part of the PR and not counted as a deviation from the "one file, one line" scope.

## Out of Scope

- Any change other than the single marker line.
- Cross-repo work, decomposition or child-issue spawning.
- Changes to pipeline behavior, stage configuration or tests.
