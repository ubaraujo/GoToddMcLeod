package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250) // rand.Intn() the max value is non-inclusive

	fmt.Printf("Var x = %v \t", x)

	if x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("between 201 and 250")
	}

}
