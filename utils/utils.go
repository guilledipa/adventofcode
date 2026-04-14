package utils

import (
	"bufio"
	"log"
	"os"
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
