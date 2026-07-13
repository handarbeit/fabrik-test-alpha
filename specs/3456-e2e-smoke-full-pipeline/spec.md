# Feature Specification: E2E Smoke — Full Pipeline

**Feature Branch**: `fabrik/issue-3456`
**Created**: 2026-07-13
**Status**: Draft
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Append a single comment line to README.md at the very end of the file: `<!-- smoke-full-pipeline-20260713-120725 -->`. That's the entire change. One file, one line."

## Background

This issue exists to smoke-test the full Fabrik pipeline end to end on a single repository: an issue should travel from Specify through Research, Plan, Implement, Review, and Validate to Done, resulting in a merged PR — with no cross-repo decomposition. The change itself is intentionally trivial (a single appended comment line in `README.md`) so that the exercise validates pipeline mechanics rather than application logic.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Full pipeline completes with a merged PR (Priority: P1)

As the Fabrik operator running this smoke test, I want a trivial, unambiguous issue to flow through every pipeline stage and land as a merged PR, so that I can confirm the full single-repo pipeline works end to end.

**Why this priority**: This is the entire purpose of the issue — there is no other scope.

**Independent Test**: Can be fully tested by observing the issue progress through each stage's board column to Done, and confirming the linked PR is merged into `main`.

**Acceptance Scenarios**:

1. **Given** the issue is specified and researched, **When** the Plan stage runs, **Then** it produces a single, non-decomposed implementation plan (no child issues spawned).
2. **Given** the plan is approved, **When** the Implement stage runs, **Then** exactly one line is appended to the very end of `README.md`: `<!-- smoke-full-pipeline-20260713-120725 -->`, and no other files are modified.
3. **Given** the change is implemented, **When** Review and Validate run, **Then** they pass and the linked PR is merged into `main`, and the issue reaches Done.

---

### Edge Cases

- If `README.md` does not end with a trailing newline, the appended line must still be added as its own new line (not concatenated onto the last existing line).
- If `README.md` already contains this exact comment line (e.g. from a re-run), Implement should not duplicate it.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The pipeline MUST process this issue through all stages (Specify, Research, Plan, Implement, Review, Validate) to Done without human intervention beyond the `fabrik:yolo` label already applied.
- **FR-002**: The Plan stage MUST NOT decompose this issue into sub-issues or child issues — this is single-repo, single-file scope.
- **FR-003**: The Implement stage MUST append exactly the line `<!-- smoke-full-pipeline-20260713-120725 -->` as a new line at the very end of `README.md`, and MUST NOT modify any other file.
- **FR-004**: The resulting PR MUST be merged into `main` as part of Validate completing (per `fabrik:yolo` auto-merge behavior).

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR associated with this issue is merged into `main` and contains a single-line diff to `README.md` only.
- **SC-002**: The issue reaches the Done column with no spawned child issues.

## Assumptions

- The repository's default base branch is `main`, as stated in the issue.
- No cross-repo work is needed; this is confined to the single repository containing `README.md`.

## Out of Scope

- Any change to files other than `README.md`.
- Any multi-repo or decomposed sub-issue work.

## Source References

- `README.md` — the only file modified by this issue
