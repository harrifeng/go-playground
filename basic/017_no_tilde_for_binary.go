package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	x, _ := strconv.ParseInt("1001", 2, 64)
	// 9 => 1001
	// 1001 => 01001
	// 01001 => 10110
	// 10110 - 1 => 10101
	// 10101 => 11010
	fmt.Println(x, ^x)
	os.Exit(0)
}
