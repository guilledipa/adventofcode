package main

import "adventofcode/2025/day/6/cephalopod"

func main() {
	part1, part2, err := cephalopod.CalculateTotal("input/data.txt")
	if err != nil {
		panic(err)
	}
	println("Part 1:", part1)
	println("Part 2:", part2)
}
