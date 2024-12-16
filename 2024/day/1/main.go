package main

import (
	"adventofcode/2024/day/1/locationids"
	"fmt"
	"log"
)

func main() {
	locs := new(locationids.LocationIDs)
	if err := locs.PopulateVectors("input/location_ids.txt"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(int(locs.ComputeTotalDistance()))
}
