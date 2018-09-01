package main

import "fmt"

func main() {
	var mf map[string]func() int
	mf = map[string]func() int{
		"m1": func() int { return 10 },
		"m2": func() int { return 20 },
	}

	fmt.Println(mf["m2"]())

	if _, isPresent := mf["m1"]; isPresent {
		fmt.Println(mf["m1"]())
	}

	delete(mf, "m1")

	if _, isPresent := mf["m1"]; isPresent {
		fmt.Println("key is present")
	}

}
