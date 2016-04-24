package main

import (
	"fmt"
	"os"
)

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color) {
	fmt.Println(c)
}

func main() {

	test(Black)
	test(Red)
	test(Blue)

	test(2)
	// cannot use x (type int) as type Color in argument to test
	// x := 1
	// test(x)
	os.Exit(0)
}
