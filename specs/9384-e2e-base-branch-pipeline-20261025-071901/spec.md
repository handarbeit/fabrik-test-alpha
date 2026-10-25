# Feature Specification: e2e base-branch pipeline (20261025-071901)

**Feature Branch**: `fabrik/issue-9384`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261025-071901 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261025-071901 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end probe of the pipeline's behavior for an item that targets a non-default base branch (the `base:<branch>` label). That contract was the subject of an earlier regression (handarbeit/fabrik#1046), whose fixes covered the issue↔pull-request linkage (#1047) and the data fed to the review gate (#1050). The code change is deliberately trivial, a one-line marker, so that any failure observed is attributable to the pipeline's handling of the non-default base and not to the change's content. The item carries `fabrik:cruise` and `base:e2e-base-branch-20261025-071901`, so it is expected to move through all stages without human approval.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

A maintainer running the base-branch regression check files this issue. The pipeline appends one marker line to a marker file and delivers it as a pull request targeting the non-default base branch, not the repository's usual default.

**Why this priority**: This is the whole change and the signal the check exists to produce.

**Independent Test**: Inspect the pull request produced for this issue: its base branch, its diff, and the end of the marker file.

**Acceptance Scenarios**:

1. **Given** the marker file `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the change is implemented, **Then** the file ends with the line `<!-- base-branch-pipeline-20261025-071901 -->` and no other line is altered.
2. **Given** the marker file does not exist on the base branch, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries the `base:e2e-base-branch-20261025-071901` label, **When** the pull request is created, **Then** it targets `e2e-base-branch-20261025-071901`.

---

### User Story 2 - Pipeline contract holds for a non-default base (Priority: P2)

The pipeline's linkage and review-gate behaviors work for an item whose pull request targets a non-default base branch.

**Why this priority**: This is the regression coverage the issue is for. It rests on the pull request existing, so it comes second.

**Independent Test**: Observe the item through Review and Validate and confirm the linkage and review-gate behaviors below.

**Acceptance Scenarios**:

1. **Given** a pull request targeting the non-default base, **When** the engine looks for the item's linked pull request, **Then** it finds the pull request (the #1047 linkage contract).
2. **Given** the item reaches a stage with a review gate, **When** the gate evaluates the pull request, **Then** it receives the pull request's review data (the #1050 data-feed contract).
3. **Given** the item has the `fabrik:cruise` label, **When** every stage completes, **Then** the item advances through the stages without human intervention and without the engine auto-merging.

---

### Edge Cases

- The marker file may already exist with or without a trailing newline; the appended line must still be on its own line at the very end.
- The marker file may be absent on the base branch; it is then created rather than treated as an error.
- If the marker line is already present, nothing is added a second time.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- base-branch-pipeline-20261025-071901 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created first and contains the marker line.
- **FR-003**: The change touches only that one file and adds only that one line.
- **FR-004**: The pull request for this issue targets the branch `e2e-base-branch-20261025-071901`.
- **FR-005**: The issue's labels are neither added nor removed by the work of this issue.
- **FR-006**: The work is a single unit and is not decomposed into child issues.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a file whose trailing HTML-comment lines record runs of this check.
- **Base branch**: `e2e-base-branch-20261025-071901`, the non-default branch the pull request must target.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The pull request's diff against its base is one file with exactly one added line, the marker line.
- **SC-002**: The pull request's base branch is `e2e-base-branch-20261025-071901`.
- **SC-003**: The engine finds the issue's linked pull request and the review gate receives its review data, with no manual intervention.
- **SC-004**: The item completes all pipeline stages with no child issues spawned and no manual label changes.

## Assumptions

- The branch `e2e-base-branch-20261025-071901` already exists on the remote, as stated in the issue.
- The marker line is wanted in the file exactly as written, with no surrounding text.
- A file that does not end in a newline gets one added before the marker line, so the marker occupies its own line.
- The work happens in the single repository this issue is filed in.

## Out of Scope *(optional)*

- Any change to the engine, its configuration, or other files.
- Adding or removing labels on this issue.
- Splitting the work into child issues.
- Merging the pull request on the engine's behalf (`fabrik:cruise` does not auto-merge).

## Source References *(optional)*

- handarbeit/fabrik#1046: regression for the `base:<branch>` pipeline contract (mentioned for context)
- handarbeit/fabrik#1047: issue↔PR linkage fix (mentioned for context)
- handarbeit/fabrik#1050: review-gate data feed fix (mentioned for context)
