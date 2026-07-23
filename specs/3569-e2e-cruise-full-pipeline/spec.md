# Feature Specification: E2E Cruise Full Pipeline Smoke Test (20260723-001046)

**Feature Branch**: `fabrik/issue-3569`
**Created**: 2026-07-22
**Status**: Draft
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line to README.md at the very end of the file: `<!-- cruise-pipeline-20260723-001046 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

Fabrik's `fabrik:cruise` label is meant to make an issue auto-advance through every pipeline stage (Specify → Research → Plan → Implement → Review → Validate) without requiring a human to manually approve each stage transition — but, unlike `fabrik:yolo`, cruise must stop short of auto-merging the resulting PR or auto-closing the issue. This issue exists purely to exercise that contract end-to-end against a real, trivial code change, confirming both halves of the behavior: the pipeline auto-advances all the way to Validate-complete, and it then waits for a human to merge the PR before the issue closes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Cruise auto-advances through all stages (Priority: P1)

As a Fabrik operator, when I label an issue `fabrik:cruise`, I expect the pipeline to move the issue through every stage automatically — without pausing for manual "proceed" approval between stages — until it reaches Validate-complete.

**Why this priority**: This is the core behavior under test; without it the cruise contract is unverified.

**Independent Test**: Create this issue with `fabrik:cruise` applied, let Fabrik process it, and observe the issue's stage labels progress from Specify through Validate without manual intervention, reaching `stage:Validate:complete`.

**Acceptance Scenarios**:

1. **Given** issue #3569 is labeled `fabrik:cruise` and Specify has produced this spec, **When** Fabrik picks up the issue on its next poll, **Then** it proceeds automatically into Research, Plan, Implement, Review, and Validate in sequence, without requiring a human comment to advance between stages.
2. **Given** the issue reaches Validate, **When** Validate completes successfully, **Then** the issue reaches `stage:Validate:complete` and a PR exists implementing the change, but the PR is not auto-merged and the issue is not auto-closed.

---

### User Story 2 - Issue closes after human merge (Priority: P2)

As a Fabrik operator, after cruise has brought the PR to a mergeable, validated state, I expect the issue to close correctly once I (a human) merge the PR myself.

**Why this priority**: Confirms the second half of the cruise contract — that Fabrik does not merge on its own, but does correctly react to and finalize a human-initiated merge.

**Independent Test**: After Validate-complete, manually merge the linked PR and confirm the issue transitions to closed/Done shortly afterward.

**Acceptance Scenarios**:

1. **Given** the PR linked to issue #3569 is validated and open, **When** a human merges the PR, **Then** Fabrik detects the merge and closes the issue / moves it to Done.

---

### Edge Cases

- What happens if Plan attempts to decompose this single-line change into multiple sub-tasks or child issues? It should not — the change is scoped to one file, one line, and Plan must treat it as a single unit of work.
- What happens if cruise is mistaken for yolo and the pipeline auto-merges the PR? This must not happen; auto-merge is exclusively a `fabrik:yolo` behavior (see [[Assumptions]]).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- cruise-pipeline-20260723-001046 -->`, to the end of `README.md`, with no other file modifications.
- **FR-002**: The Plan stage MUST NOT decompose this issue into multiple tasks or child issues — it is a single trivial change.
- **FR-003**: With `fabrik:cruise` applied, the pipeline MUST auto-advance through Research, Plan, Implement, Review, and Validate without requiring manual stage-advance approval.
- **FR-004**: The pipeline MUST NOT auto-merge the PR upon Validate-complete (cruise, unlike yolo, does not trigger auto-merge).
- **FR-005**: The pipeline MUST NOT auto-close the issue upon Validate-complete; the issue closes only after a human merges the linked PR.
- **FR-006**: This issue is confined to a single repository; no cross-repo coordination is involved.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `README.md` contains the line `<!-- cruise-pipeline-20260723-001046 -->` as its final line after Implement runs.
- **SC-002**: The issue reaches `stage:Validate:complete` without any manual "proceed" comment posted by a human between stages.
- **SC-003**: The linked PR remains open (unmerged) immediately after Validate-complete.
- **SC-004**: The issue transitions to closed/Done within one poll cycle after a human merges the PR.

## Assumptions

- `fabrik:cruise` and `fabrik:yolo` are distinct labels with distinct behavior: cruise auto-advances stages but leaves merge and issue-closure to a human; yolo additionally auto-merges the PR and closes the issue. This issue verifies cruise specifically, per issue #898.
- The repository's default base branch (`main`) is used; no `base:<branch>` override applies.
- No other concurrent changes to `README.md` are expected to conflict with this single-line append.

## Out of Scope

- Verification of `fabrik:yolo` auto-merge behavior (covered by other smoke-test issues).
- Multi-repo or cross-repo cruise behavior.
- Any change beyond the single appended HTML comment line.

## Source References

- Issue #898 — original definition of the `fabrik:cruise` pipeline contract.
