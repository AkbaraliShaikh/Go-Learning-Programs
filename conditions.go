package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 1
	y := 10

	if x < y {
		fmt.Println("y is greater")
	} else {
		fmt.Println("x is greater")
	}

	if x = 11; x < y {
		fmt.Println("y is greater")
	} else {
		fmt.Println("x is greater")
	}

	flag := true
	if flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	str := "string"
	if len(str) == 0 {
		fmt.Println("string is empty")
	} else {
		fmt.Println("Count is :", len(str))
	}

	if runtime.GOOS == "windows" {
		fmt.Println(runtime.GOOS)
	} else {
		fmt.Println(runtime.GOOS)
	}

	isNumber := 100
	switch isNumber {
	case 98, 99:
		fmt.Println(isNumber)
	case 100:
		fmt.Println(isNumber)
	default:
		fmt.Println("Default")
	}
}
