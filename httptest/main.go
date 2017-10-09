package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintln(writer, "hello")
	for i := 1; i < 10000000000; i++ {
		fmt.Println("hello")
	}

	return
}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8888", nil)
	os.Exit(0)
}
