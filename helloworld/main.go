package main

import (
	"fmt"
	"os"
)

func main() {
	v := make(map[string]int)
	v["k1"] = 7

	fmt.Println("maps: ", v, len(v))
	delete(v, "k1")
	fmt.Println("maps: ", v, len(v))
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("Usage: go run main.go <arguments>\n")
		os.Exit(1)
	}

	fmt.Printf("Hello World\nos.Arg: %v\nArguments: %v\n", args[0], args[1])
}
