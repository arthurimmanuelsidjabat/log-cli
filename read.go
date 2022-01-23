package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample.log")

	if err != nil {
		log.Fatalf("Gagal membuka file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var loglines []string

	for scanner.Scan() {
		loglines = append(loglines, scanner.Text())
	}

	file.Close()

	for _, eachline := range loglines {
		fmt.Println(eachline)
	}
}
