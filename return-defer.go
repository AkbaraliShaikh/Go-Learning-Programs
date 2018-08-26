package main

import (
	"fmt"
)

func main() {
	fmt.Println("Example #1")
	fmt.Println(foo())

	fmt.Println("Example #2")
	call := add()
	fmt.Println(call(5, 5))

	fmt.Println("Example #3")
	add := adder()
	fmt.Println(add(1))
	fmt.Println(add(10))
	fmt.Println(add(100))
}

func foo() (ret int) {

	defer func() {
		ret++
	}()

	return 1
}

func add() func(a int, b int) int {
	return func(a int, b int) int {
		return a + b
	}
}

func adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
