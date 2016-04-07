package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}
	rAddr, err := net.ResolveTCPAddr("tcp4", os.Args[1])
	checkError(err)

	lAddr, err := net.ResolveTCPAddr("tcp4", "localhost:9876")
	checkError(err)

	conn, err := net.DialTCP("tcp", lAddr, rAddr)

	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}
