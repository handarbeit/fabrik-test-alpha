# Feature Specification: e2e base-branch pipeline marker line

**Feature Branch**: `fabrik/issue-12030`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261126-231722 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261126-231722 (already set via
the base:<branch> label applied at filing — do not add or remove labels).


## Labels

fabrik:cruise, base:e2e-base-branch-20261126-231722"

## Background

This issue is a live end-to-end exercise of the pipeline for items that carry a `base:<branch>` label, meaning work that forks from and targets a branch other than the repository's original default. It is regression coverage for handarbeit/fabrik#1046. It checks that the fixes in #1047 (the issue-to-PR linkage for a PR that targets a non-default base) and #1050 (the data the review gate reads for such a PR) hold end to end.

The code change is deliberately trivial: the point is to push a minimal change through every stage on the non-default base, not to deliver a feature. The marker line is unique to this run (it carries the run timestamp), so the resulting change can be identified unambiguously.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

An operator running the e2e suite files this issue with the `base:e2e-base-branch-20261126-231722` and `fabrik:cruise` labels. The pipeline carries it through all stages unattended. The result is a PR targeting `e2e-base-branch-20261126-231722` that appends the run's marker comment to `e2e/markers/base-branch-pipeline.md`.

**Why this priority**: This is the whole change and the only thing the e2e run needs to observe.

**Independent Test**: Inspect the PR diff against `e2e-base-branch-20261126-231722`. It should contain exactly one added line in `e2e/markers/base-branch-pipeline.md`, and that line should be the marker comment.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the change is implemented, **Then** the file ends with the line `<!-- base-branch-pipeline-20261126-231722 -->` and every earlier line is unchanged.
2. **Given** the file does not exist on the base branch, **When** the change is implemented, **Then** the file is created and its content is the single marker line.
3. **Given** the item carries `base:e2e-base-branch-20261126-231722`, **When** the PR is opened, **Then** the PR targets that branch and not any other.

---

### User Story 2 - Non-default base pipeline contract is exercised (Priority: P1)

The e2e run needs the pipeline's base-branch behaviours to be observable on this item so that a regression in #1047 or #1050 shows up as a failure of the run.

**Why this priority**: This is the reason the issue exists. The marker change is only the vehicle.

**Independent Test**: After the run, check that the issue and its PR are linked, and that the review gate sees the PR's review data and does not stall or mis-report.

**Acceptance Scenarios**:

1. **Given** the PR targets the non-default base, **When** the pipeline discovers the PR for the issue, **Then** the issue and the PR are associated for pipeline purposes (#1047).
2. **Given** the PR is open on the non-default base, **When** the review gate evaluates it, **Then** it reads the PR's review data correctly and reaches a verdict (#1050).
3. **Given** the item carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances without human approval.

---

### Edge Cases

- The marker file already ends without a trailing newline: the marker must still land on its own line, and the existing last line must be unchanged.
- The marker file already contains this exact marker line (a re-run): it must not be duplicated.
- The marker file exists on the base branch but not on the repository's other branches: the base branch's version is the one that is modified.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261126-231722 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, the change MUST create it with that single line as its content.
- **FR-003**: The change MUST NOT modify any other line in that file, and MUST NOT modify any other file.
- **FR-004**: The change MUST target the base branch `e2e-base-branch-20261126-231722`, as set by the `base:` label, and the resulting PR MUST target that branch.
- **FR-005**: The labels on the issue (`fabrik:cruise`, `base:e2e-base-branch-20261126-231722`) MUST NOT be added to or removed from by the work in this issue.
- **FR-006**: The work MUST be delivered as a single unit and MUST NOT be decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261126-231722` consists of exactly one file and exactly one added line, the marker comment (or, if the file was absent, a new file whose only content is that line).
- **SC-002**: The PR's base branch is `e2e-base-branch-20261126-231722`.
- **SC-003**: The issue and its PR are linked for pipeline purposes without manual intervention, and the item completes every stage without human input.
- **SC-004**: No child issues are created for this item.

## Assumptions

- The marker file's current content, if any, is plain text or Markdown, and appending an HTML comment line is harmless to whatever reads it.
- The `base:` label and the repository's default branch both name `e2e-base-branch-20261126-231722`, so no fallback to another branch is expected.
- A re-run that finds the marker already present should leave the file unchanged rather than add a duplicate line.

## Out of Scope *(optional)*

- Any change other than the one marker line.
- Changes to the engine, stage configs, or the e2e harness itself.
- Adding or removing labels on the issue.
- Decomposing the work into sub-issues.

## Source References *(optional)*

- handarbeit/fabrik#1046 — base-branch pipeline regression
- handarbeit/fabrik#1047 — issue-to-PR linkage fix for non-default base
- handarbeit/fabrik#1050 — review-gate data feed fix
