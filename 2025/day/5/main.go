package main

import (
	"adventofcode/2025/day/5/fresh"
)

func main() {
	part1, part2, err := fresh.Checker("2025/day/5/fresh/input.txt")
	if err != nil {
		panic(err)
	}
	println("Part 1:", part1)
	println("Part 2:", part2)
}
