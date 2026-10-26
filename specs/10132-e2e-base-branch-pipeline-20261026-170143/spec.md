# Feature Specification: e2e base-branch pipeline (20261026-170143)

**Feature Branch**: `fabrik/issue-10132`
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
<!-- base-branch-pipeline-20261026-170143 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261026-170143 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is an end-to-end regression probe for the `base:<branch>` pipeline contract, i.e. how an issue pipeline behaves when the work targets a non-default base branch. It guards against a regression of handarbeit/fabrik#1046 and validates the fixes in #1047 (issue↔PR linkage) and #1050 (review-gate data feed). The code change is intentionally minimal; the value is in the pipeline running end to end against the non-default base branch `e2e-base-branch-20261026-170143`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

A single marker line is appended to the e2e marker file and delivered through the full pipeline to the non-default base branch.

**Why this priority**: It is the only deliverable; it is the artifact the e2e run observes.

**Independent Test**: After the pipeline finishes, the marker file on `e2e-base-branch-20261026-170143` ends with `<!-- base-branch-pipeline-20261026-170143 -->`, and the file differs from its prior state by exactly that one line (or consists solely of that line if it was newly created).

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the change is implemented, **Then** the file's last line is `<!-- base-branch-pipeline-20261026-170143 -->` and no other line changed.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created containing the marker line.
3. **Given** the issue carries the `base:e2e-base-branch-20261026-170143` label, **When** the PR is created, **Then** it targets that branch, not the repository default, and is linked to the issue (`Closes #10132`).

---

### Edge Cases

- The file already ends without a trailing newline: the marker must still occupy its own line.
- The marker line is already present (re-run): it must not be duplicated.
- The base branch does not exist on the remote: out of this issue's control; the engine's documented fallback behaviour applies.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- base-branch-pipeline-20261026-170143 -->` MUST be appended as the last line of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created containing the marker line.
- **FR-003**: The change MUST consist of exactly one file and one added line (plus the file's creation if absent); no other files are modified.
- **FR-004**: The marker line MUST NOT be duplicated if it is already present.
- **FR-005**: The work MUST target the base branch `e2e-base-branch-20261026-170143`, and the resulting PR MUST be linked to this issue.
- **FR-006**: The issue's labels MUST NOT be added or removed as part of this work.
- **FR-007**: The work MUST NOT be decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff against `e2e-base-branch-20261026-170143` touches exactly one file and adds exactly one line (the marker).
- **SC-002**: The PR's base branch is `e2e-base-branch-20261026-170143` and it is linked to issue #10132.
- **SC-003**: The issue progresses through all pipeline stages with no manual intervention and no spawned child issues.

## Assumptions

- The marker file may or may not already exist; either case is handled by FR-001/FR-002.
- The `base:` label and `fabrik:cruise` label are already correctly set at filing and are left untouched.
- The marker uses the exact timestamp string given in the request.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Modifying labels, engine behaviour, or other repositories.
- Decomposing the work into sub-issues.
