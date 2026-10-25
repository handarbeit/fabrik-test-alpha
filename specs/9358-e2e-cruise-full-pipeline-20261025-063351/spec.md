# Feature Specification: e2e cruise full pipeline marker (20261025-063351)

**Feature Branch**: `fabrik/issue-9358`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261025-063351 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end test fixture for the `fabrik:cruise` label (#898). Cruise should auto-advance an issue through every pipeline stage up to Validate-complete, stop there without merging the linked PR, and let the issue close correctly once a human merges the PR. The code change is deliberately trivial, one marker line in one file, so the run exercises the pipeline's behaviour and not the complexity of the work. The timestamped marker (`20261025-063351`) makes this run's change unique and identifiable.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the cruise marker file (Priority: P1)

The pipeline appends one HTML comment line, `<!-- cruise-pipeline-20261025-063351 -->`, at the end of `e2e/markers/cruise-full-pipeline.md`. If the file does not exist, it is created first.

**Why this priority**: This is the only code change. It gives the pipeline a real, reviewable diff and a PR to carry through the stages.

**Independent Test**: Inspect the PR diff. It should show exactly one added line in `e2e/markers/cruise-full-pipeline.md`, and that line is the marker above.

**Acceptance Scenarios**:

1. **Given** the marker file exists, **When** the change is implemented, **Then** the marker line is the last line of the file and all prior content is unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.

---

### User Story 2 - Cruise carries the issue to Validate-complete without merging (Priority: P1)

With `fabrik:cruise` set, the issue moves through Specify, Research, Plan, Implement, Review and Validate without manual advancement. It stops at Validate-complete and the linked PR stays unmerged.

**Why this priority**: This is the contract under test (#898).

**Independent Test**: Observe the issue's stage labels and board column after the run. All stages should be complete, the PR should be open, and no merge should have happened.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage with no human action.
2. **Given** Validate has completed, **When** the engine reaches its end-of-Validate decision, **Then** the PR is not merged and the issue is not moved to Done.

---

### User Story 3 - Issue closes correctly after a human merges the PR (Priority: P2)

After the pipeline stops at Validate-complete, a human merges the PR. The issue then closes and reaches Done.

**Why this priority**: This is the second half of the contract. It matters, but it depends on Stories 1 and 2.

**Independent Test**: Merge the PR by hand, then check that the linked issue is closed and in Done.

**Acceptance Scenarios**:

1. **Given** the PR is open and Validate is complete, **When** a human merges it, **Then** the issue is closed and moved to Done.

---

### Edge Cases

- The marker file already ends without a trailing newline. The marker must still land on its own line.
- The marker file already contains this exact marker, for example after a retry. The line must not be duplicated.
- Plan must not decompose the work into child issues; it is a single-file, single-line change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the line `<!-- cruise-pipeline-20261025-063351 -->` as the final line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created and contain the marker line.
- **FR-003**: The change MUST touch no file other than `e2e/markers/cruise-full-pipeline.md`, and MUST NOT alter existing content in that file.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The work MUST be delivered as a single issue with no child issues spawned.
- **FR-006**: Under `fabrik:cruise`, the pipeline MUST auto-advance through all stages to Validate-complete and MUST NOT merge the linked PR.
- **FR-007**: After a human merges the PR, the issue MUST close and reach Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one added line, the marker, in exactly one file.
- **SC-002**: The issue reaches Validate-complete with no manual stage advancement.
- **SC-003**: At Validate-complete the PR is open and unmerged.
- **SC-004**: After the human merge, the issue is closed and in Done.
- **SC-005**: No child issues are created.

## Assumptions

- `e2e/markers/` may or may not contain `cruise-full-pipeline.md` at this point. Both cases are handled by FR-001 and FR-002.
- Appending a line in the repository's existing text-file style (newline-terminated) satisfies "at the very end of the file".
- The cruise behaviour is defined by #898 and the existing label semantics. This issue verifies it and does not change it.

## Out of Scope

- Any change to engine, stage, or plugin behaviour.
- Edits to files other than `e2e/markers/cruise-full-pipeline.md`.
- Decomposing the work into child issues.
- Merging the PR automatically; the merge is a manual step performed by a human.

## Source References

- #898 (the `fabrik:cruise` pipeline contract)
