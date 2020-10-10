load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "library",
    srcs = ["library.go"],
    importpath = "github.com/bvk/bazel-repo-template/go/library",
    visibility = ["//visibility:public"],
    deps = ["//proto/library:libgo"],
)

go_test(
    name = "library_test",
    srcs = ["library_test.go"],
    embed = [":library"],
    deps = ["//proto/library:libgo"],
)
