package fresh

import (
	"adventofcode/utils"
	"errors"
	"log"
	"strings"
)

func Checker(file string) (int, int, error) {
	// TODO: Implement the logic to read the file and compute the results for part 1 and part 2.
	ranges, err := freshRanges(file) // Just to demonstrate reading the file, can be removed later.
	if err != nil {
		return 0, 0, err
	}
	ing, err := ingredients(file) // Just to demonstrate reading the file, can be removed later.
	if err != nil {
		return 0, 0, err
	}
	log.Println("Ranges:", ranges)
	log.Println("Ingredients:", ing)
	return 3, 0, nil
}

func freshRanges(file string) ([]string, error) {
	var lines []string
	if err := utils.ProcessLines(file, func(line string) error {
		if strings.TrimSpace(line) == "" {
			return errors.New("blank line found") // Return error to stop processing
		}
		lines = append(lines, line)
		return nil
	}); err != nil && err.Error() != "blank line found" {
		// Only return the error if it's not our "stop processing" signal
		return nil, err
	}
	return lines, nil
}

func ingredients(file string) ([]string, error) {
	var lines []string
	foundBlank := false
	if err := utils.ProcessLines(file, func(line string) error {
		if !foundBlank {
			if strings.TrimSpace(line) == "" {
				foundBlank = true
			}
			return nil // Skip lines until we find the blank line
		}
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return lines, nil
}
