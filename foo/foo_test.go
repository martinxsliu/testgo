package foo_test

import (
	"fmt"
	"testing"

	"github.com/martinxsliu/testgo/foo"
)

func TestGreeting(t *testing.T) {
	greeter := foo.Hello{}
	got := greeter.Greet()
	expected := fmt.Sprintf("Hello %s", foo.Name)
	if got != expected {
		t.Errorf("greeting is wrong. got %s, expected %s", got, expected)
	}
}
