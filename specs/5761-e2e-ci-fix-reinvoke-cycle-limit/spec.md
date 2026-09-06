# Feature Specification: e2e ci-fix-cycle-limit (20260906-054040)

**Feature Branch**: `fabrik/issue-5761`
**Created**: 2026-09-06
**Status**: Draft
**Input**: User description: "Regression coverage for Fabrik's CI-fix reinvoke cycle limit (`MaxCiFixCycles`, handarbeit/fabrik#900, #1323)."

## Background

Fabrik's CI-fix reinvoke loop retries a stage when a required CI check is
red, up to a configured `MaxCiFixCycles` limit, after which it escalates
(pauses the issue for human intervention). That escalation path has no
live e2e regression coverage — every existing CI-fix scenario in this repo
(`specs/80-e2e-ci-fix-reinvoke`, `specs/96-e2e-ci-fix-reinvoke`) exercises
the **convergent** case: a check that starts red and is later fixed. This
issue exercises the **divergent** counterpart: a check that can never be
fixed, so the reinvoke loop is guaranteed to run to the cycle limit and
trigger Fabrik's escalation behavior.

This repo already contains the test infrastructure this scenario depends
on: the `ci-fix-sentinel` job in `.github/workflows/ci.yml` unconditionally
fails any PR whose body contains the literal marker
`ci-fix-sentinel-unfixable` — no file change, workaround, or diagnosis can
satisfy it. This issue does not modify that job; it only triggers it.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - CI-fix reinvoke loop runs to the cycle limit and escalates (Priority: P1)

An operator watching the Fabrik engine process this issue's PR sees the
`ci-fix-sentinel` check fail permanently, triggering repeated CI-fix
reinvoke cycles. On each reinvoke, Implement appends a new attempt line to
`.ci-fix-attempts.log` and pushes a real commit — never diagnosing the
check or skipping the push. Once `MaxCiFixCycles` reinvokes have occurred
without the check passing, Fabrik pauses/escalates the issue rather than
looping forever.

**Why this priority**: This is the exact behavior under test — if the
reinvoke loop stalls early (no commit pushed) or never terminates, the
regression coverage is invalid.

**Independent Test**: After the issue is paused/escalated, confirm
`.ci-fix-attempts.log` on the PR branch has one `attempt-N` line per
reinvoke cycle (initial commit plus one per reinvoke), with `N` strictly
increasing and no gaps or duplicates.

**Acceptance Scenarios**:

1. **Given** the initial commit with the `ci-fix-sentinel-unfixable`
   marker in the PR body, **When** CI runs, **Then** the `ci-fix-sentinel`
   check fails
2. **Given** a failed required check, **When** Fabrik dispatches a CI-fix
   reinvoke, **Then** the agent appends exactly one new `attempt-N` line to
   `.ci-fix-attempts.log` and pushes a commit — never leaving the check
   unaddressed with no push
3. **Given** repeated reinvokes with the check still failing, **When** the
   reinvoke count reaches `MaxCiFixCycles`, **Then** Fabrik stops
   reinvoking and pauses/escalates the issue instead of looping forever
4. **Given** any reinvoke, **When** the agent acts, **Then** no file under
   `.github/` is modified and the `ci-fix-sentinel-unfixable` marker is
   never removed from the PR body

---

### Edge Cases

- `.ci-fix-attempts.log` may already contain out-of-order or unexpected
  lines from a partial prior run; the correct behavior is to parse the
  highest existing `attempt-N` and append `attempt-(N+1)`, not to assume
  the file's line count equals the number of prior invocations
- The check is permanently unfixable by design — an agent that concludes
  "nothing to do" and skips the push defeats the purpose of this test

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The initial commit MUST append exactly one line,
  `<!-- ci-fix-cycle-limit-test -->`, to
  `e2e/markers/ci-fix-reinvoke-cycle-limit.md`
- **FR-002**: The initial commit MUST create `.ci-fix-attempts.log` at the
  repository root containing exactly one line: `attempt-1`
- **FR-003**: The PR body MUST contain the literal marker
  `ci-fix-sentinel-unfixable` so the `ci-fix-sentinel` check fails on every
  CI run for this PR
- **FR-004**: On every CI-fix reinvoke, the agent MUST unconditionally
  append exactly one new `attempt-N` line to `.ci-fix-attempts.log` (N one
  greater than the highest existing attempt number) and push the commit —
  never skip the push, never attempt to fix or diagnose the check
- **FR-005**: No file under `.github/` MUST be modified on any reinvoke
- **FR-006**: The `ci-fix-sentinel-unfixable` marker MUST never be removed
  or altered in the PR body on any reinvoke

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The `ci-fix-sentinel` check fails on the initial commit and
  on every subsequent commit on this PR
- **SC-002**: `.ci-fix-attempts.log` gains exactly one new line per
  CI-fix reinvoke cycle, with strictly increasing `attempt-N` numbers
- **SC-003**: Fabrik's engine reaches `MaxCiFixCycles` reinvokes and then
  pauses/escalates the issue rather than reinvoking indefinitely
- **SC-004**: No file under `.github/` is modified at any point in the
  PR's history

## Assumptions

- The `ci-fix-sentinel` job in `.github/workflows/ci.yml` (added for
  handarbeit/fabrik#897) is already merged and requires no changes
- `MaxCiFixCycles` and the escalation/pause behavior are engine-side
  concerns in `handarbeit/fabrik`, asserted by an external e2e harness —
  not by anything in this repo
- Compliance with "always append and push, never diagnose" on every
  reinvoke depends on the agent literally following the issue and PR body
  instructions each time it is invoked; this repo has no code-level
  mechanism to enforce it

## Out of Scope

- Any change to `.github/` workflow files
- Any attempt to make the `ci-fix-sentinel` check pass
- Any file other than `e2e/markers/ci-fix-reinvoke-cycle-limit.md`,
  `.ci-fix-attempts.log`, and this spec

## Source References

- `.github/workflows/ci.yml` (`ci-fix-sentinel` job) — pre-existing,
  unmodified test infrastructure this scenario relies on
- handarbeit/fabrik#900, handarbeit/fabrik#1323 — the engine-side
  `MaxCiFixCycles` behavior this issue provides regression coverage for
- `specs/80-e2e-ci-fix-reinvoke`, `specs/96-e2e-ci-fix-reinvoke` — the
  related convergent CI-fix reinvoke scenarios this issue deliberately
  does not collide with (separate marker/log files)
