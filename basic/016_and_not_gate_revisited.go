package main

import (
	"fmt"
	"os"
)

func main() {
	a := 0
	a |= 1 << 2 // set bit 2 flag
	fmt.Printf("%b\n", a)
	a |= 1 << 6 // set bit 6 flag
	fmt.Printf("%b\n", a)
	a = a &^ (1 << 6) // clear bit 6 flag
	fmt.Printf("%b\n", a)
	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 100						  //
// 1000100					  //
// 100						  //
////////////////////////////////////////////////////
