package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("x = %v \ny = %v\n", x, y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("both less than 4")
	case x > 6 && y > 6:
		fmt.Println("both greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x is >= 4 and <= 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("none of the previous cases were met")
	}
}
