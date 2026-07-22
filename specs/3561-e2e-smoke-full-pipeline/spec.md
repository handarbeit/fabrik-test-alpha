# Feature Specification: e2e smoke full-pipeline (20260722-223810)

**Feature Branch**: `fabrik/issue-3561`
**Created**: 2026-07-22
**Status**: Draft
**Input**: User description: "e2e smoke full-pipeline (20260722-223810)"

## Background

This issue is an end-to-end smoke test of the full Fabrik single-repo pipeline: Specify → Research → Plan → Implement → Review → Validate → Done, with the linked PR merged along the way. The change itself is a single HTML comment line appended to the very end of `README.md`. Its sole purpose is to give the e2e harness a minimal, unambiguous diff to drive through every pipeline stage and confirm the issue reaches Done with a merged PR.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append full-pipeline smoke marker to README.md (Priority: P1)

The e2e harness expects exactly one new line appended at the very end of `README.md`, using a discriminator unique to this test run. The pipeline must carry the change from Specify through Implement, Review, and Validate, ending with the PR merged and the issue moved to Done.

**Why this priority**: This is the sole deliverable; nothing else is in scope.

**Independent Test**: Open README.md and confirm the last line reads exactly `<!-- smoke-full-pipeline-20260722-223810 -->` with no other lines modified.

**Acceptance Scenarios**:

1. **Given** README.md exists, **When** the change is applied, **Then** the very last line of the file is exactly `<!-- smoke-full-pipeline-20260722-223810 -->` and no other line is added, removed, or modified.
2. **Given** the PR is created and passes Review and Validate, **When** the pipeline completes, **Then** the PR is merged and the issue is moved to Done.
3. **Given** the `fabrik:yolo` label is present, **When** each stage completes, **Then** the pipeline auto-advances through all stages without requiring manual approval, and the PR is auto-merged at Validate completion.

---

### Edge Cases

- No other files may be modified (except this specification file). Any diff touching files other than `README.md` and this spec file is out of scope and incorrect.
- The comment must be appended at the very end of the file, not inserted elsewhere (e.g. not after a heading).
- This is a single-repo change — the Plan stage MUST NOT decompose this issue into sub-issues or cross-repo work.
- The marker `<!-- smoke-full-pipeline-20260722-223810 -->` must not be reused from or confused with any prior test's HTML comment.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `README.md` MUST have exactly one new line appended at the very end of the file with the content `<!-- smoke-full-pipeline-20260722-223810 -->`.
- **FR-002**: No file other than `README.md` may be modified by this change (with the exception of this specification file).
- **FR-003**: The implementation MUST NOT decompose this into sub-issues or multiple commits beyond the single-line edit.
- **FR-004**: The issue MUST progress through every pipeline stage (Research, Plan, Implement, Review, Validate) and reach Done with the linked PR merged.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: `git diff HEAD~1 HEAD -- README.md` shows exactly one line added at the end of the file: `<!-- smoke-full-pipeline-20260722-223810 -->`, with no other changes.
- **SC-002**: The merged PR diff touches only `README.md` (excluding this specification file).
- **SC-003**: The issue reaches the Done column with the linked PR merged, driven entirely by the `fabrik:yolo` auto-advance/auto-merge behavior.
- **SC-004**: The spec file `specs/3561-e2e-smoke-full-pipeline/spec.md` is committed on the feature branch.

## Assumptions

- `README.md` exists in the repository root.
- This is a single-repo smoke test; no cross-repo coordination is required.
- The e2e harness will verify full-pipeline completion externally (PR merged, issue Done) — this issue only needs to produce the correct change; harness logic is out of scope.
