package main

import (
	"net/http"
	"os"
)

func main() {
	shareFolder := "/share"
	if len(os.Args) > 1 {
		shareFolder = os.Args[1]
	}
	http.ListenAndServe(":8081", http.FileServer(http.Dir(shareFolder)))
}
