# Feature Specification: e2e convergence-race B (20260619-182103)

**Feature Branch**: `fabrik/issue-70`
**Created**: 2026-06-19
**Status**: Draft
**Input**: User description: "e2e convergence-race B (20260619-182103)"

## Background

This issue is one of a deliberately-conflicting pair filed by the e2e harness to deterministically provoke the post-Validate auto-merge race covered by handarbeit/fabrik#829. Each pair member inserts a different discriminator at the same position in README.md, so that when both branches rebase onto main, a true textual conflict is produced. The slow-CI gate (triggered by a marker in the PR body) provides a deterministic 6-minute window during which main can move under the branch, exercising the race condition.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Insert convergence-race marker into README.md (Priority: P1)

The e2e harness expects exactly one new line inserted into README.md immediately after the `# fabrik-test-alpha` heading. This single-line edit, combined with the pair partner's identical-position edit, produces a textual conflict on rebase that the harness uses to verify Fabrik's conflict-resolution and auto-merge behavior.

**Why this priority**: This is the sole deliverable; nothing else is in scope.

**Independent Test**: Open README.md and confirm line 2 reads exactly `<!-- convergence-race-B-20260619-182103 -->` with no other changes to the file.

**Acceptance Scenarios**:

1. **Given** README.md starts with `# fabrik-test-alpha` as its first line, **When** the change is applied, **Then** line 2 of README.md is exactly `<!-- convergence-race-B-20260619-182103 -->` and no other line is modified.
2. **Given** the pair partner's marker (`<!-- convergence-race-A-20260619-182103 -->`) is already present at the same position during rebase, **When** a merge conflict occurs, **Then** the conflict is resolved by keeping both marker lines (in either order) with no other changes.

---

### Edge Cases

- If a rebase conflict is encountered with the pair partner's marker at the same line, both markers must be kept (order does not matter).
- No other files may be modified. Any diff touching files other than README.md is out of scope and incorrect.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: README.md MUST have exactly one new line inserted immediately after the line containing `# fabrik-test-alpha`, with the content `<!-- convergence-race-B-20260619-182103 -->`.
- **FR-002**: No file other than README.md may be modified by this change.
- **FR-003**: The PR body MUST contain the literal string `slow-ci-required` so the test repo's slow-CI gate fires and provides the 6-minute race window.
- **FR-004**: If a rebase conflict is encountered with the pair partner's marker at the same position, the conflict MUST be resolved by retaining both lines (in either order).

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `git diff HEAD~1 HEAD -- README.md` shows exactly one line added: `<!-- convergence-race-B-20260619-182103 -->` immediately after `# fabrik-test-alpha`, and no other changes.
- **SC-002**: `git diff HEAD~1 HEAD --name-only` lists only `README.md`.
- **SC-003**: The linked PR body contains the literal string `slow-ci-required`.
- **SC-004**: The spec file `specs/70-e2e-convergence-race-b/spec.md` is committed on the feature branch (FR-002 is satisfied at the spec stage by not modifying any source files here).

## Assumptions

- README.md exists in the repository root and its first line is `# fabrik-test-alpha`.
- The pair partner issue targets the same line with a different discriminator (`convergence-race-A-*`), producing a genuine textual conflict on rebase.
- No other changes are needed; this is a one-line edit followed by a one-line commit.

## Out of Scope

- Any changes to files other than README.md (and this spec file).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes beyond the single marker insertion.
