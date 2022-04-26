workspace(name = "bazel-gotest-unsynced")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

buildtools_version = "4.2.1"

buildtools_sha256 = "385f84667b904294ba9122e237fcf7a72c9d553ea710c4669843c67f315ad32f"

# buildifier BUILD file linter.
http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = buildtools_sha256,
    strip_prefix = "buildtools-%s" % buildtools_version,
    url = "https://github.com/bazelbuild/buildtools/archive/refs/tags/%s.zip" % buildtools_version,
)

rules_go_version = "v0.31.0"

rules_go_sha265 = "f2dcd210c7095febe54b804bb1cd3a58fe8435a909db2ec04e31542631cf715c"

http_archive(
    name = "io_bazel_rules_go",
    sha256 = rules_go_sha265,
    urls = [
        "https://github.com/bazelbuild/rules_go/releases/download/{version}/rules_go-{version}.zip".format(version = rules_go_version),
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/{version}/rules_go-{version}.tar.gz".format(version = rules_go_version),
    ],
)

gazelle_version = "v0.24.0"

gazelle_sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb"

http_archive(
    name = "bazel_gazelle",
    sha256 = gazelle_sha256,
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/{version}/bazel-gazelle-{version}.tar.gz".format(version = gazelle_version),
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/{version}/bazel-gazelle-{version}.tar.gz".format(version = gazelle_version),
    ],
)

# <Go> -------------------------------------------------------------------------
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    version = "1.18",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:go_deps.bzl", "go_dependencies")

# gazelle:repository_macro go_deps.bzl%go_dependencies
go_dependencies()

gazelle_dependencies()
# </Go> ------------------------------------------------------------------------
