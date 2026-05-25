# Feature Specification: e2e yolo auto-merge

**Feature Branch**: `fabrik/issue-16`
**Created**: 2026-05-25
**Status**: Draft
**Input**: User description: "End-to-end verification of the GitHub native auto-merge path for yolo issues (#829)."

## Background

Fabrik's post-Validate convergence flow for `fabrik:yolo` issues was updated to use GitHub's native auto-merge mechanism (applying a `fabrik:auto-merge-enabled` label and enabling auto-merge on the PR) rather than the legacy poll-merge loop. This issue drives a minimal yolo PR through that new path to verify the mechanism works end-to-end: that Fabrik correctly enables GitHub auto-merge at Validate completion rather than attempting its own merge loop.

The change itself is intentionally trivial — one line appended to `README.md` — so that any pipeline failure is attributable to the auto-merge path, not to implementation complexity.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - GitHub native auto-merge is enabled and PR merges (Priority: P1)

A developer monitoring the project board sees this issue advance through all pipeline stages with `fabrik:yolo`. When Validate completes, Fabrik enables GitHub auto-merge on the PR (rather than polling). Once CI passes, GitHub merges the PR automatically and the issue moves to Done.

**Why this priority**: This is the sole purpose of the issue — verifying the new auto-merge convergence path. There is no secondary story.

**Independent Test**: After the issue moves to Done, check out `main` and confirm the final line of `README.md` is `<!-- auto-merge-yolo-20260525-225522 -->`. Confirm in the PR timeline that the merge was triggered by GitHub auto-merge, not by a Fabrik commit or API call.

**Acceptance Scenarios**:

1. **Given** the issue at Specify with `fabrik:yolo`, **When** Fabrik processes it, **Then** the issue advances to Research without requiring human input
2. **Given** the issue in the pipeline, **When** Implement runs, **Then** a PR is created on branch `fabrik/issue-16` with `Closes #16` in the body
3. **Given** the PR created by Implement, **When** Validate completes, **Then** Fabrik enables GitHub auto-merge on the PR (the `fabrik:auto-merge-enabled` label is applied) rather than running the legacy poll-merge loop
4. **Given** GitHub auto-merge enabled and CI passing, **When** all required checks complete, **Then** GitHub merges the PR and the issue moves to Done
5. **Given** `main` after the PR merges, **When** the last line of `README.md` is read, **Then** it is exactly `<!-- auto-merge-yolo-20260525-225522 -->`

---

### Edge Cases

- README.md may or may not end with a trailing newline; the appended line must be on its own line regardless
- The Plan stage must not decompose into sub-issues — this is a single-repo, single-file change

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Implement MUST append exactly `<!-- auto-merge-yolo-20260525-225522 -->` as the final line of `README.md`
- **FR-002**: No other source or configuration files MUST be modified
- **FR-003**: The Plan stage MUST NOT spawn sub-issues or decompose across repos
- **FR-004**: The PR body MUST contain `Closes #16` so Fabrik can discover and merge it
- **FR-005**: At Validate completion, Fabrik MUST enable GitHub native auto-merge on the PR rather than invoking the legacy poll-merge loop

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches the Done column on the project board
- **SC-002**: The linked PR is merged into `main` via GitHub native auto-merge
- **SC-003**: `README.md` on `main` ends with the line `<!-- auto-merge-yolo-20260525-225522 -->`
- **SC-004**: No other files (excluding this specification) are changed in the merged PR diff
- **SC-005**: The `fabrik:auto-merge-enabled` label is applied to the issue at Validate completion

## Assumptions

- The `fabrik:yolo` label is set on the issue, enabling auto-advance and auto-merge through all stages
- This is a single-repo operation — `handarbeit/fabrik-test-alpha` only; no cross-repo work
- CI on this repo runs `go build ./...` and `go test ./...`; appending a comment to README.md does not affect either
- GitHub auto-merge is enabled on the repository

## Out of Scope

- Any changes beyond appending the single comment line to `README.md`
- Cross-repo sub-issue spawning or decomposition
- Updating any Go source files, tests, or configuration
- Testing any pipeline behavior other than the yolo auto-merge convergence path

## Source References

- `README.md` — the only file modified by this issue
- Issue #829 — the Fabrik PR that introduced the GitHub native auto-merge path (referenced for context)
