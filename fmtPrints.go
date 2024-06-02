package main

import "fmt"

func main() {

	x := "Oi bom dia\ncomo vai?\tespero \"que\" tudo bem."

	// raw string literals
	y := `
	desta forma o print 
	é executado exatamente como formatado aqui
	dentro do código...
			... interessante, não?`
	fmt.Println(x)
	fmt.Println(y)

	// Sprint -> não printa na tela. Gera e retorna uma nova string sem espaços entre os argumentos
	// Sprintln -> mesma coisa da Sprint, mas com espaços entre os argumentos
	str1 := "oi"
	str2 := "bom dia"
	str3 := fmt.Sprint(str1, str2)
	str4 := fmt.Sprintln(str1, str2)
	fmt.Println(str3)
	fmt.Println(str4)
}
