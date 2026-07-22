# Feature Specification: e2e yolo auto-merge (20260722-221356)

**Feature Branch**: `fabrik/issue-3552`
**Created**: 2026-07-22
**Status**: Implemented
**Input**: User description: "e2e yolo auto-merge (20260722-221356)"

## Background

This issue is an end-to-end regression guard for the GitHub native auto-merge path for yolo issues (handarbeit/fabrik#829). The change itself is a single HTML comment line appended to the end of `README.md`. Its sole purpose is to produce a `fabrik:yolo` PR that drives through the post-Validate convergence flow, so the e2e harness can verify that Fabrik enables GitHub's native auto-merge (applying the `fabrik:auto-merge-enabled` label) rather than falling back to the legacy poll-merge loop.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append auto-merge-yolo marker to README.md (Priority: P1)

The e2e harness expects exactly one new line appended at the very end of `README.md`, using a discriminator unique to this test run. The resulting PR carries the `fabrik:yolo` label. The harness then verifies that Fabrik's post-Validate convergence flow applies the `fabrik:auto-merge-enabled` label and relies on GitHub native auto-merge, instead of running the legacy poll-merge loop.

**Why this priority**: This is the sole deliverable; nothing else is in scope.

**Independent Test**: Open `README.md` and confirm that the last line of the file reads exactly `<!-- auto-merge-yolo-20260722-221356 -->` with no other lines modified.

**Acceptance Scenarios**:

1. **Given** `README.md` exists, **When** the change is applied, **Then** the very last line of the file is exactly `<!-- auto-merge-yolo-20260722-221356 -->` and no other line is added, removed, or modified.
2. **Given** the issue carries the `fabrik:yolo` label, **When** the linked PR reaches Validate completion, **Then** Fabrik applies the `fabrik:auto-merge-enabled` label and enables GitHub native auto-merge on the PR rather than running the legacy poll-merge loop.

---

### Edge Cases

- No file other than `README.md` may be modified (except this specification file). Any diff touching other files is out of scope and incorrect.
- The comment must be appended at the very end of the file, not inserted elsewhere.
- This is a single-repo test; no cross-repo or multi-issue coordination is involved.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `README.md` MUST have exactly one new line appended at the very end of the file with the content `<!-- auto-merge-yolo-20260722-221356 -->`.
- **FR-002**: No file other than `README.md` may be modified by this change (with the exception of this specification file).
- **FR-003**: The implementation MUST NOT decompose this into sub-issues or multiple commits beyond the single-line edit.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `git diff HEAD~1 HEAD -- README.md` shows exactly one line appended at the end of the file: `<!-- auto-merge-yolo-20260722-221356 -->`, with no other changes.
- **SC-002**: The merged PR diff touches only `README.md` (excluding this specification file).
- **SC-003**: At Validate completion, Fabrik applies the `fabrik:auto-merge-enabled` label to the issue/PR and GitHub native auto-merge is enabled on the PR, rather than the legacy poll-merge loop running.
- **SC-004**: The spec file `specs/3552-e2e-yolo-auto-merge/spec.md` is committed on the feature branch.

## Assumptions

- `README.md` exists in the repository root.
- The `fabrik:yolo` label is already applied to this issue (per the issue body), triggering auto-advance and the post-Validate auto-merge convergence flow.
- The e2e harness will verify the auto-merge convergence behavior externally — this issue only needs to produce the correct PR; harness logic is out of scope.

## Out of Scope

- Any changes to files other than `README.md` (and this spec file).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes beyond the single marker insertion.
- Verification of the auto-merge convergence flow's internal logic — that is exercised by the e2e harness externally, not by this issue's implementation.
