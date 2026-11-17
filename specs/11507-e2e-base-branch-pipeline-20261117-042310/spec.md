# Feature Specification: e2e base-branch pipeline marker (20261117-042310)

**Feature Branch**: `fabrik/issue-11507`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description:

`````
## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261117-042310 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261117-042310 (already set via
the base:<branch> label applied at filing — do not add or remove labels).


## Labels

fabrik:cruise, base:e2e-base-branch-20261117-042310
`````

## Background

This issue is a live end-to-end test fixture, not a product feature. It exercises the contract for issues that carry a `base:<branch>` label, meaning the work targets a non-default base branch. It is regression coverage for handarbeit/fabrik#1046, and it validates two fixes: #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (the data feed to the review gate for such PRs).

The code change is deliberately trivial: one line in one marker file. The point is to drive an item through the whole pipeline (Specify → Research → Plan → Implement → Review → Validate) on a non-default base and observe that the pipeline behaves correctly. The change itself carries no product value. The item carries `fabrik:cruise` and `base:e2e-base-branch-20261117-042310`, which were applied at filing and must not be added, removed or altered.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An e2e harness operator files this issue with a `base:` label and expects the pipeline to produce a pull request that targets the labelled base branch. The PR must add exactly one HTML comment line to the end of the marker file.

**Why this priority**: This is the entire deliverable. Without it there is nothing for the pipeline's base-branch handling to be verified against.

**Independent Test**: After the pipeline completes, inspect the PR diff. It touches only `e2e/markers/base-branch-pipeline.md` and adds only the line `<!-- base-branch-pipeline-20261117-042310 -->` at the end of the file.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is implemented, **Then** the file's existing content is unchanged and the new line is the last line of the file.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `base:e2e-base-branch-20261117-042310`, **When** the PR is created, **Then** its base branch is `e2e-base-branch-20261117-042310`.

---

### User Story 2 - Pipeline contract for non-default base holds (Priority: P1)

The operator expects the pipeline to treat the non-default-base item correctly end to end, as covered by the #1047 and #1050 fixes.

**Why this priority**: This is the reason the test exists. It is verified by observing the pipeline, not by the content of the change.

**Independent Test**: Observe the item through the pipeline and confirm the linkage and review-gate behaviour below.

**Acceptance Scenarios**:

1. **Given** the PR targets a non-default base, **When** the pipeline looks up the issue's PR, **Then** the PR is correctly linked to this issue (#1047).
2. **Given** the linked PR is in review, **When** the review gate evaluates it, **Then** it receives the correct PR and review data (#1050).
3. **Given** the item carries `fabrik:cruise`, **When** each stage completes, **Then** the item auto-advances through the stages without being auto-merged by the engine.

---

### Edge Cases

- The marker file exists but its last line has no trailing newline: the new line must still end up on its own line, not appended to the previous line.
- The marker file already contains the same marker line: the line is not duplicated, and the result still ends with exactly one such line.
- The pipeline must not decompose this issue into child issues; it is a single-file, single-line change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- base-branch-pipeline-20261117-042310 -->` at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, the change MUST create it, containing the marker line.
- **FR-003**: The change MUST touch no file other than `e2e/markers/base-branch-pipeline.md`, and MUST add no lines other than the marker line.
- **FR-004**: The work MUST target the base branch `e2e-base-branch-20261117-042310`, and any PR MUST use that branch as its base.
- **FR-005**: The change MUST NOT add, remove or alter any label on the issue, including `fabrik:cruise` and `base:e2e-base-branch-20261117-042310`.
- **FR-006**: The work MUST NOT be decomposed into child issues.
- **FR-007**: The PR MUST include `Closes #11507` in its body so issue↔PR linkage can be verified.

### Key Entities

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a test fixture whose only purpose is to carry per-run marker lines.
- **Marker line**: an HTML comment, `<!-- base-branch-pipeline-20261117-042310 -->`, unique to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one added line in exactly one file.
- **SC-002**: The PR's base branch equals `e2e-base-branch-20261117-042310`, not the repository's other branches.
- **SC-003**: The issue is linked to its PR (Closes #11507) and the review gate sees that PR's review data, with no manual repair.
- **SC-004**: The item's labels at the end of the run match those at filing, apart from engine-managed state labels.
- **SC-005**: No child issues are created.

## Assumptions

- The repository's default base branch is `e2e-base-branch-20261117-042310`, as stated in the stage context, so the `base:` label and the repository default coincide for this run.
- A single trailing newline after the marker line is acceptable and expected, as is normal for text files.
- The marker file's existing contents, if any, are unrelated run markers that must be preserved.

## Out of Scope *(optional)*

- Any change to engine behaviour, documentation, tests or code outside the marker file.
- Changing labels, board columns or the base branch.
- Cleanup of other run markers in `e2e/markers/`.
- Follow-up fixes for #1046, #1047 or #1050 themselves; this run only verifies them.

## Source References *(optional)*

- handarbeit/fabrik#1046 (base-branch pipeline regression), #1047 (issue↔PR linkage fix), #1050 (review-gate data feed fix)
