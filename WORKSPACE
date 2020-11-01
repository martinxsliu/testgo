load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "08c3cd71857d58af3cda759112437d9e63339ac9c6e0042add43f4d94caf632d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.24.2/rules_go-v0.24.2.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.24.2/rules_go-v0.24.2.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "cdb02a887a7187ea4d5a27452311a75ed8637379a1287d8eeb952138ea485f7d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_rules_dependencies()
go_register_toolchains()
gazelle_dependencies()

http_archive(
    name = "com_github_jmhodges_bazel_gomock",
    sha256 = "cb532242627e0d2b3c8587abfa3377a72f58323248faf28490cfdada27d29578",
    strip_prefix = "bazel_gomock-dd57d5599d64ebde72aff1c91df41ef32baece51",
    urls = [
        "https://github.com/jmhodges/bazel_gomock/archive/dd57d5599d64ebde72aff1c91df41ef32baece51.zip",
    ],
)

go_repository(
    name = "com_github_golang_mock",
    build_file_proto_mode = "disable",
    commit = "9fa652df1129bef0e734c9cf9bf6dbae9ef3b9fa",
    importpath = "github.com/golang/mock",
)

go_repository(
    name = "org_golang_x_mod",
    build_file_proto_mode = "disable",
    commit = "859b3ef565e237f9f1a0fb6b55385c497545680d",
    importpath = "golang.org/x/mod",
)
