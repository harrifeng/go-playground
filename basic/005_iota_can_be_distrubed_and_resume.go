package main

import (
	"fmt"
	"os"
)

func main() {
	const (
		A = iota
		B
		C = "c"
		D
		E = iota
		F
	)
	fmt.Println(A, B, C, D, E, F)
	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 0 1 c c 4 5					  //
////////////////////////////////////////////////////
