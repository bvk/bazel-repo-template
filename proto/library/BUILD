load("@rules_proto//proto:defs.bzl", "proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@com_github_grpc_grpc//bazel:cc_grpc_library.bzl", "cc_grpc_library")

proto_library(
    name = "library",
    srcs = [
        "library.proto",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "@com_google_googleapis//google/api:annotations_proto",
        "@com_google_googleapis//google/api:resource_proto",
    ],
)

go_proto_library(
    name = "library_go_proto",
    compiler = "@io_bazel_rules_go//proto:go_grpc",
    importpath = "github.com/bvk/bazel-repo-template/proto/library",
    proto = ":library",
    visibility = ["//visibility:public"],
    deps = ["@com_google_googleapis//google/api:annotations_go_proto"],
)

go_library(
    name = "libgo",
    embed = [":library_go_proto"],
    importpath = "github.com/bvk/bazel-repo-template/proto/library",
    visibility = ["//visibility:public"],
)

cc_proto_library(
    name = "library_cc_proto",
    visibility = ["//visibility:public"],
    deps = [":library"],
)

cc_grpc_library(
    name = "library_cc_grpc",
    srcs = [":library"],
    grpc_only = True,
    deps = [":library_cc_proto"],
)

cc_library(
    name = "libcc",
    visibility = ["//visibility:public"],
    deps = [":library_cc_proto"],
)
