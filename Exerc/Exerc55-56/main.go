package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	maker string
	model string
	doors int
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{
			electric: false,
		},
		maker: "Fiat",
		model: "Uno",
		doors: 2,
		color: "white",
	}

	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		maker: "Gurgel",
		model: "Itaipu E150",
		doors: 2,
		color: "Red",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Printf("V1: %v %v %v\n", v1.maker, v1.model, v1.electric)
	fmt.Printf("V2: %v %v %v\n", v2.maker, v2.model, v2.electric)

	// EXERC 56 ANONYMOUS STRUCTS
	person := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Tom",
		friends: map[string]int{
			"Fulano":  35,
			"Ciclano": 34},
		favDrinks: []string{
			"Caipirinha",
			"Cuba Libre"},
	}

	fmt.Println(person)
}
