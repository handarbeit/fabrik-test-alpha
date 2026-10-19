# Feature Specification: e2e base-branch pipeline (20261019-140116)

**Feature Branch**: `fabrik/issue-9030`
**Created**: 2026-10-19
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
> `<!-- base-branch-pipeline-20261019-140116 -->`
>
> That is the entire change. One file, one line. Plan should NOT decompose.
>
> ## Scope
>
> Single repo only, targeting the non-default base branch e2e-base-branch-20261019-140116 (already set via
> the base:<branch> label applied at filing — do not add or remove labels).

## Background

This issue is an end-to-end regression probe for the `base:<branch>` pipeline contract, which lets an issue target a non-default base branch. It guards against a recurrence of handarbeit/fabrik#1046 and exercises two fixes: #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (the data feed to the review gate). The code change is deliberately trivial, so the only thing under test is the pipeline's handling of the base branch (branching, PR targeting, linkage, review gating, landing), not the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator files this issue with a `base:e2e-base-branch-20261019-140116` label. The pipeline runs it through every stage. It should produce a one-line addition to a marker file, delivered through a PR that targets the labelled base branch.

**Why this priority**: This is the entire purpose of the issue. It shows that the pipeline works end to end for a non-default base.

**Independent Test**: After the pipeline finishes, the marker file on the base branch ends with the expected comment line. The PR that delivered it targeted the labelled base branch.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists, **When** the change is made, **Then** its final line is exactly `<!-- base-branch-pipeline-20261019-140116 -->` and all earlier content is unchanged.
2. **Given** `e2e/markers/base-branch-pipeline.md` does not exist, **When** the change is made, **Then** the file is created and contains the marker line as its last line.
3. **Given** the issue carries the label `base:e2e-base-branch-20261019-140116`, **When** the pipeline creates the PR, **Then** the PR's base is `e2e-base-branch-20261019-140116`, not the repository default.

---

### User Story 2 - Pipeline contract holds for the non-default base (Priority: P2)

The pipeline's linkage and review-gate behaviour for a non-default base is observed. This is the regression coverage for #1047 and #1050.

**Why this priority**: These behaviours are the reason for the exercise, but they are observed outcomes of the pipeline. They are not additional deliverables.

**Independent Test**: Inspect the issue and PR after the run to confirm the PR is linked to the issue and the review gate saw the PR's review data.

**Acceptance Scenarios**:

1. **Given** a PR targeting the non-default base, **When** the pipeline discovers it, **Then** the PR is associated with this issue.
2. **Given** the PR is in the review gate, **When** the gate evaluates it, **Then** it receives the PR's review data and does not stall because of the non-default base.

---

### Edge Cases

- The marker file already exists with or without a trailing newline: the new line must start on its own line, with no existing content altered.
- The marker line already exists in the file: it must not be duplicated.
- The repository's default base branch in this environment is the same as the labelled base branch, so the change must still target the labelled base explicitly.
- The labels on the issue must not be added to or removed by the work.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261019-140116 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created, with the marker line as its content.
- **FR-003**: The change MUST touch only that one file and add only that one line; existing content is unchanged.
- **FR-004**: The work MUST be delivered as a single unit and MUST NOT be decomposed into child issues.
- **FR-005**: The work MUST target the branch `e2e-base-branch-20261019-140116` as its base, as set by the `base:` label.
- **FR-006**: The work MUST NOT add or remove issue labels.
- **FR-007**: The work is limited to a single repository.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a tracked file whose last line records the run's unique marker.
- **Base branch**: `e2e-base-branch-20261019-140116`, the non-default branch selected by the `base:<branch>` label.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the last line of `e2e/markers/base-branch-pipeline.md` on the base branch is exactly `<!-- base-branch-pipeline-20261019-140116 -->`, and the change's diff is exactly one added line (or the file's creation with only that line).
- **SC-002**: The PR that carried the change targeted `e2e-base-branch-20261019-140116` and is linked to this issue.
- **SC-003**: The issue completed every pipeline stage without a manual intervention caused by the non-default base.
- **SC-004**: The issue's labels at the end of the run match those at filing, apart from the pipeline's own state labels.

## Assumptions

- The `base:e2e-base-branch-20261019-140116` branch exists on the remote.
- Appending the marker line to an already-existing file does not need a separate blank line before it.
- The marker line is a unique per-run string, so the file may already hold earlier markers from other runs.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Fixing or changing the engine behaviour of #1046, #1047 or #1050.
- Multi-repo scenarios.
- Adding or removing issue labels.

## Source References *(optional)*

- handarbeit/fabrik#1046 (regression covered), #1047 (issue↔PR linkage fix), #1050 (review-gate data feed fix)
