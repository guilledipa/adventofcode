package main

import (
	"fmt"

	"github.com/guilledipa/adventofcode/2024/day/1/locationids"
)

func main() {
	locs := new(locationids.LocationIDs)
	if err := locs.populateVectors("input/location_ids.txt"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(int(locs.computeTotalDistance()))
}
