# Feature Specification: E2E Smoke — Full Pipeline

**Feature Branch**: `fabrik/issue-3536`
**Created**: 2026-07-20
**Status**: Draft
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Append a single comment line to README.md at the very end of the file: `<!-- smoke-full-pipeline-20260720-164630 -->`. That's the entire change. One file, one line. Single repo only — no cross-repo work. The Plan stage should NOT decompose."

## Background

This issue exists to exercise Fabrik's full single-repo pipeline end-to-end — Specify → Research → Plan → Implement → Review → Validate → Done — with a real, mergeable PR, using a change so trivial that any pipeline failure can be attributed to the orchestration itself rather than to the complexity of the change. The `fabrik:yolo` label drives full auto-advance, including auto-merge of the PR at Validate completion.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Full pipeline smoke completes with a merged PR (Priority: P1)

As a Fabrik maintainer, I want a trivial, unambiguous issue to travel through every pipeline stage and land as a merged PR, so that I can confirm the end-to-end orchestration (board transitions, worktree lifecycle, PR creation, CI, auto-merge) works correctly on a single repo.

**Why this priority**: This is the only user story — the issue exists solely to validate pipeline mechanics via the smallest possible payload.

**Independent Test**: Can be fully tested by observing the issue move through each board column to Done and confirming a PR containing the one-line `README.md` change was merged to `main`.

**Acceptance Scenarios**:

1. **Given** the issue is in the Specify column with this spec approved, **When** the pipeline advances, **Then** Research, Plan, Implement, Review, and Validate each complete without requiring manual decomposition or additional scope.
2. **Given** the Implement stage runs, **When** it makes its change, **Then** the only diff introduced is a single new line `<!-- smoke-full-pipeline-20260720-164630 -->` appended at the end of `README.md`, with no other files touched.
3. **Given** Validate completes successfully and `fabrik:yolo` is set, **When** the engine processes completion, **Then** the linked PR is auto-merged into `main` and the issue moves to Done.

### Edge Cases

- What happens if `README.md` does not end with a trailing newline? The appended line must still be added cleanly as its own final line, not concatenated onto the last existing line.
- What happens if the Plan stage attempts to spawn child issues for this? It must not — the change is scoped to a single file in a single repo and must not be decomposed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The implementation MUST append exactly one line, `<!-- smoke-full-pipeline-20260720-164630 -->`, to the end of `README.md`.
- **FR-002**: The implementation MUST NOT modify any file other than `README.md`.
- **FR-003**: The Plan stage MUST NOT decompose this issue into sub-issues or child work items.
- **FR-004**: The pipeline MUST progress through Research, Plan, Implement, Review, and Validate without requiring human input, consistent with the `fabrik:yolo` label.
- **FR-005**: The resulting PR body MUST include `Closes #3536` so Fabrik can discover PR comments per standard convention.

### Key Entities

- N/A — this change involves no data entities, only a static documentation file edit.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A PR containing only the one-line `README.md` addition is created, passes CI, and is merged into `main`.
- **SC-002**: The issue reaches the Done column with no manual intervention required after Specify approval.
- **SC-003**: No sub-issues or additional repos are touched during the pipeline run.

## Assumptions

- The repository's default base branch is `main`.
- `README.md` exists at the repository root and is a plain Markdown file safe to append an HTML comment to.
- The `fabrik:yolo` label is present on the issue for the full run, enabling auto-advance and auto-merge.

## Out of Scope

- Any functional code change — this is a documentation-only smoke test.
- Cross-repo coordination or multi-repo PR chains.
