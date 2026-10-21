# Feature Specification: e2e cruise full pipeline (20261021-202348)

**Feature Branch**: `fabrik/issue-9209`
**Created**: 2026-10-21
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261021-202348 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This is an end-to-end smoke test of the `fabrik:cruise` label contract (#898). Under cruise, an issue should auto-advance through every pipeline stage up to Validate-complete, but the engine must not auto-merge the linked PR or move the issue to Done. After a human merges the PR, the issue should close correctly. The code change is intentionally trivial so the test exercises pipeline behavior rather than feature work. The marker file already exists and holds one such comment line per previous run, so this run appends a new, uniquely timestamped line.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the run marker (Priority: P1)

A pipeline operator wants a minimal, deterministic change that gives the cruise pipeline something to implement, review and validate.

**Why this priority**: Without the change there is no PR and nothing to carry through the stages.

**Independent Test**: Inspect the last line of `e2e/markers/cruise-full-pipeline.md` on the issue branch and confirm it is the marker line, and that the diff against the base branch is exactly one added line in one file.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier marker lines, **When** the change is applied, **Then** the last line of the file is `<!-- cruise-pipeline-20261021-202348 -->` and all earlier content is unchanged.
2. **Given** the change is complete, **When** the PR diff is reviewed, **Then** it shows one file changed with one line added.

---

### User Story 2 - Cruise advances without merging (Priority: P1)

An operator running with `fabrik:cruise` wants the issue to move through all stages automatically while leaving the merge to a human.

**Why this priority**: This is the contract under test (#898).

**Independent Test**: Observe the issue reaching Validate-complete with the PR still open and the issue not in Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the engine evaluates the next step, **Then** it does not merge the PR and does not move the issue to Done.
3. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue closes correctly.

---

### Edge Cases

- The marker file is missing: it is created first, then the line is appended.
- The file does not end in a newline: the marker still lands on its own line at the end.
- Plan must not decompose this issue into child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- cruise-pipeline-20261021-202348 -->` MUST be appended as the last line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: The file MUST be created first if it does not exist, and existing content MUST be left unchanged.
- **FR-003**: The change MUST consist of exactly one file and one added line, with no other modifications.
- **FR-004**: The work MUST NOT be decomposed into child issues.
- **FR-005**: With `fabrik:cruise` present, the issue MUST auto-advance through all stages to Validate-complete without the PR being merged or the issue moved to Done by the engine.
- **FR-006**: After a human merges the PR, the issue MUST close correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against the base branch is one file with one added line, equal to the marker above.
- **SC-002**: The issue reaches Validate-complete with no human stage approvals and with the PR open (not merged).
- **SC-003**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker file already exists in the repository; the create-if-missing instruction is a fallback.
- The timestamp `20261021-202348` is unique to this run and is used literally.
- Only the single repository is involved.

## Out of Scope

- Any change other than the single appended marker line.
- Auto-merging the PR or moving the issue to Done under cruise.
- Decomposition into child issues.
- Changes to engine behavior or to the cruise contract itself.
