# Feature Specification: E2E base-branch pipeline marker (20261025-142641)

**Feature Branch**: `fabrik/issue-9641`
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
<!-- base-branch-pipeline-20261025-142641 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261025-142641 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is an end-to-end regression exercise for the `base:<branch>` (non-default base branch) pipeline contract, originally reported in handarbeit/fabrik#1046. Two fixes are being validated by driving a real item through the full pipeline against a non-default base:

- #1047: the issue↔PR linkage for a PR that targets a non-default base branch.
- #1050: the data feed into the review gate for such a PR.

The code change is deliberately trivial, so that any failure points at the pipeline's base-branch handling and not at the content of the change. The issue carries the `base:e2e-base-branch-20261025-142641` label, set at filing, and the `fabrik:cruise` label. Neither label is to be added or removed.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

A single marker line is appended to a tracked marker file and delivered through the normal pipeline (Specify → Research → Plan → Implement → Review → Validate). The pull request targets the non-default base branch `e2e-base-branch-20261025-142641`, not the repository default.

**Why this priority**: This is the whole change. It is the artifact that exercises the base-branch pipeline contract.

**Independent Test**: After the PR merges, inspect `e2e/markers/base-branch-pipeline.md` on the branch `e2e-base-branch-20261025-142641`. Its last line must be exactly `<!-- base-branch-pipeline-20261025-142641 -->`, and nothing else in the repository may differ because of this issue.

**Acceptance Scenarios**:

1. **Given** the marker file already exists on the base branch, **When** the change is implemented, **Then** the marker line is appended as the final line and all existing content is unchanged.
2. **Given** the marker file does not exist on the base branch, **When** the change is implemented, **Then** the file is created containing the marker line.
3. **Given** the issue has the `base:e2e-base-branch-20261025-142641` label, **When** the PR is created, **Then** its base is `e2e-base-branch-20261025-142641` and it links to this issue.

---

### Edge Cases

- The marker file exists but its last line has no trailing newline: the marker must still land on its own line, not be joined to the previous line.
- The marker line is already present (for example, after a retry or re-run): it must not be appended a second time.
- The file does not exist yet: it is created with only the marker line.
- Labels must not be modified by the work. This includes the `base:` label and `fabrik:cruise`.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/base-branch-pipeline.md` ends with the line `<!-- base-branch-pipeline-20261025-142641 -->`.
- **FR-002**: If the file already exists, its prior content is preserved byte-for-byte and the marker is appended after it, on its own line.
- **FR-003**: If the file does not exist, it is created containing the marker line.
- **FR-004**: The change touches only this one file and adds only this one line.
- **FR-005**: The pull request targets the base branch `e2e-base-branch-20261025-142641` and is linked to this issue.
- **FR-006**: The work is not decomposed into child issues.
- **FR-007**: The issue's labels, including `base:e2e-base-branch-20261025-142641`, are not added or removed as part of this work.

### Key Entities

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, an append-only record of base-branch pipeline e2e runs. Each run adds one HTML comment line, which is invisible when rendered.
- **Base branch**: `e2e-base-branch-20261025-142641`, the non-default branch the PR targets and the one the pipeline must treat as the base.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the PR merges, the last line of the marker file on the non-default base branch equals the marker string exactly, and the diff of the change is exactly one added line (or one new one-line file).
- **SC-002**: The PR for this issue targets `e2e-base-branch-20261025-142641` and is recognised as linked to the issue, with no manual intervention.
- **SC-003**: The item moves through every pipeline stage to completion with no pause or human input caused by the base-branch handling.

## Assumptions

- The marker directory `e2e/markers/` may or may not exist on the base branch. If it doesn't, creating it along with the file is acceptable.
- The marker line should end with a newline, as is normal for a text file.
- The `fabrik:cruise` label advances the item automatically through the stages, so no human approval of this spec is needed.
- A rerun that finds the marker line already present is a no-op and not an error.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configuration or documentation.
- Decomposing the work into child issues.
- Adding or removing labels on the issue.
- Changing the default branch or creating other branches.
- Multi-repo scenarios.

## Source References *(optional)*

- handarbeit/fabrik#1046: the base-branch pipeline contract regression this run covers.
- handarbeit/fabrik#1047: issue↔PR linkage fix being validated.
- handarbeit/fabrik#1050: review-gate data feed fix being validated.
