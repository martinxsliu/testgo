.PHONY: build
build: gazelle
	bazel build //...

.PHONY: test
test: gazelle
	bazel test //...

.PHONY: gazelle
gazelle:
	bazel run //:gazelle

.PHONY: cgo
cgo:
	gcc -c -o cgo/cgo.o cgo/cgo.c
	ar rcs cgo/cgo.a cgo/cgo.o
