package main

import "fmt"

type mystruct struct {
	counter int
}

func main() {
	myInstance := mystruct{}
	finish := make(chan bool)
	for i := 0; i < 5; i++ {
		go func(myInstance *mystruct) {
			myInstance.counter++
			finish <- true
		}(&myInstance)
	}
	for i := 0; i < 5; i++ {
		<-finish
	}
	fmt.Printf("Counter: %d\n", myInstance.counter)
}
