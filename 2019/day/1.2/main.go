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
	var totalFuel int
	if mass < 6 {
		return totalFuel
	}
	fuel = (mass / 3) - 2
	totalFuel = fuel + fuelRequirement(fuel)
	return totalFuel
}

func modulesMasses(scanner *bufio.Scanner) <-chan int {
	out := make(chan int)
	go func() {
		for scanner.Scan() {
			mass, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Printf("unable to convert to integer: %v", err)
				continue
			}
			out <- mass
		}
		close(out)
	}()
	return out
}

func fuel(modulesMasses <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for mass := range modulesMasses {
			out <- fuelRequirement(mass)
		}
		close(out)
	}()
	return out
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
	m := modulesMasses(scanner)
	out := fuel(m)

	var totalFuel int
	for f := range out {
		totalFuel = totalFuel + f
	}

	fmt.Printf("Total fuel requirements: %v", totalFuel)
}
