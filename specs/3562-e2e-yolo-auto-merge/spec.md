# Feature Specification: e2e yolo auto-merge (20260722-223810)

**Feature Branch**: `fabrik/issue-3562`
**Created**: 2026-07-22
**Status**: Draft
**Input**: User description: "e2e yolo auto-merge (20260722-223810)"

## Background

This issue is a regression guard for the GitHub native auto-merge path for yolo issues (handarbeit/fabrik#829). Yolo-labeled issues are expected to have their PR converge via GitHub's native auto-merge (Fabrik applying the `fabrik:auto-merge-enabled` label) rather than the legacy poll-merge loop, once the post-Validate convergence flow completes. The change itself is a one-line HTML comment appended to the end of `README.md`. Its sole purpose is to produce a trivial yolo PR so the e2e harness can verify Fabrik enables GitHub native auto-merge for it.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append auto-merge-yolo marker to README.md (Priority: P1)

The e2e harness expects exactly one new line appended at the very end of `README.md`, using a discriminator unique to this test run. The resulting PR is yolo-labeled (`fabrik:yolo`). The harness then verifies that, after Validate completes, Fabrik applies the `fabrik:auto-merge-enabled` label and relies on GitHub's native auto-merge rather than running the legacy poll-merge loop.

**Why this priority**: This is the sole deliverable; nothing else is in scope.

**Independent Test**: Open `README.md` and confirm that the final line of the file reads exactly `<!-- auto-merge-yolo-20260722-223810 -->` with no other lines modified.

**Acceptance Scenarios**:

1. **Given** `README.md` exists, **When** the change is applied, **Then** the very last line of the file is exactly `<!-- auto-merge-yolo-20260722-223810 -->` and no other line is added, removed, or modified.
2. **Given** the PR is created, **When** the PR is inspected, **Then** it carries the `fabrik:yolo` label and is linked to issue #3562 via `Closes #3562`.
3. **Given** Validate completes for this issue, **When** the post-Validate convergence flow runs, **Then** Fabrik applies the `fabrik:auto-merge-enabled` label to the PR instead of running the legacy poll-merge loop.

---

### Edge Cases

- No other files may be modified (except this specification file). Any diff touching files other than `README.md` and this spec file is out of scope and incorrect.
- The comment must be appended at the very end of the file, not inserted elsewhere.
- The marker `<!-- auto-merge-yolo-20260722-223810 -->` must not be reused from or confused with any prior test's HTML comment (earlier runs, e.g. `<!-- auto-merge-yolo-20260722-221356 -->`, already exist in `README.md`).
- This issue is single-repo scope only; no multi-repo or cross-repo auto-merge behavior is exercised.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `README.md` MUST have exactly one new line appended at the very end of the file with the content `<!-- auto-merge-yolo-20260722-223810 -->`.
- **FR-002**: No file other than `README.md` may be modified by this change (with the exception of this specification file).
- **FR-003**: The implementation MUST NOT decompose this into sub-issues or multiple commits beyond the single-line edit.
- **FR-004**: The PR MUST remain yolo-labeled end to end so the post-Validate convergence flow's auto-merge path is exercised.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `git diff HEAD~1 HEAD -- README.md` shows exactly one line added at the end of the file: `<!-- auto-merge-yolo-20260722-223810 -->`, with no other changes.
- **SC-002**: The PR diff touches only `README.md` (excluding this specification file).
- **SC-003**: After Validate completes, the linked PR carries the `fabrik:auto-merge-enabled` label (applied by Fabrik) rather than being merged via the legacy poll-merge loop.
- **SC-004**: The spec file `specs/3562-e2e-yolo-auto-merge/spec.md` is committed on the feature branch.

## Assumptions

- `README.md` exists in the repository root.
- The `fabrik:yolo` label is already applied to this issue (per the original issue body) and remains set throughout the pipeline.
- The e2e harness will verify the auto-merge label and convergence behavior externally — this issue only needs to produce the correct yolo PR; harness logic is out of scope.
- No other changes are needed; this is a one-line append to `README.md` (the spec file is committed separately).

## Out of Scope

- Any changes to files other than `README.md` (and this spec file).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes beyond the single marker append.
- Verification of the post-Validate convergence/auto-merge logic itself — that is exercised by the e2e harness externally, not by this issue's implementation.
- Multi-repo scenarios.

## Source References *(optional)*

- handarbeit/fabrik#829 — GitHub native auto-merge for yolo issues.
