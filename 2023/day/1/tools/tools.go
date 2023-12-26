package tools

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var calibrationValueMappings = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func calibrationValueMapKeys() []string {
	keys := make([]string, len(calibrationValueMappings))
	i := 0
	for k := range calibrationValueMappings {
		keys[i] = k
		i++
	}
	return keys
}

// ReadInputFile reads the contents of a given file provided by it's path and
// returns a list containing every line in the file or an error if the program
// was unable to read the file.
func ReadInputFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	// Lines are less than 64K
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func findDigits(calibrationValue string) []int {
	var digits []int
	for i, c := range calibrationValue {
		if unicode.IsDigit(c) {
			d, err := strconv.Atoi(string(c))
			if err != nil {
				log.Printf("Unable to convert '%s' to int: %v", string(c), err)
				continue
			}
			digits = append(digits, d)
		}
		if unicode.IsLetter(c) { // Unnecessary actually.
			for _, number := range calibrationValueMapKeys() {
				if strings.HasPrefix(calibrationValue[i:], number) {
					digits = append(digits, calibrationValueMappings[number])
				} else {
					continue
				}
			}
		}
	}
	return digits
}

// CheckCalibrationValues verifies the integrity of the calibration values
// provided in an input file.
func CheckCalibrationValues(filePath string) (int, error) {
	var checkSum int
	var value int
	lines, err := ReadInputFile(filePath)
	if err != nil {
		return checkSum, err
	}
	for _, cv := range lines {
		digits := findDigits(cv)
		switch cvs := len(digits); cvs {
		case 0:
			continue
		case 1: // 1 then we create a double-digit using the same CV.
			value, err = strconv.Atoi(fmt.Sprintf("%d%d", digits[0], digits[0]))
			if err != nil {
				return checkSum, err
			}
		default: // 2 or more CVs we create a double digit using the first and last CVs.
			value, err = strconv.Atoi(fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1]))
			if err != nil {
				return checkSum, err
			}
		}
		checkSum += value
	}
	return checkSum, nil
}
