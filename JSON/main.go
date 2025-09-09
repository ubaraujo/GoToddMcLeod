package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type personU struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	// MARSHALL

	p1 := person{
		First: "Tom",
		Last:  "Araújo",
		Age:   34,
	}

	p2 := person{
		First: "Luzanira",
		Last:  "Lima",
		Age:   67,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	byteSlice, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(byteSlice))

	// UNMARSHALL
	s := `[{"First":"Tom","Last":"Araújo","Age":34},{"First":"Luzanira","Last":"Lima","Age":67}]`
	byteSlice = []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", byteSlice)
	var peopleU []personU

	err = json.Unmarshal(byteSlice, &peopleU)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("all of the data", peopleU)

	for i, v := range peopleU {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
