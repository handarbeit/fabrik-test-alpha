# Feature Specification: E2E smoke full-pipeline marker line (20261031-162417)

**Feature Branch**: `fabrik/issue-10533`
**Created**: 2026-10-31
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261031-162417 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is a live end-to-end smoke test of the Fabrik pipeline. Its purpose is to confirm that an issue can move from Specify through Research, Plan, Implement, Review and Validate to Done, ending in a merged PR, without human intervention (the `fabrik:yolo` label enables auto-advance and auto-merge). The code change is deliberately trivial so that any failure points at the pipeline, not the change. The run is identified by the timestamp `20261031-162417`, which makes the appended line unique to this run.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator running the release gate wants proof that a single-repo issue travels through every pipeline stage and lands on the default branch. The proof is a one-line addition to a marker file that exists only because this run made it.

**Why this priority**: It is the only story. The whole value of the issue is the end-to-end traversal and the merged result.

**Independent Test**: After the pipeline finishes, read `e2e/markers/smoke-full-pipeline.md` on `main` and check that its last line is the marker line, that it appears exactly once, and that the rest of the file is unchanged.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists on `main`, **When** the pipeline completes, **Then** the file's final line is `<!-- smoke-full-pipeline-20261031-162417 -->` and all earlier content is unchanged.
2. **Given** the issue carries `fabrik:yolo`, **When** Validate completes, **Then** the linked PR is merged and the issue reaches Done with no manual action.
3. **Given** the change touches one repository, **When** Plan runs, **Then** it produces no child issues.

---

### Edge Cases

- The marker file already contains other lines from earlier smoke runs: the new line is added after them and none are altered or removed.
- The file's last line has no trailing newline: the marker still lands on its own line, not joined to the previous one.
- The marker line is already present (a re-run of the same stage): it is not added a second time.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The final line of `e2e/markers/smoke-full-pipeline.md` MUST be exactly `<!-- smoke-full-pipeline-20261031-162417 -->`.
- **FR-002**: Existing content of `e2e/markers/smoke-full-pipeline.md` MUST NOT be modified or removed.
- **FR-003**: The marker line MUST appear exactly once in the file.
- **FR-004**: The change set MUST consist of that one added line in that one file, and nothing else.
- **FR-005**: The work MUST stay within this repository, and Plan MUST NOT decompose it into child issues.
- **FR-006**: The pipeline MUST carry the issue from Specify to Done, with the linked PR merged, under `fabrik:yolo`.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, a text file that records one HTML comment line per smoke run.
- **Marker line**: the HTML comment `<!-- smoke-full-pipeline-20261031-162417 -->`, unique to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches the Done column with its linked PR merged to `main`, and no operator intervention was needed after the label was applied.
- **SC-002**: The merged PR's diff is exactly one added line in one file.
- **SC-003**: On `main` after the merge, the marker line appears once, as the last line of the file.

## Assumptions

- The marker file already exists on `main`, so no file creation is needed. The issue allows for creating it if absent.
- Marker lines from earlier runs may sit in the file; the new line goes after them.
- The Plan stage needs no decomposition because the change is one line in one file.

## Out of Scope

- Cross-repository work.
- Any change outside `e2e/markers/smoke-full-pipeline.md`.
- Changes to engine behaviour, tests, or documentation.
- Cleaning up or reformatting earlier marker lines.
