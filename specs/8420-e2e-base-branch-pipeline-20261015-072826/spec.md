# Feature Specification: e2e base-branch pipeline (20261015-072826)

**Feature Branch**: `fabrik/issue-8420`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261015-072826 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261015-072826 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end fixture. Its purpose is not the content of the change but to push a minimal, deterministic change through the full pipeline (Specify → Research → Plan → Implement → Review → Validate) on an item labelled `base:e2e-base-branch-20261015-072826`, so that the non-default-base contract is exercised for real. It is regression coverage for handarbeit/fabrik#1046, validating the fixes in #1047 (issue↔PR linkage for a PR that targets a non-default base) and #1050 (the review-gate data feed). The change is deliberately trivial so that any failure is attributable to the base-branch handling rather than to the work itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

A single marker line is appended to the base-branch pipeline marker file, delivered through a PR that targets the non-default base branch `e2e-base-branch-20261015-072826`.

**Why this priority**: This is the whole change and the signal the e2e run asserts on.

**Independent Test**: After the pipeline completes, the marker file on the target base branch ends with the exact marker line, and nothing else in the repository differs.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist on the base branch, **When** the change is implemented, **Then** the file exists and contains the marker line `<!-- base-branch-pipeline-20261015-072826 -->` as its last line.
2. **Given** the marker file already exists, **When** the change is implemented, **Then** the existing content is unchanged and the marker line is appended at the very end.

---

### User Story 2 - Pipeline honours the non-default base end to end (Priority: P2)

The work branch is forked from, rebased onto, and its PR targets `e2e-base-branch-20261015-072826` rather than the repository's default branch, and the issue↔PR linkage and review-gate data are correct for that PR.

**Why this priority**: This is the contract under regression test (#1046, #1047, #1050), though it is observed through the pipeline rather than produced by the change.

**Independent Test**: Inspect the PR created for the issue: its base is `e2e-base-branch-20261015-072826`, and the engine discovers the PR and advances the item through review.

**Acceptance Scenarios**:

1. **Given** the issue carries label `base:e2e-base-branch-20261015-072826`, **When** the Implement stage opens a PR, **Then** the PR's base branch is `e2e-base-branch-20261015-072826`.
2. **Given** that PR exists, **When** the review gate evaluates it, **Then** the PR is found and its review data is read for the correct PR.

---

### Edge Cases

- The marker file or its parent directory `e2e/markers/` does not exist: it is created first, then the line appended.
- The existing file lacks a trailing newline: the marker must still land on its own line at the end.
- The marker is already present (re-run): it must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/base-branch-pipeline.md` ends with the line `<!-- base-branch-pipeline-20261015-072826 -->`, created first if it does not exist.
- **FR-002**: The change touches exactly one file and adds exactly one line; no existing content is modified or removed.
- **FR-003**: The marker line appears exactly once in the file.
- **FR-004**: The work is delivered as a single unit; it is not decomposed into child issues.
- **FR-005**: The PR for this work targets the non-default base branch `e2e-base-branch-20261015-072826`.
- **FR-006**: No labels are added to or removed from the issue as part of the change.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The diff between the PR head and its base is exactly one added line in `e2e/markers/base-branch-pipeline.md` (or a new file containing only that line).
- **SC-002**: The PR's base branch equals `e2e-base-branch-20261015-072826`, and the pipeline discovers the PR and carries the item through the review stage without manual intervention.
- **SC-003**: The marker line occurs exactly once in the final file on the base branch after merge.

## Assumptions

- The branch `e2e-base-branch-20261015-072826` already exists on the remote.
- The `base:` and `fabrik:cruise` labels are already applied and must be left as-is.
- The marker line text is exact and case-sensitive, with no surrounding whitespace.

## Out of Scope *(optional)*

- Any change to files other than `e2e/markers/base-branch-pipeline.md`.
- Any change to engine behaviour; this issue only exercises it.
- Decomposing the work into child issues.
- Adding or removing labels.
