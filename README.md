# fabrik-test-alpha
<!-- convergence-race-A-20260619-154613 -->
<!-- convergence-race-B-20260619-154613 -->
<!-- convergence-race-A-20260616-115016 -->
<!-- convergence-race-B-20260616-115016 -->
<!-- convergence-race-B-20260526-122826 -->
<!-- convergence-race-A-20260526-122826 -->

Primary test bed for [Fabrik](https://github.com/handarbeit/fabrik), paired with [`fabrik-test-beta`](https://github.com/handarbeit/fabrik-test-beta).

This repository is not a real project. It exists purely as substrate for exercising Fabrik's multi-repo features (cross-repo sub-issue spawn, cross-repo dependency linkage, parallel multi-repo work, etc.) without polluting real project boards.

## Layout

- `main.go` — minimal CLI that will eventually consume `fabrik-test-beta`'s greeting package. Cross-repo Fabrik tests file issues here that require a change in beta first.
- `.specify/` — Spec Kit templates used by Fabrik's customized Specify stage.

## License

Apache 2.0 (matches Fabrik upstream).
<!-- smoke-full-pipeline-20260525-022137 -->
<!-- auto-merge-yolo-20260526-122826 -->
<!-- auto-merge-yolo-20260527-150017 -->
<!-- auto-merge-yolo-20260527-150653 -->

<!-- smoke-full-pipeline-20260616-115016 -->
<!-- auto-merge-yolo-20260616-115016 -->
<!-- smoke-full-pipeline-20260619-154613 -->
<!-- cruise-pipeline-20260619-154613 -->
