# Feature Specification: E2E cruise full pipeline marker (20261116-154949)

**Feature Branch**: `fabrik/issue-11381`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261116-154949 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898): an issue labelled `fabrik:cruise` should advance automatically through every pipeline stage up to Validate-complete, must not auto-merge its PR, and must close correctly once a human merges the PR. The code change is deliberately trivial (one marker line in one file) so that any failure observed is attributable to the pipeline, not the change. The marker file is a shared, append-only record of such runs; each run appends its own uniquely timestamped line.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended (Priority: P1)

A maintainer running the e2e check needs a single, uniquely identifiable marker line added to the cruise marker file, so the change is a minimal, verifiable diff.

**Why this priority**: It is the only code change; the pipeline run is meaningless without a PR carrying it.

**Independent Test**: Inspect the PR diff: exactly one file changed, with exactly one added line equal to `<!-- cruise-pipeline-20261116-154949 -->`, positioned at the end of the file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is applied, **Then** its final line is `<!-- cruise-pipeline-20261116-154949 -->` and all pre-existing lines are unchanged.
2. **Given** the file does not exist, **When** the change is applied, **Then** the file is created containing that single line.

---

### User Story 2 - Cruise pipeline contract holds (Priority: P1)

A maintainer needs to confirm the cruise label behaves as specified end to end.

**Why this priority**: This is the purpose of the issue.

**Independent Test**: Observe the issue through the pipeline: it reaches Validate-complete without human advancement, the PR remains unmerged until a human merges it, and the issue is closed after that merge.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** it advances to the next stage automatically through Validate.
2. **Given** Validate has completed, **When** no human has acted, **Then** the PR is not merged and the issue is not moved to Done.
3. **Given** a human merges the PR, **When** the merge completes, **Then** the issue closes and reaches Done.

---

### Edge Cases

- The marker file already ends without a trailing newline: the new line must still appear on its own line.
- The marker file already contains earlier run markers: they are preserved untouched.
- Plan must not decompose the work into child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- cruise-pipeline-20261116-154949 -->` MUST be the final line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: The file MUST be created if it does not exist; if it exists, existing content MUST be preserved unchanged.
- **FR-003**: The change MUST consist of exactly one file and one added line, with no other modifications.
- **FR-004**: The work MUST NOT be decomposed into child issues.
- **FR-005**: The change MUST be delivered in a single repository only.
- **FR-006**: The PR MUST NOT be auto-merged by the pipeline under `fabrik:cruise`; merging is a human action.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly 1 changed file and 1 added line matching the marker text.
- **SC-002**: The issue reaches Validate-complete via automatic advancement with zero manual stage moves.
- **SC-003**: The PR is unmerged at Validate-complete, and the issue is closed after a human merges it.

## Assumptions

- The marker file already exists at the stated path (observed prior run markers); the create-if-missing clause is a fallback only.
- The timestamp `20261116-154949` is the run identifier and is used verbatim.

## Out of Scope

- Any change beyond the single marker line.
- Multi-repo behaviour and child-issue spawning.
- Changes to engine code, stage configuration or documentation.
