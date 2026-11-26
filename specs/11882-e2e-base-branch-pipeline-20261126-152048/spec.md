# Feature Specification: E2E base-branch pipeline marker (20261126-152048)

**Feature Branch**: `fabrik/issue-11882`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261126-152048 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261126-152048 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue exists to exercise the Fabrik pipeline end to end on an item whose `base:<branch>` label points at a non-default base branch. It is regression coverage for handarbeit/fabrik#1046, validating the fixes in #1047 (issue-to-PR linkage for non-default-base PRs) and #1050 (review-gate data feed). The code change is deliberately trivial so that the pipeline's behaviour, not the change itself, is what is under test.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands via the non-default base pipeline (Priority: P1)

An operator files this issue with a `base:e2e-base-branch-20261126-152048` label. The pipeline runs it through every stage, and the resulting PR targets that base branch, not the repository's historical default. The PR adds a single marker line to a marker file.

**Why this priority**: It is the only story. It carries the whole purpose of the issue, which is verifying the base-branch pipeline contract.

**Independent Test**: After the pipeline completes, read `e2e/markers/base-branch-pipeline.md` on the PR branch. Its last line is the marker comment. Confirm the PR's base is `e2e-base-branch-20261126-152048`.

**Acceptance Scenarios**:

1. **Given** the file `e2e/markers/base-branch-pipeline.md` already exists on the base branch, **When** the change is implemented, **Then** the file ends with the line `<!-- base-branch-pipeline-20261126-152048 -->` and its earlier content is unchanged.
2. **Given** the file does not exist on the base branch, **When** the change is implemented, **Then** the file is created and contains the marker line as its final line.
3. **Given** the issue carries the `base:e2e-base-branch-20261126-152048` label, **When** the PR is created, **Then** the PR targets that branch and is linked to this issue.

---

### Edge Cases

- The file exists but does not end with a newline: the marker still goes on its own line at the end of the file.
- The marker line is already present at the end of the file: it is not duplicated.
- The `e2e/markers/` directory does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/base-branch-pipeline.md` ends with exactly one occurrence of the line `<!-- base-branch-pipeline-20261126-152048 -->` as its final line.
- **FR-002**: If the file already exists, its existing content is preserved and the marker is appended; if it does not exist, it is created with the marker as its content.
- **FR-003**: The change touches only that one file and adds only that one line. No other files are modified.
- **FR-004**: The work targets the base branch `e2e-base-branch-20261126-152048`.
- **FR-005**: No labels are added to or removed from the issue by this work.
- **FR-006**: The work is not decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261126-152048` contains exactly one changed file and one added line (plus the file creation if it was absent).
- **SC-002**: The PR's base branch is `e2e-base-branch-20261126-152048` and the PR is linked to issue #11882.
- **SC-003**: The issue completes the full pipeline with no pause for human input and no spawned child issues.

## Assumptions

- The marker file is a plain-text Markdown file, and an HTML comment line is acceptable content in it.
- The label set (`fabrik:cruise`, `base:e2e-base-branch-20261126-152048`) is intentional and stays unchanged.
- The base branch `e2e-base-branch-20261126-152048` already exists on the remote.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Changes to engine behaviour or the stage configuration.
- Decomposing the work into multiple issues.
- Adding or removing labels.

## Source References *(optional)*

- handarbeit/fabrik#1046 (base-branch pipeline contract), #1047 (issue-to-PR linkage), #1050 (review-gate data feed)
