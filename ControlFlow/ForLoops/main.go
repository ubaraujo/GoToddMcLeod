package main

import "fmt"

func main() {

	x := 40
	y := 5
	fmt.Printf("")
	fmt.Printf("x = %v \ny = %v\n", x, y)

	for i := 0; i < 5; i++ {
		fmt.Printf("counting to five: %v\t first for loop\n", i)
	}

	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	// break
	for {
		fmt.Printf("y is %v \t\t third for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	// continue
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("couting even numbers: ", i)
	}

	// nested loops
	for i := 0; i < 5; i++ {
		fmt.Println("--")
		for j := 0; j < 5; j++ {
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
	}

	// for range loop with slice
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	// for range loop with map
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}

	persons := map[int]string{
		1: "Tom",
		2: "Thata",
	}
	for k, v := range persons {
		fmt.Println("ranging over persons", k, v)
	}

	// for range loop with string
	name := "Tomzeras"
	for i, v := range name {
		fmt.Printf("index = %v \t %v \t", i, v)
		fmt.Println(string(v)) //prints the value of the char
	}
}
