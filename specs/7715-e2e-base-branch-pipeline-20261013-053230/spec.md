# Feature Specification: e2e base-branch pipeline marker (20261013-053230)

**Feature Branch**: `fabrik/issue-7715`
**Created**: 2026-10-13
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
> <!-- base-branch-pipeline-20261013-053230 -->
> ```
>
> That is the entire change. One file, one line. Plan should NOT decompose.
>
> ## Scope
>
> Single repo only, targeting the non-default base branch e2e-base-branch-20261013-053230 (already set via
> the base:<branch> label applied at filing — do not add or remove labels).
>
> Labels: fabrik:cruise, base:e2e-base-branch-20261013-053230

## Background

This issue is a live end-to-end test fixture. It exercises the `base:<branch>` pipeline, where an issue's work is forked from, and its PR targets, a branch other than the repository default. It is regression coverage for handarbeit/fabrik#1046. It validates two fixes:

- #1047: the linkage between the issue and its PR.
- #1050: the review-gate data feed.

The code change is deliberately trivial. The point is that the issue travels the full pipeline against the non-default base branch, not that the change is interesting.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended on the non-default base (Priority: P1)

A single marker line is appended to the end of `e2e/markers/base-branch-pipeline.md`. The work is delivered through a PR that targets the non-default base branch `e2e-base-branch-20261013-053230`.

**Why this priority**: This is the only deliverable. Its delivery through the pipeline is what the e2e test observes.

**Independent Test**: Inspect the PR diff and the resulting file on the base branch. The file's last line is the marker, and nothing else changed.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the change is applied, **Then** the file ends with the line `<!-- base-branch-pipeline-20261013-053230 -->` and all prior content is unchanged.
2. **Given** the file does not exist on the base branch, **When** the change is applied, **Then** the file is created containing the marker line.
3. **Given** the issue carries the `base:e2e-base-branch-20261013-053230` label, **When** the PR is opened, **Then** the PR targets that branch and is linked to this issue.

---

### Edge Cases

- The file already ends without a trailing newline. The marker must still land on its own line and not be joined to the previous line.
- The marker line is already present at the end of the file, for example on a re-run. It must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- base-branch-pipeline-20261013-053230 -->` is appended as the final line of `e2e/markers/base-branch-pipeline.md`, preserving all existing content.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created with the marker line as its content.
- **FR-003**: The change touches exactly one file and adds exactly one line (plus a separating newline if the file lacked a trailing one).
- **FR-004**: The work is delivered against the base branch `e2e-base-branch-20261013-053230`, as set by the existing `base:` label. No labels are added or removed as part of this work.
- **FR-005**: The work is not decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff for this issue contains exactly one changed file and one added marker line.
- **SC-002**: The PR targets `e2e-base-branch-20261013-053230` and is linked to issue #7715.
- **SC-003**: The pipeline carries the issue from Specify through Validate on the non-default base without manual intervention beyond the configured `fabrik:cruise` autonomy.

## Assumptions

- The marker file may or may not exist yet, and both cases are handled.
- The branch `e2e-base-branch-20261013-053230` exists on the remote, as the filing process set it up.
- No documentation, config, or test changes are needed, because the marker is itself the test artifact.

## Out of Scope

- Any change beyond the single marker line.
- Adding, removing, or modifying labels.
- Changes to engine behavior, stage configs, or docs.
- Multi-repo scenarios.
- Decomposing the work into child issues.
