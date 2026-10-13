# Feature Specification: e2e base-branch pipeline (20261013-210827)

**Feature Branch**: `fabrik/issue-7816`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261013-210827 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261013-210827 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This is a live end-to-end run that exercises the `base:<branch>` pipeline contract: an issue labelled with a non-default base branch must flow through every stage, with its PR targeting that branch. It is regression coverage for handarbeit/fabrik#1046 and validates the fixes in #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (the data feed the review gate depends on). The change itself is deliberately trivial, so any failure observed comes from the pipeline and not from the work.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

A marker line is appended to `e2e/markers/base-branch-pipeline.md`, and the resulting PR targets the `base:` branch and is linked to this issue.

**Why this priority**: This is the whole change, and its delivery through the pipeline is what the run verifies.

**Independent Test**: After the pipeline completes, the file on branch `e2e-base-branch-20261013-210827` ends with the marker line, and the PR's base is that branch.

**Acceptance Scenarios**:

1. **Given** the file `e2e/markers/base-branch-pipeline.md` does not exist, **When** the change is implemented, **Then** the file exists and contains the marker line `<!-- base-branch-pipeline-20261013-210827 -->`.
2. **Given** the file already exists, **When** the change is implemented, **Then** the marker line is the last line of the file and the prior content is unchanged.
3. **Given** the issue carries `base:e2e-base-branch-20261013-210827`, **When** the PR is created, **Then** it targets that branch rather than the repository's default, and is linked to this issue.

---

### Edge Cases

- The file exists without a trailing newline: the marker must still occupy its own line.
- The marker line is already present (for example on a re-run): it is not duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- base-branch-pipeline-20261013-210827 -->` is present exactly once, as the final line of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created containing the marker line; if it exists, its existing content is preserved.
- **FR-003**: The change touches only that one file, apart from the spec artifact the engine itself commits.
- **FR-004**: The work is delivered on a PR that targets branch `e2e-base-branch-20261013-210827`.
- **FR-005**: The issue's labels are not added to or removed from by the work, beyond the pipeline's own state labels.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged result on `e2e-base-branch-20261013-210827` differs from before by exactly one appended line in one file, plus the spec artifact.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261013-210827` and the PR is linked to this issue.
- **SC-003**: The item moves through every pipeline stage without a manual intervention caused by the non-default base.

## Assumptions

- The branch `e2e-base-branch-20261013-210827` already exists on the remote, as stated in the issue.
- The file does not currently exist on the base branch (it is absent in the working tree), so it will be created.
- The issue is a single unit of work and is not decomposed into child issues, as the issue directs.

## Out of Scope *(optional)*

- Any change to the engine, tests or docs.
- Any other file under `e2e/markers/`.
- Decomposition into sub-issues.
