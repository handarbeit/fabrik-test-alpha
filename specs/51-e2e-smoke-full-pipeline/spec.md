# Feature Specification: e2e smoke full-pipeline (20260619-154613)

**Feature Branch**: `fabrik/issue-51`
**Created**: 2026-06-19
**Status**: Draft
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR."

## Background

Fabrik's pipeline drives issues through Research → Plan → Implement → Review → Validate → Done stages, creating a PR and merging it along the way. This issue is a recurrent single-repo end-to-end smoke test: it verifies the entire pipeline executes correctly using the simplest possible change (one line appended to README.md), so that pipeline regressions are immediately visible without noise from complex implementation work.

The `fabrik:yolo` label is present, which means Fabrik will auto-advance through all stages and auto-merge the PR when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Full pipeline executes end-to-end (Priority: P1)

A developer monitoring the Fabrik project board sees this issue advance automatically through all pipeline stages (Specify → Research → Plan → Implement → Review → Validate → Done), a PR is created and merged, and the single-line README change appears on `main`.

**Why this priority**: This is the sole purpose of the issue — smoke-testing the full pipeline path. There is no secondary story.

**Independent Test**: After the issue moves to Done, check out `main` and confirm the final line of README.md is `<!-- smoke-full-pipeline-20260619-154613 -->`.

**Acceptance Scenarios**:

1. **Given** the issue at Specify, **When** Fabrik processes it with `fabrik:yolo`, **Then** the issue advances to Research without requiring human input
2. **Given** the issue in the pipeline, **When** Implement runs, **Then** a PR is created on branch `fabrik/issue-51` with `Closes #51` in the body
3. **Given** the PR created by Implement, **When** Validate completes, **Then** the PR is auto-merged and the issue moves to Done
4. **Given** `main` after the PR merges, **When** the last line of `README.md` is read, **Then** it is exactly `<!-- smoke-full-pipeline-20260619-154613 -->`

---

### Edge Cases

- README.md may or may not end with a trailing newline; the appended line must be on its own line regardless
- The Plan stage must not decompose into sub-issues — this is a single-repo, single-file change

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Implement MUST append exactly `<!-- smoke-full-pipeline-20260619-154613 -->` as the final line of `README.md`
- **FR-002**: No other source or configuration files MUST be modified
- **FR-003**: The Plan stage MUST NOT spawn sub-issues or decompose across repos
- **FR-004**: The PR body MUST contain `Closes #51` so Fabrik can discover and merge it

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches the Done column on the project board
- **SC-002**: The linked PR is merged into `main`
- **SC-003**: `README.md` on `main` ends with the line `<!-- smoke-full-pipeline-20260619-154613 -->`
- **SC-004**: No other files (excluding this specification) are changed in the merged PR diff

## Assumptions

- The `fabrik:yolo` label is set on the issue, enabling auto-advance and auto-merge through all stages
- This is a single-repo operation — `handarbeit/fabrik-test-alpha` only; no cross-repo work
- CI on this repo does not block on a README-only change

## Out of Scope

- Any changes beyond appending the single comment line to `README.md`
- Cross-repo sub-issue spawning or decomposition
- Updating any Go source files, tests, or configuration

## Source References

- `README.md` — the only file modified by this issue
