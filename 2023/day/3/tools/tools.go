package tools

import (
	"adventofcode/2023/day/1/tools"
	"fmt"
)

func GetPartNumbers(filePath string) ([]int, error) {
	lines, err := tools.ReadInputFile(filePath)
	if err != nil {
		return nil, err
	}
	for _, l := range lines {
		fmt.Println(l)
	}
	return nil, nil
}
