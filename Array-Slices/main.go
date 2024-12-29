package main

import (
	"fmt"
	"sort"
)

func main() {

	// array literal
	a := [3]int{42, 43, 44}
	fmt.Println(a)

	// iterate over an array with range
	for i, v := range a {
		fmt.Printf("index %v - value %v\n", i, v)
	}

	// let compiler count elements
	b := [...]string{"Hello", "Gophers"}
	fmt.Println(b)

	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", a)

	fmt.Println(len(a))
	fmt.Println(len(c))

	// slices
	xs := []string{"hello", "Gophers"}
	fmt.Printf("xs: %v \t len: %v \t cap: %v\n", xs, len(xs), cap(xs))

	// append slices
	xs = append(xs, "how", "are", "you")
	fmt.Printf("xs: %v \t len: %v \t cap: %v\n", xs, len(xs), cap(xs))

	// slicing slices
	fmt.Println(xs[0:3]) // not inclusive
	fmt.Println(xs[:4])
	fmt.Println(xs[:])
	fmt.Println(xs[0:])

	fmt.Println("----------------------\ntesting slice copying")
	var copia []string
	copia = xs // this creates a copy of the slice, "copia" is a new slice in memory
	fmt.Println("copia: ", copia)
	xs = append(xs, "testing", "reference")
	fmt.Println("copia after xs is altered", copia) // testing if copia is changed
	fmt.Println("xs after append:", xs)

	// delete from a slice
	fmt.Println(xs)
	xs = append(xs[:3], xs[4:]...) // removing the item "are"
	fmt.Println(xs)
	fmt.Println("-------------------------")

	// creating a slice with make()
	intSlice := make([]int, 0, 10) //creates a slice with zero items, and a capacity of 10.
	fmt.Println(intSlice)
	fmt.Printf("intSlice is of len: %v and capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 1, 2, 3)
	fmt.Println(intSlice)
	fmt.Printf("intSlice is of len: %v and capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(intSlice)
	fmt.Printf("intSlice is of len: %v and capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 11)
	fmt.Println(intSlice)
	fmt.Printf("intSlice is of len: %v and capacity %v\n", len(intSlice), cap(intSlice))

	// multidimensional slice
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)
	xxs := [][]string{jb, jm}
	fmt.Println(xxs)

	// slice internals & underlying arrays
	a1 := []int{0, 1, 2, 3, 4, 5}
	b1 := a1
	fmt.Println("a ", a1)
	fmt.Println("a ", b1)
	fmt.Println("----------------")
	a1[0] = 7
	fmt.Println("a ", a1)
	fmt.Println("b ", b1)
	fmt.Println("----------------")

	strA := []string{"a", "b", "c", "d"}
	strB := strA
	fmt.Println("a ", strA)
	fmt.Println("b ", strB)
	fmt.Println("----------------")
	// to solve this when trying to copy an array without the colaterals, use function copy()
	strC := make([]string, 4)
	copy(strC, strA)
	strA[0] = "X"
	fmt.Println("a ", strA)
	fmt.Println("b ", strB)
	fmt.Println("b ", strC)
	fmt.Println("----------------")

	// yet another example of underlying arrays in slices
	num := []float64{3, 1, 4, 2}
	fmt.Println(medianOne(num))
	fmt.Println("after medianOne: ", num)

	num2 := []float64{3, 1, 4, 2}
	fmt.Println(medianTwo(num2))
	fmt.Println("after medianTwo: ", num2)
}

func medianOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]
	}
	return (x[i-1] + x[i]) / 2
}

func medianTwo(x []float64) float64 {
	// allocate new underlying array
	n := make([]float64, len(x))
	copy(n, x)

	sort.Float64s(n)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]
	}
	return (n[i-1] + n[i]) / 2
}
