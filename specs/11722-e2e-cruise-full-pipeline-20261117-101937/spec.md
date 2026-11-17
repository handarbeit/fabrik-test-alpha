# Feature Specification: e2e cruise full pipeline (20261117-101937)

**Feature Branch**: `fabrik/issue-11722`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261117-101937 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898). Cruise should carry an issue through every pipeline stage automatically, stop at Validate-complete without merging the linked PR, and close the issue correctly once a human merges that PR. The code change is deliberately tiny so that the pipeline's behavior, not the change, is what gets exercised. It is a single marker line appended to a marker file.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Cruise carries a trivial change through the pipeline (Priority: P1)

An operator labels an issue `fabrik:cruise`. The issue advances through Specify, Research, Plan, Implement, Review and Validate with no manual stage approval. The result is a PR containing the one-line marker change, left unmerged at Validate-complete. After a human merges the PR, the issue closes.

**Why this priority**: This is the whole purpose of the issue. It confirms the cruise contract end to end.

**Independent Test**: Apply `fabrik:cruise` to the issue and observe that every stage completes without manual approval. Then confirm the PR is not auto-merged, merge it by hand, and confirm the issue closes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the linked PR is still open and unmerged and the issue has not advanced to Done.
3. **Given** a human merges the PR, **When** the merge lands, **Then** the issue closes and reaches Done.

---

### User Story 2 - The marker change is applied correctly (Priority: P1)

The implementation appends exactly one HTML comment line to the marker file `e2e/markers/cruise-full-pipeline.md`, creating the file first if it does not exist.

**Why this priority**: It is the only code change, and it is the artifact the pipeline is meant to deliver.

**Independent Test**: Inspect the PR diff and confirm it touches only `e2e/markers/cruise-full-pipeline.md` and adds only the marker line at the end of that file.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist, **When** Implement runs, **Then** the file is created and its last line is `<!-- cruise-pipeline-20261117-101937 -->`.
2. **Given** the marker file already exists, **When** Implement runs, **Then** the marker line is appended as the final line and existing content is unchanged.

---

### Edge Cases

- The marker file already ends without a trailing newline: the marker must still be on its own line at the end of the file.
- The marker line already exists in the file (for example from a re-run): it should not be duplicated.
- Plan must not decompose the work into child issues, because the change is one file and one line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST append the line `<!-- cruise-pipeline-20261117-101937 -->` as the final line of `e2e/markers/cruise-full-pipeline.md`, creating the file first if it does not exist.
- **FR-002**: The change MUST be limited to that one file and that one line.
- **FR-003**: The Plan stage MUST NOT decompose this work into child issues.
- **FR-004**: With `fabrik:cruise` present, the issue MUST auto-advance through all stages to Validate-complete without manual approval.
- **FR-005**: The linked PR MUST NOT be auto-merged, and the issue MUST NOT advance to Done at Validate completion.
- **FR-006**: After a human merges the PR, the issue MUST close correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file with exactly one added line, matching the marker text.
- **SC-002**: The issue reaches Validate-complete with no manual stage approvals.
- **SC-003**: The PR is still open and unmerged when the pipeline settles at Validate-complete.
- **SC-004**: The issue is closed after a human merges the PR.

## Assumptions

- The `fabrik:cruise` label is already present on the issue, so no operator action is needed to start the pipeline.
- The marker file lives under `e2e/markers/` in the repository root.
- The marker file's other content, if any, is irrelevant to this check and is left untouched.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Child issue spawning or decomposition.
- Multi-repo behavior.
- Testing `fabrik:yolo` or auto-merge behavior.
