# Feature Specification: e2e paused merged-PR recovery (awaiting-review 20260619-155500)

**Feature Branch**: `fabrik/issue-61`
**Created**: 2026-06-19
**Status**: Draft
**Input**: User description: "End-to-end regression guard for the #874 bug class (paused item + merged PR recovery via the settle-owner, ADR-056 D2)."

## Background

A class of bugs (#874) existed where a cruise issue could become stuck: its linked PR was merged externally while the issue held `fabrik:paused` + `fabrik:awaiting-input` (and optionally a gate label like `fabrik:awaiting-review` at the Validate stage). In this stuck state, Fabrik would not advance the issue to Done because the normal stage-completion path was blocked. ADR-056 D2 introduced a settle-owner reconciliation path that detects a merged PR linked to a paused/stuck issue and heals it directly to CLOSED without re-invoking Validate.

This issue drives the simplest possible change — appending one HTML comment line to `README.md` — through a cruise pipeline, then simulates the stuck scenario to verify the settle-owner recovery fires correctly and the issue closes without Validate being re-invoked.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Settle-owner heals stuck cruise issue when linked PR is merged externally (Priority: P1)

A project board observer sees a cruise issue that is paused and awaiting input at Validate with its linked PR already merged. Instead of remaining stuck, the issue transitions to CLOSED automatically — the settle-owner detects the merged PR and heals the issue to Done without Validate running again.

**Why this priority**: This is the sole purpose of the issue — guarding against regression of the #874 bug class where merged-PR recovery was absent.

**Independent Test**: After the issue moves to Done/CLOSED, verify no additional Validate invocation was triggered after the PR merged, and confirm `README.md` on `main` ends with `<!-- paused-merged-pr-awaiting-review-20260619-155500 -->`.

**Acceptance Scenarios**:

1. **Given** the issue at Specify with `fabrik:cruise`, **When** Fabrik processes it, **Then** the issue advances through Research → Plan → Implement without requiring human input
2. **Given** the issue in the pipeline, **When** Implement runs, **Then** a PR is created on branch `fabrik/issue-61` with `Closes #61` in the body
3. **Given** the issue in the stuck state (`fabrik:paused` + `fabrik:awaiting-input` with linked PR merged externally), **When** the settle-owner reconciliation runs, **Then** the issue is healed to CLOSED without Validate being re-invoked
4. **Given** `main` after the PR merges, **When** the last line of `README.md` is read, **Then** it is exactly `<!-- paused-merged-pr-awaiting-review-20260619-155500 -->`

---

### Edge Cases

- README.md may or may not end with a trailing newline; the appended line must be on its own line regardless
- The Plan stage must not decompose into sub-issues — this is a single-repo, single-file change
- The optional gate label (`fabrik:awaiting-review`) at Validate should not prevent settle-owner recovery

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Implement MUST append exactly `<!-- paused-merged-pr-awaiting-review-20260619-155500 -->` as the final line of `README.md`
- **FR-002**: No other source or configuration files MUST be modified
- **FR-003**: The Plan stage MUST NOT spawn sub-issues or decompose across repos
- **FR-004**: The PR body MUST contain `Closes #61` so Fabrik can discover and link it
- **FR-005**: The settle-owner MUST close the issue when its linked PR is merged and the issue is in the stuck state (`fabrik:paused` + `fabrik:awaiting-input`)

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches the Done column on the project board and is CLOSED
- **SC-002**: The issue was healed by the settle-owner without Validate being re-invoked after the PR merged
- **SC-003**: `README.md` on `main` ends with the line `<!-- paused-merged-pr-awaiting-review-20260619-155500 -->`
- **SC-004**: No other files (excluding this specification) are changed in the merged PR diff

## Assumptions

- The `fabrik:cruise` label is set on the issue, enabling auto-advance without auto-merging
- This is a single-repo operation — `handarbeit/fabrik-test-alpha` only; no cross-repo work
- CI on this repo runs `go build ./...` and `go test ./...`; appending a comment to README.md does not affect either
- The settle-owner reconciliation path (ADR-056 D2) is already implemented in Fabrik

## Out of Scope

- Any changes beyond appending the single comment line to `README.md`
- Cross-repo sub-issue spawning or decomposition
- Updating any Go source files, tests, or configuration
- Verifying any other stuck-state recovery paths beyond the paused + merged PR scenario

## Source References

- `README.md` — the only file modified by this issue
- Fabrik issue #874 — the bug class this regression guard targets
- ADR-056 D2 — the settle-owner reconciliation design decision
