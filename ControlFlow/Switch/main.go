package main

import "fmt"

func main() {
	// SEQUENCE
	fmt.Println("This is the first statement to run")
	fmt.Println("This is the second statement to run")
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement to run
	fmt.Printf("x = %v \ny = %v\n", x, y)

	// CONDITIONAL and LOGICAL OPERATORS
	if x < 42 && y < 42 { // && requires both to be true to run
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || x < 42 { // || requires one of them to be true to run
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 && y != 10 {
		fmt.Println("x is not 42 & y is not 10")
	}

	// SWITCH STATEMENTS
	switch {
	case x < 42:
		fmt.Println("x is LESS THAN 42")
	case x == 42:
		fmt.Println("x is EQUAL to 42")
	case x > 42:
		fmt.Println("x is GREATER THAN 42")
	default:
		fmt.Println("this is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("this is the default case for x")
	}

	//no default fallthrough
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("x is 41")
		fallthrough
	case 42:
		fmt.Println("x is 42")
		fallthrough
	case 43:
		fmt.Println("x is 43")
		fallthrough
	default:
		fmt.Println("printed because ALL OF THE fallthrough statements: this is the default case for x")
	}
}
