package main

import "fmt"

func main() {

	s, i, f := "Hapiness", 42, 42.42
	fmt.Printf("%v is of type %T\n", s, s)
	fmt.Printf("%v is of type %T\n", i, i)
	fmt.Printf("%v is of type %T\n", f, f)

	x, y, z := 747, 911, 90210

	fmt.Println("Decimal \t binary \t\t Hexadecimal")
	fmt.Printf("%d \t\t %b \t\t %#X\n", x, x, x)
	fmt.Printf("%d \t\t %b \t\t %#X\n", y, y, y)
	fmt.Printf("%d \t\t %b \t %#X\n", z, z, z)

	var a int8 = 127  // -128 to 127
	var b uint8 = 255 // 0 to 255

	fmt.Println("a: ", a, "overflow: ", a+1)
	fmt.Println("b: ", b, "overflow: ", b+1)
}
