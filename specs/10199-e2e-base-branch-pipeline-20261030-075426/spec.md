# Feature Specification: e2e base-branch pipeline marker (20261030-075426)

**Feature Branch**: `fabrik/issue-10199`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261030-075426 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261030-075426 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end check of the `base:<branch>` pipeline contract: an issue labelled `base:<branch>` must run through the whole pipeline against a non-default base branch. It is regression coverage for handarbeit/fabrik#1046, and it validates the fixes in #1047 (issue↔PR linkage) and #1050 (review-gate data feed). The code change itself is deliberately trivial, so that any failure points at the pipeline and not at the change. The target base branch is `e2e-base-branch-20261030-075426`, set by the `base:e2e-base-branch-20261030-075426` label applied when the issue was filed.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the run marker to the base-branch pipeline marker file (Priority: P1)

An e2e run needs a durable, uniquely identified trace in the repository showing that this issue's work was delivered through a PR targeting the non-default base branch.

**Why this priority**: This is the only change in the issue, and it is what the pipeline run is meant to carry through to the PR.

**Independent Test**: On the PR branch, check that `e2e/markers/base-branch-pipeline.md` exists, that its last line is exactly `<!-- base-branch-pipeline-20261030-075426 -->`, and that this is the only line added.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` already exists on the base branch, **When** the change is implemented, **Then** the file's existing content is unchanged and one new final line, `<!-- base-branch-pipeline-20261030-075426 -->`, has been appended.
2. **Given** `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the change is complete, **When** the PR diff is inspected, **Then** it touches exactly one file and adds exactly one line.

---

### Edge Cases

- The file exists but does not end with a newline: the marker must start on its own line and must not be joined to the previous content.
- The marker line is already present (for example, a re-run): it must not be added a second time.
- The file does not exist: it is created, and the marker is its only content.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the line `<!-- base-branch-pipeline-20261030-075426 -->` at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, the change MUST create it, and the marker line MUST be its content.
- **FR-003**: The change MUST NOT modify, remove or reorder any existing content in that file.
- **FR-004**: The change MUST touch only that one file and add only that one line.
- **FR-005**: The work MUST be delivered as a single PR that targets `e2e-base-branch-20261030-075426` and links to this issue.
- **FR-006**: The work MUST NOT be decomposed into child issues.
- **FR-007**: The change MUST NOT add or remove any labels on the issue, including `base:e2e-base-branch-20261030-075426`.

### Key Entities

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a plain-text file of HTML-comment lines that records e2e runs of this pipeline.
- **Run marker line**: `<!-- base-branch-pipeline-20261030-075426 -->`, the unique identifier of this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the PR merges, the last line of `e2e/markers/base-branch-pipeline.md` on `e2e-base-branch-20261030-075426` is exactly `<!-- base-branch-pipeline-20261030-075426 -->`.
- **SC-002**: The PR diff contains 1 changed file and 1 added line, plus a created-file entry only if the file did not already exist.
- **SC-003**: The PR's base branch is `e2e-base-branch-20261030-075426`, not the repository's original default branch.
- **SC-004**: The issue passes through every pipeline stage with no manual intervention and no child issues.

## Assumptions

- `e2e/markers/` already exists on the base branch (other marker files live there), but the file itself may not; if it is absent it is created.
- The marker is a Markdown HTML comment, so it does not change how the file renders.
- The base branch `e2e-base-branch-20261030-075426` already exists on the remote, as the issue states.
- A trailing newline after the marker line is acceptable and expected.

## Out of Scope *(optional)*

- Any change to engine code, stage configuration, tests or documentation.
- Creating, renaming or deleting labels, or changing the base branch.
- Multi-repo scope or decomposition into child issues.
- Cleaning up or editing other marker files in `e2e/markers/`.
