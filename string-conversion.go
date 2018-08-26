package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "12"
	an, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("The integer is : ", an)
	}

	x := -10
	s := strconv.Itoa(x)
	fmt.Printf("%T %v", s, s)
}
