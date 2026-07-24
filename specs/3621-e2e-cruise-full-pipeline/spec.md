# Feature Specification: E2E Cruise Full Pipeline

**Feature Branch**: `fabrik/issue-3621`
**Created**: 2026-07-24
**Status**: Draft
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line to README.md at the very end of the file: `<!-- cruise-pipeline-20260724-144031 -->`. That is the entire change. One file, one line. Plan should NOT decompose."

## Background

Fabrik's `fabrik:cruise` label (contract defined in #898) is meant to auto-advance an issue through all pipeline stages (Specify → Research → Plan → Implement → Review → Validate) without auto-merging the resulting PR, and without moving the issue to Done at Validate completion. This issue exists purely as a harness to exercise that contract end-to-end against a live GitHub Project board and confirm two things: the pipeline runs unattended through to Validate-complete, and the issue closes correctly once a human merges the PR afterward. It carries a deliberately trivial code change so the run is fast and any pipeline friction can be attributed to the orchestration logic rather than the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Cruise auto-advances through all stages without merging (Priority: P1)

As the Fabrik maintainer, I want an issue labeled `fabrik:cruise` to progress automatically from Specify through Validate without requiring manual advancement between stages, and without the pipeline merging the PR itself, so that I can confirm the cruise contract holds for a full, real run.

**Why this priority**: This is the entire purpose of the issue — without automatic stage advancement, the e2e test provides no signal.

**Independent Test**: Apply `fabrik:cruise` to an issue with a trivial one-line change and observe the pipeline run unattended from Specify through `stage:Validate:complete` with no human intervention between stages, and confirm the linked PR remains unmerged throughout.

**Acceptance Scenarios**:

1. **Given** issue #3621 is labeled `fabrik:cruise` and Specify has produced a clear spec, **When** the engine polls, **Then** it auto-advances to Research without waiting for manual approval.
2. **Given** each stage (Research, Plan, Implement, Review) completes with `FABRIK_STAGE_COMPLETE`, **When** the engine polls, **Then** it auto-advances to the next stage without a human comment or label change.
3. **Given** Validate completes successfully, **When** the engine processes completion, **Then** the issue reaches `stage:Validate:complete` and the linked PR is left open (unmerged) and the issue is left open (not moved to Done).

### Edge Cases

- Plan must not decompose this single-line change into sub-issues or multiple tasks — the change is intentionally atomic.
- If cruise auto-advance stalls at any stage (e.g. incorrectly waits for human input), that is a pipeline defect this issue is designed to surface.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The system MUST append the exact line `<!-- cruise-pipeline-20260724-144031 -->` to the end of `README.md`, and this MUST be the only content change in the PR.
- **FR-002**: The Plan stage MUST NOT decompose this change into multiple tasks or sub-issues.
- **FR-003**: With `fabrik:cruise` applied, the engine MUST auto-advance the issue through Research, Plan, Implement, Review, and Validate without requiring a human to manually move the board column or comment between stages.
- **FR-004**: Upon Validate completion, the engine MUST NOT auto-merge the linked PR (this distinguishes `fabrik:cruise` from `fabrik:yolo`).
- **FR-005**: Upon Validate completion, the engine MUST NOT move the issue to Done or close it — the issue MUST close only after a human merges the PR.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches `stage:Validate:complete` without any human comment or manual label/column change after `fabrik:cruise` is applied.
- **SC-002**: The linked PR remains open and unmerged at the point Validate completes.
- **SC-003**: After a human merges the PR, the issue closes correctly (verified as a separate, manual follow-up step outside this pipeline run).
- **SC-004**: The final diff touches exactly one file (`README.md`) with exactly one added line.

## Assumptions

- This issue targets a single repository (`fabrik-test-alpha`); no cross-repo coordination is involved.
- The reader is expected to know the `fabrik:cruise` label semantics documented in the Fabrik project's `CLAUDE.md` (auto-advance without auto-merge, contrasted with `fabrik:yolo`).
- Verifying the post-merge issue-close behavior (SC-003) happens after this automated pipeline run completes and is outside the scope of stages Specify–Validate.

## Out of Scope

- Any change to files other than `README.md`.
- Verifying `fabrik:yolo` behavior (auto-merge) — this issue is specifically about `fabrik:cruise`'s non-merging contract.
- Multi-repo or sub-issue spawning behavior.

## Source References

- Issue #898 — original `fabrik:cruise` pipeline contract definition.
- `CLAUDE.md` — `fabrik:cruise` label semantics.
