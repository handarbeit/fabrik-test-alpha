# Feature Specification: e2e smoke full-pipeline (20260713-040227)

**Feature Branch**: `fabrik/issue-3436`
**Created**: 2026-07-13
**Status**: Draft
**Input**: User description: "e2e smoke full-pipeline (20260713-040227)"

## Background

This issue is an end-to-end smoke test for the full Fabrik pipeline on a single repository. Its purpose is to verify that Fabrik can carry an issue through every stage — Specify, Research, Plan, Implement, Review, Validate — to Done with a merged PR, without any cross-repo coordination or issue decomposition. The change itself is a single-line HTML comment appended to the end of `README.md`. Its sole purpose is to produce a trivial, unambiguous diff so the e2e harness can verify full-pipeline plumbing rather than exercise any particular code path.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append full-pipeline smoke marker to README.md (Priority: P1)

The e2e harness expects exactly one new line appended at the very end of `README.md`, using a discriminator unique to this test run. The pipeline must carry this change from Specify through Implement, Review, and Validate to a merged PR without decomposing the work into sub-issues.

**Why this priority**: This is the sole deliverable; nothing else is in scope.

**Independent Test**: Open README.md and confirm that the last line of the file reads exactly `<!-- smoke-full-pipeline-20260713-040227 -->` with no other lines modified.

**Acceptance Scenarios**:

1. **Given** README.md exists, **When** the change is applied, **Then** the last line of the file is exactly `<!-- smoke-full-pipeline-20260713-040227 -->` and no other line is added, removed, or modified.
2. **Given** the issue carries the `fabrik:yolo` label, **When** Validate completes successfully, **Then** the PR is auto-merged and the issue advances to Done.
3. **Given** the single-repo scope, **When** Plan runs, **Then** it does not decompose the issue into sub-issues or child issues.

---

### Edge Cases

- No other files may be modified (except this specification file). Any diff touching files other than `README.md` and this spec file is out of scope and incorrect.
- The comment must be appended at the very end of the file, not inserted elsewhere.
- The marker `<!-- smoke-full-pipeline-20260713-040227 -->` must not be reused from or confused with any prior test's HTML comment.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `README.md` MUST have exactly one new line appended at the very end of the file with the content `<!-- smoke-full-pipeline-20260713-040227 -->`.
- **FR-002**: No file other than `README.md` may be modified by this change (with the exception of this specification file).
- **FR-003**: The implementation MUST NOT decompose this into sub-issues or multiple commits beyond the single-line edit.
- **FR-004**: The pipeline MUST carry this issue through all stages (Research, Plan, Implement, Review, Validate) to a merged PR and Done, given the `fabrik:yolo` label is present.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `git diff HEAD~1 HEAD -- README.md` shows exactly one line appended at the end of the file: `<!-- smoke-full-pipeline-20260713-040227 -->`, with no other changes.
- **SC-002**: The merged PR diff touches only `README.md` (excluding this specification file).
- **SC-003**: The issue reaches Done with the PR merged, without any child issues being spawned.
- **SC-004**: The spec file `specs/3436-e2e-smoke-full-pipeline/spec.md` is committed on the feature branch.

## Assumptions

- `README.md` exists in the repository root.
- No other changes are needed; this is a one-line append to `README.md` (the spec file is committed separately).
- Single-repo scope: no cross-repo work is involved.

## Out of Scope

- Any changes to files other than `README.md` (and this spec file).
- Decomposition into sub-issues.
- Any logic, configuration, or test changes beyond the single marker append.
