# Feature Specification: e2e base-branch pipeline (20261026-132100)

**Feature Branch**: `fabrik/issue-9977`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261026-132100 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261026-132100 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end check of the `base:<branch>` contract, where an issue carries a label that moves its worktree base, PR target and merge destination off the repository default. It is regression coverage for handarbeit/fabrik#1046. It checks that the fixes from #1047 (issue-to-PR linkage for a PR targeting a non-default base) and #1050 (the review gate receiving its data for such a PR) hold when the item runs through the whole pipeline.

The code change is deliberately trivial: one marker line in one file. The subject under test is the pipeline's handling of the non-default base, not the content of the change. A minimal, deterministic diff keeps any failure attributable to base-branch handling rather than to the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

An operator runs the e2e suite. The issue carrying `base:e2e-base-branch-20261026-132100` passes through every pipeline stage. The result is a PR, targeting that base branch rather than the repository default, whose entire diff is one appended marker line in `e2e/markers/base-branch-pipeline.md`.

**Why this priority**: This is the whole deliverable. Without it nothing about the base-branch contract can be verified.

**Independent Test**: Inspect the resulting PR. Its base is `e2e-base-branch-20261026-132100`, and its diff touches only `e2e/markers/base-branch-pipeline.md`, adding exactly the line `<!-- base-branch-pipeline-20261026-132100 -->` at the end of the file.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist on the base branch, **When** the change is implemented, **Then** the file is created containing the marker line as its last (and only) line.
2. **Given** the marker file already exists, **When** the change is implemented, **Then** the marker line is appended after all existing content and no existing line is modified or removed.
3. **Given** the issue has the `base:e2e-base-branch-20261026-132100` label, **When** the PR is created, **Then** the PR targets that branch, not the repository default.

---

### User Story 2 - Issue and PR linkage and review-gate data feed work for a non-default base (Priority: P2)

The pipeline discovers the PR for the issue and feeds the review gate even though GitHub creates no Development-panel link for a PR whose base is not the default branch.

**Why this priority**: This is the regression the test exists to guard (#1047 and #1050). It is secondary to the change landing, because it is observed through the pipeline's behaviour rather than through the diff.

**Independent Test**: Watch the item progress past Implement and the review gate without stalling on missing linkage or missing review data, and see it reach completion.

**Acceptance Scenarios**:

1. **Given** the PR targets a non-default base, **When** the engine looks up the issue's linked PR, **Then** the PR is found through its `Closes #9977` body reference.
2. **Given** the PR is open and ready, **When** the review gate evaluates it, **Then** it receives the PR's reviews and review requests and reaches a decision rather than stalling or erroring.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker line must still start on its own line, so it does not merge with the previous last line.
- The marker file exists and already contains this exact marker line: it is not duplicated, and the change is a no-op for that line.
- The change must not add labels to or remove labels from the issue, since the base label was set at filing.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261026-132100 -->`, at the end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created, containing that line.
- **FR-003**: The change MUST NOT modify or remove any existing content of the marker file.
- **FR-004**: The change MUST touch only `e2e/markers/base-branch-pipeline.md`: one file, one added line.
- **FR-005**: The work MUST be treated as a single unit and MUST NOT be decomposed into child issues.
- **FR-006**: The work MUST target the single repository named by the issue and the base branch `e2e-base-branch-20261026-132100`, as set by the existing `base:` label.
- **FR-007**: The pipeline MUST NOT add or remove labels on the issue as part of this work beyond those the engine itself manages for stage state.

### Key Entities

- **Marker file** (`e2e/markers/base-branch-pipeline.md`): a text file whose only purpose is to carry one-line markers that give each e2e run a unique, trivially verifiable diff.
- **Base branch** (`e2e-base-branch-20261026-132100`): the non-default branch that the issue's worktree, PR and merge target.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The resulting PR's diff against its base contains exactly one changed file with exactly one added line, and that line equals `<!-- base-branch-pipeline-20261026-132100 -->`.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261026-132100` and not the repository default.
- **SC-003**: The item progresses from Specify through the last stage without any manual intervention to repair issue-to-PR linkage or review-gate data.
- **SC-004**: No child issues are created for this work.

## Assumptions

- The base branch `e2e-base-branch-20261026-132100` already exists on the remote, so the engine does not fall back to the default.
- The marker file may or may not exist on that branch; both cases are covered by FR-001 and FR-002.
- The `fabrik:cruise` label means the item auto-advances through the stages without waiting for approval, and it does not auto-merge.

## Out of Scope *(optional)*

- Any change other than the single marker line, including edits to the engine, tests or documentation.
- Adding, removing or changing labels on the issue.
- Multi-repo scenarios.
- Decomposing the work into child issues.
- Changes to the pipeline's base-branch handling itself, which this issue only verifies.
