# Feature Specification: E2E smoke full-pipeline marker (20261025-064733)

**Feature Branch**: `fabrik/issue-9362`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261025-064733 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This is an end-to-end smoke test of the Fabrik pipeline. Its purpose is to confirm that a single-repo issue can move through every stage (Specify → Research → Plan → Implement → Review → Validate) and end at Done with a merged PR. The code change is intentionally trivial so that any failure points at the pipeline and not at the change. The issue carries `fabrik:yolo`, so stages auto-advance and the PR is auto-merged once Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged (Priority: P1)

An operator running the pipeline smoke wants a unique, harmless marker line added to the smoke marker file, delivered through a PR that merges, so the run proves the full pipeline works.

**Why this priority**: It is the only deliverable. A merged PR carrying exactly this line is the evidence that the pipeline completed.

**Independent Test**: After the PR merges, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch is `<!-- smoke-full-pipeline-20261025-064733 -->`, and the PR diff touches only that file and adds only that line.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists with earlier marker lines, **When** the change is implemented, **Then** the new marker line is the final line of the file and all existing lines are unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created containing the marker line.
3. **Given** the PR is open with `fabrik:yolo` set, **When** Validate completes, **Then** the PR is merged and the issue reaches Done.

---

### Edge Cases

- The file's last existing line has no trailing newline: the marker must still go on its own new line, not be joined to the previous one.
- The marker line is already present (for example on a retry): it must not be added a second time.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The exact line `<!-- smoke-full-pipeline-20261025-064733 -->` is appended as the last line of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it is created first, containing the marker line.
- **FR-003**: Existing content of the file is left unchanged.
- **FR-004**: The marker line appears exactly once in the file.
- **FR-005**: The change touches only `e2e/markers/smoke-full-pipeline.md` in this repository (plus the engine-generated spec file under `specs/`).
- **FR-006**: The work is not decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with a merged PR, with no human intervention beyond the initial labels.
- **SC-002**: On the default branch after merge, the marker line is the last line of the file and occurs exactly once.
- **SC-003**: The merged PR changes exactly one line in one non-spec file.

## Assumptions

- The marker file already exists (it holds earlier smoke-run marker lines); the create-if-missing path is a fallback only.
- A trailing-newline convention matching the rest of the file is acceptable.
- `fabrik:yolo` governs auto-advance and auto-merge, so no human approval of this spec is needed.

## Out of Scope *(optional)*

- Any cross-repo work.
- Any change to engine code, stage configuration or documentation.
- Decomposing the work into child issues during Plan.
