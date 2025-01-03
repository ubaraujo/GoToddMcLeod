package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   50,
	}

	p2 := person{
		first: "Aloy",
		last:  "Sobeck",
		age:   18,
	}

	p3 := secretAgent{
		person: person{
			first: "Conrad",
			last:  "Harthrop-Vane",
			age:   35,
		},
		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Printf("%T \t %#v\n", p2, p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p3.first, p3.last, p3.age, p3.ltk)

	// Creating Anonimous Structs
	p4 := struct {
		first string
		last  string
		age   int
	}{
		first: "Stuart",
		last:  "Thomas",
		age:   40,
	}
	// comparing the types when printing
	fmt.Printf("%T\n", p1) // prints type main.person.
	fmt.Printf("%T\n", p4) // just prints struct and its composition, because its anonymous
}
