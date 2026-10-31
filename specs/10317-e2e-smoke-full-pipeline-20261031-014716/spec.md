# Feature Specification: e2e smoke full-pipeline (20261031-014716)

**Feature Branch**: `fabrik/issue-10317`
**Created**: 2026-10-30
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
> <!-- smoke-full-pipeline-20261031-014716 -->
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

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to show that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) and finish at Done with a merged PR. The code change is deliberately trivial, so a failure points at the pipeline and not at the change. The `fabrik:yolo` label is set, so stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

A maintainer files this issue and the pipeline carries it from Specify to Done. The result is a merged PR that adds one marker comment line to the end of the smoke-test marker file.

**Why this priority**: This is the only story. A merged PR containing exactly that line shows the pipeline works end to end.

**Independent Test**: After the PR merges, read `e2e/markers/smoke-full-pipeline.md` on the default branch. Its final line is `<!-- smoke-full-pipeline-20261031-014716 -->`. The PR's diff touches only that file and adds only that line, plus the file's creation if it did not already exist.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the marker line is appended at the very end of the file and all existing content is unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and its content ends with the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** every stage completes, **Then** the PR is merged and the issue reaches Done without manual intervention.
4. **Given** the issue is single-repo, **When** the Plan stage runs, **Then** it does not split the work into child issues.

---

### Edge Cases

- The file exists but its last line has no trailing newline: the marker must still land on its own line at the end, not joined to the previous line.
- The marker line is already present at the end of the file (for example on a re-run): it must not be added a second time.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` ends with the exact line `<!-- smoke-full-pipeline-20261031-014716 -->`.
- **FR-002**: If the file already exists, its existing content is preserved and the marker line is appended after it.
- **FR-003**: If the file does not exist, it is created with the marker line as its content.
- **FR-004**: The marker line appears on its own line, exactly once.
- **FR-005**: The change touches only `e2e/markers/smoke-full-pipeline.md` and no other file.
- **FR-006**: The work stays within this single repository, and the Plan stage does not decompose it into child issues.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, the file the smoke test appends to.
- **Marker line**: the HTML comment `<!-- smoke-full-pipeline-20261031-014716 -->`, unique to this smoke run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with a merged PR and no manual intervention.
- **SC-002**: On the default branch after merge, the last line of the marker file is exactly the marker line, and it occurs exactly once in the file.
- **SC-003**: The merged PR's diff contains one file and one added line, plus the file's creation if it was new.
- **SC-004**: No child issues are spawned and no other repository is touched.

## Assumptions

- The marker line text is taken exactly as given in the request, including the timestamp `20261031-014716`.
- If the file must be created, its only content is the marker line, with no extra heading or body text.
- A trailing newline after the marker line is acceptable and expected for a text file.
- Because `fabrik:yolo` is set, no human approval of the spec is needed before Research begins.

## Out of Scope *(optional)*

- Any change to a file other than the marker file.
- Cross-repo work or decomposition into child issues.
- Changes to Fabrik engine, stage configuration, or documentation.
- Tests or tooling for the marker file.
