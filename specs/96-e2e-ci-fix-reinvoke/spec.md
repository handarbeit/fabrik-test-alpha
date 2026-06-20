# Feature Specification: e2e ci-fix-reinvoke (20260620-002323)

**Feature Branch**: `fabrik/issue-96`
**Created**: 2026-06-19
**Status**: Draft
**Input**: User description: "e2e ci-fix-reinvoke (20260620-002323)"

## Background

This issue is an end-to-end regression guard for the CI-fix reinvoke loop (handarbeit/fabrik#900). When a PR's CI check fails after the Implement stage pushes code, Fabrik should detect the failure, re-invoke the Implement stage (the "CI-fix reinvoke"), and allow the stage to push a corrective commit that makes CI pass. The triggering sequence is: an initial commit adds an HTML comment to `README.md` (deliberately causing the CI sentinel to fail); a second commit — pushed only during the CI-fix reinvoke — creates a new `SENTINEL_FIX` file that satisfies the sentinel. The e2e harness verifies that both commits are present and distinct on the final branch.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Initial commit causes CI sentinel to fail (Priority: P1)

The Implement stage pushes exactly one commit adding `<!-- ci-fix-reinvoke-initial -->` to `README.md` immediately after the `# fabrik-test-alpha` heading. The PR body carries the literal `ci-fix-sentinel-required` marker, which causes the test repo's CI sentinel check to fire and fail on this commit. No other files are changed on this commit.

**Why this priority**: This is the triggering condition for the CI-fix reinvoke path — without it, the reinvoke never fires.

**Independent Test**: After the initial Implement commit, `README.md` contains `<!-- ci-fix-reinvoke-initial -->` on its own line immediately after `# fabrik-test-alpha`, and CI reports failure.

**Acceptance Scenarios**:

1. **Given** `README.md` exists with `# fabrik-test-alpha` as a heading, **When** the initial Implement commit is applied, **Then** `<!-- ci-fix-reinvoke-initial -->` appears on its own line immediately after that heading and no other lines are changed.
2. **Given** the PR body contains the literal string `ci-fix-sentinel-required`, **When** CI runs on the initial commit, **Then** the CI sentinel check fails.

---

### User Story 2 - CI-fix reinvoke pushes a second commit that satisfies CI (Priority: P1)

When the CI-fix reinvoke fires (the stage receives a prompt about CI failure), a second commit creates a new file `SENTINEL_FIX` at the repository root containing exactly `ci-fix-satisfied`. This makes the CI sentinel pass. The two commits must remain distinct — no rebase, squash, or amend.

**Why this priority**: This is the behavior under test; the e2e harness verifies the commit count and CI outcome.

**Independent Test**: After both commits, `git log --oneline` on the branch shows exactly two Implement-authored commits, `SENTINEL_FIX` exists at the repo root with content `ci-fix-satisfied`, and CI passes on the HEAD commit.

**Acceptance Scenarios**:

1. **Given** the CI-fix reinvoke has fired, **When** the second commit is pushed, **Then** a file named exactly `SENTINEL_FIX` exists at the repository root containing exactly one line: `ci-fix-satisfied`.
2. **Given** the `SENTINEL_FIX` file is present, **When** CI runs on the HEAD commit, **Then** the CI sentinel check passes.
3. **Given** both commits are on the branch, **When** the e2e harness inspects `git log`, **Then** exactly two distinct commits are present — the initial README.md commit and the CI-fix SENTINEL_FIX commit — with no squash or rebase applied.

---

### Edge Cases

- The `SENTINEL_FIX` file MUST NOT be created on the initial commit — only on the CI-fix reinvoke.
- No file under `.github/` may be modified by either commit. The CI workflow is immutable test infrastructure.
- The two commits must not be squashed, rebased, or amended into one — the e2e harness counts them.
- Plan MUST NOT decompose this into sub-issues; the change is intentionally atomic.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The initial Implement commit MUST add exactly one line — `<!-- ci-fix-reinvoke-initial -->` — on its own line immediately after the line containing `# fabrik-test-alpha` in `README.md`, with no other file changes.
- **FR-002**: The PR body MUST contain the literal string `ci-fix-sentinel-required` so the CI sentinel check fires on the initial commit.
- **FR-003**: No file under `.github/` may be modified by either the initial commit or the CI-fix commit. The CI workflow is immutable test infrastructure.
- **FR-004**: The CI-fix reinvoke commit MUST create a file named exactly `SENTINEL_FIX` at the repository root, containing exactly one line: `ci-fix-satisfied`. No other files may be modified by this commit.
- **FR-005**: The two commits (initial + CI-fix) MUST remain as distinct, separate commits on the branch — no rebase, squash, or amend that merges them.
- **FR-006**: The Plan stage MUST NOT decompose this into sub-tasks or sub-issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the initial Implement commit, `git diff HEAD~1 HEAD -- README.md` shows exactly one line added — `<!-- ci-fix-reinvoke-initial -->` — immediately after `# fabrik-test-alpha`, with no other changes.
- **SC-002**: The PR body contains the literal string `ci-fix-sentinel-required`.
- **SC-003**: CI fails on the initial commit (sentinel fires as expected).
- **SC-004**: After the CI-fix commit, `SENTINEL_FIX` exists at the repo root containing exactly `ci-fix-satisfied`.
- **SC-005**: `git log --oneline` on the branch shows two distinct commits — one touching `README.md`, one creating `SENTINEL_FIX`.
- **SC-006**: CI passes on the HEAD commit after the CI-fix commit is pushed.

## Assumptions

- `README.md` exists in the repository root and contains a line with `# fabrik-test-alpha`.
- The test repo's CI sentinel check is already configured to fail when the PR body carries `ci-fix-sentinel-required` and to pass when `SENTINEL_FIX` is present with content `ci-fix-satisfied`.
- The CI-fix reinvoke behavior under test is already implemented in the Fabrik engine; this issue only provides the triggering change.
- The e2e harness inspects commit count and CI outcome externally.

## Out of Scope

- Any changes to files other than `README.md` (initial commit) and `SENTINEL_FIX` (CI-fix commit).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes in `.github/` or the Fabrik engine itself.
