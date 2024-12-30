package main

import "fmt"

func main() {
	map1 := make(map[string][]string)

	map1["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	map1["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	map1["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	fmt.Printf("%#v\n", map1)

	map1["araujo_tom"] = []string{"books", "games", "movies"}

	for i, v := range map1 {
		fmt.Printf("%v:\n", i)
		for i, v := range v {
			fmt.Printf("idx: %v - %v\n", i, v)
		}
	}

	fmt.Println(" --- Deleting Tom --- ")
	delete(map1, "araujo_tom")
	for i, v := range map1 {
		fmt.Printf("%v:\n", i)
		for i, v := range v {
			fmt.Printf("idx: %v - %v\n", i, v)
		}
	}

}
