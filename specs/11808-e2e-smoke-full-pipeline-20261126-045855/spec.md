# Feature Specification: e2e smoke full-pipeline (20261126-045855)

**Feature Branch**: `fabrik/issue-11808`
**Created**: 2026-11-25
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261126-045855 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. Its purpose is to confirm that a single-repo issue can move through every stage (Specify → Research → Plan → Implement → Review → Validate) and reach Done with a merged PR, with no human intervention beyond the initial `fabrik:yolo` label. The code change is deliberately trivial so that any failure points at the pipeline rather than at the work. The marker line carries a unique run identifier (`20261126-045855`), so this run's change can be told apart from other smoke runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands via a fully automated pipeline run (Priority: P1)

An operator files this issue with `fabrik:yolo`. Fabrik carries it through all stages unattended. A single marker line is appended to the smoke-test marker file, and the resulting PR is merged to the default branch.

**Why this priority**: This is the only story. It is the whole point of the smoke test.

**Independent Test**: After the run, `e2e/markers/smoke-full-pipeline.md` on `main` ends with the line `<!-- smoke-full-pipeline-20261126-045855 -->`. The issue is closed and on Done, and the linked PR is merged.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the pipeline completes, **Then** the file's last line is `<!-- smoke-full-pipeline-20261126-045855 -->` and all prior content is unchanged.
2. **Given** the marker file does not exist, **When** the pipeline completes, **Then** the file exists and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances to the next stage without a human step, and the PR is merged after Validate.
4. **Given** the issue is single-repo, **When** the Plan stage runs, **Then** it produces no child issues.

---

### Edge Cases

- The marker file exists without a trailing newline: the marker line must still land on its own line, not be joined to the previous content.
- The same marker line is already present, for example on a re-run of this issue: the file must not get a duplicate.
- Other smoke runs append their own marker lines to the same file: this run's line is appended after them and must leave theirs untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261126-045855 -->`.
- **FR-002**: If the file does not exist, it MUST be created, containing the marker line.
- **FR-003**: If the file exists, the marker line MUST be appended at the very end, with all existing content preserved unchanged.
- **FR-004**: The change MUST touch exactly one file and add exactly one line.
- **FR-005**: The work MUST stay within this repository: no cross-repo changes.
- **FR-006**: The Plan stage MUST NOT decompose the work into child issues.
- **FR-007**: The PR MUST include `Closes #11808`.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, a plain text file whose only purpose is to carry per-run marker comments.
- **Marker line**: the HTML comment `<!-- smoke-full-pipeline-20261126-045855 -->`, unique to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its PR merged, with no manual step after the initial label.
- **SC-002**: The merged diff is exactly one file with one added line, and that line is the marker line, with no other additions or deletions beyond file creation if it was needed.
- **SC-003**: No child issues are created during the run.

## Assumptions

- The default branch is `main`, and the PR targets it.
- The `fabrik:yolo` label drives auto-advance and auto-merge, so no manual approval of this spec is expected.
- The marker line is an HTML comment so it does not change how the file renders.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Cross-repo work.
- Decomposition into child issues.
- Changes to Fabrik itself or to the e2e harness.
