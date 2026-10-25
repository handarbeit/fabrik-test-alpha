# Feature Specification: e2e cruise full pipeline marker (20261025-133430)

**Feature Branch**: `fabrik/issue-9625`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261025-133430 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` contract (#898): a cruise-labelled issue should advance automatically through every stage up to Validate-complete, must not have its PR merged automatically, and should close correctly once a human merges the PR. The code change is deliberately trivial so that the pipeline behaviour is what is under test. The marker file `e2e/markers/cruise-full-pipeline.md` already exists and holds earlier runs' marker lines (the most recent is `<!-- cruise-pipeline-20261025-112932 -->`), so this run appends one more line.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

The marker file gains one new line identifying this run, and nothing else changes.

**Why this priority**: It is the only change this issue makes and gives the pipeline a real diff and PR to carry through the stages.

**Independent Test**: Diff the branch against `main`; the only change is one added line at the end of `e2e/markers/cruise-full-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier run lines, **When** the change is applied, **Then** the file ends with `<!-- cruise-pipeline-20261025-133430 -->` and all earlier lines are unchanged.
2. **Given** the change is applied, **When** the PR diff is inspected, **Then** exactly one file is modified, with exactly one added line.

---

### User Story 2 - Cruise pipeline contract holds (Priority: P1)

The issue moves through all stages automatically and stops short of merging.

**Why this priority**: This is the actual purpose of the issue (verifying #898).

**Independent Test**: Observe the issue from Specify to Validate-complete without human intervention, then merge the PR by hand and confirm the issue closes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** it advances to the next stage without human approval, up to Validate-complete.
2. **Given** Validate is complete, **When** the pipeline settles, **Then** the PR is not merged and the issue is not moved to Done by the engine.
3. **Given** a human merges the PR, **When** the engine next processes the issue, **Then** the issue ends up closed and in Done.

---

### Edge Cases

- If the marker file lacked a trailing newline, the new line must still be on its own line.
- If the file did not exist, it is created containing only the new line (it does exist today).
- Planning must not decompose this work into child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/cruise-full-pipeline.md` MUST end with the line `<!-- cruise-pipeline-20261025-133430 -->`.
- **FR-002**: Existing content of the marker file MUST be preserved unchanged; the new line is appended at the very end.
- **FR-003**: The change MUST touch only that one file and add only that one line.
- **FR-004**: The work MUST NOT be decomposed into child issues.
- **FR-005**: The engine MUST NOT automatically merge the PR or advance the issue to Done at Validate completion while `fabrik:cruise` is present.
- **FR-006**: After a human merges the PR, the issue MUST end up closed.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `main` is exactly one added line in one file.
- **SC-002**: The issue reaches Validate-complete with no manual stage advancement.
- **SC-003**: Zero automatic merges of the PR occur before a human merges it.
- **SC-004**: The issue is closed after the human merge.

## Assumptions

- The marker file already exists; if it did not, creating it would satisfy FR-001.
- The timestamp `20261025-133430` is the unique identifier of this run and is used literally.

## Out of Scope *(optional)*

- Any change to engine behaviour or to other files.
- Cleaning up earlier marker lines.
- Decomposition into child issues or multi-repo work.
