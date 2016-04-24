package main

import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(b)
	)
	fmt.Println(a, b, c)
	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// abc 3 8					  //
////////////////////////////////////////////////////
