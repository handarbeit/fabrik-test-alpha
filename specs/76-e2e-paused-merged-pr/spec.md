# Feature Specification: e2e paused merged-PR recovery (awaiting-review 20260619-202956)

**Feature Branch**: `fabrik/issue-76`
**Created**: 2026-06-19
**Status**: Draft
**Input**: User description: "e2e paused merged-PR recovery (awaiting-review 20260619-202956)"

## Background

This issue is a regression guard for the #874 bug class: a cruise issue whose linked PR is merged externally while the issue is stuck (fabrik:paused + fabrik:awaiting-input, with an optional gate label at Validate). The settle-owner (ADR-056 D2) is expected to detect the merged PR and heal the issue directly to CLOSED without re-invoking the Validate stage.

The change itself is a one-line HTML comment appended to `README.md`. Its sole purpose is to produce a PR that the e2e harness can merge externally while the issue is in the stuck state, then verify the settle-owner fires correctly.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append paused-merged-pr-recovery marker to README.md (Priority: P1)

The e2e harness expects exactly one new line appended at the end of `README.md` with a specific discriminator. The resulting PR is merged externally while the issue is held in the paused/awaiting-input stuck state. The harness then verifies that the settle-owner transitions the issue to CLOSED without Validate being invoked.

**Why this priority**: This is the sole deliverable; nothing else is in scope.

**Independent Test**: Open README.md and confirm the final line reads exactly `<!-- paused-merged-pr-awaiting-review-20260619-202956 -->` with no other changes to the file.

**Acceptance Scenarios**:

1. **Given** README.md exists in the repository root, **When** the change is applied, **Then** the last line of README.md is exactly `<!-- paused-merged-pr-awaiting-review-20260619-202956 -->` and no other line is modified.
2. **Given** the linked PR is merged externally while the issue carries fabrik:paused + fabrik:awaiting-input labels, **When** the settle-owner runs, **Then** the issue is transitioned to CLOSED without the Validate stage being invoked.

---

### Edge Cases

- No other files may be modified (except this specification file). Any diff touching files other than `README.md` and this spec file is out of scope and incorrect.
- The change must be appended at the very end of the file, not inserted at any other position.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `README.md` MUST have exactly one new line appended at the end of the file with the content `<!-- paused-merged-pr-awaiting-review-20260619-202956 -->`.
- **FR-002**: No file other than `README.md` may be modified by this change (with the exception of this specification file).
- **FR-003**: The implementation MUST NOT decompose this into sub-issues or multiple commits beyond the single-line edit.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `git diff HEAD~1 HEAD -- README.md` shows exactly one line added at the end of the file: `<!-- paused-merged-pr-awaiting-review-20260619-202956 -->`, with no other changes.
- **SC-002**: The merged PR diff touches only `README.md` (excluding this specification file).
- **SC-003**: The spec file `specs/76-e2e-paused-merged-pr/spec.md` is committed on the feature branch (FR-002 is satisfied at the spec stage by not modifying any source files here).

## Assumptions

- `README.md` exists in the repository root.
- The e2e harness will merge the linked PR externally while the issue is in the stuck state (fabrik:paused + fabrik:awaiting-input), exercising the settle-owner recovery path.
- No other changes are needed; this is a one-line append to `README.md` (the spec file is committed separately).

## Out of Scope

- Any changes to files other than `README.md` (and this spec file).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes beyond the single marker append.
- Verification of the settle-owner behavior itself — that is exercised by the e2e harness externally, not by this issue's implementation.
