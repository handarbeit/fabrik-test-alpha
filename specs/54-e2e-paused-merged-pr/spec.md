# Feature Specification: e2e paused merged-PR recovery

**Feature Branch**: `fabrik/issue-54`
**Created**: 2026-06-19
**Status**: Draft
**Input**: User description: "End-to-end regression guard for the #874 bug class (paused item + merged PR recovery via the settle-owner, ADR-056 D2)."

## Background

The #874 bug class describes a scenario where a cruise issue becomes stuck (labelled `fabrik:paused` + `fabrik:awaiting-input`, optionally also holding a gate label at Validate) and its linked PR is then merged externally. The expected recovery path (ADR-056 D2) is for the settle-owner to detect the merged-PR state and heal the issue directly to CLOSED without re-invoking Validate.

This issue creates a minimal, deliberate trigger for that code path: a trivial one-line README change causes Fabrik to open a PR, enter the stuck state, have the PR merged externally, and then rely on the settle-owner to close the issue. The change itself has no functional significance — it exists solely to exercise and guard this recovery path.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Settle-owner heals paused issue after external PR merge (Priority: P1)

A cruise issue's linked PR is merged by a human (or an external bot) while the issue holds `fabrik:paused` + `fabrik:awaiting-input` (and optionally a gate label from Validate). The settle-owner detects this on the next poll cycle and moves the issue to CLOSED without re-running the Validate stage.

**Why this priority**: This is the only scenario in scope. It directly guards the #874 regression: without the fix, issues in this state would remain stuck forever after an external merge.

**Independent Test**: Trigger the trivial README change through the full pipeline up to the stuck state, merge the PR externally, and verify the issue closes without a Validate invocation.

**Acceptance Scenarios**:

1. **Given** a cruise issue with a linked PR and labels `fabrik:paused` + `fabrik:awaiting-input`, **When** the linked PR is merged externally, **Then** the settle-owner closes the issue to CLOSED within one poll cycle without invoking Validate.

2. **Given** a cruise issue with a linked PR and labels `fabrik:paused` + `fabrik:awaiting-input` + a Validate gate label, **When** the linked PR is merged externally, **Then** the settle-owner closes the issue to CLOSED without invoking Validate.

---

### Edge Cases

- The settle-owner must handle the case where the optional gate label (e.g. `fabrik:awaiting-ci`) is also present alongside `fabrik:paused` + `fabrik:awaiting-input`.
- The issue should be closed (not re-queued or re-run) — Validate must NOT be re-invoked.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The implementation MUST append exactly one line — `<!-- paused-merged-pr-awaiting-ci-20260619-154613 -->` — to the very end of `README.md`. No other files may be changed.
- **FR-002**: The Plan stage MUST NOT decompose this into sub-tasks. A single atomic commit for the one-line change is the entire implementation.
- **FR-003**: The change MUST be delivered on branch `fabrik/issue-54` and linked to this issue via a PR that includes `Closes #54`.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The line `<!-- paused-merged-pr-awaiting-ci-20260619-154613 -->` appears at the end of `README.md` on the merged branch.
- **SC-002**: The PR is created, enters the stuck state, is merged externally, and the issue reaches CLOSED without a Validate invocation being recorded.

## Assumptions

- The settle-owner and ADR-056 D2 recovery logic already exist in the engine; this issue exercises them, it does not implement them.
- `fabrik:cruise` label is set on the issue, enabling cruise-mode behavior throughout the pipeline.
- "Plan should NOT decompose" means the Implement stage receives a single-task checklist with one commit action.

## Out of Scope

- Any changes beyond the single appended comment line in `README.md`.
- Implementing or modifying the settle-owner logic (that is covered by the referenced ADR/bug fix).
- Multi-repo scenarios.
