# Feature Specification: e2e base-branch pipeline marker (20261026-054457)

**Feature Branch**: `fabrik/issue-9705`
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
<!-- base-branch-pipeline-20261026-054457 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261026-054457 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end fixture. Its purpose is not the content of the change, but to drive an item carrying a `base:<branch>` label through the full pipeline (Specify through Done) so the non-default-base contract can be verified. It provides regression coverage for handarbeit/fabrik#1046 and validates the fixes in #1047 (issue-to-PR linkage for a PR targeting a non-default base) and #1050 (the review gate's data feed for such a PR). The change is deliberately trivial so that any failure is attributable to the pipeline, not the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

The pipeline produces a pull request that adds one marker line to `e2e/markers/base-branch-pipeline.md`, targeting the base branch `e2e-base-branch-20261026-054457`, and the item completes the pipeline.

**Why this priority**: It is the only deliverable; it exercises the whole non-default-base path.

**Independent Test**: Inspect the merged PR diff and the file on the base branch.

**Acceptance Scenarios**:

1. **Given** the file `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, **When** the change is implemented, **Then** the file exists and its last line is `<!-- base-branch-pipeline-20261026-054457 -->`.
2. **Given** the file already exists, **When** the change is implemented, **Then** the existing content is unchanged and the marker line is appended at the very end.
3. **Given** the PR is opened, **When** it is inspected, **Then** it targets `e2e-base-branch-20261026-054457` and changes exactly one file by adding exactly one line.

---

### Edge Cases

- The file exists without a trailing newline: the marker must still occupy its own final line.
- The marker line is already present: it must not be duplicated.
- The `base:` label is the only base-branch mechanism; labels are not added or removed by the change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/base-branch-pipeline.md` MUST end with the line `<!-- base-branch-pipeline-20261026-054457 -->`, created if absent.
- **FR-002**: Existing content of the file, if any, MUST be preserved, with the marker appended after it.
- **FR-003**: The marker line MUST appear exactly once in the file.
- **FR-004**: The change MUST consist of exactly one file and one added line.
- **FR-005**: The work MUST target base branch `e2e-base-branch-20261026-054457`, and the issue's labels MUST NOT be added or removed as part of the change.
- **FR-006**: The work MUST NOT be decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261026-054457` shows one file changed with one line added and none removed.
- **SC-002**: The linked PR is associated with this issue and visible to the review gate, so the item advances through the pipeline without manual intervention.

## Assumptions

- The marker file is a plain Markdown file in a directory of similar per-run marker files.
- The date-stamped token in the marker line is intentional and used verbatim.
- Single repository only.

## Out of Scope

- Any change besides the one marker line.
- Splitting the work into child issues.
- Changes to labels or to the default branch.
- Changes to engine behavior.
