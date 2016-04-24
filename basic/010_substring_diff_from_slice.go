package main

import (
	"fmt"
	"os"
)

func main() {
	s := "hello, world"
	s1 := s[:5]
	fmt.Printf("%T\n", s1)

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", a)
	a1 := a[:3]
	fmt.Printf("%T\n", a1)

	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// string					  //
// [5]int					  //
// []int					  //
////////////////////////////////////////////////////
