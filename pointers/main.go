package main

import "fmt"

func main() {
	a := "stringA"
	b := []string{"stringB"}
	c := make(map[string]string)
	c["teste"] = "value"

	var ar [6]int = [6]int{7, 1, 2, 4, 53, 8}
	fmt.Println(ar)
	fmt.Printf("size1: %d capacidade1: %d\n", len(ar), cap(ar))
	var ar2 []int = ar[1:3]
	fmt.Println(ar2)
	fmt.Printf("size2: %d capacidade2: %d\n", len(ar2), cap(ar2))
	ar2 = ar2[0 : len(ar2)+2]
	fmt.Printf("size2: %d capacidade2: %d\n", len(ar2), cap(ar2))
	for k := range ar2 {
		ar2[k] += 1
	}
	fmt.Println(ar2)

	ar3 := make([]int, 3, 9)
	fmt.Println("Array 3", ar3)
	fmt.Printf("size1: %d capacidade1: %d\n", len(ar3), cap(ar3))
	testPointerA(&a)
	testPointerB(&b)
	testPointerC(c)
	fmt.Printf("a: %s\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b: %v\n", c)
}

func testPointerA(a *string) {
	*a = "another string"
}

func testPointerB(b *[]string) {
	*b = append(*b, "another stirng")
}

func testPointerC(c map[string]string) {
	c["teste"] = "new value"
	c["teste2"] = "new value2"
}
