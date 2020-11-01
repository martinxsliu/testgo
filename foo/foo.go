package foo

import (
	"fmt"
)

type Greeter interface {
	Greet() string
}

type Hello struct{}

func (Hello) Greet() string {
	// Variable `Name` is generated by `:gen`.
	return fmt.Sprintf("Hello %s", Name)
}