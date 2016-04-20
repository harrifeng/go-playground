package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = "Hello, World.\n" + text
	fmt.Println(text)
	os.Exit(0)
}
