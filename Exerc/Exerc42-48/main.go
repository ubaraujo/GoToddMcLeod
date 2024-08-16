package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4}
	fmt.Printf("Type: %T\n", arr)
	for i := range arr {
		fmt.Printf("index value: %v\n", arr[i])
	}
	// other way, same answer
	arr2 := [5]int{} // array
	for i := 0; i < len(arr2); i++ {
		arr2[i] = i
	}
	for i, v := range arr {
		fmt.Printf("Index: %v - Value: %v - Type: %T\n", i, v, v)
	}

	// exercise 43
	slc := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("Type: %T\n", slc)
	for i, v := range slc {
		fmt.Printf("Index: %v - Value: %v - Type: %T\n", i, v, v)
	}

	// exercise 44
	fmt.Println(slc[:5])
	fmt.Println(slc[5:])
	fmt.Println(slc[2:7])
	fmt.Println(slc[1:6])

	// exercise 45
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	// exercise 46
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	// exercise 47
	/*
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`
	*/
	states := make([]string, 0, 50) // creates the array, with no values inside, capacity to 50
	fmt.Printf("Len: %v - cap: %v - Type: %T\n", len(states), cap(states), states)
	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Printf("Len: %v - cap: %v - Type: %T\n", len(states), cap(states), states)
	fmt.Println(states)

	// exercise 48]
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenny", "I'm 008"}
	sStr := [2][3]string{}
	fmt.Printf("%v\n", len(sStr))
	//for i := 0; i < 2; i++ {
	for i := range sStr {
		if i == 0 {
			for j := 0; j < 3; j++ {
				sStr[i][j] = jb[j]
			}
		} else {
			for j := 0; j < 3; j++ {
				sStr[i][j] = jm[j]
			}
		}
	}
	for i, v := range sStr {
		fmt.Println(i, v)
	}
}
