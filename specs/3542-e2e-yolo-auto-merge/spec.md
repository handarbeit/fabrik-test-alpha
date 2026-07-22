# Feature Specification: e2e yolo auto-merge (20260722-150049)

**Feature Branch**: `fabrik/issue-3542`
**Created**: 2026-07-22
**Status**: Draft
**Input**: User description: "e2e yolo auto-merge (20260722-150049)"

## Background

This issue is an end-to-end regression guard for the GitHub-native auto-merge path used by `fabrik:yolo` issues (handarbeit/fabrik#829). After Validate completes for a yolo-labeled issue, Fabrik's post-Validate convergence flow should enable GitHub's native auto-merge on the linked PR (applying the `fabrik:auto-merge-enabled` label) instead of running the legacy poll-merge loop that repeatedly checks mergeability and merges manually. The change itself is a single-line HTML comment appended to the end of `README.md`. Its sole purpose is to produce a trivial, low-risk yolo PR so the e2e harness can drive it through Validate and verify Fabrik takes the auto-merge-enabled path rather than the legacy poll-merge path.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append auto-merge-yolo marker to README.md (Priority: P1)

The e2e harness expects exactly one new line appended at the very end of `README.md`, using a discriminator unique to this test run. Because the issue carries the `fabrik:yolo` label, the resulting PR should flow through Fabrik's pipeline with auto-advance enabled, and upon Validate completion Fabrik should apply the `fabrik:auto-merge-enabled` label and enable GitHub native auto-merge on the PR rather than polling for mergeability itself.

**Why this priority**: This is the sole deliverable; nothing else is in scope.

**Independent Test**: Open `README.md` and confirm that the final line of the file reads exactly `<!-- auto-merge-yolo-20260722-150049 -->` with no other lines modified.

**Acceptance Scenarios**:

1. **Given** `README.md` exists, **When** the change is applied, **Then** the last line of the file is exactly `<!-- auto-merge-yolo-20260722-150049 -->` and no other line is added, removed, or modified.
2. **Given** the issue carries the `fabrik:yolo` label, **When** the PR is created, **Then** the pipeline auto-advances through stages without requiring manual approval at each stage.
3. **Given** Validate completes successfully, **When** Fabrik's post-Validate convergence flow runs, **Then** the `fabrik:auto-merge-enabled` label is applied and GitHub native auto-merge is enabled on the PR, rather than Fabrik running its legacy poll-merge loop.

---

### Edge Cases

- No other files may be modified (except this specification file). Any diff touching files other than `README.md` and this spec file is out of scope and incorrect.
- The comment must be appended at the very end of the file, not inserted elsewhere (e.g., not after a heading).
- The marker `<!-- auto-merge-yolo-20260722-150049 -->` must not be reused from or confused with any prior test's HTML comment (e.g. the earlier `auto-merge-yolo-20260714-125806` marker already present in the file).
- Plan should NOT decompose this issue into sub-issues or multiple commits.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `README.md` MUST have exactly one new line appended at the very end of the file with the content `<!-- auto-merge-yolo-20260722-150049 -->`.
- **FR-002**: No file other than `README.md` may be modified by this change (with the exception of this specification file).
- **FR-003**: The implementation MUST NOT decompose this into sub-issues or multiple commits beyond the single-line edit.
- **FR-004**: The issue MUST retain the `fabrik:yolo` label through the pipeline so the post-Validate convergence flow is exercised.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `git diff HEAD~1 HEAD -- README.md` shows exactly one line added at the end of the file: `<!-- auto-merge-yolo-20260722-150049 -->`, with no other changes.
- **SC-002**: The merged PR diff touches only `README.md` (excluding this specification file).
- **SC-003**: Upon Validate completion, the linked PR carries the `fabrik:auto-merge-enabled` label and has GitHub native auto-merge enabled, rather than being merged via Fabrik's legacy poll-merge loop.
- **SC-004**: The spec file `specs/3542-e2e-yolo-auto-merge/spec.md` is committed on the feature branch.

## Assumptions

- `README.md` exists in the repository root.
- The `fabrik:yolo` label is already applied to the issue (per the original issue body) and will remain in place through the pipeline.
- The e2e harness will verify the auto-merge-enabled behavior externally (label presence, GitHub auto-merge state) — this issue only needs to produce the correct PR; harness logic is out of scope.
- No other changes are needed; this is a one-line append to `README.md` (the spec file is committed separately).

## Out of Scope

- Any changes to files other than `README.md` (and this spec file).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes beyond the single marker append.
- Verification of the post-Validate convergence flow's internal logic itself — that is exercised by the e2e harness externally, not by this issue's implementation.

## Source References

- handarbeit/fabrik#829 — GitHub native auto-merge path for yolo issues.
