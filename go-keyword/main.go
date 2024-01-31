package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("One ")
	c := make(chan string)
	go testFunction(c)
	finish := <-c
	fmt.Printf("Two ")
	fmt.Printf("Finish: %v", finish)

}

func testFunction(c chan string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Checking...")
		time.Sleep(1 * time.Second)
	}
	c <- "Finsih"
}
