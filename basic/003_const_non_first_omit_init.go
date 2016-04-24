package main

import (
	"fmt"
	"os"
)

func main() {
	const (
		s = "abc"
		x // x = "abc"
	)
	fmt.Println(s, x)
	os.Exit(0)
}
