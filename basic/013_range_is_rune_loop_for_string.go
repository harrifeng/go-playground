package main

import (
	"fmt"
	"os"
)

func main() {
	s := "abc汉字"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c, ", s[i])
	}

	fmt.Println()

	for _, r := range s {
		fmt.Printf("%c, ", r)
	}

	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// a, b, c, æ, ±, , å, ­, , 		  //
// a, b, c, 汉, 字, 				  //
////////////////////////////////////////////////////
