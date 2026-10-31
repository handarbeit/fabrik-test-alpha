# Feature Specification: e2e base-branch pipeline (20261031-161639)

**Feature Branch**: `fabrik/issue-10529`
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
<!-- base-branch-pipeline-20261031-161639 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261031-161639 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end run that exercises the `base:<branch>` pipeline contract: an issue carrying a `base:<branch>` label must be worked on a branch forked from that base, and its pull request must target that base rather than the repository default. It is regression coverage for handarbeit/fabrik#1046, and validates the fixes for #1047 (issue↔PR linkage when the PR targets a non-default base) and #1050 (the review gate being fed the correct review data on such a PR).

The code change is deliberately trivial. It exists only so that the pipeline has a real, minimal diff to carry through every stage on the non-default base branch `e2e-base-branch-20261031-161639`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

An operator running the base-branch e2e wants a one-line marker added to a tracked file, delivered through the full pipeline on the non-default base, so they can confirm the pipeline honours `base:<branch>` end to end.

**Why this priority**: This is the whole change. The pipeline's handling of the non-default base is what is being verified.

**Independent Test**: After the pipeline finishes, check that the marker line is the last line of `e2e/markers/base-branch-pipeline.md` on the target base branch, and that the file differs from its prior state by exactly that one line.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the change is applied, **Then** the file's final line is `<!-- base-branch-pipeline-20261031-161639 -->` and all prior content is unchanged.
2. **Given** the file does not exist on the base branch, **When** the change is applied, **Then** the file is created and contains the marker line as its last line.
3. **Given** the issue carries the `base:e2e-base-branch-20261031-161639` label, **When** the pipeline creates its branch and PR, **Then** the branch is forked from, and the PR targets, `e2e-base-branch-20261031-161639`.

---

### Edge Cases

- The marker file already ends without a trailing newline: the marker must still land on its own line, not be joined onto the previous line.
- The marker line is already present (for example on a re-run): it must not be duplicated.
- Labels on the issue must stay as filed; nothing in the pipeline adds or removes `base:` labels.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261031-161639 -->`, as the final line of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, the change MUST create it, containing the marker line.
- **FR-003**: The change MUST modify only that one file, and MUST NOT alter any pre-existing content in it.
- **FR-004**: The work MUST be based on, and its PR MUST target, the branch named by the issue's `base:` label, `e2e-base-branch-20261031-161639`.
- **FR-005**: The issue's labels MUST NOT be added to or removed by this work.
- **FR-006**: The work MUST NOT be decomposed into child issues; it is a single change in a single repository.

### Key Entities

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a tracked file whose last line records this run's marker.
- **Base branch**: `e2e-base-branch-20261031-161639`, the non-default branch that the PR must target.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file and exactly one added line, the marker.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261031-161639`, and the PR is linked to this issue.
- **SC-003**: The issue progresses through all pipeline stages, including the review gate, without a linkage or review-data failure on the non-default base.
- **SC-004**: The marker appears exactly once in the file.

## Assumptions

- The repository's default base branch for this run is `e2e-base-branch-20261031-161639`, as stated in the pipeline context, so the label and default coincide here.
- The marker line has no trailing content, and the file ends with a newline after it.
- The directory `e2e/markers/` may need to be created if absent.

## Out of Scope

- Any change beyond the single marker line.
- Fixes to #1046, #1047 or #1050 themselves; they are already implemented, and this run only verifies them.
- Adding or removing labels, or changing board configuration.
- Splitting the work into child issues.

## Source References

- handarbeit/fabrik#1046, #1047, #1050 (the base-branch regression and its fixes)
