# Feature Specification: Verify README presence, name and license pointer (no-work-needed e2e)

**Feature Branch**: `fabrik/issue-8547`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

Verify that the FABRIK_NO_WORK_NEEDED short-circuit closes the issue without creating a PR.

## What this asks for

`README.md` in this repo already exists and contains a description of the test bed. Please verify that the README.md file is present, contains the substring \"`+`fabrik-test-alpha`+`\", and has the Apache 2.0 LICENSE pointer.

## Hint to Plan

This issue is purely a verification request. There is no code change to make. The Plan stage should recognise that and emit FABRIK_NO_WORK_NEEDED on a line of its own (instead of decomposing or proceeding to Implement).

This is an e2e regression test for #733 (the marker) and #742 (close-on-no-work)."

## Background

This issue is an end-to-end regression test for the engine's "no work needed" short-circuit (#733, the `FABRIK_NO_WORK_NEEDED` marker; #742, closing the issue on no-work). The repository already satisfies the stated requirement, so the correct outcome is that no code or documentation changes are made, no PR is created, and the issue is closed via the short-circuit.

The Specify stage has confirmed the current state of the repository:

- `README.md` exists at the repository root and begins with the heading `# fabrik-test-alpha`.
- Its `## License` section states "Apache 2.0 (matches Fabrik upstream)."
- A `LICENSE` file containing the Apache License, Version 2.0 is present at the repository root.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Confirm the README meets the stated conditions (Priority: P1)

A maintainer wants confirmation that the existing README identifies the repository by name and points to the Apache 2.0 license, without any change being made.

**Why this priority**: This is the only functional content of the issue.

**Independent Test**: Open `README.md` and check that the file exists, contains `fabrik-test-alpha`, and contains an Apache 2.0 license pointer.

**Acceptance Scenarios**:

1. **Given** the repository at its current state, **When** `README.md` is inspected, **Then** the file exists and contains the text `fabrik-test-alpha`.
2. **Given** the repository at its current state, **When** the README's License section is read, **Then** it names Apache 2.0, and a `LICENSE` file with the Apache 2.0 text exists.

---

### User Story 2 - Exercise the no-work-needed short-circuit (Priority: P1)

An engine maintainer wants the pipeline to recognise a pure verification request and finish it without producing a PR.

**Why this priority**: This is the purpose of the e2e test (#733, #742).

**Independent Test**: Run the issue through the pipeline and observe that it reaches Done and is closed with no PR and no repository commits other than the engine's own spec projection.

**Acceptance Scenarios**:

1. **Given** the verification conditions already hold, **When** the Plan stage evaluates the issue, **Then** it signals that no work is needed rather than decomposing the issue or proceeding to Implement.
2. **Given** the no-work-needed signal, **When** the engine processes it, **Then** the remaining stages are marked skipped, the issue moves to Done and is closed, and no PR is created.

---

### Edge Cases

- If the README were missing the repository name or the license pointer, the issue would stop being verification-only and would need an actual documentation change.
- The quoted substring in the request is wrapped in literal `+` and backtick characters. This is treated as formatting noise (see Assumptions).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `README.md` MUST be present at the repository root.
- **FR-002**: `README.md` MUST contain the substring `fabrik-test-alpha`.
- **FR-003**: `README.md` MUST contain a pointer to the Apache 2.0 license, and the repository MUST contain the corresponding `LICENSE` file.
- **FR-004**: Because FR-001 to FR-003 already hold, no change to repository files is required, and the pipeline SHOULD conclude via the no-work-needed short-circuit at Plan without creating a PR.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Verification of FR-001 to FR-003 passes against the current repository with zero file modifications.
- **SC-002**: The issue ends in Done and closed with zero PRs created for it.

## Assumptions

- The quoted substring "`+`fabrik-test-alpha`+`" is a templating artifact from the test harness. The intended substring is `fabrik-test-alpha`. The literal characters `` `+` `` do not appear in the README, and adding them is not intended.
- "Apache 2.0 LICENSE pointer" is satisfied by the README's License section naming Apache 2.0 together with the `LICENSE` file at the repository root.
- The Plan stage, not Specify, is where the no-work-needed signal is expected, per the hint in the request.

## Out of Scope

- Any edit to `README.md` or `LICENSE`.
- Any change to engine behavior for `FABRIK_NO_WORK_NEEDED`.
- Creating a PR or a code change of any kind.
