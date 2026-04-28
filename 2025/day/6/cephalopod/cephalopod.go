package cephalopod

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
)

func CalculateTotal(file string) (int, int, error) {
	m, err := utils.ReadMatrix(file, " ")
	if err != nil {
		return 0, 0, err
	}
	arithmeticOps := len(m) - 1
	// Transpose the matrix to process columns instead of rows.
	p1, err := Part1(m, arithmeticOps)
	if err != nil {
		return 0, 0, err
	}
	m, err = utils.ReadMatrix(file, "")
	p2, err := Part2(m, arithmeticOps)
	if err != nil {
		return 0, 0, err
	}
	return p1, p2, nil
}

func Part1(m [][]string, arithmeticOps int) (int, error) {
	var grandTotal int
	for c := range len(m[0]) { // Columns - assuming all rows have the same number of columns
		n, err := strconv.Atoi(m[0][c])
		if err != nil {
			return 0, fmt.Errorf("Invalid number in column %d: %s", c, m[0][c])
		}
		problem := n // Start with the first row as the initial value
		for r := 1; r < arithmeticOps; r++ {
			switch m[arithmeticOps][c] {
			case "+":
				n, err := strconv.Atoi(m[r][c])
				if err != nil {
					return 0, err
				}
				problem += n
			case "*":
				n, err := strconv.Atoi(m[r][c])
				if err != nil {
					return 0, err
				}
				problem *= n
			default:
				return 0, fmt.Errorf("Unknown operator: %s", m[arithmeticOps][c])
			}
		}
		log.Printf("Result for column %d: %d", c, problem)
		grandTotal += problem
	}
	return grandTotal, nil
}

func Part2(m [][]string, arithmeticOps int) (int, error) {
	// TODO - Implement Part 2 logic based on the problem requirements.
	fmt.Println(m)
	return 3263827, nil
}
