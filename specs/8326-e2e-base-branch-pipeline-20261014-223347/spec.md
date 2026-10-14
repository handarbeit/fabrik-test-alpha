# Feature Specification: e2e base-branch pipeline (20261014-223347)

**Feature Branch**: `fabrik/issue-8326`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261014-223347 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: single repo only, targeting the non-default base branch e2e-base-branch-20261014-223347 (already set via the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end probe of the `base:<branch>` pipeline contract: an issue carrying a `base:<branch>` label must be worked on, reviewed and landed against that non-default base branch rather than the repository default. It provides regression coverage for handarbeit/fabrik#1046 and validates two fixes:

- **#1047**: the issue↔PR linkage for a PR that targets a non-default base branch.
- **#1050**: the review-gate data feed for such a PR.

The code change is deliberately trivial so that any failure points at the pipeline's base-branch handling, not at the content of the change. The `base:e2e-base-branch-20261014-223347` label is applied at filing and is part of the scenario under test.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

The pipeline takes the issue through its stages with the target base branch `e2e-base-branch-20261014-223347`. It produces a PR that adds one marker line to `e2e/markers/base-branch-pipeline.md`, and that PR targets the non-default base branch.

**Why this priority**: This is the whole change, and it is the artifact the e2e run asserts on.

**Independent Test**: Inspect the PR diff against `e2e-base-branch-20261014-223347`. It should contain exactly one added line in exactly one file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, **When** the change is made, **Then** the file exists and its last line is `<!-- base-branch-pipeline-20261014-223347 -->`.
2. **Given** the file already exists on the base branch, **When** the change is made, **Then** the marker line is appended at the very end and all existing content is unchanged.
3. **Given** the PR is opened, **When** its base is inspected, **Then** it targets `e2e-base-branch-20261014-223347`.

---

### User Story 2 - Issue↔PR linkage and review-gate feed work for a non-default base (Priority: P2)

The pipeline finds the PR for this issue and feeds its review data into the review gate, even though GitHub creates no Development-panel link for a PR whose base is not the repository default.

**Why this priority**: This validates the #1047 and #1050 fixes, which is the purpose of the run. The marker line only gives the pipeline something to carry.

**Independent Test**: Confirm the pipeline discovers the PR (via `Closes #8326` in its body) and that the later stages see the PR's reviews and CI.

**Acceptance Scenarios**:

1. **Given** the PR targets a non-default base branch and has `Closes #8326` in its body, **When** later stages run, **Then** they resolve the PR as this issue's linked PR.
2. **Given** the PR has reviews or review comments, **When** the review gate is evaluated, **Then** it uses the PR's review data and does not treat the PR as missing.

---

### Edge Cases

- The marker file is absent from the base branch: it is created, containing just the marker line.
- The marker file lacks a trailing newline: the marker still goes on its own final line.
- The base branch is not the repository default: no stage may fall back to the default branch, and the PR must not target the default branch.
- Labels (`fabrik:cruise`, `base:e2e-base-branch-20261014-223347`) must not be added or removed as part of the change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the line `<!-- base-branch-pipeline-20261014-223347 -->` as the final line of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, the change MUST create it; if it exists, its prior content MUST be preserved unchanged.
- **FR-003**: The change MUST consist of exactly one file and one added line, and no other files may be modified.
- **FR-004**: The resulting PR MUST target `e2e-base-branch-20261014-223347` and MUST include `Closes #8326` in its body.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: The work MUST NOT add or remove labels on the issue as part of the change; the existing `base:` label governs the base branch.

### Key Entities

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a test artifact on the non-default base branch.
- **Base branch**: `e2e-base-branch-20261014-223347`, the non-default branch the pipeline must fork from and target.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261014-223347` is exactly one added line in one file.
- **SC-002**: The final line of the marker file on the base branch after landing is `<!-- base-branch-pipeline-20261014-223347 -->`.
- **SC-003**: The pipeline progresses the issue through all stages to completion without a manual linkage or review-gate workaround.
- **SC-004**: No PR or commit from this issue targets the repository default branch.

## Assumptions

- The base branch `e2e-base-branch-20261014-223347` exists on the remote, and the harness set it up as part of the scenario.
- Spec-kit documents such as `specs/8326-*/spec.md` are engine-generated and are not part of the "one file, one line" change counted by FR-003.
- The marker directory `e2e/markers/` is created if missing.

## Out of Scope *(optional)*

- Any change to Fabrik engine code or behavior.
- Any other file, documentation, or test change.
- Adding, removing, or changing labels.
- Multi-repo behavior.
- Decomposing into child issues.
