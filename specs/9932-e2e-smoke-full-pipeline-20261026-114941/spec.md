# Feature Specification: E2E smoke full-pipeline marker (20261026-114941)

**Feature Branch**: `fabrik/issue-9932`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261026-114941 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. The change itself is deliberately trivial so that any failure points at the pipeline (Specify → Research → Plan → Implement → Review → Validate → Done, including PR creation and merge) rather than at the work. The marker file `e2e/markers/smoke-full-pipeline.md` accumulates one HTML comment line per smoke run; this run adds the line for timestamp `20261026-114941`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended and merged through the full pipeline (Priority: P1)

A maintainer runs the smoke issue through the whole pipeline with `fabrik:yolo` set. The pipeline appends one marker line to the marker file, opens a PR linked to the issue, and merges it, leaving the issue in Done.

**Why this priority**: This is the only story; it is the entire purpose of the smoke run.

**Independent Test**: After the run, inspect `e2e/markers/smoke-full-pipeline.md` on the default branch and the issue's board status.

**Acceptance Scenarios**:

1. **Given** the marker file exists with prior marker lines, **When** the pipeline completes, **Then** the file's last line is `<!-- smoke-full-pipeline-20261026-114941 -->` and all prior lines are unchanged.
2. **Given** the marker file does not exist, **When** the pipeline completes, **Then** the file exists and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** Validate completes, **Then** the linked PR is merged and the issue reaches Done without human intervention.

---

### Edge Cases

- The marker file's last line lacks a trailing newline: the new marker must still land on its own line.
- The marker line is already present (re-run): it must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- smoke-full-pipeline-20261026-114941 -->`, at the end of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it MUST be created containing that line.
- **FR-003**: The change MUST NOT modify any other line of that file or any other file (aside from the engine-managed spec file for this issue).
- **FR-004**: The work MUST be confined to this single repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-005**: The change MUST be delivered through a PR that closes this issue and is merged.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The default branch's `e2e/markers/smoke-full-pipeline.md` ends with `<!-- smoke-full-pipeline-20261026-114941 -->`, appearing exactly once.
- **SC-002**: The PR's diff touches only that one file with a single added line (plus the engine-committed spec file).
- **SC-003**: The issue reaches Done with its PR merged, with no child issues spawned and no human intervention.

## Assumptions

- The marker file normally exists already; creating it is only a fallback.
- The marker is an HTML comment so it has no rendered effect in Markdown.
- `fabrik:yolo` provides all required approvals, including auto-advance past Specify and auto-merge.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Cross-repo work or decomposition into child issues.
- Changes to pipeline behavior or tests.
