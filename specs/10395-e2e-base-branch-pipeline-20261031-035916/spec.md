# Feature Specification: e2e base-branch pipeline marker (20261031-035916)

**Feature Branch**: `fabrik/issue-10395`
**Created**: 2026-10-31
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261031-035916 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261031-035916 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end check of the `base:<branch>` pipeline contract: an item that targets a non-default base branch must travel the whole pipeline (Specify through Validate) with its branch, pull request and review handling all resolved against that base. It is regression coverage for handarbeit/fabrik#1046, validating the fixes in #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (the review-gate data feed). The change itself is deliberately trivial so that the pipeline behavior, not the content, is what is being exercised.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

The pipeline produces a pull request, targeting base branch `e2e-base-branch-20261031-035916`, whose only change is one appended marker line in `e2e/markers/base-branch-pipeline.md`.

**Why this priority**: This is the entire deliverable and the observable artifact the e2e run verifies.

**Independent Test**: Inspect the PR diff against `e2e-base-branch-20261031-035916`: exactly one file changed, with exactly one added line equal to the marker.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` already exists on the base branch, **When** the change is implemented, **Then** the file's final line is `<!-- base-branch-pipeline-20261031-035916 -->` and all prior content is unchanged.
2. **Given** the file does not exist on the base branch, **When** the change is implemented, **Then** the file is created containing the marker line.
3. **Given** the change is complete, **When** the PR is opened, **Then** its base is `e2e-base-branch-20261031-035916` and its diff contains only that one added line in that one file.

---

### User Story 2 - Pipeline contract holds for a non-default base (Priority: P2)

The item progresses through the pipeline with the `base:e2e-base-branch-20261031-035916` label unchanged, exercising the issue↔PR linkage and review-gate data feed for a non-default base.

**Why this priority**: This is the purpose of the e2e run (regression coverage for #1046/#1047/#1050), but it is observed rather than implemented by this change.

**Independent Test**: After the run, the item's labels still include `base:e2e-base-branch-20261031-035916` and `fabrik:cruise` with none added or removed by hand, and the PR is discoverable from the issue and reaches the review gate.

**Acceptance Scenarios**:

1. **Given** the issue carries `base:e2e-base-branch-20261031-035916`, **When** the pipeline runs, **Then** the worktree branches from and the PR targets that base rather than the repository default.
2. **Given** the PR targets the non-default base, **When** the review gate evaluates it, **Then** it receives the PR's review data and the issue↔PR linkage is resolved.

---

### Edge Cases

- The marker file already exists with a trailing newline or without one: the marker must still end up on its own final line, with earlier content untouched.
- The marker line is already present (e.g. on a re-run): it must not be duplicated.
- The file's directory `e2e/markers/` does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- base-branch-pipeline-20261031-035916 -->` at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, the change MUST create it, containing the marker line.
- **FR-003**: The change MUST modify no other file and add no other line.
- **FR-004**: The marker line MUST NOT be duplicated if it is already present.
- **FR-005**: The work MUST target the base branch `e2e-base-branch-20261031-035916`, and the `base:` and `fabrik:cruise` labels MUST NOT be added, removed or changed.
- **FR-006**: The work MUST NOT be decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261031-035916` shows exactly 1 file changed and exactly 1 line added.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261031-035916`, and the issue is linked to the PR.
- **SC-003**: The item's labels `base:e2e-base-branch-20261031-035916` and `fabrik:cruise` are unchanged at the end of the run.
- **SC-004**: No child issues are spawned.

## Assumptions

- The non-default base branch `e2e-base-branch-20261031-035916` exists on the remote.
- The marker file's current content, if any, needs no change beyond the appended line.
- If the file has no trailing newline, a newline is inserted before the marker so the marker sits on its own line.

## Out of Scope *(optional)*

- Any change to engine code, tests, docs or other files.
- Adding, removing or altering labels.
- Decomposition into sub-issues.
- Changes to the default branch.
