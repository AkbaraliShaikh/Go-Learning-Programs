package main

import "fmt"

func main() {

	var arr [10]int

	fmt.Println("Example #1")
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
		fmt.Println(arr[i])
	}

	fmt.Println("Example #2")
	for i := range arr {
		fmt.Println(arr[i])
	}

	fmt.Println("Example #3")
	strArray := [...]string{"a", "b", "c", "d", "e"}
	for i := range strArray {
		fmt.Println(strArray[i])
	}
}
