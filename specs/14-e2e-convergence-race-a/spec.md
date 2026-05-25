# Feature Specification: e2e convergence-race A (20260525-225522)

**Feature Branch**: `fabrik/issue-14`
**Created**: 2026-05-25
**Status**: Draft
**Input**: User description: "Deterministic provocation of the post-Validate auto-merge race covered by handarbeit/fabrik#829. This issue is one of a deliberately-conflicting pair filed by the e2e harness."

## Background

Fabrik's auto-merge path (post-Validate) contains a race window: if `main` advances after the PR is validated but before the merge completes, the merge can fail or produce a corrupt result. Issue handarbeit/fabrik#829 tracks this bug. This issue is one half of a deliberate conflict pair created by the e2e harness to reproduce the race deterministically.

Both pair members insert a single HTML comment immediately after the `# fabrik-test-alpha` heading in `README.md`, at the same line position but with different discriminators. Because they share a position, one will conflict with the other when the second branch tries to rebase or merge onto `main`. The `slow-ci-required` marker in the PR body triggers a 6-minute CI slow-gate, giving the race a deterministic window during which `main` can move under the second PR.

The `fabrik:yolo` label is present, enabling auto-advance and auto-merge.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Race pair provokes a real textual conflict (Priority: P1)

The e2e harness watches both pair members advance through the pipeline. When the second PR attempts to rebase or merge onto `main` (which already contains the first PR's marker at the same position), a true textual conflict is produced. Fabrik must handle this correctly per fabrik#829.

**Why this priority**: This is the sole purpose of the issue pair — producing a deterministic conflict at the post-Validate merge point. There is no secondary story.

**Independent Test**: After the PR for this issue merges, check out `main` and confirm the line `<!-- convergence-race-A-20260525-225522 -->` appears immediately after `# fabrik-test-alpha` in `README.md`.

**Acceptance Scenarios**:

1. **Given** `README.md` on `main` contains `# fabrik-test-alpha` as the first line, **When** Implement runs, **Then** the line `<!-- convergence-race-A-20260525-225522 -->` is inserted immediately after it and no other lines are added or modified
2. **Given** the PR created by Implement, **When** the PR body is read, **Then** it contains the literal string `slow-ci-required` (triggering the CI slow-gate)
3. **Given** the PR created by Implement, **When** the PR body is read, **Then** it contains `Closes #14` so Fabrik can discover and auto-merge it
4. **Given** a rebase conflict where the pair partner's marker `<!-- convergence-race-B-20260525-225522 -->` is already present at the same position, **When** the conflict is resolved, **Then** both lines are kept (in either order)

---

### Edge Cases

- If the pair partner's PR merges first, `main` will already have `<!-- convergence-race-B-20260525-225522 -->` at the same position; rebase will produce a conflict that must be resolved by keeping both lines
- The CI slow-gate (triggered by `slow-ci-required` in the PR body) holds the pipeline open for ~6 minutes; the race window is intentional

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Implement MUST insert exactly `<!-- convergence-race-A-20260525-225522 -->` as a new line immediately after the line `# fabrik-test-alpha` in `README.md`
- **FR-002**: No other files (excluding this specification) MUST be modified
- **FR-003**: The PR body MUST contain the literal string `slow-ci-required` (exact, unformatted) so the CI slow-gate fires
- **FR-004**: The PR body MUST contain `Closes #14` so Fabrik can discover and auto-merge the PR
- **FR-005**: The Plan stage MUST NOT decompose into sub-issues — this is a single-repo, single-file, one-line change
- **FR-006**: If a rebase conflict is encountered where the pair partner's marker is already present at the same position, the conflict MUST be resolved by keeping both lines (in either order)

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches the Done column on the project board
- **SC-002**: The linked PR is merged into `main`
- **SC-003**: `README.md` on `main` contains `<!-- convergence-race-A-20260525-225522 -->` on the line immediately following `# fabrik-test-alpha`
- **SC-004**: No other files (excluding this specification) are changed in the merged PR diff
- **SC-005**: The merged PR body contains `slow-ci-required`

## Assumptions

- The `fabrik:yolo` label is set, enabling auto-advance through all stages and auto-merge at Validate completion
- This is a single-repo operation — `handarbeit/fabrik-test-alpha` only
- `README.md` currently begins with `# fabrik-test-alpha` as its first line
- CI on this repo runs `go build ./...` and `go test ./...`; inserting an HTML comment into `README.md` does not affect either
- The pair partner issue (inserting `<!-- convergence-race-B-20260525-225522 -->`) may or may not have merged by the time this PR is processed

## Out of Scope

- Any changes beyond inserting the single comment line into `README.md`
- Cross-repo sub-issue spawning or decomposition
- Updating any Go source files, tests, or configuration

## Source References

- `README.md` — the only file modified by this issue
- handarbeit/fabrik#829 — the auto-merge race bug this test provokes
