package main

import (
	"adventofcode/2024/day/3/multiadd"
	"log"
)

func main() {
	data, err := multiadd.ProcessData("input/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	accu := multiadd.AccumulateAllMul(data)
	log.Printf("Total multiplication: %d", accu)
}
