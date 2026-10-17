# Feature Specification: E2E base-branch pipeline marker (20261017-150628)

**Feature Branch**: `fabrik/issue-8565`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261017-150628 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261017-150628 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end smoke test of the `base:<branch>` contract, in which an issue is worked on a non-default base branch rather than the repository default. It is regression coverage for handarbeit/fabrik#1046, and it validates the fixes in #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (the review-gate data feed). The code change is intentionally trivial, so that any failure observed comes from the pipeline and not from the change. The issue carries the `fabrik:cruise` and `base:e2e-base-branch-20261017-150628` labels, which were applied at filing and must not be altered.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator running the e2e suite wants one single-line change to travel through every pipeline stage on a non-default base branch, so that issue↔PR linkage and review-gate behaviour on that base can be verified.

**Why this priority**: This is the only deliverable. It exercises the whole pipeline contract.

**Independent Test**: After the pipeline finishes, the base branch's `e2e/markers/base-branch-pipeline.md` ends with the marker line, and the PR that delivered it targeted the base branch.

**Acceptance Scenarios**:

1. **Given** the file `e2e/markers/base-branch-pipeline.md` exists, **When** the change is implemented, **Then** the file's final line is `<!-- base-branch-pipeline-20261017-150628 -->` and every pre-existing line is unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `base:e2e-base-branch-20261017-150628`, **When** the PR is opened, **Then** it targets that branch, not the repository default, and it is linked to this issue.

---

### Edge Cases

- The marker file is missing: it is created with only the marker line.
- The existing file does not end with a trailing newline: the marker still lands on its own line, not appended to the previous line.
- The marker line is already present (a re-run): it is not duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The string `<!-- base-branch-pipeline-20261017-150628 -->` appears as the last line of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created, and it contains the marker line.
- **FR-003**: The marker line appears exactly once in the file, and all existing content of the file is preserved.
- **FR-004**: The delivered PR changes only `e2e/markers/base-branch-pipeline.md`, apart from the spec artifact the engine itself commits under `specs/`.
- **FR-005**: The PR targets the `e2e-base-branch-20261017-150628` base branch and is linked to this issue.
- **FR-006**: The issue's labels are not added or removed by this work.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The marker file on the base branch ends with the marker line, which appears exactly once.
- **SC-002**: The PR's diff contains one file other than the engine's spec artifact, with one added line (or the file creation plus that line).
- **SC-003**: The PR's base branch is `e2e-base-branch-20261017-150628`, and the issue↔PR link is discoverable by the pipeline.
- **SC-004**: The pipeline completes through the review gate without manual intervention.

## Assumptions

- The marker file may or may not already exist; the spec handles both cases.
- The work needs no decomposition into child issues, as the issue states.
- Because `fabrik:cruise` is set, the stages auto-advance and the PR is not auto-merged by the engine.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Adding or removing labels.
- Changes to the engine, its docs or its tests.
- Multi-repo scenarios.
- Decomposing the work into sub-issues.
