# Feature Specification: E2E smoke full-pipeline marker (20261013-045221)

**Feature Branch**: `fabrik/issue-7692`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description:

> ## Goal
>
> End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.
>
> ## Trivial change
>
> Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):
>
> ```
> <!-- smoke-full-pipeline-20261013-045221 -->
> ```
>
> That's the entire change. One file, one line.
>
> ## Scope
>
> Single repo only — no cross-repo work. The Plan stage should NOT decompose.
>
> ## Labels
>
> fabrik:yolo

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. The goal is to confirm that a single-repo issue can move through every stage (Specify → Research → Plan → Implement → Review → Validate) and end at Done with a merged PR. The code change is deliberately trivial so that any failure points at the pipeline rather than at the change. The issue carries the `fabrik:yolo` label, so stages auto-advance and the PR is auto-merged when Validate completes.

The change appends one HTML comment line, unique to this run, to the smoke-test marker file. The run-specific token (`20261013-045221`) lets the run's change be told apart from earlier smoke runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged (Priority: P1)

An operator running the release smoke test wants the marker file to end with this run's unique comment line, delivered through a merged PR that was produced by the full pipeline.

**Why this priority**: This is the entire purpose of the issue. Without it, there is no evidence that the pipeline works end to end.

**Independent Test**: After the PR merges, read `e2e/markers/smoke-full-pipeline.md` on the default branch. Its last line is `<!-- smoke-full-pipeline-20261013-045221 -->`, and the PR diff touches only that file.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is applied, **Then** the file's final line is `<!-- smoke-full-pipeline-20261013-045221 -->` and all prior content is unchanged.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created and contains the marker line as its final line.
3. **Given** the PR is merged, **When** the issue is inspected, **Then** it is in the Done column and closed.

---

### User Story 2 - Pipeline runs without decomposition (Priority: P2)

An operator wants the whole run to stay on one issue and one PR, so the smoke result reflects a single straight-through pipeline pass.

**Why this priority**: Spawning child issues would change what is being tested and could leave unfinished work behind.

**Independent Test**: At the end of the run, no child issues were spawned from this issue and exactly one PR is linked to it.

**Acceptance Scenarios**:

1. **Given** the Plan stage runs for this issue, **When** it produces its plan, **Then** the plan is a single unit of work with no child issues spawned.
2. **Given** the issue is labelled `fabrik:yolo`, **When** each stage completes, **Then** the issue advances to the next stage without waiting for manual approval.

---

### Edge Cases

- The marker file exists but its last line has no trailing newline: the marker line must still land on its own line, not be joined to the previous content.
- The marker file exists and already contains markers from earlier smoke runs: those lines are preserved and the new line is added after them.
- The marker line is already present at the end of the file (for example a retried stage): it is not added a second time.
- The `e2e/markers/` directory does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261013-045221 -->`.
- **FR-002**: If the file does not exist, it MUST be created (including any missing parent directory) containing that line.
- **FR-003**: If the file exists, all existing content MUST be preserved unchanged, with the new line appended after it on its own line.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The change MUST consist of this one file and this one line only; no other file may be modified.
- **FR-006**: The work MUST stay within this single repository; no cross-repo work and no child issues are created.
- **FR-007**: The change MUST be delivered through a PR that links to this issue and is merged by the pipeline.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After merge, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch is exactly `<!-- smoke-full-pipeline-20261013-045221 -->`.
- **SC-002**: The merged PR's diff contains exactly one changed file with one added line, plus a missing-newline fix-up only if the previous last line lacked a trailing newline.
- **SC-003**: The issue reaches Done and is closed with a merged PR, with no manual intervention.
- **SC-004**: Zero child issues are spawned from this issue.

## Assumptions

- The marker line is added as a literal single line, with no extra blank lines before or after it.
- If the existing file's last line lacks a trailing newline, adding one so the marker sits on its own line is acceptable.
- The file is Markdown, so the HTML comment is invisible when rendered; this is intended.
- The default base branch is `main`, and no `base:` label overrides it.
- `fabrik:yolo` handles auto-advance and auto-merge, so no human approval is needed at any stage.

## Out of Scope

- Any change to files other than `e2e/markers/smoke-full-pipeline.md`.
- Cross-repo changes or decomposition into child issues.
- Changes to the smoke-test harness, pipeline configuration or stage definitions.
- Adding tests, documentation or other markers.
