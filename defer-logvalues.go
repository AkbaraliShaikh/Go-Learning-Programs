package main

import (
	"fmt"
	"io"
)

func main() {
	testFunc("Go")
}

func testFunc(s string) (n int, err error) {
	defer func() {
		fmt.Printf("testFunc(%q) = %d, %v", s, n, err)
	}()

	return 7, io.EOF
}
