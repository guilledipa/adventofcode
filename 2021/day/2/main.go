package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

type CoursePlan struct {
	Instructions []string
}

// Position is an (x,y) pair that denotes the sub position 0,0 is the surface.
//
// 0,0 ------► X
//  |
//  |
//  ▼ Y
//
type Position struct {
	X int
	Y int
}

type Submarine struct {
	Course []Position
}

func (c *CoursePlan) PopulateInstructions(file string) error {
	if len(c.Instructions) != 0 {
		return errors.New("readings have already been captured")
	}
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c.Instructions = append(c.Instructions, scanner.Text())
	}
	return scanner.Err()
}

func main() {
	c := new(CoursePlan)
	if err := d.PopulateInstructions("input/course.txt"); err != nil {
		log.Fatal(err)
	}

}
