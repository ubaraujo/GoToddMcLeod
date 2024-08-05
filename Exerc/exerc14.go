package main

import "fmt"

var numberVar = 42

const constVar = "teste"

func main() {

	fmt.Printf("numberVar: %v - %T\n", numberVar, numberVar)
	fmt.Printf("numberVar: %v - %T\n", numberVar, numberVar)

	z := 42
	fmt.Printf("Value: %v and type: %T\n", z, z)

}
