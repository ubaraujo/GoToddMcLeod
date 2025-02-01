package main

import (
	"fmt"
	"math"
)

func main() {
	s := square{
		length: 1.2,
		width:  3.5,
	}

	c := circle{
		radius: 5.0,
	}

	info(s)
	info(c)

	fmt.Println(s.area())
	fmt.Println(c.area())
}

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.length * s.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}
