package main

import (
	"adventofcode/2023/day/2/tools"
	"log"
	"os"
)

func main() {
	logInfo := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	if games, err := tools.NewGames("input.txt"); err != nil {
		log.Fatal(err)
	} else {
		logInfo.Printf("New Games: %v", games)
	}
}
