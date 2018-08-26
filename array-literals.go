package main

import "fmt"

func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{6, 7, 8, 9, 10}
	var arr3 = []int{1, 2, 3, 4, 5}

	var arrkeyval1 = [5]string{3: "a", 4: "b"}
	var arrkeyval2 = []string{2: "a", 3: "b"}

	fmt.Println("Example arr1")
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at %d is %d \n", i, arr1[i])
	}
	fmt.Println("Example arr2")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Array at %d is %d \n", i, arr2[i])
	}
	fmt.Println("Example arr3")
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("Array at %d is %d \n", i, arr3[i])
	}
	fmt.Println("Example arrkeyval1")
	for i := 0; i < len(arrkeyval1); i++ {
		fmt.Printf("Array at %d is %s \n", i, arrkeyval1[i])
	}
	fmt.Println("Example arrkeyval2")
	for i := 0; i < len(arrkeyval2); i++ {
		fmt.Printf("Array at %d is %s \n", i, arrkeyval2[i])
	}
}
