package main

import (
	"adventofcode/2024/day/4/search"
	"log"
)

func main() {
	d, err := search.NewData("input/test_data.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total findings: %d", d.CountXMAS())
}
