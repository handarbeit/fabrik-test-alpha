# Feature Specification: e2e base-branch pipeline (20260724-064143)

**Feature Branch**: `fabrik/issue-3614`
**Created**: 2026-07-24
**Status**: Approved
**Input**: User description: "e2e base-branch pipeline (20260724-064143)"

## Background

This issue is a regression guard for Fabrik's `base:<branch>` (non-default base branch) pipeline contract, covering handarbeit/fabrik#1046 and validating the fixes from #1047 (issue↔PR linkage when the base branch is not the repository default) and #1050 (review-gate data feed under a non-default base branch). The issue carries the `base:e2e-base-branch-20260724-064143` label, set at filing, which directs Fabrik to fork from, rebase onto, and target PRs at that branch instead of `main`. The change itself is a single-line HTML comment appended to the end of `README.md`. Its sole purpose is to produce a base-branch-targeted PR that the e2e harness can drive through the full pipeline to confirm issue↔PR linkage and the review-gate data feed work correctly against a non-default base branch.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append base-branch-pipeline marker to README.md (Priority: P1)

The e2e harness expects exactly one new line appended at the very end of `README.md`, using a discriminator unique to this test run. Because the issue carries the `base:e2e-base-branch-20260724-064143` label, Fabrik should fork the worktree from, rebase onto, and target the PR at `e2e-base-branch-20260724-064143` instead of the repository default branch (`main`). Because the issue also carries `fabrik:cruise`, Fabrik should auto-advance through all stages without pausing for human review, but must not auto-merge the PR or move the issue to Done at Validate completion.

**Why this priority**: This is the sole deliverable; nothing else is in scope.

**Independent Test**: Open README.md and confirm that the very last line of the file reads exactly `<!-- base-branch-pipeline-20260724-064143 -->` with no other lines modified, and that the linked PR targets `e2e-base-branch-20260724-064143` rather than `main`.

**Acceptance Scenarios**:

1. **Given** README.md exists, **When** the change is applied, **Then** the last line of the file is exactly `<!-- base-branch-pipeline-20260724-064143 -->` and no other line is added, removed, or modified.
2. **Given** the issue carries the `base:e2e-base-branch-20260724-064143` label, **When** the Implement stage creates a PR, **Then** the PR is opened against base branch `e2e-base-branch-20260724-064143`, not `main`, and correctly links back to issue #3614 via `Closes #3614`.
3. **Given** the issue carries the `fabrik:cruise` label, **When** each stage completes, **Then** Fabrik auto-advances to the next stage without pausing for human review, and does not auto-merge the PR or move the issue to Done at Validate completion.

---

### Edge Cases

- No other files may be modified (except this specification file). Any diff touching files other than `README.md` and this spec file is out of scope and incorrect.
- The comment must be appended at the very end of the file, not inserted elsewhere.
- This issue must not be decomposed into sub-issues or multiple commits — it is a single trivial change.
- The `base:<branch>` and `fabrik:cruise` labels are set at filing and must not be added or removed during this issue's lifecycle.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `README.md` MUST have exactly one new line appended at the very end of the file with the content `<!-- base-branch-pipeline-20260724-064143 -->`.
- **FR-002**: No file other than `README.md` may be modified by this change (with the exception of this specification file).
- **FR-003**: The implementation MUST NOT decompose this into sub-issues or multiple commits beyond the single-line edit.
- **FR-004**: The issue MUST remain single-repo scope; no cross-repo coordination is required.
- **FR-005**: The linked PR MUST target the base branch `e2e-base-branch-20260724-064143`, not the repository default branch `main`.
- **FR-006**: The linked PR body MUST contain `Closes #3614` so that issue↔PR linkage resolves correctly against the non-default base branch.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `git diff HEAD~1 HEAD -- README.md` shows exactly one line added at the end of the file: `<!-- base-branch-pipeline-20260724-064143 -->`, with no other changes.
- **SC-002**: The PR diff touches only `README.md` (excluding this specification file).
- **SC-003**: The spec file `specs/3614-e2e-base-branch-pipeline/spec.md` is committed on the feature branch.
- **SC-004**: The linked PR's base branch is `e2e-base-branch-20260724-064143`, confirming correct issue↔PR linkage and review-gate data feed under a non-default base branch.

## Assumptions

- `README.md` exists in the repository root.
- The base branch `e2e-base-branch-20260724-064143` already exists on the remote (created as part of test setup for this e2e run).
- This issue only needs to produce the correct base-branch-targeted PR; verification of the linkage and review-gate behavior itself is exercised by the e2e harness externally.
- No other changes are needed; this is a one-line append to `README.md` (the spec file is committed separately).

## Out of Scope

- Any changes to files other than `README.md` (and this spec file).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes beyond the single marker append.
- Verification of the issue↔PR linkage and review-gate data feed logic itself — that is exercised by the e2e harness externally, not by this issue's implementation.
- Adding or removing the `base:<branch>` or `fabrik:cruise` labels — both are set at filing and must remain untouched.
