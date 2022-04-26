# Issue context

All the go test files (`*_test.go`) are always marked as "unsynced" (when using bazel project view). Re-starting IDE,
cleaning cache, re-syncing, re-importing bazel project doesn't solve the issue.

Presumably, also by this reason it is not possible to run individual go tests from within the IDE sidebar context menu,
it shows "Nothing here" when clicking on the "play button" next to test case.

It is possible though to run full package test from the BUILD file sidebar context menu.

Issue is still reproducible with following versions:

```shell
IdeaUltimate: 2021.3.3
Platform: Mac OS X 12.1
Bazel plugin: 2022.03.22.0.5-api-version-213
Bazel: 4.2.1
```

# How to reproduce

After cloning this repo is, import the Bazel Project via `File -> Import Bazel Project`.

Import the project using existing [`.bazelproject` file](.ijwb/.bazelproject), sync project via
`Bazel -> Sync -> Sync Project with BUILD files` menu. Restart / clean IDE caches if some dependencies cannot be
resolved.

Navigate to `lib/log` directory, `logger_test.go` is marked (unsynced) in gray.