package main

import (
	"adventofcode/2023/day/1/tools"
	"log"
	"os"
)

func main() {
	logInfo := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	if checkSum, err := tools.CheckCalibrationValues("input.txt"); err != nil {
		log.Fatal(err)
	} else {
		logInfo.Printf("Calibration values checksum pass: %d", checkSum)
	}
}
