package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("This is the book ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

// Wrapping function into another function to expand functionality(polymorphism)
func logInfo(s fmt.Stringer) {
	log.Println("LOGINFO: ", s.String())
}

func main() {
	b := book{
		title: "West With the Night",
	}

	var c count = 42

	fmt.Println(b)
	fmt.Println(c)

	log.Println(b)
	log.Println(c)

	logInfo(b)
	logInfo(c)

}
