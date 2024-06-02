package main

import "fmt"

func main() {
	const name, age = "Kim", 22

	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Printf("%s is %d years old. \n And the type is %T and %T \n", name, age, name, age)
}
