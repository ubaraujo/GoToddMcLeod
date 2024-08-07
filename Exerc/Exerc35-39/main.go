package main

import (
	"fmt"
	"math/rand"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for i, v := range xi {
		fmt.Printf("Index: %v Item: %v\n", i, v)
	}

	// Exerc 36
	m := map[string]int{
		"James":   42,
		"Tomzera": 33,
	}
	for k, v := range m {
		fmt.Printf("Key: %v \tValue: %v\n", k, v)
	}

	// Exerc 37
	x := m["James"]
	fmt.Println("James:", x)

	if j, ok := m["James"]; ok {
		fmt.Println("There is a James, and his age is: ", j)
	}

	y := m["Q"]
	fmt.Println("Q:", y)

	if comma, ok := m["Q"]; !ok { //comma ok idiom
		fmt.Println("There's no Q. Comma returned:", comma)
	}

	// Exerc 38
	for i := range 100 {
		if randInt := rand.Intn(5); randInt == 3 {
			fmt.Printf("randInt is %v - Iter: %v\n", randInt, i)
		}
	}

	// Exerc 39
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
