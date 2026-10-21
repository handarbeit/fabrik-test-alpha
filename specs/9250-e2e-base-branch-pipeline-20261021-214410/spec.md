# Feature Specification: e2e base-branch pipeline (20261021-214410)

**Feature Branch**: `fabrik/issue-9250`
**Created**: 2026-10-21
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261021-214410 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: single repo only, targeting the non-default base branch e2e-base-branch-20261021-214410 (already set via the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is an end-to-end test fixture, not a product feature. It exercises the pipeline for an issue carrying a `base:<branch>` label, so that its work is forked from, rebased onto and targeted at a non-default base branch. It is regression coverage for handarbeit/fabrik#1046 and validates two fixes:

- **#1047**: the issue-to-PR linkage for PRs that target a non-default base branch.
- **#1050**: the data feed that the review gate relies on for such PRs.

The code change is deliberately minimal (one marker line in one file) so that the pipeline mechanics are what is under test, not the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

A single marker line is added to the end of `e2e/markers/base-branch-pipeline.md` through the normal pipeline. The resulting PR targets the `e2e-base-branch-20261021-214410` base branch, not the repository's original default branch.

**Why this priority**: This is the entire change, and the only observable artifact the pipeline test checks.

**Independent Test**: After the change lands, the last line of `e2e/markers/base-branch-pipeline.md` on the target base branch is exactly `<!-- base-branch-pipeline-20261021-214410 -->`, and the file has no other additions.

**Acceptance Scenarios**:

1. **Given** the file `e2e/markers/base-branch-pipeline.md` already exists on the base branch, **When** the change is implemented, **Then** exactly one line, `<!-- base-branch-pipeline-20261021-214410 -->`, is appended at the very end and existing content is unchanged.
2. **Given** the file does not exist on the base branch, **When** the change is implemented, **Then** the file is created containing the marker line.
3. **Given** the issue carries the `base:e2e-base-branch-20261021-214410` label, **When** the PR is created, **Then** the PR targets that branch.

---

### Edge Cases

- The file already ends without a trailing newline: the marker must still start on its own line, so the file must not end up with the marker joined to the previous line.
- The marker line is already present at the end of the file (for example on a re-run): it must not be duplicated.
- The labels on the issue (`fabrik:cruise`, `base:e2e-base-branch-20261021-214410`) must remain untouched by the work.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- base-branch-pipeline-20261021-214410 -->` is the last line of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created, and it contains the marker line.
- **FR-003**: The change touches exactly one file and adds exactly one line (excluding a created file's own existence); no other files, code or documentation are modified.
- **FR-004**: The marker line appears exactly once in the file.
- **FR-005**: The work is based on, and its PR targets, the branch `e2e-base-branch-20261021-214410`.
- **FR-006**: The issue's labels are neither added nor removed as part of this work.
- **FR-007**: The work is not decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261021-214410` shows one changed file with a single added line (the marker line).
- **SC-002**: The PR is discovered as linked to issue #9250 (the issue-to-PR link is recognised by the pipeline despite the non-default base).
- **SC-003**: The pipeline advances the issue through every stage without a manual intervention caused by the non-default base, including the review gate receiving its review data.
- **SC-004**: After the PR merges, the issue is closed and moved to Done.

## Assumptions

- The base branch `e2e-base-branch-20261021-214410` already exists on the remote, since it is the repository's default for this run.
- The marker file lives under `e2e/markers/` at the repository root, as stated in the request.
- Only the presence and position of the marker line matters; the file's other contents, if any, are irrelevant to the test.
- Because the stages run with `fabrik:cruise`, they advance automatically but the PR is not auto-merged.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configuration or documentation.
- Splitting the work into child issues.
- Adding or removing labels on the issue.
- Changes outside this single repository.
- Fixes for #1046, #1047 and #1050 themselves, which are already implemented; this issue only verifies them.
