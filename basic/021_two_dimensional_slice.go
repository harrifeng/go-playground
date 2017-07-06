package main

import (
	"fmt"
	"os"
)

func createIntSlice(length int) [][]int {
	sli := make([][]int, length)
	return sli
}

func main() {
	a := [][]int{}
	for i := 0; i < 3; i++ {
		a = append(a, []int{})
	}
	fmt.Println(a)
	b := make([][]int, 3)
	fmt.Println(b)

	c := [][]int{
		[]int{},
		[]int{},
		[]int{},
	}
	fmt.Println(c)
	// non-const way is to use the make function
	fmt.Println(createIntSlice(3))
	os.Exit(0)
}

// <===================OUTPUT===================>
// [[] [] []]
// [[] [] []]
// [[] [] []]
// [[] [] []]
