package main

import (
	"fmt"
)

func main() {
	fmt.Println(isGreater(10, 20))
	fmt.Println(isSmall(10, 20))
}

func isGreater(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func isSmall(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
