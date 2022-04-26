go-mod-tidy:
	go mod tidy

go-sync: go-mod-tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=go_deps.bzl%go_dependencies -prune
	bazel run //:gazelle