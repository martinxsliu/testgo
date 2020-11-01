load("@io_bazel_rules_go//go:def.bzl", "go_context")

def _gen_impl(ctx):
    # See: https://github.com/bazelbuild/rules_go/blob/master/go/toolchains.rst#go_context
    go = go_context(ctx)

    args = ctx.actions.args()
    args.add("--output", ctx.outputs.out.path)

    ctx.actions.run(
        outputs = [ctx.outputs.out],
        executable = ctx.executable._builder,
        arguments = [args],
        mnemonic = "GoGen",
        progress_message = "GoGen {}".format(ctx.label),
        use_default_shell_env = True,
    )

    return [DefaultInfo(files = depset([ctx.outputs.out]))]

gen = rule(
    implementation = _gen_impl,
    attrs = {
        "out": attr.output(
            mandatory = True,
        ),
        "_builder": attr.label(
            default = "//gen",
            executable = True,
            cfg = "exec",
        ),
        "_go_context_data": attr.label(
            default = "@io_bazel_rules_go//:go_context_data",
        ),
    },
    toolchains = ["@io_bazel_rules_go//go:toolchain"],
)
