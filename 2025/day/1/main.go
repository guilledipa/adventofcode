package main

import (
	"adventofcode/2025/day/1/dial"
	"fmt"
)

const (
	//inputData = "input/test_data2.txt"
	//inputData = "input/test_data.txt"
	inputData = "input/data.txt"
)

func main() {
	zeros := dial.UseDial(inputData, true)
	fmt.Printf("Number of zeros encountered: %d\n", zeros)
}
