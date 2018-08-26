package main

import "fmt"

func main() {
	var arr [3]int

	fmt.Println("================")
	foo1(arr)

	fmt.Println("================")
	foo2(&arr)

	fmt.Println("================")
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}
}

func foo1(a [3]int) {
	fmt.Println(a)
}

func foo2(a *[3]int) {
	fmt.Println(a)
}

func fp(a *[3]int) {
	fmt.Println(a)
}
