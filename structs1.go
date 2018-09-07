package main

import (
	"fmt"
)

type person struct {
	id   int
	name string
	age  int
}

func main() {

	prn := new(person)
	prn.id = 1
	prn.age = 25
	prn.name = "smith"

	prn1 := &person{10, "bob", 22}
	prn2 := person{id: 20, name: "alice", age: 24}

	fmt.Println(prn)
	fmt.Println(prn1)
	fmt.Println(prn2)
}
