package main

import (
	"adventofcode/2025/day/4/forklift"
	"fmt"
	"log"
)

func main() {
	rolls1, rolls, err := forklift.Forklift("input/data.txt")
	if err != nil {
		log.Fatalf("Error: %v\n", err)

	}
	fmt.Printf("Result: %d, %d\n", rolls1, rolls)
}
