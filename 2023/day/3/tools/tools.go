package tools

import (
	"adventofcode/2023/day/1/tools"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	paddingChar = "."
)

type Engine struct {
	Schematic []string
}

// findPartCoordinatePairs a slice of coordinate pair for each number found on a
// given line (string)
func findPartCoordinatePairs(str string) [][]int {
	// Define the regular expression for finding numbers
	re := regexp.MustCompile(`\d+`)
	// Find all matches in the input string
	matches := re.FindAllStringSubmatchIndex(str, -1)
	return matches
}

// Aca va el test con padding
//
// ............
// .467..114...
// ....*.......
// ...35..633..
// .......#....
// .617*.......
// ......+.58..
// ...592......
// .......755..
// ....$.*.....
// ..664.598...
// ............
func (e *Engine) LoadSchematic(filePath string) error {
	lines, err := tools.ReadInputFile(filePath)
	if err != nil {
		return err
	}
	// Agrego padding (.) para facilitar el escaneo de numeros, ya que puedo
	// escanear en todas las direcciones sin preocuparme por las condiciones
	// de borde.
	for _, l := range lines {
		e.Schematic = append(e.Schematic, fmt.Sprintf(".%s.", l))
	}
	padding := strings.Repeat(paddingChar, len(e.Schematic[0]))
	e.Schematic = append([]string{padding}, e.Schematic...)
	e.Schematic = append(e.Schematic, padding)
	return nil
}

// scanLine will return true if it finds a symbol
// Asumo que las partes pueden tener 1, 2 o 3 digitos.
// 1 Digito: 3 arriba, izquierda derecha, 3 abajo
// 2 Digitos: 4 arriba, izquierda derecha, 4 abajo
// 3 Digitos: 5 arriba, izquierda derecha, 5 abajo
func (e *Engine) scanLine(line int, partCoordinates []int) bool {
	lookAhead := partCoordinates[1] - partCoordinates[0]
	cursor := partCoordinates[0]
	// fmt.Println(partCoordinates)
	for j := line - 1; j <= line+1; j++ {
		for i := cursor - 1; i <= (cursor + lookAhead); i++ {
			// fmt.Println("Search Line: ", j, "- Cursor: ", cursor, "- LookAhead: ", lookAhead, "- Coords: ", j, i, "-", string(e.Schematic[j][i]))
			// Skip checkin the part itself.
			if j == line && (i >= partCoordinates[0] && i < partCoordinates[1]) {
				continue
			}
			// Aparentemente no existe una situacion como la siguiente:
			// .....
			// ...1.
			// ..2..
			// .....
			// Esta situacion devolveria tanto a 2 como a 1 como partes validas,
			// pero asumo que no es una condicion posible segun los datos dados.
			if e.Schematic[j][i] != '.' {
				return true
			}
		}
	}
	return false
}

func (e *Engine) recoverPartNumber(line int, partCoordinates []int) (int, error) {
	return strconv.Atoi(e.Schematic[line][partCoordinates[0]:partCoordinates[1]])
}

func GetPartNumbers(filePath string) (int, error) {
	engine := new(Engine)
	var partSum int
	if err := engine.LoadSchematic(filePath); err != nil {
		return partSum, err
	}
	for lineNumber, line := range engine.Schematic {
		// fmt.Println("Scanning: ", lineNumber, line)
		// Skip padding lines
		if lineNumber == 0 || lineNumber == len(engine.Schematic) {
			continue
		}
		partCoordinatePairs := findPartCoordinatePairs(line)     // [[0 3] [5 8]]
		for _, partCoordinatePair := range partCoordinatePairs { // [0 3]
			if !engine.scanLine(lineNumber, partCoordinatePair) {
				continue
			}
			partN, err := engine.recoverPartNumber(lineNumber, partCoordinatePair)
			if err != nil {
				return partSum, err
			}
			partSum += partN
		}
	}
	return partSum, nil
}
