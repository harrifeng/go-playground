package main

import (
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func main() {

	// composite literal can init the array or map
	s := []string{"abc", "def", "ghi"}
	m := map[int]string{1: "no error", 2: "Eio", 3: "invalid argument"}
	fmt.Println(s)
	fmt.Println(m)
	// composite literal can also be used for struct, way one, as slice
	wone := Point{1, 2}
	fmt.Println(wone)
	// composite literal can also be used for struct, way two, as slice
	wtwo := Point{y: 1, x: 2}
	fmt.Println(wtwo)
	os.Exit(0)
}

// <===================OUTPUT===================>
// [abc def ghi]
// map[1:no error 2:Eio 3:invalid argument]
// {1 2}
// {2 1}
