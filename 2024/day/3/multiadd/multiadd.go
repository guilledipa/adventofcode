// Package multiadd provides useful functions to add the multiplication of numbers.
package multiadd

import (
	"adventofcode/2024/day/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// ProcessData will parse raw data and populate the Report struct.
func ProcessData(file string) ([]byte, error) {
	data := []byte{}
	scanner, f, err := utils.CreateScanner(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	for scanner.Scan() {
		data = append(data, scanner.Bytes()...)
	}
	return data, scanner.Err()
}

// AccumulateAllMul will accumulate all the valid multiplication operations in the data.
func AccumulateAllMul(data []byte) int {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`)
	matches := re.FindAll(data, -1)
	acum := 0
	dontCount := false
	for _, match := range matches {
		// Check if we should skip the next multiplication (D3P2). Also
		// had to update the RegEx to match the new patterns.
		if string(match) == "don't()" {
			dontCount = true
			continue
		} else if string(match) == "do()" {
			dontCount = false
			continue
		}
		if dontCount {
			continue
		}
		re := regexp.MustCompile(`(\d{1,3})`)
		digits := re.FindAllString(string(match), -1)
		a, b, err := ConvertToInts(digits)
		if err != nil {
			log.Printf("error converting digits %s: %v", digits, err)
			continue
		}
		acum += a * b
	}
	return acum
}

// ConvertToInts will convert a slice of strings to a slice of ints.
func ConvertToInts(digits []string) (int, int, error) {
	if len(digits) != 2 {
		return 0, 0, fmt.Errorf("invalid number of digits")
	}
	a, err := strconv.Atoi(digits[0])
	if err != nil {
		return 0, 0, err
	}
	b, err := strconv.Atoi(digits[1])
	if err != nil {
		return 0, 0, err
	}
	return a, b, nil
}
