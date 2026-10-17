# Feature Specification: e2e smoke full-pipeline marker line (20261017-183838)

**Feature Branch**: `fabrik/issue-8657`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261017-183838 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged pull request. The content change is deliberately trivial, so that any failure points at the pipeline and not at the work itself. The issue carries the `fabrik:yolo` label, so stages auto-advance and the linked PR is merged automatically once Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator files this issue and applies `fabrik:yolo`. Fabrik takes it through every pipeline stage. The result is a merged PR in which the marker file ends with the unique marker comment line for this run.

**Why this priority**: It is the only story. It is the whole point of the smoke test.

**Independent Test**: After the PR merges, read `e2e/markers/smoke-full-pipeline.md` on the default branch. Its last line is `<!-- smoke-full-pipeline-20261017-183838 -->`. The issue is closed and sits in Done.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the file's content is unchanged except for one new final line, `<!-- smoke-full-pipeline-20261017-183838 -->`.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** every stage completes, **Then** the PR is merged automatically and the issue reaches Done with no manual step.
4. **Given** the issue is single-repo, **When** the Plan stage runs, **Then** it produces one unit of work and spawns no child issues.

---

### Edge Cases

- The file exists but its last line has no trailing newline: the marker must still land on its own line and must not be joined to the previous line.
- The marker line is already present in the file, for example after a retried run: it must not be added a second time.
- Earlier smoke runs have left other marker lines in the file: they are preserved and the new line goes after them.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- smoke-full-pipeline-20261017-183838 -->`, as the final line of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, the change MUST create it, with the marker line as its content.
- **FR-003**: The change MUST NOT modify, reorder or remove any existing content of that file.
- **FR-004**: The change MUST touch only that one file; no other file in the repository is changed.
- **FR-005**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-006**: The marker line MUST NOT be duplicated if it is already present.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff for this issue shows exactly one file changed and exactly one line added.
- **SC-002**: The last line of `e2e/markers/smoke-full-pipeline.md` on the default branch after the merge is `<!-- smoke-full-pipeline-20261017-183838 -->`.
- **SC-003**: The issue progresses from Specify to Done, with its PR merged, without manual intervention.
- **SC-004**: No child or sub-issues are created.

## Assumptions

- The `e2e/markers/` directory may or may not exist yet, and creating it is acceptable if needed.
- The marker is an HTML comment, so it does not change how the Markdown file renders.
- The timestamp suffix `20261017-183838` is the unique identifier of this smoke run and is used exactly as given.
- The default base branch is `main`.

## Out of Scope *(optional)*

- Cross-repository changes.
- Any change to engine, stage or plugin behavior.
- Any file other than `e2e/markers/smoke-full-pipeline.md`.
- Decomposition into child issues.
