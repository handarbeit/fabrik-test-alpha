# Feature Specification: e2e cruise full pipeline (20261107-202923)

**Feature Branch**: `fabrik/issue-10857`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261107-202923 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue exists to exercise the `fabrik:cruise` pipeline contract (#898) end to end. Under cruise, an issue auto-advances through every stage up to Validate-complete, but the engine does not merge the PR or move the issue to Done. A human merges the PR, and the issue then closes. The code change is deliberately trivial, so that any failure in the run points at pipeline behavior and not at the content of the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

The marker file `e2e/markers/cruise-full-pipeline.md` ends with the line `<!-- cruise-pipeline-20261107-202923 -->`. If the file did not exist, it is created containing that line.

**Why this priority**: This is the only change the issue asks for, and the cruise run needs a real PR to verify against.

**Independent Test**: Read the last line of `e2e/markers/cruise-full-pipeline.md` on the PR branch and compare it with the expected marker text.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is applied, **Then** the existing content is unchanged and the marker line is the final line of the file.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created and contains the marker line.
3. **Given** the change is applied, **When** the PR diff is inspected, **Then** exactly one file is touched and exactly one line is added.

---

### User Story 2 - Cruise pipeline contract holds (Priority: P1)

The issue carries `fabrik:cruise` and advances through all stages to Validate-complete. The PR is not auto-merged and the issue does not move to Done until a human merges the PR. After that merge, the issue closes.

**Why this priority**: This is the purpose of the issue. It verifies the contract from #898.

**Independent Test**: Observe the issue and its PR from Specify through Validate-complete. Confirm the PR is still open and unmerged. Merge it by hand and confirm the issue closes.

**Acceptance Scenarios**:

1. **Given** the issue has `fabrik:cruise`, **When** each stage completes, **Then** the engine advances to the next stage without waiting for human approval.
2. **Given** Validate has completed, **When** the engine evaluates landing, **Then** it does not merge the PR or move the issue to Done.
3. **Given** a human has merged the PR, **When** the engine next processes the issue, **Then** the issue is closed.

---

### Edge Cases

- The marker file is absent: it is created, with the marker line as its only content.
- The existing file does not end with a newline: the marker still goes on its own line at the end.
- The Plan stage must not split this work into child issues, because the change is one file and one line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/cruise-full-pipeline.md` MUST end with the line `<!-- cruise-pipeline-20261107-202923 -->`.
- **FR-002**: If the file does not exist, it MUST be created, and its content MUST be that line.
- **FR-003**: The change MUST consist of exactly one added line in exactly one file. Existing content MUST NOT be modified or removed.
- **FR-004**: The Plan stage MUST NOT decompose the work into child issues.
- **FR-005**: Under `fabrik:cruise`, the issue MUST advance through all stages to Validate-complete without human approval at intermediate stages.
- **FR-006**: The PR MUST NOT be merged automatically and the issue MUST NOT advance to Done by engine action at Validate-complete.
- **FR-007**: After a human merges the PR, the issue MUST close.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff touches exactly one file and adds exactly one line, which is the marker line.
- **SC-002**: The issue reaches Validate-complete with no human approval at any intermediate stage and with its PR still open and unmerged.
- **SC-003**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker comment text is exactly as given in the request, including the timestamp `20261107-202923`.
- The marker file lives at the repository root path `e2e/markers/`, and the target repository is the single repository this issue belongs to.
- The marker line is appended after any existing content, and a trailing newline at the end of the file is acceptable.

## Out of Scope *(optional)*

- Any change to files other than `e2e/markers/cruise-full-pipeline.md`.
- Any change to engine behavior or to the cruise contract itself.
- Cross-repo work.
- Decomposition into child issues.
