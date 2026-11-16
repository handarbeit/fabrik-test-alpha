# Feature Specification: e2e base-branch pipeline marker (20261116-165054)

**Feature Branch**: `fabrik/issue-11399`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261116-165054 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Single repo only, targeting the non-default base branch e2e-base-branch-20261116-165054 (already set via the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is an end-to-end regression check for the `base:<branch>` pipeline contract: an issue carrying a `base:<branch>` label must flow through every pipeline stage with its work branch and pull request targeting that non-default base branch. It exercises the fixes for handarbeit/fabrik#1046 (the original regression), #1047 (issue-to-PR linkage for non-default-base PRs) and #1050 (the review-gate data feed for such PRs). The code change itself is deliberately trivial so that any failure is attributable to the pipeline, not to the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands via a non-default-base PR (Priority: P1)

A maintainer running the e2e suite files this issue with the `base:e2e-base-branch-20261116-165054` label. The pipeline appends one HTML comment line to the marker file, opens a PR targeting that base branch, and carries it through review and landing.

**Why this priority**: It is the only story; it is the whole purpose of the issue.

**Independent Test**: After the pipeline completes, inspect `e2e/markers/base-branch-pipeline.md` on the base branch `e2e-base-branch-20261116-165054` and the PR's base and linkage.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist on the base branch, **When** the change is implemented, **Then** the file is created and its final line is exactly `<!-- base-branch-pipeline-20261116-165054 -->`.
2. **Given** the marker file already exists, **When** the change is implemented, **Then** the line is appended at the very end and existing content is untouched.
3. **Given** the issue carries `base:e2e-base-branch-20261116-165054`, **When** the PR is created, **Then** its base is `e2e-base-branch-20261116-165054`, not any other branch, and it references this issue with `Closes #11399`.
4. **Given** the PR targets a non-default base, **When** the pipeline reaches review and landing, **Then** the PR is correctly linked to the issue and the review gate receives the PR's review data.

---

### Edge Cases

- The marker file is absent at the time of the change: it is created containing only the marker line.
- The marker file exists without a trailing newline: the marker line must still occupy its own line.
- The change must not alter labels, including `fabrik:cruise` and the `base:` label.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/base-branch-pipeline.md` ends with the single line `<!-- base-branch-pipeline-20261116-165054 -->`, created first if it does not exist.
- **FR-002**: The change touches exactly one file and adds exactly one line (plus a separating newline if needed).
- **FR-003**: The work branch and PR target the base branch `e2e-base-branch-20261116-165054`.
- **FR-004**: The PR body includes `Closes #11399` so the issue and PR are linked.
- **FR-005**: No labels are added to or removed from the issue by this work.
- **FR-006**: The work is not decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The diff of the PR against `e2e-base-branch-20261116-165054` is exactly one file with one added content line.
- **SC-002**: The issue progresses through all pipeline stages with the PR's base reported as `e2e-base-branch-20261116-165054` throughout.
- **SC-003**: The pipeline's issue-to-PR linkage and review-gate checks succeed for the non-default-base PR without manual intervention.

## Assumptions

- The marker file's current existence is unknown to the spec; either case is handled (create or append).
- The `fabrik:cruise` label means stages auto-advance, but the PR is not auto-merged by the pipeline.

## Out of Scope

- Any change besides the one marker line.
- Changes to pipeline code or to labels.
- Multi-repo scenarios.
