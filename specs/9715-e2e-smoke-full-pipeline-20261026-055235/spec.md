# Feature Specification: E2E smoke full-pipeline (20261026-055235)

**Feature Branch**: `fabrik/issue-9715`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261026-055235 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to verify that Fabrik can take an issue through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done, ending with a merged PR. The code change is intentionally trivial so that any failure points to the pipeline rather than the change. The issue carries the `fabrik:yolo` label, so stages are expected to auto-advance and the PR is expected to be auto-merged once Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator runs the smoke issue through the pipeline. When it finishes, the marker file on the default branch ends with the unique smoke-test comment line, and the issue is in Done with its PR merged.

**Why this priority**: This is the whole purpose of the issue. It proves the pipeline works end to end.

**Independent Test**: After the PR merges, read the marker file on the default branch and confirm its last line is the marker comment. Confirm the issue is in Done and the linked PR is merged.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker comment line as its last line.
2. **Given** the marker file already exists with other content, **When** the change is implemented, **Then** the existing content is unchanged and the marker comment line is appended at the very end.
3. **Given** the issue has the `fabrik:yolo` label, **When** each stage completes, **Then** the issue advances to the next stage without manual approval, and the PR is merged when Validate completes.

---

### Edge Cases

- The marker file already contains this exact marker line (for example after a retry): it must not be duplicated.
- The existing file has no trailing newline: the marker must still appear on its own line.
- The Plan stage must not decompose the work or spawn child issues, since the change is a single line in a single file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261026-055235 -->`.
- **FR-002**: If the file does not exist, it MUST be created, and it MUST contain the marker line.
- **FR-003**: If the file exists, its prior content MUST be preserved and the marker line MUST be appended after it, on its own line.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The change MUST touch only that one file and add only that one line (plus any file creation needed).
- **FR-006**: The work MUST be confined to this single repository, with no cross-repo changes.
- **FR-007**: The Plan stage MUST NOT decompose the work into child issues.
- **FR-008**: The issue MUST reach Done with its linked PR merged.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline completes, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch is exactly the marker line, and it occurs exactly once.
- **SC-002**: The merged PR's diff contains changes to exactly one file besides the specification record, and adds exactly one non-empty line to it.
- **SC-003**: The issue ends in Done with a merged PR, and no child issues were created.

## Assumptions

- The path `e2e/markers/smoke-full-pipeline.md` is relative to the repository root.
- A leading newline or blank line may be added only if needed so the marker sits on its own line; this does not count as extra content.
- The engine's own spec-persistence file for this issue is not part of the one-file change described above.
- Auto-advance and auto-merge are provided by the `fabrik:yolo` label, so no manual approvals are needed.

## Out of Scope

- Any change other than appending the single marker line.
- Cross-repo work.
- Decomposing the work into child issues.
- Changes to Fabrik's engine, stage configuration, or documentation.
