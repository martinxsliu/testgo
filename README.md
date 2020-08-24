This repo demonstrates an issue with recompiling packages for test if a transitive dependency has a `cdep` dependency.

#### Dependency graph

```
a_test -> b -> a
          └--> cgo -> cgo.a (cc_import)
```

#### `go test ./...`

```
➜  testgo git:(master) go test ./...
ok  	github.com/martinxsliu/testgo/a	0.007s
?   	github.com/martinxsliu/testgo/b	[no test files]
?   	github.com/martinxsliu/testgo/cgo	[no test files]
```

#### `bazel build //...`

```
➜  testgo git:(master) bazel build //...
INFO: Analyzed 5 targets (0 packages loaded, 0 targets configured).
INFO: Found 5 targets...
ERROR: /Users/martin/repos/testgo/a/BUILD.bazel:10:8: GoLink a/go_default_test_/go_default_test failed (Exit 1) builder failed: error executing command bazel-out/host/bin/external/go_sdk/builder link -sdk external/go_sdk -installsuffix darwin_amd64 -arc ... (remaining 21 argument(s) skipped)

Use --sandbox_debug to see verbose messages from the sandbox
external/go_sdk/pkg/tool/darwin_amd64/link: running external/local_config_cc/cc_wrapper.sh failed: exit status 1
clang: error: no such file or directory: 'cgo/cgo.a'
clang: error: no such file or directory: 'cgo/cgo.a'

link: error running subcommand: exit status 2
INFO: Elapsed time: 0.475s, Critical Path: 0.36s
INFO: 0 processes.
FAILED: Build did NOT complete successfully
```

#### Observations

* If `a_test` directly imports `cgo` then the build succeeds.
* If `b` does not import `a` (so that it only imports `cgo`) then the build succeeds.
* It doesn't change anything if `a_test` imports `a` or not.
