package main

import (
	"fmt"
)

func main() {

	res := min(10, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	arr2 := []int{4, 3, 6, 8, 3, 5, 2, 10}

	fmt.Println("min #1")
	fmt.Println(res)

	fmt.Println("min #2")
	fmt.Println(min(arr2...))
}

func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}

	min := a[0]
	for _, arr := range a {
		if arr < min {
			min = arr
		}
	}
	return min
}
