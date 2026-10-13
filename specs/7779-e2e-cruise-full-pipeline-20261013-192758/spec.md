# Feature Specification: E2E cruise full pipeline marker (20261013-192758)

**Feature Branch**: `fabrik/issue-7779`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261013-192758 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

The `fabrik:cruise` label is meant to auto-advance an issue through every pipeline stage without auto-merging the PR. Advancing to Done is left to a human merge. The pipeline contract (#898) is verified by running a deliberately trivial issue through the whole pipeline. The marker file `e2e/markers/cruise-full-pipeline.md` already exists. Each run appends one uniquely timestamped HTML comment line, which gives the PR a real but harmless diff. The change is intentionally minimal so that any failure points at the pipeline rather than the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

An operator running the cruise E2E check needs a PR that contains exactly one added line in the marker file. That gives the pipeline a real diff to carry through all stages.

**Why this priority**: The line is the only deliverable. Without it there is no PR diff to carry through the pipeline.

**Independent Test**: Inspect the PR diff. It should touch only `e2e/markers/cruise-full-pipeline.md` and add exactly one line, `<!-- cruise-pipeline-20261013-192758 -->`, at the end of the file.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier cruise marker lines, **When** the change is implemented, **Then** the file's last line is `<!-- cruise-pipeline-20261013-192758 -->` and all earlier lines are unchanged.
2. **Given** the change is implemented, **When** the PR diff is reviewed, **Then** it contains one added line in one file and no other changes.

---

### User Story 2 - Cruise carries the issue to Validate-complete without merging (Priority: P1)

An operator wants confirmation that an issue labelled `fabrik:cruise` advances through every stage up to Validate-complete automatically, and that the PR stays open and unmerged at that point.

**Why this priority**: This is the contract under test (#898).

**Independent Test**: Observe the issue after the pipeline runs. All stages should be complete, Validate should be marked complete, and the PR should still be open.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the PR is not merged and the issue is not moved to Done.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

After the pipeline stops at Validate-complete, a human merges the PR and the issue should close correctly.

**Why this priority**: It verifies the end of the lifecycle, but it depends on Stories 1 and 2.

**Independent Test**: Merge the PR by hand and confirm the issue closes and reaches Done.

**Acceptance Scenarios**:

1. **Given** the PR is open at Validate-complete, **When** a human merges it, **Then** the linked issue closes and moves to Done.

---

### Edge Cases

- The marker file is missing. The change creates it first and then appends the line.
- The file's last line has no trailing newline. The new line must still start on its own line and not be joined to the previous one.
- The same timestamped line is already present. This is not expected, since the timestamp is unique per run.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- cruise-pipeline-20261013-192758 -->`, at the end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: The change MUST create the file first if it does not exist.
- **FR-003**: Existing content in the marker file MUST remain unchanged.
- **FR-004**: The change MUST NOT modify any file other than the marker file.
- **FR-005**: The work MUST be implemented as a single unit and MUST NOT be decomposed into child issues.
- **FR-006**: With `fabrik:cruise` present, the pipeline MUST auto-advance through all stages to Validate-complete and MUST NOT merge the PR.
- **FR-007**: After a human merges the PR, the issue MUST close.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one added line in exactly one file.
- **SC-002**: The issue reaches Validate-complete with no human advance action, and the PR is still unmerged at that point.
- **SC-003**: After a human merge, the issue is closed.
- **SC-004**: No child issues are spawned.

## Assumptions

- The marker file already exists in the repository, so no creation is needed. The creation fallback is kept for safety.
- The timestamp in the line identifies this run and is used as given.
- Only the single repository is involved. No cross-repo behaviour is exercised.

## Out of Scope *(optional)*

- Any change to the engine, stage configuration or other marker files.
- Auto-merging the PR, which cruise deliberately does not do.
- Decomposition into sub-issues.

## Source References *(optional)*

- Pipeline contract: #898
