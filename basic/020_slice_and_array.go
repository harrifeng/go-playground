package main

import (
	"fmt"
	"os"
)

func main() {
	sli := []int{1, 2, 3}
	fmt.Printf("%T, %v\n", sli, sli)
	arr := [...]int{1, 2, 3}
	fmt.Printf("%T, %v\n", arr, arr)
	os.Exit(0)
}

// <===================OUTPUT===================>
// []int, [1 2 3]
// [3]int, [1 2 3]
