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

var (
	programPath = flag.String("program_path", "../2/input/program.txt", "Intcode program file.")
)

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

	programBis := make(computer.Program, len(program))
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			_ = copy(programBis, program)
			programBis[1] = noun
			programBis[2] = verb
			result := computer.Compute(&programBis)
			if result[0] == 19690720 {
				fmt.Printf("Answer: %02d%02d -- %v", noun, verb, result)
			}
		}
	}
}
