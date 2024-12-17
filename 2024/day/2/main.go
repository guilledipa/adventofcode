package main

import (
	"adventofcode/2024/day/2/levels"
	"fmt"
	"log"
)

func main() {
	reps := new(levels.Reports)
	if err := reps.PopulateReports("input/test_input.txt"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(reps)
}
