package main

import "fmt"

func main() {
	s1_from := []int{1, 2, 3}
	s2_to := make([]int, 10)

	// copy dest = src
	fmt.Println("# copy example #")
	n := copy(s2_to, s1_from)

	fmt.Printf("Copied %d elements\n", n)
	fmt.Println(s2_to)

	// append
	fmt.Println("# append example #")
	s3 := []int{1, 2, 3}
	s3 = append(s3, 4, 5, 6)
	fmt.Println(s3)

	fmt.Println("# resize and append slices example #")
	result := appendByte([]byte{1, 2, 3}, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(result)
}

func appendByte(slice1 []byte, data ...byte) []byte {
	m := len(slice1)
	n := m + len(data)
	if n > cap(slice1) {
		newSlice := make([]byte, n+1)
		copy(newSlice, slice1)
		slice1 = newSlice
	}
	slice1 = slice1[0:n]
	copy(slice1[m:n], data)
	return slice1
}
