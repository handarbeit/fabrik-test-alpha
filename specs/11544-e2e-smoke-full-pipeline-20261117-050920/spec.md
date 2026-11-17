# Feature Specification: E2E smoke full-pipeline (20261117-050920)

**Feature Branch**: `fabrik/issue-11544`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261117-050920 -->`. That's the entire change. One file, one line. Scope: Single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This is an end-to-end smoke test of the Fabrik pipeline in a single repository. It checks that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR, with no human intervention beyond the `fabrik:yolo` label. The change itself is deliberately trivial, so a failure points at the pipeline and not at the work.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged via the full pipeline (Priority: P1)

An operator files this issue with `fabrik:yolo`. Fabrik carries it through all stages unattended. It appends one marker comment line to `e2e/markers/smoke-full-pipeline.md` and merges the resulting PR, and the issue ends in Done.

**Why this priority**: This is the only purpose of the issue. It exercises the whole pipeline end to end.

**Independent Test**: After the pipeline completes, check that `e2e/markers/smoke-full-pipeline.md` on the default branch ends with `<!-- smoke-full-pipeline-20261117-050920 -->`. Also check that the PR is merged and the issue is in Done.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is implemented, **Then** the line `<!-- smoke-full-pipeline-20261117-050920 -->` is appended as the last line and existing content is unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances automatically through to a merged PR and Done.

---

### Edge Cases

- The marker file already ends without a trailing newline: the new line must still start on its own line.
- The marker line is already present (for example on a retry): it must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` is created if it does not exist.
- **FR-002**: The line `<!-- smoke-full-pipeline-20261117-050920 -->` is appended at the very end of that file, on its own line, exactly once.
- **FR-003**: No other file is changed apart from the engine-generated spec file for this issue, and no other line in the marker file is modified.
- **FR-004**: All work happens in this one repository. The Plan stage does not decompose or spawn child issues.
- **FR-005**: The change is delivered through a PR that references this issue so Fabrik can discover it. The PR is merged and the issue reaches Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch is exactly `<!-- smoke-full-pipeline-20261117-050920 -->`, and that string appears exactly once.
- **SC-002**: The PR diff contains exactly one added line in the marker file, apart from the engine-generated spec file.
- **SC-003**: The issue reaches Done with its PR merged, with no manual intervention after the `fabrik:yolo` label.
- **SC-004**: No child issues are created.

## Assumptions

- The `e2e/markers/` directory may need to be created together with the file if it is absent.
- `fabrik:yolo` auto-advances every stage and auto-merges the PR at Validate completion.
- The Specify stage's engine-generated spec file under `specs/` is expected pipeline output and not part of the "one file" constraint.

## Out of Scope *(optional)*

- Cross-repo work.
- Decomposition into child issues.
- Any change to the engine, stage configs, or other marker files.
