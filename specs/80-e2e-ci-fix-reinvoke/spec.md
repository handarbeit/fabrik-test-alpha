# Feature Specification: e2e ci-fix-reinvoke (20260619-211012)

**Feature Branch**: `fabrik/issue-80`
**Created**: 2026-06-19
**Status**: Implemented
**Input**: User description: "e2e ci-fix-reinvoke (20260619-211012)"

## Background

This issue is an end-to-end regression guard for the CI-fix reinvoke loop (handarbeit/fabrik#900). When a PR's CI check fails after the Implement stage pushes code, Fabrik should detect the failure, re-invoke the Implement stage (the "CI-fix reinvoke"), and allow the stage to push a corrective commit that makes CI pass. A two-commit sequence in `README.md` is the minimal triggering change: the first commit deliberately causes the CI sentinel to fail; the second commit (pushed during the CI-fix reinvoke) satisfies it. The e2e harness verifies that both commits are present and distinct on the final branch.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Initial commit causes CI sentinel to fail (Priority: P1)

The Implement stage pushes exactly one commit adding `<!-- ci-fix-reinvoke-initial -->` to `README.md` immediately after the `# fabrik-test-alpha` heading. The PR body carries the literal `ci-fix-sentinel-required` marker, which causes the test repo's CI sentinel check to fire and fail on this commit.

**Why this priority**: This is the triggering condition for the CI-fix reinvoke path — without it, the reinvoke never fires.

**Independent Test**: After the initial Implement commit, `README.md` contains exactly `<!-- ci-fix-reinvoke-initial -->` on its own line immediately after `# fabrik-test-alpha`, and CI reports failure.

**Acceptance Scenarios**:

1. **Given** `README.md` exists with `# fabrik-test-alpha` as a heading, **When** the initial Implement commit is applied, **Then** `<!-- ci-fix-reinvoke-initial -->` appears on its own line immediately after that heading and no other lines are changed.
2. **Given** the PR body contains the literal string `ci-fix-sentinel-required`, **When** CI runs on the initial commit, **Then** the CI sentinel check fails.

---

### User Story 2 - CI-fix reinvoke pushes a second commit that satisfies CI (Priority: P1)

When the CI-fix reinvoke fires (the stage receives a prompt about CI failure), a second commit adds `<!-- ci-fix-sentinel-satisfied -->` immediately below the first comment. This makes the CI sentinel pass. The two commits must remain distinct — no rebase or squash.

**Why this priority**: This is the behavior under test; the e2e harness verifies the commit count and CI outcome.

**Independent Test**: After both commits, `git log --oneline` on the branch shows exactly two Implement-authored commits, and CI passes on the HEAD commit.

**Acceptance Scenarios**:

1. **Given** the CI-fix reinvoke has fired, **When** the second commit is pushed, **Then** `<!-- ci-fix-sentinel-satisfied -->` appears on its own line immediately below `<!-- ci-fix-reinvoke-initial -->` in `README.md`.
2. **Given** the two commits are present, **When** CI runs on the HEAD commit, **Then** the CI sentinel check passes.
3. **Given** both commits are on the branch, **When** the e2e harness inspects `git log`, **Then** exactly two distinct commits touch `README.md` — the initial commit and the CI-fix commit — with no squash or rebase applied.

---

### Edge Cases

- No other files may be modified by either marker commit (except this specification file, committed separately). The CI sentinel update in `.github/workflows/ci.yml` is a prerequisite commit, not a marker commit, and is not subject to this constraint.
- Plan MUST NOT decompose this into sub-issues; the change is intentionally atomic.
- The two commits must not be squashed or rebased into one — the e2e harness counts them.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The initial Implement commit MUST add exactly one line — `<!-- ci-fix-reinvoke-initial -->` — on its own line immediately after the line containing `# fabrik-test-alpha` in `README.md`, with no other file changes.
- **FR-002**: The PR body MUST contain the literal string `ci-fix-sentinel-required` so the CI sentinel check fires on the initial commit.
- **FR-003**: When the CI-fix reinvoke fires, the second commit MUST add exactly one line — `<!-- ci-fix-sentinel-satisfied -->` — immediately below `<!-- ci-fix-reinvoke-initial -->` in `README.md`, with no other file changes.
- **FR-004**: The two commits MUST remain as distinct, separate commits on the branch — no rebase, squash, or amend that merges them.
- **FR-005**: No file other than `README.md` may be modified by either marker commit (the initial commit and the CI-fix commit). The CI sentinel update and the spec file are separate prerequisite commits and are not subject to this constraint.
- **FR-006**: The Plan stage MUST NOT decompose this into sub-tasks or sub-issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the initial Implement commit, `git diff HEAD~1 HEAD -- README.md` shows exactly one line added — `<!-- ci-fix-reinvoke-initial -->` — immediately after `# fabrik-test-alpha`, with no other changes.
- **SC-002**: The PR body contains the literal string `ci-fix-sentinel-required`.
- **SC-003**: After the CI-fix commit, `git diff HEAD~1 HEAD -- README.md` shows exactly one line added — `<!-- ci-fix-sentinel-satisfied -->` — immediately below `<!-- ci-fix-reinvoke-initial -->`, with no other changes.
- **SC-004**: `git log --oneline` on the branch shows two distinct commits that touch `README.md`.
- **SC-005**: CI passes on the HEAD commit after the CI-fix commit is pushed.

## Assumptions

- `README.md` exists in the repository root and contains a line with `# fabrik-test-alpha`.
- The test repo's CI sentinel check is already configured to fail when the PR body carries `ci-fix-sentinel-required` and pass after `<!-- ci-fix-sentinel-satisfied -->` is added.
- The CI-fix reinvoke behavior under test is already implemented in the Fabrik engine; this issue only provides the triggering change.
- The e2e harness inspects commit count and CI outcome externally.

## Out of Scope

- Any changes to files other than `README.md` (and this spec file).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes unrelated to the CI sentinel update and two-commit marker sequence.
- Any unrelated CI configuration changes. (Updating the `ci-fix-sentinel` job to recognize `<!-- ci-fix-sentinel-satisfied -->` in README.md is in scope — it is the test instrument for this e2e scenario.)
