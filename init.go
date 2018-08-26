package main

import (
	"fmt"
	"runtime"
)

func init() {
	msg := "Enter a digit, e.g. 3 “+ “or %s to quit."

	if runtime.GOOS == "windows" {
		msg = fmt.Sprintf(msg, "Ctrl+Z, Enter")
	} else { // Unix like
		msg = fmt.Sprintf(msg, "Ctrl+D")
	}

	fmt.Println(msg)
}

func main() {

}
