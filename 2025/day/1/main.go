package main

import (
	"adventofcode/2025/day/1/dial"
	"fmt"
)

const (
	inputData = "input/data.txt"
)

func main() {
	zeros := dial.UseDial(inputData)
	fmt.Printf("Number of zeros encountered: %d\n", zeros)
}
