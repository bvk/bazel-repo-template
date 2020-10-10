load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "library_lib",
    srcs = ["main.go"],
    importpath = "github.com/bvk/bazel-repo-template/go/cmd/library",
    visibility = ["//visibility:private"],
    deps = [
        "//go/library",
        "//proto/library:libgo",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "library",
    embed = [":library_lib"],
    visibility = ["//visibility:public"],
)
