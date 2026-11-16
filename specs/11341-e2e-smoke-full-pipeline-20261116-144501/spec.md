# Feature Specification: e2e smoke full-pipeline (20261116-144501)

**Feature Branch**: `fabrik/issue-11341`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261116-144501 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is a live end-to-end smoke test of the Fabrik pipeline. It checks that one small, low-risk change can pass through every stage (Specify → Research → Plan → Implement → Review → Validate) and end as a merged PR with the issue in Done. The change itself is deliberately trivial so that any failure points at the pipeline, not the work. The `fabrik:yolo` label means stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator verifying the pipeline files this issue. Fabrik takes it through every stage with no human intervention, and the marker line ends up on the default branch.

**Why this priority**: This is the only purpose of the issue. A successful run proves the pipeline works end to end for a single repo.

**Independent Test**: After the run, check that `e2e/markers/smoke-full-pipeline.md` on `main` ends with the marker line, that the linked PR is merged, and that the issue is closed and in Done.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the pipeline implements this issue, **Then** the file's last line is `<!-- smoke-full-pipeline-20261116-144501 -->` and all prior content is unchanged.
2. **Given** the file does not exist, **When** the pipeline implements this issue, **Then** the file is created and contains the marker line as its final line.
3. **Given** the change is implemented, **When** the pipeline reaches the end of Validate under `fabrik:yolo`, **Then** the PR is merged and the issue moves to Done.

---

### Edge Cases

- The marker line is already present at the end of the file (for example on a re-run): it must not be appended a second time.
- The existing file does not end with a trailing newline: the marker must still land on its own line.
- Plan must not decompose the work into child issues, because the scope is a single one-line change in one repository.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261116-144501 -->`.
- **FR-002**: If the file does not exist, it MUST be created, and its only content is the marker line.
- **FR-003**: If the file exists, its existing content MUST be preserved and the marker appended at the end.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The change MUST touch only that one file in the one repository, with no other files modified and no cross-repo work.
- **FR-006**: The Plan stage MUST NOT decompose the work into child issues.
- **FR-007**: The change MUST be delivered through a PR that merges to the default branch, and the issue MUST reach Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the PR diff contains exactly one changed file and exactly one added line (plus the file's creation if it was absent).
- **SC-002**: The marker line appears exactly once as the final line of `e2e/markers/smoke-full-pipeline.md` on the default branch.
- **SC-003**: The issue reaches Done with its linked PR merged, with no human intervention and no child issues created.

## Assumptions

- The `e2e/markers/` directory is created if it does not exist, as part of creating the file.
- The marker is an HTML comment so it does not change the rendered Markdown.
- `fabrik:yolo` is expected to drive auto-advance through every stage and the auto-merge at Validate.

## Out of Scope

- Any change to other files or repositories.
- Any decomposition into child or follow-up issues.
- Any change to Fabrik engine behaviour; this issue only exercises it.
