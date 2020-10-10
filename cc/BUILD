cc_binary(
    name = "hello_cc",
    srcs = ["hello.cc"],
    visibility = ["//visibility:public"],
)

cc_test(
    name = "library_proto_test",
    srcs = ["library_proto_test.cc"],
    deps = ["//proto/library:libcc"],
)
