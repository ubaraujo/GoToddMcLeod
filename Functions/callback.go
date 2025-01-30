package main

import "fmt"

func main() {
	// passing a function in as an argument to another function
	x := doMath(42, 10, add)
	fmt.Println(x)

	y := doMath(42, 5, sub)
	fmt.Println(y)
}

func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
