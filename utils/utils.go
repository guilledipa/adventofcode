package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// CreateScanner is a helper function to scan files line by line.
// Note: I need to return the os.File filedescriptor because I can't close
// the file here; I must do it after I use the scanner.
func CreateScanner(file string) (*bufio.Scanner, *os.File, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}
	log.Println("Open file:", f.Name())
	scanner := bufio.NewScanner(f)
	return scanner, f, scanner.Err()
}

// ProcessLines is a helper function to process each line of a file with
// a provided function. The data is processed line by line, so it can handle
// large files without loading them entirely into memory.
func ProcessLines(file string, fn func(string) error) error {
	scanner, f, err := CreateScanner(file)
	if err != nil {
		return err
	}
	defer f.Close()

	for scanner.Scan() {
		if err := fn(scanner.Text()); err != nil {
			return err
		}
	}

	return scanner.Err()
}

// ReadLines is a helper function to read all lines from a file into a slice of strings.
// Data is loaded entirely into memory, so it should be used for smaller files.
func ReadLines(file string) ([]string, error) {
	var lines []string
	if err := ProcessLines(file, func(line string) error {
		lines = append(lines, line)
		return nil
	}); err != nil {
		return nil, err
	}
	return lines, nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func StrToInt(strings []string) []int {
	ints := make([]int, 0, len(strings))
	for _, s := range strings {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		ints = append(ints, i)
	}
	return ints
}
