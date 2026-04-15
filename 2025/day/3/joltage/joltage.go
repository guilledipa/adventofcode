package joltage

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
)

func MaxJoltage(bankFile string) (int, error) {
	var result int
	err := utils.ProcessLines(bankFile, func(line string) error {
		result += processJoltageLine(line)
		return nil
	})
	if err != nil {
		return 0, err
	}
	return result, nil
}

func MaxJoltage2(bankFile string) (int, error) {
	var result int
	err := utils.ProcessLines(bankFile, func(line string) error {
		joltage, err := processJoltageLine2(line)
		if err != nil {
			return err
		}
		result += joltage
		return nil
	})
	if err != nil {
		return 0, err
	}
	return result, nil
}

func processJoltageLine(line string) int {
	var unidad rune
	var decena rune
	var cursor int
	for i := 0; i < len(line)-1; i++ {
		r := rune(line[i])
		if r > decena {
			decena = r
			cursor = i
			log.Printf("Updated decena to %c at index %d\n", decena, i)
		}
	}
	for i := cursor + 1; i < len(line); i++ {
		r := rune(line[i])
		if r > unidad {
			unidad = r
			log.Printf("Updated unidad to %c at index %d\n", unidad, i)
		}
	}
	joltage := int(decena-'0')*10 + int(unidad-'0')
	log.Printf("Processed line: %q, decena: %c, unidad: %c, joltage: %d\n", line, decena, unidad, joltage)
	return joltage
}

func processJoltageLine2(line string) (int, error) {
	n := len(line)
	if n < len([12]rune{}) {
		return 0, fmt.Errorf("line too short for 12-digit joltage: %q", line)
	}
	var joltageArr [12]rune
	start := 0
	for i := 0; i < len(joltageArr); i++ {
		remaining := len(joltageArr) - i
		maxPos := n - remaining
		best := rune('0' - 1)
		bestPos := start
		for j := start; j <= maxPos; j++ {
			r := rune(line[j])
			if r > best {
				best = r
				bestPos = j
			}
		}
		joltageArr[i] = best
		start = bestPos + 1
		log.Printf("Selected digit %c for position %d at index %d\n", best, i, bestPos)
	}
	return strconv.Atoi(string(joltageArr[:]))
}
