# Feature Specification: e2e base-branch pipeline (20261017-182025)

**Feature Branch**: `fabrik/issue-8640`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261017-182025 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261017-182025 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end probe of the `base:<branch>` contract: an issue labelled with a non-default base branch must travel the whole pipeline (Specify through Validate) with its worktree, PR target, issue-to-PR linkage and review-gate data all following that base branch. It is regression coverage for handarbeit/fabrik#1046, and it validates the fixes in #1047 (issue↔PR linkage) and #1050 (review-gate data feed).

The code change is deliberately trivial, a single marker line in a single file. The change itself has no product value. Its purpose is to give the pipeline a real, minimal PR to carry on the non-default base, so that any breakage shows up in the pipeline's behaviour and not in the content of the diff. The marker carries a unique run timestamp (`20261017-182025`) so this run's change is distinguishable from earlier runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

A maintainer running the e2e suite wants a PR that appends one marker line to `e2e/markers/base-branch-pipeline.md` and targets the branch `e2e-base-branch-20261017-182025` instead of the repository's usual default.

**Why this priority**: This is the whole deliverable. Without this change there is nothing for the pipeline to carry.

**Independent Test**: Inspect the resulting PR diff against `e2e-base-branch-20261017-182025`. It must show exactly one added line, `<!-- base-branch-pipeline-20261017-182025 -->`, at the end of `e2e/markers/base-branch-pipeline.md`, and no other file changes.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` already exists on the base branch, **When** the change is implemented, **Then** the file's last line is `<!-- base-branch-pipeline-20261017-182025 -->` and all prior content is unchanged.
2. **Given** the file does not exist on the base branch, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the change is complete, **When** the PR diff is inspected, **Then** it touches exactly one file and adds exactly one line.

---

### User Story 2 - Pipeline honours the `base:<branch>` contract (Priority: P1)

The maintainer wants the issue to move through every stage while all branch-sensitive behaviour follows `e2e-base-branch-20261017-182025`, so that the #1047 and #1050 fixes are verified.

**Why this priority**: This is the purpose of the run. The marker line is only the carrier.

**Independent Test**: After the pipeline finishes, check the PR's base branch, the issue↔PR link, and the review-gate behaviour. All must match the non-default base.

**Acceptance Scenarios**:

1. **Given** the issue carries the label `base:e2e-base-branch-20261017-182025`, **When** the Implement stage opens its PR, **Then** the PR targets that branch and not the repository default.
2. **Given** the PR targets a non-default base, **When** the engine looks for the issue's linked PR, **Then** it finds the PR (#1047).
3. **Given** the PR targets a non-default base, **When** the review gate reads review data, **Then** it gets the review data for that PR (#1050).
4. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue auto-advances through the stages, with no auto-merge and no advance to Done at the end of Validate, per cruise semantics.

---

### Edge Cases

- The marker file may already end without a trailing newline. The appended line must still land on its own line, not be joined to the previous content.
- The marker file may already hold marker lines from earlier runs. Those lines must be kept, and this run's line added after them.
- The marker file may be missing from the base branch. It is then created with the marker line as its only content.
- The Plan stage must not decompose the work into child issues, because the change is a single line in a single file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261017-182025 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, the change MUST create it, with the marker line as its content.
- **FR-003**: The change MUST NOT alter, reorder or remove any pre-existing content in that file.
- **FR-004**: The change MUST touch no file other than `e2e/markers/base-branch-pipeline.md`.
- **FR-005**: The work MUST be delivered as a single PR that targets the branch `e2e-base-branch-20261017-182025`.
- **FR-006**: The issue's labels MUST NOT be added to or removed by hand. The filing-time labels `base:e2e-base-branch-20261017-182025` and `fabrik:cruise` stay as they are.
- **FR-007**: The work MUST NOT be decomposed into child issues, and this single-repo run MUST NOT involve any other repository.

### Key Entities

- **Marker file** (`e2e/markers/base-branch-pipeline.md`): An append-only record of e2e runs, one HTML comment line per run.
- **Run marker line**: `<!-- base-branch-pipeline-20261017-182025 -->`, unique to this run.
- **Non-default base branch** (`e2e-base-branch-20261017-182025`): The branch the PR must target, named by the issue's `base:` label.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261017-182025` contains exactly 1 changed file and exactly 1 added line, and that line is the run marker.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261017-182025`, and no commit lands on the repository's usual default branch as part of this run.
- **SC-003**: The engine finds the issue's linked PR at every stage that needs it, with no "no linked PR" or missing-review-data condition caused by the non-default base.
- **SC-004**: The issue reaches the end of Validate without any manual label change, and the run's e2e assertions pass.

## Assumptions

- The repository's default branch, as seen by the engine in this run, is `e2e-base-branch-20261017-182025`, as the stage context states. The target is therefore the same branch whether it is reached through the `base:` label or through the default, and the PR base is checked against that branch's name.
- The marker file uses plain Markdown and HTML comments, so appending a comment line has no visible effect.
- Earlier runs may have left marker lines in the file, and appending after them is the intended behaviour.
- Cruise semantics apply: stages auto-advance, but the PR is not auto-merged by this issue's own Validate stage.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configuration, labels or documentation.
- Editing or removing marker lines from earlier runs.
- Adding tests, tooling or scripts for the marker file.
- Multi-repo scenarios and any other base-branch variants.
- Fixes for any pipeline fault this run uncovers. Those belong in separate issues.

## Source References *(optional)*

- handarbeit/fabrik#1046 — base-branch pipeline regression this run covers
- handarbeit/fabrik#1047 — issue↔PR linkage fix being validated
- handarbeit/fabrik#1050 — review-gate data feed fix being validated
