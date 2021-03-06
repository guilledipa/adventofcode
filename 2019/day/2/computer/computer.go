package computer

import (
	"errors"
	"unicode/utf8"
)

// Program is...
type Program []int

// IsComma reports whether the character is a Unicode comma character.
func IsComma(r rune) bool {
	if r == ',' {
		return true
	}
	return false
}

// ScanInstruction is a split function for a Scanner that returns each
// comma-separated integer, with surrounding spaces deleted.
func ScanInstruction(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading commas.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !IsComma(r) {
			break
		}
	}
	// Scan until ',', marking end of instruction.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if IsComma(r) {
			return i, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

// Sum does stuff
func Sum(a, b int) int {
	return a + b
}

// Mult does stuff
func Mult(a, b int) int {
	return a * b
}

// ExecuteInstruction does stuff
func ExecuteInstruction(cursor int, program *Program) error {
	switch (*program)[cursor] {
	case 1:
		(*program)[(*program)[cursor+3]] = Sum((*program)[(*program)[cursor+1]], (*program)[(*program)[cursor+2]])
	case 2:
		(*program)[(*program)[cursor+3]] = Mult((*program)[(*program)[cursor+1]], (*program)[(*program)[cursor+2]])
	case 99:
		return errors.New("halt")
	default:
		return errors.New("unsupported instruction")
	}
	return nil
}

// Compute does stuff.
func Compute(program *Program) Program {
	for i := 0; i <= len(*program); i += 4 {
		if err := ExecuteInstruction(i, program); err != nil {
			return *program
		}
	}
	return *program
}
