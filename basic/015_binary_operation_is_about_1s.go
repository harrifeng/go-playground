package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.ParseInt("0110", 2, 64)
	b, _ := strconv.ParseInt("1011", 2, 64)
	fmt.Printf("%b\n", a&b)  // => 0010: both need to be 1
	fmt.Printf("%b\n", a|b)  // => 1111: either need to be 1
	fmt.Printf("%b\n", a^b)  // => 1101: either need to be 1
	fmt.Printf("%b\n", a&^b) // => 0100: clear all the a's 1 if b's correspoing pos is 1
	fmt.Println()
	os.Exit(0)
}

////////////////////////////////////////////////////
// <===================OUTPUT===================> //
// 10						  //
// 1111						  //
// 1101						  //
// 100						  //
////////////////////////////////////////////////////
