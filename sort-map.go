package main

import (
	"fmt"
	"sort"
)

var (
	unsortedArray = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5}
)

func main() {

	fmt.Println("unsorted array")
	for key, _ := range unsortedArray {
		println(key)
	}

	keys := make([]string, len(unsortedArray))
	i := 0
	for key, _ := range unsortedArray {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	fmt.Println("sorted array")
	for _, key := range keys {
		fmt.Printf("key = %s, val = %d\n", key, unsortedArray[key])
	}
}
