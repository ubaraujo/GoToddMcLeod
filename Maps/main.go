package main

import "fmt"

func main() {
	myMap := map[string]int{
		"Todd":    42,
		"Tomzera": 33,
		"McLeudd": 40,
	}

	fmt.Println("The age of Todd is", myMap["Todd"])
	fmt.Println(myMap)
	fmt.Printf("%#v\n", myMap)

	// another way to create a map
	otherMap := make(map[string]int)
	otherMap["Tom"] = 33
	otherMap["Steph"] = 38

	fmt.Println(otherMap)
	fmt.Printf("%#v\n", otherMap)

	// iterate over a MAP with for and range
	for k, v := range otherMap {
		fmt.Println(k, v)
	}

	for _, v := range otherMap {
		fmt.Println("Values of otherMap:", v)
	}

	for k := range myMap {
		fmt.Println(k)
	}

	// delete an object from a MAP
	delete(otherMap, "Tom")
	fmt.Println("After deleting Tom")
	fmt.Println(otherMap)
	for k, v := range otherMap {
		fmt.Println(k, v)
	}

	fmt.Println("--- accessing keys that don't exist ---")
	delete(otherMap, "test")      // won't panic
	fmt.Println(otherMap["John"]) // won't panic - it returns 0
	fmt.Println("---")

	// comma ok idiom with map
	v, ok := myMap["Tomzera"]
	if ok {
		fmt.Println("The value prints:", v)
	} else {
		fmt.Println("Key didn't exist")
	}

	// another way 'statement statement idiom'
	if v, ok := myMap["Test"]; !ok { // we teste first if not ok and treat the error
		fmt.Println("Key didn't exist!")
	} else {
		fmt.Println("The value prints:", v)
	}
}
