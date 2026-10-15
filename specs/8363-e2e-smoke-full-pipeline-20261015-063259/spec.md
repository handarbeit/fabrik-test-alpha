# Feature Specification: e2e smoke full-pipeline (20261015-063259)

**Feature Branch**: `fabrik/issue-8363`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261015-063259 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. It exists to verify that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR. The code change is deliberately trivial, so that any failure points to the pipeline and not to the change. The issue carries the `fabrik:yolo` label, so the stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged (Priority: P1)

An operator running the smoke test wants a single marker line appended to the smoke-test marker file, delivered through the normal PR flow and merged, to show the pipeline works end to end.

**Why this priority**: This is the whole purpose of the issue. Without it the smoke test proves nothing.

**Independent Test**: After the issue reaches Done, check the default branch. The file `e2e/markers/smoke-full-pipeline.md` must end with the line `<!-- smoke-full-pipeline-20261015-063259 -->`, and the PR that introduced it must be merged and linked to this issue.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is implemented, **Then** the exact line `<!-- smoke-full-pipeline-20261015-063259 -->` is the last line of the file, and all earlier content is unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the PR is merged, **When** the pipeline finishes, **Then** the issue is in Done and the PR closes it through its `Closes #8363` link.

---

### Edge Cases

- If the marker file exists but its last line has no trailing newline, the marker line must still land on its own line and not be joined to the previous content.
- If the marker line is already present in the file (for example on a retry), it must not be added a second time.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- smoke-full-pipeline-20261015-063259 -->`, at the end of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, the change MUST create it, containing the marker line.
- **FR-003**: The change MUST touch only that one file and add only that one line, with no other additions, edits or deletions.
- **FR-004**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-005**: The change MUST be delivered as a PR that closes this issue and is merged to the default branch.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches the Done column with its PR merged and no manual intervention beyond the `fabrik:yolo` label.
- **SC-002**: The merged diff is exactly one file with exactly one added line, the marker line.
- **SC-003**: No child issues are created for this issue.

## Assumptions

- The `e2e/markers/` directory may or may not exist. If it does not, creating it as part of creating the file is acceptable.
- The marker file is a plain Markdown file, so an HTML comment line is inert in it.
- `fabrik:yolo` is the only automation label needed. It lets the stages auto-advance and the PR auto-merge.

## Out of Scope

- Any change other than the single marker line.
- Cross-repo work.
- Decomposing the work into child issues.
- Changes to pipeline, engine or stage configuration.
