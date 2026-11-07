# Feature Specification: E2E smoke full-pipeline marker (20261107-193650)

**Feature Branch**: `fabrik/issue-10798`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261107-193650 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. Its purpose is to confirm that a single-repo issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR. The code change is deliberately trivial so that any failure points at the pipeline, not at the change. The `fabrik:yolo` label is set, so stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged (Priority: P1)

An operator running the smoke test wants a one-line marker added to a known file so that a merged PR proves the pipeline ran end to end.

**Why this priority**: It is the only deliverable of the issue. The marker line is the observable artifact of a successful run.

**Independent Test**: After the PR merges, read `e2e/markers/smoke-full-pipeline.md` on the default branch and check that its last line is `<!-- smoke-full-pipeline-20261107-193650 -->`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the marker line is appended at the very end and all existing content is unchanged.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the change is merged, **When** the issue is inspected, **Then** it is in Done and its PR is merged.

---

### Edge Cases

- The file exists but does not end with a newline: the marker must go on its own line, not joined to the previous content.
- The marker line is already present (for example on a re-run): it must not be appended a second time.
- The `e2e/markers/` directory does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` ends with the exact line `<!-- smoke-full-pipeline-20261107-193650 -->`, appended at the very end of the file.
- **FR-002**: If the file does not exist, it is created, along with any missing parent directory, and contains the marker line.
- **FR-003**: The marker line appears exactly once in the file.
- **FR-004**: The change touches only `e2e/markers/smoke-full-pipeline.md`, with one line added and nothing else changed.
- **FR-005**: The work stays within this single repository. Plan does not decompose it and no child issues are spawned.
- **FR-006**: The issue's PR links the issue with `Closes #10798` and merges, and the issue reaches Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged PR diff contains exactly one changed file with exactly one added line, the marker.
- **SC-002**: The issue reaches Done with a merged PR, with no manual intervention beyond the initial labels.
- **SC-003**: No child issues are created for this issue.

## Assumptions

- The marker is a self-contained HTML comment, so it does not alter how the markdown file renders.
- Whether the file already exists is unknown at specification time. The implementer handles both cases, per FR-002.
- The default base branch `main` is the target, with no `base:<branch>` override.

## Out of Scope *(optional)*

- Any change to a file other than `e2e/markers/smoke-full-pipeline.md`.
- Cross-repo work and decomposition into child issues.
- Changes to engine, stage, or documentation behavior.
- Adding tests for the marker file.
