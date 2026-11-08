# Feature Specification: e2e base-branch pipeline marker (20261108-032807)

**Feature Branch**: `fabrik/issue-11097`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261108-032807 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261108-032807 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end exercise of the `base:<branch>` pipeline contract, where an issue is worked on a non-default base branch. It is regression coverage for handarbeit/fabrik#1046 and checks the fixes for #1047 (issue↔PR linkage) and #1050 (review-gate data feed). The code change is deliberately trivial. What matters is that the item moves through every pipeline stage with the PR targeting the non-default base branch `e2e-base-branch-20261108-032807`, which is already set through the `base:` label applied when the issue was filed.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

A single marker line is appended to the base-branch pipeline marker file, and the change flows through the full pipeline as a PR targeting the non-default base branch.

**Why this priority**: This is the whole change, and the pipeline run it triggers is the purpose of the issue.

**Independent Test**: After the PR merges, `e2e/markers/base-branch-pipeline.md` on branch `e2e-base-branch-20261108-032807` ends with the line `<!-- base-branch-pipeline-20261108-032807 -->`, and the PR's base branch is `e2e-base-branch-20261108-032807`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the change is implemented, **Then** the file gains exactly one new final line, `<!-- base-branch-pipeline-20261108-032807 -->`, and all prior content is unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created containing that one line.
3. **Given** the issue carries the `base:e2e-base-branch-20261108-032807` label, **When** the pipeline creates the PR, **Then** the PR targets that branch, not the repository's other branches.

---

### Edge Cases

- If the file already ends without a trailing newline, the marker is still appended as its own final line.
- If the marker line is already present, it is not duplicated.
- The change must not add, remove or alter any labels, and must not touch any file other than the marker file, apart from engine-generated spec files.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/base-branch-pipeline.md` ends with exactly one new line, `<!-- base-branch-pipeline-20261108-032807 -->`, appended at the very end of the file.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created with that single line as its content.
- **FR-003**: The change is limited to that one file and one line, and the work is not decomposed into child issues.
- **FR-004**: The PR for this change targets the base branch `e2e-base-branch-20261108-032807` and links to this issue.
- **FR-005**: No labels are added to or removed from the issue as part of the change.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After merge, the last line of `e2e/markers/base-branch-pipeline.md` on `e2e-base-branch-20261108-032807` is exactly `<!-- base-branch-pipeline-20261108-032807 -->`.
- **SC-002**: The diff of the change touches one marker file and adds one line.
- **SC-003**: The item passes through all pipeline stages with its PR based on `e2e-base-branch-20261108-032807`, and the issue↔PR linkage and review-gate data are observed correctly.

## Assumptions

- The branch `e2e-base-branch-20261108-032807` exists on the remote and is the intended base.
- The marker is a plain Markdown HTML comment and needs no other content or formatting.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Splitting the work into child issues.
- Modifying issue labels.
- Work on any repository other than this one.
