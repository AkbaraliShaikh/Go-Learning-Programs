package main

import "fmt"

func main() {

	fmt.Println("Loop #1")
	for i := 0; i < 5; i++ {
		fmt.Print(" ", i)
	}
	fmt.Println()

	fmt.Println("Loop #2")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ", i)
		}
		fmt.Println()
	}

	str := "GO is a beautiful language!"
	for pos, char := range str {
		fmt.Printf("Character on position %d is : %c \n", pos, char)
	}

	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}
}
