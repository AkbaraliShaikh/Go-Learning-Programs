package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		defer fmt.Print(i, " ")
	}

	defer fmt.Println("===================")

	onExecuting()
	defer onExecuted()
	fmt.Println("Do Something...")
}

func onExecuting() {
	fmt.Println("Entering...")
}

func onExecuted() {
	fmt.Println("Leaving...")
}
