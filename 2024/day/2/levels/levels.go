// Package levels provide useful functions to parse levels data.
package levels

import (
	"adventofcode/2024/day/utils"
	"strconv"
	"strings"
)

// Report is a struct to hold level data and its safety.
type Report struct {
	Levels []int
	Safe   bool
}

// Reports is a representation of the raw input data.
type Reports []Report

// PopulateReports will parse raw data and populate the Report struct.
func (r *Reports) PopulateReports(file string) error {
	scanner, f, err := utils.CreateScanner(file)
	defer f.Close()
	if err != nil {
		return err
	}
	for scanner.Scan() {
		rep := new(Report)
		line := scanner.Text()
		levels := strings.Fields(line)
		for _, level := range levels {
			levelInt, err := strconv.Atoi(level)
			if err != nil {
				return err
			}
			rep.Levels = append(rep.Levels, levelInt)
		}
		rep.Safe = isSafe(rep.Levels)
		*r = append(*r, *rep)
	}
	return scanner.Err()
}

// CountSafe returns the number of safe reports.
func (r *Reports) CountSafe() int {
	var count int
	for _, rep := range *r {
		if rep.Safe {
			count++
		}
	}
	return count
}

// isSafe checks if the levels are safe based on the rules.
func isSafe(levels []int) bool {
	if checkSafety(levels) {
		return true
	}
	for i := range levels {
		modifiedLevels := make([]int, len(levels)-1)
		copy(modifiedLevels, levels[:i])
		copy(modifiedLevels[i:], levels[i+1:])
		if checkSafety(modifiedLevels) {
			return true
		}
	}
	return false
}

// checkSafety checks if the levels are safe without removing any level.
func checkSafety(levels []int) bool {
	increasing := levels[1] > levels[0] // Check initial direction
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		if increasing {
			if diff < 1 || diff > 3 {
				return false // Not increasing within range
			}
		} else {
			if diff > -1 || diff < -3 {
				return false // Not decreasing within range
			}
		}
	}
	return true
}
