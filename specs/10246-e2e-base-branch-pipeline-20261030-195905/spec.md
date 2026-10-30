# Feature Specification: E2E base-branch pipeline verification marker

**Feature Branch**: `fabrik/issue-10246`
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
<!-- base-branch-pipeline-20261030-195905 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261030-195905 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end run exercising the pipeline for an item that carries a `base:<branch>` label and so targets a non-default base branch. It provides regression coverage for handarbeit/fabrik#1046 and validates the fixes in #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (review-gate data feed). The code change is deliberately trivial, so that any failure observed comes from the pipeline's handling of the non-default base and not from the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

The pipeline carries the issue from Specify through to merge, with all work branched from, and the PR targeting, the `base:` branch `e2e-base-branch-20261030-195905`. The only content change is one marker line appended to `e2e/markers/base-branch-pipeline.md`.

**Why this priority**: This is the sole deliverable. The marker line is the evidence that the whole base-branch path worked.

**Independent Test**: After the PR merges, inspect `e2e/markers/base-branch-pipeline.md` on the branch `e2e-base-branch-20261030-195905`. Its last line must be exactly `<!-- base-branch-pipeline-20261030-195905 -->`.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist on the base branch, **When** the change is implemented, **Then** the file is created and contains the marker line as its last line.
2. **Given** the marker file already exists, **When** the change is implemented, **Then** the marker line is appended at the very end and existing content is unchanged.
3. **Given** the PR is opened, **When** its target is inspected, **Then** its base is `e2e-base-branch-20261030-195905`.

---

### User Story 2 - Pipeline contract for non-default bases is observable (Priority: P2)

The run lets an operator confirm the #1047 and #1050 behaviours for a `base:<branch>` item.

**Why this priority**: This is the purpose of the e2e run, but it is verified by observing the run, not by the content change.

**Independent Test**: Inspect the issue and PR after the run for the linkage and review-gate behaviour.

**Acceptance Scenarios**:

1. **Given** the PR targets a non-default base, **When** the pipeline runs, **Then** the engine still associates the PR with issue #10246 (per #1047).
2. **Given** the PR is under review, **When** the review gate evaluates it, **Then** it receives the PR and review data it needs (per #1050).

---

### Edge Cases

- The marker file lacks a trailing newline: the marker must still end up on its own line, as the last line.
- The marker file is created fresh: its content is only the marker line.
- The marker is already present from a prior attempt: it must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/base-branch-pipeline.md` MUST end with the line `<!-- base-branch-pipeline-20261030-195905 -->`.
- **FR-002**: If the file does not exist, it MUST be created with that line as its content.
- **FR-003**: The change MUST consist of exactly one file and one added line, with no other files or content modified.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The work MUST branch from, and the PR MUST target, `e2e-base-branch-20261030-195905`.
- **FR-006**: The issue's labels (including `fabrik:cruise` and `base:e2e-base-branch-20261030-195905`) MUST NOT be added or removed by this work.
- **FR-007**: The work MUST NOT be decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After merge, the last line of `e2e/markers/base-branch-pipeline.md` on `e2e-base-branch-20261030-195905` is the exact marker line, and it occurs once.
- **SC-002**: The PR diff shows exactly one file changed with one line added (apart from any engine-authored spec file).
- **SC-003**: The PR's base branch is `e2e-base-branch-20261030-195905` and it is linked to issue #10246.

## Assumptions

- The branch `e2e-base-branch-20261030-195905` exists on the remote, as stated in the issue.
- `e2e/markers/` may or may not already exist; creating the directory with the file is acceptable.
- The engine-authored `specs/10246-*/spec.md` projection is not counted as part of the "one file, one line" change.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Modifying labels, board configuration, or the engine.
- Decomposing the work into child issues.
- Changes to the default branch.

## Source References *(optional)*

- handarbeit/fabrik#1046 (regression being covered)
- handarbeit/fabrik#1047 (issue↔PR linkage fix)
- handarbeit/fabrik#1050 (review-gate data feed fix)
