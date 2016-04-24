package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%T\n", 'a')
	fmt.Printf("%T\n", "a")

	var c1, c2 rune = '\u6211', '们'
	fmt.Println(c1 == '我', string(c2) == "\xe4\xbb\xac")
	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// int32					  //
// string					  //
// true true					  //
////////////////////////////////////////////////////
