package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	Direction string
	Intensity int
}

type CoursePlan struct {
	RawVectors []string
	Vectors    []Vector
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
	CoursePlan
	Position
}

func (p *Position) String() string {
	return fmt.Sprintf("X: %d, Y: %d", p.X, p.Y)
}

func (c *CoursePlan) populateRawVectors(file string) error {
	if len(c.RawVectors) != 0 {
		return errors.New("readings have already been captured")
	}
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c.RawVectors = append(c.RawVectors, scanner.Text())
	}
	return scanner.Err()
}

func (c *CoursePlan) parseVectors() error {
	for _, v := range c.RawVectors {
		instruction := strings.Fields(v)
		if len(instruction) != 2 {
			return fmt.Errorf("unrecognized raw vector: %s", v)
		}
		i, err := strconv.Atoi(instruction[1])
		if err != nil {
			return fmt.Errorf("unsupported intensity: %s", instruction[1])
		}
		switch instruction[0] {
		case "up", "down", "forward":
			c.Vectors = append(c.Vectors, Vector{instruction[0], i})
		default:
			return fmt.Errorf("unsupported direction: %s", instruction[0])
		}
	}
	return nil
}

func (s *Submarine) updatePosition(vector Vector) error {
	switch vector.Direction {
	case "up":
		s.Position.Y -= vector.Intensity
	case "down":
		s.Position.Y += vector.Intensity
	case "forward":
		s.Position.X += vector.Intensity
	default:
		return fmt.Errorf("unsupported vector type: %v", vector)
	}
	return nil
}

func (s *Submarine) Navigate() (Position, error) {
	for _, v := range s.CoursePlan.Vectors {
		if err := s.updatePosition(v); err != nil {
			return Position{}, fmt.Errorf("navigation error: %v", err)
		}
	}
	return s.Position, nil
}

func SetupCourse(file string) (*Submarine, error) {
	sub := new(Submarine)
	if err := sub.CoursePlan.populateRawVectors(file); err != nil {
		return nil, fmt.Errorf("unable to populate raw vectors: %v", err)
	}
	if err := sub.CoursePlan.parseVectors(); err != nil {
		return nil, fmt.Errorf("unable to parse vectors: %v", err)
	}
	return sub, nil
}

func main() {
	sub, err := SetupCourse("input/course.txt")
	if err != nil {
		log.Fatal(err)
	}
	pos, err := sub.Navigate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pos)
}
