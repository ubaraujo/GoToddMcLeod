package main

import "fmt"

var first int

func main() {
	second, third := "2", "3"
	var idx int = 5

	for idx > first {
		fmt.Println(second, third)
		idx--
	}
}
