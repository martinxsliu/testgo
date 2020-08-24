package b

import (
	// Removing dependency on `a` will make the build succeed.
	"github.com/martinxsliu/testgo/a"
	"github.com/martinxsliu/testgo/cgo"
)

func Ten() int {
	return a.Ten()
}

func CgoTen() int {
	return cgo.Ten()
}
