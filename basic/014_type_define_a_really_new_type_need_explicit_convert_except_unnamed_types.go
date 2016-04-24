package main

import (
	"fmt"
	"os"
)

func main() {
	type bigint int64
	x := 1234
	var b bigint = bigint(x)
	var b2 int64 = int64(b)
	fmt.Println(b, b2)

	type myslice []int
	var s myslice = []int{1, 2, 3}
	var s2 []int = s
	fmt.Println(s, s2)
	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 1234 1234					  //
// [1 2 3] [1 2 3]				  //
////////////////////////////////////////////////////
