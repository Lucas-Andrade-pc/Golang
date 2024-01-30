package main

import "fmt"

func main() {
	var t string = "string"
	discoverType(t)
}
func discoverType(t any) {
	switch t.(type) {
	case string:
		fmt.Printf("Type: %s\n", t)
	}
}
