package forklift

import (
	"adventofcode/utils"
	"log"
)

func Forklift(file string) (int, int, error) {
	var rolls1 int
	var rolls int // Part 2 rolls
	var rollsPrev int
	var passes int

	m, err := utils.ReadMatrix(file, "")
	if err != nil {
		return 0, 0, err
	}
	// Part 1 counts rolls first pass only.
	for r := 0; r < len(m); r++ { // Rows
		for c := 0; c < len(m[r]); c++ { // Columns
			if m[r][c] != "@" {
				continue
			}
			if isForkliftable(m, r, c) {
				rolls1++
			}
		}
	}

	// Part 2 resets the matrix and counts rolls until no more can be removed.
	for rolls > rollsPrev && passes < 1000 || passes == 0 {
		passes++
		rollsPrev = rolls
		for r := 0; r < len(m); r++ { // Rows
			for c := 0; c < len(m[r]); c++ { // Columns
				if m[r][c] != "@" {
					continue
				}
				if isForkliftable(m, r, c) {
					rolls++
					m[r][c] = "." // Mark as removed
				}
			}
		}
	}
	return rolls1, rolls, nil
}

// isForkliftable checks if the forklift can move to the given
// position (r, c) in the matrix m.
func isForkliftable(m [][]string, r, c int) bool {
	var rolls int
	rows := len(m)
	cols := len(m[0])

	// If r,c is a corner, check the adjacent positions
	if (r == 0 || r == rows-1) && (c == 0 || c == cols-1) {
		// Top-left corner
		if r == 0 && c == 0 {
			if m[r+1][c] == "@" {
				rolls++
			}
			if m[r][c+1] == "@" {
				rolls++
			}
			if m[r+1][c+1] == "@" {
				rolls++
			}
		}
		// Top-right corner
		if r == 0 && c == cols-1 {
			if m[r+1][c] == "@" {
				rolls++
			}
			if m[r][c-1] == "@" {
				rolls++
			}
			if m[r+1][c-1] == "@" {
				rolls++
			}
		}
		// Bottom-left corner
		if r == rows-1 && c == 0 {
			if m[r-1][c] == "@" {
				rolls++
			}
			if m[r][c+1] == "@" {
				rolls++
			}
			if m[r-1][c+1] == "@" {
				rolls++
			}
		}
		// Bottom-right corner
		if r == rows-1 && c == cols-1 {
			if m[r-1][c] == "@" {
				rolls++
			}
			if m[r][c-1] == "@" {
				rolls++
			}
			if m[r-1][c-1] == "@" {
				rolls++
			}
		}
		log.Printf("Position (%d, %d) CORNER has %d rolls around it\n", r, c, rolls)
		if rolls >= 4 {
			return false
		}
		return true
	}

	// If r,c is an edge, check the adjacent positions (an edge excluding corners)
	if r == 0 || r == rows-1 || c == 0 || c == cols-1 {
		// Top edge
		if r == 0 {
			for j := c - 1; j <= c+1; j++ {
				if j >= 0 && j < cols && m[r+1][j] == "@" {
					rolls++
				}
			}
			if c > 0 && m[r][c-1] == "@" {
				rolls++
			}
			if c < cols-1 && m[r][c+1] == "@" {
				rolls++
			}
		}
		// Bottom edge
		if r == rows-1 {
			for j := c - 1; j <= c+1; j++ {
				if j >= 0 && j < cols && m[r-1][j] == "@" {
					rolls++
				}
			}
			if c > 0 && m[r][c-1] == "@" {
				rolls++
			}
			if c < cols-1 && m[r][c+1] == "@" {
				rolls++
			}
		}
		// Left edge
		if c == 0 {
			for i := r - 1; i <= r+1; i++ {
				if i >= 0 && i < rows && m[i][c+1] == "@" {
					rolls++
				}
			}
			if r > 0 && m[r-1][c] == "@" {
				rolls++
			}
			if r < rows-1 && m[r+1][c] == "@" {
				rolls++
			}
		}
		// Right edge
		if c == cols-1 {
			for i := r - 1; i <= r+1; i++ {
				if i >= 0 && i < rows && m[i][c-1] == "@" {
					rolls++
				}
			}
			if r > 0 && m[r-1][c] == "@" {
				rolls++
			}
			if r < rows-1 && m[r+1][c] == "@" {
				rolls++
			}
		}
		log.Printf("Position (%d, %d) EDGE has %d rolls around it\n", r, c, rolls)
		if rolls >= 4 {
			return false
		}
		return true
	}
	// Check inner matrix
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if i == r && j == c {
				continue // Skip the current position
			}
			if m[i][j] == "@" {
				rolls++
			}
			if rolls >= 4 {
				return false
			}
		}
	}
	log.Printf("Position (%d, %d) INNER has %d rolls around it\n", r, c, rolls)
	// 0 - 3 rolls are valid, 4 or more are not.
	if rolls >= 4 {
		return false
	}
	return true
}
