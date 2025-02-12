package main

import "fmt"

func main() {
	fmt.Println(Paradise("Hawaii"))
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// Paradise prints out the users input of paradise as a sentence.
func Paradise(loc string) string {
	return "My idea of padradise is " + loc
}
