package main

import (
	"adventofcode/2019/day/2/computer"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var programPath = flag.String("program_path", "input/program.txt", "Intcode program file.")

func main() {
	file, err := os.Open(*programPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(computer.ScanInstruction)
	var program computer.Program
	for scanner.Scan() {
		instruction, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("unable to convert to integer: %v", err)
			continue
		}
		program = append(program, instruction)
	}
	fmt.Println(computer.Compute(&program))
}
