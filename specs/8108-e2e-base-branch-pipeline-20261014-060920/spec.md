# Feature Specification: e2e base-branch pipeline (20261014-060920)

**Feature Branch**: `fabrik/issue-8108`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261014-060920 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261014-060920 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end exercise of the `base:<branch>` pipeline contract: an issue whose work targets a non-default base branch must travel the whole pipeline correctly. It is regression coverage for handarbeit/fabrik#1046, validating the fixes in #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (review-gate data feed). The code change is deliberately trivial, because the point is the pipeline behavior around it, not the content of the change. The target base branch is `e2e-base-branch-20261014-060920`, set by the `base:` label at filing.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator running the e2e suite files this issue with a `base:e2e-base-branch-20261014-060920` label. The pipeline produces a PR targeting that base branch that adds one marker line to `e2e/markers/base-branch-pipeline.md`.

**Why this priority**: This is the whole deliverable. Without it there is nothing for the pipeline to carry through its stages.

**Independent Test**: Inspect the resulting PR diff against `e2e-base-branch-20261014-060920`. It should contain exactly one added line in `e2e/markers/base-branch-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the change is implemented, **Then** the file ends with `<!-- base-branch-pipeline-20261014-060920 -->` as a new final line and no other line is altered.
2. **Given** the file does not exist on the base branch, **When** the change is implemented, **Then** the file is created containing the marker line.
3. **Given** the issue carries `base:e2e-base-branch-20261014-060920`, **When** the PR is opened, **Then** the PR targets that branch, not the repository's default branch.

---

### Edge Cases

- The marker file is missing: it must be created, not treated as an error.
- The marker file already ends without a trailing newline: the marker must still land on its own line.
- The marker line is already present (for example on a re-run): it must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261014-060920 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, the change MUST create it with the marker line as its content.
- **FR-003**: The change MUST touch only that one file, with no other file modified, added or removed.
- **FR-004**: The work MUST target the base branch `e2e-base-branch-20261014-060920`, and the resulting PR MUST be based on that branch.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: The issue's labels MUST NOT be added to or removed from as part of this work.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261014-060920` consists of exactly one file with exactly one added line, the marker.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261014-060920`, and the issue and the PR are linked to each other (regression check for #1047).
- **SC-003**: The review gate sees the PR's review data and the issue completes the pipeline stages without manual intervention (regression check for #1050).
- **SC-004**: No child issues are spawned.

## Assumptions

- The repository's configured default base is already `e2e-base-branch-20261014-060920` for this run, so the `base:` label and the default agree. The pipeline behavior under test is unaffected.
- The marker line text is exact, including the timestamp suffix `20261014-060920`.
- The `fabrik:cruise` label advances the issue through stages without human approval, and that is expected for this e2e run.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configuration or documentation.
- Any other file under `e2e/` or elsewhere.
- Multi-repo scenarios.
- Decomposition into sub-issues.
- Adding or removing labels.
