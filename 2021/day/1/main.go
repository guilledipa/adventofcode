package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type DepthReadings struct {
	Readings []int
}

func (d *DepthReadings) PopulateReadings(file string) error {
	if len(d.Readings) != 0 {
		return errors.New("readings have already been captured")
	}
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		d.Readings = append(d.Readings, depth)
	}
	return scanner.Err()
}

func (d *DepthReadings) Increases() (error, int) {
	var increases []int
	if len(d.Readings) == 0 {
		return errors.New("no readings found"), 0
	}
	for i, r := range d.Readings {
		if i == 0 {
			// First reading; ignore
			continue
		}
		if r > d.Readings[i-1] {
			increases = append(increases, r)
		}
	}
	return nil, len(increases)
}

func (d *DepthReadings) SliceSum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func (d *DepthReadings) SumSlidingWindow() (error, int) {
	window := 3
	windowReadings := new(DepthReadings)
	if len(d.Readings) == 0 {
		return errors.New("no readings found"), 0
	}
	for i := 0; i <= len(d.Readings); i++ {
		if len(d.Readings[:i]) < window {
			// Not enough readings
			continue
		}
		windowSum := windowReadings.SliceSum(d.Readings[i-window : i])
		windowReadings.Readings = append(windowReadings.Readings, windowSum)
	}
	return windowReadings.Increases()
}

func main() {
	d := new(DepthReadings)
	if err := d.PopulateReadings("input/readings.txt"); err != nil {
		log.Fatal(err)
	}
	if err, inc := d.Increases(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Number of increases: %d\n", inc)
	}
	if err, swInc := d.SumSlidingWindow(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Number of sliding window increases: %d\n", swInc)
	}
}
