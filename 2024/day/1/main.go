package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// LocationIDs hold historically significant locations.
type LocationIDs struct {
	left  []int
	right []int
}

func (l *LocationIDs) populateVectors(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		IDs := strings.Fields(line)
		leftInt, err := strconv.Atoi(IDs[0])
		if err != nil {
			return err
		}
		l.left = append(l.left, leftInt)
		rightInt, err := strconv.Atoi(IDs[1])
		if err != nil {
			return err
		}
		l.right = append(l.right, rightInt)
	}
	l.sortVectors()
	return scanner.Err()
}

func (l *LocationIDs) sortVectors() {
	sort.Ints(l.left)
	sort.Ints(l.right)
}

func (l *LocationIDs) computeTotalDistance() {
	var totalDistance float64
	for i := 0; i < len(l.left)-1; i++ {
		totalDistance += math.Abs(float64(l.right[i] - l.left[i]))
	}
	fmt.Println(totalDistance)
}

func main() {
	locs := new(LocationIDs)
	if err := locs.populateVectors("input/location_ids.txt"); err != nil {
		fmt.Println(err)
	}
}
