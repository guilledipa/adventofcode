package main

import (
	"adventofcode/2023/day/2/tools"
	"log"
	"os"
)

func main() {
	logInfo := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	games, err := tools.NewGames("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	logInfo.Printf("Valid Games ID sum: %d", games.IDSumValidGames())
}
