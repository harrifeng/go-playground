package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func postToLocal(json_body []byte) {
	time.Sleep(200 * time.Millisecond)

	log.Println("done one-->")
}

func main() {

	start := time.Now()

	backFile := flag.String("back_file", "./backup_error.log", "Input the base64 encoding file which contains missing data")

	file, err := os.Open(*backFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		http_body, err := base64.StdEncoding.DecodeString(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		postToLocal(http_body)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("repair finished")
	fmt.Println()
	os.Exit(0)
}
