package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName, lastName string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {

	// 1. struct as a value type:
	var prsn1 Person
	prsn1.firstName = "foo"
	prsn1.lastName = "bar"
	upPerson(&prsn1)
	fmt.Printf("the name of the person is %s %s \n", prsn1.firstName, prsn1.lastName)

	// 2. struct as a pointer:
	prsn2 := new(Person)
	prsn2.firstName = "foo"
	prsn2.lastName = "bar"
	upPerson(prsn2)
	fmt.Printf("the name of the person is %s %s \n", prsn2.firstName, prsn2.lastName)

	// 3. struct as a literal:
	prsn3 := &Person{"foo", "bar"}
	upPerson(prsn3)
	fmt.Printf("the name of the person is %s %s \n", prsn3.firstName, prsn3.lastName)
}
