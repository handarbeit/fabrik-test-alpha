# Feature Specification: e2e cruise full pipeline marker (20261030-203717)

**Feature Branch**: `fabrik/issue-10280`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261030-203717 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This issue is a live end-to-end probe of the `fabrik:cruise` contract (#898). Cruise should auto-advance an item through every pipeline stage up to Validate-complete, leave the PR unmerged, and let the issue close only once a human merges the PR. The probe needs a real but trivial change to carry through the pipeline. That change is one marker line appended to a tracked marker file. The unique timestamp in the line (`20261030-203717`) ties this run's change to this run's issue.

The change is deliberately minimal so that any pipeline behavior observed (stage advancement, PR creation, no auto-merge, issue closure on human merge) comes from the pipeline itself and not from the content of the change. Plan should not decompose it into child issues.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the marker file (Priority: P1)

An operator running the cruise e2e check wants this run's unique marker line present at the end of `e2e/markers/cruise-full-pipeline.md`, so the PR has a real, minimal diff to carry through the pipeline.

**Why this priority**: Without the change there is no PR diff, and the cruise pipeline contract cannot be exercised.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the PR branch and confirm its last line is exactly `<!-- cruise-pipeline-20261030-203717 -->` and that it is the only line added by the PR.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is applied, **Then** the file ends with `<!-- cruise-pipeline-20261030-203717 -->` as its final line and all previous content is unchanged.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created containing the marker line `<!-- cruise-pipeline-20261030-203717 -->`.
3. **Given** the change is applied, **When** the PR diff is viewed, **Then** it touches exactly one file and adds exactly one line.

---

### User Story 2 - Cruise carries the item through the pipeline without merging (Priority: P1)

An operator wants the `fabrik:cruise` item to advance through all stages to Validate-complete automatically, with the PR left open for a human to merge.

**Why this priority**: This is the contract under test (#898).

**Independent Test**: Observe the issue's stage progression and the PR's state once Validate completes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the PR is still open and unmerged, and the issue has not been moved to Done.
3. **Given** Validate-complete with the PR unmerged, **When** a human merges the PR, **Then** the issue closes correctly.

---

### Edge Cases

- The marker file is absent: it is created with only the marker line.
- The marker file exists but its last line lacks a trailing newline: the marker still lands on its own line and no existing line is altered.
- Plan must not split this into child issues, since the whole change is one line in one file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The PR MUST add the line `<!-- cruise-pipeline-20261030-203717 -->` to `e2e/markers/cruise-full-pipeline.md` as the last line of the file.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, the change MUST create it with that line as its content.
- **FR-003**: The change MUST NOT modify any existing content of the marker file and MUST NOT touch any other file.
- **FR-004**: The work MUST be delivered as a single issue and a single PR; Plan MUST NOT decompose it into child issues.
- **FR-005**: With `fabrik:cruise` present, the item MUST auto-advance through all stages to Validate-complete, and the PR MUST NOT be auto-merged.
- **FR-006**: After a human merges the PR, the issue MUST close correctly.

### Key Entities

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, the tracked file that receives the marker line.
- **Marker line**: `<!-- cruise-pipeline-20261030-203717 -->`, an HTML comment unique to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file with exactly one added line, matching the marker line character for character.
- **SC-002**: The item reaches Validate-complete through automatic stage advancement with no human intervention before that point.
- **SC-003**: At Validate-complete, the PR is open and unmerged.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker file may or may not already exist. Either case is handled per FR-001 and FR-002.
- Appending an HTML comment line to a Markdown file has no rendered effect and needs no other documentation updates.
- The repository is single-repo for this probe; no cross-repo behavior is exercised.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Changes to the cruise or pipeline behavior itself.
- Decomposing the work into child issues.
- Automated merging of the PR.
