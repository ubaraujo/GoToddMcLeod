package main

import "fmt"

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {
	p1 := person{
		first:    "Tom",
		last:     "Araujo",
		icecream: []string{"Maracujá", "Chocolate"},
	}

	p2 := person{
		first:    "Luzanira",
		last:     "Lima",
		icecream: []string{"Limão", "Morango"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.icecream)
	fmt.Println(p2.icecream)

	for _, v := range p1.icecream {
		fmt.Println(p1.first, "favorit is", v)
	}

	for _, v := range p2.icecream {
		fmt.Println(p2.first, "favorit is", v)
	}

	mapPerson := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range mapPerson {
		fmt.Printf("%T - %v\n", v, v.first)

		for _, v2 := range p1.icecream {
			fmt.Println(v2)
		}
	}
}
