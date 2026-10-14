# Feature Specification: e2e base-branch pipeline (20261014-030832)

**Feature Branch**: `fabrik/issue-7984`
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
> <!-- base-branch-pipeline-20261014-030832 -->
> ```
>
> That is the entire change. One file, one line. Plan should NOT decompose.
>
> ## Scope
>
> Single repo only, targeting the non-default base branch e2e-base-branch-20261014-030832 (already set via
> the base:<branch> label applied at filing — do not add or remove labels).

## Background

This issue is a live end-to-end exercise of the `base:<branch>` pipeline contract: an issue labelled with a non-default base branch must flow through the pipeline, with its work item and pull request correctly linked and its review data correctly fed to the review gate. It is regression coverage for handarbeit/fabrik#1046, validating the fixes shipped under #1047 (issue↔PR linkage) and #1050 (review-gate data feed). The change itself is deliberately trivial, so that any failure observed is attributable to the pipeline rather than to the content of the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

An operator running the e2e base-branch check files this issue with the `base:e2e-base-branch-20261014-030832` label. The pipeline produces a pull request that adds one marker line to the marker file, targets that base branch, and is linked to this issue.

**Why this priority**: It is the only deliverable. The check passes only if the change is made and flows through the pipeline against the labelled base branch.

**Independent Test**: Inspect the resulting pull request: its base is `e2e-base-branch-20261014-030832`, it is linked to this issue, and its diff is exactly one added line in `e2e/markers/base-branch-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** the marker file already exists on the base branch, **When** the change is made, **Then** the line `<!-- base-branch-pipeline-20261014-030832 -->` is the last line of the file and no other line is changed.
2. **Given** the marker file does not exist on the base branch, **When** the change is made, **Then** the file is created containing the marker line.
3. **Given** the issue carries the `base:e2e-base-branch-20261014-030832` label, **When** the pull request is opened, **Then** it targets that branch and not the repository's other branches.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker still appears on its own line at the end.
- The marker line is already present (for example on a retry): it is not added a second time.
- Labels on the issue must not be added or removed as part of this work.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261014-030832 -->`, to the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created, containing that line.
- **FR-003**: The change MUST touch only that one file and add only that one line; no other file or line changes.
- **FR-004**: The work MUST target the base branch named by the issue's `base:` label, `e2e-base-branch-20261014-030832`.
- **FR-005**: The work MUST NOT add or remove any issue labels.
- **FR-006**: The work MUST NOT be decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The pull request diff against `e2e-base-branch-20261014-030832` contains exactly one file changed and exactly one line added.
- **SC-002**: The pull request's base branch is `e2e-base-branch-20261014-030832`.
- **SC-003**: The pull request is linked to this issue, so that the issue's pipeline discovers it.
- **SC-004**: The issue passes through the pipeline stages without being paused or blocked by base-branch handling.

## Assumptions

- The marker line is added once; a rerun that finds it already present is a no-op rather than a duplicate.
- The `e2e/markers/` directory may need to be created together with the file if neither exists.
- The labels set at filing (`fabrik:cruise`, `base:e2e-base-branch-20261014-030832`) are correct and stay as they are.

## Out of Scope *(optional)*

- Any change to the engine or its base-branch handling; this issue only exercises it.
- Any change beyond the single marker line, including edits to existing content in the marker file.
- Multi-repo scenarios.
- Decomposition into child issues.
