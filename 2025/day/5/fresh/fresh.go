package fresh

import (
	"adventofcode/utils"
	"errors"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

var (
	ranges []Range
)

func Checker(file string) (int, int, error) {
	ranges, err := freshRanges(file) // Just to demonstrate reading the file, can be removed later.
	if err != nil {
		return 0, 0, err
	}
	ingredients, err := ingredients(file) // Just to demonstrate reading the file, can be removed later.
	if err != nil {
		return 0, 0, err
	}
	freshIngredients, err := getFreshIngredients(ingredients, ranges)
	if err != nil {
		return 0, 0, err
	}
	freshIDs, err := countIngredientIDs(ranges)
	if err != nil {
		return 0, 0, err
	}
	log.Println("Ranges:", ranges)
	log.Println("Ingredients:", ingredients)
	log.Println("Fresh Ingredients:", freshIngredients)
	log.Println("Fresh IDs:", freshIDs)
	return freshIngredients, freshIDs, nil
}

func countIngredientIDs(freshRanges []string) (int, error) {
	for _, r := range freshRanges { // 3-5, 2-4, 6-8
		indexes, err := parseRange(r)
		if err != nil {
			log.Printf("Error parsing range %q: %v", r, err)
			return 0, err
		}
		ranges = append(ranges, Range{Start: indexes[0], End: indexes[1]})
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})
	currentRange := ranges[0]
	mergedRanges := []Range{}
	for _, r := range ranges[1:] {
		log.Printf("Current Range: %+v, Next Range: %+v", currentRange, r)
		if r.Start <= currentRange.End+1 { // Overlapping or contiguous ranges
			if r.End > currentRange.End {
				currentRange.End = r.End // Extend the current range
			}
		} else {
			mergedRanges = append(mergedRanges, currentRange) // Save the merged current range
			currentRange = r
		}
	}
	mergedRanges = append(mergedRanges, currentRange) // Add the last range
	log.Println("Merged Ranges:", mergedRanges)
	var totalIDs int
	for _, r := range mergedRanges {
		totalIDs += r.End - r.Start + 1
	}
	return totalIDs, nil
}

// Part 1: Count how many ingredients are fresh based on the provided ranges.
func getFreshIngredients(ingredients []string, ranges []string) (int, error) {
	var freshIngredients int
	for _, ingr := range ingredients {
		i, err := strconv.Atoi(ingr)
		if err != nil {
			log.Printf("Error converting ingredient %q to integer: %v", ingr, err)
			return 0, err
		}
		for _, r := range ranges { // 3-5, 2-4, 6-8
			indexes, err := parseRange(r)
			if err != nil {
				log.Printf("Error parsing range %q: %v", r, err)
				return 0, err
			}
			if i >= indexes[0] && i <= indexes[1] {
				log.Printf("Ingredient %d is in range %q", i, r)
				freshIngredients++
				break // No need to check other ranges for this ingredient cos we know it's fresh
			}
		}
	}
	return freshIngredients, nil
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

func parseRange(s string) ([]int, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return nil, errors.New("invalid range format")
	}
	indexes := utils.StrToInt(parts)
	return indexes, nil
}
