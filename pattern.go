package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"

	if ok, _ := regexp.MatchString(pat, searchIn); ok {
		fmt.Println("match found.")
	}

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		// multiply the value by 2
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	re, _ := regexp.Compile(pat)

	str := re.ReplaceAllString(searchIn, "##.##")
	fmt.Println(str)

	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}