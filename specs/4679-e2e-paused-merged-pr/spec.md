# Feature Specification: e2e paused merged-PR recovery (awaiting-ci 20260813-030229)

**Feature Branch**: `fabrik/issue-4679`
**Created**: 2026-08-12
**Status**: Draft
**Input**: User description: "e2e paused merged-PR recovery (awaiting-ci 20260813-030229)"

## Background

This issue is an end-to-end regression guard for the #874 bug class: when a cruise issue enters the stuck state (`fabrik:paused` + `fabrik:awaiting-input`), optionally with a gate label present at Validate (e.g. `fabrik:awaiting-ci`), and its linked PR is merged externally, the settle-owner (ADR-056 D2) must detect the merged PR and heal the issue directly to CLOSED — without invoking the Validate stage. A single HTML comment appended to `e2e/markers/paused-merged-pr-recovery.md` is the minimal triggering change; it creates a PR that the e2e harness can merge externally to exercise this recovery path.

`e2e/markers/paused-merged-pr-recovery.md` already exists and carries prior marker lines from earlier instances of this same e2e scenario; this issue appends one more line, it does not create the file from scratch (though Implement must handle the case where the file is ever absent, per the issue body).

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append paused-merged-PR marker to e2e/markers/paused-merged-pr-recovery.md (Priority: P1)

The e2e harness expects exactly one new HTML comment line appended at the very end of `e2e/markers/paused-merged-pr-recovery.md`. The resulting PR is merged externally by the harness while the issue is in the stuck state, verifying that the settle-owner closes the issue without a Validate invocation.

**Why this priority**: This is the sole deliverable; nothing else is in scope.

**Independent Test**: Open `e2e/markers/paused-merged-pr-recovery.md` and confirm the last line reads exactly `<!-- paused-merged-pr-awaiting-ci-20260813-030229 -->` with no other lines changed.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/paused-merged-pr-recovery.md` exists in the repository (creating it first if it does not), **When** the change is applied, **Then** the final line of the file is exactly `<!-- paused-merged-pr-awaiting-ci-20260813-030229 -->` and no other line is modified.
2. **Given** the PR associated with this issue is merged externally while the issue carries `fabrik:paused` + `fabrik:awaiting-input` (optionally plus a gate label applied at Validate), **When** the settle-owner next runs, **Then** the issue transitions to CLOSED without a Validate stage invocation.

---

### Edge Cases

- No other files may be modified (except this specification file).
- Plan must NOT decompose this into sub-issues; the change is intentionally atomic.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/paused-merged-pr-recovery.md` MUST have exactly one new line appended at the very end of the file, with the content `<!-- paused-merged-pr-awaiting-ci-20260813-030229 -->`. If the file does not already exist, it MUST be created first.
- **FR-002**: No file other than `e2e/markers/paused-merged-pr-recovery.md` may be modified by this change (with the exception of this specification file).
- **FR-003**: The Plan stage MUST NOT decompose this into sub-tasks or sub-issues; the change is a single atomic line append.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `git diff HEAD~1 HEAD -- e2e/markers/paused-merged-pr-recovery.md` shows exactly one line added — `<!-- paused-merged-pr-awaiting-ci-20260813-030229 -->` — at the end of the file, with no other changes.
- **SC-002**: The merged PR diff touches only `e2e/markers/paused-merged-pr-recovery.md` (excluding this specification file).

## Assumptions

- `e2e/markers/paused-merged-pr-recovery.md` already exists in the repository (it does, carrying prior marker lines from earlier instances of this scenario); Implement should create it only if it is somehow absent.
- The e2e harness will merge the PR externally while the issue is in the stuck state (`fabrik:paused` + `fabrik:awaiting-input`), optionally with a gate label present at Validate.
- The settle-owner behavior under test is already implemented in the Fabrik engine; this issue only provides the triggering change.

## Out of Scope

- Any changes to files other than `e2e/markers/paused-merged-pr-recovery.md` (and this spec file).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes beyond the single marker append.
- Multi-repo scenarios.

## Source References

- Prior instances of this e2e pattern (targeting `README.md`, the earlier convention for this scenario): `specs/4082-e2e-paused-merged-pr/`, `specs/3464-e2e-paused-merged-pr/`, `specs/3460-e2e-paused-merged-pr/`.
- Prior instances already appending to `e2e/markers/paused-merged-pr-recovery.md`: see the file's existing history (marker lines dated 20260809 through 20260810).
