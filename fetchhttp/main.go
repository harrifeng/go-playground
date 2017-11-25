package main

import (
	"fmt"
	"flag"
	"time"
	"log"
	"os"
	"bufio"
	"bytes"
	"io/ioutil"
)


func getShanghaiTimeName() string {
	t := time.Now()
	c, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println("err", err.Error())
	}

	return t.In(c).Format("2006-01-02-15-04-05") + ".log"
}

func main() {
	logfile := flag.String("log", "http2.log", "Input log text file to analysis")
	flag.Parse()

	file, err := os.Open(*logfile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ret bytes.Buffer
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) > 8 && t[2:8] == "log_id" {
			ret.WriteString(t)
			ret.WriteString("\n")
		}

	}

	ioutil.WriteFile(getShanghaiTimeName(), ret.Bytes(), 0644)
	fmt.Println("Finish")
}
