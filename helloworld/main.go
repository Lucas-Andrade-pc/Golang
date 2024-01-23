package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: go run main.go <arguments>\n")
		os.Exit(1)
	}

	fmt.Printf("Hello World\nos.Arg: %v\nArguments: %v\n", args[0], args[1])
}
