package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("sum = %d \n", sum(arr[:]))
	makeSample()
}

func sum(a []int) int {
	sum := 0
	for _, item := range a {
		sum += item
	}
	return sum
}

func makeSample() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := make([]int, 5)
	slice2 := make([]int, 5, 10)
	anotherWay := new([100]int)[0:25]

	for i := 0; i < 5; i++ {
		slice1[i] = arr[i]
	}

	fmt.Printf("slice1 len = %d, cap = %d\n", len(slice1), cap(slice1))
	for index, x := range slice1 {
		fmt.Printf("slice1 index = %d, value = %d\n", index, x)
	}

	fmt.Printf("slice2 len = %d, cap = %d\n", len(slice2), cap(slice2))
	for index, x := range slice2 {
		fmt.Printf("slice2 index = %d, value = %d\n", index, x)
	}

	fmt.Printf("anotherWay len = %d, cap = %d\n", len(anotherWay), cap(anotherWay))
	for index, x := range anotherWay {
		fmt.Printf("anotherWay index = %d, value = %d\n", index, x)
	}
}
