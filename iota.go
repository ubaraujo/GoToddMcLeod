package main

import "fmt"

/*
	->IOTA é um operador utilizado apenas no bloco de definição de constantes
	->é um const inteiro não tipado
	->seu valor é o índice da sua declaração na respectiva ConstSpec. COMEÇA EM 0(ZERO)
	->como é inteiro, incrementa de 1 em 1 de acordo com a ConstSpec
*/

const (
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)

const (
	c3 = iota // c0 == 0
	c4
	c5
	c6
)

const (
	_ = iota // c0 == 0
	a
	b
	c
	d
	e
	f
)

const (
	sunday    = iota // sunday == 0
	monday           // sunday == 1
	tuesday          // sunday == 2
	wednesday        // sunday == 3
	thurday          // sunday == 4
	friday           // sunday == 5
	saturday         // sunday == 6
)

func main() {
	fmt.Println(c0, c1, c2)
	fmt.Println(c3, c4, c5, c6)
	fmt.Println()
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b\n", 1<<6, 1<<6)
	fmt.Println()
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)
	fmt.Println()
	fmt.Println(sunday)
	fmt.Println(saturday)
}
