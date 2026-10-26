# Feature Specification: e2e base-branch pipeline marker (20261026-080759)

**Feature Branch**: `fabrik/issue-9799`
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
<!-- base-branch-pipeline-20261026-080759 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261026-080759 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end fixture, not a product feature. It exercises the `base:<branch>` pipeline contract, in which an issue carrying a `base:<branch>` label is worked on a branch forked from that base, and its pull request targets that base instead of the repository default. It is regression coverage for handarbeit/fabrik#1046. It validates two fixes: #1047 (the issue-to-PR linkage for non-default-base PRs) and #1050 (the data feed to the review gate).

The change itself is deliberately trivial, so that any failure points at the pipeline and not at the content of the change. The issue is labelled `fabrik:cruise` and `base:e2e-base-branch-20261026-080759`. Those labels are part of the test setup and must be left as they are.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

The pipeline takes this issue through every stage using the non-default base branch `e2e-base-branch-20261026-080759`. The result is one pull request, targeting that base, that adds the marker line to `e2e/markers/base-branch-pipeline.md`.

**Why this priority**: This is the whole purpose of the issue. A correct one-line change on the correct base proves the contract.

**Independent Test**: Inspect the resulting PR. Its base is `e2e-base-branch-20261026-080759` and its diff contains exactly one added line in exactly one file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, **When** the change is made, **Then** the file is created and its last line is `<!-- base-branch-pipeline-20261026-080759 -->`.
2. **Given** the file already exists on the base branch, **When** the change is made, **Then** the marker line is appended at the very end and existing content is unchanged.
3. **Given** the issue carries the label `base:e2e-base-branch-20261026-080759`, **When** the PR is created, **Then** the PR targets that branch and not the repository default branch.

---

### User Story 2 - Issue and PR linkage and review-gate data on a non-default base (Priority: P2)

The pipeline associates the issue with its PR, and gives the review gate the data it needs, even though the PR targets a non-default base. This is the behaviour fixed by #1047 and #1050.

**Why this priority**: These are the regressions the fixture exists to catch. They only matter once User Story 1 produces a PR.

**Independent Test**: Check that the PR is discoverable from the issue (the issue's PR comments and review feedback are seen by the pipeline) and that the review gate evaluates the PR instead of stalling for lack of data.

**Acceptance Scenarios**:

1. **Given** the PR body carries `Closes #9799`, **When** the pipeline looks up the issue's PR, **Then** it finds this PR.
2. **Given** the PR is open against the non-default base, **When** the review gate is evaluated, **Then** it receives review data for this PR and the pipeline advances.

---

### Edge Cases

- The marker file does not exist: it is created first, then the line is appended.
- The existing file does not end in a newline: the marker is still placed on its own line.
- The marker is already present from an earlier run: it is not added twice. The timestamped marker is unique to this run, so duplication indicates a re-run.
- Plan or Implement tries to split the work into child issues: this must not happen, because the change is one line in one file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change adds exactly one line, `<!-- base-branch-pipeline-20261026-080759 -->`, at the end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, it is created, and the marker line is its content.
- **FR-003**: The marker line is on its own line, and any existing content of the file is preserved unchanged.
- **FR-004**: The change touches exactly one file. No other files are modified, added or removed.
- **FR-005**: The resulting PR targets `e2e-base-branch-20261026-080759` and its body links this issue with `Closes #9799`.
- **FR-006**: The work is not decomposed. No child issues are spawned.
- **FR-007**: The labels on the issue (`fabrik:cruise`, `base:e2e-base-branch-20261026-080759`) are not added to or removed from by the work.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The final PR diff contains exactly 1 changed file and 1 added line (or, if the file was created, a file containing only that line).
- **SC-002**: The PR base branch is `e2e-base-branch-20261026-080759` and the PR is linked to issue #9799.
- **SC-003**: The pipeline moves the issue from Specify through to completion without manual intervention and without pausing for missing linkage or review-gate data.
- **SC-004**: Zero child issues are created for this issue.

## Assumptions

- The base branch `e2e-base-branch-20261026-080759` exists on the remote, as the issue states.
- The marker file does not currently exist, so the file is created. If it does exist on the base, the line is appended per FR-003.
- Nothing else in the repository depends on the content of this marker file.
- Because `fabrik:cruise` is set, stages advance automatically, but the PR is not auto-merged.

## Out of Scope *(optional)*

- Any change to the engine or to the pipeline's handling of `base:<branch>`. This issue only verifies it.
- Multi-repo scenarios.
- Changes to labels on the issue.
- Decomposing the work into child issues.

## Source References *(optional)*

- handarbeit/fabrik#1046: regression being covered (mentioned for context)
- handarbeit/fabrik#1047: issue and PR linkage fix for non-default base
- handarbeit/fabrik#1050: review-gate data feed fix
