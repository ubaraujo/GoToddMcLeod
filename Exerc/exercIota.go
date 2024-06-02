package main

import "fmt"

/*
	O operador >> ou << aumenta ou diminui as casas do valor binário de uma variável:
	Ex: 1 em binário é = 1
	Se aplicarmos o operador 1 << 1
	1 em binário é empurrado para esquerda, e adicionado um zero ao lado == 10
	10 em binário é 2
	10 << 1 == 100 == 4
	ao shiftar mais 1 casa, o 10 binario vira 100 que equivale a 4 em decimal
*/

const (
	_      = iota             // recebe 0
	KB int = 1 << (10 * iota) // KB == 10000000000(binary) == 1024
	MB
	GB
	TB
	PB // at this point iota == 5, so PB == 100000000000000000000000000000000000000000000000000 == 1125899906842624
	EB
)

const (
	_         = iota
	teste int = 1 << iota // 1 em binário é empurrado para esquerda e criado zero no campo novo == 10(binario) == 2(decimal)
	teste2
)

func main() {
	fmt.Printf("Decimal: %d - Binary: %b\n", KB, KB)
	fmt.Printf("Decimal: %d - Binary: %b\n", MB, MB)
	fmt.Printf("Decimal: %d - Binary: %b\n", GB, GB)
	fmt.Printf("Decimal: %d - Binary: %b\n", TB, TB)
	fmt.Printf("Decimal: %d - Binary: %b\n", PB, PB)
	fmt.Printf("Antes do shift, teste valia %d em decimal e valia %b em binário\n", 1, 1)
	fmt.Printf("Após o shift, teste vale: %d em decimal, representado por: %b em binário\n", teste, teste)
	fmt.Printf("Após o shift, teste vale: %d em decimal, representado por: %b em binário\n", teste2, teste2)
	//fmt.Printf("Decimal: %d - Binary: %b\n", EB*1024, EB*1024)
}
