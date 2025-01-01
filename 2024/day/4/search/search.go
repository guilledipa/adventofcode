// Package search provides utilities for searching for XMAS in a given input.
//
// This is a visual representation of how I will break the problem apart:
//
// 1. I'll start by finding the XMAS in the boundaries of the matrix.
// 2. I'll then explore the inner matrix A' where I'll look for the XMAS pattern
// in each direction for each Aij.
// 3. I'll then combine the results from the boundaries and the inner matrix.
// 4. I'll then return the number of XMAS found.
//
// This is how the matrix will look like:
//
// A11 A12 .. A1i
// A21 A22 .. A2i
// ...
// Aj1 Aj2 .. Aji
//
// The boundaries are the following:
//
// Top: [A11 A12 .. A1i]
// Bottom: [Aj1 Aj2 .. Aji]
// Left: [A11 A21 .. Aj1]
// Right: [A1i A2i .. Aji]
//
// When I refer to "inner" matrix A', I'm referring to the following:
//
// A22 A23 .. A2(i-1)
// A32 A33 .. A3(i-1)
// ...
// Aj2 Aj3 .. A(j-1)(i-1)
package search

import (
	"adventofcode/2024/day/utils"
	"fmt"
	"strings"
)

// Data is a struct to hold the data to be searched.
type Data struct {
	Matrix [][]string
	Rows   int
	Cols   int
}

func (d *Data) CountXMAS() int {
	count := d.computeRowsAndColumns()
	count += d.computeDiagonals()
	return count
}

// NewData will parse raw data and return a Data struct with a populated matrix.
func NewData(file string) (*Data, error) {
	var d Data
	scanner, f, err := utils.CreateScanner(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	for scanner.Scan() {
		d.Matrix = append(d.Matrix, strings.Split(scanner.Text(), ""))
	}
	d.Rows = len(d.Matrix)
	d.Cols = len(d.Matrix[0])
	return &d, scanner.Err()
}

// countXMASInSlice counts the number of XMAS or SAMX in a given slice.
// This is a helper function that is used for matrix boundaries.
func countXMASInSlice(input []string) int {
	count := 0
	if len(input) == 0 {
		return count
	}
	s := strings.Join(input, "")
	count += strings.Count(s, "XMAS")
	count += strings.Count(s, "SAMX")
	return count
}

// computeRowsAndColumns computes the number of XMAS in rows and columns.
func (d *Data) computeRowsAndColumns() int {
	count := 0
	for i := 0; i < d.Rows; i++ {
		column := []string{}
		count += countXMASInSlice(d.Matrix[i])
		for j := 0; j < d.Cols; j++ {
			column = append(column, d.Matrix[i][j])
		}
		count += countXMASInSlice(column)
	}
	return count
}

// computeDiagonals computes the number of XMAS in diagonals..
// Bugeado, needs more work
func (d *Data) computeDiagonals() int {
	count := 0
	for i := 0; i < d.Rows; i++ {
		fmt.Println("i: ", i)
		for j := 0; j < d.Cols; j++ {
			fmt.Println("j: ", j)
			// Check diagonal from top-left to bottom-right
			diag1 := []string{}
			for k := 0; ; k++ {
				r := i + k
				c := j + k
				if r >= d.Rows || c >= d.Cols {
					break
				}
				diag1 = append(diag1, d.Matrix[r][c])
			}
			if len(diag1) < 4 {
				break
			}
			count += countXMASInSlice(diag1)
			fmt.Println(diag1, ": ", count)
		}
	}
	return count
}
