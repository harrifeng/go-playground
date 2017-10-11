package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/korovkin/limiter"
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

	limit := limiter.NewConcurrencyLimiter(10)
	for scanner.Scan() {
		http_body, err := base64.StdEncoding.DecodeString(scanner.Text())
		_ = http_body
		if err != nil {
			log.Fatal(err)
		}

		var postToLocal = func() {
			time.Sleep(200 * time.Millisecond)

			fmt.Println(len(http_body))
			log.Println("done one-->")
		}

		limit.Execute(postToLocal)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	limit.Wait()
	fmt.Println("<--Time used-->:", time.Since(start))
	fmt.Println("repair finished")
	fmt.Println()
	os.Exit(0)
}
