# Feature Specification: e2e yolo auto-merge (20260714-125806)

**Feature Branch**: `fabrik/issue-3466`
**Created**: 2026-07-14
**Status**: Draft
**Input**: User description: "e2e yolo auto-merge (20260714-125806)"

## Background

This issue is a regression guard for the GitHub native auto-merge path for yolo issues (handarbeit/fabrik#829). Fabrik's post-Validate convergence flow is expected to enable GitHub's native auto-merge on the linked PR (applying the `fabrik:auto-merge-enabled` label) rather than running the legacy poll-merge loop. The change itself is a single-line HTML comment appended to the end of `README.md`. Its sole purpose is to produce a yolo-labeled PR that the e2e harness can drive through Validate to confirm the auto-merge convergence path fires correctly.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append auto-merge-yolo marker to README.md (Priority: P1)

The e2e harness expects exactly one new line appended at the very end of `README.md`, using a discriminator unique to this test run. Because the issue carries the `fabrik:yolo` label, Fabrik should auto-advance the issue through all stages without pausing for human review, and at Validate completion enable GitHub native auto-merge on the linked PR rather than running the legacy poll-merge loop.

**Why this priority**: This is the sole deliverable; nothing else is in scope.

**Independent Test**: Open README.md and confirm that the very last line of the file reads exactly `<!-- auto-merge-yolo-20260714-125806 -->` with no other lines modified.

**Acceptance Scenarios**:

1. **Given** README.md exists, **When** the change is applied, **Then** the last line of the file is exactly `<!-- auto-merge-yolo-20260714-125806 -->` and no other line is added, removed, or modified.
2. **Given** the issue carries the `fabrik:yolo` label, **When** Validate completes successfully, **Then** Fabrik applies the `fabrik:auto-merge-enabled` label to the linked PR and enables GitHub native auto-merge instead of running the legacy poll-merge loop.

---

### Edge Cases

- No other files may be modified (except this specification file). Any diff touching files other than `README.md` and this spec file is out of scope and incorrect.
- The comment must be appended at the very end of the file, not inserted elsewhere.
- This issue must not be decomposed into sub-issues or multiple commits — it is a single trivial change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `README.md` MUST have exactly one new line appended at the very end of the file with the content `<!-- auto-merge-yolo-20260714-125806 -->`.
- **FR-002**: No file other than `README.md` may be modified by this change (with the exception of this specification file).
- **FR-003**: The implementation MUST NOT decompose this into sub-issues or multiple commits beyond the single-line edit.
- **FR-004**: The issue MUST remain single-repo scope; no cross-repo coordination is required.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `git diff HEAD~1 HEAD -- README.md` shows exactly one line added at the end of the file: `<!-- auto-merge-yolo-20260714-125806 -->`, with no other changes.
- **SC-002**: The merged PR diff touches only `README.md` (excluding this specification file).
- **SC-003**: The spec file `specs/3466-e2e-yolo-auto-merge/spec.md` is committed on the feature branch.
- **SC-004**: The linked PR carries the `fabrik:auto-merge-enabled` label after Validate completes, and GitHub native auto-merge is enabled on it.

## Assumptions

- `README.md` exists in the repository root.
- This issue only needs to produce the correct yolo PR; verification of the auto-merge convergence flow itself is exercised by the e2e harness externally.
- No other changes are needed; this is a one-line append to `README.md` (the spec file is committed separately).

## Out of Scope

- Any changes to files other than `README.md` (and this spec file).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes beyond the single marker append.
- Verification of the auto-merge convergence logic itself — that is exercised by the e2e harness externally, not by this issue's implementation.
