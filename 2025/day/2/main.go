package main

import (
	"adventofcode/2025/day/2/validator"
	"fmt"
)

func main() {
	p1, p2, err := validator.Validate("input/data.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Result P1: %d, P2: %d\n", p1, p2)
}
