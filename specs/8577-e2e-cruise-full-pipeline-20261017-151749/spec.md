# Feature Specification: Cruise full-pipeline e2e marker (20261017-151749)

**Feature Branch**: `fabrik/issue-8577`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261017-151749 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an end-to-end check of the `fabrik:cruise` label's pipeline contract (#898). With cruise, an item should advance automatically through every stage up to Validate-complete, but the engine must not merge the PR itself. After a human merges the PR, the issue should close correctly. The code change is deliberately trivial, so that any failure is attributable to pipeline behaviour and not to the change.

The marker file `e2e/markers/cruise-full-pipeline.md` already exists. Earlier runs each appended one timestamped HTML comment line, and the most recent line is `<!-- cruise-pipeline-20261015-065400 -->`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the run marker (Priority: P1)

An e2e harness operator needs this run's unique marker line added to the shared marker file, so the resulting PR has a minimal, recognisable diff.

**Why this priority**: It is the only change this issue makes, and the pipeline run depends on it.

**Independent Test**: Inspect the PR diff. It contains exactly one added line, `<!-- cruise-pipeline-20261017-151749 -->`, at the end of `e2e/markers/cruise-full-pipeline.md`, and no other changes to that file.

**Acceptance Scenarios**:

1. **Given** the marker file exists with prior marker lines, **When** the change is made, **Then** the new line is the last line of the file and every prior line is unchanged.
2. **Given** the marker file does not exist, **When** the change is made, **Then** the file is created containing the single marker line.

---

### User Story 2 - Cruise pipeline contract holds (Priority: P1)

The e2e harness needs to confirm that an item labelled `fabrik:cruise` advances through all stages to Validate-complete without the engine merging the PR, and that the issue closes once a human merges the PR.

**Why this priority**: This is the purpose of the run (verifying #898).

**Independent Test**: Observe the item through the pipeline. Check that it reaches Validate-complete with the PR open and unmerged. Then merge the PR by hand and check that the issue closes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** the PR remains open and unmerged, and the item is not moved to Done.
3. **Given** Validate is complete and a human merges the PR, **When** the engine next processes the item, **Then** the issue is closed and moved to Done.

---

### Edge Cases

- The marker file lacks a trailing newline: the new line must still be on its own line, not joined to the previous one.
- The marker line is already present (for example on a retry): it must not be added twice.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The final line of `e2e/markers/cruise-full-pipeline.md` MUST be exactly `<!-- cruise-pipeline-20261017-151749 -->`.
- **FR-002**: All existing content of the marker file MUST be preserved unchanged, with the new line appended after it.
- **FR-003**: The marker file MUST be created if it does not exist.
- **FR-004**: The change MUST be limited to that one file and that one line, excluding the spec artifact the engine generates. Plan MUST NOT decompose the work into child issues.
- **FR-005**: The engine MUST NOT merge the PR as part of the cruise pipeline. It stops at Validate-complete.
- **FR-006**: After a human merges the PR, the issue MUST close.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff, excluding the spec artifact, adds exactly one line to exactly one file.
- **SC-002**: The item reaches Validate-complete through automatic advancement only, with no manual stage approvals.
- **SC-003**: At Validate-complete the PR is open and unmerged.
- **SC-004**: Within one engine poll after a human merges the PR, the issue is closed.

## Assumptions

- The marker file already exists, so the "create if missing" branch is unlikely to be exercised.
- The timestamp `20261017-151749` is an opaque unique run identifier and is used verbatim.
- Validate's completion criteria are met by this minimal change.

## Out of Scope *(optional)*

- Any change to engine, stage or label behaviour.
- Edits to other marker files.
- Decomposing the work into child issues.
- Merging the PR automatically (the point of the test is that the engine does not).
