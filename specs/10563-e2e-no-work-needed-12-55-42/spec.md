# Feature Specification: E2E verification that no-work-needed short-circuits to Done

**Feature Branch**: `fabrik/issue-10563`
**Created**: 2026-10-31
**Status**: Specified
**Input**: User description: "## Goal

Verify that the FABRIK_NO_WORK_NEEDED short-circuit closes the issue without creating a PR.

## What this asks for

`README.md` in this repo already exists and contains a description of the test bed. Please verify that the README.md file is present, contains the substring "fabrik-test-alpha", and has the Apache 2.0 LICENSE pointer.

## Hint to Plan

This issue is purely a verification request. There is no code change to make. The Plan stage should recognise that and emit FABRIK_NO_WORK_NEEDED on a line of its own (instead of decomposing or proceeding to Implement).

This is an e2e regression test for #733 (the marker) and #742 (close-on-no-work)."

## Background

This is an end-to-end regression test for two engine behaviours: the `FABRIK_NO_WORK_NEEDED` marker (#733) and closing the issue when no work is needed (#742). The issue asks for a verification only. The test bed's `README.md` should already satisfy the stated conditions, so no code or documentation change is expected. The pipeline should recognise this, skip the remaining stages, create no pull request, and move the issue to Done and close it.

At Specify time, the README was checked and already meets the conditions:
- `README.md` is present at the repository root.
- It contains the substring `fabrik-test-alpha`.
- It has a `## License` section stating "Apache 2.0 (matches Fabrik upstream)", and a `LICENSE` file is present at the repository root.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Verify the README conditions without changing anything (Priority: P1)

An operator running the e2e suite wants confirmation that the test bed README is present, names the test bed, and carries the Apache 2.0 license pointer, without any file being modified.

**Why this priority**: This is the whole substance of the issue. The pipeline outcome depends on the verification being a pure read.

**Independent Test**: Inspect `README.md` and confirm the file exists, contains `fabrik-test-alpha`, and references Apache 2.0. Confirm the working tree is unchanged afterwards.

**Acceptance Scenarios**:

1. **Given** the repository at its default base branch, **When** the README is inspected, **Then** `README.md` exists and contains the substring `fabrik-test-alpha`.
2. **Given** the same repository, **When** the README is inspected, **Then** it contains an Apache 2.0 license pointer.
3. **Given** all three conditions hold, **When** the verification finishes, **Then** no file in the repository has been changed.

---

### User Story 2 - No-work short-circuit closes the issue without a PR (Priority: P1)

The pipeline should detect that nothing needs doing and finish the issue without further stages. This exercises #733 and #742.

**Why this priority**: This is the engine behaviour the e2e test is meant to protect against regression.

**Independent Test**: After the Plan stage signals completion together with the no-work marker, check that the remaining stages are marked skipped, no PR exists for the issue, the board status is Done, and the issue is closed.

**Acceptance Scenarios**:

1. **Given** Plan recognises that the issue is a pure verification, **When** it signals stage completion together with the no-work marker, **Then** the remaining stages are marked complete as skipped.
2. **Given** the no-work short-circuit has been taken, **When** the pipeline finishes, **Then** no pull request has been created for the issue.
3. **Given** the same state, **When** the pipeline finishes, **Then** the issue is on the Done column and is closed.

---

### Edge Cases

- If the README were missing or lacked `fabrik-test-alpha` or the Apache 2.0 pointer, the verification would fail and real work would be needed; the no-work path must not be taken in that case. This is not the current state.
- A stage that decomposes the issue or proceeds to Implement instead of recognising it as a pure verification would be a regression for this test.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The verification MUST confirm that `README.md` is present at the repository root.
- **FR-002**: The verification MUST confirm that `README.md` contains the substring `fabrik-test-alpha`.
- **FR-003**: The verification MUST confirm that `README.md` contains an Apache 2.0 license pointer.
- **FR-004**: The pipeline MUST NOT modify any repository file as part of this issue.
- **FR-005**: When the verification passes, the Plan stage MUST signal that no work is needed instead of decomposing the issue or proceeding to Implement.
- **FR-006**: On the no-work path, the engine MUST NOT create a pull request for this issue.
- **FR-007**: On the no-work path, the engine MUST mark the remaining pipeline stages complete as skipped, move the issue to Done, and close it.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: All three README conditions (present, names `fabrik-test-alpha`, Apache 2.0 pointer) are confirmed true.
- **SC-002**: Zero pull requests are linked to this issue when it reaches Done.
- **SC-003**: The issue ends closed and in the Done column, with no Implement, Review, or Validate stage run.
- **SC-004**: No commits other than the engine's own spec projection are added to the issue branch.

## Assumptions

- "Apache 2.0 LICENSE pointer" is satisfied by the README's License section naming Apache 2.0, together with the repository's `LICENSE` file; no specific wording is required.
- The `fabrik:yolo` label is present, so stages advance without waiting for human approval of the spec.
- The engine's own spec file projection from Specify is not counted as a change made by this issue's work.

## Out of Scope

- Editing the README or the LICENSE file.
- Any change to engine code, stage configuration, or tests.
- Verifying behaviours of #733 and #742 beyond what this single run exercises.

## Source References

- Issue #733: introduction of the `FABRIK_NO_WORK_NEEDED` marker.
- Issue #742: close-on-no-work behaviour.
