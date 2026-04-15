package joltage

import (
	"adventofcode/utils"
	"log"
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
