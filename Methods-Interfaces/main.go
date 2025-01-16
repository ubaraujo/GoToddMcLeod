package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

type human interface {
	speak()
}

// a method that receives a value of type human and calls the method speak() from the interface
func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person{
			first: "Jason",
		},
		true}

	p2 := person{
		first: "Freddy",
	}

	sa1.speak()
	p2.speak()

	saySomething(sa1)
	saySomething(p2)
}
