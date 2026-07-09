# Feature Specification: e2e cruise full pipeline

**Feature Branch**: `fabrik/issue-3371`
**Created**: 2026-07-09
**Status**: Draft
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898)."

## Background

Fabrik's `fabrik:cruise` label enables auto-advance through all pipeline stages (Specify → Research → Plan → Implement → Review → Validate) without auto-merging the PR at the end. This is distinct from `fabrik:yolo`, which both auto-advances and enables GitHub native auto-merge on Validate completion. Issue #898 introduced this cruise mode contract; this issue exercises it end-to-end to verify the engine honours the correct semantics: auto-advance fires at every stage, but the PR is deliberately left open for a human to merge.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Cruise auto-advances to Validate-complete without merging PR (Priority: P1)

A developer monitoring the project board sees this issue advance automatically through all pipeline stages without any human intervention at stage boundaries. After Validate completes, the linked PR exists on the repo but has NOT been merged — it is left open for human review and merge. Once a human merges the PR, the issue closes via the "Closes #3371" link in the PR body.

**Why this priority**: This is the sole purpose of the issue — verifying that cruise mode auto-advances through all stages but stops short of auto-merging, and that post-merge issue closure works correctly.

**Independent Test**: After the issue moves to the Validate-complete state, confirm the linked PR is open (not merged) on `handarbeit/fabrik-test-alpha`. Merge it manually and confirm the issue closes and moves to Done.

**Acceptance Scenarios**:

1. **Given** the issue at Specify with `fabrik:cruise`, **When** Fabrik processes it, **Then** the issue advances to Research without requiring human input
2. **Given** the issue in the pipeline, **When** Implement runs, **Then** a PR is created on branch `fabrik/issue-3371` with `Closes #3371` in the body
3. **Given** the PR created by Implement, **When** Validate completes, **Then** the PR remains open — Fabrik does NOT enable auto-merge and does NOT merge the PR directly
4. **Given** the open PR after Validate-complete, **When** a human merges the PR, **Then** the issue closes automatically via the GitHub "Closes" link and moves to Done
5. **Given** `main` after the PR merges, **When** the last line of `README.md` is read, **Then** it is exactly `<!-- cruise-pipeline-20260709-142728 -->`

---

### Edge Cases

- README.md may or may not end with a trailing newline; the appended line must be on its own line regardless
- The Plan stage must not decompose into sub-issues — this is a single-repo, single-file change

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Implement MUST append exactly `<!-- cruise-pipeline-20260709-142728 -->` as the final line of `README.md`
- **FR-002**: No other source or configuration files MUST be modified
- **FR-003**: The Plan stage MUST NOT spawn sub-issues or decompose across repos
- **FR-004**: The PR body MUST contain `Closes #3371` so the issue closes when the PR is merged
- **FR-005**: After Validate completes, the PR MUST remain open — Fabrik MUST NOT enable GitHub auto-merge or merge the PR directly

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue auto-advances through all stages (Specify → Research → Plan → Implement → Review → Validate) without human intervention at stage boundaries
- **SC-002**: After Validate completes, the linked PR is open and NOT merged
- **SC-003**: After a human merges the PR, the issue closes and moves to Done
- **SC-004**: `README.md` on `main` ends with the line `<!-- cruise-pipeline-20260709-142728 -->` after merge
- **SC-005**: No other files (excluding this specification) are changed in the merged PR diff

## Assumptions

- The `fabrik:cruise` label is set on the issue, enabling auto-advance but NOT auto-merge
- This is a single-repo operation — `handarbeit/fabrik-test-alpha` only; no cross-repo work
- CI on this repo runs `go build ./...` and `go test ./...`; appending a comment to README.md does not affect either
- The issue closes via GitHub's "Closes #N" mechanism when the PR is merged by a human

## Out of Scope

- Any changes beyond appending the single comment line to `README.md`
- Cross-repo sub-issue spawning or decomposition
- Updating any Go source files, tests, or configuration
- Auto-merging the PR (that is the `fabrik:yolo` contract, not cruise)

## Source References

- `README.md` — the only file modified by this issue
- Fabrik issue #898 — introduced the `fabrik:cruise` pipeline contract
