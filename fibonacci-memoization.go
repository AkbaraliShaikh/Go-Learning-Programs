package main

import (
	"fmt"
	"time"
)

const LIM = 50

var fibs [LIM]int64

func main() {

	start := time.Now()

	for i := 0; i < LIM; i++ {
		fmt.Printf("fib(%d) is: %d\n", i, fib(i))
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("fib took this amount of time: %s\n", delta)
}

func fib(n int) (res int64) {

	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 0 {
		res = 0
	} else if n == 1 {
		res = 1
	} else {
		res = fib(n-1) + fib(n-2)
	}
	fibs[n] = res
	return
}
