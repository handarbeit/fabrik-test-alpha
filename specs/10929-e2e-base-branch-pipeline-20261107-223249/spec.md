# Feature Specification: e2e base-branch pipeline marker (20261107-223249)

**Feature Branch**: `fabrik/issue-10929`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261107-223249 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261107-223249 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end test fixture, not a product feature. It exercises the pipeline for an issue carrying a `base:<branch>` label, so that work forks from, and its PR targets, a branch other than the repository default. It is regression coverage for handarbeit/fabrik#1046 and validates the fixes in #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (review-gate data feed). The code change is deliberately trivial so that the pipeline mechanics are what is being tested. The issue carries `fabrik:cruise` and `base:e2e-base-branch-20261107-223249`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator running the base-branch e2e check files this issue. The pipeline appends one marker line to a tracked file, opens a PR against the non-default base branch, and carries the issue through to completion.

**Why this priority**: This is the whole purpose of the issue. Without the marker line on the right base, the e2e check has nothing to assert on.

**Independent Test**: After the pipeline finishes, `e2e/markers/base-branch-pipeline.md` on the base branch `e2e-base-branch-20261107-223249` ends with the line `<!-- base-branch-pipeline-20261107-223249 -->`. The PR targets that branch, not the repository default.

**Acceptance Scenarios**:

1. **Given** the file `e2e/markers/base-branch-pipeline.md` already exists, **When** the change is made, **Then** exactly one new line `<!-- base-branch-pipeline-20261107-223249 -->` is added at the very end and no existing content is altered.
2. **Given** the file does not exist, **When** the change is made, **Then** the file is created containing the marker line.
3. **Given** the issue carries the `base:e2e-base-branch-20261107-223249` label, **When** a PR is opened, **Then** the PR's base is `e2e-base-branch-20261107-223249` and it is linked to this issue.

---

### Edge Cases

- The file exists but its last line has no trailing newline: the marker must still land on its own line, not appended to the previous line.
- The marker line is already present from an earlier run: it must not be added a second time.
- The base branch is missing on the remote: this is existing engine behaviour (fall back to the default and comment) and is not part of this change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261107-223249 -->`, at the end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created containing that line.
- **FR-003**: No other file may be changed and no existing line in the marker file may be modified.
- **FR-004**: The work MUST target the base branch `e2e-base-branch-20261107-223249`, as set by the existing `base:` label. Labels MUST NOT be added or removed as part of the work.
- **FR-005**: The work MUST be delivered as a single change in a single repository and MUST NOT be decomposed into child issues.

### Key Entities

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, an append-only file of HTML comment lines, one per e2e run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The diff of the resulting PR against its base is exactly one added line in exactly one file.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261107-223249`, and the PR is linked to this issue.
- **SC-003**: No child issues are created for this issue.

## Assumptions

- The marker file may or may not already exist on the base branch; both cases are handled by FR-001 and FR-002.
- The existing `base:` label and `fabrik:cruise` are correct and need no change.
- The marker file is an e2e fixture with no consumers other than the e2e check.

## Out of Scope *(optional)*

- Any change to engine behaviour, documentation, or tests.
- Adding, removing, or changing labels.
- Multi-repo work or decomposition into child issues.
- Fixing a missing remote base branch (existing engine fallback applies).
