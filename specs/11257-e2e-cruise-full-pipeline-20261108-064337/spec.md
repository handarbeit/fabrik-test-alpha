# Feature Specification: e2e cruise full pipeline (20261108-064337)

**Feature Branch**: `fabrik/issue-11257`
**Created**: 2026-11-08
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261108-064337 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end verification of the `fabrik:cruise` pipeline contract (#898). Cruise should auto-advance an item through every stage up to Validate-complete, stop there without auto-merging the linked PR, and let the issue close correctly once a human merges the PR. The code change is deliberately trivial: a single marker line in one file. That keeps the pipeline behaviour, not the change, as the thing under test. The run is identified by the timestamp `20261108-064337`, so this run's marker is distinguishable from any other.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is recorded in the cruise marker file (Priority: P1)

An operator running the cruise e2e check needs a minimal, unambiguous change that produces a real PR. The change is one marker line appended to the end of the cruise-full-pipeline marker file, creating the file first if it does not exist.

**Why this priority**: Without the change there is no PR, and without a PR the cruise contract cannot be exercised.

**Independent Test**: Inspect the PR diff. It touches only `e2e/markers/cruise-full-pipeline.md`, and the final line of that file is the marker line for this run.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist, **When** the change is made, **Then** the file is created and contains the marker line `<!-- cruise-pipeline-20261108-064337 -->` as its last line.
2. **Given** the marker file already exists with content, **When** the change is made, **Then** the existing content is unchanged and the marker line is appended as the new final line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

An operator wants confirmation that an item labelled `fabrik:cruise` moves through every pipeline stage automatically and then holds at Validate-complete with the PR left unmerged.

**Why this priority**: This is the contract under test (#898). The marker change exists only to drive it.

**Independent Test**: Observe the item progress from Specify through Validate with no human approval between stages. Confirm the linked PR is still open and unmerged after Validate completes.

**Acceptance Scenarios**:

1. **Given** the item carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without waiting for a human.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** the linked PR is not merged and the item is not moved to Done.

---

### User Story 3 - Issue closes correctly after a human merge (Priority: P2)

After the pipeline stops at Validate-complete, a human merges the PR, and the issue should close and complete as normal.

**Why this priority**: It completes the contract by showing that the human merge is what finishes the work. It depends on the earlier stories.

**Independent Test**: Merge the PR by hand and confirm the issue closes and reaches Done.

**Acceptance Scenarios**:

1. **Given** the item is at Validate-complete with an open PR, **When** a human merges the PR, **Then** the issue is closed and the item reaches Done.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker must still land on its own final line.
- The marker file is absent: it is created, so the marker is not lost.
- Plan must not split this into child issues. The work is one file and one line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The exact line `<!-- cruise-pipeline-20261108-064337 -->` is appended as the last line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it is created before the marker line is added.
- **FR-003**: Any existing content of the marker file is preserved unchanged.
- **FR-004**: The change touches no file other than the marker file.
- **FR-005**: The work is not decomposed into child issues.
- **FR-006**: With `fabrik:cruise` present, the item advances through all stages to Validate-complete without human approval between stages.
- **FR-007**: The linked PR is not auto-merged at Validate-complete, and the item is not advanced to Done by the engine.
- **FR-008**: After a human merges the PR, the issue closes correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one added line in exactly one file, `e2e/markers/cruise-full-pipeline.md`.
- **SC-002**: The final line of the marker file equals `<!-- cruise-pipeline-20261108-064337 -->`.
- **SC-003**: The item reaches Validate-complete with no human action between stages.
- **SC-004**: At Validate-complete the linked PR is still open and unmerged.
- **SC-005**: After a human merge, the issue is closed and the item is in Done.

## Assumptions

- The marker line is appended exactly as given, with no extra whitespace or surrounding text.
- A trailing newline after the marker line is acceptable and expected for a text file.
- The `e2e/markers/` directory is created if it is missing, as part of creating the file.
- Only one repository is involved; no cross-repo work is needed.

## Out of Scope

- Any change to engine behaviour, cruise semantics, or the e2e harness.
- Any file other than the cruise-full-pipeline marker file.
- Decomposition into child issues or follow-up work.
- Auto-merging the PR, which cruise deliberately does not do.
