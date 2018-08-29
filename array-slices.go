package main

import "fmt"

func main() {

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	arr2 := &arr

	// arr[start : end] // end index not included!
	slice1 := arr[5:]
	slice2 := arr[:5]
	slice3 := arr[5:9]
	slice4 := arr[:]

	fmt.Println(arr)
	fmt.Println(arr2)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	fmt.Printf("arr len = %d and cap = %d \n", len(arr), cap(arr))
	fmt.Printf("arr2 len = %d and cap = %d \n", len(arr2), cap(arr2))

	fmt.Printf("slice1 len = %d and cap = %d \n", len(slice1), cap(slice1))
	fmt.Printf("slice2 len = %d and cap = %d \n", len(slice2), cap(slice2))
	fmt.Printf("slice3 len = %d and cap = %d \n", len(slice3), cap(slice3))
	fmt.Printf("slice4 len = %d and cap = %d \n", len(slice4), cap(slice4))
}
