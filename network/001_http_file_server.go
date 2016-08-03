package main

import "net/http"

func main() {
	http.ListenAndServe(":8085", http.FileServer(http.Dir(".")))
}
