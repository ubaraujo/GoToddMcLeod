package main

import "fmt"

func main() {
	foo()
	bar("Tomzera")
	s := aloha("Todd")
	fmt.Println(s)
	fmt.Println(dogYears("Ugleiston", 33))

	//variadic parameters
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(numeros...))

}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func aloha(s string) string {
	return fmt.Sprintln("They call me Mr. ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years:"), age
}

// VARIADIC PARAMETERS

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}
