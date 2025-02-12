package main

import (
	"fmt"
	"math"
	"time"
)

// 68 - create an anonymous function
func main() {
	//68
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	// 69
	test := func(s string) string {
		return fmt.Sprintf("My name is: %s\n", s)
	}
	fmt.Println(test("Tomzera"))

	// 70
	f := outer()
	fmt.Println(f())

	// 71
	fmt.Println(printSquare(square, 3)) // this is a callback

	// 72
	x := powinator(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())

	// 73
	timeFunc(doWork)
}

// 70
func outer() func() int {
	return func() int {
		return 42
	}
}

// 71 callback
func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the number %v squared is %v\n", a, x)
}

// 72 closure
func powinator(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}

// 73 wrapper
func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func doWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Println(i)
	}
}
