# Feature Specification: e2e smoke full-pipeline (20261030-072425)

**Feature Branch**: `fabrik/issue-10160`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261030-072425 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done, ending with a merged PR. The code change is deliberately trivial, so a failure points at the pipeline rather than the change. The `fabrik:yolo` label is set so the stages auto-advance and the PR is merged automatically once Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator files this issue. Fabrik carries it through all pipeline stages without manual intervention. The result is a merged PR that adds one marker line to `e2e/markers/smoke-full-pipeline.md`.

**Why this priority**: This is the only story. It is the whole purpose of the smoke test.

**Independent Test**: After the pipeline finishes, the default branch contains `e2e/markers/smoke-full-pipeline.md` with `<!-- smoke-full-pipeline-20261030-072425 -->` as its last line. The issue is in Done and its PR is merged.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the Implement stage runs, **Then** the single line `<!-- smoke-full-pipeline-20261030-072425 -->` is appended at the very end of the file and no existing content is changed.
2. **Given** the file does not exist, **When** the Implement stage runs, **Then** the file is created and contains the marker line.
3. **Given** the PR is open and all stage gates pass, **When** Validate completes under `fabrik:yolo`, **Then** the PR is merged and the issue moves to Done.

---

### Edge Cases

- The file exists but does not end with a newline: the marker must still land on its own line at the end.
- The marker line is already present, for example on a re-run: it must not be duplicated.
- Plan must not decompose the work or spawn child issues; the work is one file and one line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- smoke-full-pipeline-20261030-072425 -->`, at the end of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it MUST be created containing that line.
- **FR-003**: The change MUST touch only that one file; no other files are modified.
- **FR-004**: Existing content of the file MUST be preserved unchanged.
- **FR-005**: The work MUST stay within this single repository; no cross-repo work and no child issues are spawned.
- **FR-006**: The PR MUST be merged and the issue MUST reach Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged, without manual intervention beyond the `fabrik:yolo` label.
- **SC-002**: The PR diff is exactly one added line in exactly one file.
- **SC-003**: The last line of `e2e/markers/smoke-full-pipeline.md` on the default branch equals `<!-- smoke-full-pipeline-20261030-072425 -->`, and that line appears once.
- **SC-004**: No child issues are created.

## Assumptions

- `fabrik:yolo` handles auto-advance between stages and the final merge.
- The `e2e/markers/` directory may need to be created together with the file if it is absent.
- A trailing newline after the marker line is acceptable.

## Out of Scope

- Any change outside `e2e/markers/smoke-full-pipeline.md`.
- Cross-repo work.
- Decomposing the work into child issues.
- Tests, documentation or code changes.
