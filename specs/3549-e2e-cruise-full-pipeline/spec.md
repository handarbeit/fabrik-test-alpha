# Feature Specification: e2e cruise full pipeline (20260722-151055)

**Feature Branch**: `fabrik/issue-3549`
**Created**: 2026-07-22
**Status**: Draft
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898)."

## Background

Fabrik's `fabrik:cruise` label drives issues through Research → Plan → Implement → Review → Validate automatically, but does NOT auto-merge the PR when Validate completes. The PR remains open for a human to merge; closing the PR (via merge) then moves the issue to Done.

This issue is a recurrent single-repo end-to-end smoke test for the cruise contract: it verifies that Fabrik auto-advances through all pipeline stages to Validate-complete without merging, that the PR remains open after Validate, and that merging the PR by a human correctly closes the issue and moves it to Done. The change is the simplest possible — one line appended to `README.md` — to keep noise minimal.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Full pipeline auto-advances to Validate-complete without merging (Priority: P1)

A developer monitoring the Fabrik project board sees this issue advance automatically through all pipeline stages (Specify → Research → Plan → Implement → Review → Validate) without any human intervention and without the PR being merged. After Validate completes, the PR is still open and the issue is in the Validate column, awaiting a human merge.

**Why this priority**: This is the core cruise contract: auto-advance through all stages, but stop short of auto-merge. If this fails, the cruise feature is broken.

**Independent Test**: After the issue reaches the Validate column, confirm on GitHub that the linked PR is still open (not merged), and that the README change is present on the PR branch.

**Acceptance Scenarios**:

1. **Given** the issue at Specify with `fabrik:cruise`, **When** Fabrik processes it, **Then** the issue advances to Research without requiring human input
2. **Given** the issue in the pipeline, **When** Implement runs, **Then** a draft PR is created on branch `fabrik/issue-3549` with `Closes #3549` in the body
3. **Given** the Validate stage is complete, **When** the stage signals `FABRIK_STAGE_COMPLETE`, **Then** the PR remains open (not auto-merged)
4. **Given** the PR still open after Validate, **When** a human merges it, **Then** the issue moves to Done

---

### User Story 2 - Single-line README change only (Priority: P2)

The Implement stage appends exactly one HTML comment line to `README.md` and no other file is modified, keeping the PR diff minimal and the test focused.

**Why this priority**: A clean single-file diff makes it easy to confirm no unintended changes were introduced and keeps the smoke test noise-free.

**Independent Test**: Inspect the PR diff — it must show exactly one added line in `README.md` and no other file changes (excluding this spec file committed in Specify).

**Acceptance Scenarios**:

1. **Given** the PR created by Implement, **When** the diff is inspected, **Then** the only changed file (excluding spec) is `README.md` with one line appended
2. **Given** `main` after the PR is merged, **When** the last line of `README.md` is read, **Then** it is exactly `<!-- cruise-pipeline-20260722-151055 -->`

---

### Edge Cases

- `README.md` may or may not end with a trailing newline; the appended line must appear on its own line regardless
- The Plan stage must not decompose into sub-issues — this is a single-repo, single-file change

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Implement MUST append exactly `<!-- cruise-pipeline-20260722-151055 -->` as the final line of `README.md`
- **FR-002**: No other source or configuration files MUST be modified
- **FR-003**: The Plan stage MUST NOT spawn sub-issues or decompose across repos
- **FR-004**: The PR body MUST contain `Closes #3549` so Fabrik can discover it
- **FR-005**: After Validate completes, the PR MUST remain open — Fabrik MUST NOT auto-merge it
- **FR-006**: After a human merges the PR, the issue MUST move to Done

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Validate-complete on the project board without any human stage-advance action
- **SC-002**: The linked PR is still open (not merged) after Validate completes
- **SC-003**: After a human merges the PR, the issue moves to Done
- **SC-004**: `README.md` on `main` ends with the line `<!-- cruise-pipeline-20260722-151055 -->` after merge
- **SC-005**: No other files (excluding this specification) are changed in the merged PR diff

## Assumptions

- The `fabrik:cruise` label is set on the issue, enabling auto-advance through all stages but NOT auto-merge
- This is a single-repo operation — `handarbeit/fabrik-test-alpha` only; no cross-repo work
- CI on this repo does not block on a README-only change
- A human reviewer will manually merge the PR after Validate completes to close the issue

## Out of Scope

- Any changes beyond appending the single comment line to `README.md`
- Cross-repo sub-issue spawning or decomposition
- Updating any Go source files, tests, or configuration
- Auto-merge behavior (that is the `fabrik:yolo` contract, not cruise)

## Source References

- `README.md` — the only file modified by this issue
- Issue #3485 — earlier analogous cruise smoke test run
