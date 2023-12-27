package main

import (
	"adventofcode/2023/day/3/tools"
	"log"
	"os"
)

func main() {
	logInfo := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	parts, err := tools.GetPartNumbers("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	logInfo.Printf("Engine parts: %v", parts)
}
