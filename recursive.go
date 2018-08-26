package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
