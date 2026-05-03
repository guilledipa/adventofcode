package main

import (
	"adventofcode/2025/day/8/junctionbox"
	"fmt"
	"log"
)

func main() {
	Part1, Part2, err := junctionbox.CalculateTotal("./input/data.txt")
	if err != nil {
		log.Fatalf("Error calculating totals: %v", err)
	}
	fmt.Println("Part 1:", Part1)
	fmt.Println("Part 2:", Part2)
}
