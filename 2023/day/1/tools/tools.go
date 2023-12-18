package tools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// readInputFile reads the contents of a given file provided by it's path and
// returns a list containing every line in the file or an error if the program
// was unable to read the file.
func readInputFile(filePath string) ([]string, error) {
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
	for _, c := range calibrationValue {
		if unicode.IsDigit(c) {
			fmt.Println("TEST: ", c)
			digits = append(digits, strconv.Atoi(c)) //FIX
		}
	}
	return digits
}

// CheckCalibrationValues verifies the integrity of the calibration values
// provided in an input file.
func CheckCalibrationValues(filePath string) (int, error) {
	var checkSum int
	lines, err := readInputFile(filePath)
	if err != nil {
		return checkSum, err
	}
	for _, cv := range lines {
		digits := findDigits(cv)
		fmt.Println(digits)
		if len(digits) < 2 {
			continue
		}
		fmt.Printf("%d-%d\n", digits[0], digits[len(digits)-1])
	}
	return checkSum, nil
}
