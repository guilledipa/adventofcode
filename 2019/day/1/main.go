package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var massFilePath = flag.String("mass_file_path", "input/modules_mass.txt", "Modules mass data file.")

func fuelRequirement(mass int) int {
	var fuel int
	if mass <= 0 {
		return fuel
	}
	fuel = (mass / 3) - 2
	return fuel
}

func main() {
	file, err := os.Open(*massFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var totalFuel int
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("unable to convert to integer: %v", err)
			continue
		}
		totalFuel = totalFuel + fuelRequirement(mass)
	}
	fmt.Printf("Total fuel requirements: %v", totalFuel)
}
