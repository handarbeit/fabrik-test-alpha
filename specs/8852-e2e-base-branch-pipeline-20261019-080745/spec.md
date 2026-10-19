# Feature Specification: e2e base-branch pipeline (20261019-080745)

**Feature Branch**: `fabrik/issue-8852`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261019-080745 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261019-080745 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end probe of the `base:<branch>` contract, in which an issue carries a `base:<branch>` label and its work is forked from, and its PR targeted at, a branch other than the repository default. It is regression coverage for handarbeit/fabrik#1046 and validates two fixes: #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (the data feed to the review gate). The code change is deliberately trivial so that the only thing under test is the pipeline's handling of the non-default base, not the content of the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

The issue moves through the pipeline and produces a PR that appends one marker line to `e2e/markers/base-branch-pipeline.md`. The PR targets `e2e-base-branch-20261019-080745`.

**Why this priority**: This is the whole change, and the pipeline behavior it exercises is the point of the test.

**Independent Test**: After the PR merges, the last line of `e2e/markers/base-branch-pipeline.md` on branch `e2e-base-branch-20261019-080745` is `<!-- base-branch-pipeline-20261019-080745 -->`, and the PR's base was that branch.

**Acceptance Scenarios**:

1. **Given** the marker file already exists on the base branch, **When** the change is implemented, **Then** exactly one line is appended at the very end and no existing content is modified.
2. **Given** the marker file does not exist on the base branch, **When** the change is implemented, **Then** the file is created containing the marker line.
3. **Given** the issue carries the `base:e2e-base-branch-20261019-080745` label, **When** the pipeline creates the PR, **Then** the PR targets that branch rather than the repository default.

---

### Edge Cases

- The marker file exists without a trailing newline: the marker must still land on its own line at the end of the file.
- The marker line is already present (for example, the stage re-runs): it must not be appended a second time.
- Any change beyond the single marker line, or any extra file, is out of bounds.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The final content of `e2e/markers/base-branch-pipeline.md` MUST end with the line `<!-- base-branch-pipeline-20261019-080745 -->`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created, containing the marker line.
- **FR-003**: The change MUST consist of exactly one file and one added line, and MUST NOT modify any other line.
- **FR-004**: The work MUST be based on, and its PR MUST target, the branch `e2e-base-branch-20261019-080745`.
- **FR-005**: The `base:` label and any other labels on the issue MUST NOT be added or removed as part of this work.
- **FR-006**: The Plan stage MUST NOT decompose this work into child issues.

### Key Entities

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, which records e2e runs by appended HTML comment lines.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff contains exactly one file and one added line, and that line is the marker.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261019-080745`, and it is linked to this issue.
- **SC-003**: The item progresses through the pipeline stages without the review gate failing for lack of data on the non-default-base PR.

## Assumptions

- The branch `e2e-base-branch-20261019-080745` exists on the remote, as the issue states it was already set.
- The marker is a plain HTML comment, so appending it has no rendering effect on the file.

## Out of Scope *(optional)*

- Any change to engine code, tests, or documentation.
- Splitting the work into child issues.
- Changing labels on the issue.
- Multi-repo scenarios.
