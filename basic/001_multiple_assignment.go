package main

import (
	"fmt"
	"os"
)

func main() {
	data, i := [3]int{0, 1, 2}, 0
	fmt.Println(data, i)
	i, data[i] = 2, 100 // ready the left, then assign from right
	fmt.Println(data, i)
	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// [0 1 2] 0					  //
// [100 1 2] 2					  //
////////////////////////////////////////////////////
