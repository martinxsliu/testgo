package a_test

import (
	"testing"

	"github.com/martinxsliu/testgo/b"
	// Uncommenting this line and running `make test` will work.
	// _ "github.com/martinxsliu/testgo/cgo"
)

func TestB(t *testing.T) {
	if b.Ten() != 10 {
		t.Fail()
	}
}
