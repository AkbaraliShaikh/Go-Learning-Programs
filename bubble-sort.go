package main

import "fmt"

func main() {

	arr := []int{2, 4, 1, 6, 8, 4, 8, 3, 5, 7, 9, 10, 30, 25, 22, 26}

	bubbleSort(arr)

	for _, item := range arr {
		fmt.Println(item)
	}
}

func bubbleSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := 0; j < len(a)-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
