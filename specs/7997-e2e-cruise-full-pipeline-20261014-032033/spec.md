# Feature Specification: e2e cruise full pipeline marker (20261014-032033)

**Feature Branch**: `fabrik/issue-7997`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261014-032033 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` contract (#898). With the `fabrik:cruise` label, an item should auto-advance through every pipeline stage up to and including Validate-complete. It must not auto-merge the PR or advance to Done on its own. After a human merges the PR, the issue should close correctly.

The content change is deliberately trivial so that the pipeline behaviour is what gets exercised. The change is one marker line appended to a single file. It is not meant to be decomposed into sub-tasks or child issues.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the cruise-pipeline marker line (Priority: P1)

An operator running the cruise e2e check needs a unique, identifiable marker line to appear at the end of the cruise marker file. The line must be unique to this run, so the resulting PR diff shows that this issue's change, and no other, went through the pipeline.

**Why this priority**: This is the only functional change. The whole e2e check depends on it producing a real, reviewable PR diff.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the issue branch. Its final line should be exactly `<!-- cruise-pipeline-20261014-032033 -->`. The diff against the base branch should add only that line, plus the file itself if it did not exist.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` already exists, **When** the change is applied, **Then** the file's last line is `<!-- cruise-pipeline-20261014-032033 -->` and all prior content is unchanged.
2. **Given** `e2e/markers/cruise-full-pipeline.md` does not exist, **When** the change is applied, **Then** the file is created and contains the marker line.
3. **Given** the change is applied, **When** the PR diff is reviewed, **Then** it touches exactly one file and adds exactly one line, apart from the file's creation if needed.

---

### User Story 2 - Cruise carries the item to Validate-complete without merging (Priority: P1)

An operator wants the pipeline to carry this issue through all stages to Validate-complete automatically. The PR stays open for a human to merge.

**Why this priority**: This is the purpose of the issue. It is the cruise contract under test.

**Independent Test**: With only `fabrik:cruise` applied, observe the item auto-advance Specify → Research → Plan → Implement → Review → Validate. Confirm that the PR is not merged and the issue does not move to Done when Validate completes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** the PR remains open and unmerged and the issue is not moved to Done.
3. **Given** a human merges the PR after Validate-complete, **When** the merge lands, **Then** the issue closes and reaches Done.

---

### Edge Cases

- The marker file already ends without a trailing newline. The appended marker must still sit on its own final line, and existing content must not be corrupted.
- The marker line already exists in the file, for example after a retry. It must not be duplicated.
- Plan must not decompose this work or spawn child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- cruise-pipeline-20261014-032033 -->` at the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, the change MUST create it first, so that the marker line is the file's last line.
- **FR-003**: The change MUST NOT modify any pre-existing content of the marker file, and MUST NOT modify any other file.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The work MUST stay as a single issue and a single PR. Plan MUST NOT decompose it or spawn child issues.
- **FR-006**: The PR MUST NOT be auto-merged by the pipeline. With `fabrik:cruise`, the item reaches Validate-complete and then waits for a human to merge.
- **FR-007**: After a human merges the PR, the issue MUST close.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The final line of `e2e/markers/cruise-full-pipeline.md` on the merged branch equals `<!-- cruise-pipeline-20261014-032033 -->`, and the marker string occurs exactly once.
- **SC-002**: The PR diff contains exactly one changed file with exactly one added marker line, plus the file's creation if it was absent.
- **SC-003**: The item passes through all six stages (Specify, Research, Plan, Implement, Review, Validate) with no human intervention before Validate-complete.
- **SC-004**: At Validate-complete, the PR is still open and unmerged and the issue is not in Done. After a human merges the PR, the issue is closed.

## Assumptions

- The marker file lives in the managed repository at `e2e/markers/cruise-full-pipeline.md`, and its existing format is a free-form markdown file of marker comments.
- The timestamp `20261014-032033` in the marker is a run identifier. It is used as given and does not need to match the date the spec was written.
- The repository's default base branch (`main`) is the PR target, since no `base:` label is present.

## Out of Scope *(optional)*

- Any change to engine, pipeline or cruise behaviour. This issue only observes it.
- Any file other than `e2e/markers/cruise-full-pipeline.md`.
- Decomposition into child issues or multi-repo work.
- Auto-merging the PR. A human performs the merge.
