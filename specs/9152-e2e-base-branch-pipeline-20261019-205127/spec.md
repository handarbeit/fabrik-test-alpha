# Feature Specification: E2E base-branch pipeline marker (20261019-205127)

**Feature Branch**: `fabrik/issue-9152`
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
<!-- base-branch-pipeline-20261019-205127 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261019-205127 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end fixture. Its purpose is not the content of the change but the pipeline it travels through. It checks that an issue labelled `base:<branch>` (a non-default base branch) runs through every stage with correct behaviour, as regression coverage for handarbeit/fabrik#1046. It specifically validates two earlier fixes:

- **#1047**: the issue↔PR linkage for a PR that targets a non-default base branch.
- **#1050**: the data feed into the review gate for such a PR.

The code change is deliberately trivial (one marker line in one file). Any failure in the run is therefore attributable to the pipeline and not to the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

An operator running the base-branch e2e check wants the issue to produce a PR containing exactly one appended marker line, targeting the non-default base branch `e2e-base-branch-20261019-205127`. The PR must be linked to the issue and processed by the review gate.

**Why this priority**: This is the whole change. A PR with the right content against the right base is what the e2e check observes.

**Independent Test**: Inspect the resulting PR's diff and base branch. The diff is one appended line in `e2e/markers/base-branch-pipeline.md`, and the base is `e2e-base-branch-20261019-205127`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` already exists, **When** the change is made, **Then** the file ends with the line `<!-- base-branch-pipeline-20261019-205127 -->`, all prior content is unchanged, and no other file changes.
2. **Given** the file does not exist, **When** the change is made, **Then** the file is created containing the marker line as its final line, and no other file changes.
3. **Given** the issue carries the `base:e2e-base-branch-20261019-205127` label, **When** the PR is opened, **Then** it targets that branch and not the repository's `main`.

---

### User Story 2 - Pipeline contract holds for a non-default base (Priority: P2)

The operator wants the pipeline to behave correctly for a `base:<branch>` item end to end. This is the regression the e2e exists to catch.

**Why this priority**: It is the purpose of the fixture, but it is verified by observing engine behaviour, not by anything the change itself implements.

**Independent Test**: After the run, check the linkage and review-gate observations below.

**Acceptance Scenarios**:

1. **Given** a PR targeting the non-default base, **When** it is opened, **Then** the issue and PR are linked (#1047) so the engine can discover the PR and its comments.
2. **Given** the PR is ready for review, **When** the review gate evaluates it, **Then** it receives the PR's review data (#1050) and does not stall for lack of it.
3. **Given** the issue was filed with its labels, **When** the pipeline runs, **Then** the `base:` label and the other labels are neither added nor removed by the work done for this issue.

---

### Edge Cases

- The marker file exists but does not end with a trailing newline. The marker line must still be appended as its own line, never joined onto the previous line.
- The marker file already contains this exact marker line. It must not be duplicated (see Assumptions).
- The default branch of the repository is the same as the base branch. The base must not be altered for this run.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261019-205127 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, the change MUST create it, with the marker line as its content.
- **FR-003**: The change MUST NOT modify any file other than `e2e/markers/base-branch-pipeline.md`, and MUST NOT alter existing lines in that file.
- **FR-004**: The resulting PR MUST target the base branch `e2e-base-branch-20261019-205127` and MUST reference this issue so the issue↔PR linkage works.
- **FR-005**: The work MUST NOT add or remove any labels on the issue, including the `base:` label.
- **FR-006**: The work MUST NOT be decomposed into child issues. It is a single, one-file, one-line change in a single repository.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, an append-only record of base-branch e2e runs. Each run appends one uniquely stamped HTML comment line.
- **Base branch**: `e2e-base-branch-20261019-205127`, the non-default branch the issue's work forks from and the PR targets, selected by the `base:` label.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff against the base branch is exactly one added line, and it equals `<!-- base-branch-pipeline-20261019-205127 -->`.
- **SC-002**: The PR's base branch equals `e2e-base-branch-20261019-205127` and not the repository default.
- **SC-003**: The PR is discoverable from the issue (linked), and the review gate completes using the PR's review data, with no manual intervention.
- **SC-004**: The issue's label set shows no additions or removals attributable to this work.

## Assumptions

- Appending the marker line is idempotent: if the exact line is already the last line, nothing further is appended.
- The marker line is added as its own line, so a missing trailing newline is supplied first.
- The single-repository scope means the repository that hosts the issue, and no other.
- The issue is run with `fabrik:cruise`, so stages advance without waiting for human approval of this spec. The spec is complete, so no approval is needed.

## Out of Scope *(optional)*

- Any change to engine, pipeline, or review-gate code. This issue only exercises existing behaviour.
- Fixing #1046, #1047 or #1050. They are already delivered, and this is regression coverage for them.
- Splitting the work into multiple issues or child issues.
- Adding or removing issue labels, or changing the base branch.
- Content in the marker file beyond the single stamped line.

## Source References *(optional)*

- handarbeit/fabrik#1046 (base-branch pipeline regression)
- handarbeit/fabrik#1047 (issue↔PR linkage for non-default base)
- handarbeit/fabrik#1050 (review-gate data feed)
