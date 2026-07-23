# Feature Specification: e2e cruise full pipeline (20260723-035603)

**Feature Branch**: `fabrik/issue-3574`
**Created**: 2026-07-22
**Status**: Draft
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line to README.md at the very end of the file. Verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

Fabrik's pipeline drives issues through Specify → Research → Plan → Implement → Review → Validate → Done stages. The `fabrik:cruise` label (contract defined in #898) is a variant of `fabrik:yolo`: it auto-advances an issue through all stages without human intervention, but — unlike `yolo` — it does **not** auto-merge the linked PR or auto-advance the issue to Done when Validate completes. A human must merge the PR manually, after which the issue is expected to close correctly.

This issue is a recurrent single-repo end-to-end sentinel test for the `cruise` contract specifically (as distinct from the `yolo` contract already covered by other e2e sentinels such as #51). It uses the simplest possible change — one line appended to README.md — so that any regression in `cruise` behavior is immediately visible without noise from complex implementation work.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Cruise auto-advances without auto-merging (Priority: P1)

A developer monitoring the Fabrik project board sees this issue advance automatically through all pipeline stages (Specify → Research → Plan → Implement → Review → Validate) without requiring human input, a PR is created, and the pipeline halts at Validate-complete with the PR left unmerged and the issue not yet moved to Done.

**Why this priority**: This is the core contract under test — `cruise` must behave like `yolo` for stage auto-advance but must stop short of merging.

**Independent Test**: After the issue reaches Validate-complete, confirm the linked PR exists, is open (not merged), and the issue is still open and not in the Done column.

**Acceptance Scenarios**:

1. **Given** the issue at Specify with only `fabrik:cruise` set, **When** Fabrik processes it, **Then** the issue advances through Research, Plan, Implement, and Review without requiring human input
2. **Given** the issue in the pipeline, **When** Implement runs, **Then** a PR is created on branch `fabrik/issue-3574` with `Closes #3574` in the body
3. **Given** the PR created by Implement, **When** Validate completes, **Then** the PR remains unmerged and the issue does NOT auto-advance to Done

---

### User Story 2 - Issue closes correctly after human merge (Priority: P2)

After Validate-complete, a human reviews and merges the PR manually. The issue must then close correctly, consistent with Fabrik's normal PR-merge-driven closure behavior.

**Why this priority**: Confirms the second half of the `cruise` contract — that stopping short of auto-merge doesn't leave the issue in a state where manual completion is impossible or produces incorrect state.

**Independent Test**: After a human merges the PR, confirm the issue closes and `main` contains the appended README line.

**Acceptance Scenarios**:

1. **Given** the PR is open and unmerged after Validate-complete, **When** a human merges the PR, **Then** the issue closes
2. **Given** the merged PR, **When** the last line of `README.md` on `main` is read, **Then** it is exactly `<!-- cruise-pipeline-20260723-035603 -->`

---

### Edge Cases

- README.md may or may not end with a trailing newline; the appended line must be on its own line regardless
- The Plan stage must not decompose into sub-issues — this is a single-repo, single-file change
- If both `fabrik:cruise` and `fabrik:yolo` were present, `yolo` would take precedence (per documented label semantics), but this issue carries only `fabrik:cruise`

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Implement MUST append exactly `<!-- cruise-pipeline-20260723-035603 -->` as the final line of `README.md`
- **FR-002**: No other source or configuration files MUST be modified
- **FR-003**: The Plan stage MUST NOT spawn sub-issues or decompose across repos
- **FR-004**: The PR body MUST contain `Closes #3574` so Fabrik can discover the PR and close the issue when it is merged
- **FR-005**: Fabrik MUST auto-advance the issue through all stages up to and including Validate-complete without requiring human input, driven solely by `fabrik:cruise`
- **FR-006**: Fabrik MUST NOT auto-merge the PR and MUST NOT auto-advance the issue to Done when Validate completes while only `fabrik:cruise` is set
- **FR-007**: After a human manually merges the PR, the issue MUST close

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Validate-complete with no human interaction required for any stage transition
- **SC-002**: At Validate-complete, the linked PR is open (unmerged) and the issue is open (not in Done)
- **SC-003**: After a human merges the PR, the issue closes
- **SC-004**: `README.md` on `main` ends with the line `<!-- cruise-pipeline-20260723-035603 -->`
- **SC-005**: No other files (excluding this specification) are changed in the merged PR diff

## Assumptions

- The `fabrik:cruise` label (and no `fabrik:yolo` label) is set on the issue, enabling auto-advance without auto-merge through all stages
- This is a single-repo operation — `handarbeit/fabrik-test-alpha` only; no cross-repo work
- CI on this repo does not block on a README-only change
- "Closes correctly" for User Story 2 means the issue transitions to closed state via the normal PR-merge-driven mechanism; no additional manual board move is required beyond merging the PR

## Out of Scope

- Any changes beyond appending the single comment line to `README.md`
- Cross-repo sub-issue spawning or decomposition
- Updating any Go source files, tests, or configuration
- Verifying `fabrik:yolo` behavior (covered by other e2e sentinels, e.g. #51)

## Source References

- `README.md` — the only file modified by this issue
- #898 — defines the `fabrik:cruise` pipeline contract
- #51 — prior e2e sentinel covering the `fabrik:yolo` contract for comparison
