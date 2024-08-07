package main

import "fmt"

func main() {
	for i := range 5 {
		for j := range 5 {
			fmt.Printf("Outer Loop: %v | Inner Loop: %v\n", i, j)
		}
	}
}
