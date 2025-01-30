package main

import "fmt"

func main() {
	foo() // named function

	// the definition of a anonymous function is func(p params) (r returns) { code }
	func() {
		fmt.Println("Anonymous func ran")
	}() // this is the execution command and the params goes there

	func(s string) {
		fmt.Println("This is an anonymous function showing my name:", s)
	}("Tom")

	// functions as an expressions, passed to variables.
	xPrint := func() {
		fmt.Println("Anonymous func ran")
	}

	yPrint := func(s string) {
		fmt.Println("This is an anonymous func showing my name:", s)
	}

	xPrint()
	yPrint("Tom")

	x := func(x, y int) int {
		return x + y
	}

	y := x(10, 5)
	fmt.Println("Anonymous function passed to a variable: ", y)
}

func foo() {
	fmt.Println("Foo ran")
}
