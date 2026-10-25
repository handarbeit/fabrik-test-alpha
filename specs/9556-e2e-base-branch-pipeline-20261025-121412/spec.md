# Feature Specification: E2E base-branch pipeline marker (20261025-121412)

**Feature Branch**: `fabrik/issue-9556`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261025-121412 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: Single repo only, targeting the non-default base branch e2e-base-branch-20261025-121412 (already set via the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end test fixture. It exists to exercise the `base:<branch>` pipeline contract: an issue whose work targets a non-default base branch must still flow through every stage, with its pull request correctly linked to the issue (#1047) and the review gate receiving the PR's review data (#1050). This is regression coverage for #1046. The code change is deliberately trivial so that the pipeline behaviour, not the content, is what is being verified.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

A maintainer running the e2e suite files this issue against the branch `e2e-base-branch-20261025-121412` and expects the pipeline to produce a PR containing exactly one added marker line.

**Why this priority**: This is the whole change and the signal the e2e test asserts on.

**Independent Test**: Inspect the PR diff against `e2e-base-branch-20261025-121412`; it contains only the added marker line in `e2e/markers/base-branch-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists, **When** the change is made, **Then** its final line is `<!-- base-branch-pipeline-20261025-121412 -->` and all prior content is unchanged.
2. **Given** the file does not exist, **When** the change is made, **Then** the file is created containing the marker line.
3. **Given** the issue carries the `base:e2e-base-branch-20261025-121412` label, **When** the PR is created, **Then** it targets that branch, not the repository's other branches.

---

### User Story 2 - Pipeline contract holds for a non-default base (Priority: P2)

The pipeline completes with the PR linked to the issue and the review gate seeing the PR's review data.

**Why this priority**: This is the regression being verified (#1047, #1050), though it depends on Story 1 producing a PR.

**Independent Test**: Observe the issue reaching completion with its PR recognized as linked and the review gate evaluating that PR's reviews.

**Acceptance Scenarios**:

1. **Given** a PR targeting the non-default base, **When** the pipeline runs, **Then** the PR is discovered as linked to this issue.
2. **Given** the linked PR, **When** the review gate runs, **Then** it reads the PR's review data rather than treating the PR as absent.

---

### Edge Cases

- The marker file already ends without a trailing newline: the marker must still occupy its own final line.
- The marker line is already present (a re-run): it must not be duplicated.
- Labels on the issue must be left as they are; the change must not add or remove any.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261025-121412 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`, creating the file first if it does not exist.
- **FR-002**: The change MUST NOT modify any other file or any existing content of the marker file.
- **FR-003**: The marker line MUST appear at most once in the file.
- **FR-004**: The work MUST target the base branch `e2e-base-branch-20261025-121412`.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: The change MUST NOT add or remove issue labels.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261025-121412` consists of exactly one file and one added line (or, if the file was new, one file whose only content is the marker line).
- **SC-002**: The PR's base branch is `e2e-base-branch-20261025-121412`.
- **SC-003**: The issue reaches completion through the pipeline with its PR linked to it and no children spawned.

## Assumptions

- The branch `e2e-base-branch-20261025-121412` already exists on the remote.
- The issue's existing labels (`fabrik:cruise`, `base:e2e-base-branch-20261025-121412`) are correct and are not to be changed.
- Whether the marker file already exists is unknown at specification time; either case is covered by FR-001.

## Out of Scope *(optional)*

- Any change to engine, pipeline or test-harness code.
- Any file other than `e2e/markers/base-branch-pipeline.md`.
- Decomposing the work into sub-issues.
