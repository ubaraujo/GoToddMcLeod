package main

import "fmt"

type myInt int

var x myInt
var y int

func main() {
	fmt.Printf("x: %v\ttipo: %T\n", x, x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("y: %v\ttipo: %T\n", y, y)
}
