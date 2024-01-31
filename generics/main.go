package main

import (
	"fmt"
	"reflect"
)

func main() {
	var t1 int = 123
	//var t2 float64 = 120.345
	fmt.Printf("Sun: %d (type: %s)\n", plusOne(t1), reflect.TypeOf(plusOne(t1)))
	//fmt.Printf("Sun: %d\n", plusOne(t2))
}

func plusOne[V int](t V) V {
	return t + 1
}
