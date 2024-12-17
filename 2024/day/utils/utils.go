package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// CreateScanner is a helper function to scan files line by line.
// Note: I need to return the os.File filedescriptor because I can't close
// the file here; I must do it after I use the scanner.
func CreateScanner(file string) (*bufio.Scanner, *os.File, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(f.Name())
	scanner := bufio.NewScanner(f)
	return scanner, f, scanner.Err()
}

// AbsDistance returns the absolute distance between two numbers.
func AbsDistance(a, b int) int {
	dist := math.Abs(float64(a - b))
	return int(dist)
}
