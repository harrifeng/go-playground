package main

import (
	"fmt"
	"os"
)

func main() {
	b := new(bool)
	fmt.Printf("%T ", b)
	fmt.Println(*b)

	i := new(int)
	fmt.Printf("%T ", i)
	fmt.Println(*i)

	s := new(string)
	fmt.Printf("%T ", s)
	fmt.Println(*s)

	type Color int
	c := new(Color)
	fmt.Printf("%T ", c)
	fmt.Println(*c)

	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// *bool false					  //
// *int 0					  //
// *string 					  //
// *main.Color 0				  //
////////////////////////////////////////////////////
