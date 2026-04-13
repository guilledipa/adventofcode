package main

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"os"
)

const (
	// inputData = "input/data.txt"
	inputData = "input/test_data.txt"
	dialInit  = 50
)

func main() {
	fmt.Println("hello world!")
	s, f, err := utils.CreateScanner(inputData)
	if err != nil {
		log.Fatalf("Unable to read input data: %q", err)
	}
	defer f.Close()
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
