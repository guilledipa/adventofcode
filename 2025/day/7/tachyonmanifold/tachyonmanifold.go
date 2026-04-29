package tachyonmanifold

import (
	"adventofcode/utils"
)

func CalculateTotal(file string) (int, int, error) {
	var p1Splits, p2 int
	manifold, err := utils.ReadMatrix(file, "")
	if err != nil {
		return 0, 0, err
	}
	for r := range manifold {
		for c := range manifold[r] {
			switch manifold[r][c] {
			case "S":
				continue
			case "^":
				if manifold[r-1][c] == "|" { // Split
					manifold[r][c-1] = "|"
					manifold[r][c+1] = "|"
					p1Splits++
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
	return p1Splits, p2, nil
}
