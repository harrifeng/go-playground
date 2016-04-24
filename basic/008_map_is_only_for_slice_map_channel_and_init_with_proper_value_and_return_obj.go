package main

import "fmt"

// receive-only channel is often used in function
func printChannel(rc <-chan string) {
	fmt.Println(<-rc)
}

func main() {

	s := make([]int, 10, 100) // can have a third cap para
	fmt.Printf("%T \tlen: %d cap: %d\n", s, len(s), cap(s))
	fmt.Println("-----------------------")

	m := make(map[int]int, 2)
	// Map: An initial allocation is made according to the size but the
	// resulting map has length 0. The size may be omitted, in which case
	// a small starting size is allocated
	fmt.Printf("%T \tlen: %d\n", m, len(m))

	fmt.Println("-----------------------")

	c := make(chan string, 1)
	fmt.Printf("%T \tlen: %d\n", c, len(c))
	c <- "abc"
	fmt.Printf("%T \tlen: %d\n", c, len(c))
	<-c
	fmt.Printf("%T \tlen: %d\n", c, len(c))

	fmt.Println("-----------------------")

	c2 := make(chan<- string, 1)
	fmt.Printf("%T \tlen: %d\n", c, len(c2))
	c2 <- "def"
	fmt.Printf("%T \tlen: %d\n", c, len(c2))

	fmt.Println("-----------------------")
	c3 := make(chan string, 1)
	fmt.Printf("%T \tlen: %d\n", c, len(c3))
	c3 <- "from c3"
	fmt.Printf("%T \tlen: %d\n", c, len(c3))
	printChannel(c3)
}

///////////////////////////////////////////////////////////
// <===================OUTPUT===================>	 //
// []int 	len: 10 cap: 100			 //
// -----------------------				 //
// map[int]int 	len: 0					 //
// -----------------------				 //
// chan string 	len: 0					 //
// chan string 	len: 1					 //
// chan string 	len: 0					 //
// -----------------------				 //
// chan string 	len: 0					 //
// chan string 	len: 1					 //
// -----------------------				 //
// chan string 	len: 0					 //
// chan string 	len: 1					 //
// from c3						 //
///////////////////////////////////////////////////////////
