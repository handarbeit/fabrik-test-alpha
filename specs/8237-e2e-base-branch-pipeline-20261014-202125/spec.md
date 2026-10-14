# Feature Specification: e2e base-branch pipeline (20261014-202125)

**Feature Branch**: `fabrik/issue-8237`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description:

> ## Goal
>
> End-to-end verification of the base:<branch> (non-default base branch)
> pipeline contract — regression coverage for handarbeit/fabrik#1046,
> validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
> fixes.
>
> ## Trivial change
>
> Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):
>
> ```
> <!-- base-branch-pipeline-20261014-202125 -->
> ```
>
> That is the entire change. One file, one line. Plan should NOT decompose.
>
> ## Scope
>
> Single repo only, targeting the non-default base branch e2e-base-branch-20261014-202125 (already set via
> the base:<branch> label applied at filing — do not add or remove labels).
>
> ## Labels
>
> fabrik:cruise, base:e2e-base-branch-20261014-202125

## Background

This issue is a live end-to-end exercise of the pipeline's handling of issues that carry a `base:<branch>` label, meaning work is forked from, and its pull request targets, a branch other than the repository default. It is regression coverage for handarbeit/fabrik#1046, and it checks the fixes in #1047 (the issue-to-PR linkage for non-default bases) and #1050 (the review-gate data feed for such PRs). The code change is deliberately trivial, so that any failure points at the pipeline's base-branch handling and not at the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An item labelled `base:e2e-base-branch-20261014-202125` moves through the pipeline. One marker line is appended to `e2e/markers/base-branch-pipeline.md`. The resulting pull request targets the non-default base branch and is linked to this issue.

**Why this priority**: This is the whole purpose of the issue. It is the only deliverable and the thing the e2e run verifies.

**Independent Test**: After the pipeline completes, read `e2e/markers/base-branch-pipeline.md` on the base branch `e2e-base-branch-20261014-202125`. It must end with the marker line, and the PR must have targeted that branch and be linked to this issue.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist on the base branch, **When** the change is implemented, **Then** the file is created and contains the marker line `<!-- base-branch-pipeline-20261014-202125 -->`.
2. **Given** the marker file already exists, **When** the change is implemented, **Then** the marker line is appended at the very end of the file and the existing content is unchanged.
3. **Given** the issue carries the `base:e2e-base-branch-20261014-202125` label, **When** the PR is created, **Then** it targets that branch (not the repository default) and is linked to this issue.
4. **Given** the PR is open against the non-default base, **When** the review stage evaluates it, **Then** the review gate receives the PR's review data and the item proceeds through the pipeline.

---

### Edge Cases

- The marker file or its parent directories do not exist yet. They are created.
- The marker file exists but its last line has no trailing newline. The marker still ends up on its own line at the end.
- The marker line is already present. It is not duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261014-202125 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created containing the marker line.
- **FR-003**: The change MUST touch only that one file and add only that one line. No other files are changed.
- **FR-004**: The work MUST be based on, and its PR MUST target, the branch `e2e-base-branch-20261014-202125`.
- **FR-005**: The PR MUST be linked to this issue so that the pipeline can discover it.
- **FR-006**: The work MUST NOT be decomposed into child issues.
- **FR-007**: No labels may be added to or removed from the issue as part of this work, other than the labels the engine manages itself.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The marker line is the last line of `e2e/markers/base-branch-pipeline.md` on `e2e-base-branch-20261014-202125`, and the diff for the change is one added line (or one new one-line file).
- **SC-002**: The PR for this issue targets `e2e-base-branch-20261014-202125` and is linked to the issue.
- **SC-003**: The item moves through the pipeline stages for a non-default base without manual intervention, and no children are spawned.

## Assumptions

- The marker line is written exactly as given, and the file ends with a newline after it.
- If the file exists, it is a Markdown file where an HTML comment at the end is harmless.
- The base branch `e2e-base-branch-20261014-202125` already exists on the remote, as stated in the issue.
- The `fabrik:cruise` label means the pipeline advances through stages automatically and does not auto-merge.

## Out of Scope *(optional)*

- Changes to any file other than `e2e/markers/base-branch-pipeline.md`.
- Changes to the engine, its tests or its documentation.
- Adding or removing labels by hand.
- Splitting the work into multiple issues or child issues.
