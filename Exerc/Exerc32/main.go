package main

import "fmt"

func main() {
	for i := 0; i < 10; {
		fmt.Println("i: ", i)
		i++
		if i > 5 {
			break
		}
	}
}
