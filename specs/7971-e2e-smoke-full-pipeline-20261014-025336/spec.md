# Feature Specification: E2E smoke full-pipeline marker (20261014-025336)

**Feature Branch**: `fabrik/issue-7971`
**Created**: 2026-10-13
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261014-025336 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This is an end-to-end smoke test of the Fabrik pipeline on a single repository. The goal is to confirm that an issue can progress through every stage (Specify → Research → Plan → Implement → Review → Validate) to Done with its PR merged. The code change is deliberately trivial so that any failure points at the pipeline rather than the change. The issue carries `fabrik:yolo`, so stages auto-advance and the PR auto-merges on Validate completion.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the smoke marker line (Priority: P1)

An operator running the release smoke wants a unique, timestamped marker line appended to the smoke-full-pipeline marker file, delivered through a merged PR, to prove the pipeline works end to end.

**Why this priority**: It is the only change; the pipeline run depends on it.

**Independent Test**: After merge, the marker file on the default branch ends with the line `<!-- smoke-full-pipeline-20261014-025336 -->`, and the diff of the PR touches only that file with a single added line.

**Acceptance Scenarios**:

1. **Given** the marker file exists with prior marker lines, **When** the change is applied, **Then** the new marker line is the last line of the file and all existing lines are unchanged.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created containing the new marker line.
3. **Given** the PR is merged, **When** the issue is checked, **Then** it is in Done and its PR is merged.

---

### Edge Cases

- The file's last line lacks a trailing newline: the new marker must still start on its own line.
- The marker line already exists in the file (e.g. a re-run): it must not be duplicated.
- Concurrent smoke runs append to the same file: a conflicting rebase must be resolved while preserving all marker lines.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261014-025336 -->`, appended after all existing content.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: Existing content of the file MUST NOT be modified or removed.
- **FR-004**: The change MUST consist of exactly one file and one added line, in this single repository.
- **FR-005**: The Plan stage MUST NOT decompose the work into child issues.
- **FR-006**: The change MUST be delivered via a PR that closes this issue and is merged, with the issue ending in Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged, with no human intervention beyond the `fabrik:yolo` label.
- **SC-002**: The merged PR's diff is exactly one added line in one file.
- **SC-003**: No child issues are spawned.

## Assumptions

- The marker file most likely exists already; the create-if-missing path is a fallback only.
- The marker is an HTML comment so it has no visible effect on rendered Markdown.
- No documentation or test changes are needed for a marker-file edit.

## Out of Scope

- Any cross-repo work.
- Any change to engine code, stage configuration, or documentation.
- Decomposing the work into sub-issues.
- Cleaning up or deduplicating earlier marker lines.
END_PLACEHOLDER
