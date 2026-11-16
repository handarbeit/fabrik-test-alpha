# Feature Specification: e2e base-branch pipeline marker (20261116-145227)

**Feature Branch**: `fabrik/issue-11346`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261116-145227 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261116-145227 (already set via
the base:<branch> label applied at filing — do not add or remove labels).
"

## Background

This issue is a live end-to-end exercise of the `base:<branch>` pipeline contract: an issue carrying a `base:<branch>` label must be forked from, and its PR must target, that non-default base branch. It is regression coverage for handarbeit/fabrik#1046, and it validates the fixes in #1047 (issue↔PR linkage) and #1050 (review-gate data feed). The code change is deliberately trivial, because the subject under test is the pipeline's behaviour on a non-default base, not the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

An operator files a trivial issue labelled `base:e2e-base-branch-20261116-145227`. The pipeline appends one marker line to `e2e/markers/base-branch-pipeline.md` and delivers it through a PR that targets that base branch, not the repository's main branch.

**Why this priority**: This is the whole change, and its delivery against the non-default base is what the test verifies.

**Independent Test**: After the pipeline finishes, inspect the PR diff against `e2e-base-branch-20261116-145227`. It should contain exactly one added line in `e2e/markers/base-branch-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist on the base branch, **When** the change is made, **Then** the file is created and its last line is `<!-- base-branch-pipeline-20261116-145227 -->`.
2. **Given** the marker file already exists on the base branch, **When** the change is made, **Then** the marker line is appended at the very end and existing content is unchanged.
3. **Given** the issue carries `base:e2e-base-branch-20261116-145227`, **When** the PR is created, **Then** the PR's base is that branch and the issue↔PR linkage is established (#1047).

---

### Edge Cases

- The marker file may be missing on the base branch; it is created in that case.
- The existing file may lack a trailing newline; the marker must still occupy its own final line.
- The marker line must not be duplicated if a stage re-runs.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- base-branch-pipeline-20261116-145227 -->` is appended as the last line of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created, with the marker line as its content.
- **FR-003**: The change touches only that one file and adds only that one line.
- **FR-004**: The work targets the non-default base branch `e2e-base-branch-20261116-145227` (fork point, rebase target and PR base), using the existing `base:` label without adding or removing labels.
- **FR-005**: The work is not decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged PR diff against `e2e-base-branch-20261116-145227` is exactly one file with one added line.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261116-145227` and it is linked to this issue.
- **SC-003**: The issue's labels are unchanged from filing: `fabrik:cruise` and `base:e2e-base-branch-20261116-145227`.
- **SC-004**: The pipeline reaches Done with no manual intervention and no child issues.

## Assumptions

- The branch `e2e-base-branch-20261116-145227` exists on the remote.
- A single trailing-newline-terminated line is the expected format for the appended marker.

## Out of Scope

- Any change besides the one marker line.
- Adding or removing labels.
- Multi-repo behaviour.
- Decomposing the work into child issues.
- Fixing #1046, #1047 or #1050; this issue only exercises them.
