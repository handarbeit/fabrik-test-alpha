# Feature Specification: e2e base-branch pipeline marker

**Feature Branch**: `fabrik/issue-11750`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261117-104501 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261117-104501 (already set via
the base:<branch> label applied at filing — do not add or remove labels).


## Labels

fabrik:cruise, base:e2e-base-branch-20261117-104501"

## Background

This issue is a live end-to-end test of the `base:<branch>` pipeline contract: an issue carrying a `base:<branch>` label must flow through every pipeline stage with its worktree, pull request and review gate all targeting that non-default base branch. It is regression coverage for handarbeit/fabrik#1046, and it validates the fixes in #1047 (issue-to-PR linkage for non-default-base PRs) and #1050 (the data feed for the review gate). The code change itself is deliberately trivial, so that any failure in the run can be attributed to the pipeline and not to the content of the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

A single marker line is appended to a marker file, delivered through a pull request that targets the branch named by the `base:` label rather than the repository default.

**Why this priority**: This is the whole change, and delivering it through the base-branch pipeline is the point of the test.

**Independent Test**: After the pipeline completes, inspect `e2e/markers/base-branch-pipeline.md` on branch `e2e-base-branch-20261117-104501` and confirm its final line is the marker comment and that the diff against the base is that one line (plus file creation if the file was absent).

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` already exists on the base branch, **When** the change is made, **Then** exactly one line, `<!-- base-branch-pipeline-20261117-104501 -->`, is appended at the very end of the file and no existing content is altered.
2. **Given** the file does not exist on the base branch, **When** the change is made, **Then** the file is created and contains the marker line as its final line.
3. **Given** the issue carries the `base:e2e-base-branch-20261117-104501` label, **When** the pipeline opens the pull request, **Then** the PR's base branch is `e2e-base-branch-20261117-104501`.

---

### Edge Cases

- The file exists but does not end in a newline: the marker must still land on its own final line.
- The marker line is already present (a re-run): it must not be duplicated.
- The `base:` label must not be added, removed or altered by the change or by the pipeline run.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- base-branch-pipeline-20261117-104501 -->` to the end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, the change MUST create it, with the marker line as its final line.
- **FR-003**: The change MUST touch no file other than `e2e/markers/base-branch-pipeline.md` and MUST add no line other than the marker line.
- **FR-004**: The work MUST target the non-default base branch `e2e-base-branch-20261117-104501` and be limited to this single repository.
- **FR-005**: The pull request MUST be linked to this issue and MUST target the branch named by the issue's `base:` label.
- **FR-006**: The labels on the issue MUST NOT be added to or removed by this work.
- **FR-007**: The work MUST be treated as a single unit and MUST NOT be decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The marker file on the base branch ends with exactly one occurrence of `<!-- base-branch-pipeline-20261117-104501 -->` after the pipeline completes.
- **SC-002**: The change's diff against the base branch is one added line (plus the file header if the file was newly created) in one file.
- **SC-003**: The pull request produced by the pipeline targets `e2e-base-branch-20261117-104501` and is recognised as linked to this issue.
- **SC-004**: The issue reaches the end of the pipeline with no manual intervention and no child issues created.

## Assumptions

- `e2e/markers/` may not exist; creating the directory and file together is acceptable.
- Because the issue's own labels declare the base branch, the repository default shown to the pipeline may coincide with it; the requirement is simply that the PR targets `e2e-base-branch-20261117-104501`.
- The marker line is plain Markdown with an HTML comment and needs no rendering or formatting beyond the exact text given.

## Out of Scope *(optional)*

- Any change to the engine, its tests, or its documentation.
- Any other file or any content beyond the single marker line.
- Decomposing the work into multiple issues.
- Adding or removing labels on the issue.
- Other repositories.
