workspace(name = "com_github_bvk_bazel_repo_template")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

#
# Direct dependencies for 3rd party bazel repositories.
#

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go",
    tag = "v0.24.3",
)

git_repository(
    name = "rules_cc",
    remote = "https://github.com/bazelbuild/rules_cc",
    commit = "f95239adde29680236afa22b4abaf1d04234f61a",
)

git_repository(
    name = "rules_java",
    remote = "https://github.com/bazelbuild/rules_java",
    tag = "0.1.1",
)

git_repository(
    name = "rules_proto",
    remote = "https://github.com/bazelbuild/rules_proto",
    commit = "40298556293ae502c66579620a7ce867d5f57311",
)

git_repository(
    name = "rules_pkg",
    remote = "https://github.com/bazelbuild/rules_pkg/pkg",
    tag = "v0.2.6-1",
)

git_repository(
    name = "bazel_gazelle",
    remote = "https://github.com/bazelbuild/bazel-gazelle",
    tag = "v0.22.2",
)

git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker",
    tag = "v0.14.4",
)

git_repository(
    name = "com_google_googleapis",
    remote = "https://github.com/googleapis/googleapis",
    commit = "eabe7c0fde64b1451df6ea171b2009238b0df07c",
)

git_repository(
    name = "com_google_protobuf",
    remote = "https://github.com/google/protobuf",
    tag = "v3.12.0",
)

git_repository(
    name = "io_grpc_grpc_java",
    remote = "https://github.com/grpc/grpc-java",
    tag = "v1.32.1",
)

git_repository(
    name = "com_github_grpc_grpc",
    remote = "https://github.com/grpc/grpc",
    tag = "v1.32.0",
)

#
# Rules/dependencies for Protobufs.
#

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
protobuf_deps()

load("@rules_proto//proto:repositories.bzl",
     "rules_proto_dependencies",
     "rules_proto_toolchains")
rules_proto_dependencies()
rules_proto_toolchains()

#
# Rules/dependencies for Go.
#

load("@io_bazel_rules_go//go:deps.bzl",
     "go_register_toolchains",
     "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains(go_version = "1.15")

#
# Rules/dependencies for C++.
#

load("@rules_cc//cc:repositories.bzl", "rules_cc_dependencies")
rules_cc_dependencies()

#
# Rules/dependencies for C++.
#

load("@rules_java//java:repositories.bzl",
     "rules_java_dependencies",
     "rules_java_toolchains")
rules_java_dependencies()
rules_java_toolchains()

#
# Rules/dependencies for Gazelle.
#

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

load("//:deps.bzl", "go_dependencies")
# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

#
# Rules/dependencies for containers.
#

load("@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)
container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl",
    container_deps = "deps",
)
container_deps()

load("@io_bazel_rules_docker//repositories:pip_repositories.bzl", "pip_deps")
pip_deps()

load("@io_bazel_rules_docker//cc:image.bzl",
    cc_image_repositories = "repositories",
)
cc_image_repositories()

load("@io_bazel_rules_docker//container:container.bzl", "container_pull")

container_pull(
    name = "debian_base",
    digest = "sha256:e5768afa5429b85ac75de67efc98a4bf53e4ef0f7388667fb34c89d481d82b00",
    registry = "gcr.io",
    repository = "distroless/base-debian10",
)

#
# Additional rules/dependencies for third_party/googleapis
#

http_archive(
    name = "gapic_generator_csharp",
    strip_prefix = "gapic-generator-csharp-1.3.1",
    sha256 = "1118448d4c3db2a443e0bc3bf240de749adbae3311703e264d14213041b6f1c3",
    urls = ["https://github.com/googleapis/gapic-generator-csharp/archive/v1.3.1.zip"],
)

load("@gapic_generator_csharp//:repositories.bzl", "gapic_generator_csharp_repositories")
gapic_generator_csharp_repositories()

load("@com_google_googleapis//:repository_rules.bzl", "switched_rules_by_language")
switched_rules_by_language(
    name = "com_google_googleapis_imports",
    cc = True,
    csharp = True,
    gapic = False,
    go = True,
    grpc = True,
    java = True,
    nodejs = True,
    php = True,
    python = True,
    ruby = True,
)

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")
grpc_deps()
load("@com_github_grpc_grpc//bazel:grpc_extra_deps.bzl", "grpc_extra_deps")
grpc_extra_deps()

http_archive(
    name = "gapic_generator_ruby",
    strip_prefix = "gapic-generator-ruby-9cc95d8f4e05bdfe4ea9c67e6ca670f27c01c8f2",
    sha256 = "cfa78149b54f69f991b9f5fd1f5f8c68ead8b8b52679a9bd48afdfb3f5a32e41",
    urls = ["https://github.com/googleapis/gapic-generator-ruby/archive/9cc95d8f4e05bdfe4ea9c67e6ca670f27c01c8f2.zip"],
)

load("@gapic_generator_ruby//rules_ruby_gapic:repositories.bzl",
     "gapic_generator_ruby_repositories")
gapic_generator_ruby_repositories()
