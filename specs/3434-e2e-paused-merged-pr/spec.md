# Feature Specification: e2e paused merged-PR recovery (awaiting-ci 20260713-035959)

**Feature Branch**: `fabrik/issue-3434`
**Created**: 2026-07-13
**Status**: Draft
**Input**: User description: "e2e paused merged-PR recovery (awaiting-ci 20260713-035959)"

## Background

This issue is an end-to-end regression guard for the #874 bug class: when a cruise issue enters the stuck state (`fabrik:paused` + `fabrik:awaiting-input`) with the `fabrik:awaiting-ci` gate label present at Validate, and its linked PR is merged externally, the settle-owner (ADR-056 D2) must detect the merged PR and heal the issue directly to CLOSED — without Validate having been (re-)invoked. A single HTML comment appended to `README.md` is the minimal triggering change; it creates a PR that the e2e harness can merge externally to exercise the awaiting-ci gate-label recovery path.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append paused-merged-PR marker to README.md (Priority: P1)

The e2e harness expects exactly one new HTML comment line appended at the very end of `README.md`. The resulting PR is merged externally by the harness while the issue is in the stuck state with the `fabrik:awaiting-ci` gate label present, verifying that the settle-owner closes the issue without Validate having been invoked.

**Why this priority**: This is the sole deliverable; nothing else is in scope.

**Independent Test**: Open `README.md` and confirm the last line reads exactly `<!-- paused-merged-pr-awaiting-ci-20260713-035959 -->` with no other lines changed.

**Acceptance Scenarios**:

1. **Given** `README.md` exists in the repository root, **When** the change is applied, **Then** the final line of `README.md` is exactly `<!-- paused-merged-pr-awaiting-ci-20260713-035959 -->` and no other line is modified.
2. **Given** the PR associated with this issue is merged externally while the issue carries `fabrik:paused` + `fabrik:awaiting-input` + `fabrik:awaiting-ci`, **When** the settle-owner next runs, **Then** the issue transitions to CLOSED without Validate having been invoked.

---

### Edge Cases

- No other files may be modified (except this specification file).
- Plan must NOT decompose this into sub-issues; the change is intentionally atomic.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `README.md` MUST have exactly one new line appended at the very end of the file, with the content `<!-- paused-merged-pr-awaiting-ci-20260713-035959 -->`.
- **FR-002**: No file other than `README.md` may be modified by this change (with the exception of this specification file).
- **FR-003**: The Plan stage MUST NOT decompose this into sub-tasks or sub-issues; the change is a single atomic line append.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `git diff HEAD~1 HEAD -- README.md` shows exactly one line added — `<!-- paused-merged-pr-awaiting-ci-20260713-035959 -->` — at the end of the file, with no other changes.
- **SC-002**: The merged PR diff touches only `README.md` (excluding this specification file).

## Assumptions

- `README.md` exists in the repository root.
- The e2e harness will merge the PR externally while the issue is in the stuck state (`fabrik:paused` + `fabrik:awaiting-input`) with the `fabrik:awaiting-ci` gate label present at Validate.
- The settle-owner behavior under test is already implemented in the Fabrik engine; this issue only provides the triggering change.

## Out of Scope

- Any changes to files other than `README.md` (and this spec file).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes beyond the single marker append.
- Gate-label variants other than awaiting-ci (no-gate-label and awaiting-review covered by other issues).
