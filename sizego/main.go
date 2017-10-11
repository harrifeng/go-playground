package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func main() {

	start := time.Now()

	backFile := flag.String("back_file", "./backup_error.log", "Input the base64 encoding file which contains missing data")

	file, err := os.Open(*backFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup
	var tokens = make(chan struct{}, 10)

	for scanner.Scan() {
		http_body, err := base64.StdEncoding.DecodeString(scanner.Text())
		_ = http_body
		if err != nil {
			log.Fatal(err)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			tokens <- struct{}{}
			time.Sleep(200 * time.Millisecond)

			fmt.Println(len(http_body))
			log.Println("done one-->")
			<-tokens
		}()

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	fmt.Println("<--Time used-->:", time.Since(start))
	fmt.Println("repair finished")
	fmt.Println()
	os.Exit(0)
}
