.PHONY: build
build:
	NAME=World bazel build //...

.PHONY: test
test:
	NAME=World bazel test //...

.PHONY: gazelle
gazelle:
	bazel run //:gazelle
