package main

import "fmt"

func main() {
	for i := range 100 {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println("Odd number:", i)
		}
	}
}
