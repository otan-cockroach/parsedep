workspace(name = "com_github_cockroach_parsedep")


load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8a993815a6a7cfe47e46238383f0e36a80b9ce0ac482f855c7288a31565b5661",
    strip_prefix = "cockroachdb-rules_go-58cb947",
    urls = [
        # cockroachdb/rules_go as of 58cb94707783f529462d4c6d3c698933a1022942
        # (upstream release-0.29 plus a few patches).
        "https://storage.googleapis.com/public-bazel-artifacts/bazel/cockroachdb-rules_go-v0.27.0-56-g58cb947.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.17.6")

gazelle_dependencies()

"""
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
git_repository(
  name = "cockroach",
  remote = "https://github.com/cockroachdb/cockroach",
  commit = "fc58be83c2a7824ceee657071160f1cee0b0a615",
  shallow_since = "1650494624 +0000",
)

load("@cockroach//:DEPS.bzl", "go_deps")
go_deps()


http_archive(
    name = "bazel_gomock",
    sha256 = "692421b0c5e04ae4bc0bfff42fb1ce8671fe68daee2b8d8ea94657bb1fcddc0a",
    strip_prefix = "bazel_gomock-fde78c91cf1783cc1e33ba278922ba67a6ee2a84",
    urls = [
        "https://storage.googleapis.com/public-bazel-artifacts/bazel/bazel_gomock-fde78c91cf1783cc1e33ba278922ba67a6ee2a84.tar.gz",
    ],
)
"""
