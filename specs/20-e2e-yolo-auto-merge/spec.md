# Feature Specification: e2e yolo auto-merge

**Feature Branch**: `fabrik/issue-20`
**Created**: 2026-05-26
**Status**: Implemented
**Input**: User description: "End-to-end verification of the GitHub native auto-merge path for yolo issues (#829)."

## Background

Fabrik historically merged yolo PRs via a poll-merge loop. Issue #829 introduced a GitHub-native auto-merge path: when Validate completes on a yolo issue, Fabrik enables GitHub's built-in auto-merge on the PR (applying the `fabrik:auto-merge-enabled` label) and lets GitHub merge the PR once CI passes, rather than polling and merging directly.

This issue drives the simplest possible change — appending one HTML comment line to `README.md` — through the full yolo pipeline to verify the new convergence flow fires correctly and that the legacy poll-merge loop is not triggered.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Full yolo pipeline uses GitHub native auto-merge (Priority: P1)

A developer monitoring the project board sees this issue advance automatically through all pipeline stages (Specify → Research → Plan → Implement → Review → Validate → Done), a PR is created, GitHub auto-merge is enabled on that PR (not merged directly by Fabrik), and the single-line README change appears on `main` once CI passes.

**Why this priority**: This is the sole purpose of the issue — verifying the post-Validate convergence flow uses GitHub native auto-merge rather than the legacy poll-merge loop.

**Independent Test**: After the issue moves to Done, check out `main` and confirm the final line of README.md is `<!-- auto-merge-yolo-20260526-122826 -->`. Verify the `fabrik:auto-merge-enabled` label was applied to the PR.

**Acceptance Scenarios**:

1. **Given** the issue at Specify with `fabrik:yolo`, **When** Fabrik processes it, **Then** the issue advances to Research without requiring human input
2. **Given** the issue in the pipeline, **When** Implement runs, **Then** a PR is created on branch `fabrik/issue-20` with `Closes #20` in the body
3. **Given** the PR created by Implement, **When** Validate completes, **Then** Fabrik enables GitHub auto-merge on the PR (applies `fabrik:auto-merge-enabled` label) rather than merging directly
4. **Given** `main` after the PR merges via GitHub auto-merge, **When** the last line of `README.md` is read, **Then** it is exactly `<!-- auto-merge-yolo-20260526-122826 -->`

---

### Edge Cases

- README.md may or may not end with a trailing newline; the appended line must be on its own line regardless
- The Plan stage must not decompose into sub-issues — this is a single-repo, single-file change

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Implement MUST append exactly `<!-- auto-merge-yolo-20260526-122826 -->` as the final line of `README.md`
- **FR-002**: No other source or configuration files MUST be modified
- **FR-003**: The Plan stage MUST NOT spawn sub-issues or decompose across repos
- **FR-004**: The PR body MUST contain `Closes #20` so Fabrik can discover and merge it

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches the Done column on the project board
- **SC-002**: The linked PR is merged into `main` via GitHub native auto-merge (not the legacy poll-merge loop)
- **SC-003**: `README.md` on `main` ends with the line `<!-- auto-merge-yolo-20260526-122826 -->`
- **SC-004**: No other files (excluding this specification) are changed in the merged PR diff

## Assumptions

- The `fabrik:yolo` label is set on the issue, enabling auto-advance and GitHub native auto-merge when Validate completes
- This is a single-repo operation — `handarbeit/fabrik-test-alpha` only; no cross-repo work
- CI on this repo runs `go build ./...` and `go test ./...`; appending a comment to README.md does not affect either
- GitHub auto-merge is enabled on the repository

## Out of Scope

- Any changes beyond appending the single comment line to `README.md`
- Cross-repo sub-issue spawning or decomposition
- Updating any Go source files, tests, or configuration
- Verifying the legacy poll-merge loop (that path is superseded)

## Source References

- `README.md` — the only file modified by this issue
- Fabrik issue #829 — introduced GitHub native auto-merge path for yolo issues
