# Feature Specification: e2e conjunctive-ci-review-gate (20260620-151932)

**Feature Branch**: `fabrik/issue-106`
**Created**: 2026-06-20
**Status**: Draft
**Input**: User description: "e2e conjunctive-ci-review-gate (20260620-151932)"

## Background

This issue is a regression guard for the conjunctive CI∧review gate (handarbeit/fabrik#895, ADR-056 D2). The gate requires both the CI gate and the review gate to be satisfied before Validate advances the issue to Done. The change itself is a one-line HTML comment inserted into `README.md`, immediately after the `# fabrik-test-alpha` heading. Its sole purpose is to produce a PR that triggers the slow-ci-required check (~10 minutes) so the e2e harness can verify that the issue is held at the conjunctive gate until both CI and review requirements are met.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Insert conjunctive-ci-review-gate marker in README.md (Priority: P1)

The e2e harness expects exactly one new line inserted in `README.md` immediately after the line containing `# fabrik-test-alpha`, using the discriminator unique to this test run. The resulting PR must carry the `slow-ci-required` marker in its body so the slow-gate CI required check fires. The harness then verifies that the conjunctive CI∧review gate holds the issue at Validate until both CI passes and the review gate clears.

**Why this priority**: This is the sole deliverable; nothing else is in scope.

**Independent Test**: Open README.md and confirm that the line immediately after `# fabrik-test-alpha` reads exactly `<!-- conjunctive-ci-review-gate-test -->` with no other lines modified.

**Acceptance Scenarios**:

1. **Given** README.md exists with a line containing `# fabrik-test-alpha`, **When** the change is applied, **Then** the very next line is exactly `<!-- conjunctive-ci-review-gate-test -->` and no other line is added, removed, or modified.
2. **Given** the PR is created, **When** the PR body is inspected, **Then** it contains the literal text `slow-ci-required` and does NOT contain `ci-fix-sentinel-required`.
3. **Given** the PR is open with the slow-gate CI check running, **When** the conjunctive gate evaluates, **Then** the issue remains at Validate until both CI passes and the review gate is satisfied.

---

### Edge Cases

- No other files may be modified (except this specification file). Any diff touching files other than `README.md` and this spec file is out of scope and incorrect.
- The comment must be inserted immediately after the `# fabrik-test-alpha` heading line, not appended at the end of the file and not placed anywhere else.
- The marker `<!-- conjunctive-ci-review-gate-test -->` must not be reused from or confused with any prior test's HTML comment.
- The PR body MUST NOT include `ci-fix-sentinel-required` — only `slow-ci-required`.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `README.md` MUST have exactly one new line inserted immediately after the line containing `# fabrik-test-alpha` with the content `<!-- conjunctive-ci-review-gate-test -->`.
- **FR-002**: No file other than `README.md` may be modified by this change (with the exception of this specification file).
- **FR-003**: The PR body MUST contain the literal text `slow-ci-required` to trigger the slow-gate CI required check.
- **FR-004**: The PR body MUST NOT contain `ci-fix-sentinel-required`.
- **FR-005**: The implementation MUST NOT decompose this into sub-issues or multiple commits beyond the single-line edit.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `git diff HEAD~1 HEAD -- README.md` shows exactly one line added immediately after `# fabrik-test-alpha`: `<!-- conjunctive-ci-review-gate-test -->`, with no other changes.
- **SC-002**: The merged PR diff touches only `README.md` (excluding this specification file).
- **SC-003**: The PR body contains `slow-ci-required` and does not contain `ci-fix-sentinel-required`.
- **SC-004**: The spec file `specs/106-e2e-conjunctive-ci-review/spec.md` is committed on the feature branch.

## Assumptions

- `README.md` exists in the repository root and contains a line with `# fabrik-test-alpha`.
- The test repo's slow-gate CI check is enrolled as a required check and fires when the PR body contains `slow-ci-required`.
- The e2e harness will verify conjunctive gate behaviour externally — this issue only needs to produce the correct PR; harness logic is out of scope.
- No other changes are needed; this is a one-line insertion to `README.md` (the spec file is committed separately).

## Out of Scope

- Any changes to files other than `README.md` (and this spec file).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes beyond the single marker insertion.
- Verification of the conjunctive gate logic itself — that is exercised by the e2e harness externally, not by this issue's implementation.
