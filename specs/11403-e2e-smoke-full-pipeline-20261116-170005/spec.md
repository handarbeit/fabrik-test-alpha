# Feature Specification: e2e smoke full-pipeline (20261116-170005)

**Feature Branch**: `fabrik/issue-11403`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261116-170005 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can progress through every stage (Specify → Research → Plan → Implement → Review → Validate) and end in Done with a merged PR. The code change is deliberately trivial so that any failure is attributable to the pipeline rather than to the work itself. The issue carries the `fabrik:yolo` label, so stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator running the e2e smoke wants a single marker line added to the smoke-test marker file and delivered via a merged PR, with no human intervention beyond the initial label.

**Why this priority**: This is the sole purpose of the issue. A merged PR containing exactly this change proves the pipeline works end to end.

**Independent Test**: After the pipeline finishes, inspect the merged PR's diff and the marker file on the default branch.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and its final line is `<!-- smoke-full-pipeline-20261116-170005 -->`.
2. **Given** the marker file already exists, **When** the change is implemented, **Then** the existing content is unchanged and the marker line is added at the very end of the file.
3. **Given** the change is implemented, **When** the PR is opened, **Then** the PR diff touches only the marker file, and adds only that one line (plus the file's creation, if it was new).
4. **Given** the issue carries `fabrik:yolo`, **When** Validate completes, **Then** the PR is merged and the issue moves to Done.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker line must still be on its own line at the end of the file.
- The marker file already contains markers from earlier smoke runs: they are preserved; only this run's marker is appended.
- The marker line is already present at the end of the file (a re-run): the change is a no-op and must not duplicate the line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261116-170005 -->` after the change is merged.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: If the file exists, its prior content MUST be preserved unchanged, and the marker line MUST be appended at the very end.
- **FR-004**: The change MUST consist of exactly one file and one added line (excluding the file's creation); no other files may be modified.
- **FR-005**: The work MUST be confined to this single repository; the Plan stage MUST NOT decompose it into child issues.
- **FR-006**: The PR MUST reference this issue with a closing reference so Fabrik can link it.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, an append-only record of smoke-run markers.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its PR merged, with no manual intervention after the `fabrik:yolo` label is applied.
- **SC-002**: The merged PR diff contains exactly one changed file, with exactly one added content line.
- **SC-003**: The last line of the marker file on the default branch equals `<!-- smoke-full-pipeline-20261116-170005 -->`, appearing exactly once.
- **SC-004**: No child issues are created by the Plan stage.

## Assumptions

- The repository's default base branch is `main`.
- The marker is an HTML comment so it does not render in the Markdown file.
- If the file must be created, the marker line is its only content.
- Appending the marker is idempotent: if the identical line is already the last line, nothing further is added.

## Out of Scope

- Any cross-repo work.
- Any change to engine code, stage configuration or documentation.
- Decomposing the work into child issues.
- Modifying or cleaning up markers from earlier smoke runs.
