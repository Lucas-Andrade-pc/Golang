package main

import "fmt"

func main() {
	var t string = "string"
	var t2 *string = &t
	discoverType(t2)
	var t3 = 123
	discoverType(t3)
	discoverType(nil)
}
func discoverType(t any) {
	switch v := t.(type) {
	case string:
		result := v + "..."
		fmt.Printf("Type: %v\n", result)
	case *string:
		fmt.Printf("Type: %v\n", *v)
	case int:
		fmt.Printf("Type: %d\n", v)
	default:
		fmt.Println("Type not found")
	}
}
