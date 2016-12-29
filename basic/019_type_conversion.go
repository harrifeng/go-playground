package main

import (
	"fmt"
	"os"
)

func main() {

	s := "ABCDE"
	arr := []byte(s)
	fmt.Println(s)
	fmt.Printf("%v\n", arr)
	fmt.Println(string(arr))
	os.Exit(0)
}

// <===================OUTPUT===================>
// ABCDE
// [65 66 67 68 69]
// ABCDE
