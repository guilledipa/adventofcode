package main

import (
	"adventofcode/2025/day/3/joltage"
	"fmt"
)

func main() {
	maxJoltage, err := joltage.MaxJoltage("input/data.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Result: %d\n", maxJoltage)
	maxJoltage2, err := joltage.MaxJoltage2("input/data.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Result2: %d\n", maxJoltage2)
}
