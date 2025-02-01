package main

import "fmt"

func main() {
	arrTest := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(arrTest...))
	fmt.Println(bar(arrTest))

	// erxerc 60
	defer fmt.Println("Testing defer 3")
	defer fmt.Println("Testing defer 2")
	defer fmt.Println("Testing defer 1")
	fmt.Println("Testing defer 0")

	// exerc 61
	p1 := Person{
		first: "Tom",
		age:   33,
	}
	p1.speak()
}

// exerc 61
type Person struct {
	first string
	age   int
}

func (p Person) speak() {
	fmt.Printf("My name is %v, I'm %v years old...\n", p.first, p.age)
}

func foo(nums ...int) int {
	t := 0
	for _, v := range nums {
		t += v
	}
	return t
}

func bar(nums []int) int {
	t := 0
	for _, v := range nums {
		t += v
	}
	return t
}
