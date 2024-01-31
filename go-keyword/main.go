package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("One\n")
	c := make(chan string)
	go testFunction(c)
	finish := <-c
	fmt.Printf("Two \n")
	fmt.Printf("Finish: %v\n", finish)

}

func testFunction(c chan string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Checking...%d\n", i)
		time.Sleep(1 * time.Second)
	}
	c <- "Finsih"
}
