package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := range 42 {

		fmt.Printf("Iteration %v - ", i)

		x := rand.Intn(5)

		switch x {
		case 0:
			fmt.Println("value is zero")
		case 1:
			fmt.Println("value is one")
		case 2:
			fmt.Println("value is two")
		case 3:
			fmt.Println("value is three")
		case 4:
			fmt.Println("value is four")
		}
	}
}
