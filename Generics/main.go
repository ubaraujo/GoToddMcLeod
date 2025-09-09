package main

import (
	"fmt"

	"golang.org/x/exp/constraints" // necessary to dwload this package: "go get golang.org/x/exp/constraints"
)

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

func addT[T int | float64](a, b T) T {
	return a + b
}

type myNumbers interface {
	~int | ~float64
}

func addN[T myNumbers](a, b T) T {
	return a + b
}

type myConstraints interface {
	constraints.Integer | constraints.Float
}

func addC[T myConstraints](a, b T) T {
	return a + b
}

type myAlias int

func main() {
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))

	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.2, 2.2))

	fmt.Println(addT(10, 2.2))
	fmt.Println(addN(10, 5.2))

	var n myAlias = 42
	fmt.Println(addN(n, 2))

	fmt.Println(addC(5.5, 10))
}
