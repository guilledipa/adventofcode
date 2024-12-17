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

// BUG!!!!
func isSafe(levels []int) bool {
	for i := 0; i < len(levels)-1; i++ {
		if levels[i] > levels[i+1] { // Increasing
			if utils.AbsDistance(levels[i], levels[i+1]) > 3 {
				return false // Distance is greater than 3
			}
			continue
		} else if levels[i] < levels[i+1] { // Decreasing
			if utils.AbsDistance(levels[i], levels[i+1]) > 3 {
				return false // Distance is greater than 3
			}
			continue
		} else {
			return false // Equal numbers
		}
	}
	return true
}
