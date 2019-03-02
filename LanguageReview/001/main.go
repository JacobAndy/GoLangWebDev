package main

import "fmt"

// declares variable, and assigns it that values 0 value
var x int

/**

PACKAGE LEVEL SCOPE

**/

type person struct {
	// upper case means this will be visible outside of the package
	// Fname string
	// Lname string
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

// polymorphism, take in many different types and it can change depending on what it takes in
func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good Morning Jacob"`)
}
func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

func main() {
	/*
		block level scope
	*/
	x = 7
	fmt.Printf("%T\n", x)

	//slice
	xi := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(xi)

	//map
	// maps are for key values and quick look ups by key
	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	nP := person{fname: "Jacob", lname: "Anderson"}
	// can also be done this way:
	// person{"Jacob", "Anderson"}
	fmt.Println(nP)

	// nP.speak()
	saySomething(nP)

	sa1 := secretAgent{
		person: person{
			fname: "James",
			lname: "Bond",
		},
		licenseToKill: true,
	}
	fmt.Println(sa1)

	//prints James Bond says, "Shaken, not stirred."
	// since we are referring to the type connected to secretAgent
	// sa1.speak()

	// prints James Bond says, "Good Morning Jacob"
	// since we are referring to the type connected to person
	// sa1.person.speak()
	saySomething(sa1)
}
