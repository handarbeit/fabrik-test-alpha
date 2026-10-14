# Feature Specification: e2e smoke full-pipeline (20261014-060957)

**Feature Branch**: `fabrik/issue-8110`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261014-060957 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can move through every stage (Specify → Research → Plan → Implement → Review → Validate) to Done with a merged PR. The code change is deliberately trivial, so that any failure points at the pipeline and not at the work. The `fabrik:yolo` label is applied so the stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended and merged through the full pipeline (Priority: P1)

An operator files this issue with `fabrik:yolo`. Fabrik takes it through all stages unattended. The result is a merged PR that adds one marker comment line to the end of `e2e/markers/smoke-full-pipeline.md`.

**Why this priority**: This is the only story. It is the whole purpose of the smoke run.

**Independent Test**: After the run, the issue is in Done and closed, its PR is merged, and the last line of `e2e/markers/smoke-full-pipeline.md` on the base branch is `<!-- smoke-full-pipeline-20261014-060957 -->`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists, **When** the change is implemented, **Then** the line `<!-- smoke-full-pipeline-20261014-060957 -->` is appended as the new last line and the existing content is unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances to the next stage without human input, and the PR is merged once Validate completes.
4. **Given** the PR is merged, **When** the pipeline finishes, **Then** the issue is in Done.

---

### Edge Cases

- The file already ends without a trailing newline: the marker must still land on its own final line, not be glued onto the previous line.
- The marker line is already present (for example, a re-run): the change must not add a duplicate line.
- The Plan stage must not split the work into child issues. The change is one file and one line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- smoke-full-pipeline-20261014-060957 -->`, at the very end of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it is created, and it contains the marker line.
- **FR-003**: The change touches no file other than `e2e/markers/smoke-full-pipeline.md`. The one exception is the spec file the engine writes under `specs/` during Specify.
- **FR-004**: Any pre-existing content of the marker file is preserved unchanged.
- **FR-005**: The work is confined to this single repository, and the Plan stage does not decompose it or spawn child issues.
- **FR-006**: The change reaches the base branch through a PR that links this issue, and the issue ends in Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged, without human intervention beyond applying the `fabrik:yolo` label.
- **SC-002**: On the base branch after merge, the final line of `e2e/markers/smoke-full-pipeline.md` equals `<!-- smoke-full-pipeline-20261014-060957 -->`, and that line appears exactly once.
- **SC-003**: The merged PR's diff contains one added content line in one file, apart from the engine-written spec file.
- **SC-004**: No child issues are created.

## Assumptions

- The `e2e/markers/` directory may or may not exist. If it does not, it is created along with the file.
- A markdown HTML comment is an acceptable marker format, since the issue specifies it.
- The timestamp suffix `20261014-060957` identifies this smoke run and is used literally.
- The engine-written `specs/8110-*/spec.md` file in the PR is expected pipeline behavior, not scope creep.

## Out of Scope *(optional)*

- Cross-repo work.
- Any change other than the single marker line.
- Decomposing the work into child issues.
- Changes to Fabrik engine code, stage configs, or docs.
