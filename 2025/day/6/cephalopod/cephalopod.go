package cephalopod

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
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
	// re-read the matrix, now with a different delimeter cos now
	// every position matters.
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
	var grandTotal int
	var columnTotal int
	var partial []int
	// Process right to left.
	for c := len(m[0]) - 1; c >= 0; c-- {
		var sb strings.Builder
		var opSb strings.Builder
		for r := range arithmeticOps {
			sb.WriteString(m[r][c])
		}
		if c < len(m[arithmeticOps]) {
			opSb.WriteString(m[arithmeticOps][c])
		}
		sn := strings.TrimSpace(sb.String())
		if sn == "" {
			continue
		}
		n, err := strconv.Atoi(sn)
		if err != nil {
			return 0, err
		}
		partial = append(partial, n)
		op := strings.TrimSpace(opSb.String())
		if op != "" {
			for _, p := range partial {
				switch op {
				case "+":
					columnTotal += p
				case "*":
					if columnTotal == 0 {
						columnTotal = p
					} else {
						columnTotal *= p
					}
				default:
					return 0, fmt.Errorf("Unknown operator: %s", op)
				}
			}
			grandTotal += columnTotal
			columnTotal = 0 // Reset column total after processing
			partial = nil   // Reset partial after processing
		}
	}
	return grandTotal, nil
}
