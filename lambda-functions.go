package main

import "fmt"

func main() {

	func() {
		sum := 0
		for i := 1; i <= 5; i++ {
			sum += i
		}
		fmt.Println("Sum : ", sum)
	}()

	f()
}

func f() {
	for i := 0; i < 10; i++ {
		g := func(i int) {
			fmt.Printf("%d ", i)
		}
		g(i)
		fmt.Printf(" g is of type %T and value is %v\n", g, g)
	}
}
