load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "hello_go",
    embed = [":hello_lib"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "hello_lib",
    srcs = ["main.go"],
    importpath = "github.com/bvk/bazel-repo-template/go/cmd/hello",
    visibility = ["//visibility:private"],
)
