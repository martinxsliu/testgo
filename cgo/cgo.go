package cgo

// #cgo CFLAGS: -I${SRCDIR}
// #cgo LDFLAGS: ${SRCDIR}/cgo.a
//
// #include <cgo.h>
import "C"

// Using the below and removing the `cdeps` attr will work.

/*
int ten() {
	return 10;
}
*/
//import "C"

func Ten() int {
	return int(C.ten())
}
