# Feature Specification: E2E Smoke — Full Pipeline

**Feature Branch**: `fabrik/issue-3617`
**Created**: 2026-07-24
**Status**: Draft
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Append a single comment line to README.md at the very end of the file: `<!-- smoke-full-pipeline-20260724-143634 -->`. That's the entire change. One file, one line."

## Background

This issue exists to exercise the full Fabrik SDLC pipeline (Specify → Research → Plan → Implement → Review → Validate → Done) end-to-end on a single repository, confirming that an issue can be carried from creation through a merged PR without manual intervention. The change itself carries no product value — it exists solely as a deterministic, trivially-verifiable payload for the smoke test.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Pipeline carries a trivial change to a merged PR (Priority: P1)

As the operator running this smoke test, I need Fabrik to take this issue through every pipeline stage and land the specified one-line change in `main` via a merged PR, so that I can confirm the pipeline works end-to-end without manual steps.

**Why this priority**: This is the entire purpose of the issue — there is no secondary scenario.

**Independent Test**: After the pipeline reaches Done, `main` contains the exact comment line appended at the end of `README.md`, and the issue is closed with a merged PR that references it.

**Acceptance Scenarios**:

1. **Given** the issue is on the board in Specify, **When** the pipeline runs unattended (this issue carries `fabrik:yolo`), **Then** it progresses through Research, Plan, Implement, Review, and Validate without requiring manual comments, and the linked PR is merged.
2. **Given** the PR is merged, **When** `README.md` is inspected on `main`, **Then** its final line is exactly `<!-- smoke-full-pipeline-20260724-143634 -->` with no other content changes.

### Edge Cases

- If `README.md` already ends with a blank line or lacks a trailing newline, the implementer must still ensure the comment is the last line of the file with a single trailing newline.
- No other files should be modified; any additional diff content is a deviation from spec.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The Plan stage MUST NOT decompose this issue into sub-issues (single repo, single trivial change).
- **FR-002**: The Implement stage MUST append exactly one line, `<!-- smoke-full-pipeline-20260724-143634 -->`, to the end of `README.md`, and MUST NOT modify any other file.
- **FR-003**: The pipeline MUST carry the issue through to a merged PR and to Done without manual intervention, per the `fabrik:yolo` label.

### Key Entities

- N/A — this is a documentation-only, single-file change with no data entities.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `README.md` on `main` ends with the line `<!-- smoke-full-pipeline-20260724-143634 -->` after the pipeline completes.
- **SC-002**: Exactly one file (`README.md`) is changed across the entire PR diff.
- **SC-003**: The issue reaches the Done column with a merged PR, with no `fabrik:paused` or `fabrik:blocked` labels remaining.

## Assumptions

- "Append... at the very end of the file" means the comment line becomes the last line of `README.md`; no other existing content is altered, reordered, or removed.
- No tests, documentation updates elsewhere, or `docs/llms-full.txt` regeneration are needed, since none of the canonical doc pages are touched.

## Out of Scope

- Any cross-repo work or multi-issue decomposition.
- Any change to files other than `README.md`.

## Source References

- Issue #3617
