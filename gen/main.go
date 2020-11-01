package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	output := flag.String("output", "", "output path")
	flag.Parse()

	name, ok := os.LookupEnv("NAME")
	if !ok {
		panic("NAME not found!")
	}

	err := ioutil.WriteFile(*output, []byte(fmt.Sprintf("package foo\nconst Name = \"%s\"", name)), 0644)
	if err != nil {
		panic(err)
	}
}
