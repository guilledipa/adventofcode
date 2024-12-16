// package locationids provides functions to process location ids.
package locationids

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// LocationIDs hold historically significant locations.
type LocationIDs struct {
	Left  []int
	Right []int
}

// PopulateVectors parses a file and populates the righ and left fields in the
// LocationIDs struct
func (l *LocationIDs) PopulateVectors(file string) error {
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
		l.Left = append(l.Left, leftInt)
		rightInt, err := strconv.Atoi(IDs[1])
		if err != nil {
			return err
		}
		l.Right = append(l.Right, rightInt)
	}
	l.sortVectors()
	return scanner.Err()
}

func (l *LocationIDs) sortVectors() {
	sort.Ints(l.Left)
	sort.Ints(l.Right)
}

func (l *LocationIDs) ComputeTotalDistance() float64 {
	var totalDistance float64
	for i := 0; i < len(l.Left); i++ {
		totalDistance += math.Abs(float64(l.Right[i] - l.Left[i]))
	}
	return totalDistance
}
