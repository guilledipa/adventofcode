package tachyonmanifold

import (
	"adventofcode/utils"
)

type VisitedState struct {
	row, col int
}

func CalculateTotal(file string) (int, int, error) {
	var p1, p2 int
	var startCol int
	manifold, err := utils.ReadMatrix(file, "")
	if err != nil {
		return 0, 0, err
	}
	for r := range manifold {
		for c := range manifold[r] {
			switch manifold[r][c] {
			case "S":
				startCol = c // For Part 2.
				continue
			case "^":
				if manifold[r-1][c] == "|" { // Split
					manifold[r][c-1] = "|"
					manifold[r][c+1] = "|"
					p1++
				}
			case ".":
				if r > 0 && (manifold[r-1][c] == "S" || manifold[r-1][c] == "|") {
					manifold[r][c] = "|"
				}
			default: // "|"
				continue
			}
		}
	}
	p2 = countTimelines(manifold, 1, startCol, make(map[VisitedState]int))
	return p1, p2, nil
}

// countTimelines uses DFS to count timelines from a given
// position and direction with memoization.
func countTimelines(manifold [][]string, row, col int, memo map[VisitedState]int) int {
	cols := len(manifold[0])
	rows := len(manifold)
	if row < 0 || row >= rows || col < 0 || col >= cols {
		return 1 // Exited the manifold, count as a timeline
	}

	state := VisitedState{row, col}
	if val, ok := memo[state]; ok {
		return val
	}

	var result int
	cell := manifold[row][col]
	switch cell {
	case "^":
		result = countTimelines(manifold, row+1, col-1, memo) +
			countTimelines(manifold, row+1, col+1, memo)
	default:
		result = countTimelines(manifold, row+1, col, memo)
	}
	memo[state] = result
	return result
}
