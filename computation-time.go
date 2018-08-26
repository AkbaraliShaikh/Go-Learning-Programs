package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	doSomething()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("doSomething took this amount of time: %s \n", delta)
}

func doSomething() int {
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum += i
	}
	return sum
}
