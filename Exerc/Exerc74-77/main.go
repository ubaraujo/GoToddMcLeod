package main

import "fmt"

// Exercise 75
var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

// Exercise 76
type dog struct {
	first string
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d
}

// Exercise 77
func newName(d dog, s string) string {
	d.first = s
	return d.first
}

func newNameP(d *dog, s string) {
	d.first = s
}

func main() {
	// 74
	var value int = 42
	fmt.Println("Value:", value)
	fmt.Println(&value)

	// 75
	fmt.Printf("Value of a: %v - Type of a: %T\n", a, a)
	fmt.Printf("Value of b: %v - Type of b: %T\n", b, b)
	fmt.Printf("Value of c: %v - Type of c: %T\n", c, c)
	fmt.Printf("Value of d: %v - Type of d: %T\n", d, d)
	fmt.Println(*a, *b, *c, *d)

	// 76
	d1 := dog{"Titan"}
	youngRun(d1)
	d2 := dog{"Padget"}
	youngRun(d2)
	d2 = d2.changeName("Rover")
	youngRun(d2)

	// 77
	fmt.Println("d2 before:", d2.first)
	fmt.Println(newName(d2, "Toto"))
	fmt.Println("d2 after:", d2.first)
	newNameP(&d2, "Toto Novo")
	fmt.Println("d2 after with pointer:", d2.first)
}
